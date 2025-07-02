package repositories

import (
	"database/sql"

	"github.com/mufferio/concave-tech/entities"
)

type UserRepository interface {
	FindAll() ([]entities.User, error)
}

type userRepo struct{ db *sql.DB }

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepo{db}
}

func (r *userRepo) FindAll() ([]entities.User, error) {
	rows, err := r.db.Query(`
		SELECT id, firstname, lastname, date_of_birth, income
		FROM "user"
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entities.User
	for rows.Next() {
		var u entities.User
		if err := rows.Scan(&u.ID, &u.Firstname, &u.Lastname, &u.DateOfBirth, &u.Income); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, rows.Err()
}

func (r *userRepo) FindByUsername(username string) (*entities.User, error) {
	const q = `
	  SELECT
	    c.id, c.username, c.password_hash, c.created_at,
	    p.id, p.first_name, p.last_name, p.date_of_birth, p.income
	  FROM credential c
	  JOIN profile   p ON p.credential_id = c.id
	  WHERE c.username = $1
	`

	var u entities.User
	err := r.db.QueryRow(q, username).Scan(
		&u.Credential.ID, &u.Credential.Username, &u.Credential.PasswordHash, &u.Credential.CreatedAt,
		&u.Profile.ID, &u.Profile.FirstName, &u.Profile.LastName, &u.Profile.DateOfBirth, &u.Profile.Income,
	)
	if err == sql.ErrNoRows {
		return nil, nil // not found
	}
	if err != nil {
		return nil, err
	}
	return &u, nil
}
