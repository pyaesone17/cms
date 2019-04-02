package posts_test

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/pyaesone17/blog"
	posts "github.com/pyaesone17/blog/app/bundles/posts"
)

var _ = Describe("Posts", func() {
	var (
		svc *blog.Service
	)

	BeforeEach(func() {
		svc = blog.NewBlogService()
		svc.Boot()
	})

	Describe("Categorizing book length", func() {
		Context("Does api work properly", func() {
			It("should say hi", func() {
				con := posts.NewController(svc.App)
				ts := httptest.NewServer(http.HandlerFunc(con.Create))
				defer ts.Close()

				res, err := http.Get(ts.URL)
				if err != nil {
					log.Fatal(err)
				}
				greeting, err := ioutil.ReadAll(res.Body)
				res.Body.Close()
				if err != nil {
					log.Fatal(err)
				}

				Expect("Hi").To(Equal(string(greeting)))
			})
		})
	})
})
