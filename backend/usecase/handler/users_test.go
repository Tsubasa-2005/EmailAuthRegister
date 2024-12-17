package handler_test

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"testing"

	"github.com/Tsubasa-2005/EmailAuthResister/pkg/infra"
	"github.com/Tsubasa-2005/EmailAuthResister/pkg/infra/rdb"
	"github.com/Tsubasa-2005/EmailAuthResister/pkg/testutil"
	"github.com/Tsubasa-2005/EmailAuthResister/pkg/testutil/fixture"
	"github.com/Tsubasa-2005/EmailAuthResister/ui/api"
	"github.com/Tsubasa-2005/EmailAuthResister/usecase/handler"
	"github.com/stretchr/testify/require"
)

func TestHandler_GetAllUsers(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	dbConn := infra.ConnectDB(ctx)

	t.Run("Invalid token", func(t *testing.T) {
		t.Parallel()

		queryParams := url.Values{}
		queryParams.Set("page", "1")
		path := "/users?" + queryParams.Encode()

		resp := testutil.SetupAndRequest(t, ctx, http.MethodGet, path, "token", nil)
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			require.NoError(t, err)
		}(resp.Body)

		require.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	})

	t.Run("Invalid page param", func(t *testing.T) {
		t.Parallel()

		_, token := fixture.CreateUser(t, ctx, dbConn, nil)

		queryParams := url.Values{}
		queryParams.Set("page", "100")
		path := "/users?" + queryParams.Encode()

		resp := testutil.SetupAndRequest(t, ctx, http.MethodGet, path, token, nil)
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			require.NoError(t, err)
		}(resp.Body)
		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)

		require.Equal(t, http.StatusBadRequest, resp.StatusCode)

		var res400 api.BadRequest
		require.NoError(t, json.Unmarshal(body, &res400))

		require.Equal(t, "Page is out of range", res400.Message)
	})

	t.Run("Success", func(t *testing.T) {
		t.Parallel()

		_, token := fixture.CreateUser(t, ctx, dbConn, nil)

		queryParams := url.Values{}
		queryParams.Set("page", "1")
		path := "/users?" + queryParams.Encode()

		resp := testutil.SetupAndRequest(t, ctx, http.MethodGet, path, token, nil)
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			require.NoError(t, err)
		}(resp.Body)
		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)

		require.Equal(t, http.StatusOK, resp.StatusCode)

		var res200 api.GetAllUsersOK
		require.NoError(t, json.Unmarshal(body, &res200))

		for _, user := range res200.Users {
			u, err := rdb.New(dbConn).TestGetUser(ctx, int64(user.ID))
			require.NoError(t, err)

			require.Equal(t, api.Name(u.Name), user.Name)
			require.Equal(t, api.Email(u.Email), user.Email)
		}

		count, err := rdb.New(dbConn).TestCountActiveUsers(ctx)
		require.NoError(t, err)

		totalPage := int((count + handler.Limit - 1) / handler.Limit)
		require.Equal(t, totalPage, res200.TotalPage)
	})
}
