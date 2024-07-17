# IntegrationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Key** | **string** |  | 
**Uuid** | **string** |  | 
**Name** | **string** |  | 
**ConnectorStoreKey** | Pointer to **string** |  | [optional] 
**ConnectorId** | Pointer to **string** |  | [optional] 
**AuthType** | Pointer to **map[string]interface{}** |  | [optional] 
**AuthOptions** | Pointer to **[]string** |  | [optional] 
**OAuthCallbackUri** | **string** |  | 
**ParametersSchema** | Pointer to **map[string]interface{}** |  | [optional] 
**HasDefaultParameters** | Pointer to **bool** |  | [optional] 
**HasMissingParameters** | Pointer to **bool** |  | [optional] 
**HasDocumentation** | Pointer to **bool** |  | [optional] 
**HasOperations** | Pointer to **bool** |  | [optional] 
**HasData** | Pointer to **bool** |  | [optional] 
**HasEvents** | Pointer to **bool** |  | [optional] 
**HasGlobalWebhooks** | Pointer to **bool** |  | [optional] 
**HasUdm** | Pointer to **bool** |  | [optional] 
**AreParametersCustomized** | Pointer to **bool** |  | [optional] 
**BaseUri** | **string** |  | 
**ConnectorVersion** | Pointer to **string** |  | [optional] 
**IsTest** | Pointer to **bool** |  | [optional] 
**LogoUri** | **string** |  | 
**ArchivedAt** | Pointer to **time.Time** |  | [optional] 
**ConnectionMode** | Pointer to **string** |  | [optional] 

## Methods

### NewIntegrationDto

`func NewIntegrationDto(id string, key string, uuid string, name string, oAuthCallbackUri string, baseUri string, logoUri string, ) *IntegrationDto`

NewIntegrationDto instantiates a new IntegrationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationDtoWithDefaults

`func NewIntegrationDtoWithDefaults() *IntegrationDto`

NewIntegrationDtoWithDefaults instantiates a new IntegrationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IntegrationDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IntegrationDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IntegrationDto) SetId(v string)`

SetId sets Id field to given value.


### GetKey

`func (o *IntegrationDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *IntegrationDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *IntegrationDto) SetKey(v string)`

SetKey sets Key field to given value.


### GetUuid

`func (o *IntegrationDto) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *IntegrationDto) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *IntegrationDto) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetName

`func (o *IntegrationDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IntegrationDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IntegrationDto) SetName(v string)`

SetName sets Name field to given value.


### GetConnectorStoreKey

`func (o *IntegrationDto) GetConnectorStoreKey() string`

GetConnectorStoreKey returns the ConnectorStoreKey field if non-nil, zero value otherwise.

### GetConnectorStoreKeyOk

`func (o *IntegrationDto) GetConnectorStoreKeyOk() (*string, bool)`

GetConnectorStoreKeyOk returns a tuple with the ConnectorStoreKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorStoreKey

`func (o *IntegrationDto) SetConnectorStoreKey(v string)`

SetConnectorStoreKey sets ConnectorStoreKey field to given value.

### HasConnectorStoreKey

`func (o *IntegrationDto) HasConnectorStoreKey() bool`

HasConnectorStoreKey returns a boolean if a field has been set.

### GetConnectorId

`func (o *IntegrationDto) GetConnectorId() string`

GetConnectorId returns the ConnectorId field if non-nil, zero value otherwise.

### GetConnectorIdOk

`func (o *IntegrationDto) GetConnectorIdOk() (*string, bool)`

GetConnectorIdOk returns a tuple with the ConnectorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorId

`func (o *IntegrationDto) SetConnectorId(v string)`

SetConnectorId sets ConnectorId field to given value.

### HasConnectorId

`func (o *IntegrationDto) HasConnectorId() bool`

HasConnectorId returns a boolean if a field has been set.

### GetAuthType

`func (o *IntegrationDto) GetAuthType() map[string]interface{}`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *IntegrationDto) GetAuthTypeOk() (*map[string]interface{}, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *IntegrationDto) SetAuthType(v map[string]interface{})`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *IntegrationDto) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetAuthOptions

`func (o *IntegrationDto) GetAuthOptions() []string`

GetAuthOptions returns the AuthOptions field if non-nil, zero value otherwise.

### GetAuthOptionsOk

`func (o *IntegrationDto) GetAuthOptionsOk() (*[]string, bool)`

GetAuthOptionsOk returns a tuple with the AuthOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthOptions

`func (o *IntegrationDto) SetAuthOptions(v []string)`

SetAuthOptions sets AuthOptions field to given value.

### HasAuthOptions

`func (o *IntegrationDto) HasAuthOptions() bool`

HasAuthOptions returns a boolean if a field has been set.

### GetOAuthCallbackUri

`func (o *IntegrationDto) GetOAuthCallbackUri() string`

GetOAuthCallbackUri returns the OAuthCallbackUri field if non-nil, zero value otherwise.

### GetOAuthCallbackUriOk

`func (o *IntegrationDto) GetOAuthCallbackUriOk() (*string, bool)`

GetOAuthCallbackUriOk returns a tuple with the OAuthCallbackUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAuthCallbackUri

`func (o *IntegrationDto) SetOAuthCallbackUri(v string)`

SetOAuthCallbackUri sets OAuthCallbackUri field to given value.


### GetParametersSchema

`func (o *IntegrationDto) GetParametersSchema() map[string]interface{}`

GetParametersSchema returns the ParametersSchema field if non-nil, zero value otherwise.

### GetParametersSchemaOk

`func (o *IntegrationDto) GetParametersSchemaOk() (*map[string]interface{}, bool)`

GetParametersSchemaOk returns a tuple with the ParametersSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParametersSchema

`func (o *IntegrationDto) SetParametersSchema(v map[string]interface{})`

SetParametersSchema sets ParametersSchema field to given value.

### HasParametersSchema

`func (o *IntegrationDto) HasParametersSchema() bool`

HasParametersSchema returns a boolean if a field has been set.

### GetHasDefaultParameters

`func (o *IntegrationDto) GetHasDefaultParameters() bool`

GetHasDefaultParameters returns the HasDefaultParameters field if non-nil, zero value otherwise.

### GetHasDefaultParametersOk

`func (o *IntegrationDto) GetHasDefaultParametersOk() (*bool, bool)`

GetHasDefaultParametersOk returns a tuple with the HasDefaultParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDefaultParameters

`func (o *IntegrationDto) SetHasDefaultParameters(v bool)`

SetHasDefaultParameters sets HasDefaultParameters field to given value.

### HasHasDefaultParameters

`func (o *IntegrationDto) HasHasDefaultParameters() bool`

HasHasDefaultParameters returns a boolean if a field has been set.

### GetHasMissingParameters

`func (o *IntegrationDto) GetHasMissingParameters() bool`

GetHasMissingParameters returns the HasMissingParameters field if non-nil, zero value otherwise.

### GetHasMissingParametersOk

`func (o *IntegrationDto) GetHasMissingParametersOk() (*bool, bool)`

GetHasMissingParametersOk returns a tuple with the HasMissingParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMissingParameters

`func (o *IntegrationDto) SetHasMissingParameters(v bool)`

SetHasMissingParameters sets HasMissingParameters field to given value.

### HasHasMissingParameters

`func (o *IntegrationDto) HasHasMissingParameters() bool`

HasHasMissingParameters returns a boolean if a field has been set.

### GetHasDocumentation

`func (o *IntegrationDto) GetHasDocumentation() bool`

GetHasDocumentation returns the HasDocumentation field if non-nil, zero value otherwise.

### GetHasDocumentationOk

`func (o *IntegrationDto) GetHasDocumentationOk() (*bool, bool)`

GetHasDocumentationOk returns a tuple with the HasDocumentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDocumentation

`func (o *IntegrationDto) SetHasDocumentation(v bool)`

SetHasDocumentation sets HasDocumentation field to given value.

### HasHasDocumentation

`func (o *IntegrationDto) HasHasDocumentation() bool`

HasHasDocumentation returns a boolean if a field has been set.

### GetHasOperations

`func (o *IntegrationDto) GetHasOperations() bool`

GetHasOperations returns the HasOperations field if non-nil, zero value otherwise.

### GetHasOperationsOk

`func (o *IntegrationDto) GetHasOperationsOk() (*bool, bool)`

GetHasOperationsOk returns a tuple with the HasOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasOperations

`func (o *IntegrationDto) SetHasOperations(v bool)`

SetHasOperations sets HasOperations field to given value.

### HasHasOperations

`func (o *IntegrationDto) HasHasOperations() bool`

HasHasOperations returns a boolean if a field has been set.

### GetHasData

`func (o *IntegrationDto) GetHasData() bool`

GetHasData returns the HasData field if non-nil, zero value otherwise.

### GetHasDataOk

`func (o *IntegrationDto) GetHasDataOk() (*bool, bool)`

GetHasDataOk returns a tuple with the HasData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasData

`func (o *IntegrationDto) SetHasData(v bool)`

SetHasData sets HasData field to given value.

### HasHasData

`func (o *IntegrationDto) HasHasData() bool`

HasHasData returns a boolean if a field has been set.

### GetHasEvents

`func (o *IntegrationDto) GetHasEvents() bool`

GetHasEvents returns the HasEvents field if non-nil, zero value otherwise.

### GetHasEventsOk

`func (o *IntegrationDto) GetHasEventsOk() (*bool, bool)`

GetHasEventsOk returns a tuple with the HasEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasEvents

`func (o *IntegrationDto) SetHasEvents(v bool)`

SetHasEvents sets HasEvents field to given value.

### HasHasEvents

`func (o *IntegrationDto) HasHasEvents() bool`

HasHasEvents returns a boolean if a field has been set.

### GetHasGlobalWebhooks

`func (o *IntegrationDto) GetHasGlobalWebhooks() bool`

GetHasGlobalWebhooks returns the HasGlobalWebhooks field if non-nil, zero value otherwise.

### GetHasGlobalWebhooksOk

`func (o *IntegrationDto) GetHasGlobalWebhooksOk() (*bool, bool)`

GetHasGlobalWebhooksOk returns a tuple with the HasGlobalWebhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasGlobalWebhooks

`func (o *IntegrationDto) SetHasGlobalWebhooks(v bool)`

SetHasGlobalWebhooks sets HasGlobalWebhooks field to given value.

### HasHasGlobalWebhooks

`func (o *IntegrationDto) HasHasGlobalWebhooks() bool`

HasHasGlobalWebhooks returns a boolean if a field has been set.

### GetHasUdm

`func (o *IntegrationDto) GetHasUdm() bool`

GetHasUdm returns the HasUdm field if non-nil, zero value otherwise.

### GetHasUdmOk

`func (o *IntegrationDto) GetHasUdmOk() (*bool, bool)`

GetHasUdmOk returns a tuple with the HasUdm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasUdm

`func (o *IntegrationDto) SetHasUdm(v bool)`

SetHasUdm sets HasUdm field to given value.

### HasHasUdm

`func (o *IntegrationDto) HasHasUdm() bool`

HasHasUdm returns a boolean if a field has been set.

### GetAreParametersCustomized

`func (o *IntegrationDto) GetAreParametersCustomized() bool`

GetAreParametersCustomized returns the AreParametersCustomized field if non-nil, zero value otherwise.

### GetAreParametersCustomizedOk

`func (o *IntegrationDto) GetAreParametersCustomizedOk() (*bool, bool)`

GetAreParametersCustomizedOk returns a tuple with the AreParametersCustomized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreParametersCustomized

`func (o *IntegrationDto) SetAreParametersCustomized(v bool)`

SetAreParametersCustomized sets AreParametersCustomized field to given value.

### HasAreParametersCustomized

`func (o *IntegrationDto) HasAreParametersCustomized() bool`

HasAreParametersCustomized returns a boolean if a field has been set.

### GetBaseUri

`func (o *IntegrationDto) GetBaseUri() string`

GetBaseUri returns the BaseUri field if non-nil, zero value otherwise.

### GetBaseUriOk

`func (o *IntegrationDto) GetBaseUriOk() (*string, bool)`

GetBaseUriOk returns a tuple with the BaseUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUri

`func (o *IntegrationDto) SetBaseUri(v string)`

SetBaseUri sets BaseUri field to given value.


### GetConnectorVersion

`func (o *IntegrationDto) GetConnectorVersion() string`

GetConnectorVersion returns the ConnectorVersion field if non-nil, zero value otherwise.

### GetConnectorVersionOk

`func (o *IntegrationDto) GetConnectorVersionOk() (*string, bool)`

GetConnectorVersionOk returns a tuple with the ConnectorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorVersion

`func (o *IntegrationDto) SetConnectorVersion(v string)`

SetConnectorVersion sets ConnectorVersion field to given value.

### HasConnectorVersion

`func (o *IntegrationDto) HasConnectorVersion() bool`

HasConnectorVersion returns a boolean if a field has been set.

### GetIsTest

`func (o *IntegrationDto) GetIsTest() bool`

GetIsTest returns the IsTest field if non-nil, zero value otherwise.

### GetIsTestOk

`func (o *IntegrationDto) GetIsTestOk() (*bool, bool)`

GetIsTestOk returns a tuple with the IsTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTest

`func (o *IntegrationDto) SetIsTest(v bool)`

SetIsTest sets IsTest field to given value.

### HasIsTest

`func (o *IntegrationDto) HasIsTest() bool`

HasIsTest returns a boolean if a field has been set.

### GetLogoUri

`func (o *IntegrationDto) GetLogoUri() string`

GetLogoUri returns the LogoUri field if non-nil, zero value otherwise.

### GetLogoUriOk

`func (o *IntegrationDto) GetLogoUriOk() (*string, bool)`

GetLogoUriOk returns a tuple with the LogoUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUri

`func (o *IntegrationDto) SetLogoUri(v string)`

SetLogoUri sets LogoUri field to given value.


### GetArchivedAt

`func (o *IntegrationDto) GetArchivedAt() time.Time`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *IntegrationDto) GetArchivedAtOk() (*time.Time, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *IntegrationDto) SetArchivedAt(v time.Time)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *IntegrationDto) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### GetConnectionMode

`func (o *IntegrationDto) GetConnectionMode() string`

GetConnectionMode returns the ConnectionMode field if non-nil, zero value otherwise.

### GetConnectionModeOk

`func (o *IntegrationDto) GetConnectionModeOk() (*string, bool)`

GetConnectionModeOk returns a tuple with the ConnectionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionMode

`func (o *IntegrationDto) SetConnectionMode(v string)`

SetConnectionMode sets ConnectionMode field to given value.

### HasConnectionMode

`func (o *IntegrationDto) HasConnectionMode() bool`

HasConnectionMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


