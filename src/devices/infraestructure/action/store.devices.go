package action

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/r0x16/ThunderForce/src/shared/infraestructure/drivers"
)

/**
 * StoreDevicesAction stores a device
 *
 * @param c
 * @return error
 */
func StoreDevicesAction(c echo.Context, bundle *drivers.ApplicationBundle) error {
	return c.String(http.StatusOK, "Create device")
}
