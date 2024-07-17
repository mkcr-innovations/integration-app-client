# FlowsControllerListFlows200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | Pointer to **string** |  | [optional] 
**Items** | [**[]FlowDto**](FlowDto.md) |  | 

## Methods

### NewFlowsControllerListFlows200Response

`func NewFlowsControllerListFlows200Response(items []FlowDto, ) *FlowsControllerListFlows200Response`

NewFlowsControllerListFlows200Response instantiates a new FlowsControllerListFlows200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowsControllerListFlows200ResponseWithDefaults

`func NewFlowsControllerListFlows200ResponseWithDefaults() *FlowsControllerListFlows200Response`

NewFlowsControllerListFlows200ResponseWithDefaults instantiates a new FlowsControllerListFlows200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursor

`func (o *FlowsControllerListFlows200Response) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *FlowsControllerListFlows200Response) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *FlowsControllerListFlows200Response) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *FlowsControllerListFlows200Response) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetItems

`func (o *FlowsControllerListFlows200Response) GetItems() []FlowDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *FlowsControllerListFlows200Response) GetItemsOk() (*[]FlowDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *FlowsControllerListFlows200Response) SetItems(v []FlowDto)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


