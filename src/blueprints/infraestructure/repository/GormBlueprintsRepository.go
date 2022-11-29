package repository

import (
	"github.com/r0x16/ThunderForce/src/blueprints/domain/model"
	"github.com/r0x16/ThunderForce/src/blueprints/domain/repository"
	ext1 "github.com/r0x16/ThunderForce/src/devices/domain/model"
	"github.com/r0x16/ThunderForce/src/shared/domain"
	"gorm.io/gorm"
)

type GormBlueprintRepository struct {
	Db *gorm.DB
}

var _ repository.BlueprintRepository = &GormBlueprintRepository{}

func NewGormBlueprintRepository(db *gorm.DB) *GormBlueprintRepository {
	repo := &GormBlueprintRepository{Db: db}
	repo.Db.AutoMigrate(&model.Blueprint{})
	return repo
}

// Delete implements repository.BlueprintRepository
func (*GormBlueprintRepository) Delete(blueprint *model.Blueprint) error {
	panic("unimplemented")
}

// FindAll implements repository.BlueprintRepository
func (*GormBlueprintRepository) FindAll() ([]*model.Blueprint, error) {
	panic("unimplemented")
}

// FindByDevice implements repository.BlueprintRepository
func (*GormBlueprintRepository) FindByDevice(device *ext1.Device) ([]*model.Blueprint, error) {
	panic("unimplemented")
}

// FindById implements repository.BlueprintRepository
func (*GormBlueprintRepository) FindById(id string) (*model.Blueprint, error) {
	panic("unimplemented")
}

// FindByName implements repository.BlueprintRepository
func (*GormBlueprintRepository) FindByName(name string) (*model.Blueprint, error) {
	panic("unimplemented")
}

// FindByService implements repository.BlueprintRepository
func (*GormBlueprintRepository) FindByService(service *model.Service) ([]*model.Blueprint, error) {
	panic("unimplemented")
}

// FindFiltered implements repository.BlueprintRepository
func (*GormBlueprintRepository) FindFiltered(filter domain.RepositoryFilter) ([]*model.Blueprint, error) {
	panic("unimplemented")
}

// Store implements repository.BlueprintRepository
func (*GormBlueprintRepository) Store(blueprint *model.Blueprint) error {
	panic("unimplemented")
}

// Update implements repository.BlueprintRepository
func (*GormBlueprintRepository) Update(blueprint *model.Blueprint) error {
	panic("unimplemented")
}
