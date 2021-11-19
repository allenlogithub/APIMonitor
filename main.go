package main

import (
	"fmt"

	"modules/export"
	"modules/jobs"
	"modules/readers"
)

func main() {
	// load args
	args := readers.GetArgs()
	requestList := readers.ReadSettings(args.ConfigPath)

	// run
	var res *[]interface{}
	if requestList.Async {
		res = jobs.Run(&requestList, requestList.Rounds*len(requestList.Cases))
	} else {
		res = jobs.Run(&requestList, requestList.Rounds)
	}

	// wrtie data to a CSV
	CSVErr := export.CSVWrapper(res).ToCSV(args.OutputPath + "/statusGo_raw")
	if CSVErr != nil {
		fmt.Println(CSVErr)
	}
}
