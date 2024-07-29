// Package persistence -----------------------------
// @file      : user_repository.go
// @author    : Jimmy
// @contact   : 2114272829@qq.com
// @time      : 2024/7/29 下午1:37
// -------------------------------------------
package persistence

import (
	"context"
	"database/sql"
	"github.com/ZUCCzwp/kitex_ddd_example/internal/domain/entity"

	"github.com/ZUCCzwp/kitex_ddd_example/internal/domain/repository"
	_ "github.com/go-sql-driver/mysql" // driver
)

// userRepository Implements repository.UserRepository
type userRepository struct {
	conn *sql.DB
}

// NewUserRepository returns initialized UserRepositoryImpl
func NewUserRepository(conn *sql.DB) repository.UserRepository {
	return &userRepository{conn: conn}
}

// Get returns domain.User
func (r *userRepository) Get(ctx context.Context, id int) (*entity.User, error) {
	row, err := r.queryRow(ctx, "select id, name from users where id=?", id)
	if err != nil {
		return nil, err
	}
	u := &entity.User{}
	err = row.Scan(&u.ID, &u.Name)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// GetAll returns list of domain.User
func (r *userRepository) GetAll(ctx context.Context) ([]*entity.User, error) {
	rows, err := r.query(ctx, "select id, name from users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	us := make([]*entity.User, 0)
	for rows.Next() {
		u := &entity.User{}
		err = rows.Scan(&u.ID, &u.Name)
		if err != nil {
			return nil, err
		}
		us = append(us, u)
	}
	return us, nil
}

// Save saves domain.User to storage
func (r *userRepository) Save(ctx context.Context, u *entity.User) error {
	stmt, err := r.conn.Prepare("insert into users (name) values (?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, u.Name)
	return err
}

func (r *userRepository) query(ctx context.Context, q string, args ...interface{}) (*sql.Rows, error) {
	stmt, err := r.conn.Prepare(q)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return stmt.QueryContext(ctx, args...)
}

func (r *userRepository) queryRow(ctx context.Context, q string, args ...interface{}) (*sql.Row, error) {
	stmt, err := r.conn.Prepare(q)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return stmt.QueryRowContext(ctx, args...), nil
}
