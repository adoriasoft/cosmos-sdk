package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/simapp/x/nameservice/types"
	"github.com/spf13/cobra"
)

func GetQueryCmd(storeKey string) *cobra.Command {
	nameserviceQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the nameservice module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	nameserviceQueryCmd.AddCommand(
		GetCmdResolveName(storeKey),
		GetCmdWhois(storeKey),
		GetCmdNames(storeKey),
	)

	return nameserviceQueryCmd
}

// GetCmdResolveName queries information about a name
func GetCmdResolveName(queryRoute string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "resolve [name]",
		Short: "resolve name",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			name := args[0]

			res, _, err := clientCtx.QueryWithData(fmt.Sprintf("custom/%s/resolve/%s", queryRoute, name), nil)
			if err != nil {
				return nil
			}
			var out types.QueryResResolve
			clientCtx.JSONMarshaler.MustUnmarshalJSON(res, &out)
			return clientCtx.PrintOutput(out)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdWhois queries information about a domain
func GetCmdWhois(queryRoute string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "whois [name]",
		Short: "Query whois info of name",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}
			name := args[0]

			res, _, err := clientCtx.QueryWithData(fmt.Sprintf("custom/%s/whois/%s", queryRoute, name), nil)
			if err != nil {
				return nil
			}

			var out types.Whois
			clientCtx.JSONMarshaler.MustUnmarshalJSON(res, &out)
			return clientCtx.PrintOutput(out)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdNames queries a list of all names
func GetCmdNames(queryRoute string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "names",
		Short: "names",
		// Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			res, _, err := clientCtx.QueryWithData(fmt.Sprintf("custom/%s/names", queryRoute), nil)
			if err != nil {
				return nil
			}

			var out types.QueryResNames
			clientCtx.JSONMarshaler.MustUnmarshalJSON(res, &out)
			return clientCtx.PrintOutput(out)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
