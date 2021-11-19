package jobs

import (
	"modules/requests"
)

func worker1(jobs <-chan requests.RequestConfig, results chan<- requests.ResponseData) {
	for job := range jobs {
		resp := requests.PerformRequest(job)
		results <- resp
	}
}

func worker2(jobs <-chan []requests.RequestConfig, results chan<- []requests.ResponseData) {
	for job := range jobs {
		var resultList []requests.ResponseData
		for _, step := range job {
			resp := requests.PerformRequest(step)
			resultList = append(resultList, resp)
		}
		results <- resultList
	}
}

func Run(data *requests.AppConfig, size int) *[]interface{} {
	workers := data.Workers
	numJobs := len(data.Cases)

	for i, c := range data.Cases {
		data.Cases[i].Url = data.Domain + c.Route
	}

	var res []interface{}

	switch {
	case data.Async:
		results := make(chan requests.ResponseData, size)
		jobs := make(chan requests.RequestConfig, size)
		for w := 1; w <= workers; w++ {
			go worker1(jobs, results)
		}

		for i := 0; i < size; i++ {
			d := i % numJobs
			r := data.Cases[d]

			jobs <- r
		}
		close(jobs)

		for i := 0; i < size; i++ {
			res = append(res, <-results)
		}
	default:
		results := make(chan []requests.ResponseData, size)
		jobs := make(chan []requests.RequestConfig, size)
		for w := 1; w <= workers; w++ {
			go worker2(jobs, results)
		}

		for i := 0; i < size; i++ {
			jobs <- data.Cases
		}
		close(jobs)

		for i := 0; i < size; i++ {
			for _, result := range <-results {
				res = append(res, result)
			}
		}
	}

	return &res
}
