package controllers

import (
	"github.com/ChargePi/ocpp-manager/ocpp_v201/component"
	"github.com/ChargePi/ocpp-manager/ocpp_v201/variables"
)

const (
	VariableNameSmartChargingEnabled             variables.VariableName = "Enabled"
	VariableNameSmartChargingAvailable           variables.VariableName = "Available"
	VariableNameACPhaseSwitchingSupported        variables.VariableName = "ACPhaseSwitchingSupported"
	VariableNameChargingProfileStackLevel        variables.VariableName = "ProfileStackLevel"
	VariableNameChargingScheduleChargingRateUnit variables.VariableName = "RateUnit"
	VariableNamePeriodsPerSchedule               variables.VariableName = "PeriodsPerSchedule"
	VariableNameExternalControlSignalsEnabled    variables.VariableName = "ExternalControlSignalsEnabled"
	VariableNameNotifyChargingLimitWithSchedules variables.VariableName = "NotifyChargingLimitWithSchedules"
	VariableNamePhases3to1                       variables.VariableName = "Phases3to1"
	VariableChargingProfileEntries               variables.VariableName = "Entries"
	VariableLimitChangeSignificance              variables.VariableName = "LimitChangeSignificance"
)

func requiredSmartChargingVariables() []variables.VariableName {
	return []variables.VariableName{
		VariableNameChargingProfileStackLevel,
		VariableNameChargingScheduleChargingRateUnit,
		VariableNamePeriodsPerSchedule,
		VariableChargingProfileEntries,
		VariableLimitChangeSignificance,
	}
}

func optionalSmartChargingVariables() []variables.VariableName {
	return []variables.VariableName{
		VariableNameSmartChargingEnabled,
		VariableNameSmartChargingAvailable,
		VariableNameACPhaseSwitchingSupported,
		VariableNameExternalControlSignalsEnabled,
		VariableNameNotifyChargingLimitWithSchedules,
		VariableNamePhases3to1,
	}
}

func supportedSmartChargingVariables() []variables.VariableName {
	return append(requiredSmartChargingVariables(), optionalSmartChargingVariables()...)
}

type SmartChargingCtrlr struct {
	variables          map[variables.VariableName]variables.Variable
	requiredVariables  []variables.VariableName
	supportedVariables []variables.VariableName
}

func (s *SmartChargingCtrlr) GetName() component.ComponentName {
	//TODO implement me
	panic("implement me")
}

func (s *SmartChargingCtrlr) GetInstanceId() string {
	//TODO implement me
	panic("implement me")
}

func (s *SmartChargingCtrlr) RegisterSubComponent(component component.Component) {
	//TODO implement me
	panic("implement me")
}

func (s *SmartChargingCtrlr) UnregisterSubComponent(component component.Component) {
	//TODO implement me
	panic("implement me")
}

func (s *SmartChargingCtrlr) GetSubComponents() []component.Component {
	//TODO implement me
	panic("implement me")
}

func (s *SmartChargingCtrlr) GetRequiredVariables() []variables.Variable {
	//TODO implement me
	panic("implement me")
}

func (s *SmartChargingCtrlr) GetSupportedVariables() []variables.Variable {
	//TODO implement me
	panic("implement me")
}

func (s *SmartChargingCtrlr) GetVariable(key variables.VariableName, opts ...component.GetSetVariableOption) (*variables.Variable, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SmartChargingCtrlr) UpdateVariable(variable variables.VariableName, attribute string, value interface{}, opts ...component.GetSetVariableOption) error {
	//TODO implement me
	panic("implement me")
}

func (s *SmartChargingCtrlr) Validate(key variables.VariableName) bool {
	//TODO implement me
	panic("implement me")
}

func NewSmartChargingCtrlr() *SmartChargingCtrlr {
	return &SmartChargingCtrlr{
		variables:          make(map[variables.VariableName]variables.Variable),
		requiredVariables:  requiredSmartChargingVariables(),
		supportedVariables: supportedSmartChargingVariables(),
	}
}
