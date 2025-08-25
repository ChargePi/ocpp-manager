package controllers

import (
	"fmt"
	"sync"

	"github.com/ChargePi/ocpp-manager/ocpp_v201/component"
	"github.com/ChargePi/ocpp-manager/ocpp_v201/variables"
)

const (
	VariableNameEVConnectionTimeOut      variables.VariableName = "EVConnectionTimeOut"
	VariableNameStopTxOnEVSideDisconnect variables.VariableName = "StopTxOnEVSideDisconnect"
	VariableNameTxBeforeAcceptedEnabled  variables.VariableName = "TxBeforeAcceptedEnabled"
	VariableNameTxStartPoint             variables.VariableName = "TxStartPoint"
	VariableNameTxStopPoint              variables.VariableName = "TxStopPoint"
	VariableNameMaxEnergyOnInvalidId     variables.VariableName = "MaxEnergyOnInvalidId"
	VariableNameStopTxOnInvalidId        variables.VariableName = "StopTxOnInvalidId"
)

func requiredTxCtrlrVariables() []variables.VariableName {
	return []variables.VariableName{
		VariableNameTxStartPoint,
		VariableNameTxStopPoint,
	}
}

func optionalTxCtrlrVariables() []variables.VariableName {
	return []variables.VariableName{
		VariableNameTxBeforeAcceptedEnabled,
		VariableNameMaxEnergyOnInvalidId,
	}
}

func supportedTxCtrlrVariables() []variables.VariableName {
	return append(requiredTxCtrlrVariables(), optionalTxCtrlrVariables()...)
}

type TxCtrlr struct {
	mu                 sync.RWMutex
	variables          map[variables.VariableName]*variables.Variable
	requiredVariables  []variables.VariableName
	supportedVariables []variables.VariableName
	validator          *variableValidator
	instanceId         string
}

func (t *TxCtrlr) GetName() component.Name {
	return component.ComponentNameTxCtrlr
}

func (t *TxCtrlr) GetInstanceId() string {
	return t.instanceId
}

func (t *TxCtrlr) RegisterSubComponent(component component.Component) {
	// No-op: controllers do not support sub-components
}

func (t *TxCtrlr) UnregisterSubComponent(component component.Component) {
	// No-op: controllers do not support sub-components
}

func (t *TxCtrlr) GetSubComponents() []component.Component {
	return []component.Component{}
}

func (t *TxCtrlr) GetRequiredVariables() []variables.VariableName {
	t.mu.RLock()
	defer t.mu.RUnlock()

	return t.requiredVariables
}

func (t *TxCtrlr) GetSupportedVariables() []variables.VariableName {
	t.mu.RLock()
	defer t.mu.RUnlock()

	return t.supportedVariables
}

func (t *TxCtrlr) GetVariable(key variables.VariableName, opts ...component.GetSetVariableOption) (*variables.Variable, error) {
	if !t.validator.IsVariableSupported(key) {
		return nil, fmt.Errorf("variable %s is not supported by this controller", key)
	}

	t.mu.RLock()
	defer t.mu.RUnlock()

	variable, exists := t.variables[key]
	if !exists {
		return nil, fmt.Errorf("variable %s not found", key)
	}
	return variable, nil
}

func (t *TxCtrlr) UpdateVariable(variable variables.VariableName, attribute string, value interface{}, opts ...component.GetSetVariableOption) error {
	if !t.validator.IsVariableSupported(variable) {
		return fmt.Errorf("variable %s is not supported by this controller", variable)
	}

	t.mu.Lock()
	defer t.mu.Unlock()

	v, exists := t.variables[variable]
	if !exists {
		return fmt.Errorf("variable %s not found", variable)
	}

	return v.UpdateVariableAttribute(attribute, value)
}

func (t *TxCtrlr) Validate(key variables.VariableName) bool {

	if !t.validator.IsVariableSupported(key) {
		return false
	}
	t.mu.RLock()
	defer t.mu.RUnlock()

	v, exists := t.variables[key]
	if !exists {
		return false
	}

	return v.Validate()
}

func NewTxCtrlr() *TxCtrlr {
	ctrlr := &TxCtrlr{
		variables:          make(map[variables.VariableName]*variables.Variable),
		requiredVariables:  requiredTxCtrlrVariables(),
		supportedVariables: supportedTxCtrlrVariables(),
		instanceId:         "tx-ctrlr",
	}
	ctrlr.validator = newVariableValidator(ctrlr)
	return ctrlr
}
