package handler

import (
	"context"

	"github.com/Tsubasa-2005/EmailAuthResister/pkg/infra/rdb"
	"github.com/Tsubasa-2005/EmailAuthResister/pkg/util"
	"github.com/Tsubasa-2005/EmailAuthResister/ui/api"
	"github.com/go-faster/errors"
	"github.com/jackc/pgx/v5/pgtype"
)

func (h *Handler) CompleteUserRegistration(ctx context.Context, req *api.CompleteUserRegistrationReq) (api.CompleteUserRegistrationRes, error) {
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

	if verification.Email != req.Email {
		return &api.BadRequest{
			Message: "invalid email",
		}, nil
	}

	hashPass, err := util.GeneratePasswordHash(req.Password)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate password hash")
	}

	// 先に削除することでユーザーが作成されてtokenの削除が失敗したときにデータが残らないようにする。
	if err := h.repo.DeleteEmailVerificationTokensByEmail(ctx, req.Email); err != nil {
		return nil, errors.Wrap(err, "failed to delete email verification token")
	}

	user, err := h.repo.CreateUser(ctx, rdb.CreateUserParams{
		Email:     req.Email,
		Password:  hashPass,
		Name:      req.Name,
		CreatedAt: util.Timestamptz(util.NowJST()),
		UpdatedAt: util.Timestamptz(util.NowJST()),
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to update user")
	}

	token, err := util.GenerateToken("EmailAuthRegister", util.UserToken{
		ID:   user.ID,
		Name: user.Name,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate token")
	}

	return &api.CompleteUserRegistrationOK{
		Token: token,
	}, nil
}
