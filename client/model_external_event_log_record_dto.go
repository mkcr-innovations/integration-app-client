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

// checks if the ExternalEventLogRecordDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalEventLogRecordDto{}

// ExternalEventLogRecordDto struct for ExternalEventLogRecordDto
type ExternalEventLogRecordDto struct {
	Id string `json:"id"`
	UserId string `json:"userId"`
	User *CustomerDto `json:"user,omitempty"`
	ExternalEventId string `json:"externalEventId"`
	ExternalEventSubscriptionId string `json:"externalEventSubscriptionId"`
	ConnectionId string `json:"connectionId"`
	Connection *ConnectionDto `json:"connection,omitempty"`
	IntegrationId string `json:"integrationId"`
	Integration *IntegrationDto `json:"integration,omitempty"`
	Payload map[string]interface{} `json:"payload"`
	LaunchedFlowRunIds []string `json:"launchedFlowRunIds"`
	Error map[string]interface{} `json:"error,omitempty"`
	Logs map[string]interface{} `json:"logs,omitempty"`
}

type _ExternalEventLogRecordDto ExternalEventLogRecordDto

// NewExternalEventLogRecordDto instantiates a new ExternalEventLogRecordDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalEventLogRecordDto(id string, userId string, externalEventId string, externalEventSubscriptionId string, connectionId string, integrationId string, payload map[string]interface{}, launchedFlowRunIds []string) *ExternalEventLogRecordDto {
	this := ExternalEventLogRecordDto{}
	this.Id = id
	this.UserId = userId
	this.ExternalEventId = externalEventId
	this.ExternalEventSubscriptionId = externalEventSubscriptionId
	this.ConnectionId = connectionId
	this.IntegrationId = integrationId
	this.Payload = payload
	this.LaunchedFlowRunIds = launchedFlowRunIds
	return &this
}

// NewExternalEventLogRecordDtoWithDefaults instantiates a new ExternalEventLogRecordDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalEventLogRecordDtoWithDefaults() *ExternalEventLogRecordDto {
	this := ExternalEventLogRecordDto{}
	return &this
}

// GetId returns the Id field value
func (o *ExternalEventLogRecordDto) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ExternalEventLogRecordDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ExternalEventLogRecordDto) SetId(v string) {
	o.Id = v
}

// GetUserId returns the UserId field value
func (o *ExternalEventLogRecordDto) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *ExternalEventLogRecordDto) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *ExternalEventLogRecordDto) SetUserId(v string) {
	o.UserId = v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *ExternalEventLogRecordDto) GetUser() CustomerDto {
	if o == nil || IsNil(o.User) {
		var ret CustomerDto
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalEventLogRecordDto) GetUserOk() (*CustomerDto, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *ExternalEventLogRecordDto) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given CustomerDto and assigns it to the User field.
func (o *ExternalEventLogRecordDto) SetUser(v CustomerDto) {
	o.User = &v
}

// GetExternalEventId returns the ExternalEventId field value
func (o *ExternalEventLogRecordDto) GetExternalEventId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalEventId
}

// GetExternalEventIdOk returns a tuple with the ExternalEventId field value
// and a boolean to check if the value has been set.
func (o *ExternalEventLogRecordDto) GetExternalEventIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalEventId, true
}

// SetExternalEventId sets field value
func (o *ExternalEventLogRecordDto) SetExternalEventId(v string) {
	o.ExternalEventId = v
}

// GetExternalEventSubscriptionId returns the ExternalEventSubscriptionId field value
func (o *ExternalEventLogRecordDto) GetExternalEventSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalEventSubscriptionId
}

// GetExternalEventSubscriptionIdOk returns a tuple with the ExternalEventSubscriptionId field value
// and a boolean to check if the value has been set.
func (o *ExternalEventLogRecordDto) GetExternalEventSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalEventSubscriptionId, true
}

// SetExternalEventSubscriptionId sets field value
func (o *ExternalEventLogRecordDto) SetExternalEventSubscriptionId(v string) {
	o.ExternalEventSubscriptionId = v
}

// GetConnectionId returns the ConnectionId field value
func (o *ExternalEventLogRecordDto) GetConnectionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionId
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value
// and a boolean to check if the value has been set.
func (o *ExternalEventLogRecordDto) GetConnectionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionId, true
}

// SetConnectionId sets field value
func (o *ExternalEventLogRecordDto) SetConnectionId(v string) {
	o.ConnectionId = v
}

// GetConnection returns the Connection field value if set, zero value otherwise.
func (o *ExternalEventLogRecordDto) GetConnection() ConnectionDto {
	if o == nil || IsNil(o.Connection) {
		var ret ConnectionDto
		return ret
	}
	return *o.Connection
}

// GetConnectionOk returns a tuple with the Connection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalEventLogRecordDto) GetConnectionOk() (*ConnectionDto, bool) {
	if o == nil || IsNil(o.Connection) {
		return nil, false
	}
	return o.Connection, true
}

// HasConnection returns a boolean if a field has been set.
func (o *ExternalEventLogRecordDto) HasConnection() bool {
	if o != nil && !IsNil(o.Connection) {
		return true
	}

	return false
}

// SetConnection gets a reference to the given ConnectionDto and assigns it to the Connection field.
func (o *ExternalEventLogRecordDto) SetConnection(v ConnectionDto) {
	o.Connection = &v
}

// GetIntegrationId returns the IntegrationId field value
func (o *ExternalEventLogRecordDto) GetIntegrationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IntegrationId
}

// GetIntegrationIdOk returns a tuple with the IntegrationId field value
// and a boolean to check if the value has been set.
func (o *ExternalEventLogRecordDto) GetIntegrationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntegrationId, true
}

// SetIntegrationId sets field value
func (o *ExternalEventLogRecordDto) SetIntegrationId(v string) {
	o.IntegrationId = v
}

// GetIntegration returns the Integration field value if set, zero value otherwise.
func (o *ExternalEventLogRecordDto) GetIntegration() IntegrationDto {
	if o == nil || IsNil(o.Integration) {
		var ret IntegrationDto
		return ret
	}
	return *o.Integration
}

// GetIntegrationOk returns a tuple with the Integration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalEventLogRecordDto) GetIntegrationOk() (*IntegrationDto, bool) {
	if o == nil || IsNil(o.Integration) {
		return nil, false
	}
	return o.Integration, true
}

// HasIntegration returns a boolean if a field has been set.
func (o *ExternalEventLogRecordDto) HasIntegration() bool {
	if o != nil && !IsNil(o.Integration) {
		return true
	}

	return false
}

// SetIntegration gets a reference to the given IntegrationDto and assigns it to the Integration field.
func (o *ExternalEventLogRecordDto) SetIntegration(v IntegrationDto) {
	o.Integration = &v
}

// GetPayload returns the Payload field value
func (o *ExternalEventLogRecordDto) GetPayload() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value
// and a boolean to check if the value has been set.
func (o *ExternalEventLogRecordDto) GetPayloadOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Payload, true
}

// SetPayload sets field value
func (o *ExternalEventLogRecordDto) SetPayload(v map[string]interface{}) {
	o.Payload = v
}

// GetLaunchedFlowRunIds returns the LaunchedFlowRunIds field value
func (o *ExternalEventLogRecordDto) GetLaunchedFlowRunIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.LaunchedFlowRunIds
}

// GetLaunchedFlowRunIdsOk returns a tuple with the LaunchedFlowRunIds field value
// and a boolean to check if the value has been set.
func (o *ExternalEventLogRecordDto) GetLaunchedFlowRunIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LaunchedFlowRunIds, true
}

// SetLaunchedFlowRunIds sets field value
func (o *ExternalEventLogRecordDto) SetLaunchedFlowRunIds(v []string) {
	o.LaunchedFlowRunIds = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ExternalEventLogRecordDto) GetError() map[string]interface{} {
	if o == nil || IsNil(o.Error) {
		var ret map[string]interface{}
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalEventLogRecordDto) GetErrorOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Error) {
		return map[string]interface{}{}, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ExternalEventLogRecordDto) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given map[string]interface{} and assigns it to the Error field.
func (o *ExternalEventLogRecordDto) SetError(v map[string]interface{}) {
	o.Error = v
}

// GetLogs returns the Logs field value if set, zero value otherwise.
func (o *ExternalEventLogRecordDto) GetLogs() map[string]interface{} {
	if o == nil || IsNil(o.Logs) {
		var ret map[string]interface{}
		return ret
	}
	return o.Logs
}

// GetLogsOk returns a tuple with the Logs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalEventLogRecordDto) GetLogsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Logs) {
		return map[string]interface{}{}, false
	}
	return o.Logs, true
}

// HasLogs returns a boolean if a field has been set.
func (o *ExternalEventLogRecordDto) HasLogs() bool {
	if o != nil && !IsNil(o.Logs) {
		return true
	}

	return false
}

// SetLogs gets a reference to the given map[string]interface{} and assigns it to the Logs field.
func (o *ExternalEventLogRecordDto) SetLogs(v map[string]interface{}) {
	o.Logs = v
}

func (o ExternalEventLogRecordDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalEventLogRecordDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["userId"] = o.UserId
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	toSerialize["externalEventId"] = o.ExternalEventId
	toSerialize["externalEventSubscriptionId"] = o.ExternalEventSubscriptionId
	toSerialize["connectionId"] = o.ConnectionId
	if !IsNil(o.Connection) {
		toSerialize["connection"] = o.Connection
	}
	toSerialize["integrationId"] = o.IntegrationId
	if !IsNil(o.Integration) {
		toSerialize["integration"] = o.Integration
	}
	toSerialize["payload"] = o.Payload
	toSerialize["launchedFlowRunIds"] = o.LaunchedFlowRunIds
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.Logs) {
		toSerialize["logs"] = o.Logs
	}
	return toSerialize, nil
}

func (o *ExternalEventLogRecordDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"userId",
		"externalEventId",
		"externalEventSubscriptionId",
		"connectionId",
		"integrationId",
		"payload",
		"launchedFlowRunIds",
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

	varExternalEventLogRecordDto := _ExternalEventLogRecordDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExternalEventLogRecordDto)

	if err != nil {
		return err
	}

	*o = ExternalEventLogRecordDto(varExternalEventLogRecordDto)

	return err
}

type NullableExternalEventLogRecordDto struct {
	value *ExternalEventLogRecordDto
	isSet bool
}

func (v NullableExternalEventLogRecordDto) Get() *ExternalEventLogRecordDto {
	return v.value
}

func (v *NullableExternalEventLogRecordDto) Set(val *ExternalEventLogRecordDto) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalEventLogRecordDto) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalEventLogRecordDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalEventLogRecordDto(val *ExternalEventLogRecordDto) *NullableExternalEventLogRecordDto {
	return &NullableExternalEventLogRecordDto{value: val, isSet: true}
}

func (v NullableExternalEventLogRecordDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalEventLogRecordDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


