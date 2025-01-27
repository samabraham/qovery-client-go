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

// ContainerResponse struct for ContainerResponse
type ContainerResponse struct {
	Id          string                       `json:"id"`
	CreatedAt   time.Time                    `json:"created_at"`
	UpdatedAt   *time.Time                   `json:"updated_at,omitempty"`
	Storage     []ServiceStorageStorageInner `json:"storage,omitempty"`
	Environment ReferenceObject              `json:"environment"`
	Registry    ReferenceObject              `json:"registry"`
	// Maximum cpu that can be allocated to the container based on organization cluster configuration. unit is millicores (m). 1000m = 1 cpu
	MaximumCpu int32 `json:"maximum_cpu"`
	// Maximum memory that can be allocated to the container based on organization cluster configuration. unit is MB. 1024 MB = 1GB
	MaximumMemory int32 `json:"maximum_memory"`
	// name is case insensitive
	Name string `json:"name"`
	// name of the image container
	ImageName string `json:"image_name"`
	// tag of the image container
	Tag       string   `json:"tag"`
	Arguments []string `json:"arguments,omitempty"`
	// optional entrypoint when launching container
	Entrypoint *string `json:"entrypoint,omitempty"`
	// unit is millicores (m). 1000m = 1 cpu
	Cpu int32 `json:"cpu"`
	// unit is MB. 1024 MB = 1GB
	Memory int32 `json:"memory"`
	// Minimum number of instances running. This resource auto-scale based on the CPU and Memory consumption. Note: 0 means that there is no container running.
	MinRunningInstances int32 `json:"min_running_instances"`
	// Maximum number of instances running. This resource auto-scale based on the CPU and Memory consumption. Note: -1 means that there is no limit.
	MaxRunningInstances int32 `json:"max_running_instances"`
	// Indicates if the 'environment preview option' is enabled for this container.   If enabled, a preview environment will be automatically cloned when `/preview` endpoint is called.   If not specified, it takes the value of the `auto_preview` property from the associated environment.
	AutoPreview bool          `json:"auto_preview"`
	Ports       []ServicePort `json:"ports,omitempty"`
}

// NewContainerResponse instantiates a new ContainerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerResponse(id string, createdAt time.Time, environment ReferenceObject, registry ReferenceObject, maximumCpu int32, maximumMemory int32, name string, imageName string, tag string, cpu int32, memory int32, minRunningInstances int32, maxRunningInstances int32, autoPreview bool) *ContainerResponse {
	this := ContainerResponse{}
	this.Id = id
	this.CreatedAt = createdAt
	this.Environment = environment
	this.Registry = registry
	this.MaximumCpu = maximumCpu
	this.MaximumMemory = maximumMemory
	this.Name = name
	this.ImageName = imageName
	this.Tag = tag
	this.Cpu = cpu
	this.Memory = memory
	this.MinRunningInstances = minRunningInstances
	this.MaxRunningInstances = maxRunningInstances
	this.AutoPreview = autoPreview
	return &this
}

// NewContainerResponseWithDefaults instantiates a new ContainerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerResponseWithDefaults() *ContainerResponse {
	this := ContainerResponse{}
	var minRunningInstances int32 = 1
	this.MinRunningInstances = minRunningInstances
	var maxRunningInstances int32 = 1
	this.MaxRunningInstances = maxRunningInstances
	return &this
}

// GetId returns the Id field value
func (o *ContainerResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ContainerResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ContainerResponse) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ContainerResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ContainerResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ContainerResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ContainerResponse) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ContainerResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ContainerResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *ContainerResponse) GetStorage() []ServiceStorageStorageInner {
	if o == nil || o.Storage == nil {
		var ret []ServiceStorageStorageInner
		return ret
	}
	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerResponse) GetStorageOk() ([]ServiceStorageStorageInner, bool) {
	if o == nil || o.Storage == nil {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *ContainerResponse) HasStorage() bool {
	if o != nil && o.Storage != nil {
		return true
	}

	return false
}

// SetStorage gets a reference to the given []ServiceStorageStorageInner and assigns it to the Storage field.
func (o *ContainerResponse) SetStorage(v []ServiceStorageStorageInner) {
	o.Storage = v
}

// GetEnvironment returns the Environment field value
func (o *ContainerResponse) GetEnvironment() ReferenceObject {
	if o == nil {
		var ret ReferenceObject
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *ContainerResponse) GetEnvironmentOk() (*ReferenceObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *ContainerResponse) SetEnvironment(v ReferenceObject) {
	o.Environment = v
}

// GetRegistry returns the Registry field value
func (o *ContainerResponse) GetRegistry() ReferenceObject {
	if o == nil {
		var ret ReferenceObject
		return ret
	}

	return o.Registry
}

// GetRegistryOk returns a tuple with the Registry field value
// and a boolean to check if the value has been set.
func (o *ContainerResponse) GetRegistryOk() (*ReferenceObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Registry, true
}

// SetRegistry sets field value
func (o *ContainerResponse) SetRegistry(v ReferenceObject) {
	o.Registry = v
}

// GetMaximumCpu returns the MaximumCpu field value
func (o *ContainerResponse) GetMaximumCpu() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaximumCpu
}

// GetMaximumCpuOk returns a tuple with the MaximumCpu field value
// and a boolean to check if the value has been set.
func (o *ContainerResponse) GetMaximumCpuOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaximumCpu, true
}

// SetMaximumCpu sets field value
func (o *ContainerResponse) SetMaximumCpu(v int32) {
	o.MaximumCpu = v
}

// GetMaximumMemory returns the MaximumMemory field value
func (o *ContainerResponse) GetMaximumMemory() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaximumMemory
}

// GetMaximumMemoryOk returns a tuple with the MaximumMemory field value
// and a boolean to check if the value has been set.
func (o *ContainerResponse) GetMaximumMemoryOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaximumMemory, true
}

// SetMaximumMemory sets field value
func (o *ContainerResponse) SetMaximumMemory(v int32) {
	o.MaximumMemory = v
}

// GetName returns the Name field value
func (o *ContainerResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ContainerResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ContainerResponse) SetName(v string) {
	o.Name = v
}

// GetImageName returns the ImageName field value
func (o *ContainerResponse) GetImageName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageName
}

// GetImageNameOk returns a tuple with the ImageName field value
// and a boolean to check if the value has been set.
func (o *ContainerResponse) GetImageNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageName, true
}

// SetImageName sets field value
func (o *ContainerResponse) SetImageName(v string) {
	o.ImageName = v
}

// GetTag returns the Tag field value
func (o *ContainerResponse) GetTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value
// and a boolean to check if the value has been set.
func (o *ContainerResponse) GetTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tag, true
}

// SetTag sets field value
func (o *ContainerResponse) SetTag(v string) {
	o.Tag = v
}

// GetArguments returns the Arguments field value if set, zero value otherwise.
func (o *ContainerResponse) GetArguments() []string {
	if o == nil || o.Arguments == nil {
		var ret []string
		return ret
	}
	return o.Arguments
}

// GetArgumentsOk returns a tuple with the Arguments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerResponse) GetArgumentsOk() ([]string, bool) {
	if o == nil || o.Arguments == nil {
		return nil, false
	}
	return o.Arguments, true
}

// HasArguments returns a boolean if a field has been set.
func (o *ContainerResponse) HasArguments() bool {
	if o != nil && o.Arguments != nil {
		return true
	}

	return false
}

// SetArguments gets a reference to the given []string and assigns it to the Arguments field.
func (o *ContainerResponse) SetArguments(v []string) {
	o.Arguments = v
}

// GetEntrypoint returns the Entrypoint field value if set, zero value otherwise.
func (o *ContainerResponse) GetEntrypoint() string {
	if o == nil || o.Entrypoint == nil {
		var ret string
		return ret
	}
	return *o.Entrypoint
}

// GetEntrypointOk returns a tuple with the Entrypoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerResponse) GetEntrypointOk() (*string, bool) {
	if o == nil || o.Entrypoint == nil {
		return nil, false
	}
	return o.Entrypoint, true
}

// HasEntrypoint returns a boolean if a field has been set.
func (o *ContainerResponse) HasEntrypoint() bool {
	if o != nil && o.Entrypoint != nil {
		return true
	}

	return false
}

// SetEntrypoint gets a reference to the given string and assigns it to the Entrypoint field.
func (o *ContainerResponse) SetEntrypoint(v string) {
	o.Entrypoint = &v
}

// GetCpu returns the Cpu field value
func (o *ContainerResponse) GetCpu() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value
// and a boolean to check if the value has been set.
func (o *ContainerResponse) GetCpuOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cpu, true
}

// SetCpu sets field value
func (o *ContainerResponse) SetCpu(v int32) {
	o.Cpu = v
}

// GetMemory returns the Memory field value
func (o *ContainerResponse) GetMemory() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value
// and a boolean to check if the value has been set.
func (o *ContainerResponse) GetMemoryOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Memory, true
}

// SetMemory sets field value
func (o *ContainerResponse) SetMemory(v int32) {
	o.Memory = v
}

// GetMinRunningInstances returns the MinRunningInstances field value
func (o *ContainerResponse) GetMinRunningInstances() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MinRunningInstances
}

// GetMinRunningInstancesOk returns a tuple with the MinRunningInstances field value
// and a boolean to check if the value has been set.
func (o *ContainerResponse) GetMinRunningInstancesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinRunningInstances, true
}

// SetMinRunningInstances sets field value
func (o *ContainerResponse) SetMinRunningInstances(v int32) {
	o.MinRunningInstances = v
}

// GetMaxRunningInstances returns the MaxRunningInstances field value
func (o *ContainerResponse) GetMaxRunningInstances() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxRunningInstances
}

// GetMaxRunningInstancesOk returns a tuple with the MaxRunningInstances field value
// and a boolean to check if the value has been set.
func (o *ContainerResponse) GetMaxRunningInstancesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxRunningInstances, true
}

// SetMaxRunningInstances sets field value
func (o *ContainerResponse) SetMaxRunningInstances(v int32) {
	o.MaxRunningInstances = v
}

// GetAutoPreview returns the AutoPreview field value
func (o *ContainerResponse) GetAutoPreview() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AutoPreview
}

// GetAutoPreviewOk returns a tuple with the AutoPreview field value
// and a boolean to check if the value has been set.
func (o *ContainerResponse) GetAutoPreviewOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoPreview, true
}

// SetAutoPreview sets field value
func (o *ContainerResponse) SetAutoPreview(v bool) {
	o.AutoPreview = v
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *ContainerResponse) GetPorts() []ServicePort {
	if o == nil || o.Ports == nil {
		var ret []ServicePort
		return ret
	}
	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerResponse) GetPortsOk() ([]ServicePort, bool) {
	if o == nil || o.Ports == nil {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *ContainerResponse) HasPorts() bool {
	if o != nil && o.Ports != nil {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []ServicePort and assigns it to the Ports field.
func (o *ContainerResponse) SetPorts(v []ServicePort) {
	o.Ports = v
}

func (o ContainerResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.Storage != nil {
		toSerialize["storage"] = o.Storage
	}
	if true {
		toSerialize["environment"] = o.Environment
	}
	if true {
		toSerialize["registry"] = o.Registry
	}
	if true {
		toSerialize["maximum_cpu"] = o.MaximumCpu
	}
	if true {
		toSerialize["maximum_memory"] = o.MaximumMemory
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["image_name"] = o.ImageName
	}
	if true {
		toSerialize["tag"] = o.Tag
	}
	if o.Arguments != nil {
		toSerialize["arguments"] = o.Arguments
	}
	if o.Entrypoint != nil {
		toSerialize["entrypoint"] = o.Entrypoint
	}
	if true {
		toSerialize["cpu"] = o.Cpu
	}
	if true {
		toSerialize["memory"] = o.Memory
	}
	if true {
		toSerialize["min_running_instances"] = o.MinRunningInstances
	}
	if true {
		toSerialize["max_running_instances"] = o.MaxRunningInstances
	}
	if true {
		toSerialize["auto_preview"] = o.AutoPreview
	}
	if o.Ports != nil {
		toSerialize["ports"] = o.Ports
	}
	return json.Marshal(toSerialize)
}

type NullableContainerResponse struct {
	value *ContainerResponse
	isSet bool
}

func (v NullableContainerResponse) Get() *ContainerResponse {
	return v.value
}

func (v *NullableContainerResponse) Set(val *ContainerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerResponse(val *ContainerResponse) *NullableContainerResponse {
	return &NullableContainerResponse{value: val, isSet: true}
}

func (v NullableContainerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
