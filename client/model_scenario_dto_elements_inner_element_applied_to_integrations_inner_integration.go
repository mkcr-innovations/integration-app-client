/*
Integration.app API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration{}

// ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration struct for ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration
type ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration struct {
	Id *string `json:"id,omitempty"`
	Key *string `json:"key,omitempty"`
	Name *string `json:"name,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
	ConnectorId *string `json:"connectorId,omitempty"`
	AuthType *string `json:"authType,omitempty"`
	HasDefaultParameters *bool `json:"hasDefaultParameters,omitempty"`
	HasMissingParameters *bool `json:"hasMissingParameters,omitempty"`
	HasDocumentation *bool `json:"hasDocumentation,omitempty"`
	HasOperations *bool `json:"hasOperations,omitempty"`
	HasData *bool `json:"hasData,omitempty"`
	HasEvents *bool `json:"hasEvents,omitempty"`
	HasGlobalWebhooks *bool `json:"hasGlobalWebhooks,omitempty"`
	HasUdm *bool `json:"hasUdm,omitempty"`
	AreParametersCustomized *bool `json:"areParametersCustomized,omitempty"`
	BaseUri *string `json:"baseUri,omitempty"`
	LogoUri *string `json:"logoUri,omitempty"`
}

// NewScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration instantiates a new ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration() *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration {
	this := ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration{}
	return &this
}

// NewScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegrationWithDefaults instantiates a new ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegrationWithDefaults() *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration {
	this := ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) SetId(v string) {
	o.Id = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) SetKey(v string) {
	o.Key = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) SetName(v string) {
	o.Name = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) SetUuid(v string) {
	o.Uuid = &v
}

// GetConnectorId returns the ConnectorId field value if set, zero value otherwise.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) GetConnectorId() string {
	if o == nil || IsNil(o.ConnectorId) {
		var ret string
		return ret
	}
	return *o.ConnectorId
}

// GetConnectorIdOk returns a tuple with the ConnectorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) GetConnectorIdOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectorId) {
		return nil, false
	}
	return o.ConnectorId, true
}

// HasConnectorId returns a boolean if a field has been set.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) HasConnectorId() bool {
	if o != nil && !IsNil(o.ConnectorId) {
		return true
	}

	return false
}

// SetConnectorId gets a reference to the given string and assigns it to the ConnectorId field.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) SetConnectorId(v string) {
	o.ConnectorId = &v
}

// GetAuthType returns the AuthType field value if set, zero value otherwise.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) GetAuthType() string {
	if o == nil || IsNil(o.AuthType) {
		var ret string
		return ret
	}
	return *o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) GetAuthTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AuthType) {
		return nil, false
	}
	return o.AuthType, true
}

// HasAuthType returns a boolean if a field has been set.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) HasAuthType() bool {
	if o != nil && !IsNil(o.AuthType) {
		return true
	}

	return false
}

// SetAuthType gets a reference to the given string and assigns it to the AuthType field.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) SetAuthType(v string) {
	o.AuthType = &v
}

// GetHasDefaultParameters returns the HasDefaultParameters field value if set, zero value otherwise.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) GetHasDefaultParameters() bool {
	if o == nil || IsNil(o.HasDefaultParameters) {
		var ret bool
		return ret
	}
	return *o.HasDefaultParameters
}

// GetHasDefaultParametersOk returns a tuple with the HasDefaultParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) GetHasDefaultParametersOk() (*bool, bool) {
	if o == nil || IsNil(o.HasDefaultParameters) {
		return nil, false
	}
	return o.HasDefaultParameters, true
}

// HasHasDefaultParameters returns a boolean if a field has been set.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) HasHasDefaultParameters() bool {
	if o != nil && !IsNil(o.HasDefaultParameters) {
		return true
	}

	return false
}

// SetHasDefaultParameters gets a reference to the given bool and assigns it to the HasDefaultParameters field.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) SetHasDefaultParameters(v bool) {
	o.HasDefaultParameters = &v
}

// GetHasMissingParameters returns the HasMissingParameters field value if set, zero value otherwise.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) GetHasMissingParameters() bool {
	if o == nil || IsNil(o.HasMissingParameters) {
		var ret bool
		return ret
	}
	return *o.HasMissingParameters
}

// GetHasMissingParametersOk returns a tuple with the HasMissingParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) GetHasMissingParametersOk() (*bool, bool) {
	if o == nil || IsNil(o.HasMissingParameters) {
		return nil, false
	}
	return o.HasMissingParameters, true
}

// HasHasMissingParameters returns a boolean if a field has been set.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) HasHasMissingParameters() bool {
	if o != nil && !IsNil(o.HasMissingParameters) {
		return true
	}

	return false
}

// SetHasMissingParameters gets a reference to the given bool and assigns it to the HasMissingParameters field.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) SetHasMissingParameters(v bool) {
	o.HasMissingParameters = &v
}

// GetHasDocumentation returns the HasDocumentation field value if set, zero value otherwise.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) GetHasDocumentation() bool {
	if o == nil || IsNil(o.HasDocumentation) {
		var ret bool
		return ret
	}
	return *o.HasDocumentation
}

// GetHasDocumentationOk returns a tuple with the HasDocumentation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) GetHasDocumentationOk() (*bool, bool) {
	if o == nil || IsNil(o.HasDocumentation) {
		return nil, false
	}
	return o.HasDocumentation, true
}

// HasHasDocumentation returns a boolean if a field has been set.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) HasHasDocumentation() bool {
	if o != nil && !IsNil(o.HasDocumentation) {
		return true
	}

	return false
}

// SetHasDocumentation gets a reference to the given bool and assigns it to the HasDocumentation field.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) SetHasDocumentation(v bool) {
	o.HasDocumentation = &v
}

// GetHasOperations returns the HasOperations field value if set, zero value otherwise.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) GetHasOperations() bool {
	if o == nil || IsNil(o.HasOperations) {
		var ret bool
		return ret
	}
	return *o.HasOperations
}

// GetHasOperationsOk returns a tuple with the HasOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) GetHasOperationsOk() (*bool, bool) {
	if o == nil || IsNil(o.HasOperations) {
		return nil, false
	}
	return o.HasOperations, true
}

// HasHasOperations returns a boolean if a field has been set.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) HasHasOperations() bool {
	if o != nil && !IsNil(o.HasOperations) {
		return true
	}

	return false
}

// SetHasOperations gets a reference to the given bool and assigns it to the HasOperations field.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) SetHasOperations(v bool) {
	o.HasOperations = &v
}

// GetHasData returns the HasData field value if set, zero value otherwise.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) GetHasData() bool {
	if o == nil || IsNil(o.HasData) {
		var ret bool
		return ret
	}
	return *o.HasData
}

// GetHasDataOk returns a tuple with the HasData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) GetHasDataOk() (*bool, bool) {
	if o == nil || IsNil(o.HasData) {
		return nil, false
	}
	return o.HasData, true
}

// HasHasData returns a boolean if a field has been set.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) HasHasData() bool {
	if o != nil && !IsNil(o.HasData) {
		return true
	}

	return false
}

// SetHasData gets a reference to the given bool and assigns it to the HasData field.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) SetHasData(v bool) {
	o.HasData = &v
}

// GetHasEvents returns the HasEvents field value if set, zero value otherwise.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) GetHasEvents() bool {
	if o == nil || IsNil(o.HasEvents) {
		var ret bool
		return ret
	}
	return *o.HasEvents
}

// GetHasEventsOk returns a tuple with the HasEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) GetHasEventsOk() (*bool, bool) {
	if o == nil || IsNil(o.HasEvents) {
		return nil, false
	}
	return o.HasEvents, true
}

// HasHasEvents returns a boolean if a field has been set.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) HasHasEvents() bool {
	if o != nil && !IsNil(o.HasEvents) {
		return true
	}

	return false
}

// SetHasEvents gets a reference to the given bool and assigns it to the HasEvents field.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) SetHasEvents(v bool) {
	o.HasEvents = &v
}

// GetHasGlobalWebhooks returns the HasGlobalWebhooks field value if set, zero value otherwise.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) GetHasGlobalWebhooks() bool {
	if o == nil || IsNil(o.HasGlobalWebhooks) {
		var ret bool
		return ret
	}
	return *o.HasGlobalWebhooks
}

// GetHasGlobalWebhooksOk returns a tuple with the HasGlobalWebhooks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) GetHasGlobalWebhooksOk() (*bool, bool) {
	if o == nil || IsNil(o.HasGlobalWebhooks) {
		return nil, false
	}
	return o.HasGlobalWebhooks, true
}

// HasHasGlobalWebhooks returns a boolean if a field has been set.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) HasHasGlobalWebhooks() bool {
	if o != nil && !IsNil(o.HasGlobalWebhooks) {
		return true
	}

	return false
}

// SetHasGlobalWebhooks gets a reference to the given bool and assigns it to the HasGlobalWebhooks field.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) SetHasGlobalWebhooks(v bool) {
	o.HasGlobalWebhooks = &v
}

// GetHasUdm returns the HasUdm field value if set, zero value otherwise.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) GetHasUdm() bool {
	if o == nil || IsNil(o.HasUdm) {
		var ret bool
		return ret
	}
	return *o.HasUdm
}

// GetHasUdmOk returns a tuple with the HasUdm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) GetHasUdmOk() (*bool, bool) {
	if o == nil || IsNil(o.HasUdm) {
		return nil, false
	}
	return o.HasUdm, true
}

// HasHasUdm returns a boolean if a field has been set.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) HasHasUdm() bool {
	if o != nil && !IsNil(o.HasUdm) {
		return true
	}

	return false
}

// SetHasUdm gets a reference to the given bool and assigns it to the HasUdm field.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) SetHasUdm(v bool) {
	o.HasUdm = &v
}

// GetAreParametersCustomized returns the AreParametersCustomized field value if set, zero value otherwise.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) GetAreParametersCustomized() bool {
	if o == nil || IsNil(o.AreParametersCustomized) {
		var ret bool
		return ret
	}
	return *o.AreParametersCustomized
}

// GetAreParametersCustomizedOk returns a tuple with the AreParametersCustomized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) GetAreParametersCustomizedOk() (*bool, bool) {
	if o == nil || IsNil(o.AreParametersCustomized) {
		return nil, false
	}
	return o.AreParametersCustomized, true
}

// HasAreParametersCustomized returns a boolean if a field has been set.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) HasAreParametersCustomized() bool {
	if o != nil && !IsNil(o.AreParametersCustomized) {
		return true
	}

	return false
}

// SetAreParametersCustomized gets a reference to the given bool and assigns it to the AreParametersCustomized field.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) SetAreParametersCustomized(v bool) {
	o.AreParametersCustomized = &v
}

// GetBaseUri returns the BaseUri field value if set, zero value otherwise.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) GetBaseUri() string {
	if o == nil || IsNil(o.BaseUri) {
		var ret string
		return ret
	}
	return *o.BaseUri
}

// GetBaseUriOk returns a tuple with the BaseUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) GetBaseUriOk() (*string, bool) {
	if o == nil || IsNil(o.BaseUri) {
		return nil, false
	}
	return o.BaseUri, true
}

// HasBaseUri returns a boolean if a field has been set.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) HasBaseUri() bool {
	if o != nil && !IsNil(o.BaseUri) {
		return true
	}

	return false
}

// SetBaseUri gets a reference to the given string and assigns it to the BaseUri field.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) SetBaseUri(v string) {
	o.BaseUri = &v
}

// GetLogoUri returns the LogoUri field value if set, zero value otherwise.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) GetLogoUri() string {
	if o == nil || IsNil(o.LogoUri) {
		var ret string
		return ret
	}
	return *o.LogoUri
}

// GetLogoUriOk returns a tuple with the LogoUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) GetLogoUriOk() (*string, bool) {
	if o == nil || IsNil(o.LogoUri) {
		return nil, false
	}
	return o.LogoUri, true
}

// HasLogoUri returns a boolean if a field has been set.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) HasLogoUri() bool {
	if o != nil && !IsNil(o.LogoUri) {
		return true
	}

	return false
}

// SetLogoUri gets a reference to the given string and assigns it to the LogoUri field.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) SetLogoUri(v string) {
	o.LogoUri = &v
}

func (o ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !IsNil(o.ConnectorId) {
		toSerialize["connectorId"] = o.ConnectorId
	}
	if !IsNil(o.AuthType) {
		toSerialize["authType"] = o.AuthType
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
	if !IsNil(o.BaseUri) {
		toSerialize["baseUri"] = o.BaseUri
	}
	if !IsNil(o.LogoUri) {
		toSerialize["logoUri"] = o.LogoUri
	}
	return toSerialize, nil
}

type NullableScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration struct {
	value *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration
	isSet bool
}

func (v NullableScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) Get() *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration {
	return v.value
}

func (v *NullableScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) Set(val *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) {
	v.value = val
	v.isSet = true
}

func (v NullableScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) IsSet() bool {
	return v.isSet
}

func (v *NullableScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration(val *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) *NullableScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration {
	return &NullableScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration{value: val, isSet: true}
}

func (v NullableScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


