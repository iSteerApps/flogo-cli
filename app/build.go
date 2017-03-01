package app

import (
	"flag"

	"github.com/TIBCOSoftware/flogo-cli/cli"
	"os"
	"fmt"
)

var optBuild = &cli.OptionInfo{
	Name:      "build",
	UsageLine: "build [-o][-i]",
	Short:     "build the flogo application",
	Long: `Build the flogo application.

Options:
    -o   optimize for embedded flows
    -i   incorporate config into application
`,
}

func init() {
	CommandRegistry.RegisterCommand(&cmdBuild{option: optBuild})
}

type cmdBuild struct {
	option     *cli.OptionInfo
	optimize   bool
	includeCfg bool
}

// HasOptionInfo implementation of cli.HasOptionInfo.OptionInfo
func (c *cmdBuild) OptionInfo() *cli.OptionInfo {
	return c.option
}

// AddFlags implementation of cli.Command.AddFlags
func (c *cmdBuild) AddFlags(fs *flag.FlagSet) {
	fs.BoolVar(&(c.optimize), "o", false, "optimize build")
	fs.BoolVar(&(c.includeCfg), "i", false, "include config")
}

// Exec implementation of cli.Command.Exec
func (c *cmdBuild) Exec(args []string) error {

	appDir, err := os.Getwd()

	if err != nil {
		fmt.Fprint(os.Stderr, "Error: Unable to determine working directory\n\n")
		os.Exit(2)
	}

	return BuildApp(SetupExistingProjectEnv(appDir), nil)
}
