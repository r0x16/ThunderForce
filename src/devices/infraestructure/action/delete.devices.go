package action

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

/**
 * DeleteDevicesAction deletes a device
 *
 * @param c
 * @return error
 */
func DeleteDevicesAction(c echo.Context) error {
	return c.String(http.StatusOK, "Delete device")
}
