package commands

import (
	"os"
	"runtime"

	"github.com/yuyangjack/dockercli/cli/command"
	"github.com/yuyangjack/dockercli/cli/command/builder"
	"github.com/yuyangjack/dockercli/cli/command/checkpoint"
	"github.com/yuyangjack/dockercli/cli/command/config"
	"github.com/yuyangjack/dockercli/cli/command/container"
	"github.com/yuyangjack/dockercli/cli/command/engine"
	"github.com/yuyangjack/dockercli/cli/command/image"
	"github.com/yuyangjack/dockercli/cli/command/manifest"
	"github.com/yuyangjack/dockercli/cli/command/network"
	"github.com/yuyangjack/dockercli/cli/command/node"
	"github.com/yuyangjack/dockercli/cli/command/plugin"
	"github.com/yuyangjack/dockercli/cli/command/registry"
	"github.com/yuyangjack/dockercli/cli/command/secret"
	"github.com/yuyangjack/dockercli/cli/command/service"
	"github.com/yuyangjack/dockercli/cli/command/stack"
	"github.com/yuyangjack/dockercli/cli/command/swarm"
	"github.com/yuyangjack/dockercli/cli/command/system"
	"github.com/yuyangjack/dockercli/cli/command/trust"
	"github.com/yuyangjack/dockercli/cli/command/volume"
	"github.com/spf13/cobra"
)

// AddCommands adds all the commands from cli/command to the root command
func AddCommands(cmd *cobra.Command, dockerCli command.Cli) {
	cmd.AddCommand(
		// checkpoint
		checkpoint.NewCheckpointCommand(dockerCli),

		// config
		config.NewConfigCommand(dockerCli),

		// container
		container.NewContainerCommand(dockerCli),
		container.NewRunCommand(dockerCli),

		// image
		image.NewImageCommand(dockerCli),
		image.NewBuildCommand(dockerCli),

		// builder
		builder.NewBuilderCommand(dockerCli),

		// manifest
		manifest.NewManifestCommand(dockerCli),

		// network
		network.NewNetworkCommand(dockerCli),

		// node
		node.NewNodeCommand(dockerCli),

		// plugin
		plugin.NewPluginCommand(dockerCli),

		// registry
		registry.NewLoginCommand(dockerCli),
		registry.NewLogoutCommand(dockerCli),
		registry.NewSearchCommand(dockerCli),

		// secret
		secret.NewSecretCommand(dockerCli),

		// service
		service.NewServiceCommand(dockerCli),

		// system
		system.NewSystemCommand(dockerCli),
		system.NewVersionCommand(dockerCli),

		// stack
		stack.NewStackCommand(dockerCli),
		stack.NewTopLevelDeployCommand(dockerCli),

		// swarm
		swarm.NewSwarmCommand(dockerCli),

		// trust
		trust.NewTrustCommand(dockerCli),

		// volume
		volume.NewVolumeCommand(dockerCli),

		// legacy commands may be hidden
		hide(system.NewEventsCommand(dockerCli)),
		hide(system.NewInfoCommand(dockerCli)),
		hide(system.NewInspectCommand(dockerCli)),
		hide(container.NewAttachCommand(dockerCli)),
		hide(container.NewCommitCommand(dockerCli)),
		hide(container.NewCopyCommand(dockerCli)),
		hide(container.NewCreateCommand(dockerCli)),
		hide(container.NewDiffCommand(dockerCli)),
		hide(container.NewExecCommand(dockerCli)),
		hide(container.NewExportCommand(dockerCli)),
		hide(container.NewKillCommand(dockerCli)),
		hide(container.NewLogsCommand(dockerCli)),
		hide(container.NewPauseCommand(dockerCli)),
		hide(container.NewPortCommand(dockerCli)),
		hide(container.NewPsCommand(dockerCli)),
		hide(container.NewRenameCommand(dockerCli)),
		hide(container.NewRestartCommand(dockerCli)),
		hide(container.NewRmCommand(dockerCli)),
		hide(container.NewStartCommand(dockerCli)),
		hide(container.NewStatsCommand(dockerCli)),
		hide(container.NewStopCommand(dockerCli)),
		hide(container.NewTopCommand(dockerCli)),
		hide(container.NewUnpauseCommand(dockerCli)),
		hide(container.NewUpdateCommand(dockerCli)),
		hide(container.NewWaitCommand(dockerCli)),
		hide(image.NewHistoryCommand(dockerCli)),
		hide(image.NewImagesCommand(dockerCli)),
		hide(image.NewImportCommand(dockerCli)),
		hide(image.NewLoadCommand(dockerCli)),
		hide(image.NewPullCommand(dockerCli)),
		hide(image.NewPushCommand(dockerCli)),
		hide(image.NewRemoveCommand(dockerCli)),
		hide(image.NewSaveCommand(dockerCli)),
		hide(image.NewTagCommand(dockerCli)),
	)
	if runtime.GOOS == "linux" {
		// engine
		cmd.AddCommand(engine.NewEngineCommand(dockerCli))
	}
}

func hide(cmd *cobra.Command) *cobra.Command {
	// If the environment variable with name "DOCKER_HIDE_LEGACY_COMMANDS" is not empty,
	// these legacy commands (such as `docker ps`, `docker exec`, etc)
	// will not be shown in output console.
	if os.Getenv("DOCKER_HIDE_LEGACY_COMMANDS") == "" {
		return cmd
	}
	cmdCopy := *cmd
	cmdCopy.Hidden = true
	cmdCopy.Aliases = []string{}
	return &cmdCopy
}
