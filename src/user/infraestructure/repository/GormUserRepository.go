package repository

import (
	"github.com/r0x16/GroundForce/src/user/domain/model"
	i "github.com/r0x16/GroundForce/src/user/domain/repository"
)

type GormUserRepository struct {
}

var _ i.UserRepository = &GormUserRepository{}

/**
 * Search all the users in the database
 */
func (*GormUserRepository) FindAll() ([]*model.User, error) {
	panic("unimplemented")
}

/**
 * Search a user by its email
 */
func (*GormUserRepository) FindByEmail(email string) (*model.User, error) {
	panic("unimplemented")
}

/**
 * Search a user by its id
 */
func (*GormUserRepository) FindById(id uint) (*model.User, error) {
	panic("unimplemented")
}

/**
 * Search a user by its username
 */
func (*GormUserRepository) FindByUsername(username string) (*model.User, error) {
	panic("unimplemented")
}

/**
 * Search a user by its uuid
 */
func (*GormUserRepository) FindByUuid(uuid string) (*model.User, error) {
	panic("unimplemented")
}

/**
 * Store a user in the database
 */
func (*GormUserRepository) Store(user *model.User) error {
	panic("unimplemented")
}
