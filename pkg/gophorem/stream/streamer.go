package streamer

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/ShiraazMoollatjie/gophorem/pkg/gophorem"
)

type Streamer struct {
	f                 *gophorem.Client
	articleChannel    chan (gophorem.Article)
	articleCache      map[int]time.Time
	articleTimeToLive time.Duration
	mu                sync.Mutex
}

func NewStreamer(client *gophorem.Client) *Streamer {
	s := Streamer{
		f:                 client,
		articleChannel:    make(chan (gophorem.Article)),
		articleCache:      make(map[int]time.Time),
		articleTimeToLive: time.Hour,
	}

	go s.cleanupCache()

	return &s
}

func (s *Streamer) Articles(ctx context.Context, args gophorem.Arguments) chan (gophorem.Article) {
	go func() {
		for {
			al, err := s.f.Articles(ctx, args)
			if err != nil {
				log.Printf("gophorem/streamer: error while streaming articles %+v", err)
			}

			for _, a := range al {
				at := a.Article
				// Articles swap around the tag lists, so this needs to be done.
				at.TagList = a.TagList
				at.Tags = a.Tags

				if !s.exists(at.ID) {
					s.articleChannel <- at
					s.addToCache(at.ID)
				}

			}

			time.Sleep(1 * time.Minute)
		}

	}()

	return s.articleChannel
}

func (s *Streamer) addToCache(id int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.articleCache[id] = time.Now()
}

func (s *Streamer) exists(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.articleCache[id]
	return ok
}

func (s *Streamer) removeFromCache(id int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.articleCache, id)
}

func (s *Streamer) cleanupCache() {
	for {
		for k, v := range s.articleCache {
			if time.Now().Sub(v) > s.articleTimeToLive {
				s.removeFromCache(k)
			}
		}

		time.Sleep(s.articleTimeToLive)
	}
}
