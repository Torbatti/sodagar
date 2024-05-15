package cmd

import (
	"errors"

	"github.com/openjogd/sodagar/core"
	"github.com/spf13/cobra"
)

func NewAdminCommand(app core.App) *cobra.Command {
	command := &cobra.Command{
		Use:   "admin",
		Short: "Manages admin accounts",
	}

	command.AddCommand(adminCreateCommand(app))
	command.AddCommand(adminUpdateCommand(app))
	command.AddCommand(adminDeleteCommand(app))

	return command
}

func adminCreateCommand(app core.App) *cobra.Command {
	command := &cobra.Command{
		Use:          "create",
		Example:      "admin create test@example.com 1234567890",
		Short:        "Creates a new admin account",
		SilenceUsage: true,
		RunE: func(command *cobra.Command, args []string) error {
			return errors.New("TODO")
		},
	}

	return command
}

func adminUpdateCommand(app core.App) *cobra.Command {
	command := &cobra.Command{
		Use:          "update",
		Example:      "admin update test@example.com 1234567890",
		Short:        "Changes the password of a single admin account",
		SilenceUsage: true,
		RunE: func(command *cobra.Command, args []string) error {
			return errors.New("TODO")
		},
	}

	return command
}

func adminDeleteCommand(app core.App) *cobra.Command {
	command := &cobra.Command{
		Use:          "delete",
		Example:      "admin delete test@example.com",
		Short:        "Deletes an existing admin account",
		SilenceUsage: true,
		RunE: func(command *cobra.Command, args []string) error {
			return errors.New("TODO")
		},
	}

	return command
}
