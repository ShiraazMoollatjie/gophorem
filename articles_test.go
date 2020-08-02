package devtogo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPublishedArticle(t *testing.T) {
	var res Article
	b := unmarshalGoldenFileBytes(t, "article.json", &res)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/articles/167919", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}))
	client := NewClient(withBaseURL(ts.URL))
	article, err := client.PublishedArticle(167919)
	assert.NoError(t, err)
	assert.Equal(t, &res, article)
}

func TestPublishedArticleByPath(t *testing.T) {
	var res Article
	b := unmarshalGoldenFileBytes(t, "article.json", &res)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/articles/devteam/using-go-is-awesome", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}))
	client := NewClient(withBaseURL(ts.URL))
	article, err := client.PublishedArticleByPath("devteam", "using-go-is-awesome")
	assert.NoError(t, err)
	assert.Equal(t, &res, article)
}

func TestGetArticles(t *testing.T) {
	var res Articles
	b := unmarshalGoldenFileBytes(t, "articles.json", &res)

	tests := []struct {
		name                string
		arguments           Arguments
		expectedQueryParams string
	}{
		{"No params", Defaults(), ""},
		{"Page param", Arguments{"page": "1"}, "page=1"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, "/articles?"+test.expectedQueryParams, r.URL.String())
				w.WriteHeader(http.StatusOK)
				w.Write(b)
			}))

			client := NewClient(withBaseURL(ts.URL))
			articles, err := client.Articles(test.arguments)
			assert.NoError(t, err)
			assert.Equal(t, res, articles)
		})
	}
}
func TestGetVideoArticles(t *testing.T) {
	var res VideoArticles
	b := unmarshalGoldenFileBytes(t, "videos.json", &res)

	tests := []struct {
		name                string
		arguments           Arguments
		expectedQueryParams string
	}{
		{"No params", Defaults(), ""},
		{"Page param", Arguments{"page": "1"}, "page=1"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, "/videos?"+test.expectedQueryParams, r.URL.String())
				w.WriteHeader(http.StatusOK)
				w.Write(b)
			}))

			client := NewClient(withBaseURL(ts.URL))
			articles, err := client.VideoArticles(test.arguments)
			assert.NoError(t, err)
			assert.Equal(t, res, articles)
		})
	}
}

func TestGetMyArticles(t *testing.T) {
	var res Articles
	b := unmarshalGoldenFileBytes(t, "articles.json", &res)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/articles/me?", r.URL.String())
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}))

	client := NewClient(withBaseURL(ts.URL))
	articles, err := client.MyArticles(Defaults())
	assert.NoError(t, err)
	assert.Equal(t, res, articles)
}

func TestGetMyPublishedArticles(t *testing.T) {
	var res Articles
	b := unmarshalGoldenFileBytes(t, "articles.json", &res)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/articles/me/published?", r.URL.String())
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}))

	client := NewClient(withBaseURL(ts.URL))
	articles, err := client.MyPublishedArticles(Defaults())
	assert.NoError(t, err)
	assert.Equal(t, res, articles)
}

func TestGetMyUnpublishedArticles(t *testing.T) {
	var res Articles
	b := unmarshalGoldenFileBytes(t, "articles.json", &res)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/articles/me/unpublished?", r.URL.String())
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}))

	client := NewClient(withBaseURL(ts.URL))
	articles, err := client.MyUnpublishedArticles(Defaults())
	assert.NoError(t, err)
	assert.Equal(t, res, articles)
}

func TestGetAllMyArticles(t *testing.T) {
	var res Articles
	b := unmarshalGoldenFileBytes(t, "myarticles.json", &res)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/articles/me/all?", r.URL.String())
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}))

	client := NewClient(withBaseURL(ts.URL))
	articles, err := client.AllMyArticles(Defaults())
	assert.NoError(t, err)
	assert.Equal(t, res, articles)
}

func TestCreateArticle(t *testing.T) {
	var res Article
	b := unmarshalGoldenFileBytes(t, "create_article.json", &res)
	testArticle := CreateArticleReq{
		Tags:         []string{"go", "help"},
		Series:       "api",
		Published:    false,
		BodyMarkdown: "This is some markdown",
		Title:        "My First Post via API",
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/articles", r.URL.Path)
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "myApiKey", r.Header.Get("api-key"))
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

		rb, err := ioutil.ReadAll(r.Body)
		assert.NoError(t, err)

		var car ArticleReq
		assert.NoError(t, json.Unmarshal(rb, &car))
		assert.Equal(t, ArticleReq{Article: testArticle}, car)

		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}))

	client := NewClient(withBaseURL(ts.URL), WithApiKey("myApiKey"))
	articles, err := client.CreateArticle(testArticle)
	assert.NoError(t, err)
	assert.Equal(t, res, articles)
}

func TestCreateArticleNoSeriesField(t *testing.T) {
	var res Article
	b := unmarshalGoldenFileBytes(t, "create_article.json", &res)
	testArticle := CreateArticleReq{
		Tags:         []string{"go", "help"},
		Published:    false,
		BodyMarkdown: "This is some markdown",
		Title:        "My First Post via API",
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/articles", r.URL.Path)
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "myApiKey", r.Header.Get("api-key"))
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

		rb, err := ioutil.ReadAll(r.Body)
		assert.NoError(t, err)

		var car ArticleReq
		assert.NoError(t, json.Unmarshal(rb, &car))
		assert.Equal(t, ArticleReq{Article: testArticle}, car)

		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}))

	client := NewClient(withBaseURL(ts.URL), WithApiKey("myApiKey"))
	articles, err := client.CreateArticle(testArticle)
	assert.NoError(t, err)
	assert.Equal(t, res, articles)
}

func TestUpdateArticle(t *testing.T) {
	var res Article
	b := unmarshalGoldenFileBytes(t, "create_article.json", &res)
	testArticle := CreateArticleReq{
		Tags:         []string{"go", "help"},
		Series:       "api",
		Published:    false,
		BodyMarkdown: "This is some markdown",
		Title:        "My First Post via API",
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/articles/1000", r.URL.Path)
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Equal(t, "myApiKey", r.Header.Get("api-key"))
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

		rb, err := ioutil.ReadAll(r.Body)
		assert.NoError(t, err)

		var car ArticleReq
		assert.NoError(t, json.Unmarshal(rb, &car))
		assert.Equal(t, ArticleReq{Article: testArticle}, car)

		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}))

	client := NewClient(withBaseURL(ts.URL), WithApiKey("myApiKey"))
	articles, err := client.UpdateArticle(1000, testArticle)
	assert.NoError(t, err)
	assert.Equal(t, res, articles)
}

func unmarshalGoldenFileBytes(t *testing.T, filename string, payload interface{}) []byte {
	p := filepath.Join("testdata", filename)
	b, err := ioutil.ReadFile(p)
	assert.NoError(t, err)

	err = json.Unmarshal(b, &payload)
	assert.NoError(t, err)

	return b
}
