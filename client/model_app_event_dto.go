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

// checks if the AppEventDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppEventDto{}

// AppEventDto struct for AppEventDto
type AppEventDto struct {
	Id string `json:"id"`
	UserId *string `json:"userId,omitempty"`
	User *CustomerDto `json:"user,omitempty"`
	AppEventTypeId *string `json:"appEventTypeId,omitempty"`
	AppEventType *AppEventTypeDto `json:"appEventType,omitempty"`
	AppEventSubscriptionId *string `json:"appEventSubscriptionId,omitempty"`
	AppEventSubscription *AppEventSubscriptionDto `json:"appEventSubscription,omitempty"`
	Event map[string]interface{} `json:"event,omitempty"`
	Datetime *string `json:"datetime,omitempty"`
	LaunchedFlowRunIds []string `json:"launchedFlowRunIds,omitempty"`
}

type _AppEventDto AppEventDto

// NewAppEventDto instantiates a new AppEventDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppEventDto(id string) *AppEventDto {
	this := AppEventDto{}
	this.Id = id
	return &this
}

// NewAppEventDtoWithDefaults instantiates a new AppEventDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppEventDtoWithDefaults() *AppEventDto {
	this := AppEventDto{}
	return &this
}

// GetId returns the Id field value
func (o *AppEventDto) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppEventDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppEventDto) SetId(v string) {
	o.Id = v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *AppEventDto) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventDto) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *AppEventDto) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *AppEventDto) SetUserId(v string) {
	o.UserId = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *AppEventDto) GetUser() CustomerDto {
	if o == nil || IsNil(o.User) {
		var ret CustomerDto
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventDto) GetUserOk() (*CustomerDto, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *AppEventDto) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given CustomerDto and assigns it to the User field.
func (o *AppEventDto) SetUser(v CustomerDto) {
	o.User = &v
}

// GetAppEventTypeId returns the AppEventTypeId field value if set, zero value otherwise.
func (o *AppEventDto) GetAppEventTypeId() string {
	if o == nil || IsNil(o.AppEventTypeId) {
		var ret string
		return ret
	}
	return *o.AppEventTypeId
}

// GetAppEventTypeIdOk returns a tuple with the AppEventTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventDto) GetAppEventTypeIdOk() (*string, bool) {
	if o == nil || IsNil(o.AppEventTypeId) {
		return nil, false
	}
	return o.AppEventTypeId, true
}

// HasAppEventTypeId returns a boolean if a field has been set.
func (o *AppEventDto) HasAppEventTypeId() bool {
	if o != nil && !IsNil(o.AppEventTypeId) {
		return true
	}

	return false
}

// SetAppEventTypeId gets a reference to the given string and assigns it to the AppEventTypeId field.
func (o *AppEventDto) SetAppEventTypeId(v string) {
	o.AppEventTypeId = &v
}

// GetAppEventType returns the AppEventType field value if set, zero value otherwise.
func (o *AppEventDto) GetAppEventType() AppEventTypeDto {
	if o == nil || IsNil(o.AppEventType) {
		var ret AppEventTypeDto
		return ret
	}
	return *o.AppEventType
}

// GetAppEventTypeOk returns a tuple with the AppEventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventDto) GetAppEventTypeOk() (*AppEventTypeDto, bool) {
	if o == nil || IsNil(o.AppEventType) {
		return nil, false
	}
	return o.AppEventType, true
}

// HasAppEventType returns a boolean if a field has been set.
func (o *AppEventDto) HasAppEventType() bool {
	if o != nil && !IsNil(o.AppEventType) {
		return true
	}

	return false
}

// SetAppEventType gets a reference to the given AppEventTypeDto and assigns it to the AppEventType field.
func (o *AppEventDto) SetAppEventType(v AppEventTypeDto) {
	o.AppEventType = &v
}

// GetAppEventSubscriptionId returns the AppEventSubscriptionId field value if set, zero value otherwise.
func (o *AppEventDto) GetAppEventSubscriptionId() string {
	if o == nil || IsNil(o.AppEventSubscriptionId) {
		var ret string
		return ret
	}
	return *o.AppEventSubscriptionId
}

// GetAppEventSubscriptionIdOk returns a tuple with the AppEventSubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventDto) GetAppEventSubscriptionIdOk() (*string, bool) {
	if o == nil || IsNil(o.AppEventSubscriptionId) {
		return nil, false
	}
	return o.AppEventSubscriptionId, true
}

// HasAppEventSubscriptionId returns a boolean if a field has been set.
func (o *AppEventDto) HasAppEventSubscriptionId() bool {
	if o != nil && !IsNil(o.AppEventSubscriptionId) {
		return true
	}

	return false
}

// SetAppEventSubscriptionId gets a reference to the given string and assigns it to the AppEventSubscriptionId field.
func (o *AppEventDto) SetAppEventSubscriptionId(v string) {
	o.AppEventSubscriptionId = &v
}

// GetAppEventSubscription returns the AppEventSubscription field value if set, zero value otherwise.
func (o *AppEventDto) GetAppEventSubscription() AppEventSubscriptionDto {
	if o == nil || IsNil(o.AppEventSubscription) {
		var ret AppEventSubscriptionDto
		return ret
	}
	return *o.AppEventSubscription
}

// GetAppEventSubscriptionOk returns a tuple with the AppEventSubscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventDto) GetAppEventSubscriptionOk() (*AppEventSubscriptionDto, bool) {
	if o == nil || IsNil(o.AppEventSubscription) {
		return nil, false
	}
	return o.AppEventSubscription, true
}

// HasAppEventSubscription returns a boolean if a field has been set.
func (o *AppEventDto) HasAppEventSubscription() bool {
	if o != nil && !IsNil(o.AppEventSubscription) {
		return true
	}

	return false
}

// SetAppEventSubscription gets a reference to the given AppEventSubscriptionDto and assigns it to the AppEventSubscription field.
func (o *AppEventDto) SetAppEventSubscription(v AppEventSubscriptionDto) {
	o.AppEventSubscription = &v
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *AppEventDto) GetEvent() map[string]interface{} {
	if o == nil || IsNil(o.Event) {
		var ret map[string]interface{}
		return ret
	}
	return o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventDto) GetEventOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Event) {
		return map[string]interface{}{}, false
	}
	return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *AppEventDto) HasEvent() bool {
	if o != nil && !IsNil(o.Event) {
		return true
	}

	return false
}

// SetEvent gets a reference to the given map[string]interface{} and assigns it to the Event field.
func (o *AppEventDto) SetEvent(v map[string]interface{}) {
	o.Event = v
}

// GetDatetime returns the Datetime field value if set, zero value otherwise.
func (o *AppEventDto) GetDatetime() string {
	if o == nil || IsNil(o.Datetime) {
		var ret string
		return ret
	}
	return *o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventDto) GetDatetimeOk() (*string, bool) {
	if o == nil || IsNil(o.Datetime) {
		return nil, false
	}
	return o.Datetime, true
}

// HasDatetime returns a boolean if a field has been set.
func (o *AppEventDto) HasDatetime() bool {
	if o != nil && !IsNil(o.Datetime) {
		return true
	}

	return false
}

// SetDatetime gets a reference to the given string and assigns it to the Datetime field.
func (o *AppEventDto) SetDatetime(v string) {
	o.Datetime = &v
}

// GetLaunchedFlowRunIds returns the LaunchedFlowRunIds field value if set, zero value otherwise.
func (o *AppEventDto) GetLaunchedFlowRunIds() []string {
	if o == nil || IsNil(o.LaunchedFlowRunIds) {
		var ret []string
		return ret
	}
	return o.LaunchedFlowRunIds
}

// GetLaunchedFlowRunIdsOk returns a tuple with the LaunchedFlowRunIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventDto) GetLaunchedFlowRunIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.LaunchedFlowRunIds) {
		return nil, false
	}
	return o.LaunchedFlowRunIds, true
}

// HasLaunchedFlowRunIds returns a boolean if a field has been set.
func (o *AppEventDto) HasLaunchedFlowRunIds() bool {
	if o != nil && !IsNil(o.LaunchedFlowRunIds) {
		return true
	}

	return false
}

// SetLaunchedFlowRunIds gets a reference to the given []string and assigns it to the LaunchedFlowRunIds field.
func (o *AppEventDto) SetLaunchedFlowRunIds(v []string) {
	o.LaunchedFlowRunIds = v
}

func (o AppEventDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppEventDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !IsNil(o.AppEventTypeId) {
		toSerialize["appEventTypeId"] = o.AppEventTypeId
	}
	if !IsNil(o.AppEventType) {
		toSerialize["appEventType"] = o.AppEventType
	}
	if !IsNil(o.AppEventSubscriptionId) {
		toSerialize["appEventSubscriptionId"] = o.AppEventSubscriptionId
	}
	if !IsNil(o.AppEventSubscription) {
		toSerialize["appEventSubscription"] = o.AppEventSubscription
	}
	if !IsNil(o.Event) {
		toSerialize["event"] = o.Event
	}
	if !IsNil(o.Datetime) {
		toSerialize["datetime"] = o.Datetime
	}
	if !IsNil(o.LaunchedFlowRunIds) {
		toSerialize["launchedFlowRunIds"] = o.LaunchedFlowRunIds
	}
	return toSerialize, nil
}

func (o *AppEventDto) UnmarshalJSON(data []byte) (err error) {
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

	varAppEventDto := _AppEventDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppEventDto)

	if err != nil {
		return err
	}

	*o = AppEventDto(varAppEventDto)

	return err
}

type NullableAppEventDto struct {
	value *AppEventDto
	isSet bool
}

func (v NullableAppEventDto) Get() *AppEventDto {
	return v.value
}

func (v *NullableAppEventDto) Set(val *AppEventDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAppEventDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAppEventDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppEventDto(val *AppEventDto) *NullableAppEventDto {
	return &NullableAppEventDto{value: val, isSet: true}
}

func (v NullableAppEventDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppEventDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


