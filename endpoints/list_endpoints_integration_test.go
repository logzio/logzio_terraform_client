package endpoints_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntegrationEndpoints_RetrieveEndpoints(t *testing.T) {
	underTest, err := setupEndpointsIntegrationTest()
	if assert.NoError(t, err) {
		endpoints, err := underTest.ListEndpoints()
		assert.NoError(t, err)
		assert.NotNil(t, endpoints)
	}
}
