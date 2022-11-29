package app

import (
	"github.com/google/uuid"
	"github.com/r0x16/ThunderForce/src/blueprints/domain/model"
	"github.com/r0x16/ThunderForce/src/blueprints/domain/repository"
)

/**
* Application service for persisting blueprints
**/
type BlueprintPersister struct {
	repository repository.BlueprintRepository
}

/**
* Blueprint data transfer object
**/
type BlueprintData struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

/**
* Creates a BlueprintPersister use case with the given BlueprintRepository
* @param blueprints the BlueprintRepository to use
* @return a BlueprintPersister instance
**/
func CreateBlueprint(blueprints repository.BlueprintRepository) *BlueprintPersister {
	return &BlueprintPersister{
		repository: blueprints,
	}
}

/**
* Persists a blueprint
* @param blueprint the blueprint to persist
* @return an error if something goes wrong
**/
func (persister *BlueprintPersister) Persist(data *BlueprintData) (*model.Blueprint, error) {
	blueprint := &model.Blueprint{
		Name:        data.Name,
		Description: data.Description,
	}
	blueprint.ID = uuid.New()

	err := persister.repository.Store(blueprint)

	return blueprint, err
}
