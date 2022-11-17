package action

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/r0x16/ThunderForce/src/devices/app"
	"github.com/r0x16/ThunderForce/src/shared/infraestructure/drivers"
)

/**
 * ListDevicesAction lists all devices
 *
 * @param c
 * @return error
 */
func ListDevicesAction(c echo.Context, bundle *drivers.ApplicationBundle) error {
	application := getLister(bundle)

	devices, err := application.List()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, devices)
}

/**
* Get application use case
*
* @param repository the device repository
* @return app.DeviceLister
 */
func getLister(bundle *drivers.ApplicationBundle) *app.DeviceLister {
	repository := DeviceRepository(bundle.Database.Connection)

	return app.ListDevices(repository)
}
