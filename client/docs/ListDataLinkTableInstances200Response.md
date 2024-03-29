# ListDataLinkTableInstances200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]DataLinkTableInstanceDto**](DataLinkTableInstanceDto.md) |  | 
**Cursor** | Pointer to **string** |  | [optional] 

## Methods

### NewListDataLinkTableInstances200Response

`func NewListDataLinkTableInstances200Response(items []DataLinkTableInstanceDto, ) *ListDataLinkTableInstances200Response`

NewListDataLinkTableInstances200Response instantiates a new ListDataLinkTableInstances200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDataLinkTableInstances200ResponseWithDefaults

`func NewListDataLinkTableInstances200ResponseWithDefaults() *ListDataLinkTableInstances200Response`

NewListDataLinkTableInstances200ResponseWithDefaults instantiates a new ListDataLinkTableInstances200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ListDataLinkTableInstances200Response) GetItems() []DataLinkTableInstanceDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListDataLinkTableInstances200Response) GetItemsOk() (*[]DataLinkTableInstanceDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListDataLinkTableInstances200Response) SetItems(v []DataLinkTableInstanceDto)`

SetItems sets Items field to given value.


### GetCursor

`func (o *ListDataLinkTableInstances200Response) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *ListDataLinkTableInstances200Response) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *ListDataLinkTableInstances200Response) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *ListDataLinkTableInstances200Response) HasCursor() bool`

HasCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


