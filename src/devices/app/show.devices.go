package app

import (
	"github.com/r0x16/ThunderForce/src/devices/domain/repository"
	"github.com/r0x16/ThunderForce/src/shared/domain/model"
)

/**
* Application service for showing devices
**/
type DeviceShower struct {
	deviceRepository repository.DeviceRepository
}

/**
* Creates a DeviceShower use case with the given DeviceRepository
* @param devices the DeviceRepository to use
* @return a DeviceShower instance
**/
func ShowDevice(devices repository.DeviceRepository) *DeviceShower {
	return &DeviceShower{
		deviceRepository: devices,
	}
}

/**
* Shows a device based on their id (uuid)
* @param id the id of the device
* @return the device
**/
func (shower *DeviceShower) Show(id string) (*model.Device, error) {
	device, err := shower.deviceRepository.FindById(id)
	if err != nil {
		return nil, err
	}

	return device, nil
}

/**
* Shows a device based on their IP Address
* @param ip the IP Address of the device
* @return the device
**/
func (shower *DeviceShower) ShowByIP(ip string) (*model.Device, error) {
	device, err := shower.deviceRepository.FindByIP(ip)
	if err != nil {
		return nil, err
	}

	return device, nil
}

/**
* Shows a device based on their name
* @param name the name of the device
* @return the device
**/
func (shower *DeviceShower) ShowByName(name string) (*model.Device, error) {
	device, err := shower.deviceRepository.FindByName(name)
	if err != nil {
		return nil, err
	}

	return device, nil
}
