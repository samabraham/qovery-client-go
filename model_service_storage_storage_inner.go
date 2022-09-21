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

// ServiceStorageStorageInner struct for ServiceStorageStorageInner
type ServiceStorageStorageInner struct {
	Id   string          `json:"id"`
	Type StorageTypeEnum `json:"type"`
	// unit is GB
	Size       int32  `json:"size"`
	MountPoint string `json:"mount_point"`
}

// NewServiceStorageStorageInner instantiates a new ServiceStorageStorageInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceStorageStorageInner(id string, type_ StorageTypeEnum, size int32, mountPoint string) *ServiceStorageStorageInner {
	this := ServiceStorageStorageInner{}
	this.Id = id
	this.Type = type_
	this.Size = size
	this.MountPoint = mountPoint
	return &this
}

// NewServiceStorageStorageInnerWithDefaults instantiates a new ServiceStorageStorageInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceStorageStorageInnerWithDefaults() *ServiceStorageStorageInner {
	this := ServiceStorageStorageInner{}
	return &this
}

// GetId returns the Id field value
func (o *ServiceStorageStorageInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ServiceStorageStorageInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ServiceStorageStorageInner) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *ServiceStorageStorageInner) GetType() StorageTypeEnum {
	if o == nil {
		var ret StorageTypeEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ServiceStorageStorageInner) GetTypeOk() (*StorageTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ServiceStorageStorageInner) SetType(v StorageTypeEnum) {
	o.Type = v
}

// GetSize returns the Size field value
func (o *ServiceStorageStorageInner) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *ServiceStorageStorageInner) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *ServiceStorageStorageInner) SetSize(v int32) {
	o.Size = v
}

// GetMountPoint returns the MountPoint field value
func (o *ServiceStorageStorageInner) GetMountPoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MountPoint
}

// GetMountPointOk returns a tuple with the MountPoint field value
// and a boolean to check if the value has been set.
func (o *ServiceStorageStorageInner) GetMountPointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MountPoint, true
}

// SetMountPoint sets field value
func (o *ServiceStorageStorageInner) SetMountPoint(v string) {
	o.MountPoint = v
}

func (o ServiceStorageStorageInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["size"] = o.Size
	}
	if true {
		toSerialize["mount_point"] = o.MountPoint
	}
	return json.Marshal(toSerialize)
}

type NullableServiceStorageStorageInner struct {
	value *ServiceStorageStorageInner
	isSet bool
}

func (v NullableServiceStorageStorageInner) Get() *ServiceStorageStorageInner {
	return v.value
}

func (v *NullableServiceStorageStorageInner) Set(val *ServiceStorageStorageInner) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceStorageStorageInner) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceStorageStorageInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceStorageStorageInner(val *ServiceStorageStorageInner) *NullableServiceStorageStorageInner {
	return &NullableServiceStorageStorageInner{value: val, isSet: true}
}

func (v NullableServiceStorageStorageInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceStorageStorageInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}