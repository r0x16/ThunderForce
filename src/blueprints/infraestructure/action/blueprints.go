package action

import (
	"github.com/r0x16/ThunderForce/src/blueprints/infraestructure/repository"
	"gorm.io/gorm"
)

/**
* Repository Singleton instance
 */
var blueprintRepository *repository.GormBlueprintRepository = nil

/**
* This method initializes the repository Singleton pattern instance
* @param db the database connection
* @return *GormBlueprintRepository the repository
**/
func BlueprintRepository(db *gorm.DB) *repository.GormBlueprintRepository {
	if blueprintRepository == nil {
		blueprintRepository = repository.NewGormBlueprintRepository(db)
	}
	return blueprintRepository
}
