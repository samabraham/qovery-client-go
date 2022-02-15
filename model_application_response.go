/*
[BETA] Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is in Beta and still in progress. Some endpoints are not available yet.

API version: 1.0.0
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"encoding/json"
	"time"
)

// ApplicationResponse struct for ApplicationResponse
type ApplicationResponse struct {
	Environment   *ReferenceObject                  `json:"environment,omitempty"`
	GitRepository *ApplicationGitRepositoryResponse `json:"git_repository,omitempty"`
	// Maximum cpu that can be allocated to the application based on organization cluster configuration. unit is millicores (m). 1000m = 1 cpu
	MaximumCpu *int32 `json:"maximum_cpu,omitempty"`
	// Maximum memory that can be allocated to the application based on organization cluster configuration. unit is MB. 1024 MB = 1GB
	MaximumMemory *int32 `json:"maximum_memory,omitempty"`
	// name is case insensitive
	Name *string `json:"name,omitempty"`
	// give a description to this application
	Description *string `json:"description,omitempty"`
	// `DOCKER` requires `dockerfile_path` `BUILDPACKS` does not require any `dockerfile_path`
	BuildMode *string `json:"build_mode,omitempty"`
	// The path of the associated Dockerfile. Only if you are using build_mode = DOCKER
	DockerfilePath *string `json:"dockerfile_path,omitempty"`
	// Development language of the application
	BuildpackLanguage *string `json:"buildpack_language,omitempty"`
	// unit is millicores (m). 1000m = 1 cpu
	Cpu *int32 `json:"cpu,omitempty"`
	// unit is MB. 1024 MB = 1GB
	Memory *int32 `json:"memory,omitempty"`
	// Minimum number of instances running. This resource auto-scale based on the CPU and Memory consumption. Note: 0 means that there is no application running.
	MinRunningInstances *int32 `json:"min_running_instances,omitempty"`
	// Maximum number of instances running. This resource auto-scale based on the CPU and Memory consumption. Note: -1 means that there is no limit.
	MaxRunningInstances *int32       `json:"max_running_instances,omitempty"`
	Healthcheck         *Healthcheck `json:"healthcheck,omitempty"`
	// Specify if the environment preview option is activated or not for this application. If activated, a preview environment will be automatically cloned at each pull request.
	AutoPreview *bool                               `json:"auto_preview,omitempty"`
	Id          string                              `json:"id"`
	CreatedAt   time.Time                           `json:"created_at"`
	UpdatedAt   *time.Time                          `json:"updated_at,omitempty"`
	Storage     []ApplicationStorageResponseStorage `json:"storage,omitempty"`
	Ports       []ApplicationPortResponsePorts      `json:"ports,omitempty"`
}

// NewApplicationResponse instantiates a new ApplicationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationResponse(id string, createdAt time.Time) *ApplicationResponse {
	this := ApplicationResponse{}
	this.Id = id
	this.CreatedAt = createdAt
	return &this
}

// NewApplicationResponseWithDefaults instantiates a new ApplicationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationResponseWithDefaults() *ApplicationResponse {
	this := ApplicationResponse{}
	var maximumCpu int32 = 250
	this.MaximumCpu = &maximumCpu
	var maximumMemory int32 = 256
	this.MaximumMemory = &maximumMemory
	var buildMode string = "BUILDPACKS"
	this.BuildMode = &buildMode
	var cpu int32 = 250
	this.Cpu = &cpu
	var memory int32 = 256
	this.Memory = &memory
	var minRunningInstances int32 = 1
	this.MinRunningInstances = &minRunningInstances
	var maxRunningInstances int32 = 1
	this.MaxRunningInstances = &maxRunningInstances
	var autoPreview bool = true
	this.AutoPreview = &autoPreview
	return &this
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *ApplicationResponse) GetEnvironment() ReferenceObject {
	if o == nil || o.Environment == nil {
		var ret ReferenceObject
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResponse) GetEnvironmentOk() (*ReferenceObject, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *ApplicationResponse) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ReferenceObject and assigns it to the Environment field.
func (o *ApplicationResponse) SetEnvironment(v ReferenceObject) {
	o.Environment = &v
}

// GetGitRepository returns the GitRepository field value if set, zero value otherwise.
func (o *ApplicationResponse) GetGitRepository() ApplicationGitRepositoryResponse {
	if o == nil || o.GitRepository == nil {
		var ret ApplicationGitRepositoryResponse
		return ret
	}
	return *o.GitRepository
}

// GetGitRepositoryOk returns a tuple with the GitRepository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResponse) GetGitRepositoryOk() (*ApplicationGitRepositoryResponse, bool) {
	if o == nil || o.GitRepository == nil {
		return nil, false
	}
	return o.GitRepository, true
}

// HasGitRepository returns a boolean if a field has been set.
func (o *ApplicationResponse) HasGitRepository() bool {
	if o != nil && o.GitRepository != nil {
		return true
	}

	return false
}

// SetGitRepository gets a reference to the given ApplicationGitRepositoryResponse and assigns it to the GitRepository field.
func (o *ApplicationResponse) SetGitRepository(v ApplicationGitRepositoryResponse) {
	o.GitRepository = &v
}

// GetMaximumCpu returns the MaximumCpu field value if set, zero value otherwise.
func (o *ApplicationResponse) GetMaximumCpu() int32 {
	if o == nil || o.MaximumCpu == nil {
		var ret int32
		return ret
	}
	return *o.MaximumCpu
}

// GetMaximumCpuOk returns a tuple with the MaximumCpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResponse) GetMaximumCpuOk() (*int32, bool) {
	if o == nil || o.MaximumCpu == nil {
		return nil, false
	}
	return o.MaximumCpu, true
}

// HasMaximumCpu returns a boolean if a field has been set.
func (o *ApplicationResponse) HasMaximumCpu() bool {
	if o != nil && o.MaximumCpu != nil {
		return true
	}

	return false
}

// SetMaximumCpu gets a reference to the given int32 and assigns it to the MaximumCpu field.
func (o *ApplicationResponse) SetMaximumCpu(v int32) {
	o.MaximumCpu = &v
}

// GetMaximumMemory returns the MaximumMemory field value if set, zero value otherwise.
func (o *ApplicationResponse) GetMaximumMemory() int32 {
	if o == nil || o.MaximumMemory == nil {
		var ret int32
		return ret
	}
	return *o.MaximumMemory
}

// GetMaximumMemoryOk returns a tuple with the MaximumMemory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResponse) GetMaximumMemoryOk() (*int32, bool) {
	if o == nil || o.MaximumMemory == nil {
		return nil, false
	}
	return o.MaximumMemory, true
}

// HasMaximumMemory returns a boolean if a field has been set.
func (o *ApplicationResponse) HasMaximumMemory() bool {
	if o != nil && o.MaximumMemory != nil {
		return true
	}

	return false
}

// SetMaximumMemory gets a reference to the given int32 and assigns it to the MaximumMemory field.
func (o *ApplicationResponse) SetMaximumMemory(v int32) {
	o.MaximumMemory = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApplicationResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApplicationResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApplicationResponse) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApplicationResponse) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApplicationResponse) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApplicationResponse) SetDescription(v string) {
	o.Description = &v
}

// GetBuildMode returns the BuildMode field value if set, zero value otherwise.
func (o *ApplicationResponse) GetBuildMode() string {
	if o == nil || o.BuildMode == nil {
		var ret string
		return ret
	}
	return *o.BuildMode
}

// GetBuildModeOk returns a tuple with the BuildMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResponse) GetBuildModeOk() (*string, bool) {
	if o == nil || o.BuildMode == nil {
		return nil, false
	}
	return o.BuildMode, true
}

// HasBuildMode returns a boolean if a field has been set.
func (o *ApplicationResponse) HasBuildMode() bool {
	if o != nil && o.BuildMode != nil {
		return true
	}

	return false
}

// SetBuildMode gets a reference to the given string and assigns it to the BuildMode field.
func (o *ApplicationResponse) SetBuildMode(v string) {
	o.BuildMode = &v
}

// GetDockerfilePath returns the DockerfilePath field value if set, zero value otherwise.
func (o *ApplicationResponse) GetDockerfilePath() string {
	if o == nil || o.DockerfilePath == nil {
		var ret string
		return ret
	}
	return *o.DockerfilePath
}

// GetDockerfilePathOk returns a tuple with the DockerfilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResponse) GetDockerfilePathOk() (*string, bool) {
	if o == nil || o.DockerfilePath == nil {
		return nil, false
	}
	return o.DockerfilePath, true
}

// HasDockerfilePath returns a boolean if a field has been set.
func (o *ApplicationResponse) HasDockerfilePath() bool {
	if o != nil && o.DockerfilePath != nil {
		return true
	}

	return false
}

// SetDockerfilePath gets a reference to the given string and assigns it to the DockerfilePath field.
func (o *ApplicationResponse) SetDockerfilePath(v string) {
	o.DockerfilePath = &v
}

// GetBuildpackLanguage returns the BuildpackLanguage field value if set, zero value otherwise.
func (o *ApplicationResponse) GetBuildpackLanguage() string {
	if o == nil || o.BuildpackLanguage == nil {
		var ret string
		return ret
	}
	return *o.BuildpackLanguage
}

// GetBuildpackLanguageOk returns a tuple with the BuildpackLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResponse) GetBuildpackLanguageOk() (*string, bool) {
	if o == nil || o.BuildpackLanguage == nil {
		return nil, false
	}
	return o.BuildpackLanguage, true
}

// HasBuildpackLanguage returns a boolean if a field has been set.
func (o *ApplicationResponse) HasBuildpackLanguage() bool {
	if o != nil && o.BuildpackLanguage != nil {
		return true
	}

	return false
}

// SetBuildpackLanguage gets a reference to the given string and assigns it to the BuildpackLanguage field.
func (o *ApplicationResponse) SetBuildpackLanguage(v string) {
	o.BuildpackLanguage = &v
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *ApplicationResponse) GetCpu() int32 {
	if o == nil || o.Cpu == nil {
		var ret int32
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResponse) GetCpuOk() (*int32, bool) {
	if o == nil || o.Cpu == nil {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *ApplicationResponse) HasCpu() bool {
	if o != nil && o.Cpu != nil {
		return true
	}

	return false
}

// SetCpu gets a reference to the given int32 and assigns it to the Cpu field.
func (o *ApplicationResponse) SetCpu(v int32) {
	o.Cpu = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *ApplicationResponse) GetMemory() int32 {
	if o == nil || o.Memory == nil {
		var ret int32
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResponse) GetMemoryOk() (*int32, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *ApplicationResponse) HasMemory() bool {
	if o != nil && o.Memory != nil {
		return true
	}

	return false
}

// SetMemory gets a reference to the given int32 and assigns it to the Memory field.
func (o *ApplicationResponse) SetMemory(v int32) {
	o.Memory = &v
}

// GetMinRunningInstances returns the MinRunningInstances field value if set, zero value otherwise.
func (o *ApplicationResponse) GetMinRunningInstances() int32 {
	if o == nil || o.MinRunningInstances == nil {
		var ret int32
		return ret
	}
	return *o.MinRunningInstances
}

// GetMinRunningInstancesOk returns a tuple with the MinRunningInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResponse) GetMinRunningInstancesOk() (*int32, bool) {
	if o == nil || o.MinRunningInstances == nil {
		return nil, false
	}
	return o.MinRunningInstances, true
}

// HasMinRunningInstances returns a boolean if a field has been set.
func (o *ApplicationResponse) HasMinRunningInstances() bool {
	if o != nil && o.MinRunningInstances != nil {
		return true
	}

	return false
}

// SetMinRunningInstances gets a reference to the given int32 and assigns it to the MinRunningInstances field.
func (o *ApplicationResponse) SetMinRunningInstances(v int32) {
	o.MinRunningInstances = &v
}

// GetMaxRunningInstances returns the MaxRunningInstances field value if set, zero value otherwise.
func (o *ApplicationResponse) GetMaxRunningInstances() int32 {
	if o == nil || o.MaxRunningInstances == nil {
		var ret int32
		return ret
	}
	return *o.MaxRunningInstances
}

// GetMaxRunningInstancesOk returns a tuple with the MaxRunningInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResponse) GetMaxRunningInstancesOk() (*int32, bool) {
	if o == nil || o.MaxRunningInstances == nil {
		return nil, false
	}
	return o.MaxRunningInstances, true
}

// HasMaxRunningInstances returns a boolean if a field has been set.
func (o *ApplicationResponse) HasMaxRunningInstances() bool {
	if o != nil && o.MaxRunningInstances != nil {
		return true
	}

	return false
}

// SetMaxRunningInstances gets a reference to the given int32 and assigns it to the MaxRunningInstances field.
func (o *ApplicationResponse) SetMaxRunningInstances(v int32) {
	o.MaxRunningInstances = &v
}

// GetHealthcheck returns the Healthcheck field value if set, zero value otherwise.
func (o *ApplicationResponse) GetHealthcheck() Healthcheck {
	if o == nil || o.Healthcheck == nil {
		var ret Healthcheck
		return ret
	}
	return *o.Healthcheck
}

// GetHealthcheckOk returns a tuple with the Healthcheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResponse) GetHealthcheckOk() (*Healthcheck, bool) {
	if o == nil || o.Healthcheck == nil {
		return nil, false
	}
	return o.Healthcheck, true
}

// HasHealthcheck returns a boolean if a field has been set.
func (o *ApplicationResponse) HasHealthcheck() bool {
	if o != nil && o.Healthcheck != nil {
		return true
	}

	return false
}

// SetHealthcheck gets a reference to the given Healthcheck and assigns it to the Healthcheck field.
func (o *ApplicationResponse) SetHealthcheck(v Healthcheck) {
	o.Healthcheck = &v
}

// GetAutoPreview returns the AutoPreview field value if set, zero value otherwise.
func (o *ApplicationResponse) GetAutoPreview() bool {
	if o == nil || o.AutoPreview == nil {
		var ret bool
		return ret
	}
	return *o.AutoPreview
}

// GetAutoPreviewOk returns a tuple with the AutoPreview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResponse) GetAutoPreviewOk() (*bool, bool) {
	if o == nil || o.AutoPreview == nil {
		return nil, false
	}
	return o.AutoPreview, true
}

// HasAutoPreview returns a boolean if a field has been set.
func (o *ApplicationResponse) HasAutoPreview() bool {
	if o != nil && o.AutoPreview != nil {
		return true
	}

	return false
}

// SetAutoPreview gets a reference to the given bool and assigns it to the AutoPreview field.
func (o *ApplicationResponse) SetAutoPreview(v bool) {
	o.AutoPreview = &v
}

// GetId returns the Id field value
func (o *ApplicationResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ApplicationResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ApplicationResponse) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ApplicationResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ApplicationResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ApplicationResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ApplicationResponse) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ApplicationResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ApplicationResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *ApplicationResponse) GetStorage() []ApplicationStorageResponseStorage {
	if o == nil || o.Storage == nil {
		var ret []ApplicationStorageResponseStorage
		return ret
	}
	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResponse) GetStorageOk() ([]ApplicationStorageResponseStorage, bool) {
	if o == nil || o.Storage == nil {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *ApplicationResponse) HasStorage() bool {
	if o != nil && o.Storage != nil {
		return true
	}

	return false
}

// SetStorage gets a reference to the given []ApplicationStorageResponseStorage and assigns it to the Storage field.
func (o *ApplicationResponse) SetStorage(v []ApplicationStorageResponseStorage) {
	o.Storage = v
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *ApplicationResponse) GetPorts() []ApplicationPortResponsePorts {
	if o == nil || o.Ports == nil {
		var ret []ApplicationPortResponsePorts
		return ret
	}
	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResponse) GetPortsOk() ([]ApplicationPortResponsePorts, bool) {
	if o == nil || o.Ports == nil {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *ApplicationResponse) HasPorts() bool {
	if o != nil && o.Ports != nil {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []ApplicationPortResponsePorts and assigns it to the Ports field.
func (o *ApplicationResponse) SetPorts(v []ApplicationPortResponsePorts) {
	o.Ports = v
}

func (o ApplicationResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	if o.GitRepository != nil {
		toSerialize["git_repository"] = o.GitRepository
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
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.BuildMode != nil {
		toSerialize["build_mode"] = o.BuildMode
	}
	if o.DockerfilePath != nil {
		toSerialize["dockerfile_path"] = o.DockerfilePath
	}
	if o.BuildpackLanguage != nil {
		toSerialize["buildpack_language"] = o.BuildpackLanguage
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
	if o.AutoPreview != nil {
		toSerialize["auto_preview"] = o.AutoPreview
	}
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
	if o.Ports != nil {
		toSerialize["ports"] = o.Ports
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationResponse struct {
	value *ApplicationResponse
	isSet bool
}

func (v NullableApplicationResponse) Get() *ApplicationResponse {
	return v.value
}

func (v *NullableApplicationResponse) Set(val *ApplicationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationResponse(val *ApplicationResponse) *NullableApplicationResponse {
	return &NullableApplicationResponse{value: val, isSet: true}
}

func (v NullableApplicationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
