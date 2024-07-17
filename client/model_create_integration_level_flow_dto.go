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

// checks if the CreateIntegrationLevelFlowDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateIntegrationLevelFlowDto{}

// CreateIntegrationLevelFlowDto struct for CreateIntegrationLevelFlowDto
type CreateIntegrationLevelFlowDto struct {
	Key string `json:"key"`
	Name string `json:"name"`
	ArchivedAt *string `json:"archivedAt,omitempty"`
	Nodes map[string]interface{} `json:"nodes,omitempty"`
	ParametersSchema map[string]interface{} `json:"parametersSchema,omitempty"`
	AutoCreateInstances *bool `json:"autoCreateInstances,omitempty"`
}

type _CreateIntegrationLevelFlowDto CreateIntegrationLevelFlowDto

// NewCreateIntegrationLevelFlowDto instantiates a new CreateIntegrationLevelFlowDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateIntegrationLevelFlowDto(key string, name string) *CreateIntegrationLevelFlowDto {
	this := CreateIntegrationLevelFlowDto{}
	this.Key = key
	this.Name = name
	return &this
}

// NewCreateIntegrationLevelFlowDtoWithDefaults instantiates a new CreateIntegrationLevelFlowDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateIntegrationLevelFlowDtoWithDefaults() *CreateIntegrationLevelFlowDto {
	this := CreateIntegrationLevelFlowDto{}
	return &this
}

// GetKey returns the Key field value
func (o *CreateIntegrationLevelFlowDto) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *CreateIntegrationLevelFlowDto) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *CreateIntegrationLevelFlowDto) SetKey(v string) {
	o.Key = v
}

// GetName returns the Name field value
func (o *CreateIntegrationLevelFlowDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateIntegrationLevelFlowDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateIntegrationLevelFlowDto) SetName(v string) {
	o.Name = v
}

// GetArchivedAt returns the ArchivedAt field value if set, zero value otherwise.
func (o *CreateIntegrationLevelFlowDto) GetArchivedAt() string {
	if o == nil || IsNil(o.ArchivedAt) {
		var ret string
		return ret
	}
	return *o.ArchivedAt
}

// GetArchivedAtOk returns a tuple with the ArchivedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIntegrationLevelFlowDto) GetArchivedAtOk() (*string, bool) {
	if o == nil || IsNil(o.ArchivedAt) {
		return nil, false
	}
	return o.ArchivedAt, true
}

// HasArchivedAt returns a boolean if a field has been set.
func (o *CreateIntegrationLevelFlowDto) HasArchivedAt() bool {
	if o != nil && !IsNil(o.ArchivedAt) {
		return true
	}

	return false
}

// SetArchivedAt gets a reference to the given string and assigns it to the ArchivedAt field.
func (o *CreateIntegrationLevelFlowDto) SetArchivedAt(v string) {
	o.ArchivedAt = &v
}

// GetNodes returns the Nodes field value if set, zero value otherwise.
func (o *CreateIntegrationLevelFlowDto) GetNodes() map[string]interface{} {
	if o == nil || IsNil(o.Nodes) {
		var ret map[string]interface{}
		return ret
	}
	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIntegrationLevelFlowDto) GetNodesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Nodes) {
		return map[string]interface{}{}, false
	}
	return o.Nodes, true
}

// HasNodes returns a boolean if a field has been set.
func (o *CreateIntegrationLevelFlowDto) HasNodes() bool {
	if o != nil && !IsNil(o.Nodes) {
		return true
	}

	return false
}

// SetNodes gets a reference to the given map[string]interface{} and assigns it to the Nodes field.
func (o *CreateIntegrationLevelFlowDto) SetNodes(v map[string]interface{}) {
	o.Nodes = v
}

// GetParametersSchema returns the ParametersSchema field value if set, zero value otherwise.
func (o *CreateIntegrationLevelFlowDto) GetParametersSchema() map[string]interface{} {
	if o == nil || IsNil(o.ParametersSchema) {
		var ret map[string]interface{}
		return ret
	}
	return o.ParametersSchema
}

// GetParametersSchemaOk returns a tuple with the ParametersSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIntegrationLevelFlowDto) GetParametersSchemaOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ParametersSchema) {
		return map[string]interface{}{}, false
	}
	return o.ParametersSchema, true
}

// HasParametersSchema returns a boolean if a field has been set.
func (o *CreateIntegrationLevelFlowDto) HasParametersSchema() bool {
	if o != nil && !IsNil(o.ParametersSchema) {
		return true
	}

	return false
}

// SetParametersSchema gets a reference to the given map[string]interface{} and assigns it to the ParametersSchema field.
func (o *CreateIntegrationLevelFlowDto) SetParametersSchema(v map[string]interface{}) {
	o.ParametersSchema = v
}

// GetAutoCreateInstances returns the AutoCreateInstances field value if set, zero value otherwise.
func (o *CreateIntegrationLevelFlowDto) GetAutoCreateInstances() bool {
	if o == nil || IsNil(o.AutoCreateInstances) {
		var ret bool
		return ret
	}
	return *o.AutoCreateInstances
}

// GetAutoCreateInstancesOk returns a tuple with the AutoCreateInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIntegrationLevelFlowDto) GetAutoCreateInstancesOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoCreateInstances) {
		return nil, false
	}
	return o.AutoCreateInstances, true
}

// HasAutoCreateInstances returns a boolean if a field has been set.
func (o *CreateIntegrationLevelFlowDto) HasAutoCreateInstances() bool {
	if o != nil && !IsNil(o.AutoCreateInstances) {
		return true
	}

	return false
}

// SetAutoCreateInstances gets a reference to the given bool and assigns it to the AutoCreateInstances field.
func (o *CreateIntegrationLevelFlowDto) SetAutoCreateInstances(v bool) {
	o.AutoCreateInstances = &v
}

func (o CreateIntegrationLevelFlowDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateIntegrationLevelFlowDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["name"] = o.Name
	if !IsNil(o.ArchivedAt) {
		toSerialize["archivedAt"] = o.ArchivedAt
	}
	if !IsNil(o.Nodes) {
		toSerialize["nodes"] = o.Nodes
	}
	if !IsNil(o.ParametersSchema) {
		toSerialize["parametersSchema"] = o.ParametersSchema
	}
	if !IsNil(o.AutoCreateInstances) {
		toSerialize["autoCreateInstances"] = o.AutoCreateInstances
	}
	return toSerialize, nil
}

func (o *CreateIntegrationLevelFlowDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"key",
		"name",
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

	varCreateIntegrationLevelFlowDto := _CreateIntegrationLevelFlowDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateIntegrationLevelFlowDto)

	if err != nil {
		return err
	}

	*o = CreateIntegrationLevelFlowDto(varCreateIntegrationLevelFlowDto)

	return err
}

type NullableCreateIntegrationLevelFlowDto struct {
	value *CreateIntegrationLevelFlowDto
	isSet bool
}

func (v NullableCreateIntegrationLevelFlowDto) Get() *CreateIntegrationLevelFlowDto {
	return v.value
}

func (v *NullableCreateIntegrationLevelFlowDto) Set(val *CreateIntegrationLevelFlowDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateIntegrationLevelFlowDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateIntegrationLevelFlowDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateIntegrationLevelFlowDto(val *CreateIntegrationLevelFlowDto) *NullableCreateIntegrationLevelFlowDto {
	return &NullableCreateIntegrationLevelFlowDto{value: val, isSet: true}
}

func (v NullableCreateIntegrationLevelFlowDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateIntegrationLevelFlowDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

