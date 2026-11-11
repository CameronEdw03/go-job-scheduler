package queue

import (
	"go-job-scheduler/internal/job"
	"sync"
)

type Queue struct {
	Jobs  []job.Job
	mutex sync.Mutex
}

func (q *Queue) AddJob(newJob job.Job) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	q.Jobs = append(q.Jobs, newJob)
}

func (q *Queue) GetNextJob() (job.Job, bool) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	if len(q.Jobs) == 0 {
		return job.Job{}, false
	}

	nextJob := q.Jobs[0]
	q.Jobs = q.Jobs[1:]

	return nextJob, true
}
