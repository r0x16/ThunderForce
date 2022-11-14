package action

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

/**
 * ListDevicesAction lists all devices
 *
 * @param c
 * @return error
 */
func ListDevicesAction(c echo.Context) error {
	return c.String(http.StatusOK, "Devices Module")
}
