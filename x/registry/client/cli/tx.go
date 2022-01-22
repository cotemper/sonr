package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/sonr-io/sonr/x/registry/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	listSeparator              = ","
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdRegisterService())
	cmd.AddCommand(CmdRegisterName())
	cmd.AddCommand(CmdAccessName())
	cmd.AddCommand(CmdUpdateName())
	cmd.AddCommand(CmdAccessService())
	cmd.AddCommand(CmdUpdateService())
	cmd.AddCommand(CmdCreateWhois())
	cmd.AddCommand(CmdUpdateWhois())
	cmd.AddCommand(CmdDeleteWhois())
	cmd.AddCommand(CmdCreateWhatis())
	cmd.AddCommand(CmdUpdateWhatis())
	cmd.AddCommand(CmdDeleteWhatis())
	cmd.AddCommand(CmdCreateThereis())
	cmd.AddCommand(CmdUpdateThereis())
	cmd.AddCommand(CmdDeleteThereis())
	// this line is used by starport scaffolding # 1

	return cmd
}
