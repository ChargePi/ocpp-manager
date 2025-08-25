package controllers

import (
	"fmt"
	"sync"

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
		VariableNameCertificateEntries,
		VariableNameOrganizationName,
	}
}

func supportedSecurityVariables() []variables.VariableName {
	return append(requiredSecurityVariables(), optionalSecurityVariables()...)
}

type SecurityCtrlr struct {
	mu                 sync.RWMutex
	variables          map[variables.VariableName]*variables.Variable
	requiredVariables  []variables.VariableName
	supportedVariables []variables.VariableName
	instanceId         string
	validator          *variableValidator
}

func (s *SecurityCtrlr) GetName() component.ComponentName {
	return component.ComponentNameSecurityCtrlr
}

func (s *SecurityCtrlr) GetInstanceId() string {
	return s.instanceId
}

func (s *SecurityCtrlr) RegisterSubComponent(component component.Component) {
	// No-op: controllers do not support sub-components
}

func (s *SecurityCtrlr) UnregisterSubComponent(component component.Component) {
	// No-op: controllers do not support sub-components
}

func (s *SecurityCtrlr) GetSubComponents() []component.Component {
	// Controllers do not support sub-components, always return empty slice
	return []component.Component{}
}

func (s *SecurityCtrlr) GetRequiredVariables() []variables.VariableName {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.requiredVariables
}

func (s *SecurityCtrlr) GetSupportedVariables() []variables.VariableName {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.supportedVariables
}

func (s *SecurityCtrlr) GetVariable(key variables.VariableName, opts ...component.GetSetVariableOption) (*variables.Variable, error) {
	if !s.validator.IsVariableSupported(key) {
		return nil, fmt.Errorf("variable %s is not supported", key)
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	variable, exists := s.variables[key]
	if !exists {
		return nil, fmt.Errorf("variable %s not found", key)
	}

	return variable, nil
}

func (s *SecurityCtrlr) UpdateVariable(variable variables.VariableName, attribute string, value interface{}, opts ...component.GetSetVariableOption) error {
	if !s.validator.IsVariableSupported(variable) {
		return fmt.Errorf("variable %s is not supported", variable)
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	v, exists := s.variables[variable]
	if !exists {
		return fmt.Errorf("variable %s not found", variable)
	}

	return v.UpdateVariableAttribute(attribute, value)
}

func (s *SecurityCtrlr) Validate(key variables.VariableName) bool {
	if !s.validator.IsVariableSupported(key) {
		return false
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	v, exists := s.variables[key]
	if !exists {
		return false
	}

	return v.Validate()
}

func NewSecurityCtrlr() *SecurityCtrlr {
	ctrlr := &SecurityCtrlr{
		mu:                 sync.RWMutex{},
		variables:          make(map[variables.VariableName]*variables.Variable),
		requiredVariables:  requiredSecurityVariables(),
		supportedVariables: supportedSecurityVariables(),
		instanceId:         "security-ctrlr",
	}

	ctrlr.validator = newVariableValidator(ctrlr)

	return ctrlr
}
