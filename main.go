package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spy16/wisenotes/server"
	"github.com/spy16/wisenotes/storage"

	_ "github.com/libsql/libsql-client-go/libsql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

//go:generate sqlc generate

func main() {
	cmd := &cobra.Command{
		Use:   "wisenotes",
		Short: "A wise self-organizing note taking app.",
	}

	var addr, store, logLevel, logFormat string
	cmd.PersistentFlags().StringVarP(&addr, "addr", "a", ":8080", "Address to listen on")
	cmd.PersistentFlags().StringVarP(&logLevel, "log-level", "l", "info", "Log level")
	cmd.PersistentFlags().StringVarP(&logFormat, "log-format", "F", "text", "Log format")
	cmd.PersistentFlags().StringVarP(&store, "store", "s", "file::memory:", "Database connection spec (Use file: prefix for sqlite3 or use libsql:// for libsql client)")

	cmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		initLogger(logLevel, logFormat)
	}

	cmd.Run = func(cmd *cobra.Command, args []string) {
		qu, dbClose, err := storage.Open(store)
		if err != nil {
			log.Fatalf("failed to open db: %v", err)
		}
		defer dbClose()

		log.Infof("starting server on %s...", addr)
		if err := server.Serve(cmd.Context(), ":8080", qu); err != nil {
			log.Fatalf("failed to start server: %v", err)
		}
	}

	_ = cmd.Execute()
}

func initLogger(level, format string) {
	lvl, err := log.ParseLevel(level)
	if err != nil {
		lvl = log.InfoLevel
	}
	log.SetLevel(lvl)

	if format == "json" {
		log.SetFormatter(&log.JSONFormatter{})
	} else {
		log.SetFormatter(&log.TextFormatter{})
	}
}
