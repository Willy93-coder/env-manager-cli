package db

import "github.com/willy93-coder/env-manager-cli/internal/models"

func (db *DB) CreateEnvironment(projectID int64, key, value string) (*models.Environment, error) {
	// TODO
	return nil, nil
}

func (db *DB) CreateEnvironments(projectID int64, envs map[string]string) error {
	// TODO
	return nil
}

func (db *DB) GetEnvironments(projectID int64) ([]models.Environment, error) {
	// TODO
	return nil, nil
}

func (db *DB) UpdateEnvironment(projectID int64, key, value string) (*models.Environment, error) {
	// TODO
	return nil, nil
}

func (db *DB) UpdateEnvironments(projectID int64, envs map[string]string) error {
	// TODO
	return nil
}

func (db *DB) DeleteEnvironment(projectID int64, key string) error {
	// TODO
	return nil
}

func (db *DB) DeleteEnvironments(projectID int64) error {
	// TODO
	return nil
}
