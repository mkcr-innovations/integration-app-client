# ExternalEventPullDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**UserId** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**CustomerDto**](CustomerDto.md) |  | [optional] 
**ExternalEventSubscriptionId** | Pointer to **string** |  | [optional] 
**ExternalEventSubscription** | Pointer to **map[string]interface{}** |  | [optional] 
**ConnectionId** | Pointer to **string** |  | [optional] 
**Connection** | Pointer to [**ConnectionDto**](ConnectionDto.md) |  | [optional] 
**IntegrationId** | Pointer to **string** |  | [optional] 
**Integration** | Pointer to [**IntegrationDto**](IntegrationDto.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StartDatetime** | Pointer to **string** |  | [optional] 
**EndDatetime** | Pointer to **string** |  | [optional] 
**IsFullScan** | Pointer to **bool** |  | [optional] 
**CollectedEventIds** | Pointer to **[]string** |  | [optional] 
**Error** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewExternalEventPullDto

`func NewExternalEventPullDto(id string, ) *ExternalEventPullDto`

NewExternalEventPullDto instantiates a new ExternalEventPullDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalEventPullDtoWithDefaults

`func NewExternalEventPullDtoWithDefaults() *ExternalEventPullDto`

NewExternalEventPullDtoWithDefaults instantiates a new ExternalEventPullDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExternalEventPullDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalEventPullDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalEventPullDto) SetId(v string)`

SetId sets Id field to given value.


### GetUserId

`func (o *ExternalEventPullDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ExternalEventPullDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ExternalEventPullDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ExternalEventPullDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUser

`func (o *ExternalEventPullDto) GetUser() CustomerDto`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ExternalEventPullDto) GetUserOk() (*CustomerDto, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ExternalEventPullDto) SetUser(v CustomerDto)`

SetUser sets User field to given value.

### HasUser

`func (o *ExternalEventPullDto) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetExternalEventSubscriptionId

`func (o *ExternalEventPullDto) GetExternalEventSubscriptionId() string`

GetExternalEventSubscriptionId returns the ExternalEventSubscriptionId field if non-nil, zero value otherwise.

### GetExternalEventSubscriptionIdOk

`func (o *ExternalEventPullDto) GetExternalEventSubscriptionIdOk() (*string, bool)`

GetExternalEventSubscriptionIdOk returns a tuple with the ExternalEventSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalEventSubscriptionId

`func (o *ExternalEventPullDto) SetExternalEventSubscriptionId(v string)`

SetExternalEventSubscriptionId sets ExternalEventSubscriptionId field to given value.

### HasExternalEventSubscriptionId

`func (o *ExternalEventPullDto) HasExternalEventSubscriptionId() bool`

HasExternalEventSubscriptionId returns a boolean if a field has been set.

### GetExternalEventSubscription

`func (o *ExternalEventPullDto) GetExternalEventSubscription() map[string]interface{}`

GetExternalEventSubscription returns the ExternalEventSubscription field if non-nil, zero value otherwise.

### GetExternalEventSubscriptionOk

`func (o *ExternalEventPullDto) GetExternalEventSubscriptionOk() (*map[string]interface{}, bool)`

GetExternalEventSubscriptionOk returns a tuple with the ExternalEventSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalEventSubscription

`func (o *ExternalEventPullDto) SetExternalEventSubscription(v map[string]interface{})`

SetExternalEventSubscription sets ExternalEventSubscription field to given value.

### HasExternalEventSubscription

`func (o *ExternalEventPullDto) HasExternalEventSubscription() bool`

HasExternalEventSubscription returns a boolean if a field has been set.

### GetConnectionId

`func (o *ExternalEventPullDto) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *ExternalEventPullDto) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *ExternalEventPullDto) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *ExternalEventPullDto) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetConnection

`func (o *ExternalEventPullDto) GetConnection() ConnectionDto`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *ExternalEventPullDto) GetConnectionOk() (*ConnectionDto, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *ExternalEventPullDto) SetConnection(v ConnectionDto)`

SetConnection sets Connection field to given value.

### HasConnection

`func (o *ExternalEventPullDto) HasConnection() bool`

HasConnection returns a boolean if a field has been set.

### GetIntegrationId

`func (o *ExternalEventPullDto) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *ExternalEventPullDto) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *ExternalEventPullDto) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.

### HasIntegrationId

`func (o *ExternalEventPullDto) HasIntegrationId() bool`

HasIntegrationId returns a boolean if a field has been set.

### GetIntegration

`func (o *ExternalEventPullDto) GetIntegration() IntegrationDto`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *ExternalEventPullDto) GetIntegrationOk() (*IntegrationDto, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *ExternalEventPullDto) SetIntegration(v IntegrationDto)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *ExternalEventPullDto) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetStatus

`func (o *ExternalEventPullDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExternalEventPullDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExternalEventPullDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ExternalEventPullDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStartDatetime

`func (o *ExternalEventPullDto) GetStartDatetime() string`

GetStartDatetime returns the StartDatetime field if non-nil, zero value otherwise.

### GetStartDatetimeOk

`func (o *ExternalEventPullDto) GetStartDatetimeOk() (*string, bool)`

GetStartDatetimeOk returns a tuple with the StartDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDatetime

`func (o *ExternalEventPullDto) SetStartDatetime(v string)`

SetStartDatetime sets StartDatetime field to given value.

### HasStartDatetime

`func (o *ExternalEventPullDto) HasStartDatetime() bool`

HasStartDatetime returns a boolean if a field has been set.

### GetEndDatetime

`func (o *ExternalEventPullDto) GetEndDatetime() string`

GetEndDatetime returns the EndDatetime field if non-nil, zero value otherwise.

### GetEndDatetimeOk

`func (o *ExternalEventPullDto) GetEndDatetimeOk() (*string, bool)`

GetEndDatetimeOk returns a tuple with the EndDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDatetime

`func (o *ExternalEventPullDto) SetEndDatetime(v string)`

SetEndDatetime sets EndDatetime field to given value.

### HasEndDatetime

`func (o *ExternalEventPullDto) HasEndDatetime() bool`

HasEndDatetime returns a boolean if a field has been set.

### GetIsFullScan

`func (o *ExternalEventPullDto) GetIsFullScan() bool`

GetIsFullScan returns the IsFullScan field if non-nil, zero value otherwise.

### GetIsFullScanOk

`func (o *ExternalEventPullDto) GetIsFullScanOk() (*bool, bool)`

GetIsFullScanOk returns a tuple with the IsFullScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFullScan

`func (o *ExternalEventPullDto) SetIsFullScan(v bool)`

SetIsFullScan sets IsFullScan field to given value.

### HasIsFullScan

`func (o *ExternalEventPullDto) HasIsFullScan() bool`

HasIsFullScan returns a boolean if a field has been set.

### GetCollectedEventIds

`func (o *ExternalEventPullDto) GetCollectedEventIds() []string`

GetCollectedEventIds returns the CollectedEventIds field if non-nil, zero value otherwise.

### GetCollectedEventIdsOk

`func (o *ExternalEventPullDto) GetCollectedEventIdsOk() (*[]string, bool)`

GetCollectedEventIdsOk returns a tuple with the CollectedEventIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectedEventIds

`func (o *ExternalEventPullDto) SetCollectedEventIds(v []string)`

SetCollectedEventIds sets CollectedEventIds field to given value.

### HasCollectedEventIds

`func (o *ExternalEventPullDto) HasCollectedEventIds() bool`

HasCollectedEventIds returns a boolean if a field has been set.

### GetError

`func (o *ExternalEventPullDto) GetError() map[string]interface{}`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ExternalEventPullDto) GetErrorOk() (*map[string]interface{}, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ExternalEventPullDto) SetError(v map[string]interface{})`

SetError sets Error field to given value.

### HasError

`func (o *ExternalEventPullDto) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


