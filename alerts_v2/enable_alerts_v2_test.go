package alerts_v2_test

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"testing"
)

func TestAlertsV2_EnableAlert(t *testing.T) {
	underTest, err, teardown := setupAlertsTest()
	defer teardown()

	enableAlert := getAlertType()

	mux.HandleFunc("/v2/alerts/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(enableAlert.AlertId, 10))
		w.WriteHeader(http.StatusNoContent)
		w.Header().Set("Content-Type", "application/json")
	})

	alert, err := underTest.EnableAlert(enableAlert)

	assert.NoError(t, err)
	assert.NotNil(t, alert)
}

func TestAlertsV2_EnableAlertNotExist(t *testing.T) {
	underTest, err, teardown := setupAlertsTest()
	defer teardown()

	enableAlert := getAlertType()

	mux.HandleFunc("/v2/alerts/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
	})

	alert, err := underTest.EnableAlert(enableAlert)

	assert.Error(t, err)
	assert.Nil(t, alert)
}
