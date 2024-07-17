# FieldMappingInstancesControllerListFieldMappingInstances200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | Pointer to **string** |  | [optional] 
**Items** | [**[]FieldMappingInstanceDto**](FieldMappingInstanceDto.md) |  | 

## Methods

### NewFieldMappingInstancesControllerListFieldMappingInstances200Response

`func NewFieldMappingInstancesControllerListFieldMappingInstances200Response(items []FieldMappingInstanceDto, ) *FieldMappingInstancesControllerListFieldMappingInstances200Response`

NewFieldMappingInstancesControllerListFieldMappingInstances200Response instantiates a new FieldMappingInstancesControllerListFieldMappingInstances200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldMappingInstancesControllerListFieldMappingInstances200ResponseWithDefaults

`func NewFieldMappingInstancesControllerListFieldMappingInstances200ResponseWithDefaults() *FieldMappingInstancesControllerListFieldMappingInstances200Response`

NewFieldMappingInstancesControllerListFieldMappingInstances200ResponseWithDefaults instantiates a new FieldMappingInstancesControllerListFieldMappingInstances200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursor

`func (o *FieldMappingInstancesControllerListFieldMappingInstances200Response) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *FieldMappingInstancesControllerListFieldMappingInstances200Response) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *FieldMappingInstancesControllerListFieldMappingInstances200Response) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *FieldMappingInstancesControllerListFieldMappingInstances200Response) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetItems

`func (o *FieldMappingInstancesControllerListFieldMappingInstances200Response) GetItems() []FieldMappingInstanceDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *FieldMappingInstancesControllerListFieldMappingInstances200Response) GetItemsOk() (*[]FieldMappingInstanceDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *FieldMappingInstancesControllerListFieldMappingInstances200Response) SetItems(v []FieldMappingInstanceDto)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


