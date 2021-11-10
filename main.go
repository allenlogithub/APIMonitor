package main

import (
	// "errors"
	"fmt"

	"modules/readers"
	"modules/requests"
)

type (
	config = requests.RequestConfig
)

func worker(jobs <-chan requests.RequestConfig, results chan<- string) {
	for job := range jobs {
		err := requests.PerformRequest(job)
		if err != nil {
			// errors.New("Error in" + job.Url + "...")
			results <- "Failed  " + err.Error()
		} else {
			results <- "Success " + job.Url
		}
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
	results := make(chan string, numJobs*rounds)

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
