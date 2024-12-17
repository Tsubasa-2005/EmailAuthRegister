package cmd

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/Tsubasa-2005/EmailAuthResister/pkg/gmail"
	"github.com/Tsubasa-2005/EmailAuthResister/pkg/infra"
	"github.com/Tsubasa-2005/EmailAuthResister/pkg/middleware"
	"github.com/Tsubasa-2005/EmailAuthResister/ui/api"
	"github.com/Tsubasa-2005/EmailAuthResister/usecase/handler"
	"github.com/go-faster/errors"
	"github.com/spf13/cobra"
)

func serverCmd(ctx context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "Start the HTTP server",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runServer(ctx)
		},
	}
}

func runServer(ctx context.Context) error {
	dbConn := infra.ConnectDB(ctx)
	hostname := "smtp.gmail.com"
	port := 587
	username := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")

	gmailClient, err := gmail.NewConnect(hostname, port, username, password)
	if err != nil {
		return errors.Wrap(err, "failed to create Gmail client")
	}
	h := handler.NewHandler(dbConn, gmailClient)
	s := handler.NewSecurityHandler()

	srv, err := api.NewServer(h, s, api.WithMiddleware(middleware.Logging()))
	if err != nil {
		return err
	}

	corsHandler := middleware.Cors(srv)

	addr := ":8080"
	slog.InfoContext(ctx, "Starting HTTP Server", slog.String("addr", addr))
	if err := http.ListenAndServe(addr, corsHandler); err != nil {
		return err
	}

	return nil
}
