package repository

import (
	"github.com/r0x16/ThunderForce/src/shared/domain"
	"github.com/r0x16/ThunderForce/src/shared/domain/model"
)

/**
 * This interface represents the blueprint repository
 * @param FindAllBlueprints returns all blueprints
 * @param FindBlueprintById returns a blueprint by id
 * @param FindBlueprintByName returns a blueprint by name
 * @param FindBlueprintsByDevice returns all blueprints by device
 * @param SaveBlueprint saves a blueprint
 * @param DeleteBlueprint deletes a blueprint
 */
type BlueprintRepository interface {
	FindAll() ([]*model.Blueprint, error)
	FindFiltered(filter domain.RepositoryFilter) ([]*model.Blueprint, error)
	FindById(id string) (*model.Blueprint, error)
	FindByName(name string) (*model.Blueprint, error)
	FindByService(service *model.Service) ([]*model.Blueprint, error)
	FindByDevice(device *model.Device) ([]*model.Blueprint, error)
	Store(blueprint *model.Blueprint) error
	Update(blueprint *model.Blueprint) error
	Delete(blueprint *model.Blueprint) error
}
