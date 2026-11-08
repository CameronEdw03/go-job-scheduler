package main

import (
	"fmt"
	"go-job-scheduler/internal/job"
)

func main() {
	newJob := job.NewJob("job-001", "database-backup")

	fmt.Println(newJob)
}
