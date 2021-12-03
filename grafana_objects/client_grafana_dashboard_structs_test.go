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
	assert.NoError(t, err)

	obj := grafana_objects.Timepicker{}
	err = json.Unmarshal([]byte(timepickerFile), &obj)
	assert.NoError(t, err)
}

func Test_ValidateTemplatingStruct(t *testing.T) {
	templatingFile, err := ioutil.ReadFile("testdata/fixtures/templating.json")
	assert.NoError(t, err)

	obj := grafana_objects.TemplatingItem{}
	err = json.Unmarshal([]byte(templatingFile), &obj)
	assert.NoError(t, err)
}

func Test_ValidateDashboardObjectsStruct(t *testing.T) {
	dashboardFile, err := ioutil.ReadFile("testdata/fixtures/dashboard.json")
	assert.NoError(t, err)

	obj := grafana_objects.DashboardObject{}
	err = json.Unmarshal([]byte(dashboardFile), &obj)
	assert.NoError(t, err)
}
