package archive_logs_test

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"testing"
)

func TestArchiveLogs_DeleteArchive(t *testing.T) {
	underTest, teardown, err := setupArchiveLogsTest()
	assert.NoError(t, err)
	defer teardown()

	id := int32(1234)

	mux.HandleFunc(archiveApiBasePath+"/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Contains(t, r.URL.String(), strconv.FormatInt(int64(id), 10))
		w.WriteHeader(http.StatusCreated)
	})

	err = underTest.DeleteArchiveLogs(id)
	assert.NoError(t, err)
}

func TestArchiveLogs_DeleteArchiveIdNotFound(t *testing.T) {
	underTest, teardown, err := setupArchiveLogsTest()
	assert.NoError(t, err)
	defer teardown()

	id := int32(1234)

	mux.HandleFunc(archiveApiBasePath+"/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		w.WriteHeader(http.StatusNotFound)
	})

	err = underTest.DeleteArchiveLogs(id)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed with missing archive")
}

func TestArchiveLogs_DeleteArchiveApiFail(t *testing.T) {
	underTest, teardown, err := setupArchiveLogsTest()
	assert.NoError(t, err)
	defer teardown()

	id := int32(1234)

	mux.HandleFunc(archiveApiBasePath+"/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		w.WriteHeader(http.StatusInternalServerError)
	})

	err = underTest.DeleteArchiveLogs(id)
	assert.Error(t, err)
}
