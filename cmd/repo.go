package cmd

import (
	"os"

	"gitprovider/pkg"
	"gitprovider/pkg/pretty"
	"gitprovider/pkg/providers"
)

// This command helps to
// get some information about repository
// gitprovider repo ProjectA/service-api
// {
// 	"id": 132
// 	"name": "service-api"
// 	"hooks": [
// 		"merge": "http://example/mr-hook?token=12345"
// 	]
// }

// get repo
// patch repo
// set hook

// RepoCmd ...
var RepoCmd = &pkg.CliCmd{
	Name: "repo",
	SubCmds: []pkg.CliCmd{
		*RepoGetCmd,
	},
}

// RepoGetOpts ...
var RepoGetOpts struct {
	Common CommonOpts `desc:"Global options"`
}

// RepoGetCmd ...
var RepoGetCmd = &pkg.CliCmd{
	Name: "get",
	Opts: &RepoGetOpts,
	Run: func() error {
		gitlabProvider := &providers.Gitlab{
			Server: RepoGetOpts.Common.Server,
			Token:  RepoGetOpts.Common.Token,
		}
		repos, err := gitlabProvider.Repos()
		if err != nil {
			return err
		}
		pretty.JSON(os.Stdout, repos)
		return nil
	},
}
