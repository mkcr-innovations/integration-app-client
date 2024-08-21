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

// checks if the CreateCustomerDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCustomerDto{}

// CreateCustomerDto struct for CreateCustomerDto
type CreateCustomerDto struct {
	InternalId *string `json:"internalId,omitempty"`
	Name *string `json:"name,omitempty"`
	Fields map[string]interface{} `json:"fields,omitempty"`
	Credentials map[string]interface{} `json:"credentials,omitempty"`
	IsTest *bool `json:"isTest,omitempty"`
}

// NewCreateCustomerDto instantiates a new CreateCustomerDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCustomerDto() *CreateCustomerDto {
	this := CreateCustomerDto{}
	return &this
}

// NewCreateCustomerDtoWithDefaults instantiates a new CreateCustomerDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCustomerDtoWithDefaults() *CreateCustomerDto {
	this := CreateCustomerDto{}
	return &this
}

// GetInternalId returns the InternalId field value if set, zero value otherwise.
func (o *CreateCustomerDto) GetInternalId() string {
	if o == nil || IsNil(o.InternalId) {
		var ret string
		return ret
	}
	return *o.InternalId
}

// GetInternalIdOk returns a tuple with the InternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCustomerDto) GetInternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.InternalId) {
		return nil, false
	}
	return o.InternalId, true
}

// HasInternalId returns a boolean if a field has been set.
func (o *CreateCustomerDto) HasInternalId() bool {
	if o != nil && !IsNil(o.InternalId) {
		return true
	}

	return false
}

// SetInternalId gets a reference to the given string and assigns it to the InternalId field.
func (o *CreateCustomerDto) SetInternalId(v string) {
	o.InternalId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateCustomerDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCustomerDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateCustomerDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateCustomerDto) SetName(v string) {
	o.Name = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *CreateCustomerDto) GetFields() map[string]interface{} {
	if o == nil || IsNil(o.Fields) {
		var ret map[string]interface{}
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCustomerDto) GetFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Fields) {
		return map[string]interface{}{}, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *CreateCustomerDto) HasFields() bool {
	if o != nil && !IsNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given map[string]interface{} and assigns it to the Fields field.
func (o *CreateCustomerDto) SetFields(v map[string]interface{}) {
	o.Fields = v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *CreateCustomerDto) GetCredentials() map[string]interface{} {
	if o == nil || IsNil(o.Credentials) {
		var ret map[string]interface{}
		return ret
	}
	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCustomerDto) GetCredentialsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Credentials) {
		return map[string]interface{}{}, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *CreateCustomerDto) HasCredentials() bool {
	if o != nil && !IsNil(o.Credentials) {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given map[string]interface{} and assigns it to the Credentials field.
func (o *CreateCustomerDto) SetCredentials(v map[string]interface{}) {
	o.Credentials = v
}

// GetIsTest returns the IsTest field value if set, zero value otherwise.
func (o *CreateCustomerDto) GetIsTest() bool {
	if o == nil || IsNil(o.IsTest) {
		var ret bool
		return ret
	}
	return *o.IsTest
}

// GetIsTestOk returns a tuple with the IsTest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCustomerDto) GetIsTestOk() (*bool, bool) {
	if o == nil || IsNil(o.IsTest) {
		return nil, false
	}
	return o.IsTest, true
}

// HasIsTest returns a boolean if a field has been set.
func (o *CreateCustomerDto) HasIsTest() bool {
	if o != nil && !IsNil(o.IsTest) {
		return true
	}

	return false
}

// SetIsTest gets a reference to the given bool and assigns it to the IsTest field.
func (o *CreateCustomerDto) SetIsTest(v bool) {
	o.IsTest = &v
}

func (o CreateCustomerDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCustomerDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InternalId) {
		toSerialize["internalId"] = o.InternalId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Fields) {
		toSerialize["fields"] = o.Fields
	}
	if !IsNil(o.Credentials) {
		toSerialize["credentials"] = o.Credentials
	}
	if !IsNil(o.IsTest) {
		toSerialize["isTest"] = o.IsTest
	}
	return toSerialize, nil
}

type NullableCreateCustomerDto struct {
	value *CreateCustomerDto
	isSet bool
}

func (v NullableCreateCustomerDto) Get() *CreateCustomerDto {
	return v.value
}

func (v *NullableCreateCustomerDto) Set(val *CreateCustomerDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCustomerDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCustomerDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCustomerDto(val *CreateCustomerDto) *NullableCreateCustomerDto {
	return &NullableCreateCustomerDto{value: val, isSet: true}
}

func (v NullableCreateCustomerDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCustomerDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


