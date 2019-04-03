package posts

import (
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	validator "gopkg.in/go-playground/validator.v9"

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
	post := &models.Post{}

	con.GetContent(post, r)

	err := con.Validate(post)

	if err != nil {
		// translate all error at once
		errs := err.(validator.ValidationErrors)
		trans, _ := core.UniversalTranslator.GetTranslator("en")

		customerrors := make([]*jsonapi.ErrorObject, 0)
		for _, e := range errs {
			// can translate each error one at a time.
			fmt.Println(e.Translate(trans))
			customerror := &jsonapi.ErrorObject{
				Title:  e.Field(),
				Detail: e.Translate(trans),
				Status: "422",
				Code:   core.VALIDATIONERROR,
			}
			customerrors = append(customerrors, customerror)
		}

		con.SendCustomError(w, customerrors, http.StatusUnprocessableEntity)
		return
	}

	con.postservice.Create(post, nil)
	con.SendJSON(w, post)
}

func (con *Controller) Update(w http.ResponseWriter, r *http.Request) {

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
			con.SendCustomError(w, customerrors, http.StatusNotFound)
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
