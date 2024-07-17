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

// checks if the ActionInstancesControllerListActionInstances200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActionInstancesControllerListActionInstances200Response{}

// ActionInstancesControllerListActionInstances200Response struct for ActionInstancesControllerListActionInstances200Response
type ActionInstancesControllerListActionInstances200Response struct {
	Cursor *string `json:"cursor,omitempty"`
	Items []ActionInstanceDto `json:"items"`
}

type _ActionInstancesControllerListActionInstances200Response ActionInstancesControllerListActionInstances200Response

// NewActionInstancesControllerListActionInstances200Response instantiates a new ActionInstancesControllerListActionInstances200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionInstancesControllerListActionInstances200Response(items []ActionInstanceDto) *ActionInstancesControllerListActionInstances200Response {
	this := ActionInstancesControllerListActionInstances200Response{}
	this.Items = items
	return &this
}

// NewActionInstancesControllerListActionInstances200ResponseWithDefaults instantiates a new ActionInstancesControllerListActionInstances200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionInstancesControllerListActionInstances200ResponseWithDefaults() *ActionInstancesControllerListActionInstances200Response {
	this := ActionInstancesControllerListActionInstances200Response{}
	return &this
}

// GetCursor returns the Cursor field value if set, zero value otherwise.
func (o *ActionInstancesControllerListActionInstances200Response) GetCursor() string {
	if o == nil || IsNil(o.Cursor) {
		var ret string
		return ret
	}
	return *o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionInstancesControllerListActionInstances200Response) GetCursorOk() (*string, bool) {
	if o == nil || IsNil(o.Cursor) {
		return nil, false
	}
	return o.Cursor, true
}

// HasCursor returns a boolean if a field has been set.
func (o *ActionInstancesControllerListActionInstances200Response) HasCursor() bool {
	if o != nil && !IsNil(o.Cursor) {
		return true
	}

	return false
}

// SetCursor gets a reference to the given string and assigns it to the Cursor field.
func (o *ActionInstancesControllerListActionInstances200Response) SetCursor(v string) {
	o.Cursor = &v
}

// GetItems returns the Items field value
func (o *ActionInstancesControllerListActionInstances200Response) GetItems() []ActionInstanceDto {
	if o == nil {
		var ret []ActionInstanceDto
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ActionInstancesControllerListActionInstances200Response) GetItemsOk() ([]ActionInstanceDto, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *ActionInstancesControllerListActionInstances200Response) SetItems(v []ActionInstanceDto) {
	o.Items = v
}

func (o ActionInstancesControllerListActionInstances200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActionInstancesControllerListActionInstances200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cursor) {
		toSerialize["cursor"] = o.Cursor
	}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

func (o *ActionInstancesControllerListActionInstances200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"items",
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

	varActionInstancesControllerListActionInstances200Response := _ActionInstancesControllerListActionInstances200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varActionInstancesControllerListActionInstances200Response)

	if err != nil {
		return err
	}

	*o = ActionInstancesControllerListActionInstances200Response(varActionInstancesControllerListActionInstances200Response)

	return err
}

type NullableActionInstancesControllerListActionInstances200Response struct {
	value *ActionInstancesControllerListActionInstances200Response
	isSet bool
}

func (v NullableActionInstancesControllerListActionInstances200Response) Get() *ActionInstancesControllerListActionInstances200Response {
	return v.value
}

func (v *NullableActionInstancesControllerListActionInstances200Response) Set(val *ActionInstancesControllerListActionInstances200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableActionInstancesControllerListActionInstances200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableActionInstancesControllerListActionInstances200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionInstancesControllerListActionInstances200Response(val *ActionInstancesControllerListActionInstances200Response) *NullableActionInstancesControllerListActionInstances200Response {
	return &NullableActionInstancesControllerListActionInstances200Response{value: val, isSet: true}
}

func (v NullableActionInstancesControllerListActionInstances200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionInstancesControllerListActionInstances200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

