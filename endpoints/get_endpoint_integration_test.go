package endpoints_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntegrationEndpoints_GetEndpointIdNotFound(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()
	if assert.NoError(t, err) {
		endpoint, err := underTest.GetEndpoint(int64(1234567))
		assert.Error(t, err)
		assert.Nil(t, endpoint)
	}
}