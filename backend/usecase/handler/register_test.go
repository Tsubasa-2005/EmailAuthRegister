package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/Tsubasa-2005/EmailAuthResister/pkg/infra"
	"github.com/Tsubasa-2005/EmailAuthResister/pkg/infra/rdb"
	"github.com/Tsubasa-2005/EmailAuthResister/pkg/testutil"
	"github.com/Tsubasa-2005/EmailAuthResister/pkg/testutil/fixture"
	"github.com/Tsubasa-2005/EmailAuthResister/pkg/util"
	"github.com/Tsubasa-2005/EmailAuthResister/ui/api"
	"github.com/Tsubasa-2005/EmailAuthResister/usecase/handler"
	"github.com/go-faker/faker/v4"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
)

func TestHandler_CompleteUserRegistration(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	dbConn := infra.ConnectDB(ctx)

	t.Run("Invalid token", func(t *testing.T) {
		t.Parallel()

		verification := fixture.CreateEmailVerificationToken(t, ctx, dbConn, func(target *rdb.TestCreateEmailVerificationTokenParams) {
			(*target).ExpiresAt = util.Timestamptz(util.NowJST().Add(-24 * time.Hour))
		})

		postParams := api.CompleteUserRegistrationReq{
			Token:    verification.Token.Bytes,
			Email:    verification.Email,
			Name:     faker.Name(),
			Password: faker.Password(),
		}
		body, err := postParams.MarshalJSON()
		require.NoError(t, err)

		resp := testutil.SetupAndRequest(t, ctx, http.MethodPost, "/complete-registration", "", bytes.NewBuffer(body))
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			require.NoError(t, err)
		}(resp.Body)
		require.Equal(t, http.StatusUnauthorized, resp.StatusCode)
		body, err = io.ReadAll(resp.Body)
		require.NoError(t, err)

		var res401 api.Unauthorized
		require.NoError(t, json.Unmarshal(body, &res401))

		require.Equal(t, "expired token", res401.Message)
	})

	t.Run("Invalid email", func(t *testing.T) {
		t.Parallel()

		verification := fixture.CreateEmailVerificationToken(t, ctx, dbConn, nil)

		postParams := api.CompleteUserRegistrationReq{
			Token:    verification.Token.Bytes,
			Email:    faker.Email(),
			Name:     faker.Name(),
			Password: faker.Password(),
		}
		body, err := postParams.MarshalJSON()
		require.NoError(t, err)

		resp := testutil.SetupAndRequest(t, ctx, http.MethodPost, "/complete-registration", "", bytes.NewBuffer(body))
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			require.NoError(t, err)
		}(resp.Body)
		require.Equal(t, http.StatusBadRequest, resp.StatusCode)
		body, err = io.ReadAll(resp.Body)
		require.NoError(t, err)

		var res400 api.BadRequest
		require.NoError(t, json.Unmarshal(body, &res400))

		require.Equal(t, "invalid email", res400.Message)
	})

	t.Run("Success", func(t *testing.T) {
		t.Parallel()

		verification := fixture.CreateEmailVerificationToken(t, ctx, dbConn, nil)

		postParams := api.CompleteUserRegistrationReq{
			Token:    verification.Token.Bytes,
			Email:    verification.Email,
			Name:     faker.Name(),
			Password: faker.Password(),
		}
		body, err := postParams.MarshalJSON()
		require.NoError(t, err)

		resp := testutil.SetupAndRequest(t, ctx, http.MethodPost, "/complete-registration", "", bytes.NewBuffer(body))
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			require.NoError(t, err)
		}(resp.Body)
		require.Equal(t, http.StatusOK, resp.StatusCode)
		body, err = io.ReadAll(resp.Body)
		require.NoError(t, err)

		token, err := util.ParseToken(resp.Header.Get("Set-Cookie"))
		require.NoError(t, err)

		claims := token.Claims.(jwt.MapClaims)

		id, err := handler.GetInt64FromClaims(claims, "user_id")
		require.NoError(t, err)

		_, ok := claims["name"].(string)
		require.True(t, ok)

		require.NoError(t, rdb.New(dbConn).TestDeleteUser(ctx, id))

		exists, err := rdb.New(dbConn).TestExistsEmailVerificationTokenByEmail(ctx, postParams.Email)
		require.NoError(t, err)

		require.Equal(t, false, exists)
	})
}
