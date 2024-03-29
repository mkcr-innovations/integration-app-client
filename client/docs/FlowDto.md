# FlowDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revision** | **string** |  | 
**UniversalFlowId** | Pointer to **string** |  | [optional] 
**IntegrationId** | Pointer to **string** |  | [optional] 
**Integration** | Pointer to [**IntegrationDto**](IntegrationDto.md) |  | [optional] 
**UniversalFlowRevision** | Pointer to **string** |  | [optional] 
**ParametersSchema** | Pointer to **map[string]interface{}** |  | [optional] 
**Nodes** | **map[string]interface{}** |  | 
**AutoCreateInstances** | Pointer to **bool** |  | [optional] 
**Customized** | Pointer to **bool** |  | [optional] 
**AppliedToIntegrations** | Pointer to **[]string** |  | [optional] 
**Dependencies** | Pointer to **[]map[string]interface{}** |  | [optional] 
**IsDeployed** | Pointer to **bool** |  | [optional] 

## Methods

### NewFlowDto

`func NewFlowDto(revision string, nodes map[string]interface{}, ) *FlowDto`

NewFlowDto instantiates a new FlowDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowDtoWithDefaults

`func NewFlowDtoWithDefaults() *FlowDto`

NewFlowDtoWithDefaults instantiates a new FlowDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevision

`func (o *FlowDto) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *FlowDto) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *FlowDto) SetRevision(v string)`

SetRevision sets Revision field to given value.


### GetUniversalFlowId

`func (o *FlowDto) GetUniversalFlowId() string`

GetUniversalFlowId returns the UniversalFlowId field if non-nil, zero value otherwise.

### GetUniversalFlowIdOk

`func (o *FlowDto) GetUniversalFlowIdOk() (*string, bool)`

GetUniversalFlowIdOk returns a tuple with the UniversalFlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniversalFlowId

`func (o *FlowDto) SetUniversalFlowId(v string)`

SetUniversalFlowId sets UniversalFlowId field to given value.

### HasUniversalFlowId

`func (o *FlowDto) HasUniversalFlowId() bool`

HasUniversalFlowId returns a boolean if a field has been set.

### GetIntegrationId

`func (o *FlowDto) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *FlowDto) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *FlowDto) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.

### HasIntegrationId

`func (o *FlowDto) HasIntegrationId() bool`

HasIntegrationId returns a boolean if a field has been set.

### GetIntegration

`func (o *FlowDto) GetIntegration() IntegrationDto`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *FlowDto) GetIntegrationOk() (*IntegrationDto, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *FlowDto) SetIntegration(v IntegrationDto)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *FlowDto) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetUniversalFlowRevision

`func (o *FlowDto) GetUniversalFlowRevision() string`

GetUniversalFlowRevision returns the UniversalFlowRevision field if non-nil, zero value otherwise.

### GetUniversalFlowRevisionOk

`func (o *FlowDto) GetUniversalFlowRevisionOk() (*string, bool)`

GetUniversalFlowRevisionOk returns a tuple with the UniversalFlowRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniversalFlowRevision

`func (o *FlowDto) SetUniversalFlowRevision(v string)`

SetUniversalFlowRevision sets UniversalFlowRevision field to given value.

### HasUniversalFlowRevision

`func (o *FlowDto) HasUniversalFlowRevision() bool`

HasUniversalFlowRevision returns a boolean if a field has been set.

### GetParametersSchema

`func (o *FlowDto) GetParametersSchema() map[string]interface{}`

GetParametersSchema returns the ParametersSchema field if non-nil, zero value otherwise.

### GetParametersSchemaOk

`func (o *FlowDto) GetParametersSchemaOk() (*map[string]interface{}, bool)`

GetParametersSchemaOk returns a tuple with the ParametersSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParametersSchema

`func (o *FlowDto) SetParametersSchema(v map[string]interface{})`

SetParametersSchema sets ParametersSchema field to given value.

### HasParametersSchema

`func (o *FlowDto) HasParametersSchema() bool`

HasParametersSchema returns a boolean if a field has been set.

### GetNodes

`func (o *FlowDto) GetNodes() map[string]interface{}`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *FlowDto) GetNodesOk() (*map[string]interface{}, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *FlowDto) SetNodes(v map[string]interface{})`

SetNodes sets Nodes field to given value.


### GetAutoCreateInstances

`func (o *FlowDto) GetAutoCreateInstances() bool`

GetAutoCreateInstances returns the AutoCreateInstances field if non-nil, zero value otherwise.

### GetAutoCreateInstancesOk

`func (o *FlowDto) GetAutoCreateInstancesOk() (*bool, bool)`

GetAutoCreateInstancesOk returns a tuple with the AutoCreateInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCreateInstances

`func (o *FlowDto) SetAutoCreateInstances(v bool)`

SetAutoCreateInstances sets AutoCreateInstances field to given value.

### HasAutoCreateInstances

`func (o *FlowDto) HasAutoCreateInstances() bool`

HasAutoCreateInstances returns a boolean if a field has been set.

### GetCustomized

`func (o *FlowDto) GetCustomized() bool`

GetCustomized returns the Customized field if non-nil, zero value otherwise.

### GetCustomizedOk

`func (o *FlowDto) GetCustomizedOk() (*bool, bool)`

GetCustomizedOk returns a tuple with the Customized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomized

`func (o *FlowDto) SetCustomized(v bool)`

SetCustomized sets Customized field to given value.

### HasCustomized

`func (o *FlowDto) HasCustomized() bool`

HasCustomized returns a boolean if a field has been set.

### GetAppliedToIntegrations

`func (o *FlowDto) GetAppliedToIntegrations() []string`

GetAppliedToIntegrations returns the AppliedToIntegrations field if non-nil, zero value otherwise.

### GetAppliedToIntegrationsOk

`func (o *FlowDto) GetAppliedToIntegrationsOk() (*[]string, bool)`

GetAppliedToIntegrationsOk returns a tuple with the AppliedToIntegrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedToIntegrations

`func (o *FlowDto) SetAppliedToIntegrations(v []string)`

SetAppliedToIntegrations sets AppliedToIntegrations field to given value.

### HasAppliedToIntegrations

`func (o *FlowDto) HasAppliedToIntegrations() bool`

HasAppliedToIntegrations returns a boolean if a field has been set.

### GetDependencies

`func (o *FlowDto) GetDependencies() []map[string]interface{}`

GetDependencies returns the Dependencies field if non-nil, zero value otherwise.

### GetDependenciesOk

`func (o *FlowDto) GetDependenciesOk() (*[]map[string]interface{}, bool)`

GetDependenciesOk returns a tuple with the Dependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencies

`func (o *FlowDto) SetDependencies(v []map[string]interface{})`

SetDependencies sets Dependencies field to given value.

### HasDependencies

`func (o *FlowDto) HasDependencies() bool`

HasDependencies returns a boolean if a field has been set.

### GetIsDeployed

`func (o *FlowDto) GetIsDeployed() bool`

GetIsDeployed returns the IsDeployed field if non-nil, zero value otherwise.

### GetIsDeployedOk

`func (o *FlowDto) GetIsDeployedOk() (*bool, bool)`

GetIsDeployedOk returns a tuple with the IsDeployed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeployed

`func (o *FlowDto) SetIsDeployed(v bool)`

SetIsDeployed sets IsDeployed field to given value.

### HasIsDeployed

`func (o *FlowDto) HasIsDeployed() bool`

HasIsDeployed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


