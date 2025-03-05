package chat

import (
	"context"
	"net/http"

	"github.com/sharveshml/chatapp/logger"
	"github.com/sharveshml/chatapp/web"
)

type app struct {
	log  *logger.Logger
	chat *chat.Chat
}

func New(log *logger.Logger, chat *chat.Chat) *app {
	return &app{
		log:  log,
		chat: chat,
	}
}

func (a *app) connect(ctx context.Context, r *http.Request) {
	usr, err := a.chat.Handshake(ctx, web.GetWriter(ctx), r)
	if err != nil {
		return errs.Newf(errs.FailedPrecondition, "handshake failed: %s", err)
	}
	defer usr.Conn.Close()

	a.chat.ListenSocket(ctx, usr)

	return web.NewNoResponse()
}
