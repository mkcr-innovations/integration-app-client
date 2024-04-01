/*
Integration Engine API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerElement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerElement{}

// ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerElement struct for ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerElement
type ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerElement struct {
	Id *string `json:"id,omitempty"`
	IntegrationId *string `json:"integrationId,omitempty"`
	Customized *bool `json:"customized,omitempty"`
}

// NewScenarioDtoElementsInnerElementAppliedToIntegrationsInnerElement instantiates a new ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerElement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScenarioDtoElementsInnerElementAppliedToIntegrationsInnerElement() *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerElement {
	this := ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerElement{}
	return &this
}

// NewScenarioDtoElementsInnerElementAppliedToIntegrationsInnerElementWithDefaults instantiates a new ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerElement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScenarioDtoElementsInnerElementAppliedToIntegrationsInnerElementWithDefaults() *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerElement {
	this := ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerElement{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerElement) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerElement) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerElement) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerElement) SetId(v string) {
	o.Id = &v
}

// GetIntegrationId returns the IntegrationId field value if set, zero value otherwise.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerElement) GetIntegrationId() string {
	if o == nil || IsNil(o.IntegrationId) {
		var ret string
		return ret
	}
	return *o.IntegrationId
}

// GetIntegrationIdOk returns a tuple with the IntegrationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerElement) GetIntegrationIdOk() (*string, bool) {
	if o == nil || IsNil(o.IntegrationId) {
		return nil, false
	}
	return o.IntegrationId, true
}

// HasIntegrationId returns a boolean if a field has been set.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerElement) HasIntegrationId() bool {
	if o != nil && !IsNil(o.IntegrationId) {
		return true
	}

	return false
}

// SetIntegrationId gets a reference to the given string and assigns it to the IntegrationId field.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerElement) SetIntegrationId(v string) {
	o.IntegrationId = &v
}

// GetCustomized returns the Customized field value if set, zero value otherwise.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerElement) GetCustomized() bool {
	if o == nil || IsNil(o.Customized) {
		var ret bool
		return ret
	}
	return *o.Customized
}

// GetCustomizedOk returns a tuple with the Customized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerElement) GetCustomizedOk() (*bool, bool) {
	if o == nil || IsNil(o.Customized) {
		return nil, false
	}
	return o.Customized, true
}

// HasCustomized returns a boolean if a field has been set.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerElement) HasCustomized() bool {
	if o != nil && !IsNil(o.Customized) {
		return true
	}

	return false
}

// SetCustomized gets a reference to the given bool and assigns it to the Customized field.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerElement) SetCustomized(v bool) {
	o.Customized = &v
}

func (o ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerElement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerElement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IntegrationId) {
		toSerialize["integrationId"] = o.IntegrationId
	}
	if !IsNil(o.Customized) {
		toSerialize["customized"] = o.Customized
	}
	return toSerialize, nil
}

type NullableScenarioDtoElementsInnerElementAppliedToIntegrationsInnerElement struct {
	value *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerElement
	isSet bool
}

func (v NullableScenarioDtoElementsInnerElementAppliedToIntegrationsInnerElement) Get() *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerElement {
	return v.value
}

func (v *NullableScenarioDtoElementsInnerElementAppliedToIntegrationsInnerElement) Set(val *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerElement) {
	v.value = val
	v.isSet = true
}

func (v NullableScenarioDtoElementsInnerElementAppliedToIntegrationsInnerElement) IsSet() bool {
	return v.isSet
}

func (v *NullableScenarioDtoElementsInnerElementAppliedToIntegrationsInnerElement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScenarioDtoElementsInnerElementAppliedToIntegrationsInnerElement(val *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerElement) *NullableScenarioDtoElementsInnerElementAppliedToIntegrationsInnerElement {
	return &NullableScenarioDtoElementsInnerElementAppliedToIntegrationsInnerElement{value: val, isSet: true}
}

func (v NullableScenarioDtoElementsInnerElementAppliedToIntegrationsInnerElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScenarioDtoElementsInnerElementAppliedToIntegrationsInnerElement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


