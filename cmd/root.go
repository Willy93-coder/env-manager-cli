// Copyright © 2026 willy93-coder

// Package cmd contains all CLI commands for the env-manager application
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/willy93-coder/env-manager-cli/internal/config"
	"github.com/willy93-coder/env-manager-cli/internal/db"
)

// Package-level variables
var (
	cfg      *config.Config
	database *db.DB
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "env-manager",
	Short: "Manage .env files with PostgreSQL backend",
	Long: `env-manager is a CLI tool to manage environment variables
	across multiple projects, stored securely in PostgreSQL.

	Example:
		env-manager init my-project
		env-manager set my-project DB_HOST localhost
		env-manager get my-project DB_HOST`,
	// PersistentPreRunE runs before ANY subcommand
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// 'setup' is the only command that does not need a DB connection
		// It creates the config file, so it cannot depend on it existing.
		if cmd.Name() == "setup" {
			return nil
		}

		// Load config from ~/.env-manager/config.yaml
		loadedCfg, err := config.LoadDefault()
		if err != nil {
			return err
		}
		cfg = loadedCfg

		// Connect to PostgreSQL using loaded config
		connectedDB, err := db.New(cfg.DSN())
		if err != nil {
			return err
		}
		database = connectedDB

		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Subcommands
	rootCmd.AddCommand(setupCmd)
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(setCmd)
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(listCmd)
}
