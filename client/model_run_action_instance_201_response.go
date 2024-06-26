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

// checks if the RunActionInstance201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RunActionInstance201Response{}

// RunActionInstance201Response struct for RunActionInstance201Response
type RunActionInstance201Response struct {
	Output map[string]interface{} `json:"output,omitempty"`
	Logs []RunActionInstance201ResponseLogsInner `json:"logs,omitempty"`
}

// NewRunActionInstance201Response instantiates a new RunActionInstance201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunActionInstance201Response() *RunActionInstance201Response {
	this := RunActionInstance201Response{}
	return &this
}

// NewRunActionInstance201ResponseWithDefaults instantiates a new RunActionInstance201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunActionInstance201ResponseWithDefaults() *RunActionInstance201Response {
	this := RunActionInstance201Response{}
	return &this
}

// GetOutput returns the Output field value if set, zero value otherwise.
func (o *RunActionInstance201Response) GetOutput() map[string]interface{} {
	if o == nil || IsNil(o.Output) {
		var ret map[string]interface{}
		return ret
	}
	return o.Output
}

// GetOutputOk returns a tuple with the Output field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunActionInstance201Response) GetOutputOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Output) {
		return map[string]interface{}{}, false
	}
	return o.Output, true
}

// HasOutput returns a boolean if a field has been set.
func (o *RunActionInstance201Response) HasOutput() bool {
	if o != nil && !IsNil(o.Output) {
		return true
	}

	return false
}

// SetOutput gets a reference to the given map[string]interface{} and assigns it to the Output field.
func (o *RunActionInstance201Response) SetOutput(v map[string]interface{}) {
	o.Output = v
}

// GetLogs returns the Logs field value if set, zero value otherwise.
func (o *RunActionInstance201Response) GetLogs() []RunActionInstance201ResponseLogsInner {
	if o == nil || IsNil(o.Logs) {
		var ret []RunActionInstance201ResponseLogsInner
		return ret
	}
	return o.Logs
}

// GetLogsOk returns a tuple with the Logs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunActionInstance201Response) GetLogsOk() ([]RunActionInstance201ResponseLogsInner, bool) {
	if o == nil || IsNil(o.Logs) {
		return nil, false
	}
	return o.Logs, true
}

// HasLogs returns a boolean if a field has been set.
func (o *RunActionInstance201Response) HasLogs() bool {
	if o != nil && !IsNil(o.Logs) {
		return true
	}

	return false
}

// SetLogs gets a reference to the given []RunActionInstance201ResponseLogsInner and assigns it to the Logs field.
func (o *RunActionInstance201Response) SetLogs(v []RunActionInstance201ResponseLogsInner) {
	o.Logs = v
}

func (o RunActionInstance201Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RunActionInstance201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Output) {
		toSerialize["output"] = o.Output
	}
	if !IsNil(o.Logs) {
		toSerialize["logs"] = o.Logs
	}
	return toSerialize, nil
}

type NullableRunActionInstance201Response struct {
	value *RunActionInstance201Response
	isSet bool
}

func (v NullableRunActionInstance201Response) Get() *RunActionInstance201Response {
	return v.value
}

func (v *NullableRunActionInstance201Response) Set(val *RunActionInstance201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableRunActionInstance201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableRunActionInstance201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunActionInstance201Response(val *RunActionInstance201Response) *NullableRunActionInstance201Response {
	return &NullableRunActionInstance201Response{value: val, isSet: true}
}

func (v NullableRunActionInstance201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunActionInstance201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


