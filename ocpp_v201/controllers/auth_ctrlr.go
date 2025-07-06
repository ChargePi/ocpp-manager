package controllers

import (
	"github.com/ChargePi/ocpp-manager/ocpp_v201/component"
	"github.com/ChargePi/ocpp-manager/ocpp_v201/variables"
)

const (
	VariableNameAuthEnabled                   variables.VariableName = "Enabled"
	VariableNameAdditionalInfoItemsPerMessage variables.VariableName = "AdditionalInfoItemsPerMessage"
	VariableNameOfflineTxForUnknownIdEnabled  variables.VariableName = "OfflineTxForUnknownIdEnabled"
	VariableNameAuthorizeRemoteStart          variables.VariableName = "AuthorizeRemoteStart"
	VariableNameLocalAuthorizeOffline         variables.VariableName = "LocalAuthorizeOffline"
	VariableNameLocalPreAuthorize             variables.VariableName = "LocalPreAuthorize"
	VariableNameMasterPassGroupId             variables.VariableName = "MasterPassGroupId"
	VariableNameDisableRemoteAuthorization    variables.VariableName = "DisableRemoteAuthorization"
)

func requiredAuthVariables() []variables.VariableName {
	return []variables.VariableName{
		VariableNameAuthorizeRemoteStart,
		VariableNameLocalAuthorizeOffline,
		VariableNameLocalPreAuthorize,
	}
}

func optionalAuthVariables() []variables.VariableName {
	return []variables.VariableName{
		VariableNameAuthEnabled,
		VariableNameAdditionalInfoItemsPerMessage,
		VariableNameOfflineTxForUnknownIdEnabled,
		VariableNameMasterPassGroupId,
		VariableNameDisableRemoteAuthorization,
	}
}

func supportedAuthVariables() []variables.VariableName {
	return append(requiredAuthVariables(), optionalAuthVariables()...)
}

type AuthCtrlr struct {
	variables          map[variables.VariableName]variables.Variable
	requiredVariables  []variables.VariableName
	supportedVariables []variables.VariableName
}

func (a *AuthCtrlr) GetName() component.ComponentName {
	//TODO implement me
	panic("implement me")
}

func (a *AuthCtrlr) GetInstanceId() string {
	//TODO implement me
	panic("implement me")
}

func (a *AuthCtrlr) RegisterSubComponent(component component.Component) {
	//TODO implement me
	panic("implement me")
}

func (a *AuthCtrlr) UnregisterSubComponent(component component.Component) {
	//TODO implement me
	panic("implement me")
}

func (a *AuthCtrlr) GetSubComponents() []component.Component {
	//TODO implement me
	panic("implement me")
}

func (a *AuthCtrlr) GetRequiredVariables() []variables.Variable {
	//TODO implement me
	panic("implement me")
}

func (a *AuthCtrlr) GetSupportedVariables() []variables.Variable {
	//TODO implement me
	panic("implement me")
}

func (a *AuthCtrlr) GetVariable(key variables.VariableName, opts ...component.GetSetVariableOption) (*variables.Variable, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AuthCtrlr) UpdateVariable(variable variables.VariableName, attribute string, value interface{}, opts ...component.GetSetVariableOption) error {
	//TODO implement me
	panic("implement me")
}

func (a *AuthCtrlr) Validate(key variables.VariableName) bool {
	//TODO implement me
	panic("implement me")
}

func NewAuthCtrlr() *AuthCtrlr {
	return &AuthCtrlr{
		variables:          make(map[variables.VariableName]variables.Variable),
		requiredVariables:  requiredAuthVariables(),
		supportedVariables: supportedAuthVariables(),
	}
}
