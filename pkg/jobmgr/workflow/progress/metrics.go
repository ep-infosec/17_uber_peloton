// Copyright (c) 2019 Uber Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package progress

import "github.com/uber-go/tally"

type Metrics struct {
	TotalWorkflow       tally.Gauge
	TotalActiveWorkflow tally.Gauge
	TotalStaleWorkflow  tally.Gauge
	ProcessDuration     tally.Timer

	GetJobRuntimeFailure tally.Gauge
}

func NewMetrics(scope tally.Scope) *Metrics {
	progressScope := scope.SubScope("workflow_progress")
	workflowScope := progressScope.SubScope("workflow")
	return &Metrics{
		TotalWorkflow: workflowScope.Tagged(map[string]string{
			"workflow_type": "all",
		}).Gauge("total_workflow"),
		TotalActiveWorkflow: workflowScope.Tagged(map[string]string{
			"workflow_type": "active",
		}).Gauge("total_workflow"),
		TotalStaleWorkflow: workflowScope.Tagged(map[string]string{
			"workflow_type": "stale",
		}).Gauge("total_workflow"),
		ProcessDuration: progressScope.Timer("duration"),

		GetJobRuntimeFailure: progressScope.Gauge("job_runtime_failure"),
	}
}
