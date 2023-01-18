/*
Copyright © 2023 Sebastian Meyer HERE sebastian.meyer@meyer-itconsulting.de
*/
package sb

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"github.com/spf13/cobra"
)

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "send a message to a queue",
	Long:  `example: sb send -c <connectionstring> -q <queue> -m <message>`,
	Run: func(cmd *cobra.Command, args []string) {
		send()
	},
}

var (
	connectionstring string
	queue            string
	message          string
	contentType string
)

func init() {

	sendCmd.Flags().StringVarP(&connectionstring, "connectionstring", "c", "", "The connection string to the Service Bus Namespace")
	sendCmd.Flags().StringVarP(&queue, "queue", "q", "", "The name of the queue to send the message to")
	sendCmd.Flags().StringVarP(&message, "message", "m", "", "The message to send to the queue. '{\"key\":\"value\"}'")
	sendCmd.Flags().StringVarP(&contentType, "contenttype", "t", "application/json", "The content type of the message.")

	if err := sendCmd.MarkFlagRequired("connectionstring"); err != nil {
		panic(err)
	}

	if err := sendCmd.MarkFlagRequired("queue"); err != nil {
		panic(err)
	}

	if err := sendCmd.MarkFlagRequired("message"); err != nil {
		panic(err)
	}

	SbCmd.AddCommand(sendCmd)
}

func send() {
	client, err := getClient()

	if err != nil {
		panic("Could not get client")
	}

	sender, err := client.NewSender(queue, nil)
	if err != nil {
		panic("Could not get sender")
	}

	sbMessage := &azservicebus.Message{Body: []byte(message)}
	sbMessage.ContentType = &contentType

	err = sender.SendMessage(context.TODO(), sbMessage, nil)
	if err != nil {
		panic("Could not send message")
	}

}

func getClient() (*azservicebus.Client, error) {
	return azservicebus.NewClientFromConnectionString(connectionstring, nil)
}
