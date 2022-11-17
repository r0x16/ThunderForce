package action

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/r0x16/ThunderForce/src/devices/app"
	"github.com/r0x16/ThunderForce/src/shared/infraestructure/drivers"
)

/**
 * StoreDevicesAction stores a device
 *
 * @param c
 * @return error
 */
func StoreDevicesAction(c echo.Context, bundle *drivers.ApplicationBundle) error {
	application := getPersister(bundle)

	deviceData := &app.DeviceData{}
	if err := c.Bind(deviceData); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	device, err := application.Persist(deviceData)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, device)
}

/**
 * Get application use case
 *
 * @param repository the device repository
 * @return app.DevicePersister
 */
func getPersister(bundle *drivers.ApplicationBundle) *app.DevicePersister {
	repository := DeviceRepository(bundle.Database.Connection)

	return app.CreateDevice(repository)
}
