# UpdateIntegrationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewUpdateIntegrationDto

`func NewUpdateIntegrationDto() *UpdateIntegrationDto`

NewUpdateIntegrationDto instantiates a new UpdateIntegrationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateIntegrationDtoWithDefaults

`func NewUpdateIntegrationDtoWithDefaults() *UpdateIntegrationDto`

NewUpdateIntegrationDtoWithDefaults instantiates a new UpdateIntegrationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *UpdateIntegrationDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UpdateIntegrationDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UpdateIntegrationDto) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *UpdateIntegrationDto) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *UpdateIntegrationDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateIntegrationDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateIntegrationDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateIntegrationDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBaseUri

`func (o *UpdateIntegrationDto) GetBaseUri() string`

GetBaseUri returns the BaseUri field if non-nil, zero value otherwise.

### GetBaseUriOk

`func (o *UpdateIntegrationDto) GetBaseUriOk() (*string, bool)`

GetBaseUriOk returns a tuple with the BaseUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUri

`func (o *UpdateIntegrationDto) SetBaseUri(v string)`

SetBaseUri sets BaseUri field to given value.

### HasBaseUri

`func (o *UpdateIntegrationDto) HasBaseUri() bool`

HasBaseUri returns a boolean if a field has been set.

### GetIsTest

`func (o *UpdateIntegrationDto) GetIsTest() bool`

GetIsTest returns the IsTest field if non-nil, zero value otherwise.

### GetIsTestOk

`func (o *UpdateIntegrationDto) GetIsTestOk() (*bool, bool)`

GetIsTestOk returns a tuple with the IsTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTest

`func (o *UpdateIntegrationDto) SetIsTest(v bool)`

SetIsTest sets IsTest field to given value.

### HasIsTest

`func (o *UpdateIntegrationDto) HasIsTest() bool`

HasIsTest returns a boolean if a field has been set.

### GetAuthType

`func (o *UpdateIntegrationDto) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *UpdateIntegrationDto) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *UpdateIntegrationDto) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *UpdateIntegrationDto) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetOAuthCallbackUri

`func (o *UpdateIntegrationDto) GetOAuthCallbackUri() string`

GetOAuthCallbackUri returns the OAuthCallbackUri field if non-nil, zero value otherwise.

### GetOAuthCallbackUriOk

`func (o *UpdateIntegrationDto) GetOAuthCallbackUriOk() (*string, bool)`

GetOAuthCallbackUriOk returns a tuple with the OAuthCallbackUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAuthCallbackUri

`func (o *UpdateIntegrationDto) SetOAuthCallbackUri(v string)`

SetOAuthCallbackUri sets OAuthCallbackUri field to given value.

### HasOAuthCallbackUri

`func (o *UpdateIntegrationDto) HasOAuthCallbackUri() bool`

HasOAuthCallbackUri returns a boolean if a field has been set.

### GetParameters

`func (o *UpdateIntegrationDto) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *UpdateIntegrationDto) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *UpdateIntegrationDto) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *UpdateIntegrationDto) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetEncryptedParameters

`func (o *UpdateIntegrationDto) GetEncryptedParameters() map[string]interface{}`

GetEncryptedParameters returns the EncryptedParameters field if non-nil, zero value otherwise.

### GetEncryptedParametersOk

`func (o *UpdateIntegrationDto) GetEncryptedParametersOk() (*map[string]interface{}, bool)`

GetEncryptedParametersOk returns a tuple with the EncryptedParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedParameters

`func (o *UpdateIntegrationDto) SetEncryptedParameters(v map[string]interface{})`

SetEncryptedParameters sets EncryptedParameters field to given value.

### HasEncryptedParameters

`func (o *UpdateIntegrationDto) HasEncryptedParameters() bool`

HasEncryptedParameters returns a boolean if a field has been set.

### GetLogoBase64

`func (o *UpdateIntegrationDto) GetLogoBase64() string`

GetLogoBase64 returns the LogoBase64 field if non-nil, zero value otherwise.

### GetLogoBase64Ok

`func (o *UpdateIntegrationDto) GetLogoBase64Ok() (*string, bool)`

GetLogoBase64Ok returns a tuple with the LogoBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoBase64

`func (o *UpdateIntegrationDto) SetLogoBase64(v string)`

SetLogoBase64 sets LogoBase64 field to given value.

### HasLogoBase64

`func (o *UpdateIntegrationDto) HasLogoBase64() bool`

HasLogoBase64 returns a boolean if a field has been set.

### GetLogoUri

`func (o *UpdateIntegrationDto) GetLogoUri() string`

GetLogoUri returns the LogoUri field if non-nil, zero value otherwise.

### GetLogoUriOk

`func (o *UpdateIntegrationDto) GetLogoUriOk() (*string, bool)`

GetLogoUriOk returns a tuple with the LogoUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUri

`func (o *UpdateIntegrationDto) SetLogoUri(v string)`

SetLogoUri sets LogoUri field to given value.

### HasLogoUri

`func (o *UpdateIntegrationDto) HasLogoUri() bool`

HasLogoUri returns a boolean if a field has been set.

### GetArchivedAt

`func (o *UpdateIntegrationDto) GetArchivedAt() time.Time`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *UpdateIntegrationDto) GetArchivedAtOk() (*time.Time, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *UpdateIntegrationDto) SetArchivedAt(v time.Time)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *UpdateIntegrationDto) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


