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

// ListContainerDeploymentHistory200ResponseAllOf struct for ListContainerDeploymentHistory200ResponseAllOf
type ListContainerDeploymentHistory200ResponseAllOf struct {
	Results []DeploymentHistoryContainer `json:"results,omitempty"`
}

// NewListContainerDeploymentHistory200ResponseAllOf instantiates a new ListContainerDeploymentHistory200ResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListContainerDeploymentHistory200ResponseAllOf() *ListContainerDeploymentHistory200ResponseAllOf {
	this := ListContainerDeploymentHistory200ResponseAllOf{}
	return &this
}

// NewListContainerDeploymentHistory200ResponseAllOfWithDefaults instantiates a new ListContainerDeploymentHistory200ResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListContainerDeploymentHistory200ResponseAllOfWithDefaults() *ListContainerDeploymentHistory200ResponseAllOf {
	this := ListContainerDeploymentHistory200ResponseAllOf{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ListContainerDeploymentHistory200ResponseAllOf) GetResults() []DeploymentHistoryContainer {
	if o == nil || o.Results == nil {
		var ret []DeploymentHistoryContainer
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListContainerDeploymentHistory200ResponseAllOf) GetResultsOk() ([]DeploymentHistoryContainer, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ListContainerDeploymentHistory200ResponseAllOf) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []DeploymentHistoryContainer and assigns it to the Results field.
func (o *ListContainerDeploymentHistory200ResponseAllOf) SetResults(v []DeploymentHistoryContainer) {
	o.Results = v
}

func (o ListContainerDeploymentHistory200ResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableListContainerDeploymentHistory200ResponseAllOf struct {
	value *ListContainerDeploymentHistory200ResponseAllOf
	isSet bool
}

func (v NullableListContainerDeploymentHistory200ResponseAllOf) Get() *ListContainerDeploymentHistory200ResponseAllOf {
	return v.value
}

func (v *NullableListContainerDeploymentHistory200ResponseAllOf) Set(val *ListContainerDeploymentHistory200ResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableListContainerDeploymentHistory200ResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableListContainerDeploymentHistory200ResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListContainerDeploymentHistory200ResponseAllOf(val *ListContainerDeploymentHistory200ResponseAllOf) *NullableListContainerDeploymentHistory200ResponseAllOf {
	return &NullableListContainerDeploymentHistory200ResponseAllOf{value: val, isSet: true}
}

func (v NullableListContainerDeploymentHistory200ResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListContainerDeploymentHistory200ResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
