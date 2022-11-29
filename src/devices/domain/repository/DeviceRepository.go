package repository

import (
	"github.com/r0x16/ThunderForce/src/devices/domain/model"
	"github.com/r0x16/ThunderForce/src/shared/domain"

	ext1 "github.com/r0x16/ThunderForce/src/blueprints/domain/model"
)

/**
 * This interface represents a repository for devices
 * @param FindAll gets all devices
 * @param FindFiltered gets a filtered list of devices
 * @param FindById gets a device by its id
 * @param FindByName gets a device by its name
 * @param FindByIP gets a device by its IP
 * @param Store persists a device
 * @param Delete deletes a device
**/
type DeviceRepository interface {
	FindAll() ([]*model.Device, error)
	FindFiltered(filter domain.RepositoryFilter) ([]*model.Device, error)
	FindById(id string) (*model.Device, error)
	FindByIP(ip string) (*model.Device, error)
	FindByName(name string) (*model.Device, error)
	FindByBlueprint(blueprint *ext1.Blueprint) ([]*model.Device, error)
	Store(device *model.Device) error
	Update(device *model.Device) error
	Delete(device *model.Device) error
}
