package controllers

import (
	"github.com/ChargePi/ocpp-manager/ocpp_v201/component"
	"github.com/ChargePi/ocpp-manager/ocpp_v201/variables"
)

const (
	VariableNameSampledDataEnabled                     variables.VariableName = "Enabled"
	VariableNameSampledDataAvailable                   variables.VariableName = "Available"
	VariableNameSampledDataSignReadings                variables.VariableName = "SignReadings"
	VariableNameSampledDataTxEndedMeasurands           variables.VariableName = "TxEndedMeasurands"
	VariableNameSampledDataTxEndedInterval             variables.VariableName = "TxEndedInterval"
	VariableNameSampledDataTxStartedMeasurands         variables.VariableName = "TxStartedMeasurands"
	VariableNameSampledDataTxUpdatedMeasurands         variables.VariableName = "TxUpdatedMeasurands"
	VariableNameSampledDataTxUpdatedInterval           variables.VariableName = "TxUpdatedInterval"
	VariableNameSampledDataRegisterValuesWithoutPhases variables.VariableName = "RegisterValuesWithoutPhases"
)

func requiredSampledDataVariables() []variables.VariableName {
	return []variables.VariableName{
		VariableNameSampledDataTxEndedMeasurands,
		VariableNameSampledDataTxEndedInterval,
		VariableNameSampledDataTxStartedMeasurands,
		VariableNameSampledDataTxUpdatedMeasurands,
		VariableNameSampledDataTxUpdatedInterval,
	}
}

func optionalSampledDataVariables() []variables.VariableName {
	return []variables.VariableName{
		VariableNameSampledDataEnabled,
		VariableNameSampledDataAvailable,
		VariableNameSampledDataSignReadings,
		VariableNameSampledDataRegisterValuesWithoutPhases,
	}
}

func supportedSampledDataVariables() []variables.VariableName {
	return append(requiredSampledDataVariables(), optionalSampledDataVariables()...)
}

type SampledDataCtrlr struct {
	variables          map[variables.VariableName]variables.Variable
	requiredVariables  []variables.VariableName
	supportedVariables []variables.VariableName
}

func (s *SampledDataCtrlr) GetName() component.ComponentName {
	//TODO implement me
	panic("implement me")
}

func (s *SampledDataCtrlr) GetInstanceId() string {
	//TODO implement me
	panic("implement me")
}

func (s *SampledDataCtrlr) RegisterSubComponent(component component.Component) {
	//TODO implement me
	panic("implement me")
}

func (s *SampledDataCtrlr) UnregisterSubComponent(component component.Component) {
	//TODO implement me
	panic("implement me")
}

func (s *SampledDataCtrlr) GetSubComponents() []component.Component {
	//TODO implement me
	panic("implement me")
}

func (s *SampledDataCtrlr) GetRequiredVariables() []variables.Variable {
	//TODO implement me
	panic("implement me")
}

func (s *SampledDataCtrlr) GetSupportedVariables() []variables.Variable {
	//TODO implement me
	panic("implement me")
}

func (s *SampledDataCtrlr) GetVariable(key variables.VariableName, opts ...component.GetSetVariableOption) (*variables.Variable, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SampledDataCtrlr) UpdateVariable(variable variables.VariableName, attribute string, value interface{}, opts ...component.GetSetVariableOption) error {
	//TODO implement me
	panic("implement me")
}

func (s *SampledDataCtrlr) Validate(key variables.VariableName) bool {
	//TODO implement me
	panic("implement me")
}

func NewSampledDataCtrlr() *SampledDataCtrlr {
	return &SampledDataCtrlr{
		variables:          make(map[variables.VariableName]variables.Variable),
		requiredVariables:  requiredSampledDataVariables(),
		supportedVariables: supportedSampledDataVariables(),
	}
}
