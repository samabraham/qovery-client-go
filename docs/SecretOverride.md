# SecretOverride

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Key** | **string** |  | 
**Scope** | [**APIVariableScopeEnum**](APIVariableScopeEnum.md) |  | 

## Methods

### NewSecretOverride

`func NewSecretOverride(id string, key string, scope APIVariableScopeEnum, ) *SecretOverride`

NewSecretOverride instantiates a new SecretOverride object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretOverrideWithDefaults

`func NewSecretOverrideWithDefaults() *SecretOverride`

NewSecretOverrideWithDefaults instantiates a new SecretOverride object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SecretOverride) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecretOverride) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecretOverride) SetId(v string)`

SetId sets Id field to given value.


### GetKey

`func (o *SecretOverride) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SecretOverride) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SecretOverride) SetKey(v string)`

SetKey sets Key field to given value.


### GetScope

`func (o *SecretOverride) GetScope() APIVariableScopeEnum`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *SecretOverride) GetScopeOk() (*APIVariableScopeEnum, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *SecretOverride) SetScope(v APIVariableScopeEnum)`

SetScope sets Scope field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


