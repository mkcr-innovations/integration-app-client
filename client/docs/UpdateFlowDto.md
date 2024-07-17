# UpdateFlowDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ArchivedAt** | Pointer to **string** |  | [optional] 
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

### GetKey

`func (o *UpdateFlowDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UpdateFlowDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UpdateFlowDto) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *UpdateFlowDto) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *UpdateFlowDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateFlowDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateFlowDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateFlowDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetArchivedAt

`func (o *UpdateFlowDto) GetArchivedAt() string`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *UpdateFlowDto) GetArchivedAtOk() (*string, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *UpdateFlowDto) SetArchivedAt(v string)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *UpdateFlowDto) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

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


