# ListAppDataSchemaInstances200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]AppDataSchemaInstanceDto**](AppDataSchemaInstanceDto.md) |  | 
**Cursor** | Pointer to **string** |  | [optional] 

## Methods

### NewListAppDataSchemaInstances200Response

`func NewListAppDataSchemaInstances200Response(items []AppDataSchemaInstanceDto, ) *ListAppDataSchemaInstances200Response`

NewListAppDataSchemaInstances200Response instantiates a new ListAppDataSchemaInstances200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAppDataSchemaInstances200ResponseWithDefaults

`func NewListAppDataSchemaInstances200ResponseWithDefaults() *ListAppDataSchemaInstances200Response`

NewListAppDataSchemaInstances200ResponseWithDefaults instantiates a new ListAppDataSchemaInstances200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ListAppDataSchemaInstances200Response) GetItems() []AppDataSchemaInstanceDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListAppDataSchemaInstances200Response) GetItemsOk() (*[]AppDataSchemaInstanceDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListAppDataSchemaInstances200Response) SetItems(v []AppDataSchemaInstanceDto)`

SetItems sets Items field to given value.


### GetCursor

`func (o *ListAppDataSchemaInstances200Response) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *ListAppDataSchemaInstances200Response) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *ListAppDataSchemaInstances200Response) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *ListAppDataSchemaInstances200Response) HasCursor() bool`

HasCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


