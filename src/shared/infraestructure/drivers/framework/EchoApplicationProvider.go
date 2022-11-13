package framework

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/r0x16/ThunderForce/src/shared/domain"
	"github.com/r0x16/ThunderForce/src/shared/infraestructure/drivers"
	"github.com/r0x16/ThunderForce/src/shared/infraestructure/module"

	devices "github.com/r0x16/ThunderForce/src/devices/infraestructure/module"
)

type EchoApplicationProvider struct {
	Bundle *drivers.ApplicationBundle
}

var _ domain.ApplicationProvider = &EchoApplicationProvider{}

// Creates a new Echo server to serve http requests and response
func (app *EchoApplicationProvider) Boot() {
	server := echo.New()
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())

	app.Bundle.Server = server
}

// Provides the list of Echo modules to bootstrap all the routes
func (app *EchoApplicationProvider) ProvideModules() []domain.ApplicationModule {
	return []domain.ApplicationModule{
		&module.MainModule{Bundle: app.Bundle},
		&devices.DevicesModule{Bundle: app.Bundle},
	}
}

// Runs the HTTP server in the especified port and listens to errors
func (app *EchoApplicationProvider) Run() error {
	err := app.Bundle.Server.Start(":" + os.Getenv("GROUND_PORT"))
	// Start server
	app.Bundle.Server.Logger.Fatal(
		err,
	)

	return err
}
