package web

import (
	"context"
	"fmt"
	"net/http"
)

type App struct {
	log     Logger
	mux     *http.ServeMux
	mw      []MidFunc
	origins []string
}

type Encoder interface {
	Encode() (data []byte, contentType string, err error)
}

func NewApp(log Logger, mw ...MidFunc) *App {
	return &App{
		log: log,
		mux: http.NewServeMux(),
		mw:  mw,
	}
}

func (a *App) ServeHTTP(r *http.Request, w http.ResponseWriter) {
	a.mux.ServeHTTP(w, r)
}

func (a *App) EnableCORS(origins []string) {
	a.origins = origins

	handler := func(r *http.Request, w http.ResponseWriter) Encoder {
		return nil
	}
	handler = wrapMiddleware([]MidFunc{a.corsHandler}, handler)

	a.HandlerFuncNoMid(http.MethodOptions, "", "/", handler)
}

func (a *App) corsHandler(webHandler HandlerFunc) HandlerFunc {

	h := func(r *http.Request, ctx context.Context) Encoder {

		w := GetWriter(ctx)
		reqOrigin := r.Header.Get("Origin")

		for _, origin := range a.origins {
			if origin == "*" || origin == reqOrigin {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				break
			}
		}

		w.Header().Set("Access-Control-Allow-Methods", "POST, PATCH, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.Header().Set("Access-Control-Max-Age", "86400")

		return webHandler(r, ctx)
	}

	return h
}

func (a *App) HandlerFuncNoMid(method string, group string, path string, handlerFunc HandlerFunc) {
	h := func(r *http.Request, w http.ResponseWriter) {
		ctx := setWriter(r.Context(), w)

		resp := handlerFunc(ctx, r)

		if err := Respond(ctx, w, resp); err != nil {
			a.log(ctx, "web-respond", "ERROR", err)
			return
		}
	}

	finalPath := path
	if group != "" {
		finalPath = "/" + group + path
	}
	finalPath = fmt.Sprintf("%s %s", method, finalPath)

	a.mux.HandleFunc(finalPath, h)
}
