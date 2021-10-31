package alerts_v2_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"testing"
)

// Get an existing alert
// Status-Code: 200
// Content-Type: application/json
// Actual content: JSON
func TestAlertsV2_GetAlert(t *testing.T) {
	underTest, err, teardown := setupAlertsTest()
	defer teardown()

	alertId := int64(1234567)

	mux.HandleFunc("/v2/alerts/", func(w http.ResponseWriter, r *http.Request) {
		assert.Contains(t, r.URL.String(), strconv.FormatInt(alertId, 10))
		assert.Equal(t, http.MethodGet, r.Method)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("get_alert.json"))
	})

	alert, err := underTest.GetAlert(alertId)

	assert.NoError(t, err)
	assert.NotNil(t, alert)
	assert.Equal(t, alertId, alert.AlertId)
}

// Get a non-existing alert
func TestAlertsV2_GetAlertNotExist(t *testing.T) {
	underTest, err, teardown := setupAlertsTest()
	defer teardown()

	mux.HandleFunc("/v2/alerts/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("get_alert_not_exist.txt"))
	})

	alert, err := underTest.GetAlert(int64(1234567))

	assert.Error(t, err)
	assert.Nil(t, alert)
	assert.Contains(t, err.Error(), "failed with missing alert")
}
