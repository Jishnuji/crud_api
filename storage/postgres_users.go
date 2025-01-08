package storage

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
)

type PostrgresUser struct {
	DB *sql.DB
}

func NewPostrgresUser(db *sql.DB) *PostrgresUser {
	return &PostrgresUser{DB: db}
}

func (storage *PostrgresUser) CreateUser(user User) (User, error) {
	user.ID = uuid.New()
	user.Created = time.Now()

	query := `INSERT INTO users (id, firstname, lastname, email, age, created) 
              VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := storage.DB.Exec(query, user.ID, user.Firstname, user.Lastname, user.Email, user.Age, user.Created)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (storage *PostrgresUser) GetUserByID(id uuid.UUID) (User, error) {
	query := `SELECT id, firstname, lastname, email, age, created FROM users WHERE id = $1`
	var user User
	err := storage.DB.QueryRow(query, id).Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Email, &user.Age, &user.Created)
	if err != nil {
		if err == sql.ErrNoRows {
			return User{}, errors.New("user not found")
		}
		return User{}, err
	}
	return user, nil
}

func (storage *PostrgresUser) UpdateUser(id uuid.UUID, user User) (User, error) {
	query := `UPDATE users SET firstname = $1, lastname = $2, email = $3, age = $4 WHERE id = $5`
	_, err := storage.DB.Exec(query, user.Firstname, user.Lastname, user.Email, user.Age, id)
	if err != nil {
		return User{}, err
	}
	user.ID = id
	return user, nil
}
