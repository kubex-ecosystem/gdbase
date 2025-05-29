package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/rafa-mori/gdbase/cmd/cli"
	gl "github.com/rafa-mori/gdbase/logger"
	"github.com/rafa-mori/gdbase/version"
	"github.com/spf13/cobra"
)

type GDBase struct {
	parentCmdName string
	printBanner   bool
	certPath      string
	keyPath       string
	configPath    string
}

func (m *GDBase) Alias() string {
	return ""
}
func (m *GDBase) ShortDescription() string {
	return "GDBase: GKBX Database and Docker manager/service. "
}
func (m *GDBase) LongDescription() string {
	return `GDBase: Is a tool to manage GKBX database and Docker services. It provides many DB flavors like MySQL, PostgreSQL, MongoDB, Redis, etc. It also provides Docker services like Docker Swarm, Docker Compose, etc. It is a command line tool that can be used to manage GKBX database and Docker services.`
}
func (m *GDBase) Usage() string {
	return "gdbase [command] [args]"
}
func (m *GDBase) Examples() []string {
	return []string{"gdbase [command] [args]", "gdbase database user auth'", "gdbase db roles list"}
}
func (m *GDBase) Active() bool {
	return true
}
func (m *GDBase) Module() string {
	return "gdbase"
}
func (m *GDBase) Execute() error {
	dbChanData := make(chan interface{})
	defer close(dbChanData)

	if spyderErr := m.Command().Execute(); spyderErr != nil {
		gl.Log("error", spyderErr.Error())
		return spyderErr
	} else {
		return nil
	}
}
func (m *GDBase) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use: m.Module(),
		//Aliases:     []string{m.Alias(), "w", "wb", "webServer", "http"},
		Example:     m.concatenateExamples(),
		Annotations: m.getDescriptions(nil, true),
		Version:     version.GetVersion(),
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}

	cmd.AddCommand(version.CliCommand())

	cmd.AddCommand(cli.DockerCmd())

	cmd.AddCommand(cli.DatabaseCmd())

	cmd.AddCommand(cli.UtilsCmds())

	cmd.AddCommand(cli.SshCmds())

	setUsageDefinition(cmd)

	for _, c := range cmd.Commands() {
		setUsageDefinition(c)
		if !strings.Contains(strings.Join(os.Args, " "), c.Use) {
			if c.Short == "" {
				c.Short = c.Annotations["description"]
			}
		}
	}

	return cmd
}
func (m *GDBase) preRunEMethod(cmd *cobra.Command, args []string) error {
	gl.Log("debug", fmt.Sprintf("PreRunE: %s", cmd.Name()))

	return nil
}
func (m *GDBase) getDescriptions(descriptionArg []string, _ bool) map[string]string {
	return cli.GetDescriptions(descriptionArg, m.printBanner)
}
func (m *GDBase) SetParentCmdName(rtCmd string) {
	m.parentCmdName = rtCmd
}
func (m *GDBase) concatenateExamples() string {
	examples := ""
	rtCmd := m.parentCmdName
	if rtCmd != "" {
		rtCmd = rtCmd + " "
	}
	for _, example := range m.Examples() {
		examples += rtCmd + example + "\n  "
	}
	return examples
}

func RegX() *GDBase {
	var configPath = os.Getenv("GODOBASE_CONFIGFILE")
	var keyPath = os.Getenv("GODOBASE_KEYFILE")
	var certPath = os.Getenv("GODOBASE_CERTFILE")
	var printBannerV = os.Getenv("GODOBASE_PRINTBANNER")
	if printBannerV == "" {
		printBannerV = "true"
	}

	return &GDBase{
		configPath:  configPath,
		keyPath:     keyPath,
		certPath:    certPath,
		printBanner: strings.ToLower(printBannerV) == "true",
	}
}
