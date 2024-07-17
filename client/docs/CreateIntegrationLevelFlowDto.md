# CreateIntegrationLevelFlowDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Name** | **string** |  | 
**ArchivedAt** | Pointer to **string** |  | [optional] 
**Nodes** | Pointer to **map[string]interface{}** |  | [optional] 
**ParametersSchema** | Pointer to **map[string]interface{}** |  | [optional] 
**AutoCreateInstances** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreateIntegrationLevelFlowDto

`func NewCreateIntegrationLevelFlowDto(key string, name string, ) *CreateIntegrationLevelFlowDto`

NewCreateIntegrationLevelFlowDto instantiates a new CreateIntegrationLevelFlowDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIntegrationLevelFlowDtoWithDefaults

`func NewCreateIntegrationLevelFlowDtoWithDefaults() *CreateIntegrationLevelFlowDto`

NewCreateIntegrationLevelFlowDtoWithDefaults instantiates a new CreateIntegrationLevelFlowDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *CreateIntegrationLevelFlowDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateIntegrationLevelFlowDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateIntegrationLevelFlowDto) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *CreateIntegrationLevelFlowDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateIntegrationLevelFlowDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateIntegrationLevelFlowDto) SetName(v string)`

SetName sets Name field to given value.


### GetArchivedAt

`func (o *CreateIntegrationLevelFlowDto) GetArchivedAt() string`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *CreateIntegrationLevelFlowDto) GetArchivedAtOk() (*string, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *CreateIntegrationLevelFlowDto) SetArchivedAt(v string)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *CreateIntegrationLevelFlowDto) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### GetNodes

`func (o *CreateIntegrationLevelFlowDto) GetNodes() map[string]interface{}`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *CreateIntegrationLevelFlowDto) GetNodesOk() (*map[string]interface{}, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *CreateIntegrationLevelFlowDto) SetNodes(v map[string]interface{})`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *CreateIntegrationLevelFlowDto) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetParametersSchema

`func (o *CreateIntegrationLevelFlowDto) GetParametersSchema() map[string]interface{}`

GetParametersSchema returns the ParametersSchema field if non-nil, zero value otherwise.

### GetParametersSchemaOk

`func (o *CreateIntegrationLevelFlowDto) GetParametersSchemaOk() (*map[string]interface{}, bool)`

GetParametersSchemaOk returns a tuple with the ParametersSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParametersSchema

`func (o *CreateIntegrationLevelFlowDto) SetParametersSchema(v map[string]interface{})`

SetParametersSchema sets ParametersSchema field to given value.

### HasParametersSchema

`func (o *CreateIntegrationLevelFlowDto) HasParametersSchema() bool`

HasParametersSchema returns a boolean if a field has been set.

### GetAutoCreateInstances

`func (o *CreateIntegrationLevelFlowDto) GetAutoCreateInstances() bool`

GetAutoCreateInstances returns the AutoCreateInstances field if non-nil, zero value otherwise.

### GetAutoCreateInstancesOk

`func (o *CreateIntegrationLevelFlowDto) GetAutoCreateInstancesOk() (*bool, bool)`

GetAutoCreateInstancesOk returns a tuple with the AutoCreateInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCreateInstances

`func (o *CreateIntegrationLevelFlowDto) SetAutoCreateInstances(v bool)`

SetAutoCreateInstances sets AutoCreateInstances field to given value.

### HasAutoCreateInstances

`func (o *CreateIntegrationLevelFlowDto) HasAutoCreateInstances() bool`

HasAutoCreateInstances returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


