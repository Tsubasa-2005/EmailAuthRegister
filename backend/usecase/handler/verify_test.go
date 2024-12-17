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
	"github.com/stretchr/testify/require"
)

func TestHandler_VerifyEmail(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	dbConn := infra.ConnectDB(ctx)

	t.Run("Invalid token", func(t *testing.T) {
		t.Parallel()

		verification := fixture.CreateEmailVerificationToken(t, ctx, dbConn, func(target *rdb.TestCreateEmailVerificationTokenParams) {
			(*target).ExpiresAt = util.Timestamptz(util.NowJST().Add(-24 * time.Hour))
		})

		postParams := api.VerifyEmailReq{
			Token: verification.Token.Bytes,
		}
		body, err := postParams.MarshalJSON()
		require.NoError(t, err)

		resp := testutil.SetupAndRequest(t, ctx, http.MethodPost, "/verify-email", "", bytes.NewBuffer(body))
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			require.NoError(t, err)
		}(resp.Body)
		body, err = io.ReadAll(resp.Body)
		require.NoError(t, err)

		require.Equal(t, http.StatusUnauthorized, resp.StatusCode)

		var res401 api.Unauthorized
		require.NoError(t, json.Unmarshal(body, &res401))

		require.Equal(t, "expired token", res401.Message)
	})

	t.Run("Success", func(t *testing.T) {
		t.Parallel()

		verification := fixture.CreateEmailVerificationToken(t, ctx, dbConn, nil)

		postParams := api.VerifyEmailReq{
			Token: verification.Token.Bytes,
		}
		body, err := postParams.MarshalJSON()
		require.NoError(t, err)

		resp := testutil.SetupAndRequest(t, ctx, http.MethodPost, "/verify-email", "", bytes.NewBuffer(body))
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			require.NoError(t, err)
		}(resp.Body)
		body, err = io.ReadAll(resp.Body)
		require.NoError(t, err)

		require.Equal(t, http.StatusOK, resp.StatusCode)

		var res200 api.VerifyEmailOK
		require.NoError(t, json.Unmarshal(body, &res200))

		require.Equal(t, verification.Email, res200.Email)
	})
}
