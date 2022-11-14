package action

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

/**
 * StoreDevicesAction stores a device
 *
 * @param c
 * @return error
 */
func StoreDevicesAction(c echo.Context) error {
	return c.String(http.StatusOK, "Create device")
}
