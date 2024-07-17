# FieldMappingInstanceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Revision** | **string** |  | 
**UserId** | **string** |  | 
**User** | Pointer to [**CustomerDto**](CustomerDto.md) |  | [optional] 
**ConnectionId** | **string** |  | 
**Connection** | Pointer to [**ConnectionDto**](ConnectionDto.md) |  | [optional] 
**IntegrationId** | **string** |  | 
**Integration** | Pointer to [**IntegrationDto**](IntegrationDto.md) |  | [optional] 
**InstanceKey** | Pointer to **string** |  | [optional] 
**FieldMappingId** | Pointer to **string** |  | [optional] 
**FieldMappingRevision** | Pointer to **string** |  | [optional] 
**FieldMapping** | Pointer to [**FieldMappingDto**](FieldMappingDto.md) |  | [optional] 
**Direction** | Pointer to **string** |  | [optional] 
**DataSourceInstanceId** | Pointer to **string** |  | [optional] 
**DataSourceInstance** | Pointer to **map[string]interface{}** |  | [optional] 
**DataSourceSchema** | Pointer to **map[string]interface{}** |  | [optional] 
**AppSchema** | Pointer to **map[string]interface{}** |  | [optional] 
**ImportValue** | Pointer to **map[string]interface{}** |  | [optional] 
**ExportValue** | Pointer to **map[string]interface{}** |  | [optional] 
**UnifiedImportValue** | Pointer to **map[string]interface{}** |  | [optional] 
**UnifiedExportValue** | Pointer to **map[string]interface{}** |  | [optional] 
**FrozenImportFields** | Pointer to **map[string]interface{}** |  | [optional] 
**FrozenExportFields** | Pointer to **map[string]interface{}** |  | [optional] 
**IsCustomized** | Pointer to **bool** |  | [optional] 
**Error** | Pointer to **map[string]interface{}** |  | [optional] 
**ExternalSchema** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewFieldMappingInstanceDto

`func NewFieldMappingInstanceDto(id string, name string, revision string, userId string, connectionId string, integrationId string, ) *FieldMappingInstanceDto`

NewFieldMappingInstanceDto instantiates a new FieldMappingInstanceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldMappingInstanceDtoWithDefaults

`func NewFieldMappingInstanceDtoWithDefaults() *FieldMappingInstanceDto`

NewFieldMappingInstanceDtoWithDefaults instantiates a new FieldMappingInstanceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FieldMappingInstanceDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FieldMappingInstanceDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FieldMappingInstanceDto) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *FieldMappingInstanceDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FieldMappingInstanceDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FieldMappingInstanceDto) SetName(v string)`

SetName sets Name field to given value.


### GetRevision

`func (o *FieldMappingInstanceDto) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *FieldMappingInstanceDto) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *FieldMappingInstanceDto) SetRevision(v string)`

SetRevision sets Revision field to given value.


### GetUserId

`func (o *FieldMappingInstanceDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *FieldMappingInstanceDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *FieldMappingInstanceDto) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetUser

`func (o *FieldMappingInstanceDto) GetUser() CustomerDto`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *FieldMappingInstanceDto) GetUserOk() (*CustomerDto, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *FieldMappingInstanceDto) SetUser(v CustomerDto)`

SetUser sets User field to given value.

### HasUser

`func (o *FieldMappingInstanceDto) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetConnectionId

`func (o *FieldMappingInstanceDto) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *FieldMappingInstanceDto) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *FieldMappingInstanceDto) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetConnection

`func (o *FieldMappingInstanceDto) GetConnection() ConnectionDto`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *FieldMappingInstanceDto) GetConnectionOk() (*ConnectionDto, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *FieldMappingInstanceDto) SetConnection(v ConnectionDto)`

SetConnection sets Connection field to given value.

### HasConnection

`func (o *FieldMappingInstanceDto) HasConnection() bool`

HasConnection returns a boolean if a field has been set.

### GetIntegrationId

`func (o *FieldMappingInstanceDto) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *FieldMappingInstanceDto) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *FieldMappingInstanceDto) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.


### GetIntegration

`func (o *FieldMappingInstanceDto) GetIntegration() IntegrationDto`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *FieldMappingInstanceDto) GetIntegrationOk() (*IntegrationDto, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *FieldMappingInstanceDto) SetIntegration(v IntegrationDto)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *FieldMappingInstanceDto) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetInstanceKey

`func (o *FieldMappingInstanceDto) GetInstanceKey() string`

GetInstanceKey returns the InstanceKey field if non-nil, zero value otherwise.

### GetInstanceKeyOk

`func (o *FieldMappingInstanceDto) GetInstanceKeyOk() (*string, bool)`

GetInstanceKeyOk returns a tuple with the InstanceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceKey

`func (o *FieldMappingInstanceDto) SetInstanceKey(v string)`

SetInstanceKey sets InstanceKey field to given value.

### HasInstanceKey

`func (o *FieldMappingInstanceDto) HasInstanceKey() bool`

HasInstanceKey returns a boolean if a field has been set.

### GetFieldMappingId

`func (o *FieldMappingInstanceDto) GetFieldMappingId() string`

GetFieldMappingId returns the FieldMappingId field if non-nil, zero value otherwise.

### GetFieldMappingIdOk

`func (o *FieldMappingInstanceDto) GetFieldMappingIdOk() (*string, bool)`

GetFieldMappingIdOk returns a tuple with the FieldMappingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldMappingId

`func (o *FieldMappingInstanceDto) SetFieldMappingId(v string)`

SetFieldMappingId sets FieldMappingId field to given value.

### HasFieldMappingId

`func (o *FieldMappingInstanceDto) HasFieldMappingId() bool`

HasFieldMappingId returns a boolean if a field has been set.

### GetFieldMappingRevision

`func (o *FieldMappingInstanceDto) GetFieldMappingRevision() string`

GetFieldMappingRevision returns the FieldMappingRevision field if non-nil, zero value otherwise.

### GetFieldMappingRevisionOk

`func (o *FieldMappingInstanceDto) GetFieldMappingRevisionOk() (*string, bool)`

GetFieldMappingRevisionOk returns a tuple with the FieldMappingRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldMappingRevision

`func (o *FieldMappingInstanceDto) SetFieldMappingRevision(v string)`

SetFieldMappingRevision sets FieldMappingRevision field to given value.

### HasFieldMappingRevision

`func (o *FieldMappingInstanceDto) HasFieldMappingRevision() bool`

HasFieldMappingRevision returns a boolean if a field has been set.

### GetFieldMapping

`func (o *FieldMappingInstanceDto) GetFieldMapping() FieldMappingDto`

GetFieldMapping returns the FieldMapping field if non-nil, zero value otherwise.

### GetFieldMappingOk

`func (o *FieldMappingInstanceDto) GetFieldMappingOk() (*FieldMappingDto, bool)`

GetFieldMappingOk returns a tuple with the FieldMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldMapping

`func (o *FieldMappingInstanceDto) SetFieldMapping(v FieldMappingDto)`

SetFieldMapping sets FieldMapping field to given value.

### HasFieldMapping

`func (o *FieldMappingInstanceDto) HasFieldMapping() bool`

HasFieldMapping returns a boolean if a field has been set.

### GetDirection

`func (o *FieldMappingInstanceDto) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *FieldMappingInstanceDto) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *FieldMappingInstanceDto) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *FieldMappingInstanceDto) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetDataSourceInstanceId

`func (o *FieldMappingInstanceDto) GetDataSourceInstanceId() string`

GetDataSourceInstanceId returns the DataSourceInstanceId field if non-nil, zero value otherwise.

### GetDataSourceInstanceIdOk

`func (o *FieldMappingInstanceDto) GetDataSourceInstanceIdOk() (*string, bool)`

GetDataSourceInstanceIdOk returns a tuple with the DataSourceInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceInstanceId

`func (o *FieldMappingInstanceDto) SetDataSourceInstanceId(v string)`

SetDataSourceInstanceId sets DataSourceInstanceId field to given value.

### HasDataSourceInstanceId

`func (o *FieldMappingInstanceDto) HasDataSourceInstanceId() bool`

HasDataSourceInstanceId returns a boolean if a field has been set.

### GetDataSourceInstance

`func (o *FieldMappingInstanceDto) GetDataSourceInstance() map[string]interface{}`

GetDataSourceInstance returns the DataSourceInstance field if non-nil, zero value otherwise.

### GetDataSourceInstanceOk

`func (o *FieldMappingInstanceDto) GetDataSourceInstanceOk() (*map[string]interface{}, bool)`

GetDataSourceInstanceOk returns a tuple with the DataSourceInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceInstance

`func (o *FieldMappingInstanceDto) SetDataSourceInstance(v map[string]interface{})`

SetDataSourceInstance sets DataSourceInstance field to given value.

### HasDataSourceInstance

`func (o *FieldMappingInstanceDto) HasDataSourceInstance() bool`

HasDataSourceInstance returns a boolean if a field has been set.

### GetDataSourceSchema

`func (o *FieldMappingInstanceDto) GetDataSourceSchema() map[string]interface{}`

GetDataSourceSchema returns the DataSourceSchema field if non-nil, zero value otherwise.

### GetDataSourceSchemaOk

`func (o *FieldMappingInstanceDto) GetDataSourceSchemaOk() (*map[string]interface{}, bool)`

GetDataSourceSchemaOk returns a tuple with the DataSourceSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceSchema

`func (o *FieldMappingInstanceDto) SetDataSourceSchema(v map[string]interface{})`

SetDataSourceSchema sets DataSourceSchema field to given value.

### HasDataSourceSchema

`func (o *FieldMappingInstanceDto) HasDataSourceSchema() bool`

HasDataSourceSchema returns a boolean if a field has been set.

### GetAppSchema

`func (o *FieldMappingInstanceDto) GetAppSchema() map[string]interface{}`

GetAppSchema returns the AppSchema field if non-nil, zero value otherwise.

### GetAppSchemaOk

`func (o *FieldMappingInstanceDto) GetAppSchemaOk() (*map[string]interface{}, bool)`

GetAppSchemaOk returns a tuple with the AppSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppSchema

`func (o *FieldMappingInstanceDto) SetAppSchema(v map[string]interface{})`

SetAppSchema sets AppSchema field to given value.

### HasAppSchema

`func (o *FieldMappingInstanceDto) HasAppSchema() bool`

HasAppSchema returns a boolean if a field has been set.

### GetImportValue

`func (o *FieldMappingInstanceDto) GetImportValue() map[string]interface{}`

GetImportValue returns the ImportValue field if non-nil, zero value otherwise.

### GetImportValueOk

`func (o *FieldMappingInstanceDto) GetImportValueOk() (*map[string]interface{}, bool)`

GetImportValueOk returns a tuple with the ImportValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportValue

`func (o *FieldMappingInstanceDto) SetImportValue(v map[string]interface{})`

SetImportValue sets ImportValue field to given value.

### HasImportValue

`func (o *FieldMappingInstanceDto) HasImportValue() bool`

HasImportValue returns a boolean if a field has been set.

### GetExportValue

`func (o *FieldMappingInstanceDto) GetExportValue() map[string]interface{}`

GetExportValue returns the ExportValue field if non-nil, zero value otherwise.

### GetExportValueOk

`func (o *FieldMappingInstanceDto) GetExportValueOk() (*map[string]interface{}, bool)`

GetExportValueOk returns a tuple with the ExportValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportValue

`func (o *FieldMappingInstanceDto) SetExportValue(v map[string]interface{})`

SetExportValue sets ExportValue field to given value.

### HasExportValue

`func (o *FieldMappingInstanceDto) HasExportValue() bool`

HasExportValue returns a boolean if a field has been set.

### GetUnifiedImportValue

`func (o *FieldMappingInstanceDto) GetUnifiedImportValue() map[string]interface{}`

GetUnifiedImportValue returns the UnifiedImportValue field if non-nil, zero value otherwise.

### GetUnifiedImportValueOk

`func (o *FieldMappingInstanceDto) GetUnifiedImportValueOk() (*map[string]interface{}, bool)`

GetUnifiedImportValueOk returns a tuple with the UnifiedImportValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnifiedImportValue

`func (o *FieldMappingInstanceDto) SetUnifiedImportValue(v map[string]interface{})`

SetUnifiedImportValue sets UnifiedImportValue field to given value.

### HasUnifiedImportValue

`func (o *FieldMappingInstanceDto) HasUnifiedImportValue() bool`

HasUnifiedImportValue returns a boolean if a field has been set.

### GetUnifiedExportValue

`func (o *FieldMappingInstanceDto) GetUnifiedExportValue() map[string]interface{}`

GetUnifiedExportValue returns the UnifiedExportValue field if non-nil, zero value otherwise.

### GetUnifiedExportValueOk

`func (o *FieldMappingInstanceDto) GetUnifiedExportValueOk() (*map[string]interface{}, bool)`

GetUnifiedExportValueOk returns a tuple with the UnifiedExportValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnifiedExportValue

`func (o *FieldMappingInstanceDto) SetUnifiedExportValue(v map[string]interface{})`

SetUnifiedExportValue sets UnifiedExportValue field to given value.

### HasUnifiedExportValue

`func (o *FieldMappingInstanceDto) HasUnifiedExportValue() bool`

HasUnifiedExportValue returns a boolean if a field has been set.

### GetFrozenImportFields

`func (o *FieldMappingInstanceDto) GetFrozenImportFields() map[string]interface{}`

GetFrozenImportFields returns the FrozenImportFields field if non-nil, zero value otherwise.

### GetFrozenImportFieldsOk

`func (o *FieldMappingInstanceDto) GetFrozenImportFieldsOk() (*map[string]interface{}, bool)`

GetFrozenImportFieldsOk returns a tuple with the FrozenImportFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrozenImportFields

`func (o *FieldMappingInstanceDto) SetFrozenImportFields(v map[string]interface{})`

SetFrozenImportFields sets FrozenImportFields field to given value.

### HasFrozenImportFields

`func (o *FieldMappingInstanceDto) HasFrozenImportFields() bool`

HasFrozenImportFields returns a boolean if a field has been set.

### GetFrozenExportFields

`func (o *FieldMappingInstanceDto) GetFrozenExportFields() map[string]interface{}`

GetFrozenExportFields returns the FrozenExportFields field if non-nil, zero value otherwise.

### GetFrozenExportFieldsOk

`func (o *FieldMappingInstanceDto) GetFrozenExportFieldsOk() (*map[string]interface{}, bool)`

GetFrozenExportFieldsOk returns a tuple with the FrozenExportFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrozenExportFields

`func (o *FieldMappingInstanceDto) SetFrozenExportFields(v map[string]interface{})`

SetFrozenExportFields sets FrozenExportFields field to given value.

### HasFrozenExportFields

`func (o *FieldMappingInstanceDto) HasFrozenExportFields() bool`

HasFrozenExportFields returns a boolean if a field has been set.

### GetIsCustomized

`func (o *FieldMappingInstanceDto) GetIsCustomized() bool`

GetIsCustomized returns the IsCustomized field if non-nil, zero value otherwise.

### GetIsCustomizedOk

`func (o *FieldMappingInstanceDto) GetIsCustomizedOk() (*bool, bool)`

GetIsCustomizedOk returns a tuple with the IsCustomized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCustomized

`func (o *FieldMappingInstanceDto) SetIsCustomized(v bool)`

SetIsCustomized sets IsCustomized field to given value.

### HasIsCustomized

`func (o *FieldMappingInstanceDto) HasIsCustomized() bool`

HasIsCustomized returns a boolean if a field has been set.

### GetError

`func (o *FieldMappingInstanceDto) GetError() map[string]interface{}`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *FieldMappingInstanceDto) GetErrorOk() (*map[string]interface{}, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *FieldMappingInstanceDto) SetError(v map[string]interface{})`

SetError sets Error field to given value.

### HasError

`func (o *FieldMappingInstanceDto) HasError() bool`

HasError returns a boolean if a field has been set.

### GetExternalSchema

`func (o *FieldMappingInstanceDto) GetExternalSchema() map[string]interface{}`

GetExternalSchema returns the ExternalSchema field if non-nil, zero value otherwise.

### GetExternalSchemaOk

`func (o *FieldMappingInstanceDto) GetExternalSchemaOk() (*map[string]interface{}, bool)`

GetExternalSchemaOk returns a tuple with the ExternalSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSchema

`func (o *FieldMappingInstanceDto) SetExternalSchema(v map[string]interface{})`

SetExternalSchema sets ExternalSchema field to given value.

### HasExternalSchema

`func (o *FieldMappingInstanceDto) HasExternalSchema() bool`

HasExternalSchema returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


