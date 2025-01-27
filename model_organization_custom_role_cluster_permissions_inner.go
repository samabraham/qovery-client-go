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

// OrganizationCustomRoleClusterPermissionsInner struct for OrganizationCustomRoleClusterPermissionsInner
type OrganizationCustomRoleClusterPermissionsInner struct {
	ClusterId   *string                                  `json:"cluster_id,omitempty"`
	ClusterName *string                                  `json:"cluster_name,omitempty"`
	Permission  *OrganizationCustomRoleClusterPermission `json:"permission,omitempty"`
}

// NewOrganizationCustomRoleClusterPermissionsInner instantiates a new OrganizationCustomRoleClusterPermissionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationCustomRoleClusterPermissionsInner() *OrganizationCustomRoleClusterPermissionsInner {
	this := OrganizationCustomRoleClusterPermissionsInner{}
	return &this
}

// NewOrganizationCustomRoleClusterPermissionsInnerWithDefaults instantiates a new OrganizationCustomRoleClusterPermissionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationCustomRoleClusterPermissionsInnerWithDefaults() *OrganizationCustomRoleClusterPermissionsInner {
	this := OrganizationCustomRoleClusterPermissionsInner{}
	return &this
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise.
func (o *OrganizationCustomRoleClusterPermissionsInner) GetClusterId() string {
	if o == nil || o.ClusterId == nil {
		var ret string
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationCustomRoleClusterPermissionsInner) GetClusterIdOk() (*string, bool) {
	if o == nil || o.ClusterId == nil {
		return nil, false
	}
	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *OrganizationCustomRoleClusterPermissionsInner) HasClusterId() bool {
	if o != nil && o.ClusterId != nil {
		return true
	}

	return false
}

// SetClusterId gets a reference to the given string and assigns it to the ClusterId field.
func (o *OrganizationCustomRoleClusterPermissionsInner) SetClusterId(v string) {
	o.ClusterId = &v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *OrganizationCustomRoleClusterPermissionsInner) GetClusterName() string {
	if o == nil || o.ClusterName == nil {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationCustomRoleClusterPermissionsInner) GetClusterNameOk() (*string, bool) {
	if o == nil || o.ClusterName == nil {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *OrganizationCustomRoleClusterPermissionsInner) HasClusterName() bool {
	if o != nil && o.ClusterName != nil {
		return true
	}

	return false
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *OrganizationCustomRoleClusterPermissionsInner) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetPermission returns the Permission field value if set, zero value otherwise.
func (o *OrganizationCustomRoleClusterPermissionsInner) GetPermission() OrganizationCustomRoleClusterPermission {
	if o == nil || o.Permission == nil {
		var ret OrganizationCustomRoleClusterPermission
		return ret
	}
	return *o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationCustomRoleClusterPermissionsInner) GetPermissionOk() (*OrganizationCustomRoleClusterPermission, bool) {
	if o == nil || o.Permission == nil {
		return nil, false
	}
	return o.Permission, true
}

// HasPermission returns a boolean if a field has been set.
func (o *OrganizationCustomRoleClusterPermissionsInner) HasPermission() bool {
	if o != nil && o.Permission != nil {
		return true
	}

	return false
}

// SetPermission gets a reference to the given OrganizationCustomRoleClusterPermission and assigns it to the Permission field.
func (o *OrganizationCustomRoleClusterPermissionsInner) SetPermission(v OrganizationCustomRoleClusterPermission) {
	o.Permission = &v
}

func (o OrganizationCustomRoleClusterPermissionsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClusterId != nil {
		toSerialize["cluster_id"] = o.ClusterId
	}
	if o.ClusterName != nil {
		toSerialize["cluster_name"] = o.ClusterName
	}
	if o.Permission != nil {
		toSerialize["permission"] = o.Permission
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationCustomRoleClusterPermissionsInner struct {
	value *OrganizationCustomRoleClusterPermissionsInner
	isSet bool
}

func (v NullableOrganizationCustomRoleClusterPermissionsInner) Get() *OrganizationCustomRoleClusterPermissionsInner {
	return v.value
}

func (v *NullableOrganizationCustomRoleClusterPermissionsInner) Set(val *OrganizationCustomRoleClusterPermissionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationCustomRoleClusterPermissionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationCustomRoleClusterPermissionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationCustomRoleClusterPermissionsInner(val *OrganizationCustomRoleClusterPermissionsInner) *NullableOrganizationCustomRoleClusterPermissionsInner {
	return &NullableOrganizationCustomRoleClusterPermissionsInner{value: val, isSet: true}
}

func (v NullableOrganizationCustomRoleClusterPermissionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationCustomRoleClusterPermissionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
