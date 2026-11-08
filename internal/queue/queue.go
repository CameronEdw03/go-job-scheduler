package queue

import "go-job-scheduler/internal/job"

type Queue struct {
	Jobs []job.Job
}

func (q *Queue) AddJob(newJob job.Job) {
	q.Jobs = append(q.Jobs, newJob)
}

func (q *Queue) GetNextJob() (job.Job, bool) {
	if len(q.Jobs) == 0 {
		return job.Job{}, false
	}

	nextJob := q.Jobs[0]

	q.Jobs = q.Jobs[1:]

	return nextJob, true
}
