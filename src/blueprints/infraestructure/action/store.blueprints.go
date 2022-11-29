package action

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/r0x16/ThunderForce/src/blueprints/app"
	"github.com/r0x16/ThunderForce/src/shared/infraestructure/drivers"
)

/**
 * This method stores a blueprint in the database
 * @param c the echo context
 * @param bundle the application bundle
 * @return error the error if any
**/
func StoreBlueprintsAction(c echo.Context, bundle *drivers.ApplicationBundle) error {
	application := getPersister(bundle)

	blueprintData := &app.BlueprintData{}
	if err := c.Bind(blueprintData); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	blueprint, err := application.Persist(blueprintData)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, blueprint)

}

/**
 * This method gets the application use case
 * @param bundle the application bundle
 * @return *app.BlueprintPersister the application use case
**/
func getPersister(bundle *drivers.ApplicationBundle) *app.BlueprintPersister {
	repository := BlueprintRepository(bundle.Database.Connection)

	return app.CreateBlueprint(repository)
}
