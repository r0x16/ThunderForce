package app

import (
	"github.com/r0x16/ThunderForce/src/blueprints/domain/model"
	"github.com/r0x16/ThunderForce/src/blueprints/domain/repository"
)

/**
* Application service for deleting blueprints
*
**/
type BlueprintDeleter struct {
	repository repository.BlueprintRepository
}

/**
* Creates a BlueprintDeleter use case with the given BlueprintRepository
* @param repository the BlueprintRepository to use
* @return a BlueprintDeleter instance
*
**/
func DeleteBlueprint(repository repository.BlueprintRepository) *BlueprintDeleter {
	return &BlueprintDeleter{
		repository: repository,
	}
}

/**
* Get blueprint to update
* @param id the id of the blueprint
* @return the blueprint
*
**/
func (deleter *BlueprintDeleter) getBlueprintToDelete(id string) (*model.Blueprint, error) {
	blueprint, err := deleter.repository.FindById(id)
	if err != nil {
		return nil, err
	}

	return blueprint, nil
}

/**
* Deletes a blueprint
* @param id the blueprint to delete
* @return an error if something goes wrong
*
**/
func (deleter *BlueprintDeleter) Delete(id string) error {
	blueprint, err := deleter.getBlueprintToDelete(id)

	if err != nil {
		return err
	}

	return deleter.repository.Delete(blueprint)
}
