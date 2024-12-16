package handler

import (
	"context"

	"github.com/Tsubasa-2005/EmailAuthResister/pkg/gmail"
	"github.com/Tsubasa-2005/EmailAuthResister/pkg/infra/rdb"
	"github.com/Tsubasa-2005/EmailAuthResister/ui/api"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Handler struct {
	dbConn      *pgxpool.Pool
	repo        *rdb.Queries
	gmailClient *gmail.Client
}

func (h *Handler) GetAllUsers(ctx context.Context) ([]api.User, error) {
	//TODO implement me
	panic("implement me")
}

type SecurityHandler struct{}

func NewHandler(dbConn *pgxpool.Pool, gmailClient *gmail.Client) *Handler {
	return &Handler{
		dbConn:      dbConn,
		repo:        rdb.New(dbConn),
		gmailClient: gmailClient,
	}
}

func NewSecurityHandler() *SecurityHandler {
	return &SecurityHandler{}
}
