package testutil

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

// SetupAndRequest
// Usage Examples:
//
//  1. For POST or PUT requests, the body can be prepared as follows:
//
//     ```go
//     postParams := api.SomeReq{Key: "Value"}
//
//     body, err := postParams.MarshalJSON()
//
//     require.NoError(t, err)
//
//     resp := SetupAndRequest(t, ctx, http.MethodPost, "/users", "", bytes.NewBuffer(body))
//     ```
//
//  2. For GET requests with query parameters, the path can be constructed as follows:
//
//     ```go
//     queryParams := url.Values{}
//
//     queryParams.Set("limit", "10")
//
//     queryParams.Set("offset", "0")
//
//     path := "/path?" + queryParams.Encode()
//
//     resp := SetupAndRequest(t, ctx, http.MethodGet, path, "", nil)
//     ```
func SetupAndRequest(t *testing.T, ctx context.Context, method, path, token string, body io.Reader) *http.Response {
	t.Helper()

	router, err := SetupRouter(ctx)
	require.NoError(t, err)

	server := httptest.NewServer(router)
	defer server.Close()

	if method == http.MethodGet && body != nil {
		t.Fatalf("GET method must not have a body")
	}

	req, err := http.NewRequest(method, server.URL+path, body)
	require.NoError(t, err)

	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	if method == http.MethodPost || method == http.MethodPut {
		req.Header.Set("Content-Type", "application/json")
	}

	client := server.Client()
	resp, err := client.Do(req)
	require.NoError(t, err)

	return resp
}
