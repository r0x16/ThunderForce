package app

import (
	"github.com/r0x16/ThunderForce/src/blueprints/domain/repository"
	"github.com/r0x16/ThunderForce/src/shared/domain/model"
)

/**
* Application service for updating blueprints
**/
type BlueprintUpdater struct {
	repository repository.BlueprintRepository
}

/**
* Blueprint data transfer object
**/
type BlueprintUpdateData struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

/**
* Creates a BlueprintUpdater use case with the given BlueprintRepository
* @param blueprints the BlueprintRepository to use
* @return a BlueprintUpdater instance
**/
func CreateBlueprintUpdater(blueprints repository.BlueprintRepository) *BlueprintUpdater {
	return &BlueprintUpdater{
		repository: blueprints,
	}
}

/**
* Get blueprint to update
* @param id the blueprint id
* @return the blueprint
* @return an error if something goes wrong
**/
func (updater *BlueprintUpdater) getBlueprintToUpdate(id string) (*model.Blueprint, error) {
	blueprint, err := updater.repository.FindById(id)
	if err != nil {
		return nil, err
	}

	return blueprint, nil
}

/**
* Updates a blueprint
* @param blueprint the blueprint to update
* @return an error if something goes wrong
**/
func (updater *BlueprintUpdater) Update(id string, data *BlueprintUpdateData) (*model.Blueprint, error) {
	blueprint, err := updater.getBlueprintToUpdate(id)
	if err != nil {
		return nil, err
	}

	blueprint.Name = data.Name
	blueprint.Description = data.Description

	err = updater.repository.Update(blueprint)

	return blueprint, err

}
