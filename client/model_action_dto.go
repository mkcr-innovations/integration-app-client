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

// checks if the ActionDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActionDto{}

// ActionDto struct for ActionDto
type ActionDto struct {
	Id string `json:"id"`
	Key *string `json:"key,omitempty"`
	Name *string `json:"name,omitempty"`
	ArchivedAt *string `json:"archivedAt,omitempty"`
	Revision *string `json:"revision,omitempty"`
	IsCustomized *bool `json:"isCustomized,omitempty"`
	ParentId *string `json:"parentId,omitempty"`
	ParentRevision *string `json:"parentRevision,omitempty"`
	IntegrationId *string `json:"integrationId,omitempty"`
	Integration *IntegrationDto `json:"integration,omitempty"`
	Type *string `json:"type,omitempty"`
	InputSchema map[string]interface{} `json:"inputSchema,omitempty"`
	Config map[string]interface{} `json:"config,omitempty"`
	OutputMapping map[string]interface{} `json:"outputMapping,omitempty"`
	CustomOutputSchema map[string]interface{} `json:"customOutputSchema,omitempty"`
	DefaultOutputSchema map[string]interface{} `json:"defaultOutputSchema,omitempty"`
	TransformedOutputSchema map[string]interface{} `json:"transformedOutputSchema,omitempty"`
	OutputSchema map[string]interface{} `json:"outputSchema,omitempty"`
	Dependencies []map[string]interface{} `json:"dependencies,omitempty"`
	IsDeployed *bool `json:"isDeployed,omitempty"`
	ConfigurationState *string `json:"configurationState,omitempty"`
	ConfigurationStateMessage *string `json:"configurationStateMessage,omitempty"`
	// 
	AppliedToIntegrations []ActionWithIntegrationDto `json:"appliedToIntegrations,omitempty"`
}

type _ActionDto ActionDto

// NewActionDto instantiates a new ActionDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionDto(id string) *ActionDto {
	this := ActionDto{}
	this.Id = id
	return &this
}

// NewActionDtoWithDefaults instantiates a new ActionDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionDtoWithDefaults() *ActionDto {
	this := ActionDto{}
	return &this
}

// GetId returns the Id field value
func (o *ActionDto) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ActionDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ActionDto) SetId(v string) {
	o.Id = v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *ActionDto) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionDto) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ActionDto) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *ActionDto) SetKey(v string) {
	o.Key = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ActionDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ActionDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ActionDto) SetName(v string) {
	o.Name = &v
}

// GetArchivedAt returns the ArchivedAt field value if set, zero value otherwise.
func (o *ActionDto) GetArchivedAt() string {
	if o == nil || IsNil(o.ArchivedAt) {
		var ret string
		return ret
	}
	return *o.ArchivedAt
}

// GetArchivedAtOk returns a tuple with the ArchivedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionDto) GetArchivedAtOk() (*string, bool) {
	if o == nil || IsNil(o.ArchivedAt) {
		return nil, false
	}
	return o.ArchivedAt, true
}

// HasArchivedAt returns a boolean if a field has been set.
func (o *ActionDto) HasArchivedAt() bool {
	if o != nil && !IsNil(o.ArchivedAt) {
		return true
	}

	return false
}

// SetArchivedAt gets a reference to the given string and assigns it to the ArchivedAt field.
func (o *ActionDto) SetArchivedAt(v string) {
	o.ArchivedAt = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *ActionDto) GetRevision() string {
	if o == nil || IsNil(o.Revision) {
		var ret string
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionDto) GetRevisionOk() (*string, bool) {
	if o == nil || IsNil(o.Revision) {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *ActionDto) HasRevision() bool {
	if o != nil && !IsNil(o.Revision) {
		return true
	}

	return false
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *ActionDto) SetRevision(v string) {
	o.Revision = &v
}

// GetIsCustomized returns the IsCustomized field value if set, zero value otherwise.
func (o *ActionDto) GetIsCustomized() bool {
	if o == nil || IsNil(o.IsCustomized) {
		var ret bool
		return ret
	}
	return *o.IsCustomized
}

// GetIsCustomizedOk returns a tuple with the IsCustomized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionDto) GetIsCustomizedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCustomized) {
		return nil, false
	}
	return o.IsCustomized, true
}

// HasIsCustomized returns a boolean if a field has been set.
func (o *ActionDto) HasIsCustomized() bool {
	if o != nil && !IsNil(o.IsCustomized) {
		return true
	}

	return false
}

// SetIsCustomized gets a reference to the given bool and assigns it to the IsCustomized field.
func (o *ActionDto) SetIsCustomized(v bool) {
	o.IsCustomized = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *ActionDto) GetParentId() string {
	if o == nil || IsNil(o.ParentId) {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionDto) GetParentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentId) {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *ActionDto) HasParentId() bool {
	if o != nil && !IsNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *ActionDto) SetParentId(v string) {
	o.ParentId = &v
}

// GetParentRevision returns the ParentRevision field value if set, zero value otherwise.
func (o *ActionDto) GetParentRevision() string {
	if o == nil || IsNil(o.ParentRevision) {
		var ret string
		return ret
	}
	return *o.ParentRevision
}

// GetParentRevisionOk returns a tuple with the ParentRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionDto) GetParentRevisionOk() (*string, bool) {
	if o == nil || IsNil(o.ParentRevision) {
		return nil, false
	}
	return o.ParentRevision, true
}

// HasParentRevision returns a boolean if a field has been set.
func (o *ActionDto) HasParentRevision() bool {
	if o != nil && !IsNil(o.ParentRevision) {
		return true
	}

	return false
}

// SetParentRevision gets a reference to the given string and assigns it to the ParentRevision field.
func (o *ActionDto) SetParentRevision(v string) {
	o.ParentRevision = &v
}

// GetIntegrationId returns the IntegrationId field value if set, zero value otherwise.
func (o *ActionDto) GetIntegrationId() string {
	if o == nil || IsNil(o.IntegrationId) {
		var ret string
		return ret
	}
	return *o.IntegrationId
}

// GetIntegrationIdOk returns a tuple with the IntegrationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionDto) GetIntegrationIdOk() (*string, bool) {
	if o == nil || IsNil(o.IntegrationId) {
		return nil, false
	}
	return o.IntegrationId, true
}

// HasIntegrationId returns a boolean if a field has been set.
func (o *ActionDto) HasIntegrationId() bool {
	if o != nil && !IsNil(o.IntegrationId) {
		return true
	}

	return false
}

// SetIntegrationId gets a reference to the given string and assigns it to the IntegrationId field.
func (o *ActionDto) SetIntegrationId(v string) {
	o.IntegrationId = &v
}

// GetIntegration returns the Integration field value if set, zero value otherwise.
func (o *ActionDto) GetIntegration() IntegrationDto {
	if o == nil || IsNil(o.Integration) {
		var ret IntegrationDto
		return ret
	}
	return *o.Integration
}

// GetIntegrationOk returns a tuple with the Integration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionDto) GetIntegrationOk() (*IntegrationDto, bool) {
	if o == nil || IsNil(o.Integration) {
		return nil, false
	}
	return o.Integration, true
}

// HasIntegration returns a boolean if a field has been set.
func (o *ActionDto) HasIntegration() bool {
	if o != nil && !IsNil(o.Integration) {
		return true
	}

	return false
}

// SetIntegration gets a reference to the given IntegrationDto and assigns it to the Integration field.
func (o *ActionDto) SetIntegration(v IntegrationDto) {
	o.Integration = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ActionDto) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionDto) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ActionDto) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ActionDto) SetType(v string) {
	o.Type = &v
}

// GetInputSchema returns the InputSchema field value if set, zero value otherwise.
func (o *ActionDto) GetInputSchema() map[string]interface{} {
	if o == nil || IsNil(o.InputSchema) {
		var ret map[string]interface{}
		return ret
	}
	return o.InputSchema
}

// GetInputSchemaOk returns a tuple with the InputSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionDto) GetInputSchemaOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.InputSchema) {
		return map[string]interface{}{}, false
	}
	return o.InputSchema, true
}

// HasInputSchema returns a boolean if a field has been set.
func (o *ActionDto) HasInputSchema() bool {
	if o != nil && !IsNil(o.InputSchema) {
		return true
	}

	return false
}

// SetInputSchema gets a reference to the given map[string]interface{} and assigns it to the InputSchema field.
func (o *ActionDto) SetInputSchema(v map[string]interface{}) {
	o.InputSchema = v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *ActionDto) GetConfig() map[string]interface{} {
	if o == nil || IsNil(o.Config) {
		var ret map[string]interface{}
		return ret
	}
	return o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionDto) GetConfigOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Config) {
		return map[string]interface{}{}, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *ActionDto) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given map[string]interface{} and assigns it to the Config field.
func (o *ActionDto) SetConfig(v map[string]interface{}) {
	o.Config = v
}

// GetOutputMapping returns the OutputMapping field value if set, zero value otherwise.
func (o *ActionDto) GetOutputMapping() map[string]interface{} {
	if o == nil || IsNil(o.OutputMapping) {
		var ret map[string]interface{}
		return ret
	}
	return o.OutputMapping
}

// GetOutputMappingOk returns a tuple with the OutputMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionDto) GetOutputMappingOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.OutputMapping) {
		return map[string]interface{}{}, false
	}
	return o.OutputMapping, true
}

// HasOutputMapping returns a boolean if a field has been set.
func (o *ActionDto) HasOutputMapping() bool {
	if o != nil && !IsNil(o.OutputMapping) {
		return true
	}

	return false
}

// SetOutputMapping gets a reference to the given map[string]interface{} and assigns it to the OutputMapping field.
func (o *ActionDto) SetOutputMapping(v map[string]interface{}) {
	o.OutputMapping = v
}

// GetCustomOutputSchema returns the CustomOutputSchema field value if set, zero value otherwise.
func (o *ActionDto) GetCustomOutputSchema() map[string]interface{} {
	if o == nil || IsNil(o.CustomOutputSchema) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomOutputSchema
}

// GetCustomOutputSchemaOk returns a tuple with the CustomOutputSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionDto) GetCustomOutputSchemaOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomOutputSchema) {
		return map[string]interface{}{}, false
	}
	return o.CustomOutputSchema, true
}

// HasCustomOutputSchema returns a boolean if a field has been set.
func (o *ActionDto) HasCustomOutputSchema() bool {
	if o != nil && !IsNil(o.CustomOutputSchema) {
		return true
	}

	return false
}

// SetCustomOutputSchema gets a reference to the given map[string]interface{} and assigns it to the CustomOutputSchema field.
func (o *ActionDto) SetCustomOutputSchema(v map[string]interface{}) {
	o.CustomOutputSchema = v
}

// GetDefaultOutputSchema returns the DefaultOutputSchema field value if set, zero value otherwise.
func (o *ActionDto) GetDefaultOutputSchema() map[string]interface{} {
	if o == nil || IsNil(o.DefaultOutputSchema) {
		var ret map[string]interface{}
		return ret
	}
	return o.DefaultOutputSchema
}

// GetDefaultOutputSchemaOk returns a tuple with the DefaultOutputSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionDto) GetDefaultOutputSchemaOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.DefaultOutputSchema) {
		return map[string]interface{}{}, false
	}
	return o.DefaultOutputSchema, true
}

// HasDefaultOutputSchema returns a boolean if a field has been set.
func (o *ActionDto) HasDefaultOutputSchema() bool {
	if o != nil && !IsNil(o.DefaultOutputSchema) {
		return true
	}

	return false
}

// SetDefaultOutputSchema gets a reference to the given map[string]interface{} and assigns it to the DefaultOutputSchema field.
func (o *ActionDto) SetDefaultOutputSchema(v map[string]interface{}) {
	o.DefaultOutputSchema = v
}

// GetTransformedOutputSchema returns the TransformedOutputSchema field value if set, zero value otherwise.
func (o *ActionDto) GetTransformedOutputSchema() map[string]interface{} {
	if o == nil || IsNil(o.TransformedOutputSchema) {
		var ret map[string]interface{}
		return ret
	}
	return o.TransformedOutputSchema
}

// GetTransformedOutputSchemaOk returns a tuple with the TransformedOutputSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionDto) GetTransformedOutputSchemaOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.TransformedOutputSchema) {
		return map[string]interface{}{}, false
	}
	return o.TransformedOutputSchema, true
}

// HasTransformedOutputSchema returns a boolean if a field has been set.
func (o *ActionDto) HasTransformedOutputSchema() bool {
	if o != nil && !IsNil(o.TransformedOutputSchema) {
		return true
	}

	return false
}

// SetTransformedOutputSchema gets a reference to the given map[string]interface{} and assigns it to the TransformedOutputSchema field.
func (o *ActionDto) SetTransformedOutputSchema(v map[string]interface{}) {
	o.TransformedOutputSchema = v
}

// GetOutputSchema returns the OutputSchema field value if set, zero value otherwise.
func (o *ActionDto) GetOutputSchema() map[string]interface{} {
	if o == nil || IsNil(o.OutputSchema) {
		var ret map[string]interface{}
		return ret
	}
	return o.OutputSchema
}

// GetOutputSchemaOk returns a tuple with the OutputSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionDto) GetOutputSchemaOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.OutputSchema) {
		return map[string]interface{}{}, false
	}
	return o.OutputSchema, true
}

// HasOutputSchema returns a boolean if a field has been set.
func (o *ActionDto) HasOutputSchema() bool {
	if o != nil && !IsNil(o.OutputSchema) {
		return true
	}

	return false
}

// SetOutputSchema gets a reference to the given map[string]interface{} and assigns it to the OutputSchema field.
func (o *ActionDto) SetOutputSchema(v map[string]interface{}) {
	o.OutputSchema = v
}

// GetDependencies returns the Dependencies field value if set, zero value otherwise.
func (o *ActionDto) GetDependencies() []map[string]interface{} {
	if o == nil || IsNil(o.Dependencies) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Dependencies
}

// GetDependenciesOk returns a tuple with the Dependencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionDto) GetDependenciesOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Dependencies) {
		return nil, false
	}
	return o.Dependencies, true
}

// HasDependencies returns a boolean if a field has been set.
func (o *ActionDto) HasDependencies() bool {
	if o != nil && !IsNil(o.Dependencies) {
		return true
	}

	return false
}

// SetDependencies gets a reference to the given []map[string]interface{} and assigns it to the Dependencies field.
func (o *ActionDto) SetDependencies(v []map[string]interface{}) {
	o.Dependencies = v
}

// GetIsDeployed returns the IsDeployed field value if set, zero value otherwise.
func (o *ActionDto) GetIsDeployed() bool {
	if o == nil || IsNil(o.IsDeployed) {
		var ret bool
		return ret
	}
	return *o.IsDeployed
}

// GetIsDeployedOk returns a tuple with the IsDeployed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionDto) GetIsDeployedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDeployed) {
		return nil, false
	}
	return o.IsDeployed, true
}

// HasIsDeployed returns a boolean if a field has been set.
func (o *ActionDto) HasIsDeployed() bool {
	if o != nil && !IsNil(o.IsDeployed) {
		return true
	}

	return false
}

// SetIsDeployed gets a reference to the given bool and assigns it to the IsDeployed field.
func (o *ActionDto) SetIsDeployed(v bool) {
	o.IsDeployed = &v
}

// GetConfigurationState returns the ConfigurationState field value if set, zero value otherwise.
func (o *ActionDto) GetConfigurationState() string {
	if o == nil || IsNil(o.ConfigurationState) {
		var ret string
		return ret
	}
	return *o.ConfigurationState
}

// GetConfigurationStateOk returns a tuple with the ConfigurationState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionDto) GetConfigurationStateOk() (*string, bool) {
	if o == nil || IsNil(o.ConfigurationState) {
		return nil, false
	}
	return o.ConfigurationState, true
}

// HasConfigurationState returns a boolean if a field has been set.
func (o *ActionDto) HasConfigurationState() bool {
	if o != nil && !IsNil(o.ConfigurationState) {
		return true
	}

	return false
}

// SetConfigurationState gets a reference to the given string and assigns it to the ConfigurationState field.
func (o *ActionDto) SetConfigurationState(v string) {
	o.ConfigurationState = &v
}

// GetConfigurationStateMessage returns the ConfigurationStateMessage field value if set, zero value otherwise.
func (o *ActionDto) GetConfigurationStateMessage() string {
	if o == nil || IsNil(o.ConfigurationStateMessage) {
		var ret string
		return ret
	}
	return *o.ConfigurationStateMessage
}

// GetConfigurationStateMessageOk returns a tuple with the ConfigurationStateMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionDto) GetConfigurationStateMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ConfigurationStateMessage) {
		return nil, false
	}
	return o.ConfigurationStateMessage, true
}

// HasConfigurationStateMessage returns a boolean if a field has been set.
func (o *ActionDto) HasConfigurationStateMessage() bool {
	if o != nil && !IsNil(o.ConfigurationStateMessage) {
		return true
	}

	return false
}

// SetConfigurationStateMessage gets a reference to the given string and assigns it to the ConfigurationStateMessage field.
func (o *ActionDto) SetConfigurationStateMessage(v string) {
	o.ConfigurationStateMessage = &v
}

// GetAppliedToIntegrations returns the AppliedToIntegrations field value if set, zero value otherwise.
func (o *ActionDto) GetAppliedToIntegrations() []ActionWithIntegrationDto {
	if o == nil || IsNil(o.AppliedToIntegrations) {
		var ret []ActionWithIntegrationDto
		return ret
	}
	return o.AppliedToIntegrations
}

// GetAppliedToIntegrationsOk returns a tuple with the AppliedToIntegrations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionDto) GetAppliedToIntegrationsOk() ([]ActionWithIntegrationDto, bool) {
	if o == nil || IsNil(o.AppliedToIntegrations) {
		return nil, false
	}
	return o.AppliedToIntegrations, true
}

// HasAppliedToIntegrations returns a boolean if a field has been set.
func (o *ActionDto) HasAppliedToIntegrations() bool {
	if o != nil && !IsNil(o.AppliedToIntegrations) {
		return true
	}

	return false
}

// SetAppliedToIntegrations gets a reference to the given []ActionWithIntegrationDto and assigns it to the AppliedToIntegrations field.
func (o *ActionDto) SetAppliedToIntegrations(v []ActionWithIntegrationDto) {
	o.AppliedToIntegrations = v
}

func (o ActionDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActionDto) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.IsCustomized) {
		toSerialize["isCustomized"] = o.IsCustomized
	}
	if !IsNil(o.ParentId) {
		toSerialize["parentId"] = o.ParentId
	}
	if !IsNil(o.ParentRevision) {
		toSerialize["parentRevision"] = o.ParentRevision
	}
	if !IsNil(o.IntegrationId) {
		toSerialize["integrationId"] = o.IntegrationId
	}
	if !IsNil(o.Integration) {
		toSerialize["integration"] = o.Integration
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.InputSchema) {
		toSerialize["inputSchema"] = o.InputSchema
	}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.OutputMapping) {
		toSerialize["outputMapping"] = o.OutputMapping
	}
	if !IsNil(o.CustomOutputSchema) {
		toSerialize["customOutputSchema"] = o.CustomOutputSchema
	}
	if !IsNil(o.DefaultOutputSchema) {
		toSerialize["defaultOutputSchema"] = o.DefaultOutputSchema
	}
	if !IsNil(o.TransformedOutputSchema) {
		toSerialize["transformedOutputSchema"] = o.TransformedOutputSchema
	}
	if !IsNil(o.OutputSchema) {
		toSerialize["outputSchema"] = o.OutputSchema
	}
	if !IsNil(o.Dependencies) {
		toSerialize["dependencies"] = o.Dependencies
	}
	if !IsNil(o.IsDeployed) {
		toSerialize["isDeployed"] = o.IsDeployed
	}
	if !IsNil(o.ConfigurationState) {
		toSerialize["configurationState"] = o.ConfigurationState
	}
	if !IsNil(o.ConfigurationStateMessage) {
		toSerialize["configurationStateMessage"] = o.ConfigurationStateMessage
	}
	if !IsNil(o.AppliedToIntegrations) {
		toSerialize["appliedToIntegrations"] = o.AppliedToIntegrations
	}
	return toSerialize, nil
}

func (o *ActionDto) UnmarshalJSON(data []byte) (err error) {
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

	varActionDto := _ActionDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varActionDto)

	if err != nil {
		return err
	}

	*o = ActionDto(varActionDto)

	return err
}

type NullableActionDto struct {
	value *ActionDto
	isSet bool
}

func (v NullableActionDto) Get() *ActionDto {
	return v.value
}

func (v *NullableActionDto) Set(val *ActionDto) {
	v.value = val
	v.isSet = true
}

func (v NullableActionDto) IsSet() bool {
	return v.isSet
}

func (v *NullableActionDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionDto(val *ActionDto) *NullableActionDto {
	return &NullableActionDto{value: val, isSet: true}
}

func (v NullableActionDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


