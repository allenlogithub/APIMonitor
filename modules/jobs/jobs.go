package jobs

import (
	"modules/requests"
)

func worker(jobs <-chan requests.RequestConfig, results chan<- requests.ResponseData) {
	for job := range jobs {
		resp := requests.PerformRequest(job)
		results <- resp
	}
}

// func worker2(jobs <-chan requests.AppConfig.Cases, results chan<- []requests.ResponseData) {
// 	var resultList []results.ResponseData
// 	for job := range jobs {
// 		for step := range job {
// 			resp := requests.PerformRequest(step)
// 			resultList = append(resultList, resp)
// 		}
// 	}
// }

func Run(data requests.AppConfig, size int) *[]interface{} {
	workers := data.Workers
	rounds := data.Rounds
	numJobs := len(data.Cases)

	jobs := make(chan requests.RequestConfig, size)
	results := make(chan requests.ResponseData, size)

	for w := 1; w <= workers; w++ {
		go worker(jobs, results)
	}

	for i := 0; i < numJobs*rounds; i++ {
		d := i % numJobs
		r := data.Cases[d]
		r.Url = data.Domain + data.Cases[d].Route

		jobs <- r
	}
	close(jobs)

	var res []interface{}
	for i := 0; i < numJobs*rounds; i++ {
		res = append(res, <-results)
	}

	return &res
}

