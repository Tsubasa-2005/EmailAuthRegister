package handler_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/Tsubasa-2005/EmailAuthResister/pkg/testutil"
	"github.com/Tsubasa-2005/EmailAuthResister/pkg/util"
	"github.com/Tsubasa-2005/EmailAuthResister/ui/api"
	"github.com/stretchr/testify/require"
)

func TestHandler_SendEmailVerification(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	email, err := util.GetEnv("TEST_EMAIL")
	require.NoError(t, err)

	postParams := api.SendEmailVerificationReq{
		Email: email,
	}
	body, err := postParams.MarshalJSON()
	require.NoError(t, err)

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
