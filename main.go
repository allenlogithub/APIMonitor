package main

import (
	// "log"
	"errors"
	"fmt"

	"modules/requests"
)

func main() {
	domain := "http://172.17.0.2:80"

	// case 1
	r := requests.RequestConfig{
		Url:         domain + "/comment/getAll",
		RequestType: "GET",
		EstElapse:   100,
	}
	err := requests.PerformRequest(r)
	if err != nil {
		errors.New("Error in" + r.Url + "...")
	}
	fmt.Println(err)

	// case 2
	UrlParams := make(map[string]string)
	UrlParams["start_id"] = "0"
	UrlParams["results_per_page"] = "2"
	r = requests.RequestConfig{
		Url:         domain + "/comment/get",
		RequestType: "GET",
		EstElapse:   100,
		UrlParams:   UrlParams,
	}
	err = requests.PerformRequest(r)
	if err != nil {
		errors.New("Error in" + r.Url + "...")
	}
	fmt.Println(err)

	// case 3
	Headers, FormParams := make(map[string]string), make(map[string]string)
	Headers["Content-Type"] = "application/json"
	FormParams["name"] = "beauty"
	FormParams["comment"] = "where?"
	r = requests.RequestConfig {
		Url: domain + "/comment/add",
		RequestType: "POST",
		EstElapse: 100,
		Headers: Headers,
		FormParams: FormParams,
	}
	err = requests.PerformRequest(r)
	if err != nil {
		errors.New("Error in" + r.Url + "...")
	}
	fmt.Println(err)

	// case 4
	Headers, FormParams := make(map[string]string), make(map[string]string)
	Headers["Content-Type"] = "application/x-www-form-urlencoded"
	FormParams["A"] = "AAAAA"
	FormParams["B"] = "BBBB"
	r = requests.RequestConfig {
		Url: domain + "/comment/testForm",
		RequestType: "POST",
		EstElapse: 100,
		Headers: Headers,
		FormParams: FormParams,
	}
	err = requests.PerformRequest(r)
	if err != nil {
		errors.New("Error in" + r.Url + "...")
	}
	fmt.Println(err)
}
