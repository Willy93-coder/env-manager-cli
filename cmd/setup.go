package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/willy93-coder/env-manager-cli/internal/config"
	"golang.org/x/term"
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup database",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		host, err := promptString("Database host", "localhost")
		if err != nil {
			return err
		}

		port, err := promptInt("Database port", 5432)
		if err != nil {
			return err
		}

		user, err := promptString("Database user", "postgres")
		if err != nil {
			return err
		}

		password, err := promptPassword("Database password")
		if err != nil {
			return err
		}

		fmt.Println()

		dbname, err := promptString("Database name", "env_manager")
		if err != nil {
			return err
		}

		sslmode, err := promptString("sslmode", "disable")
		if err != nil {
			return err
		}

		cfg := &config.Config{
			Database: config.DatabaseConfig{
				Host:     host,
				Port:     port,
				User:     user,
				Password: password,
				DBName:   dbname,
				SSLMode:  sslmode,
			},
		}

		err = config.SaveDefault(cfg)
		if err != nil {
			return err
		}

		fmt.Println("Database config created successfully")

		return nil
	},
}

func promptString(label, defaultValue string) (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("%s [%s]: ", label, defaultValue)
	line, err := reader.ReadString('\n')

	if err != nil {
		return "", fmt.Errorf("could not read input: %w", err)
	}

	line = strings.TrimSpace(line)

	if line == "" {
		return defaultValue, nil
	}

	return line, nil
}

func promptInt(label string, defaultValue int) (int, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("%s [%d]: ", label, defaultValue)
	line, err := reader.ReadString('\n')

	if err != nil {
		return 0, fmt.Errorf("could not read input: %w", err)
	}

	line = strings.TrimSpace(line)

	if line == "" {
		return defaultValue, nil
	}

	value, err := strconv.Atoi(line)

	if err != nil {
		return 0, fmt.Errorf("invalid number: %w", err)
	}

	return value, nil
}

func promptPassword(label string) (string, error) {
	fmt.Printf("%s: ", label)
	line, err := term.ReadPassword(int(os.Stdin.Fd()))

	if err != nil {
		return "", fmt.Errorf("could not read input: %w", err)
	}

	password := string(line)

	return password, nil
}
