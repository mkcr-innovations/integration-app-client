/*
Integration.app API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the IntegrationDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IntegrationDto{}

// IntegrationDto struct for IntegrationDto
type IntegrationDto struct {
	Id string `json:"id"`
	Key string `json:"key"`
	Uuid string `json:"uuid"`
	Name string `json:"name"`
	ConnectorStoreKey *string `json:"connectorStoreKey,omitempty"`
	ConnectorId *string `json:"connectorId,omitempty"`
	AuthType map[string]interface{} `json:"authType,omitempty"`
	AuthOptions []string `json:"authOptions,omitempty"`
	OAuthCallbackUri string `json:"oAuthCallbackUri"`
	ParametersSchema map[string]interface{} `json:"parametersSchema,omitempty"`
	HasDefaultParameters *bool `json:"hasDefaultParameters,omitempty"`
	HasMissingParameters *bool `json:"hasMissingParameters,omitempty"`
	HasDocumentation *bool `json:"hasDocumentation,omitempty"`
	HasOperations *bool `json:"hasOperations,omitempty"`
	HasData *bool `json:"hasData,omitempty"`
	HasEvents *bool `json:"hasEvents,omitempty"`
	HasGlobalWebhooks *bool `json:"hasGlobalWebhooks,omitempty"`
	HasUdm *bool `json:"hasUdm,omitempty"`
	AreParametersCustomized *bool `json:"areParametersCustomized,omitempty"`
	BaseUri string `json:"baseUri"`
	ConnectorVersion *string `json:"connectorVersion,omitempty"`
	IsTest *bool `json:"isTest,omitempty"`
	LogoUri string `json:"logoUri"`
	ArchivedAt *time.Time `json:"archivedAt,omitempty"`
	// Deprecated
	ConnectionMode *string `json:"connectionMode,omitempty"`
}

type _IntegrationDto IntegrationDto

// NewIntegrationDto instantiates a new IntegrationDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationDto(id string, key string, uuid string, name string, oAuthCallbackUri string, baseUri string, logoUri string) *IntegrationDto {
	this := IntegrationDto{}
	this.Id = id
	this.Key = key
	this.Uuid = uuid
	this.Name = name
	this.OAuthCallbackUri = oAuthCallbackUri
	this.BaseUri = baseUri
	this.LogoUri = logoUri
	return &this
}

// NewIntegrationDtoWithDefaults instantiates a new IntegrationDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationDtoWithDefaults() *IntegrationDto {
	this := IntegrationDto{}
	return &this
}

// GetId returns the Id field value
func (o *IntegrationDto) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IntegrationDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *IntegrationDto) SetId(v string) {
	o.Id = v
}

// GetKey returns the Key field value
func (o *IntegrationDto) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *IntegrationDto) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *IntegrationDto) SetKey(v string) {
	o.Key = v
}

// GetUuid returns the Uuid field value
func (o *IntegrationDto) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *IntegrationDto) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *IntegrationDto) SetUuid(v string) {
	o.Uuid = v
}

// GetName returns the Name field value
func (o *IntegrationDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IntegrationDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IntegrationDto) SetName(v string) {
	o.Name = v
}

// GetConnectorStoreKey returns the ConnectorStoreKey field value if set, zero value otherwise.
func (o *IntegrationDto) GetConnectorStoreKey() string {
	if o == nil || IsNil(o.ConnectorStoreKey) {
		var ret string
		return ret
	}
	return *o.ConnectorStoreKey
}

// GetConnectorStoreKeyOk returns a tuple with the ConnectorStoreKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationDto) GetConnectorStoreKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectorStoreKey) {
		return nil, false
	}
	return o.ConnectorStoreKey, true
}

// HasConnectorStoreKey returns a boolean if a field has been set.
func (o *IntegrationDto) HasConnectorStoreKey() bool {
	if o != nil && !IsNil(o.ConnectorStoreKey) {
		return true
	}

	return false
}

// SetConnectorStoreKey gets a reference to the given string and assigns it to the ConnectorStoreKey field.
func (o *IntegrationDto) SetConnectorStoreKey(v string) {
	o.ConnectorStoreKey = &v
}

// GetConnectorId returns the ConnectorId field value if set, zero value otherwise.
func (o *IntegrationDto) GetConnectorId() string {
	if o == nil || IsNil(o.ConnectorId) {
		var ret string
		return ret
	}
	return *o.ConnectorId
}

// GetConnectorIdOk returns a tuple with the ConnectorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationDto) GetConnectorIdOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectorId) {
		return nil, false
	}
	return o.ConnectorId, true
}

// HasConnectorId returns a boolean if a field has been set.
func (o *IntegrationDto) HasConnectorId() bool {
	if o != nil && !IsNil(o.ConnectorId) {
		return true
	}

	return false
}

// SetConnectorId gets a reference to the given string and assigns it to the ConnectorId field.
func (o *IntegrationDto) SetConnectorId(v string) {
	o.ConnectorId = &v
}

// GetAuthType returns the AuthType field value if set, zero value otherwise.
func (o *IntegrationDto) GetAuthType() map[string]interface{} {
	if o == nil || IsNil(o.AuthType) {
		var ret map[string]interface{}
		return ret
	}
	return o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationDto) GetAuthTypeOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AuthType) {
		return map[string]interface{}{}, false
	}
	return o.AuthType, true
}

// HasAuthType returns a boolean if a field has been set.
func (o *IntegrationDto) HasAuthType() bool {
	if o != nil && !IsNil(o.AuthType) {
		return true
	}

	return false
}

// SetAuthType gets a reference to the given map[string]interface{} and assigns it to the AuthType field.
func (o *IntegrationDto) SetAuthType(v map[string]interface{}) {
	o.AuthType = v
}

// GetAuthOptions returns the AuthOptions field value if set, zero value otherwise.
func (o *IntegrationDto) GetAuthOptions() []string {
	if o == nil || IsNil(o.AuthOptions) {
		var ret []string
		return ret
	}
	return o.AuthOptions
}

// GetAuthOptionsOk returns a tuple with the AuthOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationDto) GetAuthOptionsOk() ([]string, bool) {
	if o == nil || IsNil(o.AuthOptions) {
		return nil, false
	}
	return o.AuthOptions, true
}

// HasAuthOptions returns a boolean if a field has been set.
func (o *IntegrationDto) HasAuthOptions() bool {
	if o != nil && !IsNil(o.AuthOptions) {
		return true
	}

	return false
}

// SetAuthOptions gets a reference to the given []string and assigns it to the AuthOptions field.
func (o *IntegrationDto) SetAuthOptions(v []string) {
	o.AuthOptions = v
}

// GetOAuthCallbackUri returns the OAuthCallbackUri field value
func (o *IntegrationDto) GetOAuthCallbackUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OAuthCallbackUri
}

// GetOAuthCallbackUriOk returns a tuple with the OAuthCallbackUri field value
// and a boolean to check if the value has been set.
func (o *IntegrationDto) GetOAuthCallbackUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OAuthCallbackUri, true
}

// SetOAuthCallbackUri sets field value
func (o *IntegrationDto) SetOAuthCallbackUri(v string) {
	o.OAuthCallbackUri = v
}

// GetParametersSchema returns the ParametersSchema field value if set, zero value otherwise.
func (o *IntegrationDto) GetParametersSchema() map[string]interface{} {
	if o == nil || IsNil(o.ParametersSchema) {
		var ret map[string]interface{}
		return ret
	}
	return o.ParametersSchema
}

// GetParametersSchemaOk returns a tuple with the ParametersSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationDto) GetParametersSchemaOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ParametersSchema) {
		return map[string]interface{}{}, false
	}
	return o.ParametersSchema, true
}

// HasParametersSchema returns a boolean if a field has been set.
func (o *IntegrationDto) HasParametersSchema() bool {
	if o != nil && !IsNil(o.ParametersSchema) {
		return true
	}

	return false
}

// SetParametersSchema gets a reference to the given map[string]interface{} and assigns it to the ParametersSchema field.
func (o *IntegrationDto) SetParametersSchema(v map[string]interface{}) {
	o.ParametersSchema = v
}

// GetHasDefaultParameters returns the HasDefaultParameters field value if set, zero value otherwise.
func (o *IntegrationDto) GetHasDefaultParameters() bool {
	if o == nil || IsNil(o.HasDefaultParameters) {
		var ret bool
		return ret
	}
	return *o.HasDefaultParameters
}

// GetHasDefaultParametersOk returns a tuple with the HasDefaultParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationDto) GetHasDefaultParametersOk() (*bool, bool) {
	if o == nil || IsNil(o.HasDefaultParameters) {
		return nil, false
	}
	return o.HasDefaultParameters, true
}

// HasHasDefaultParameters returns a boolean if a field has been set.
func (o *IntegrationDto) HasHasDefaultParameters() bool {
	if o != nil && !IsNil(o.HasDefaultParameters) {
		return true
	}

	return false
}

// SetHasDefaultParameters gets a reference to the given bool and assigns it to the HasDefaultParameters field.
func (o *IntegrationDto) SetHasDefaultParameters(v bool) {
	o.HasDefaultParameters = &v
}

// GetHasMissingParameters returns the HasMissingParameters field value if set, zero value otherwise.
func (o *IntegrationDto) GetHasMissingParameters() bool {
	if o == nil || IsNil(o.HasMissingParameters) {
		var ret bool
		return ret
	}
	return *o.HasMissingParameters
}

// GetHasMissingParametersOk returns a tuple with the HasMissingParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationDto) GetHasMissingParametersOk() (*bool, bool) {
	if o == nil || IsNil(o.HasMissingParameters) {
		return nil, false
	}
	return o.HasMissingParameters, true
}

// HasHasMissingParameters returns a boolean if a field has been set.
func (o *IntegrationDto) HasHasMissingParameters() bool {
	if o != nil && !IsNil(o.HasMissingParameters) {
		return true
	}

	return false
}

// SetHasMissingParameters gets a reference to the given bool and assigns it to the HasMissingParameters field.
func (o *IntegrationDto) SetHasMissingParameters(v bool) {
	o.HasMissingParameters = &v
}

// GetHasDocumentation returns the HasDocumentation field value if set, zero value otherwise.
func (o *IntegrationDto) GetHasDocumentation() bool {
	if o == nil || IsNil(o.HasDocumentation) {
		var ret bool
		return ret
	}
	return *o.HasDocumentation
}

// GetHasDocumentationOk returns a tuple with the HasDocumentation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationDto) GetHasDocumentationOk() (*bool, bool) {
	if o == nil || IsNil(o.HasDocumentation) {
		return nil, false
	}
	return o.HasDocumentation, true
}

// HasHasDocumentation returns a boolean if a field has been set.
func (o *IntegrationDto) HasHasDocumentation() bool {
	if o != nil && !IsNil(o.HasDocumentation) {
		return true
	}

	return false
}

// SetHasDocumentation gets a reference to the given bool and assigns it to the HasDocumentation field.
func (o *IntegrationDto) SetHasDocumentation(v bool) {
	o.HasDocumentation = &v
}

// GetHasOperations returns the HasOperations field value if set, zero value otherwise.
func (o *IntegrationDto) GetHasOperations() bool {
	if o == nil || IsNil(o.HasOperations) {
		var ret bool
		return ret
	}
	return *o.HasOperations
}

// GetHasOperationsOk returns a tuple with the HasOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationDto) GetHasOperationsOk() (*bool, bool) {
	if o == nil || IsNil(o.HasOperations) {
		return nil, false
	}
	return o.HasOperations, true
}

// HasHasOperations returns a boolean if a field has been set.
func (o *IntegrationDto) HasHasOperations() bool {
	if o != nil && !IsNil(o.HasOperations) {
		return true
	}

	return false
}

// SetHasOperations gets a reference to the given bool and assigns it to the HasOperations field.
func (o *IntegrationDto) SetHasOperations(v bool) {
	o.HasOperations = &v
}

// GetHasData returns the HasData field value if set, zero value otherwise.
func (o *IntegrationDto) GetHasData() bool {
	if o == nil || IsNil(o.HasData) {
		var ret bool
		return ret
	}
	return *o.HasData
}

// GetHasDataOk returns a tuple with the HasData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationDto) GetHasDataOk() (*bool, bool) {
	if o == nil || IsNil(o.HasData) {
		return nil, false
	}
	return o.HasData, true
}

// HasHasData returns a boolean if a field has been set.
func (o *IntegrationDto) HasHasData() bool {
	if o != nil && !IsNil(o.HasData) {
		return true
	}

	return false
}

// SetHasData gets a reference to the given bool and assigns it to the HasData field.
func (o *IntegrationDto) SetHasData(v bool) {
	o.HasData = &v
}

// GetHasEvents returns the HasEvents field value if set, zero value otherwise.
func (o *IntegrationDto) GetHasEvents() bool {
	if o == nil || IsNil(o.HasEvents) {
		var ret bool
		return ret
	}
	return *o.HasEvents
}

// GetHasEventsOk returns a tuple with the HasEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationDto) GetHasEventsOk() (*bool, bool) {
	if o == nil || IsNil(o.HasEvents) {
		return nil, false
	}
	return o.HasEvents, true
}

// HasHasEvents returns a boolean if a field has been set.
func (o *IntegrationDto) HasHasEvents() bool {
	if o != nil && !IsNil(o.HasEvents) {
		return true
	}

	return false
}

// SetHasEvents gets a reference to the given bool and assigns it to the HasEvents field.
func (o *IntegrationDto) SetHasEvents(v bool) {
	o.HasEvents = &v
}

// GetHasGlobalWebhooks returns the HasGlobalWebhooks field value if set, zero value otherwise.
func (o *IntegrationDto) GetHasGlobalWebhooks() bool {
	if o == nil || IsNil(o.HasGlobalWebhooks) {
		var ret bool
		return ret
	}
	return *o.HasGlobalWebhooks
}

// GetHasGlobalWebhooksOk returns a tuple with the HasGlobalWebhooks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationDto) GetHasGlobalWebhooksOk() (*bool, bool) {
	if o == nil || IsNil(o.HasGlobalWebhooks) {
		return nil, false
	}
	return o.HasGlobalWebhooks, true
}

// HasHasGlobalWebhooks returns a boolean if a field has been set.
func (o *IntegrationDto) HasHasGlobalWebhooks() bool {
	if o != nil && !IsNil(o.HasGlobalWebhooks) {
		return true
	}

	return false
}

// SetHasGlobalWebhooks gets a reference to the given bool and assigns it to the HasGlobalWebhooks field.
func (o *IntegrationDto) SetHasGlobalWebhooks(v bool) {
	o.HasGlobalWebhooks = &v
}

// GetHasUdm returns the HasUdm field value if set, zero value otherwise.
func (o *IntegrationDto) GetHasUdm() bool {
	if o == nil || IsNil(o.HasUdm) {
		var ret bool
		return ret
	}
	return *o.HasUdm
}

// GetHasUdmOk returns a tuple with the HasUdm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationDto) GetHasUdmOk() (*bool, bool) {
	if o == nil || IsNil(o.HasUdm) {
		return nil, false
	}
	return o.HasUdm, true
}

// HasHasUdm returns a boolean if a field has been set.
func (o *IntegrationDto) HasHasUdm() bool {
	if o != nil && !IsNil(o.HasUdm) {
		return true
	}

	return false
}

// SetHasUdm gets a reference to the given bool and assigns it to the HasUdm field.
func (o *IntegrationDto) SetHasUdm(v bool) {
	o.HasUdm = &v
}

// GetAreParametersCustomized returns the AreParametersCustomized field value if set, zero value otherwise.
func (o *IntegrationDto) GetAreParametersCustomized() bool {
	if o == nil || IsNil(o.AreParametersCustomized) {
		var ret bool
		return ret
	}
	return *o.AreParametersCustomized
}

// GetAreParametersCustomizedOk returns a tuple with the AreParametersCustomized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationDto) GetAreParametersCustomizedOk() (*bool, bool) {
	if o == nil || IsNil(o.AreParametersCustomized) {
		return nil, false
	}
	return o.AreParametersCustomized, true
}

// HasAreParametersCustomized returns a boolean if a field has been set.
func (o *IntegrationDto) HasAreParametersCustomized() bool {
	if o != nil && !IsNil(o.AreParametersCustomized) {
		return true
	}

	return false
}

// SetAreParametersCustomized gets a reference to the given bool and assigns it to the AreParametersCustomized field.
func (o *IntegrationDto) SetAreParametersCustomized(v bool) {
	o.AreParametersCustomized = &v
}

// GetBaseUri returns the BaseUri field value
func (o *IntegrationDto) GetBaseUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BaseUri
}

// GetBaseUriOk returns a tuple with the BaseUri field value
// and a boolean to check if the value has been set.
func (o *IntegrationDto) GetBaseUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseUri, true
}

// SetBaseUri sets field value
func (o *IntegrationDto) SetBaseUri(v string) {
	o.BaseUri = v
}

// GetConnectorVersion returns the ConnectorVersion field value if set, zero value otherwise.
func (o *IntegrationDto) GetConnectorVersion() string {
	if o == nil || IsNil(o.ConnectorVersion) {
		var ret string
		return ret
	}
	return *o.ConnectorVersion
}

// GetConnectorVersionOk returns a tuple with the ConnectorVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationDto) GetConnectorVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectorVersion) {
		return nil, false
	}
	return o.ConnectorVersion, true
}

// HasConnectorVersion returns a boolean if a field has been set.
func (o *IntegrationDto) HasConnectorVersion() bool {
	if o != nil && !IsNil(o.ConnectorVersion) {
		return true
	}

	return false
}

// SetConnectorVersion gets a reference to the given string and assigns it to the ConnectorVersion field.
func (o *IntegrationDto) SetConnectorVersion(v string) {
	o.ConnectorVersion = &v
}

// GetIsTest returns the IsTest field value if set, zero value otherwise.
func (o *IntegrationDto) GetIsTest() bool {
	if o == nil || IsNil(o.IsTest) {
		var ret bool
		return ret
	}
	return *o.IsTest
}

// GetIsTestOk returns a tuple with the IsTest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationDto) GetIsTestOk() (*bool, bool) {
	if o == nil || IsNil(o.IsTest) {
		return nil, false
	}
	return o.IsTest, true
}

// HasIsTest returns a boolean if a field has been set.
func (o *IntegrationDto) HasIsTest() bool {
	if o != nil && !IsNil(o.IsTest) {
		return true
	}

	return false
}

// SetIsTest gets a reference to the given bool and assigns it to the IsTest field.
func (o *IntegrationDto) SetIsTest(v bool) {
	o.IsTest = &v
}

// GetLogoUri returns the LogoUri field value
func (o *IntegrationDto) GetLogoUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogoUri
}

// GetLogoUriOk returns a tuple with the LogoUri field value
// and a boolean to check if the value has been set.
func (o *IntegrationDto) GetLogoUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogoUri, true
}

// SetLogoUri sets field value
func (o *IntegrationDto) SetLogoUri(v string) {
	o.LogoUri = v
}

// GetArchivedAt returns the ArchivedAt field value if set, zero value otherwise.
func (o *IntegrationDto) GetArchivedAt() time.Time {
	if o == nil || IsNil(o.ArchivedAt) {
		var ret time.Time
		return ret
	}
	return *o.ArchivedAt
}

// GetArchivedAtOk returns a tuple with the ArchivedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationDto) GetArchivedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ArchivedAt) {
		return nil, false
	}
	return o.ArchivedAt, true
}

// HasArchivedAt returns a boolean if a field has been set.
func (o *IntegrationDto) HasArchivedAt() bool {
	if o != nil && !IsNil(o.ArchivedAt) {
		return true
	}

	return false
}

// SetArchivedAt gets a reference to the given time.Time and assigns it to the ArchivedAt field.
func (o *IntegrationDto) SetArchivedAt(v time.Time) {
	o.ArchivedAt = &v
}

// GetConnectionMode returns the ConnectionMode field value if set, zero value otherwise.
// Deprecated
func (o *IntegrationDto) GetConnectionMode() string {
	if o == nil || IsNil(o.ConnectionMode) {
		var ret string
		return ret
	}
	return *o.ConnectionMode
}

// GetConnectionModeOk returns a tuple with the ConnectionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *IntegrationDto) GetConnectionModeOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectionMode) {
		return nil, false
	}
	return o.ConnectionMode, true
}

// HasConnectionMode returns a boolean if a field has been set.
func (o *IntegrationDto) HasConnectionMode() bool {
	if o != nil && !IsNil(o.ConnectionMode) {
		return true
	}

	return false
}

// SetConnectionMode gets a reference to the given string and assigns it to the ConnectionMode field.
// Deprecated
func (o *IntegrationDto) SetConnectionMode(v string) {
	o.ConnectionMode = &v
}

func (o IntegrationDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IntegrationDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["key"] = o.Key
	toSerialize["uuid"] = o.Uuid
	toSerialize["name"] = o.Name
	if !IsNil(o.ConnectorStoreKey) {
		toSerialize["connectorStoreKey"] = o.ConnectorStoreKey
	}
	if !IsNil(o.ConnectorId) {
		toSerialize["connectorId"] = o.ConnectorId
	}
	if !IsNil(o.AuthType) {
		toSerialize["authType"] = o.AuthType
	}
	if !IsNil(o.AuthOptions) {
		toSerialize["authOptions"] = o.AuthOptions
	}
	toSerialize["oAuthCallbackUri"] = o.OAuthCallbackUri
	if !IsNil(o.ParametersSchema) {
		toSerialize["parametersSchema"] = o.ParametersSchema
	}
	if !IsNil(o.HasDefaultParameters) {
		toSerialize["hasDefaultParameters"] = o.HasDefaultParameters
	}
	if !IsNil(o.HasMissingParameters) {
		toSerialize["hasMissingParameters"] = o.HasMissingParameters
	}
	if !IsNil(o.HasDocumentation) {
		toSerialize["hasDocumentation"] = o.HasDocumentation
	}
	if !IsNil(o.HasOperations) {
		toSerialize["hasOperations"] = o.HasOperations
	}
	if !IsNil(o.HasData) {
		toSerialize["hasData"] = o.HasData
	}
	if !IsNil(o.HasEvents) {
		toSerialize["hasEvents"] = o.HasEvents
	}
	if !IsNil(o.HasGlobalWebhooks) {
		toSerialize["hasGlobalWebhooks"] = o.HasGlobalWebhooks
	}
	if !IsNil(o.HasUdm) {
		toSerialize["hasUdm"] = o.HasUdm
	}
	if !IsNil(o.AreParametersCustomized) {
		toSerialize["areParametersCustomized"] = o.AreParametersCustomized
	}
	toSerialize["baseUri"] = o.BaseUri
	if !IsNil(o.ConnectorVersion) {
		toSerialize["connectorVersion"] = o.ConnectorVersion
	}
	if !IsNil(o.IsTest) {
		toSerialize["isTest"] = o.IsTest
	}
	toSerialize["logoUri"] = o.LogoUri
	if !IsNil(o.ArchivedAt) {
		toSerialize["archivedAt"] = o.ArchivedAt
	}
	if !IsNil(o.ConnectionMode) {
		toSerialize["connectionMode"] = o.ConnectionMode
	}
	return toSerialize, nil
}

func (o *IntegrationDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"key",
		"uuid",
		"name",
		"oAuthCallbackUri",
		"baseUri",
		"logoUri",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varIntegrationDto := _IntegrationDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varIntegrationDto)

	if err != nil {
		return err
	}

	*o = IntegrationDto(varIntegrationDto)

	return err
}

type NullableIntegrationDto struct {
	value *IntegrationDto
	isSet bool
}

func (v NullableIntegrationDto) Get() *IntegrationDto {
	return v.value
}

func (v *NullableIntegrationDto) Set(val *IntegrationDto) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationDto) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationDto(val *IntegrationDto) *NullableIntegrationDto {
	return &NullableIntegrationDto{value: val, isSet: true}
}

func (v NullableIntegrationDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


