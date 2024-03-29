# PublishConnectorToIntegrationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectorId** | **string** |  | 
**Parameters** | Pointer to **map[string]interface{}** |  | [optional] 
**EncryptedParameters** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPublishConnectorToIntegrationDto

`func NewPublishConnectorToIntegrationDto(connectorId string, ) *PublishConnectorToIntegrationDto`

NewPublishConnectorToIntegrationDto instantiates a new PublishConnectorToIntegrationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublishConnectorToIntegrationDtoWithDefaults

`func NewPublishConnectorToIntegrationDtoWithDefaults() *PublishConnectorToIntegrationDto`

NewPublishConnectorToIntegrationDtoWithDefaults instantiates a new PublishConnectorToIntegrationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectorId

`func (o *PublishConnectorToIntegrationDto) GetConnectorId() string`

GetConnectorId returns the ConnectorId field if non-nil, zero value otherwise.

### GetConnectorIdOk

`func (o *PublishConnectorToIntegrationDto) GetConnectorIdOk() (*string, bool)`

GetConnectorIdOk returns a tuple with the ConnectorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorId

`func (o *PublishConnectorToIntegrationDto) SetConnectorId(v string)`

SetConnectorId sets ConnectorId field to given value.


### GetParameters

`func (o *PublishConnectorToIntegrationDto) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *PublishConnectorToIntegrationDto) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *PublishConnectorToIntegrationDto) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *PublishConnectorToIntegrationDto) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetEncryptedParameters

`func (o *PublishConnectorToIntegrationDto) GetEncryptedParameters() map[string]interface{}`

GetEncryptedParameters returns the EncryptedParameters field if non-nil, zero value otherwise.

### GetEncryptedParametersOk

`func (o *PublishConnectorToIntegrationDto) GetEncryptedParametersOk() (*map[string]interface{}, bool)`

GetEncryptedParametersOk returns a tuple with the EncryptedParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedParameters

`func (o *PublishConnectorToIntegrationDto) SetEncryptedParameters(v map[string]interface{})`

SetEncryptedParameters sets EncryptedParameters field to given value.

### HasEncryptedParameters

`func (o *PublishConnectorToIntegrationDto) HasEncryptedParameters() bool`

HasEncryptedParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


