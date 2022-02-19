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
)

// DoCredentialsRequest struct for DoCredentialsRequest
type DoCredentialsRequest struct {
	Name            string  `json:"name"`
	Token           *string `json:"token,omitempty"`
	SpacesAccessId  *string `json:"spaces_access_id,omitempty"`
	SpacesSecretKey *string `json:"spaces_secret_key,omitempty"`
}

// NewDoCredentialsRequest instantiates a new DoCredentialsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDoCredentialsRequest(name string) *DoCredentialsRequest {
	this := DoCredentialsRequest{}
	this.Name = name
	return &this
}

// NewDoCredentialsRequestWithDefaults instantiates a new DoCredentialsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDoCredentialsRequestWithDefaults() *DoCredentialsRequest {
	this := DoCredentialsRequest{}
	return &this
}

// GetName returns the Name field value
func (o *DoCredentialsRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DoCredentialsRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DoCredentialsRequest) SetName(v string) {
	o.Name = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *DoCredentialsRequest) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DoCredentialsRequest) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *DoCredentialsRequest) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *DoCredentialsRequest) SetToken(v string) {
	o.Token = &v
}

// GetSpacesAccessId returns the SpacesAccessId field value if set, zero value otherwise.
func (o *DoCredentialsRequest) GetSpacesAccessId() string {
	if o == nil || o.SpacesAccessId == nil {
		var ret string
		return ret
	}
	return *o.SpacesAccessId
}

// GetSpacesAccessIdOk returns a tuple with the SpacesAccessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DoCredentialsRequest) GetSpacesAccessIdOk() (*string, bool) {
	if o == nil || o.SpacesAccessId == nil {
		return nil, false
	}
	return o.SpacesAccessId, true
}

// HasSpacesAccessId returns a boolean if a field has been set.
func (o *DoCredentialsRequest) HasSpacesAccessId() bool {
	if o != nil && o.SpacesAccessId != nil {
		return true
	}

	return false
}

// SetSpacesAccessId gets a reference to the given string and assigns it to the SpacesAccessId field.
func (o *DoCredentialsRequest) SetSpacesAccessId(v string) {
	o.SpacesAccessId = &v
}

// GetSpacesSecretKey returns the SpacesSecretKey field value if set, zero value otherwise.
func (o *DoCredentialsRequest) GetSpacesSecretKey() string {
	if o == nil || o.SpacesSecretKey == nil {
		var ret string
		return ret
	}
	return *o.SpacesSecretKey
}

// GetSpacesSecretKeyOk returns a tuple with the SpacesSecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DoCredentialsRequest) GetSpacesSecretKeyOk() (*string, bool) {
	if o == nil || o.SpacesSecretKey == nil {
		return nil, false
	}
	return o.SpacesSecretKey, true
}

// HasSpacesSecretKey returns a boolean if a field has been set.
func (o *DoCredentialsRequest) HasSpacesSecretKey() bool {
	if o != nil && o.SpacesSecretKey != nil {
		return true
	}

	return false
}

// SetSpacesSecretKey gets a reference to the given string and assigns it to the SpacesSecretKey field.
func (o *DoCredentialsRequest) SetSpacesSecretKey(v string) {
	o.SpacesSecretKey = &v
}

func (o DoCredentialsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.SpacesAccessId != nil {
		toSerialize["spaces_access_id"] = o.SpacesAccessId
	}
	if o.SpacesSecretKey != nil {
		toSerialize["spaces_secret_key"] = o.SpacesSecretKey
	}
	return json.Marshal(toSerialize)
}

type NullableDoCredentialsRequest struct {
	value *DoCredentialsRequest
	isSet bool
}

func (v NullableDoCredentialsRequest) Get() *DoCredentialsRequest {
	return v.value
}

func (v *NullableDoCredentialsRequest) Set(val *DoCredentialsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDoCredentialsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDoCredentialsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDoCredentialsRequest(val *DoCredentialsRequest) *NullableDoCredentialsRequest {
	return &NullableDoCredentialsRequest{value: val, isSet: true}
}

func (v NullableDoCredentialsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDoCredentialsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
