package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a variable by key",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		project, err := database.GetProjectByName(args[0])
		if err != nil {
			return err
		}

		env, err := database.GetEnvironmentByKey(project.ID, args[1])
		if err != nil {
			return err
		}

		fmt.Printf("%s=%s\n", env.Key, env.Value)

		return nil
	},
}
