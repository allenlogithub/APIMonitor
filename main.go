package main

import (
	"errors"
	"fmt"

	"modules/requests"
	"modules/readers"
)

func main() {
	// load request config
	requestList := readers.ReadSettings("/source/docs/examples.json")

	// perform tests
	for i := 0; i < len(requestList.Cases); i++ {
		r := requestList.Cases[i]
		r.Url = requestList.Domain + requestList.Cases[i].Route
		err := requests.PerformRequest(r)
		if err != nil {
			errors.New("Error in" + r.Url + "...")
		}
		fmt.Println(err)
	}
}
