package app

import "github.com/r0x16/ThunderForce/src/blueprints/domain/repository"

/**
* Application service for showing blueprints
**/
type BlueprintShower struct {
	repository repository.BlueprintRepository
}

/**
* Creates a BlueprintShower use case with the given BlueprintRepository
* @param blueprints the BlueprintRepository to use
* @return a BlueprintShower instance
**/
func ShowBlueprint(blueprints repository.BlueprintRepository) *BlueprintShower {
	return &BlueprintShower{
		repository: blueprints,
	}
}

/**
* Shows a blueprint based on their id (uuid)
* @param id the id of the blueprint
* @return the blueprint
**/
func (shower *BlueprintShower) Show(id string) (interface{}, error) {
	blueprint, err := shower.repository.FindById(id)
	if err != nil {
		return nil, err
	}

	return blueprint, nil
}
