# CreateCustomerDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InternalId** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Fields** | Pointer to **map[string]interface{}** |  | [optional] 
**Credentials** | Pointer to **map[string]interface{}** |  | [optional] 
**IsTest** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreateCustomerDto

`func NewCreateCustomerDto(internalId string, ) *CreateCustomerDto`

NewCreateCustomerDto instantiates a new CreateCustomerDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCustomerDtoWithDefaults

`func NewCreateCustomerDtoWithDefaults() *CreateCustomerDto`

NewCreateCustomerDtoWithDefaults instantiates a new CreateCustomerDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInternalId

`func (o *CreateCustomerDto) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *CreateCustomerDto) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *CreateCustomerDto) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.


### GetName

`func (o *CreateCustomerDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateCustomerDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateCustomerDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateCustomerDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFields

`func (o *CreateCustomerDto) GetFields() map[string]interface{}`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *CreateCustomerDto) GetFieldsOk() (*map[string]interface{}, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *CreateCustomerDto) SetFields(v map[string]interface{})`

SetFields sets Fields field to given value.

### HasFields

`func (o *CreateCustomerDto) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetCredentials

`func (o *CreateCustomerDto) GetCredentials() map[string]interface{}`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *CreateCustomerDto) GetCredentialsOk() (*map[string]interface{}, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *CreateCustomerDto) SetCredentials(v map[string]interface{})`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *CreateCustomerDto) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetIsTest

`func (o *CreateCustomerDto) GetIsTest() bool`

GetIsTest returns the IsTest field if non-nil, zero value otherwise.

### GetIsTestOk

`func (o *CreateCustomerDto) GetIsTestOk() (*bool, bool)`

GetIsTestOk returns a tuple with the IsTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTest

`func (o *CreateCustomerDto) SetIsTest(v bool)`

SetIsTest sets IsTest field to given value.

### HasIsTest

`func (o *CreateCustomerDto) HasIsTest() bool`

HasIsTest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


