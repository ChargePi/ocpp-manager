package ocpp_v201

import "github.com/ChargePi/ocpp-manager/ocpp_v201/component"

type ChargingStation struct {
	ID                string
	components        map[component.ComponentName]component.Component // Charging station (top level) specific components
	controllerManager Manager
}
