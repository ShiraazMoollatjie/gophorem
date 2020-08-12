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
	al, err := cl.Articles(ctx, gophorem.Defaults())
	if err != nil {
		log.Fatalf("something went wrong: %+v", err)
	}

	fmt.Printf("All Articles: %+v", al)

	a, err := cl.PublishedArticle(ctx, 416009)
	if err != nil {
		log.Fatalf("something went wrong: %+v", err)
	}

	fmt.Printf("Article: %+v", a.ID)
}
