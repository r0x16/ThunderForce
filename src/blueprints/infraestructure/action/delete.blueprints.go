package action

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/r0x16/ThunderForce/src/blueprints/app"
	"github.com/r0x16/ThunderForce/src/shared/infraestructure/drivers"
)

/**
 * This method delete a blueprint in the database
 * @param c the echo context
 * @param bundle the application bundle
 * @return error the error if any
**/
func DeleteBlueprintsAction(c echo.Context, bundle *drivers.ApplicationBundle) error {
	id := c.Param("id")

	application := getDeleter(bundle)

	err := application.Delete(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, nil)
}

/**
 * This method gets the application use case
 * @param bundle the application bundle
 * @return *app.BlueprintDeleter the application use case
**/
func getDeleter(bundle *drivers.ApplicationBundle) *app.BlueprintDeleter {
	repository := BlueprintRepository(bundle.Database.Connection)

	return app.DeleteBlueprint(repository)
}
