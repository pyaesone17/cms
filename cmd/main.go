package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/pyaesone17/blog"
	"github.com/pyaesone17/blog/internal"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var version string

func main() {

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile | log.Lmicroseconds)

	viper.AddConfigPath("../config") // optionally look for config in the working directory
	err := viper.ReadInConfig()      // Find and read the config file

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

	log.Printf("STARTUP: %s version %s", os.Args[0], version)
	log.Printf("Listening on: %s", viper.GetString("address"))

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	svc := blog.NewBlogService(&internal.App{Config: viper.GetViper(), Log: logger, Db: client})
	svc.ListenAndServe()
}
