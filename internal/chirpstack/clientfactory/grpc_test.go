package clientfactory

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBearerCredentialsMetadata(t *testing.T) {
	credentials := bearerCreds{token: "test-token"}

	md, err := credentials.GetRequestMetadata(context.Background())

	require.NoError(t, err)
	assert.Equal(t, "Bearer test-token", md["authorization"])
	assert.Equal(t, "chipctl-1", md["user-agent"])
}
