package pg

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/nimoism/ad-rotator/internal/entity"
)

type UserGroupRepo struct {
	db *sql.DB
}

func NewUserGroupRepo(db *sql.DB) *UserGroupRepo {
	return &UserGroupRepo{db: db}
}

func (r *UserGroupRepo) UsersGroups(ctx context.Context) ([]entity.UserGroup, error) {
	query := "SELECT id, name FROM user_group ORDER BY id"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("all user group query error: %w", err)
	}
	defer rows.Close()
	ugs := make([]entity.UserGroup, 0)
	for rows.Next() {
		ug := entity.UserGroup{}
		if err = rows.Scan(&ug.ID, &ug.Name); err != nil {
			return nil, fmt.Errorf("user group db mapping error: %w", err)
		}
		ugs = append(ugs, ug)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("getting users groups error: %w", err)
	}
	return ugs, nil
}

func (r *UserGroupRepo) CreateUserGroup(ctx context.Context, ug *entity.UserGroup) error {
	query := "INSERT INTO user_group (name) VALUES ($1) RETURNING id"
	result := r.db.QueryRowContext(ctx, query, ug.Name)
	if err := result.Scan(&ug.ID); err != nil {
		return err
	}
	return result.Err()
}

func (r *UserGroupRepo) UpdateUserGroup(ctx context.Context, ug *entity.UserGroup) error {
	query := "UPDATE user_group SET name = $2 WHERE id = $1"
	_, err := r.db.ExecContext(ctx, query, ug.ID, ug.Name)
	return err
}

func (r *UserGroupRepo) DeleteUserGroup(ctx context.Context, id int) error {
	query := "DELETE FROM user_group WHERE id = $1"
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
