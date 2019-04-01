package posts

import (
	"github.com/pyaesone17/blog/app/core"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Bundle struct {
	routes []core.Route
}

func NewBundle(config *viper.Viper, log *logrus.Logger) core.Bundle {
	c := NewController(config, log)
	r := []core.Route{
		core.Route{
			Method:  "POST",
			Path:    "/blogs",
			Handler: c.Create,
		},
		core.Route{
			Method:  "GET",
			Path:    "/blogs",
			Handler: c.Create,
		},
	}
	return &Bundle{r}
}

func (b *Bundle) GetRoutes() []core.Route {
	return b.routes
}
