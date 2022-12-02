package model

import (
	"github.com/r0x16/ThunderForce/src/shared/domain"
)

/**
 * This model represents a blueprint in the application
 * @param BaseModel the base model for all models
 * @param Name the name of the blueprint
 * @param Description the description of the blueprint
 * @param Services the services associated with the blueprint
 *
**/
type Blueprint struct {
	domain.BaseModel
	Name        string    `gorm:"uniqueIndex;not null" json:"name"`
	Description string    `json:"description"`
	Services    []Service `gorm:"many2many:services_blueprints" json:"services"`
	Devices     []Device  `gorm:"many2many:devices_blueprints" json:"devices"`
}
