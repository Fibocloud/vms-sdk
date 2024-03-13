package marketing

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func (v *Client) httpRequest(url string, body interface{}) (response []byte, err error) {
	var requestByte []byte
	var requestBody *bytes.Reader
	if body == nil {
		requestBody = bytes.NewReader(nil)
	} else {
		requestByte, _ = json.Marshal(body)
		requestBody = bytes.NewReader(requestByte)
	}

	req, _ := http.NewRequest("POST", v.Endpoint+url, requestBody)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("x-api-key", v.APIKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	response, _ = ioutil.ReadAll(res.Body)
	if res.StatusCode != 200 {
		return nil, errors.New("response")
	}

	return response, nil
}
