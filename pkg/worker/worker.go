package worker

import (
	"context"
	"fmt"
	"os/exec"
	"sync"
	"syscall"

	logger "log"

	c "github.com/jaylane/job-scheduler/pkg/worker/config"
	log "github.com/jaylane/job-scheduler/pkg/worker/log"

	"errors"

	"github.com/google/uuid"
	j "github.com/jaylane/job-scheduler/pkg/job"
)

// Worker interface defines the methods for manipulating jobs
type Worker interface {
	// StartJob starts a job with the given command and arguments
	StartJob(command j.Command) (jobID string, err error)
	// StopJob stops a job with the given jobID
	StopJob(jobID string) (status j.Status, err error)
	// GetJobStatus returns the status of a job with the given jobID
	GetJobStatus(jobID string) (status j.Status, err error)
	// StreamJobOutput streams the output of a job with the given jobID
	StreamJobOutput(ctx context.Context, jobID string) (outchan chan string, err error)
}

type worker struct {
	logger log.Logger
	jobs   map[string]*j.Job
	m      sync.RWMutex
}

// NewWorker returns a new Worker instance
func NewWorker(conf c.Config) Worker {
	return &worker{
		logger: log.NewLogger(c.Config{}),
		jobs:   make(map[string]*j.Job),
	}
}

func (w *worker) StartJob(command j.Command) (jobID string, err error) {
	logger.Printf("Starting job with command %s and args %v", command.Name, command.Args)
	cmd := exec.Command(command.Name, command.Args...)
	job := *&j.Job{
		ID:     uuid.NewString(),
		Cmd:    cmd,
		Output: make([]byte, 0),
		Process: &j.Process{
			PID: cmd.ProcessState.Pid(),
		},
	}
	jobID = job.ID
	w.m.Lock()
	w.jobs[jobID] = &job
	w.m.Unlock()

	go func() {
		if err := job.Cmd.Run(); err != nil {
			logger.Printf("Error running job: %s", err)
		}

		process := j.Process{
			PID:      job.Cmd.ProcessState.Pid(),
			ExitCode: job.Cmd.ProcessState.ExitCode(),
		}
		w.m.Lock()
		job.Process = &process
		w.m.Unlock()
	}()

	return jobID, nil
}

func (w *worker) StopJob(jobID string) (status j.Status, err error) {
	w.m.RLock()
	defer w.m.RUnlock()

	job, err := w.getJob(jobID)
	if err != nil {
		return j.UNKNOWN, err
	}

	if job.IsRunning() {
		job.Cmd.Process.Signal(syscall.SIGTERM)
		w.m.Lock()
		job.Process.Status = j.TERMINATED
		w.m.Unlock()
		return job.Process.Status, nil
	}

	return j.STOPPED, errors.New("job has already stopped")
}

func (w *worker) GetJobStatus(jobID string) (status j.Status, err error) {
	w.m.RLock()
	defer w.m.RUnlock()
	job, err := w.getJob(jobID)
	if err != nil {
		return j.UNKNOWN, err
	}
	return *&job.Process.Status, nil
}

func (w *worker) StreamJobOutput(ctx context.Context, jobID string) (outchan chan string, err error) {
	w.m.RLock()
	job, err := w.getJob(jobID)
	w.m.RUnlock()
	if err != nil {
		return nil, err
	}

	return w.logger.Tailf(ctx, job.ID)
}

func (w *worker) getJob(jobID string) (job *j.Job, err error) {
	if job, exists := w.jobs[jobID]; exists {
		return job, nil
	}
	return nil, fmt.Errorf("job with ID %s not found", jobID)
}
