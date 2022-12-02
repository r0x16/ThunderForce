package model

import (
	"github.com/r0x16/ThunderForce/src/shared/domain"
)

/**
 * This model represents a device in the application
 * @param BaseModel the base model for all models
 * @param Uuid the uuid of the device
 * @param Name the name of the device
 * @param Description the description of the device
 * @param Type the type of the device
 * @param Status the status of the device
 * @param Credentials the credentials of the device
 *
**/
type Device struct {
	domain.BaseModel
	Name        string       `gorm:"uniqueIndex;not null" json:"name"`
	Description string       `json:"description"`
	Type        string       `gorm:"not null" json:"type"`
	IP          string       `gorm:"uniqueIndex;not null" json:"ip"`
	Enabled     bool         `gorm:"not null;default:true" json:"enabled"`
	Status      string       `gorm:"not null;default:'stateless'" json:"status"`
	Credentials []Credential `json:"credentials"`
	Blueprints  []Blueprint  `gorm:"many2many:devices_blueprints" json:"blueprints"`
}
