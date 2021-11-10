package main

import (
	"fmt"

	"modules/readers"
	"modules/requests"
)

type (
	config = requests.RequestConfig
	respD = requests.ResponseData
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
	jobs := make(chan config, numJobs*rounds)
	results := make(chan respD, numJobs*rounds)

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

	for i := 0; i < numJobs*rounds; i++ {
		// <-results
		x := <-results
		fmt.Println(x)
	}
}
