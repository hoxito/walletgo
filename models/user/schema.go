package user

import (
	"database/sql"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// User data structure
type User struct {
	UserId   string       `json:"userId"  validate:"required,max=50"`
	Name     string       `json:"name" validate:"required,max=50"`
	Email    string       `json:"email" `
	Password string       `json:"password" validate:"required,max=50"`
	Deleted  sql.NullTime `json:"deleted" `
	Created  sql.NullTime `json:"created" validate:"required"`
}

// User data structure
type Login struct {
	Name     string `json:"name" validate:"required,max=50"`
	Password string `json:"password" validate:"required,max=50"`
}

// NewUser Nueva instancia de usuario
func NewUser() *User {
	var now sql.NullTime
	now.Time = time.Now()
	now.Valid = true
	return &User{
		UserId:  uuid.NewString(),
		Created: now,
	}
}

// SetPassword (encripts it) TODO
func (e *User) SetPassword(plainPwd string) error {

	return nil
}

// ValidatePassword  TODO
func (e *User) ValidatePassword(plainPwd string) error {

	return nil
}

// ValidateSchema validates user structure
func (e *User) ValidateSchema() error {
	return validator.New().Struct(e)
}

func (e *UserRequest) ValidateSchema() error {
	return validator.New().Struct(e)
}
