# gophorem
[![Go Report Card](https://goreportcard.com/badge/github.com/ShiraazMoollatjie/gophorem?style=flat-square)](https://goreportcard.com/report/github.com/ShiraazMoollatjie/gophorem)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/ShiraazMoollatjie/gophorem/pkg/gophorem)
[![Build status](https://ci.appveyor.com/api/projects/status/qiyndko2krd4ltep?svg=true)](https://ci.appveyor.com/project/ShiraazMoollatjie/gophorem)

![gophorem](gophorem_logo.png)

gophorem is a REST API Wrapper for a forem api written in go.

# Usage
Import the package into your go file:

```go
import "github.com/ShiraazMoollatjie/gophorem/pkg/gophorem"
```

Thereafter, create a client and specify your API token:
```go
cl := gophorem.NewClient(gophorem.WithApiKey("MY_API_KEY"))
```

It is also possible to not use an API key for anonymous operations.

# Getting started

To retrieve a list of articles, simply use the `GetArticles` function:
```go
articles, err := cl.GetArticles(gophorem.Defaults())
```
It is also possible for us to add query parameters. For example, it's useful to retrieve articles for a specific `tag`.
The way to do this would be:
```go
al, err := cl.GetArticles(gophorem.Arguments{
		"tag": "go",
	})
```

To retrieve a single article, you need to specify the `article id`:
```go
article, err := client.GetArticle("167919")
```

# Features

* Authentication via API Key
* Convenience functions to create a devto client
* The entire forem api set of endpoints including:
  * Article management
  * Listing management
  * Webhook management
  * Listing podcasts, video articles,

# More examples

See the `examples` package for more examples on forem api usages.

# Reference
https://docs.dev.to/api/