# IntegrationsControllerListIntegrations200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | Pointer to **string** |  | [optional] 
**Items** | [**[]IntegrationDto**](IntegrationDto.md) |  | 

## Methods

### NewIntegrationsControllerListIntegrations200Response

`func NewIntegrationsControllerListIntegrations200Response(items []IntegrationDto, ) *IntegrationsControllerListIntegrations200Response`

NewIntegrationsControllerListIntegrations200Response instantiates a new IntegrationsControllerListIntegrations200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationsControllerListIntegrations200ResponseWithDefaults

`func NewIntegrationsControllerListIntegrations200ResponseWithDefaults() *IntegrationsControllerListIntegrations200Response`

NewIntegrationsControllerListIntegrations200ResponseWithDefaults instantiates a new IntegrationsControllerListIntegrations200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursor

`func (o *IntegrationsControllerListIntegrations200Response) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *IntegrationsControllerListIntegrations200Response) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *IntegrationsControllerListIntegrations200Response) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *IntegrationsControllerListIntegrations200Response) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetItems

`func (o *IntegrationsControllerListIntegrations200Response) GetItems() []IntegrationDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *IntegrationsControllerListIntegrations200Response) GetItemsOk() (*[]IntegrationDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *IntegrationsControllerListIntegrations200Response) SetItems(v []IntegrationDto)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


