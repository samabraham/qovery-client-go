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

// CommitPaginatedResponseList struct for CommitPaginatedResponseList
type CommitPaginatedResponseList struct {
	Results  []CommitResponse `json:"results,omitempty"`
	Page     float32          `json:"page"`
	PageSize float32          `json:"page_size"`
}

// NewCommitPaginatedResponseList instantiates a new CommitPaginatedResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommitPaginatedResponseList(page float32, pageSize float32) *CommitPaginatedResponseList {
	this := CommitPaginatedResponseList{}
	this.Page = page
	this.PageSize = pageSize
	return &this
}

// NewCommitPaginatedResponseListWithDefaults instantiates a new CommitPaginatedResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommitPaginatedResponseListWithDefaults() *CommitPaginatedResponseList {
	this := CommitPaginatedResponseList{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *CommitPaginatedResponseList) GetResults() []CommitResponse {
	if o == nil || o.Results == nil {
		var ret []CommitResponse
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitPaginatedResponseList) GetResultsOk() ([]CommitResponse, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *CommitPaginatedResponseList) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []CommitResponse and assigns it to the Results field.
func (o *CommitPaginatedResponseList) SetResults(v []CommitResponse) {
	o.Results = v
}

// GetPage returns the Page field value
func (o *CommitPaginatedResponseList) GetPage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *CommitPaginatedResponseList) GetPageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *CommitPaginatedResponseList) SetPage(v float32) {
	o.Page = v
}

// GetPageSize returns the PageSize field value
func (o *CommitPaginatedResponseList) GetPageSize() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value
// and a boolean to check if the value has been set.
func (o *CommitPaginatedResponseList) GetPageSizeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageSize, true
}

// SetPageSize sets field value
func (o *CommitPaginatedResponseList) SetPageSize(v float32) {
	o.PageSize = v
}

func (o CommitPaginatedResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	if true {
		toSerialize["page"] = o.Page
	}
	if true {
		toSerialize["page_size"] = o.PageSize
	}
	return json.Marshal(toSerialize)
}

type NullableCommitPaginatedResponseList struct {
	value *CommitPaginatedResponseList
	isSet bool
}

func (v NullableCommitPaginatedResponseList) Get() *CommitPaginatedResponseList {
	return v.value
}

func (v *NullableCommitPaginatedResponseList) Set(val *CommitPaginatedResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableCommitPaginatedResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableCommitPaginatedResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommitPaginatedResponseList(val *CommitPaginatedResponseList) *NullableCommitPaginatedResponseList {
	return &NullableCommitPaginatedResponseList{value: val, isSet: true}
}

func (v NullableCommitPaginatedResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommitPaginatedResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
