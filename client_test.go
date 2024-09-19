package rippling_test

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.stellar.af/go-rippling"
)

func Test_Client(t *testing.T) {
	key := os.Getenv("RIPPLING_KEY")
	if key == "" {
		panic("Missing RIPPLING_KEY environment variable")
	}
	client, err := rippling.New(key)
	require.NoError(t, err)
	ctx := context.Background()
	t.Run("get leave requests", func(t *testing.T) {
		req := &rippling.GetLeaveRequestsParams{}
		status := "APPROVED"
		req.Status = &status
		res, err := client.GetLeaveRequests(ctx, req)
		require.NoError(t, err)
		data, err := rippling.ParseGetLeaveRequestsRes(res)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, data.StatusCode())
		require.NotNil(t, data.JSON200)
		assert.GreaterOrEqual(t, len(*data.JSON200), 2)
	})
}
