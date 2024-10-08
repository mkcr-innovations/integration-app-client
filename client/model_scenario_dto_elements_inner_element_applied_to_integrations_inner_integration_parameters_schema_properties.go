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

// checks if the ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegrationParametersSchemaProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegrationParametersSchemaProperties{}

// ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegrationParametersSchemaProperties struct for ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegrationParametersSchemaProperties
type ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegrationParametersSchemaProperties struct {
	ClientId *string `json:"clientId,omitempty"`
	ClientSecret *string `json:"clientSecret,omitempty"`
	Scopes []string `json:"scopes,omitempty"`
}

// NewScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegrationParametersSchemaProperties instantiates a new ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegrationParametersSchemaProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegrationParametersSchemaProperties() *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegrationParametersSchemaProperties {
	this := ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegrationParametersSchemaProperties{}
	return &this
}

// NewScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegrationParametersSchemaPropertiesWithDefaults instantiates a new ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegrationParametersSchemaProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegrationParametersSchemaPropertiesWithDefaults() *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegrationParametersSchemaProperties {
	this := ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegrationParametersSchemaProperties{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegrationParametersSchemaProperties) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegrationParametersSchemaProperties) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegrationParametersSchemaProperties) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegrationParametersSchemaProperties) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegrationParametersSchemaProperties) GetClientSecret() string {
	if o == nil || IsNil(o.ClientSecret) {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegrationParametersSchemaProperties) GetClientSecretOk() (*string, bool) {
	if o == nil || IsNil(o.ClientSecret) {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegrationParametersSchemaProperties) HasClientSecret() bool {
	if o != nil && !IsNil(o.ClientSecret) {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegrationParametersSchemaProperties) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegrationParametersSchemaProperties) GetScopes() []string {
	if o == nil || IsNil(o.Scopes) {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegrationParametersSchemaProperties) GetScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegrationParametersSchemaProperties) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegrationParametersSchemaProperties) SetScopes(v []string) {
	o.Scopes = v
}

func (o ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegrationParametersSchemaProperties) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegrationParametersSchemaProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClientId) {
		toSerialize["clientId"] = o.ClientId
	}
	if !IsNil(o.ClientSecret) {
		toSerialize["clientSecret"] = o.ClientSecret
	}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	return toSerialize, nil
}

type NullableScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegrationParametersSchemaProperties struct {
	value *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegrationParametersSchemaProperties
	isSet bool
}

func (v NullableScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegrationParametersSchemaProperties) Get() *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegrationParametersSchemaProperties {
	return v.value
}

func (v *NullableScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegrationParametersSchemaProperties) Set(val *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegrationParametersSchemaProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegrationParametersSchemaProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegrationParametersSchemaProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegrationParametersSchemaProperties(val *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegrationParametersSchemaProperties) *NullableScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegrationParametersSchemaProperties {
	return &NullableScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegrationParametersSchemaProperties{value: val, isSet: true}
}

func (v NullableScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegrationParametersSchemaProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegrationParametersSchemaProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


