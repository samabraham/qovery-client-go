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
	"time"
)

// DeploymentHistoryDatabase struct for DeploymentHistoryDatabase
type DeploymentHistoryDatabase struct {
	Id        string                  `json:"id"`
	CreatedAt time.Time               `json:"created_at"`
	UpdatedAt *time.Time              `json:"updated_at,omitempty"`
	Name      *string                 `json:"name,omitempty"`
	Status    *GlobalDeploymentStatus `json:"status,omitempty"`
}

// NewDeploymentHistoryDatabase instantiates a new DeploymentHistoryDatabase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentHistoryDatabase(id string, createdAt time.Time) *DeploymentHistoryDatabase {
	this := DeploymentHistoryDatabase{}
	this.Id = id
	this.CreatedAt = createdAt
	return &this
}

// NewDeploymentHistoryDatabaseWithDefaults instantiates a new DeploymentHistoryDatabase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentHistoryDatabaseWithDefaults() *DeploymentHistoryDatabase {
	this := DeploymentHistoryDatabase{}
	return &this
}

// GetId returns the Id field value
func (o *DeploymentHistoryDatabase) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryDatabase) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DeploymentHistoryDatabase) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *DeploymentHistoryDatabase) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryDatabase) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *DeploymentHistoryDatabase) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *DeploymentHistoryDatabase) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryDatabase) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *DeploymentHistoryDatabase) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *DeploymentHistoryDatabase) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DeploymentHistoryDatabase) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryDatabase) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DeploymentHistoryDatabase) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DeploymentHistoryDatabase) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DeploymentHistoryDatabase) GetStatus() GlobalDeploymentStatus {
	if o == nil || o.Status == nil {
		var ret GlobalDeploymentStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryDatabase) GetStatusOk() (*GlobalDeploymentStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DeploymentHistoryDatabase) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given GlobalDeploymentStatus and assigns it to the Status field.
func (o *DeploymentHistoryDatabase) SetStatus(v GlobalDeploymentStatus) {
	o.Status = &v
}

func (o DeploymentHistoryDatabase) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableDeploymentHistoryDatabase struct {
	value *DeploymentHistoryDatabase
	isSet bool
}

func (v NullableDeploymentHistoryDatabase) Get() *DeploymentHistoryDatabase {
	return v.value
}

func (v *NullableDeploymentHistoryDatabase) Set(val *DeploymentHistoryDatabase) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentHistoryDatabase) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentHistoryDatabase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentHistoryDatabase(val *DeploymentHistoryDatabase) *NullableDeploymentHistoryDatabase {
	return &NullableDeploymentHistoryDatabase{value: val, isSet: true}
}

func (v NullableDeploymentHistoryDatabase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentHistoryDatabase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}