package app

import (
	"github.com/r0x16/ThunderForce/src/devices/domain/model"
	"github.com/r0x16/ThunderForce/src/devices/domain/repository"
)

/**
* Application service for updating devices
**/
type DeviceUpdater struct {
	deviceRepository repository.DeviceRepository
}

/**
* Creates a DeviceUpdater use case with the given DeviceRepository
* @param devices the DeviceRepository to use
* @return a DeviceUpdater instance
**/
func CreateDeviceUpdater(deviceRepository repository.DeviceRepository) *DeviceUpdater {
	return &DeviceUpdater{
		deviceRepository: deviceRepository,
	}
}

/**
* Updates a device
* @param device the device to update
* @return an error if something goes wrong
**/
func (updater *DeviceUpdater) Update(device *model.Device) error {
	return updater.deviceRepository.Update(device)
}
