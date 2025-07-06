package controllers

import (
	"github.com/ChargePi/ocpp-manager/ocpp_v201/component"
	"github.com/ChargePi/ocpp-manager/ocpp_v201/variables"
)

const (
	VariableNameCustomImplementationEnabled variables.VariableName = "CustomImplementationEnabled"
)

type CustomizationCtrlr struct{}

func (c *CustomizationCtrlr) GetName() component.ComponentName {
	//TODO implement me
	panic("implement me")
}

func (c *CustomizationCtrlr) GetInstanceId() string {
	//TODO implement me
	panic("implement me")
}

func (c *CustomizationCtrlr) RegisterSubComponent(component component.Component) {
	//TODO implement me
	panic("implement me")
}

func (c *CustomizationCtrlr) UnregisterSubComponent(component component.Component) {
	//TODO implement me
	panic("implement me")
}

func (c *CustomizationCtrlr) GetSubComponents() []component.Component {
	//TODO implement me
	panic("implement me")
}

func (c *CustomizationCtrlr) GetRequiredVariables() []variables.Variable {
	//TODO implement me
	panic("implement me")
}

func (c *CustomizationCtrlr) GetSupportedVariables() []variables.Variable {
	//TODO implement me
	panic("implement me")
}

func (c *CustomizationCtrlr) GetVariable(key variables.VariableName, opts ...component.GetSetVariableOption) (*variables.Variable, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CustomizationCtrlr) UpdateVariable(variable variables.VariableName, attribute string, value interface{}, opts ...component.GetSetVariableOption) error {
	//TODO implement me
	panic("implement me")
}

func (c *CustomizationCtrlr) Validate(key variables.VariableName) bool {
	//TODO implement me
	panic("implement me")
}

func NewCustomizationCtrlr() *CustomizationCtrlr {
	return &CustomizationCtrlr{}
}
