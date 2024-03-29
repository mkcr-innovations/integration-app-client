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

// checks if the RunFlowPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RunFlowPayload{}

// RunFlowPayload struct for RunFlowPayload
type RunFlowPayload struct {
	NodeKey *string `json:"nodeKey,omitempty"`
	Input map[string]interface{} `json:"input,omitempty"`
}

// NewRunFlowPayload instantiates a new RunFlowPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunFlowPayload() *RunFlowPayload {
	this := RunFlowPayload{}
	return &this
}

// NewRunFlowPayloadWithDefaults instantiates a new RunFlowPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunFlowPayloadWithDefaults() *RunFlowPayload {
	this := RunFlowPayload{}
	return &this
}

// GetNodeKey returns the NodeKey field value if set, zero value otherwise.
func (o *RunFlowPayload) GetNodeKey() string {
	if o == nil || IsNil(o.NodeKey) {
		var ret string
		return ret
	}
	return *o.NodeKey
}

// GetNodeKeyOk returns a tuple with the NodeKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunFlowPayload) GetNodeKeyOk() (*string, bool) {
	if o == nil || IsNil(o.NodeKey) {
		return nil, false
	}
	return o.NodeKey, true
}

// HasNodeKey returns a boolean if a field has been set.
func (o *RunFlowPayload) HasNodeKey() bool {
	if o != nil && !IsNil(o.NodeKey) {
		return true
	}

	return false
}

// SetNodeKey gets a reference to the given string and assigns it to the NodeKey field.
func (o *RunFlowPayload) SetNodeKey(v string) {
	o.NodeKey = &v
}

// GetInput returns the Input field value if set, zero value otherwise.
func (o *RunFlowPayload) GetInput() map[string]interface{} {
	if o == nil || IsNil(o.Input) {
		var ret map[string]interface{}
		return ret
	}
	return o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunFlowPayload) GetInputOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Input) {
		return map[string]interface{}{}, false
	}
	return o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *RunFlowPayload) HasInput() bool {
	if o != nil && !IsNil(o.Input) {
		return true
	}

	return false
}

// SetInput gets a reference to the given map[string]interface{} and assigns it to the Input field.
func (o *RunFlowPayload) SetInput(v map[string]interface{}) {
	o.Input = v
}

func (o RunFlowPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RunFlowPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NodeKey) {
		toSerialize["nodeKey"] = o.NodeKey
	}
	if !IsNil(o.Input) {
		toSerialize["input"] = o.Input
	}
	return toSerialize, nil
}

type NullableRunFlowPayload struct {
	value *RunFlowPayload
	isSet bool
}

func (v NullableRunFlowPayload) Get() *RunFlowPayload {
	return v.value
}

func (v *NullableRunFlowPayload) Set(val *RunFlowPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableRunFlowPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableRunFlowPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunFlowPayload(val *RunFlowPayload) *NullableRunFlowPayload {
	return &NullableRunFlowPayload{value: val, isSet: true}
}

func (v NullableRunFlowPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunFlowPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


