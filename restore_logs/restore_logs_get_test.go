package restore_logs_test

import (
	"fmt"
	"github.com/logzio/logzio_terraform_client/restore_logs"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"testing"
)

func TestRestoreLogs_GetRestoreOperation(t *testing.T) {
	underTest, teardown, err := setupRestoreLogsTest()
	defer teardown()
	assert.NoError(t, err)
	restoreId := int32(1234)

	mux.HandleFunc(restoreApiBasePath+"/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(restoreId), 10))
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("get_restore.json"))
	})

	restore, err := underTest.GetRestoreOperation(restoreId)
	assert.NoError(t, err)
	assert.NotNil(t, restore)
	assert.Equal(t, restoreId, restore.Id)
	assert.Equal(t, "test_account", restore.AccountName)
	assert.Equal(t, int32(1234567), restore.AccountId)
	assert.Equal(t, 2.0, restore.RestoredVolumeGb)
	assert.Equal(t, restore_logs.RestoreStatusActive, restore.Status)
	assert.Equal(t, 1634437185.000000000, restore.StartTime)
	assert.Equal(t, 1634444385.000000000, restore.EndTime)
	assert.Equal(t, 1634479972.000000000, restore.CreatedAt)
	assert.Equal(t, 1634479976.000000000, restore.StartedAt)
	assert.Equal(t, 1634480039.000000000, restore.FinishedAt)
	assert.Equal(t, 1634912039.000000000, restore.ExpiresAt)
}

func TestRestoreLogs_GetRestoreOperationIdNotFound(t *testing.T) {
	underTest, teardown, err := setupRestoreLogsTest()
	defer teardown()
	assert.NoError(t, err)
	restoreId := int32(1234)

	mux.HandleFunc(restoreApiBasePath+"/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(restoreId), 10))
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, fixture("get_restore_id_not_found.txt"))
	})

	restore, err := underTest.GetRestoreOperation(restoreId)
	assert.Error(t, err)
	assert.Nil(t, restore)
}

func TestRestoreLogs_GetRestoreOperationApiFail(t *testing.T) {
	underTest, teardown, err := setupRestoreLogsTest()
	defer teardown()
	assert.NoError(t, err)
	restoreId := int32(1234)

	mux.HandleFunc(restoreApiBasePath+"/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(restoreId), 10))
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, fixture("restore_api_fail.txt"))
	})

	restore, err := underTest.GetRestoreOperation(restoreId)
	assert.Error(t, err)
	assert.Nil(t, restore)
}
