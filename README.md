# Distributed Job Scheduler

A distributed job scheduling system built with **Go** that manages, queues, and executes background tasks using a worker-based architecture.

The goal of this project is to build a scalable job processing system capable of handling individual jobs and batch workloads while exploring distributed systems, concurrency, and cloud-native infrastructure.

## Project Goals

- Build a job queue to manage pending tasks.
- Develop workers that retrieve and execute jobs.
- Support individual and batch job submissions.
- Implement concurrent job processing.
- Track job status, execution history, and failures.
- Introduce retry mechanisms for failed jobs.
- Deploy the system using Docker and Kubernetes.
- Implement CI/CD pipelines for automated testing and deployment.

## Planned Architecture

The system will eventually consist of the following components:

| Component | Responsibility |
|---|---|
| API | Accepts job and batch submissions |
| Job Manager | Creates and manages jobs |
| Queue | Stores pending jobs |
| Workers | Retrieve and execute jobs |
| Scheduler | Controls job execution and scheduling |
| Database | Persists job information and execution history |

## Technology Stack

- **Language:** Go
- **API:** Go standard library / HTTP
- **Queue:** In-memory queue initially, Redis planned
- **Database:** PostgreSQL planned
- **Containerization:** Docker
- **Orchestration:** Kubernetes
- **CI/CD:** GitHub Actions
- **Monitoring:** Prometheus and Grafana planned

## Project Structure

```text
go-job-scheduler/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── job/
│   ├── queue/
│   └── worker/
├── go.mod
└── README.md
```

## Planned Features

- [x] Initial Go project setup
- [x] Job data structure
- [ ] Job creation and management
- [ ] In-memory job queue
- [ ] Worker implementation
- [ ] Concurrent job execution
- [ ] Batch job processing
- [ ] Job status tracking
- [ ] Retry and failure handling
- [ ] REST API endpoints
- [ ] Redis-backed distributed queue
- [ ] PostgreSQL persistence
- [ ] Docker containerization
- [ ] Kubernetes deployment
- [ ] CI/CD pipeline
- [ ] Monitoring and observability

## Running Locally

Clone the repository:

```bash
git clone YOUR_REPOSITORY_URL
cd go-job-scheduler
```

Run the application:

```bash
go run ./cmd/api
```

Build the project:

```bash
go build ./...
```

## Learning Objectives

This project is also an opportunity to explore:

- Go programming and application structure
- Goroutines and concurrency
- Distributed job processing
- Queue management
- Fault tolerance and retries
- Containerization and Kubernetes
- CI/CD and observability

---

**Status:** In Development
