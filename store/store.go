package store

import (
	"database/sql"
	"log"
	"user-service/errors"
	"user-service/models"
)

type store struct {
	db *sql.DB
}

func New(db *sql.DB) store {
	return store{db: db}
}

func (s *store) CreateUser(user *models.User) error {
	query := `INSERT INTO users (email,password,name,address,latitude,longitude,phone_number) VALUES ($1, $2, $3,$4,$5,$6,$7)`
	_, err := s.db.Exec(query, user.Email, user.Password, user.Name, user.Address, user.Latitude, user.Longitude, user.PhoneNumber)
	if err != nil {
		log.Printf("Failed to create user: %v", err)
		return &errors.InternalServerError{Message: "Query Execution Failed"}
	}
	return nil
}

func (s *store) GetUserByEmail(email string) (*models.User, error) {
	query := `SELECT id, email,password,name,address,latitude,longitude,phone_number FROM users WHERE email = $1`
	row := s.db.QueryRow(query, email)

	user := &models.User{}
	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.Name, &user.Address, &user.Latitude, &user.Longitude,
		&user.PhoneNumber)
	if err != nil {
		log.Printf("Failed to get user by username: %v", err)
		return nil, &errors.InternalServerError{Message: "Query Execution Failed"}
	}
	return user, nil
}
