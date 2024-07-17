# FlowExportDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**IntegrationKeys** | Pointer to **[]string** |  | [optional] 
**ParametersSchema** | Pointer to **map[string]interface{}** |  | [optional] 
**Nodes** | Pointer to **map[string]interface{}** |  | [optional] 
**AutoCreateInstances** | Pointer to **bool** |  | [optional] 

## Methods

### NewFlowExportDto

`func NewFlowExportDto(name string, ) *FlowExportDto`

NewFlowExportDto instantiates a new FlowExportDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowExportDtoWithDefaults

`func NewFlowExportDtoWithDefaults() *FlowExportDto`

NewFlowExportDtoWithDefaults instantiates a new FlowExportDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *FlowExportDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FlowExportDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FlowExportDto) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *FlowExportDto) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *FlowExportDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FlowExportDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FlowExportDto) SetName(v string)`

SetName sets Name field to given value.


### GetIntegrationKeys

`func (o *FlowExportDto) GetIntegrationKeys() []string`

GetIntegrationKeys returns the IntegrationKeys field if non-nil, zero value otherwise.

### GetIntegrationKeysOk

`func (o *FlowExportDto) GetIntegrationKeysOk() (*[]string, bool)`

GetIntegrationKeysOk returns a tuple with the IntegrationKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationKeys

`func (o *FlowExportDto) SetIntegrationKeys(v []string)`

SetIntegrationKeys sets IntegrationKeys field to given value.

### HasIntegrationKeys

`func (o *FlowExportDto) HasIntegrationKeys() bool`

HasIntegrationKeys returns a boolean if a field has been set.

### GetParametersSchema

`func (o *FlowExportDto) GetParametersSchema() map[string]interface{}`

GetParametersSchema returns the ParametersSchema field if non-nil, zero value otherwise.

### GetParametersSchemaOk

`func (o *FlowExportDto) GetParametersSchemaOk() (*map[string]interface{}, bool)`

GetParametersSchemaOk returns a tuple with the ParametersSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParametersSchema

`func (o *FlowExportDto) SetParametersSchema(v map[string]interface{})`

SetParametersSchema sets ParametersSchema field to given value.

### HasParametersSchema

`func (o *FlowExportDto) HasParametersSchema() bool`

HasParametersSchema returns a boolean if a field has been set.

### GetNodes

`func (o *FlowExportDto) GetNodes() map[string]interface{}`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *FlowExportDto) GetNodesOk() (*map[string]interface{}, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *FlowExportDto) SetNodes(v map[string]interface{})`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *FlowExportDto) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetAutoCreateInstances

`func (o *FlowExportDto) GetAutoCreateInstances() bool`

GetAutoCreateInstances returns the AutoCreateInstances field if non-nil, zero value otherwise.

### GetAutoCreateInstancesOk

`func (o *FlowExportDto) GetAutoCreateInstancesOk() (*bool, bool)`

GetAutoCreateInstancesOk returns a tuple with the AutoCreateInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCreateInstances

`func (o *FlowExportDto) SetAutoCreateInstances(v bool)`

SetAutoCreateInstances sets AutoCreateInstances field to given value.

### HasAutoCreateInstances

`func (o *FlowExportDto) HasAutoCreateInstances() bool`

HasAutoCreateInstances returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


