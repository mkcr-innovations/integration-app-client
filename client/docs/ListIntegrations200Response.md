# ListIntegrations200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]IntegrationDto**](IntegrationDto.md) |  | 
**Cursor** | Pointer to **string** |  | [optional] 

## Methods

### NewListIntegrations200Response

`func NewListIntegrations200Response(items []IntegrationDto, ) *ListIntegrations200Response`

NewListIntegrations200Response instantiates a new ListIntegrations200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListIntegrations200ResponseWithDefaults

`func NewListIntegrations200ResponseWithDefaults() *ListIntegrations200Response`

NewListIntegrations200ResponseWithDefaults instantiates a new ListIntegrations200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ListIntegrations200Response) GetItems() []IntegrationDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListIntegrations200Response) GetItemsOk() (*[]IntegrationDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListIntegrations200Response) SetItems(v []IntegrationDto)`

SetItems sets Items field to given value.


### GetCursor

`func (o *ListIntegrations200Response) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *ListIntegrations200Response) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *ListIntegrations200Response) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *ListIntegrations200Response) HasCursor() bool`

HasCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


