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

// VariableImportResponseSuccessfulImportedVariables struct for VariableImportResponseSuccessfulImportedVariables
type VariableImportResponseSuccessfulImportedVariables struct {
	Name string `json:"name"`
	// Optional if the variable is secret
	Value    *string `json:"value,omitempty"`
	Scope    string  `json:"scope"`
	IsSecret bool    `json:"is_secret"`
}

// NewVariableImportResponseSuccessfulImportedVariables instantiates a new VariableImportResponseSuccessfulImportedVariables object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVariableImportResponseSuccessfulImportedVariables(name string, scope string, isSecret bool) *VariableImportResponseSuccessfulImportedVariables {
	this := VariableImportResponseSuccessfulImportedVariables{}
	this.Name = name
	this.Scope = scope
	this.IsSecret = isSecret
	return &this
}

// NewVariableImportResponseSuccessfulImportedVariablesWithDefaults instantiates a new VariableImportResponseSuccessfulImportedVariables object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVariableImportResponseSuccessfulImportedVariablesWithDefaults() *VariableImportResponseSuccessfulImportedVariables {
	this := VariableImportResponseSuccessfulImportedVariables{}
	return &this
}

// GetName returns the Name field value
func (o *VariableImportResponseSuccessfulImportedVariables) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VariableImportResponseSuccessfulImportedVariables) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VariableImportResponseSuccessfulImportedVariables) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *VariableImportResponseSuccessfulImportedVariables) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableImportResponseSuccessfulImportedVariables) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *VariableImportResponseSuccessfulImportedVariables) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *VariableImportResponseSuccessfulImportedVariables) SetValue(v string) {
	o.Value = &v
}

// GetScope returns the Scope field value
func (o *VariableImportResponseSuccessfulImportedVariables) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *VariableImportResponseSuccessfulImportedVariables) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *VariableImportResponseSuccessfulImportedVariables) SetScope(v string) {
	o.Scope = v
}

// GetIsSecret returns the IsSecret field value
func (o *VariableImportResponseSuccessfulImportedVariables) GetIsSecret() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSecret
}

// GetIsSecretOk returns a tuple with the IsSecret field value
// and a boolean to check if the value has been set.
func (o *VariableImportResponseSuccessfulImportedVariables) GetIsSecretOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSecret, true
}

// SetIsSecret sets field value
func (o *VariableImportResponseSuccessfulImportedVariables) SetIsSecret(v bool) {
	o.IsSecret = v
}

func (o VariableImportResponseSuccessfulImportedVariables) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if true {
		toSerialize["scope"] = o.Scope
	}
	if true {
		toSerialize["is_secret"] = o.IsSecret
	}
	return json.Marshal(toSerialize)
}

type NullableVariableImportResponseSuccessfulImportedVariables struct {
	value *VariableImportResponseSuccessfulImportedVariables
	isSet bool
}

func (v NullableVariableImportResponseSuccessfulImportedVariables) Get() *VariableImportResponseSuccessfulImportedVariables {
	return v.value
}

func (v *NullableVariableImportResponseSuccessfulImportedVariables) Set(val *VariableImportResponseSuccessfulImportedVariables) {
	v.value = val
	v.isSet = true
}

func (v NullableVariableImportResponseSuccessfulImportedVariables) IsSet() bool {
	return v.isSet
}

func (v *NullableVariableImportResponseSuccessfulImportedVariables) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariableImportResponseSuccessfulImportedVariables(val *VariableImportResponseSuccessfulImportedVariables) *NullableVariableImportResponseSuccessfulImportedVariables {
	return &NullableVariableImportResponseSuccessfulImportedVariables{value: val, isSet: true}
}

func (v NullableVariableImportResponseSuccessfulImportedVariables) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVariableImportResponseSuccessfulImportedVariables) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
