package alerts_v2_test

import (
	"encoding/json"
	"fmt"
	"github.com/logzio/logzio_terraform_client/alerts_v2"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"strconv"
	"testing"
)

func TestAlertsV2_UpdateAlert(t *testing.T) {
	underTest, err, teardown := setupAlertsTest()
	defer teardown()

	alertId := int64(1234567)

	mux.HandleFunc("/v2/alerts/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(alertId, 10))

		jsonBytes, _ := io.ReadAll(r.Body)
		var target alerts_v2.CreateAlertType
		err = json.Unmarshal(jsonBytes, &target)
		assert.NotNil(t, target)
		assert.NotEmpty(t, target.Title)
		assert.NotEmpty(t, target.Description)
		assert.NotEmpty(t, target.SubComponents[0].QueryDefinition.Query)
		assert.NotEmpty(t, target.SubComponents[0].Trigger.SeverityThresholdTiers)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("update_alert.json"))
	})

	updateAlertRequest := getCreateAlertType()
	updateAlertRequest.Title = "test update alert"

	updatedAlert, err := underTest.UpdateAlert(alertId, updateAlertRequest)

	assert.NoError(t, err)
	assert.NotNil(t, updatedAlert)
}

func TestAlerts_UpdateAlert_NotExist(t *testing.T) {
	underTest, err, teardown := setupAlertsTest()
	defer teardown()

	alertId := int64(1234567)

	mux.HandleFunc("/v2/alerts/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("update_alert_not_exist.txt"))
	})

	updateAlertRequest := getCreateAlertType()
	updatedAlert, err := underTest.UpdateAlert(alertId, updateAlertRequest)

	assert.Error(t, err)
	assert.Nil(t, updatedAlert)
	assert.Contains(t, err.Error(), "failed with missing alert")
}
