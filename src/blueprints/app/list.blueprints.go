package app

import (
	"github.com/r0x16/ThunderForce/src/blueprints/domain/model"
	"github.com/r0x16/ThunderForce/src/blueprints/domain/repository"
	"github.com/r0x16/ThunderForce/src/shared/domain"
)

/**
* Application service for listing blueprints
**/
type BlueprintLister struct {
	repository repository.BlueprintRepository
}

/**
* Creates a BlueprintLister use case with the given BlueprintRepository
* @param blueprints the BlueprintRepository to use
* @return a BlueprintLister instance
**/
func ListBlueprints(blueprints repository.BlueprintRepository) *BlueprintLister {
	return &BlueprintLister{
		repository: blueprints,
	}
}

/**
* Lists all blueprints
* @return a list of blueprints
**/
func (lister *BlueprintLister) List() ([]*model.Blueprint, error) {
	blueprints, err := lister.repository.FindAll()
	if err != nil {
		return nil, err
	}

	return blueprints, nil
}

/**
* Gets a filterted list of blueprints
* @param filter the filter to apply
* @return a list of blueprints
**/
func (lister *BlueprintLister) Filter(filter domain.RepositoryFilter) ([]*model.Blueprint, error) {
	blueprints, err := lister.repository.FindFiltered(filter)
	if err != nil {
		return nil, err
	}

	return blueprints, nil
}
