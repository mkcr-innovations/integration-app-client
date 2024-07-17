# ActionExportDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntegrationKeys** | Pointer to **[]string** |  | [optional] 
**Name** | **string** |  | 
**Key** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**InputSchema** | Pointer to **map[string]interface{}** |  | [optional] 
**OutputMapping** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewActionExportDto

`func NewActionExportDto(name string, ) *ActionExportDto`

NewActionExportDto instantiates a new ActionExportDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionExportDtoWithDefaults

`func NewActionExportDtoWithDefaults() *ActionExportDto`

NewActionExportDtoWithDefaults instantiates a new ActionExportDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntegrationKeys

`func (o *ActionExportDto) GetIntegrationKeys() []string`

GetIntegrationKeys returns the IntegrationKeys field if non-nil, zero value otherwise.

### GetIntegrationKeysOk

`func (o *ActionExportDto) GetIntegrationKeysOk() (*[]string, bool)`

GetIntegrationKeysOk returns a tuple with the IntegrationKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationKeys

`func (o *ActionExportDto) SetIntegrationKeys(v []string)`

SetIntegrationKeys sets IntegrationKeys field to given value.

### HasIntegrationKeys

`func (o *ActionExportDto) HasIntegrationKeys() bool`

HasIntegrationKeys returns a boolean if a field has been set.

### GetName

`func (o *ActionExportDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ActionExportDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ActionExportDto) SetName(v string)`

SetName sets Name field to given value.


### GetKey

`func (o *ActionExportDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ActionExportDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ActionExportDto) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ActionExportDto) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetType

`func (o *ActionExportDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ActionExportDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ActionExportDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ActionExportDto) HasType() bool`

HasType returns a boolean if a field has been set.

### GetConfig

`func (o *ActionExportDto) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ActionExportDto) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ActionExportDto) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ActionExportDto) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetInputSchema

`func (o *ActionExportDto) GetInputSchema() map[string]interface{}`

GetInputSchema returns the InputSchema field if non-nil, zero value otherwise.

### GetInputSchemaOk

`func (o *ActionExportDto) GetInputSchemaOk() (*map[string]interface{}, bool)`

GetInputSchemaOk returns a tuple with the InputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputSchema

`func (o *ActionExportDto) SetInputSchema(v map[string]interface{})`

SetInputSchema sets InputSchema field to given value.

### HasInputSchema

`func (o *ActionExportDto) HasInputSchema() bool`

HasInputSchema returns a boolean if a field has been set.

### GetOutputMapping

`func (o *ActionExportDto) GetOutputMapping() map[string]interface{}`

GetOutputMapping returns the OutputMapping field if non-nil, zero value otherwise.

### GetOutputMappingOk

`func (o *ActionExportDto) GetOutputMappingOk() (*map[string]interface{}, bool)`

GetOutputMappingOk returns a tuple with the OutputMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputMapping

`func (o *ActionExportDto) SetOutputMapping(v map[string]interface{})`

SetOutputMapping sets OutputMapping field to given value.

### HasOutputMapping

`func (o *ActionExportDto) HasOutputMapping() bool`

HasOutputMapping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


