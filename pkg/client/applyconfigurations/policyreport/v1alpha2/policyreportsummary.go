/*
Copyright The Kubernetes Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha2

// PolicyReportSummaryApplyConfiguration represents an declarative configuration of the PolicyReportSummary type for use
// with apply.
type PolicyReportSummaryApplyConfiguration struct {
	Pass  *int `json:"pass,omitempty"`
	Fail  *int `json:"fail,omitempty"`
	Warn  *int `json:"warn,omitempty"`
	Error *int `json:"error,omitempty"`
	Skip  *int `json:"skip,omitempty"`
}

// PolicyReportSummaryApplyConfiguration constructs an declarative configuration of the PolicyReportSummary type for use with
// apply.
func PolicyReportSummary() *PolicyReportSummaryApplyConfiguration {
	return &PolicyReportSummaryApplyConfiguration{}
}

// WithPass sets the Pass field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Pass field is set to the value of the last call.
func (b *PolicyReportSummaryApplyConfiguration) WithPass(value int) *PolicyReportSummaryApplyConfiguration {
	b.Pass = &value
	return b
}

// WithFail sets the Fail field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Fail field is set to the value of the last call.
func (b *PolicyReportSummaryApplyConfiguration) WithFail(value int) *PolicyReportSummaryApplyConfiguration {
	b.Fail = &value
	return b
}

// WithWarn sets the Warn field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Warn field is set to the value of the last call.
func (b *PolicyReportSummaryApplyConfiguration) WithWarn(value int) *PolicyReportSummaryApplyConfiguration {
	b.Warn = &value
	return b
}

// WithError sets the Error field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Error field is set to the value of the last call.
func (b *PolicyReportSummaryApplyConfiguration) WithError(value int) *PolicyReportSummaryApplyConfiguration {
	b.Error = &value
	return b
}

// WithSkip sets the Skip field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Skip field is set to the value of the last call.
func (b *PolicyReportSummaryApplyConfiguration) WithSkip(value int) *PolicyReportSummaryApplyConfiguration {
	b.Skip = &value
	return b
}
