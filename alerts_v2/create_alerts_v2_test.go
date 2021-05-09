package alerts_v2_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/logzio/logzio_terraform_client/alerts_v2"
	"github.com/stretchr/testify/assert"
)

func TestAlertsV2_CreateAlert(t *testing.T) {
	underTest, err, teardown := setupAlertsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v2/alerts", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)

			jsonBytes, _ := ioutil.ReadAll(r.Body)
			var target alerts_v2.CreateAlertType
			err = json.Unmarshal(jsonBytes, &target)
			assert.NotNil(t, target)
			assert.NotEmpty(t, target.Title)
			assert.NotEmpty(t, target.Description)
			assert.NotEmpty(t, target.SubComponents[0].QueryDefinition.Query)
			assert.NotEmpty(t, target.SubComponents[0].Trigger.SeverityThresholdTiers)
			assert.NotEmpty(t, target.Tags)
			assert.Equal(t, 2, len(target.Tags))

			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("create_alert.json"))
			w.WriteHeader(http.StatusOK)
		})

		testAlert := getCreateAlertType()

		alert, err := underTest.CreateAlert(testAlert)
		assert.NoError(t, err)
		assert.Equal(t, int64(1234567), alert.AlertId)
		assert.Equal(t, float32(10.0), alert.SubComponents[0].Trigger.SeverityThresholdTiers[alerts_v2.SeverityHigh])
		assert.Equal(t, float32(5.0), alert.SubComponents[0].Trigger.SeverityThresholdTiers[alerts_v2.SeverityInfo])
		assert.Equal(t, "some", alert.Tags[0])
		assert.Equal(t, "words", alert.Tags[1])
	}
}

func TestAlertsV2_CreateAlertAPIFail(t *testing.T) {
	underTest, err, teardown := setupAlertsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v2/alerts", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("create_alert_failed.txt"))
		})
	}

	_, err = underTest.CreateAlert(getCreateAlertType())
	assert.Error(t, err)
}

func TestAlertsV2_CreateAlertNoTitle(t *testing.T) {
	underTest, err, teardown := setupAlertsTest()
	defer teardown()

	createAlertType := getCreateAlertType()
	createAlertType.Title = ""
	_, err = underTest.CreateAlert(createAlertType)
	assert.Error(t, err)
}
