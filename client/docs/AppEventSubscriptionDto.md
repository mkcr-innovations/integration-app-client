# AppEventSubscriptionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Revision** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**CustomerDto**](CustomerDto.md) |  | [optional] 
**InstanceKey** | Pointer to **string** |  | [optional] 
**AppEventTypeId** | Pointer to **string** |  | [optional] 
**AppEventType** | Pointer to [**AppEventTypeDto**](AppEventTypeDto.md) |  | [optional] 
**Schema** | Pointer to **map[string]interface{}** |  | [optional] 
**IsSubscribed** | Pointer to **bool** |  | [optional] 
**WebhookUri** | Pointer to **string** |  | [optional] 
**SubscriptionRequest** | Pointer to **map[string]interface{}** |  | [optional] 
**SubscriptionResponse** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewAppEventSubscriptionDto

`func NewAppEventSubscriptionDto(id string, ) *AppEventSubscriptionDto`

NewAppEventSubscriptionDto instantiates a new AppEventSubscriptionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppEventSubscriptionDtoWithDefaults

`func NewAppEventSubscriptionDtoWithDefaults() *AppEventSubscriptionDto`

NewAppEventSubscriptionDtoWithDefaults instantiates a new AppEventSubscriptionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppEventSubscriptionDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppEventSubscriptionDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppEventSubscriptionDto) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AppEventSubscriptionDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppEventSubscriptionDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppEventSubscriptionDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AppEventSubscriptionDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRevision

`func (o *AppEventSubscriptionDto) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *AppEventSubscriptionDto) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *AppEventSubscriptionDto) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *AppEventSubscriptionDto) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetUserId

`func (o *AppEventSubscriptionDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AppEventSubscriptionDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AppEventSubscriptionDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *AppEventSubscriptionDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUser

`func (o *AppEventSubscriptionDto) GetUser() CustomerDto`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AppEventSubscriptionDto) GetUserOk() (*CustomerDto, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AppEventSubscriptionDto) SetUser(v CustomerDto)`

SetUser sets User field to given value.

### HasUser

`func (o *AppEventSubscriptionDto) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetInstanceKey

`func (o *AppEventSubscriptionDto) GetInstanceKey() string`

GetInstanceKey returns the InstanceKey field if non-nil, zero value otherwise.

### GetInstanceKeyOk

`func (o *AppEventSubscriptionDto) GetInstanceKeyOk() (*string, bool)`

GetInstanceKeyOk returns a tuple with the InstanceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceKey

`func (o *AppEventSubscriptionDto) SetInstanceKey(v string)`

SetInstanceKey sets InstanceKey field to given value.

### HasInstanceKey

`func (o *AppEventSubscriptionDto) HasInstanceKey() bool`

HasInstanceKey returns a boolean if a field has been set.

### GetAppEventTypeId

`func (o *AppEventSubscriptionDto) GetAppEventTypeId() string`

GetAppEventTypeId returns the AppEventTypeId field if non-nil, zero value otherwise.

### GetAppEventTypeIdOk

`func (o *AppEventSubscriptionDto) GetAppEventTypeIdOk() (*string, bool)`

GetAppEventTypeIdOk returns a tuple with the AppEventTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppEventTypeId

`func (o *AppEventSubscriptionDto) SetAppEventTypeId(v string)`

SetAppEventTypeId sets AppEventTypeId field to given value.

### HasAppEventTypeId

`func (o *AppEventSubscriptionDto) HasAppEventTypeId() bool`

HasAppEventTypeId returns a boolean if a field has been set.

### GetAppEventType

`func (o *AppEventSubscriptionDto) GetAppEventType() AppEventTypeDto`

GetAppEventType returns the AppEventType field if non-nil, zero value otherwise.

### GetAppEventTypeOk

`func (o *AppEventSubscriptionDto) GetAppEventTypeOk() (*AppEventTypeDto, bool)`

GetAppEventTypeOk returns a tuple with the AppEventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppEventType

`func (o *AppEventSubscriptionDto) SetAppEventType(v AppEventTypeDto)`

SetAppEventType sets AppEventType field to given value.

### HasAppEventType

`func (o *AppEventSubscriptionDto) HasAppEventType() bool`

HasAppEventType returns a boolean if a field has been set.

### GetSchema

`func (o *AppEventSubscriptionDto) GetSchema() map[string]interface{}`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *AppEventSubscriptionDto) GetSchemaOk() (*map[string]interface{}, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *AppEventSubscriptionDto) SetSchema(v map[string]interface{})`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *AppEventSubscriptionDto) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetIsSubscribed

`func (o *AppEventSubscriptionDto) GetIsSubscribed() bool`

GetIsSubscribed returns the IsSubscribed field if non-nil, zero value otherwise.

### GetIsSubscribedOk

`func (o *AppEventSubscriptionDto) GetIsSubscribedOk() (*bool, bool)`

GetIsSubscribedOk returns a tuple with the IsSubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubscribed

`func (o *AppEventSubscriptionDto) SetIsSubscribed(v bool)`

SetIsSubscribed sets IsSubscribed field to given value.

### HasIsSubscribed

`func (o *AppEventSubscriptionDto) HasIsSubscribed() bool`

HasIsSubscribed returns a boolean if a field has been set.

### GetWebhookUri

`func (o *AppEventSubscriptionDto) GetWebhookUri() string`

GetWebhookUri returns the WebhookUri field if non-nil, zero value otherwise.

### GetWebhookUriOk

`func (o *AppEventSubscriptionDto) GetWebhookUriOk() (*string, bool)`

GetWebhookUriOk returns a tuple with the WebhookUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUri

`func (o *AppEventSubscriptionDto) SetWebhookUri(v string)`

SetWebhookUri sets WebhookUri field to given value.

### HasWebhookUri

`func (o *AppEventSubscriptionDto) HasWebhookUri() bool`

HasWebhookUri returns a boolean if a field has been set.

### GetSubscriptionRequest

`func (o *AppEventSubscriptionDto) GetSubscriptionRequest() map[string]interface{}`

GetSubscriptionRequest returns the SubscriptionRequest field if non-nil, zero value otherwise.

### GetSubscriptionRequestOk

`func (o *AppEventSubscriptionDto) GetSubscriptionRequestOk() (*map[string]interface{}, bool)`

GetSubscriptionRequestOk returns a tuple with the SubscriptionRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionRequest

`func (o *AppEventSubscriptionDto) SetSubscriptionRequest(v map[string]interface{})`

SetSubscriptionRequest sets SubscriptionRequest field to given value.

### HasSubscriptionRequest

`func (o *AppEventSubscriptionDto) HasSubscriptionRequest() bool`

HasSubscriptionRequest returns a boolean if a field has been set.

### GetSubscriptionResponse

`func (o *AppEventSubscriptionDto) GetSubscriptionResponse() map[string]interface{}`

GetSubscriptionResponse returns the SubscriptionResponse field if non-nil, zero value otherwise.

### GetSubscriptionResponseOk

`func (o *AppEventSubscriptionDto) GetSubscriptionResponseOk() (*map[string]interface{}, bool)`

GetSubscriptionResponseOk returns a tuple with the SubscriptionResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionResponse

`func (o *AppEventSubscriptionDto) SetSubscriptionResponse(v map[string]interface{})`

SetSubscriptionResponse sets SubscriptionResponse field to given value.

### HasSubscriptionResponse

`func (o *AppEventSubscriptionDto) HasSubscriptionResponse() bool`

HasSubscriptionResponse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


