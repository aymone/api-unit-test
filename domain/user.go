package domain

import "errors"

type User struct {
	ID string `json:"id" bson:"_id"`
}

var (
	ErrEmptyID       error = errors.New("empty id")
	ErrZeroID        error = errors.New("zero id")
	ErrDuplicatedKey error = errors.New("duplicated key error")
)

func (u User) Validate() error {
	switch u.ID {
	case "":
		return ErrEmptyID
	case "0":
		return ErrZeroID
	}

	return nil
}
