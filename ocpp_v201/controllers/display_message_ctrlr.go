package controllers

import (
	"github.com/ChargePi/ocpp-manager/ocpp_v201/component"
	"github.com/ChargePi/ocpp-manager/ocpp_v201/variables"
)

const (
	VariableNameDisplayMessageEnabled             variables.VariableName = "Enabled"
	VariableNameDisplayMessageAvailable           variables.VariableName = "Available"
	VariableNameNumberOfDisplayMessages           variables.VariableName = "DisplayMessages"
	VariableNameDisplayMessageSupportedFormats    variables.VariableName = "SupportedFormats"
	VariableNameDisplayMessageSupportedPriorities variables.VariableName = "SupportedPriorities"
)

func requiredDisplayMessageVariables() []variables.VariableName {
	return []variables.VariableName{
		VariableNameNumberOfDisplayMessages,
		VariableNameDisplayMessageSupportedFormats,
		VariableNameDisplayMessageSupportedPriorities,
	}
}

func optionalDisplayMessageVariables() []variables.VariableName {
	return []variables.VariableName{
		VariableNameDisplayMessageEnabled,
		VariableNameDisplayMessageAvailable,
	}
}

func supportedDisplayMessageVariables() []variables.VariableName {
	return append(requiredDisplayMessageVariables(), optionalDisplayMessageVariables()...)
}

type DisplayCtrlr struct {
	variables          map[variables.VariableName]variables.Variable
	requiredVariables  []variables.VariableName
	supportedVariables []variables.VariableName
}

func (d *DisplayCtrlr) GetName() component.ComponentName {
	//TODO implement me
	panic("implement me")
}

func (d *DisplayCtrlr) GetInstanceId() string {
	//TODO implement me
	panic("implement me")
}

func (d *DisplayCtrlr) RegisterSubComponent(component component.Component) {
	//TODO implement me
	panic("implement me")
}

func (d *DisplayCtrlr) UnregisterSubComponent(component component.Component) {
	//TODO implement me
	panic("implement me")
}

func (d *DisplayCtrlr) GetSubComponents() []component.Component {
	//TODO implement me
	panic("implement me")
}

func (d *DisplayCtrlr) GetRequiredVariables() []variables.Variable {
	//TODO implement me
	panic("implement me")
}

func (d *DisplayCtrlr) GetSupportedVariables() []variables.Variable {
	//TODO implement me
	panic("implement me")
}

func (d *DisplayCtrlr) GetVariable(key variables.VariableName, opts ...component.GetSetVariableOption) (*variables.Variable, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DisplayCtrlr) UpdateVariable(variable variables.VariableName, attribute string, value interface{}, opts ...component.GetSetVariableOption) error {
	//TODO implement me
	panic("implement me")
}

func (d *DisplayCtrlr) Validate(key variables.VariableName) bool {
	//TODO implement me
	panic("implement me")
}

func NewDisplayCtrlr() *DisplayCtrlr {
	return &DisplayCtrlr{
		variables:          make(map[variables.VariableName]variables.Variable),
		requiredVariables:  requiredDisplayMessageVariables(),
		supportedVariables: supportedDisplayMessageVariables(),
	}
}
