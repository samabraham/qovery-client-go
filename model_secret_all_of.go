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

// SecretAllOf struct for SecretAllOf
type SecretAllOf struct {
	// key is case sensitive
	Key              *string                      `json:"key,omitempty"`
	OverriddenSecret *OverriddenSecret            `json:"overridden_secret,omitempty"`
	AliasedSecret    *AliasedSecret               `json:"aliased_secret,omitempty"`
	Scope            EnvironmentVariableScopeEnum `json:"scope"`
}

// NewSecretAllOf instantiates a new SecretAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecretAllOf(scope EnvironmentVariableScopeEnum) *SecretAllOf {
	this := SecretAllOf{}
	this.Scope = scope
	return &this
}

// NewSecretAllOfWithDefaults instantiates a new SecretAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecretAllOfWithDefaults() *SecretAllOf {
	this := SecretAllOf{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *SecretAllOf) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretAllOf) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *SecretAllOf) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *SecretAllOf) SetKey(v string) {
	o.Key = &v
}

// GetOverriddenSecret returns the OverriddenSecret field value if set, zero value otherwise.
func (o *SecretAllOf) GetOverriddenSecret() OverriddenSecret {
	if o == nil || o.OverriddenSecret == nil {
		var ret OverriddenSecret
		return ret
	}
	return *o.OverriddenSecret
}

// GetOverriddenSecretOk returns a tuple with the OverriddenSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretAllOf) GetOverriddenSecretOk() (*OverriddenSecret, bool) {
	if o == nil || o.OverriddenSecret == nil {
		return nil, false
	}
	return o.OverriddenSecret, true
}

// HasOverriddenSecret returns a boolean if a field has been set.
func (o *SecretAllOf) HasOverriddenSecret() bool {
	if o != nil && o.OverriddenSecret != nil {
		return true
	}

	return false
}

// SetOverriddenSecret gets a reference to the given OverriddenSecret and assigns it to the OverriddenSecret field.
func (o *SecretAllOf) SetOverriddenSecret(v OverriddenSecret) {
	o.OverriddenSecret = &v
}

// GetAliasedSecret returns the AliasedSecret field value if set, zero value otherwise.
func (o *SecretAllOf) GetAliasedSecret() AliasedSecret {
	if o == nil || o.AliasedSecret == nil {
		var ret AliasedSecret
		return ret
	}
	return *o.AliasedSecret
}

// GetAliasedSecretOk returns a tuple with the AliasedSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretAllOf) GetAliasedSecretOk() (*AliasedSecret, bool) {
	if o == nil || o.AliasedSecret == nil {
		return nil, false
	}
	return o.AliasedSecret, true
}

// HasAliasedSecret returns a boolean if a field has been set.
func (o *SecretAllOf) HasAliasedSecret() bool {
	if o != nil && o.AliasedSecret != nil {
		return true
	}

	return false
}

// SetAliasedSecret gets a reference to the given AliasedSecret and assigns it to the AliasedSecret field.
func (o *SecretAllOf) SetAliasedSecret(v AliasedSecret) {
	o.AliasedSecret = &v
}

// GetScope returns the Scope field value
func (o *SecretAllOf) GetScope() EnvironmentVariableScopeEnum {
	if o == nil {
		var ret EnvironmentVariableScopeEnum
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *SecretAllOf) GetScopeOk() (*EnvironmentVariableScopeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *SecretAllOf) SetScope(v EnvironmentVariableScopeEnum) {
	o.Scope = v
}

func (o SecretAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.OverriddenSecret != nil {
		toSerialize["overridden_secret"] = o.OverriddenSecret
	}
	if o.AliasedSecret != nil {
		toSerialize["aliased_secret"] = o.AliasedSecret
	}
	if true {
		toSerialize["scope"] = o.Scope
	}
	return json.Marshal(toSerialize)
}

type NullableSecretAllOf struct {
	value *SecretAllOf
	isSet bool
}

func (v NullableSecretAllOf) Get() *SecretAllOf {
	return v.value
}

func (v *NullableSecretAllOf) Set(val *SecretAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSecretAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSecretAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecretAllOf(val *SecretAllOf) *NullableSecretAllOf {
	return &NullableSecretAllOf{value: val, isSet: true}
}

func (v NullableSecretAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecretAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}