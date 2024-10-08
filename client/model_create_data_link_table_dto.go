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

// checks if the CreateDataLinkTableDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateDataLinkTableDto{}

// CreateDataLinkTableDto struct for CreateDataLinkTableDto
type CreateDataLinkTableDto struct {
	Key string `json:"key"`
	Name *string `json:"name,omitempty"`
	ArchivedAt *string `json:"archivedAt,omitempty"`
}

type _CreateDataLinkTableDto CreateDataLinkTableDto

// NewCreateDataLinkTableDto instantiates a new CreateDataLinkTableDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDataLinkTableDto(key string) *CreateDataLinkTableDto {
	this := CreateDataLinkTableDto{}
	this.Key = key
	return &this
}

// NewCreateDataLinkTableDtoWithDefaults instantiates a new CreateDataLinkTableDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDataLinkTableDtoWithDefaults() *CreateDataLinkTableDto {
	this := CreateDataLinkTableDto{}
	return &this
}

// GetKey returns the Key field value
func (o *CreateDataLinkTableDto) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *CreateDataLinkTableDto) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *CreateDataLinkTableDto) SetKey(v string) {
	o.Key = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateDataLinkTableDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDataLinkTableDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateDataLinkTableDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateDataLinkTableDto) SetName(v string) {
	o.Name = &v
}

// GetArchivedAt returns the ArchivedAt field value if set, zero value otherwise.
func (o *CreateDataLinkTableDto) GetArchivedAt() string {
	if o == nil || IsNil(o.ArchivedAt) {
		var ret string
		return ret
	}
	return *o.ArchivedAt
}

// GetArchivedAtOk returns a tuple with the ArchivedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDataLinkTableDto) GetArchivedAtOk() (*string, bool) {
	if o == nil || IsNil(o.ArchivedAt) {
		return nil, false
	}
	return o.ArchivedAt, true
}

// HasArchivedAt returns a boolean if a field has been set.
func (o *CreateDataLinkTableDto) HasArchivedAt() bool {
	if o != nil && !IsNil(o.ArchivedAt) {
		return true
	}

	return false
}

// SetArchivedAt gets a reference to the given string and assigns it to the ArchivedAt field.
func (o *CreateDataLinkTableDto) SetArchivedAt(v string) {
	o.ArchivedAt = &v
}

func (o CreateDataLinkTableDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateDataLinkTableDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ArchivedAt) {
		toSerialize["archivedAt"] = o.ArchivedAt
	}
	return toSerialize, nil
}

func (o *CreateDataLinkTableDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"key",
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

	varCreateDataLinkTableDto := _CreateDataLinkTableDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateDataLinkTableDto)

	if err != nil {
		return err
	}

	*o = CreateDataLinkTableDto(varCreateDataLinkTableDto)

	return err
}

type NullableCreateDataLinkTableDto struct {
	value *CreateDataLinkTableDto
	isSet bool
}

func (v NullableCreateDataLinkTableDto) Get() *CreateDataLinkTableDto {
	return v.value
}

func (v *NullableCreateDataLinkTableDto) Set(val *CreateDataLinkTableDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDataLinkTableDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDataLinkTableDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDataLinkTableDto(val *CreateDataLinkTableDto) *NullableCreateDataLinkTableDto {
	return &NullableCreateDataLinkTableDto{value: val, isSet: true}
}

func (v NullableCreateDataLinkTableDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDataLinkTableDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


