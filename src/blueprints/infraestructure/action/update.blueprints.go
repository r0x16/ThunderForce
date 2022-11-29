package action

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/r0x16/ThunderForce/src/blueprints/app"
	"github.com/r0x16/ThunderForce/src/shared/infraestructure/drivers"
)

/**
 * This method updates a blueprint in the database
 * @param c the echo context
 * @param bundle the application bundle
 * @return error the error if any
**/
func UpdateBlueprintsAction(c echo.Context, bundle *drivers.ApplicationBundle) error {
	id := c.Param("id")

	application := getUpdater(bundle)

	blueprintData := &app.BlueprintUpdateData{}
	if err := c.Bind(blueprintData); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	blueprint, err := application.Update(id, blueprintData)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, blueprint)

}

/**
 * This method gets the application use case
 * @param bundle the application bundle
 * @return *app.BlueprintUpdater the application use case
**/
func getUpdater(bundle *drivers.ApplicationBundle) *app.BlueprintUpdater {
	repository := BlueprintRepository(bundle.Database.Connection)

	return app.CreateBlueprintUpdater(repository)
}
