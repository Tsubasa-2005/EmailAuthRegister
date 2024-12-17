package handler

import (
	"context"

	"github.com/Tsubasa-2005/EmailAuthResister/ui/api"
)

func (h *Handler) Ping(_ context.Context) (*api.PingOK, error) {
	return &api.PingOK{
		Message: "pong",
	}, nil
}
