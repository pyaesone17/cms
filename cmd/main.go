package main

import (
	"log"
	"os"

	"github.com/pyaesone17/blog"
)

var version string

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	svc := blog.NewBlogService()
	svc.Boot(dir)
	svc.ListenAndServe()
}
