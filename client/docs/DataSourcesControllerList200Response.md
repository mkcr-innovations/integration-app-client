# DataSourcesControllerList200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | Pointer to **string** |  | [optional] 
**Items** | [**[]DataSourceDto**](DataSourceDto.md) |  | 

## Methods

### NewDataSourcesControllerList200Response

`func NewDataSourcesControllerList200Response(items []DataSourceDto, ) *DataSourcesControllerList200Response`

NewDataSourcesControllerList200Response instantiates a new DataSourcesControllerList200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataSourcesControllerList200ResponseWithDefaults

`func NewDataSourcesControllerList200ResponseWithDefaults() *DataSourcesControllerList200Response`

NewDataSourcesControllerList200ResponseWithDefaults instantiates a new DataSourcesControllerList200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursor

`func (o *DataSourcesControllerList200Response) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *DataSourcesControllerList200Response) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *DataSourcesControllerList200Response) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *DataSourcesControllerList200Response) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetItems

`func (o *DataSourcesControllerList200Response) GetItems() []DataSourceDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *DataSourcesControllerList200Response) GetItemsOk() (*[]DataSourceDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *DataSourcesControllerList200Response) SetItems(v []DataSourceDto)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


