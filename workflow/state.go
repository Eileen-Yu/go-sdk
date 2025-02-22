/*
Copyright 2024 The Dapr Authors
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package workflow

import (
	"github.com/dapr/durabletask-go/api"
	"github.com/dapr/durabletask-go/api/protos"
)

type Status int

const (
	StatusRunning Status = iota
	StatusCompleted
	StatusContinuedAsNew
	StatusFailed
	StatusCanceled
	StatusTerminated
	StatusPending
	StatusSuspended
	StatusUnknown
)

// String returns the runtime status as a string.
func (s Status) String() string {
	status := [...]string{
		"RUNNING",
		"COMPLETED",
		"CONTINUED_AS_NEW",
		"FAILED",
		"CANCELED",
		"TERMINATED",
		"PENDING",
		"SUSPENDED",
	}
	if s > StatusSuspended || s < StatusRunning {
		return "UNKNOWN"
	}
	return status[s]
}

func (s Status) RuntimeStatus() api.OrchestrationStatus {
	switch s {
	case StatusRunning:
		return api.RUNTIME_STATUS_RUNNING
	case StatusCompleted:
		return api.RUNTIME_STATUS_COMPLETED
	case StatusContinuedAsNew:
		return api.RUNTIME_STATUS_CONTINUED_AS_NEW
	case StatusFailed:
		return api.RUNTIME_STATUS_FAILED
	case StatusCanceled:
		return api.RUNTIME_STATUS_CANCELED
	case StatusTerminated:
		return api.RUNTIME_STATUS_TERMINATED
	case StatusPending:
		return api.RUNTIME_STATUS_PENDING
	case StatusSuspended:
		return api.RUNTIME_STATUS_SUSPENDED
	}
	return -1
}

type WorkflowState struct {
	Metadata protos.OrchestrationMetadata
}

// RuntimeStatus returns the status from a workflow state.
func (wfs *WorkflowState) RuntimeStatus() Status {
	s := Status(wfs.Metadata.GetRuntimeStatus().Number())
	return s
}

func convertStatusSlice(ss []Status) []api.OrchestrationStatus {
	out := []api.OrchestrationStatus{}
	for _, s := range ss {
		out = append(out, s.RuntimeStatus())
	}
	return out
}
