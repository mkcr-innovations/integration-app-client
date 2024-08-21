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

// checks if the CreateFieldMappingDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateFieldMappingDto{}

// CreateFieldMappingDto struct for CreateFieldMappingDto
type CreateFieldMappingDto struct {
	Key string `json:"key"`
	Name *string `json:"name,omitempty"`
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
	IntegrationId *string `json:"integrationId,omitempty"`
}

type _CreateFieldMappingDto CreateFieldMappingDto

// NewCreateFieldMappingDto instantiates a new CreateFieldMappingDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFieldMappingDto(key string) *CreateFieldMappingDto {
	this := CreateFieldMappingDto{}
	this.Key = key
	return &this
}

// NewCreateFieldMappingDtoWithDefaults instantiates a new CreateFieldMappingDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFieldMappingDtoWithDefaults() *CreateFieldMappingDto {
	this := CreateFieldMappingDto{}
	return &this
}

// GetKey returns the Key field value
func (o *CreateFieldMappingDto) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *CreateFieldMappingDto) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *CreateFieldMappingDto) SetKey(v string) {
	o.Key = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateFieldMappingDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFieldMappingDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateFieldMappingDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateFieldMappingDto) SetName(v string) {
	o.Name = &v
}

// GetArchivedAt returns the ArchivedAt field value if set, zero value otherwise.
func (o *CreateFieldMappingDto) GetArchivedAt() string {
	if o == nil || IsNil(o.ArchivedAt) {
		var ret string
		return ret
	}
	return *o.ArchivedAt
}

// GetArchivedAtOk returns a tuple with the ArchivedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFieldMappingDto) GetArchivedAtOk() (*string, bool) {
	if o == nil || IsNil(o.ArchivedAt) {
		return nil, false
	}
	return o.ArchivedAt, true
}

// HasArchivedAt returns a boolean if a field has been set.
func (o *CreateFieldMappingDto) HasArchivedAt() bool {
	if o != nil && !IsNil(o.ArchivedAt) {
		return true
	}

	return false
}

// SetArchivedAt gets a reference to the given string and assigns it to the ArchivedAt field.
func (o *CreateFieldMappingDto) SetArchivedAt(v string) {
	o.ArchivedAt = &v
}

// GetDataSourceKey returns the DataSourceKey field value if set, zero value otherwise.
func (o *CreateFieldMappingDto) GetDataSourceKey() string {
	if o == nil || IsNil(o.DataSourceKey) {
		var ret string
		return ret
	}
	return *o.DataSourceKey
}

// GetDataSourceKeyOk returns a tuple with the DataSourceKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFieldMappingDto) GetDataSourceKeyOk() (*string, bool) {
	if o == nil || IsNil(o.DataSourceKey) {
		return nil, false
	}
	return o.DataSourceKey, true
}

// HasDataSourceKey returns a boolean if a field has been set.
func (o *CreateFieldMappingDto) HasDataSourceKey() bool {
	if o != nil && !IsNil(o.DataSourceKey) {
		return true
	}

	return false
}

// SetDataSourceKey gets a reference to the given string and assigns it to the DataSourceKey field.
func (o *CreateFieldMappingDto) SetDataSourceKey(v string) {
	o.DataSourceKey = &v
}

// GetDataSourceId returns the DataSourceId field value if set, zero value otherwise.
func (o *CreateFieldMappingDto) GetDataSourceId() string {
	if o == nil || IsNil(o.DataSourceId) {
		var ret string
		return ret
	}
	return *o.DataSourceId
}

// GetDataSourceIdOk returns a tuple with the DataSourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFieldMappingDto) GetDataSourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.DataSourceId) {
		return nil, false
	}
	return o.DataSourceId, true
}

// HasDataSourceId returns a boolean if a field has been set.
func (o *CreateFieldMappingDto) HasDataSourceId() bool {
	if o != nil && !IsNil(o.DataSourceId) {
		return true
	}

	return false
}

// SetDataSourceId gets a reference to the given string and assigns it to the DataSourceId field.
func (o *CreateFieldMappingDto) SetDataSourceId(v string) {
	o.DataSourceId = &v
}

// GetAppSchema returns the AppSchema field value if set, zero value otherwise.
func (o *CreateFieldMappingDto) GetAppSchema() map[string]interface{} {
	if o == nil || IsNil(o.AppSchema) {
		var ret map[string]interface{}
		return ret
	}
	return o.AppSchema
}

// GetAppSchemaOk returns a tuple with the AppSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFieldMappingDto) GetAppSchemaOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AppSchema) {
		return map[string]interface{}{}, false
	}
	return o.AppSchema, true
}

// HasAppSchema returns a boolean if a field has been set.
func (o *CreateFieldMappingDto) HasAppSchema() bool {
	if o != nil && !IsNil(o.AppSchema) {
		return true
	}

	return false
}

// SetAppSchema gets a reference to the given map[string]interface{} and assigns it to the AppSchema field.
func (o *CreateFieldMappingDto) SetAppSchema(v map[string]interface{}) {
	o.AppSchema = v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *CreateFieldMappingDto) GetDirection() string {
	if o == nil || IsNil(o.Direction) {
		var ret string
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFieldMappingDto) GetDirectionOk() (*string, bool) {
	if o == nil || IsNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *CreateFieldMappingDto) HasDirection() bool {
	if o != nil && !IsNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given string and assigns it to the Direction field.
func (o *CreateFieldMappingDto) SetDirection(v string) {
	o.Direction = &v
}

// GetDefaultImportValue returns the DefaultImportValue field value if set, zero value otherwise.
func (o *CreateFieldMappingDto) GetDefaultImportValue() map[string]interface{} {
	if o == nil || IsNil(o.DefaultImportValue) {
		var ret map[string]interface{}
		return ret
	}
	return o.DefaultImportValue
}

// GetDefaultImportValueOk returns a tuple with the DefaultImportValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFieldMappingDto) GetDefaultImportValueOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.DefaultImportValue) {
		return map[string]interface{}{}, false
	}
	return o.DefaultImportValue, true
}

// HasDefaultImportValue returns a boolean if a field has been set.
func (o *CreateFieldMappingDto) HasDefaultImportValue() bool {
	if o != nil && !IsNil(o.DefaultImportValue) {
		return true
	}

	return false
}

// SetDefaultImportValue gets a reference to the given map[string]interface{} and assigns it to the DefaultImportValue field.
func (o *CreateFieldMappingDto) SetDefaultImportValue(v map[string]interface{}) {
	o.DefaultImportValue = v
}

// GetExportValue returns the ExportValue field value if set, zero value otherwise.
func (o *CreateFieldMappingDto) GetExportValue() map[string]interface{} {
	if o == nil || IsNil(o.ExportValue) {
		var ret map[string]interface{}
		return ret
	}
	return o.ExportValue
}

// GetExportValueOk returns a tuple with the ExportValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFieldMappingDto) GetExportValueOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ExportValue) {
		return map[string]interface{}{}, false
	}
	return o.ExportValue, true
}

// HasExportValue returns a boolean if a field has been set.
func (o *CreateFieldMappingDto) HasExportValue() bool {
	if o != nil && !IsNil(o.ExportValue) {
		return true
	}

	return false
}

// SetExportValue gets a reference to the given map[string]interface{} and assigns it to the ExportValue field.
func (o *CreateFieldMappingDto) SetExportValue(v map[string]interface{}) {
	o.ExportValue = v
}

// GetImportValue returns the ImportValue field value if set, zero value otherwise.
func (o *CreateFieldMappingDto) GetImportValue() map[string]interface{} {
	if o == nil || IsNil(o.ImportValue) {
		var ret map[string]interface{}
		return ret
	}
	return o.ImportValue
}

// GetImportValueOk returns a tuple with the ImportValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFieldMappingDto) GetImportValueOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ImportValue) {
		return map[string]interface{}{}, false
	}
	return o.ImportValue, true
}

// HasImportValue returns a boolean if a field has been set.
func (o *CreateFieldMappingDto) HasImportValue() bool {
	if o != nil && !IsNil(o.ImportValue) {
		return true
	}

	return false
}

// SetImportValue gets a reference to the given map[string]interface{} and assigns it to the ImportValue field.
func (o *CreateFieldMappingDto) SetImportValue(v map[string]interface{}) {
	o.ImportValue = v
}

// GetDefaultExportValue returns the DefaultExportValue field value if set, zero value otherwise.
func (o *CreateFieldMappingDto) GetDefaultExportValue() map[string]interface{} {
	if o == nil || IsNil(o.DefaultExportValue) {
		var ret map[string]interface{}
		return ret
	}
	return o.DefaultExportValue
}

// GetDefaultExportValueOk returns a tuple with the DefaultExportValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFieldMappingDto) GetDefaultExportValueOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.DefaultExportValue) {
		return map[string]interface{}{}, false
	}
	return o.DefaultExportValue, true
}

// HasDefaultExportValue returns a boolean if a field has been set.
func (o *CreateFieldMappingDto) HasDefaultExportValue() bool {
	if o != nil && !IsNil(o.DefaultExportValue) {
		return true
	}

	return false
}

// SetDefaultExportValue gets a reference to the given map[string]interface{} and assigns it to the DefaultExportValue field.
func (o *CreateFieldMappingDto) SetDefaultExportValue(v map[string]interface{}) {
	o.DefaultExportValue = v
}

// GetFrozenImportFields returns the FrozenImportFields field value if set, zero value otherwise.
func (o *CreateFieldMappingDto) GetFrozenImportFields() map[string]interface{} {
	if o == nil || IsNil(o.FrozenImportFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.FrozenImportFields
}

// GetFrozenImportFieldsOk returns a tuple with the FrozenImportFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFieldMappingDto) GetFrozenImportFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.FrozenImportFields) {
		return map[string]interface{}{}, false
	}
	return o.FrozenImportFields, true
}

// HasFrozenImportFields returns a boolean if a field has been set.
func (o *CreateFieldMappingDto) HasFrozenImportFields() bool {
	if o != nil && !IsNil(o.FrozenImportFields) {
		return true
	}

	return false
}

// SetFrozenImportFields gets a reference to the given map[string]interface{} and assigns it to the FrozenImportFields field.
func (o *CreateFieldMappingDto) SetFrozenImportFields(v map[string]interface{}) {
	o.FrozenImportFields = v
}

// GetFrozenExportFields returns the FrozenExportFields field value if set, zero value otherwise.
func (o *CreateFieldMappingDto) GetFrozenExportFields() map[string]interface{} {
	if o == nil || IsNil(o.FrozenExportFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.FrozenExportFields
}

// GetFrozenExportFieldsOk returns a tuple with the FrozenExportFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFieldMappingDto) GetFrozenExportFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.FrozenExportFields) {
		return map[string]interface{}{}, false
	}
	return o.FrozenExportFields, true
}

// HasFrozenExportFields returns a boolean if a field has been set.
func (o *CreateFieldMappingDto) HasFrozenExportFields() bool {
	if o != nil && !IsNil(o.FrozenExportFields) {
		return true
	}

	return false
}

// SetFrozenExportFields gets a reference to the given map[string]interface{} and assigns it to the FrozenExportFields field.
func (o *CreateFieldMappingDto) SetFrozenExportFields(v map[string]interface{}) {
	o.FrozenExportFields = v
}

// GetFrozenUnifiedExportFields returns the FrozenUnifiedExportFields field value if set, zero value otherwise.
func (o *CreateFieldMappingDto) GetFrozenUnifiedExportFields() map[string]interface{} {
	if o == nil || IsNil(o.FrozenUnifiedExportFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.FrozenUnifiedExportFields
}

// GetFrozenUnifiedExportFieldsOk returns a tuple with the FrozenUnifiedExportFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFieldMappingDto) GetFrozenUnifiedExportFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.FrozenUnifiedExportFields) {
		return map[string]interface{}{}, false
	}
	return o.FrozenUnifiedExportFields, true
}

// HasFrozenUnifiedExportFields returns a boolean if a field has been set.
func (o *CreateFieldMappingDto) HasFrozenUnifiedExportFields() bool {
	if o != nil && !IsNil(o.FrozenUnifiedExportFields) {
		return true
	}

	return false
}

// SetFrozenUnifiedExportFields gets a reference to the given map[string]interface{} and assigns it to the FrozenUnifiedExportFields field.
func (o *CreateFieldMappingDto) SetFrozenUnifiedExportFields(v map[string]interface{}) {
	o.FrozenUnifiedExportFields = v
}

// GetIntegrationId returns the IntegrationId field value if set, zero value otherwise.
func (o *CreateFieldMappingDto) GetIntegrationId() string {
	if o == nil || IsNil(o.IntegrationId) {
		var ret string
		return ret
	}
	return *o.IntegrationId
}

// GetIntegrationIdOk returns a tuple with the IntegrationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFieldMappingDto) GetIntegrationIdOk() (*string, bool) {
	if o == nil || IsNil(o.IntegrationId) {
		return nil, false
	}
	return o.IntegrationId, true
}

// HasIntegrationId returns a boolean if a field has been set.
func (o *CreateFieldMappingDto) HasIntegrationId() bool {
	if o != nil && !IsNil(o.IntegrationId) {
		return true
	}

	return false
}

// SetIntegrationId gets a reference to the given string and assigns it to the IntegrationId field.
func (o *CreateFieldMappingDto) SetIntegrationId(v string) {
	o.IntegrationId = &v
}

func (o CreateFieldMappingDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateFieldMappingDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
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
	if !IsNil(o.IntegrationId) {
		toSerialize["integrationId"] = o.IntegrationId
	}
	return toSerialize, nil
}

func (o *CreateFieldMappingDto) UnmarshalJSON(data []byte) (err error) {
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

	varCreateFieldMappingDto := _CreateFieldMappingDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateFieldMappingDto)

	if err != nil {
		return err
	}

	*o = CreateFieldMappingDto(varCreateFieldMappingDto)

	return err
}

type NullableCreateFieldMappingDto struct {
	value *CreateFieldMappingDto
	isSet bool
}

func (v NullableCreateFieldMappingDto) Get() *CreateFieldMappingDto {
	return v.value
}

func (v *NullableCreateFieldMappingDto) Set(val *CreateFieldMappingDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFieldMappingDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFieldMappingDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFieldMappingDto(val *CreateFieldMappingDto) *NullableCreateFieldMappingDto {
	return &NullableCreateFieldMappingDto{value: val, isSet: true}
}

func (v NullableCreateFieldMappingDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFieldMappingDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


