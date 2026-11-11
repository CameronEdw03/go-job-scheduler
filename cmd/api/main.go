package main

import (
	"fmt"
	"go-job-scheduler/internal/job"
	"go-job-scheduler/internal/queue"
	"go-job-scheduler/internal/worker"
	"sync"
)

func main() {
	newJob := job.NewJob("job-001", "database-backup")
	q := queue.Queue{}

	q.AddJob(newJob)
	q.AddJob(job.NewJob("Send Email", "email"))
	q.AddJob(job.NewJob("Generate Invoice", "invoice"))
	q.AddJob(job.NewJob("Data Cleanup", "cleanup"))

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)

		go func(workerID int) {
			defer wg.Done()

			w := worker.Worker{
				ID:    workerID,
				Queue: &q,
			}
			w.ProcessNextJob()
		}(i)
	}

	wg.Wait()
	fmt.Println(newJob)
}
