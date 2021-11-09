package main

import (
	"errors"
	"fmt"

	"modules/readers"
	"modules/requests"
)

func main() {
	// load args
	args := readers.GetArgs()
	requestList := readers.ReadSettings(args.ConfigPath)

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
