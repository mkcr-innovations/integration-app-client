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

// checks if the FieldMappingInstanceDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FieldMappingInstanceDto{}

// FieldMappingInstanceDto struct for FieldMappingInstanceDto
type FieldMappingInstanceDto struct {
	Id string `json:"id"`
	Name *string `json:"name,omitempty"`
	Revision *string `json:"revision,omitempty"`
	UserId *string `json:"userId,omitempty"`
	User *CustomerDto `json:"user,omitempty"`
	ConnectionId *string `json:"connectionId,omitempty"`
	Connection *ConnectionDto `json:"connection,omitempty"`
	IntegrationId *string `json:"integrationId,omitempty"`
	Integration *IntegrationDto `json:"integration,omitempty"`
	InstanceKey *string `json:"instanceKey,omitempty"`
	FieldMappingId *string `json:"fieldMappingId,omitempty"`
	FieldMappingRevision *string `json:"fieldMappingRevision,omitempty"`
	FieldMapping *FieldMappingDto `json:"fieldMapping,omitempty"`
	Direction *string `json:"direction,omitempty"`
	DataSourceInstanceId *string `json:"dataSourceInstanceId,omitempty"`
	DataSourceInstance map[string]interface{} `json:"dataSourceInstance,omitempty"`
	DataSourceSchema map[string]interface{} `json:"dataSourceSchema,omitempty"`
	AppSchema map[string]interface{} `json:"appSchema,omitempty"`
	ImportValue map[string]interface{} `json:"importValue,omitempty"`
	ExportValue map[string]interface{} `json:"exportValue,omitempty"`
	UnifiedImportValue map[string]interface{} `json:"unifiedImportValue,omitempty"`
	UnifiedExportValue map[string]interface{} `json:"unifiedExportValue,omitempty"`
	FrozenImportFields map[string]interface{} `json:"frozenImportFields,omitempty"`
	FrozenExportFields map[string]interface{} `json:"frozenExportFields,omitempty"`
	IsCustomized *bool `json:"isCustomized,omitempty"`
	Error map[string]interface{} `json:"error,omitempty"`
	// Deprecated
	ExternalSchema map[string]interface{} `json:"externalSchema,omitempty"`
}

type _FieldMappingInstanceDto FieldMappingInstanceDto

// NewFieldMappingInstanceDto instantiates a new FieldMappingInstanceDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFieldMappingInstanceDto(id string) *FieldMappingInstanceDto {
	this := FieldMappingInstanceDto{}
	this.Id = id
	return &this
}

// NewFieldMappingInstanceDtoWithDefaults instantiates a new FieldMappingInstanceDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFieldMappingInstanceDtoWithDefaults() *FieldMappingInstanceDto {
	this := FieldMappingInstanceDto{}
	return &this
}

// GetId returns the Id field value
func (o *FieldMappingInstanceDto) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FieldMappingInstanceDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FieldMappingInstanceDto) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FieldMappingInstanceDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldMappingInstanceDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FieldMappingInstanceDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FieldMappingInstanceDto) SetName(v string) {
	o.Name = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *FieldMappingInstanceDto) GetRevision() string {
	if o == nil || IsNil(o.Revision) {
		var ret string
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldMappingInstanceDto) GetRevisionOk() (*string, bool) {
	if o == nil || IsNil(o.Revision) {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *FieldMappingInstanceDto) HasRevision() bool {
	if o != nil && !IsNil(o.Revision) {
		return true
	}

	return false
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *FieldMappingInstanceDto) SetRevision(v string) {
	o.Revision = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *FieldMappingInstanceDto) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldMappingInstanceDto) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *FieldMappingInstanceDto) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *FieldMappingInstanceDto) SetUserId(v string) {
	o.UserId = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *FieldMappingInstanceDto) GetUser() CustomerDto {
	if o == nil || IsNil(o.User) {
		var ret CustomerDto
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldMappingInstanceDto) GetUserOk() (*CustomerDto, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *FieldMappingInstanceDto) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given CustomerDto and assigns it to the User field.
func (o *FieldMappingInstanceDto) SetUser(v CustomerDto) {
	o.User = &v
}

// GetConnectionId returns the ConnectionId field value if set, zero value otherwise.
func (o *FieldMappingInstanceDto) GetConnectionId() string {
	if o == nil || IsNil(o.ConnectionId) {
		var ret string
		return ret
	}
	return *o.ConnectionId
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldMappingInstanceDto) GetConnectionIdOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectionId) {
		return nil, false
	}
	return o.ConnectionId, true
}

// HasConnectionId returns a boolean if a field has been set.
func (o *FieldMappingInstanceDto) HasConnectionId() bool {
	if o != nil && !IsNil(o.ConnectionId) {
		return true
	}

	return false
}

// SetConnectionId gets a reference to the given string and assigns it to the ConnectionId field.
func (o *FieldMappingInstanceDto) SetConnectionId(v string) {
	o.ConnectionId = &v
}

// GetConnection returns the Connection field value if set, zero value otherwise.
func (o *FieldMappingInstanceDto) GetConnection() ConnectionDto {
	if o == nil || IsNil(o.Connection) {
		var ret ConnectionDto
		return ret
	}
	return *o.Connection
}

// GetConnectionOk returns a tuple with the Connection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldMappingInstanceDto) GetConnectionOk() (*ConnectionDto, bool) {
	if o == nil || IsNil(o.Connection) {
		return nil, false
	}
	return o.Connection, true
}

// HasConnection returns a boolean if a field has been set.
func (o *FieldMappingInstanceDto) HasConnection() bool {
	if o != nil && !IsNil(o.Connection) {
		return true
	}

	return false
}

// SetConnection gets a reference to the given ConnectionDto and assigns it to the Connection field.
func (o *FieldMappingInstanceDto) SetConnection(v ConnectionDto) {
	o.Connection = &v
}

// GetIntegrationId returns the IntegrationId field value if set, zero value otherwise.
func (o *FieldMappingInstanceDto) GetIntegrationId() string {
	if o == nil || IsNil(o.IntegrationId) {
		var ret string
		return ret
	}
	return *o.IntegrationId
}

// GetIntegrationIdOk returns a tuple with the IntegrationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldMappingInstanceDto) GetIntegrationIdOk() (*string, bool) {
	if o == nil || IsNil(o.IntegrationId) {
		return nil, false
	}
	return o.IntegrationId, true
}

// HasIntegrationId returns a boolean if a field has been set.
func (o *FieldMappingInstanceDto) HasIntegrationId() bool {
	if o != nil && !IsNil(o.IntegrationId) {
		return true
	}

	return false
}

// SetIntegrationId gets a reference to the given string and assigns it to the IntegrationId field.
func (o *FieldMappingInstanceDto) SetIntegrationId(v string) {
	o.IntegrationId = &v
}

// GetIntegration returns the Integration field value if set, zero value otherwise.
func (o *FieldMappingInstanceDto) GetIntegration() IntegrationDto {
	if o == nil || IsNil(o.Integration) {
		var ret IntegrationDto
		return ret
	}
	return *o.Integration
}

// GetIntegrationOk returns a tuple with the Integration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldMappingInstanceDto) GetIntegrationOk() (*IntegrationDto, bool) {
	if o == nil || IsNil(o.Integration) {
		return nil, false
	}
	return o.Integration, true
}

// HasIntegration returns a boolean if a field has been set.
func (o *FieldMappingInstanceDto) HasIntegration() bool {
	if o != nil && !IsNil(o.Integration) {
		return true
	}

	return false
}

// SetIntegration gets a reference to the given IntegrationDto and assigns it to the Integration field.
func (o *FieldMappingInstanceDto) SetIntegration(v IntegrationDto) {
	o.Integration = &v
}

// GetInstanceKey returns the InstanceKey field value if set, zero value otherwise.
func (o *FieldMappingInstanceDto) GetInstanceKey() string {
	if o == nil || IsNil(o.InstanceKey) {
		var ret string
		return ret
	}
	return *o.InstanceKey
}

// GetInstanceKeyOk returns a tuple with the InstanceKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldMappingInstanceDto) GetInstanceKeyOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceKey) {
		return nil, false
	}
	return o.InstanceKey, true
}

// HasInstanceKey returns a boolean if a field has been set.
func (o *FieldMappingInstanceDto) HasInstanceKey() bool {
	if o != nil && !IsNil(o.InstanceKey) {
		return true
	}

	return false
}

// SetInstanceKey gets a reference to the given string and assigns it to the InstanceKey field.
func (o *FieldMappingInstanceDto) SetInstanceKey(v string) {
	o.InstanceKey = &v
}

// GetFieldMappingId returns the FieldMappingId field value if set, zero value otherwise.
func (o *FieldMappingInstanceDto) GetFieldMappingId() string {
	if o == nil || IsNil(o.FieldMappingId) {
		var ret string
		return ret
	}
	return *o.FieldMappingId
}

// GetFieldMappingIdOk returns a tuple with the FieldMappingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldMappingInstanceDto) GetFieldMappingIdOk() (*string, bool) {
	if o == nil || IsNil(o.FieldMappingId) {
		return nil, false
	}
	return o.FieldMappingId, true
}

// HasFieldMappingId returns a boolean if a field has been set.
func (o *FieldMappingInstanceDto) HasFieldMappingId() bool {
	if o != nil && !IsNil(o.FieldMappingId) {
		return true
	}

	return false
}

// SetFieldMappingId gets a reference to the given string and assigns it to the FieldMappingId field.
func (o *FieldMappingInstanceDto) SetFieldMappingId(v string) {
	o.FieldMappingId = &v
}

// GetFieldMappingRevision returns the FieldMappingRevision field value if set, zero value otherwise.
func (o *FieldMappingInstanceDto) GetFieldMappingRevision() string {
	if o == nil || IsNil(o.FieldMappingRevision) {
		var ret string
		return ret
	}
	return *o.FieldMappingRevision
}

// GetFieldMappingRevisionOk returns a tuple with the FieldMappingRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldMappingInstanceDto) GetFieldMappingRevisionOk() (*string, bool) {
	if o == nil || IsNil(o.FieldMappingRevision) {
		return nil, false
	}
	return o.FieldMappingRevision, true
}

// HasFieldMappingRevision returns a boolean if a field has been set.
func (o *FieldMappingInstanceDto) HasFieldMappingRevision() bool {
	if o != nil && !IsNil(o.FieldMappingRevision) {
		return true
	}

	return false
}

// SetFieldMappingRevision gets a reference to the given string and assigns it to the FieldMappingRevision field.
func (o *FieldMappingInstanceDto) SetFieldMappingRevision(v string) {
	o.FieldMappingRevision = &v
}

// GetFieldMapping returns the FieldMapping field value if set, zero value otherwise.
func (o *FieldMappingInstanceDto) GetFieldMapping() FieldMappingDto {
	if o == nil || IsNil(o.FieldMapping) {
		var ret FieldMappingDto
		return ret
	}
	return *o.FieldMapping
}

// GetFieldMappingOk returns a tuple with the FieldMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldMappingInstanceDto) GetFieldMappingOk() (*FieldMappingDto, bool) {
	if o == nil || IsNil(o.FieldMapping) {
		return nil, false
	}
	return o.FieldMapping, true
}

// HasFieldMapping returns a boolean if a field has been set.
func (o *FieldMappingInstanceDto) HasFieldMapping() bool {
	if o != nil && !IsNil(o.FieldMapping) {
		return true
	}

	return false
}

// SetFieldMapping gets a reference to the given FieldMappingDto and assigns it to the FieldMapping field.
func (o *FieldMappingInstanceDto) SetFieldMapping(v FieldMappingDto) {
	o.FieldMapping = &v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *FieldMappingInstanceDto) GetDirection() string {
	if o == nil || IsNil(o.Direction) {
		var ret string
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldMappingInstanceDto) GetDirectionOk() (*string, bool) {
	if o == nil || IsNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *FieldMappingInstanceDto) HasDirection() bool {
	if o != nil && !IsNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given string and assigns it to the Direction field.
func (o *FieldMappingInstanceDto) SetDirection(v string) {
	o.Direction = &v
}

// GetDataSourceInstanceId returns the DataSourceInstanceId field value if set, zero value otherwise.
func (o *FieldMappingInstanceDto) GetDataSourceInstanceId() string {
	if o == nil || IsNil(o.DataSourceInstanceId) {
		var ret string
		return ret
	}
	return *o.DataSourceInstanceId
}

// GetDataSourceInstanceIdOk returns a tuple with the DataSourceInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldMappingInstanceDto) GetDataSourceInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.DataSourceInstanceId) {
		return nil, false
	}
	return o.DataSourceInstanceId, true
}

// HasDataSourceInstanceId returns a boolean if a field has been set.
func (o *FieldMappingInstanceDto) HasDataSourceInstanceId() bool {
	if o != nil && !IsNil(o.DataSourceInstanceId) {
		return true
	}

	return false
}

// SetDataSourceInstanceId gets a reference to the given string and assigns it to the DataSourceInstanceId field.
func (o *FieldMappingInstanceDto) SetDataSourceInstanceId(v string) {
	o.DataSourceInstanceId = &v
}

// GetDataSourceInstance returns the DataSourceInstance field value if set, zero value otherwise.
func (o *FieldMappingInstanceDto) GetDataSourceInstance() map[string]interface{} {
	if o == nil || IsNil(o.DataSourceInstance) {
		var ret map[string]interface{}
		return ret
	}
	return o.DataSourceInstance
}

// GetDataSourceInstanceOk returns a tuple with the DataSourceInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldMappingInstanceDto) GetDataSourceInstanceOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.DataSourceInstance) {
		return map[string]interface{}{}, false
	}
	return o.DataSourceInstance, true
}

// HasDataSourceInstance returns a boolean if a field has been set.
func (o *FieldMappingInstanceDto) HasDataSourceInstance() bool {
	if o != nil && !IsNil(o.DataSourceInstance) {
		return true
	}

	return false
}

// SetDataSourceInstance gets a reference to the given map[string]interface{} and assigns it to the DataSourceInstance field.
func (o *FieldMappingInstanceDto) SetDataSourceInstance(v map[string]interface{}) {
	o.DataSourceInstance = v
}

// GetDataSourceSchema returns the DataSourceSchema field value if set, zero value otherwise.
func (o *FieldMappingInstanceDto) GetDataSourceSchema() map[string]interface{} {
	if o == nil || IsNil(o.DataSourceSchema) {
		var ret map[string]interface{}
		return ret
	}
	return o.DataSourceSchema
}

// GetDataSourceSchemaOk returns a tuple with the DataSourceSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldMappingInstanceDto) GetDataSourceSchemaOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.DataSourceSchema) {
		return map[string]interface{}{}, false
	}
	return o.DataSourceSchema, true
}

// HasDataSourceSchema returns a boolean if a field has been set.
func (o *FieldMappingInstanceDto) HasDataSourceSchema() bool {
	if o != nil && !IsNil(o.DataSourceSchema) {
		return true
	}

	return false
}

// SetDataSourceSchema gets a reference to the given map[string]interface{} and assigns it to the DataSourceSchema field.
func (o *FieldMappingInstanceDto) SetDataSourceSchema(v map[string]interface{}) {
	o.DataSourceSchema = v
}

// GetAppSchema returns the AppSchema field value if set, zero value otherwise.
func (o *FieldMappingInstanceDto) GetAppSchema() map[string]interface{} {
	if o == nil || IsNil(o.AppSchema) {
		var ret map[string]interface{}
		return ret
	}
	return o.AppSchema
}

// GetAppSchemaOk returns a tuple with the AppSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldMappingInstanceDto) GetAppSchemaOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AppSchema) {
		return map[string]interface{}{}, false
	}
	return o.AppSchema, true
}

// HasAppSchema returns a boolean if a field has been set.
func (o *FieldMappingInstanceDto) HasAppSchema() bool {
	if o != nil && !IsNil(o.AppSchema) {
		return true
	}

	return false
}

// SetAppSchema gets a reference to the given map[string]interface{} and assigns it to the AppSchema field.
func (o *FieldMappingInstanceDto) SetAppSchema(v map[string]interface{}) {
	o.AppSchema = v
}

// GetImportValue returns the ImportValue field value if set, zero value otherwise.
func (o *FieldMappingInstanceDto) GetImportValue() map[string]interface{} {
	if o == nil || IsNil(o.ImportValue) {
		var ret map[string]interface{}
		return ret
	}
	return o.ImportValue
}

// GetImportValueOk returns a tuple with the ImportValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldMappingInstanceDto) GetImportValueOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ImportValue) {
		return map[string]interface{}{}, false
	}
	return o.ImportValue, true
}

// HasImportValue returns a boolean if a field has been set.
func (o *FieldMappingInstanceDto) HasImportValue() bool {
	if o != nil && !IsNil(o.ImportValue) {
		return true
	}

	return false
}

// SetImportValue gets a reference to the given map[string]interface{} and assigns it to the ImportValue field.
func (o *FieldMappingInstanceDto) SetImportValue(v map[string]interface{}) {
	o.ImportValue = v
}

// GetExportValue returns the ExportValue field value if set, zero value otherwise.
func (o *FieldMappingInstanceDto) GetExportValue() map[string]interface{} {
	if o == nil || IsNil(o.ExportValue) {
		var ret map[string]interface{}
		return ret
	}
	return o.ExportValue
}

// GetExportValueOk returns a tuple with the ExportValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldMappingInstanceDto) GetExportValueOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ExportValue) {
		return map[string]interface{}{}, false
	}
	return o.ExportValue, true
}

// HasExportValue returns a boolean if a field has been set.
func (o *FieldMappingInstanceDto) HasExportValue() bool {
	if o != nil && !IsNil(o.ExportValue) {
		return true
	}

	return false
}

// SetExportValue gets a reference to the given map[string]interface{} and assigns it to the ExportValue field.
func (o *FieldMappingInstanceDto) SetExportValue(v map[string]interface{}) {
	o.ExportValue = v
}

// GetUnifiedImportValue returns the UnifiedImportValue field value if set, zero value otherwise.
func (o *FieldMappingInstanceDto) GetUnifiedImportValue() map[string]interface{} {
	if o == nil || IsNil(o.UnifiedImportValue) {
		var ret map[string]interface{}
		return ret
	}
	return o.UnifiedImportValue
}

// GetUnifiedImportValueOk returns a tuple with the UnifiedImportValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldMappingInstanceDto) GetUnifiedImportValueOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.UnifiedImportValue) {
		return map[string]interface{}{}, false
	}
	return o.UnifiedImportValue, true
}

// HasUnifiedImportValue returns a boolean if a field has been set.
func (o *FieldMappingInstanceDto) HasUnifiedImportValue() bool {
	if o != nil && !IsNil(o.UnifiedImportValue) {
		return true
	}

	return false
}

// SetUnifiedImportValue gets a reference to the given map[string]interface{} and assigns it to the UnifiedImportValue field.
func (o *FieldMappingInstanceDto) SetUnifiedImportValue(v map[string]interface{}) {
	o.UnifiedImportValue = v
}

// GetUnifiedExportValue returns the UnifiedExportValue field value if set, zero value otherwise.
func (o *FieldMappingInstanceDto) GetUnifiedExportValue() map[string]interface{} {
	if o == nil || IsNil(o.UnifiedExportValue) {
		var ret map[string]interface{}
		return ret
	}
	return o.UnifiedExportValue
}

// GetUnifiedExportValueOk returns a tuple with the UnifiedExportValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldMappingInstanceDto) GetUnifiedExportValueOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.UnifiedExportValue) {
		return map[string]interface{}{}, false
	}
	return o.UnifiedExportValue, true
}

// HasUnifiedExportValue returns a boolean if a field has been set.
func (o *FieldMappingInstanceDto) HasUnifiedExportValue() bool {
	if o != nil && !IsNil(o.UnifiedExportValue) {
		return true
	}

	return false
}

// SetUnifiedExportValue gets a reference to the given map[string]interface{} and assigns it to the UnifiedExportValue field.
func (o *FieldMappingInstanceDto) SetUnifiedExportValue(v map[string]interface{}) {
	o.UnifiedExportValue = v
}

// GetFrozenImportFields returns the FrozenImportFields field value if set, zero value otherwise.
func (o *FieldMappingInstanceDto) GetFrozenImportFields() map[string]interface{} {
	if o == nil || IsNil(o.FrozenImportFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.FrozenImportFields
}

// GetFrozenImportFieldsOk returns a tuple with the FrozenImportFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldMappingInstanceDto) GetFrozenImportFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.FrozenImportFields) {
		return map[string]interface{}{}, false
	}
	return o.FrozenImportFields, true
}

// HasFrozenImportFields returns a boolean if a field has been set.
func (o *FieldMappingInstanceDto) HasFrozenImportFields() bool {
	if o != nil && !IsNil(o.FrozenImportFields) {
		return true
	}

	return false
}

// SetFrozenImportFields gets a reference to the given map[string]interface{} and assigns it to the FrozenImportFields field.
func (o *FieldMappingInstanceDto) SetFrozenImportFields(v map[string]interface{}) {
	o.FrozenImportFields = v
}

// GetFrozenExportFields returns the FrozenExportFields field value if set, zero value otherwise.
func (o *FieldMappingInstanceDto) GetFrozenExportFields() map[string]interface{} {
	if o == nil || IsNil(o.FrozenExportFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.FrozenExportFields
}

// GetFrozenExportFieldsOk returns a tuple with the FrozenExportFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldMappingInstanceDto) GetFrozenExportFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.FrozenExportFields) {
		return map[string]interface{}{}, false
	}
	return o.FrozenExportFields, true
}

// HasFrozenExportFields returns a boolean if a field has been set.
func (o *FieldMappingInstanceDto) HasFrozenExportFields() bool {
	if o != nil && !IsNil(o.FrozenExportFields) {
		return true
	}

	return false
}

// SetFrozenExportFields gets a reference to the given map[string]interface{} and assigns it to the FrozenExportFields field.
func (o *FieldMappingInstanceDto) SetFrozenExportFields(v map[string]interface{}) {
	o.FrozenExportFields = v
}

// GetIsCustomized returns the IsCustomized field value if set, zero value otherwise.
func (o *FieldMappingInstanceDto) GetIsCustomized() bool {
	if o == nil || IsNil(o.IsCustomized) {
		var ret bool
		return ret
	}
	return *o.IsCustomized
}

// GetIsCustomizedOk returns a tuple with the IsCustomized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldMappingInstanceDto) GetIsCustomizedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCustomized) {
		return nil, false
	}
	return o.IsCustomized, true
}

// HasIsCustomized returns a boolean if a field has been set.
func (o *FieldMappingInstanceDto) HasIsCustomized() bool {
	if o != nil && !IsNil(o.IsCustomized) {
		return true
	}

	return false
}

// SetIsCustomized gets a reference to the given bool and assigns it to the IsCustomized field.
func (o *FieldMappingInstanceDto) SetIsCustomized(v bool) {
	o.IsCustomized = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *FieldMappingInstanceDto) GetError() map[string]interface{} {
	if o == nil || IsNil(o.Error) {
		var ret map[string]interface{}
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldMappingInstanceDto) GetErrorOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Error) {
		return map[string]interface{}{}, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *FieldMappingInstanceDto) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given map[string]interface{} and assigns it to the Error field.
func (o *FieldMappingInstanceDto) SetError(v map[string]interface{}) {
	o.Error = v
}

// GetExternalSchema returns the ExternalSchema field value if set, zero value otherwise.
// Deprecated
func (o *FieldMappingInstanceDto) GetExternalSchema() map[string]interface{} {
	if o == nil || IsNil(o.ExternalSchema) {
		var ret map[string]interface{}
		return ret
	}
	return o.ExternalSchema
}

// GetExternalSchemaOk returns a tuple with the ExternalSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *FieldMappingInstanceDto) GetExternalSchemaOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ExternalSchema) {
		return map[string]interface{}{}, false
	}
	return o.ExternalSchema, true
}

// HasExternalSchema returns a boolean if a field has been set.
func (o *FieldMappingInstanceDto) HasExternalSchema() bool {
	if o != nil && !IsNil(o.ExternalSchema) {
		return true
	}

	return false
}

// SetExternalSchema gets a reference to the given map[string]interface{} and assigns it to the ExternalSchema field.
// Deprecated
func (o *FieldMappingInstanceDto) SetExternalSchema(v map[string]interface{}) {
	o.ExternalSchema = v
}

func (o FieldMappingInstanceDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FieldMappingInstanceDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Revision) {
		toSerialize["revision"] = o.Revision
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !IsNil(o.ConnectionId) {
		toSerialize["connectionId"] = o.ConnectionId
	}
	if !IsNil(o.Connection) {
		toSerialize["connection"] = o.Connection
	}
	if !IsNil(o.IntegrationId) {
		toSerialize["integrationId"] = o.IntegrationId
	}
	if !IsNil(o.Integration) {
		toSerialize["integration"] = o.Integration
	}
	if !IsNil(o.InstanceKey) {
		toSerialize["instanceKey"] = o.InstanceKey
	}
	if !IsNil(o.FieldMappingId) {
		toSerialize["fieldMappingId"] = o.FieldMappingId
	}
	if !IsNil(o.FieldMappingRevision) {
		toSerialize["fieldMappingRevision"] = o.FieldMappingRevision
	}
	if !IsNil(o.FieldMapping) {
		toSerialize["fieldMapping"] = o.FieldMapping
	}
	if !IsNil(o.Direction) {
		toSerialize["direction"] = o.Direction
	}
	if !IsNil(o.DataSourceInstanceId) {
		toSerialize["dataSourceInstanceId"] = o.DataSourceInstanceId
	}
	if !IsNil(o.DataSourceInstance) {
		toSerialize["dataSourceInstance"] = o.DataSourceInstance
	}
	if !IsNil(o.DataSourceSchema) {
		toSerialize["dataSourceSchema"] = o.DataSourceSchema
	}
	if !IsNil(o.AppSchema) {
		toSerialize["appSchema"] = o.AppSchema
	}
	if !IsNil(o.ImportValue) {
		toSerialize["importValue"] = o.ImportValue
	}
	if !IsNil(o.ExportValue) {
		toSerialize["exportValue"] = o.ExportValue
	}
	if !IsNil(o.UnifiedImportValue) {
		toSerialize["unifiedImportValue"] = o.UnifiedImportValue
	}
	if !IsNil(o.UnifiedExportValue) {
		toSerialize["unifiedExportValue"] = o.UnifiedExportValue
	}
	if !IsNil(o.FrozenImportFields) {
		toSerialize["frozenImportFields"] = o.FrozenImportFields
	}
	if !IsNil(o.FrozenExportFields) {
		toSerialize["frozenExportFields"] = o.FrozenExportFields
	}
	if !IsNil(o.IsCustomized) {
		toSerialize["isCustomized"] = o.IsCustomized
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.ExternalSchema) {
		toSerialize["externalSchema"] = o.ExternalSchema
	}
	return toSerialize, nil
}

func (o *FieldMappingInstanceDto) UnmarshalJSON(data []byte) (err error) {
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

	varFieldMappingInstanceDto := _FieldMappingInstanceDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFieldMappingInstanceDto)

	if err != nil {
		return err
	}

	*o = FieldMappingInstanceDto(varFieldMappingInstanceDto)

	return err
}

type NullableFieldMappingInstanceDto struct {
	value *FieldMappingInstanceDto
	isSet bool
}

func (v NullableFieldMappingInstanceDto) Get() *FieldMappingInstanceDto {
	return v.value
}

func (v *NullableFieldMappingInstanceDto) Set(val *FieldMappingInstanceDto) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldMappingInstanceDto) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldMappingInstanceDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldMappingInstanceDto(val *FieldMappingInstanceDto) *NullableFieldMappingInstanceDto {
	return &NullableFieldMappingInstanceDto{value: val, isSet: true}
}

func (v NullableFieldMappingInstanceDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldMappingInstanceDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


