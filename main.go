package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/efirs/test/model"
	"github.com/efirs/test/route"

	"github.com/gin-gonic/gin"
	"github.com/tigrisdata/tigris-client-go/tigris"
)

var Branch string // set duing build time from the Git branch

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	fmt.Printf("Tigris: Project=test, Branch=%s URL: %s\n", Branch, os.Getenv("TIGRIS_URL"))

	client, err := tigris.NewClient(ctx, &tigris.Config{Project: "test", Branch: Branch})

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
