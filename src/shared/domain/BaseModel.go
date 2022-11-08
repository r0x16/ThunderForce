package domain

import (
	"time"

	"gorm.io/gorm"
)

/**
 * This model contains the base data for all models
 * @param ID the numeric incremental id of the model
 * @param CreatedAt the date when the model was created
 * @param UpdatedAt the date when the model was updated
 * @param DeletedAt the date when the model was deleted
 */
type BaseModel struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
