package main

import (
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/sonr-io/sonr/app"
	cmd "github.com/tendermint/spm/cosmoscmd"
)

var isCLI bool = false
var isHighway bool = false
var isDashboard bool = false
var isSonrd bool = true

func main() {
	rootCmd, _ := cmd.NewRootCmd(
		app.Name,
		app.AccountAddressPrefix,
		app.DefaultNodeHome,
		app.Name,
		app.ModuleBasics,
		app.New,
		cmd.WithEnvPrefix("SONR_CHAIN"),
		// this line is used by starport scaffolding # root/arguments

	)
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
