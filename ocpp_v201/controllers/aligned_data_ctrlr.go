package controllers

import (
	"fmt"
	"sync"

	"github.com/ChargePi/ocpp-manager/ocpp_v201/component"
	"github.com/ChargePi/ocpp-manager/ocpp_v201/variables"
)

const (
	VariableNameAlignedDataEnabled           variables.VariableName = "Enabled"
	VariableNameAlignedDataAvailable         variables.VariableName = "Available"
	VariableNameAlignedDataMeasurands        variables.VariableName = "Measurands"
	VariableNameAlignedDataInterval          variables.VariableName = "Interval"
	VariableNameAlignedDataSendDuringIdle    variables.VariableName = "SendDuringIdle"
	VariableNameAlignedDataSignReadings      variables.VariableName = "SignReadings"
	VariableNameAlignedDataTxEndedMeasurands variables.VariableName = "TxEndedMeasurands"
	VariableNameAlignedDataTxEndedInterval   variables.VariableName = "TxEndedInterval"
)

func requiredAlignedDataVariables() []variables.VariableName {
	return []variables.VariableName{
		VariableNameAlignedDataMeasurands,
		VariableNameAlignedDataInterval,
		VariableNameAlignedDataTxEndedMeasurands,
		VariableNameAlignedDataTxEndedInterval,
	}
}

func optionalAlignedDataVariables() []variables.VariableName {
	return []variables.VariableName{
		VariableNameAlignedDataEnabled,
		VariableNameAlignedDataAvailable,
		VariableNameAlignedDataSendDuringIdle,
		VariableNameAlignedDataSignReadings,
	}
}

func supportedAlignedDataVariables() []variables.VariableName {
	return append(requiredAlignedDataVariables(), optionalAlignedDataVariables()...)
}

type AlignedDataCtrlr struct {
	mu                 sync.RWMutex
	variables          map[variables.VariableName]*variables.Variable
	requiredVariables  []variables.VariableName
	supportedVariables []variables.VariableName
	instanceId         string
	validator          *variableValidator
}

func (a *AlignedDataCtrlr) GetName() component.ComponentName {
	return component.ComponentNameAlignedDataCtrlr
}

func (a *AlignedDataCtrlr) GetInstanceId() string {
	return a.instanceId
}

func (a *AlignedDataCtrlr) RegisterSubComponent(component component.Component) {
	// No-op: controllers do not support sub-components
}

func (a *AlignedDataCtrlr) UnregisterSubComponent(component component.Component) {
	// No-op: controllers do not support sub-components
}

func (a *AlignedDataCtrlr) GetSubComponents() []component.Component {
	// Controllers do not support sub-components, always return empty slice
	return []component.Component{}
}

func (a *AlignedDataCtrlr) GetRequiredVariables() []variables.VariableName {
	a.mu.RLock()
	defer a.mu.RUnlock()

	return a.requiredVariables
}

func (a *AlignedDataCtrlr) GetSupportedVariables() []variables.VariableName {
	a.mu.RLock()
	defer a.mu.RUnlock()

	return a.supportedVariables
}

func (a *AlignedDataCtrlr) GetVariable(key variables.VariableName, opts ...component.GetSetVariableOption) (*variables.Variable, error) {
	if !a.validator.IsVariableSupported(key) {
		return nil, fmt.Errorf("variable %s is not supported", key)
	}

	a.mu.RLock()
	defer a.mu.RUnlock()

	variable, exists := a.variables[key]
	if !exists {
		return nil, fmt.Errorf("variable %s not found", key)
	}

	return variable, nil
}

func (a *AlignedDataCtrlr) UpdateVariable(variable variables.VariableName, attribute string, value interface{}, opts ...component.GetSetVariableOption) error {
	if !a.validator.IsVariableSupported(variable) {
		return fmt.Errorf("variable %s is not supported", variable)
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	v, exists := a.variables[variable]
	if !exists {
		return fmt.Errorf("variable %s not found", variable)
	}

	return v.UpdateVariableAttribute(attribute, value)
}

func (a *AlignedDataCtrlr) Validate(key variables.VariableName) bool {
	if !a.validator.IsVariableSupported(key) {
		return false
	}

	a.mu.RLock()
	defer a.mu.RUnlock()

	v, exists := a.variables[key]
	if !exists {
		return false
	}

	return v.Validate()
}

func NewAlignedDataCtrlr() *AlignedDataCtrlr {
	ctrlr := &AlignedDataCtrlr{
		mu:                 sync.RWMutex{},
		variables:          make(map[variables.VariableName]*variables.Variable),
		requiredVariables:  requiredAlignedDataVariables(),
		supportedVariables: supportedAlignedDataVariables(),
		instanceId:         "aligned-data-ctrlr",
	}

	ctrlr.validator = newVariableValidator(ctrlr)
	return ctrlr
}
