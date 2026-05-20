package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import .env file to database",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		project, err := database.GetProjectByName(args[0])
		if err != nil {
			return err
		}

		environments := make(map[string]string)
		// read .env file
		content, err := os.ReadFile(args[1])
		if err != nil {
			return fmt.Errorf("cannot read %s file %w", args[1], err)
		}
		for value := range strings.SplitSeq(string(content), "\n") {
			if value == "" || strings.HasPrefix(value, "#") {
				continue
			}
			pairs := strings.SplitN(value, "=", 2)
			if len(pairs) != 2 {
				continue
			}
			environments[pairs[0]] = pairs[1]
		}

		err = database.CreateEnvironments(project.ID, environments)
		if err != nil {
			return err
		}

		fmt.Printf("Successfully imported file in %s project", args[0])

		return nil
	},
}
