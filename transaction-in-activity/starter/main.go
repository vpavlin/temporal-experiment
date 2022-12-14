package main

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	demo "github.com/vpavlin/temporal-experiment/transaction-in-activity"
	"go.temporal.io/sdk/client"
)

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

	params := demo.Mint{
		Contract: "0xE0c3698dEe722e2FccD5cf4582007F700205De6e",
		To:       "0x90e922b6D839335b5F12a1ECe75b5E04EbCC23Ef", //TODO: Fail for 0x93A36A162Ab993C678F8ff247aA90f003F96eAB1
		Quantity: 1,
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
