package fixture

import (
	"context"
	"testing"

	"github.com/Tsubasa-2005/EmailAuthResister/pkg/infra/rdb"
	"github.com/Tsubasa-2005/EmailAuthResister/pkg/util"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/require"
)

func CreateUser(t *testing.T, ctx context.Context, db rdb.DBTX, f func(target *rdb.TestCreateUserParams)) rdb.User {
	t.Helper()

	target := &rdb.TestCreateUserParams{
		Email:     faker.Email(),
		Password:  faker.Password(),
		Name:      faker.Name(),
		CreatedAt: util.Timestamptz(util.NowJST()),
		UpdatedAt: util.Timestamptz(util.NowJST()),
	}

	if f != nil {
		f(target)
	}

	created, err := rdb.New(db).TestCreateUser(ctx, rdb.TestCreateUserParams{
		Email:     target.Email,
		Password:  target.Password,
		Name:      target.Name,
		CreatedAt: target.CreatedAt,
		UpdatedAt: target.UpdatedAt,
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		require.NoError(t, rdb.New(db).TestDeleteUser(ctx, created.ID))
	})

	return created
}
