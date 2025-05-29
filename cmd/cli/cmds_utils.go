package cli

import (
	"github.com/rafa-mori/gdbase/utils/helpers"
	"github.com/spf13/cobra"
)

// SshCmdsList retorna uma lista de comandos Cobra relacionados a SSH.
// Retorna um slice de ponteiros para comandos Cobra.
func UtilsCmds() *cobra.Command {
	uCmd := &cobra.Command{
		Use:     "utils",
		Aliases: []string{"u", "util"},
		Short:   "Configura os utilitários do sistema",
		Long:    "Configura os utilitários do sistema",
	}
	uCmd.AddCommand(installUtilsCmd())
	uCmd.AddCommand(uninstallUtilsCmd())
	return uCmd
}

// sshTunnelServiceCmd cria um comando Cobra para configurar um serviço de túnel SSH em segundo plano.
// Retorna um ponteiro para o comando Cobra configurado.
func installUtilsCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:  "install",
		Long: "Install the bash helpers.",
		Run: func(cmd *cobra.Command, args []string) {
			helpers.InstallBashHelpers()
		},
	}
	return rootCmd
}

// sshTunnelCmd cria um comando Cobra para configurar um túnel SSH.
// Retorna um ponteiro para o comando Cobra configurado.
func uninstallUtilsCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:  "uninstall",
		Long: "Uninstall the bash helpers.",
		Run: func(cmd *cobra.Command, args []string) {
			helpers.UninstallBashHelpers()
		},
	}

	return rootCmd
}
