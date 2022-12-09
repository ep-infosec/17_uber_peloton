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

package ptoa

import (
	"github.com/uber/peloton/.gen/peloton/api/v1alpha/pod"
	"github.com/uber/peloton/.gen/thrift/aurora/api"
	"go.uber.org/thriftrw/ptr"
)

// NewRange converts peloton's instance id list to aurora's range
func NewRange(instanceIDRange []*pod.InstanceIDRange) []*api.Range {
	ranges := make([]*api.Range, 0, len(instanceIDRange))
	for _, instanceRange := range instanceIDRange {
		ranges = append(ranges, &api.Range{
			First: ptr.Int32(int32(instanceRange.GetFrom())),
			Last:  ptr.Int32(int32(instanceRange.GetTo())),
		})
	}
	return ranges
}
