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

// RewardClaimResponse struct for RewardClaimResponse
type RewardClaimResponse struct {
	Type *string `json:"type,omitempty"`
	Code *string `json:"code,omitempty"`
}

// NewRewardClaimResponse instantiates a new RewardClaimResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRewardClaimResponse() *RewardClaimResponse {
	this := RewardClaimResponse{}
	return &this
}

// NewRewardClaimResponseWithDefaults instantiates a new RewardClaimResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRewardClaimResponseWithDefaults() *RewardClaimResponse {
	this := RewardClaimResponse{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RewardClaimResponse) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RewardClaimResponse) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RewardClaimResponse) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RewardClaimResponse) SetType(v string) {
	o.Type = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *RewardClaimResponse) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RewardClaimResponse) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *RewardClaimResponse) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *RewardClaimResponse) SetCode(v string) {
	o.Code = &v
}

func (o RewardClaimResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	return json.Marshal(toSerialize)
}

type NullableRewardClaimResponse struct {
	value *RewardClaimResponse
	isSet bool
}

func (v NullableRewardClaimResponse) Get() *RewardClaimResponse {
	return v.value
}

func (v *NullableRewardClaimResponse) Set(val *RewardClaimResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRewardClaimResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRewardClaimResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRewardClaimResponse(val *RewardClaimResponse) *NullableRewardClaimResponse {
	return &NullableRewardClaimResponse{value: val, isSet: true}
}

func (v NullableRewardClaimResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRewardClaimResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
