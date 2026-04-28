package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup database",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
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
