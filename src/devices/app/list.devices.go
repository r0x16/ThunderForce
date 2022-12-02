package app

import (
	"github.com/r0x16/ThunderForce/src/devices/domain/repository"
	"github.com/r0x16/ThunderForce/src/shared/domain"
	"github.com/r0x16/ThunderForce/src/shared/domain/model"
)

/**
* Application service for listing devices
**/
type DeviceLister struct {
	deviceRepository repository.DeviceRepository
}

/**
* Creates a DeviceLister use case with the given DeviceRepository
* @param devices the DeviceRepository to use
* @return a DeviceLister instance
**/
func ListDevices(devices repository.DeviceRepository) *DeviceLister {
	return &DeviceLister{
		deviceRepository: devices,
	}
}

/**
* Lists all devices
* @return a list of devices
**/
func (lister *DeviceLister) List() ([]*model.Device, error) {
	devices, err := lister.deviceRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return devices, nil
}

/**
* Gets a filterted list of devices
* @param filter the filter to apply
* @return a list of devices
**/
func (lister *DeviceLister) Filter(filter domain.RepositoryFilter) ([]*model.Device, error) {
	devices, err := lister.deviceRepository.FindFiltered(filter)
	if err != nil {
		return nil, err
	}

	return devices, nil
}
