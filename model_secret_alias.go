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

// SecretAlias struct for SecretAlias
type SecretAlias struct {
	Id    string               `json:"id"`
	Key   string               `json:"key"`
	Scope APIVariableScopeEnum `json:"scope"`
}

// NewSecretAlias instantiates a new SecretAlias object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecretAlias(id string, key string, scope APIVariableScopeEnum) *SecretAlias {
	this := SecretAlias{}
	this.Id = id
	this.Key = key
	this.Scope = scope
	return &this
}

// NewSecretAliasWithDefaults instantiates a new SecretAlias object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecretAliasWithDefaults() *SecretAlias {
	this := SecretAlias{}
	return &this
}

// GetId returns the Id field value
func (o *SecretAlias) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SecretAlias) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SecretAlias) SetId(v string) {
	o.Id = v
}

// GetKey returns the Key field value
func (o *SecretAlias) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *SecretAlias) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *SecretAlias) SetKey(v string) {
	o.Key = v
}

// GetScope returns the Scope field value
func (o *SecretAlias) GetScope() APIVariableScopeEnum {
	if o == nil {
		var ret APIVariableScopeEnum
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *SecretAlias) GetScopeOk() (*APIVariableScopeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *SecretAlias) SetScope(v APIVariableScopeEnum) {
	o.Scope = v
}

func (o SecretAlias) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["scope"] = o.Scope
	}
	return json.Marshal(toSerialize)
}

type NullableSecretAlias struct {
	value *SecretAlias
	isSet bool
}

func (v NullableSecretAlias) Get() *SecretAlias {
	return v.value
}

func (v *NullableSecretAlias) Set(val *SecretAlias) {
	v.value = val
	v.isSet = true
}

func (v NullableSecretAlias) IsSet() bool {
	return v.isSet
}

func (v *NullableSecretAlias) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecretAlias(val *SecretAlias) *NullableSecretAlias {
	return &NullableSecretAlias{value: val, isSet: true}
}

func (v NullableSecretAlias) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecretAlias) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
