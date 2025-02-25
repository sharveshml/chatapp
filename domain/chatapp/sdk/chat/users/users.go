package users

import (
	"sync"

	"github.com/google/uuid"
	"github.com/sharveshml/chatapp/domain/chatapp/sdk/chat"
	"github.com/sharveshml/chatapp/logger"
)

type Users struct {
	log     *logger.Logger
	user    map[uuid.UUID]chat.User
	muUsers sync.RWMutex
}
