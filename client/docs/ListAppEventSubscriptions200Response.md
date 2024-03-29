# ListAppEventSubscriptions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]AppEventSubscriptionDto**](AppEventSubscriptionDto.md) |  | 
**Cursor** | Pointer to **string** |  | [optional] 

## Methods

### NewListAppEventSubscriptions200Response

`func NewListAppEventSubscriptions200Response(items []AppEventSubscriptionDto, ) *ListAppEventSubscriptions200Response`

NewListAppEventSubscriptions200Response instantiates a new ListAppEventSubscriptions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAppEventSubscriptions200ResponseWithDefaults

`func NewListAppEventSubscriptions200ResponseWithDefaults() *ListAppEventSubscriptions200Response`

NewListAppEventSubscriptions200ResponseWithDefaults instantiates a new ListAppEventSubscriptions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ListAppEventSubscriptions200Response) GetItems() []AppEventSubscriptionDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListAppEventSubscriptions200Response) GetItemsOk() (*[]AppEventSubscriptionDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListAppEventSubscriptions200Response) SetItems(v []AppEventSubscriptionDto)`

SetItems sets Items field to given value.


### GetCursor

`func (o *ListAppEventSubscriptions200Response) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *ListAppEventSubscriptions200Response) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *ListAppEventSubscriptions200Response) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *ListAppEventSubscriptions200Response) HasCursor() bool`

HasCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


