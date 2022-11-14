package action

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

/**
 * ShowDevicesAction shows a device
 *
 * @param c
 * @return error
 */
func ShowDevicesAction(c echo.Context) error {
	return c.String(http.StatusOK, "Show Device")
}
