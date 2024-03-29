# DataLinkTableInstanceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**DataLinkTableId** | **string** |  | 
**DataLinkTable** | Pointer to [**DataLinkTableDto**](DataLinkTableDto.md) |  | [optional] 
**UserId** | **string** |  | 
**User** | Pointer to [**CustomerDto**](CustomerDto.md) |  | [optional] 
**ConnectionId** | **string** |  | 
**Connection** | Pointer to [**ConnectionDto**](ConnectionDto.md) |  | [optional] 
**IntegrationId** | **string** |  | 
**Integration** | Pointer to [**IntegrationDto**](IntegrationDto.md) |  | [optional] 
**InstanceKey** | Pointer to **string** |  | [optional] 

## Methods

### NewDataLinkTableInstanceDto

`func NewDataLinkTableInstanceDto(id string, name string, dataLinkTableId string, userId string, connectionId string, integrationId string, ) *DataLinkTableInstanceDto`

NewDataLinkTableInstanceDto instantiates a new DataLinkTableInstanceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataLinkTableInstanceDtoWithDefaults

`func NewDataLinkTableInstanceDtoWithDefaults() *DataLinkTableInstanceDto`

NewDataLinkTableInstanceDtoWithDefaults instantiates a new DataLinkTableInstanceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DataLinkTableInstanceDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataLinkTableInstanceDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataLinkTableInstanceDto) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *DataLinkTableInstanceDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataLinkTableInstanceDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataLinkTableInstanceDto) SetName(v string)`

SetName sets Name field to given value.


### GetDataLinkTableId

`func (o *DataLinkTableInstanceDto) GetDataLinkTableId() string`

GetDataLinkTableId returns the DataLinkTableId field if non-nil, zero value otherwise.

### GetDataLinkTableIdOk

`func (o *DataLinkTableInstanceDto) GetDataLinkTableIdOk() (*string, bool)`

GetDataLinkTableIdOk returns a tuple with the DataLinkTableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLinkTableId

`func (o *DataLinkTableInstanceDto) SetDataLinkTableId(v string)`

SetDataLinkTableId sets DataLinkTableId field to given value.


### GetDataLinkTable

`func (o *DataLinkTableInstanceDto) GetDataLinkTable() DataLinkTableDto`

GetDataLinkTable returns the DataLinkTable field if non-nil, zero value otherwise.

### GetDataLinkTableOk

`func (o *DataLinkTableInstanceDto) GetDataLinkTableOk() (*DataLinkTableDto, bool)`

GetDataLinkTableOk returns a tuple with the DataLinkTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLinkTable

`func (o *DataLinkTableInstanceDto) SetDataLinkTable(v DataLinkTableDto)`

SetDataLinkTable sets DataLinkTable field to given value.

### HasDataLinkTable

`func (o *DataLinkTableInstanceDto) HasDataLinkTable() bool`

HasDataLinkTable returns a boolean if a field has been set.

### GetUserId

`func (o *DataLinkTableInstanceDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *DataLinkTableInstanceDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *DataLinkTableInstanceDto) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetUser

`func (o *DataLinkTableInstanceDto) GetUser() CustomerDto`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *DataLinkTableInstanceDto) GetUserOk() (*CustomerDto, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *DataLinkTableInstanceDto) SetUser(v CustomerDto)`

SetUser sets User field to given value.

### HasUser

`func (o *DataLinkTableInstanceDto) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetConnectionId

`func (o *DataLinkTableInstanceDto) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *DataLinkTableInstanceDto) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *DataLinkTableInstanceDto) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetConnection

`func (o *DataLinkTableInstanceDto) GetConnection() ConnectionDto`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *DataLinkTableInstanceDto) GetConnectionOk() (*ConnectionDto, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *DataLinkTableInstanceDto) SetConnection(v ConnectionDto)`

SetConnection sets Connection field to given value.

### HasConnection

`func (o *DataLinkTableInstanceDto) HasConnection() bool`

HasConnection returns a boolean if a field has been set.

### GetIntegrationId

`func (o *DataLinkTableInstanceDto) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *DataLinkTableInstanceDto) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *DataLinkTableInstanceDto) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.


### GetIntegration

`func (o *DataLinkTableInstanceDto) GetIntegration() IntegrationDto`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *DataLinkTableInstanceDto) GetIntegrationOk() (*IntegrationDto, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *DataLinkTableInstanceDto) SetIntegration(v IntegrationDto)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *DataLinkTableInstanceDto) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetInstanceKey

`func (o *DataLinkTableInstanceDto) GetInstanceKey() string`

GetInstanceKey returns the InstanceKey field if non-nil, zero value otherwise.

### GetInstanceKeyOk

`func (o *DataLinkTableInstanceDto) GetInstanceKeyOk() (*string, bool)`

GetInstanceKeyOk returns a tuple with the InstanceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceKey

`func (o *DataLinkTableInstanceDto) SetInstanceKey(v string)`

SetInstanceKey sets InstanceKey field to given value.

### HasInstanceKey

`func (o *DataLinkTableInstanceDto) HasInstanceKey() bool`

HasInstanceKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


