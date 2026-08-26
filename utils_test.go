package logzio_client

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCallLogzioApi_TransportErrorDoesNotPanic(t *testing.T) {
	res, err := CallLogzioApi(LogzioApiCallDetails{
		ApiToken:     "token",
		HttpMethod:   http.MethodGet,
		Url:          "http://127.0.0.1:1/",
		SuccessCodes: []int{http.StatusOK},
		NotFoundCode: http.StatusNotFound,
		ApiAction:    "Test",
		ResourceName: "test",
	})

	assert.Error(t, err)
	assert.Nil(t, res)
}
