# ListAppDataSchemas200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]AppDataSchemaDto**](AppDataSchemaDto.md) |  | 
**Cursor** | Pointer to **string** |  | [optional] 

## Methods

### NewListAppDataSchemas200Response

`func NewListAppDataSchemas200Response(items []AppDataSchemaDto, ) *ListAppDataSchemas200Response`

NewListAppDataSchemas200Response instantiates a new ListAppDataSchemas200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAppDataSchemas200ResponseWithDefaults

`func NewListAppDataSchemas200ResponseWithDefaults() *ListAppDataSchemas200Response`

NewListAppDataSchemas200ResponseWithDefaults instantiates a new ListAppDataSchemas200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ListAppDataSchemas200Response) GetItems() []AppDataSchemaDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListAppDataSchemas200Response) GetItemsOk() (*[]AppDataSchemaDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListAppDataSchemas200Response) SetItems(v []AppDataSchemaDto)`

SetItems sets Items field to given value.


### GetCursor

`func (o *ListAppDataSchemas200Response) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *ListAppDataSchemas200Response) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *ListAppDataSchemas200Response) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *ListAppDataSchemas200Response) HasCursor() bool`

HasCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


