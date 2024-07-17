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

// checks if the FlowRunDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FlowRunDto{}

// FlowRunDto struct for FlowRunDto
type FlowRunDto struct {
	Id string `json:"id"`
	FlowInstanceId string `json:"flowInstanceId"`
	UniversalFlowId *string `json:"universalFlowId,omitempty"`
	FlowInstance *FlowInstanceDto `json:"flowInstance,omitempty"`
	IntegrationId *string `json:"integrationId,omitempty"`
	Integration *IntegrationDto `json:"integration,omitempty"`
	ConnectionId *string `json:"connectionId,omitempty"`
	Connection *ConnectionDto `json:"connection,omitempty"`
	StartNodeKey string `json:"startNodeKey"`
	UserId string `json:"userId"`
	User *CustomerDto `json:"user,omitempty"`
	Input map[string]interface{} `json:"input,omitempty"`
	State string `json:"state"`
	StartTime string `json:"startTime"`
	EndTime *string `json:"endTime,omitempty"`
	Errors []map[string]interface{} `json:"errors,omitempty"`
	Nodes map[string]interface{} `json:"nodes"`
	LaunchedBy *FlorRunLaunchedByDto `json:"launchedBy,omitempty"`
}

type _FlowRunDto FlowRunDto

// NewFlowRunDto instantiates a new FlowRunDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlowRunDto(id string, flowInstanceId string, startNodeKey string, userId string, state string, startTime string, nodes map[string]interface{}) *FlowRunDto {
	this := FlowRunDto{}
	this.Id = id
	this.FlowInstanceId = flowInstanceId
	this.StartNodeKey = startNodeKey
	this.UserId = userId
	this.State = state
	this.StartTime = startTime
	this.Nodes = nodes
	return &this
}

// NewFlowRunDtoWithDefaults instantiates a new FlowRunDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlowRunDtoWithDefaults() *FlowRunDto {
	this := FlowRunDto{}
	return &this
}

// GetId returns the Id field value
func (o *FlowRunDto) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FlowRunDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FlowRunDto) SetId(v string) {
	o.Id = v
}

// GetFlowInstanceId returns the FlowInstanceId field value
func (o *FlowRunDto) GetFlowInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FlowInstanceId
}

// GetFlowInstanceIdOk returns a tuple with the FlowInstanceId field value
// and a boolean to check if the value has been set.
func (o *FlowRunDto) GetFlowInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FlowInstanceId, true
}

// SetFlowInstanceId sets field value
func (o *FlowRunDto) SetFlowInstanceId(v string) {
	o.FlowInstanceId = v
}

// GetUniversalFlowId returns the UniversalFlowId field value if set, zero value otherwise.
func (o *FlowRunDto) GetUniversalFlowId() string {
	if o == nil || IsNil(o.UniversalFlowId) {
		var ret string
		return ret
	}
	return *o.UniversalFlowId
}

// GetUniversalFlowIdOk returns a tuple with the UniversalFlowId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowRunDto) GetUniversalFlowIdOk() (*string, bool) {
	if o == nil || IsNil(o.UniversalFlowId) {
		return nil, false
	}
	return o.UniversalFlowId, true
}

// HasUniversalFlowId returns a boolean if a field has been set.
func (o *FlowRunDto) HasUniversalFlowId() bool {
	if o != nil && !IsNil(o.UniversalFlowId) {
		return true
	}

	return false
}

// SetUniversalFlowId gets a reference to the given string and assigns it to the UniversalFlowId field.
func (o *FlowRunDto) SetUniversalFlowId(v string) {
	o.UniversalFlowId = &v
}

// GetFlowInstance returns the FlowInstance field value if set, zero value otherwise.
func (o *FlowRunDto) GetFlowInstance() FlowInstanceDto {
	if o == nil || IsNil(o.FlowInstance) {
		var ret FlowInstanceDto
		return ret
	}
	return *o.FlowInstance
}

// GetFlowInstanceOk returns a tuple with the FlowInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowRunDto) GetFlowInstanceOk() (*FlowInstanceDto, bool) {
	if o == nil || IsNil(o.FlowInstance) {
		return nil, false
	}
	return o.FlowInstance, true
}

// HasFlowInstance returns a boolean if a field has been set.
func (o *FlowRunDto) HasFlowInstance() bool {
	if o != nil && !IsNil(o.FlowInstance) {
		return true
	}

	return false
}

// SetFlowInstance gets a reference to the given FlowInstanceDto and assigns it to the FlowInstance field.
func (o *FlowRunDto) SetFlowInstance(v FlowInstanceDto) {
	o.FlowInstance = &v
}

// GetIntegrationId returns the IntegrationId field value if set, zero value otherwise.
func (o *FlowRunDto) GetIntegrationId() string {
	if o == nil || IsNil(o.IntegrationId) {
		var ret string
		return ret
	}
	return *o.IntegrationId
}

// GetIntegrationIdOk returns a tuple with the IntegrationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowRunDto) GetIntegrationIdOk() (*string, bool) {
	if o == nil || IsNil(o.IntegrationId) {
		return nil, false
	}
	return o.IntegrationId, true
}

// HasIntegrationId returns a boolean if a field has been set.
func (o *FlowRunDto) HasIntegrationId() bool {
	if o != nil && !IsNil(o.IntegrationId) {
		return true
	}

	return false
}

// SetIntegrationId gets a reference to the given string and assigns it to the IntegrationId field.
func (o *FlowRunDto) SetIntegrationId(v string) {
	o.IntegrationId = &v
}

// GetIntegration returns the Integration field value if set, zero value otherwise.
func (o *FlowRunDto) GetIntegration() IntegrationDto {
	if o == nil || IsNil(o.Integration) {
		var ret IntegrationDto
		return ret
	}
	return *o.Integration
}

// GetIntegrationOk returns a tuple with the Integration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowRunDto) GetIntegrationOk() (*IntegrationDto, bool) {
	if o == nil || IsNil(o.Integration) {
		return nil, false
	}
	return o.Integration, true
}

// HasIntegration returns a boolean if a field has been set.
func (o *FlowRunDto) HasIntegration() bool {
	if o != nil && !IsNil(o.Integration) {
		return true
	}

	return false
}

// SetIntegration gets a reference to the given IntegrationDto and assigns it to the Integration field.
func (o *FlowRunDto) SetIntegration(v IntegrationDto) {
	o.Integration = &v
}

// GetConnectionId returns the ConnectionId field value if set, zero value otherwise.
func (o *FlowRunDto) GetConnectionId() string {
	if o == nil || IsNil(o.ConnectionId) {
		var ret string
		return ret
	}
	return *o.ConnectionId
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowRunDto) GetConnectionIdOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectionId) {
		return nil, false
	}
	return o.ConnectionId, true
}

// HasConnectionId returns a boolean if a field has been set.
func (o *FlowRunDto) HasConnectionId() bool {
	if o != nil && !IsNil(o.ConnectionId) {
		return true
	}

	return false
}

// SetConnectionId gets a reference to the given string and assigns it to the ConnectionId field.
func (o *FlowRunDto) SetConnectionId(v string) {
	o.ConnectionId = &v
}

// GetConnection returns the Connection field value if set, zero value otherwise.
func (o *FlowRunDto) GetConnection() ConnectionDto {
	if o == nil || IsNil(o.Connection) {
		var ret ConnectionDto
		return ret
	}
	return *o.Connection
}

// GetConnectionOk returns a tuple with the Connection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowRunDto) GetConnectionOk() (*ConnectionDto, bool) {
	if o == nil || IsNil(o.Connection) {
		return nil, false
	}
	return o.Connection, true
}

// HasConnection returns a boolean if a field has been set.
func (o *FlowRunDto) HasConnection() bool {
	if o != nil && !IsNil(o.Connection) {
		return true
	}

	return false
}

// SetConnection gets a reference to the given ConnectionDto and assigns it to the Connection field.
func (o *FlowRunDto) SetConnection(v ConnectionDto) {
	o.Connection = &v
}

// GetStartNodeKey returns the StartNodeKey field value
func (o *FlowRunDto) GetStartNodeKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartNodeKey
}

// GetStartNodeKeyOk returns a tuple with the StartNodeKey field value
// and a boolean to check if the value has been set.
func (o *FlowRunDto) GetStartNodeKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartNodeKey, true
}

// SetStartNodeKey sets field value
func (o *FlowRunDto) SetStartNodeKey(v string) {
	o.StartNodeKey = v
}

// GetUserId returns the UserId field value
func (o *FlowRunDto) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *FlowRunDto) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *FlowRunDto) SetUserId(v string) {
	o.UserId = v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *FlowRunDto) GetUser() CustomerDto {
	if o == nil || IsNil(o.User) {
		var ret CustomerDto
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowRunDto) GetUserOk() (*CustomerDto, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *FlowRunDto) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given CustomerDto and assigns it to the User field.
func (o *FlowRunDto) SetUser(v CustomerDto) {
	o.User = &v
}

// GetInput returns the Input field value if set, zero value otherwise.
func (o *FlowRunDto) GetInput() map[string]interface{} {
	if o == nil || IsNil(o.Input) {
		var ret map[string]interface{}
		return ret
	}
	return o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowRunDto) GetInputOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Input) {
		return map[string]interface{}{}, false
	}
	return o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *FlowRunDto) HasInput() bool {
	if o != nil && !IsNil(o.Input) {
		return true
	}

	return false
}

// SetInput gets a reference to the given map[string]interface{} and assigns it to the Input field.
func (o *FlowRunDto) SetInput(v map[string]interface{}) {
	o.Input = v
}

// GetState returns the State field value
func (o *FlowRunDto) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *FlowRunDto) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *FlowRunDto) SetState(v string) {
	o.State = v
}

// GetStartTime returns the StartTime field value
func (o *FlowRunDto) GetStartTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *FlowRunDto) GetStartTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *FlowRunDto) SetStartTime(v string) {
	o.StartTime = v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *FlowRunDto) GetEndTime() string {
	if o == nil || IsNil(o.EndTime) {
		var ret string
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowRunDto) GetEndTimeOk() (*string, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *FlowRunDto) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given string and assigns it to the EndTime field.
func (o *FlowRunDto) SetEndTime(v string) {
	o.EndTime = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *FlowRunDto) GetErrors() []map[string]interface{} {
	if o == nil || IsNil(o.Errors) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowRunDto) GetErrorsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *FlowRunDto) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []map[string]interface{} and assigns it to the Errors field.
func (o *FlowRunDto) SetErrors(v []map[string]interface{}) {
	o.Errors = v
}

// GetNodes returns the Nodes field value
func (o *FlowRunDto) GetNodes() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value
// and a boolean to check if the value has been set.
func (o *FlowRunDto) GetNodesOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Nodes, true
}

// SetNodes sets field value
func (o *FlowRunDto) SetNodes(v map[string]interface{}) {
	o.Nodes = v
}

// GetLaunchedBy returns the LaunchedBy field value if set, zero value otherwise.
func (o *FlowRunDto) GetLaunchedBy() FlorRunLaunchedByDto {
	if o == nil || IsNil(o.LaunchedBy) {
		var ret FlorRunLaunchedByDto
		return ret
	}
	return *o.LaunchedBy
}

// GetLaunchedByOk returns a tuple with the LaunchedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowRunDto) GetLaunchedByOk() (*FlorRunLaunchedByDto, bool) {
	if o == nil || IsNil(o.LaunchedBy) {
		return nil, false
	}
	return o.LaunchedBy, true
}

// HasLaunchedBy returns a boolean if a field has been set.
func (o *FlowRunDto) HasLaunchedBy() bool {
	if o != nil && !IsNil(o.LaunchedBy) {
		return true
	}

	return false
}

// SetLaunchedBy gets a reference to the given FlorRunLaunchedByDto and assigns it to the LaunchedBy field.
func (o *FlowRunDto) SetLaunchedBy(v FlorRunLaunchedByDto) {
	o.LaunchedBy = &v
}

func (o FlowRunDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlowRunDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["flowInstanceId"] = o.FlowInstanceId
	if !IsNil(o.UniversalFlowId) {
		toSerialize["universalFlowId"] = o.UniversalFlowId
	}
	if !IsNil(o.FlowInstance) {
		toSerialize["flowInstance"] = o.FlowInstance
	}
	if !IsNil(o.IntegrationId) {
		toSerialize["integrationId"] = o.IntegrationId
	}
	if !IsNil(o.Integration) {
		toSerialize["integration"] = o.Integration
	}
	if !IsNil(o.ConnectionId) {
		toSerialize["connectionId"] = o.ConnectionId
	}
	if !IsNil(o.Connection) {
		toSerialize["connection"] = o.Connection
	}
	toSerialize["startNodeKey"] = o.StartNodeKey
	toSerialize["userId"] = o.UserId
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !IsNil(o.Input) {
		toSerialize["input"] = o.Input
	}
	toSerialize["state"] = o.State
	toSerialize["startTime"] = o.StartTime
	if !IsNil(o.EndTime) {
		toSerialize["endTime"] = o.EndTime
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	toSerialize["nodes"] = o.Nodes
	if !IsNil(o.LaunchedBy) {
		toSerialize["launchedBy"] = o.LaunchedBy
	}
	return toSerialize, nil
}

func (o *FlowRunDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"flowInstanceId",
		"startNodeKey",
		"userId",
		"state",
		"startTime",
		"nodes",
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

	varFlowRunDto := _FlowRunDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFlowRunDto)

	if err != nil {
		return err
	}

	*o = FlowRunDto(varFlowRunDto)

	return err
}

type NullableFlowRunDto struct {
	value *FlowRunDto
	isSet bool
}

func (v NullableFlowRunDto) Get() *FlowRunDto {
	return v.value
}

func (v *NullableFlowRunDto) Set(val *FlowRunDto) {
	v.value = val
	v.isSet = true
}

func (v NullableFlowRunDto) IsSet() bool {
	return v.isSet
}

func (v *NullableFlowRunDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlowRunDto(val *FlowRunDto) *NullableFlowRunDto {
	return &NullableFlowRunDto{value: val, isSet: true}
}

func (v NullableFlowRunDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlowRunDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


