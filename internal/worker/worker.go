package worker

import (
	"fmt"
	"go-job-scheduler/internal/queue"
)

type Worker struct {
	ID    int
	Queue *queue.Queue
}

func (w *Worker) ProcessNextJob() {
	nextJob, found := w.Queue.GetNextJob()

	if !found {
		return
	}
	nextJob.UpdateStatus("processing")
	fmt.Println("Worker processing:", nextJob)

	executeJob(nextJob.Type)

	nextJob.UpdateStatus("Completed")
	fmt.Println("Worker completed:", nextJob)

}

func executeJob(jobType string) {
	switch jobType {
	case "database-backup":
		fmt.Println("Running database backup")

	case "email":
		fmt.Println("Sending email")

	case "invoice":
		fmt.Println("Generating invoice")

	case "cleanup":
		fmt.Println("Cleaning up data")

	default:
		fmt.Println("Unknown job type")
	}
}
