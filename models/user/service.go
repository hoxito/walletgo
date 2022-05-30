package user

import (
	"fmt"
)

type UserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" `
	Password string `json:"password" validate:"required"`
}

// newOS Creates a new user
func CreateUser(user *UserRequest) (*User, error) {

	if err := user.ValidateSchema(); err != nil {
		return nil, err
	}
	newUser := NewUser()
	newUser.Name = user.Name
	newUser.Email = user.Email
	newUser.Password = user.Password

	fmt.Println("User Created")
	newUser, err := insert(newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (e *User) SetState(state string) {
	//TODO
}

// Get wrapper
func Get(UserId string) (*User, error) {
	return findByID(UserId)
}

//  wrapper

func Users() ([]*User, error) {
	return findAll()
}
