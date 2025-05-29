package cli

import (
	"fmt"

	l "github.com/faelmori/logz"
	"github.com/rafa-mori/gdbase/factory"
	s "github.com/rafa-mori/gdbase/internal/services"
	"github.com/spf13/cobra"
)

func DockerCmd() *cobra.Command {
	var configFile string
	short := "Docker management commands for GodoBase"
	long := "Docker management commands for GodoBase"
	cmd := &cobra.Command{
		Use:   "docker",
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

	cmd.AddCommand(startDockerCmd())

	cmd.AddCommand(stopDockerCmd())

	cmd.AddCommand(statusDockerCmd())

	return cmd
}

func startDockerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start Docker",
		Long:  "Start Docker service",
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
	cmd := &cobra.Command{
		Use:   "stop",
		Short: "Stop Docker",
		Long:  "Stop Docker service",
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
			return
		},
	}
	return cmd
}

func statusDockerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "status",
		Short: "Status Docker",
		Long:  "Status Docker service",
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
