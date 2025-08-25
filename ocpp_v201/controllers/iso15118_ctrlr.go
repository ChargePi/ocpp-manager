package controllers

import (
	"fmt"
	"sync"

	"github.com/ChargePi/ocpp-manager/ocpp_v201/component"
	"github.com/ChargePi/ocpp-manager/ocpp_v201/variables"
)

const (
	VariableNameCentralContractValidationAllowed               variables.VariableName = "CentralContractValidationAllowed"
	VariableNameContractValidationOffline                      variables.VariableName = "ContractValidationOffline"
	VariableNameProtocolSupportedByEV                          variables.VariableName = "ProtocolSupportedByEV"
	VariableNameProtocolAgreed                                 variables.VariableName = "ProtocolAgreed"
	VariableNameISO15118PnCEnabled                             variables.VariableName = "PnCEnabled"
	VariableNameISO15118V2GCertificateInstallationEnabled      variables.VariableName = "V2GCertificateInstallationEnabled"
	VariableNameISO15118ContractCertificateInstallationEnabled variables.VariableName = "ContractCertificateInstallationEnabled"
	VariableNameISO15118RequestMeteringReceipt                 variables.VariableName = "RequestMeteringReceipt"
	VariableNameISO15118SeccId                                 variables.VariableName = "SeccId"
	VariableNameISO15118CountryName                            variables.VariableName = "CountryName"
	VariableNameISO15118EvseId                                 variables.VariableName = "ISO15118EvseId"
)

func requiredVariablesISO15118Ctrlr() []variables.VariableName {
	return []variables.VariableName{
		VariableNameCentralContractValidationAllowed,
		VariableNameContractValidationOffline,
		VariableNameProtocolSupportedByEV,
		VariableNameProtocolAgreed,
	}
}

func optionalVariablesISO15118Ctrlr() []variables.VariableName {
	return []variables.VariableName{
		VariableNameISO15118PnCEnabled,
		VariableNameISO15118V2GCertificateInstallationEnabled,
		VariableNameISO15118ContractCertificateInstallationEnabled,
		VariableNameISO15118RequestMeteringReceipt,
		VariableNameISO15118SeccId,
		VariableNameISO15118CountryName,
		VariableNameISO15118EvseId,
	}
}

func supportedVariablesISO15118Ctrlr() []variables.VariableName {
	return append(requiredVariablesISO15118Ctrlr(), optionalVariablesISO15118Ctrlr()...)
}

type ISO15118Ctrlr struct {
	mu                 sync.RWMutex
	variables          map[variables.VariableName]*variables.Variable
	requiredVariables  []variables.VariableName
	supportedVariables []variables.VariableName
	instanceId         string
	validator          *variableValidator
}

func (I *ISO15118Ctrlr) GetName() component.ComponentName {
	return component.ComponentNameISO15118Ctrlr
}

func (I *ISO15118Ctrlr) GetInstanceId() string {
	return I.instanceId
}

func (I *ISO15118Ctrlr) RegisterSubComponent(component component.Component) {
	// No-op: controllers do not support sub-components
}

func (I *ISO15118Ctrlr) UnregisterSubComponent(component component.Component) {
	// No-op: controllers do not support sub-components
}

func (I *ISO15118Ctrlr) GetSubComponents() []component.Component {
	// Controllers do not support sub-components, always return empty slice
	return []component.Component{}
}

func (I *ISO15118Ctrlr) GetRequiredVariables() []variables.VariableName {
	I.mu.RLock()
	defer I.mu.RUnlock()

	return I.requiredVariables
}

func (I *ISO15118Ctrlr) GetSupportedVariables() []variables.VariableName {
	I.mu.RLock()
	defer I.mu.RUnlock()

	return I.supportedVariables
}

func (I *ISO15118Ctrlr) GetVariable(key variables.VariableName, opts ...component.GetSetVariableOption) (*variables.Variable, error) {
	if !I.validator.IsVariableSupported(key) {
		return nil, fmt.Errorf("variable %s is not supported", key)
	}

	I.mu.RLock()
	defer I.mu.RUnlock()

	variable, exists := I.variables[key]
	if !exists {
		return nil, fmt.Errorf("variable %s not found", key)
	}
	return variable, nil
}

func (I *ISO15118Ctrlr) UpdateVariable(variable variables.VariableName, attribute string, value interface{}, opts ...component.GetSetVariableOption) error {
	if !I.validator.IsVariableSupported(variable) {
		return fmt.Errorf("variable %s is not supported", variable)
	}

	I.mu.Lock()
	defer I.mu.Unlock()

	v, exists := I.variables[variable]
	if !exists {
		return fmt.Errorf("variable %s not found", variable)
	}

	return v.UpdateVariableAttribute(attribute, value)
}

func (I *ISO15118Ctrlr) Validate(key variables.VariableName) bool {
	if !I.validator.IsVariableSupported(key) {
		return false
	}

	I.mu.RLock()
	defer I.mu.RUnlock()

	v, exists := I.variables[key]
	if !exists {
		return false
	}

	return v.Validate()
}

func NewISO15118Ctrlr() *ISO15118Ctrlr {
	ctrlr := &ISO15118Ctrlr{
		mu:                 sync.RWMutex{},
		variables:          map[variables.VariableName]*variables.Variable{},
		requiredVariables:  requiredVariablesISO15118Ctrlr(),
		supportedVariables: supportedVariablesISO15118Ctrlr(),
		instanceId:         "iso15118-ctrlr",
	}

	ctrlr.validator = newVariableValidator(ctrlr)

	return ctrlr
}
