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

// CurrentCost struct for CurrentCost
type CurrentCost struct {
	Plan *string `json:"plan,omitempty"`
	// number of days remaining before the end of the trial period
	RemainingTrialDay *int32            `json:"remaining_trial_day,omitempty"`
	RemainingCredits  *RemainingCredits `json:"remaining_credits,omitempty"`
	Cost              *Cost             `json:"cost,omitempty"`
}

// NewCurrentCost instantiates a new CurrentCost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCurrentCost() *CurrentCost {
	this := CurrentCost{}
	return &this
}

// NewCurrentCostWithDefaults instantiates a new CurrentCost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCurrentCostWithDefaults() *CurrentCost {
	this := CurrentCost{}
	return &this
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *CurrentCost) GetPlan() string {
	if o == nil || o.Plan == nil {
		var ret string
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentCost) GetPlanOk() (*string, bool) {
	if o == nil || o.Plan == nil {
		return nil, false
	}
	return o.Plan, true
}

// HasPlan returns a boolean if a field has been set.
func (o *CurrentCost) HasPlan() bool {
	if o != nil && o.Plan != nil {
		return true
	}

	return false
}

// SetPlan gets a reference to the given string and assigns it to the Plan field.
func (o *CurrentCost) SetPlan(v string) {
	o.Plan = &v
}

// GetRemainingTrialDay returns the RemainingTrialDay field value if set, zero value otherwise.
func (o *CurrentCost) GetRemainingTrialDay() int32 {
	if o == nil || o.RemainingTrialDay == nil {
		var ret int32
		return ret
	}
	return *o.RemainingTrialDay
}

// GetRemainingTrialDayOk returns a tuple with the RemainingTrialDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentCost) GetRemainingTrialDayOk() (*int32, bool) {
	if o == nil || o.RemainingTrialDay == nil {
		return nil, false
	}
	return o.RemainingTrialDay, true
}

// HasRemainingTrialDay returns a boolean if a field has been set.
func (o *CurrentCost) HasRemainingTrialDay() bool {
	if o != nil && o.RemainingTrialDay != nil {
		return true
	}

	return false
}

// SetRemainingTrialDay gets a reference to the given int32 and assigns it to the RemainingTrialDay field.
func (o *CurrentCost) SetRemainingTrialDay(v int32) {
	o.RemainingTrialDay = &v
}

// GetRemainingCredits returns the RemainingCredits field value if set, zero value otherwise.
func (o *CurrentCost) GetRemainingCredits() RemainingCredits {
	if o == nil || o.RemainingCredits == nil {
		var ret RemainingCredits
		return ret
	}
	return *o.RemainingCredits
}

// GetRemainingCreditsOk returns a tuple with the RemainingCredits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentCost) GetRemainingCreditsOk() (*RemainingCredits, bool) {
	if o == nil || o.RemainingCredits == nil {
		return nil, false
	}
	return o.RemainingCredits, true
}

// HasRemainingCredits returns a boolean if a field has been set.
func (o *CurrentCost) HasRemainingCredits() bool {
	if o != nil && o.RemainingCredits != nil {
		return true
	}

	return false
}

// SetRemainingCredits gets a reference to the given RemainingCredits and assigns it to the RemainingCredits field.
func (o *CurrentCost) SetRemainingCredits(v RemainingCredits) {
	o.RemainingCredits = &v
}

// GetCost returns the Cost field value if set, zero value otherwise.
func (o *CurrentCost) GetCost() Cost {
	if o == nil || o.Cost == nil {
		var ret Cost
		return ret
	}
	return *o.Cost
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentCost) GetCostOk() (*Cost, bool) {
	if o == nil || o.Cost == nil {
		return nil, false
	}
	return o.Cost, true
}

// HasCost returns a boolean if a field has been set.
func (o *CurrentCost) HasCost() bool {
	if o != nil && o.Cost != nil {
		return true
	}

	return false
}

// SetCost gets a reference to the given Cost and assigns it to the Cost field.
func (o *CurrentCost) SetCost(v Cost) {
	o.Cost = &v
}

func (o CurrentCost) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Plan != nil {
		toSerialize["plan"] = o.Plan
	}
	if o.RemainingTrialDay != nil {
		toSerialize["remaining_trial_day"] = o.RemainingTrialDay
	}
	if o.RemainingCredits != nil {
		toSerialize["remaining_credits"] = o.RemainingCredits
	}
	if o.Cost != nil {
		toSerialize["cost"] = o.Cost
	}
	return json.Marshal(toSerialize)
}

type NullableCurrentCost struct {
	value *CurrentCost
	isSet bool
}

func (v NullableCurrentCost) Get() *CurrentCost {
	return v.value
}

func (v *NullableCurrentCost) Set(val *CurrentCost) {
	v.value = val
	v.isSet = true
}

func (v NullableCurrentCost) IsSet() bool {
	return v.isSet
}

func (v *NullableCurrentCost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCurrentCost(val *CurrentCost) *NullableCurrentCost {
	return &NullableCurrentCost{value: val, isSet: true}
}

func (v NullableCurrentCost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCurrentCost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
