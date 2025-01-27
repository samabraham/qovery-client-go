# SecretAlias

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Key** | **string** |  | 
**Scope** | [**APIVariableScopeEnum**](APIVariableScopeEnum.md) |  | 

## Methods

### NewSecretAlias

`func NewSecretAlias(id string, key string, scope APIVariableScopeEnum, ) *SecretAlias`

NewSecretAlias instantiates a new SecretAlias object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretAliasWithDefaults

`func NewSecretAliasWithDefaults() *SecretAlias`

NewSecretAliasWithDefaults instantiates a new SecretAlias object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SecretAlias) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecretAlias) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecretAlias) SetId(v string)`

SetId sets Id field to given value.


### GetKey

`func (o *SecretAlias) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SecretAlias) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SecretAlias) SetKey(v string)`

SetKey sets Key field to given value.


### GetScope

`func (o *SecretAlias) GetScope() APIVariableScopeEnum`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *SecretAlias) GetScopeOk() (*APIVariableScopeEnum, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *SecretAlias) SetScope(v APIVariableScopeEnum)`

SetScope sets Scope field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


