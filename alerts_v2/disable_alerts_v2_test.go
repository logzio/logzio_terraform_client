package alerts_v2_test

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"testing"
)

func TestAlertsV2_DisableAlert(t *testing.T) {
	underTest, err, teardown := setupAlertsTest()
	defer teardown()

	disableAlert := getAlertType()

	mux.HandleFunc("/v2/alerts/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(disableAlert.AlertId, 10))
		w.WriteHeader(http.StatusNoContent)
		w.Header().Set("Content-Type", "application/json")
	})

	alert, err := underTest.EnableAlert(disableAlert)

	assert.NoError(t, err)
	assert.NotNil(t, alert)
}

func TestAlertsV2_DisableAlertNotExist(t *testing.T) {
	underTest, err, teardown := setupAlertsTest()
	defer teardown()

	disableAlert := getAlertType()

	mux.HandleFunc("/v2/alerts/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
	})

	alert, err := underTest.EnableAlert(disableAlert)

	assert.Error(t, err)
	assert.Nil(t, alert)
	assert.Contains(t, err.Error(), "failed with missing alert")
}
