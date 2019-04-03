package blog

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/pyaesone17/blog/app/bundles/posts"
	"github.com/pyaesone17/blog/app/core"
	"github.com/pyaesone17/blog/internal"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Service holds
type Service struct {
	App *internal.App
	srv *http.Server
}

// NewBlogService is the constructor
func NewBlogService() *Service {
	return &Service{}
}

func (s *Service) Boot() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile | log.Lmicroseconds)
	viper.AddConfigPath("/Users/nyanwin/go/src/github.com/pyaesone17/blog/config") // optionally look for config in the working directory
	err := viper.ReadInConfig()                                                    // Find and read the config file

	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{}

	file, err := os.OpenFile("logrus.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		logger.Out = file
	} else {
		logrus.Info("Failed to log to file, using default stderr")
	}

	log.Printf("STARTUP: %s version %s", os.Args[0], "1")
	log.Printf("Listening on: %s", viper.GetString("address"))

	clientOptions := options.Client().ApplyURI(viper.GetString("mongodb"))
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	s.App = &internal.App{
		Db:     client,
		Log:    logger,
		Config: viper.GetViper(),
	}
}

// ListenAndServe will listen the port
func (s *Service) ListenAndServe() {
	r := chi.NewRouter()

	for _, b := range registerBundles(s.App) {
		for _, route := range b.GetRoutes() {
			log.Printf("adding handler for \"%s %s\"", route.Method, route.Path)
			r.Method(route.Method, route.Path, route.Handler)
		}
	}

	err := http.ListenAndServe(s.App.Config.GetString("address"), r)
	if err != nil {
		log.Fatal(err)
	}

	return
}

// Stop will stop running the server
func (s *Service) Stop() {
	if s.srv != nil {
		s.srv.Close()
	}
}

func registerBundles(app *internal.App) []core.Bundle {
	return []core.Bundle{
		posts.NewBundle(app),
	}
}
