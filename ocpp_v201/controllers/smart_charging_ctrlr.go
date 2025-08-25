package controllers

import (
	"fmt"
	"sync"

	"github.com/ChargePi/ocpp-manager/ocpp_v201/component"
	"github.com/ChargePi/ocpp-manager/ocpp_v201/variables"
)

const (
	VariableNameSmartChargingEnabled             variables.VariableName = "Enabled"
	VariableNameSmartChargingAvailable           variables.VariableName = "Available"
	VariableNameACPhaseSwitchingSupported        variables.VariableName = "ACPhaseSwitchingSupported"
	VariableNameChargingProfileStackLevel        variables.VariableName = "ProfileStackLevel"
	VariableNameChargingScheduleChargingRateUnit variables.VariableName = "RateUnit"
	VariableNamePeriodsPerSchedule               variables.VariableName = "PeriodsPerSchedule"
	VariableNameExternalControlSignalsEnabled    variables.VariableName = "ExternalControlSignalsEnabled"
	VariableNameNotifyChargingLimitWithSchedules variables.VariableName = "NotifyChargingLimitWithSchedules"
	VariableNamePhases3to1                       variables.VariableName = "Phases3to1"
	VariableChargingProfileEntries               variables.VariableName = "Entries"
	VariableLimitChangeSignificance              variables.VariableName = "LimitChangeSignificance"
)

func requiredSmartChargingVariables() []variables.VariableName {
	return []variables.VariableName{
		VariableNameChargingProfileStackLevel,
		VariableNameChargingScheduleChargingRateUnit,
		VariableNamePeriodsPerSchedule,
		VariableChargingProfileEntries,
		VariableLimitChangeSignificance,
	}
}

func optionalSmartChargingVariables() []variables.VariableName {
	return []variables.VariableName{
		VariableNameSmartChargingEnabled,
		VariableNameSmartChargingAvailable,
		VariableNameACPhaseSwitchingSupported,
		VariableNameExternalControlSignalsEnabled,
		VariableNameNotifyChargingLimitWithSchedules,
		VariableNamePhases3to1,
	}
}

func supportedSmartChargingVariables() []variables.VariableName {
	return append(requiredSmartChargingVariables(), optionalSmartChargingVariables()...)
}

type SmartChargingCtrlr struct {
	mu                 sync.RWMutex
	variables          map[variables.VariableName]*variables.Variable
	requiredVariables  []variables.VariableName
	supportedVariables []variables.VariableName
	instanceId         string
	validator          *variableValidator
}

func (s *SmartChargingCtrlr) GetName() component.ComponentName {
	return component.ComponentNameSmartChargingCtrlr
}

func (s *SmartChargingCtrlr) GetInstanceId() string {
	return s.instanceId
}

func (s *SmartChargingCtrlr) RegisterSubComponent(component component.Component) {
	// No-op: controllers do not support sub-components
}

func (s *SmartChargingCtrlr) UnregisterSubComponent(component component.Component) {
	// No-op: controllers do not support sub-components
}

func (s *SmartChargingCtrlr) GetSubComponents() []component.Component {
	// Controllers do not support sub-components, always return empty slice
	return []component.Component{}
}

func (s *SmartChargingCtrlr) GetRequiredVariables() []variables.VariableName {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.requiredVariables
}

func (s *SmartChargingCtrlr) GetSupportedVariables() []variables.VariableName {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.supportedVariables
}

func (s *SmartChargingCtrlr) GetVariable(key variables.VariableName, opts ...component.GetSetVariableOption) (*variables.Variable, error) {
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

func (s *SmartChargingCtrlr) UpdateVariable(variable variables.VariableName, attribute string, value interface{}, opts ...component.GetSetVariableOption) error {
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

func (s *SmartChargingCtrlr) Validate(key variables.VariableName) bool {
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

func NewSmartChargingCtrlr() *SmartChargingCtrlr {
	ctrlr := &SmartChargingCtrlr{
		mu:                 sync.RWMutex{},
		variables:          make(map[variables.VariableName]*variables.Variable),
		requiredVariables:  requiredSmartChargingVariables(),
		supportedVariables: supportedSmartChargingVariables(),
		instanceId:         "smart-charging-ctrlr",
	}

	ctrlr.validator = newVariableValidator(ctrlr)

	return ctrlr
}
