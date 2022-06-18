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

// CreateOrganizationWebhook201ResponseAllOf struct for CreateOrganizationWebhook201ResponseAllOf
type CreateOrganizationWebhook201ResponseAllOf struct {
	Kind *Object `json:"kind,omitempty"`
	// Set the public HTTP or HTTPS endpoint that will receive the specified events. The target URL must starts with `http://` or `https://`
	TargetUrl       *string `json:"target_url,omitempty"`
	TargetSecretSet *bool   `json:"target_secret_set,omitempty"`
	Description     *string `json:"description,omitempty"`
	// Turn on or off your endpoint.
	Enabled                *bool                 `json:"enabled,omitempty"`
	Events                 []Object              `json:"events,omitempty"`
	ProjectIdFilter        []string              `json:"project_id_filter,omitempty"`
	EnvironmentTypesFilter []EnvironmentModeEnum `json:"environment_types_filter,omitempty"`
}

// NewCreateOrganizationWebhook201ResponseAllOf instantiates a new CreateOrganizationWebhook201ResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationWebhook201ResponseAllOf() *CreateOrganizationWebhook201ResponseAllOf {
	this := CreateOrganizationWebhook201ResponseAllOf{}
	return &this
}

// NewCreateOrganizationWebhook201ResponseAllOfWithDefaults instantiates a new CreateOrganizationWebhook201ResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationWebhook201ResponseAllOfWithDefaults() *CreateOrganizationWebhook201ResponseAllOf {
	this := CreateOrganizationWebhook201ResponseAllOf{}
	return &this
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *CreateOrganizationWebhook201ResponseAllOf) GetKind() Object {
	if o == nil || o.Kind == nil {
		var ret Object
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationWebhook201ResponseAllOf) GetKindOk() (*Object, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *CreateOrganizationWebhook201ResponseAllOf) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given Object and assigns it to the Kind field.
func (o *CreateOrganizationWebhook201ResponseAllOf) SetKind(v Object) {
	o.Kind = &v
}

// GetTargetUrl returns the TargetUrl field value if set, zero value otherwise.
func (o *CreateOrganizationWebhook201ResponseAllOf) GetTargetUrl() string {
	if o == nil || o.TargetUrl == nil {
		var ret string
		return ret
	}
	return *o.TargetUrl
}

// GetTargetUrlOk returns a tuple with the TargetUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationWebhook201ResponseAllOf) GetTargetUrlOk() (*string, bool) {
	if o == nil || o.TargetUrl == nil {
		return nil, false
	}
	return o.TargetUrl, true
}

// HasTargetUrl returns a boolean if a field has been set.
func (o *CreateOrganizationWebhook201ResponseAllOf) HasTargetUrl() bool {
	if o != nil && o.TargetUrl != nil {
		return true
	}

	return false
}

// SetTargetUrl gets a reference to the given string and assigns it to the TargetUrl field.
func (o *CreateOrganizationWebhook201ResponseAllOf) SetTargetUrl(v string) {
	o.TargetUrl = &v
}

// GetTargetSecretSet returns the TargetSecretSet field value if set, zero value otherwise.
func (o *CreateOrganizationWebhook201ResponseAllOf) GetTargetSecretSet() bool {
	if o == nil || o.TargetSecretSet == nil {
		var ret bool
		return ret
	}
	return *o.TargetSecretSet
}

// GetTargetSecretSetOk returns a tuple with the TargetSecretSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationWebhook201ResponseAllOf) GetTargetSecretSetOk() (*bool, bool) {
	if o == nil || o.TargetSecretSet == nil {
		return nil, false
	}
	return o.TargetSecretSet, true
}

// HasTargetSecretSet returns a boolean if a field has been set.
func (o *CreateOrganizationWebhook201ResponseAllOf) HasTargetSecretSet() bool {
	if o != nil && o.TargetSecretSet != nil {
		return true
	}

	return false
}

// SetTargetSecretSet gets a reference to the given bool and assigns it to the TargetSecretSet field.
func (o *CreateOrganizationWebhook201ResponseAllOf) SetTargetSecretSet(v bool) {
	o.TargetSecretSet = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateOrganizationWebhook201ResponseAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationWebhook201ResponseAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateOrganizationWebhook201ResponseAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateOrganizationWebhook201ResponseAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CreateOrganizationWebhook201ResponseAllOf) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationWebhook201ResponseAllOf) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CreateOrganizationWebhook201ResponseAllOf) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CreateOrganizationWebhook201ResponseAllOf) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *CreateOrganizationWebhook201ResponseAllOf) GetEvents() []Object {
	if o == nil || o.Events == nil {
		var ret []Object
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationWebhook201ResponseAllOf) GetEventsOk() ([]Object, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *CreateOrganizationWebhook201ResponseAllOf) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []Object and assigns it to the Events field.
func (o *CreateOrganizationWebhook201ResponseAllOf) SetEvents(v []Object) {
	o.Events = v
}

// GetProjectIdFilter returns the ProjectIdFilter field value if set, zero value otherwise.
func (o *CreateOrganizationWebhook201ResponseAllOf) GetProjectIdFilter() []string {
	if o == nil || o.ProjectIdFilter == nil {
		var ret []string
		return ret
	}
	return o.ProjectIdFilter
}

// GetProjectIdFilterOk returns a tuple with the ProjectIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationWebhook201ResponseAllOf) GetProjectIdFilterOk() ([]string, bool) {
	if o == nil || o.ProjectIdFilter == nil {
		return nil, false
	}
	return o.ProjectIdFilter, true
}

// HasProjectIdFilter returns a boolean if a field has been set.
func (o *CreateOrganizationWebhook201ResponseAllOf) HasProjectIdFilter() bool {
	if o != nil && o.ProjectIdFilter != nil {
		return true
	}

	return false
}

// SetProjectIdFilter gets a reference to the given []string and assigns it to the ProjectIdFilter field.
func (o *CreateOrganizationWebhook201ResponseAllOf) SetProjectIdFilter(v []string) {
	o.ProjectIdFilter = v
}

// GetEnvironmentTypesFilter returns the EnvironmentTypesFilter field value if set, zero value otherwise.
func (o *CreateOrganizationWebhook201ResponseAllOf) GetEnvironmentTypesFilter() []EnvironmentModeEnum {
	if o == nil || o.EnvironmentTypesFilter == nil {
		var ret []EnvironmentModeEnum
		return ret
	}
	return o.EnvironmentTypesFilter
}

// GetEnvironmentTypesFilterOk returns a tuple with the EnvironmentTypesFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationWebhook201ResponseAllOf) GetEnvironmentTypesFilterOk() ([]EnvironmentModeEnum, bool) {
	if o == nil || o.EnvironmentTypesFilter == nil {
		return nil, false
	}
	return o.EnvironmentTypesFilter, true
}

// HasEnvironmentTypesFilter returns a boolean if a field has been set.
func (o *CreateOrganizationWebhook201ResponseAllOf) HasEnvironmentTypesFilter() bool {
	if o != nil && o.EnvironmentTypesFilter != nil {
		return true
	}

	return false
}

// SetEnvironmentTypesFilter gets a reference to the given []EnvironmentModeEnum and assigns it to the EnvironmentTypesFilter field.
func (o *CreateOrganizationWebhook201ResponseAllOf) SetEnvironmentTypesFilter(v []EnvironmentModeEnum) {
	o.EnvironmentTypesFilter = v
}

func (o CreateOrganizationWebhook201ResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.TargetUrl != nil {
		toSerialize["target_url"] = o.TargetUrl
	}
	if o.TargetSecretSet != nil {
		toSerialize["target_secret_set"] = o.TargetSecretSet
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	if o.ProjectIdFilter != nil {
		toSerialize["project_id_filter"] = o.ProjectIdFilter
	}
	if o.EnvironmentTypesFilter != nil {
		toSerialize["environment_types_filter"] = o.EnvironmentTypesFilter
	}
	return json.Marshal(toSerialize)
}

type NullableCreateOrganizationWebhook201ResponseAllOf struct {
	value *CreateOrganizationWebhook201ResponseAllOf
	isSet bool
}

func (v NullableCreateOrganizationWebhook201ResponseAllOf) Get() *CreateOrganizationWebhook201ResponseAllOf {
	return v.value
}

func (v *NullableCreateOrganizationWebhook201ResponseAllOf) Set(val *CreateOrganizationWebhook201ResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationWebhook201ResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationWebhook201ResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationWebhook201ResponseAllOf(val *CreateOrganizationWebhook201ResponseAllOf) *NullableCreateOrganizationWebhook201ResponseAllOf {
	return &NullableCreateOrganizationWebhook201ResponseAllOf{value: val, isSet: true}
}

func (v NullableCreateOrganizationWebhook201ResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationWebhook201ResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
