package user

import (
	"context"
	"database/sql"
)

// interface for passing sb/transaction object that performs queries
type IDBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt,error)
	QueryContext(context.Context, string, ...interface{})(*sql.Rows,error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

type repository struct {
	db IDBTX
}

func(r *repository) CreateUser(ctx context.Context, user *User) (*User, error) {

	var lastInsertId int
	query := "INSERT INTO USERS users(username,password_hash,email) VALUES ($1,$2,$3) returning id"
	err := r.db.QueryRowContext(ctx,query,user.Username,user.Password,user.Email).Scan(&lastInsertId)
	if err != nil {
		return &User{}, err
	}
	user.Id = int64(lastInsertId)
	return user, nil
}

func NewRepository(db IDBTX) IRepository {
	return &repository{db:db}
}
