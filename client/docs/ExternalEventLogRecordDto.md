# ExternalEventLogRecordDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**UserId** | **string** |  | 
**User** | Pointer to [**CustomerDto**](CustomerDto.md) |  | [optional] 
**ExternalEventId** | **string** |  | 
**ExternalEventSubscriptionId** | **string** |  | 
**ConnectionId** | **string** |  | 
**Connection** | Pointer to [**ConnectionDto**](ConnectionDto.md) |  | [optional] 
**IntegrationId** | **string** |  | 
**Integration** | Pointer to [**IntegrationDto**](IntegrationDto.md) |  | [optional] 
**Payload** | **map[string]interface{}** |  | 
**LaunchedFlowRunIds** | **[]string** |  | 
**Error** | Pointer to **map[string]interface{}** |  | [optional] 
**Logs** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewExternalEventLogRecordDto

`func NewExternalEventLogRecordDto(id string, userId string, externalEventId string, externalEventSubscriptionId string, connectionId string, integrationId string, payload map[string]interface{}, launchedFlowRunIds []string, ) *ExternalEventLogRecordDto`

NewExternalEventLogRecordDto instantiates a new ExternalEventLogRecordDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalEventLogRecordDtoWithDefaults

`func NewExternalEventLogRecordDtoWithDefaults() *ExternalEventLogRecordDto`

NewExternalEventLogRecordDtoWithDefaults instantiates a new ExternalEventLogRecordDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExternalEventLogRecordDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalEventLogRecordDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalEventLogRecordDto) SetId(v string)`

SetId sets Id field to given value.


### GetUserId

`func (o *ExternalEventLogRecordDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ExternalEventLogRecordDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ExternalEventLogRecordDto) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetUser

`func (o *ExternalEventLogRecordDto) GetUser() CustomerDto`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ExternalEventLogRecordDto) GetUserOk() (*CustomerDto, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ExternalEventLogRecordDto) SetUser(v CustomerDto)`

SetUser sets User field to given value.

### HasUser

`func (o *ExternalEventLogRecordDto) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetExternalEventId

`func (o *ExternalEventLogRecordDto) GetExternalEventId() string`

GetExternalEventId returns the ExternalEventId field if non-nil, zero value otherwise.

### GetExternalEventIdOk

`func (o *ExternalEventLogRecordDto) GetExternalEventIdOk() (*string, bool)`

GetExternalEventIdOk returns a tuple with the ExternalEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalEventId

`func (o *ExternalEventLogRecordDto) SetExternalEventId(v string)`

SetExternalEventId sets ExternalEventId field to given value.


### GetExternalEventSubscriptionId

`func (o *ExternalEventLogRecordDto) GetExternalEventSubscriptionId() string`

GetExternalEventSubscriptionId returns the ExternalEventSubscriptionId field if non-nil, zero value otherwise.

### GetExternalEventSubscriptionIdOk

`func (o *ExternalEventLogRecordDto) GetExternalEventSubscriptionIdOk() (*string, bool)`

GetExternalEventSubscriptionIdOk returns a tuple with the ExternalEventSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalEventSubscriptionId

`func (o *ExternalEventLogRecordDto) SetExternalEventSubscriptionId(v string)`

SetExternalEventSubscriptionId sets ExternalEventSubscriptionId field to given value.


### GetConnectionId

`func (o *ExternalEventLogRecordDto) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *ExternalEventLogRecordDto) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *ExternalEventLogRecordDto) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetConnection

`func (o *ExternalEventLogRecordDto) GetConnection() ConnectionDto`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *ExternalEventLogRecordDto) GetConnectionOk() (*ConnectionDto, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *ExternalEventLogRecordDto) SetConnection(v ConnectionDto)`

SetConnection sets Connection field to given value.

### HasConnection

`func (o *ExternalEventLogRecordDto) HasConnection() bool`

HasConnection returns a boolean if a field has been set.

### GetIntegrationId

`func (o *ExternalEventLogRecordDto) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *ExternalEventLogRecordDto) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *ExternalEventLogRecordDto) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.


### GetIntegration

`func (o *ExternalEventLogRecordDto) GetIntegration() IntegrationDto`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *ExternalEventLogRecordDto) GetIntegrationOk() (*IntegrationDto, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *ExternalEventLogRecordDto) SetIntegration(v IntegrationDto)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *ExternalEventLogRecordDto) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetPayload

`func (o *ExternalEventLogRecordDto) GetPayload() map[string]interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *ExternalEventLogRecordDto) GetPayloadOk() (*map[string]interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *ExternalEventLogRecordDto) SetPayload(v map[string]interface{})`

SetPayload sets Payload field to given value.


### GetLaunchedFlowRunIds

`func (o *ExternalEventLogRecordDto) GetLaunchedFlowRunIds() []string`

GetLaunchedFlowRunIds returns the LaunchedFlowRunIds field if non-nil, zero value otherwise.

### GetLaunchedFlowRunIdsOk

`func (o *ExternalEventLogRecordDto) GetLaunchedFlowRunIdsOk() (*[]string, bool)`

GetLaunchedFlowRunIdsOk returns a tuple with the LaunchedFlowRunIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchedFlowRunIds

`func (o *ExternalEventLogRecordDto) SetLaunchedFlowRunIds(v []string)`

SetLaunchedFlowRunIds sets LaunchedFlowRunIds field to given value.


### GetError

`func (o *ExternalEventLogRecordDto) GetError() map[string]interface{}`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ExternalEventLogRecordDto) GetErrorOk() (*map[string]interface{}, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ExternalEventLogRecordDto) SetError(v map[string]interface{})`

SetError sets Error field to given value.

### HasError

`func (o *ExternalEventLogRecordDto) HasError() bool`

HasError returns a boolean if a field has been set.

### GetLogs

`func (o *ExternalEventLogRecordDto) GetLogs() map[string]interface{}`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *ExternalEventLogRecordDto) GetLogsOk() (*map[string]interface{}, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *ExternalEventLogRecordDto) SetLogs(v map[string]interface{})`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *ExternalEventLogRecordDto) HasLogs() bool`

HasLogs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


