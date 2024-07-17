# ActionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Key** | **string** |  | 
**Name** | **string** |  | 
**ArchivedAt** | Pointer to **string** |  | [optional] 
**Revision** | **string** |  | 
**IsCustomized** | Pointer to **bool** |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**ParentRevision** | Pointer to **string** |  | [optional] 
**IntegrationId** | Pointer to **string** |  | [optional] 
**Integration** | Pointer to [**IntegrationDto**](IntegrationDto.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**InputSchema** | Pointer to **map[string]interface{}** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**OutputMapping** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomOutputSchema** | Pointer to **map[string]interface{}** |  | [optional] 
**DefaultOutputSchema** | Pointer to **map[string]interface{}** |  | [optional] 
**TransformedOutputSchema** | Pointer to **map[string]interface{}** |  | [optional] 
**OutputSchema** | Pointer to **map[string]interface{}** |  | [optional] 
**Dependencies** | Pointer to **[]map[string]interface{}** |  | [optional] 
**IsDeployed** | Pointer to **bool** |  | [optional] 
**ConfigurationState** | Pointer to **string** |  | [optional] 
**ConfigurationStateMessage** | Pointer to **string** |  | [optional] 
**AppliedToIntegrations** | Pointer to [**[]ActionWithIntegrationDto**](ActionWithIntegrationDto.md) |  | [optional] 

## Methods

### NewActionDto

`func NewActionDto(id string, key string, name string, revision string, ) *ActionDto`

NewActionDto instantiates a new ActionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionDtoWithDefaults

`func NewActionDtoWithDefaults() *ActionDto`

NewActionDtoWithDefaults instantiates a new ActionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ActionDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActionDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActionDto) SetId(v string)`

SetId sets Id field to given value.


### GetKey

`func (o *ActionDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ActionDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ActionDto) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *ActionDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ActionDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ActionDto) SetName(v string)`

SetName sets Name field to given value.


### GetArchivedAt

`func (o *ActionDto) GetArchivedAt() string`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *ActionDto) GetArchivedAtOk() (*string, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *ActionDto) SetArchivedAt(v string)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *ActionDto) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### GetRevision

`func (o *ActionDto) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ActionDto) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ActionDto) SetRevision(v string)`

SetRevision sets Revision field to given value.


### GetIsCustomized

`func (o *ActionDto) GetIsCustomized() bool`

GetIsCustomized returns the IsCustomized field if non-nil, zero value otherwise.

### GetIsCustomizedOk

`func (o *ActionDto) GetIsCustomizedOk() (*bool, bool)`

GetIsCustomizedOk returns a tuple with the IsCustomized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCustomized

`func (o *ActionDto) SetIsCustomized(v bool)`

SetIsCustomized sets IsCustomized field to given value.

### HasIsCustomized

`func (o *ActionDto) HasIsCustomized() bool`

HasIsCustomized returns a boolean if a field has been set.

### GetParentId

`func (o *ActionDto) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *ActionDto) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *ActionDto) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *ActionDto) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetParentRevision

`func (o *ActionDto) GetParentRevision() string`

GetParentRevision returns the ParentRevision field if non-nil, zero value otherwise.

### GetParentRevisionOk

`func (o *ActionDto) GetParentRevisionOk() (*string, bool)`

GetParentRevisionOk returns a tuple with the ParentRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRevision

`func (o *ActionDto) SetParentRevision(v string)`

SetParentRevision sets ParentRevision field to given value.

### HasParentRevision

`func (o *ActionDto) HasParentRevision() bool`

HasParentRevision returns a boolean if a field has been set.

### GetIntegrationId

`func (o *ActionDto) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *ActionDto) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *ActionDto) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.

### HasIntegrationId

`func (o *ActionDto) HasIntegrationId() bool`

HasIntegrationId returns a boolean if a field has been set.

### GetIntegration

`func (o *ActionDto) GetIntegration() IntegrationDto`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *ActionDto) GetIntegrationOk() (*IntegrationDto, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *ActionDto) SetIntegration(v IntegrationDto)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *ActionDto) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

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

### HasType

`func (o *ActionDto) HasType() bool`

HasType returns a boolean if a field has been set.

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

### GetConfigurationState

`func (o *ActionDto) GetConfigurationState() string`

GetConfigurationState returns the ConfigurationState field if non-nil, zero value otherwise.

### GetConfigurationStateOk

`func (o *ActionDto) GetConfigurationStateOk() (*string, bool)`

GetConfigurationStateOk returns a tuple with the ConfigurationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationState

`func (o *ActionDto) SetConfigurationState(v string)`

SetConfigurationState sets ConfigurationState field to given value.

### HasConfigurationState

`func (o *ActionDto) HasConfigurationState() bool`

HasConfigurationState returns a boolean if a field has been set.

### GetConfigurationStateMessage

`func (o *ActionDto) GetConfigurationStateMessage() string`

GetConfigurationStateMessage returns the ConfigurationStateMessage field if non-nil, zero value otherwise.

### GetConfigurationStateMessageOk

`func (o *ActionDto) GetConfigurationStateMessageOk() (*string, bool)`

GetConfigurationStateMessageOk returns a tuple with the ConfigurationStateMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationStateMessage

`func (o *ActionDto) SetConfigurationStateMessage(v string)`

SetConfigurationStateMessage sets ConfigurationStateMessage field to given value.

### HasConfigurationStateMessage

`func (o *ActionDto) HasConfigurationStateMessage() bool`

HasConfigurationStateMessage returns a boolean if a field has been set.

### GetAppliedToIntegrations

`func (o *ActionDto) GetAppliedToIntegrations() []ActionWithIntegrationDto`

GetAppliedToIntegrations returns the AppliedToIntegrations field if non-nil, zero value otherwise.

### GetAppliedToIntegrationsOk

`func (o *ActionDto) GetAppliedToIntegrationsOk() (*[]ActionWithIntegrationDto, bool)`

GetAppliedToIntegrationsOk returns a tuple with the AppliedToIntegrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedToIntegrations

`func (o *ActionDto) SetAppliedToIntegrations(v []ActionWithIntegrationDto)`

SetAppliedToIntegrations sets AppliedToIntegrations field to given value.

### HasAppliedToIntegrations

`func (o *ActionDto) HasAppliedToIntegrations() bool`

HasAppliedToIntegrations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


