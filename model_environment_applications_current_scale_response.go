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

// EnvironmentApplicationsCurrentScaleResponse struct for EnvironmentApplicationsCurrentScaleResponse
type EnvironmentApplicationsCurrentScaleResponse struct {
	Application               *string    `json:"application,omitempty"`
	Min                       *int32     `json:"min,omitempty"`
	Max                       *int32     `json:"max,omitempty"`
	Running                   *int32     `json:"running,omitempty"`
	RunningInPercent          *float32   `json:"running_in_percent,omitempty"`
	WarningThresholdInPercent *float32   `json:"warning_threshold_in_percent,omitempty"`
	AlertThresholdInPercent   *float32   `json:"alert_threshold_in_percent,omitempty"`
	Status                    *string    `json:"status,omitempty"`
	UpdatedAt                 *time.Time `json:"updated_at,omitempty"`
}

// NewEnvironmentApplicationsCurrentScaleResponse instantiates a new EnvironmentApplicationsCurrentScaleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentApplicationsCurrentScaleResponse() *EnvironmentApplicationsCurrentScaleResponse {
	this := EnvironmentApplicationsCurrentScaleResponse{}
	return &this
}

// NewEnvironmentApplicationsCurrentScaleResponseWithDefaults instantiates a new EnvironmentApplicationsCurrentScaleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentApplicationsCurrentScaleResponseWithDefaults() *EnvironmentApplicationsCurrentScaleResponse {
	this := EnvironmentApplicationsCurrentScaleResponse{}
	return &this
}

// GetApplication returns the Application field value if set, zero value otherwise.
func (o *EnvironmentApplicationsCurrentScaleResponse) GetApplication() string {
	if o == nil || o.Application == nil {
		var ret string
		return ret
	}
	return *o.Application
}

// GetApplicationOk returns a tuple with the Application field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentApplicationsCurrentScaleResponse) GetApplicationOk() (*string, bool) {
	if o == nil || o.Application == nil {
		return nil, false
	}
	return o.Application, true
}

// HasApplication returns a boolean if a field has been set.
func (o *EnvironmentApplicationsCurrentScaleResponse) HasApplication() bool {
	if o != nil && o.Application != nil {
		return true
	}

	return false
}

// SetApplication gets a reference to the given string and assigns it to the Application field.
func (o *EnvironmentApplicationsCurrentScaleResponse) SetApplication(v string) {
	o.Application = &v
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *EnvironmentApplicationsCurrentScaleResponse) GetMin() int32 {
	if o == nil || o.Min == nil {
		var ret int32
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentApplicationsCurrentScaleResponse) GetMinOk() (*int32, bool) {
	if o == nil || o.Min == nil {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *EnvironmentApplicationsCurrentScaleResponse) HasMin() bool {
	if o != nil && o.Min != nil {
		return true
	}

	return false
}

// SetMin gets a reference to the given int32 and assigns it to the Min field.
func (o *EnvironmentApplicationsCurrentScaleResponse) SetMin(v int32) {
	o.Min = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *EnvironmentApplicationsCurrentScaleResponse) GetMax() int32 {
	if o == nil || o.Max == nil {
		var ret int32
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentApplicationsCurrentScaleResponse) GetMaxOk() (*int32, bool) {
	if o == nil || o.Max == nil {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *EnvironmentApplicationsCurrentScaleResponse) HasMax() bool {
	if o != nil && o.Max != nil {
		return true
	}

	return false
}

// SetMax gets a reference to the given int32 and assigns it to the Max field.
func (o *EnvironmentApplicationsCurrentScaleResponse) SetMax(v int32) {
	o.Max = &v
}

// GetRunning returns the Running field value if set, zero value otherwise.
func (o *EnvironmentApplicationsCurrentScaleResponse) GetRunning() int32 {
	if o == nil || o.Running == nil {
		var ret int32
		return ret
	}
	return *o.Running
}

// GetRunningOk returns a tuple with the Running field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentApplicationsCurrentScaleResponse) GetRunningOk() (*int32, bool) {
	if o == nil || o.Running == nil {
		return nil, false
	}
	return o.Running, true
}

// HasRunning returns a boolean if a field has been set.
func (o *EnvironmentApplicationsCurrentScaleResponse) HasRunning() bool {
	if o != nil && o.Running != nil {
		return true
	}

	return false
}

// SetRunning gets a reference to the given int32 and assigns it to the Running field.
func (o *EnvironmentApplicationsCurrentScaleResponse) SetRunning(v int32) {
	o.Running = &v
}

// GetRunningInPercent returns the RunningInPercent field value if set, zero value otherwise.
func (o *EnvironmentApplicationsCurrentScaleResponse) GetRunningInPercent() float32 {
	if o == nil || o.RunningInPercent == nil {
		var ret float32
		return ret
	}
	return *o.RunningInPercent
}

// GetRunningInPercentOk returns a tuple with the RunningInPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentApplicationsCurrentScaleResponse) GetRunningInPercentOk() (*float32, bool) {
	if o == nil || o.RunningInPercent == nil {
		return nil, false
	}
	return o.RunningInPercent, true
}

// HasRunningInPercent returns a boolean if a field has been set.
func (o *EnvironmentApplicationsCurrentScaleResponse) HasRunningInPercent() bool {
	if o != nil && o.RunningInPercent != nil {
		return true
	}

	return false
}

// SetRunningInPercent gets a reference to the given float32 and assigns it to the RunningInPercent field.
func (o *EnvironmentApplicationsCurrentScaleResponse) SetRunningInPercent(v float32) {
	o.RunningInPercent = &v
}

// GetWarningThresholdInPercent returns the WarningThresholdInPercent field value if set, zero value otherwise.
func (o *EnvironmentApplicationsCurrentScaleResponse) GetWarningThresholdInPercent() float32 {
	if o == nil || o.WarningThresholdInPercent == nil {
		var ret float32
		return ret
	}
	return *o.WarningThresholdInPercent
}

// GetWarningThresholdInPercentOk returns a tuple with the WarningThresholdInPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentApplicationsCurrentScaleResponse) GetWarningThresholdInPercentOk() (*float32, bool) {
	if o == nil || o.WarningThresholdInPercent == nil {
		return nil, false
	}
	return o.WarningThresholdInPercent, true
}

// HasWarningThresholdInPercent returns a boolean if a field has been set.
func (o *EnvironmentApplicationsCurrentScaleResponse) HasWarningThresholdInPercent() bool {
	if o != nil && o.WarningThresholdInPercent != nil {
		return true
	}

	return false
}

// SetWarningThresholdInPercent gets a reference to the given float32 and assigns it to the WarningThresholdInPercent field.
func (o *EnvironmentApplicationsCurrentScaleResponse) SetWarningThresholdInPercent(v float32) {
	o.WarningThresholdInPercent = &v
}

// GetAlertThresholdInPercent returns the AlertThresholdInPercent field value if set, zero value otherwise.
func (o *EnvironmentApplicationsCurrentScaleResponse) GetAlertThresholdInPercent() float32 {
	if o == nil || o.AlertThresholdInPercent == nil {
		var ret float32
		return ret
	}
	return *o.AlertThresholdInPercent
}

// GetAlertThresholdInPercentOk returns a tuple with the AlertThresholdInPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentApplicationsCurrentScaleResponse) GetAlertThresholdInPercentOk() (*float32, bool) {
	if o == nil || o.AlertThresholdInPercent == nil {
		return nil, false
	}
	return o.AlertThresholdInPercent, true
}

// HasAlertThresholdInPercent returns a boolean if a field has been set.
func (o *EnvironmentApplicationsCurrentScaleResponse) HasAlertThresholdInPercent() bool {
	if o != nil && o.AlertThresholdInPercent != nil {
		return true
	}

	return false
}

// SetAlertThresholdInPercent gets a reference to the given float32 and assigns it to the AlertThresholdInPercent field.
func (o *EnvironmentApplicationsCurrentScaleResponse) SetAlertThresholdInPercent(v float32) {
	o.AlertThresholdInPercent = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *EnvironmentApplicationsCurrentScaleResponse) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentApplicationsCurrentScaleResponse) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *EnvironmentApplicationsCurrentScaleResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *EnvironmentApplicationsCurrentScaleResponse) SetStatus(v string) {
	o.Status = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *EnvironmentApplicationsCurrentScaleResponse) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentApplicationsCurrentScaleResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *EnvironmentApplicationsCurrentScaleResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *EnvironmentApplicationsCurrentScaleResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o EnvironmentApplicationsCurrentScaleResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Application != nil {
		toSerialize["application"] = o.Application
	}
	if o.Min != nil {
		toSerialize["min"] = o.Min
	}
	if o.Max != nil {
		toSerialize["max"] = o.Max
	}
	if o.Running != nil {
		toSerialize["running"] = o.Running
	}
	if o.RunningInPercent != nil {
		toSerialize["running_in_percent"] = o.RunningInPercent
	}
	if o.WarningThresholdInPercent != nil {
		toSerialize["warning_threshold_in_percent"] = o.WarningThresholdInPercent
	}
	if o.AlertThresholdInPercent != nil {
		toSerialize["alert_threshold_in_percent"] = o.AlertThresholdInPercent
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironmentApplicationsCurrentScaleResponse struct {
	value *EnvironmentApplicationsCurrentScaleResponse
	isSet bool
}

func (v NullableEnvironmentApplicationsCurrentScaleResponse) Get() *EnvironmentApplicationsCurrentScaleResponse {
	return v.value
}

func (v *NullableEnvironmentApplicationsCurrentScaleResponse) Set(val *EnvironmentApplicationsCurrentScaleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentApplicationsCurrentScaleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentApplicationsCurrentScaleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentApplicationsCurrentScaleResponse(val *EnvironmentApplicationsCurrentScaleResponse) *NullableEnvironmentApplicationsCurrentScaleResponse {
	return &NullableEnvironmentApplicationsCurrentScaleResponse{value: val, isSet: true}
}

func (v NullableEnvironmentApplicationsCurrentScaleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentApplicationsCurrentScaleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
