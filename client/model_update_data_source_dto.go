/*
Integration Engine API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the UpdateDataSourceDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDataSourceDto{}

// UpdateDataSourceDto struct for UpdateDataSourceDto
type UpdateDataSourceDto struct {
	Udm *string `json:"udm,omitempty"`
	PullUpdatesIntervalSeconds *float32 `json:"pullUpdatesIntervalSeconds,omitempty"`
	FullSyncIntervalSeconds *float32 `json:"fullSyncIntervalSeconds,omitempty"`
	DefaultPath *string `json:"defaultPath,omitempty"`
	CollectionKey map[string]interface{} `json:"collectionKey,omitempty"`
	CollectionParameters map[string]interface{} `json:"collectionParameters,omitempty"`
}

// NewUpdateDataSourceDto instantiates a new UpdateDataSourceDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDataSourceDto() *UpdateDataSourceDto {
	this := UpdateDataSourceDto{}
	return &this
}

// NewUpdateDataSourceDtoWithDefaults instantiates a new UpdateDataSourceDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDataSourceDtoWithDefaults() *UpdateDataSourceDto {
	this := UpdateDataSourceDto{}
	return &this
}

// GetUdm returns the Udm field value if set, zero value otherwise.
func (o *UpdateDataSourceDto) GetUdm() string {
	if o == nil || IsNil(o.Udm) {
		var ret string
		return ret
	}
	return *o.Udm
}

// GetUdmOk returns a tuple with the Udm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDataSourceDto) GetUdmOk() (*string, bool) {
	if o == nil || IsNil(o.Udm) {
		return nil, false
	}
	return o.Udm, true
}

// HasUdm returns a boolean if a field has been set.
func (o *UpdateDataSourceDto) HasUdm() bool {
	if o != nil && !IsNil(o.Udm) {
		return true
	}

	return false
}

// SetUdm gets a reference to the given string and assigns it to the Udm field.
func (o *UpdateDataSourceDto) SetUdm(v string) {
	o.Udm = &v
}

// GetPullUpdatesIntervalSeconds returns the PullUpdatesIntervalSeconds field value if set, zero value otherwise.
func (o *UpdateDataSourceDto) GetPullUpdatesIntervalSeconds() float32 {
	if o == nil || IsNil(o.PullUpdatesIntervalSeconds) {
		var ret float32
		return ret
	}
	return *o.PullUpdatesIntervalSeconds
}

// GetPullUpdatesIntervalSecondsOk returns a tuple with the PullUpdatesIntervalSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDataSourceDto) GetPullUpdatesIntervalSecondsOk() (*float32, bool) {
	if o == nil || IsNil(o.PullUpdatesIntervalSeconds) {
		return nil, false
	}
	return o.PullUpdatesIntervalSeconds, true
}

// HasPullUpdatesIntervalSeconds returns a boolean if a field has been set.
func (o *UpdateDataSourceDto) HasPullUpdatesIntervalSeconds() bool {
	if o != nil && !IsNil(o.PullUpdatesIntervalSeconds) {
		return true
	}

	return false
}

// SetPullUpdatesIntervalSeconds gets a reference to the given float32 and assigns it to the PullUpdatesIntervalSeconds field.
func (o *UpdateDataSourceDto) SetPullUpdatesIntervalSeconds(v float32) {
	o.PullUpdatesIntervalSeconds = &v
}

// GetFullSyncIntervalSeconds returns the FullSyncIntervalSeconds field value if set, zero value otherwise.
func (o *UpdateDataSourceDto) GetFullSyncIntervalSeconds() float32 {
	if o == nil || IsNil(o.FullSyncIntervalSeconds) {
		var ret float32
		return ret
	}
	return *o.FullSyncIntervalSeconds
}

// GetFullSyncIntervalSecondsOk returns a tuple with the FullSyncIntervalSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDataSourceDto) GetFullSyncIntervalSecondsOk() (*float32, bool) {
	if o == nil || IsNil(o.FullSyncIntervalSeconds) {
		return nil, false
	}
	return o.FullSyncIntervalSeconds, true
}

// HasFullSyncIntervalSeconds returns a boolean if a field has been set.
func (o *UpdateDataSourceDto) HasFullSyncIntervalSeconds() bool {
	if o != nil && !IsNil(o.FullSyncIntervalSeconds) {
		return true
	}

	return false
}

// SetFullSyncIntervalSeconds gets a reference to the given float32 and assigns it to the FullSyncIntervalSeconds field.
func (o *UpdateDataSourceDto) SetFullSyncIntervalSeconds(v float32) {
	o.FullSyncIntervalSeconds = &v
}

// GetDefaultPath returns the DefaultPath field value if set, zero value otherwise.
func (o *UpdateDataSourceDto) GetDefaultPath() string {
	if o == nil || IsNil(o.DefaultPath) {
		var ret string
		return ret
	}
	return *o.DefaultPath
}

// GetDefaultPathOk returns a tuple with the DefaultPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDataSourceDto) GetDefaultPathOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultPath) {
		return nil, false
	}
	return o.DefaultPath, true
}

// HasDefaultPath returns a boolean if a field has been set.
func (o *UpdateDataSourceDto) HasDefaultPath() bool {
	if o != nil && !IsNil(o.DefaultPath) {
		return true
	}

	return false
}

// SetDefaultPath gets a reference to the given string and assigns it to the DefaultPath field.
func (o *UpdateDataSourceDto) SetDefaultPath(v string) {
	o.DefaultPath = &v
}

// GetCollectionKey returns the CollectionKey field value if set, zero value otherwise.
func (o *UpdateDataSourceDto) GetCollectionKey() map[string]interface{} {
	if o == nil || IsNil(o.CollectionKey) {
		var ret map[string]interface{}
		return ret
	}
	return o.CollectionKey
}

// GetCollectionKeyOk returns a tuple with the CollectionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDataSourceDto) GetCollectionKeyOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CollectionKey) {
		return map[string]interface{}{}, false
	}
	return o.CollectionKey, true
}

// HasCollectionKey returns a boolean if a field has been set.
func (o *UpdateDataSourceDto) HasCollectionKey() bool {
	if o != nil && !IsNil(o.CollectionKey) {
		return true
	}

	return false
}

// SetCollectionKey gets a reference to the given map[string]interface{} and assigns it to the CollectionKey field.
func (o *UpdateDataSourceDto) SetCollectionKey(v map[string]interface{}) {
	o.CollectionKey = v
}

// GetCollectionParameters returns the CollectionParameters field value if set, zero value otherwise.
func (o *UpdateDataSourceDto) GetCollectionParameters() map[string]interface{} {
	if o == nil || IsNil(o.CollectionParameters) {
		var ret map[string]interface{}
		return ret
	}
	return o.CollectionParameters
}

// GetCollectionParametersOk returns a tuple with the CollectionParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDataSourceDto) GetCollectionParametersOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CollectionParameters) {
		return map[string]interface{}{}, false
	}
	return o.CollectionParameters, true
}

// HasCollectionParameters returns a boolean if a field has been set.
func (o *UpdateDataSourceDto) HasCollectionParameters() bool {
	if o != nil && !IsNil(o.CollectionParameters) {
		return true
	}

	return false
}

// SetCollectionParameters gets a reference to the given map[string]interface{} and assigns it to the CollectionParameters field.
func (o *UpdateDataSourceDto) SetCollectionParameters(v map[string]interface{}) {
	o.CollectionParameters = v
}

func (o UpdateDataSourceDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDataSourceDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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

type NullableUpdateDataSourceDto struct {
	value *UpdateDataSourceDto
	isSet bool
}

func (v NullableUpdateDataSourceDto) Get() *UpdateDataSourceDto {
	return v.value
}

func (v *NullableUpdateDataSourceDto) Set(val *UpdateDataSourceDto) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDataSourceDto) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDataSourceDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDataSourceDto(val *UpdateDataSourceDto) *NullableUpdateDataSourceDto {
	return &NullableUpdateDataSourceDto{value: val, isSet: true}
}

func (v NullableUpdateDataSourceDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDataSourceDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


