package db

import (
	"context"
	"fmt"

	"github.com/willy93-coder/env-manager-cli/internal/models"
)

func (db *DB) CreateEnvironment(projectID int64, key, value string) (*models.Environment, error) {
	query := `
		INSERT INTO environments (project_id, key, value)
		VALUES ($1, $2, $3)
		RETURNING id, project_id, key, value, created_at, updated_at
	`
	var environment models.Environment
	ctx := context.Background()
	row := db.QueryRowContext(ctx, query, projectID, key, value)
	err := row.Scan(&environment.ID, &environment.ProjectID, &environment.Key, &environment.Value, &environment.CreatedAt, &environment.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("could not create an environment: %w", err)
	}
	return &environment, nil
}

func (db *DB) CreateEnvironments(projectID int64, envs map[string]string) error {
	query := `
		INSERT INTO environments (project_id, key, value)
		VALUES ($1, $2, $3)
	`
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("could not begin transaction: %w", err)
	}
	defer tx.Rollback()
	for key, value := range envs {
		_, err := tx.ExecContext(ctx, query, projectID, key, value)
		if err != nil {
			return fmt.Errorf("could not insert environment %s: %w", key, err)
		}
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("could not commit transaction: %w", err)
	}
	return nil
}

func (db *DB) GetEnvironments(projectID int64) ([]models.Environment, error) {
	query := `
		SELECT * FROM environments
		WHERE project_id = $1
	`
	ctx := context.Background()
	rows, err := db.QueryContext(ctx, query, projectID)
	if err != nil {
		return nil, fmt.Errorf("could not get environments: %w", err)
	}
	defer rows.Close()
	environments := make([]models.Environment, 0)

	for rows.Next() {
		var environment models.Environment
		if err := rows.Scan(&environment.ID, &environment.ProjectID, &environment.Key, &environment.Value, &environment.CreatedAt, &environment.UpdatedAt); err != nil {
			return nil, fmt.Errorf("could not parse environment: %w", err)
		}
		environments = append(environments, environment)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating environment: %w", err)
	}
	return environments, nil
}

func (db *DB) GetEnvironmentByKey(projectID int64, key string) (*models.Environment, error) {
	query := `
		SELECT * FROM environments
		WHERE project_id = $1 AND key = $2
	`

	var environment models.Environment
	ctx := context.Background()
	row := db.QueryRowContext(ctx, query, projectID, key)
	err := row.Scan(&environment.ID, &environment.ProjectID, &environment.Key, &environment.Value, &environment.CreatedAt, &environment.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("could not get key: %w", err)
	}

	return &environment, nil
}

func (db *DB) UpdateEnvironment(projectID int64, key, value string) (*models.Environment, error) {
	query := `
		UPDATE environments
		SET value = $1
		WHERE project_id = $2 AND key = $3
		RETURNING id, project_id, key, value, created_at, updated_at
	`
	var environment models.Environment
	ctx := context.Background()
	row := db.QueryRowContext(ctx, query, value, projectID, key)
	err := row.Scan(&environment.ID, &environment.ProjectID, &environment.Key, &environment.Value, &environment.CreatedAt, &environment.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("could not update environment: %w", err)
	}
	return &environment, nil
}

func (db *DB) UpdateEnvironments(projectID int64, envs map[string]string) error {
	query := `
		UPDATE environments
		SET value = $1
		WHERE project_id = $2 AND key = $3
	`
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("could not begin transaction: %w", err)
	}
	defer tx.Rollback()
	for key, value := range envs {
		_, err := tx.ExecContext(ctx, query, value, projectID, key)
		if err != nil {
			return fmt.Errorf("could not update environment %s: %w", key, err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("could not commit transaction: %w", err)
	}

	return nil
}

func (db *DB) DeleteEnvironment(projectID int64, key string) error {
	query := `
		DELETE FROM environments
		WHERE project_id = $1 AND key = $2
	`
	ctx := context.Background()
	result, err := db.ExecContext(ctx, query, projectID, key)
	if err != nil {
		return fmt.Errorf("could not delete key: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not delete environment: %w", err)
	}
	if rows != 1 {
		return fmt.Errorf("expected to affect 1 row: %d", rows)
	}

	return nil
}

func (db *DB) DeleteEnvironments(projectID int64) error {
	query := `
		DELETE FROM environments
		WHERE project_id = $1
	`
	ctx := context.Background()
	_, err := db.ExecContext(ctx, query, projectID)
	if err != nil {
		return fmt.Errorf("could not delete environments: %w", err)
	}

	return nil
}
