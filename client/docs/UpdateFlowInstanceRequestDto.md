# UpdateFlowInstanceRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Parameters** | Pointer to **map[string]interface{}** |  | [optional] 
**Nodes** | Pointer to **map[string]interface{}** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**ArchivedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewUpdateFlowInstanceRequestDto

`func NewUpdateFlowInstanceRequestDto() *UpdateFlowInstanceRequestDto`

NewUpdateFlowInstanceRequestDto instantiates a new UpdateFlowInstanceRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFlowInstanceRequestDtoWithDefaults

`func NewUpdateFlowInstanceRequestDtoWithDefaults() *UpdateFlowInstanceRequestDto`

NewUpdateFlowInstanceRequestDtoWithDefaults instantiates a new UpdateFlowInstanceRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateFlowInstanceRequestDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateFlowInstanceRequestDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateFlowInstanceRequestDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateFlowInstanceRequestDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParameters

`func (o *UpdateFlowInstanceRequestDto) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *UpdateFlowInstanceRequestDto) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *UpdateFlowInstanceRequestDto) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *UpdateFlowInstanceRequestDto) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetNodes

`func (o *UpdateFlowInstanceRequestDto) GetNodes() map[string]interface{}`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *UpdateFlowInstanceRequestDto) GetNodesOk() (*map[string]interface{}, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *UpdateFlowInstanceRequestDto) SetNodes(v map[string]interface{})`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *UpdateFlowInstanceRequestDto) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateFlowInstanceRequestDto) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateFlowInstanceRequestDto) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateFlowInstanceRequestDto) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateFlowInstanceRequestDto) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetArchivedAt

`func (o *UpdateFlowInstanceRequestDto) GetArchivedAt() time.Time`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *UpdateFlowInstanceRequestDto) GetArchivedAtOk() (*time.Time, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *UpdateFlowInstanceRequestDto) SetArchivedAt(v time.Time)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *UpdateFlowInstanceRequestDto) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


