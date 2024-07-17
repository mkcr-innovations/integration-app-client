# AppEventSubscriptionsControllerListAppEventSubscriptions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | Pointer to **string** |  | [optional] 
**Items** | [**[]AppEventSubscriptionDto**](AppEventSubscriptionDto.md) |  | 

## Methods

### NewAppEventSubscriptionsControllerListAppEventSubscriptions200Response

`func NewAppEventSubscriptionsControllerListAppEventSubscriptions200Response(items []AppEventSubscriptionDto, ) *AppEventSubscriptionsControllerListAppEventSubscriptions200Response`

NewAppEventSubscriptionsControllerListAppEventSubscriptions200Response instantiates a new AppEventSubscriptionsControllerListAppEventSubscriptions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppEventSubscriptionsControllerListAppEventSubscriptions200ResponseWithDefaults

`func NewAppEventSubscriptionsControllerListAppEventSubscriptions200ResponseWithDefaults() *AppEventSubscriptionsControllerListAppEventSubscriptions200Response`

NewAppEventSubscriptionsControllerListAppEventSubscriptions200ResponseWithDefaults instantiates a new AppEventSubscriptionsControllerListAppEventSubscriptions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursor

`func (o *AppEventSubscriptionsControllerListAppEventSubscriptions200Response) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *AppEventSubscriptionsControllerListAppEventSubscriptions200Response) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *AppEventSubscriptionsControllerListAppEventSubscriptions200Response) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *AppEventSubscriptionsControllerListAppEventSubscriptions200Response) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetItems

`func (o *AppEventSubscriptionsControllerListAppEventSubscriptions200Response) GetItems() []AppEventSubscriptionDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AppEventSubscriptionsControllerListAppEventSubscriptions200Response) GetItemsOk() (*[]AppEventSubscriptionDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AppEventSubscriptionsControllerListAppEventSubscriptions200Response) SetItems(v []AppEventSubscriptionDto)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


