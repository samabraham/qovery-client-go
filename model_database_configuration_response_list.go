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

// DatabaseConfigurationResponseList struct for DatabaseConfigurationResponseList
type DatabaseConfigurationResponseList struct {
	Results []DatabaseConfigurationResponse `json:"results,omitempty"`
}

// NewDatabaseConfigurationResponseList instantiates a new DatabaseConfigurationResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseConfigurationResponseList() *DatabaseConfigurationResponseList {
	this := DatabaseConfigurationResponseList{}
	return &this
}

// NewDatabaseConfigurationResponseListWithDefaults instantiates a new DatabaseConfigurationResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseConfigurationResponseListWithDefaults() *DatabaseConfigurationResponseList {
	this := DatabaseConfigurationResponseList{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *DatabaseConfigurationResponseList) GetResults() []DatabaseConfigurationResponse {
	if o == nil || o.Results == nil {
		var ret []DatabaseConfigurationResponse
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseConfigurationResponseList) GetResultsOk() ([]DatabaseConfigurationResponse, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *DatabaseConfigurationResponseList) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []DatabaseConfigurationResponse and assigns it to the Results field.
func (o *DatabaseConfigurationResponseList) SetResults(v []DatabaseConfigurationResponse) {
	o.Results = v
}

func (o DatabaseConfigurationResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableDatabaseConfigurationResponseList struct {
	value *DatabaseConfigurationResponseList
	isSet bool
}

func (v NullableDatabaseConfigurationResponseList) Get() *DatabaseConfigurationResponseList {
	return v.value
}

func (v *NullableDatabaseConfigurationResponseList) Set(val *DatabaseConfigurationResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseConfigurationResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseConfigurationResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseConfigurationResponseList(val *DatabaseConfigurationResponseList) *NullableDatabaseConfigurationResponseList {
	return &NullableDatabaseConfigurationResponseList{value: val, isSet: true}
}

func (v NullableDatabaseConfigurationResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseConfigurationResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
