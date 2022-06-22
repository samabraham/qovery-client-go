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

// OrganizationWebhookCreateResponseAllOf struct for OrganizationWebhookCreateResponseAllOf
type OrganizationWebhookCreateResponseAllOf struct {
	Kind *Kind `json:"kind,omitempty"`
	// Set the public HTTP or HTTPS endpoint that will receive the specified events. The target URL must starts with `http://` or `https://`
	TargetUrl       *string `json:"target_url,omitempty"`
	TargetSecretSet *bool   `json:"target_secret_set,omitempty"`
	Description     *string `json:"description,omitempty"`
	// Turn on or off your endpoint.
	Enabled                *bool                 `json:"enabled,omitempty"`
	Events                 []Items               `json:"events,omitempty"`
	ProjectIdFilter        []string              `json:"project_id_filter,omitempty"`
	EnvironmentTypesFilter []EnvironmentModeEnum `json:"environment_types_filter,omitempty"`
}

// NewOrganizationWebhookCreateResponseAllOf instantiates a new OrganizationWebhookCreateResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationWebhookCreateResponseAllOf() *OrganizationWebhookCreateResponseAllOf {
	this := OrganizationWebhookCreateResponseAllOf{}
	return &this
}

// NewOrganizationWebhookCreateResponseAllOfWithDefaults instantiates a new OrganizationWebhookCreateResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationWebhookCreateResponseAllOfWithDefaults() *OrganizationWebhookCreateResponseAllOf {
	this := OrganizationWebhookCreateResponseAllOf{}
	return &this
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *OrganizationWebhookCreateResponseAllOf) GetKind() Kind {
	if o == nil || o.Kind == nil {
		var ret Kind
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationWebhookCreateResponseAllOf) GetKindOk() (*Kind, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *OrganizationWebhookCreateResponseAllOf) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given Kind and assigns it to the Kind field.
func (o *OrganizationWebhookCreateResponseAllOf) SetKind(v Kind) {
	o.Kind = &v
}

// GetTargetUrl returns the TargetUrl field value if set, zero value otherwise.
func (o *OrganizationWebhookCreateResponseAllOf) GetTargetUrl() string {
	if o == nil || o.TargetUrl == nil {
		var ret string
		return ret
	}
	return *o.TargetUrl
}

// GetTargetUrlOk returns a tuple with the TargetUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationWebhookCreateResponseAllOf) GetTargetUrlOk() (*string, bool) {
	if o == nil || o.TargetUrl == nil {
		return nil, false
	}
	return o.TargetUrl, true
}

// HasTargetUrl returns a boolean if a field has been set.
func (o *OrganizationWebhookCreateResponseAllOf) HasTargetUrl() bool {
	if o != nil && o.TargetUrl != nil {
		return true
	}

	return false
}

// SetTargetUrl gets a reference to the given string and assigns it to the TargetUrl field.
func (o *OrganizationWebhookCreateResponseAllOf) SetTargetUrl(v string) {
	o.TargetUrl = &v
}

// GetTargetSecretSet returns the TargetSecretSet field value if set, zero value otherwise.
func (o *OrganizationWebhookCreateResponseAllOf) GetTargetSecretSet() bool {
	if o == nil || o.TargetSecretSet == nil {
		var ret bool
		return ret
	}
	return *o.TargetSecretSet
}

// GetTargetSecretSetOk returns a tuple with the TargetSecretSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationWebhookCreateResponseAllOf) GetTargetSecretSetOk() (*bool, bool) {
	if o == nil || o.TargetSecretSet == nil {
		return nil, false
	}
	return o.TargetSecretSet, true
}

// HasTargetSecretSet returns a boolean if a field has been set.
func (o *OrganizationWebhookCreateResponseAllOf) HasTargetSecretSet() bool {
	if o != nil && o.TargetSecretSet != nil {
		return true
	}

	return false
}

// SetTargetSecretSet gets a reference to the given bool and assigns it to the TargetSecretSet field.
func (o *OrganizationWebhookCreateResponseAllOf) SetTargetSecretSet(v bool) {
	o.TargetSecretSet = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OrganizationWebhookCreateResponseAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationWebhookCreateResponseAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OrganizationWebhookCreateResponseAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OrganizationWebhookCreateResponseAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *OrganizationWebhookCreateResponseAllOf) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationWebhookCreateResponseAllOf) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *OrganizationWebhookCreateResponseAllOf) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *OrganizationWebhookCreateResponseAllOf) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *OrganizationWebhookCreateResponseAllOf) GetEvents() []Items {
	if o == nil || o.Events == nil {
		var ret []Items
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationWebhookCreateResponseAllOf) GetEventsOk() ([]Items, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *OrganizationWebhookCreateResponseAllOf) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []Items and assigns it to the Events field.
func (o *OrganizationWebhookCreateResponseAllOf) SetEvents(v []Items) {
	o.Events = v
}

// GetProjectIdFilter returns the ProjectIdFilter field value if set, zero value otherwise.
func (o *OrganizationWebhookCreateResponseAllOf) GetProjectIdFilter() []string {
	if o == nil || o.ProjectIdFilter == nil {
		var ret []string
		return ret
	}
	return o.ProjectIdFilter
}

// GetProjectIdFilterOk returns a tuple with the ProjectIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationWebhookCreateResponseAllOf) GetProjectIdFilterOk() ([]string, bool) {
	if o == nil || o.ProjectIdFilter == nil {
		return nil, false
	}
	return o.ProjectIdFilter, true
}

// HasProjectIdFilter returns a boolean if a field has been set.
func (o *OrganizationWebhookCreateResponseAllOf) HasProjectIdFilter() bool {
	if o != nil && o.ProjectIdFilter != nil {
		return true
	}

	return false
}

// SetProjectIdFilter gets a reference to the given []string and assigns it to the ProjectIdFilter field.
func (o *OrganizationWebhookCreateResponseAllOf) SetProjectIdFilter(v []string) {
	o.ProjectIdFilter = v
}

// GetEnvironmentTypesFilter returns the EnvironmentTypesFilter field value if set, zero value otherwise.
func (o *OrganizationWebhookCreateResponseAllOf) GetEnvironmentTypesFilter() []EnvironmentModeEnum {
	if o == nil || o.EnvironmentTypesFilter == nil {
		var ret []EnvironmentModeEnum
		return ret
	}
	return o.EnvironmentTypesFilter
}

// GetEnvironmentTypesFilterOk returns a tuple with the EnvironmentTypesFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationWebhookCreateResponseAllOf) GetEnvironmentTypesFilterOk() ([]EnvironmentModeEnum, bool) {
	if o == nil || o.EnvironmentTypesFilter == nil {
		return nil, false
	}
	return o.EnvironmentTypesFilter, true
}

// HasEnvironmentTypesFilter returns a boolean if a field has been set.
func (o *OrganizationWebhookCreateResponseAllOf) HasEnvironmentTypesFilter() bool {
	if o != nil && o.EnvironmentTypesFilter != nil {
		return true
	}

	return false
}

// SetEnvironmentTypesFilter gets a reference to the given []EnvironmentModeEnum and assigns it to the EnvironmentTypesFilter field.
func (o *OrganizationWebhookCreateResponseAllOf) SetEnvironmentTypesFilter(v []EnvironmentModeEnum) {
	o.EnvironmentTypesFilter = v
}

func (o OrganizationWebhookCreateResponseAllOf) MarshalJSON() ([]byte, error) {
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

type NullableOrganizationWebhookCreateResponseAllOf struct {
	value *OrganizationWebhookCreateResponseAllOf
	isSet bool
}

func (v NullableOrganizationWebhookCreateResponseAllOf) Get() *OrganizationWebhookCreateResponseAllOf {
	return v.value
}

func (v *NullableOrganizationWebhookCreateResponseAllOf) Set(val *OrganizationWebhookCreateResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationWebhookCreateResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationWebhookCreateResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationWebhookCreateResponseAllOf(val *OrganizationWebhookCreateResponseAllOf) *NullableOrganizationWebhookCreateResponseAllOf {
	return &NullableOrganizationWebhookCreateResponseAllOf{value: val, isSet: true}
}

func (v NullableOrganizationWebhookCreateResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationWebhookCreateResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}