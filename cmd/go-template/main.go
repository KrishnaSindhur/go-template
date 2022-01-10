package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/KrishnaSindhur/go-template/pkg/handler"

	"github.com/KrishnaSindhur/go-template/pkg/constants"
	"github.com/gorilla/mux"
	lgr "github.com/sirupsen/logrus"
)

const (
	help    = ``
	version = `1.0.0`
)

const (
	HealthPath = "/health"
)

func main() {
	program := os.Args[0]
	if len(os.Args) <= 1 {
		lgr.Fatalf("No arguments: %v, provided run help command for more details", program)
		os.Exit(1)
	}
	cmd := os.Args[1]
	switch strings.ToLower(cmd) {
	case "serve":
		serve()
	case "migrate":
		migrate()
	case "help":
		fmt.Println(help)
	case "version":
		fmt.Println(version)
	default:
		fmt.Println("unknown command, for command details pass help argument")
	}
}

func migrate() {
	return
}

func serve() {
	// initialize the config

	lgr.Infof("%s started", constants.ServiceName)

	//initilaize the postgres

	router := mux.NewRouter()

	router.HandleFunc(HealthPath, handler.Health).Methods((http.MethodGet))

	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
	})

	router.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
	})

	// should read port from config
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", "8080"),
		Handler: router,
	}

	go func() {
		lgr.Infof("starting server in port: %v", "8080")
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			lgr.Panic("failed to start server")
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	<-stop
	lgr.Info("shutting server down")
	if err := server.Shutdown(context.Background()); err != nil {
		lgr.Error("server did not shut down carefully")
	} else {
		lgr.Info("server down")
	}
}
