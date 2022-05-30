package custerror

import (
	"fmt"
	"strconv"
)

// - Some common errors-

// Unauthorized
var Unauthorized = NewCustom(401, "Unauthorized")

// NotFound
var NotFound = NewCustom(400, "Entity not found")

// Internal
var Internal = NewCustom(500, "Internal server error")

// NewCustom creates a new errCustom
func NewCustom(status int, message string) Custom {
	return &errCustom{
		status:  status,
		Message: message,
	}
}

//  - Algunas definiciones necesarias -

//Errors Interface to cast errors that contains a status code and an error message
type Custom interface {
	Status() int
	Error() string
}

// errCustom personalized http error
type errCustom struct {
	status  int
	Message string `json:"error"`
}

//Returns the error message
func (e *errCustom) Error() string {
	return fmt.Sprintf(e.Message)
}

// Returns http status code
func (e *errCustom) Status() int {
	return e.status
}

// Returns http status code and error message
func (e *errCustom) String() string {
	return strconv.Itoa(e.Status()) + e.Error()
}
