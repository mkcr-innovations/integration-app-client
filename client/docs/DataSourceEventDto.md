# DataSourceEventDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**UserId** | **string** |  | 
**User** | Pointer to [**CustomerDto**](CustomerDto.md) |  | [optional] 
**DataSourceInstanceId** | **string** |  | 
**DataSourceInstance** | Pointer to [**DataSourceInstanceDto**](DataSourceInstanceDto.md) |  | [optional] 
**ConnectionId** | **string** |  | 
**Connection** | Pointer to [**ConnectionDto**](ConnectionDto.md) |  | [optional] 
**IntegrationId** | **string** |  | 
**Integration** | Pointer to [**IntegrationDto**](IntegrationDto.md) |  | [optional] 
**Event** | **map[string]interface{}** |  | 
**Datetime** | **string** |  | 
**LaunchedFlowRunIds** | **[]string** |  | 
**Error** | Pointer to **map[string]interface{}** |  | [optional] 
**Logs** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewDataSourceEventDto

`func NewDataSourceEventDto(id string, userId string, dataSourceInstanceId string, connectionId string, integrationId string, event map[string]interface{}, datetime string, launchedFlowRunIds []string, ) *DataSourceEventDto`

NewDataSourceEventDto instantiates a new DataSourceEventDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataSourceEventDtoWithDefaults

`func NewDataSourceEventDtoWithDefaults() *DataSourceEventDto`

NewDataSourceEventDtoWithDefaults instantiates a new DataSourceEventDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DataSourceEventDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataSourceEventDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataSourceEventDto) SetId(v string)`

SetId sets Id field to given value.


### GetUserId

`func (o *DataSourceEventDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *DataSourceEventDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *DataSourceEventDto) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetUser

`func (o *DataSourceEventDto) GetUser() CustomerDto`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *DataSourceEventDto) GetUserOk() (*CustomerDto, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *DataSourceEventDto) SetUser(v CustomerDto)`

SetUser sets User field to given value.

### HasUser

`func (o *DataSourceEventDto) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetDataSourceInstanceId

`func (o *DataSourceEventDto) GetDataSourceInstanceId() string`

GetDataSourceInstanceId returns the DataSourceInstanceId field if non-nil, zero value otherwise.

### GetDataSourceInstanceIdOk

`func (o *DataSourceEventDto) GetDataSourceInstanceIdOk() (*string, bool)`

GetDataSourceInstanceIdOk returns a tuple with the DataSourceInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceInstanceId

`func (o *DataSourceEventDto) SetDataSourceInstanceId(v string)`

SetDataSourceInstanceId sets DataSourceInstanceId field to given value.


### GetDataSourceInstance

`func (o *DataSourceEventDto) GetDataSourceInstance() DataSourceInstanceDto`

GetDataSourceInstance returns the DataSourceInstance field if non-nil, zero value otherwise.

### GetDataSourceInstanceOk

`func (o *DataSourceEventDto) GetDataSourceInstanceOk() (*DataSourceInstanceDto, bool)`

GetDataSourceInstanceOk returns a tuple with the DataSourceInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceInstance

`func (o *DataSourceEventDto) SetDataSourceInstance(v DataSourceInstanceDto)`

SetDataSourceInstance sets DataSourceInstance field to given value.

### HasDataSourceInstance

`func (o *DataSourceEventDto) HasDataSourceInstance() bool`

HasDataSourceInstance returns a boolean if a field has been set.

### GetConnectionId

`func (o *DataSourceEventDto) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *DataSourceEventDto) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *DataSourceEventDto) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetConnection

`func (o *DataSourceEventDto) GetConnection() ConnectionDto`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *DataSourceEventDto) GetConnectionOk() (*ConnectionDto, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *DataSourceEventDto) SetConnection(v ConnectionDto)`

SetConnection sets Connection field to given value.

### HasConnection

`func (o *DataSourceEventDto) HasConnection() bool`

HasConnection returns a boolean if a field has been set.

### GetIntegrationId

`func (o *DataSourceEventDto) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *DataSourceEventDto) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *DataSourceEventDto) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.


### GetIntegration

`func (o *DataSourceEventDto) GetIntegration() IntegrationDto`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *DataSourceEventDto) GetIntegrationOk() (*IntegrationDto, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *DataSourceEventDto) SetIntegration(v IntegrationDto)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *DataSourceEventDto) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetEvent

`func (o *DataSourceEventDto) GetEvent() map[string]interface{}`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *DataSourceEventDto) GetEventOk() (*map[string]interface{}, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *DataSourceEventDto) SetEvent(v map[string]interface{})`

SetEvent sets Event field to given value.


### GetDatetime

`func (o *DataSourceEventDto) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *DataSourceEventDto) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *DataSourceEventDto) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetLaunchedFlowRunIds

`func (o *DataSourceEventDto) GetLaunchedFlowRunIds() []string`

GetLaunchedFlowRunIds returns the LaunchedFlowRunIds field if non-nil, zero value otherwise.

### GetLaunchedFlowRunIdsOk

`func (o *DataSourceEventDto) GetLaunchedFlowRunIdsOk() (*[]string, bool)`

GetLaunchedFlowRunIdsOk returns a tuple with the LaunchedFlowRunIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchedFlowRunIds

`func (o *DataSourceEventDto) SetLaunchedFlowRunIds(v []string)`

SetLaunchedFlowRunIds sets LaunchedFlowRunIds field to given value.


### GetError

`func (o *DataSourceEventDto) GetError() map[string]interface{}`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *DataSourceEventDto) GetErrorOk() (*map[string]interface{}, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *DataSourceEventDto) SetError(v map[string]interface{})`

SetError sets Error field to given value.

### HasError

`func (o *DataSourceEventDto) HasError() bool`

HasError returns a boolean if a field has been set.

### GetLogs

`func (o *DataSourceEventDto) GetLogs() map[string]interface{}`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *DataSourceEventDto) GetLogsOk() (*map[string]interface{}, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *DataSourceEventDto) SetLogs(v map[string]interface{})`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *DataSourceEventDto) HasLogs() bool`

HasLogs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


