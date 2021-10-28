package logzio_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/logzio/logzio_terraform_client/client"
	"io/ioutil"
	"net/http"
)

const (
	serviceSuccess   int = http.StatusOK
	serviceNoContent int = http.StatusNoContent // This is like StatusOK with no body in response
	serviceCreated   int = http.StatusCreated
)

type LogzioApiCallDetails struct {
	ApiToken     string
	HttpMethod   string
	Url          string
	Body         []byte
	SuccessCodes []int
	NotFoundCode int
	ResourceId   interface{}
	ApiAction    string
	ResourceName string
}

func AddHttpHeaders(apiToken string, req *http.Request) {
	req.Header.Add("X-API-TOKEN", apiToken)
	req.Header.Add("Content-Type", "application/json")
}

func Contains(slice []string, s string) bool {
	for _, value := range slice {
		if value == s {
			return true
		}
	}
	return false
}

func CheckValidStatus(response *http.Response, status []int) bool {
	for x := 0; x < len(status); x++ {
		if response.StatusCode == status[x] {
			return true
		}
	}
	return false
}

func CreateHttpRequest(req *http.Request) (map[string]interface{}, error) {
	jsonBytes, err := CreateHttpRequestBytesResponse(req)
	if err != nil {
		return nil, err
	}

	var target map[string]interface{}
	if len(jsonBytes) > 0 {
		err = json.Unmarshal(jsonBytes, &target)
		if err != nil {
			return nil, err
		}
	}
	return target, nil
}

func CreateHttpRequestBytesResponse(req *http.Request) ([]byte, error) {
	httpClient := client.GetHttpClient(req)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	jsonBytes, err := ioutil.ReadAll(resp.Body)

	if !CheckValidStatus(resp, []int{serviceSuccess, serviceNoContent, serviceCreated}) {
		return nil, fmt.Errorf("%d %s", resp.StatusCode, jsonBytes)
	}

	return jsonBytes, nil
}

func CallLogzioApi(logzioCall LogzioApiCallDetails) ([]byte, error) {
	req, err := buildApiRequest(logzioCall.ApiToken, logzioCall.HttpMethod, logzioCall.Url, logzioCall.Body)
	if err != nil {
		return nil, err
	}

	httpClient := client.GetHttpClient(req)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	jsonBytes, _ := ioutil.ReadAll(resp.Body)
	if !CheckValidStatus(resp, logzioCall.SuccessCodes) {
		if resp.StatusCode == logzioCall.NotFoundCode {
			return nil, fmt.Errorf("API call %s failed with missing %s %d, data: %s",
				logzioCall.ApiAction, logzioCall.ResourceName, logzioCall.ResourceId, jsonBytes)
		}

		return nil, fmt.Errorf("API call %s failed with status code %d, data: %s",
			logzioCall.ApiAction, resp.StatusCode, jsonBytes)
	}

	return jsonBytes, nil
}

func buildApiRequest(apiToken string, method string, url string, body []byte) (*http.Request, error) {
	var req *http.Request
	var err error
	if body == nil {
		req, err = http.NewRequest(method, url, nil)

	} else {
		req, err = http.NewRequest(method, url, bytes.NewBuffer(body))
	}

	if err != nil {
		return nil, err
	}

	AddHttpHeaders(apiToken, req)
	return req, err
}
