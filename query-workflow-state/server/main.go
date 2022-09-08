package main

import (
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	queryworkflowstate "github.com/vpavlin/temporal-experiment/query-workflow-state"
	"go.temporal.io/sdk/client"
)

func main() {
	e := echo.New()

	e.Use(
		middleware.Logger(), // Log everything to stdout
	)

	temporal, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer temporal.Close()

	e.GET("/mint/:id/status", func(c echo.Context) error {
		id := c.Param("id")

		response, err := temporal.QueryWorkflow(context.Background(), id, "", queryworkflowstate.QueryType)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		var responseBody interface{}

		err = response.Get(&responseBody)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, responseBody)
	})

	log.Fatal(e.Start("localhost:8888"))
}
