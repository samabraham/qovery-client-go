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

// ReferenceObjectStatus struct for ReferenceObjectStatus
type ReferenceObjectStatus struct {
	Id    string    `json:"id"`
	State StateEnum `json:"state"`
	// message related to the state
	Message                 *string                     `json:"message,omitempty"`
	ServiceDeploymentStatus ServiceDeploymentStatusEnum `json:"service_deployment_status"`
	LastDeploymentDate      *time.Time                  `json:"last_deployment_date,omitempty"`
}

// NewReferenceObjectStatus instantiates a new ReferenceObjectStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReferenceObjectStatus(id string, state StateEnum, serviceDeploymentStatus ServiceDeploymentStatusEnum) *ReferenceObjectStatus {
	this := ReferenceObjectStatus{}
	this.Id = id
	this.State = state
	this.ServiceDeploymentStatus = serviceDeploymentStatus
	return &this
}

// NewReferenceObjectStatusWithDefaults instantiates a new ReferenceObjectStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReferenceObjectStatusWithDefaults() *ReferenceObjectStatus {
	this := ReferenceObjectStatus{}
	return &this
}

// GetId returns the Id field value
func (o *ReferenceObjectStatus) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ReferenceObjectStatus) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ReferenceObjectStatus) SetId(v string) {
	o.Id = v
}

// GetState returns the State field value
func (o *ReferenceObjectStatus) GetState() StateEnum {
	if o == nil {
		var ret StateEnum
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *ReferenceObjectStatus) GetStateOk() (*StateEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *ReferenceObjectStatus) SetState(v StateEnum) {
	o.State = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ReferenceObjectStatus) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferenceObjectStatus) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ReferenceObjectStatus) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ReferenceObjectStatus) SetMessage(v string) {
	o.Message = &v
}

// GetServiceDeploymentStatus returns the ServiceDeploymentStatus field value
func (o *ReferenceObjectStatus) GetServiceDeploymentStatus() ServiceDeploymentStatusEnum {
	if o == nil {
		var ret ServiceDeploymentStatusEnum
		return ret
	}

	return o.ServiceDeploymentStatus
}

// GetServiceDeploymentStatusOk returns a tuple with the ServiceDeploymentStatus field value
// and a boolean to check if the value has been set.
func (o *ReferenceObjectStatus) GetServiceDeploymentStatusOk() (*ServiceDeploymentStatusEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceDeploymentStatus, true
}

// SetServiceDeploymentStatus sets field value
func (o *ReferenceObjectStatus) SetServiceDeploymentStatus(v ServiceDeploymentStatusEnum) {
	o.ServiceDeploymentStatus = v
}

// GetLastDeploymentDate returns the LastDeploymentDate field value if set, zero value otherwise.
func (o *ReferenceObjectStatus) GetLastDeploymentDate() time.Time {
	if o == nil || o.LastDeploymentDate == nil {
		var ret time.Time
		return ret
	}
	return *o.LastDeploymentDate
}

// GetLastDeploymentDateOk returns a tuple with the LastDeploymentDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferenceObjectStatus) GetLastDeploymentDateOk() (*time.Time, bool) {
	if o == nil || o.LastDeploymentDate == nil {
		return nil, false
	}
	return o.LastDeploymentDate, true
}

// HasLastDeploymentDate returns a boolean if a field has been set.
func (o *ReferenceObjectStatus) HasLastDeploymentDate() bool {
	if o != nil && o.LastDeploymentDate != nil {
		return true
	}

	return false
}

// SetLastDeploymentDate gets a reference to the given time.Time and assigns it to the LastDeploymentDate field.
func (o *ReferenceObjectStatus) SetLastDeploymentDate(v time.Time) {
	o.LastDeploymentDate = &v
}

func (o ReferenceObjectStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["state"] = o.State
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if true {
		toSerialize["service_deployment_status"] = o.ServiceDeploymentStatus
	}
	if o.LastDeploymentDate != nil {
		toSerialize["last_deployment_date"] = o.LastDeploymentDate
	}
	return json.Marshal(toSerialize)
}

type NullableReferenceObjectStatus struct {
	value *ReferenceObjectStatus
	isSet bool
}

func (v NullableReferenceObjectStatus) Get() *ReferenceObjectStatus {
	return v.value
}

func (v *NullableReferenceObjectStatus) Set(val *ReferenceObjectStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableReferenceObjectStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableReferenceObjectStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReferenceObjectStatus(val *ReferenceObjectStatus) *NullableReferenceObjectStatus {
	return &NullableReferenceObjectStatus{value: val, isSet: true}
}

func (v NullableReferenceObjectStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReferenceObjectStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
