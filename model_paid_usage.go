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
)

// PaidUsage struct for PaidUsage
type PaidUsage struct {
	PaidUsage *PaidUsageResponse `json:"paid_usage,omitempty"`
}

// NewPaidUsage instantiates a new PaidUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaidUsage() *PaidUsage {
	this := PaidUsage{}
	return &this
}

// NewPaidUsageWithDefaults instantiates a new PaidUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaidUsageWithDefaults() *PaidUsage {
	this := PaidUsage{}
	return &this
}

// GetPaidUsage returns the PaidUsage field value if set, zero value otherwise.
func (o *PaidUsage) GetPaidUsage() PaidUsageResponse {
	if o == nil || o.PaidUsage == nil {
		var ret PaidUsageResponse
		return ret
	}
	return *o.PaidUsage
}

// GetPaidUsageOk returns a tuple with the PaidUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaidUsage) GetPaidUsageOk() (*PaidUsageResponse, bool) {
	if o == nil || o.PaidUsage == nil {
		return nil, false
	}
	return o.PaidUsage, true
}

// HasPaidUsage returns a boolean if a field has been set.
func (o *PaidUsage) HasPaidUsage() bool {
	if o != nil && o.PaidUsage != nil {
		return true
	}

	return false
}

// SetPaidUsage gets a reference to the given PaidUsageResponse and assigns it to the PaidUsage field.
func (o *PaidUsage) SetPaidUsage(v PaidUsageResponse) {
	o.PaidUsage = &v
}

func (o PaidUsage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PaidUsage != nil {
		toSerialize["paid_usage"] = o.PaidUsage
	}
	return json.Marshal(toSerialize)
}

type NullablePaidUsage struct {
	value *PaidUsage
	isSet bool
}

func (v NullablePaidUsage) Get() *PaidUsage {
	return v.value
}

func (v *NullablePaidUsage) Set(val *PaidUsage) {
	v.value = val
	v.isSet = true
}

func (v NullablePaidUsage) IsSet() bool {
	return v.isSet
}

func (v *NullablePaidUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaidUsage(val *PaidUsage) *NullablePaidUsage {
	return &NullablePaidUsage{value: val, isSet: true}
}

func (v NullablePaidUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaidUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
