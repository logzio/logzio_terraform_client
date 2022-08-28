package restore_logs_test

import (
	"encoding/json"
	"fmt"
	"github.com/logzio/logzio_terraform_client/restore_logs"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestRestoreLogs_InitiateRestore(t *testing.T) {
	underTest, teardown, err := setupRestoreLogsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc(restoreApiBasePath, func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			jsonBytes, _ := ioutil.ReadAll(r.Body)
			var target restore_logs.InitiateRestore
			err = json.Unmarshal(jsonBytes, &target)
			assert.NoError(t, err)
			assert.NotNil(t, target)
			assert.NotEmpty(t, target.AccountName)
			assert.NotEmpty(t, target.UserName)
			assert.NotZero(t, target.StartTime)
			assert.NotZero(t, target.EndTime)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("initiate_restore.json"))
		})

		initiateRestore := getInitiateRestoreOperationTest()
		restore, err := underTest.InitiateRestoreOperation(initiateRestore)
		assert.NoError(t, err)
		assert.NotEmpty(t, restore)
		assert.NotZero(t, restore.Id)
		assert.Equal(t, initiateRestore.AccountName, restore.AccountName)
		assert.Equal(t, initiateRestore.StartTime, int64(restore.StartTime))
		assert.Equal(t, initiateRestore.EndTime, int64(restore.EndTime))
		assert.Equal(t, restore_logs.RestoreStatusInProgress, restore.Status)
	}
}

func TestRestoreLogs_InitiateRestoreApiFail(t *testing.T) {
	underTest, teardown, err := setupRestoreLogsTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc(restoreApiBasePath, func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			jsonBytes, _ := ioutil.ReadAll(r.Body)
			var target restore_logs.InitiateRestore
			err = json.Unmarshal(jsonBytes, &target)
			assert.NoError(t, err)
			assert.NotNil(t, target)
			assert.NotEmpty(t, target.AccountName)
			assert.NotZero(t, target.StartTime)
			assert.NotZero(t, target.EndTime)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, fixture("restore_api_fail.txt"))
		})

		initiateRestore := getInitiateRestoreOperationTest()
		restore, err := underTest.InitiateRestoreOperation(initiateRestore)
		assert.Error(t, err)
		assert.Nil(t, restore)
	}
}
