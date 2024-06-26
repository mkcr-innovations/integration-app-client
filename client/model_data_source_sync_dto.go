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

// checks if the DataSourceSyncDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataSourceSyncDto{}

// DataSourceSyncDto struct for DataSourceSyncDto
type DataSourceSyncDto struct {
	Id string `json:"id"`
	UserId string `json:"userId"`
	User *CustomerDto `json:"user,omitempty"`
	DataSourceId string `json:"dataSourceId"`
	DataSource *DataSourceDto `json:"dataSource,omitempty"`
	DataSourceInstanceId string `json:"dataSourceInstanceId"`
	DataSourceInstance *DataSourceInstanceDto `json:"dataSourceInstance,omitempty"`
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

type _DataSourceSyncDto DataSourceSyncDto

// NewDataSourceSyncDto instantiates a new DataSourceSyncDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataSourceSyncDto(id string, userId string, dataSourceId string, dataSourceInstanceId string, connectionId string, integrationId string, status string, startDatetime string, endDatetime string, isFullScan bool, collectedEventIds []string) *DataSourceSyncDto {
	this := DataSourceSyncDto{}
	this.Id = id
	this.UserId = userId
	this.DataSourceId = dataSourceId
	this.DataSourceInstanceId = dataSourceInstanceId
	this.ConnectionId = connectionId
	this.IntegrationId = integrationId
	this.Status = status
	this.StartDatetime = startDatetime
	this.EndDatetime = endDatetime
	this.IsFullScan = isFullScan
	this.CollectedEventIds = collectedEventIds
	return &this
}

// NewDataSourceSyncDtoWithDefaults instantiates a new DataSourceSyncDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataSourceSyncDtoWithDefaults() *DataSourceSyncDto {
	this := DataSourceSyncDto{}
	return &this
}

// GetId returns the Id field value
func (o *DataSourceSyncDto) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DataSourceSyncDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DataSourceSyncDto) SetId(v string) {
	o.Id = v
}

// GetUserId returns the UserId field value
func (o *DataSourceSyncDto) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *DataSourceSyncDto) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *DataSourceSyncDto) SetUserId(v string) {
	o.UserId = v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *DataSourceSyncDto) GetUser() CustomerDto {
	if o == nil || IsNil(o.User) {
		var ret CustomerDto
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourceSyncDto) GetUserOk() (*CustomerDto, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *DataSourceSyncDto) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given CustomerDto and assigns it to the User field.
func (o *DataSourceSyncDto) SetUser(v CustomerDto) {
	o.User = &v
}

// GetDataSourceId returns the DataSourceId field value
func (o *DataSourceSyncDto) GetDataSourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DataSourceId
}

// GetDataSourceIdOk returns a tuple with the DataSourceId field value
// and a boolean to check if the value has been set.
func (o *DataSourceSyncDto) GetDataSourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSourceId, true
}

// SetDataSourceId sets field value
func (o *DataSourceSyncDto) SetDataSourceId(v string) {
	o.DataSourceId = v
}

// GetDataSource returns the DataSource field value if set, zero value otherwise.
func (o *DataSourceSyncDto) GetDataSource() DataSourceDto {
	if o == nil || IsNil(o.DataSource) {
		var ret DataSourceDto
		return ret
	}
	return *o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourceSyncDto) GetDataSourceOk() (*DataSourceDto, bool) {
	if o == nil || IsNil(o.DataSource) {
		return nil, false
	}
	return o.DataSource, true
}

// HasDataSource returns a boolean if a field has been set.
func (o *DataSourceSyncDto) HasDataSource() bool {
	if o != nil && !IsNil(o.DataSource) {
		return true
	}

	return false
}

// SetDataSource gets a reference to the given DataSourceDto and assigns it to the DataSource field.
func (o *DataSourceSyncDto) SetDataSource(v DataSourceDto) {
	o.DataSource = &v
}

// GetDataSourceInstanceId returns the DataSourceInstanceId field value
func (o *DataSourceSyncDto) GetDataSourceInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DataSourceInstanceId
}

// GetDataSourceInstanceIdOk returns a tuple with the DataSourceInstanceId field value
// and a boolean to check if the value has been set.
func (o *DataSourceSyncDto) GetDataSourceInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSourceInstanceId, true
}

// SetDataSourceInstanceId sets field value
func (o *DataSourceSyncDto) SetDataSourceInstanceId(v string) {
	o.DataSourceInstanceId = v
}

// GetDataSourceInstance returns the DataSourceInstance field value if set, zero value otherwise.
func (o *DataSourceSyncDto) GetDataSourceInstance() DataSourceInstanceDto {
	if o == nil || IsNil(o.DataSourceInstance) {
		var ret DataSourceInstanceDto
		return ret
	}
	return *o.DataSourceInstance
}

// GetDataSourceInstanceOk returns a tuple with the DataSourceInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourceSyncDto) GetDataSourceInstanceOk() (*DataSourceInstanceDto, bool) {
	if o == nil || IsNil(o.DataSourceInstance) {
		return nil, false
	}
	return o.DataSourceInstance, true
}

// HasDataSourceInstance returns a boolean if a field has been set.
func (o *DataSourceSyncDto) HasDataSourceInstance() bool {
	if o != nil && !IsNil(o.DataSourceInstance) {
		return true
	}

	return false
}

// SetDataSourceInstance gets a reference to the given DataSourceInstanceDto and assigns it to the DataSourceInstance field.
func (o *DataSourceSyncDto) SetDataSourceInstance(v DataSourceInstanceDto) {
	o.DataSourceInstance = &v
}

// GetConnectionId returns the ConnectionId field value
func (o *DataSourceSyncDto) GetConnectionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionId
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value
// and a boolean to check if the value has been set.
func (o *DataSourceSyncDto) GetConnectionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionId, true
}

// SetConnectionId sets field value
func (o *DataSourceSyncDto) SetConnectionId(v string) {
	o.ConnectionId = v
}

// GetConnection returns the Connection field value if set, zero value otherwise.
func (o *DataSourceSyncDto) GetConnection() ConnectionDto {
	if o == nil || IsNil(o.Connection) {
		var ret ConnectionDto
		return ret
	}
	return *o.Connection
}

// GetConnectionOk returns a tuple with the Connection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourceSyncDto) GetConnectionOk() (*ConnectionDto, bool) {
	if o == nil || IsNil(o.Connection) {
		return nil, false
	}
	return o.Connection, true
}

// HasConnection returns a boolean if a field has been set.
func (o *DataSourceSyncDto) HasConnection() bool {
	if o != nil && !IsNil(o.Connection) {
		return true
	}

	return false
}

// SetConnection gets a reference to the given ConnectionDto and assigns it to the Connection field.
func (o *DataSourceSyncDto) SetConnection(v ConnectionDto) {
	o.Connection = &v
}

// GetIntegrationId returns the IntegrationId field value
func (o *DataSourceSyncDto) GetIntegrationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IntegrationId
}

// GetIntegrationIdOk returns a tuple with the IntegrationId field value
// and a boolean to check if the value has been set.
func (o *DataSourceSyncDto) GetIntegrationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntegrationId, true
}

// SetIntegrationId sets field value
func (o *DataSourceSyncDto) SetIntegrationId(v string) {
	o.IntegrationId = v
}

// GetIntegration returns the Integration field value if set, zero value otherwise.
func (o *DataSourceSyncDto) GetIntegration() IntegrationDto {
	if o == nil || IsNil(o.Integration) {
		var ret IntegrationDto
		return ret
	}
	return *o.Integration
}

// GetIntegrationOk returns a tuple with the Integration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourceSyncDto) GetIntegrationOk() (*IntegrationDto, bool) {
	if o == nil || IsNil(o.Integration) {
		return nil, false
	}
	return o.Integration, true
}

// HasIntegration returns a boolean if a field has been set.
func (o *DataSourceSyncDto) HasIntegration() bool {
	if o != nil && !IsNil(o.Integration) {
		return true
	}

	return false
}

// SetIntegration gets a reference to the given IntegrationDto and assigns it to the Integration field.
func (o *DataSourceSyncDto) SetIntegration(v IntegrationDto) {
	o.Integration = &v
}

// GetStatus returns the Status field value
func (o *DataSourceSyncDto) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DataSourceSyncDto) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *DataSourceSyncDto) SetStatus(v string) {
	o.Status = v
}

// GetStartDatetime returns the StartDatetime field value
func (o *DataSourceSyncDto) GetStartDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartDatetime
}

// GetStartDatetimeOk returns a tuple with the StartDatetime field value
// and a boolean to check if the value has been set.
func (o *DataSourceSyncDto) GetStartDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDatetime, true
}

// SetStartDatetime sets field value
func (o *DataSourceSyncDto) SetStartDatetime(v string) {
	o.StartDatetime = v
}

// GetEndDatetime returns the EndDatetime field value
func (o *DataSourceSyncDto) GetEndDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndDatetime
}

// GetEndDatetimeOk returns a tuple with the EndDatetime field value
// and a boolean to check if the value has been set.
func (o *DataSourceSyncDto) GetEndDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndDatetime, true
}

// SetEndDatetime sets field value
func (o *DataSourceSyncDto) SetEndDatetime(v string) {
	o.EndDatetime = v
}

// GetIsFullScan returns the IsFullScan field value
func (o *DataSourceSyncDto) GetIsFullScan() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsFullScan
}

// GetIsFullScanOk returns a tuple with the IsFullScan field value
// and a boolean to check if the value has been set.
func (o *DataSourceSyncDto) GetIsFullScanOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsFullScan, true
}

// SetIsFullScan sets field value
func (o *DataSourceSyncDto) SetIsFullScan(v bool) {
	o.IsFullScan = v
}

// GetCollectedEventIds returns the CollectedEventIds field value
func (o *DataSourceSyncDto) GetCollectedEventIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.CollectedEventIds
}

// GetCollectedEventIdsOk returns a tuple with the CollectedEventIds field value
// and a boolean to check if the value has been set.
func (o *DataSourceSyncDto) GetCollectedEventIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CollectedEventIds, true
}

// SetCollectedEventIds sets field value
func (o *DataSourceSyncDto) SetCollectedEventIds(v []string) {
	o.CollectedEventIds = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *DataSourceSyncDto) GetError() map[string]interface{} {
	if o == nil || IsNil(o.Error) {
		var ret map[string]interface{}
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourceSyncDto) GetErrorOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Error) {
		return map[string]interface{}{}, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *DataSourceSyncDto) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given map[string]interface{} and assigns it to the Error field.
func (o *DataSourceSyncDto) SetError(v map[string]interface{}) {
	o.Error = v
}

func (o DataSourceSyncDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataSourceSyncDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["userId"] = o.UserId
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	toSerialize["dataSourceId"] = o.DataSourceId
	if !IsNil(o.DataSource) {
		toSerialize["dataSource"] = o.DataSource
	}
	toSerialize["dataSourceInstanceId"] = o.DataSourceInstanceId
	if !IsNil(o.DataSourceInstance) {
		toSerialize["dataSourceInstance"] = o.DataSourceInstance
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

func (o *DataSourceSyncDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"userId",
		"dataSourceId",
		"dataSourceInstanceId",
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

	varDataSourceSyncDto := _DataSourceSyncDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDataSourceSyncDto)

	if err != nil {
		return err
	}

	*o = DataSourceSyncDto(varDataSourceSyncDto)

	return err
}

type NullableDataSourceSyncDto struct {
	value *DataSourceSyncDto
	isSet bool
}

func (v NullableDataSourceSyncDto) Get() *DataSourceSyncDto {
	return v.value
}

func (v *NullableDataSourceSyncDto) Set(val *DataSourceSyncDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDataSourceSyncDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDataSourceSyncDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataSourceSyncDto(val *DataSourceSyncDto) *NullableDataSourceSyncDto {
	return &NullableDataSourceSyncDto{value: val, isSet: true}
}

func (v NullableDataSourceSyncDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataSourceSyncDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


