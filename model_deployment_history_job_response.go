/*
Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is stable and still in development.

API version: 1.0.3
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"encoding/json"
	"time"
)

// DeploymentHistoryJobResponse struct for DeploymentHistoryJobResponse
type DeploymentHistoryJobResponse struct {
	Id        string     `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// name of the job
	Name       *string                                    `json:"name,omitempty"`
	Status     *DeploymentHistoryStatusEnum               `json:"status,omitempty"`
	ImageName  *string                                    `json:"image_name,omitempty"`
	Tag        *string                                    `json:"tag,omitempty"`
	Commit     *Commit                                    `json:"commit,omitempty"`
	Schedule   *DeploymentHistoryJobResponseAllOfSchedule `json:"schedule,omitempty"`
	Arguments  []string                                   `json:"arguments,omitempty"`
	Entrypoint *string                                    `json:"entrypoint,omitempty"`
}

// NewDeploymentHistoryJobResponse instantiates a new DeploymentHistoryJobResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentHistoryJobResponse(id string, createdAt time.Time) *DeploymentHistoryJobResponse {
	this := DeploymentHistoryJobResponse{}
	this.Id = id
	this.CreatedAt = createdAt
	return &this
}

// NewDeploymentHistoryJobResponseWithDefaults instantiates a new DeploymentHistoryJobResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentHistoryJobResponseWithDefaults() *DeploymentHistoryJobResponse {
	this := DeploymentHistoryJobResponse{}
	return &this
}

// GetId returns the Id field value
func (o *DeploymentHistoryJobResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryJobResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DeploymentHistoryJobResponse) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *DeploymentHistoryJobResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryJobResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *DeploymentHistoryJobResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *DeploymentHistoryJobResponse) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryJobResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *DeploymentHistoryJobResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *DeploymentHistoryJobResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DeploymentHistoryJobResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryJobResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DeploymentHistoryJobResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DeploymentHistoryJobResponse) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DeploymentHistoryJobResponse) GetStatus() DeploymentHistoryStatusEnum {
	if o == nil || o.Status == nil {
		var ret DeploymentHistoryStatusEnum
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryJobResponse) GetStatusOk() (*DeploymentHistoryStatusEnum, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DeploymentHistoryJobResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given DeploymentHistoryStatusEnum and assigns it to the Status field.
func (o *DeploymentHistoryJobResponse) SetStatus(v DeploymentHistoryStatusEnum) {
	o.Status = &v
}

// GetImageName returns the ImageName field value if set, zero value otherwise.
func (o *DeploymentHistoryJobResponse) GetImageName() string {
	if o == nil || o.ImageName == nil {
		var ret string
		return ret
	}
	return *o.ImageName
}

// GetImageNameOk returns a tuple with the ImageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryJobResponse) GetImageNameOk() (*string, bool) {
	if o == nil || o.ImageName == nil {
		return nil, false
	}
	return o.ImageName, true
}

// HasImageName returns a boolean if a field has been set.
func (o *DeploymentHistoryJobResponse) HasImageName() bool {
	if o != nil && o.ImageName != nil {
		return true
	}

	return false
}

// SetImageName gets a reference to the given string and assigns it to the ImageName field.
func (o *DeploymentHistoryJobResponse) SetImageName(v string) {
	o.ImageName = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *DeploymentHistoryJobResponse) GetTag() string {
	if o == nil || o.Tag == nil {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryJobResponse) GetTagOk() (*string, bool) {
	if o == nil || o.Tag == nil {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *DeploymentHistoryJobResponse) HasTag() bool {
	if o != nil && o.Tag != nil {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *DeploymentHistoryJobResponse) SetTag(v string) {
	o.Tag = &v
}

// GetCommit returns the Commit field value if set, zero value otherwise.
func (o *DeploymentHistoryJobResponse) GetCommit() Commit {
	if o == nil || o.Commit == nil {
		var ret Commit
		return ret
	}
	return *o.Commit
}

// GetCommitOk returns a tuple with the Commit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryJobResponse) GetCommitOk() (*Commit, bool) {
	if o == nil || o.Commit == nil {
		return nil, false
	}
	return o.Commit, true
}

// HasCommit returns a boolean if a field has been set.
func (o *DeploymentHistoryJobResponse) HasCommit() bool {
	if o != nil && o.Commit != nil {
		return true
	}

	return false
}

// SetCommit gets a reference to the given Commit and assigns it to the Commit field.
func (o *DeploymentHistoryJobResponse) SetCommit(v Commit) {
	o.Commit = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *DeploymentHistoryJobResponse) GetSchedule() DeploymentHistoryJobResponseAllOfSchedule {
	if o == nil || o.Schedule == nil {
		var ret DeploymentHistoryJobResponseAllOfSchedule
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryJobResponse) GetScheduleOk() (*DeploymentHistoryJobResponseAllOfSchedule, bool) {
	if o == nil || o.Schedule == nil {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *DeploymentHistoryJobResponse) HasSchedule() bool {
	if o != nil && o.Schedule != nil {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given DeploymentHistoryJobResponseAllOfSchedule and assigns it to the Schedule field.
func (o *DeploymentHistoryJobResponse) SetSchedule(v DeploymentHistoryJobResponseAllOfSchedule) {
	o.Schedule = &v
}

// GetArguments returns the Arguments field value if set, zero value otherwise.
func (o *DeploymentHistoryJobResponse) GetArguments() []string {
	if o == nil || o.Arguments == nil {
		var ret []string
		return ret
	}
	return o.Arguments
}

// GetArgumentsOk returns a tuple with the Arguments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryJobResponse) GetArgumentsOk() ([]string, bool) {
	if o == nil || o.Arguments == nil {
		return nil, false
	}
	return o.Arguments, true
}

// HasArguments returns a boolean if a field has been set.
func (o *DeploymentHistoryJobResponse) HasArguments() bool {
	if o != nil && o.Arguments != nil {
		return true
	}

	return false
}

// SetArguments gets a reference to the given []string and assigns it to the Arguments field.
func (o *DeploymentHistoryJobResponse) SetArguments(v []string) {
	o.Arguments = v
}

// GetEntrypoint returns the Entrypoint field value if set, zero value otherwise.
func (o *DeploymentHistoryJobResponse) GetEntrypoint() string {
	if o == nil || o.Entrypoint == nil {
		var ret string
		return ret
	}
	return *o.Entrypoint
}

// GetEntrypointOk returns a tuple with the Entrypoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryJobResponse) GetEntrypointOk() (*string, bool) {
	if o == nil || o.Entrypoint == nil {
		return nil, false
	}
	return o.Entrypoint, true
}

// HasEntrypoint returns a boolean if a field has been set.
func (o *DeploymentHistoryJobResponse) HasEntrypoint() bool {
	if o != nil && o.Entrypoint != nil {
		return true
	}

	return false
}

// SetEntrypoint gets a reference to the given string and assigns it to the Entrypoint field.
func (o *DeploymentHistoryJobResponse) SetEntrypoint(v string) {
	o.Entrypoint = &v
}

func (o DeploymentHistoryJobResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.ImageName != nil {
		toSerialize["image_name"] = o.ImageName
	}
	if o.Tag != nil {
		toSerialize["tag"] = o.Tag
	}
	if o.Commit != nil {
		toSerialize["commit"] = o.Commit
	}
	if o.Schedule != nil {
		toSerialize["schedule"] = o.Schedule
	}
	if o.Arguments != nil {
		toSerialize["arguments"] = o.Arguments
	}
	if o.Entrypoint != nil {
		toSerialize["entrypoint"] = o.Entrypoint
	}
	return json.Marshal(toSerialize)
}

type NullableDeploymentHistoryJobResponse struct {
	value *DeploymentHistoryJobResponse
	isSet bool
}

func (v NullableDeploymentHistoryJobResponse) Get() *DeploymentHistoryJobResponse {
	return v.value
}

func (v *NullableDeploymentHistoryJobResponse) Set(val *DeploymentHistoryJobResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentHistoryJobResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentHistoryJobResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentHistoryJobResponse(val *DeploymentHistoryJobResponse) *NullableDeploymentHistoryJobResponse {
	return &NullableDeploymentHistoryJobResponse{value: val, isSet: true}
}

func (v NullableDeploymentHistoryJobResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentHistoryJobResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
