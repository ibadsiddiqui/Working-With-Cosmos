package client

import (
	"github.com/cosmos/cosmos-sdk/client"
	nameservicecmd "github.com/ibadsiddiqui/Working-With-Cosmos/x/nameservice/client/cli"
	"github.com/spf13/cobra"
	amino "github.com/tendermint/go-amino"
)

// ModuleClient exoirts all client functionality from this module
type ModuleClient struct {
	storekey string
	cdc *amino.Codec
}

func NewModuleClient(storekey string, cdc *amino.Codec) ModuleClient {
	return ModuleClient(storekey, cdc)
}

// GetQueryCmd - returns the cli commands for this module
func (mc ModuleClient) GetQueryCmd() *cobra.Command  {

	// Group gov queries under a subcommand
	govQueryCmd = &cobra.Command {
		Use: 	"nameservice",
		Short: 	"Querying commands for the nameservice module",
	}

	govQueryCmd.AddCommand(client.GetCommands(
		nameservicecmd.GetCmdResolveName(mc.storekey, mc.cdc),
		nameservicecmd.GetCmdWhois(mc.storekey, mc.cdc)
	)...)
	return govQueryCmd
}


// GetTxCmd - returns the transaction commands for this module
func (mc ModuleClient) GetTxCmd() *cobra.Command  {
	govTxCmd := &cobra.Command {
		Use: 	"nameservice",
		Short:	"Nameservice transactions subcommands"
	}
	govTxCmd.AddCommand(client.PostCommands(
		nameservicecmd.GetCmdBuyName(mc.cdc),
		nameservicecmd.GetCmdSetName(mc.cdc)
	)...)
	return govTxCmd
}