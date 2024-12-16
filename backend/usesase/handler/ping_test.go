package handler_test

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/Tsubasa-2005/EmailAuthResister/pkg/testutil"
	"github.com/Tsubasa-2005/EmailAuthResister/ui/api"
	"github.com/stretchr/testify/require"
)

func TestHandler_Ping(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	resp := testutil.SetupAndRequest(t, ctx, http.MethodGet, "/ping", "", nil)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		require.NoError(t, err)
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	t.Run("Success ping", func(t *testing.T) {
		t.Parallel()

		require.Equal(t, http.StatusOK, resp.StatusCode)

		var res200 api.PingOK
		require.NoError(t, json.Unmarshal(body, &res200))

		require.Equal(t, "pong", res200.Message)
	})
}
