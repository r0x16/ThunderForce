package action

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/r0x16/ThunderForce/src/devices/app"
	"github.com/r0x16/ThunderForce/src/shared/infraestructure/drivers"
)

/**
 * DeleteDevicesAction deletes a device
 *
 * @param c the echo context
 * @param bundle the application bundle
 * @return error
 */
func DeleteDevicesAction(c echo.Context, bundle *drivers.ApplicationBundle) error {
	id := c.Param("id")

	application := getDeleter(bundle)

	err := application.Delete(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, nil)
}

/**
 * Get application use case
 *
 * @param bundle the application bundle
 * @return app.DeviceDeleter
 */
func getDeleter(bundle *drivers.ApplicationBundle) *app.DeviceDeleter {
	repository := DeviceRepository(bundle.Database.Connection)

	return app.DeleteDevice(repository)
}
