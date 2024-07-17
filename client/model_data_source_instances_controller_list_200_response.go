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

// checks if the DataSourceInstancesControllerList200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataSourceInstancesControllerList200Response{}

// DataSourceInstancesControllerList200Response struct for DataSourceInstancesControllerList200Response
type DataSourceInstancesControllerList200Response struct {
	Cursor *string `json:"cursor,omitempty"`
	Items []DataSourceInstanceDto `json:"items"`
}

type _DataSourceInstancesControllerList200Response DataSourceInstancesControllerList200Response

// NewDataSourceInstancesControllerList200Response instantiates a new DataSourceInstancesControllerList200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataSourceInstancesControllerList200Response(items []DataSourceInstanceDto) *DataSourceInstancesControllerList200Response {
	this := DataSourceInstancesControllerList200Response{}
	this.Items = items
	return &this
}

// NewDataSourceInstancesControllerList200ResponseWithDefaults instantiates a new DataSourceInstancesControllerList200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataSourceInstancesControllerList200ResponseWithDefaults() *DataSourceInstancesControllerList200Response {
	this := DataSourceInstancesControllerList200Response{}
	return &this
}

// GetCursor returns the Cursor field value if set, zero value otherwise.
func (o *DataSourceInstancesControllerList200Response) GetCursor() string {
	if o == nil || IsNil(o.Cursor) {
		var ret string
		return ret
	}
	return *o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourceInstancesControllerList200Response) GetCursorOk() (*string, bool) {
	if o == nil || IsNil(o.Cursor) {
		return nil, false
	}
	return o.Cursor, true
}

// HasCursor returns a boolean if a field has been set.
func (o *DataSourceInstancesControllerList200Response) HasCursor() bool {
	if o != nil && !IsNil(o.Cursor) {
		return true
	}

	return false
}

// SetCursor gets a reference to the given string and assigns it to the Cursor field.
func (o *DataSourceInstancesControllerList200Response) SetCursor(v string) {
	o.Cursor = &v
}

// GetItems returns the Items field value
func (o *DataSourceInstancesControllerList200Response) GetItems() []DataSourceInstanceDto {
	if o == nil {
		var ret []DataSourceInstanceDto
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *DataSourceInstancesControllerList200Response) GetItemsOk() ([]DataSourceInstanceDto, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *DataSourceInstancesControllerList200Response) SetItems(v []DataSourceInstanceDto) {
	o.Items = v
}

func (o DataSourceInstancesControllerList200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataSourceInstancesControllerList200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cursor) {
		toSerialize["cursor"] = o.Cursor
	}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

func (o *DataSourceInstancesControllerList200Response) UnmarshalJSON(data []byte) (err error) {
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

	varDataSourceInstancesControllerList200Response := _DataSourceInstancesControllerList200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDataSourceInstancesControllerList200Response)

	if err != nil {
		return err
	}

	*o = DataSourceInstancesControllerList200Response(varDataSourceInstancesControllerList200Response)

	return err
}

type NullableDataSourceInstancesControllerList200Response struct {
	value *DataSourceInstancesControllerList200Response
	isSet bool
}

func (v NullableDataSourceInstancesControllerList200Response) Get() *DataSourceInstancesControllerList200Response {
	return v.value
}

func (v *NullableDataSourceInstancesControllerList200Response) Set(val *DataSourceInstancesControllerList200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDataSourceInstancesControllerList200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDataSourceInstancesControllerList200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataSourceInstancesControllerList200Response(val *DataSourceInstancesControllerList200Response) *NullableDataSourceInstancesControllerList200Response {
	return &NullableDataSourceInstancesControllerList200Response{value: val, isSet: true}
}

func (v NullableDataSourceInstancesControllerList200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataSourceInstancesControllerList200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


