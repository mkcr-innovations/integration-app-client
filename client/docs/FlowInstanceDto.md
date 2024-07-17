# FlowInstanceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**UserId** | **string** |  | 
**User** | Pointer to [**CustomerDto**](CustomerDto.md) |  | [optional] 
**InstanceKey** | Pointer to **string** |  | [optional] 
**ConnectionId** | **string** |  | 
**IntegrationId** | **string** |  | 
**Integration** | Pointer to [**IntegrationDto**](IntegrationDto.md) |  | [optional] 
**FlowId** | **string** |  | 
**Flow** | Pointer to [**FlowDto**](FlowDto.md) |  | [optional] 
**UniversalFlowId** | Pointer to **string** |  | [optional] 
**FlowRevision** | **string** |  | 
**Outdated** | Pointer to **bool** |  | [optional] 
**Name** | **string** |  | 
**ParametersSchema** | Pointer to **map[string]interface{}** |  | [optional] 
**Parameters** | Pointer to **map[string]interface{}** |  | [optional] 
**Nodes** | Pointer to **map[string]interface{}** |  | [optional] 
**Enabled** | **bool** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**ArchivedAt** | Pointer to **time.Time** |  | [optional] 
**Dependencies** | Pointer to [**[]IntegrationElementInstanceDependencyDto**](IntegrationElementInstanceDependencyDto.md) |  | [optional] 
**State** | **string** |  | 
**Connection** | Pointer to [**ConnectionDto**](ConnectionDto.md) |  | [optional] 

## Methods

### NewFlowInstanceDto

`func NewFlowInstanceDto(id string, userId string, connectionId string, integrationId string, flowId string, flowRevision string, name string, enabled bool, createdAt string, updatedAt string, state string, ) *FlowInstanceDto`

NewFlowInstanceDto instantiates a new FlowInstanceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowInstanceDtoWithDefaults

`func NewFlowInstanceDtoWithDefaults() *FlowInstanceDto`

NewFlowInstanceDtoWithDefaults instantiates a new FlowInstanceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FlowInstanceDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FlowInstanceDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FlowInstanceDto) SetId(v string)`

SetId sets Id field to given value.


### GetUserId

`func (o *FlowInstanceDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *FlowInstanceDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *FlowInstanceDto) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetUser

`func (o *FlowInstanceDto) GetUser() CustomerDto`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *FlowInstanceDto) GetUserOk() (*CustomerDto, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *FlowInstanceDto) SetUser(v CustomerDto)`

SetUser sets User field to given value.

### HasUser

`func (o *FlowInstanceDto) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetInstanceKey

`func (o *FlowInstanceDto) GetInstanceKey() string`

GetInstanceKey returns the InstanceKey field if non-nil, zero value otherwise.

### GetInstanceKeyOk

`func (o *FlowInstanceDto) GetInstanceKeyOk() (*string, bool)`

GetInstanceKeyOk returns a tuple with the InstanceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceKey

`func (o *FlowInstanceDto) SetInstanceKey(v string)`

SetInstanceKey sets InstanceKey field to given value.

### HasInstanceKey

`func (o *FlowInstanceDto) HasInstanceKey() bool`

HasInstanceKey returns a boolean if a field has been set.

### GetConnectionId

`func (o *FlowInstanceDto) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *FlowInstanceDto) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *FlowInstanceDto) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetIntegrationId

`func (o *FlowInstanceDto) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *FlowInstanceDto) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *FlowInstanceDto) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.


### GetIntegration

`func (o *FlowInstanceDto) GetIntegration() IntegrationDto`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *FlowInstanceDto) GetIntegrationOk() (*IntegrationDto, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *FlowInstanceDto) SetIntegration(v IntegrationDto)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *FlowInstanceDto) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetFlowId

`func (o *FlowInstanceDto) GetFlowId() string`

GetFlowId returns the FlowId field if non-nil, zero value otherwise.

### GetFlowIdOk

`func (o *FlowInstanceDto) GetFlowIdOk() (*string, bool)`

GetFlowIdOk returns a tuple with the FlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowId

`func (o *FlowInstanceDto) SetFlowId(v string)`

SetFlowId sets FlowId field to given value.


### GetFlow

`func (o *FlowInstanceDto) GetFlow() FlowDto`

GetFlow returns the Flow field if non-nil, zero value otherwise.

### GetFlowOk

`func (o *FlowInstanceDto) GetFlowOk() (*FlowDto, bool)`

GetFlowOk returns a tuple with the Flow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlow

`func (o *FlowInstanceDto) SetFlow(v FlowDto)`

SetFlow sets Flow field to given value.

### HasFlow

`func (o *FlowInstanceDto) HasFlow() bool`

HasFlow returns a boolean if a field has been set.

### GetUniversalFlowId

`func (o *FlowInstanceDto) GetUniversalFlowId() string`

GetUniversalFlowId returns the UniversalFlowId field if non-nil, zero value otherwise.

### GetUniversalFlowIdOk

`func (o *FlowInstanceDto) GetUniversalFlowIdOk() (*string, bool)`

GetUniversalFlowIdOk returns a tuple with the UniversalFlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniversalFlowId

`func (o *FlowInstanceDto) SetUniversalFlowId(v string)`

SetUniversalFlowId sets UniversalFlowId field to given value.

### HasUniversalFlowId

`func (o *FlowInstanceDto) HasUniversalFlowId() bool`

HasUniversalFlowId returns a boolean if a field has been set.

### GetFlowRevision

`func (o *FlowInstanceDto) GetFlowRevision() string`

GetFlowRevision returns the FlowRevision field if non-nil, zero value otherwise.

### GetFlowRevisionOk

`func (o *FlowInstanceDto) GetFlowRevisionOk() (*string, bool)`

GetFlowRevisionOk returns a tuple with the FlowRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowRevision

`func (o *FlowInstanceDto) SetFlowRevision(v string)`

SetFlowRevision sets FlowRevision field to given value.


### GetOutdated

`func (o *FlowInstanceDto) GetOutdated() bool`

GetOutdated returns the Outdated field if non-nil, zero value otherwise.

### GetOutdatedOk

`func (o *FlowInstanceDto) GetOutdatedOk() (*bool, bool)`

GetOutdatedOk returns a tuple with the Outdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutdated

`func (o *FlowInstanceDto) SetOutdated(v bool)`

SetOutdated sets Outdated field to given value.

### HasOutdated

`func (o *FlowInstanceDto) HasOutdated() bool`

HasOutdated returns a boolean if a field has been set.

### GetName

`func (o *FlowInstanceDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FlowInstanceDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FlowInstanceDto) SetName(v string)`

SetName sets Name field to given value.


### GetParametersSchema

`func (o *FlowInstanceDto) GetParametersSchema() map[string]interface{}`

GetParametersSchema returns the ParametersSchema field if non-nil, zero value otherwise.

### GetParametersSchemaOk

`func (o *FlowInstanceDto) GetParametersSchemaOk() (*map[string]interface{}, bool)`

GetParametersSchemaOk returns a tuple with the ParametersSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParametersSchema

`func (o *FlowInstanceDto) SetParametersSchema(v map[string]interface{})`

SetParametersSchema sets ParametersSchema field to given value.

### HasParametersSchema

`func (o *FlowInstanceDto) HasParametersSchema() bool`

HasParametersSchema returns a boolean if a field has been set.

### GetParameters

`func (o *FlowInstanceDto) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *FlowInstanceDto) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *FlowInstanceDto) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *FlowInstanceDto) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetNodes

`func (o *FlowInstanceDto) GetNodes() map[string]interface{}`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *FlowInstanceDto) GetNodesOk() (*map[string]interface{}, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *FlowInstanceDto) SetNodes(v map[string]interface{})`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *FlowInstanceDto) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetEnabled

`func (o *FlowInstanceDto) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FlowInstanceDto) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FlowInstanceDto) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetCreatedAt

`func (o *FlowInstanceDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FlowInstanceDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FlowInstanceDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *FlowInstanceDto) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FlowInstanceDto) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FlowInstanceDto) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetArchivedAt

`func (o *FlowInstanceDto) GetArchivedAt() time.Time`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *FlowInstanceDto) GetArchivedAtOk() (*time.Time, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *FlowInstanceDto) SetArchivedAt(v time.Time)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *FlowInstanceDto) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### GetDependencies

`func (o *FlowInstanceDto) GetDependencies() []IntegrationElementInstanceDependencyDto`

GetDependencies returns the Dependencies field if non-nil, zero value otherwise.

### GetDependenciesOk

`func (o *FlowInstanceDto) GetDependenciesOk() (*[]IntegrationElementInstanceDependencyDto, bool)`

GetDependenciesOk returns a tuple with the Dependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencies

`func (o *FlowInstanceDto) SetDependencies(v []IntegrationElementInstanceDependencyDto)`

SetDependencies sets Dependencies field to given value.

### HasDependencies

`func (o *FlowInstanceDto) HasDependencies() bool`

HasDependencies returns a boolean if a field has been set.

### GetState

`func (o *FlowInstanceDto) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FlowInstanceDto) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FlowInstanceDto) SetState(v string)`

SetState sets State field to given value.


### GetConnection

`func (o *FlowInstanceDto) GetConnection() ConnectionDto`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *FlowInstanceDto) GetConnectionOk() (*ConnectionDto, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *FlowInstanceDto) SetConnection(v ConnectionDto)`

SetConnection sets Connection field to given value.

### HasConnection

`func (o *FlowInstanceDto) HasConnection() bool`

HasConnection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


