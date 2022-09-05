package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/google/uuid"
	demo "github.com/vpavlin/temporal-experiment/transaction-check-split"
	"go.temporal.io/sdk/client"
)

func NewMint(path string) (demo.Mint, error) {
	var r demo.Mint

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return r, err
	}

	err = json.Unmarshal(data, &r)
	if err != nil {
		return r, err
	}

	return r, nil
}

func main() {
	// The client is a heavyweight object that should be created once per process.
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	id := uuid.NewString()

	workflowOptions := client.StartWorkflowOptions{
		ID:        fmt.Sprintf("MINT_%s", id),
		TaskQueue: "COLLECTION_QUEUE",
	}

	var params demo.Mint
	if len(os.Args) < 2 {
		log.Fatalln("You need to provide Mint Config")
	}

	params, err = NewMint(os.Args[1])
	if err != nil {
		log.Fatalln("Failed to load mint config", os.Args[1])
	}

	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, demo.Workflow, params)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}

	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())

	var result demo.MintResult
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable to get workflow result", err)
	}

	log.Println("Workflow result: ", result)
}
