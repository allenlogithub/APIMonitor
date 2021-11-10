package requests

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	FormContentType = "application/x-www-form-urlencoded"
	JsonContentType = "application/json"
)

type AppConfig struct {
	Async  bool            `json:"async"`
	Domain string          `json:"domain"`
	Cases  []RequestConfig `json:"cases"`
}

type RequestConfig struct {
	Url         string
	Route       string            `json:"route"`
	RequestType string            `json:"request_type"`
	EstElapse   int64             `json:"est_elapse"`
	UrlParams   map[string]string `json:"url_params"`
	FormParams  map[string]string `json:"form_params"`
	Headers     map[string]string `json:"headers"`
}

func PerformRequest(requestConfig RequestConfig) error {
	// get body
	body := getBody(requestConfig)

	// prepare request
	request, reqErr := http.NewRequest(requestConfig.RequestType, requestConfig.Url, body)
	if reqErr != nil {
		return errors.New("Error in http.NewRequest, Url:" + requestConfig.Url)
	}

	// add headers
	if len(requestConfig.Headers) != 0 {
		addHeaders(request, requestConfig.Headers)
	}

	// add paramss
	if len(requestConfig.UrlParams) != 0 {
		request.URL.RawQuery = getUrlParams(requestConfig.UrlParams)
	}

	// send request
	client := &http.Client{}
	start := time.Now()
	resp, respErr := client.Do(request)
	if respErr != nil {
		return errors.New("Error in client.Do, Url:" + requestConfig.Url)
	}
	defer resp.Body.Close()

	// response time: elapsed
	elapsed := time.Since(start)
	if (elapsed.Nanoseconds() / 1000000) > requestConfig.EstElapse {
		fmt.Println("Url:", requestConfig.Url, "'s expected elapse time was longer than estamation.")
	}
	fmt.Println("Elapse:", elapsed.Nanoseconds()/1000000, "ms")

	// load return
	if resp.StatusCode == 200 {
		fmt.Println("Test Success, StatusCode:", resp.StatusCode)		
	} else {
		fmt.Println("Test failed, StatusCode:", resp.StatusCode)
	}
	// fmt.Println(responseToString(requestConfig, resp))

	return nil
}

func responseToString(requestConfig RequestConfig, response *http.Response) string {
	buf := new(bytes.Buffer)
	_, bufErr := buf.ReadFrom(response.Body)
	if bufErr != nil {
		return " "
	}
	return buf.String()
}

func addHeaders(req *http.Request, headers map[string]string) {
	for k, v := range headers {
		req.Header.Add(k, v)
	}
}

func getUrlParams(params map[string]string) string {
	urlParams := url.Values{}
	for k, v := range params {
		urlParams.Set(k, v)
	}
	return urlParams.Encode()
}

func getJsonParams(params map[string]string) (io.Reader, error) {
	data, err := json.Marshal(params)
	if err != nil {
		err = errors.New("Invalid Parameters for Content-Type application/json : " + err.Error())
		return nil, err
	}
	return bytes.NewBuffer(data), nil
}

func getBody(cfg RequestConfig) (io.Reader) {
	var body io.Reader
	if len(cfg.FormParams) != 0 {
		switch ct := cfg.Headers["Content-Type"]; ct {
		case "application/x-www-form-urlencoded":
			body = bytes.NewBufferString(getUrlParams(cfg.FormParams))
		case "application/json":
			body, _ = getJsonParams(cfg.FormParams)
		default:
			body = nil
		}
	} else {
		body = nil
	}

	return body
}

// func PerformTest(requestConfig RequestConfig) error {
// 	performRequest
// }
