package controllers

import (
	"fmt"
	"sync"

	"github.com/ChargePi/ocpp-manager/ocpp_v201/component"
	"github.com/ChargePi/ocpp-manager/ocpp_v201/variables"
)

const (
	VariableNameDefaultMessageTimeout                  variables.VariableName = "DefaultMessageTimeout"
	VariableNameNetworkProfileConnectionAttempts       variables.VariableName = "NetworkProfileConnectionAttempts"
	VariableNameNetworkConfigurationPriority           variables.VariableName = "NetworkConfigurationPriority"
	VariableNameHeartbeatInterval                      variables.VariableName = "HeartbeatInterval"
	VariableNameFileTransferProtocols                  variables.VariableName = "FileTransferProtocols"
	VariableNameMessageTimeout                         variables.VariableName = "MessageTimeout"
	VariableNameActiveNetworkProfile                   variables.VariableName = "ActiveNetworkProfile"
	VariableNameOfflineThreshold                       variables.VariableName = "OfflineThreshold"
	VariableNameQueueAllMessages                       variables.VariableName = "QueueAllMessages"
	VariableNameMessageAttempts                        variables.VariableName = "MessageAttempts"
	VariableNameMessageAttemptInterval                 variables.VariableName = "MessageAttemptInterval"
	VariableNameMessageAttemptsTransactionEvent        variables.VariableName = "MessageAttemptsTransactionEvent"
	VariableNameMessageAttemptIntervalTransactionEvent variables.VariableName = "MessageAttemptIntervalTransactionEvent"
	VariableNameUnlockOnEVSideDisconnect               variables.VariableName = "UnlockOnEVSideDisconnect"
	VariableNameWebSocketPingInterval                  variables.VariableName = "WebSocketPingInterval"
	VariableNameResetRetries                           variables.VariableName = "ResetRetries"
	VariableNameFieldLength                            variables.VariableName = "FieldLength"
	VariableNamePublicKeyWithSignedMeterValue          variables.VariableName = "PublicKeyWithSignedMeterValue"
)

func requiredOcppCommCtrlrVariables() []variables.VariableName {
	return []variables.VariableName{
		VariableNameDefaultMessageTimeout,
		VariableNameFileTransferProtocols,
		VariableNameOfflineThreshold,
		VariableNameNetworkProfileConnectionAttempts,
		VariableNameMessageAttempts,
		VariableNameMessageAttemptInterval,
		VariableNameMessageAttemptsTransactionEvent,
		VariableNameMessageAttemptIntervalTransactionEvent,
		VariableNameUnlockOnEVSideDisconnect,
		VariableNameResetRetries,
	}
}

func optionalOcppCommCtrlrVariables() []variables.VariableName {
	return []variables.VariableName{
		VariableNameNetworkConfigurationPriority,
		VariableNameHeartbeatInterval,
		VariableNameMessageTimeout,
		VariableNameActiveNetworkProfile,
		VariableNameQueueAllMessages,
		VariableNameWebSocketPingInterval,
		VariableNameFieldLength,
		VariableNamePublicKeyWithSignedMeterValue,
	}
}

func supportedOcppCommCtrlrVariables() []variables.VariableName {
	return append(requiredOcppCommCtrlrVariables(), optionalOcppCommCtrlrVariables()...)
}

type OCPPCommCtrlr struct {
	mu                 sync.RWMutex
	variables          map[variables.VariableName]*variables.Variable
	requiredVariables  []variables.VariableName
	supportedVariables []variables.VariableName
	instanceId         string
	validator          *variableValidator
}

func (ctrlr *OCPPCommCtrlr) GetName() component.ComponentName {
	return component.ComponentNameOCPPCommCtrlr
}

func (ctrlr *OCPPCommCtrlr) GetInstanceId() string {
	return ctrlr.instanceId
}

func (ctrlr *OCPPCommCtrlr) RegisterSubComponent(component component.Component) {
	// No-op: controllers do not support sub-components
}

func (ctrlr *OCPPCommCtrlr) UnregisterSubComponent(component component.Component) {
	// No-op: controllers do not support sub-components
}

func (ctrlr *OCPPCommCtrlr) GetSubComponents() []component.Component {
	// Controllers do not support sub-components, always return empty slice
	return []component.Component{}
}

func (ctrlr *OCPPCommCtrlr) GetRequiredVariables() []variables.VariableName {
	ctrlr.mu.RLock()
	defer ctrlr.mu.RUnlock()

	return ctrlr.requiredVariables
}

func (ctrlr *OCPPCommCtrlr) GetSupportedVariables() []variables.VariableName {
	ctrlr.mu.RLock()
	defer ctrlr.mu.RUnlock()

	return ctrlr.supportedVariables
}

func (ctrlr *OCPPCommCtrlr) GetVariable(key variables.VariableName, opts ...component.GetSetVariableOption) (*variables.Variable, error) {
	if !ctrlr.validator.IsVariableSupported(key) {
		return nil, fmt.Errorf("variable %s is not supported", key)
	}

	ctrlr.mu.RLock()
	defer ctrlr.mu.RUnlock()

	variable, exists := ctrlr.variables[key]
	if !exists {
		return nil, fmt.Errorf("variable %s not found", key)
	}
	return variable, nil
}

func (ctrlr *OCPPCommCtrlr) UpdateVariable(variable variables.VariableName, attribute string, value interface{}, opts ...component.GetSetVariableOption) error {
	if !ctrlr.validator.IsVariableSupported(variable) {
		return fmt.Errorf("variable %s is not supported", variable)
	}

	ctrlr.mu.Lock()
	defer ctrlr.mu.Unlock()

	v, exists := ctrlr.variables[variable]
	if !exists {
		return fmt.Errorf("variable %s not found", variable)
	}

	return v.UpdateVariableAttribute(attribute, value)
}

func (ctrlr *OCPPCommCtrlr) Validate(key variables.VariableName) bool {
	if !ctrlr.validator.IsVariableSupported(key) {
		return false
	}

	ctrlr.mu.RLock()
	defer ctrlr.mu.RUnlock()

	v, exists := ctrlr.variables[key]
	if !exists {
		return false
	}

	return v.Validate()
}

func NewOCPPCommCtrlr() *OCPPCommCtrlr {
	ctrlr := &OCPPCommCtrlr{
		mu:                 sync.RWMutex{},
		variables:          map[variables.VariableName]*variables.Variable{},
		requiredVariables:  requiredOcppCommCtrlrVariables(),
		supportedVariables: supportedOcppCommCtrlrVariables(),
		instanceId:         "ocpp-comm-ctrlr-1",
	}

	ctrlr.validator = newVariableValidator(ctrlr)

	return ctrlr
}
