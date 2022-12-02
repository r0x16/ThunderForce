package repository

import (
	"github.com/r0x16/ThunderForce/src/devices/domain/repository"
	"github.com/r0x16/ThunderForce/src/shared/domain"
	"github.com/r0x16/ThunderForce/src/shared/domain/model"
	"gorm.io/gorm"
)

/**
* Object that implements the DeviceRepository interface
* persisting the data in a Gorm database
* @param db the GORM database connection
**/
type GormDeviceRepository struct {
	db *gorm.DB
}

// GormDeviceRepository implements the DeviceRepository interface
var _ repository.DeviceRepository = &GormDeviceRepository{}

/**
 * This method initializes the repository
 * @param db the database connection
 * @return *GormDeviceRepository the repository
**/
func NewGormDeviceRepository(db *gorm.DB) *GormDeviceRepository {
	repo := &GormDeviceRepository{
		db: db,
	}
	repo.db.AutoMigrate(&model.Device{})
	return repo
}

/**
 * This method delete a device in the database
 * @param device the device to save
 * @return error the error if any
**/
func (r *GormDeviceRepository) Delete(device *model.Device) error {
	result := r.db.Delete(&device)
	return result.Error
}

/**
 * This method list all devices in the database
 * @return []*model.Device the list of devices
 * @return error the error if any
**/
func (r *GormDeviceRepository) FindAll() ([]*model.Device, error) {
	devices := []*model.Device{}
	result := r.db.Find(&devices)
	return devices, result.Error
}

/**
 * This method find a device by their IP Address in the database
 * @param IP Address the id of the device
 * @return *model.Device the device
 * @return error the error if any
**/
func (r *GormDeviceRepository) FindByIP(ip string) (*model.Device, error) {
	device := model.Device{}
	result := r.db.Where("ip = ?", ip).Limit(1).Find(&device)
	return &device, result.Error
}

/**
 * This method find a device by id in the database
 * @param name the name of the device
 * @return *model.Device the device
 * @return error the error if any
**/
func (r *GormDeviceRepository) FindById(id string) (*model.Device, error) {
	device := model.Device{}
	result := r.db.Where("id = ?", id).Limit(1).Find(&device)
	return &device, result.Error
}

/**
 * This method find a device by their name in the database
 * @param name the name of the device
 * @return *model.Device the device
 * @return error the error if any
**/
func (r *GormDeviceRepository) FindByName(name string) (*model.Device, error) {
	device := model.Device{}
	result := r.db.Where("name = ?", name).Limit(1).Find(&device)
	return &device, result.Error
}

/**
 * This method find devices by given blueprint
 * @param blueprint the blueprint to search
 * @return []*model.Device the list of devices
 * @return error the error if any
**/
func (*GormDeviceRepository) FindByBlueprint(blueprint *model.Blueprint) ([]*model.Device, error) {
	panic("unimplemented")
}

/**
 * This method find devices based on a filter in the database
 * @param filter the filter to apply
 * @return []*model.Device the list of devices
 * @return error the error if any
**/
func (r *GormDeviceRepository) FindFiltered(filter domain.RepositoryFilter) ([]*model.Device, error) {
	devices := []*model.Device{}
	result := r.db.Where(filter.GetFilter(), filter.GetData()...).Find(&devices)
	return devices, result.Error
}

/**
 * This method stores a device in the database
 * @param device the device to save
 * @return error the error if any
**/
func (r *GormDeviceRepository) Store(device *model.Device) error {
	result := r.db.Create(&device)
	return result.Error
}

/**
 * This method updates a device in the database
 * @param device the device to update
 * @return error the error if any
**/
func (r *GormDeviceRepository) Update(device *model.Device) error {
	result := r.db.Save(&device)
	return result.Error
}
