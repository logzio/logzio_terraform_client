package logzio_client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const deleteServiceUrl string = "%s/v1/alerts/%d"
const deleteServiceMethod string = "DELETE"

func buildDeleteApiRequest(apiToken string, alertId int64) (*http.Request, error) {
	baseUrl := getLogzioBaseUrl()
	req, err := http.NewRequest(deleteServiceMethod, fmt.Sprintf(deleteServiceUrl, baseUrl, alertId), nil)
	addHttpHeaders(apiToken, req)
	log.Printf("%s::%s", "buildDeleteApiRequest", req.URL.Path)

	return req, err
}

func (c *Client) DeleteAlert(alertId int64) error {
	req, _ := buildDeleteApiRequest(c.apiToken, alertId)

	var client http.Client
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	data, _ := ioutil.ReadAll(resp.Body)
	s, _ := prettyprint(data)
	log.Printf("%s::%s", "DeleteAlert", s)


	if !checkValidStatus(resp, []int{200}) {
		return fmt.Errorf("API call %s failed with status code %d, data: %s", "DeleteAlert", resp.StatusCode, s)
	}

	str := fmt.Sprintf("%s", s)
	if strings.Contains(str, "no alert id") {
		return fmt.Errorf("API call %s failed with missing alert %d, data: %s", "DeleteAlert", alertId, s)
	}

	return nil
}
