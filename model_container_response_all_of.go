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

// ContainerResponseAllOf struct for ContainerResponseAllOf
type ContainerResponseAllOf struct {
	Environment *ReferenceObject `json:"environment,omitempty"`
	// Maximum cpu that can be allocated to the container based on organization cluster configuration. unit is millicores (m). 1000m = 1 cpu
	MaximumCpu *int32 `json:"maximum_cpu,omitempty"`
	// Maximum memory that can be allocated to the container based on organization cluster configuration. unit is MB. 1024 MB = 1GB
	MaximumMemory *int32 `json:"maximum_memory,omitempty"`
	// name is case insensitive
	Name *string `json:"name,omitempty"`
	// give a description to this container
	Description NullableString `json:"description,omitempty"`
	// id of the linked registry
	RegistryId *string `json:"registry_id,omitempty"`
	// name of the image container
	ImageName *string `json:"image_name,omitempty"`
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

// NewContainerResponseAllOf instantiates a new ContainerResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerResponseAllOf() *ContainerResponseAllOf {
	this := ContainerResponseAllOf{}
	var maximumCpu int32 = 250
	this.MaximumCpu = &maximumCpu
	var maximumMemory int32 = 256
	this.MaximumMemory = &maximumMemory
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

// NewContainerResponseAllOfWithDefaults instantiates a new ContainerResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerResponseAllOfWithDefaults() *ContainerResponseAllOf {
	this := ContainerResponseAllOf{}
	var maximumCpu int32 = 250
	this.MaximumCpu = &maximumCpu
	var maximumMemory int32 = 256
	this.MaximumMemory = &maximumMemory
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

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *ContainerResponseAllOf) GetEnvironment() ReferenceObject {
	if o == nil || o.Environment == nil {
		var ret ReferenceObject
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerResponseAllOf) GetEnvironmentOk() (*ReferenceObject, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *ContainerResponseAllOf) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ReferenceObject and assigns it to the Environment field.
func (o *ContainerResponseAllOf) SetEnvironment(v ReferenceObject) {
	o.Environment = &v
}

// GetMaximumCpu returns the MaximumCpu field value if set, zero value otherwise.
func (o *ContainerResponseAllOf) GetMaximumCpu() int32 {
	if o == nil || o.MaximumCpu == nil {
		var ret int32
		return ret
	}
	return *o.MaximumCpu
}

// GetMaximumCpuOk returns a tuple with the MaximumCpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerResponseAllOf) GetMaximumCpuOk() (*int32, bool) {
	if o == nil || o.MaximumCpu == nil {
		return nil, false
	}
	return o.MaximumCpu, true
}

// HasMaximumCpu returns a boolean if a field has been set.
func (o *ContainerResponseAllOf) HasMaximumCpu() bool {
	if o != nil && o.MaximumCpu != nil {
		return true
	}

	return false
}

// SetMaximumCpu gets a reference to the given int32 and assigns it to the MaximumCpu field.
func (o *ContainerResponseAllOf) SetMaximumCpu(v int32) {
	o.MaximumCpu = &v
}

// GetMaximumMemory returns the MaximumMemory field value if set, zero value otherwise.
func (o *ContainerResponseAllOf) GetMaximumMemory() int32 {
	if o == nil || o.MaximumMemory == nil {
		var ret int32
		return ret
	}
	return *o.MaximumMemory
}

// GetMaximumMemoryOk returns a tuple with the MaximumMemory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerResponseAllOf) GetMaximumMemoryOk() (*int32, bool) {
	if o == nil || o.MaximumMemory == nil {
		return nil, false
	}
	return o.MaximumMemory, true
}

// HasMaximumMemory returns a boolean if a field has been set.
func (o *ContainerResponseAllOf) HasMaximumMemory() bool {
	if o != nil && o.MaximumMemory != nil {
		return true
	}

	return false
}

// SetMaximumMemory gets a reference to the given int32 and assigns it to the MaximumMemory field.
func (o *ContainerResponseAllOf) SetMaximumMemory(v int32) {
	o.MaximumMemory = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ContainerResponseAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerResponseAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ContainerResponseAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ContainerResponseAllOf) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContainerResponseAllOf) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContainerResponseAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ContainerResponseAllOf) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ContainerResponseAllOf) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ContainerResponseAllOf) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ContainerResponseAllOf) UnsetDescription() {
	o.Description.Unset()
}

// GetRegistryId returns the RegistryId field value if set, zero value otherwise.
func (o *ContainerResponseAllOf) GetRegistryId() string {
	if o == nil || o.RegistryId == nil {
		var ret string
		return ret
	}
	return *o.RegistryId
}

// GetRegistryIdOk returns a tuple with the RegistryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerResponseAllOf) GetRegistryIdOk() (*string, bool) {
	if o == nil || o.RegistryId == nil {
		return nil, false
	}
	return o.RegistryId, true
}

// HasRegistryId returns a boolean if a field has been set.
func (o *ContainerResponseAllOf) HasRegistryId() bool {
	if o != nil && o.RegistryId != nil {
		return true
	}

	return false
}

// SetRegistryId gets a reference to the given string and assigns it to the RegistryId field.
func (o *ContainerResponseAllOf) SetRegistryId(v string) {
	o.RegistryId = &v
}

// GetImageName returns the ImageName field value if set, zero value otherwise.
func (o *ContainerResponseAllOf) GetImageName() string {
	if o == nil || o.ImageName == nil {
		var ret string
		return ret
	}
	return *o.ImageName
}

// GetImageNameOk returns a tuple with the ImageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerResponseAllOf) GetImageNameOk() (*string, bool) {
	if o == nil || o.ImageName == nil {
		return nil, false
	}
	return o.ImageName, true
}

// HasImageName returns a boolean if a field has been set.
func (o *ContainerResponseAllOf) HasImageName() bool {
	if o != nil && o.ImageName != nil {
		return true
	}

	return false
}

// SetImageName gets a reference to the given string and assigns it to the ImageName field.
func (o *ContainerResponseAllOf) SetImageName(v string) {
	o.ImageName = &v
}

// GetArguments returns the Arguments field value if set, zero value otherwise.
func (o *ContainerResponseAllOf) GetArguments() string {
	if o == nil || o.Arguments == nil {
		var ret string
		return ret
	}
	return *o.Arguments
}

// GetArgumentsOk returns a tuple with the Arguments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerResponseAllOf) GetArgumentsOk() (*string, bool) {
	if o == nil || o.Arguments == nil {
		return nil, false
	}
	return o.Arguments, true
}

// HasArguments returns a boolean if a field has been set.
func (o *ContainerResponseAllOf) HasArguments() bool {
	if o != nil && o.Arguments != nil {
		return true
	}

	return false
}

// SetArguments gets a reference to the given string and assigns it to the Arguments field.
func (o *ContainerResponseAllOf) SetArguments(v string) {
	o.Arguments = &v
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *ContainerResponseAllOf) GetCpu() int32 {
	if o == nil || o.Cpu == nil {
		var ret int32
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerResponseAllOf) GetCpuOk() (*int32, bool) {
	if o == nil || o.Cpu == nil {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *ContainerResponseAllOf) HasCpu() bool {
	if o != nil && o.Cpu != nil {
		return true
	}

	return false
}

// SetCpu gets a reference to the given int32 and assigns it to the Cpu field.
func (o *ContainerResponseAllOf) SetCpu(v int32) {
	o.Cpu = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *ContainerResponseAllOf) GetMemory() int32 {
	if o == nil || o.Memory == nil {
		var ret int32
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerResponseAllOf) GetMemoryOk() (*int32, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *ContainerResponseAllOf) HasMemory() bool {
	if o != nil && o.Memory != nil {
		return true
	}

	return false
}

// SetMemory gets a reference to the given int32 and assigns it to the Memory field.
func (o *ContainerResponseAllOf) SetMemory(v int32) {
	o.Memory = &v
}

// GetMinRunningInstances returns the MinRunningInstances field value if set, zero value otherwise.
func (o *ContainerResponseAllOf) GetMinRunningInstances() int32 {
	if o == nil || o.MinRunningInstances == nil {
		var ret int32
		return ret
	}
	return *o.MinRunningInstances
}

// GetMinRunningInstancesOk returns a tuple with the MinRunningInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerResponseAllOf) GetMinRunningInstancesOk() (*int32, bool) {
	if o == nil || o.MinRunningInstances == nil {
		return nil, false
	}
	return o.MinRunningInstances, true
}

// HasMinRunningInstances returns a boolean if a field has been set.
func (o *ContainerResponseAllOf) HasMinRunningInstances() bool {
	if o != nil && o.MinRunningInstances != nil {
		return true
	}

	return false
}

// SetMinRunningInstances gets a reference to the given int32 and assigns it to the MinRunningInstances field.
func (o *ContainerResponseAllOf) SetMinRunningInstances(v int32) {
	o.MinRunningInstances = &v
}

// GetMaxRunningInstances returns the MaxRunningInstances field value if set, zero value otherwise.
func (o *ContainerResponseAllOf) GetMaxRunningInstances() int32 {
	if o == nil || o.MaxRunningInstances == nil {
		var ret int32
		return ret
	}
	return *o.MaxRunningInstances
}

// GetMaxRunningInstancesOk returns a tuple with the MaxRunningInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerResponseAllOf) GetMaxRunningInstancesOk() (*int32, bool) {
	if o == nil || o.MaxRunningInstances == nil {
		return nil, false
	}
	return o.MaxRunningInstances, true
}

// HasMaxRunningInstances returns a boolean if a field has been set.
func (o *ContainerResponseAllOf) HasMaxRunningInstances() bool {
	if o != nil && o.MaxRunningInstances != nil {
		return true
	}

	return false
}

// SetMaxRunningInstances gets a reference to the given int32 and assigns it to the MaxRunningInstances field.
func (o *ContainerResponseAllOf) SetMaxRunningInstances(v int32) {
	o.MaxRunningInstances = &v
}

// GetHealthcheck returns the Healthcheck field value if set, zero value otherwise.
func (o *ContainerResponseAllOf) GetHealthcheck() Healthcheck {
	if o == nil || o.Healthcheck == nil {
		var ret Healthcheck
		return ret
	}
	return *o.Healthcheck
}

// GetHealthcheckOk returns a tuple with the Healthcheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerResponseAllOf) GetHealthcheckOk() (*Healthcheck, bool) {
	if o == nil || o.Healthcheck == nil {
		return nil, false
	}
	return o.Healthcheck, true
}

// HasHealthcheck returns a boolean if a field has been set.
func (o *ContainerResponseAllOf) HasHealthcheck() bool {
	if o != nil && o.Healthcheck != nil {
		return true
	}

	return false
}

// SetHealthcheck gets a reference to the given Healthcheck and assigns it to the Healthcheck field.
func (o *ContainerResponseAllOf) SetHealthcheck(v Healthcheck) {
	o.Healthcheck = &v
}

func (o ContainerResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	if o.MaximumCpu != nil {
		toSerialize["maximum_cpu"] = o.MaximumCpu
	}
	if o.MaximumMemory != nil {
		toSerialize["maximum_memory"] = o.MaximumMemory
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.RegistryId != nil {
		toSerialize["registry_id"] = o.RegistryId
	}
	if o.ImageName != nil {
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

type NullableContainerResponseAllOf struct {
	value *ContainerResponseAllOf
	isSet bool
}

func (v NullableContainerResponseAllOf) Get() *ContainerResponseAllOf {
	return v.value
}

func (v *NullableContainerResponseAllOf) Set(val *ContainerResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerResponseAllOf(val *ContainerResponseAllOf) *NullableContainerResponseAllOf {
	return &NullableContainerResponseAllOf{value: val, isSet: true}
}

func (v NullableContainerResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}