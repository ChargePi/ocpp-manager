package ocpp_v201

import "github.com/ChargePi/ocpp-manager/ocpp_v201/component"

type ChargingStation struct {
	components map[component.ComponentName]component.Component // Charging station (top level) specific components
}

type EVSE struct {
	ID         int
	components map[component.ComponentName]component.Component
}

type Connector struct {
	ID         int
	components map[component.ComponentName]component.Component
}
