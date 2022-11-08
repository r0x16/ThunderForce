package model

import "github.com/r0x16/GroundForce/src/shared/domain"

/**
 * This model represents a user in the application
 * @param BaseModel the base model for all models
 * @param Username the username of the user
 * @param Email the email of the user
 * @param Password the password of the user
 * @param Uuid the uuid of the user
 *
 * UUID identification over numeric incremental id is used to prevent
 * the user from knowing the number of users in the database
 *
 */
type User struct {
	domain.BaseModel
	Uuid        string       `gorm:"uniqueIndex"`
	Username    string       `gorm:"uniqueIndex"`
	Credentials []Credential `json:"ignored"`
	Email       string       `gorm:"uniqueIndex"`
}

/**
 * This model represents a user's credential
 * @param BaseModel the base model for all models
 * @param Mode type of credentials, e.g. password, token, etc.
 * @param Data the value of the credential
 * @param User the user that owns this credential
 */
type Credential struct {
	domain.BaseModel
	Mode   string
	Data   string
	UserID uint
}
