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

// OrganizationCustomRoleUpdateRequest struct for OrganizationCustomRoleUpdateRequest
type OrganizationCustomRoleUpdateRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	// Should contain an entry for every existing cluster
	ClusterPermissions []OrganizationCustomRoleUpdateRequestClusterPermissionsInner `json:"cluster_permissions"`
	// Should contain an entry for every existing project
	ProjectPermissions []OrganizationCustomRoleUpdateRequestProjectPermissionsInner `json:"project_permissions"`
}

// NewOrganizationCustomRoleUpdateRequest instantiates a new OrganizationCustomRoleUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationCustomRoleUpdateRequest(name string, clusterPermissions []OrganizationCustomRoleUpdateRequestClusterPermissionsInner, projectPermissions []OrganizationCustomRoleUpdateRequestProjectPermissionsInner) *OrganizationCustomRoleUpdateRequest {
	this := OrganizationCustomRoleUpdateRequest{}
	this.Name = name
	this.ClusterPermissions = clusterPermissions
	this.ProjectPermissions = projectPermissions
	return &this
}

// NewOrganizationCustomRoleUpdateRequestWithDefaults instantiates a new OrganizationCustomRoleUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationCustomRoleUpdateRequestWithDefaults() *OrganizationCustomRoleUpdateRequest {
	this := OrganizationCustomRoleUpdateRequest{}
	return &this
}

// GetName returns the Name field value
func (o *OrganizationCustomRoleUpdateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OrganizationCustomRoleUpdateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OrganizationCustomRoleUpdateRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OrganizationCustomRoleUpdateRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationCustomRoleUpdateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OrganizationCustomRoleUpdateRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OrganizationCustomRoleUpdateRequest) SetDescription(v string) {
	o.Description = &v
}

// GetClusterPermissions returns the ClusterPermissions field value
func (o *OrganizationCustomRoleUpdateRequest) GetClusterPermissions() []OrganizationCustomRoleUpdateRequestClusterPermissionsInner {
	if o == nil {
		var ret []OrganizationCustomRoleUpdateRequestClusterPermissionsInner
		return ret
	}

	return o.ClusterPermissions
}

// GetClusterPermissionsOk returns a tuple with the ClusterPermissions field value
// and a boolean to check if the value has been set.
func (o *OrganizationCustomRoleUpdateRequest) GetClusterPermissionsOk() ([]OrganizationCustomRoleUpdateRequestClusterPermissionsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClusterPermissions, true
}

// SetClusterPermissions sets field value
func (o *OrganizationCustomRoleUpdateRequest) SetClusterPermissions(v []OrganizationCustomRoleUpdateRequestClusterPermissionsInner) {
	o.ClusterPermissions = v
}

// GetProjectPermissions returns the ProjectPermissions field value
func (o *OrganizationCustomRoleUpdateRequest) GetProjectPermissions() []OrganizationCustomRoleUpdateRequestProjectPermissionsInner {
	if o == nil {
		var ret []OrganizationCustomRoleUpdateRequestProjectPermissionsInner
		return ret
	}

	return o.ProjectPermissions
}

// GetProjectPermissionsOk returns a tuple with the ProjectPermissions field value
// and a boolean to check if the value has been set.
func (o *OrganizationCustomRoleUpdateRequest) GetProjectPermissionsOk() ([]OrganizationCustomRoleUpdateRequestProjectPermissionsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectPermissions, true
}

// SetProjectPermissions sets field value
func (o *OrganizationCustomRoleUpdateRequest) SetProjectPermissions(v []OrganizationCustomRoleUpdateRequestProjectPermissionsInner) {
	o.ProjectPermissions = v
}

func (o OrganizationCustomRoleUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["cluster_permissions"] = o.ClusterPermissions
	}
	if true {
		toSerialize["project_permissions"] = o.ProjectPermissions
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationCustomRoleUpdateRequest struct {
	value *OrganizationCustomRoleUpdateRequest
	isSet bool
}

func (v NullableOrganizationCustomRoleUpdateRequest) Get() *OrganizationCustomRoleUpdateRequest {
	return v.value
}

func (v *NullableOrganizationCustomRoleUpdateRequest) Set(val *OrganizationCustomRoleUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationCustomRoleUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationCustomRoleUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationCustomRoleUpdateRequest(val *OrganizationCustomRoleUpdateRequest) *NullableOrganizationCustomRoleUpdateRequest {
	return &NullableOrganizationCustomRoleUpdateRequest{value: val, isSet: true}
}

func (v NullableOrganizationCustomRoleUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationCustomRoleUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
