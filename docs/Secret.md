# Secret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Key** | **string** | key is case sensitive | 
**OverriddenSecret** | Pointer to [**SecretOverride**](SecretOverride.md) |  | [optional] 
**AliasedSecret** | Pointer to [**SecretAlias**](SecretAlias.md) |  | [optional] 
**Scope** | [**APIVariableScopeEnum**](APIVariableScopeEnum.md) |  | 
**ServiceId** | Pointer to **string** | present only for &#x60;BUILT_IN&#x60; variable | [optional] 
**ServiceName** | Pointer to **string** | present only for &#x60;BUILT_IN&#x60; variable | [optional] 
**ServiceType** | Pointer to [**LinkedServiceTypeEnum**](LinkedServiceTypeEnum.md) |  | [optional] 

## Methods

### NewSecret

`func NewSecret(id string, createdAt time.Time, key string, scope APIVariableScopeEnum, ) *Secret`

NewSecret instantiates a new Secret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretWithDefaults

`func NewSecretWithDefaults() *Secret`

NewSecretWithDefaults instantiates a new Secret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Secret) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Secret) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Secret) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *Secret) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Secret) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Secret) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Secret) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Secret) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Secret) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Secret) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetKey

`func (o *Secret) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Secret) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Secret) SetKey(v string)`

SetKey sets Key field to given value.


### GetOverriddenSecret

`func (o *Secret) GetOverriddenSecret() SecretOverride`

GetOverriddenSecret returns the OverriddenSecret field if non-nil, zero value otherwise.

### GetOverriddenSecretOk

`func (o *Secret) GetOverriddenSecretOk() (*SecretOverride, bool)`

GetOverriddenSecretOk returns a tuple with the OverriddenSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverriddenSecret

`func (o *Secret) SetOverriddenSecret(v SecretOverride)`

SetOverriddenSecret sets OverriddenSecret field to given value.

### HasOverriddenSecret

`func (o *Secret) HasOverriddenSecret() bool`

HasOverriddenSecret returns a boolean if a field has been set.

### GetAliasedSecret

`func (o *Secret) GetAliasedSecret() SecretAlias`

GetAliasedSecret returns the AliasedSecret field if non-nil, zero value otherwise.

### GetAliasedSecretOk

`func (o *Secret) GetAliasedSecretOk() (*SecretAlias, bool)`

GetAliasedSecretOk returns a tuple with the AliasedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasedSecret

`func (o *Secret) SetAliasedSecret(v SecretAlias)`

SetAliasedSecret sets AliasedSecret field to given value.

### HasAliasedSecret

`func (o *Secret) HasAliasedSecret() bool`

HasAliasedSecret returns a boolean if a field has been set.

### GetScope

`func (o *Secret) GetScope() APIVariableScopeEnum`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *Secret) GetScopeOk() (*APIVariableScopeEnum, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *Secret) SetScope(v APIVariableScopeEnum)`

SetScope sets Scope field to given value.


### GetServiceId

`func (o *Secret) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *Secret) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *Secret) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *Secret) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetServiceName

`func (o *Secret) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *Secret) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *Secret) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *Secret) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetServiceType

`func (o *Secret) GetServiceType() LinkedServiceTypeEnum`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *Secret) GetServiceTypeOk() (*LinkedServiceTypeEnum, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *Secret) SetServiceType(v LinkedServiceTypeEnum)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *Secret) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


