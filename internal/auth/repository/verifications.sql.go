// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: verifications.sql

package auth_repository

import (
	"context"

	"github.com/google/uuid"
)

const cleanUpVerifications = `-- name: CleanUpVerifications :exec
DELETE FROM verifications WHERE expires_at < now()
`

// Clean up expired verifications
func (q *Queries) CleanUpVerifications(ctx context.Context) error {
	_, err := q.exec(ctx, q.cleanUpVerificationsStmt, cleanUpVerifications)
	return err
}

const deleteVerificationByID = `-- name: DeleteVerificationByID :exec
DELETE FROM verifications WHERE id = $1
`

// Delete a verification by ID
func (q *Queries) DeleteVerificationByID(ctx context.Context, id uuid.UUID) error {
	_, err := q.exec(ctx, q.deleteVerificationByIDStmt, deleteVerificationByID, id)
	return err
}

const findVerificationByID = `-- name: FindVerificationByID :one
SELECT id, user_id, verification_type, email, otp_hash, created_at, expires_at FROM verifications WHERE id = $1 AND expires_at > now()
`

// Find a verification by ID
func (q *Queries) FindVerificationByID(ctx context.Context, id uuid.UUID) (Verification, error) {
	row := q.queryRow(ctx, q.findVerificationByIDStmt, findVerificationByID, id)
	var i Verification
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.VerificationType,
		&i.Email,
		&i.OtpHash,
		&i.CreatedAt,
		&i.ExpiresAt,
	)
	return i, err
}

const storeOrUpdateVerification = `-- name: StoreOrUpdateVerification :one
INSERT INTO verifications (user_id, verification_type, email, otp_hash) VALUES ($1, $2, $3, $4) RETURNING id
`

type StoreOrUpdateVerificationParams struct {
	UserID           uuid.UUID `json:"user_id"`
	VerificationType string    `json:"verification_type"`
	Email            string    `json:"email"`
	OtpHash          []byte    `json:"otp_hash"`
}

// Store or update a user's verification code.
func (q *Queries) StoreOrUpdateVerification(ctx context.Context, arg StoreOrUpdateVerificationParams) (uuid.UUID, error) {
	row := q.queryRow(ctx, q.storeOrUpdateVerificationStmt, storeOrUpdateVerification,
		arg.UserID,
		arg.VerificationType,
		arg.Email,
		arg.OtpHash,
	)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}
