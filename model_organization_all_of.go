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

// OrganizationAllOf struct for OrganizationAllOf
type OrganizationAllOf struct {
	// uuid of the user owning the organization
	Owner *string `json:"owner,omitempty"`
}

// NewOrganizationAllOf instantiates a new OrganizationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationAllOf() *OrganizationAllOf {
	this := OrganizationAllOf{}
	return &this
}

// NewOrganizationAllOfWithDefaults instantiates a new OrganizationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationAllOfWithDefaults() *OrganizationAllOf {
	this := OrganizationAllOf{}
	return &this
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *OrganizationAllOf) GetOwner() string {
	if o == nil || o.Owner == nil {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationAllOf) GetOwnerOk() (*string, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *OrganizationAllOf) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *OrganizationAllOf) SetOwner(v string) {
	o.Owner = &v
}

func (o OrganizationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationAllOf struct {
	value *OrganizationAllOf
	isSet bool
}

func (v NullableOrganizationAllOf) Get() *OrganizationAllOf {
	return v.value
}

func (v *NullableOrganizationAllOf) Set(val *OrganizationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationAllOf(val *OrganizationAllOf) *NullableOrganizationAllOf {
	return &NullableOrganizationAllOf{value: val, isSet: true}
}

func (v NullableOrganizationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}