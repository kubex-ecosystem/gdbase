package cli

import (
	"os"

	"github.com/spf13/cobra"
)

func DatabaseCmd() *cobra.Command {
	var configFile string
	shortDesc := "Database management commands for GodoBase"
	longDesc := "Database management commands for GodoBase"
	cmd := &cobra.Command{
		Use:         "database",
		Short:       shortDesc,
		Long:        longDesc,
		Annotations: GetDescriptions([]string{shortDesc, longDesc}, (os.Getenv("GDBASE_HIDEBANNER") == "true")),
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
	shortDesc := "Start Database services"
	longDesc := "Start Database services"
	cmd := &cobra.Command{
		Use:         "start",
		Short:       shortDesc,
		Long:        longDesc,
		Annotations: GetDescriptions([]string{shortDesc, longDesc}, (os.Getenv("GDBASE_HIDEBANNER") == "true")),
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}
	return cmd
}

func stopDatabaseCmd() *cobra.Command {

	shortDesc := "Stop Docker"
	longDesc := "Stop Docker service"

	cmd := &cobra.Command{
		Use:         "stop",
		Short:       shortDesc,
		Long:        longDesc,
		Annotations: GetDescriptions([]string{shortDesc, longDesc}, (os.Getenv("GDBASE_HIDEBANNER") == "true")),
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}
	return cmd
}

func statusDatabaseCmd() *cobra.Command {

	shortDesc := "Status Docker"
	longDesc := "Status Docker service"

	cmd := &cobra.Command{
		Use:         "status",
		Short:       shortDesc,
		Long:        longDesc,
		Annotations: GetDescriptions([]string{shortDesc, longDesc}, (os.Getenv("GDBASE_HIDEBANNER") == "true")),
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}
	return cmd
}
