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

// checks if the CreateScenarioDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateScenarioDto{}

// CreateScenarioDto struct for CreateScenarioDto
type CreateScenarioDto struct {
	Name string `json:"name"`
	Key string `json:"key"`
	ShortDescription *string `json:"shortDescription,omitempty"`
	Description *string `json:"description,omitempty"`
	Elements []map[string]interface{} `json:"elements,omitempty"`
	Todos []map[string]interface{} `json:"todos,omitempty"`
}

type _CreateScenarioDto CreateScenarioDto

// NewCreateScenarioDto instantiates a new CreateScenarioDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateScenarioDto(name string, key string) *CreateScenarioDto {
	this := CreateScenarioDto{}
	this.Name = name
	this.Key = key
	return &this
}

// NewCreateScenarioDtoWithDefaults instantiates a new CreateScenarioDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateScenarioDtoWithDefaults() *CreateScenarioDto {
	this := CreateScenarioDto{}
	return &this
}

// GetName returns the Name field value
func (o *CreateScenarioDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateScenarioDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateScenarioDto) SetName(v string) {
	o.Name = v
}

// GetKey returns the Key field value
func (o *CreateScenarioDto) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *CreateScenarioDto) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *CreateScenarioDto) SetKey(v string) {
	o.Key = v
}

// GetShortDescription returns the ShortDescription field value if set, zero value otherwise.
func (o *CreateScenarioDto) GetShortDescription() string {
	if o == nil || IsNil(o.ShortDescription) {
		var ret string
		return ret
	}
	return *o.ShortDescription
}

// GetShortDescriptionOk returns a tuple with the ShortDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateScenarioDto) GetShortDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ShortDescription) {
		return nil, false
	}
	return o.ShortDescription, true
}

// HasShortDescription returns a boolean if a field has been set.
func (o *CreateScenarioDto) HasShortDescription() bool {
	if o != nil && !IsNil(o.ShortDescription) {
		return true
	}

	return false
}

// SetShortDescription gets a reference to the given string and assigns it to the ShortDescription field.
func (o *CreateScenarioDto) SetShortDescription(v string) {
	o.ShortDescription = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateScenarioDto) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateScenarioDto) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateScenarioDto) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateScenarioDto) SetDescription(v string) {
	o.Description = &v
}

// GetElements returns the Elements field value if set, zero value otherwise.
func (o *CreateScenarioDto) GetElements() []map[string]interface{} {
	if o == nil || IsNil(o.Elements) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Elements
}

// GetElementsOk returns a tuple with the Elements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateScenarioDto) GetElementsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Elements) {
		return nil, false
	}
	return o.Elements, true
}

// HasElements returns a boolean if a field has been set.
func (o *CreateScenarioDto) HasElements() bool {
	if o != nil && !IsNil(o.Elements) {
		return true
	}

	return false
}

// SetElements gets a reference to the given []map[string]interface{} and assigns it to the Elements field.
func (o *CreateScenarioDto) SetElements(v []map[string]interface{}) {
	o.Elements = v
}

// GetTodos returns the Todos field value if set, zero value otherwise.
func (o *CreateScenarioDto) GetTodos() []map[string]interface{} {
	if o == nil || IsNil(o.Todos) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Todos
}

// GetTodosOk returns a tuple with the Todos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateScenarioDto) GetTodosOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Todos) {
		return nil, false
	}
	return o.Todos, true
}

// HasTodos returns a boolean if a field has been set.
func (o *CreateScenarioDto) HasTodos() bool {
	if o != nil && !IsNil(o.Todos) {
		return true
	}

	return false
}

// SetTodos gets a reference to the given []map[string]interface{} and assigns it to the Todos field.
func (o *CreateScenarioDto) SetTodos(v []map[string]interface{}) {
	o.Todos = v
}

func (o CreateScenarioDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateScenarioDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["key"] = o.Key
	if !IsNil(o.ShortDescription) {
		toSerialize["shortDescription"] = o.ShortDescription
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Elements) {
		toSerialize["elements"] = o.Elements
	}
	if !IsNil(o.Todos) {
		toSerialize["todos"] = o.Todos
	}
	return toSerialize, nil
}

func (o *CreateScenarioDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varCreateScenarioDto := _CreateScenarioDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateScenarioDto)

	if err != nil {
		return err
	}

	*o = CreateScenarioDto(varCreateScenarioDto)

	return err
}

type NullableCreateScenarioDto struct {
	value *CreateScenarioDto
	isSet bool
}

func (v NullableCreateScenarioDto) Get() *CreateScenarioDto {
	return v.value
}

func (v *NullableCreateScenarioDto) Set(val *CreateScenarioDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateScenarioDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateScenarioDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateScenarioDto(val *CreateScenarioDto) *NullableCreateScenarioDto {
	return &NullableCreateScenarioDto{value: val, isSet: true}
}

func (v NullableCreateScenarioDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateScenarioDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


