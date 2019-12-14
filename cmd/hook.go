package cmd

import (
	"encoding/json"
	"fmt"
	"gitprovider/pkg"
	"gitprovider/pkg/pretty"
	"gitprovider/pkg/providers"
	"os"
)

// HookCmd hook command
var HookCmd = &pkg.CliCmd{
	Name: "hook",
	SubCmds: []pkg.CliCmd{
		*HookApplyCmd,
		*HookGetCmd,
	},
}

// HookGetOpts ...
var HookGetOpts struct {
	RepoName string     `desc:"name of project"`
	Common   CommonOpts `desc:"Global options"`
}

// HookGetCmd ...
var HookGetCmd = &pkg.CliCmd{
	Name: "get",
	Opts: &HookGetOpts,
	Run: func() error {
		gitlabProvider := &providers.Gitlab{
			Server: HookGetOpts.Common.Server,
			Token:  HookGetOpts.Common.Token,
		}
		repos, err := gitlabProvider.Repos()
		if err != nil {
			return err
		}
		var targetRepo *providers.Repo
		for _, repo := range repos {
			if repo.FullName == HookGetOpts.RepoName {
				targetRepo = &repo
				break
			}
		}
		if targetRepo == nil {
			return fmt.Errorf("repo '%s' not found", HookGetOpts.RepoName)
		}
		hooks, err := gitlabProvider.Hooks(*targetRepo)
		if err != nil {
			return err
		}
		pretty.JSON(os.Stdout, hooks)
		return nil
	},
}

// HookApplyOpts ...
var HookApplyOpts struct {
	FileName string     `desc:"that contains the hook to apply"`
	RepoName string     `desc:"name of project"`
	Common   CommonOpts `desc:"Global options"`
}

// HookApplyCmd ...
var HookApplyCmd = &pkg.CliCmd{
	Name: "apply",
	Opts: &HookApplyOpts,
	Run: func() error {
		fmt.Printf("apply hook from '%s'\n", HookApplyOpts.FileName)
		hooks := []providers.Hook{}
		if err := json.NewDecoder(os.Stdin).Decode(&hooks); err != nil {
			return err
		}
		pretty.JSON(os.Stdout, hooks)
		return nil
	},
}
