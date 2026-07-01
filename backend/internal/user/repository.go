package user

import (
	"context"

	md "backend/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetByID(id string) (*md.User, error) {

	query := `
		SELECT
			id,
			email,
			username,
			password_hash,
			google_id,
			role,
			created_at
		FROM users
		WHERE id = $1
	`

	row := r.db.QueryRow(context.Background(), query, id)

	var user md.User

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.Username,
		&user.PasswordHash,
		&user.GoogleID,
		&user.Role,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) GetAllUsers() ([]md.User, error) {

	query := `
		SELECT
			id,
			email,
			username,
			role,
			created_at
		FROM users
	`

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []md.User

	for rows.Next() {

		var user md.User

		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.Username,
			&user.Role,
			&user.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *Repository) CreateUser(user *md.User) error {
    query := `
        INSERT INTO users (id, email, username, password_hash, role, created_at)
        VALUES ($1, $2, $3, $4, $5, $6)
    `
    _, err := r.db.Exec(context.Background(), query,
        user.ID,
        user.Email,
        user.Username,
        user.PasswordHash,
        user.Role,
        user.CreatedAt,
    )
    return err
}

func (r *Repository) GetByEmail(email string) (*md.User, error) {
    query := `
        SELECT id, email, username, password_hash, role, created_at
        FROM users WHERE email = $1
    `
    row := r.db.QueryRow(context.Background(), query, email)

    var user md.User
    err := row.Scan(
        &user.ID, &user.Email, &user.Username,
        &user.PasswordHash, &user.Role, &user.CreatedAt,
    )
    if err != nil {
        return nil, err
    }
    return &user, nil
}