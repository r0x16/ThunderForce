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

	devices.GET("/", m.Bundle.ActionInjection(action.ListDevicesAction))

	devices.GET("/:id", m.Bundle.ActionInjection(action.ShowDevicesAction))

	devices.POST("/", m.Bundle.ActionInjection(action.StoreDevicesAction))

	devices.PUT("/:id", m.Bundle.ActionInjection(action.UpdateDevicesAction))

	devices.DELETE("/:id", m.Bundle.ActionInjection(action.DeleteDevicesAction))

}
