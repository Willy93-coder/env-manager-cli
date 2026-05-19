package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export project environment variables from the database",
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

		dir, err := os.Getwd()
		if err != nil {
			return err
		}

		var sb strings.Builder
		for _, env := range environments {
			sb.WriteString(env.Key + "=" + env.Value + "\n")
		}

		content := sb.String()

		err = os.WriteFile(filepath.Join(dir, ".env"), []byte(content), 0600)
		if err != nil {
			return err
		}

		fmt.Printf("Successfully exported .env file in %s", dir)

		return nil
	},
}
