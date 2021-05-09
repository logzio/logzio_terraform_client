package alerts_v2_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"testing"
)

func TestAlertsV2_DeleteAlert(t *testing.T) {
	underTest, err, teardown := setupAlertsTest()
	defer teardown()

	alertId := int64(1234567)

	mux.HandleFunc("/v2/alerts/", func(w http.ResponseWriter, r *http.Request) {
		assert.Contains(t, r.URL.String(), strconv.FormatInt(alertId, 10))
		assert.Equal(t, http.MethodDelete, r.Method)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("delete_alert.txt"))
		w.WriteHeader(http.StatusOK)
	})

	err = underTest.DeleteAlert(alertId)
	assert.NoError(t, err)
}

func TestAlertsV2_DeleteMissingAlert(t *testing.T) {
	underTest, err, teardown := setupAlertsTest()
	defer teardown()

	mux.HandleFunc("/v2/alerts/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("delete_alert_not_exist.txt"))
	})

	err = underTest.DeleteAlert(int64(1234567))
	assert.Error(t, err)
}