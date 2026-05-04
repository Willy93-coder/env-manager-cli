package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a new project in the database",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		project, err := database.CreateProject(args[0], nil)
		if err != nil {
			return err
		}

		fmt.Printf("Project '%s' created successfully\n", project.Name)

		return nil
	},
}
