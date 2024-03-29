# ActionInstanceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**InputSchema** | Pointer to **map[string]interface{}** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**DefaultOutputSchema** | Pointer to **map[string]interface{}** |  | [optional] 
**OutputSchema** | Pointer to **map[string]interface{}** |  | [optional] 
**State** | **string** |  | 
**Dependencies** | [**[]IntegrationElementInstanceDependencyDto**](IntegrationElementInstanceDependencyDto.md) |  | 
**Errors** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewActionInstanceDto

`func NewActionInstanceDto(type_ string, state string, dependencies []IntegrationElementInstanceDependencyDto, ) *ActionInstanceDto`

NewActionInstanceDto instantiates a new ActionInstanceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionInstanceDtoWithDefaults

`func NewActionInstanceDtoWithDefaults() *ActionInstanceDto`

NewActionInstanceDtoWithDefaults instantiates a new ActionInstanceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ActionInstanceDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ActionInstanceDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ActionInstanceDto) SetType(v string)`

SetType sets Type field to given value.


### GetInputSchema

`func (o *ActionInstanceDto) GetInputSchema() map[string]interface{}`

GetInputSchema returns the InputSchema field if non-nil, zero value otherwise.

### GetInputSchemaOk

`func (o *ActionInstanceDto) GetInputSchemaOk() (*map[string]interface{}, bool)`

GetInputSchemaOk returns a tuple with the InputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputSchema

`func (o *ActionInstanceDto) SetInputSchema(v map[string]interface{})`

SetInputSchema sets InputSchema field to given value.

### HasInputSchema

`func (o *ActionInstanceDto) HasInputSchema() bool`

HasInputSchema returns a boolean if a field has been set.

### GetConfig

`func (o *ActionInstanceDto) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ActionInstanceDto) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ActionInstanceDto) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ActionInstanceDto) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetDefaultOutputSchema

`func (o *ActionInstanceDto) GetDefaultOutputSchema() map[string]interface{}`

GetDefaultOutputSchema returns the DefaultOutputSchema field if non-nil, zero value otherwise.

### GetDefaultOutputSchemaOk

`func (o *ActionInstanceDto) GetDefaultOutputSchemaOk() (*map[string]interface{}, bool)`

GetDefaultOutputSchemaOk returns a tuple with the DefaultOutputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOutputSchema

`func (o *ActionInstanceDto) SetDefaultOutputSchema(v map[string]interface{})`

SetDefaultOutputSchema sets DefaultOutputSchema field to given value.

### HasDefaultOutputSchema

`func (o *ActionInstanceDto) HasDefaultOutputSchema() bool`

HasDefaultOutputSchema returns a boolean if a field has been set.

### GetOutputSchema

`func (o *ActionInstanceDto) GetOutputSchema() map[string]interface{}`

GetOutputSchema returns the OutputSchema field if non-nil, zero value otherwise.

### GetOutputSchemaOk

`func (o *ActionInstanceDto) GetOutputSchemaOk() (*map[string]interface{}, bool)`

GetOutputSchemaOk returns a tuple with the OutputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputSchema

`func (o *ActionInstanceDto) SetOutputSchema(v map[string]interface{})`

SetOutputSchema sets OutputSchema field to given value.

### HasOutputSchema

`func (o *ActionInstanceDto) HasOutputSchema() bool`

HasOutputSchema returns a boolean if a field has been set.

### GetState

`func (o *ActionInstanceDto) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ActionInstanceDto) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ActionInstanceDto) SetState(v string)`

SetState sets State field to given value.


### GetDependencies

`func (o *ActionInstanceDto) GetDependencies() []IntegrationElementInstanceDependencyDto`

GetDependencies returns the Dependencies field if non-nil, zero value otherwise.

### GetDependenciesOk

`func (o *ActionInstanceDto) GetDependenciesOk() (*[]IntegrationElementInstanceDependencyDto, bool)`

GetDependenciesOk returns a tuple with the Dependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencies

`func (o *ActionInstanceDto) SetDependencies(v []IntegrationElementInstanceDependencyDto)`

SetDependencies sets Dependencies field to given value.


### GetErrors

`func (o *ActionInstanceDto) GetErrors() []map[string]interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ActionInstanceDto) GetErrorsOk() (*[]map[string]interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ActionInstanceDto) SetErrors(v []map[string]interface{})`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ActionInstanceDto) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


