# OrganizationWebhookCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | Define the type of the webhook. &#x60;SLACK&#x60; is a special webhook type to push notifications directly to slack. The &#x60;target_url&#x60; must be a Slack compatible endpoint. | 
**TargetUrl** | **string** | Set the public HTTP or HTTPS endpoint that will receive the specified events. The target URL must starts with &#x60;http://&#x60; or &#x60;https://&#x60;  | 
**TargetSecret** | Pointer to **string** | Make sure you receive a payload to sign the Qovery request with your secret. Qovery will add a HTTP header &#x60;Qovery-Signature: &lt;Your Secret&gt;&#x60; to every webhook requests sent to your target URL.  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** | Turn on or off your endpoint. | [optional] 
**Events** | **[]string** |  | 
**ProjectIdFilter** | Pointer to **[]string** |  | [optional] 
**EnvironmentTypesFilter** | Pointer to [**[]EnvironmentModeEnum**](EnvironmentModeEnum.md) |  | [optional] 

## Methods

### NewOrganizationWebhookCreateRequest

`func NewOrganizationWebhookCreateRequest(kind string, targetUrl string, events []string, ) *OrganizationWebhookCreateRequest`

NewOrganizationWebhookCreateRequest instantiates a new OrganizationWebhookCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationWebhookCreateRequestWithDefaults

`func NewOrganizationWebhookCreateRequestWithDefaults() *OrganizationWebhookCreateRequest`

NewOrganizationWebhookCreateRequestWithDefaults instantiates a new OrganizationWebhookCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *OrganizationWebhookCreateRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *OrganizationWebhookCreateRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *OrganizationWebhookCreateRequest) SetKind(v string)`

SetKind sets Kind field to given value.


### GetTargetUrl

`func (o *OrganizationWebhookCreateRequest) GetTargetUrl() string`

GetTargetUrl returns the TargetUrl field if non-nil, zero value otherwise.

### GetTargetUrlOk

`func (o *OrganizationWebhookCreateRequest) GetTargetUrlOk() (*string, bool)`

GetTargetUrlOk returns a tuple with the TargetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUrl

`func (o *OrganizationWebhookCreateRequest) SetTargetUrl(v string)`

SetTargetUrl sets TargetUrl field to given value.


### GetTargetSecret

`func (o *OrganizationWebhookCreateRequest) GetTargetSecret() string`

GetTargetSecret returns the TargetSecret field if non-nil, zero value otherwise.

### GetTargetSecretOk

`func (o *OrganizationWebhookCreateRequest) GetTargetSecretOk() (*string, bool)`

GetTargetSecretOk returns a tuple with the TargetSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSecret

`func (o *OrganizationWebhookCreateRequest) SetTargetSecret(v string)`

SetTargetSecret sets TargetSecret field to given value.

### HasTargetSecret

`func (o *OrganizationWebhookCreateRequest) HasTargetSecret() bool`

HasTargetSecret returns a boolean if a field has been set.

### GetDescription

`func (o *OrganizationWebhookCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrganizationWebhookCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrganizationWebhookCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OrganizationWebhookCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *OrganizationWebhookCreateRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OrganizationWebhookCreateRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OrganizationWebhookCreateRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *OrganizationWebhookCreateRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEvents

`func (o *OrganizationWebhookCreateRequest) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *OrganizationWebhookCreateRequest) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *OrganizationWebhookCreateRequest) SetEvents(v []string)`

SetEvents sets Events field to given value.


### GetProjectIdFilter

`func (o *OrganizationWebhookCreateRequest) GetProjectIdFilter() []string`

GetProjectIdFilter returns the ProjectIdFilter field if non-nil, zero value otherwise.

### GetProjectIdFilterOk

`func (o *OrganizationWebhookCreateRequest) GetProjectIdFilterOk() (*[]string, bool)`

GetProjectIdFilterOk returns a tuple with the ProjectIdFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIdFilter

`func (o *OrganizationWebhookCreateRequest) SetProjectIdFilter(v []string)`

SetProjectIdFilter sets ProjectIdFilter field to given value.

### HasProjectIdFilter

`func (o *OrganizationWebhookCreateRequest) HasProjectIdFilter() bool`

HasProjectIdFilter returns a boolean if a field has been set.

### GetEnvironmentTypesFilter

`func (o *OrganizationWebhookCreateRequest) GetEnvironmentTypesFilter() []EnvironmentModeEnum`

GetEnvironmentTypesFilter returns the EnvironmentTypesFilter field if non-nil, zero value otherwise.

### GetEnvironmentTypesFilterOk

`func (o *OrganizationWebhookCreateRequest) GetEnvironmentTypesFilterOk() (*[]EnvironmentModeEnum, bool)`

GetEnvironmentTypesFilterOk returns a tuple with the EnvironmentTypesFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentTypesFilter

`func (o *OrganizationWebhookCreateRequest) SetEnvironmentTypesFilter(v []EnvironmentModeEnum)`

SetEnvironmentTypesFilter sets EnvironmentTypesFilter field to given value.

### HasEnvironmentTypesFilter

`func (o *OrganizationWebhookCreateRequest) HasEnvironmentTypesFilter() bool`

HasEnvironmentTypesFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

