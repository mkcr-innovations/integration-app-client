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

// checks if the ScenarioDtoElementsInnerElementAppliedToIntegrationsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScenarioDtoElementsInnerElementAppliedToIntegrationsInner{}

// ScenarioDtoElementsInnerElementAppliedToIntegrationsInner struct for ScenarioDtoElementsInnerElementAppliedToIntegrationsInner
type ScenarioDtoElementsInnerElementAppliedToIntegrationsInner struct {
	Element *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerElement `json:"element,omitempty"`
	Integration *ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration `json:"integration,omitempty"`
}

// NewScenarioDtoElementsInnerElementAppliedToIntegrationsInner instantiates a new ScenarioDtoElementsInnerElementAppliedToIntegrationsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScenarioDtoElementsInnerElementAppliedToIntegrationsInner() *ScenarioDtoElementsInnerElementAppliedToIntegrationsInner {
	this := ScenarioDtoElementsInnerElementAppliedToIntegrationsInner{}
	return &this
}

// NewScenarioDtoElementsInnerElementAppliedToIntegrationsInnerWithDefaults instantiates a new ScenarioDtoElementsInnerElementAppliedToIntegrationsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScenarioDtoElementsInnerElementAppliedToIntegrationsInnerWithDefaults() *ScenarioDtoElementsInnerElementAppliedToIntegrationsInner {
	this := ScenarioDtoElementsInnerElementAppliedToIntegrationsInner{}
	return &this
}

// GetElement returns the Element field value if set, zero value otherwise.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInner) GetElement() ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerElement {
	if o == nil || IsNil(o.Element) {
		var ret ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerElement
		return ret
	}
	return *o.Element
}

// GetElementOk returns a tuple with the Element field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInner) GetElementOk() (*ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerElement, bool) {
	if o == nil || IsNil(o.Element) {
		return nil, false
	}
	return o.Element, true
}

// HasElement returns a boolean if a field has been set.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInner) HasElement() bool {
	if o != nil && !IsNil(o.Element) {
		return true
	}

	return false
}

// SetElement gets a reference to the given ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerElement and assigns it to the Element field.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInner) SetElement(v ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerElement) {
	o.Element = &v
}

// GetIntegration returns the Integration field value if set, zero value otherwise.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInner) GetIntegration() ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration {
	if o == nil || IsNil(o.Integration) {
		var ret ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration
		return ret
	}
	return *o.Integration
}

// GetIntegrationOk returns a tuple with the Integration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInner) GetIntegrationOk() (*ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration, bool) {
	if o == nil || IsNil(o.Integration) {
		return nil, false
	}
	return o.Integration, true
}

// HasIntegration returns a boolean if a field has been set.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInner) HasIntegration() bool {
	if o != nil && !IsNil(o.Integration) {
		return true
	}

	return false
}

// SetIntegration gets a reference to the given ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration and assigns it to the Integration field.
func (o *ScenarioDtoElementsInnerElementAppliedToIntegrationsInner) SetIntegration(v ScenarioDtoElementsInnerElementAppliedToIntegrationsInnerIntegration) {
	o.Integration = &v
}

func (o ScenarioDtoElementsInnerElementAppliedToIntegrationsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScenarioDtoElementsInnerElementAppliedToIntegrationsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Element) {
		toSerialize["element"] = o.Element
	}
	if !IsNil(o.Integration) {
		toSerialize["integration"] = o.Integration
	}
	return toSerialize, nil
}

type NullableScenarioDtoElementsInnerElementAppliedToIntegrationsInner struct {
	value *ScenarioDtoElementsInnerElementAppliedToIntegrationsInner
	isSet bool
}

func (v NullableScenarioDtoElementsInnerElementAppliedToIntegrationsInner) Get() *ScenarioDtoElementsInnerElementAppliedToIntegrationsInner {
	return v.value
}

func (v *NullableScenarioDtoElementsInnerElementAppliedToIntegrationsInner) Set(val *ScenarioDtoElementsInnerElementAppliedToIntegrationsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableScenarioDtoElementsInnerElementAppliedToIntegrationsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableScenarioDtoElementsInnerElementAppliedToIntegrationsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScenarioDtoElementsInnerElementAppliedToIntegrationsInner(val *ScenarioDtoElementsInnerElementAppliedToIntegrationsInner) *NullableScenarioDtoElementsInnerElementAppliedToIntegrationsInner {
	return &NullableScenarioDtoElementsInnerElementAppliedToIntegrationsInner{value: val, isSet: true}
}

func (v NullableScenarioDtoElementsInnerElementAppliedToIntegrationsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScenarioDtoElementsInnerElementAppliedToIntegrationsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


