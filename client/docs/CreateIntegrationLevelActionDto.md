# CreateIntegrationLevelActionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Name** | **string** |  | 
**ArchivedAt** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**InputSchema** | Pointer to **map[string]interface{}** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**OutputMapping** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomOutputSchema** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCreateIntegrationLevelActionDto

`func NewCreateIntegrationLevelActionDto(key string, name string, ) *CreateIntegrationLevelActionDto`

NewCreateIntegrationLevelActionDto instantiates a new CreateIntegrationLevelActionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIntegrationLevelActionDtoWithDefaults

`func NewCreateIntegrationLevelActionDtoWithDefaults() *CreateIntegrationLevelActionDto`

NewCreateIntegrationLevelActionDtoWithDefaults instantiates a new CreateIntegrationLevelActionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *CreateIntegrationLevelActionDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateIntegrationLevelActionDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateIntegrationLevelActionDto) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *CreateIntegrationLevelActionDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateIntegrationLevelActionDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateIntegrationLevelActionDto) SetName(v string)`

SetName sets Name field to given value.


### GetArchivedAt

`func (o *CreateIntegrationLevelActionDto) GetArchivedAt() string`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *CreateIntegrationLevelActionDto) GetArchivedAtOk() (*string, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *CreateIntegrationLevelActionDto) SetArchivedAt(v string)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *CreateIntegrationLevelActionDto) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### GetType

`func (o *CreateIntegrationLevelActionDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateIntegrationLevelActionDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateIntegrationLevelActionDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateIntegrationLevelActionDto) HasType() bool`

HasType returns a boolean if a field has been set.

### GetInputSchema

`func (o *CreateIntegrationLevelActionDto) GetInputSchema() map[string]interface{}`

GetInputSchema returns the InputSchema field if non-nil, zero value otherwise.

### GetInputSchemaOk

`func (o *CreateIntegrationLevelActionDto) GetInputSchemaOk() (*map[string]interface{}, bool)`

GetInputSchemaOk returns a tuple with the InputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputSchema

`func (o *CreateIntegrationLevelActionDto) SetInputSchema(v map[string]interface{})`

SetInputSchema sets InputSchema field to given value.

### HasInputSchema

`func (o *CreateIntegrationLevelActionDto) HasInputSchema() bool`

HasInputSchema returns a boolean if a field has been set.

### GetConfig

`func (o *CreateIntegrationLevelActionDto) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CreateIntegrationLevelActionDto) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CreateIntegrationLevelActionDto) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CreateIntegrationLevelActionDto) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetOutputMapping

`func (o *CreateIntegrationLevelActionDto) GetOutputMapping() map[string]interface{}`

GetOutputMapping returns the OutputMapping field if non-nil, zero value otherwise.

### GetOutputMappingOk

`func (o *CreateIntegrationLevelActionDto) GetOutputMappingOk() (*map[string]interface{}, bool)`

GetOutputMappingOk returns a tuple with the OutputMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputMapping

`func (o *CreateIntegrationLevelActionDto) SetOutputMapping(v map[string]interface{})`

SetOutputMapping sets OutputMapping field to given value.

### HasOutputMapping

`func (o *CreateIntegrationLevelActionDto) HasOutputMapping() bool`

HasOutputMapping returns a boolean if a field has been set.

### GetCustomOutputSchema

`func (o *CreateIntegrationLevelActionDto) GetCustomOutputSchema() map[string]interface{}`

GetCustomOutputSchema returns the CustomOutputSchema field if non-nil, zero value otherwise.

### GetCustomOutputSchemaOk

`func (o *CreateIntegrationLevelActionDto) GetCustomOutputSchemaOk() (*map[string]interface{}, bool)`

GetCustomOutputSchemaOk returns a tuple with the CustomOutputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomOutputSchema

`func (o *CreateIntegrationLevelActionDto) SetCustomOutputSchema(v map[string]interface{})`

SetCustomOutputSchema sets CustomOutputSchema field to given value.

### HasCustomOutputSchema

`func (o *CreateIntegrationLevelActionDto) HasCustomOutputSchema() bool`

HasCustomOutputSchema returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


