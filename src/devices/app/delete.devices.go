package app

import (
	"github.com/r0x16/ThunderForce/src/devices/domain/model"
	"github.com/r0x16/ThunderForce/src/devices/domain/repository"
)

/**
* Application service for deleting devices
**/
type DeviceDeleter struct {
	deviceRepository repository.DeviceRepository
}

/**
* Creates a DeviceDeleter use case with the given DeviceRepository
* @param devices the DeviceRepository to use
* @return a DeviceDeleter instance
**/
func DeleteDevice(deviceRepository repository.DeviceRepository) *DeviceDeleter {
	return &DeviceDeleter{
		deviceRepository: deviceRepository,
	}
}

/**
* Get device to update
* @param id the id of the device
* @return the device
**/
func (deleter *DeviceDeleter) getDeviceToDelete(id string) (*model.Device, error) {
	device, err := deleter.deviceRepository.FindById(id)
	if err != nil {
		return nil, err
	}

	return device, nil
}

/**
* Deletes a device
* @param device the device to delete
* @return an error if something goes wrong
**/
func (deleter *DeviceDeleter) Delete(id string) error {
	device, err := deleter.getDeviceToDelete(id)

	if err != nil {
		return err
	}

	return deleter.deviceRepository.Delete(device)
}
