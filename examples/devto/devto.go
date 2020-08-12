package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ShiraazMoollatjie/devtogo/pkg/gophorem"
)

func main() {
	cl := gophorem.NewDevtoClient(gophorem.WithAPIKey("MY_API_KEY"))
	ctx := context.Background()

	// Retrieve all the go articles.
	al, err := cl.Articles(ctx, gophorem.Arguments{
		"tag": "go",
	})
	if err != nil {
		log.Fatalf("something went wrong: %+v", err)
	}

	fmt.Printf("All Articles: %+v", al)
}
