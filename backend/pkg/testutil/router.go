package testutil

import (
	"context"
	"os"

	"net/http"

	"github.com/Tsubasa-2005/EmailAuthResister/pkg/gmail"
	"github.com/Tsubasa-2005/EmailAuthResister/pkg/infra"
	"github.com/Tsubasa-2005/EmailAuthResister/ui/api"
	"github.com/Tsubasa-2005/EmailAuthResister/usesase/handler"
	"github.com/taxio/errors"
)

func SetupRouter(ctx context.Context) (http.Handler, error) {
	dbConn := infra.ConnectDB(ctx)
	hostname := "smtp.gmail.com"
	port := 587
	username := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")

	gmailClient, err := gmail.NewConnect(hostname, port, username, password)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create Gmail client")
	}
	h := handler.NewHandler(dbConn, gmailClient)
	s := handler.NewSecurityHandler()
	srv, err := api.NewServer(h, s, api.WithMiddleware())
	if err != nil {
		return nil, err
	}

	// Wrap the server with the CORS middleware
	corsHandler := testEnableCORS(srv)

	return corsHandler, nil
}

func testEnableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")                                // Allow all origins, adjust as necessary
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE") // Allowed methods
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		// Check if the request is for CORS OPTIONS (pre-flight)
		if r.Method == "OPTIONS" {
			// Just add headers and send response
			w.WriteHeader(http.StatusOK)
			return
		}

		// Serve the next handler
		next.ServeHTTP(w, r)
	})
}
