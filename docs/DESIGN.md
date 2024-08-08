---
authors: Jason Lane (jason.j.lane@gmail.com)
state: draft
---

# What

A prototype job worker service that provides an API to run arbitrary Linux processes.

# Why

To level set my engineering skill, and spend 2 weeks working on a project to get a feel for the day to day of an engineer at Teleport.

# Details

This project will consist of 3 main parts per the level 4 requirements spec laid out in the [challenge document](https://github.com/gravitational/careers/blob/main/challenges/systems/challenge-1.md).

## Library

The library will consist of 2 Go packages that will be used by the API and the CLI to run the processes.

### Worker package

The worker will be the top level struct in the library. Workers will be able to injest a configuration object which will define:
* Log output directory
* Log file name
* Log chunk size
* Logger

Workers will have 2 properties

a mutex

and a map of jobs (a map of jobID (uuid) to a pointer to the specific job.)

The worker struct will have 4 methods:.
    * StartJob - (command string, args []string) (jobID string) - Starts a job 
      * command is a program executable available via PATH or an absolute path to the program 
      * args are arguments to supply to the command
      * jobID is a string version of a uuidv4 of the created job
    * StopJob (jobID string) (status job.Status)
      * jobID is the ID of the job to stop
      * status is the status of the job after stopping expect(Terminated/Stopped)
    * GetJobStatus (jobID string) (status job.Status)
      * jobID is the ID of the job to get status on
      * status is the status of the job 
    * StreamJobOutput(jobID string) (output string)
      * jobID is the ID of the job to stream the output of
      * output is 

### Job package

The job package will define the structs containing the information about the specific processes spawned by the worker.

Example of the structs that will make up the Job package

```golang
type Process struct {
	PID int `json:"pid"`
	ExitCode int `json:"exit_code"`
	Status Status `json:"status"`
}

type Status int

const (
	UNKNOWN Status = iota
	RUNNING
	STOPPED
)

func (s Status) String() string {
	switch s {
	case RUNNING:
		return "running"
	case STOPPED:
		return "stopped"
	default:
		return "unknown"
	}
}

func (s Status) EnumIndex() int {
	return int(s)
}

type Job struct {
    ID string 
    Cmd *exec.Cmd 
    Output []byte
    Process *Process
    StartTime time.Time
    EndTime time.Time
    mu sync.Mutex
}
```

## API
The API will be a gRPC server that will implement the Worker service.

```protobuf
syntax = "proto3";

option go_package = "github.com/jaylane/job-scheduler/pkg/worker";

service Worker {
    // StartJob starts a job.
    rpc StartJob(StartJobRequest) returns (StartJobReponse);
    // StopJob stops a job.
    rpc StopJob(StopJobRequest) returns (StopJobResponse);
    // GetJobStatus gets the status of a job.
    rpc GetJobStatus(GetJobStatusRequest) returns (GetJobStatusResponse);
    /* StreamJobOutput streams the output of a job.
       Note: If the job is still active when StreamOutputRequest 
       is sent the cli will tail the output. */ 
    rpc StreamJobOutput(StreamJobOutputRequest) returns (StreamJobOutputResponse);
}

message StartJobRequest {
    // The name or path of the command to run.
    string command = 1;
    // The arguments to pass to the command.
    repeated string args = 2;
}

message StartJobReponse {
    // The id of the job that was started.
    string id = 1;
    // The status of the job that was started.
    string status = 2;
}

message StopJobRequest {
    // The id of the job to stop.
    string id = 1;
}

message StopJobResponse {
    // status is a readable string representation of the process' exit_code
    string status = 1;
}

message GetJobStatusRequest {
    // The id of the job to get the status of.
    string id = 1;
}

message GetJobStatusResponse {  
    // The process id of the job
    int32 pid = 1;
    // status is a readable string representation of the process' exit_code
    string status = 2;
}

message StreamJobOutputRequest {
    // The id of the job to stream the output of.
    string id = 1;
}

message StreamJobOutputResponse {
    // The output of the job.
    string output = 1;
}


```

### CLI
The CLI will use the API package







