package domain

import (
	"time"

	"gorm.io/gorm"
)

/**
 * This model contains the base data for all models
 * @param ID the numeric incremental id of the model
 * @param Uuid the uuid of the model
 * @param CreatedAt the date when the model was created
 * @param UpdatedAt the date when the model was updated
 * @param DeletedAt the date when the model was deleted
 *
 * UUID identification over numeric incremental id is used to prevent
 * the user from knowing the number of elements in the table
 *
 */
type BaseModel struct {
	ID        string `gorm:"primaryKey;type:uuid;default:UUID()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
