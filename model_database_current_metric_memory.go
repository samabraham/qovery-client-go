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
)

// DatabaseCurrentMetricMemory struct for DatabaseCurrentMetricMemory
type DatabaseCurrentMetricMemory struct {
	RequestedInMb             *int32                     `json:"requested_in_mb,omitempty"`
	ConsumedInMb              *int32                     `json:"consumed_in_mb,omitempty"`
	ConsumedInPercent         *float32                   `json:"consumed_in_percent,omitempty"`
	WarningThresholdInPercent *float32                   `json:"warning_threshold_in_percent,omitempty"`
	AlertThresholdInPercent   *float32                   `json:"alert_threshold_in_percent,omitempty"`
	Status                    *ThresholdMetricStatusEnum `json:"status,omitempty"`
}

// NewDatabaseCurrentMetricMemory instantiates a new DatabaseCurrentMetricMemory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseCurrentMetricMemory() *DatabaseCurrentMetricMemory {
	this := DatabaseCurrentMetricMemory{}
	return &this
}

// NewDatabaseCurrentMetricMemoryWithDefaults instantiates a new DatabaseCurrentMetricMemory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseCurrentMetricMemoryWithDefaults() *DatabaseCurrentMetricMemory {
	this := DatabaseCurrentMetricMemory{}
	return &this
}

// GetRequestedInMb returns the RequestedInMb field value if set, zero value otherwise.
func (o *DatabaseCurrentMetricMemory) GetRequestedInMb() int32 {
	if o == nil || o.RequestedInMb == nil {
		var ret int32
		return ret
	}
	return *o.RequestedInMb
}

// GetRequestedInMbOk returns a tuple with the RequestedInMb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseCurrentMetricMemory) GetRequestedInMbOk() (*int32, bool) {
	if o == nil || o.RequestedInMb == nil {
		return nil, false
	}
	return o.RequestedInMb, true
}

// HasRequestedInMb returns a boolean if a field has been set.
func (o *DatabaseCurrentMetricMemory) HasRequestedInMb() bool {
	if o != nil && o.RequestedInMb != nil {
		return true
	}

	return false
}

// SetRequestedInMb gets a reference to the given int32 and assigns it to the RequestedInMb field.
func (o *DatabaseCurrentMetricMemory) SetRequestedInMb(v int32) {
	o.RequestedInMb = &v
}

// GetConsumedInMb returns the ConsumedInMb field value if set, zero value otherwise.
func (o *DatabaseCurrentMetricMemory) GetConsumedInMb() int32 {
	if o == nil || o.ConsumedInMb == nil {
		var ret int32
		return ret
	}
	return *o.ConsumedInMb
}

// GetConsumedInMbOk returns a tuple with the ConsumedInMb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseCurrentMetricMemory) GetConsumedInMbOk() (*int32, bool) {
	if o == nil || o.ConsumedInMb == nil {
		return nil, false
	}
	return o.ConsumedInMb, true
}

// HasConsumedInMb returns a boolean if a field has been set.
func (o *DatabaseCurrentMetricMemory) HasConsumedInMb() bool {
	if o != nil && o.ConsumedInMb != nil {
		return true
	}

	return false
}

// SetConsumedInMb gets a reference to the given int32 and assigns it to the ConsumedInMb field.
func (o *DatabaseCurrentMetricMemory) SetConsumedInMb(v int32) {
	o.ConsumedInMb = &v
}

// GetConsumedInPercent returns the ConsumedInPercent field value if set, zero value otherwise.
func (o *DatabaseCurrentMetricMemory) GetConsumedInPercent() float32 {
	if o == nil || o.ConsumedInPercent == nil {
		var ret float32
		return ret
	}
	return *o.ConsumedInPercent
}

// GetConsumedInPercentOk returns a tuple with the ConsumedInPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseCurrentMetricMemory) GetConsumedInPercentOk() (*float32, bool) {
	if o == nil || o.ConsumedInPercent == nil {
		return nil, false
	}
	return o.ConsumedInPercent, true
}

// HasConsumedInPercent returns a boolean if a field has been set.
func (o *DatabaseCurrentMetricMemory) HasConsumedInPercent() bool {
	if o != nil && o.ConsumedInPercent != nil {
		return true
	}

	return false
}

// SetConsumedInPercent gets a reference to the given float32 and assigns it to the ConsumedInPercent field.
func (o *DatabaseCurrentMetricMemory) SetConsumedInPercent(v float32) {
	o.ConsumedInPercent = &v
}

// GetWarningThresholdInPercent returns the WarningThresholdInPercent field value if set, zero value otherwise.
func (o *DatabaseCurrentMetricMemory) GetWarningThresholdInPercent() float32 {
	if o == nil || o.WarningThresholdInPercent == nil {
		var ret float32
		return ret
	}
	return *o.WarningThresholdInPercent
}

// GetWarningThresholdInPercentOk returns a tuple with the WarningThresholdInPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseCurrentMetricMemory) GetWarningThresholdInPercentOk() (*float32, bool) {
	if o == nil || o.WarningThresholdInPercent == nil {
		return nil, false
	}
	return o.WarningThresholdInPercent, true
}

// HasWarningThresholdInPercent returns a boolean if a field has been set.
func (o *DatabaseCurrentMetricMemory) HasWarningThresholdInPercent() bool {
	if o != nil && o.WarningThresholdInPercent != nil {
		return true
	}

	return false
}

// SetWarningThresholdInPercent gets a reference to the given float32 and assigns it to the WarningThresholdInPercent field.
func (o *DatabaseCurrentMetricMemory) SetWarningThresholdInPercent(v float32) {
	o.WarningThresholdInPercent = &v
}

// GetAlertThresholdInPercent returns the AlertThresholdInPercent field value if set, zero value otherwise.
func (o *DatabaseCurrentMetricMemory) GetAlertThresholdInPercent() float32 {
	if o == nil || o.AlertThresholdInPercent == nil {
		var ret float32
		return ret
	}
	return *o.AlertThresholdInPercent
}

// GetAlertThresholdInPercentOk returns a tuple with the AlertThresholdInPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseCurrentMetricMemory) GetAlertThresholdInPercentOk() (*float32, bool) {
	if o == nil || o.AlertThresholdInPercent == nil {
		return nil, false
	}
	return o.AlertThresholdInPercent, true
}

// HasAlertThresholdInPercent returns a boolean if a field has been set.
func (o *DatabaseCurrentMetricMemory) HasAlertThresholdInPercent() bool {
	if o != nil && o.AlertThresholdInPercent != nil {
		return true
	}

	return false
}

// SetAlertThresholdInPercent gets a reference to the given float32 and assigns it to the AlertThresholdInPercent field.
func (o *DatabaseCurrentMetricMemory) SetAlertThresholdInPercent(v float32) {
	o.AlertThresholdInPercent = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DatabaseCurrentMetricMemory) GetStatus() ThresholdMetricStatusEnum {
	if o == nil || o.Status == nil {
		var ret ThresholdMetricStatusEnum
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseCurrentMetricMemory) GetStatusOk() (*ThresholdMetricStatusEnum, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DatabaseCurrentMetricMemory) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ThresholdMetricStatusEnum and assigns it to the Status field.
func (o *DatabaseCurrentMetricMemory) SetStatus(v ThresholdMetricStatusEnum) {
	o.Status = &v
}

func (o DatabaseCurrentMetricMemory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RequestedInMb != nil {
		toSerialize["requested_in_mb"] = o.RequestedInMb
	}
	if o.ConsumedInMb != nil {
		toSerialize["consumed_in_mb"] = o.ConsumedInMb
	}
	if o.ConsumedInPercent != nil {
		toSerialize["consumed_in_percent"] = o.ConsumedInPercent
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
	return json.Marshal(toSerialize)
}

type NullableDatabaseCurrentMetricMemory struct {
	value *DatabaseCurrentMetricMemory
	isSet bool
}

func (v NullableDatabaseCurrentMetricMemory) Get() *DatabaseCurrentMetricMemory {
	return v.value
}

func (v *NullableDatabaseCurrentMetricMemory) Set(val *DatabaseCurrentMetricMemory) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseCurrentMetricMemory) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseCurrentMetricMemory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseCurrentMetricMemory(val *DatabaseCurrentMetricMemory) *NullableDatabaseCurrentMetricMemory {
	return &NullableDatabaseCurrentMetricMemory{value: val, isSet: true}
}

func (v NullableDatabaseCurrentMetricMemory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseCurrentMetricMemory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}