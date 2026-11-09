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
	fmt.Println("Worker processing:", nextJob)

}
