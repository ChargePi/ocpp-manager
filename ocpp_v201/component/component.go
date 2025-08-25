package component

import "github.com/ChargePi/ocpp-manager/ocpp_v201/variables"

type Component interface {

	// GetName Essentially a component type.
	GetName() ComponentName

	// GetInstanceId returns the unique instance ID of this component.
	GetInstanceId() string

	// RegisterSubComponent registers a sub-component to this component.
	RegisterSubComponent(component Component)

	// UnregisterSubComponent unregisters a sub-component from this component.
	UnregisterSubComponent(component Component)

	// GetSubComponents returns all sub-components of this component.
	GetSubComponents() []Component

	// GetRequiredVariables returns required variables for this component
	GetRequiredVariables() []variables.VariableName

	// GetSupportedVariables returns supported variables (both required and optional) for this component
	GetSupportedVariables() []variables.VariableName

	// GetVariable retrieves a variable by its name.
	GetVariable(key variables.VariableName, opts ...GetSetVariableOption) (*variables.Variable, error)

	// UpdateVariable updates a variable's attribute with a new value.
	UpdateVariable(variable variables.VariableName, attribute string, value interface{}, opts ...GetSetVariableOption) error

	// Validate checks if the variable is valid for this component.
	Validate(key variables.VariableName) bool
}

type ComponentName string

const (
	ComponentNameOCPPCommCtrlr       ComponentName = "OCPPCommCtrlr"
	ComponentNameLocalAuthListCtrlr  ComponentName = "LocalAuthListCtrlr"
	ComponentNameTxCtrlr             ComponentName = "TxCtrlr"
	ComponentNameDeviceDataCtrlr     ComponentName = "DeviceDataCtrlr"
	ComponentNameSecurityCtrlr       ComponentName = "SecurityCtrlr"
	ComponentNameClockCtrlr          ComponentName = "ClockCtrlr"
	ComponentNameCustomizationCtrlr  ComponentName = "CustomizationCtrlr"
	ComponentNameSampledDataCtrlr    ComponentName = "SampledDataCtrlr"
	ComponentNameAlignedDataCtrlr    ComponentName = "AlignedDataCtrlr"
	ComponentNameReservationCtrlr    ComponentName = "ReservationCtrlr"
	ComponentNameSmartChargingCtrlr  ComponentName = "SmartChargingCtrlr"
	ComponentNameTariffCostCtrlr     ComponentName = "TariffCostCtrlr"
	ComponentNameMonitoringCtrlr     ComponentName = "MonitoringCtrlr"
	ComponentNameDisplayMessageCtrlr ComponentName = "DisplayMessageCtrlr"
	ComponentNameISO15118Ctrlr       ComponentName = "ISO15118Ctrlr"
	ComponentNameAuthCtrlr           ComponentName = "AuthCtrlr"
	ComponentNameAuthCacheCtrlr      ComponentName = "AuthCacheCtrlr"
	ComponentNameChargingStation     ComponentName = "ChargingStation"
	ComponentNameEVSE                ComponentName = "EVSE"
	ComponentNameConnector           ComponentName = "Connector"
	ComponentNameConnectedEV         ComponentName = "ConnectedEV"
)
