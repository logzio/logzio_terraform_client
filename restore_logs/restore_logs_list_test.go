package restore_logs_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestRestoreLogs_ListRestoreOperations(t *testing.T) {
	underTest, teardown, err := setupRestoreLogsTest()
	defer teardown()

	mux.HandleFunc(restoreApiBasePath, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("list_restore.json"))
	})

	restores, err := underTest.ListRestoreOperations()
	assert.NoError(t, err)
	assert.NotNil(t, restores)
	assert.Equal(t, 2, len(restores))
}

func TestRestoreLogs_ListRestoreOperationsApiFail(t *testing.T) {
	underTest, teardown, err := setupRestoreLogsTest()
	defer teardown()

	mux.HandleFunc(restoreApiBasePath, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, fixture("restore_api_fail.txt"))
	})

	restores, err := underTest.ListRestoreOperations()
	assert.Error(t, err)
	assert.Nil(t, restores)
}
