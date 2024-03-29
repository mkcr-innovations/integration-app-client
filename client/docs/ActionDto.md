# ActionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**InputSchema** | Pointer to **map[string]interface{}** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**OutputMapping** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomOutputSchema** | Pointer to **map[string]interface{}** |  | [optional] 
**DefaultOutputSchema** | Pointer to **map[string]interface{}** |  | [optional] 
**TransformedOutputSchema** | Pointer to **map[string]interface{}** |  | [optional] 
**OutputSchema** | Pointer to **map[string]interface{}** |  | [optional] 
**IsConfigurationMissing** | Pointer to **bool** |  | [optional] 
**Dependencies** | Pointer to **[]map[string]interface{}** |  | [optional] 
**IsDeployed** | Pointer to **bool** |  | [optional] 

## Methods

### NewActionDto

`func NewActionDto(type_ string, ) *ActionDto`

NewActionDto instantiates a new ActionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionDtoWithDefaults

`func NewActionDtoWithDefaults() *ActionDto`

NewActionDtoWithDefaults instantiates a new ActionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ActionDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ActionDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ActionDto) SetType(v string)`

SetType sets Type field to given value.


### GetInputSchema

`func (o *ActionDto) GetInputSchema() map[string]interface{}`

GetInputSchema returns the InputSchema field if non-nil, zero value otherwise.

### GetInputSchemaOk

`func (o *ActionDto) GetInputSchemaOk() (*map[string]interface{}, bool)`

GetInputSchemaOk returns a tuple with the InputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputSchema

`func (o *ActionDto) SetInputSchema(v map[string]interface{})`

SetInputSchema sets InputSchema field to given value.

### HasInputSchema

`func (o *ActionDto) HasInputSchema() bool`

HasInputSchema returns a boolean if a field has been set.

### GetConfig

`func (o *ActionDto) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ActionDto) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ActionDto) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ActionDto) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetOutputMapping

`func (o *ActionDto) GetOutputMapping() map[string]interface{}`

GetOutputMapping returns the OutputMapping field if non-nil, zero value otherwise.

### GetOutputMappingOk

`func (o *ActionDto) GetOutputMappingOk() (*map[string]interface{}, bool)`

GetOutputMappingOk returns a tuple with the OutputMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputMapping

`func (o *ActionDto) SetOutputMapping(v map[string]interface{})`

SetOutputMapping sets OutputMapping field to given value.

### HasOutputMapping

`func (o *ActionDto) HasOutputMapping() bool`

HasOutputMapping returns a boolean if a field has been set.

### GetCustomOutputSchema

`func (o *ActionDto) GetCustomOutputSchema() map[string]interface{}`

GetCustomOutputSchema returns the CustomOutputSchema field if non-nil, zero value otherwise.

### GetCustomOutputSchemaOk

`func (o *ActionDto) GetCustomOutputSchemaOk() (*map[string]interface{}, bool)`

GetCustomOutputSchemaOk returns a tuple with the CustomOutputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomOutputSchema

`func (o *ActionDto) SetCustomOutputSchema(v map[string]interface{})`

SetCustomOutputSchema sets CustomOutputSchema field to given value.

### HasCustomOutputSchema

`func (o *ActionDto) HasCustomOutputSchema() bool`

HasCustomOutputSchema returns a boolean if a field has been set.

### GetDefaultOutputSchema

`func (o *ActionDto) GetDefaultOutputSchema() map[string]interface{}`

GetDefaultOutputSchema returns the DefaultOutputSchema field if non-nil, zero value otherwise.

### GetDefaultOutputSchemaOk

`func (o *ActionDto) GetDefaultOutputSchemaOk() (*map[string]interface{}, bool)`

GetDefaultOutputSchemaOk returns a tuple with the DefaultOutputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOutputSchema

`func (o *ActionDto) SetDefaultOutputSchema(v map[string]interface{})`

SetDefaultOutputSchema sets DefaultOutputSchema field to given value.

### HasDefaultOutputSchema

`func (o *ActionDto) HasDefaultOutputSchema() bool`

HasDefaultOutputSchema returns a boolean if a field has been set.

### GetTransformedOutputSchema

`func (o *ActionDto) GetTransformedOutputSchema() map[string]interface{}`

GetTransformedOutputSchema returns the TransformedOutputSchema field if non-nil, zero value otherwise.

### GetTransformedOutputSchemaOk

`func (o *ActionDto) GetTransformedOutputSchemaOk() (*map[string]interface{}, bool)`

GetTransformedOutputSchemaOk returns a tuple with the TransformedOutputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformedOutputSchema

`func (o *ActionDto) SetTransformedOutputSchema(v map[string]interface{})`

SetTransformedOutputSchema sets TransformedOutputSchema field to given value.

### HasTransformedOutputSchema

`func (o *ActionDto) HasTransformedOutputSchema() bool`

HasTransformedOutputSchema returns a boolean if a field has been set.

### GetOutputSchema

`func (o *ActionDto) GetOutputSchema() map[string]interface{}`

GetOutputSchema returns the OutputSchema field if non-nil, zero value otherwise.

### GetOutputSchemaOk

`func (o *ActionDto) GetOutputSchemaOk() (*map[string]interface{}, bool)`

GetOutputSchemaOk returns a tuple with the OutputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputSchema

`func (o *ActionDto) SetOutputSchema(v map[string]interface{})`

SetOutputSchema sets OutputSchema field to given value.

### HasOutputSchema

`func (o *ActionDto) HasOutputSchema() bool`

HasOutputSchema returns a boolean if a field has been set.

### GetIsConfigurationMissing

`func (o *ActionDto) GetIsConfigurationMissing() bool`

GetIsConfigurationMissing returns the IsConfigurationMissing field if non-nil, zero value otherwise.

### GetIsConfigurationMissingOk

`func (o *ActionDto) GetIsConfigurationMissingOk() (*bool, bool)`

GetIsConfigurationMissingOk returns a tuple with the IsConfigurationMissing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConfigurationMissing

`func (o *ActionDto) SetIsConfigurationMissing(v bool)`

SetIsConfigurationMissing sets IsConfigurationMissing field to given value.

### HasIsConfigurationMissing

`func (o *ActionDto) HasIsConfigurationMissing() bool`

HasIsConfigurationMissing returns a boolean if a field has been set.

### GetDependencies

`func (o *ActionDto) GetDependencies() []map[string]interface{}`

GetDependencies returns the Dependencies field if non-nil, zero value otherwise.

### GetDependenciesOk

`func (o *ActionDto) GetDependenciesOk() (*[]map[string]interface{}, bool)`

GetDependenciesOk returns a tuple with the Dependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencies

`func (o *ActionDto) SetDependencies(v []map[string]interface{})`

SetDependencies sets Dependencies field to given value.

### HasDependencies

`func (o *ActionDto) HasDependencies() bool`

HasDependencies returns a boolean if a field has been set.

### GetIsDeployed

`func (o *ActionDto) GetIsDeployed() bool`

GetIsDeployed returns the IsDeployed field if non-nil, zero value otherwise.

### GetIsDeployedOk

`func (o *ActionDto) GetIsDeployedOk() (*bool, bool)`

GetIsDeployedOk returns a tuple with the IsDeployed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeployed

`func (o *ActionDto) SetIsDeployed(v bool)`

SetIsDeployed sets IsDeployed field to given value.

### HasIsDeployed

`func (o *ActionDto) HasIsDeployed() bool`

HasIsDeployed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


