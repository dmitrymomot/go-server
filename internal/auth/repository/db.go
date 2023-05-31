// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package auth_repository

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.cleanUpTokensStmt, err = db.PrepareContext(ctx, cleanUpTokens); err != nil {
		return nil, fmt.Errorf("error preparing query CleanUpTokens: %w", err)
	}
	if q.cleanUpVerificationsStmt, err = db.PrepareContext(ctx, cleanUpVerifications); err != nil {
		return nil, fmt.Errorf("error preparing query CleanUpVerifications: %w", err)
	}
	if q.createUserStmt, err = db.PrepareContext(ctx, createUser); err != nil {
		return nil, fmt.Errorf("error preparing query CreateUser: %w", err)
	}
	if q.deleteTokenByAccessTokenIDStmt, err = db.PrepareContext(ctx, deleteTokenByAccessTokenID); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteTokenByAccessTokenID: %w", err)
	}
	if q.deleteTokenByRefreshTokenIDStmt, err = db.PrepareContext(ctx, deleteTokenByRefreshTokenID); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteTokenByRefreshTokenID: %w", err)
	}
	if q.deleteTokensByUserIDStmt, err = db.PrepareContext(ctx, deleteTokensByUserID); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteTokensByUserID: %w", err)
	}
	if q.deleteUserByIDStmt, err = db.PrepareContext(ctx, deleteUserByID); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteUserByID: %w", err)
	}
	if q.findTokenByAccessTokenIDStmt, err = db.PrepareContext(ctx, findTokenByAccessTokenID); err != nil {
		return nil, fmt.Errorf("error preparing query FindTokenByAccessTokenID: %w", err)
	}
	if q.findTokenByRefreshTokenIDStmt, err = db.PrepareContext(ctx, findTokenByRefreshTokenID); err != nil {
		return nil, fmt.Errorf("error preparing query FindTokenByRefreshTokenID: %w", err)
	}
	if q.findUserByEmailStmt, err = db.PrepareContext(ctx, findUserByEmail); err != nil {
		return nil, fmt.Errorf("error preparing query FindUserByEmail: %w", err)
	}
	if q.findUserByIDStmt, err = db.PrepareContext(ctx, findUserByID); err != nil {
		return nil, fmt.Errorf("error preparing query FindUserByID: %w", err)
	}
	if q.findVerificationByIDStmt, err = db.PrepareContext(ctx, findVerificationByID); err != nil {
		return nil, fmt.Errorf("error preparing query FindVerificationByID: %w", err)
	}
	if q.refreshTokenStmt, err = db.PrepareContext(ctx, refreshToken); err != nil {
		return nil, fmt.Errorf("error preparing query RefreshToken: %w", err)
	}
	if q.storeOrUpdateVerificationStmt, err = db.PrepareContext(ctx, storeOrUpdateVerification); err != nil {
		return nil, fmt.Errorf("error preparing query StoreOrUpdateVerification: %w", err)
	}
	if q.storeTokenStmt, err = db.PrepareContext(ctx, storeToken); err != nil {
		return nil, fmt.Errorf("error preparing query StoreToken: %w", err)
	}
	if q.updateUserEmailByIDStmt, err = db.PrepareContext(ctx, updateUserEmailByID); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateUserEmailByID: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.cleanUpTokensStmt != nil {
		if cerr := q.cleanUpTokensStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing cleanUpTokensStmt: %w", cerr)
		}
	}
	if q.cleanUpVerificationsStmt != nil {
		if cerr := q.cleanUpVerificationsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing cleanUpVerificationsStmt: %w", cerr)
		}
	}
	if q.createUserStmt != nil {
		if cerr := q.createUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createUserStmt: %w", cerr)
		}
	}
	if q.deleteTokenByAccessTokenIDStmt != nil {
		if cerr := q.deleteTokenByAccessTokenIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteTokenByAccessTokenIDStmt: %w", cerr)
		}
	}
	if q.deleteTokenByRefreshTokenIDStmt != nil {
		if cerr := q.deleteTokenByRefreshTokenIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteTokenByRefreshTokenIDStmt: %w", cerr)
		}
	}
	if q.deleteTokensByUserIDStmt != nil {
		if cerr := q.deleteTokensByUserIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteTokensByUserIDStmt: %w", cerr)
		}
	}
	if q.deleteUserByIDStmt != nil {
		if cerr := q.deleteUserByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteUserByIDStmt: %w", cerr)
		}
	}
	if q.findTokenByAccessTokenIDStmt != nil {
		if cerr := q.findTokenByAccessTokenIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findTokenByAccessTokenIDStmt: %w", cerr)
		}
	}
	if q.findTokenByRefreshTokenIDStmt != nil {
		if cerr := q.findTokenByRefreshTokenIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findTokenByRefreshTokenIDStmt: %w", cerr)
		}
	}
	if q.findUserByEmailStmt != nil {
		if cerr := q.findUserByEmailStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findUserByEmailStmt: %w", cerr)
		}
	}
	if q.findUserByIDStmt != nil {
		if cerr := q.findUserByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findUserByIDStmt: %w", cerr)
		}
	}
	if q.findVerificationByIDStmt != nil {
		if cerr := q.findVerificationByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findVerificationByIDStmt: %w", cerr)
		}
	}
	if q.refreshTokenStmt != nil {
		if cerr := q.refreshTokenStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing refreshTokenStmt: %w", cerr)
		}
	}
	if q.storeOrUpdateVerificationStmt != nil {
		if cerr := q.storeOrUpdateVerificationStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing storeOrUpdateVerificationStmt: %w", cerr)
		}
	}
	if q.storeTokenStmt != nil {
		if cerr := q.storeTokenStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing storeTokenStmt: %w", cerr)
		}
	}
	if q.updateUserEmailByIDStmt != nil {
		if cerr := q.updateUserEmailByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateUserEmailByIDStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                              DBTX
	tx                              *sql.Tx
	cleanUpTokensStmt               *sql.Stmt
	cleanUpVerificationsStmt        *sql.Stmt
	createUserStmt                  *sql.Stmt
	deleteTokenByAccessTokenIDStmt  *sql.Stmt
	deleteTokenByRefreshTokenIDStmt *sql.Stmt
	deleteTokensByUserIDStmt        *sql.Stmt
	deleteUserByIDStmt              *sql.Stmt
	findTokenByAccessTokenIDStmt    *sql.Stmt
	findTokenByRefreshTokenIDStmt   *sql.Stmt
	findUserByEmailStmt             *sql.Stmt
	findUserByIDStmt                *sql.Stmt
	findVerificationByIDStmt        *sql.Stmt
	refreshTokenStmt                *sql.Stmt
	storeOrUpdateVerificationStmt   *sql.Stmt
	storeTokenStmt                  *sql.Stmt
	updateUserEmailByIDStmt         *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                              tx,
		tx:                              tx,
		cleanUpTokensStmt:               q.cleanUpTokensStmt,
		cleanUpVerificationsStmt:        q.cleanUpVerificationsStmt,
		createUserStmt:                  q.createUserStmt,
		deleteTokenByAccessTokenIDStmt:  q.deleteTokenByAccessTokenIDStmt,
		deleteTokenByRefreshTokenIDStmt: q.deleteTokenByRefreshTokenIDStmt,
		deleteTokensByUserIDStmt:        q.deleteTokensByUserIDStmt,
		deleteUserByIDStmt:              q.deleteUserByIDStmt,
		findTokenByAccessTokenIDStmt:    q.findTokenByAccessTokenIDStmt,
		findTokenByRefreshTokenIDStmt:   q.findTokenByRefreshTokenIDStmt,
		findUserByEmailStmt:             q.findUserByEmailStmt,
		findUserByIDStmt:                q.findUserByIDStmt,
		findVerificationByIDStmt:        q.findVerificationByIDStmt,
		refreshTokenStmt:                q.refreshTokenStmt,
		storeOrUpdateVerificationStmt:   q.storeOrUpdateVerificationStmt,
		storeTokenStmt:                  q.storeTokenStmt,
		updateUserEmailByIDStmt:         q.updateUserEmailByIDStmt,
	}
}
