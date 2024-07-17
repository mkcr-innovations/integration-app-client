# UpdateFlowInstanceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Parameters** | Pointer to **map[string]interface{}** |  | [optional] 
**Nodes** | Pointer to **map[string]interface{}** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**ArchivedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewUpdateFlowInstanceDto

`func NewUpdateFlowInstanceDto() *UpdateFlowInstanceDto`

NewUpdateFlowInstanceDto instantiates a new UpdateFlowInstanceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFlowInstanceDtoWithDefaults

`func NewUpdateFlowInstanceDtoWithDefaults() *UpdateFlowInstanceDto`

NewUpdateFlowInstanceDtoWithDefaults instantiates a new UpdateFlowInstanceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateFlowInstanceDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateFlowInstanceDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateFlowInstanceDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateFlowInstanceDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParameters

`func (o *UpdateFlowInstanceDto) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *UpdateFlowInstanceDto) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *UpdateFlowInstanceDto) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *UpdateFlowInstanceDto) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetNodes

`func (o *UpdateFlowInstanceDto) GetNodes() map[string]interface{}`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *UpdateFlowInstanceDto) GetNodesOk() (*map[string]interface{}, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *UpdateFlowInstanceDto) SetNodes(v map[string]interface{})`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *UpdateFlowInstanceDto) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateFlowInstanceDto) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateFlowInstanceDto) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateFlowInstanceDto) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateFlowInstanceDto) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetArchivedAt

`func (o *UpdateFlowInstanceDto) GetArchivedAt() time.Time`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *UpdateFlowInstanceDto) GetArchivedAtOk() (*time.Time, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *UpdateFlowInstanceDto) SetArchivedAt(v time.Time)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *UpdateFlowInstanceDto) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


