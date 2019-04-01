package blog

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/pyaesone17/blog/app/bundles/posts"
	"github.com/pyaesone17/blog/app/core"
	"github.com/pyaesone17/blog/internal"
)

// BlogService holds
type BlogService struct {
	App *internal.App
	srv *http.Server
}

// NewBlogService is the constructor
func NewBlogService(app *internal.App) *BlogService {
	return &BlogService{
		App: app,
	}
}

// ListenAndServe will listen the port
func (s *BlogService) ListenAndServe() {
	r := chi.NewRouter()
	for _, b := range registerBundles(s.App) {
		for _, route := range b.GetRoutes() {
			log.Printf("adding handler for \"%s %s\"", route.Method, route.Path)
			r.Method(route.Method, route.Path, route.Handler)
		}
	}

	http.ListenAndServe(s.App.Config.GetString("address"), r)
	return
}

// Stop will stop running the server
func (s *BlogService) Stop() {
	if s.srv != nil {
		s.srv.Close()
	}
}

func registerBundles(app *internal.App) []core.Bundle {
	return []core.Bundle{
		posts.NewBundle(app),
	}
}
