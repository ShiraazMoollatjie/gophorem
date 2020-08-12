package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ShiraazMoollatjie/devtogo/pkg/gophorem"
)

func main() {
	cl := gophorem.NewClient("https://forem.dev/api", gophorem.WithAPIKey("MY_API_KEY"))
	ctx := context.Background()
	al, err := cl.Articles(ctx, gophorem.Defaults())
	if err != nil {
		log.Fatalf("something went wrong: %+v", err)
	}

	fmt.Printf("All Articles: %+v", al)
}
