package worker

import (
	"errors"
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

	err := executeJob(nextJob.Type)

	if err != nil {
		nextJob.UpdateStatus("failed")
		fmt.Println("Job failed:", err)
		return
	}

	nextJob.UpdateStatus("Completed")
	fmt.Println("Worker completed:", nextJob)

}

func executeJob(jobType string) error {
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
		return errors.New("Unknown job type")
	}
	return nil
}
