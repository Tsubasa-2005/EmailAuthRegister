package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/Tsubasa-2005/EmailAuthResister/pkg/domain"
	"github.com/Tsubasa-2005/EmailAuthResister/pkg/infra/rdb"
	"github.com/Tsubasa-2005/EmailAuthResister/pkg/util"
	"github.com/Tsubasa-2005/EmailAuthResister/ui/api"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/taxio/errors"
)

func (h *Handler) SendEmailVerification(ctx context.Context, req *api.SendEmailVerificationReq) (api.SendEmailVerificationRes, error) {
	token := uuid.New()

	if _, err := h.repo.CreateEmailVerificationToken(ctx, rdb.CreateEmailVerificationTokenParams{
		Token: pgtype.UUID{
			Bytes: token,
			Valid: true,
		},
		Email:     req.Email,
		ExpiresAt: util.Timestamptz(util.NowJST().Add(24 * time.Hour)),
		CreatedAt: util.Timestamptz(util.NowJST()),
	}); err != nil {
		return nil, errors.Wrap(err, "failed to create email verification token")
	}

	verificationURL := fmt.Sprintf("https://your-domain.com/verify?token=%s", token.String())

	template, err := domain.LoadEmailVerificationTemplates()
	if err != nil {
		return nil, errors.Wrap(err, "failed to load email templates")
	}

	body, err := domain.GenerateEmailVerificationBody(template.Body, domain.EmailVerificationData{
		URL: verificationURL,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate email body")
	}

	if err := h.gmailClient.SendEmail(req.Email, template.Subject, body); err != nil {
		return nil, errors.Wrap(err, "failed to send email")
	}

	return &api.SendEmailVerificationOK{}, nil
}
