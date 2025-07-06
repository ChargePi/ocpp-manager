package controllers

import (
	"github.com/ChargePi/ocpp-manager/ocpp_v201/component"
	"github.com/ChargePi/ocpp-manager/ocpp_v201/variables"
)

const (
	VariableNameAuthCacheEnabled              variables.VariableName = "Enabled"
	VariableNameAuthCacheAvailable            variables.VariableName = "Available"
	VariableNameAuthCacheLifeTime             variables.VariableName = "LifeTime"
	VariableNameAuthCacheStorage              variables.VariableName = "Storage"
	VariableNameAuthCachePolicy               variables.VariableName = "Policy"
	VariableNameAuthCacheDisablePostAuthorize variables.VariableName = "DisablePostAuthorize"
)

func requiredAuthCacheVariables() []variables.VariableName {
	return []variables.VariableName{}
}

func optionalAuthCacheVariables() []variables.VariableName {
	return []variables.VariableName{
		VariableNameAuthCacheAvailable,
		VariableNameAuthCacheEnabled,
		VariableNameAuthCacheLifeTime,
		VariableNameAuthCacheStorage,
		VariableNameAuthCachePolicy,
		VariableNameAuthCacheDisablePostAuthorize,
	}
}

func supportedAuthCacheVariables() []variables.VariableName {
	return append(requiredAuthCacheVariables(), optionalAuthCacheVariables()...)
}

type AuthCacheCtrlr struct {
	variables          map[variables.VariableName]variables.Variable
	supportedVariables []variables.VariableName
	requiredVariables  []variables.VariableName
}

func (a *AuthCacheCtrlr) GetName() component.ComponentName {
	//TODO implement me
	panic("implement me")
}

func (a *AuthCacheCtrlr) GetInstanceId() string {
	//TODO implement me
	panic("implement me")
}

func (a *AuthCacheCtrlr) RegisterSubComponent(component component.Component) {
	//TODO implement me
	panic("implement me")
}

func (a *AuthCacheCtrlr) UnregisterSubComponent(component component.Component) {
	//TODO implement me
	panic("implement me")
}

func (a *AuthCacheCtrlr) GetSubComponents() []component.Component {
	//TODO implement me
	panic("implement me")
}

func (a *AuthCacheCtrlr) GetRequiredVariables() []variables.Variable {
	//TODO implement me
	panic("implement me")
}

func (a *AuthCacheCtrlr) GetSupportedVariables() []variables.Variable {
	//TODO implement me
	panic("implement me")
}

func (a *AuthCacheCtrlr) GetVariable(key variables.VariableName, opts ...component.GetSetVariableOption) (*variables.Variable, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AuthCacheCtrlr) UpdateVariable(variable variables.VariableName, attribute string, value interface{}, opts ...component.GetSetVariableOption) error {
	//TODO implement me
	panic("implement me")
}

func (a *AuthCacheCtrlr) Validate(key variables.VariableName) bool {
	//TODO implement me
	panic("implement me")
}

func NewAuthCacheCtrlr() *AuthCacheCtrlr {
	return &AuthCacheCtrlr{
		variables:          make(map[variables.VariableName]variables.Variable),
		supportedVariables: supportedAuthCacheVariables(),
		requiredVariables:  requiredAuthCacheVariables(),
	}
}
