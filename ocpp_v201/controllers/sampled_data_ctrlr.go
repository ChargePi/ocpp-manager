package controllers

import (
	"fmt"
	"sync"

	"github.com/ChargePi/ocpp-manager/ocpp_v201/component"
	"github.com/ChargePi/ocpp-manager/ocpp_v201/variables"
)

const (
	VariableNameSampledDataEnabled                     variables.VariableName = "Enabled"
	VariableNameSampledDataAvailable                   variables.VariableName = "Available"
	VariableNameSampledDataSignReadings                variables.VariableName = "SignReadings"
	VariableNameSampledDataTxEndedMeasurands           variables.VariableName = "TxEndedMeasurands"
	VariableNameSampledDataTxEndedInterval             variables.VariableName = "TxEndedInterval"
	VariableNameSampledDataTxStartedMeasurands         variables.VariableName = "TxStartedMeasurands"
	VariableNameSampledDataTxUpdatedMeasurands         variables.VariableName = "TxUpdatedMeasurands"
	VariableNameSampledDataTxUpdatedInterval           variables.VariableName = "TxUpdatedInterval"
	VariableNameSampledDataRegisterValuesWithoutPhases variables.VariableName = "RegisterValuesWithoutPhases"
)

func requiredSampledDataVariables() []variables.VariableName {
	return []variables.VariableName{
		VariableNameSampledDataTxEndedMeasurands,
		VariableNameSampledDataTxEndedInterval,
		VariableNameSampledDataTxStartedMeasurands,
		VariableNameSampledDataTxUpdatedMeasurands,
		VariableNameSampledDataTxUpdatedInterval,
	}
}

func optionalSampledDataVariables() []variables.VariableName {
	return []variables.VariableName{
		VariableNameSampledDataEnabled,
		VariableNameSampledDataAvailable,
		VariableNameSampledDataSignReadings,
		VariableNameSampledDataRegisterValuesWithoutPhases,
	}
}

func supportedSampledDataVariables() []variables.VariableName {
	return append(requiredSampledDataVariables(), optionalSampledDataVariables()...)
}

type SampledDataCtrlr struct {
	mu                 sync.RWMutex
	variables          map[variables.VariableName]*variables.Variable
	requiredVariables  []variables.VariableName
	supportedVariables []variables.VariableName
	instanceId         string
	validator          *variableValidator
}

func (s *SampledDataCtrlr) GetName() component.Name {
	return component.ComponentNameSampledDataCtrlr
}

func (s *SampledDataCtrlr) GetInstanceId() string {
	return s.instanceId
}

func (s *SampledDataCtrlr) RegisterSubComponent(component component.Component) {
	// No-op: controllers do not support sub-components
}

func (s *SampledDataCtrlr) UnregisterSubComponent(component component.Component) {
	// No-op: controllers do not support sub-components
}

func (s *SampledDataCtrlr) GetSubComponents() []component.Component {
	// Controllers do not support sub-components, always return empty slice
	return []component.Component{}
}

func (s *SampledDataCtrlr) GetRequiredVariables() []variables.VariableName {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.requiredVariables
}

func (s *SampledDataCtrlr) GetSupportedVariables() []variables.VariableName {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.supportedVariables
}

func (s *SampledDataCtrlr) GetVariable(key variables.VariableName, opts ...component.GetSetVariableOption) (*variables.Variable, error) {
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

func (s *SampledDataCtrlr) UpdateVariable(variable variables.VariableName, attribute string, value interface{}, opts ...component.GetSetVariableOption) error {
	if !s.validator.IsVariableSupported(variable) {
		return fmt.Errorf("variable %s is not supported", variable)
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	v, exists := s.variables[variable]
	if !exists {
		return fmt.Errorf("variable %s not found", variable)
	}

	return v.UpdateVariableAttribute(attribute, value)
}

func (s *SampledDataCtrlr) Validate(key variables.VariableName) bool {
	if !s.validator.IsVariableSupported(key) {
		return false
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	v, exists := s.variables[key]
	if !exists {
		return false
	}

	return v.Validate()
}

func NewSampledDataCtrlr() *SampledDataCtrlr {
	ctrlr := &SampledDataCtrlr{
		mu:                 sync.RWMutex{},
		variables:          make(map[variables.VariableName]*variables.Variable),
		requiredVariables:  requiredSampledDataVariables(),
		supportedVariables: supportedSampledDataVariables(),
		instanceId:         "sampled-data-ctrlr",
	}

	ctrlr.validator = newVariableValidator(ctrlr)

	return ctrlr
}
