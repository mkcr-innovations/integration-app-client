/*
Integration.app API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the FlorRunLaunchedByDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FlorRunLaunchedByDto{}

// FlorRunLaunchedByDto struct for FlorRunLaunchedByDto
type FlorRunLaunchedByDto struct {
	Type string `json:"type"`
	EventId *string `json:"eventId,omitempty"`
}

type _FlorRunLaunchedByDto FlorRunLaunchedByDto

// NewFlorRunLaunchedByDto instantiates a new FlorRunLaunchedByDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlorRunLaunchedByDto(type_ string) *FlorRunLaunchedByDto {
	this := FlorRunLaunchedByDto{}
	this.Type = type_
	return &this
}

// NewFlorRunLaunchedByDtoWithDefaults instantiates a new FlorRunLaunchedByDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlorRunLaunchedByDtoWithDefaults() *FlorRunLaunchedByDto {
	this := FlorRunLaunchedByDto{}
	return &this
}

// GetType returns the Type field value
func (o *FlorRunLaunchedByDto) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FlorRunLaunchedByDto) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FlorRunLaunchedByDto) SetType(v string) {
	o.Type = v
}

// GetEventId returns the EventId field value if set, zero value otherwise.
func (o *FlorRunLaunchedByDto) GetEventId() string {
	if o == nil || IsNil(o.EventId) {
		var ret string
		return ret
	}
	return *o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlorRunLaunchedByDto) GetEventIdOk() (*string, bool) {
	if o == nil || IsNil(o.EventId) {
		return nil, false
	}
	return o.EventId, true
}

// HasEventId returns a boolean if a field has been set.
func (o *FlorRunLaunchedByDto) HasEventId() bool {
	if o != nil && !IsNil(o.EventId) {
		return true
	}

	return false
}

// SetEventId gets a reference to the given string and assigns it to the EventId field.
func (o *FlorRunLaunchedByDto) SetEventId(v string) {
	o.EventId = &v
}

func (o FlorRunLaunchedByDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlorRunLaunchedByDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.EventId) {
		toSerialize["eventId"] = o.EventId
	}
	return toSerialize, nil
}

func (o *FlorRunLaunchedByDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varFlorRunLaunchedByDto := _FlorRunLaunchedByDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFlorRunLaunchedByDto)

	if err != nil {
		return err
	}

	*o = FlorRunLaunchedByDto(varFlorRunLaunchedByDto)

	return err
}

type NullableFlorRunLaunchedByDto struct {
	value *FlorRunLaunchedByDto
	isSet bool
}

func (v NullableFlorRunLaunchedByDto) Get() *FlorRunLaunchedByDto {
	return v.value
}

func (v *NullableFlorRunLaunchedByDto) Set(val *FlorRunLaunchedByDto) {
	v.value = val
	v.isSet = true
}

func (v NullableFlorRunLaunchedByDto) IsSet() bool {
	return v.isSet
}

func (v *NullableFlorRunLaunchedByDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlorRunLaunchedByDto(val *FlorRunLaunchedByDto) *NullableFlorRunLaunchedByDto {
	return &NullableFlorRunLaunchedByDto{value: val, isSet: true}
}

func (v NullableFlorRunLaunchedByDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlorRunLaunchedByDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


