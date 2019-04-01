package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pyaesone17/blog"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
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

	svc := blog.NewBlogService(viper.GetViper(), logger)
	svc.ListenAndServe()
}
