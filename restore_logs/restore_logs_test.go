package restore_logs_test

import (
	"fmt"
	"github.com/logzio/logzio_terraform_client/archive_logs"
	"github.com/logzio/logzio_terraform_client/restore_logs"
	"github.com/logzio/logzio_terraform_client/test_utils"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"time"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
)

const (
	withArchive    = true
	withoutArchive = false

	restoreApiBasePath = "/archive/restore"
)

/* setupRestoreLogsIntegrationTest sets up the resources that are needed to test the Restore API.
The function retrieves the api token, base url, and creates a new restore logs client.
To initiate a restore operation, an active archive needs to be connected to the Logz.io account, so the function also
creates an archive, and retrieves the function that deletes it.
!! NOTE that it is the caller's responsibility to use the delete function that's being returned. !!
*/
func setupRestoreLogsIntegrationTest(createArchive bool) (*restore_logs.RestoreClient, func(), error) {
	apiToken, err := test_utils.GetApiToken()
	if err != nil {
		return nil, nil, err
	}
	baseUrl := test_utils.GetLogzIoBaseUrl()
	underTest, err := restore_logs.New(apiToken, baseUrl)
	if err != nil {
		return nil, nil, err
	}

	if createArchive {
		deleteArchiveFunc, err := createTestArchiveDeleteFunc(apiToken, baseUrl)
		return underTest, deleteArchiveFunc, err
	}

	return underTest, nil, nil
}

// createTestArchiveDeleteFunc creates a test archive, and return a function that deletes that account
func createTestArchiveDeleteFunc(apiToken string, baseUrl string) (func(), error) {
	createArchive, err := test_utils.GetCreateOrUpdateArchiveLogs(archive_logs.StorageTypeS3)
	if err != nil {
		return nil, err
	}

	archiveClient, err := archive_logs.New(apiToken, baseUrl)
	if err != nil {
		return nil, err
	}

	archive, err := archiveClient.SetupArchive(createArchive)
	if err != nil {
		return nil, err
	}

	// allow some time for the archive to be created
	time.Sleep(1 * time.Second)
	return func() {
		err := archiveClient.DeleteArchiveLogs(archive.Id)
		if err != nil {
			log.Printf("error occurred while trying to delete test archive for restore tests")
			log.Printf("archive id: %d", archive.Id)
			log.Printf("error: %s", err.Error())
		}
	}, nil
}

func getInitiateRestoreOperationIntegrationTest() restore_logs.InitiateRestore {
	currentTime := time.Now()
	hourAgo := currentTime.Add(-time.Hour)
	return restore_logs.InitiateRestore{
		AccountName: fmt.Sprintf("test-account-%s", currentTime.Format("2006-01-02,15:04:05")),
		UserName:    os.Getenv(test_utils.EnvLogzioEmail),
		StartTime:   hourAgo.Unix(),
		EndTime:     currentTime.Unix(),
	}
}

func getInitiateRestoreOperationTest() restore_logs.InitiateRestore {
	return restore_logs.InitiateRestore{
		AccountName: "test_account",
		UserName:    "test@test.com",
		StartTime:   1634437185,
		EndTime:     1634444385,
	}
}

func setupRestoreLogsTest() (*restore_logs.RestoreClient, func(), error) {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	apiToken := "SOME_API_TOKEN"
	underTest, err := restore_logs.New(apiToken, server.URL)

	return underTest, func() {
		server.Close()
	}, err
}

func fixture(path string) string {
	b, err := ioutil.ReadFile("testdata/fixtures/" + path)
	if err != nil {
		panic(err)
	}
	return string(b)
}
