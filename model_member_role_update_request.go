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

// MemberRoleUpdateRequest struct for MemberRoleUpdateRequest
type MemberRoleUpdateRequest struct {
	// specify the git provider user id
	UserId *string `json:"user_id,omitempty"`
	// used to specify an organization custom role, otherwise `null`
	CustomRoleId    *string            `json:"custom_role_id,omitempty"`
	DefaultRoleName *DefaultMemberRole `json:"default_role_name,omitempty"`
}

// NewMemberRoleUpdateRequest instantiates a new MemberRoleUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemberRoleUpdateRequest() *MemberRoleUpdateRequest {
	this := MemberRoleUpdateRequest{}
	return &this
}

// NewMemberRoleUpdateRequestWithDefaults instantiates a new MemberRoleUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemberRoleUpdateRequestWithDefaults() *MemberRoleUpdateRequest {
	this := MemberRoleUpdateRequest{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *MemberRoleUpdateRequest) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberRoleUpdateRequest) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *MemberRoleUpdateRequest) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *MemberRoleUpdateRequest) SetUserId(v string) {
	o.UserId = &v
}

// GetCustomRoleId returns the CustomRoleId field value if set, zero value otherwise.
func (o *MemberRoleUpdateRequest) GetCustomRoleId() string {
	if o == nil || o.CustomRoleId == nil {
		var ret string
		return ret
	}
	return *o.CustomRoleId
}

// GetCustomRoleIdOk returns a tuple with the CustomRoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberRoleUpdateRequest) GetCustomRoleIdOk() (*string, bool) {
	if o == nil || o.CustomRoleId == nil {
		return nil, false
	}
	return o.CustomRoleId, true
}

// HasCustomRoleId returns a boolean if a field has been set.
func (o *MemberRoleUpdateRequest) HasCustomRoleId() bool {
	if o != nil && o.CustomRoleId != nil {
		return true
	}

	return false
}

// SetCustomRoleId gets a reference to the given string and assigns it to the CustomRoleId field.
func (o *MemberRoleUpdateRequest) SetCustomRoleId(v string) {
	o.CustomRoleId = &v
}

// GetDefaultRoleName returns the DefaultRoleName field value if set, zero value otherwise.
func (o *MemberRoleUpdateRequest) GetDefaultRoleName() DefaultMemberRole {
	if o == nil || o.DefaultRoleName == nil {
		var ret DefaultMemberRole
		return ret
	}
	return *o.DefaultRoleName
}

// GetDefaultRoleNameOk returns a tuple with the DefaultRoleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberRoleUpdateRequest) GetDefaultRoleNameOk() (*DefaultMemberRole, bool) {
	if o == nil || o.DefaultRoleName == nil {
		return nil, false
	}
	return o.DefaultRoleName, true
}

// HasDefaultRoleName returns a boolean if a field has been set.
func (o *MemberRoleUpdateRequest) HasDefaultRoleName() bool {
	if o != nil && o.DefaultRoleName != nil {
		return true
	}

	return false
}

// SetDefaultRoleName gets a reference to the given DefaultMemberRole and assigns it to the DefaultRoleName field.
func (o *MemberRoleUpdateRequest) SetDefaultRoleName(v DefaultMemberRole) {
	o.DefaultRoleName = &v
}

func (o MemberRoleUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UserId != nil {
		toSerialize["user_id"] = o.UserId
	}
	if o.CustomRoleId != nil {
		toSerialize["custom_role_id"] = o.CustomRoleId
	}
	if o.DefaultRoleName != nil {
		toSerialize["default_role_name"] = o.DefaultRoleName
	}
	return json.Marshal(toSerialize)
}

type NullableMemberRoleUpdateRequest struct {
	value *MemberRoleUpdateRequest
	isSet bool
}

func (v NullableMemberRoleUpdateRequest) Get() *MemberRoleUpdateRequest {
	return v.value
}

func (v *NullableMemberRoleUpdateRequest) Set(val *MemberRoleUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMemberRoleUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMemberRoleUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemberRoleUpdateRequest(val *MemberRoleUpdateRequest) *NullableMemberRoleUpdateRequest {
	return &NullableMemberRoleUpdateRequest{value: val, isSet: true}
}

func (v NullableMemberRoleUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemberRoleUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}