package controllers

import (
	"github.com/ChargePi/ocpp-manager/ocpp_v201/component"
	"github.com/ChargePi/ocpp-manager/ocpp_v201/variables"
)

const (
	VariableNameDateTime                         variables.VariableName = "DateTime"
	VariableNameNtpSource                        variables.VariableName = "NtpSource"
	VariableNameNtpServerUri                     variables.VariableName = "NtpServerUri"
	VariableNameTimeOffset                       variables.VariableName = "TimeOffset"
	VariableNameNextTimeOffsetTransitionDateTime variables.VariableName = "NextTimeOffsetTransitionDateTime"
	VariableNameTimeOffsetNextTransition         variables.VariableName = "TimeOffsetNextTransition"
	VariableNameTimeSource                       variables.VariableName = "TimeSource"
	VariableNameTimeZone                         variables.VariableName = "TimeZone"
	VariableNameTimeAdjustmentReportingThreshold variables.VariableName = "TimeAdjustmentReportingThreshold"
)

func requiredClockVariables() []variables.VariableName {
	return []variables.VariableName{
		VariableNameDateTime,
		VariableNameTimeSource,
	}
}

func optionalClockVariables() []variables.VariableName {
	return []variables.VariableName{
		VariableNameTimeZone,
		VariableNameNtpSource,
		VariableNameNtpServerUri,
		VariableNameTimeOffsetNextTransition,
		VariableNameNextTimeOffsetTransitionDateTime,
		VariableNameTimeOffset,
		VariableNameTimeAdjustmentReportingThreshold,
	}
}

// supportedClockVariables returns a list of all variables supported by the ClockCtrlr.
func supportedClockVariables() []variables.VariableName {
	return append(requiredClockVariables(), optionalClockVariables()...)
}

type ClockCtrlr struct {
	variables          map[variables.VariableName]variables.Variable
	requiredVariables  []variables.VariableName
	supportedVariables []variables.VariableName
}

func (c *ClockCtrlr) GetName() component.ComponentName {
	//TODO implement me
	panic("implement me")
}

func (c *ClockCtrlr) GetInstanceId() string {
	//TODO implement me
	panic("implement me")
}

func (c *ClockCtrlr) RegisterSubComponent(component component.Component) {
	//TODO implement me
	panic("implement me")
}

func (c *ClockCtrlr) UnregisterSubComponent(component component.Component) {
	//TODO implement me
	panic("implement me")
}

func (c *ClockCtrlr) GetSubComponents() []component.Component {
	//TODO implement me
	panic("implement me")
}

func (c *ClockCtrlr) GetRequiredVariables() []variables.Variable {
	//TODO implement me
	panic("implement me")
}

func (c *ClockCtrlr) GetSupportedVariables() []variables.Variable {
	//TODO implement me
	panic("implement me")
}

func (c *ClockCtrlr) GetVariable(key variables.VariableName, opts ...component.GetSetVariableOption) (*variables.Variable, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ClockCtrlr) UpdateVariable(variable variables.VariableName, attribute string, value interface{}, opts ...component.GetSetVariableOption) error {
	//TODO implement me
	panic("implement me")
}

func (c *ClockCtrlr) Validate(key variables.VariableName) bool {
	//TODO implement me
	panic("implement me")
}

func NewClockCtrlr() *ClockCtrlr {
	return &ClockCtrlr{
		variables:          make(map[variables.VariableName]variables.Variable),
		requiredVariables:  requiredClockVariables(),
		supportedVariables: supportedClockVariables(),
	}
}
