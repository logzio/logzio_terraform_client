package grafana_objects_test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/logzio/logzio_terraform_client/grafana_objects"
)

func Test_ValidateTimepickerStruct(t *testing.T) {
	timepickerFile := fixture("timepicker.json")

	obj := grafana_objects.Timepicker{}
	err := json.Unmarshal([]byte(timepickerFile), &obj)
	assert.NoError(t, err)
}

func Test_ValidateTemplatingStruct(t *testing.T) {
	templatingFile := fixture("templating.json")

	obj := grafana_objects.TemplatingItem{}
	err := json.Unmarshal([]byte(templatingFile), &obj)
	assert.NoError(t, err)
}

func Test_ValidateDashboardObjectsStruct(t *testing.T) {
	dashboardFile := fixture("dashboard.json")

	obj := grafana_objects.DashboardObject{}
	err := json.Unmarshal([]byte(dashboardFile), &obj)
	assert.NoError(t, err)
}
