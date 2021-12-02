package grafana_objects_test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"

	"github.com/logzio/logzio_terraform_client/grafana_objects"
)

func Test_ValidateTimepickerStruct(t *testing.T) {
	timepickerFile, err := ioutil.ReadFile("testdata/fixtures/timepicker.json")

	resp := grafana_objects.Timepicker{}
	err = json.Unmarshal([]byte(timepickerFile), &resp)
	assert.NoError(t, err)
}

func Test_ValidateTemplatingStruct(t *testing.T) {
	templatingFile, err := ioutil.ReadFile("testdata/fixtures/templating.json")

	resp := grafana_objects.TemplatingItem{}
	err = json.Unmarshal([]byte(templatingFile), &resp)
	assert.NoError(t, err)
}

func Test_ValidateDashboardObjectsStruct(t *testing.T) {
	dashboardFile, err := ioutil.ReadFile("testdata/fixtures/dashboard.json")

	resp := grafana_objects.DashboardObject{}
	err = json.Unmarshal([]byte(dashboardFile), &resp)
	assert.NoError(t, err)
}
