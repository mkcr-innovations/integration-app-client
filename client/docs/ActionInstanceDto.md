# ActionInstanceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**InstanceKey** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**CustomerDto**](CustomerDto.md) |  | [optional] 
**ConnectionId** | Pointer to **string** |  | [optional] 
**Connection** | Pointer to [**ConnectionDto**](ConnectionDto.md) |  | [optional] 
**IntegrationId** | Pointer to **string** |  | [optional] 
**Integration** | Pointer to [**IntegrationDto**](IntegrationDto.md) |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**ParentRevision** | Pointer to **string** |  | [optional] 
**Revision** | Pointer to **string** |  | [optional] 
**IsCustomized** | Pointer to **bool** |  | [optional] 
**IsOutdated** | Pointer to **bool** |  | [optional] 
**Error** | Pointer to **map[string]interface{}** |  | [optional] 
**ArchivedAt** | Pointer to **string** |  | [optional] 
**Parent** | Pointer to [**ActionDto**](ActionDto.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**InputSchema** | Pointer to **map[string]interface{}** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**DefaultOutputSchema** | Pointer to **map[string]interface{}** |  | [optional] 
**OutputSchema** | Pointer to **map[string]interface{}** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Dependencies** | Pointer to [**[]IntegrationElementInstanceDependencyDto**](IntegrationElementInstanceDependencyDto.md) |  | [optional] 
**Errors** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewActionInstanceDto

`func NewActionInstanceDto(id string, ) *ActionInstanceDto`

NewActionInstanceDto instantiates a new ActionInstanceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionInstanceDtoWithDefaults

`func NewActionInstanceDtoWithDefaults() *ActionInstanceDto`

NewActionInstanceDtoWithDefaults instantiates a new ActionInstanceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ActionInstanceDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActionInstanceDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActionInstanceDto) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ActionInstanceDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ActionInstanceDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ActionInstanceDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ActionInstanceDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetInstanceKey

`func (o *ActionInstanceDto) GetInstanceKey() string`

GetInstanceKey returns the InstanceKey field if non-nil, zero value otherwise.

### GetInstanceKeyOk

`func (o *ActionInstanceDto) GetInstanceKeyOk() (*string, bool)`

GetInstanceKeyOk returns a tuple with the InstanceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceKey

`func (o *ActionInstanceDto) SetInstanceKey(v string)`

SetInstanceKey sets InstanceKey field to given value.

### HasInstanceKey

`func (o *ActionInstanceDto) HasInstanceKey() bool`

HasInstanceKey returns a boolean if a field has been set.

### GetUserId

`func (o *ActionInstanceDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ActionInstanceDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ActionInstanceDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ActionInstanceDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUser

`func (o *ActionInstanceDto) GetUser() CustomerDto`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ActionInstanceDto) GetUserOk() (*CustomerDto, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ActionInstanceDto) SetUser(v CustomerDto)`

SetUser sets User field to given value.

### HasUser

`func (o *ActionInstanceDto) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetConnectionId

`func (o *ActionInstanceDto) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *ActionInstanceDto) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *ActionInstanceDto) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *ActionInstanceDto) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetConnection

`func (o *ActionInstanceDto) GetConnection() ConnectionDto`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *ActionInstanceDto) GetConnectionOk() (*ConnectionDto, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *ActionInstanceDto) SetConnection(v ConnectionDto)`

SetConnection sets Connection field to given value.

### HasConnection

`func (o *ActionInstanceDto) HasConnection() bool`

HasConnection returns a boolean if a field has been set.

### GetIntegrationId

`func (o *ActionInstanceDto) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *ActionInstanceDto) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *ActionInstanceDto) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.

### HasIntegrationId

`func (o *ActionInstanceDto) HasIntegrationId() bool`

HasIntegrationId returns a boolean if a field has been set.

### GetIntegration

`func (o *ActionInstanceDto) GetIntegration() IntegrationDto`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *ActionInstanceDto) GetIntegrationOk() (*IntegrationDto, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *ActionInstanceDto) SetIntegration(v IntegrationDto)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *ActionInstanceDto) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetParentId

`func (o *ActionInstanceDto) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *ActionInstanceDto) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *ActionInstanceDto) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *ActionInstanceDto) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetParentRevision

`func (o *ActionInstanceDto) GetParentRevision() string`

GetParentRevision returns the ParentRevision field if non-nil, zero value otherwise.

### GetParentRevisionOk

`func (o *ActionInstanceDto) GetParentRevisionOk() (*string, bool)`

GetParentRevisionOk returns a tuple with the ParentRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRevision

`func (o *ActionInstanceDto) SetParentRevision(v string)`

SetParentRevision sets ParentRevision field to given value.

### HasParentRevision

`func (o *ActionInstanceDto) HasParentRevision() bool`

HasParentRevision returns a boolean if a field has been set.

### GetRevision

`func (o *ActionInstanceDto) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ActionInstanceDto) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ActionInstanceDto) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *ActionInstanceDto) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetIsCustomized

`func (o *ActionInstanceDto) GetIsCustomized() bool`

GetIsCustomized returns the IsCustomized field if non-nil, zero value otherwise.

### GetIsCustomizedOk

`func (o *ActionInstanceDto) GetIsCustomizedOk() (*bool, bool)`

GetIsCustomizedOk returns a tuple with the IsCustomized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCustomized

`func (o *ActionInstanceDto) SetIsCustomized(v bool)`

SetIsCustomized sets IsCustomized field to given value.

### HasIsCustomized

`func (o *ActionInstanceDto) HasIsCustomized() bool`

HasIsCustomized returns a boolean if a field has been set.

### GetIsOutdated

`func (o *ActionInstanceDto) GetIsOutdated() bool`

GetIsOutdated returns the IsOutdated field if non-nil, zero value otherwise.

### GetIsOutdatedOk

`func (o *ActionInstanceDto) GetIsOutdatedOk() (*bool, bool)`

GetIsOutdatedOk returns a tuple with the IsOutdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOutdated

`func (o *ActionInstanceDto) SetIsOutdated(v bool)`

SetIsOutdated sets IsOutdated field to given value.

### HasIsOutdated

`func (o *ActionInstanceDto) HasIsOutdated() bool`

HasIsOutdated returns a boolean if a field has been set.

### GetError

`func (o *ActionInstanceDto) GetError() map[string]interface{}`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ActionInstanceDto) GetErrorOk() (*map[string]interface{}, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ActionInstanceDto) SetError(v map[string]interface{})`

SetError sets Error field to given value.

### HasError

`func (o *ActionInstanceDto) HasError() bool`

HasError returns a boolean if a field has been set.

### GetArchivedAt

`func (o *ActionInstanceDto) GetArchivedAt() string`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *ActionInstanceDto) GetArchivedAtOk() (*string, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *ActionInstanceDto) SetArchivedAt(v string)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *ActionInstanceDto) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### GetParent

`func (o *ActionInstanceDto) GetParent() ActionDto`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *ActionInstanceDto) GetParentOk() (*ActionDto, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *ActionInstanceDto) SetParent(v ActionDto)`

SetParent sets Parent field to given value.

### HasParent

`func (o *ActionInstanceDto) HasParent() bool`

HasParent returns a boolean if a field has been set.

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

### HasType

`func (o *ActionInstanceDto) HasType() bool`

HasType returns a boolean if a field has been set.

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

### HasState

`func (o *ActionInstanceDto) HasState() bool`

HasState returns a boolean if a field has been set.

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

### HasDependencies

`func (o *ActionInstanceDto) HasDependencies() bool`

HasDependencies returns a boolean if a field has been set.

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


