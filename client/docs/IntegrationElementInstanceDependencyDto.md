# IntegrationElementInstanceDependencyDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**InstanceId** | Pointer to **string** |  | [optional] 
**Instance** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewIntegrationElementInstanceDependencyDto

`func NewIntegrationElementInstanceDependencyDto(type_ string, ) *IntegrationElementInstanceDependencyDto`

NewIntegrationElementInstanceDependencyDto instantiates a new IntegrationElementInstanceDependencyDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationElementInstanceDependencyDtoWithDefaults

`func NewIntegrationElementInstanceDependencyDtoWithDefaults() *IntegrationElementInstanceDependencyDto`

NewIntegrationElementInstanceDependencyDtoWithDefaults instantiates a new IntegrationElementInstanceDependencyDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *IntegrationElementInstanceDependencyDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IntegrationElementInstanceDependencyDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IntegrationElementInstanceDependencyDto) SetType(v string)`

SetType sets Type field to given value.


### GetInstanceId

`func (o *IntegrationElementInstanceDependencyDto) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *IntegrationElementInstanceDependencyDto) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *IntegrationElementInstanceDependencyDto) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *IntegrationElementInstanceDependencyDto) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetInstance

`func (o *IntegrationElementInstanceDependencyDto) GetInstance() map[string]interface{}`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *IntegrationElementInstanceDependencyDto) GetInstanceOk() (*map[string]interface{}, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *IntegrationElementInstanceDependencyDto) SetInstance(v map[string]interface{})`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *IntegrationElementInstanceDependencyDto) HasInstance() bool`

HasInstance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


