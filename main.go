package main

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"test/model"
	"test/route"
	"github.com/tigrisdata/tigris-client-go/tigris"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := tigris.NewClient(ctx, &tigris.Config{Project: "test"})

	if err != nil {
		panic(err)
	}
	defer client.Close()

	db, err := client.OpenDatabase(ctx,
		&model.Unconfirmed{},
	)
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	route.SetupUnconfirmedCRUD[model.Unconfirmed](r, db, "unconfirmed")

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
