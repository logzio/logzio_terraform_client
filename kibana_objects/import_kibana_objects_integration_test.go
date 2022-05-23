package kibana_objects_test

import (
	"github.com/logzio/logzio_terraform_client/kibana_objects"
	"github.com/logzio/logzio_terraform_client/test_utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntegrationKibanaObjects_ImportKibanaObject(t *testing.T) {
	underTest, err := setupKibanaObjectImportIntegrationTest()

	if assert.NoError(t, err) {
		importKibanaObjectRequest, err := getImportRequest()
		if assert.NoError(t, err) {
			kibanaObject, err := underTest.ImportKibanaObject(importKibanaObjectRequest)
			if assert.NoError(t, err) && assert.NotNil(t, kibanaObject) && assert.NotZero(t, len(kibanaObject.Created)) {
				// test with override:
				override := true
				importKibanaObjectRequest.Override = &override
				kibanaObjectOverride, err := underTest.ImportKibanaObject(importKibanaObjectRequest)
				assert.NoError(t, err)
				assert.NotNil(t, kibanaObjectOverride)
				assert.NotZero(t, len(kibanaObjectOverride.Updated))

				// At the moment cannot delete the object via API. Need to delete manually.
			}
		}
	}
}

func setupKibanaObjectImportIntegrationTest() (*kibana_objects.KibanaObjectsClient, error) {
	apiToken, err := test_utils.GetApiToken()
	if err != nil {
		return nil, err
	}
	underTest, err := kibana_objects.New(apiToken, test_utils.GetLogzIoBaseUrl())
	return underTest, nil
}
