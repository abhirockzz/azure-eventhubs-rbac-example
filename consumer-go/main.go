package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	eventhub "github.com/Azure/azure-event-hubs-go/v3"
)

func main() {

	ehnamespace := os.Getenv("EVENTHUBS_NAMESPACE")
	ehname := os.Getenv("EVENTHUB_NAME")

	hub, err := eventhub.NewHubWithNamespaceNameAndEnvironment(ehnamespace, ehname)
	if err != nil {
		log.Fatal("failed to create EH client ", err)
	}
	fmt.Println("crwated EH client")

	info, err := hub.GetRuntimeInformation(context.Background())
	if err != nil {
		log.Fatal("failed to get runtime info ", err)
	}

	fmt.Println("got runtime info")

	for _, pid := range info.PartitionIDs {
		_, err := hub.Receive(context.Background(), pid, func(ctx context.Context, event *eventhub.Event) error {
			log.Println("recieved message from EH - ", string(event.Data))
			return nil
		}, eventhub.ReceiveWithLatestOffset())
		if err != nil {
			log.Println("Receive failed ", err)
		}
	}

	exit := make(chan os.Signal)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	<-exit
	fmt.Println("bye...")
}
