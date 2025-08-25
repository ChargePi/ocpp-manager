package controllers

import (
	"errors"
	"fmt"
	"sync"

	"github.com/ChargePi/ocpp-manager/ocpp_v201/component"
	"github.com/ChargePi/ocpp-manager/ocpp_v201/variables"
)

const (
	VariableNameTariffEnabled            variables.VariableName = "Enabled"
	VariableNameTariffAvailable          variables.VariableName = "Available"
	VariableNameTariffFallbackMessage    variables.VariableName = "TariffFallbackMessage"
	VariableNameCostEnabled              variables.VariableName = "Enabled"
	VariableNameCostAvailable            variables.VariableName = "Available"
	VariableNameTotalCostFallbackMessage variables.VariableName = "TotalCostFallbackMessage"
	VariableNameCurrency                 variables.VariableName = "Currency"
)

func requiredVariablesTariffCostCtrlr() []variables.VariableName {
	return []variables.VariableName{
		VariableNameTariffFallbackMessage,
		VariableNameTotalCostFallbackMessage,
		VariableNameCurrency,
	}
}

func optionalVariablesTariffCostCtrlr() []variables.VariableName {
	return []variables.VariableName{
		VariableNameTariffEnabled,
		VariableNameTariffAvailable,
		VariableNameCostEnabled,
		VariableNameCostAvailable,
	}
}

func supportedVariablesTariffCostCtrlr() []variables.VariableName {
	return append(requiredVariablesTariffCostCtrlr(), optionalVariablesTariffCostCtrlr()...)
}

type TariffCostCtrlr struct {
	mu                 sync.RWMutex
	variables          map[variables.VariableName]*variables.Variable
	requiredVariables  []variables.VariableName
	supportedVariables []variables.VariableName
	instanceId         string
	validator          *variableValidator
}

func (t *TariffCostCtrlr) GetName() component.Name {
	return component.ComponentNameTariffCostCtrlr
}

func (t *TariffCostCtrlr) GetInstanceId() string {
	return t.instanceId
}

func (t *TariffCostCtrlr) RegisterSubComponent(component component.Component) {
}

func (t *TariffCostCtrlr) UnregisterSubComponent(component component.Component) {
}

func (t *TariffCostCtrlr) GetSubComponents() []component.Component {
	return []component.Component{}
}

func (t *TariffCostCtrlr) GetRequiredVariables() []variables.VariableName {
	t.mu.RLock()
	defer t.mu.RUnlock()

	return t.requiredVariables
}

func (t *TariffCostCtrlr) GetSupportedVariables() []variables.VariableName {
	t.mu.RLock()
	defer t.mu.RUnlock()

	return t.supportedVariables
}

func (t *TariffCostCtrlr) GetVariable(key variables.VariableName, opts ...component.GetSetVariableOption) (*variables.Variable, error) {
	if !t.validator.IsVariableSupported(key) {
		return nil, fmt.Errorf("variable %s is not supported", key)
	}

	t.mu.RLock()
	defer t.mu.RUnlock()

	variable, exists := t.variables[key]
	if !exists {
		return nil, errors.New("variable does not exist")
	}

	return variable, nil
}

func (t *TariffCostCtrlr) UpdateVariable(variable variables.VariableName, attribute string, value interface{}, opts ...component.GetSetVariableOption) error {
	if !t.validator.IsVariableSupported(variable) {
		return fmt.Errorf("variable %s is not supported", variable)
	}

	t.mu.Lock()
	defer t.mu.Unlock()
	if _, exists := t.variables[variable]; !exists {
		return errors.New("variable does not exist")
	}

	variableInstance := t.variables[variable]
	err := variableInstance.UpdateVariableAttribute(attribute, value)
	if err != nil {
		return err
	}

	t.variables[variable] = variableInstance
	return nil
}

func (t *TariffCostCtrlr) Validate(key variables.VariableName) bool {
	if !t.validator.IsVariableSupported(key) {
		return false
	}

	t.mu.RLock()
	defer t.mu.RUnlock()
	variable, exists := t.variables[key]
	if !exists {
		return false
	}

	return variable.Validate()
}

func NewTariffCostCtrlr() *TariffCostCtrlr {
	ctrlr := &TariffCostCtrlr{
		mu:                 sync.RWMutex{},
		variables:          make(map[variables.VariableName]*variables.Variable),
		requiredVariables:  requiredVariablesTariffCostCtrlr(),
		supportedVariables: supportedVariablesTariffCostCtrlr(),
		instanceId:         "tariff-cost-ctrlr",
	}

	ctrlr.validator = newVariableValidator(ctrlr)

	return ctrlr
}
