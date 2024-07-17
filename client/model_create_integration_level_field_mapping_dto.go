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

// checks if the CreateIntegrationLevelFieldMappingDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateIntegrationLevelFieldMappingDto{}

// CreateIntegrationLevelFieldMappingDto struct for CreateIntegrationLevelFieldMappingDto
type CreateIntegrationLevelFieldMappingDto struct {
	Key string `json:"key"`
	Name string `json:"name"`
	ArchivedAt *string `json:"archivedAt,omitempty"`
	DataSourceKey *string `json:"dataSourceKey,omitempty"`
	DataSourceId *string `json:"dataSourceId,omitempty"`
	AppSchema map[string]interface{} `json:"appSchema,omitempty"`
	Direction *string `json:"direction,omitempty"`
	DefaultImportValue map[string]interface{} `json:"defaultImportValue,omitempty"`
	ExportValue map[string]interface{} `json:"exportValue,omitempty"`
	ImportValue map[string]interface{} `json:"importValue,omitempty"`
	DefaultExportValue map[string]interface{} `json:"defaultExportValue,omitempty"`
	FrozenImportFields map[string]interface{} `json:"frozenImportFields,omitempty"`
	FrozenExportFields map[string]interface{} `json:"frozenExportFields,omitempty"`
	FrozenUnifiedExportFields map[string]interface{} `json:"frozenUnifiedExportFields,omitempty"`
}

type _CreateIntegrationLevelFieldMappingDto CreateIntegrationLevelFieldMappingDto

// NewCreateIntegrationLevelFieldMappingDto instantiates a new CreateIntegrationLevelFieldMappingDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateIntegrationLevelFieldMappingDto(key string, name string) *CreateIntegrationLevelFieldMappingDto {
	this := CreateIntegrationLevelFieldMappingDto{}
	this.Key = key
	this.Name = name
	return &this
}

// NewCreateIntegrationLevelFieldMappingDtoWithDefaults instantiates a new CreateIntegrationLevelFieldMappingDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateIntegrationLevelFieldMappingDtoWithDefaults() *CreateIntegrationLevelFieldMappingDto {
	this := CreateIntegrationLevelFieldMappingDto{}
	return &this
}

// GetKey returns the Key field value
func (o *CreateIntegrationLevelFieldMappingDto) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *CreateIntegrationLevelFieldMappingDto) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *CreateIntegrationLevelFieldMappingDto) SetKey(v string) {
	o.Key = v
}

// GetName returns the Name field value
func (o *CreateIntegrationLevelFieldMappingDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateIntegrationLevelFieldMappingDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateIntegrationLevelFieldMappingDto) SetName(v string) {
	o.Name = v
}

// GetArchivedAt returns the ArchivedAt field value if set, zero value otherwise.
func (o *CreateIntegrationLevelFieldMappingDto) GetArchivedAt() string {
	if o == nil || IsNil(o.ArchivedAt) {
		var ret string
		return ret
	}
	return *o.ArchivedAt
}

// GetArchivedAtOk returns a tuple with the ArchivedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIntegrationLevelFieldMappingDto) GetArchivedAtOk() (*string, bool) {
	if o == nil || IsNil(o.ArchivedAt) {
		return nil, false
	}
	return o.ArchivedAt, true
}

// HasArchivedAt returns a boolean if a field has been set.
func (o *CreateIntegrationLevelFieldMappingDto) HasArchivedAt() bool {
	if o != nil && !IsNil(o.ArchivedAt) {
		return true
	}

	return false
}

// SetArchivedAt gets a reference to the given string and assigns it to the ArchivedAt field.
func (o *CreateIntegrationLevelFieldMappingDto) SetArchivedAt(v string) {
	o.ArchivedAt = &v
}

// GetDataSourceKey returns the DataSourceKey field value if set, zero value otherwise.
func (o *CreateIntegrationLevelFieldMappingDto) GetDataSourceKey() string {
	if o == nil || IsNil(o.DataSourceKey) {
		var ret string
		return ret
	}
	return *o.DataSourceKey
}

// GetDataSourceKeyOk returns a tuple with the DataSourceKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIntegrationLevelFieldMappingDto) GetDataSourceKeyOk() (*string, bool) {
	if o == nil || IsNil(o.DataSourceKey) {
		return nil, false
	}
	return o.DataSourceKey, true
}

// HasDataSourceKey returns a boolean if a field has been set.
func (o *CreateIntegrationLevelFieldMappingDto) HasDataSourceKey() bool {
	if o != nil && !IsNil(o.DataSourceKey) {
		return true
	}

	return false
}

// SetDataSourceKey gets a reference to the given string and assigns it to the DataSourceKey field.
func (o *CreateIntegrationLevelFieldMappingDto) SetDataSourceKey(v string) {
	o.DataSourceKey = &v
}

// GetDataSourceId returns the DataSourceId field value if set, zero value otherwise.
func (o *CreateIntegrationLevelFieldMappingDto) GetDataSourceId() string {
	if o == nil || IsNil(o.DataSourceId) {
		var ret string
		return ret
	}
	return *o.DataSourceId
}

// GetDataSourceIdOk returns a tuple with the DataSourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIntegrationLevelFieldMappingDto) GetDataSourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.DataSourceId) {
		return nil, false
	}
	return o.DataSourceId, true
}

// HasDataSourceId returns a boolean if a field has been set.
func (o *CreateIntegrationLevelFieldMappingDto) HasDataSourceId() bool {
	if o != nil && !IsNil(o.DataSourceId) {
		return true
	}

	return false
}

// SetDataSourceId gets a reference to the given string and assigns it to the DataSourceId field.
func (o *CreateIntegrationLevelFieldMappingDto) SetDataSourceId(v string) {
	o.DataSourceId = &v
}

// GetAppSchema returns the AppSchema field value if set, zero value otherwise.
func (o *CreateIntegrationLevelFieldMappingDto) GetAppSchema() map[string]interface{} {
	if o == nil || IsNil(o.AppSchema) {
		var ret map[string]interface{}
		return ret
	}
	return o.AppSchema
}

// GetAppSchemaOk returns a tuple with the AppSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIntegrationLevelFieldMappingDto) GetAppSchemaOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AppSchema) {
		return map[string]interface{}{}, false
	}
	return o.AppSchema, true
}

// HasAppSchema returns a boolean if a field has been set.
func (o *CreateIntegrationLevelFieldMappingDto) HasAppSchema() bool {
	if o != nil && !IsNil(o.AppSchema) {
		return true
	}

	return false
}

// SetAppSchema gets a reference to the given map[string]interface{} and assigns it to the AppSchema field.
func (o *CreateIntegrationLevelFieldMappingDto) SetAppSchema(v map[string]interface{}) {
	o.AppSchema = v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *CreateIntegrationLevelFieldMappingDto) GetDirection() string {
	if o == nil || IsNil(o.Direction) {
		var ret string
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIntegrationLevelFieldMappingDto) GetDirectionOk() (*string, bool) {
	if o == nil || IsNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *CreateIntegrationLevelFieldMappingDto) HasDirection() bool {
	if o != nil && !IsNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given string and assigns it to the Direction field.
func (o *CreateIntegrationLevelFieldMappingDto) SetDirection(v string) {
	o.Direction = &v
}

// GetDefaultImportValue returns the DefaultImportValue field value if set, zero value otherwise.
func (o *CreateIntegrationLevelFieldMappingDto) GetDefaultImportValue() map[string]interface{} {
	if o == nil || IsNil(o.DefaultImportValue) {
		var ret map[string]interface{}
		return ret
	}
	return o.DefaultImportValue
}

// GetDefaultImportValueOk returns a tuple with the DefaultImportValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIntegrationLevelFieldMappingDto) GetDefaultImportValueOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.DefaultImportValue) {
		return map[string]interface{}{}, false
	}
	return o.DefaultImportValue, true
}

// HasDefaultImportValue returns a boolean if a field has been set.
func (o *CreateIntegrationLevelFieldMappingDto) HasDefaultImportValue() bool {
	if o != nil && !IsNil(o.DefaultImportValue) {
		return true
	}

	return false
}

// SetDefaultImportValue gets a reference to the given map[string]interface{} and assigns it to the DefaultImportValue field.
func (o *CreateIntegrationLevelFieldMappingDto) SetDefaultImportValue(v map[string]interface{}) {
	o.DefaultImportValue = v
}

// GetExportValue returns the ExportValue field value if set, zero value otherwise.
func (o *CreateIntegrationLevelFieldMappingDto) GetExportValue() map[string]interface{} {
	if o == nil || IsNil(o.ExportValue) {
		var ret map[string]interface{}
		return ret
	}
	return o.ExportValue
}

// GetExportValueOk returns a tuple with the ExportValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIntegrationLevelFieldMappingDto) GetExportValueOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ExportValue) {
		return map[string]interface{}{}, false
	}
	return o.ExportValue, true
}

// HasExportValue returns a boolean if a field has been set.
func (o *CreateIntegrationLevelFieldMappingDto) HasExportValue() bool {
	if o != nil && !IsNil(o.ExportValue) {
		return true
	}

	return false
}

// SetExportValue gets a reference to the given map[string]interface{} and assigns it to the ExportValue field.
func (o *CreateIntegrationLevelFieldMappingDto) SetExportValue(v map[string]interface{}) {
	o.ExportValue = v
}

// GetImportValue returns the ImportValue field value if set, zero value otherwise.
func (o *CreateIntegrationLevelFieldMappingDto) GetImportValue() map[string]interface{} {
	if o == nil || IsNil(o.ImportValue) {
		var ret map[string]interface{}
		return ret
	}
	return o.ImportValue
}

// GetImportValueOk returns a tuple with the ImportValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIntegrationLevelFieldMappingDto) GetImportValueOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ImportValue) {
		return map[string]interface{}{}, false
	}
	return o.ImportValue, true
}

// HasImportValue returns a boolean if a field has been set.
func (o *CreateIntegrationLevelFieldMappingDto) HasImportValue() bool {
	if o != nil && !IsNil(o.ImportValue) {
		return true
	}

	return false
}

// SetImportValue gets a reference to the given map[string]interface{} and assigns it to the ImportValue field.
func (o *CreateIntegrationLevelFieldMappingDto) SetImportValue(v map[string]interface{}) {
	o.ImportValue = v
}

// GetDefaultExportValue returns the DefaultExportValue field value if set, zero value otherwise.
func (o *CreateIntegrationLevelFieldMappingDto) GetDefaultExportValue() map[string]interface{} {
	if o == nil || IsNil(o.DefaultExportValue) {
		var ret map[string]interface{}
		return ret
	}
	return o.DefaultExportValue
}

// GetDefaultExportValueOk returns a tuple with the DefaultExportValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIntegrationLevelFieldMappingDto) GetDefaultExportValueOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.DefaultExportValue) {
		return map[string]interface{}{}, false
	}
	return o.DefaultExportValue, true
}

// HasDefaultExportValue returns a boolean if a field has been set.
func (o *CreateIntegrationLevelFieldMappingDto) HasDefaultExportValue() bool {
	if o != nil && !IsNil(o.DefaultExportValue) {
		return true
	}

	return false
}

// SetDefaultExportValue gets a reference to the given map[string]interface{} and assigns it to the DefaultExportValue field.
func (o *CreateIntegrationLevelFieldMappingDto) SetDefaultExportValue(v map[string]interface{}) {
	o.DefaultExportValue = v
}

// GetFrozenImportFields returns the FrozenImportFields field value if set, zero value otherwise.
func (o *CreateIntegrationLevelFieldMappingDto) GetFrozenImportFields() map[string]interface{} {
	if o == nil || IsNil(o.FrozenImportFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.FrozenImportFields
}

// GetFrozenImportFieldsOk returns a tuple with the FrozenImportFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIntegrationLevelFieldMappingDto) GetFrozenImportFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.FrozenImportFields) {
		return map[string]interface{}{}, false
	}
	return o.FrozenImportFields, true
}

// HasFrozenImportFields returns a boolean if a field has been set.
func (o *CreateIntegrationLevelFieldMappingDto) HasFrozenImportFields() bool {
	if o != nil && !IsNil(o.FrozenImportFields) {
		return true
	}

	return false
}

// SetFrozenImportFields gets a reference to the given map[string]interface{} and assigns it to the FrozenImportFields field.
func (o *CreateIntegrationLevelFieldMappingDto) SetFrozenImportFields(v map[string]interface{}) {
	o.FrozenImportFields = v
}

// GetFrozenExportFields returns the FrozenExportFields field value if set, zero value otherwise.
func (o *CreateIntegrationLevelFieldMappingDto) GetFrozenExportFields() map[string]interface{} {
	if o == nil || IsNil(o.FrozenExportFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.FrozenExportFields
}

// GetFrozenExportFieldsOk returns a tuple with the FrozenExportFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIntegrationLevelFieldMappingDto) GetFrozenExportFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.FrozenExportFields) {
		return map[string]interface{}{}, false
	}
	return o.FrozenExportFields, true
}

// HasFrozenExportFields returns a boolean if a field has been set.
func (o *CreateIntegrationLevelFieldMappingDto) HasFrozenExportFields() bool {
	if o != nil && !IsNil(o.FrozenExportFields) {
		return true
	}

	return false
}

// SetFrozenExportFields gets a reference to the given map[string]interface{} and assigns it to the FrozenExportFields field.
func (o *CreateIntegrationLevelFieldMappingDto) SetFrozenExportFields(v map[string]interface{}) {
	o.FrozenExportFields = v
}

// GetFrozenUnifiedExportFields returns the FrozenUnifiedExportFields field value if set, zero value otherwise.
func (o *CreateIntegrationLevelFieldMappingDto) GetFrozenUnifiedExportFields() map[string]interface{} {
	if o == nil || IsNil(o.FrozenUnifiedExportFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.FrozenUnifiedExportFields
}

// GetFrozenUnifiedExportFieldsOk returns a tuple with the FrozenUnifiedExportFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIntegrationLevelFieldMappingDto) GetFrozenUnifiedExportFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.FrozenUnifiedExportFields) {
		return map[string]interface{}{}, false
	}
	return o.FrozenUnifiedExportFields, true
}

// HasFrozenUnifiedExportFields returns a boolean if a field has been set.
func (o *CreateIntegrationLevelFieldMappingDto) HasFrozenUnifiedExportFields() bool {
	if o != nil && !IsNil(o.FrozenUnifiedExportFields) {
		return true
	}

	return false
}

// SetFrozenUnifiedExportFields gets a reference to the given map[string]interface{} and assigns it to the FrozenUnifiedExportFields field.
func (o *CreateIntegrationLevelFieldMappingDto) SetFrozenUnifiedExportFields(v map[string]interface{}) {
	o.FrozenUnifiedExportFields = v
}

func (o CreateIntegrationLevelFieldMappingDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateIntegrationLevelFieldMappingDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["name"] = o.Name
	if !IsNil(o.ArchivedAt) {
		toSerialize["archivedAt"] = o.ArchivedAt
	}
	if !IsNil(o.DataSourceKey) {
		toSerialize["dataSourceKey"] = o.DataSourceKey
	}
	if !IsNil(o.DataSourceId) {
		toSerialize["dataSourceId"] = o.DataSourceId
	}
	if !IsNil(o.AppSchema) {
		toSerialize["appSchema"] = o.AppSchema
	}
	if !IsNil(o.Direction) {
		toSerialize["direction"] = o.Direction
	}
	if !IsNil(o.DefaultImportValue) {
		toSerialize["defaultImportValue"] = o.DefaultImportValue
	}
	if !IsNil(o.ExportValue) {
		toSerialize["exportValue"] = o.ExportValue
	}
	if !IsNil(o.ImportValue) {
		toSerialize["importValue"] = o.ImportValue
	}
	if !IsNil(o.DefaultExportValue) {
		toSerialize["defaultExportValue"] = o.DefaultExportValue
	}
	if !IsNil(o.FrozenImportFields) {
		toSerialize["frozenImportFields"] = o.FrozenImportFields
	}
	if !IsNil(o.FrozenExportFields) {
		toSerialize["frozenExportFields"] = o.FrozenExportFields
	}
	if !IsNil(o.FrozenUnifiedExportFields) {
		toSerialize["frozenUnifiedExportFields"] = o.FrozenUnifiedExportFields
	}
	return toSerialize, nil
}

func (o *CreateIntegrationLevelFieldMappingDto) UnmarshalJSON(data []byte) (err error) {
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

	varCreateIntegrationLevelFieldMappingDto := _CreateIntegrationLevelFieldMappingDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateIntegrationLevelFieldMappingDto)

	if err != nil {
		return err
	}

	*o = CreateIntegrationLevelFieldMappingDto(varCreateIntegrationLevelFieldMappingDto)

	return err
}

type NullableCreateIntegrationLevelFieldMappingDto struct {
	value *CreateIntegrationLevelFieldMappingDto
	isSet bool
}

func (v NullableCreateIntegrationLevelFieldMappingDto) Get() *CreateIntegrationLevelFieldMappingDto {
	return v.value
}

func (v *NullableCreateIntegrationLevelFieldMappingDto) Set(val *CreateIntegrationLevelFieldMappingDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateIntegrationLevelFieldMappingDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateIntegrationLevelFieldMappingDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateIntegrationLevelFieldMappingDto(val *CreateIntegrationLevelFieldMappingDto) *NullableCreateIntegrationLevelFieldMappingDto {
	return &NullableCreateIntegrationLevelFieldMappingDto{value: val, isSet: true}
}

func (v NullableCreateIntegrationLevelFieldMappingDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateIntegrationLevelFieldMappingDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

