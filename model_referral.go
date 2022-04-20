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

// Referral struct for Referral
type Referral struct {
	TotalInvited   *int32  `json:"total_invited,omitempty"`
	InvitationLink *string `json:"invitation_link,omitempty"`
}

// NewReferral instantiates a new Referral object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReferral() *Referral {
	this := Referral{}
	return &this
}

// NewReferralWithDefaults instantiates a new Referral object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReferralWithDefaults() *Referral {
	this := Referral{}
	return &this
}

// GetTotalInvited returns the TotalInvited field value if set, zero value otherwise.
func (o *Referral) GetTotalInvited() int32 {
	if o == nil || o.TotalInvited == nil {
		var ret int32
		return ret
	}
	return *o.TotalInvited
}

// GetTotalInvitedOk returns a tuple with the TotalInvited field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Referral) GetTotalInvitedOk() (*int32, bool) {
	if o == nil || o.TotalInvited == nil {
		return nil, false
	}
	return o.TotalInvited, true
}

// HasTotalInvited returns a boolean if a field has been set.
func (o *Referral) HasTotalInvited() bool {
	if o != nil && o.TotalInvited != nil {
		return true
	}

	return false
}

// SetTotalInvited gets a reference to the given int32 and assigns it to the TotalInvited field.
func (o *Referral) SetTotalInvited(v int32) {
	o.TotalInvited = &v
}

// GetInvitationLink returns the InvitationLink field value if set, zero value otherwise.
func (o *Referral) GetInvitationLink() string {
	if o == nil || o.InvitationLink == nil {
		var ret string
		return ret
	}
	return *o.InvitationLink
}

// GetInvitationLinkOk returns a tuple with the InvitationLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Referral) GetInvitationLinkOk() (*string, bool) {
	if o == nil || o.InvitationLink == nil {
		return nil, false
	}
	return o.InvitationLink, true
}

// HasInvitationLink returns a boolean if a field has been set.
func (o *Referral) HasInvitationLink() bool {
	if o != nil && o.InvitationLink != nil {
		return true
	}

	return false
}

// SetInvitationLink gets a reference to the given string and assigns it to the InvitationLink field.
func (o *Referral) SetInvitationLink(v string) {
	o.InvitationLink = &v
}

func (o Referral) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TotalInvited != nil {
		toSerialize["total_invited"] = o.TotalInvited
	}
	if o.InvitationLink != nil {
		toSerialize["invitation_link"] = o.InvitationLink
	}
	return json.Marshal(toSerialize)
}

type NullableReferral struct {
	value *Referral
	isSet bool
}

func (v NullableReferral) Get() *Referral {
	return v.value
}

func (v *NullableReferral) Set(val *Referral) {
	v.value = val
	v.isSet = true
}

func (v NullableReferral) IsSet() bool {
	return v.isSet
}

func (v *NullableReferral) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReferral(val *Referral) *NullableReferral {
	return &NullableReferral{value: val, isSet: true}
}

func (v NullableReferral) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReferral) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}