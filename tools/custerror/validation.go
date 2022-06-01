package custerror

import (
	"encoding/json"
	"fmt"
)

// NewValidation creates a validation error for a field
func NewValidation() Validation {
	return &errValidation{
		Messages: []errField{},
	}
}

// Validation custom errors interface
type Validation interface {
	Add(path string, message string) Validation
	Size() int
	Error() string
}

// errField defines the path and error message
type errField struct {
	Path    string `json:"path"`
	Message string `json:"message"`
}

// ErrValidation parameter validation error
type errValidation struct {
	Messages []errField `json:"messages"`
}

func (e *errValidation) Error() string {
	body, err := json.Marshal(e)
	if err != nil {
		return fmt.Sprintf("ErrValidation that cant be marshaled")
	}
	return fmt.Sprintf(string(body))
}

// Add agrega errores a un validation error
func (e *errValidation) Add(path string, message string) Validation {
	err := errField{
		Path:    path,
		Message: message,
	}
	e.Messages = append(e.Messages, err)
	return e
}

// Size devuelve la cantidad de errores
func (e *errValidation) Size() int {
	return len(e.Messages)
}
