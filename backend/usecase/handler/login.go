package handler

import (
	"context"

	"github.com/Tsubasa-2005/EmailAuthResister/pkg/util"
	"github.com/Tsubasa-2005/EmailAuthResister/ui/api"
	"github.com/go-faster/errors"
)

func (h *Handler) Login(ctx context.Context, req *api.LoginReq) (api.LoginRes, error) {
	user, err := h.repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return &api.BadRequest{
			Message: "Email or password is incorrect.",
		}, nil
	}
	if util.CompareHashAndPassword(user.Password, req.Password) != nil {
		return &api.BadRequest{
			Message: "Email or password is incorrect.",
		}, nil
	}

	token, err := util.GenerateToken("EmailAuthResister", util.UserToken{
		ID:   user.ID,
		Name: user.Name,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate token")
	}

	return &api.LoginOK{
		SetCookie: token,
	}, nil
}
