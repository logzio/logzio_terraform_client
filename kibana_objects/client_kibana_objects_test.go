package kibana_objects_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/logzio/logzio_terraform_client/kibana_objects"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
)

func fixture(path string) string {
	b, err := ioutil.ReadFile("testdata/fixtures/" + path)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func setupKibanaObjectsTest() (*kibana_objects.KibanaObjectsClient, error, func()) {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	apiToken := "SOME_API_TOKEN"
	underTest, _ := kibana_objects.New(apiToken, server.URL)

	return underTest, nil, func() {
		server.Close()
	}
}
