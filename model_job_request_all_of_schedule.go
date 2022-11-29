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

// JobRequestAllOfSchedule struct for JobRequestAllOfSchedule
type JobRequestAllOfSchedule struct {
	Event *JobScheduleEvent `json:"event,omitempty"`
	// Can only be set if the event is CRON. Represent the cron format for the job schedule without seconds. For example: `* * * * *` represent the cron to launch the job every minute. See https://crontab.guru/ to WISIWIG interface. Timezone is UTC
	ScheduledAt NullableString `json:"scheduled_at,omitempty"`
}

// NewJobRequestAllOfSchedule instantiates a new JobRequestAllOfSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobRequestAllOfSchedule() *JobRequestAllOfSchedule {
	this := JobRequestAllOfSchedule{}
	return &this
}

// NewJobRequestAllOfScheduleWithDefaults instantiates a new JobRequestAllOfSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobRequestAllOfScheduleWithDefaults() *JobRequestAllOfSchedule {
	this := JobRequestAllOfSchedule{}
	return &this
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *JobRequestAllOfSchedule) GetEvent() JobScheduleEvent {
	if o == nil || o.Event == nil {
		var ret JobScheduleEvent
		return ret
	}
	return *o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobRequestAllOfSchedule) GetEventOk() (*JobScheduleEvent, bool) {
	if o == nil || o.Event == nil {
		return nil, false
	}
	return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *JobRequestAllOfSchedule) HasEvent() bool {
	if o != nil && o.Event != nil {
		return true
	}

	return false
}

// SetEvent gets a reference to the given JobScheduleEvent and assigns it to the Event field.
func (o *JobRequestAllOfSchedule) SetEvent(v JobScheduleEvent) {
	o.Event = &v
}

// GetScheduledAt returns the ScheduledAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JobRequestAllOfSchedule) GetScheduledAt() string {
	if o == nil || o.ScheduledAt.Get() == nil {
		var ret string
		return ret
	}
	return *o.ScheduledAt.Get()
}

// GetScheduledAtOk returns a tuple with the ScheduledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JobRequestAllOfSchedule) GetScheduledAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScheduledAt.Get(), o.ScheduledAt.IsSet()
}

// HasScheduledAt returns a boolean if a field has been set.
func (o *JobRequestAllOfSchedule) HasScheduledAt() bool {
	if o != nil && o.ScheduledAt.IsSet() {
		return true
	}

	return false
}

// SetScheduledAt gets a reference to the given NullableString and assigns it to the ScheduledAt field.
func (o *JobRequestAllOfSchedule) SetScheduledAt(v string) {
	o.ScheduledAt.Set(&v)
}

// SetScheduledAtNil sets the value for ScheduledAt to be an explicit nil
func (o *JobRequestAllOfSchedule) SetScheduledAtNil() {
	o.ScheduledAt.Set(nil)
}

// UnsetScheduledAt ensures that no value is present for ScheduledAt, not even an explicit nil
func (o *JobRequestAllOfSchedule) UnsetScheduledAt() {
	o.ScheduledAt.Unset()
}

func (o JobRequestAllOfSchedule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Event != nil {
		toSerialize["event"] = o.Event
	}
	if o.ScheduledAt.IsSet() {
		toSerialize["scheduled_at"] = o.ScheduledAt.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableJobRequestAllOfSchedule struct {
	value *JobRequestAllOfSchedule
	isSet bool
}

func (v NullableJobRequestAllOfSchedule) Get() *JobRequestAllOfSchedule {
	return v.value
}

func (v *NullableJobRequestAllOfSchedule) Set(val *JobRequestAllOfSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableJobRequestAllOfSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableJobRequestAllOfSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobRequestAllOfSchedule(val *JobRequestAllOfSchedule) *NullableJobRequestAllOfSchedule {
	return &NullableJobRequestAllOfSchedule{value: val, isSet: true}
}

func (v NullableJobRequestAllOfSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobRequestAllOfSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
