package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/gorilla/mux"

	"github.com/joepk90/graphql-auth/internal/auth"
	"github.com/joepk90/graphql-auth/internal/middleware"
	"github.com/joepk90/graphql-auth/internal/policy"
	"github.com/joepk90/graphql-auth/internal/stats"

	cli "github.com/jawher/mow.cli"
	log "github.com/sirupsen/logrus"
	_ "google.golang.org/grpc/xds"
)

const (
	appName = "graphql-auth"
	appDesc = "GraphQL Auth used for handling Frontend Authorization Operations"
)

func main() {
	app := cli.App(appName, appDesc)

	httpBind := app.Int(cli.IntOpt{
		Name:   "http",
		Desc:   "the port to bind the GraphQL server to",
		EnvVar: "HTTP_BIND",
		Value:  8080,
	})
	opsPort := app.Int(cli.IntOpt{
		Name:   "ops-port",
		Desc:   "The HTTP ops port",
		EnvVar: "OPS_PORT",
		Value:  8081,
	})
	logLevel := app.String(cli.StringOpt{
		Name:   "log-level",
		Desc:   "log level [debug|info|warn|error]",
		EnvVar: "LOG_LEVEL",
		Value:  "info",
	})

	logFormat := app.String(cli.StringOpt{
		Name:   "log-format",
		Desc:   "Log format, if set to text will use text as logging format, otherwise will use json",
		EnvVar: "LOG_FORMAT",
		Value:  "text",
	})

	authServiceUrl := app.String(cli.StringOpt{
		Name:   "auth-service-url",
		Desc:   "URL of the authorisation service",
		EnvVar: "AUTH_SERVICE_URL",
		Value:  "http:/localhost:8080",
	})

	localAuthToken := app.String(cli.StringOpt{
		Name:   "auth-token",
		Desc:   "auth token used for development purposes",
		EnvVar: "LOCAL_AUTH_TOKEN",
		Value:  "",
	})

	app.Action = func() {
		ctx := context.Background()

		configureLogger(*logLevel, *logFormat)

		metrics := &stats.PrometheusMetrics{}

		authService := auth.NewAuthService(*authServiceUrl)
		authorizer := auth.NewAuthorizer(authService)
		service := policy.NewService(metrics, &authorizer)
		schema, err := service.ToSchema()
		if err != nil {
			log.WithError(err).Panic("unable to create GraphQL schema")
		}

		router := mux.NewRouter()

		go func() {
			router.Handle(
				"/graphql",
				middleware.CORSHTTPMiddleware(
					middleware.CORSHTTPMiddleware(
						middleware.HTTPHandler(true)(
							middleware.GQLHTTPMiddleware(schema),
						),
					),
				),
			).Methods(http.MethodGet, http.MethodPost, http.MethodOptions)
			router.Handle("/", middleware.GQLiHTTPMiddleware(*localAuthToken))
			s := &http.Server{
				Handler: router,
				Addr:    fmt.Sprintf(":%d", *httpBind),
			}
			log.Fatal(s.ListenAndServe())
		}()

		opsServer := initialiseOpsServer(opsPort)
		go startOpsServer(opsServer)
		defer opsServer.Shutdown(ctx)

		waitForShutdown()
	}

	if err := app.Run(os.Args); err != nil {
		log.Panic(err)
	}
}

func initialiseOpsServer(opsPort *int) *http.Server {
	return &http.Server{
		Addr: fmt.Sprintf(":%d", *opsPort),
	}
}

func startOpsServer(opsServer *http.Server) {
	opsErr := opsServer.ListenAndServe()
	switch opsErr {
	case http.ErrServerClosed:
		log.WithError(opsErr).Warn("ops server shutdown")
	default:
		log.WithError(opsErr).Panic("unable to start ops http server")
	}
}

func configureLogger(level, format string) {
	l, err := log.ParseLevel(level)
	if err != nil {
		log.WithFields(log.Fields{"log_level": level}).
			WithError(err).
			Panic("invalid log level")
	}
	log.SetLevel(l)

	format = strings.ToLower(format)
	if format != "text" && format != "json" {
		log.Panicf("invalid log format: %s", format)
	}
	if format == "json" {
		log.SetFormatter(&log.JSONFormatter{})
	}
}

func waitForShutdown() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan
	log.Warn("shutting down")
}
