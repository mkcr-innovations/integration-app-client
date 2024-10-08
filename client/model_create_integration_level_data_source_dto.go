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

// checks if the CreateIntegrationLevelDataSourceDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateIntegrationLevelDataSourceDto{}

// CreateIntegrationLevelDataSourceDto struct for CreateIntegrationLevelDataSourceDto
type CreateIntegrationLevelDataSourceDto struct {
	Key string `json:"key"`
	Name *string `json:"name,omitempty"`
	ArchivedAt *string `json:"archivedAt,omitempty"`
	Udm *string `json:"udm,omitempty"`
	PullUpdatesIntervalSeconds *float32 `json:"pullUpdatesIntervalSeconds,omitempty"`
	FullSyncIntervalSeconds *float32 `json:"fullSyncIntervalSeconds,omitempty"`
	DefaultPath *string `json:"defaultPath,omitempty"`
	CollectionKey *string `json:"collectionKey,omitempty"`
	CollectionParameters map[string]interface{} `json:"collectionParameters,omitempty"`
}

type _CreateIntegrationLevelDataSourceDto CreateIntegrationLevelDataSourceDto

// NewCreateIntegrationLevelDataSourceDto instantiates a new CreateIntegrationLevelDataSourceDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateIntegrationLevelDataSourceDto(key string) *CreateIntegrationLevelDataSourceDto {
	this := CreateIntegrationLevelDataSourceDto{}
	this.Key = key
	return &this
}

// NewCreateIntegrationLevelDataSourceDtoWithDefaults instantiates a new CreateIntegrationLevelDataSourceDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateIntegrationLevelDataSourceDtoWithDefaults() *CreateIntegrationLevelDataSourceDto {
	this := CreateIntegrationLevelDataSourceDto{}
	return &this
}

// GetKey returns the Key field value
func (o *CreateIntegrationLevelDataSourceDto) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *CreateIntegrationLevelDataSourceDto) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *CreateIntegrationLevelDataSourceDto) SetKey(v string) {
	o.Key = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateIntegrationLevelDataSourceDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIntegrationLevelDataSourceDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateIntegrationLevelDataSourceDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateIntegrationLevelDataSourceDto) SetName(v string) {
	o.Name = &v
}

// GetArchivedAt returns the ArchivedAt field value if set, zero value otherwise.
func (o *CreateIntegrationLevelDataSourceDto) GetArchivedAt() string {
	if o == nil || IsNil(o.ArchivedAt) {
		var ret string
		return ret
	}
	return *o.ArchivedAt
}

// GetArchivedAtOk returns a tuple with the ArchivedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIntegrationLevelDataSourceDto) GetArchivedAtOk() (*string, bool) {
	if o == nil || IsNil(o.ArchivedAt) {
		return nil, false
	}
	return o.ArchivedAt, true
}

// HasArchivedAt returns a boolean if a field has been set.
func (o *CreateIntegrationLevelDataSourceDto) HasArchivedAt() bool {
	if o != nil && !IsNil(o.ArchivedAt) {
		return true
	}

	return false
}

// SetArchivedAt gets a reference to the given string and assigns it to the ArchivedAt field.
func (o *CreateIntegrationLevelDataSourceDto) SetArchivedAt(v string) {
	o.ArchivedAt = &v
}

// GetUdm returns the Udm field value if set, zero value otherwise.
func (o *CreateIntegrationLevelDataSourceDto) GetUdm() string {
	if o == nil || IsNil(o.Udm) {
		var ret string
		return ret
	}
	return *o.Udm
}

// GetUdmOk returns a tuple with the Udm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIntegrationLevelDataSourceDto) GetUdmOk() (*string, bool) {
	if o == nil || IsNil(o.Udm) {
		return nil, false
	}
	return o.Udm, true
}

// HasUdm returns a boolean if a field has been set.
func (o *CreateIntegrationLevelDataSourceDto) HasUdm() bool {
	if o != nil && !IsNil(o.Udm) {
		return true
	}

	return false
}

// SetUdm gets a reference to the given string and assigns it to the Udm field.
func (o *CreateIntegrationLevelDataSourceDto) SetUdm(v string) {
	o.Udm = &v
}

// GetPullUpdatesIntervalSeconds returns the PullUpdatesIntervalSeconds field value if set, zero value otherwise.
func (o *CreateIntegrationLevelDataSourceDto) GetPullUpdatesIntervalSeconds() float32 {
	if o == nil || IsNil(o.PullUpdatesIntervalSeconds) {
		var ret float32
		return ret
	}
	return *o.PullUpdatesIntervalSeconds
}

// GetPullUpdatesIntervalSecondsOk returns a tuple with the PullUpdatesIntervalSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIntegrationLevelDataSourceDto) GetPullUpdatesIntervalSecondsOk() (*float32, bool) {
	if o == nil || IsNil(o.PullUpdatesIntervalSeconds) {
		return nil, false
	}
	return o.PullUpdatesIntervalSeconds, true
}

// HasPullUpdatesIntervalSeconds returns a boolean if a field has been set.
func (o *CreateIntegrationLevelDataSourceDto) HasPullUpdatesIntervalSeconds() bool {
	if o != nil && !IsNil(o.PullUpdatesIntervalSeconds) {
		return true
	}

	return false
}

// SetPullUpdatesIntervalSeconds gets a reference to the given float32 and assigns it to the PullUpdatesIntervalSeconds field.
func (o *CreateIntegrationLevelDataSourceDto) SetPullUpdatesIntervalSeconds(v float32) {
	o.PullUpdatesIntervalSeconds = &v
}

// GetFullSyncIntervalSeconds returns the FullSyncIntervalSeconds field value if set, zero value otherwise.
func (o *CreateIntegrationLevelDataSourceDto) GetFullSyncIntervalSeconds() float32 {
	if o == nil || IsNil(o.FullSyncIntervalSeconds) {
		var ret float32
		return ret
	}
	return *o.FullSyncIntervalSeconds
}

// GetFullSyncIntervalSecondsOk returns a tuple with the FullSyncIntervalSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIntegrationLevelDataSourceDto) GetFullSyncIntervalSecondsOk() (*float32, bool) {
	if o == nil || IsNil(o.FullSyncIntervalSeconds) {
		return nil, false
	}
	return o.FullSyncIntervalSeconds, true
}

// HasFullSyncIntervalSeconds returns a boolean if a field has been set.
func (o *CreateIntegrationLevelDataSourceDto) HasFullSyncIntervalSeconds() bool {
	if o != nil && !IsNil(o.FullSyncIntervalSeconds) {
		return true
	}

	return false
}

// SetFullSyncIntervalSeconds gets a reference to the given float32 and assigns it to the FullSyncIntervalSeconds field.
func (o *CreateIntegrationLevelDataSourceDto) SetFullSyncIntervalSeconds(v float32) {
	o.FullSyncIntervalSeconds = &v
}

// GetDefaultPath returns the DefaultPath field value if set, zero value otherwise.
func (o *CreateIntegrationLevelDataSourceDto) GetDefaultPath() string {
	if o == nil || IsNil(o.DefaultPath) {
		var ret string
		return ret
	}
	return *o.DefaultPath
}

// GetDefaultPathOk returns a tuple with the DefaultPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIntegrationLevelDataSourceDto) GetDefaultPathOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultPath) {
		return nil, false
	}
	return o.DefaultPath, true
}

// HasDefaultPath returns a boolean if a field has been set.
func (o *CreateIntegrationLevelDataSourceDto) HasDefaultPath() bool {
	if o != nil && !IsNil(o.DefaultPath) {
		return true
	}

	return false
}

// SetDefaultPath gets a reference to the given string and assigns it to the DefaultPath field.
func (o *CreateIntegrationLevelDataSourceDto) SetDefaultPath(v string) {
	o.DefaultPath = &v
}

// GetCollectionKey returns the CollectionKey field value if set, zero value otherwise.
func (o *CreateIntegrationLevelDataSourceDto) GetCollectionKey() string {
	if o == nil || IsNil(o.CollectionKey) {
		var ret string
		return ret
	}
	return *o.CollectionKey
}

// GetCollectionKeyOk returns a tuple with the CollectionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIntegrationLevelDataSourceDto) GetCollectionKeyOk() (*string, bool) {
	if o == nil || IsNil(o.CollectionKey) {
		return nil, false
	}
	return o.CollectionKey, true
}

// HasCollectionKey returns a boolean if a field has been set.
func (o *CreateIntegrationLevelDataSourceDto) HasCollectionKey() bool {
	if o != nil && !IsNil(o.CollectionKey) {
		return true
	}

	return false
}

// SetCollectionKey gets a reference to the given string and assigns it to the CollectionKey field.
func (o *CreateIntegrationLevelDataSourceDto) SetCollectionKey(v string) {
	o.CollectionKey = &v
}

// GetCollectionParameters returns the CollectionParameters field value if set, zero value otherwise.
func (o *CreateIntegrationLevelDataSourceDto) GetCollectionParameters() map[string]interface{} {
	if o == nil || IsNil(o.CollectionParameters) {
		var ret map[string]interface{}
		return ret
	}
	return o.CollectionParameters
}

// GetCollectionParametersOk returns a tuple with the CollectionParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIntegrationLevelDataSourceDto) GetCollectionParametersOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CollectionParameters) {
		return map[string]interface{}{}, false
	}
	return o.CollectionParameters, true
}

// HasCollectionParameters returns a boolean if a field has been set.
func (o *CreateIntegrationLevelDataSourceDto) HasCollectionParameters() bool {
	if o != nil && !IsNil(o.CollectionParameters) {
		return true
	}

	return false
}

// SetCollectionParameters gets a reference to the given map[string]interface{} and assigns it to the CollectionParameters field.
func (o *CreateIntegrationLevelDataSourceDto) SetCollectionParameters(v map[string]interface{}) {
	o.CollectionParameters = v
}

func (o CreateIntegrationLevelDataSourceDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateIntegrationLevelDataSourceDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ArchivedAt) {
		toSerialize["archivedAt"] = o.ArchivedAt
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
	if !IsNil(o.DefaultPath) {
		toSerialize["defaultPath"] = o.DefaultPath
	}
	if !IsNil(o.CollectionKey) {
		toSerialize["collectionKey"] = o.CollectionKey
	}
	if !IsNil(o.CollectionParameters) {
		toSerialize["collectionParameters"] = o.CollectionParameters
	}
	return toSerialize, nil
}

func (o *CreateIntegrationLevelDataSourceDto) UnmarshalJSON(data []byte) (err error) {
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

	varCreateIntegrationLevelDataSourceDto := _CreateIntegrationLevelDataSourceDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateIntegrationLevelDataSourceDto)

	if err != nil {
		return err
	}

	*o = CreateIntegrationLevelDataSourceDto(varCreateIntegrationLevelDataSourceDto)

	return err
}

type NullableCreateIntegrationLevelDataSourceDto struct {
	value *CreateIntegrationLevelDataSourceDto
	isSet bool
}

func (v NullableCreateIntegrationLevelDataSourceDto) Get() *CreateIntegrationLevelDataSourceDto {
	return v.value
}

func (v *NullableCreateIntegrationLevelDataSourceDto) Set(val *CreateIntegrationLevelDataSourceDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateIntegrationLevelDataSourceDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateIntegrationLevelDataSourceDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateIntegrationLevelDataSourceDto(val *CreateIntegrationLevelDataSourceDto) *NullableCreateIntegrationLevelDataSourceDto {
	return &NullableCreateIntegrationLevelDataSourceDto{value: val, isSet: true}
}

func (v NullableCreateIntegrationLevelDataSourceDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateIntegrationLevelDataSourceDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


