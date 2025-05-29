package cli

import (
	"log"
	"os/exec"
	"strings"

	"github.com/rafa-mori/gdbase/utils"
	"github.com/spf13/cobra"
)

// SshCmdsList retorna uma lista de comandos Cobra relacionados a SSH.
// Retorna um slice de ponteiros para comandos Cobra.
func SshCmds() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:     "ssh",
		Aliases: []string{"s", "ss"},
		Short:   "Configura os utilitários SSH do sistema",
		Long:    "Configura os utilitários SSH do sistema",
	}

	rootCmd.AddCommand(sshTunnelCmd())
	rootCmd.AddCommand(sshTunnelServiceCmd())

	return rootCmd
}

// sshTunnelCmd cria um comando Cobra para configurar um túnel SSH.
// Retorna um ponteiro para o comando Cobra configurado.
func sshTunnelCmd() *cobra.Command {
	var sshUser, sshCert, sshPassword, sshAddress, sshPort string
	var tunnels []string
	var background bool

	rootCmd := &cobra.Command{
		Use:     "tunnel",
		Aliases: []string{"tun", "t"},
		Short:   "Configura um túnel SSH",
		RunE: func(cmd *cobra.Command, args []string) error {
			if background {
				sshCmdRun := exec.Command("kbx", "u", "s", "tunnel-service-background", "--sshUser", sshUser, "--sshCert", sshCert, "--sshPassword", sshPassword, "--sshAddress", sshAddress, "--sshPort", sshPort, "--tunnels", strings.Join(tunnels, ","))
				sshCmdRunErr := sshCmdRun.Start()
				if sshCmdRunErr != nil {
					log.Println("Erro ao iniciar o serviço de túnel SSH:", sshCmdRunErr)
					return nil
				}
				//processReleaseErr := sshCmdRun.Process.Release()
				//if processReleaseErr != nil {
				//	log.Println("Erro ao liberar o processo do serviço de túnel SSH:", processReleaseErr)
				//	return nil
				//}
				log.Println("Serviço de túnel SSH iniciado em segundo plano")
				return nil
			}

			return utils.SshTunnel(sshUser, sshCert, sshPassword, sshAddress, sshPort, tunnels...)
		},
	}

	rootCmd.Flags().BoolVarP(&background, "background", "b", false, "Executar em segundo plano")
	rootCmd.Flags().StringVarP(&sshUser, "login", "l", "", "Usuário SSH")
	rootCmd.Flags().StringVarP(&sshCert, "cert", "i", "", "Certificado SSH")
	rootCmd.Flags().StringVarP(&sshPassword, "secret", "s", "", "Senha SSH")
	rootCmd.Flags().StringVarP(&sshAddress, "host", "t", "", "Endereço SSH")
	rootCmd.Flags().StringVarP(&sshPort, "port", "p", "", "Porta SSH")
	rootCmd.Flags().StringSliceVarP(&tunnels, "tunnels", "L", []string{}, "Túneis")

	return rootCmd
}

// sshTunnelServiceCmd cria um comando Cobra para configurar um serviço de túnel SSH em segundo plano.
// Retorna um ponteiro para o comando Cobra configurado.
func sshTunnelServiceCmd() *cobra.Command {
	var sshUser, sshCert, sshPassword, sshAddress, sshPort string
	var tunnels []string
	rootCmd := &cobra.Command{
		Use:    "tunnel-service-background",
		Hidden: true,
		Run: func(cmd *cobra.Command, args []string) {
			_ = utils.SshTunnel(sshUser, sshCert, sshPassword, sshAddress, sshPort, tunnels...)
		},
	}
	rootCmd.Flags().StringVarP(&sshUser, "sshUser", "l", "", "Usuário SSH")
	rootCmd.Flags().StringVarP(&sshCert, "sshCert", "i", "", "Certificado SSH")
	rootCmd.Flags().StringVarP(&sshPassword, "sshPassword", "s", "", "Senha SSH")
	rootCmd.Flags().StringVarP(&sshAddress, "sshAddress", "t", "", "Endereço SSH")
	rootCmd.Flags().StringVarP(&sshPort, "sshPort", "p", "", "Porta SSH")
	rootCmd.Flags().StringSliceVarP(&tunnels, "tunnels", "L", []string{}, "Túneis")
	return rootCmd
}
