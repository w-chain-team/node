package root

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/w-chain-team/node/command/backup"
	"github.com/w-chain-team/node/command/bridge"
	"github.com/w-chain-team/node/command/genesis"
	"github.com/w-chain-team/node/command/helper"
	"github.com/w-chain-team/node/command/ibft"
	"github.com/w-chain-team/node/command/license"
	"github.com/w-chain-team/node/command/monitor"
	"github.com/w-chain-team/node/command/peers"
	"github.com/w-chain-team/node/command/polybft"
	"github.com/w-chain-team/node/command/polybftsecrets"
	"github.com/w-chain-team/node/command/regenesis"
	"github.com/w-chain-team/node/command/rootchain"
	"github.com/w-chain-team/node/command/secrets"
	"github.com/w-chain-team/node/command/server"
	"github.com/w-chain-team/node/command/status"
	"github.com/w-chain-team/node/command/txpool"
	"github.com/w-chain-team/node/command/version"
)

type RootCommand struct {
	baseCmd *cobra.Command
}

func NewRootCommand() *RootCommand {
	rootCommand := &RootCommand{
		baseCmd: &cobra.Command{
			Short: "Polygon Edge is a framework for building Ethereum-compatible Blockchain networks",
		},
	}

	helper.RegisterJSONOutputFlag(rootCommand.baseCmd)

	rootCommand.registerSubCommands()

	return rootCommand
}

func (rc *RootCommand) registerSubCommands() {
	rc.baseCmd.AddCommand(
		version.GetCommand(),
		txpool.GetCommand(),
		status.GetCommand(),
		secrets.GetCommand(),
		peers.GetCommand(),
		rootchain.GetCommand(),
		monitor.GetCommand(),
		ibft.GetCommand(),
		backup.GetCommand(),
		genesis.GetCommand(),
		server.GetCommand(),
		license.GetCommand(),
		polybftsecrets.GetCommand(),
		polybft.GetCommand(),
		bridge.GetCommand(),
		regenesis.GetCommand(),
	)
}

func (rc *RootCommand) Execute() {
	if err := rc.baseCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}
}
