# CreateIntegrationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectorStoreKey** | Pointer to **string** |  | [optional] 
**ConnectorId** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**BaseUri** | Pointer to **string** |  | [optional] 
**IsTest** | Pointer to **bool** |  | [optional] 
**AuthType** | Pointer to **string** |  | [optional] 
**OAuthCallbackUri** | Pointer to **string** |  | [optional] 
**Parameters** | Pointer to **map[string]interface{}** |  | [optional] 
**EncryptedParameters** | Pointer to **map[string]interface{}** |  | [optional] 
**LogoBase64** | Pointer to **string** |  | [optional] 
**LogoUri** | Pointer to **string** |  | [optional] 
**ArchivedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCreateIntegrationDto

`func NewCreateIntegrationDto() *CreateIntegrationDto`

NewCreateIntegrationDto instantiates a new CreateIntegrationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIntegrationDtoWithDefaults

`func NewCreateIntegrationDtoWithDefaults() *CreateIntegrationDto`

NewCreateIntegrationDtoWithDefaults instantiates a new CreateIntegrationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectorStoreKey

`func (o *CreateIntegrationDto) GetConnectorStoreKey() string`

GetConnectorStoreKey returns the ConnectorStoreKey field if non-nil, zero value otherwise.

### GetConnectorStoreKeyOk

`func (o *CreateIntegrationDto) GetConnectorStoreKeyOk() (*string, bool)`

GetConnectorStoreKeyOk returns a tuple with the ConnectorStoreKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorStoreKey

`func (o *CreateIntegrationDto) SetConnectorStoreKey(v string)`

SetConnectorStoreKey sets ConnectorStoreKey field to given value.

### HasConnectorStoreKey

`func (o *CreateIntegrationDto) HasConnectorStoreKey() bool`

HasConnectorStoreKey returns a boolean if a field has been set.

### GetConnectorId

`func (o *CreateIntegrationDto) GetConnectorId() string`

GetConnectorId returns the ConnectorId field if non-nil, zero value otherwise.

### GetConnectorIdOk

`func (o *CreateIntegrationDto) GetConnectorIdOk() (*string, bool)`

GetConnectorIdOk returns a tuple with the ConnectorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorId

`func (o *CreateIntegrationDto) SetConnectorId(v string)`

SetConnectorId sets ConnectorId field to given value.

### HasConnectorId

`func (o *CreateIntegrationDto) HasConnectorId() bool`

HasConnectorId returns a boolean if a field has been set.

### GetKey

`func (o *CreateIntegrationDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateIntegrationDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateIntegrationDto) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CreateIntegrationDto) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *CreateIntegrationDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateIntegrationDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateIntegrationDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateIntegrationDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBaseUri

`func (o *CreateIntegrationDto) GetBaseUri() string`

GetBaseUri returns the BaseUri field if non-nil, zero value otherwise.

### GetBaseUriOk

`func (o *CreateIntegrationDto) GetBaseUriOk() (*string, bool)`

GetBaseUriOk returns a tuple with the BaseUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUri

`func (o *CreateIntegrationDto) SetBaseUri(v string)`

SetBaseUri sets BaseUri field to given value.

### HasBaseUri

`func (o *CreateIntegrationDto) HasBaseUri() bool`

HasBaseUri returns a boolean if a field has been set.

### GetIsTest

`func (o *CreateIntegrationDto) GetIsTest() bool`

GetIsTest returns the IsTest field if non-nil, zero value otherwise.

### GetIsTestOk

`func (o *CreateIntegrationDto) GetIsTestOk() (*bool, bool)`

GetIsTestOk returns a tuple with the IsTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTest

`func (o *CreateIntegrationDto) SetIsTest(v bool)`

SetIsTest sets IsTest field to given value.

### HasIsTest

`func (o *CreateIntegrationDto) HasIsTest() bool`

HasIsTest returns a boolean if a field has been set.

### GetAuthType

`func (o *CreateIntegrationDto) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *CreateIntegrationDto) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *CreateIntegrationDto) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *CreateIntegrationDto) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetOAuthCallbackUri

`func (o *CreateIntegrationDto) GetOAuthCallbackUri() string`

GetOAuthCallbackUri returns the OAuthCallbackUri field if non-nil, zero value otherwise.

### GetOAuthCallbackUriOk

`func (o *CreateIntegrationDto) GetOAuthCallbackUriOk() (*string, bool)`

GetOAuthCallbackUriOk returns a tuple with the OAuthCallbackUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAuthCallbackUri

`func (o *CreateIntegrationDto) SetOAuthCallbackUri(v string)`

SetOAuthCallbackUri sets OAuthCallbackUri field to given value.

### HasOAuthCallbackUri

`func (o *CreateIntegrationDto) HasOAuthCallbackUri() bool`

HasOAuthCallbackUri returns a boolean if a field has been set.

### GetParameters

`func (o *CreateIntegrationDto) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *CreateIntegrationDto) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *CreateIntegrationDto) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *CreateIntegrationDto) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetEncryptedParameters

`func (o *CreateIntegrationDto) GetEncryptedParameters() map[string]interface{}`

GetEncryptedParameters returns the EncryptedParameters field if non-nil, zero value otherwise.

### GetEncryptedParametersOk

`func (o *CreateIntegrationDto) GetEncryptedParametersOk() (*map[string]interface{}, bool)`

GetEncryptedParametersOk returns a tuple with the EncryptedParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedParameters

`func (o *CreateIntegrationDto) SetEncryptedParameters(v map[string]interface{})`

SetEncryptedParameters sets EncryptedParameters field to given value.

### HasEncryptedParameters

`func (o *CreateIntegrationDto) HasEncryptedParameters() bool`

HasEncryptedParameters returns a boolean if a field has been set.

### GetLogoBase64

`func (o *CreateIntegrationDto) GetLogoBase64() string`

GetLogoBase64 returns the LogoBase64 field if non-nil, zero value otherwise.

### GetLogoBase64Ok

`func (o *CreateIntegrationDto) GetLogoBase64Ok() (*string, bool)`

GetLogoBase64Ok returns a tuple with the LogoBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoBase64

`func (o *CreateIntegrationDto) SetLogoBase64(v string)`

SetLogoBase64 sets LogoBase64 field to given value.

### HasLogoBase64

`func (o *CreateIntegrationDto) HasLogoBase64() bool`

HasLogoBase64 returns a boolean if a field has been set.

### GetLogoUri

`func (o *CreateIntegrationDto) GetLogoUri() string`

GetLogoUri returns the LogoUri field if non-nil, zero value otherwise.

### GetLogoUriOk

`func (o *CreateIntegrationDto) GetLogoUriOk() (*string, bool)`

GetLogoUriOk returns a tuple with the LogoUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUri

`func (o *CreateIntegrationDto) SetLogoUri(v string)`

SetLogoUri sets LogoUri field to given value.

### HasLogoUri

`func (o *CreateIntegrationDto) HasLogoUri() bool`

HasLogoUri returns a boolean if a field has been set.

### GetArchivedAt

`func (o *CreateIntegrationDto) GetArchivedAt() time.Time`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *CreateIntegrationDto) GetArchivedAtOk() (*time.Time, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *CreateIntegrationDto) SetArchivedAt(v time.Time)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *CreateIntegrationDto) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


