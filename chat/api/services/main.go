package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/ardanlabs/conf/v3"
	"github.com/sharveshml/chatapp/chat/app/sdk/mux"
	"github.com/sharveshml/chatapp/chat/foundation/logger"
	"github.com/sharveshml/chatapp/chat/foundation/web"
)

var build = "dev"

func main() {
	var log *logger.Logger

	traceIdFn := func(ctx context.Context) string {
		return web.GetTraceID(ctx).String()
	}

	log = logger.New(os.Stdout, logger.LevelInfo, "CAP", traceIdFn)

	ctx := context.Background()

	if err := run(ctx, log); err != nil {
		log.Error(ctx, "startup", "err", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, log *logger.Logger) error {
	log.Info(ctx, "startup", runtime.GOMAXPROCS(0))

	cfg := struct {
		conf.Version
		Web struct {
			ReadTimeout     time.Duration
			WriteTimeout    time.Duration
			IdleTimeout     time.Duration
			ShutdownTimeout time.Duration
			APIHost         string
		}
	}{
		Version: conf.Version{
			Build: build,
			Desc:  "CAP",
		},
	}

	const Prefix = "SALES"
	help, err := conf.Parse(Prefix, &cfg)

	if err != nil {
		if errors.Is(err, conf.ErrHelpWanted) {
			fmt.Println(help)
			return nil
		}
		return fmt.Errorf("parsing error: %w", err)
	}

	log.Info(ctx, "App Starting", "version", cfg.Build)
	defer log.Info(ctx, "Shutdown Complete")

	out, err := conf.String(&cfg)
	if err != nil {
		return fmt.Errorf("generating config for output: %w", err)
	}
	log.Info(ctx, "startup", "config", out)

	log.BuildInfo(ctx)

	log.Info(ctx, "startup", "initializing V1 API support")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	cfgMux := mux.Config{
		Log: log,
	}

	WebAPI := mux.WebAPI(cfgMux)

	api := http.Server{
		Addr:         cfg.Web.APIHost,
		Handler:      WebAPI,
		ReadTimeout:  cfg.Web.ReadTimeout,
		WriteTimeout: cfg.Web.WriteTimeout,
		IdleTimeout:  cfg.Web.IdleTimeout,
		ErrorLog:     logger.NewStdLogger(log, logger.LevelError),
	}

	serverErrors := make(chan error, 1)

	go func() {
		log.Info(ctx, "startup", "status", "api router started", "host", api.Addr)

		serverErrors <- api.ListenAndServe()
	}()

	// -------------------------------------------------------------------------
	// Shutdown

	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)

	case sig := <-shutdown:
		log.Info(ctx, "shutdown", "status", "shutdown started", "signal", sig)
		defer log.Info(ctx, "shutdown", "status", "shutdown complete", "signal", sig)

		ctx, cancel := context.WithTimeout(ctx, cfg.Web.ShutdownTimeout)
		defer cancel()

		if err := api.Shutdown(ctx); err != nil {
			api.Close()
			return fmt.Errorf("could not stop server gracefully: %w", err)
		}
	}

	return nil
}
