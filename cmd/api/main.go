package main

import (
	"fmt"
	"go-job-scheduler/internal/job"
	"go-job-scheduler/internal/queue"
	"go-job-scheduler/internal/worker"
)

func main() {
	newJob := job.NewJob("job-001", "database-backup")
	q := queue.Queue{}

	q.AddJob(newJob)
	q.AddJob(job.NewJob("Send Email", "email"))
	q.AddJob(job.NewJob("Generate Invoice", "invoice"))
	q.AddJob(job.NewJob("Data Cleanup", "cleanup"))

	for i := 1; i <= 3; i++ {
		w := worker.Worker{
			ID:    1,
			Queue: &q,
		}
		w.ProcessNextJob()
	}

	fmt.Println(newJob)
}
