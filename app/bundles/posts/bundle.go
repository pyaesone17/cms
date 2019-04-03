package posts

import (
	"github.com/pyaesone17/blog/app/core"
	"github.com/pyaesone17/blog/internal"
)

type Bundle struct {
	routes []core.Route
}

func NewBundle(app *internal.App) core.Bundle {
	c := NewController(app)
	r := []core.Route{
		core.Route{
			Method:  "POST",
			Path:    "/blogs",
			Handler: c.Create,
		},
		core.Route{
			Method:  "PUT",
			Path:    "/blogs",
			Handler: c.Update,
		},
		core.Route{
			Method:  "GET",
			Path:    "/blogs",
			Handler: c.Get,
		},
		core.Route{
			Method:  "GET",
			Path:    "/blogs/{post_id}",
			Handler: c.Show,
		},
	}
	return &Bundle{r}
}

func (b *Bundle) GetRoutes() []core.Route {
	return b.routes
}
