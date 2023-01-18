/*
Copyright © 2023 Sebastian Meyer HERE sebastian.meyer@meyer-itconsulting.de
*/
package sb

import (
	"github.com/spf13/cobra"
)

// sbCmd represents the sbctl command
var SbCmd = &cobra.Command{
	Use:   "sb",
	Short: "use to interact with a Service Bus Namespace in Azure",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
}
