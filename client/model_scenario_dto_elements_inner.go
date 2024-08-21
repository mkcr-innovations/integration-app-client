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

// checks if the ScenarioDtoElementsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScenarioDtoElementsInner{}

// ScenarioDtoElementsInner struct for ScenarioDtoElementsInner
type ScenarioDtoElementsInner struct {
	Id *string `json:"id,omitempty"`
	Type *string `json:"type,omitempty"`
	Element *ScenarioDtoElementsInnerElement `json:"element,omitempty"`
}

// NewScenarioDtoElementsInner instantiates a new ScenarioDtoElementsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScenarioDtoElementsInner() *ScenarioDtoElementsInner {
	this := ScenarioDtoElementsInner{}
	return &this
}

// NewScenarioDtoElementsInnerWithDefaults instantiates a new ScenarioDtoElementsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScenarioDtoElementsInnerWithDefaults() *ScenarioDtoElementsInner {
	this := ScenarioDtoElementsInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ScenarioDtoElementsInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScenarioDtoElementsInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ScenarioDtoElementsInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ScenarioDtoElementsInner) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ScenarioDtoElementsInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScenarioDtoElementsInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ScenarioDtoElementsInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ScenarioDtoElementsInner) SetType(v string) {
	o.Type = &v
}

// GetElement returns the Element field value if set, zero value otherwise.
func (o *ScenarioDtoElementsInner) GetElement() ScenarioDtoElementsInnerElement {
	if o == nil || IsNil(o.Element) {
		var ret ScenarioDtoElementsInnerElement
		return ret
	}
	return *o.Element
}

// GetElementOk returns a tuple with the Element field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScenarioDtoElementsInner) GetElementOk() (*ScenarioDtoElementsInnerElement, bool) {
	if o == nil || IsNil(o.Element) {
		return nil, false
	}
	return o.Element, true
}

// HasElement returns a boolean if a field has been set.
func (o *ScenarioDtoElementsInner) HasElement() bool {
	if o != nil && !IsNil(o.Element) {
		return true
	}

	return false
}

// SetElement gets a reference to the given ScenarioDtoElementsInnerElement and assigns it to the Element field.
func (o *ScenarioDtoElementsInner) SetElement(v ScenarioDtoElementsInnerElement) {
	o.Element = &v
}

func (o ScenarioDtoElementsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScenarioDtoElementsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Element) {
		toSerialize["element"] = o.Element
	}
	return toSerialize, nil
}

type NullableScenarioDtoElementsInner struct {
	value *ScenarioDtoElementsInner
	isSet bool
}

func (v NullableScenarioDtoElementsInner) Get() *ScenarioDtoElementsInner {
	return v.value
}

func (v *NullableScenarioDtoElementsInner) Set(val *ScenarioDtoElementsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableScenarioDtoElementsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableScenarioDtoElementsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScenarioDtoElementsInner(val *ScenarioDtoElementsInner) *NullableScenarioDtoElementsInner {
	return &NullableScenarioDtoElementsInner{value: val, isSet: true}
}

func (v NullableScenarioDtoElementsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScenarioDtoElementsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


