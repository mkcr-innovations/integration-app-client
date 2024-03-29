# ListActions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]ActionDto**](ActionDto.md) |  | 
**Cursor** | Pointer to **string** |  | [optional] 

## Methods

### NewListActions200Response

`func NewListActions200Response(items []ActionDto, ) *ListActions200Response`

NewListActions200Response instantiates a new ListActions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListActions200ResponseWithDefaults

`func NewListActions200ResponseWithDefaults() *ListActions200Response`

NewListActions200ResponseWithDefaults instantiates a new ListActions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ListActions200Response) GetItems() []ActionDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListActions200Response) GetItemsOk() (*[]ActionDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListActions200Response) SetItems(v []ActionDto)`

SetItems sets Items field to given value.


### GetCursor

`func (o *ListActions200Response) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *ListActions200Response) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *ListActions200Response) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *ListActions200Response) HasCursor() bool`

HasCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


