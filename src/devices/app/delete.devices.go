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
* Deletes a device
* @param device the device to delete
* @return an error if something goes wrong
**/
func (deleter *DeviceDeleter) Delete(device *model.Device) error {
	return deleter.deviceRepository.Delete(device)
}
