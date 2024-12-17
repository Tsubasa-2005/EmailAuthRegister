package handler

import (
	"github.com/Tsubasa-2005/EmailAuthResister/pkg/gmail"
	"github.com/Tsubasa-2005/EmailAuthResister/pkg/infra/rdb"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Handler struct {
	dbConn      *pgxpool.Pool
	repo        *rdb.Queries
	gmailClient *gmail.Client
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
