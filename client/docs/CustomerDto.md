# CustomerDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**InternalId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**IsTest** | Pointer to **bool** |  | [optional] 
**Fields** | Pointer to **map[string]interface{}** |  | [optional] 
**Credentials** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**LastActiveAt** | Pointer to **string** |  | [optional] 
**IsBillable** | Pointer to **bool** |  | [optional] 

## Methods

### NewCustomerDto

`func NewCustomerDto(id string, ) *CustomerDto`

NewCustomerDto instantiates a new CustomerDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerDtoWithDefaults

`func NewCustomerDtoWithDefaults() *CustomerDto`

NewCustomerDtoWithDefaults instantiates a new CustomerDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomerDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomerDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomerDto) SetId(v string)`

SetId sets Id field to given value.


### GetInternalId

`func (o *CustomerDto) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *CustomerDto) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *CustomerDto) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *CustomerDto) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetName

`func (o *CustomerDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomerDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomerDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomerDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIsTest

`func (o *CustomerDto) GetIsTest() bool`

GetIsTest returns the IsTest field if non-nil, zero value otherwise.

### GetIsTestOk

`func (o *CustomerDto) GetIsTestOk() (*bool, bool)`

GetIsTestOk returns a tuple with the IsTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTest

`func (o *CustomerDto) SetIsTest(v bool)`

SetIsTest sets IsTest field to given value.

### HasIsTest

`func (o *CustomerDto) HasIsTest() bool`

HasIsTest returns a boolean if a field has been set.

### GetFields

`func (o *CustomerDto) GetFields() map[string]interface{}`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *CustomerDto) GetFieldsOk() (*map[string]interface{}, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *CustomerDto) SetFields(v map[string]interface{})`

SetFields sets Fields field to given value.

### HasFields

`func (o *CustomerDto) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetCredentials

`func (o *CustomerDto) GetCredentials() map[string]interface{}`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *CustomerDto) GetCredentialsOk() (*map[string]interface{}, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *CustomerDto) SetCredentials(v map[string]interface{})`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *CustomerDto) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CustomerDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CustomerDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CustomerDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CustomerDto) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastActiveAt

`func (o *CustomerDto) GetLastActiveAt() string`

GetLastActiveAt returns the LastActiveAt field if non-nil, zero value otherwise.

### GetLastActiveAtOk

`func (o *CustomerDto) GetLastActiveAtOk() (*string, bool)`

GetLastActiveAtOk returns a tuple with the LastActiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActiveAt

`func (o *CustomerDto) SetLastActiveAt(v string)`

SetLastActiveAt sets LastActiveAt field to given value.

### HasLastActiveAt

`func (o *CustomerDto) HasLastActiveAt() bool`

HasLastActiveAt returns a boolean if a field has been set.

### GetIsBillable

`func (o *CustomerDto) GetIsBillable() bool`

GetIsBillable returns the IsBillable field if non-nil, zero value otherwise.

### GetIsBillableOk

`func (o *CustomerDto) GetIsBillableOk() (*bool, bool)`

GetIsBillableOk returns a tuple with the IsBillable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBillable

`func (o *CustomerDto) SetIsBillable(v bool)`

SetIsBillable sets IsBillable field to given value.

### HasIsBillable

`func (o *CustomerDto) HasIsBillable() bool`

HasIsBillable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


