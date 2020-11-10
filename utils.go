package logzio_client

import (
	"encoding/json"
	"fmt"
	"github.com/logzio/logzio_terraform_client/client"
	"io/ioutil"
	"net/http"
)

const (
	serviceSuccess   int = http.StatusOK
	serviceNoContest int = http.StatusNoContent // This is like StatusOK with no body in response
)

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
	httpClient := client.GetHttpClient(req)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	jsonBytes, err := ioutil.ReadAll(resp.Body)
	if !CheckValidStatus(resp, []int{serviceSuccess, serviceNoContest}) {
		return nil, fmt.Errorf("%d %s", resp.StatusCode, jsonBytes)
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
