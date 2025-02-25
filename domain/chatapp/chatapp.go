package chatapp

import (
	"context"
	"fmt"
	"net/http"

	"github.com/sharveshml/chatapp/web"
)

type app struct {
	log  *logger.Logger
	chat *chat.Chat
}

func newApp() *app {
	return &app{
		log:  log,
		chat: chat,
	}
}

func (a *app) connect(ctx context.Context, r *http.Request) {
	usr, err := a.chat.Handshake(ctx, web.GetWriter(ctx), r)
	if err != nil {
		return fmt.Errorf("Handshake failed")
	}
	defer usr.Conn.Close()

	a.chat.Listen(ctx, usr)

	return web.NewNoResponse()
}
