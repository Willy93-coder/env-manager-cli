package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set an environment variable for a project",
	Args:  cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		project, err := database.GetProjectByName(args[0])
		if err != nil {
			return err
		}

		environment, err := database.CreateEnvironment(project.ID, args[1], args[2])
		if err != nil {
			return err
		}

		fmt.Printf("'%s' set for project '%s'\n", environment.Key, project.Name)

		return nil
	},
}
