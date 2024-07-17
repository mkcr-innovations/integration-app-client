# ActionsControllerListActions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | Pointer to **string** |  | [optional] 
**Items** | [**[]ActionDto**](ActionDto.md) |  | 

## Methods

### NewActionsControllerListActions200Response

`func NewActionsControllerListActions200Response(items []ActionDto, ) *ActionsControllerListActions200Response`

NewActionsControllerListActions200Response instantiates a new ActionsControllerListActions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionsControllerListActions200ResponseWithDefaults

`func NewActionsControllerListActions200ResponseWithDefaults() *ActionsControllerListActions200Response`

NewActionsControllerListActions200ResponseWithDefaults instantiates a new ActionsControllerListActions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursor

`func (o *ActionsControllerListActions200Response) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *ActionsControllerListActions200Response) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *ActionsControllerListActions200Response) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *ActionsControllerListActions200Response) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetItems

`func (o *ActionsControllerListActions200Response) GetItems() []ActionDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ActionsControllerListActions200Response) GetItemsOk() (*[]ActionDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ActionsControllerListActions200Response) SetItems(v []ActionDto)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


