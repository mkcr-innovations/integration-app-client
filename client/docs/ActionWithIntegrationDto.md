# ActionWithIntegrationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Element** | Pointer to [**ActionDto**](ActionDto.md) |  | [optional] 
**Integration** | Pointer to [**IntegrationDto**](IntegrationDto.md) |  | [optional] 

## Methods

### NewActionWithIntegrationDto

`func NewActionWithIntegrationDto() *ActionWithIntegrationDto`

NewActionWithIntegrationDto instantiates a new ActionWithIntegrationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionWithIntegrationDtoWithDefaults

`func NewActionWithIntegrationDtoWithDefaults() *ActionWithIntegrationDto`

NewActionWithIntegrationDtoWithDefaults instantiates a new ActionWithIntegrationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElement

`func (o *ActionWithIntegrationDto) GetElement() ActionDto`

GetElement returns the Element field if non-nil, zero value otherwise.

### GetElementOk

`func (o *ActionWithIntegrationDto) GetElementOk() (*ActionDto, bool)`

GetElementOk returns a tuple with the Element field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElement

`func (o *ActionWithIntegrationDto) SetElement(v ActionDto)`

SetElement sets Element field to given value.

### HasElement

`func (o *ActionWithIntegrationDto) HasElement() bool`

HasElement returns a boolean if a field has been set.

### GetIntegration

`func (o *ActionWithIntegrationDto) GetIntegration() IntegrationDto`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *ActionWithIntegrationDto) GetIntegrationOk() (*IntegrationDto, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *ActionWithIntegrationDto) SetIntegration(v IntegrationDto)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *ActionWithIntegrationDto) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


