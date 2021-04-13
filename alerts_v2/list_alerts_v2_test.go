package alerts_v2_test

import (
"fmt"
"github.com/stretchr/testify/assert"
"net/http"
"testing"
)

func TestAlertsV2_ListAlerts(t *testing.T) {
	underTest, err, teardown := setupAlertsTest()
	defer teardown()

	mux.HandleFunc("/v2/alerts", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("list_alerts.json"))
	})

	alerts, err := underTest.ListAlerts()

	assert.NoError(t, err)
	assert.NotNil(t, alerts)
	assert.NotEmpty(t, alerts)
}
