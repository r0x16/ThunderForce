package module

import (
	"github.com/r0x16/ThunderForce/src/blueprints/infraestructure/action"
	"github.com/r0x16/ThunderForce/src/shared/domain"
	"github.com/r0x16/ThunderForce/src/shared/infraestructure/drivers"
)

/**
 * BlueprintsModule represents the links between blueprints routes and actions
 */
type BlueprintsModule struct {
	Bundle *drivers.ApplicationBundle
}

var _ domain.ApplicationModule = &BlueprintsModule{}

/**
 * Initializes blueprints module and set up their routes
**/
func (m *BlueprintsModule) Setup() {
	blueprints := m.Bundle.Server.Group("/blueprints")

	blueprints.GET("", m.Bundle.ActionInjection(action.ListBlueprintsAction))
	blueprints.GET("/:id", m.Bundle.ActionInjection(action.ShowBlueprintsAction))
	blueprints.POST("", m.Bundle.ActionInjection(action.StoreBlueprintsAction))
	blueprints.PUT("/:id", m.Bundle.ActionInjection(action.UpdateBlueprintsAction))
	blueprints.DELETE("/:id", m.Bundle.ActionInjection(action.DeleteBlueprintsAction))
}
