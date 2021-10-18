package archive_logs_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestArchiveLogs_ListArchives(t *testing.T) {
	underTest, err, teardown := setupArchiveLogsTest()
	defer teardown()

	mux.HandleFunc(archiveApiBasePath, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, fixture("list_archives.json"))
	})

	archives, err := underTest.ListArchiveLog()
	assert.NoError(t, err)
	assert.NotNil(t, archives)
	assert.Equal(t, 1, len(archives))
}

func TestArchiveLogs_ListArchivesApiFail(t *testing.T) {
	underTest, err, teardown := setupArchiveLogsTest()
	defer teardown()

	mux.HandleFunc("/v2/archive/settings", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, fixture("archive_api_fail.txt"))
	})

	archives, err := underTest.ListArchiveLog()
	assert.Error(t, err)
	assert.Nil(t, archives)
}
