/*
Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is stable and still in development.

API version: 1.0.1
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"encoding/json"
	"time"
)

// EnvironmentDeploymentRuleEditRequest struct for EnvironmentDeploymentRuleEditRequest
type EnvironmentDeploymentRuleEditRequest struct {
	AutoDeploy *bool     `json:"auto_deploy,omitempty"`
	AutoDelete *bool     `json:"auto_delete,omitempty"`
	AutoStop   *bool     `json:"auto_stop,omitempty"`
	Timezone   string    `json:"timezone"`
	StartTime  time.Time `json:"start_time"`
	StopTime   time.Time `json:"stop_time"`
	Weekdays   []string  `json:"weekdays"`
}

// NewEnvironmentDeploymentRuleEditRequest instantiates a new EnvironmentDeploymentRuleEditRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentDeploymentRuleEditRequest(timezone string, startTime time.Time, stopTime time.Time, weekdays []string) *EnvironmentDeploymentRuleEditRequest {
	this := EnvironmentDeploymentRuleEditRequest{}
	var autoDeploy bool = true
	this.AutoDeploy = &autoDeploy
	var autoDelete bool = false
	this.AutoDelete = &autoDelete
	var autoStop bool = false
	this.AutoStop = &autoStop
	this.Timezone = timezone
	this.StartTime = startTime
	this.StopTime = stopTime
	this.Weekdays = weekdays
	return &this
}

// NewEnvironmentDeploymentRuleEditRequestWithDefaults instantiates a new EnvironmentDeploymentRuleEditRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentDeploymentRuleEditRequestWithDefaults() *EnvironmentDeploymentRuleEditRequest {
	this := EnvironmentDeploymentRuleEditRequest{}
	var autoDeploy bool = true
	this.AutoDeploy = &autoDeploy
	var autoDelete bool = false
	this.AutoDelete = &autoDelete
	var autoStop bool = false
	this.AutoStop = &autoStop
	return &this
}

// GetAutoDeploy returns the AutoDeploy field value if set, zero value otherwise.
func (o *EnvironmentDeploymentRuleEditRequest) GetAutoDeploy() bool {
	if o == nil || o.AutoDeploy == nil {
		var ret bool
		return ret
	}
	return *o.AutoDeploy
}

// GetAutoDeployOk returns a tuple with the AutoDeploy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentDeploymentRuleEditRequest) GetAutoDeployOk() (*bool, bool) {
	if o == nil || o.AutoDeploy == nil {
		return nil, false
	}
	return o.AutoDeploy, true
}

// HasAutoDeploy returns a boolean if a field has been set.
func (o *EnvironmentDeploymentRuleEditRequest) HasAutoDeploy() bool {
	if o != nil && o.AutoDeploy != nil {
		return true
	}

	return false
}

// SetAutoDeploy gets a reference to the given bool and assigns it to the AutoDeploy field.
func (o *EnvironmentDeploymentRuleEditRequest) SetAutoDeploy(v bool) {
	o.AutoDeploy = &v
}

// GetAutoDelete returns the AutoDelete field value if set, zero value otherwise.
func (o *EnvironmentDeploymentRuleEditRequest) GetAutoDelete() bool {
	if o == nil || o.AutoDelete == nil {
		var ret bool
		return ret
	}
	return *o.AutoDelete
}

// GetAutoDeleteOk returns a tuple with the AutoDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentDeploymentRuleEditRequest) GetAutoDeleteOk() (*bool, bool) {
	if o == nil || o.AutoDelete == nil {
		return nil, false
	}
	return o.AutoDelete, true
}

// HasAutoDelete returns a boolean if a field has been set.
func (o *EnvironmentDeploymentRuleEditRequest) HasAutoDelete() bool {
	if o != nil && o.AutoDelete != nil {
		return true
	}

	return false
}

// SetAutoDelete gets a reference to the given bool and assigns it to the AutoDelete field.
func (o *EnvironmentDeploymentRuleEditRequest) SetAutoDelete(v bool) {
	o.AutoDelete = &v
}

// GetAutoStop returns the AutoStop field value if set, zero value otherwise.
func (o *EnvironmentDeploymentRuleEditRequest) GetAutoStop() bool {
	if o == nil || o.AutoStop == nil {
		var ret bool
		return ret
	}
	return *o.AutoStop
}

// GetAutoStopOk returns a tuple with the AutoStop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentDeploymentRuleEditRequest) GetAutoStopOk() (*bool, bool) {
	if o == nil || o.AutoStop == nil {
		return nil, false
	}
	return o.AutoStop, true
}

// HasAutoStop returns a boolean if a field has been set.
func (o *EnvironmentDeploymentRuleEditRequest) HasAutoStop() bool {
	if o != nil && o.AutoStop != nil {
		return true
	}

	return false
}

// SetAutoStop gets a reference to the given bool and assigns it to the AutoStop field.
func (o *EnvironmentDeploymentRuleEditRequest) SetAutoStop(v bool) {
	o.AutoStop = &v
}

// GetTimezone returns the Timezone field value
func (o *EnvironmentDeploymentRuleEditRequest) GetTimezone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value
// and a boolean to check if the value has been set.
func (o *EnvironmentDeploymentRuleEditRequest) GetTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timezone, true
}

// SetTimezone sets field value
func (o *EnvironmentDeploymentRuleEditRequest) SetTimezone(v string) {
	o.Timezone = v
}

// GetStartTime returns the StartTime field value
func (o *EnvironmentDeploymentRuleEditRequest) GetStartTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *EnvironmentDeploymentRuleEditRequest) GetStartTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *EnvironmentDeploymentRuleEditRequest) SetStartTime(v time.Time) {
	o.StartTime = v
}

// GetStopTime returns the StopTime field value
func (o *EnvironmentDeploymentRuleEditRequest) GetStopTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StopTime
}

// GetStopTimeOk returns a tuple with the StopTime field value
// and a boolean to check if the value has been set.
func (o *EnvironmentDeploymentRuleEditRequest) GetStopTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StopTime, true
}

// SetStopTime sets field value
func (o *EnvironmentDeploymentRuleEditRequest) SetStopTime(v time.Time) {
	o.StopTime = v
}

// GetWeekdays returns the Weekdays field value
func (o *EnvironmentDeploymentRuleEditRequest) GetWeekdays() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Weekdays
}

// GetWeekdaysOk returns a tuple with the Weekdays field value
// and a boolean to check if the value has been set.
func (o *EnvironmentDeploymentRuleEditRequest) GetWeekdaysOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Weekdays, true
}

// SetWeekdays sets field value
func (o *EnvironmentDeploymentRuleEditRequest) SetWeekdays(v []string) {
	o.Weekdays = v
}

func (o EnvironmentDeploymentRuleEditRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AutoDeploy != nil {
		toSerialize["auto_deploy"] = o.AutoDeploy
	}
	if o.AutoDelete != nil {
		toSerialize["auto_delete"] = o.AutoDelete
	}
	if o.AutoStop != nil {
		toSerialize["auto_stop"] = o.AutoStop
	}
	if true {
		toSerialize["timezone"] = o.Timezone
	}
	if true {
		toSerialize["start_time"] = o.StartTime
	}
	if true {
		toSerialize["stop_time"] = o.StopTime
	}
	if true {
		toSerialize["weekdays"] = o.Weekdays
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironmentDeploymentRuleEditRequest struct {
	value *EnvironmentDeploymentRuleEditRequest
	isSet bool
}

func (v NullableEnvironmentDeploymentRuleEditRequest) Get() *EnvironmentDeploymentRuleEditRequest {
	return v.value
}

func (v *NullableEnvironmentDeploymentRuleEditRequest) Set(val *EnvironmentDeploymentRuleEditRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentDeploymentRuleEditRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentDeploymentRuleEditRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentDeploymentRuleEditRequest(val *EnvironmentDeploymentRuleEditRequest) *NullableEnvironmentDeploymentRuleEditRequest {
	return &NullableEnvironmentDeploymentRuleEditRequest{value: val, isSet: true}
}

func (v NullableEnvironmentDeploymentRuleEditRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentDeploymentRuleEditRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
