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

// checks if the DataSourceExportDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataSourceExportDto{}

// DataSourceExportDto struct for DataSourceExportDto
type DataSourceExportDto struct {
	Key *string `json:"key,omitempty"`
	Name string `json:"name"`
	IntegrationKeys []string `json:"integrationKeys,omitempty"`
	CollectionKey *string `json:"collectionKey,omitempty"`
	CollectionParameters map[string]interface{} `json:"collectionParameters,omitempty"`
	Udm *string `json:"udm,omitempty"`
	PullUpdatesIntervalSeconds *float32 `json:"pullUpdatesIntervalSeconds,omitempty"`
	FullSyncIntervalSeconds *float32 `json:"fullSyncIntervalSeconds,omitempty"`
}

type _DataSourceExportDto DataSourceExportDto

// NewDataSourceExportDto instantiates a new DataSourceExportDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataSourceExportDto(name string) *DataSourceExportDto {
	this := DataSourceExportDto{}
	this.Name = name
	return &this
}

// NewDataSourceExportDtoWithDefaults instantiates a new DataSourceExportDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataSourceExportDtoWithDefaults() *DataSourceExportDto {
	this := DataSourceExportDto{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *DataSourceExportDto) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourceExportDto) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *DataSourceExportDto) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *DataSourceExportDto) SetKey(v string) {
	o.Key = &v
}

// GetName returns the Name field value
func (o *DataSourceExportDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DataSourceExportDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DataSourceExportDto) SetName(v string) {
	o.Name = v
}

// GetIntegrationKeys returns the IntegrationKeys field value if set, zero value otherwise.
func (o *DataSourceExportDto) GetIntegrationKeys() []string {
	if o == nil || IsNil(o.IntegrationKeys) {
		var ret []string
		return ret
	}
	return o.IntegrationKeys
}

// GetIntegrationKeysOk returns a tuple with the IntegrationKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourceExportDto) GetIntegrationKeysOk() ([]string, bool) {
	if o == nil || IsNil(o.IntegrationKeys) {
		return nil, false
	}
	return o.IntegrationKeys, true
}

// HasIntegrationKeys returns a boolean if a field has been set.
func (o *DataSourceExportDto) HasIntegrationKeys() bool {
	if o != nil && !IsNil(o.IntegrationKeys) {
		return true
	}

	return false
}

// SetIntegrationKeys gets a reference to the given []string and assigns it to the IntegrationKeys field.
func (o *DataSourceExportDto) SetIntegrationKeys(v []string) {
	o.IntegrationKeys = v
}

// GetCollectionKey returns the CollectionKey field value if set, zero value otherwise.
func (o *DataSourceExportDto) GetCollectionKey() string {
	if o == nil || IsNil(o.CollectionKey) {
		var ret string
		return ret
	}
	return *o.CollectionKey
}

// GetCollectionKeyOk returns a tuple with the CollectionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourceExportDto) GetCollectionKeyOk() (*string, bool) {
	if o == nil || IsNil(o.CollectionKey) {
		return nil, false
	}
	return o.CollectionKey, true
}

// HasCollectionKey returns a boolean if a field has been set.
func (o *DataSourceExportDto) HasCollectionKey() bool {
	if o != nil && !IsNil(o.CollectionKey) {
		return true
	}

	return false
}

// SetCollectionKey gets a reference to the given string and assigns it to the CollectionKey field.
func (o *DataSourceExportDto) SetCollectionKey(v string) {
	o.CollectionKey = &v
}

// GetCollectionParameters returns the CollectionParameters field value if set, zero value otherwise.
func (o *DataSourceExportDto) GetCollectionParameters() map[string]interface{} {
	if o == nil || IsNil(o.CollectionParameters) {
		var ret map[string]interface{}
		return ret
	}
	return o.CollectionParameters
}

// GetCollectionParametersOk returns a tuple with the CollectionParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourceExportDto) GetCollectionParametersOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CollectionParameters) {
		return map[string]interface{}{}, false
	}
	return o.CollectionParameters, true
}

// HasCollectionParameters returns a boolean if a field has been set.
func (o *DataSourceExportDto) HasCollectionParameters() bool {
	if o != nil && !IsNil(o.CollectionParameters) {
		return true
	}

	return false
}

// SetCollectionParameters gets a reference to the given map[string]interface{} and assigns it to the CollectionParameters field.
func (o *DataSourceExportDto) SetCollectionParameters(v map[string]interface{}) {
	o.CollectionParameters = v
}

// GetUdm returns the Udm field value if set, zero value otherwise.
func (o *DataSourceExportDto) GetUdm() string {
	if o == nil || IsNil(o.Udm) {
		var ret string
		return ret
	}
	return *o.Udm
}

// GetUdmOk returns a tuple with the Udm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourceExportDto) GetUdmOk() (*string, bool) {
	if o == nil || IsNil(o.Udm) {
		return nil, false
	}
	return o.Udm, true
}

// HasUdm returns a boolean if a field has been set.
func (o *DataSourceExportDto) HasUdm() bool {
	if o != nil && !IsNil(o.Udm) {
		return true
	}

	return false
}

// SetUdm gets a reference to the given string and assigns it to the Udm field.
func (o *DataSourceExportDto) SetUdm(v string) {
	o.Udm = &v
}

// GetPullUpdatesIntervalSeconds returns the PullUpdatesIntervalSeconds field value if set, zero value otherwise.
func (o *DataSourceExportDto) GetPullUpdatesIntervalSeconds() float32 {
	if o == nil || IsNil(o.PullUpdatesIntervalSeconds) {
		var ret float32
		return ret
	}
	return *o.PullUpdatesIntervalSeconds
}

// GetPullUpdatesIntervalSecondsOk returns a tuple with the PullUpdatesIntervalSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourceExportDto) GetPullUpdatesIntervalSecondsOk() (*float32, bool) {
	if o == nil || IsNil(o.PullUpdatesIntervalSeconds) {
		return nil, false
	}
	return o.PullUpdatesIntervalSeconds, true
}

// HasPullUpdatesIntervalSeconds returns a boolean if a field has been set.
func (o *DataSourceExportDto) HasPullUpdatesIntervalSeconds() bool {
	if o != nil && !IsNil(o.PullUpdatesIntervalSeconds) {
		return true
	}

	return false
}

// SetPullUpdatesIntervalSeconds gets a reference to the given float32 and assigns it to the PullUpdatesIntervalSeconds field.
func (o *DataSourceExportDto) SetPullUpdatesIntervalSeconds(v float32) {
	o.PullUpdatesIntervalSeconds = &v
}

// GetFullSyncIntervalSeconds returns the FullSyncIntervalSeconds field value if set, zero value otherwise.
func (o *DataSourceExportDto) GetFullSyncIntervalSeconds() float32 {
	if o == nil || IsNil(o.FullSyncIntervalSeconds) {
		var ret float32
		return ret
	}
	return *o.FullSyncIntervalSeconds
}

// GetFullSyncIntervalSecondsOk returns a tuple with the FullSyncIntervalSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourceExportDto) GetFullSyncIntervalSecondsOk() (*float32, bool) {
	if o == nil || IsNil(o.FullSyncIntervalSeconds) {
		return nil, false
	}
	return o.FullSyncIntervalSeconds, true
}

// HasFullSyncIntervalSeconds returns a boolean if a field has been set.
func (o *DataSourceExportDto) HasFullSyncIntervalSeconds() bool {
	if o != nil && !IsNil(o.FullSyncIntervalSeconds) {
		return true
	}

	return false
}

// SetFullSyncIntervalSeconds gets a reference to the given float32 and assigns it to the FullSyncIntervalSeconds field.
func (o *DataSourceExportDto) SetFullSyncIntervalSeconds(v float32) {
	o.FullSyncIntervalSeconds = &v
}

func (o DataSourceExportDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataSourceExportDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.IntegrationKeys) {
		toSerialize["integrationKeys"] = o.IntegrationKeys
	}
	if !IsNil(o.CollectionKey) {
		toSerialize["collectionKey"] = o.CollectionKey
	}
	if !IsNil(o.CollectionParameters) {
		toSerialize["collectionParameters"] = o.CollectionParameters
	}
	if !IsNil(o.Udm) {
		toSerialize["udm"] = o.Udm
	}
	if !IsNil(o.PullUpdatesIntervalSeconds) {
		toSerialize["pullUpdatesIntervalSeconds"] = o.PullUpdatesIntervalSeconds
	}
	if !IsNil(o.FullSyncIntervalSeconds) {
		toSerialize["fullSyncIntervalSeconds"] = o.FullSyncIntervalSeconds
	}
	return toSerialize, nil
}

func (o *DataSourceExportDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varDataSourceExportDto := _DataSourceExportDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDataSourceExportDto)

	if err != nil {
		return err
	}

	*o = DataSourceExportDto(varDataSourceExportDto)

	return err
}

type NullableDataSourceExportDto struct {
	value *DataSourceExportDto
	isSet bool
}

func (v NullableDataSourceExportDto) Get() *DataSourceExportDto {
	return v.value
}

func (v *NullableDataSourceExportDto) Set(val *DataSourceExportDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDataSourceExportDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDataSourceExportDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataSourceExportDto(val *DataSourceExportDto) *NullableDataSourceExportDto {
	return &NullableDataSourceExportDto{value: val, isSet: true}
}

func (v NullableDataSourceExportDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataSourceExportDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

