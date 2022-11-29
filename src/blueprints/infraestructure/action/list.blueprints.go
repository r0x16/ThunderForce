package action

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/r0x16/ThunderForce/src/blueprints/app"
	"github.com/r0x16/ThunderForce/src/shared/infraestructure/drivers"
)

/**
 * This method lists all the blueprints in the database
 * @param c the echo context
 * @param bundle the application bundle
 * @return error the error if any
**/
func ListBlueprintsAction(c echo.Context, bundle *drivers.ApplicationBundle) error {
	application := getLister(bundle)

	blueprints, err := application.List()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, blueprints)
}

/**
 * This method gets the application use case
 * @param bundle the application bundle
 * @return *app.BlueprintLister the application use case
**/
func getLister(bundle *drivers.ApplicationBundle) *app.BlueprintLister {
	repository := BlueprintRepository(bundle.Database.Connection)

	return app.ListBlueprints(repository)
}
