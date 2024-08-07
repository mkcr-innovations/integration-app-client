/*
Integration.app API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the UpdateDataSourceInstanceDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDataSourceInstanceDto{}

// UpdateDataSourceInstanceDto struct for UpdateDataSourceInstanceDto
type UpdateDataSourceInstanceDto struct {
	// Deprecated
	Path *string `json:"path,omitempty"`
	CollectionKey *string `json:"collectionKey,omitempty"`
	CollectionParameters map[string]interface{} `json:"collectionParameters,omitempty"`
	PullUpdatesIntervalSeconds *float32 `json:"pullUpdatesIntervalSeconds,omitempty"`
	FullSyncIntervalSeconds *float32 `json:"fullSyncIntervalSeconds,omitempty"`
}

// NewUpdateDataSourceInstanceDto instantiates a new UpdateDataSourceInstanceDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDataSourceInstanceDto() *UpdateDataSourceInstanceDto {
	this := UpdateDataSourceInstanceDto{}
	return &this
}

// NewUpdateDataSourceInstanceDtoWithDefaults instantiates a new UpdateDataSourceInstanceDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDataSourceInstanceDtoWithDefaults() *UpdateDataSourceInstanceDto {
	this := UpdateDataSourceInstanceDto{}
	return &this
}

// GetPath returns the Path field value if set, zero value otherwise.
// Deprecated
func (o *UpdateDataSourceInstanceDto) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *UpdateDataSourceInstanceDto) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *UpdateDataSourceInstanceDto) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
// Deprecated
func (o *UpdateDataSourceInstanceDto) SetPath(v string) {
	o.Path = &v
}

// GetCollectionKey returns the CollectionKey field value if set, zero value otherwise.
func (o *UpdateDataSourceInstanceDto) GetCollectionKey() string {
	if o == nil || IsNil(o.CollectionKey) {
		var ret string
		return ret
	}
	return *o.CollectionKey
}

// GetCollectionKeyOk returns a tuple with the CollectionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDataSourceInstanceDto) GetCollectionKeyOk() (*string, bool) {
	if o == nil || IsNil(o.CollectionKey) {
		return nil, false
	}
	return o.CollectionKey, true
}

// HasCollectionKey returns a boolean if a field has been set.
func (o *UpdateDataSourceInstanceDto) HasCollectionKey() bool {
	if o != nil && !IsNil(o.CollectionKey) {
		return true
	}

	return false
}

// SetCollectionKey gets a reference to the given string and assigns it to the CollectionKey field.
func (o *UpdateDataSourceInstanceDto) SetCollectionKey(v string) {
	o.CollectionKey = &v
}

// GetCollectionParameters returns the CollectionParameters field value if set, zero value otherwise.
func (o *UpdateDataSourceInstanceDto) GetCollectionParameters() map[string]interface{} {
	if o == nil || IsNil(o.CollectionParameters) {
		var ret map[string]interface{}
		return ret
	}
	return o.CollectionParameters
}

// GetCollectionParametersOk returns a tuple with the CollectionParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDataSourceInstanceDto) GetCollectionParametersOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CollectionParameters) {
		return map[string]interface{}{}, false
	}
	return o.CollectionParameters, true
}

// HasCollectionParameters returns a boolean if a field has been set.
func (o *UpdateDataSourceInstanceDto) HasCollectionParameters() bool {
	if o != nil && !IsNil(o.CollectionParameters) {
		return true
	}

	return false
}

// SetCollectionParameters gets a reference to the given map[string]interface{} and assigns it to the CollectionParameters field.
func (o *UpdateDataSourceInstanceDto) SetCollectionParameters(v map[string]interface{}) {
	o.CollectionParameters = v
}

// GetPullUpdatesIntervalSeconds returns the PullUpdatesIntervalSeconds field value if set, zero value otherwise.
func (o *UpdateDataSourceInstanceDto) GetPullUpdatesIntervalSeconds() float32 {
	if o == nil || IsNil(o.PullUpdatesIntervalSeconds) {
		var ret float32
		return ret
	}
	return *o.PullUpdatesIntervalSeconds
}

// GetPullUpdatesIntervalSecondsOk returns a tuple with the PullUpdatesIntervalSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDataSourceInstanceDto) GetPullUpdatesIntervalSecondsOk() (*float32, bool) {
	if o == nil || IsNil(o.PullUpdatesIntervalSeconds) {
		return nil, false
	}
	return o.PullUpdatesIntervalSeconds, true
}

// HasPullUpdatesIntervalSeconds returns a boolean if a field has been set.
func (o *UpdateDataSourceInstanceDto) HasPullUpdatesIntervalSeconds() bool {
	if o != nil && !IsNil(o.PullUpdatesIntervalSeconds) {
		return true
	}

	return false
}

// SetPullUpdatesIntervalSeconds gets a reference to the given float32 and assigns it to the PullUpdatesIntervalSeconds field.
func (o *UpdateDataSourceInstanceDto) SetPullUpdatesIntervalSeconds(v float32) {
	o.PullUpdatesIntervalSeconds = &v
}

// GetFullSyncIntervalSeconds returns the FullSyncIntervalSeconds field value if set, zero value otherwise.
func (o *UpdateDataSourceInstanceDto) GetFullSyncIntervalSeconds() float32 {
	if o == nil || IsNil(o.FullSyncIntervalSeconds) {
		var ret float32
		return ret
	}
	return *o.FullSyncIntervalSeconds
}

// GetFullSyncIntervalSecondsOk returns a tuple with the FullSyncIntervalSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDataSourceInstanceDto) GetFullSyncIntervalSecondsOk() (*float32, bool) {
	if o == nil || IsNil(o.FullSyncIntervalSeconds) {
		return nil, false
	}
	return o.FullSyncIntervalSeconds, true
}

// HasFullSyncIntervalSeconds returns a boolean if a field has been set.
func (o *UpdateDataSourceInstanceDto) HasFullSyncIntervalSeconds() bool {
	if o != nil && !IsNil(o.FullSyncIntervalSeconds) {
		return true
	}

	return false
}

// SetFullSyncIntervalSeconds gets a reference to the given float32 and assigns it to the FullSyncIntervalSeconds field.
func (o *UpdateDataSourceInstanceDto) SetFullSyncIntervalSeconds(v float32) {
	o.FullSyncIntervalSeconds = &v
}

func (o UpdateDataSourceInstanceDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDataSourceInstanceDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.CollectionKey) {
		toSerialize["collectionKey"] = o.CollectionKey
	}
	if !IsNil(o.CollectionParameters) {
		toSerialize["collectionParameters"] = o.CollectionParameters
	}
	if !IsNil(o.PullUpdatesIntervalSeconds) {
		toSerialize["pullUpdatesIntervalSeconds"] = o.PullUpdatesIntervalSeconds
	}
	if !IsNil(o.FullSyncIntervalSeconds) {
		toSerialize["fullSyncIntervalSeconds"] = o.FullSyncIntervalSeconds
	}
	return toSerialize, nil
}

type NullableUpdateDataSourceInstanceDto struct {
	value *UpdateDataSourceInstanceDto
	isSet bool
}

func (v NullableUpdateDataSourceInstanceDto) Get() *UpdateDataSourceInstanceDto {
	return v.value
}

func (v *NullableUpdateDataSourceInstanceDto) Set(val *UpdateDataSourceInstanceDto) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDataSourceInstanceDto) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDataSourceInstanceDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDataSourceInstanceDto(val *UpdateDataSourceInstanceDto) *NullableUpdateDataSourceInstanceDto {
	return &NullableUpdateDataSourceInstanceDto{value: val, isSet: true}
}

func (v NullableUpdateDataSourceInstanceDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDataSourceInstanceDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


