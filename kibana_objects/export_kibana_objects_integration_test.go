package kibana_objects_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntegrationKibanaObjects_ExportKibanaObject(t *testing.T) {
	underTest, err := setupKibanaObjectsIntegrationTest()

	if assert.NoError(t, err) {
		exportRequest := getExportRequest()
		exportResponse, err := underTest.ExportKibanaObject(exportRequest)
		assert.NoError(t, err)
		assert.NotNil(t, exportResponse)
	}
}
