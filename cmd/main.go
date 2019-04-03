package main

import (
	"log"
	"os"

	"github.com/pyaesone17/blog/boot"
)

var version string

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	svc := boot.NewBlogService(dir)
	svc.Boot()
	svc.ListenAndServe()
}
