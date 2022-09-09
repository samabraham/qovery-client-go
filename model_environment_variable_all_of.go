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

// EnvironmentVariableAllOf struct for EnvironmentVariableAllOf
type EnvironmentVariableAllOf struct {
	OverriddenVariable *EnvironmentVariableOverride `json:"overridden_variable,omitempty"`
	AliasedVariable    *EnvironmentVariableAlias    `json:"aliased_variable,omitempty"`
	Scope              APIVariableScopeEnum         `json:"scope"`
	// present only for `BUILT_IN` variable
	ServiceId *string `json:"service_id,omitempty"`
	// present only for `BUILT_IN` variable
	ServiceName *string                `json:"service_name,omitempty"`
	ServiceType *LinkedServiceTypeEnum `json:"service_type,omitempty"`
}

// NewEnvironmentVariableAllOf instantiates a new EnvironmentVariableAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentVariableAllOf(scope APIVariableScopeEnum) *EnvironmentVariableAllOf {
	this := EnvironmentVariableAllOf{}
	this.Scope = scope
	return &this
}

// NewEnvironmentVariableAllOfWithDefaults instantiates a new EnvironmentVariableAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentVariableAllOfWithDefaults() *EnvironmentVariableAllOf {
	this := EnvironmentVariableAllOf{}
	return &this
}

// GetOverriddenVariable returns the OverriddenVariable field value if set, zero value otherwise.
func (o *EnvironmentVariableAllOf) GetOverriddenVariable() EnvironmentVariableOverride {
	if o == nil || o.OverriddenVariable == nil {
		var ret EnvironmentVariableOverride
		return ret
	}
	return *o.OverriddenVariable
}

// GetOverriddenVariableOk returns a tuple with the OverriddenVariable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentVariableAllOf) GetOverriddenVariableOk() (*EnvironmentVariableOverride, bool) {
	if o == nil || o.OverriddenVariable == nil {
		return nil, false
	}
	return o.OverriddenVariable, true
}

// HasOverriddenVariable returns a boolean if a field has been set.
func (o *EnvironmentVariableAllOf) HasOverriddenVariable() bool {
	if o != nil && o.OverriddenVariable != nil {
		return true
	}

	return false
}

// SetOverriddenVariable gets a reference to the given EnvironmentVariableOverride and assigns it to the OverriddenVariable field.
func (o *EnvironmentVariableAllOf) SetOverriddenVariable(v EnvironmentVariableOverride) {
	o.OverriddenVariable = &v
}

// GetAliasedVariable returns the AliasedVariable field value if set, zero value otherwise.
func (o *EnvironmentVariableAllOf) GetAliasedVariable() EnvironmentVariableAlias {
	if o == nil || o.AliasedVariable == nil {
		var ret EnvironmentVariableAlias
		return ret
	}
	return *o.AliasedVariable
}

// GetAliasedVariableOk returns a tuple with the AliasedVariable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentVariableAllOf) GetAliasedVariableOk() (*EnvironmentVariableAlias, bool) {
	if o == nil || o.AliasedVariable == nil {
		return nil, false
	}
	return o.AliasedVariable, true
}

// HasAliasedVariable returns a boolean if a field has been set.
func (o *EnvironmentVariableAllOf) HasAliasedVariable() bool {
	if o != nil && o.AliasedVariable != nil {
		return true
	}

	return false
}

// SetAliasedVariable gets a reference to the given EnvironmentVariableAlias and assigns it to the AliasedVariable field.
func (o *EnvironmentVariableAllOf) SetAliasedVariable(v EnvironmentVariableAlias) {
	o.AliasedVariable = &v
}

// GetScope returns the Scope field value
func (o *EnvironmentVariableAllOf) GetScope() APIVariableScopeEnum {
	if o == nil {
		var ret APIVariableScopeEnum
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *EnvironmentVariableAllOf) GetScopeOk() (*APIVariableScopeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *EnvironmentVariableAllOf) SetScope(v APIVariableScopeEnum) {
	o.Scope = v
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *EnvironmentVariableAllOf) GetServiceId() string {
	if o == nil || o.ServiceId == nil {
		var ret string
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentVariableAllOf) GetServiceIdOk() (*string, bool) {
	if o == nil || o.ServiceId == nil {
		return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *EnvironmentVariableAllOf) HasServiceId() bool {
	if o != nil && o.ServiceId != nil {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given string and assigns it to the ServiceId field.
func (o *EnvironmentVariableAllOf) SetServiceId(v string) {
	o.ServiceId = &v
}

// GetServiceName returns the ServiceName field value if set, zero value otherwise.
func (o *EnvironmentVariableAllOf) GetServiceName() string {
	if o == nil || o.ServiceName == nil {
		var ret string
		return ret
	}
	return *o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentVariableAllOf) GetServiceNameOk() (*string, bool) {
	if o == nil || o.ServiceName == nil {
		return nil, false
	}
	return o.ServiceName, true
}

// HasServiceName returns a boolean if a field has been set.
func (o *EnvironmentVariableAllOf) HasServiceName() bool {
	if o != nil && o.ServiceName != nil {
		return true
	}

	return false
}

// SetServiceName gets a reference to the given string and assigns it to the ServiceName field.
func (o *EnvironmentVariableAllOf) SetServiceName(v string) {
	o.ServiceName = &v
}

// GetServiceType returns the ServiceType field value if set, zero value otherwise.
func (o *EnvironmentVariableAllOf) GetServiceType() LinkedServiceTypeEnum {
	if o == nil || o.ServiceType == nil {
		var ret LinkedServiceTypeEnum
		return ret
	}
	return *o.ServiceType
}

// GetServiceTypeOk returns a tuple with the ServiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentVariableAllOf) GetServiceTypeOk() (*LinkedServiceTypeEnum, bool) {
	if o == nil || o.ServiceType == nil {
		return nil, false
	}
	return o.ServiceType, true
}

// HasServiceType returns a boolean if a field has been set.
func (o *EnvironmentVariableAllOf) HasServiceType() bool {
	if o != nil && o.ServiceType != nil {
		return true
	}

	return false
}

// SetServiceType gets a reference to the given LinkedServiceTypeEnum and assigns it to the ServiceType field.
func (o *EnvironmentVariableAllOf) SetServiceType(v LinkedServiceTypeEnum) {
	o.ServiceType = &v
}

func (o EnvironmentVariableAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OverriddenVariable != nil {
		toSerialize["overridden_variable"] = o.OverriddenVariable
	}
	if o.AliasedVariable != nil {
		toSerialize["aliased_variable"] = o.AliasedVariable
	}
	if true {
		toSerialize["scope"] = o.Scope
	}
	if o.ServiceId != nil {
		toSerialize["service_id"] = o.ServiceId
	}
	if o.ServiceName != nil {
		toSerialize["service_name"] = o.ServiceName
	}
	if o.ServiceType != nil {
		toSerialize["service_type"] = o.ServiceType
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironmentVariableAllOf struct {
	value *EnvironmentVariableAllOf
	isSet bool
}

func (v NullableEnvironmentVariableAllOf) Get() *EnvironmentVariableAllOf {
	return v.value
}

func (v *NullableEnvironmentVariableAllOf) Set(val *EnvironmentVariableAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentVariableAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentVariableAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentVariableAllOf(val *EnvironmentVariableAllOf) *NullableEnvironmentVariableAllOf {
	return &NullableEnvironmentVariableAllOf{value: val, isSet: true}
}

func (v NullableEnvironmentVariableAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentVariableAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
