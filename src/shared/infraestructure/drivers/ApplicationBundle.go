package drivers

import (
	"github.com/labstack/echo/v4"
	"github.com/r0x16/GroundForce/src/shared/infraestructure/drivers/db"
)

type ApplicationBundle struct {
	Server   *echo.Echo
	Database *db.GormMysqlDatabaseProvider
}
