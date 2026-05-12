package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all environment variables",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		project, err := database.GetProjectByName(args[0])
		if err != nil {
			return err
		}

		environments, err := database.GetEnvironments(project.ID)
		if err != nil {
			return err
		}

		for _, env := range environments {
			fmt.Printf("%s=%s\n", env.Key, env.Value)
		}

		return nil
	},
}
