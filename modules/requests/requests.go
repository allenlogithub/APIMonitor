package requests

import (
	"net/http"
	// "log"
	"bytes"
	"errors"
	"fmt"
	"time"
)

const (
	FormContentType = "application/x-www-form-urlencoded"
	JsonContentType = "application/json"
)

type RequestConfig struct {
	Url         string
	RequestType string
	EstElapse   int64
	UrlParams   map[string]string
	Headers     map[string]string
}

func PerformRequest(requestConfig RequestConfig) error {
	//  prepare request
	request, reqErr := http.NewRequest(requestConfig.RequestType, requestConfig.Url, nil)
	if reqErr != nil {
		return errors.New("Error in http.NewRequest, Url:" + requestConfig.Url)
	}

	fmt.Println(len(requestConfig.Headers))
	if len(requestConfig.Headers) != 0 {
		fmt.Println("INNNN")
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
	fmt.Println(elapsed.Nanoseconds()/1000000, "ms")
	// load return
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

// func PerformTest(requestConfig RequestConfig) error {
// 	performRequest
// }
