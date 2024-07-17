# DataSourceWithIntegrationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Element** | [**DataSourceDto**](DataSourceDto.md) |  | 
**Integration** | [**IntegrationDto**](IntegrationDto.md) |  | 

## Methods

### NewDataSourceWithIntegrationDto

`func NewDataSourceWithIntegrationDto(element DataSourceDto, integration IntegrationDto, ) *DataSourceWithIntegrationDto`

NewDataSourceWithIntegrationDto instantiates a new DataSourceWithIntegrationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataSourceWithIntegrationDtoWithDefaults

`func NewDataSourceWithIntegrationDtoWithDefaults() *DataSourceWithIntegrationDto`

NewDataSourceWithIntegrationDtoWithDefaults instantiates a new DataSourceWithIntegrationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElement

`func (o *DataSourceWithIntegrationDto) GetElement() DataSourceDto`

GetElement returns the Element field if non-nil, zero value otherwise.

### GetElementOk

`func (o *DataSourceWithIntegrationDto) GetElementOk() (*DataSourceDto, bool)`

GetElementOk returns a tuple with the Element field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElement

`func (o *DataSourceWithIntegrationDto) SetElement(v DataSourceDto)`

SetElement sets Element field to given value.


### GetIntegration

`func (o *DataSourceWithIntegrationDto) GetIntegration() IntegrationDto`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *DataSourceWithIntegrationDto) GetIntegrationOk() (*IntegrationDto, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *DataSourceWithIntegrationDto) SetIntegration(v IntegrationDto)`

SetIntegration sets Integration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


