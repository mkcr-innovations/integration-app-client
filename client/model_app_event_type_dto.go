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

// checks if the AppEventTypeDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppEventTypeDto{}

// AppEventTypeDto struct for AppEventTypeDto
type AppEventTypeDto struct {
	Id string `json:"id"`
	Key *string `json:"key,omitempty"`
	Name *string `json:"name,omitempty"`
	ArchivedAt *string `json:"archivedAt,omitempty"`
	Revision *string `json:"revision,omitempty"`
	PublishedRevision *string `json:"publishedRevision,omitempty"`
	GlobalWebhookUri *string `json:"globalWebhookUri,omitempty"`
	Example map[string]interface{} `json:"example,omitempty"`
	Schema map[string]interface{} `json:"schema,omitempty"`
	SubscribeRequest map[string]interface{} `json:"subscribeRequest,omitempty"`
	UserIdFormula map[string]interface{} `json:"userIdFormula,omitempty"`
}

type _AppEventTypeDto AppEventTypeDto

// NewAppEventTypeDto instantiates a new AppEventTypeDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppEventTypeDto(id string) *AppEventTypeDto {
	this := AppEventTypeDto{}
	this.Id = id
	return &this
}

// NewAppEventTypeDtoWithDefaults instantiates a new AppEventTypeDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppEventTypeDtoWithDefaults() *AppEventTypeDto {
	this := AppEventTypeDto{}
	return &this
}

// GetId returns the Id field value
func (o *AppEventTypeDto) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppEventTypeDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppEventTypeDto) SetId(v string) {
	o.Id = v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *AppEventTypeDto) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventTypeDto) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *AppEventTypeDto) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *AppEventTypeDto) SetKey(v string) {
	o.Key = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AppEventTypeDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventTypeDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AppEventTypeDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AppEventTypeDto) SetName(v string) {
	o.Name = &v
}

// GetArchivedAt returns the ArchivedAt field value if set, zero value otherwise.
func (o *AppEventTypeDto) GetArchivedAt() string {
	if o == nil || IsNil(o.ArchivedAt) {
		var ret string
		return ret
	}
	return *o.ArchivedAt
}

// GetArchivedAtOk returns a tuple with the ArchivedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventTypeDto) GetArchivedAtOk() (*string, bool) {
	if o == nil || IsNil(o.ArchivedAt) {
		return nil, false
	}
	return o.ArchivedAt, true
}

// HasArchivedAt returns a boolean if a field has been set.
func (o *AppEventTypeDto) HasArchivedAt() bool {
	if o != nil && !IsNil(o.ArchivedAt) {
		return true
	}

	return false
}

// SetArchivedAt gets a reference to the given string and assigns it to the ArchivedAt field.
func (o *AppEventTypeDto) SetArchivedAt(v string) {
	o.ArchivedAt = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *AppEventTypeDto) GetRevision() string {
	if o == nil || IsNil(o.Revision) {
		var ret string
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventTypeDto) GetRevisionOk() (*string, bool) {
	if o == nil || IsNil(o.Revision) {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *AppEventTypeDto) HasRevision() bool {
	if o != nil && !IsNil(o.Revision) {
		return true
	}

	return false
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *AppEventTypeDto) SetRevision(v string) {
	o.Revision = &v
}

// GetPublishedRevision returns the PublishedRevision field value if set, zero value otherwise.
func (o *AppEventTypeDto) GetPublishedRevision() string {
	if o == nil || IsNil(o.PublishedRevision) {
		var ret string
		return ret
	}
	return *o.PublishedRevision
}

// GetPublishedRevisionOk returns a tuple with the PublishedRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventTypeDto) GetPublishedRevisionOk() (*string, bool) {
	if o == nil || IsNil(o.PublishedRevision) {
		return nil, false
	}
	return o.PublishedRevision, true
}

// HasPublishedRevision returns a boolean if a field has been set.
func (o *AppEventTypeDto) HasPublishedRevision() bool {
	if o != nil && !IsNil(o.PublishedRevision) {
		return true
	}

	return false
}

// SetPublishedRevision gets a reference to the given string and assigns it to the PublishedRevision field.
func (o *AppEventTypeDto) SetPublishedRevision(v string) {
	o.PublishedRevision = &v
}

// GetGlobalWebhookUri returns the GlobalWebhookUri field value if set, zero value otherwise.
func (o *AppEventTypeDto) GetGlobalWebhookUri() string {
	if o == nil || IsNil(o.GlobalWebhookUri) {
		var ret string
		return ret
	}
	return *o.GlobalWebhookUri
}

// GetGlobalWebhookUriOk returns a tuple with the GlobalWebhookUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventTypeDto) GetGlobalWebhookUriOk() (*string, bool) {
	if o == nil || IsNil(o.GlobalWebhookUri) {
		return nil, false
	}
	return o.GlobalWebhookUri, true
}

// HasGlobalWebhookUri returns a boolean if a field has been set.
func (o *AppEventTypeDto) HasGlobalWebhookUri() bool {
	if o != nil && !IsNil(o.GlobalWebhookUri) {
		return true
	}

	return false
}

// SetGlobalWebhookUri gets a reference to the given string and assigns it to the GlobalWebhookUri field.
func (o *AppEventTypeDto) SetGlobalWebhookUri(v string) {
	o.GlobalWebhookUri = &v
}

// GetExample returns the Example field value if set, zero value otherwise.
func (o *AppEventTypeDto) GetExample() map[string]interface{} {
	if o == nil || IsNil(o.Example) {
		var ret map[string]interface{}
		return ret
	}
	return o.Example
}

// GetExampleOk returns a tuple with the Example field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventTypeDto) GetExampleOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Example) {
		return map[string]interface{}{}, false
	}
	return o.Example, true
}

// HasExample returns a boolean if a field has been set.
func (o *AppEventTypeDto) HasExample() bool {
	if o != nil && !IsNil(o.Example) {
		return true
	}

	return false
}

// SetExample gets a reference to the given map[string]interface{} and assigns it to the Example field.
func (o *AppEventTypeDto) SetExample(v map[string]interface{}) {
	o.Example = v
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *AppEventTypeDto) GetSchema() map[string]interface{} {
	if o == nil || IsNil(o.Schema) {
		var ret map[string]interface{}
		return ret
	}
	return o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventTypeDto) GetSchemaOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Schema) {
		return map[string]interface{}{}, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *AppEventTypeDto) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given map[string]interface{} and assigns it to the Schema field.
func (o *AppEventTypeDto) SetSchema(v map[string]interface{}) {
	o.Schema = v
}

// GetSubscribeRequest returns the SubscribeRequest field value if set, zero value otherwise.
func (o *AppEventTypeDto) GetSubscribeRequest() map[string]interface{} {
	if o == nil || IsNil(o.SubscribeRequest) {
		var ret map[string]interface{}
		return ret
	}
	return o.SubscribeRequest
}

// GetSubscribeRequestOk returns a tuple with the SubscribeRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventTypeDto) GetSubscribeRequestOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.SubscribeRequest) {
		return map[string]interface{}{}, false
	}
	return o.SubscribeRequest, true
}

// HasSubscribeRequest returns a boolean if a field has been set.
func (o *AppEventTypeDto) HasSubscribeRequest() bool {
	if o != nil && !IsNil(o.SubscribeRequest) {
		return true
	}

	return false
}

// SetSubscribeRequest gets a reference to the given map[string]interface{} and assigns it to the SubscribeRequest field.
func (o *AppEventTypeDto) SetSubscribeRequest(v map[string]interface{}) {
	o.SubscribeRequest = v
}

// GetUserIdFormula returns the UserIdFormula field value if set, zero value otherwise.
func (o *AppEventTypeDto) GetUserIdFormula() map[string]interface{} {
	if o == nil || IsNil(o.UserIdFormula) {
		var ret map[string]interface{}
		return ret
	}
	return o.UserIdFormula
}

// GetUserIdFormulaOk returns a tuple with the UserIdFormula field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventTypeDto) GetUserIdFormulaOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.UserIdFormula) {
		return map[string]interface{}{}, false
	}
	return o.UserIdFormula, true
}

// HasUserIdFormula returns a boolean if a field has been set.
func (o *AppEventTypeDto) HasUserIdFormula() bool {
	if o != nil && !IsNil(o.UserIdFormula) {
		return true
	}

	return false
}

// SetUserIdFormula gets a reference to the given map[string]interface{} and assigns it to the UserIdFormula field.
func (o *AppEventTypeDto) SetUserIdFormula(v map[string]interface{}) {
	o.UserIdFormula = v
}

func (o AppEventTypeDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppEventTypeDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ArchivedAt) {
		toSerialize["archivedAt"] = o.ArchivedAt
	}
	if !IsNil(o.Revision) {
		toSerialize["revision"] = o.Revision
	}
	if !IsNil(o.PublishedRevision) {
		toSerialize["publishedRevision"] = o.PublishedRevision
	}
	if !IsNil(o.GlobalWebhookUri) {
		toSerialize["globalWebhookUri"] = o.GlobalWebhookUri
	}
	if !IsNil(o.Example) {
		toSerialize["example"] = o.Example
	}
	if !IsNil(o.Schema) {
		toSerialize["schema"] = o.Schema
	}
	if !IsNil(o.SubscribeRequest) {
		toSerialize["subscribeRequest"] = o.SubscribeRequest
	}
	if !IsNil(o.UserIdFormula) {
		toSerialize["userIdFormula"] = o.UserIdFormula
	}
	return toSerialize, nil
}

func (o *AppEventTypeDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varAppEventTypeDto := _AppEventTypeDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppEventTypeDto)

	if err != nil {
		return err
	}

	*o = AppEventTypeDto(varAppEventTypeDto)

	return err
}

type NullableAppEventTypeDto struct {
	value *AppEventTypeDto
	isSet bool
}

func (v NullableAppEventTypeDto) Get() *AppEventTypeDto {
	return v.value
}

func (v *NullableAppEventTypeDto) Set(val *AppEventTypeDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAppEventTypeDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAppEventTypeDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppEventTypeDto(val *AppEventTypeDto) *NullableAppEventTypeDto {
	return &NullableAppEventTypeDto{value: val, isSet: true}
}

func (v NullableAppEventTypeDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppEventTypeDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


