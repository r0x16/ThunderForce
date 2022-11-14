package action

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

/**
 * UpdateDevicesAction updates a device
 *
 * @param c
 * @return error
 */
func UpdateDevicesAction(c echo.Context) error {
	return c.String(http.StatusOK, "Update device")
}
