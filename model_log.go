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

// Log struct for Log
type Log struct {
	Id        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Message   string    `json:"message"`
	PodName   *string   `json:"pod_name,omitempty"`
	Version   *string   `json:"version,omitempty"`
}

// NewLog instantiates a new Log object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLog(id string, createdAt time.Time, message string) *Log {
	this := Log{}
	this.Id = id
	this.CreatedAt = createdAt
	this.Message = message
	return &this
}

// NewLogWithDefaults instantiates a new Log object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogWithDefaults() *Log {
	this := Log{}
	return &this
}

// GetId returns the Id field value
func (o *Log) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Log) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Log) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Log) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Log) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Log) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetMessage returns the Message field value
func (o *Log) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *Log) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *Log) SetMessage(v string) {
	o.Message = v
}

// GetPodName returns the PodName field value if set, zero value otherwise.
func (o *Log) GetPodName() string {
	if o == nil || o.PodName == nil {
		var ret string
		return ret
	}
	return *o.PodName
}

// GetPodNameOk returns a tuple with the PodName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Log) GetPodNameOk() (*string, bool) {
	if o == nil || o.PodName == nil {
		return nil, false
	}
	return o.PodName, true
}

// HasPodName returns a boolean if a field has been set.
func (o *Log) HasPodName() bool {
	if o != nil && o.PodName != nil {
		return true
	}

	return false
}

// SetPodName gets a reference to the given string and assigns it to the PodName field.
func (o *Log) SetPodName(v string) {
	o.PodName = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Log) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Log) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Log) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *Log) SetVersion(v string) {
	o.Version = &v
}

func (o Log) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["message"] = o.Message
	}
	if o.PodName != nil {
		toSerialize["pod_name"] = o.PodName
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableLog struct {
	value *Log
	isSet bool
}

func (v NullableLog) Get() *Log {
	return v.value
}

func (v *NullableLog) Set(val *Log) {
	v.value = val
	v.isSet = true
}

func (v NullableLog) IsSet() bool {
	return v.isSet
}

func (v *NullableLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLog(val *Log) *NullableLog {
	return &NullableLog{value: val, isSet: true}
}

func (v NullableLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
