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

// checks if the ScenarioDtoElementsInnerElementConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScenarioDtoElementsInnerElementConfig{}

// ScenarioDtoElementsInnerElementConfig struct for ScenarioDtoElementsInnerElementConfig
type ScenarioDtoElementsInnerElementConfig struct {
	DataSource *ScenarioDtoElementsInnerElementConfigDataSource `json:"dataSource,omitempty"`
	Cursor *ScenarioDtoElementsInnerElementConfigCursor `json:"cursor,omitempty"`
}

// NewScenarioDtoElementsInnerElementConfig instantiates a new ScenarioDtoElementsInnerElementConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScenarioDtoElementsInnerElementConfig() *ScenarioDtoElementsInnerElementConfig {
	this := ScenarioDtoElementsInnerElementConfig{}
	return &this
}

// NewScenarioDtoElementsInnerElementConfigWithDefaults instantiates a new ScenarioDtoElementsInnerElementConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScenarioDtoElementsInnerElementConfigWithDefaults() *ScenarioDtoElementsInnerElementConfig {
	this := ScenarioDtoElementsInnerElementConfig{}
	return &this
}

// GetDataSource returns the DataSource field value if set, zero value otherwise.
func (o *ScenarioDtoElementsInnerElementConfig) GetDataSource() ScenarioDtoElementsInnerElementConfigDataSource {
	if o == nil || IsNil(o.DataSource) {
		var ret ScenarioDtoElementsInnerElementConfigDataSource
		return ret
	}
	return *o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScenarioDtoElementsInnerElementConfig) GetDataSourceOk() (*ScenarioDtoElementsInnerElementConfigDataSource, bool) {
	if o == nil || IsNil(o.DataSource) {
		return nil, false
	}
	return o.DataSource, true
}

// HasDataSource returns a boolean if a field has been set.
func (o *ScenarioDtoElementsInnerElementConfig) HasDataSource() bool {
	if o != nil && !IsNil(o.DataSource) {
		return true
	}

	return false
}

// SetDataSource gets a reference to the given ScenarioDtoElementsInnerElementConfigDataSource and assigns it to the DataSource field.
func (o *ScenarioDtoElementsInnerElementConfig) SetDataSource(v ScenarioDtoElementsInnerElementConfigDataSource) {
	o.DataSource = &v
}

// GetCursor returns the Cursor field value if set, zero value otherwise.
func (o *ScenarioDtoElementsInnerElementConfig) GetCursor() ScenarioDtoElementsInnerElementConfigCursor {
	if o == nil || IsNil(o.Cursor) {
		var ret ScenarioDtoElementsInnerElementConfigCursor
		return ret
	}
	return *o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScenarioDtoElementsInnerElementConfig) GetCursorOk() (*ScenarioDtoElementsInnerElementConfigCursor, bool) {
	if o == nil || IsNil(o.Cursor) {
		return nil, false
	}
	return o.Cursor, true
}

// HasCursor returns a boolean if a field has been set.
func (o *ScenarioDtoElementsInnerElementConfig) HasCursor() bool {
	if o != nil && !IsNil(o.Cursor) {
		return true
	}

	return false
}

// SetCursor gets a reference to the given ScenarioDtoElementsInnerElementConfigCursor and assigns it to the Cursor field.
func (o *ScenarioDtoElementsInnerElementConfig) SetCursor(v ScenarioDtoElementsInnerElementConfigCursor) {
	o.Cursor = &v
}

func (o ScenarioDtoElementsInnerElementConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScenarioDtoElementsInnerElementConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DataSource) {
		toSerialize["dataSource"] = o.DataSource
	}
	if !IsNil(o.Cursor) {
		toSerialize["cursor"] = o.Cursor
	}
	return toSerialize, nil
}

type NullableScenarioDtoElementsInnerElementConfig struct {
	value *ScenarioDtoElementsInnerElementConfig
	isSet bool
}

func (v NullableScenarioDtoElementsInnerElementConfig) Get() *ScenarioDtoElementsInnerElementConfig {
	return v.value
}

func (v *NullableScenarioDtoElementsInnerElementConfig) Set(val *ScenarioDtoElementsInnerElementConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableScenarioDtoElementsInnerElementConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableScenarioDtoElementsInnerElementConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScenarioDtoElementsInnerElementConfig(val *ScenarioDtoElementsInnerElementConfig) *NullableScenarioDtoElementsInnerElementConfig {
	return &NullableScenarioDtoElementsInnerElementConfig{value: val, isSet: true}
}

func (v NullableScenarioDtoElementsInnerElementConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScenarioDtoElementsInnerElementConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

