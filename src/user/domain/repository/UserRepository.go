package repository

import "github.com/r0x16/GroundForce/src/user/domain/model"

/**
 * This interface defines the methods that a user repository must implement
 * @param FindAll returns all the users in the persistence layer
 * @param FindById returns a user by its numeric incremental id
 * @param FindByUuid returns a user by its uuid
 * @param FindByUsername returns a user by its username
 * @param FindByEmail returns a user by its email
 * @param Store stores a user in the persistence layer
 */
type UserRepository interface {
	FindAll() ([]*model.User, error)
	FindById(id uint) (*model.User, error)
	FindByUuid(uuid string) (*model.User, error)
	FindByUsername(username string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	Store(user *model.User) error
}
