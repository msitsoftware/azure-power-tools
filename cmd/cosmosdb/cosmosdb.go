/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cosmosdb

import (
	"github.com/spf13/cobra"
)

// cosmosdbCmd represents the cosmosdb command
var CosmosdbCmd = &cobra.Command{
	Use:   "cosmosdb",
	Short: "Use to interact with a Cosmos DB in Azure",
	Long:  `With this command you can interact with a Cosmos DB in Azure.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
}
