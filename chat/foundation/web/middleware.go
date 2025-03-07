package web

type MidFunc func(hander HandlerFunc) HandlerFunc

func wrapMiddleware(middlewares []MidFunc, handler HandlerFunc) HandlerFunc {
	for i := len(middlewares) - 1; i >= 0; i-- {
		if middlewares[i] != nil {
			handler = middlewares[i](handler)
		}
	}
	return handler
}
