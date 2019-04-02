package posts

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/go-chi/chi"
	"github.com/pkg/errors"
	"github.com/pyaesone17/blog/app/models"
	"github.com/pyaesone17/blog/app/services"
	"github.com/pyaesone17/blog/internal"
	"github.com/sirupsen/logrus"
)

type Controller struct {
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
	w.Write([]byte("Hi"))
}

func (con *Controller) Get(w http.ResponseWriter, r *http.Request) {

	postid := chi.URLParam(r, "post_id") // from a route like /users/{userID}
	post, err := con.postservice.Get(postid)

	if err != nil {
		if errors.Cause(err) == mongo.ErrNoDocuments {
			fmt.Println(err)
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Not found"))
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
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went wrong"))
		return
	}

	data, _ := json.Marshal(post)
	w.Write([]byte(data))
}
