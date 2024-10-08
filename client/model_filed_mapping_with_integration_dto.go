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

// checks if the FiledMappingWithIntegrationDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FiledMappingWithIntegrationDto{}

// FiledMappingWithIntegrationDto struct for FiledMappingWithIntegrationDto
type FiledMappingWithIntegrationDto struct {
	Element *FieldMappingDto `json:"element,omitempty"`
	Integration *IntegrationDto `json:"integration,omitempty"`
}

// NewFiledMappingWithIntegrationDto instantiates a new FiledMappingWithIntegrationDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFiledMappingWithIntegrationDto() *FiledMappingWithIntegrationDto {
	this := FiledMappingWithIntegrationDto{}
	return &this
}

// NewFiledMappingWithIntegrationDtoWithDefaults instantiates a new FiledMappingWithIntegrationDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFiledMappingWithIntegrationDtoWithDefaults() *FiledMappingWithIntegrationDto {
	this := FiledMappingWithIntegrationDto{}
	return &this
}

// GetElement returns the Element field value if set, zero value otherwise.
func (o *FiledMappingWithIntegrationDto) GetElement() FieldMappingDto {
	if o == nil || IsNil(o.Element) {
		var ret FieldMappingDto
		return ret
	}
	return *o.Element
}

// GetElementOk returns a tuple with the Element field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiledMappingWithIntegrationDto) GetElementOk() (*FieldMappingDto, bool) {
	if o == nil || IsNil(o.Element) {
		return nil, false
	}
	return o.Element, true
}

// HasElement returns a boolean if a field has been set.
func (o *FiledMappingWithIntegrationDto) HasElement() bool {
	if o != nil && !IsNil(o.Element) {
		return true
	}

	return false
}

// SetElement gets a reference to the given FieldMappingDto and assigns it to the Element field.
func (o *FiledMappingWithIntegrationDto) SetElement(v FieldMappingDto) {
	o.Element = &v
}

// GetIntegration returns the Integration field value if set, zero value otherwise.
func (o *FiledMappingWithIntegrationDto) GetIntegration() IntegrationDto {
	if o == nil || IsNil(o.Integration) {
		var ret IntegrationDto
		return ret
	}
	return *o.Integration
}

// GetIntegrationOk returns a tuple with the Integration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiledMappingWithIntegrationDto) GetIntegrationOk() (*IntegrationDto, bool) {
	if o == nil || IsNil(o.Integration) {
		return nil, false
	}
	return o.Integration, true
}

// HasIntegration returns a boolean if a field has been set.
func (o *FiledMappingWithIntegrationDto) HasIntegration() bool {
	if o != nil && !IsNil(o.Integration) {
		return true
	}

	return false
}

// SetIntegration gets a reference to the given IntegrationDto and assigns it to the Integration field.
func (o *FiledMappingWithIntegrationDto) SetIntegration(v IntegrationDto) {
	o.Integration = &v
}

func (o FiledMappingWithIntegrationDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FiledMappingWithIntegrationDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Element) {
		toSerialize["element"] = o.Element
	}
	if !IsNil(o.Integration) {
		toSerialize["integration"] = o.Integration
	}
	return toSerialize, nil
}

type NullableFiledMappingWithIntegrationDto struct {
	value *FiledMappingWithIntegrationDto
	isSet bool
}

func (v NullableFiledMappingWithIntegrationDto) Get() *FiledMappingWithIntegrationDto {
	return v.value
}

func (v *NullableFiledMappingWithIntegrationDto) Set(val *FiledMappingWithIntegrationDto) {
	v.value = val
	v.isSet = true
}

func (v NullableFiledMappingWithIntegrationDto) IsSet() bool {
	return v.isSet
}

func (v *NullableFiledMappingWithIntegrationDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFiledMappingWithIntegrationDto(val *FiledMappingWithIntegrationDto) *NullableFiledMappingWithIntegrationDto {
	return &NullableFiledMappingWithIntegrationDto{value: val, isSet: true}
}

func (v NullableFiledMappingWithIntegrationDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFiledMappingWithIntegrationDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


