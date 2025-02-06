package web

import "net/http"

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

type MidFunc func(hander HandlerFunc) HandlerFunc

func wrappedMiddleware(middlewares []MidFunc, handler HandlerFunc) HandlerFunc {
	for i := len(middlewares) - 1; i >= 0; i-- {
		if middlewares[i] != nil {
			handler = middlewares[i](handler)
		}
	}
	return handler
}
