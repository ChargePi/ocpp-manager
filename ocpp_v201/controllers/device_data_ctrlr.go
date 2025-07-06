package controllers

import (
	"github.com/ChargePi/ocpp-manager/ocpp_v201/component"
	"github.com/ChargePi/ocpp-manager/ocpp_v201/variables"
)

const (
	VariableNameBytesPerMessage        variables.VariableName = "BytesPerMessage"
	VariableNameConfigurationValueSize variables.VariableName = "ConfigurationValueSize"
	VariableNameReportingValueSize     variables.VariableName = "ReportingValueSize"
	VariableNameItemsPerMessage        variables.VariableName = "ItemsPerMessage"
)

func requiredVariables() []variables.VariableName {
	return []variables.VariableName{
		VariableNameBytesPerMessage,
		VariableNameConfigurationValueSize,
		VariableNameReportingValueSize,
		VariableNameItemsPerMessage,
	}
}

func optionalVariables() []variables.VariableName {
	return []variables.VariableName{}
}

func supportedVariables() []variables.VariableName {
	return append(requiredVariables(), optionalVariables()...)
}

type DeviceDataCtrlr struct {
	variables          map[variables.VariableName]variables.Variable
	requiredVariables  []variables.VariableName
	supportedVariables []variables.VariableName
}

func (d *DeviceDataCtrlr) GetName() component.ComponentName {
	//TODO implement me
	panic("implement me")
}

func (d *DeviceDataCtrlr) GetInstanceId() string {
	//TODO implement me
	panic("implement me")
}

func (d *DeviceDataCtrlr) RegisterSubComponent(component component.Component) {
	//TODO implement me
	panic("implement me")
}

func (d *DeviceDataCtrlr) UnregisterSubComponent(component component.Component) {
	//TODO implement me
	panic("implement me")
}

func (d *DeviceDataCtrlr) GetSubComponents() []component.Component {
	//TODO implement me
	panic("implement me")
}

func (d *DeviceDataCtrlr) GetRequiredVariables() []variables.Variable {
	//TODO implement me
	panic("implement me")
}

func (d *DeviceDataCtrlr) GetSupportedVariables() []variables.Variable {
	//TODO implement me
	panic("implement me")
}

func (d *DeviceDataCtrlr) GetVariable(key variables.VariableName, opts ...component.GetSetVariableOption) (*variables.Variable, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DeviceDataCtrlr) UpdateVariable(variable variables.VariableName, attribute string, value interface{}, opts ...component.GetSetVariableOption) error {
	//TODO implement me
	panic("implement me")
}

func (d *DeviceDataCtrlr) Validate(key variables.VariableName) bool {
	//TODO implement me
	panic("implement me")
}

func NewDeviceDataCtrlr() *DeviceDataCtrlr {
	return &DeviceDataCtrlr{}
}
