package main

import (
	"context"
	"fmt"

	"github.com/ShiraazMoollatjie/gophorem/pkg/gophorem"
	streamer "github.com/ShiraazMoollatjie/gophorem/pkg/gophorem/stream"
)

func main() {
	cl := gophorem.NewDevtoClient(gophorem.WithAPIKey("MY_API_KEY"))
	ctx := context.Background()

	s := streamer.NewStreamer(cl)

	ch := s.Articles(ctx, gophorem.Arguments{
		"state": "fresh",
	})

	for a := range ch {
		fmt.Printf("Received article ID: %d, Title: %s, Username: %s URL: %s \n", a.ID, a.Title, a.User.Username, a.URL)
	}
}
