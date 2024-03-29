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

// checks if the ExternalEventPullDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalEventPullDto{}

// ExternalEventPullDto struct for ExternalEventPullDto
type ExternalEventPullDto struct {
	Id string `json:"id"`
	UserId string `json:"userId"`
	User *CustomerDto `json:"user,omitempty"`
	ExternalEventSubscriptionId string `json:"externalEventSubscriptionId"`
	ExternalEventSubscription map[string]interface{} `json:"externalEventSubscription,omitempty"`
	ConnectionId string `json:"connectionId"`
	Connection *ConnectionDto `json:"connection,omitempty"`
	IntegrationId string `json:"integrationId"`
	Integration *IntegrationDto `json:"integration,omitempty"`
	Status string `json:"status"`
	StartDatetime string `json:"startDatetime"`
	EndDatetime string `json:"endDatetime"`
	IsFullScan bool `json:"isFullScan"`
	CollectedEventIds []string `json:"collectedEventIds"`
	Error map[string]interface{} `json:"error,omitempty"`
}

type _ExternalEventPullDto ExternalEventPullDto

// NewExternalEventPullDto instantiates a new ExternalEventPullDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalEventPullDto(id string, userId string, externalEventSubscriptionId string, connectionId string, integrationId string, status string, startDatetime string, endDatetime string, isFullScan bool, collectedEventIds []string) *ExternalEventPullDto {
	this := ExternalEventPullDto{}
	this.Id = id
	this.UserId = userId
	this.ExternalEventSubscriptionId = externalEventSubscriptionId
	this.ConnectionId = connectionId
	this.IntegrationId = integrationId
	this.Status = status
	this.StartDatetime = startDatetime
	this.EndDatetime = endDatetime
	this.IsFullScan = isFullScan
	this.CollectedEventIds = collectedEventIds
	return &this
}

// NewExternalEventPullDtoWithDefaults instantiates a new ExternalEventPullDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalEventPullDtoWithDefaults() *ExternalEventPullDto {
	this := ExternalEventPullDto{}
	return &this
}

// GetId returns the Id field value
func (o *ExternalEventPullDto) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ExternalEventPullDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ExternalEventPullDto) SetId(v string) {
	o.Id = v
}

// GetUserId returns the UserId field value
func (o *ExternalEventPullDto) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *ExternalEventPullDto) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *ExternalEventPullDto) SetUserId(v string) {
	o.UserId = v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *ExternalEventPullDto) GetUser() CustomerDto {
	if o == nil || IsNil(o.User) {
		var ret CustomerDto
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalEventPullDto) GetUserOk() (*CustomerDto, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *ExternalEventPullDto) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given CustomerDto and assigns it to the User field.
func (o *ExternalEventPullDto) SetUser(v CustomerDto) {
	o.User = &v
}

// GetExternalEventSubscriptionId returns the ExternalEventSubscriptionId field value
func (o *ExternalEventPullDto) GetExternalEventSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalEventSubscriptionId
}

// GetExternalEventSubscriptionIdOk returns a tuple with the ExternalEventSubscriptionId field value
// and a boolean to check if the value has been set.
func (o *ExternalEventPullDto) GetExternalEventSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalEventSubscriptionId, true
}

// SetExternalEventSubscriptionId sets field value
func (o *ExternalEventPullDto) SetExternalEventSubscriptionId(v string) {
	o.ExternalEventSubscriptionId = v
}

// GetExternalEventSubscription returns the ExternalEventSubscription field value if set, zero value otherwise.
func (o *ExternalEventPullDto) GetExternalEventSubscription() map[string]interface{} {
	if o == nil || IsNil(o.ExternalEventSubscription) {
		var ret map[string]interface{}
		return ret
	}
	return o.ExternalEventSubscription
}

// GetExternalEventSubscriptionOk returns a tuple with the ExternalEventSubscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalEventPullDto) GetExternalEventSubscriptionOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ExternalEventSubscription) {
		return map[string]interface{}{}, false
	}
	return o.ExternalEventSubscription, true
}

// HasExternalEventSubscription returns a boolean if a field has been set.
func (o *ExternalEventPullDto) HasExternalEventSubscription() bool {
	if o != nil && !IsNil(o.ExternalEventSubscription) {
		return true
	}

	return false
}

// SetExternalEventSubscription gets a reference to the given map[string]interface{} and assigns it to the ExternalEventSubscription field.
func (o *ExternalEventPullDto) SetExternalEventSubscription(v map[string]interface{}) {
	o.ExternalEventSubscription = v
}

// GetConnectionId returns the ConnectionId field value
func (o *ExternalEventPullDto) GetConnectionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionId
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value
// and a boolean to check if the value has been set.
func (o *ExternalEventPullDto) GetConnectionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionId, true
}

// SetConnectionId sets field value
func (o *ExternalEventPullDto) SetConnectionId(v string) {
	o.ConnectionId = v
}

// GetConnection returns the Connection field value if set, zero value otherwise.
func (o *ExternalEventPullDto) GetConnection() ConnectionDto {
	if o == nil || IsNil(o.Connection) {
		var ret ConnectionDto
		return ret
	}
	return *o.Connection
}

// GetConnectionOk returns a tuple with the Connection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalEventPullDto) GetConnectionOk() (*ConnectionDto, bool) {
	if o == nil || IsNil(o.Connection) {
		return nil, false
	}
	return o.Connection, true
}

// HasConnection returns a boolean if a field has been set.
func (o *ExternalEventPullDto) HasConnection() bool {
	if o != nil && !IsNil(o.Connection) {
		return true
	}

	return false
}

// SetConnection gets a reference to the given ConnectionDto and assigns it to the Connection field.
func (o *ExternalEventPullDto) SetConnection(v ConnectionDto) {
	o.Connection = &v
}

// GetIntegrationId returns the IntegrationId field value
func (o *ExternalEventPullDto) GetIntegrationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IntegrationId
}

// GetIntegrationIdOk returns a tuple with the IntegrationId field value
// and a boolean to check if the value has been set.
func (o *ExternalEventPullDto) GetIntegrationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntegrationId, true
}

// SetIntegrationId sets field value
func (o *ExternalEventPullDto) SetIntegrationId(v string) {
	o.IntegrationId = v
}

// GetIntegration returns the Integration field value if set, zero value otherwise.
func (o *ExternalEventPullDto) GetIntegration() IntegrationDto {
	if o == nil || IsNil(o.Integration) {
		var ret IntegrationDto
		return ret
	}
	return *o.Integration
}

// GetIntegrationOk returns a tuple with the Integration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalEventPullDto) GetIntegrationOk() (*IntegrationDto, bool) {
	if o == nil || IsNil(o.Integration) {
		return nil, false
	}
	return o.Integration, true
}

// HasIntegration returns a boolean if a field has been set.
func (o *ExternalEventPullDto) HasIntegration() bool {
	if o != nil && !IsNil(o.Integration) {
		return true
	}

	return false
}

// SetIntegration gets a reference to the given IntegrationDto and assigns it to the Integration field.
func (o *ExternalEventPullDto) SetIntegration(v IntegrationDto) {
	o.Integration = &v
}

// GetStatus returns the Status field value
func (o *ExternalEventPullDto) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ExternalEventPullDto) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ExternalEventPullDto) SetStatus(v string) {
	o.Status = v
}

// GetStartDatetime returns the StartDatetime field value
func (o *ExternalEventPullDto) GetStartDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartDatetime
}

// GetStartDatetimeOk returns a tuple with the StartDatetime field value
// and a boolean to check if the value has been set.
func (o *ExternalEventPullDto) GetStartDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDatetime, true
}

// SetStartDatetime sets field value
func (o *ExternalEventPullDto) SetStartDatetime(v string) {
	o.StartDatetime = v
}

// GetEndDatetime returns the EndDatetime field value
func (o *ExternalEventPullDto) GetEndDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndDatetime
}

// GetEndDatetimeOk returns a tuple with the EndDatetime field value
// and a boolean to check if the value has been set.
func (o *ExternalEventPullDto) GetEndDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndDatetime, true
}

// SetEndDatetime sets field value
func (o *ExternalEventPullDto) SetEndDatetime(v string) {
	o.EndDatetime = v
}

// GetIsFullScan returns the IsFullScan field value
func (o *ExternalEventPullDto) GetIsFullScan() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsFullScan
}

// GetIsFullScanOk returns a tuple with the IsFullScan field value
// and a boolean to check if the value has been set.
func (o *ExternalEventPullDto) GetIsFullScanOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsFullScan, true
}

// SetIsFullScan sets field value
func (o *ExternalEventPullDto) SetIsFullScan(v bool) {
	o.IsFullScan = v
}

// GetCollectedEventIds returns the CollectedEventIds field value
func (o *ExternalEventPullDto) GetCollectedEventIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.CollectedEventIds
}

// GetCollectedEventIdsOk returns a tuple with the CollectedEventIds field value
// and a boolean to check if the value has been set.
func (o *ExternalEventPullDto) GetCollectedEventIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CollectedEventIds, true
}

// SetCollectedEventIds sets field value
func (o *ExternalEventPullDto) SetCollectedEventIds(v []string) {
	o.CollectedEventIds = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ExternalEventPullDto) GetError() map[string]interface{} {
	if o == nil || IsNil(o.Error) {
		var ret map[string]interface{}
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalEventPullDto) GetErrorOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Error) {
		return map[string]interface{}{}, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ExternalEventPullDto) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given map[string]interface{} and assigns it to the Error field.
func (o *ExternalEventPullDto) SetError(v map[string]interface{}) {
	o.Error = v
}

func (o ExternalEventPullDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalEventPullDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["userId"] = o.UserId
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	toSerialize["externalEventSubscriptionId"] = o.ExternalEventSubscriptionId
	if !IsNil(o.ExternalEventSubscription) {
		toSerialize["externalEventSubscription"] = o.ExternalEventSubscription
	}
	toSerialize["connectionId"] = o.ConnectionId
	if !IsNil(o.Connection) {
		toSerialize["connection"] = o.Connection
	}
	toSerialize["integrationId"] = o.IntegrationId
	if !IsNil(o.Integration) {
		toSerialize["integration"] = o.Integration
	}
	toSerialize["status"] = o.Status
	toSerialize["startDatetime"] = o.StartDatetime
	toSerialize["endDatetime"] = o.EndDatetime
	toSerialize["isFullScan"] = o.IsFullScan
	toSerialize["collectedEventIds"] = o.CollectedEventIds
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

func (o *ExternalEventPullDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"userId",
		"externalEventSubscriptionId",
		"connectionId",
		"integrationId",
		"status",
		"startDatetime",
		"endDatetime",
		"isFullScan",
		"collectedEventIds",
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

	varExternalEventPullDto := _ExternalEventPullDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExternalEventPullDto)

	if err != nil {
		return err
	}

	*o = ExternalEventPullDto(varExternalEventPullDto)

	return err
}

type NullableExternalEventPullDto struct {
	value *ExternalEventPullDto
	isSet bool
}

func (v NullableExternalEventPullDto) Get() *ExternalEventPullDto {
	return v.value
}

func (v *NullableExternalEventPullDto) Set(val *ExternalEventPullDto) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalEventPullDto) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalEventPullDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalEventPullDto(val *ExternalEventPullDto) *NullableExternalEventPullDto {
	return &NullableExternalEventPullDto{value: val, isSet: true}
}

func (v NullableExternalEventPullDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalEventPullDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


