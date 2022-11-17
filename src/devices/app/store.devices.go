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

type DeviceData struct {
	Name        string `json:"name"`
	IP          string `json:"ip"`
	Description string `json:"description"`
	Type        string `json:"type"`
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
func (persister *DevicePersister) Persist(data *DeviceData) (*model.Device, error) {
	device := &model.Device{
		Name:        data.Name,
		IP:          data.IP,
		Description: data.Description,
		Type:        data.Type,
	}

	err := persister.deviceRepository.Store(device)

	return device, err
}
