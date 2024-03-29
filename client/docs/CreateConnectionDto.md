# CreateConnectionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntegrationId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Credentials** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateConnectionDto

`func NewCreateConnectionDto() *CreateConnectionDto`

NewCreateConnectionDto instantiates a new CreateConnectionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateConnectionDtoWithDefaults

`func NewCreateConnectionDtoWithDefaults() *CreateConnectionDto`

NewCreateConnectionDtoWithDefaults instantiates a new CreateConnectionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntegrationId

`func (o *CreateConnectionDto) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *CreateConnectionDto) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *CreateConnectionDto) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.

### HasIntegrationId

`func (o *CreateConnectionDto) HasIntegrationId() bool`

HasIntegrationId returns a boolean if a field has been set.

### GetName

`func (o *CreateConnectionDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateConnectionDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateConnectionDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateConnectionDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCredentials

`func (o *CreateConnectionDto) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *CreateConnectionDto) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *CreateConnectionDto) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *CreateConnectionDto) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


