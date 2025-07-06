package controllers

import (
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
	variables          map[variables.VariableName]variables.Variable
	requiredVariables  []variables.VariableName
	supportedVariables []variables.VariableName
}

func (a *AlignedDataCtrlr) GetName() component.ComponentName {
	//TODO implement me
	panic("implement me")
}

func (a *AlignedDataCtrlr) GetInstanceId() string {
	//TODO implement me
	panic("implement me")
}

func (a *AlignedDataCtrlr) RegisterSubComponent(component component.Component) {
	//TODO implement me
	panic("implement me")
}

func (a *AlignedDataCtrlr) UnregisterSubComponent(component component.Component) {
	//TODO implement me
	panic("implement me")
}

func (a *AlignedDataCtrlr) GetSubComponents() []component.Component {
	//TODO implement me
	panic("implement me")
}

func (a *AlignedDataCtrlr) GetRequiredVariables() []variables.Variable {
	//TODO implement me
	panic("implement me")
}

func (a *AlignedDataCtrlr) GetSupportedVariables() []variables.Variable {
	//TODO implement me
	panic("implement me")
}

func (a *AlignedDataCtrlr) GetVariable(key variables.VariableName, opts ...component.GetSetVariableOption) (*variables.Variable, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AlignedDataCtrlr) UpdateVariable(variable variables.VariableName, attribute string, value interface{}, opts ...component.GetSetVariableOption) error {
	//TODO implement me
	panic("implement me")
}

func (a *AlignedDataCtrlr) Validate(key variables.VariableName) bool {
	//TODO implement me
	panic("implement me")
}

func NewAlignedDataCtrlr() *AlignedDataCtrlr {
	return &AlignedDataCtrlr{
		variables:          make(map[variables.VariableName]variables.Variable),
		requiredVariables:  requiredAlignedDataVariables(),
		supportedVariables: supportedAlignedDataVariables(),
	}
}
