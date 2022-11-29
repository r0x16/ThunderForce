package model

import "github.com/r0x16/ThunderForce/src/shared/domain"

/**
 * This model represents a service in the application
 * @param BaseModel the base model for all models
 * @param Code the name of the service
 * @param Description the description of the service
 * @param Blueprints the blueprints associated with the service
 */
type Service struct {
	domain.BaseModel
	Code        string      `gorm:"uniqueIndex;not null" json:"name"`
	Description string      `json:"description"`
	Blueprints  []Blueprint `gorm:"many2many:services_blueprints" json:"blueprints"`
}
