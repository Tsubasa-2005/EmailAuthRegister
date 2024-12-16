package handler

import (
	"context"

	"github.com/Tsubasa-2005/EmailAuthResister/pkg/util"
	"github.com/Tsubasa-2005/EmailAuthResister/ui/api"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/taxio/errors"
)

func (h *Handler) VerifyEmail(ctx context.Context, req *api.VerifyEmailReq) (api.VerifyEmailRes, error) {
	verification, err := h.repo.GetEmailVerificationToken(ctx, pgtype.UUID{
		Bytes: req.Token,
		Valid: true,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to get email verification info")
	}

	if verification.ExpiresAt.Time.Before(util.NowJST()) {
		return &api.Unauthorized{
			Message: "expired token",
		}, nil
	}

	return &api.VerifyEmailOK{
		Email: verification.Email,
	}, nil
}
