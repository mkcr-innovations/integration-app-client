# ListAppEventTypes200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]AppEventTypeDto**](AppEventTypeDto.md) |  | 
**Cursor** | Pointer to **string** |  | [optional] 

## Methods

### NewListAppEventTypes200Response

`func NewListAppEventTypes200Response(items []AppEventTypeDto, ) *ListAppEventTypes200Response`

NewListAppEventTypes200Response instantiates a new ListAppEventTypes200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAppEventTypes200ResponseWithDefaults

`func NewListAppEventTypes200ResponseWithDefaults() *ListAppEventTypes200Response`

NewListAppEventTypes200ResponseWithDefaults instantiates a new ListAppEventTypes200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ListAppEventTypes200Response) GetItems() []AppEventTypeDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListAppEventTypes200Response) GetItemsOk() (*[]AppEventTypeDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListAppEventTypes200Response) SetItems(v []AppEventTypeDto)`

SetItems sets Items field to given value.


### GetCursor

`func (o *ListAppEventTypes200Response) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *ListAppEventTypes200Response) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *ListAppEventTypes200Response) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *ListAppEventTypes200Response) HasCursor() bool`

HasCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


