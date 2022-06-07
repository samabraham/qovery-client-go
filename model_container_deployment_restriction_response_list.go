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

// ContainerDeploymentRestrictionResponseList struct for ContainerDeploymentRestrictionResponseList
type ContainerDeploymentRestrictionResponseList struct {
	DeploymentRestrictions []ContainerDeploymentRestriction `json:"deployment_restrictions,omitempty"`
}

// NewContainerDeploymentRestrictionResponseList instantiates a new ContainerDeploymentRestrictionResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerDeploymentRestrictionResponseList() *ContainerDeploymentRestrictionResponseList {
	this := ContainerDeploymentRestrictionResponseList{}
	return &this
}

// NewContainerDeploymentRestrictionResponseListWithDefaults instantiates a new ContainerDeploymentRestrictionResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerDeploymentRestrictionResponseListWithDefaults() *ContainerDeploymentRestrictionResponseList {
	this := ContainerDeploymentRestrictionResponseList{}
	return &this
}

// GetDeploymentRestrictions returns the DeploymentRestrictions field value if set, zero value otherwise.
func (o *ContainerDeploymentRestrictionResponseList) GetDeploymentRestrictions() []ContainerDeploymentRestriction {
	if o == nil || o.DeploymentRestrictions == nil {
		var ret []ContainerDeploymentRestriction
		return ret
	}
	return o.DeploymentRestrictions
}

// GetDeploymentRestrictionsOk returns a tuple with the DeploymentRestrictions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerDeploymentRestrictionResponseList) GetDeploymentRestrictionsOk() ([]ContainerDeploymentRestriction, bool) {
	if o == nil || o.DeploymentRestrictions == nil {
		return nil, false
	}
	return o.DeploymentRestrictions, true
}

// HasDeploymentRestrictions returns a boolean if a field has been set.
func (o *ContainerDeploymentRestrictionResponseList) HasDeploymentRestrictions() bool {
	if o != nil && o.DeploymentRestrictions != nil {
		return true
	}

	return false
}

// SetDeploymentRestrictions gets a reference to the given []ContainerDeploymentRestriction and assigns it to the DeploymentRestrictions field.
func (o *ContainerDeploymentRestrictionResponseList) SetDeploymentRestrictions(v []ContainerDeploymentRestriction) {
	o.DeploymentRestrictions = v
}

func (o ContainerDeploymentRestrictionResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeploymentRestrictions != nil {
		toSerialize["deployment_restrictions"] = o.DeploymentRestrictions
	}
	return json.Marshal(toSerialize)
}

type NullableContainerDeploymentRestrictionResponseList struct {
	value *ContainerDeploymentRestrictionResponseList
	isSet bool
}

func (v NullableContainerDeploymentRestrictionResponseList) Get() *ContainerDeploymentRestrictionResponseList {
	return v.value
}

func (v *NullableContainerDeploymentRestrictionResponseList) Set(val *ContainerDeploymentRestrictionResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerDeploymentRestrictionResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerDeploymentRestrictionResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerDeploymentRestrictionResponseList(val *ContainerDeploymentRestrictionResponseList) *NullableContainerDeploymentRestrictionResponseList {
	return &NullableContainerDeploymentRestrictionResponseList{value: val, isSet: true}
}

func (v NullableContainerDeploymentRestrictionResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerDeploymentRestrictionResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}