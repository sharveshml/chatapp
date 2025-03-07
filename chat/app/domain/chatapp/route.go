package chatapp

import (
	"net/http"

	"github.com/sharveshml/chatapp/chat/foundation/web"
)

func Routes(app *web.App) {
	api := newApp()

	app.HandlerFunc(http.MethodGet, "", "/test", api.test)
}
