package scheduler

import (
	"sync"
)

var (
	mutex     sync.Mutex
	JobParams = []JobParam{}
)

type JobParam struct {
	City string
	Pn   int
	Kd   string
}

func NewJobScheduler() *JobParam {
	return &JobParam{}
}

func (j *JobParam) Pop() *JobParam {
	mutex.Lock()
	length := len(JobParams)
	if length < 1 {
		mutex.Unlock()
		return nil
	}

	job := JobParams[length-1]
	JobParams = JobParams[:length-1]

	mutex.Unlock()
	return &job
}

func (j *JobParam) Append(city string, pn int, kd string) {
	mutex.Lock()
	JobParams = append(JobParams, JobParam{
		City: city,
		Pn:   pn,
		Kd:   kd,
	})

	mutex.Unlock()
}

func (j *JobParam) Count() int {
	return len(JobParams)
}
