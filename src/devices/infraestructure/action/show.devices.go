package action

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/r0x16/ThunderForce/src/devices/app"
	"github.com/r0x16/ThunderForce/src/shared/infraestructure/drivers"
)

/**
 * ShowDevicesAction shows a device
 *
 * @param c echo.Context
 * @param bundle the application bundle
 * @return error
 */
func ShowDevicesAction(c echo.Context, bundle *drivers.ApplicationBundle) error {
	id := c.Param("id")

	application := getShower(bundle)

	device, err := application.Show(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, device)
}

/**
* ShowDevicesByIPAction shows a device
*
* @param c echo.Context
* @param bundle the application bundle
* @return error
 */
func ShowDevicesByIpAction(c echo.Context, bundle *drivers.ApplicationBundle) error {
	ip := c.Param("ip")

	application := getShower(bundle)

	device, err := application.ShowByIP(ip)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, device)
}

/**
* ShowDevicesByNameAction shows a device
*
* @param c echo.Context
* @param bundle the application bundle
* @return error
 */
func ShowDevicesByNameAction(c echo.Context, bundle *drivers.ApplicationBundle) error {
	name := c.Param("name")

	application := getShower(bundle)

	device, err := application.ShowByName(name)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, device)
}

/**
* Get application use case
*
* @param repository the device repository
* @return app.DeviceShower
 */
func getShower(bundle *drivers.ApplicationBundle) *app.DeviceShower {
	repository := DeviceRepository(bundle.Database.Connection)

	return app.ShowDevice(repository)
}
