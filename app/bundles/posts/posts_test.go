package posts_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/go-chi/chi"
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
		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		svc = blog.NewBlogService()
		svc.Boot(dir + "/../../../")
	})

	Describe("Post Controller", func() {
		Context("Post Controller Create Api", func() {
			It("Does create api work properly", func() {
				con := posts.NewController(svc.App)
				ts := httptest.NewServer(http.HandlerFunc(con.Create))
				defer ts.Close()

				var jsonStr = []byte(`{"title": "nice", "content": "content"}`)
				res, err := http.Post(ts.URL, "application/json", bytes.NewBuffer(jsonStr))
				if err != nil {
					log.Fatal(err)
				}
				result, err := ioutil.ReadAll(res.Body)
				res.Body.Close()
				if err != nil {
					log.Fatal(err)
				}

				if res.StatusCode != 200 {
					fmt.Println(string(result))
				}

				Expect(http.StatusOK).To(Equal(res.StatusCode))
			})

			It("Does api throw validation error", func() {
				con := posts.NewController(svc.App)
				ts := httptest.NewServer(http.HandlerFunc(con.Create))
				defer ts.Close()

				var jsonStr = []byte(`{"title": "nice"}`)
				res, err := http.Post(ts.URL, "application/json", bytes.NewBuffer(jsonStr))
				if err != nil {
					log.Fatal(err)
				}
				_, err = ioutil.ReadAll(res.Body)
				res.Body.Close()
				if err != nil {
					log.Fatal(err)
				}

				Expect(http.StatusUnprocessableEntity).To(Equal(res.StatusCode))
			})
		})

		Context("Post Controller Show Api", func() {
			It("should say created correctly", func() {
				type Post struct{ ID string }
				type Model struct {
					Data Post `json:"data"`
				}

				con := posts.NewController(svc.App)
				ts := httptest.NewServer(http.HandlerFunc(con.Create))
				defer ts.Close()

				var jsonStr = []byte(`{"title": "nice", "content": "content"}`)
				res, err := http.Post(ts.URL, "application/json", bytes.NewBuffer(jsonStr))
				if err != nil {
					log.Fatal(err)
				}

				result, err := ioutil.ReadAll(res.Body)
				res.Body.Close()

				var model1 Model
				json.Unmarshal(result, &model1)

				m := chi.NewRouter()
				m.Get("/blogs/{post_id}", con.Show)

				ts = httptest.NewServer(m)
				defer ts.Close()

				res, err = http.Get(ts.URL + "/blogs/" + model1.Data.ID)
				result, err = ioutil.ReadAll(res.Body)
				res.Body.Close()

				if err != nil {
					log.Fatal(err)
				}

				var model2 Model
				json.Unmarshal(result, &model2)

				Expect(http.StatusOK).To(Equal(res.StatusCode))
				Expect(model1.Data.ID).To(Equal(model2.Data.ID))
			})
		})

		Context("Post Controller Get Api", func() {
			It("should say created correctly", func() {
				con := posts.NewController(svc.App)
				ts := httptest.NewServer(http.HandlerFunc(con.Get))
				defer ts.Close()

				var jsonStr = []byte(`{"title": "nice", "content": "content"}`)
				res, err := http.Post(ts.URL, "application/json", bytes.NewBuffer(jsonStr))
				if err != nil {
					log.Fatal(err)
				}
				result, err := ioutil.ReadAll(res.Body)
				res.Body.Close()
				if err != nil {
					log.Fatal(err)
				}

				type Post struct {
					ID string
				}

				type Model struct {
					Data []Post `json:"data"`
				}
				var model Model

				json.Unmarshal(result, &model)

				Expect(http.StatusOK).To(Equal(res.StatusCode))
				Expect(model.Data).ShouldNot(HaveLen(0))
			})
		})
	})
})
