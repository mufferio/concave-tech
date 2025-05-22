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
