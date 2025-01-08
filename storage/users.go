package storage

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Email     string    `json:"email"`
	Age       uint      `json:"age"`
	Created   time.Time `json:"created"`
}

type UserStorage interface {
	CreateUser(user User) (User, error)
	GetUserByID(id uuid.UUID) (User, error)
	UpdateUser(id uuid.UUID, user User) (User, error)
}
