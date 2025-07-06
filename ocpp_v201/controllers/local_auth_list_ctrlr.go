package controllers

import (
	"github.com/ChargePi/ocpp-manager/ocpp_v201/component"
	"github.com/ChargePi/ocpp-manager/ocpp_v201/variables"
)

const (
	VariableNameLocalAuthListEnabled                variables.VariableName = "Enabled"
	VariableNameLocalAuthListEntries                variables.VariableName = "Entries"
	VariableNameLocalAuthListItemsPerMessage        variables.VariableName = "ItemsPerMessage"
	VariableNameLocalAuthListBytesPerMessage        variables.VariableName = "BytesPerMessage"
	VariableNameLocalAuthListStorage                variables.VariableName = "Storage"
	VariableNameLocalAuthListDisablePostAuthorize   variables.VariableName = "DisablePostAuthorize"
	VariableNameLocalAuthListSupportsExpiryDateTime variables.VariableName = "SupportsExpiryDateTime"
)

func requiredLocalAuthListVariables() []variables.VariableName {
	return []variables.VariableName{}
}

func optionalLocalAuthListVariables() []variables.VariableName {
	return []variables.VariableName{
		VariableNameLocalAuthListItemsPerMessage,
		VariableNameLocalAuthListBytesPerMessage,
		VariableNameLocalAuthListStorage,
		VariableNameLocalAuthListDisablePostAuthorize,
		VariableNameLocalAuthListSupportsExpiryDateTime,
		VariableNameLocalAuthListEnabled,
		VariableNameLocalAuthListEntries,
	}
}

func supportedLocalAuthListVariables() []variables.VariableName {
	return append(requiredLocalAuthListVariables(), optionalLocalAuthListVariables()...)
}

type LocalAuthListCtrlr struct {
	variables          map[variables.VariableName]variables.Variable
	requiredVariables  []variables.VariableName
	supportedVariables []variables.VariableName
}

func (l *LocalAuthListCtrlr) GetName() component.ComponentName {
	//TODO implement me
	panic("implement me")
}

func (l *LocalAuthListCtrlr) GetInstanceId() string {
	//TODO implement me
	panic("implement me")
}

func (l *LocalAuthListCtrlr) RegisterSubComponent(component component.Component) {
	//TODO implement me
	panic("implement me")
}

func (l *LocalAuthListCtrlr) UnregisterSubComponent(component component.Component) {
	//TODO implement me
	panic("implement me")
}

func (l *LocalAuthListCtrlr) GetSubComponents() []component.Component {
	//TODO implement me
	panic("implement me")
}

func (l *LocalAuthListCtrlr) GetRequiredVariables() []variables.Variable {
	//TODO implement me
	panic("implement me")
}

func (l *LocalAuthListCtrlr) GetSupportedVariables() []variables.Variable {
	//TODO implement me
	panic("implement me")
}

func (l *LocalAuthListCtrlr) GetVariable(key variables.VariableName, opts ...component.GetSetVariableOption) (*variables.Variable, error) {
	//TODO implement me
	panic("implement me")
}

func (l *LocalAuthListCtrlr) UpdateVariable(variable variables.VariableName, attribute string, value interface{}, opts ...component.GetSetVariableOption) error {
	//TODO implement me
	panic("implement me")
}

func (l *LocalAuthListCtrlr) Validate(key variables.VariableName) bool {
	//TODO implement me
	panic("implement me")
}

func NewLocalAuthListCtrlr() *LocalAuthListCtrlr {
	return &LocalAuthListCtrlr{
		variables:          make(map[variables.VariableName]variables.Variable),
		requiredVariables:  requiredLocalAuthListVariables(),
		supportedVariables: supportedLocalAuthListVariables(),
	}
}
