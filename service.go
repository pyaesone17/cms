package blog

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/pyaesone17/blog/app/bundles/posts"
	"github.com/pyaesone17/blog/app/core"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// BlogService holds
type BlogService struct {
	config *viper.Viper
	srv    *http.Server
	log    *logrus.Logger
}

// NewBlogService is the constructor
func NewBlogService(viper *viper.Viper, logrus *logrus.Logger) *BlogService {
	return &BlogService{
		config: viper,
		log:    logrus,
	}
}

// ListenAndServe will listen the port
func (s *BlogService) ListenAndServe() {
	r := chi.NewRouter()
	for _, b := range registerBundles(s.config, s.log) {
		for _, route := range b.GetRoutes() {
			log.Printf("adding handler for \"%s %s\"", route.Method, route.Path)
			r.Method(route.Method, route.Path, route.Handler)
		}
	}

	http.ListenAndServe(s.config.GetString("address"), r)
	return
}

// Stop will stop running the server
func (s *BlogService) Stop() {
	if s.srv != nil {
		s.srv.Close()
	}
}

func registerBundles(config *viper.Viper, log *logrus.Logger) []core.Bundle {
	return []core.Bundle{
		posts.NewBundle(config, log),
	}
}
