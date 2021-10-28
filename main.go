package main

import (
	// "log"
	"errors"
	"fmt"

	"modules/requests"
)

func main() {
	// domain := "http://172.17.0.2:80"
	

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

	r = requests.RequestConfig{
		Url:         domain + "/comment/get",
		RequestType: "GET",
		EstElapse:   100,
	}
	err = requests.PerformRequest(r)
	if err != nil {
		errors.New("Error in" + r.Url + "...")
	}
	fmt.Println(err)

	Headers := make(map[string]string)
	Headers["aaa"] = "application/x-www-form-urlencoded"
	r = requests.RequestConfig {
		Url: domain + "/comment/get",
		RequestType: "GET",
		EstElapse: 100,
		Headers: Headers,
	}
	err = requests.PerformRequest(r)
	if err != nil {
		errors.New("Error in" + r.Url + "...")
	}
	fmt.Println(err)
}
