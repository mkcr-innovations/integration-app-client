# CreateActionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**ArchivedAt** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**InputSchema** | Pointer to **map[string]interface{}** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**OutputMapping** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomOutputSchema** | Pointer to **map[string]interface{}** |  | [optional] 
**IntegrationId** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateActionDto

`func NewCreateActionDto(key string, ) *CreateActionDto`

NewCreateActionDto instantiates a new CreateActionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateActionDtoWithDefaults

`func NewCreateActionDtoWithDefaults() *CreateActionDto`

NewCreateActionDtoWithDefaults instantiates a new CreateActionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *CreateActionDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateActionDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateActionDto) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *CreateActionDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateActionDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateActionDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateActionDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetArchivedAt

`func (o *CreateActionDto) GetArchivedAt() string`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *CreateActionDto) GetArchivedAtOk() (*string, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *CreateActionDto) SetArchivedAt(v string)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *CreateActionDto) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### GetType

`func (o *CreateActionDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateActionDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateActionDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateActionDto) HasType() bool`

HasType returns a boolean if a field has been set.

### GetInputSchema

`func (o *CreateActionDto) GetInputSchema() map[string]interface{}`

GetInputSchema returns the InputSchema field if non-nil, zero value otherwise.

### GetInputSchemaOk

`func (o *CreateActionDto) GetInputSchemaOk() (*map[string]interface{}, bool)`

GetInputSchemaOk returns a tuple with the InputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputSchema

`func (o *CreateActionDto) SetInputSchema(v map[string]interface{})`

SetInputSchema sets InputSchema field to given value.

### HasInputSchema

`func (o *CreateActionDto) HasInputSchema() bool`

HasInputSchema returns a boolean if a field has been set.

### GetConfig

`func (o *CreateActionDto) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CreateActionDto) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CreateActionDto) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CreateActionDto) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetOutputMapping

`func (o *CreateActionDto) GetOutputMapping() map[string]interface{}`

GetOutputMapping returns the OutputMapping field if non-nil, zero value otherwise.

### GetOutputMappingOk

`func (o *CreateActionDto) GetOutputMappingOk() (*map[string]interface{}, bool)`

GetOutputMappingOk returns a tuple with the OutputMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputMapping

`func (o *CreateActionDto) SetOutputMapping(v map[string]interface{})`

SetOutputMapping sets OutputMapping field to given value.

### HasOutputMapping

`func (o *CreateActionDto) HasOutputMapping() bool`

HasOutputMapping returns a boolean if a field has been set.

### GetCustomOutputSchema

`func (o *CreateActionDto) GetCustomOutputSchema() map[string]interface{}`

GetCustomOutputSchema returns the CustomOutputSchema field if non-nil, zero value otherwise.

### GetCustomOutputSchemaOk

`func (o *CreateActionDto) GetCustomOutputSchemaOk() (*map[string]interface{}, bool)`

GetCustomOutputSchemaOk returns a tuple with the CustomOutputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomOutputSchema

`func (o *CreateActionDto) SetCustomOutputSchema(v map[string]interface{})`

SetCustomOutputSchema sets CustomOutputSchema field to given value.

### HasCustomOutputSchema

`func (o *CreateActionDto) HasCustomOutputSchema() bool`

HasCustomOutputSchema returns a boolean if a field has been set.

### GetIntegrationId

`func (o *CreateActionDto) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *CreateActionDto) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *CreateActionDto) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.

### HasIntegrationId

`func (o *CreateActionDto) HasIntegrationId() bool`

HasIntegrationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


