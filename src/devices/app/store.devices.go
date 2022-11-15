package app

import (
	"github.com/r0x16/ThunderForce/src/devices/domain/model"
	"github.com/r0x16/ThunderForce/src/devices/domain/repository"
)

/**
* Application service for persist devices
**/
type DevicePersister struct {
	deviceRepository repository.DeviceRepository
}

/**
* Creates a DevicePersister use case with the given DeviceRepository
 * @param devices the DeviceRepository to use
 * @return a DevicePersister instance
**/
func CreateDevice(deviceRepository repository.DeviceRepository) *DevicePersister {
	return &DevicePersister{
		deviceRepository: deviceRepository,
	}
}

/**
* Persists a device
* @param device the device to persist
* @return an error if something goes wrong
**/
func (persister *DevicePersister) Persist(device *model.Device) error {
	return persister.deviceRepository.Store(device)
}
