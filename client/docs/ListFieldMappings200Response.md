# ListFieldMappings200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]FieldMappingDto**](FieldMappingDto.md) |  | 
**Cursor** | Pointer to **string** |  | [optional] 

## Methods

### NewListFieldMappings200Response

`func NewListFieldMappings200Response(items []FieldMappingDto, ) *ListFieldMappings200Response`

NewListFieldMappings200Response instantiates a new ListFieldMappings200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFieldMappings200ResponseWithDefaults

`func NewListFieldMappings200ResponseWithDefaults() *ListFieldMappings200Response`

NewListFieldMappings200ResponseWithDefaults instantiates a new ListFieldMappings200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ListFieldMappings200Response) GetItems() []FieldMappingDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListFieldMappings200Response) GetItemsOk() (*[]FieldMappingDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListFieldMappings200Response) SetItems(v []FieldMappingDto)`

SetItems sets Items field to given value.


### GetCursor

`func (o *ListFieldMappings200Response) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *ListFieldMappings200Response) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *ListFieldMappings200Response) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *ListFieldMappings200Response) HasCursor() bool`

HasCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


