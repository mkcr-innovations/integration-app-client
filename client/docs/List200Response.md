# List200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]ExternalEventPullDto**](ExternalEventPullDto.md) |  | 
**Cursor** | Pointer to **string** |  | [optional] 

## Methods

### NewList200Response

`func NewList200Response(items []ExternalEventPullDto, ) *List200Response`

NewList200Response instantiates a new List200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewList200ResponseWithDefaults

`func NewList200ResponseWithDefaults() *List200Response`

NewList200ResponseWithDefaults instantiates a new List200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *List200Response) GetItems() []ExternalEventPullDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *List200Response) GetItemsOk() (*[]ExternalEventPullDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *List200Response) SetItems(v []ExternalEventPullDto)`

SetItems sets Items field to given value.


### GetCursor

`func (o *List200Response) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *List200Response) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *List200Response) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *List200Response) HasCursor() bool`

HasCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


