# ListFlows200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]FlowDto**](FlowDto.md) |  | 
**Cursor** | Pointer to **string** |  | [optional] 

## Methods

### NewListFlows200Response

`func NewListFlows200Response(items []FlowDto, ) *ListFlows200Response`

NewListFlows200Response instantiates a new ListFlows200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFlows200ResponseWithDefaults

`func NewListFlows200ResponseWithDefaults() *ListFlows200Response`

NewListFlows200ResponseWithDefaults instantiates a new ListFlows200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ListFlows200Response) GetItems() []FlowDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListFlows200Response) GetItemsOk() (*[]FlowDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListFlows200Response) SetItems(v []FlowDto)`

SetItems sets Items field to given value.


### GetCursor

`func (o *ListFlows200Response) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *ListFlows200Response) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *ListFlows200Response) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *ListFlows200Response) HasCursor() bool`

HasCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


