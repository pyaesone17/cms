package posts

import (
	"net/http"

	"github.com/pyaesone17/blog/internal"
	"github.com/sirupsen/logrus"
)

type Controller struct {
	app *internal.App
}

func NewController(app *internal.App) *Controller {
	return &Controller{
		app: app,
	}
}

func (con *Controller) Create(w http.ResponseWriter, r *http.Request) {
	con.app.Log.WithFields(logrus.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")

	w.Write([]byte("Hi"))
}
