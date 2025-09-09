package cli

import (
	"fmt"
	"os"

	"github.com/kubex-ecosystem/gdbase/factory"
	s "github.com/kubex-ecosystem/gdbase/internal/services"
	l "github.com/kubex-ecosystem/logz"
	"github.com/spf13/cobra"
)

func DockerCmd() *cobra.Command {
	var configFile string

	shortDesc := "Docker management commands for GodoBase"
	longDesc := "Docker management commands for GodoBase"

	cmd := &cobra.Command{
		Use:         "docker",
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

	cmd.AddCommand(startDockerCmd())

	cmd.AddCommand(stopDockerCmd())

	cmd.AddCommand(statusDockerCmd())

	return cmd
}

func startDockerCmd() *cobra.Command {

	shortDesc := "Start Docker"
	longDesc := "Start Docker service"

	cmd := &cobra.Command{
		Use:         "start",
		Short:       shortDesc,
		Long:        longDesc,
		Annotations: GetDescriptions([]string{shortDesc, longDesc}, (os.Getenv("GDBASE_HIDEBANNER") == "true")),
		Run: func(cmd *cobra.Command, args []string) {
			dkr, dkrErr := factory.NewDockerService(nil, l.GetLogger("GodoBase"))
			if dkrErr != nil {
				fmt.Printf("Error starting Docker service: %v\n", dkrErr)
				return
			}
			dkrErr = dkr.Initialize()
			if dkrErr != nil {
				fmt.Printf("Error initializing Docker service: %v\n", dkrErr)
				return
			}
			dkrErr = s.SetupDatabaseServices(dkr, nil)
		},
	}
	return cmd
}

func stopDockerCmd() *cobra.Command {

	shortDesc := "Stop Docker"
	longDesc := "Stop Docker service"

	cmd := &cobra.Command{
		Use:         "stop",
		Short:       shortDesc,
		Long:        longDesc,
		Annotations: GetDescriptions([]string{shortDesc, longDesc}, (os.Getenv("GDBASE_HIDEBANNER") == "true")),
		Run: func(cmd *cobra.Command, args []string) {
			//dkr, dkrErr := factory.NewDockerService(nil, l.GetLogger("GodoBase"))
			//if dkrErr != nil {
			//	fmt.Printf("Error stopping Docker service: %v\n", dkrErr)
			//	return
			//}
			//dkrErr = dkr
			//if dkrErr != nil {
			//	fmt.Printf("Error stopping Docker service: %v\n", dkrErr)
			//	return
			//}
			// return
		},
	}
	return cmd
}

func statusDockerCmd() *cobra.Command {

	shortDesc := "Status Docker"
	longDesc := "Status Docker service"

	cmd := &cobra.Command{
		Use:         "status",
		Short:       shortDesc,
		Long:        longDesc,
		Annotations: GetDescriptions([]string{shortDesc, longDesc}, (os.Getenv("GDBASE_HIDEBANNER") == "true")),
		Run: func(cmd *cobra.Command, args []string) {
			//dkr, dkrErr := factory.NewDockerService(nil, l.GetLogger("GodoBase"))
			//if dkrErr != nil {
			//	fmt.Printf("Error getting Docker status: %v\n", dkrErr)
			//	return
			//}
			//dkrErr = dkr.Status()
			//if dkrErr != nil {
			//	fmt.Printf("Error getting Docker status: %v\n", dkrErr)
			//	return
			//}
			_ = cmd.Help()
		},
	}
	return cmd
}
