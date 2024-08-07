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

// checks if the OperationRunResponseDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OperationRunResponseDto{}

// OperationRunResponseDto struct for OperationRunResponseDto
type OperationRunResponseDto struct {
	Output map[string]interface{} `json:"output,omitempty"`
	Logs []map[string]interface{} `json:"logs,omitempty"`
}

// NewOperationRunResponseDto instantiates a new OperationRunResponseDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperationRunResponseDto() *OperationRunResponseDto {
	this := OperationRunResponseDto{}
	return &this
}

// NewOperationRunResponseDtoWithDefaults instantiates a new OperationRunResponseDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperationRunResponseDtoWithDefaults() *OperationRunResponseDto {
	this := OperationRunResponseDto{}
	return &this
}

// GetOutput returns the Output field value if set, zero value otherwise.
func (o *OperationRunResponseDto) GetOutput() map[string]interface{} {
	if o == nil || IsNil(o.Output) {
		var ret map[string]interface{}
		return ret
	}
	return o.Output
}

// GetOutputOk returns a tuple with the Output field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationRunResponseDto) GetOutputOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Output) {
		return map[string]interface{}{}, false
	}
	return o.Output, true
}

// HasOutput returns a boolean if a field has been set.
func (o *OperationRunResponseDto) HasOutput() bool {
	if o != nil && !IsNil(o.Output) {
		return true
	}

	return false
}

// SetOutput gets a reference to the given map[string]interface{} and assigns it to the Output field.
func (o *OperationRunResponseDto) SetOutput(v map[string]interface{}) {
	o.Output = v
}

// GetLogs returns the Logs field value if set, zero value otherwise.
func (o *OperationRunResponseDto) GetLogs() []map[string]interface{} {
	if o == nil || IsNil(o.Logs) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Logs
}

// GetLogsOk returns a tuple with the Logs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationRunResponseDto) GetLogsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Logs) {
		return nil, false
	}
	return o.Logs, true
}

// HasLogs returns a boolean if a field has been set.
func (o *OperationRunResponseDto) HasLogs() bool {
	if o != nil && !IsNil(o.Logs) {
		return true
	}

	return false
}

// SetLogs gets a reference to the given []map[string]interface{} and assigns it to the Logs field.
func (o *OperationRunResponseDto) SetLogs(v []map[string]interface{}) {
	o.Logs = v
}

func (o OperationRunResponseDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OperationRunResponseDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Output) {
		toSerialize["output"] = o.Output
	}
	if !IsNil(o.Logs) {
		toSerialize["logs"] = o.Logs
	}
	return toSerialize, nil
}

type NullableOperationRunResponseDto struct {
	value *OperationRunResponseDto
	isSet bool
}

func (v NullableOperationRunResponseDto) Get() *OperationRunResponseDto {
	return v.value
}

func (v *NullableOperationRunResponseDto) Set(val *OperationRunResponseDto) {
	v.value = val
	v.isSet = true
}

func (v NullableOperationRunResponseDto) IsSet() bool {
	return v.isSet
}

func (v *NullableOperationRunResponseDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperationRunResponseDto(val *OperationRunResponseDto) *NullableOperationRunResponseDto {
	return &NullableOperationRunResponseDto{value: val, isSet: true}
}

func (v NullableOperationRunResponseDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperationRunResponseDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


