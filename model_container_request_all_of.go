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

// ContainerRequestAllOf struct for ContainerRequestAllOf
type ContainerRequestAllOf struct {
	// name is case insensitive
	Name string `json:"name"`
	// give a description to this container
	Description NullableString `json:"description,omitempty"`
	// id of the linked registry
	RegistryId string `json:"registry_id"`
	// name of the image container
	ImageName string  `json:"image_name"`
	Arguments *string `json:"arguments,omitempty"`
	// unit is millicores (m). 1000m = 1 cpu
	Cpu *int32 `json:"cpu,omitempty"`
	// unit is MB. 1024 MB = 1GB
	Memory *int32 `json:"memory,omitempty"`
	// Minimum number of instances running. This resource auto-scale based on the CPU and Memory consumption. Note: 0 means that there is no container running.
	MinRunningInstances *int32 `json:"min_running_instances,omitempty"`
	// Maximum number of instances running. This resource auto-scale based on the CPU and Memory consumption. Note: -1 means that there is no limit.
	MaxRunningInstances *int32       `json:"max_running_instances,omitempty"`
	Healthcheck         *Healthcheck `json:"healthcheck,omitempty"`
}

// NewContainerRequestAllOf instantiates a new ContainerRequestAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerRequestAllOf(name string, registryId string, imageName string) *ContainerRequestAllOf {
	this := ContainerRequestAllOf{}
	this.Name = name
	this.RegistryId = registryId
	this.ImageName = imageName
	var cpu int32 = 250
	this.Cpu = &cpu
	var memory int32 = 256
	this.Memory = &memory
	var minRunningInstances int32 = 1
	this.MinRunningInstances = &minRunningInstances
	var maxRunningInstances int32 = 1
	this.MaxRunningInstances = &maxRunningInstances
	return &this
}

// NewContainerRequestAllOfWithDefaults instantiates a new ContainerRequestAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerRequestAllOfWithDefaults() *ContainerRequestAllOf {
	this := ContainerRequestAllOf{}
	var cpu int32 = 250
	this.Cpu = &cpu
	var memory int32 = 256
	this.Memory = &memory
	var minRunningInstances int32 = 1
	this.MinRunningInstances = &minRunningInstances
	var maxRunningInstances int32 = 1
	this.MaxRunningInstances = &maxRunningInstances
	return &this
}

// GetName returns the Name field value
func (o *ContainerRequestAllOf) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ContainerRequestAllOf) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ContainerRequestAllOf) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContainerRequestAllOf) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContainerRequestAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ContainerRequestAllOf) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ContainerRequestAllOf) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ContainerRequestAllOf) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ContainerRequestAllOf) UnsetDescription() {
	o.Description.Unset()
}

// GetRegistryId returns the RegistryId field value
func (o *ContainerRequestAllOf) GetRegistryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegistryId
}

// GetRegistryIdOk returns a tuple with the RegistryId field value
// and a boolean to check if the value has been set.
func (o *ContainerRequestAllOf) GetRegistryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegistryId, true
}

// SetRegistryId sets field value
func (o *ContainerRequestAllOf) SetRegistryId(v string) {
	o.RegistryId = v
}

// GetImageName returns the ImageName field value
func (o *ContainerRequestAllOf) GetImageName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageName
}

// GetImageNameOk returns a tuple with the ImageName field value
// and a boolean to check if the value has been set.
func (o *ContainerRequestAllOf) GetImageNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageName, true
}

// SetImageName sets field value
func (o *ContainerRequestAllOf) SetImageName(v string) {
	o.ImageName = v
}

// GetArguments returns the Arguments field value if set, zero value otherwise.
func (o *ContainerRequestAllOf) GetArguments() string {
	if o == nil || o.Arguments == nil {
		var ret string
		return ret
	}
	return *o.Arguments
}

// GetArgumentsOk returns a tuple with the Arguments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerRequestAllOf) GetArgumentsOk() (*string, bool) {
	if o == nil || o.Arguments == nil {
		return nil, false
	}
	return o.Arguments, true
}

// HasArguments returns a boolean if a field has been set.
func (o *ContainerRequestAllOf) HasArguments() bool {
	if o != nil && o.Arguments != nil {
		return true
	}

	return false
}

// SetArguments gets a reference to the given string and assigns it to the Arguments field.
func (o *ContainerRequestAllOf) SetArguments(v string) {
	o.Arguments = &v
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *ContainerRequestAllOf) GetCpu() int32 {
	if o == nil || o.Cpu == nil {
		var ret int32
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerRequestAllOf) GetCpuOk() (*int32, bool) {
	if o == nil || o.Cpu == nil {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *ContainerRequestAllOf) HasCpu() bool {
	if o != nil && o.Cpu != nil {
		return true
	}

	return false
}

// SetCpu gets a reference to the given int32 and assigns it to the Cpu field.
func (o *ContainerRequestAllOf) SetCpu(v int32) {
	o.Cpu = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *ContainerRequestAllOf) GetMemory() int32 {
	if o == nil || o.Memory == nil {
		var ret int32
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerRequestAllOf) GetMemoryOk() (*int32, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *ContainerRequestAllOf) HasMemory() bool {
	if o != nil && o.Memory != nil {
		return true
	}

	return false
}

// SetMemory gets a reference to the given int32 and assigns it to the Memory field.
func (o *ContainerRequestAllOf) SetMemory(v int32) {
	o.Memory = &v
}

// GetMinRunningInstances returns the MinRunningInstances field value if set, zero value otherwise.
func (o *ContainerRequestAllOf) GetMinRunningInstances() int32 {
	if o == nil || o.MinRunningInstances == nil {
		var ret int32
		return ret
	}
	return *o.MinRunningInstances
}

// GetMinRunningInstancesOk returns a tuple with the MinRunningInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerRequestAllOf) GetMinRunningInstancesOk() (*int32, bool) {
	if o == nil || o.MinRunningInstances == nil {
		return nil, false
	}
	return o.MinRunningInstances, true
}

// HasMinRunningInstances returns a boolean if a field has been set.
func (o *ContainerRequestAllOf) HasMinRunningInstances() bool {
	if o != nil && o.MinRunningInstances != nil {
		return true
	}

	return false
}

// SetMinRunningInstances gets a reference to the given int32 and assigns it to the MinRunningInstances field.
func (o *ContainerRequestAllOf) SetMinRunningInstances(v int32) {
	o.MinRunningInstances = &v
}

// GetMaxRunningInstances returns the MaxRunningInstances field value if set, zero value otherwise.
func (o *ContainerRequestAllOf) GetMaxRunningInstances() int32 {
	if o == nil || o.MaxRunningInstances == nil {
		var ret int32
		return ret
	}
	return *o.MaxRunningInstances
}

// GetMaxRunningInstancesOk returns a tuple with the MaxRunningInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerRequestAllOf) GetMaxRunningInstancesOk() (*int32, bool) {
	if o == nil || o.MaxRunningInstances == nil {
		return nil, false
	}
	return o.MaxRunningInstances, true
}

// HasMaxRunningInstances returns a boolean if a field has been set.
func (o *ContainerRequestAllOf) HasMaxRunningInstances() bool {
	if o != nil && o.MaxRunningInstances != nil {
		return true
	}

	return false
}

// SetMaxRunningInstances gets a reference to the given int32 and assigns it to the MaxRunningInstances field.
func (o *ContainerRequestAllOf) SetMaxRunningInstances(v int32) {
	o.MaxRunningInstances = &v
}

// GetHealthcheck returns the Healthcheck field value if set, zero value otherwise.
func (o *ContainerRequestAllOf) GetHealthcheck() Healthcheck {
	if o == nil || o.Healthcheck == nil {
		var ret Healthcheck
		return ret
	}
	return *o.Healthcheck
}

// GetHealthcheckOk returns a tuple with the Healthcheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerRequestAllOf) GetHealthcheckOk() (*Healthcheck, bool) {
	if o == nil || o.Healthcheck == nil {
		return nil, false
	}
	return o.Healthcheck, true
}

// HasHealthcheck returns a boolean if a field has been set.
func (o *ContainerRequestAllOf) HasHealthcheck() bool {
	if o != nil && o.Healthcheck != nil {
		return true
	}

	return false
}

// SetHealthcheck gets a reference to the given Healthcheck and assigns it to the Healthcheck field.
func (o *ContainerRequestAllOf) SetHealthcheck(v Healthcheck) {
	o.Healthcheck = &v
}

func (o ContainerRequestAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if true {
		toSerialize["registry_id"] = o.RegistryId
	}
	if true {
		toSerialize["image_name"] = o.ImageName
	}
	if o.Arguments != nil {
		toSerialize["arguments"] = o.Arguments
	}
	if o.Cpu != nil {
		toSerialize["cpu"] = o.Cpu
	}
	if o.Memory != nil {
		toSerialize["memory"] = o.Memory
	}
	if o.MinRunningInstances != nil {
		toSerialize["min_running_instances"] = o.MinRunningInstances
	}
	if o.MaxRunningInstances != nil {
		toSerialize["max_running_instances"] = o.MaxRunningInstances
	}
	if o.Healthcheck != nil {
		toSerialize["healthcheck"] = o.Healthcheck
	}
	return json.Marshal(toSerialize)
}

type NullableContainerRequestAllOf struct {
	value *ContainerRequestAllOf
	isSet bool
}

func (v NullableContainerRequestAllOf) Get() *ContainerRequestAllOf {
	return v.value
}

func (v *NullableContainerRequestAllOf) Set(val *ContainerRequestAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerRequestAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerRequestAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerRequestAllOf(val *ContainerRequestAllOf) *NullableContainerRequestAllOf {
	return &NullableContainerRequestAllOf{value: val, isSet: true}
}

func (v NullableContainerRequestAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerRequestAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
