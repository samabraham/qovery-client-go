/*
Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is stable and still in development.

API version: 1.0.1
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"encoding/json"
	"time"
)

// MemberResponse struct for MemberResponse
type MemberResponse struct {
	Name              *string `json:"name,omitempty"`
	Nickname          *string `json:"nickname,omitempty"`
	Email             string  `json:"email"`
	ProfilePictureUrl *string `json:"profile_picture_url,omitempty"`
	// last time the user was connected
	LastActivityAt *time.Time `json:"last_activity_at,omitempty"`
	Role           *string    `json:"role,omitempty"`
	Id             string     `json:"id"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at,omitempty"`
}

// NewMemberResponse instantiates a new MemberResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemberResponse(email string, id string, createdAt time.Time) *MemberResponse {
	this := MemberResponse{}
	this.Id = id
	this.CreatedAt = createdAt
	return &this
}

// NewMemberResponseWithDefaults instantiates a new MemberResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemberResponseWithDefaults() *MemberResponse {
	this := MemberResponse{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MemberResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MemberResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MemberResponse) SetName(v string) {
	o.Name = &v
}

// GetNickname returns the Nickname field value if set, zero value otherwise.
func (o *MemberResponse) GetNickname() string {
	if o == nil || o.Nickname == nil {
		var ret string
		return ret
	}
	return *o.Nickname
}

// GetNicknameOk returns a tuple with the Nickname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberResponse) GetNicknameOk() (*string, bool) {
	if o == nil || o.Nickname == nil {
		return nil, false
	}
	return o.Nickname, true
}

// HasNickname returns a boolean if a field has been set.
func (o *MemberResponse) HasNickname() bool {
	if o != nil && o.Nickname != nil {
		return true
	}

	return false
}

// SetNickname gets a reference to the given string and assigns it to the Nickname field.
func (o *MemberResponse) SetNickname(v string) {
	o.Nickname = &v
}

// GetEmail returns the Email field value
func (o *MemberResponse) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *MemberResponse) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *MemberResponse) SetEmail(v string) {
	o.Email = v
}

// GetProfilePictureUrl returns the ProfilePictureUrl field value if set, zero value otherwise.
func (o *MemberResponse) GetProfilePictureUrl() string {
	if o == nil || o.ProfilePictureUrl == nil {
		var ret string
		return ret
	}
	return *o.ProfilePictureUrl
}

// GetProfilePictureUrlOk returns a tuple with the ProfilePictureUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberResponse) GetProfilePictureUrlOk() (*string, bool) {
	if o == nil || o.ProfilePictureUrl == nil {
		return nil, false
	}
	return o.ProfilePictureUrl, true
}

// HasProfilePictureUrl returns a boolean if a field has been set.
func (o *MemberResponse) HasProfilePictureUrl() bool {
	if o != nil && o.ProfilePictureUrl != nil {
		return true
	}

	return false
}

// SetProfilePictureUrl gets a reference to the given string and assigns it to the ProfilePictureUrl field.
func (o *MemberResponse) SetProfilePictureUrl(v string) {
	o.ProfilePictureUrl = &v
}

// GetLastActivityAt returns the LastActivityAt field value if set, zero value otherwise.
func (o *MemberResponse) GetLastActivityAt() time.Time {
	if o == nil || o.LastActivityAt == nil {
		var ret time.Time
		return ret
	}
	return *o.LastActivityAt
}

// GetLastActivityAtOk returns a tuple with the LastActivityAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberResponse) GetLastActivityAtOk() (*time.Time, bool) {
	if o == nil || o.LastActivityAt == nil {
		return nil, false
	}
	return o.LastActivityAt, true
}

// HasLastActivityAt returns a boolean if a field has been set.
func (o *MemberResponse) HasLastActivityAt() bool {
	if o != nil && o.LastActivityAt != nil {
		return true
	}

	return false
}

// SetLastActivityAt gets a reference to the given time.Time and assigns it to the LastActivityAt field.
func (o *MemberResponse) SetLastActivityAt(v time.Time) {
	o.LastActivityAt = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *MemberResponse) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberResponse) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *MemberResponse) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *MemberResponse) SetRole(v string) {
	o.Role = &v
}

// GetId returns the Id field value
func (o *MemberResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MemberResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MemberResponse) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *MemberResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *MemberResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *MemberResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *MemberResponse) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *MemberResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *MemberResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o MemberResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Nickname != nil {
		toSerialize["nickname"] = o.Nickname
	}
	if true {
		toSerialize["email"] = o.Email
	}
	if o.ProfilePictureUrl != nil {
		toSerialize["profile_picture_url"] = o.ProfilePictureUrl
	}
	if o.LastActivityAt != nil {
		toSerialize["last_activity_at"] = o.LastActivityAt
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
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
	return json.Marshal(toSerialize)
}

type NullableMemberResponse struct {
	value *MemberResponse
	isSet bool
}

func (v NullableMemberResponse) Get() *MemberResponse {
	return v.value
}

func (v *NullableMemberResponse) Set(val *MemberResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMemberResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMemberResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemberResponse(val *MemberResponse) *NullableMemberResponse {
	return &NullableMemberResponse{value: val, isSet: true}
}

func (v NullableMemberResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemberResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
