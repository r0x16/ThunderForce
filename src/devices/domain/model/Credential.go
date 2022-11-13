package model

import "github.com/r0x16/ThunderForce/src/shared/domain"

/**
 * This model represents a device in the application
 * @param BaseModel the base model for all models
 * @param Type the type of connection used (ssh, telnet, etc)
 * @param Data json data associated to the type of connection
 * @param Device the device that owns this credential
 *
**/
type Credential struct {
	domain.BaseModel
	Type     string `gorm:"not null"`
	Data     string `gorm:"type:json;not null"`
	Device   Device `gorm:"foreignKey:DeviceID;references:ID"`
	DeviceID uint
}
