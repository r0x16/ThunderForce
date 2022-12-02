package app

import (
	"github.com/r0x16/ThunderForce/src/devices/domain/repository"
	"github.com/r0x16/ThunderForce/src/shared/domain/model"
)

/**
* Application service for updating devices
**/
type DeviceUpdater struct {
	deviceRepository repository.DeviceRepository
}

type DeviceUpdateData struct {
	Name        string `json:"name"`
	IP          string `json:"ip"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Enabled     bool   `json:"active"`
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
* Get device to update
* @param id the id of the device
* @return the device
**/
func (updater *DeviceUpdater) getDeviceToUpdate(id string) (*model.Device, error) {
	device, err := updater.deviceRepository.FindById(id)
	if err != nil {
		return nil, err
	}

	return device, nil
}

/**
* Updates a device
* @param device the device to update
* @return an error if something goes wrong
**/
func (updater *DeviceUpdater) Update(id string, data *DeviceUpdateData) (*model.Device, error) {
	device, err := updater.getDeviceToUpdate(id)
	if err != nil {
		return nil, err
	}

	device.Name = data.Name
	device.IP = data.IP
	device.Description = data.Description
	device.Type = data.Type
	device.Enabled = data.Enabled

	err = updater.deviceRepository.Update(device)

	return device, err
}
