package cli

import (
	"github.com/spf13/cobra"
)

func DatabaseCmd() *cobra.Command {
	var configFile string
	short := "Database management commands for GodoBase"
	long := "Database management commands for GodoBase"
	cmd := &cobra.Command{
		Use:   "database",
		Short: short,
		Long:  long,
		Run: func(cmd *cobra.Command, args []string) {
			err := cmd.Help()
			if err != nil {
				return
			}
		},
	}
	cmd.Flags().StringVar(&configFile, "config-file", "config.yaml", "Path to configuration file")

	cmd.AddCommand(startDatabaseCmd())

	cmd.AddCommand(stopDatabaseCmd())

	cmd.AddCommand(statusDatabaseCmd())

	return cmd
}

func startDatabaseCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start Database services",
		Long:  "Start Database services",
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}
	return cmd
}

func stopDatabaseCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stop",
		Short: "Stop Docker",
		Long:  "Stop Docker service",
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}
	return cmd
}

func statusDatabaseCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "status",
		Short: "Status Docker",
		Long:  "Status Docker service",
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}
	return cmd
}
