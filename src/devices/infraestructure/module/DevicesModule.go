package module

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/r0x16/ThunderForce/src/shared/domain"
	"github.com/r0x16/ThunderForce/src/shared/infraestructure/drivers"
)

type DevicesModule struct {
	Bundle *drivers.ApplicationBundle
}

var _ domain.ApplicationModule = &DevicesModule{}

func (m *DevicesModule) Setup() {
	m.Bundle.Server.GET("/devices", func(c echo.Context) error {
		return c.String(http.StatusOK, "Devices Module")
	})
}
