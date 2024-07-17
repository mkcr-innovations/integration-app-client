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

// checks if the AppDataSchemasControllerListAppDataSchemas200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppDataSchemasControllerListAppDataSchemas200Response{}

// AppDataSchemasControllerListAppDataSchemas200Response struct for AppDataSchemasControllerListAppDataSchemas200Response
type AppDataSchemasControllerListAppDataSchemas200Response struct {
	Cursor *string `json:"cursor,omitempty"`
	Items []AppDataSchemaDto `json:"items"`
}

type _AppDataSchemasControllerListAppDataSchemas200Response AppDataSchemasControllerListAppDataSchemas200Response

// NewAppDataSchemasControllerListAppDataSchemas200Response instantiates a new AppDataSchemasControllerListAppDataSchemas200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppDataSchemasControllerListAppDataSchemas200Response(items []AppDataSchemaDto) *AppDataSchemasControllerListAppDataSchemas200Response {
	this := AppDataSchemasControllerListAppDataSchemas200Response{}
	this.Items = items
	return &this
}

// NewAppDataSchemasControllerListAppDataSchemas200ResponseWithDefaults instantiates a new AppDataSchemasControllerListAppDataSchemas200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppDataSchemasControllerListAppDataSchemas200ResponseWithDefaults() *AppDataSchemasControllerListAppDataSchemas200Response {
	this := AppDataSchemasControllerListAppDataSchemas200Response{}
	return &this
}

// GetCursor returns the Cursor field value if set, zero value otherwise.
func (o *AppDataSchemasControllerListAppDataSchemas200Response) GetCursor() string {
	if o == nil || IsNil(o.Cursor) {
		var ret string
		return ret
	}
	return *o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDataSchemasControllerListAppDataSchemas200Response) GetCursorOk() (*string, bool) {
	if o == nil || IsNil(o.Cursor) {
		return nil, false
	}
	return o.Cursor, true
}

// HasCursor returns a boolean if a field has been set.
func (o *AppDataSchemasControllerListAppDataSchemas200Response) HasCursor() bool {
	if o != nil && !IsNil(o.Cursor) {
		return true
	}

	return false
}

// SetCursor gets a reference to the given string and assigns it to the Cursor field.
func (o *AppDataSchemasControllerListAppDataSchemas200Response) SetCursor(v string) {
	o.Cursor = &v
}

// GetItems returns the Items field value
func (o *AppDataSchemasControllerListAppDataSchemas200Response) GetItems() []AppDataSchemaDto {
	if o == nil {
		var ret []AppDataSchemaDto
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *AppDataSchemasControllerListAppDataSchemas200Response) GetItemsOk() ([]AppDataSchemaDto, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *AppDataSchemasControllerListAppDataSchemas200Response) SetItems(v []AppDataSchemaDto) {
	o.Items = v
}

func (o AppDataSchemasControllerListAppDataSchemas200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppDataSchemasControllerListAppDataSchemas200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cursor) {
		toSerialize["cursor"] = o.Cursor
	}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

func (o *AppDataSchemasControllerListAppDataSchemas200Response) UnmarshalJSON(data []byte) (err error) {
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

	varAppDataSchemasControllerListAppDataSchemas200Response := _AppDataSchemasControllerListAppDataSchemas200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppDataSchemasControllerListAppDataSchemas200Response)

	if err != nil {
		return err
	}

	*o = AppDataSchemasControllerListAppDataSchemas200Response(varAppDataSchemasControllerListAppDataSchemas200Response)

	return err
}

type NullableAppDataSchemasControllerListAppDataSchemas200Response struct {
	value *AppDataSchemasControllerListAppDataSchemas200Response
	isSet bool
}

func (v NullableAppDataSchemasControllerListAppDataSchemas200Response) Get() *AppDataSchemasControllerListAppDataSchemas200Response {
	return v.value
}

func (v *NullableAppDataSchemasControllerListAppDataSchemas200Response) Set(val *AppDataSchemasControllerListAppDataSchemas200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAppDataSchemasControllerListAppDataSchemas200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAppDataSchemasControllerListAppDataSchemas200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppDataSchemasControllerListAppDataSchemas200Response(val *AppDataSchemasControllerListAppDataSchemas200Response) *NullableAppDataSchemasControllerListAppDataSchemas200Response {
	return &NullableAppDataSchemasControllerListAppDataSchemas200Response{value: val, isSet: true}
}

func (v NullableAppDataSchemasControllerListAppDataSchemas200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppDataSchemasControllerListAppDataSchemas200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

