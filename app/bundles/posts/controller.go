package posts

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Controller struct {
	config *viper.Viper
	log    *logrus.Logger
}

func NewController(config *viper.Viper, log *logrus.Logger) *Controller {
	return &Controller{
		config: config,
		log:    log,
	}
}

func (con *Controller) Create(w http.ResponseWriter, r *http.Request) {
	con.log.WithFields(logrus.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")

	w.Write([]byte("Hi"))
}
