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

// ContainerRegistryResponse struct for ContainerRegistryResponse
type ContainerRegistryResponse struct {
	Id          string                     `json:"id"`
	CreatedAt   time.Time                  `json:"created_at"`
	UpdatedAt   *time.Time                 `json:"updated_at,omitempty"`
	Name        *string                    `json:"name,omitempty"`
	Kind        *ContainerRegistryKindEnum `json:"kind,omitempty"`
	Description *string                    `json:"description,omitempty"`
	// URL of the container registry
	Url *string `json:"url,omitempty"`
}

// NewContainerRegistryResponse instantiates a new ContainerRegistryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerRegistryResponse(id string, createdAt time.Time) *ContainerRegistryResponse {
	this := ContainerRegistryResponse{}
	this.Id = id
	this.CreatedAt = createdAt
	return &this
}

// NewContainerRegistryResponseWithDefaults instantiates a new ContainerRegistryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerRegistryResponseWithDefaults() *ContainerRegistryResponse {
	this := ContainerRegistryResponse{}
	return &this
}

// GetId returns the Id field value
func (o *ContainerRegistryResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ContainerRegistryResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ContainerRegistryResponse) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ContainerRegistryResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ContainerRegistryResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ContainerRegistryResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ContainerRegistryResponse) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerRegistryResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ContainerRegistryResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ContainerRegistryResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ContainerRegistryResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerRegistryResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ContainerRegistryResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ContainerRegistryResponse) SetName(v string) {
	o.Name = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *ContainerRegistryResponse) GetKind() ContainerRegistryKindEnum {
	if o == nil || o.Kind == nil {
		var ret ContainerRegistryKindEnum
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerRegistryResponse) GetKindOk() (*ContainerRegistryKindEnum, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *ContainerRegistryResponse) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given ContainerRegistryKindEnum and assigns it to the Kind field.
func (o *ContainerRegistryResponse) SetKind(v ContainerRegistryKindEnum) {
	o.Kind = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ContainerRegistryResponse) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerRegistryResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ContainerRegistryResponse) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ContainerRegistryResponse) SetDescription(v string) {
	o.Description = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ContainerRegistryResponse) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerRegistryResponse) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ContainerRegistryResponse) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ContainerRegistryResponse) SetUrl(v string) {
	o.Url = &v
}

func (o ContainerRegistryResponse) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullableContainerRegistryResponse struct {
	value *ContainerRegistryResponse
	isSet bool
}

func (v NullableContainerRegistryResponse) Get() *ContainerRegistryResponse {
	return v.value
}

func (v *NullableContainerRegistryResponse) Set(val *ContainerRegistryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerRegistryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerRegistryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerRegistryResponse(val *ContainerRegistryResponse) *NullableContainerRegistryResponse {
	return &NullableContainerRegistryResponse{value: val, isSet: true}
}

func (v NullableContainerRegistryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerRegistryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
