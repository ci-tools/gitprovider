package cmd

import (
	"gitprovider/pkg"
)

// CommonOpts ...
type CommonOpts struct {
	Server string `required:"true" desc:"git server endpoint"`
	Token  string `desc:"authentification token"`
}

// Cli container
var Cli = pkg.CliCmd{
	Name: "gitprovider",
	SubCmds: []pkg.CliCmd{
		*RepoCmd,
		*HookCmd,
	},
}

// Execute cli
func Execute() error {
	return Cli.Execute()
}
