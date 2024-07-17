# AppDataSchemasControllerListAppDataSchemas200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | Pointer to **string** |  | [optional] 
**Items** | [**[]AppDataSchemaDto**](AppDataSchemaDto.md) |  | 

## Methods

### NewAppDataSchemasControllerListAppDataSchemas200Response

`func NewAppDataSchemasControllerListAppDataSchemas200Response(items []AppDataSchemaDto, ) *AppDataSchemasControllerListAppDataSchemas200Response`

NewAppDataSchemasControllerListAppDataSchemas200Response instantiates a new AppDataSchemasControllerListAppDataSchemas200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppDataSchemasControllerListAppDataSchemas200ResponseWithDefaults

`func NewAppDataSchemasControllerListAppDataSchemas200ResponseWithDefaults() *AppDataSchemasControllerListAppDataSchemas200Response`

NewAppDataSchemasControllerListAppDataSchemas200ResponseWithDefaults instantiates a new AppDataSchemasControllerListAppDataSchemas200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursor

`func (o *AppDataSchemasControllerListAppDataSchemas200Response) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *AppDataSchemasControllerListAppDataSchemas200Response) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *AppDataSchemasControllerListAppDataSchemas200Response) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *AppDataSchemasControllerListAppDataSchemas200Response) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetItems

`func (o *AppDataSchemasControllerListAppDataSchemas200Response) GetItems() []AppDataSchemaDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AppDataSchemasControllerListAppDataSchemas200Response) GetItemsOk() (*[]AppDataSchemaDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AppDataSchemasControllerListAppDataSchemas200Response) SetItems(v []AppDataSchemaDto)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


