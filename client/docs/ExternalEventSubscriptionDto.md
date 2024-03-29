# ExternalEventSubscriptionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**UserId** | **string** |  | 
**User** | Pointer to [**CustomerDto**](CustomerDto.md) |  | [optional] 
**ConnectionId** | **string** |  | 
**Connection** | Pointer to [**ConnectionDto**](ConnectionDto.md) |  | [optional] 
**IntegrationId** | **string** |  | 
**Integration** | Pointer to [**IntegrationDto**](IntegrationDto.md) |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**IsRealTime** | Pointer to **bool** |  | [optional] 
**RequiresPull** | Pointer to **bool** |  | [optional] 
**RequiresFullSync** | Pointer to **bool** |  | [optional] 
**StateData** | Pointer to **map[string]interface{}** |  | [optional] 
**PullUpdatesIntervalSeconds** | Pointer to **float32** |  | [optional] 
**FullSyncIntervalSeconds** | Pointer to **float32** |  | [optional] 
**NextPullEventsTimestamp** | Pointer to **float32** |  | [optional] 
**NextRefreshTimestamp** | Pointer to **float32** |  | [optional] 
**Error** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**ArchivedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewExternalEventSubscriptionDto

`func NewExternalEventSubscriptionDto(id string, userId string, connectionId string, integrationId string, ) *ExternalEventSubscriptionDto`

NewExternalEventSubscriptionDto instantiates a new ExternalEventSubscriptionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalEventSubscriptionDtoWithDefaults

`func NewExternalEventSubscriptionDtoWithDefaults() *ExternalEventSubscriptionDto`

NewExternalEventSubscriptionDtoWithDefaults instantiates a new ExternalEventSubscriptionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExternalEventSubscriptionDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalEventSubscriptionDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalEventSubscriptionDto) SetId(v string)`

SetId sets Id field to given value.


### GetUserId

`func (o *ExternalEventSubscriptionDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ExternalEventSubscriptionDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ExternalEventSubscriptionDto) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetUser

`func (o *ExternalEventSubscriptionDto) GetUser() CustomerDto`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ExternalEventSubscriptionDto) GetUserOk() (*CustomerDto, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ExternalEventSubscriptionDto) SetUser(v CustomerDto)`

SetUser sets User field to given value.

### HasUser

`func (o *ExternalEventSubscriptionDto) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetConnectionId

`func (o *ExternalEventSubscriptionDto) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *ExternalEventSubscriptionDto) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *ExternalEventSubscriptionDto) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetConnection

`func (o *ExternalEventSubscriptionDto) GetConnection() ConnectionDto`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *ExternalEventSubscriptionDto) GetConnectionOk() (*ConnectionDto, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *ExternalEventSubscriptionDto) SetConnection(v ConnectionDto)`

SetConnection sets Connection field to given value.

### HasConnection

`func (o *ExternalEventSubscriptionDto) HasConnection() bool`

HasConnection returns a boolean if a field has been set.

### GetIntegrationId

`func (o *ExternalEventSubscriptionDto) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *ExternalEventSubscriptionDto) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *ExternalEventSubscriptionDto) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.


### GetIntegration

`func (o *ExternalEventSubscriptionDto) GetIntegration() IntegrationDto`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *ExternalEventSubscriptionDto) GetIntegrationOk() (*IntegrationDto, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *ExternalEventSubscriptionDto) SetIntegration(v IntegrationDto)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *ExternalEventSubscriptionDto) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetConfig

`func (o *ExternalEventSubscriptionDto) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ExternalEventSubscriptionDto) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ExternalEventSubscriptionDto) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ExternalEventSubscriptionDto) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetStatus

`func (o *ExternalEventSubscriptionDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExternalEventSubscriptionDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExternalEventSubscriptionDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ExternalEventSubscriptionDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetIsRealTime

`func (o *ExternalEventSubscriptionDto) GetIsRealTime() bool`

GetIsRealTime returns the IsRealTime field if non-nil, zero value otherwise.

### GetIsRealTimeOk

`func (o *ExternalEventSubscriptionDto) GetIsRealTimeOk() (*bool, bool)`

GetIsRealTimeOk returns a tuple with the IsRealTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRealTime

`func (o *ExternalEventSubscriptionDto) SetIsRealTime(v bool)`

SetIsRealTime sets IsRealTime field to given value.

### HasIsRealTime

`func (o *ExternalEventSubscriptionDto) HasIsRealTime() bool`

HasIsRealTime returns a boolean if a field has been set.

### GetRequiresPull

`func (o *ExternalEventSubscriptionDto) GetRequiresPull() bool`

GetRequiresPull returns the RequiresPull field if non-nil, zero value otherwise.

### GetRequiresPullOk

`func (o *ExternalEventSubscriptionDto) GetRequiresPullOk() (*bool, bool)`

GetRequiresPullOk returns a tuple with the RequiresPull field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresPull

`func (o *ExternalEventSubscriptionDto) SetRequiresPull(v bool)`

SetRequiresPull sets RequiresPull field to given value.

### HasRequiresPull

`func (o *ExternalEventSubscriptionDto) HasRequiresPull() bool`

HasRequiresPull returns a boolean if a field has been set.

### GetRequiresFullSync

`func (o *ExternalEventSubscriptionDto) GetRequiresFullSync() bool`

GetRequiresFullSync returns the RequiresFullSync field if non-nil, zero value otherwise.

### GetRequiresFullSyncOk

`func (o *ExternalEventSubscriptionDto) GetRequiresFullSyncOk() (*bool, bool)`

GetRequiresFullSyncOk returns a tuple with the RequiresFullSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresFullSync

`func (o *ExternalEventSubscriptionDto) SetRequiresFullSync(v bool)`

SetRequiresFullSync sets RequiresFullSync field to given value.

### HasRequiresFullSync

`func (o *ExternalEventSubscriptionDto) HasRequiresFullSync() bool`

HasRequiresFullSync returns a boolean if a field has been set.

### GetStateData

`func (o *ExternalEventSubscriptionDto) GetStateData() map[string]interface{}`

GetStateData returns the StateData field if non-nil, zero value otherwise.

### GetStateDataOk

`func (o *ExternalEventSubscriptionDto) GetStateDataOk() (*map[string]interface{}, bool)`

GetStateDataOk returns a tuple with the StateData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateData

`func (o *ExternalEventSubscriptionDto) SetStateData(v map[string]interface{})`

SetStateData sets StateData field to given value.

### HasStateData

`func (o *ExternalEventSubscriptionDto) HasStateData() bool`

HasStateData returns a boolean if a field has been set.

### GetPullUpdatesIntervalSeconds

`func (o *ExternalEventSubscriptionDto) GetPullUpdatesIntervalSeconds() float32`

GetPullUpdatesIntervalSeconds returns the PullUpdatesIntervalSeconds field if non-nil, zero value otherwise.

### GetPullUpdatesIntervalSecondsOk

`func (o *ExternalEventSubscriptionDto) GetPullUpdatesIntervalSecondsOk() (*float32, bool)`

GetPullUpdatesIntervalSecondsOk returns a tuple with the PullUpdatesIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullUpdatesIntervalSeconds

`func (o *ExternalEventSubscriptionDto) SetPullUpdatesIntervalSeconds(v float32)`

SetPullUpdatesIntervalSeconds sets PullUpdatesIntervalSeconds field to given value.

### HasPullUpdatesIntervalSeconds

`func (o *ExternalEventSubscriptionDto) HasPullUpdatesIntervalSeconds() bool`

HasPullUpdatesIntervalSeconds returns a boolean if a field has been set.

### GetFullSyncIntervalSeconds

`func (o *ExternalEventSubscriptionDto) GetFullSyncIntervalSeconds() float32`

GetFullSyncIntervalSeconds returns the FullSyncIntervalSeconds field if non-nil, zero value otherwise.

### GetFullSyncIntervalSecondsOk

`func (o *ExternalEventSubscriptionDto) GetFullSyncIntervalSecondsOk() (*float32, bool)`

GetFullSyncIntervalSecondsOk returns a tuple with the FullSyncIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullSyncIntervalSeconds

`func (o *ExternalEventSubscriptionDto) SetFullSyncIntervalSeconds(v float32)`

SetFullSyncIntervalSeconds sets FullSyncIntervalSeconds field to given value.

### HasFullSyncIntervalSeconds

`func (o *ExternalEventSubscriptionDto) HasFullSyncIntervalSeconds() bool`

HasFullSyncIntervalSeconds returns a boolean if a field has been set.

### GetNextPullEventsTimestamp

`func (o *ExternalEventSubscriptionDto) GetNextPullEventsTimestamp() float32`

GetNextPullEventsTimestamp returns the NextPullEventsTimestamp field if non-nil, zero value otherwise.

### GetNextPullEventsTimestampOk

`func (o *ExternalEventSubscriptionDto) GetNextPullEventsTimestampOk() (*float32, bool)`

GetNextPullEventsTimestampOk returns a tuple with the NextPullEventsTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPullEventsTimestamp

`func (o *ExternalEventSubscriptionDto) SetNextPullEventsTimestamp(v float32)`

SetNextPullEventsTimestamp sets NextPullEventsTimestamp field to given value.

### HasNextPullEventsTimestamp

`func (o *ExternalEventSubscriptionDto) HasNextPullEventsTimestamp() bool`

HasNextPullEventsTimestamp returns a boolean if a field has been set.

### GetNextRefreshTimestamp

`func (o *ExternalEventSubscriptionDto) GetNextRefreshTimestamp() float32`

GetNextRefreshTimestamp returns the NextRefreshTimestamp field if non-nil, zero value otherwise.

### GetNextRefreshTimestampOk

`func (o *ExternalEventSubscriptionDto) GetNextRefreshTimestampOk() (*float32, bool)`

GetNextRefreshTimestampOk returns a tuple with the NextRefreshTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRefreshTimestamp

`func (o *ExternalEventSubscriptionDto) SetNextRefreshTimestamp(v float32)`

SetNextRefreshTimestamp sets NextRefreshTimestamp field to given value.

### HasNextRefreshTimestamp

`func (o *ExternalEventSubscriptionDto) HasNextRefreshTimestamp() bool`

HasNextRefreshTimestamp returns a boolean if a field has been set.

### GetError

`func (o *ExternalEventSubscriptionDto) GetError() map[string]interface{}`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ExternalEventSubscriptionDto) GetErrorOk() (*map[string]interface{}, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ExternalEventSubscriptionDto) SetError(v map[string]interface{})`

SetError sets Error field to given value.

### HasError

`func (o *ExternalEventSubscriptionDto) HasError() bool`

HasError returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ExternalEventSubscriptionDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ExternalEventSubscriptionDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ExternalEventSubscriptionDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ExternalEventSubscriptionDto) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetArchivedAt

`func (o *ExternalEventSubscriptionDto) GetArchivedAt() string`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *ExternalEventSubscriptionDto) GetArchivedAtOk() (*string, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *ExternalEventSubscriptionDto) SetArchivedAt(v string)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *ExternalEventSubscriptionDto) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


