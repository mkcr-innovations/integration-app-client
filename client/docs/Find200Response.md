# Find200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]ExternalEventSubscriptionDto**](ExternalEventSubscriptionDto.md) |  | 
**Cursor** | Pointer to **string** |  | [optional] 

## Methods

### NewFind200Response

`func NewFind200Response(items []ExternalEventSubscriptionDto, ) *Find200Response`

NewFind200Response instantiates a new Find200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFind200ResponseWithDefaults

`func NewFind200ResponseWithDefaults() *Find200Response`

NewFind200ResponseWithDefaults instantiates a new Find200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *Find200Response) GetItems() []ExternalEventSubscriptionDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Find200Response) GetItemsOk() (*[]ExternalEventSubscriptionDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Find200Response) SetItems(v []ExternalEventSubscriptionDto)`

SetItems sets Items field to given value.


### GetCursor

`func (o *Find200Response) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *Find200Response) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *Find200Response) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *Find200Response) HasCursor() bool`

HasCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


