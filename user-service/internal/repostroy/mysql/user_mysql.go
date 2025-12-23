package mysql

import (
	"database/sql"
	"errors"
	domain "user-service/internal/domen"
)

type userRepository struct {
	db *sql.DB
}


func NewUserRepository(db *sql.DB) domain.UserRepository {
	return &userRepository{db: db}
}


func (r *userRepository) GetByTelegramID(telegramID int64) (*domain.User, error) {
	query := `
		SELECT id, telegram_id, role, first_name
		FROM users
		WHERE telegram_id = ?
		LIMIT 1
	`

	row := r.db.QueryRow(query, telegramID)

	var u domain.User
	err := row.Scan(&u.ID, &u.TelegramID, &u.Role, &u.FirstName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &u, nil
}


func (r *userRepository) Create(user *domain.User) (*domain.User, error) {
	query := `
		INSERT INTO users (telegram_id, role, first_name)
		VALUES (?, ?, ?)
	`

	result, err := r.db.Exec(query, user.TelegramID, user.Role, user.FirstName)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	user.ID = id
	return user, nil
}

func (r *userRepository) DeleteByTelegramID(telegramID int64) (bool, error) {
	
	user, err := r.GetByTelegramID(telegramID)
	if err != nil {
		return false, err
	}
	if user == nil {
	
		return false, nil
	}

	
	query := `
		DELETE FROM users
		WHERE telegram_id = ?
	`
	result, err := r.db.Exec(query, telegramID)
	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	
	return rowsAffected > 0, nil
}
