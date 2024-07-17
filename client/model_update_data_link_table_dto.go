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

// checks if the UpdateDataLinkTableDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDataLinkTableDto{}

// UpdateDataLinkTableDto struct for UpdateDataLinkTableDto
type UpdateDataLinkTableDto struct {
	Key *string `json:"key,omitempty"`
	Name *string `json:"name,omitempty"`
	ArchivedAt *string `json:"archivedAt,omitempty"`
}

// NewUpdateDataLinkTableDto instantiates a new UpdateDataLinkTableDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDataLinkTableDto() *UpdateDataLinkTableDto {
	this := UpdateDataLinkTableDto{}
	return &this
}

// NewUpdateDataLinkTableDtoWithDefaults instantiates a new UpdateDataLinkTableDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDataLinkTableDtoWithDefaults() *UpdateDataLinkTableDto {
	this := UpdateDataLinkTableDto{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *UpdateDataLinkTableDto) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDataLinkTableDto) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *UpdateDataLinkTableDto) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *UpdateDataLinkTableDto) SetKey(v string) {
	o.Key = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateDataLinkTableDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDataLinkTableDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateDataLinkTableDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateDataLinkTableDto) SetName(v string) {
	o.Name = &v
}

// GetArchivedAt returns the ArchivedAt field value if set, zero value otherwise.
func (o *UpdateDataLinkTableDto) GetArchivedAt() string {
	if o == nil || IsNil(o.ArchivedAt) {
		var ret string
		return ret
	}
	return *o.ArchivedAt
}

// GetArchivedAtOk returns a tuple with the ArchivedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDataLinkTableDto) GetArchivedAtOk() (*string, bool) {
	if o == nil || IsNil(o.ArchivedAt) {
		return nil, false
	}
	return o.ArchivedAt, true
}

// HasArchivedAt returns a boolean if a field has been set.
func (o *UpdateDataLinkTableDto) HasArchivedAt() bool {
	if o != nil && !IsNil(o.ArchivedAt) {
		return true
	}

	return false
}

// SetArchivedAt gets a reference to the given string and assigns it to the ArchivedAt field.
func (o *UpdateDataLinkTableDto) SetArchivedAt(v string) {
	o.ArchivedAt = &v
}

func (o UpdateDataLinkTableDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDataLinkTableDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ArchivedAt) {
		toSerialize["archivedAt"] = o.ArchivedAt
	}
	return toSerialize, nil
}

type NullableUpdateDataLinkTableDto struct {
	value *UpdateDataLinkTableDto
	isSet bool
}

func (v NullableUpdateDataLinkTableDto) Get() *UpdateDataLinkTableDto {
	return v.value
}

func (v *NullableUpdateDataLinkTableDto) Set(val *UpdateDataLinkTableDto) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDataLinkTableDto) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDataLinkTableDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDataLinkTableDto(val *UpdateDataLinkTableDto) *NullableUpdateDataLinkTableDto {
	return &NullableUpdateDataLinkTableDto{value: val, isSet: true}
}

func (v NullableUpdateDataLinkTableDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDataLinkTableDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


