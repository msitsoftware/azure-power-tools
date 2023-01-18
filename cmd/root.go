/*
Copyright © 2023 Sebastian Meyer
*/
package cmd

import (
	"msitsoftware/azure-power-tools/cmd/cosmosdb"
	"msitsoftware/azure-power-tools/cmd/sb"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "azure-power-tools",
	Short: "CLI to boost developer productivity",
	Long:  `This CLI is a collection of tools to boost developer productivity.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addSubCommands() {
	rootCmd.AddCommand(sb.SbCmd)
	rootCmd.AddCommand(cosmosdb.CosmosdbCmd)
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addSubCommands()
}
