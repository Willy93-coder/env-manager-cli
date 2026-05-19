package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an environment variable from a project",
	Args:  cobra.RangeArgs(1, 2),
	RunE: func(cmd *cobra.Command, args []string) error {
		project, err := database.GetProjectByName(args[0])
		if err != nil {
			return err
		}
		all, err := cmd.Flags().GetBool("all")
		if err != nil {
			return err
		}
		if !all && len(args) < 2 {
			return fmt.Errorf("key is required when --all flag is not set")
		}
		if all {
			err = database.DeleteEnvironments(project.ID)
			if err != nil {
				return err
			}
			fmt.Printf("Successfully deleted environment variables from %s project\n", args[0])
			return nil
		}

		err = database.DeleteEnvironment(project.ID, args[1])
		if err != nil {
			return err
		}

		fmt.Printf("Successfully deleted %s from %s project\n", args[1], args[0])
		return nil
	},
}
