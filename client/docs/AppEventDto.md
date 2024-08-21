# AppEventDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**UserId** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**CustomerDto**](CustomerDto.md) |  | [optional] 
**AppEventTypeId** | Pointer to **string** |  | [optional] 
**AppEventType** | Pointer to [**AppEventTypeDto**](AppEventTypeDto.md) |  | [optional] 
**AppEventSubscriptionId** | Pointer to **string** |  | [optional] 
**AppEventSubscription** | Pointer to [**AppEventSubscriptionDto**](AppEventSubscriptionDto.md) |  | [optional] 
**Event** | Pointer to **map[string]interface{}** |  | [optional] 
**Datetime** | Pointer to **string** |  | [optional] 
**LaunchedFlowRunIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAppEventDto

`func NewAppEventDto(id string, ) *AppEventDto`

NewAppEventDto instantiates a new AppEventDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppEventDtoWithDefaults

`func NewAppEventDtoWithDefaults() *AppEventDto`

NewAppEventDtoWithDefaults instantiates a new AppEventDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppEventDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppEventDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppEventDto) SetId(v string)`

SetId sets Id field to given value.


### GetUserId

`func (o *AppEventDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AppEventDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AppEventDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *AppEventDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUser

`func (o *AppEventDto) GetUser() CustomerDto`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AppEventDto) GetUserOk() (*CustomerDto, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AppEventDto) SetUser(v CustomerDto)`

SetUser sets User field to given value.

### HasUser

`func (o *AppEventDto) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetAppEventTypeId

`func (o *AppEventDto) GetAppEventTypeId() string`

GetAppEventTypeId returns the AppEventTypeId field if non-nil, zero value otherwise.

### GetAppEventTypeIdOk

`func (o *AppEventDto) GetAppEventTypeIdOk() (*string, bool)`

GetAppEventTypeIdOk returns a tuple with the AppEventTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppEventTypeId

`func (o *AppEventDto) SetAppEventTypeId(v string)`

SetAppEventTypeId sets AppEventTypeId field to given value.

### HasAppEventTypeId

`func (o *AppEventDto) HasAppEventTypeId() bool`

HasAppEventTypeId returns a boolean if a field has been set.

### GetAppEventType

`func (o *AppEventDto) GetAppEventType() AppEventTypeDto`

GetAppEventType returns the AppEventType field if non-nil, zero value otherwise.

### GetAppEventTypeOk

`func (o *AppEventDto) GetAppEventTypeOk() (*AppEventTypeDto, bool)`

GetAppEventTypeOk returns a tuple with the AppEventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppEventType

`func (o *AppEventDto) SetAppEventType(v AppEventTypeDto)`

SetAppEventType sets AppEventType field to given value.

### HasAppEventType

`func (o *AppEventDto) HasAppEventType() bool`

HasAppEventType returns a boolean if a field has been set.

### GetAppEventSubscriptionId

`func (o *AppEventDto) GetAppEventSubscriptionId() string`

GetAppEventSubscriptionId returns the AppEventSubscriptionId field if non-nil, zero value otherwise.

### GetAppEventSubscriptionIdOk

`func (o *AppEventDto) GetAppEventSubscriptionIdOk() (*string, bool)`

GetAppEventSubscriptionIdOk returns a tuple with the AppEventSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppEventSubscriptionId

`func (o *AppEventDto) SetAppEventSubscriptionId(v string)`

SetAppEventSubscriptionId sets AppEventSubscriptionId field to given value.

### HasAppEventSubscriptionId

`func (o *AppEventDto) HasAppEventSubscriptionId() bool`

HasAppEventSubscriptionId returns a boolean if a field has been set.

### GetAppEventSubscription

`func (o *AppEventDto) GetAppEventSubscription() AppEventSubscriptionDto`

GetAppEventSubscription returns the AppEventSubscription field if non-nil, zero value otherwise.

### GetAppEventSubscriptionOk

`func (o *AppEventDto) GetAppEventSubscriptionOk() (*AppEventSubscriptionDto, bool)`

GetAppEventSubscriptionOk returns a tuple with the AppEventSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppEventSubscription

`func (o *AppEventDto) SetAppEventSubscription(v AppEventSubscriptionDto)`

SetAppEventSubscription sets AppEventSubscription field to given value.

### HasAppEventSubscription

`func (o *AppEventDto) HasAppEventSubscription() bool`

HasAppEventSubscription returns a boolean if a field has been set.

### GetEvent

`func (o *AppEventDto) GetEvent() map[string]interface{}`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *AppEventDto) GetEventOk() (*map[string]interface{}, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *AppEventDto) SetEvent(v map[string]interface{})`

SetEvent sets Event field to given value.

### HasEvent

`func (o *AppEventDto) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetDatetime

`func (o *AppEventDto) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *AppEventDto) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *AppEventDto) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.

### HasDatetime

`func (o *AppEventDto) HasDatetime() bool`

HasDatetime returns a boolean if a field has been set.

### GetLaunchedFlowRunIds

`func (o *AppEventDto) GetLaunchedFlowRunIds() []string`

GetLaunchedFlowRunIds returns the LaunchedFlowRunIds field if non-nil, zero value otherwise.

### GetLaunchedFlowRunIdsOk

`func (o *AppEventDto) GetLaunchedFlowRunIdsOk() (*[]string, bool)`

GetLaunchedFlowRunIdsOk returns a tuple with the LaunchedFlowRunIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchedFlowRunIds

`func (o *AppEventDto) SetLaunchedFlowRunIds(v []string)`

SetLaunchedFlowRunIds sets LaunchedFlowRunIds field to given value.

### HasLaunchedFlowRunIds

`func (o *AppEventDto) HasLaunchedFlowRunIds() bool`

HasLaunchedFlowRunIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


