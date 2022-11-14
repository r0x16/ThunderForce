package module

import (
	"github.com/r0x16/ThunderForce/src/devices/infraestructure/action"
	"github.com/r0x16/ThunderForce/src/shared/domain"
	"github.com/r0x16/ThunderForce/src/shared/infraestructure/drivers"
)

/**
 * DevicesModule represents the links between devices routes and actions
 */
type DevicesModule struct {
	Bundle *drivers.ApplicationBundle
}

var _ domain.ApplicationModule = &DevicesModule{}

/**
 *
 * Initializes devices module and set up their routes
 *
 */
func (m *DevicesModule) Setup() {
	devices := m.Bundle.Server.Group("/devices")
	devices.GET("", action.ListDevicesAction)
	devices.GET(":id", action.ShowDevicesAction)
	devices.POST("", action.StoreDevicesAction)
	devices.PUT(":id", action.UpdateDevicesAction)
	devices.DELETE(":id", action.DeleteDevicesAction)
}
