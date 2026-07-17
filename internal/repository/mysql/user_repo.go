package mysql

import (
	"Pair_Project_Clothing_Ecommerce/internal/domain"
	"database/sql"
	"errors"
	"fmt"
)

type userRepo struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) domain.UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) Login(email, password string) (*domain.User, error) {
	var user domain.User

	query := `
		SELECT id, email, role
		FROM users
		WHERE email = ? AND password = ?
	`

	err := r.db.QueryRow(query, email, password).
		Scan(&user.ID, &user.Email, &user.Role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("email atau password salah")
		}
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) GetAll() ([]domain.UserReportDTO, error) {
	query := `
		SELECT
			u.id,
			u.email,
			u.role,
			COALESCE(p.full_name, '-'),
			COALESCE(p.phone, '-')
		FROM users u
		LEFT JOIN profiles p
			ON u.id = p.user_id
		ORDER BY u.id
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []domain.UserReportDTO{}

	for rows.Next() {
		var user domain.UserReportDTO

		if err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.Role,
			&user.FullName,
			&user.Phone,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepo) Create(req domain.CreateUserRequest) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	result, err := tx.Exec(
		`
		INSERT INTO users (email, password, role)
		VALUES (?, ?, ?)
		`,
		req.Email,
		req.Password,
		req.Role,
	)
	if err != nil {
		return fmt.Errorf("gagal membuat user: %w", err)
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	_, err = tx.Exec(
		`
		INSERT INTO profiles (user_id, full_name, phone, address)
		VALUES (?, ?, ?, ?)
		`,
		userID,
		req.FullName,
		req.Phone,
		req.Address,
	)
	if err != nil {
		return fmt.Errorf("gagal membuat profile: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *userRepo) UpdateRole(userID int, role string) error {
	result, err := r.db.Exec(
		`
		UPDATE users
		SET role = ?
		WHERE id = ?
		`,
		role,
		userID,
	)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return errors.New("user tidak ditemukan")
	}

	return nil
}

func (r *userRepo) Delete(userID int) error {
	result, err := r.db.Exec(
		`
		DELETE FROM users
		WHERE id = ?
		`,
		userID,
	)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return errors.New("user tidak ditemukan")
	}

	return nil
}
