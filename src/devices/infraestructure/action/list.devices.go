package action

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/r0x16/ThunderForce/src/shared/infraestructure/drivers"
)

/**
 * ListDevicesAction lists all devices
 *
 * @param c
 * @return error
 */
func ListDevicesAction(c echo.Context, bundle *drivers.ApplicationBundle) error {
	return c.String(http.StatusOK, "Devices Module")
}
