// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: users.sql

package auth_repository

import (
	"context"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (email) VALUES ($1) ON CONFLICT (email) DO NOTHING RETURNING id
`

// Store or update a user
func (q *Queries) CreateUser(ctx context.Context, email string) (uuid.UUID, error) {
	row := q.queryRow(ctx, q.createUserStmt, createUser, email)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const deleteUserByID = `-- name: DeleteUserByID :exec
DELETE FROM users WHERE id = $1
`

// Delete a user by ID
func (q *Queries) DeleteUserByID(ctx context.Context, id uuid.UUID) error {
	_, err := q.exec(ctx, q.deleteUserByIDStmt, deleteUserByID, id)
	return err
}

const findUserByEmail = `-- name: FindUserByEmail :one
SELECT id, email, verified, updated_at, created_at FROM users WHERE email = $1
`

// Find a user by email
func (q *Queries) FindUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.queryRow(ctx, q.findUserByEmailStmt, findUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Verified,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const findUserByID = `-- name: FindUserByID :one
SELECT id, email, verified, updated_at, created_at FROM users WHERE id = $1
`

// Find a user by ID
func (q *Queries) FindUserByID(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.queryRow(ctx, q.findUserByIDStmt, findUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Verified,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const updateUserEmailByID = `-- name: UpdateUserEmailByID :exec
UPDATE users SET email = $1, verified = $3 WHERE id = $2
`

type UpdateUserEmailByIDParams struct {
	Email    string    `json:"email"`
	ID       uuid.UUID `json:"id"`
	Verified bool      `json:"verified"`
}

// Update a user's email by ID
func (q *Queries) UpdateUserEmailByID(ctx context.Context, arg UpdateUserEmailByIDParams) error {
	_, err := q.exec(ctx, q.updateUserEmailByIDStmt, updateUserEmailByID, arg.Email, arg.ID, arg.Verified)
	return err
}

const updateUserVerificationStatusByID = `-- name: UpdateUserVerificationStatusByID :exec
UPDATE users SET verified = $1 WHERE id = $2
`

type UpdateUserVerificationStatusByIDParams struct {
	Verified bool      `json:"verified"`
	ID       uuid.UUID `json:"id"`
}

// Update a user's verification status by ID
func (q *Queries) UpdateUserVerificationStatusByID(ctx context.Context, arg UpdateUserVerificationStatusByIDParams) error {
	_, err := q.exec(ctx, q.updateUserVerificationStatusByIDStmt, updateUserVerificationStatusByID, arg.Verified, arg.ID)
	return err
}
