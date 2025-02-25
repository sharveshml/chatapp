package chatapp

import (
	"net/http"

	"github.com/sharveshml/chatapp/web"
)

func Route(app *web.App) {
	api := newApp()

	app.HandlerFunc(http.MethodGet)
}
