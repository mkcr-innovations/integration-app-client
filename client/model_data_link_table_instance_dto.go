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

// checks if the DataLinkTableInstanceDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataLinkTableInstanceDto{}

// DataLinkTableInstanceDto struct for DataLinkTableInstanceDto
type DataLinkTableInstanceDto struct {
	Id string `json:"id"`
	Name string `json:"name"`
	DataLinkTableId string `json:"dataLinkTableId"`
	DataLinkTable *DataLinkTableDto `json:"dataLinkTable,omitempty"`
	UserId string `json:"userId"`
	User *CustomerDto `json:"user,omitempty"`
	ConnectionId string `json:"connectionId"`
	Connection *ConnectionDto `json:"connection,omitempty"`
	IntegrationId string `json:"integrationId"`
	Integration *IntegrationDto `json:"integration,omitempty"`
	InstanceKey *string `json:"instanceKey,omitempty"`
}

type _DataLinkTableInstanceDto DataLinkTableInstanceDto

// NewDataLinkTableInstanceDto instantiates a new DataLinkTableInstanceDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataLinkTableInstanceDto(id string, name string, dataLinkTableId string, userId string, connectionId string, integrationId string) *DataLinkTableInstanceDto {
	this := DataLinkTableInstanceDto{}
	this.Id = id
	this.Name = name
	this.DataLinkTableId = dataLinkTableId
	this.UserId = userId
	this.ConnectionId = connectionId
	this.IntegrationId = integrationId
	return &this
}

// NewDataLinkTableInstanceDtoWithDefaults instantiates a new DataLinkTableInstanceDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataLinkTableInstanceDtoWithDefaults() *DataLinkTableInstanceDto {
	this := DataLinkTableInstanceDto{}
	return &this
}

// GetId returns the Id field value
func (o *DataLinkTableInstanceDto) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DataLinkTableInstanceDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DataLinkTableInstanceDto) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *DataLinkTableInstanceDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DataLinkTableInstanceDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DataLinkTableInstanceDto) SetName(v string) {
	o.Name = v
}

// GetDataLinkTableId returns the DataLinkTableId field value
func (o *DataLinkTableInstanceDto) GetDataLinkTableId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DataLinkTableId
}

// GetDataLinkTableIdOk returns a tuple with the DataLinkTableId field value
// and a boolean to check if the value has been set.
func (o *DataLinkTableInstanceDto) GetDataLinkTableIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataLinkTableId, true
}

// SetDataLinkTableId sets field value
func (o *DataLinkTableInstanceDto) SetDataLinkTableId(v string) {
	o.DataLinkTableId = v
}

// GetDataLinkTable returns the DataLinkTable field value if set, zero value otherwise.
func (o *DataLinkTableInstanceDto) GetDataLinkTable() DataLinkTableDto {
	if o == nil || IsNil(o.DataLinkTable) {
		var ret DataLinkTableDto
		return ret
	}
	return *o.DataLinkTable
}

// GetDataLinkTableOk returns a tuple with the DataLinkTable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataLinkTableInstanceDto) GetDataLinkTableOk() (*DataLinkTableDto, bool) {
	if o == nil || IsNil(o.DataLinkTable) {
		return nil, false
	}
	return o.DataLinkTable, true
}

// HasDataLinkTable returns a boolean if a field has been set.
func (o *DataLinkTableInstanceDto) HasDataLinkTable() bool {
	if o != nil && !IsNil(o.DataLinkTable) {
		return true
	}

	return false
}

// SetDataLinkTable gets a reference to the given DataLinkTableDto and assigns it to the DataLinkTable field.
func (o *DataLinkTableInstanceDto) SetDataLinkTable(v DataLinkTableDto) {
	o.DataLinkTable = &v
}

// GetUserId returns the UserId field value
func (o *DataLinkTableInstanceDto) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *DataLinkTableInstanceDto) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *DataLinkTableInstanceDto) SetUserId(v string) {
	o.UserId = v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *DataLinkTableInstanceDto) GetUser() CustomerDto {
	if o == nil || IsNil(o.User) {
		var ret CustomerDto
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataLinkTableInstanceDto) GetUserOk() (*CustomerDto, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *DataLinkTableInstanceDto) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given CustomerDto and assigns it to the User field.
func (o *DataLinkTableInstanceDto) SetUser(v CustomerDto) {
	o.User = &v
}

// GetConnectionId returns the ConnectionId field value
func (o *DataLinkTableInstanceDto) GetConnectionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionId
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value
// and a boolean to check if the value has been set.
func (o *DataLinkTableInstanceDto) GetConnectionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionId, true
}

// SetConnectionId sets field value
func (o *DataLinkTableInstanceDto) SetConnectionId(v string) {
	o.ConnectionId = v
}

// GetConnection returns the Connection field value if set, zero value otherwise.
func (o *DataLinkTableInstanceDto) GetConnection() ConnectionDto {
	if o == nil || IsNil(o.Connection) {
		var ret ConnectionDto
		return ret
	}
	return *o.Connection
}

// GetConnectionOk returns a tuple with the Connection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataLinkTableInstanceDto) GetConnectionOk() (*ConnectionDto, bool) {
	if o == nil || IsNil(o.Connection) {
		return nil, false
	}
	return o.Connection, true
}

// HasConnection returns a boolean if a field has been set.
func (o *DataLinkTableInstanceDto) HasConnection() bool {
	if o != nil && !IsNil(o.Connection) {
		return true
	}

	return false
}

// SetConnection gets a reference to the given ConnectionDto and assigns it to the Connection field.
func (o *DataLinkTableInstanceDto) SetConnection(v ConnectionDto) {
	o.Connection = &v
}

// GetIntegrationId returns the IntegrationId field value
func (o *DataLinkTableInstanceDto) GetIntegrationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IntegrationId
}

// GetIntegrationIdOk returns a tuple with the IntegrationId field value
// and a boolean to check if the value has been set.
func (o *DataLinkTableInstanceDto) GetIntegrationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntegrationId, true
}

// SetIntegrationId sets field value
func (o *DataLinkTableInstanceDto) SetIntegrationId(v string) {
	o.IntegrationId = v
}

// GetIntegration returns the Integration field value if set, zero value otherwise.
func (o *DataLinkTableInstanceDto) GetIntegration() IntegrationDto {
	if o == nil || IsNil(o.Integration) {
		var ret IntegrationDto
		return ret
	}
	return *o.Integration
}

// GetIntegrationOk returns a tuple with the Integration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataLinkTableInstanceDto) GetIntegrationOk() (*IntegrationDto, bool) {
	if o == nil || IsNil(o.Integration) {
		return nil, false
	}
	return o.Integration, true
}

// HasIntegration returns a boolean if a field has been set.
func (o *DataLinkTableInstanceDto) HasIntegration() bool {
	if o != nil && !IsNil(o.Integration) {
		return true
	}

	return false
}

// SetIntegration gets a reference to the given IntegrationDto and assigns it to the Integration field.
func (o *DataLinkTableInstanceDto) SetIntegration(v IntegrationDto) {
	o.Integration = &v
}

// GetInstanceKey returns the InstanceKey field value if set, zero value otherwise.
func (o *DataLinkTableInstanceDto) GetInstanceKey() string {
	if o == nil || IsNil(o.InstanceKey) {
		var ret string
		return ret
	}
	return *o.InstanceKey
}

// GetInstanceKeyOk returns a tuple with the InstanceKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataLinkTableInstanceDto) GetInstanceKeyOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceKey) {
		return nil, false
	}
	return o.InstanceKey, true
}

// HasInstanceKey returns a boolean if a field has been set.
func (o *DataLinkTableInstanceDto) HasInstanceKey() bool {
	if o != nil && !IsNil(o.InstanceKey) {
		return true
	}

	return false
}

// SetInstanceKey gets a reference to the given string and assigns it to the InstanceKey field.
func (o *DataLinkTableInstanceDto) SetInstanceKey(v string) {
	o.InstanceKey = &v
}

func (o DataLinkTableInstanceDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataLinkTableInstanceDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["dataLinkTableId"] = o.DataLinkTableId
	if !IsNil(o.DataLinkTable) {
		toSerialize["dataLinkTable"] = o.DataLinkTable
	}
	toSerialize["userId"] = o.UserId
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	toSerialize["connectionId"] = o.ConnectionId
	if !IsNil(o.Connection) {
		toSerialize["connection"] = o.Connection
	}
	toSerialize["integrationId"] = o.IntegrationId
	if !IsNil(o.Integration) {
		toSerialize["integration"] = o.Integration
	}
	if !IsNil(o.InstanceKey) {
		toSerialize["instanceKey"] = o.InstanceKey
	}
	return toSerialize, nil
}

func (o *DataLinkTableInstanceDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"dataLinkTableId",
		"userId",
		"connectionId",
		"integrationId",
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

	varDataLinkTableInstanceDto := _DataLinkTableInstanceDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDataLinkTableInstanceDto)

	if err != nil {
		return err
	}

	*o = DataLinkTableInstanceDto(varDataLinkTableInstanceDto)

	return err
}

type NullableDataLinkTableInstanceDto struct {
	value *DataLinkTableInstanceDto
	isSet bool
}

func (v NullableDataLinkTableInstanceDto) Get() *DataLinkTableInstanceDto {
	return v.value
}

func (v *NullableDataLinkTableInstanceDto) Set(val *DataLinkTableInstanceDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDataLinkTableInstanceDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDataLinkTableInstanceDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataLinkTableInstanceDto(val *DataLinkTableInstanceDto) *NullableDataLinkTableInstanceDto {
	return &NullableDataLinkTableInstanceDto{value: val, isSet: true}
}

func (v NullableDataLinkTableInstanceDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataLinkTableInstanceDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


