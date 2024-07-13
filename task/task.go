package task

import (
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/google/uuid"
)

type State int

const(
	Pending State = iota
	Scheduled 
	Running
	Completed
	Failed
)

// Task struct represents a task that a user wants to run on the cluster
// that a user wants to run on our cluster.
type Task struct{
	ID				uuid.UUID
	Name			string
	state			State
	Image			string
	Memory			int
	Disk			int
	ExposedPorts	nat.PortSet
	RestartPolicy	string
	StartTime		time.Time
	FinishTime		time.Time
}

// how does a user tell the
// system to stop a task? For this purpose, let’s introduce the
// TaskEvent struct.
type TaskEvent struct{
	ID 				uuid.UUID
	State			State
	Timestamp		time.Time
	Task			Task
}