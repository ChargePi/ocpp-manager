package controllers

import (
	"github.com/ChargePi/ocpp-manager/ocpp_v201/component"
	"github.com/ChargePi/ocpp-manager/ocpp_v201/variables"
)

const (
	VariableNameBasicAuthPassword              variables.VariableName = "BasicAuthPassword"
	VariableNameIdentity                       variables.VariableName = "Identity"
	VariableNameOrganizationName               variables.VariableName = "OrganizationName"
	VariableNameCertificateEntries             variables.VariableName = "CertificateEntries"
	VariableNameAdditionalRootCertificateCheck variables.VariableName = "AdditionalRootCertificateCheck"
	VariableNameSecurityProfile                variables.VariableName = "SecurityProfile"
	VariableNameMaxCertificateChainSize        variables.VariableName = "MaxCertificateChainSize"
	VariableNameCertSigningWaitMinimum         variables.VariableName = "CertSigningWaitMinimum"
	VariableNameCertSigningRepeatTimes         variables.VariableName = "CertSigningRepeatTimes"
)

func requiredSecurityVariables() []variables.VariableName {
	return []variables.VariableName{
		VariableNameOrganizationName,
		VariableNameCertificateEntries,
		VariableNameSecurityProfile,
	}
}

func optionalSecurityVariables() []variables.VariableName {
	return []variables.VariableName{
		VariableNameBasicAuthPassword,
		VariableNameIdentity,
		VariableNameAdditionalRootCertificateCheck,
		VariableNameCertSigningRepeatTimes,
		VariableNameMaxCertificateChainSize,
		VariableNameCertSigningWaitMinimum,
	}
}

func supportedSecurityVariables() []variables.VariableName {
	return append(requiredSecurityVariables(), optionalSecurityVariables()...)
}

type SecurityCtrlr struct {
	variables          map[variables.VariableName]variables.Variable
	requiredVariables  []variables.VariableName
	supportedVariables []variables.VariableName
}

func (s *SecurityCtrlr) GetName() component.ComponentName {
	//TODO implement me
	panic("implement me")
}

func (s *SecurityCtrlr) GetInstanceId() string {
	//TODO implement me
	panic("implement me")
}

func (s *SecurityCtrlr) RegisterSubComponent(component component.Component) {
	//TODO implement me
	panic("implement me")
}

func (s *SecurityCtrlr) UnregisterSubComponent(component component.Component) {
	//TODO implement me
	panic("implement me")
}

func (s *SecurityCtrlr) GetSubComponents() []component.Component {
	//TODO implement me
	panic("implement me")
}

func (s *SecurityCtrlr) GetRequiredVariables() []variables.Variable {
	//TODO implement me
	panic("implement me")
}

func (s *SecurityCtrlr) GetSupportedVariables() []variables.Variable {
	//TODO implement me
	panic("implement me")
}

func (s *SecurityCtrlr) GetVariable(key variables.VariableName, opts ...component.GetSetVariableOption) (*variables.Variable, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SecurityCtrlr) UpdateVariable(variable variables.VariableName, attribute string, value interface{}, opts ...component.GetSetVariableOption) error {
	//TODO implement me
	panic("implement me")
}

func (s *SecurityCtrlr) Validate(key variables.VariableName) bool {
	//TODO implement me
	panic("implement me")
}

func NewSecurityCtrlr() *SecurityCtrlr {
	return &SecurityCtrlr{
		variables:          make(map[variables.VariableName]variables.Variable),
		requiredVariables:  requiredSecurityVariables(),
		supportedVariables: supportedSecurityVariables(),
	}
}
