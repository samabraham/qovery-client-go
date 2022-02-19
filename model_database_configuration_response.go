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

// DatabaseConfigurationResponse struct for DatabaseConfigurationResponse
type DatabaseConfigurationResponse struct {
	DatabaseType *string               `json:"database_type,omitempty"`
	Version      []DatabaseVersionMode `json:"version,omitempty"`
}

// NewDatabaseConfigurationResponse instantiates a new DatabaseConfigurationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseConfigurationResponse() *DatabaseConfigurationResponse {
	this := DatabaseConfigurationResponse{}
	return &this
}

// NewDatabaseConfigurationResponseWithDefaults instantiates a new DatabaseConfigurationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseConfigurationResponseWithDefaults() *DatabaseConfigurationResponse {
	this := DatabaseConfigurationResponse{}
	return &this
}

// GetDatabaseType returns the DatabaseType field value if set, zero value otherwise.
func (o *DatabaseConfigurationResponse) GetDatabaseType() string {
	if o == nil || o.DatabaseType == nil {
		var ret string
		return ret
	}
	return *o.DatabaseType
}

// GetDatabaseTypeOk returns a tuple with the DatabaseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseConfigurationResponse) GetDatabaseTypeOk() (*string, bool) {
	if o == nil || o.DatabaseType == nil {
		return nil, false
	}
	return o.DatabaseType, true
}

// HasDatabaseType returns a boolean if a field has been set.
func (o *DatabaseConfigurationResponse) HasDatabaseType() bool {
	if o != nil && o.DatabaseType != nil {
		return true
	}

	return false
}

// SetDatabaseType gets a reference to the given string and assigns it to the DatabaseType field.
func (o *DatabaseConfigurationResponse) SetDatabaseType(v string) {
	o.DatabaseType = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *DatabaseConfigurationResponse) GetVersion() []DatabaseVersionMode {
	if o == nil || o.Version == nil {
		var ret []DatabaseVersionMode
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseConfigurationResponse) GetVersionOk() ([]DatabaseVersionMode, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *DatabaseConfigurationResponse) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given []DatabaseVersionMode and assigns it to the Version field.
func (o *DatabaseConfigurationResponse) SetVersion(v []DatabaseVersionMode) {
	o.Version = v
}

func (o DatabaseConfigurationResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DatabaseType != nil {
		toSerialize["database_type"] = o.DatabaseType
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableDatabaseConfigurationResponse struct {
	value *DatabaseConfigurationResponse
	isSet bool
}

func (v NullableDatabaseConfigurationResponse) Get() *DatabaseConfigurationResponse {
	return v.value
}

func (v *NullableDatabaseConfigurationResponse) Set(val *DatabaseConfigurationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseConfigurationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseConfigurationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseConfigurationResponse(val *DatabaseConfigurationResponse) *NullableDatabaseConfigurationResponse {
	return &NullableDatabaseConfigurationResponse{value: val, isSet: true}
}

func (v NullableDatabaseConfigurationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseConfigurationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
