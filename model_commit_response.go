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
	"time"
)

// CommitResponse struct for CommitResponse
type CommitResponse struct {
	CreatedAt       time.Time `json:"created_at"`
	GitCommitId     string    `json:"git_commit_id"`
	Tag             string    `json:"tag"`
	Message         string    `json:"message"`
	AuthorName      string    `json:"author_name"`
	AuthorAvatarUrl *string   `json:"author_avatar_url,omitempty"`
	CommitPageUrl   *string   `json:"commit_page_url,omitempty"`
}

// NewCommitResponse instantiates a new CommitResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommitResponse(createdAt time.Time, gitCommitId string, tag string, message string, authorName string) *CommitResponse {
	this := CommitResponse{}
	this.CreatedAt = createdAt
	this.GitCommitId = gitCommitId
	this.Tag = tag
	this.Message = message
	this.AuthorName = authorName
	return &this
}

// NewCommitResponseWithDefaults instantiates a new CommitResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommitResponseWithDefaults() *CommitResponse {
	this := CommitResponse{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *CommitResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CommitResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *CommitResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetGitCommitId returns the GitCommitId field value
func (o *CommitResponse) GetGitCommitId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GitCommitId
}

// GetGitCommitIdOk returns a tuple with the GitCommitId field value
// and a boolean to check if the value has been set.
func (o *CommitResponse) GetGitCommitIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GitCommitId, true
}

// SetGitCommitId sets field value
func (o *CommitResponse) SetGitCommitId(v string) {
	o.GitCommitId = v
}

// GetTag returns the Tag field value
func (o *CommitResponse) GetTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value
// and a boolean to check if the value has been set.
func (o *CommitResponse) GetTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tag, true
}

// SetTag sets field value
func (o *CommitResponse) SetTag(v string) {
	o.Tag = v
}

// GetMessage returns the Message field value
func (o *CommitResponse) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *CommitResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *CommitResponse) SetMessage(v string) {
	o.Message = v
}

// GetAuthorName returns the AuthorName field value
func (o *CommitResponse) GetAuthorName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthorName
}

// GetAuthorNameOk returns a tuple with the AuthorName field value
// and a boolean to check if the value has been set.
func (o *CommitResponse) GetAuthorNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorName, true
}

// SetAuthorName sets field value
func (o *CommitResponse) SetAuthorName(v string) {
	o.AuthorName = v
}

// GetAuthorAvatarUrl returns the AuthorAvatarUrl field value if set, zero value otherwise.
func (o *CommitResponse) GetAuthorAvatarUrl() string {
	if o == nil || o.AuthorAvatarUrl == nil {
		var ret string
		return ret
	}
	return *o.AuthorAvatarUrl
}

// GetAuthorAvatarUrlOk returns a tuple with the AuthorAvatarUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitResponse) GetAuthorAvatarUrlOk() (*string, bool) {
	if o == nil || o.AuthorAvatarUrl == nil {
		return nil, false
	}
	return o.AuthorAvatarUrl, true
}

// HasAuthorAvatarUrl returns a boolean if a field has been set.
func (o *CommitResponse) HasAuthorAvatarUrl() bool {
	if o != nil && o.AuthorAvatarUrl != nil {
		return true
	}

	return false
}

// SetAuthorAvatarUrl gets a reference to the given string and assigns it to the AuthorAvatarUrl field.
func (o *CommitResponse) SetAuthorAvatarUrl(v string) {
	o.AuthorAvatarUrl = &v
}

// GetCommitPageUrl returns the CommitPageUrl field value if set, zero value otherwise.
func (o *CommitResponse) GetCommitPageUrl() string {
	if o == nil || o.CommitPageUrl == nil {
		var ret string
		return ret
	}
	return *o.CommitPageUrl
}

// GetCommitPageUrlOk returns a tuple with the CommitPageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitResponse) GetCommitPageUrlOk() (*string, bool) {
	if o == nil || o.CommitPageUrl == nil {
		return nil, false
	}
	return o.CommitPageUrl, true
}

// HasCommitPageUrl returns a boolean if a field has been set.
func (o *CommitResponse) HasCommitPageUrl() bool {
	if o != nil && o.CommitPageUrl != nil {
		return true
	}

	return false
}

// SetCommitPageUrl gets a reference to the given string and assigns it to the CommitPageUrl field.
func (o *CommitResponse) SetCommitPageUrl(v string) {
	o.CommitPageUrl = &v
}

func (o CommitResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["git_commit_id"] = o.GitCommitId
	}
	if true {
		toSerialize["tag"] = o.Tag
	}
	if true {
		toSerialize["message"] = o.Message
	}
	if true {
		toSerialize["author_name"] = o.AuthorName
	}
	if o.AuthorAvatarUrl != nil {
		toSerialize["author_avatar_url"] = o.AuthorAvatarUrl
	}
	if o.CommitPageUrl != nil {
		toSerialize["commit_page_url"] = o.CommitPageUrl
	}
	return json.Marshal(toSerialize)
}

type NullableCommitResponse struct {
	value *CommitResponse
	isSet bool
}

func (v NullableCommitResponse) Get() *CommitResponse {
	return v.value
}

func (v *NullableCommitResponse) Set(val *CommitResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCommitResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCommitResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommitResponse(val *CommitResponse) *NullableCommitResponse {
	return &NullableCommitResponse{value: val, isSet: true}
}

func (v NullableCommitResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommitResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
