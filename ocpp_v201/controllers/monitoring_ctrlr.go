package controllers

import (
	"github.com/ChargePi/ocpp-manager/ocpp_v201/component"
	"github.com/ChargePi/ocpp-manager/ocpp_v201/variables"
)

const (
	VariableNameMonitoringEnabled                      variables.VariableName = "Enabled"
	VariableNameMonitoringAvailable                    variables.VariableName = "Available"
	VariableNameItemsPerMessageClearVariableMonitoring variables.VariableName = "ItemsPerMessage"
	VariableNameItemsPerMessageSetVariableMonitoring   variables.VariableName = "ItemsPerMessage"
	VariableNameClearVariableMonitoring                variables.VariableName = "BytesPerMessage"
	VariableNameBytesPerMessageSetVariableMonitoring   variables.VariableName = "BytesPerMessage"
	VariableNameOfflineMonitoringEventQueuingSeverity  variables.VariableName = "OfflineQueuingSeverity"
	VariableNameActiveMonitoringBase                   variables.VariableName = "ActiveMonitoringBase"
	VariableNameActiveMonitoringLevel                  variables.VariableName = "ActiveMonitoringLevel"
)

func requiredMonitoringVariables() []variables.VariableName {
	return []variables.VariableName{
		VariableNameItemsPerMessageSetVariableMonitoring,
		VariableNameBytesPerMessageSetVariableMonitoring,
	}
}

func optionalMonitoringVariables() []variables.VariableName {
	return []variables.VariableName{
		VariableNameClearVariableMonitoring,
		VariableNameMonitoringEnabled,
		VariableNameMonitoringAvailable,
		VariableNameItemsPerMessageClearVariableMonitoring,
		VariableNameOfflineMonitoringEventQueuingSeverity,
		VariableNameActiveMonitoringBase,
		VariableNameActiveMonitoringLevel,
	}
}

func supportedMonitoringVariables() []variables.VariableName {
	return append(requiredMonitoringVariables(), optionalMonitoringVariables()...)
}

type MonitoringCtrlr struct {
	variables          map[variables.VariableName]variables.Variable
	requiredVariables  []variables.VariableName
	supportedVariables []variables.VariableName
}

func (m *MonitoringCtrlr) GetName() component.ComponentName {
	//TODO implement me
	panic("implement me")
}

func (m *MonitoringCtrlr) GetInstanceId() string {
	//TODO implement me
	panic("implement me")
}

func (m *MonitoringCtrlr) RegisterSubComponent(component component.Component) {
	//TODO implement me
	panic("implement me")
}

func (m *MonitoringCtrlr) UnregisterSubComponent(component component.Component) {
	//TODO implement me
	panic("implement me")
}

func (m *MonitoringCtrlr) GetSubComponents() []component.Component {
	//TODO implement me
	panic("implement me")
}

func (m *MonitoringCtrlr) GetRequiredVariables() []variables.Variable {
	//TODO implement me
	panic("implement me")
}

func (m *MonitoringCtrlr) GetSupportedVariables() []variables.Variable {
	//TODO implement me
	panic("implement me")
}

func (m *MonitoringCtrlr) GetVariable(key variables.VariableName, opts ...component.GetSetVariableOption) (*variables.Variable, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MonitoringCtrlr) UpdateVariable(variable variables.VariableName, attribute string, value interface{}, opts ...component.GetSetVariableOption) error {
	//TODO implement me
	panic("implement me")
}

func (m *MonitoringCtrlr) Validate(key variables.VariableName) bool {
	//TODO implement me
	panic("implement me")
}

func NewMonitoringCtrlr() *MonitoringCtrlr {
	return &MonitoringCtrlr{
		variables:          make(map[variables.VariableName]variables.Variable),
		requiredVariables:  requiredMonitoringVariables(),
		supportedVariables: supportedMonitoringVariables(),
	}
}
