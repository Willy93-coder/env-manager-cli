package db

import "github.com/willy93-coder/env-manager-cli/internal/models"

func (db *DB) CreateProject(name string, description *string) (*models.Project, error) {
	// TODO
	return nil, nil
}

func (db *DB) GetProjectByName(name string) (*models.Project, error) {
	// TODO
	return nil, nil
}

func (db *DB) GetProjects() ([]models.Project, error) {
	// TODO
	return nil, nil
}

func (db *DB) UpdateProject(id int64, name, description *string) (*models.Project, error) {
	// TODO
	return nil, nil
}

func (db *DB) DeleteProject(name string) error {
	// TODO
	return nil
}
