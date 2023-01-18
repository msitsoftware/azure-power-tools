/*
Copyright © 2023 Sebastian Meyer
*/
package cmd

import (
	"msitsoftware/azure-power-tools/cmd/sb"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "azure-power-tools",
	Short: "CLI to boost developer productivity",
	Long: `This CLI is a collection of tools to boost developer productivity.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addSubCommands() {
	rootCmd.AddCommand(sb.SbCmd)
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.azure-power-tools.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addSubCommands()
}
