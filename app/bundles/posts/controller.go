package posts

import (
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/go-chi/chi"
	"github.com/google/jsonapi"
	"github.com/pkg/errors"
	"github.com/pyaesone17/blog/app/core"
	"github.com/pyaesone17/blog/app/models"
	"github.com/pyaesone17/blog/app/services"
	"github.com/pyaesone17/blog/internal"
	"github.com/sirupsen/logrus"
)

type Controller struct {
	*core.Controller
	app         *internal.App
	postservice *services.PostService
}

func NewController(app *internal.App) *Controller {
	return &Controller{
		app:         app,
		postservice: services.NewPostService(app),
	}
}

func (con *Controller) Create(w http.ResponseWriter, r *http.Request) {

	con.app.Log.WithFields(logrus.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")

	con.postservice.Create(&models.Post{Title: "Blog"}, nil)
	con.SendJSON(w, "Hi")
}

func (con *Controller) Get(w http.ResponseWriter, r *http.Request) {

	posts, err := con.postservice.Get()

	if err != nil {
		switch err := errors.Cause(err).(type) {
		default:
			fmt.Println(err)
			con.app.Log.WithFields(logrus.Fields{
				"message": fmt.Sprintf("%+v", err),
			}).Error("A walrus appears")
		}

		log.Printf("%+v\n", err)
		con.SendServerError(w)
		return
	}

	con.SendJSON(w, posts)
}

func (con *Controller) Show(w http.ResponseWriter, r *http.Request) {

	postid := chi.URLParam(r, "post_id") // from a route like /users/{userID}
	post, err := con.postservice.FindPost(postid)

	if err != nil {
		if errors.Cause(err) == mongo.ErrNoDocuments {
			fmt.Println(err)
			customerrors := []*jsonapi.ErrorObject{{
				Title:  "Post Not Found",
				Detail: fmt.Sprintf("The post %s is not found on our record.", postid),
				Status: "404",
				Code:   core.CODENOTFOUNDERROR,
			}}
			con.SendCustomError(w, customerrors)
			return
		}

		switch err := errors.Cause(err).(type) {
		default:
			fmt.Println(err)
			con.app.Log.WithFields(logrus.Fields{
				"message": fmt.Sprintf("%+v", err),
			}).Error("A walrus appears")
		}

		log.Printf("%+v\n", err)
		con.SendServerError(w)
		return
	}

	con.SendJSON(w, post)
}
