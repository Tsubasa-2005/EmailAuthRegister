package fixture

import (
	"context"
	"testing"
	"time"

	"github.com/Tsubasa-2005/EmailAuthResister/pkg/infra/rdb"
	"github.com/Tsubasa-2005/EmailAuthResister/pkg/util"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func CreateEmailVerificationToken(t *testing.T, ctx context.Context, db rdb.DBTX, f func(target *rdb.TestCreateEmailVerificationTokenParams)) rdb.EmailVerificationToken {
	t.Helper()

	target := &rdb.TestCreateEmailVerificationTokenParams{
		Token: pgtype.UUID{
			Bytes: uuid.New(),
			Valid: true,
		},
		Email:     faker.Email(),
		ExpiresAt: util.Timestamptz(util.NowJST().Add(24 * time.Hour)),
		CreatedAt: util.Timestamptz(util.NowJST()),
	}

	if f != nil {
		f(target)
	}

	created, err := rdb.New(db).TestCreateEmailVerificationToken(ctx, rdb.TestCreateEmailVerificationTokenParams{
		Token:     target.Token,
		Email:     target.Email,
		ExpiresAt: target.ExpiresAt,
		CreatedAt: target.CreatedAt,
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		require.NoError(t, rdb.New(db).TestDeleteEmailVerificationToken(ctx, created.Token))
	})

	return created
}
