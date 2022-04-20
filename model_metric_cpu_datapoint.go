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

// MetricCPUDatapoint struct for MetricCPUDatapoint
type MetricCPUDatapoint struct {
	CreatedAt         time.Time `json:"created_at"`
	RequestedInNumber *float32  `json:"requested_in_number,omitempty"`
	ConsumedInNumber  float32   `json:"consumed_in_number"`
	ConsumedInPercent float32   `json:"consumed_in_percent"`
}

// NewMetricCPUDatapoint instantiates a new MetricCPUDatapoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricCPUDatapoint(createdAt time.Time, consumedInNumber float32, consumedInPercent float32) *MetricCPUDatapoint {
	this := MetricCPUDatapoint{}
	this.CreatedAt = createdAt
	this.ConsumedInNumber = consumedInNumber
	this.ConsumedInPercent = consumedInPercent
	return &this
}

// NewMetricCPUDatapointWithDefaults instantiates a new MetricCPUDatapoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricCPUDatapointWithDefaults() *MetricCPUDatapoint {
	this := MetricCPUDatapoint{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *MetricCPUDatapoint) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *MetricCPUDatapoint) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *MetricCPUDatapoint) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetRequestedInNumber returns the RequestedInNumber field value if set, zero value otherwise.
func (o *MetricCPUDatapoint) GetRequestedInNumber() float32 {
	if o == nil || o.RequestedInNumber == nil {
		var ret float32
		return ret
	}
	return *o.RequestedInNumber
}

// GetRequestedInNumberOk returns a tuple with the RequestedInNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricCPUDatapoint) GetRequestedInNumberOk() (*float32, bool) {
	if o == nil || o.RequestedInNumber == nil {
		return nil, false
	}
	return o.RequestedInNumber, true
}

// HasRequestedInNumber returns a boolean if a field has been set.
func (o *MetricCPUDatapoint) HasRequestedInNumber() bool {
	if o != nil && o.RequestedInNumber != nil {
		return true
	}

	return false
}

// SetRequestedInNumber gets a reference to the given float32 and assigns it to the RequestedInNumber field.
func (o *MetricCPUDatapoint) SetRequestedInNumber(v float32) {
	o.RequestedInNumber = &v
}

// GetConsumedInNumber returns the ConsumedInNumber field value
func (o *MetricCPUDatapoint) GetConsumedInNumber() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ConsumedInNumber
}

// GetConsumedInNumberOk returns a tuple with the ConsumedInNumber field value
// and a boolean to check if the value has been set.
func (o *MetricCPUDatapoint) GetConsumedInNumberOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConsumedInNumber, true
}

// SetConsumedInNumber sets field value
func (o *MetricCPUDatapoint) SetConsumedInNumber(v float32) {
	o.ConsumedInNumber = v
}

// GetConsumedInPercent returns the ConsumedInPercent field value
func (o *MetricCPUDatapoint) GetConsumedInPercent() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ConsumedInPercent
}

// GetConsumedInPercentOk returns a tuple with the ConsumedInPercent field value
// and a boolean to check if the value has been set.
func (o *MetricCPUDatapoint) GetConsumedInPercentOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConsumedInPercent, true
}

// SetConsumedInPercent sets field value
func (o *MetricCPUDatapoint) SetConsumedInPercent(v float32) {
	o.ConsumedInPercent = v
}

func (o MetricCPUDatapoint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.RequestedInNumber != nil {
		toSerialize["requested_in_number"] = o.RequestedInNumber
	}
	if true {
		toSerialize["consumed_in_number"] = o.ConsumedInNumber
	}
	if true {
		toSerialize["consumed_in_percent"] = o.ConsumedInPercent
	}
	return json.Marshal(toSerialize)
}

type NullableMetricCPUDatapoint struct {
	value *MetricCPUDatapoint
	isSet bool
}

func (v NullableMetricCPUDatapoint) Get() *MetricCPUDatapoint {
	return v.value
}

func (v *NullableMetricCPUDatapoint) Set(val *MetricCPUDatapoint) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricCPUDatapoint) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricCPUDatapoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricCPUDatapoint(val *MetricCPUDatapoint) *NullableMetricCPUDatapoint {
	return &NullableMetricCPUDatapoint{value: val, isSet: true}
}

func (v NullableMetricCPUDatapoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricCPUDatapoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}