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
	"github.com/stretchr/testify/require"
)

func TestHandler_SendEmailVerification(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	dbConn := infra.ConnectDB(ctx)

	email, err := util.GetEnv("TEST_EMAIL")
	require.NoError(t, err)

	postParams := api.SendEmailVerificationReq{
		Email: email,
	}
	body, err := postParams.MarshalJSON()
	require.NoError(t, err)

	t.Run("Email already used", func(t *testing.T) {
		t.Parallel()

		fixture.CreateUser(t, ctx, dbConn, func(target *rdb.TestCreateUserParams) {
			(*target).Email = email
		})

		resp := testutil.SetupAndRequest(t, ctx, http.MethodPost, "/send-verification", "", bytes.NewBuffer(body))
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			require.NoError(t, err)
		}(resp.Body)
		body, err = io.ReadAll(resp.Body)
		require.NoError(t, err)

		require.Equal(t, http.StatusBadRequest, resp.StatusCode)

		var res400 api.BadRequest
		require.NoError(t, json.Unmarshal(body, &res400))

		require.Equal(t, "The email address is already in use", res400.Message)
	})

	t.Run("Success", func(t *testing.T) {
		t.Parallel()

		resp := testutil.SetupAndRequest(t, ctx, http.MethodPost, "/send-verification", "", bytes.NewBuffer(body))
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			require.NoError(t, err)
		}(resp.Body)

		require.Equal(t, http.StatusOK, resp.StatusCode)
	})
}
