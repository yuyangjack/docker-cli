package engine

import (
	"github.com/yuyangjack/dockercli/cli"
	"github.com/yuyangjack/dockercli/cli/command"
	"github.com/spf13/cobra"
)

// NewEngineCommand returns a cobra command for `engine` subcommands
func NewEngineCommand(dockerCli command.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "engine COMMAND",
		Short: "Manage the docker engine",
		Args:  cli.NoArgs,
		RunE:  command.ShowHelp(dockerCli.Err()),
	}
	cmd.AddCommand(
		newActivateCommand(dockerCli),
		newCheckForUpdatesCommand(dockerCli),
		newUpdateCommand(dockerCli),
	)
	return cmd
}
