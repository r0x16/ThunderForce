package action

import (
	"github.com/r0x16/ThunderForce/src/devices/infraestructure/repository"
	"gorm.io/gorm"
)

/**
 * Repository singleton instance
 */
var deviceRepository *repository.GormDeviceRepository = nil

/**
 * This method initializes the repository Singleton pattern instance
 * @param db the database connection
 * @return *GormDeviceRepository the repository
**/
func DeviceRepository(db *gorm.DB) *repository.GormDeviceRepository {
	if deviceRepository == nil {
		deviceRepository = repository.NewGormDeviceRepository(db)
	}
	return deviceRepository
}
