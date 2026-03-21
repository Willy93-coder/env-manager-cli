package db

import (
	"context"
	"fmt"

	"github.com/willy93-coder/env-manager-cli/internal/models"
)

func (db *DB) CreateProject(name string, description *string) (*models.Project, error) {
	query := `
			INSERT INTO projects (name, description) 
			VALUES($1, $2) 
			RETURNING id, name, description, created_at, updated_at
	`
	var project models.Project
	ctx := context.Background()
	row := db.QueryRowContext(ctx, query, name, description)
	err := row.Scan(&project.ID, &project.Name, &project.Description, &project.CreatedAt, &project.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("could not create project: %w", err)
	}
	return &project, nil
}

func (db *DB) GetProjectByName(name string) (*models.Project, error) {
	query := `
		SELECT * FROM projects
		WHERE name = $1
	`
	var project models.Project
	ctx := context.Background()
	row := db.QueryRowContext(ctx, query, name)
	err := row.Scan(&project.ID, &project.Name, &project.Description, &project.CreatedAt, &project.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("could not get project %s: %w", name, err)
	}
	return &project, nil
}

func (db *DB) GetProjects() ([]models.Project, error) {
	query := `
			SELECT * from projects
	`

	ctx := context.Background()
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("could not get all projects: %w", err)
	}
	defer rows.Close()
	projects := make([]models.Project, 0)

	for rows.Next() {
		var project models.Project
		if err := rows.Scan(&project.ID, &project.Name, &project.Description, &project.CreatedAt, &project.UpdatedAt); err != nil {
			return nil, fmt.Errorf("could not parse project: %w", err)
		}
		projects = append(projects, project)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating project: %w", err)
	}
	return projects, nil
}

func (db *DB) UpdateProject(id int64, name, description *string) (*models.Project, error) {
	query := `
		UPDATE projects
		SET name = $1, description = $2
		WHERE id = $3
		RETURNING id, name, description, created_at, updated_at
	`

	var project models.Project
	ctx := context.Background()
	row := db.QueryRowContext(ctx, query, name, description, id)
	err := row.Scan(&project.ID, &project.Name, &project.Description, &project.CreatedAt, &project.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("could not update project: %w", err)
	}

	return &project, nil
}

func (db *DB) DeleteProject(name string) error {
	query := `
		DELETE FROM projects
		WHERE name = $1
	`
	ctx := context.Background()
	result, err := db.ExecContext(ctx, query, name)
	if err != nil {
		return fmt.Errorf("could not delete project: %w", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not delete project: %w", err)
	}
	if rows != 1 {
		return fmt.Errorf("expected to affect 1 row: %d", rows)
	}
	return nil
}
