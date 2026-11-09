package queue

import (
	"go-job-scheduler/internal/job"
	"testing"
)

func TestAddJob(t *testing.T) {
	q := Queue{}

	newJob := job.NewJob("Generate Report", "report")

	q.AddJob(newJob)

	if len(q.Jobs) != 1 {
		t.Errorf("Expected 1 job, got %d", len(q.Jobs))
	}
}
