# ConnectionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**UserId** | **string** |  | 
**IsTest** | Pointer to **bool** |  | [optional] 
**User** | Pointer to [**CustomerDto**](CustomerDto.md) |  | [optional] 
**IntegrationId** | **string** |  | 
**Integration** | Pointer to [**IntegrationDto**](IntegrationDto.md) |  | [optional] 
**Disconnected** | Pointer to **bool** |  | [optional] 
**Credentials** | Pointer to **string** |  | [optional] 
**IsDefunct** | Pointer to **bool** |  | [optional] 
**Error** | Pointer to **map[string]interface{}** |  | [optional] 
**Name** | **string** |  | 
**CreatedAt** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**LastActiveAt** | Pointer to **string** |  | [optional] 
**ArchivedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewConnectionDto

`func NewConnectionDto(id string, userId string, integrationId string, name string, ) *ConnectionDto`

NewConnectionDto instantiates a new ConnectionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionDtoWithDefaults

`func NewConnectionDtoWithDefaults() *ConnectionDto`

NewConnectionDtoWithDefaults instantiates a new ConnectionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConnectionDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConnectionDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConnectionDto) SetId(v string)`

SetId sets Id field to given value.


### GetUserId

`func (o *ConnectionDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ConnectionDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ConnectionDto) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetIsTest

`func (o *ConnectionDto) GetIsTest() bool`

GetIsTest returns the IsTest field if non-nil, zero value otherwise.

### GetIsTestOk

`func (o *ConnectionDto) GetIsTestOk() (*bool, bool)`

GetIsTestOk returns a tuple with the IsTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTest

`func (o *ConnectionDto) SetIsTest(v bool)`

SetIsTest sets IsTest field to given value.

### HasIsTest

`func (o *ConnectionDto) HasIsTest() bool`

HasIsTest returns a boolean if a field has been set.

### GetUser

`func (o *ConnectionDto) GetUser() CustomerDto`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ConnectionDto) GetUserOk() (*CustomerDto, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ConnectionDto) SetUser(v CustomerDto)`

SetUser sets User field to given value.

### HasUser

`func (o *ConnectionDto) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetIntegrationId

`func (o *ConnectionDto) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *ConnectionDto) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *ConnectionDto) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.


### GetIntegration

`func (o *ConnectionDto) GetIntegration() IntegrationDto`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *ConnectionDto) GetIntegrationOk() (*IntegrationDto, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *ConnectionDto) SetIntegration(v IntegrationDto)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *ConnectionDto) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetDisconnected

`func (o *ConnectionDto) GetDisconnected() bool`

GetDisconnected returns the Disconnected field if non-nil, zero value otherwise.

### GetDisconnectedOk

`func (o *ConnectionDto) GetDisconnectedOk() (*bool, bool)`

GetDisconnectedOk returns a tuple with the Disconnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnected

`func (o *ConnectionDto) SetDisconnected(v bool)`

SetDisconnected sets Disconnected field to given value.

### HasDisconnected

`func (o *ConnectionDto) HasDisconnected() bool`

HasDisconnected returns a boolean if a field has been set.

### GetCredentials

`func (o *ConnectionDto) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *ConnectionDto) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *ConnectionDto) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *ConnectionDto) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetIsDefunct

`func (o *ConnectionDto) GetIsDefunct() bool`

GetIsDefunct returns the IsDefunct field if non-nil, zero value otherwise.

### GetIsDefunctOk

`func (o *ConnectionDto) GetIsDefunctOk() (*bool, bool)`

GetIsDefunctOk returns a tuple with the IsDefunct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefunct

`func (o *ConnectionDto) SetIsDefunct(v bool)`

SetIsDefunct sets IsDefunct field to given value.

### HasIsDefunct

`func (o *ConnectionDto) HasIsDefunct() bool`

HasIsDefunct returns a boolean if a field has been set.

### GetError

`func (o *ConnectionDto) GetError() map[string]interface{}`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ConnectionDto) GetErrorOk() (*map[string]interface{}, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ConnectionDto) SetError(v map[string]interface{})`

SetError sets Error field to given value.

### HasError

`func (o *ConnectionDto) HasError() bool`

HasError returns a boolean if a field has been set.

### GetName

`func (o *ConnectionDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectionDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectionDto) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedAt

`func (o *ConnectionDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ConnectionDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ConnectionDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ConnectionDto) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ConnectionDto) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ConnectionDto) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ConnectionDto) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ConnectionDto) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetLastActiveAt

`func (o *ConnectionDto) GetLastActiveAt() string`

GetLastActiveAt returns the LastActiveAt field if non-nil, zero value otherwise.

### GetLastActiveAtOk

`func (o *ConnectionDto) GetLastActiveAtOk() (*string, bool)`

GetLastActiveAtOk returns a tuple with the LastActiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActiveAt

`func (o *ConnectionDto) SetLastActiveAt(v string)`

SetLastActiveAt sets LastActiveAt field to given value.

### HasLastActiveAt

`func (o *ConnectionDto) HasLastActiveAt() bool`

HasLastActiveAt returns a boolean if a field has been set.

### GetArchivedAt

`func (o *ConnectionDto) GetArchivedAt() string`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *ConnectionDto) GetArchivedAtOk() (*string, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *ConnectionDto) SetArchivedAt(v string)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *ConnectionDto) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


