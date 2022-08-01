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

// ServicePortRequest struct for ServicePortRequest
type ServicePortRequest struct {
	Ports []ServicePortRequestPortsInner `json:"ports,omitempty"`
}

// NewServicePortRequest instantiates a new ServicePortRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServicePortRequest() *ServicePortRequest {
	this := ServicePortRequest{}
	return &this
}

// NewServicePortRequestWithDefaults instantiates a new ServicePortRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServicePortRequestWithDefaults() *ServicePortRequest {
	this := ServicePortRequest{}
	return &this
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *ServicePortRequest) GetPorts() []ServicePortRequestPortsInner {
	if o == nil || o.Ports == nil {
		var ret []ServicePortRequestPortsInner
		return ret
	}
	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePortRequest) GetPortsOk() ([]ServicePortRequestPortsInner, bool) {
	if o == nil || o.Ports == nil {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *ServicePortRequest) HasPorts() bool {
	if o != nil && o.Ports != nil {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []ServicePortRequestPortsInner and assigns it to the Ports field.
func (o *ServicePortRequest) SetPorts(v []ServicePortRequestPortsInner) {
	o.Ports = v
}

func (o ServicePortRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Ports != nil {
		toSerialize["ports"] = o.Ports
	}
	return json.Marshal(toSerialize)
}

type NullableServicePortRequest struct {
	value *ServicePortRequest
	isSet bool
}

func (v NullableServicePortRequest) Get() *ServicePortRequest {
	return v.value
}

func (v *NullableServicePortRequest) Set(val *ServicePortRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableServicePortRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableServicePortRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServicePortRequest(val *ServicePortRequest) *NullableServicePortRequest {
	return &NullableServicePortRequest{value: val, isSet: true}
}

func (v NullableServicePortRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServicePortRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}