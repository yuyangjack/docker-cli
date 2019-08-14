package builder

import (
	"github.com/spf13/cobra"

	"github.com/yuyangjack/dockercli/cli"
	"github.com/yuyangjack/dockercli/cli/command"
)

// NewBuilderCommand returns a cobra command for `builder` subcommands
func NewBuilderCommand(dockerCli command.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "builder",
		Short: "Manage builds",
		Args:  cli.NoArgs,
		RunE:  command.ShowHelp(dockerCli.Err()),
	}
	cmd.AddCommand(
		NewPruneCommand(dockerCli),
	)
	return cmd
}
