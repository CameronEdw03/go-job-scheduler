package job

import "time"

type Job struct {
	ID        string
	Type      string
	Status    string
	CreatedAt time.Time
}

func NewJob(id string, jobType string) Job {
	return Job{
		ID:        id,
		Type:      jobType,
		Status:    "pending",
		CreatedAt: time.Now(),
	}
}

func (j *Job) UpdateStatus(status string) {
	j.Status = status
}
