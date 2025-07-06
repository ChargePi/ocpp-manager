package controllers

import (
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
	variables          map[variables.VariableName]variables.Variable
	requiredVariables  []variables.VariableName
	supportedVariables []variables.VariableName
}

func (t *TariffCostCtrlr) GetName() component.ComponentName {
	//TODO implement me
	panic("implement me")
}

func (t *TariffCostCtrlr) GetInstanceId() string {
	//TODO implement me
	panic("implement me")
}

func (t *TariffCostCtrlr) RegisterSubComponent(component component.Component) {
	//TODO implement me
	panic("implement me")
}

func (t *TariffCostCtrlr) UnregisterSubComponent(component component.Component) {
	//TODO implement me
	panic("implement me")
}

func (t *TariffCostCtrlr) GetSubComponents() []component.Component {
	//TODO implement me
	panic("implement me")
}

func (t *TariffCostCtrlr) GetRequiredVariables() []variables.Variable {
	//TODO implement me
	panic("implement me")
}

func (t *TariffCostCtrlr) GetSupportedVariables() []variables.Variable {
	//TODO implement me
	panic("implement me")
}

func (t *TariffCostCtrlr) GetVariable(key variables.VariableName, opts ...component.GetSetVariableOption) (*variables.Variable, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TariffCostCtrlr) UpdateVariable(variable variables.VariableName, attribute string, value interface{}, opts ...component.GetSetVariableOption) error {
	//TODO implement me
	panic("implement me")
}

func (t *TariffCostCtrlr) Validate(key variables.VariableName) bool {
	//TODO implement me
	panic("implement me")
}

func NewTariffCostCtrlr() *TariffCostCtrlr {
	return &TariffCostCtrlr{
		variables:          make(map[variables.VariableName]variables.Variable),
		requiredVariables:  requiredVariablesTariffCostCtrlr(),
		supportedVariables: supportedVariablesTariffCostCtrlr(),
	}
}
