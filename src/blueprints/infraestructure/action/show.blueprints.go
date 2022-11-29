package action

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/r0x16/ThunderForce/src/blueprints/app"
	"github.com/r0x16/ThunderForce/src/shared/infraestructure/drivers"
)

/**
 * This method shows a blueprint in the database
 * @param c the echo context
 * @param bundle the application bundle
 * @return error the error if any
**/
func ShowBlueprintsAction(c echo.Context, bundle *drivers.ApplicationBundle) error {
	id := c.Param("id")

	application := getShower(bundle)

	blueprint, err := application.Show(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, blueprint)
}

/**
 * This method gets the application use case
 * @param bundle the application bundle
 * @return *app.BlueprintShower the application use case
**/
func getShower(bundle *drivers.ApplicationBundle) *app.BlueprintShower {
	repo := BlueprintRepository(bundle.Database.Connection)

	return app.ShowBlueprint(repo)
}
