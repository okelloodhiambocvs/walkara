package sqlite

import (
	"database/sql"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(id, email, password string) error {
	_, err := r.DB.Exec(`
	INSERT INTO users (id, email, password)
	VALUES (?, ?, ?)
	`, id, email, password)

	return err
}

func (r *UserRepository) GetUserByEmail(email string) (string, string, error) {
	var id, password string

	err := r.DB.QueryRow(`
	SELECT id, password FROM users WHERE email = ?
	`, email).Scan(&id, &password)

	return id, password, err
}