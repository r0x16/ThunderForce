package action

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/r0x16/ThunderForce/src/devices/app"
	"github.com/r0x16/ThunderForce/src/shared/infraestructure/drivers"
)

/**
 * UpdateDevicesAction updates a device
 *
 * @param c
 * @return error
 */
func UpdateDevicesAction(c echo.Context, bundle *drivers.ApplicationBundle) error {
	id := c.Param("id")

	application := getUpdater(bundle)

	deviceData := &app.DeviceUpdateData{}
	if err := c.Bind(deviceData); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	device, err := application.Update(id, deviceData)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, device)
}

/**
 * Get application use case
 *
 * @param repository the device repository
 * @return app.DeviceUpdater
 */
func getUpdater(bundle *drivers.ApplicationBundle) *app.DeviceUpdater {
	repository := DeviceRepository(bundle.Database.Connection)

	return app.CreateDeviceUpdater(repository)
}
