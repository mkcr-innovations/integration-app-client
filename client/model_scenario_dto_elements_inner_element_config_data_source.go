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

// checks if the ScenarioDtoElementsInnerElementConfigDataSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScenarioDtoElementsInnerElementConfigDataSource{}

// ScenarioDtoElementsInnerElementConfigDataSource struct for ScenarioDtoElementsInnerElementConfigDataSource
type ScenarioDtoElementsInnerElementConfigDataSource struct {
	Key *string `json:"key,omitempty"`
}

// NewScenarioDtoElementsInnerElementConfigDataSource instantiates a new ScenarioDtoElementsInnerElementConfigDataSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScenarioDtoElementsInnerElementConfigDataSource() *ScenarioDtoElementsInnerElementConfigDataSource {
	this := ScenarioDtoElementsInnerElementConfigDataSource{}
	return &this
}

// NewScenarioDtoElementsInnerElementConfigDataSourceWithDefaults instantiates a new ScenarioDtoElementsInnerElementConfigDataSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScenarioDtoElementsInnerElementConfigDataSourceWithDefaults() *ScenarioDtoElementsInnerElementConfigDataSource {
	this := ScenarioDtoElementsInnerElementConfigDataSource{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *ScenarioDtoElementsInnerElementConfigDataSource) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScenarioDtoElementsInnerElementConfigDataSource) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ScenarioDtoElementsInnerElementConfigDataSource) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *ScenarioDtoElementsInnerElementConfigDataSource) SetKey(v string) {
	o.Key = &v
}

func (o ScenarioDtoElementsInnerElementConfigDataSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScenarioDtoElementsInnerElementConfigDataSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	return toSerialize, nil
}

type NullableScenarioDtoElementsInnerElementConfigDataSource struct {
	value *ScenarioDtoElementsInnerElementConfigDataSource
	isSet bool
}

func (v NullableScenarioDtoElementsInnerElementConfigDataSource) Get() *ScenarioDtoElementsInnerElementConfigDataSource {
	return v.value
}

func (v *NullableScenarioDtoElementsInnerElementConfigDataSource) Set(val *ScenarioDtoElementsInnerElementConfigDataSource) {
	v.value = val
	v.isSet = true
}

func (v NullableScenarioDtoElementsInnerElementConfigDataSource) IsSet() bool {
	return v.isSet
}

func (v *NullableScenarioDtoElementsInnerElementConfigDataSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScenarioDtoElementsInnerElementConfigDataSource(val *ScenarioDtoElementsInnerElementConfigDataSource) *NullableScenarioDtoElementsInnerElementConfigDataSource {
	return &NullableScenarioDtoElementsInnerElementConfigDataSource{value: val, isSet: true}
}

func (v NullableScenarioDtoElementsInnerElementConfigDataSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScenarioDtoElementsInnerElementConfigDataSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


