package restore_logs_test

import (
	"fmt"
	"github.com/logzio/logzio_terraform_client/restore_logs"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"testing"
)

func TestRestoreLogs_DeleteRestoreOperation(t *testing.T) {
	underTest, teardown, err := setupRestoreLogsTest()
	assert.NoError(t, err)
	defer teardown()

	id := int32(1234)

	mux.HandleFunc(restoreApiBasePath+"/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(id), 10))
		fmt.Fprint(w, fixture("delete_restore.json"))
	})

	deleted, err := underTest.DeleteRestoreOperation(id)
	assert.NoError(t, err)
	assert.NotNil(t, deleted)
	assert.Equal(t, restore_logs.RestoreStatusAborted, deleted.Status)
}

func TestRestoreLogs_DeleteRestoreOperationIdNotFound(t *testing.T) {
	underTest, teardown, err := setupRestoreLogsTest()
	assert.NoError(t, err)
	defer teardown()

	id := int32(1234)

	mux.HandleFunc(restoreApiBasePath+"/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		w.WriteHeader(http.StatusNotFound)
	})

	deleted, err := underTest.DeleteRestoreOperation(id)
	assert.Error(t, err)
	assert.Nil(t, deleted)
}

func TestRestoreLogs_DeleteRestoreOperationApiFail(t *testing.T) {
	underTest, teardown, err := setupRestoreLogsTest()
	assert.NoError(t, err)
	defer teardown()

	id := int32(1234)

	mux.HandleFunc(restoreApiBasePath+"/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		w.WriteHeader(http.StatusInternalServerError)
	})

	deleted, err := underTest.DeleteRestoreOperation(id)
	assert.Error(t, err)
	assert.Nil(t, deleted)
}
