package handler

import (
	"context"

	"github.com/Tsubasa-2005/EmailAuthResister/pkg/infra/rdb"
	"github.com/Tsubasa-2005/EmailAuthResister/pkg/util"
	"github.com/Tsubasa-2005/EmailAuthResister/ui/api"
	"github.com/taxio/errors"
)

func (h *Handler) GetAllUsers(ctx context.Context, params api.GetAllUsersParams) (api.GetAllUsersRes, error) {
	if _, ok := ctx.Value("user").(util.UserToken); !ok {
		return nil, errors.New("user not found")
	}

	count, err := h.repo.CountActiveUsers(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to count active users")
	}

	totalPage := int((count + Limit - 1) / Limit)

	if totalPage < params.Page {
		return &api.BadRequest{
			Message: "Page is out of range",
		}, nil
	}

	p := rdb.GetUsersParams{
		Limit:  Limit,
		Offset: int32(Limit * (params.Page - 1)),
	}

	users, err := h.repo.GetUsers(ctx, p)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get users")
	}

	res := make([]api.User, 0, len(users))
	for _, user := range users {
		res = append(res, api.User{
			ID:    api.ID(user.ID),
			Name:  api.Name(user.Name),
			Email: api.Email(user.Email),
		})
	}

	return &api.GetAllUsersOK{
		Users:     res,
		TotalPage: totalPage,
	}, nil
}
