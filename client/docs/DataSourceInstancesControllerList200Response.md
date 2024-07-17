# DataSourceInstancesControllerList200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | Pointer to **string** |  | [optional] 
**Items** | [**[]DataSourceInstanceDto**](DataSourceInstanceDto.md) |  | 

## Methods

### NewDataSourceInstancesControllerList200Response

`func NewDataSourceInstancesControllerList200Response(items []DataSourceInstanceDto, ) *DataSourceInstancesControllerList200Response`

NewDataSourceInstancesControllerList200Response instantiates a new DataSourceInstancesControllerList200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataSourceInstancesControllerList200ResponseWithDefaults

`func NewDataSourceInstancesControllerList200ResponseWithDefaults() *DataSourceInstancesControllerList200Response`

NewDataSourceInstancesControllerList200ResponseWithDefaults instantiates a new DataSourceInstancesControllerList200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursor

`func (o *DataSourceInstancesControllerList200Response) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *DataSourceInstancesControllerList200Response) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *DataSourceInstancesControllerList200Response) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *DataSourceInstancesControllerList200Response) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetItems

`func (o *DataSourceInstancesControllerList200Response) GetItems() []DataSourceInstanceDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *DataSourceInstancesControllerList200Response) GetItemsOk() (*[]DataSourceInstanceDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *DataSourceInstancesControllerList200Response) SetItems(v []DataSourceInstanceDto)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


