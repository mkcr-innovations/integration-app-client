# FlowInstancesControllerListFlowInstances200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | Pointer to **string** |  | [optional] 
**Items** | [**[]FlowInstanceDto**](FlowInstanceDto.md) |  | 

## Methods

### NewFlowInstancesControllerListFlowInstances200Response

`func NewFlowInstancesControllerListFlowInstances200Response(items []FlowInstanceDto, ) *FlowInstancesControllerListFlowInstances200Response`

NewFlowInstancesControllerListFlowInstances200Response instantiates a new FlowInstancesControllerListFlowInstances200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowInstancesControllerListFlowInstances200ResponseWithDefaults

`func NewFlowInstancesControllerListFlowInstances200ResponseWithDefaults() *FlowInstancesControllerListFlowInstances200Response`

NewFlowInstancesControllerListFlowInstances200ResponseWithDefaults instantiates a new FlowInstancesControllerListFlowInstances200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursor

`func (o *FlowInstancesControllerListFlowInstances200Response) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *FlowInstancesControllerListFlowInstances200Response) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *FlowInstancesControllerListFlowInstances200Response) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *FlowInstancesControllerListFlowInstances200Response) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetItems

`func (o *FlowInstancesControllerListFlowInstances200Response) GetItems() []FlowInstanceDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *FlowInstancesControllerListFlowInstances200Response) GetItemsOk() (*[]FlowInstanceDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *FlowInstancesControllerListFlowInstances200Response) SetItems(v []FlowInstanceDto)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


