package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"

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

func TestHandler_Login(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	dbConn := infra.ConnectDB(ctx)

	t.Run("Invalid email", func(t *testing.T) {
		t.Parallel()

		postParams := api.LoginReq{
			Email:    faker.Email(),
			Password: faker.Password(),
		}
		body, err := postParams.MarshalJSON()
		require.NoError(t, err)

		resp := testutil.SetupAndRequest(t, ctx, http.MethodPost, "/login", "", bytes.NewBuffer(body))
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			require.NoError(t, err)
		}(resp.Body)
		require.Equal(t, http.StatusBadRequest, resp.StatusCode)
		body, err = io.ReadAll(resp.Body)
		require.NoError(t, err)

		var res400 api.BadRequest
		require.NoError(t, json.Unmarshal(body, &res400))

		require.Equal(t, "Email or password is incorrect.", res400.Message)
	})

	t.Run("Invalid password", func(t *testing.T) {
		t.Parallel()

		pass := faker.Password()

		user, _ := fixture.CreateUser(t, ctx, dbConn, func(target *rdb.TestCreateUserParams) {
			(*target).Password = pass
		})

		postParams := api.LoginReq{
			Email:    user.Email,
			Password: "invalid",
		}
		body, err := postParams.MarshalJSON()
		require.NoError(t, err)

		resp := testutil.SetupAndRequest(t, ctx, http.MethodPost, "/login", "", bytes.NewBuffer(body))
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			require.NoError(t, err)
		}(resp.Body)
		require.Equal(t, http.StatusBadRequest, resp.StatusCode)
		body, err = io.ReadAll(resp.Body)
		require.NoError(t, err)

		var res400 api.BadRequest
		require.NoError(t, json.Unmarshal(body, &res400))

		require.Equal(t, "Email or password is incorrect.", res400.Message)
	})

	t.Run("Success", func(t *testing.T) {
		t.Parallel()

		pass := faker.Password()

		user, _ := fixture.CreateUser(t, ctx, dbConn, func(target *rdb.TestCreateUserParams) {
			(*target).Password = pass
		})

		postParams := api.LoginReq{
			Email:    user.Email,
			Password: pass,
		}
		body, err := postParams.MarshalJSON()
		require.NoError(t, err)

		resp := testutil.SetupAndRequest(t, ctx, http.MethodPost, "/login", "", bytes.NewBuffer(body))
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			require.NoError(t, err)
		}(resp.Body)
		require.Equal(t, http.StatusOK, resp.StatusCode)

		token, err := util.ParseToken(resp.Header.Get("Set-Cookie"))
		require.NoError(t, err)

		claims := token.Claims.(jwt.MapClaims)

		id, err := handler.GetInt64FromClaims(claims, "user_id")
		require.NoError(t, err)

		require.Equal(t, user.ID, id)

		name, ok := claims["name"].(string)
		require.True(t, ok)

		require.Equal(t, user.Name, name)
	})
}
