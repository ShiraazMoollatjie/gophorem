package main

import (
	"fmt"
	"log"

	"github.com/ShiraazMoollatjie/devtogo/pkg/gophorem"
)

func main() {
	cl := gophorem.NewDevtoClient(gophorem.WithApiKey("MY_API_KEY"))
	al, err := cl.Articles(gophorem.Defaults())
	if err != nil {
		log.Fatalf("something went wrong: %+v", err)
	}

	fmt.Printf("All Articles: %+v", al)

	a, err := cl.PublishedArticle(416009)
	if err != nil {
		log.Fatalf("something went wrong: %+v", err)
	}

	fmt.Printf("Article: %+v", a.ID)
}
