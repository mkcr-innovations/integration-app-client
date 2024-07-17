# FieldMappingsControllerList200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | Pointer to **string** |  | [optional] 
**Items** | [**[]FieldMappingDto**](FieldMappingDto.md) |  | 

## Methods

### NewFieldMappingsControllerList200Response

`func NewFieldMappingsControllerList200Response(items []FieldMappingDto, ) *FieldMappingsControllerList200Response`

NewFieldMappingsControllerList200Response instantiates a new FieldMappingsControllerList200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldMappingsControllerList200ResponseWithDefaults

`func NewFieldMappingsControllerList200ResponseWithDefaults() *FieldMappingsControllerList200Response`

NewFieldMappingsControllerList200ResponseWithDefaults instantiates a new FieldMappingsControllerList200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursor

`func (o *FieldMappingsControllerList200Response) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *FieldMappingsControllerList200Response) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *FieldMappingsControllerList200Response) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *FieldMappingsControllerList200Response) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetItems

`func (o *FieldMappingsControllerList200Response) GetItems() []FieldMappingDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *FieldMappingsControllerList200Response) GetItemsOk() (*[]FieldMappingDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *FieldMappingsControllerList200Response) SetItems(v []FieldMappingDto)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


