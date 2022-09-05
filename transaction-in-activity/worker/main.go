package main

import (
	"log"

	"github.com/vpavlin/temporal-experiment/config"
	transactioninactivity "github.com/vpavlin/temporal-experiment/transaction-in-activity"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	config, err := config.NewConfig(".config.json")
	if err != nil {
		log.Fatalln("Failed to load config: ", err)
	}

	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()
	// This worker hosts both Worker and Activity functions
	w := worker.New(c, "COLLECTION_QUEUE", worker.Options{})

	b, err := transactioninactivity.NewBlockchain(config.Rpc, 80001, config.PrivateKey)
	if err != nil {
		log.Fatalln("Failed to instantiate blockchain wrapper: ", err)
	}

	w.RegisterActivity(b.Mint)
	w.RegisterWorkflow(transactioninactivity.Workflow)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}
