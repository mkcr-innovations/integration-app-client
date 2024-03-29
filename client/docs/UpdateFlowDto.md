# UpdateFlowDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nodes** | Pointer to **map[string]interface{}** |  | [optional] 
**ParametersSchema** | Pointer to **map[string]interface{}** |  | [optional] 
**AutoCreateInstances** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateFlowDto

`func NewUpdateFlowDto() *UpdateFlowDto`

NewUpdateFlowDto instantiates a new UpdateFlowDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFlowDtoWithDefaults

`func NewUpdateFlowDtoWithDefaults() *UpdateFlowDto`

NewUpdateFlowDtoWithDefaults instantiates a new UpdateFlowDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodes

`func (o *UpdateFlowDto) GetNodes() map[string]interface{}`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *UpdateFlowDto) GetNodesOk() (*map[string]interface{}, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *UpdateFlowDto) SetNodes(v map[string]interface{})`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *UpdateFlowDto) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetParametersSchema

`func (o *UpdateFlowDto) GetParametersSchema() map[string]interface{}`

GetParametersSchema returns the ParametersSchema field if non-nil, zero value otherwise.

### GetParametersSchemaOk

`func (o *UpdateFlowDto) GetParametersSchemaOk() (*map[string]interface{}, bool)`

GetParametersSchemaOk returns a tuple with the ParametersSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParametersSchema

`func (o *UpdateFlowDto) SetParametersSchema(v map[string]interface{})`

SetParametersSchema sets ParametersSchema field to given value.

### HasParametersSchema

`func (o *UpdateFlowDto) HasParametersSchema() bool`

HasParametersSchema returns a boolean if a field has been set.

### GetAutoCreateInstances

`func (o *UpdateFlowDto) GetAutoCreateInstances() bool`

GetAutoCreateInstances returns the AutoCreateInstances field if non-nil, zero value otherwise.

### GetAutoCreateInstancesOk

`func (o *UpdateFlowDto) GetAutoCreateInstancesOk() (*bool, bool)`

GetAutoCreateInstancesOk returns a tuple with the AutoCreateInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCreateInstances

`func (o *UpdateFlowDto) SetAutoCreateInstances(v bool)`

SetAutoCreateInstances sets AutoCreateInstances field to given value.

### HasAutoCreateInstances

`func (o *UpdateFlowDto) HasAutoCreateInstances() bool`

HasAutoCreateInstances returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


