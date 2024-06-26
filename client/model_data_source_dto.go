/*
Integration Engine API

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

// checks if the DataSourceDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataSourceDto{}

// DataSourceDto struct for DataSourceDto
type DataSourceDto struct {
	IntegrationId *string `json:"integrationId,omitempty"`
	Revision string `json:"revision"`
	Customized *bool `json:"customized,omitempty"`
	UniversalDataSourceId *string `json:"universalDataSourceId,omitempty"`
	UniversalDataSourceRevision *string `json:"universalDataSourceRevision,omitempty"`
	CollectionKey string `json:"collectionKey"`
	CollectionParameters map[string]interface{} `json:"collectionParameters,omitempty"`
	// Deprecated
	DefaultPath *string `json:"defaultPath,omitempty"`
	Udm *string `json:"udm,omitempty"`
	PullUpdatesIntervalSeconds *float32 `json:"pullUpdatesIntervalSeconds,omitempty"`
	FullSyncIntervalSeconds *float32 `json:"fullSyncIntervalSeconds,omitempty"`
	AppliedToIntegrations []string `json:"appliedToIntegrations,omitempty"`
}

type _DataSourceDto DataSourceDto

// NewDataSourceDto instantiates a new DataSourceDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataSourceDto(revision string, collectionKey string) *DataSourceDto {
	this := DataSourceDto{}
	this.Revision = revision
	this.CollectionKey = collectionKey
	return &this
}

// NewDataSourceDtoWithDefaults instantiates a new DataSourceDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataSourceDtoWithDefaults() *DataSourceDto {
	this := DataSourceDto{}
	return &this
}

// GetIntegrationId returns the IntegrationId field value if set, zero value otherwise.
func (o *DataSourceDto) GetIntegrationId() string {
	if o == nil || IsNil(o.IntegrationId) {
		var ret string
		return ret
	}
	return *o.IntegrationId
}

// GetIntegrationIdOk returns a tuple with the IntegrationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourceDto) GetIntegrationIdOk() (*string, bool) {
	if o == nil || IsNil(o.IntegrationId) {
		return nil, false
	}
	return o.IntegrationId, true
}

// HasIntegrationId returns a boolean if a field has been set.
func (o *DataSourceDto) HasIntegrationId() bool {
	if o != nil && !IsNil(o.IntegrationId) {
		return true
	}

	return false
}

// SetIntegrationId gets a reference to the given string and assigns it to the IntegrationId field.
func (o *DataSourceDto) SetIntegrationId(v string) {
	o.IntegrationId = &v
}

// GetRevision returns the Revision field value
func (o *DataSourceDto) GetRevision() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value
// and a boolean to check if the value has been set.
func (o *DataSourceDto) GetRevisionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Revision, true
}

// SetRevision sets field value
func (o *DataSourceDto) SetRevision(v string) {
	o.Revision = v
}

// GetCustomized returns the Customized field value if set, zero value otherwise.
func (o *DataSourceDto) GetCustomized() bool {
	if o == nil || IsNil(o.Customized) {
		var ret bool
		return ret
	}
	return *o.Customized
}

// GetCustomizedOk returns a tuple with the Customized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourceDto) GetCustomizedOk() (*bool, bool) {
	if o == nil || IsNil(o.Customized) {
		return nil, false
	}
	return o.Customized, true
}

// HasCustomized returns a boolean if a field has been set.
func (o *DataSourceDto) HasCustomized() bool {
	if o != nil && !IsNil(o.Customized) {
		return true
	}

	return false
}

// SetCustomized gets a reference to the given bool and assigns it to the Customized field.
func (o *DataSourceDto) SetCustomized(v bool) {
	o.Customized = &v
}

// GetUniversalDataSourceId returns the UniversalDataSourceId field value if set, zero value otherwise.
func (o *DataSourceDto) GetUniversalDataSourceId() string {
	if o == nil || IsNil(o.UniversalDataSourceId) {
		var ret string
		return ret
	}
	return *o.UniversalDataSourceId
}

// GetUniversalDataSourceIdOk returns a tuple with the UniversalDataSourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourceDto) GetUniversalDataSourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.UniversalDataSourceId) {
		return nil, false
	}
	return o.UniversalDataSourceId, true
}

// HasUniversalDataSourceId returns a boolean if a field has been set.
func (o *DataSourceDto) HasUniversalDataSourceId() bool {
	if o != nil && !IsNil(o.UniversalDataSourceId) {
		return true
	}

	return false
}

// SetUniversalDataSourceId gets a reference to the given string and assigns it to the UniversalDataSourceId field.
func (o *DataSourceDto) SetUniversalDataSourceId(v string) {
	o.UniversalDataSourceId = &v
}

// GetUniversalDataSourceRevision returns the UniversalDataSourceRevision field value if set, zero value otherwise.
func (o *DataSourceDto) GetUniversalDataSourceRevision() string {
	if o == nil || IsNil(o.UniversalDataSourceRevision) {
		var ret string
		return ret
	}
	return *o.UniversalDataSourceRevision
}

// GetUniversalDataSourceRevisionOk returns a tuple with the UniversalDataSourceRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourceDto) GetUniversalDataSourceRevisionOk() (*string, bool) {
	if o == nil || IsNil(o.UniversalDataSourceRevision) {
		return nil, false
	}
	return o.UniversalDataSourceRevision, true
}

// HasUniversalDataSourceRevision returns a boolean if a field has been set.
func (o *DataSourceDto) HasUniversalDataSourceRevision() bool {
	if o != nil && !IsNil(o.UniversalDataSourceRevision) {
		return true
	}

	return false
}

// SetUniversalDataSourceRevision gets a reference to the given string and assigns it to the UniversalDataSourceRevision field.
func (o *DataSourceDto) SetUniversalDataSourceRevision(v string) {
	o.UniversalDataSourceRevision = &v
}

// GetCollectionKey returns the CollectionKey field value
func (o *DataSourceDto) GetCollectionKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CollectionKey
}

// GetCollectionKeyOk returns a tuple with the CollectionKey field value
// and a boolean to check if the value has been set.
func (o *DataSourceDto) GetCollectionKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectionKey, true
}

// SetCollectionKey sets field value
func (o *DataSourceDto) SetCollectionKey(v string) {
	o.CollectionKey = v
}

// GetCollectionParameters returns the CollectionParameters field value if set, zero value otherwise.
func (o *DataSourceDto) GetCollectionParameters() map[string]interface{} {
	if o == nil || IsNil(o.CollectionParameters) {
		var ret map[string]interface{}
		return ret
	}
	return o.CollectionParameters
}

// GetCollectionParametersOk returns a tuple with the CollectionParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourceDto) GetCollectionParametersOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CollectionParameters) {
		return map[string]interface{}{}, false
	}
	return o.CollectionParameters, true
}

// HasCollectionParameters returns a boolean if a field has been set.
func (o *DataSourceDto) HasCollectionParameters() bool {
	if o != nil && !IsNil(o.CollectionParameters) {
		return true
	}

	return false
}

// SetCollectionParameters gets a reference to the given map[string]interface{} and assigns it to the CollectionParameters field.
func (o *DataSourceDto) SetCollectionParameters(v map[string]interface{}) {
	o.CollectionParameters = v
}

// GetDefaultPath returns the DefaultPath field value if set, zero value otherwise.
// Deprecated
func (o *DataSourceDto) GetDefaultPath() string {
	if o == nil || IsNil(o.DefaultPath) {
		var ret string
		return ret
	}
	return *o.DefaultPath
}

// GetDefaultPathOk returns a tuple with the DefaultPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *DataSourceDto) GetDefaultPathOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultPath) {
		return nil, false
	}
	return o.DefaultPath, true
}

// HasDefaultPath returns a boolean if a field has been set.
func (o *DataSourceDto) HasDefaultPath() bool {
	if o != nil && !IsNil(o.DefaultPath) {
		return true
	}

	return false
}

// SetDefaultPath gets a reference to the given string and assigns it to the DefaultPath field.
// Deprecated
func (o *DataSourceDto) SetDefaultPath(v string) {
	o.DefaultPath = &v
}

// GetUdm returns the Udm field value if set, zero value otherwise.
func (o *DataSourceDto) GetUdm() string {
	if o == nil || IsNil(o.Udm) {
		var ret string
		return ret
	}
	return *o.Udm
}

// GetUdmOk returns a tuple with the Udm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourceDto) GetUdmOk() (*string, bool) {
	if o == nil || IsNil(o.Udm) {
		return nil, false
	}
	return o.Udm, true
}

// HasUdm returns a boolean if a field has been set.
func (o *DataSourceDto) HasUdm() bool {
	if o != nil && !IsNil(o.Udm) {
		return true
	}

	return false
}

// SetUdm gets a reference to the given string and assigns it to the Udm field.
func (o *DataSourceDto) SetUdm(v string) {
	o.Udm = &v
}

// GetPullUpdatesIntervalSeconds returns the PullUpdatesIntervalSeconds field value if set, zero value otherwise.
func (o *DataSourceDto) GetPullUpdatesIntervalSeconds() float32 {
	if o == nil || IsNil(o.PullUpdatesIntervalSeconds) {
		var ret float32
		return ret
	}
	return *o.PullUpdatesIntervalSeconds
}

// GetPullUpdatesIntervalSecondsOk returns a tuple with the PullUpdatesIntervalSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourceDto) GetPullUpdatesIntervalSecondsOk() (*float32, bool) {
	if o == nil || IsNil(o.PullUpdatesIntervalSeconds) {
		return nil, false
	}
	return o.PullUpdatesIntervalSeconds, true
}

// HasPullUpdatesIntervalSeconds returns a boolean if a field has been set.
func (o *DataSourceDto) HasPullUpdatesIntervalSeconds() bool {
	if o != nil && !IsNil(o.PullUpdatesIntervalSeconds) {
		return true
	}

	return false
}

// SetPullUpdatesIntervalSeconds gets a reference to the given float32 and assigns it to the PullUpdatesIntervalSeconds field.
func (o *DataSourceDto) SetPullUpdatesIntervalSeconds(v float32) {
	o.PullUpdatesIntervalSeconds = &v
}

// GetFullSyncIntervalSeconds returns the FullSyncIntervalSeconds field value if set, zero value otherwise.
func (o *DataSourceDto) GetFullSyncIntervalSeconds() float32 {
	if o == nil || IsNil(o.FullSyncIntervalSeconds) {
		var ret float32
		return ret
	}
	return *o.FullSyncIntervalSeconds
}

// GetFullSyncIntervalSecondsOk returns a tuple with the FullSyncIntervalSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourceDto) GetFullSyncIntervalSecondsOk() (*float32, bool) {
	if o == nil || IsNil(o.FullSyncIntervalSeconds) {
		return nil, false
	}
	return o.FullSyncIntervalSeconds, true
}

// HasFullSyncIntervalSeconds returns a boolean if a field has been set.
func (o *DataSourceDto) HasFullSyncIntervalSeconds() bool {
	if o != nil && !IsNil(o.FullSyncIntervalSeconds) {
		return true
	}

	return false
}

// SetFullSyncIntervalSeconds gets a reference to the given float32 and assigns it to the FullSyncIntervalSeconds field.
func (o *DataSourceDto) SetFullSyncIntervalSeconds(v float32) {
	o.FullSyncIntervalSeconds = &v
}

// GetAppliedToIntegrations returns the AppliedToIntegrations field value if set, zero value otherwise.
func (o *DataSourceDto) GetAppliedToIntegrations() []string {
	if o == nil || IsNil(o.AppliedToIntegrations) {
		var ret []string
		return ret
	}
	return o.AppliedToIntegrations
}

// GetAppliedToIntegrationsOk returns a tuple with the AppliedToIntegrations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourceDto) GetAppliedToIntegrationsOk() ([]string, bool) {
	if o == nil || IsNil(o.AppliedToIntegrations) {
		return nil, false
	}
	return o.AppliedToIntegrations, true
}

// HasAppliedToIntegrations returns a boolean if a field has been set.
func (o *DataSourceDto) HasAppliedToIntegrations() bool {
	if o != nil && !IsNil(o.AppliedToIntegrations) {
		return true
	}

	return false
}

// SetAppliedToIntegrations gets a reference to the given []string and assigns it to the AppliedToIntegrations field.
func (o *DataSourceDto) SetAppliedToIntegrations(v []string) {
	o.AppliedToIntegrations = v
}

func (o DataSourceDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataSourceDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IntegrationId) {
		toSerialize["integrationId"] = o.IntegrationId
	}
	toSerialize["revision"] = o.Revision
	if !IsNil(o.Customized) {
		toSerialize["customized"] = o.Customized
	}
	if !IsNil(o.UniversalDataSourceId) {
		toSerialize["universalDataSourceId"] = o.UniversalDataSourceId
	}
	if !IsNil(o.UniversalDataSourceRevision) {
		toSerialize["universalDataSourceRevision"] = o.UniversalDataSourceRevision
	}
	toSerialize["collectionKey"] = o.CollectionKey
	if !IsNil(o.CollectionParameters) {
		toSerialize["collectionParameters"] = o.CollectionParameters
	}
	if !IsNil(o.DefaultPath) {
		toSerialize["defaultPath"] = o.DefaultPath
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
	if !IsNil(o.AppliedToIntegrations) {
		toSerialize["appliedToIntegrations"] = o.AppliedToIntegrations
	}
	return toSerialize, nil
}

func (o *DataSourceDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"revision",
		"collectionKey",
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

	varDataSourceDto := _DataSourceDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDataSourceDto)

	if err != nil {
		return err
	}

	*o = DataSourceDto(varDataSourceDto)

	return err
}

type NullableDataSourceDto struct {
	value *DataSourceDto
	isSet bool
}

func (v NullableDataSourceDto) Get() *DataSourceDto {
	return v.value
}

func (v *NullableDataSourceDto) Set(val *DataSourceDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDataSourceDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDataSourceDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataSourceDto(val *DataSourceDto) *NullableDataSourceDto {
	return &NullableDataSourceDto{value: val, isSet: true}
}

func (v NullableDataSourceDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataSourceDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


