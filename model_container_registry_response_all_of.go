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

// ContainerRegistryResponseAllOf struct for ContainerRegistryResponseAllOf
type ContainerRegistryResponseAllOf struct {
	Name        *string                    `json:"name,omitempty"`
	Kind        *ContainerRegistryKindEnum `json:"kind,omitempty"`
	Description *string                    `json:"description,omitempty"`
	// URL of the container registry
	Url     *string      `json:"url,omitempty"`
	Cluster NullableBase `json:"cluster,omitempty"`
}

// NewContainerRegistryResponseAllOf instantiates a new ContainerRegistryResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerRegistryResponseAllOf() *ContainerRegistryResponseAllOf {
	this := ContainerRegistryResponseAllOf{}
	return &this
}

// NewContainerRegistryResponseAllOfWithDefaults instantiates a new ContainerRegistryResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerRegistryResponseAllOfWithDefaults() *ContainerRegistryResponseAllOf {
	this := ContainerRegistryResponseAllOf{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ContainerRegistryResponseAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerRegistryResponseAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ContainerRegistryResponseAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ContainerRegistryResponseAllOf) SetName(v string) {
	o.Name = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *ContainerRegistryResponseAllOf) GetKind() ContainerRegistryKindEnum {
	if o == nil || o.Kind == nil {
		var ret ContainerRegistryKindEnum
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerRegistryResponseAllOf) GetKindOk() (*ContainerRegistryKindEnum, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *ContainerRegistryResponseAllOf) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given ContainerRegistryKindEnum and assigns it to the Kind field.
func (o *ContainerRegistryResponseAllOf) SetKind(v ContainerRegistryKindEnum) {
	o.Kind = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ContainerRegistryResponseAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerRegistryResponseAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ContainerRegistryResponseAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ContainerRegistryResponseAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ContainerRegistryResponseAllOf) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerRegistryResponseAllOf) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ContainerRegistryResponseAllOf) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ContainerRegistryResponseAllOf) SetUrl(v string) {
	o.Url = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContainerRegistryResponseAllOf) GetCluster() Base {
	if o == nil || o.Cluster.Get() == nil {
		var ret Base
		return ret
	}
	return *o.Cluster.Get()
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContainerRegistryResponseAllOf) GetClusterOk() (*Base, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cluster.Get(), o.Cluster.IsSet()
}

// HasCluster returns a boolean if a field has been set.
func (o *ContainerRegistryResponseAllOf) HasCluster() bool {
	if o != nil && o.Cluster.IsSet() {
		return true
	}

	return false
}

// SetCluster gets a reference to the given NullableBase and assigns it to the Cluster field.
func (o *ContainerRegistryResponseAllOf) SetCluster(v Base) {
	o.Cluster.Set(&v)
}

// SetClusterNil sets the value for Cluster to be an explicit nil
func (o *ContainerRegistryResponseAllOf) SetClusterNil() {
	o.Cluster.Set(nil)
}

// UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
func (o *ContainerRegistryResponseAllOf) UnsetCluster() {
	o.Cluster.Unset()
}

func (o ContainerRegistryResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.Cluster.IsSet() {
		toSerialize["cluster"] = o.Cluster.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableContainerRegistryResponseAllOf struct {
	value *ContainerRegistryResponseAllOf
	isSet bool
}

func (v NullableContainerRegistryResponseAllOf) Get() *ContainerRegistryResponseAllOf {
	return v.value
}

func (v *NullableContainerRegistryResponseAllOf) Set(val *ContainerRegistryResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerRegistryResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerRegistryResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerRegistryResponseAllOf(val *ContainerRegistryResponseAllOf) *NullableContainerRegistryResponseAllOf {
	return &NullableContainerRegistryResponseAllOf{value: val, isSet: true}
}

func (v NullableContainerRegistryResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerRegistryResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
