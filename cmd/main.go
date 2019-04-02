package main

import (
	"github.com/pyaesone17/blog"
)

var version string

func main() {
	svc := blog.NewBlogService()
	svc.Boot()
	svc.ListenAndServe()
}
