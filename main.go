package main

import (
	"fmt"

	"modules/export"
	"modules/readers"
	"modules/requests"
)

func worker(jobs <-chan requests.RequestConfig, results chan<- requests.ResponseData) {
	for job := range jobs {
		resp := requests.PerformRequest(job)
		results <- resp
	}
}

func main() {
	// load args
	args := readers.GetArgs()
	requestList := readers.ReadSettings(args.ConfigPath)
	// requestList := readers.ReadSettings("/source/docs/examples.json")

	rounds := requestList.Rounds
	workers := requestList.Workers

	numJobs := len(requestList.Cases)
	jobs := make(chan requests.RequestConfig, numJobs*rounds)
	results := make(chan requests.ResponseData, numJobs*rounds)

	for w := 1; w <= workers; w++ {
		go worker(jobs, results)
	}

	for i := 0; i < numJobs*rounds; i++ {
		d := i % numJobs
		r := requestList.Cases[d]
		r.Url = requestList.Domain + requestList.Cases[d].Route

		jobs <- r
	}
	close(jobs)

	res := []requests.ResponseData{}
	for i := 0; i < numJobs*rounds; i++ {
		// x := <-results
		// fmt.Println(x)
		res = append(res, <-results)
	}

	// wrtie data to a CSV
	CSVErr := export.ToCSV(res, "/source/docs/examples.json")
	if CSVErr != nil {
		fmt.Println(CSVErr)
	}
}
