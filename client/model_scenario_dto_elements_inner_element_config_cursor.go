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

// checks if the ScenarioDtoElementsInnerElementConfigCursor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScenarioDtoElementsInnerElementConfigCursor{}

// ScenarioDtoElementsInnerElementConfigCursor struct for ScenarioDtoElementsInnerElementConfigCursor
type ScenarioDtoElementsInnerElementConfigCursor struct {
	Var *string `json:"$var,omitempty"`
}

// NewScenarioDtoElementsInnerElementConfigCursor instantiates a new ScenarioDtoElementsInnerElementConfigCursor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScenarioDtoElementsInnerElementConfigCursor() *ScenarioDtoElementsInnerElementConfigCursor {
	this := ScenarioDtoElementsInnerElementConfigCursor{}
	return &this
}

// NewScenarioDtoElementsInnerElementConfigCursorWithDefaults instantiates a new ScenarioDtoElementsInnerElementConfigCursor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScenarioDtoElementsInnerElementConfigCursorWithDefaults() *ScenarioDtoElementsInnerElementConfigCursor {
	this := ScenarioDtoElementsInnerElementConfigCursor{}
	return &this
}

// GetVar returns the Var field value if set, zero value otherwise.
func (o *ScenarioDtoElementsInnerElementConfigCursor) GetVar() string {
	if o == nil || IsNil(o.Var) {
		var ret string
		return ret
	}
	return *o.Var
}

// GetVarOk returns a tuple with the Var field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScenarioDtoElementsInnerElementConfigCursor) GetVarOk() (*string, bool) {
	if o == nil || IsNil(o.Var) {
		return nil, false
	}
	return o.Var, true
}

// HasVar returns a boolean if a field has been set.
func (o *ScenarioDtoElementsInnerElementConfigCursor) HasVar() bool {
	if o != nil && !IsNil(o.Var) {
		return true
	}

	return false
}

// SetVar gets a reference to the given string and assigns it to the Var field.
func (o *ScenarioDtoElementsInnerElementConfigCursor) SetVar(v string) {
	o.Var = &v
}

func (o ScenarioDtoElementsInnerElementConfigCursor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScenarioDtoElementsInnerElementConfigCursor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Var) {
		toSerialize["$var"] = o.Var
	}
	return toSerialize, nil
}

type NullableScenarioDtoElementsInnerElementConfigCursor struct {
	value *ScenarioDtoElementsInnerElementConfigCursor
	isSet bool
}

func (v NullableScenarioDtoElementsInnerElementConfigCursor) Get() *ScenarioDtoElementsInnerElementConfigCursor {
	return v.value
}

func (v *NullableScenarioDtoElementsInnerElementConfigCursor) Set(val *ScenarioDtoElementsInnerElementConfigCursor) {
	v.value = val
	v.isSet = true
}

func (v NullableScenarioDtoElementsInnerElementConfigCursor) IsSet() bool {
	return v.isSet
}

func (v *NullableScenarioDtoElementsInnerElementConfigCursor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScenarioDtoElementsInnerElementConfigCursor(val *ScenarioDtoElementsInnerElementConfigCursor) *NullableScenarioDtoElementsInnerElementConfigCursor {
	return &NullableScenarioDtoElementsInnerElementConfigCursor{value: val, isSet: true}
}

func (v NullableScenarioDtoElementsInnerElementConfigCursor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScenarioDtoElementsInnerElementConfigCursor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


