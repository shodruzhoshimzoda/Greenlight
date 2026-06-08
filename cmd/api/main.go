package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

// api version
const version = "1.0.0"

// appilcation configuration

type config struct {
	port 	int
	env 	string
} 
type application struct {
	cfg 	config
	logger  *slog.Logger
}



func main() {
	var cfg config
	
	flag.IntVar(&cfg.port, "port", 8080, "API server port" )
	flag.StringVar(&cfg.env, "env", "development", "Environment development|staging|production")
	flag.Parse()


	// initializing logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))


	app := &application{
		cfg: cfg,
		logger: logger,
	}


	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", cfg.port),
		Handler: app.routes() ,
		IdleTimeout: time.Minute,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog: slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	logger.Info("starting server on port:", "addr", srv.Addr, "env",cfg.env)
	err := srv.ListenAndServe()

	logger.Error(err.Error())

	os.Exit(1)
	
}
