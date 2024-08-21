# CreateFieldMappingDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**ArchivedAt** | Pointer to **string** |  | [optional] 
**DataSourceKey** | Pointer to **string** |  | [optional] 
**DataSourceId** | Pointer to **string** |  | [optional] 
**AppSchema** | Pointer to **map[string]interface{}** |  | [optional] 
**Direction** | Pointer to **string** |  | [optional] 
**DefaultImportValue** | Pointer to **map[string]interface{}** |  | [optional] 
**ExportValue** | Pointer to **map[string]interface{}** |  | [optional] 
**ImportValue** | Pointer to **map[string]interface{}** |  | [optional] 
**DefaultExportValue** | Pointer to **map[string]interface{}** |  | [optional] 
**FrozenImportFields** | Pointer to **map[string]interface{}** |  | [optional] 
**FrozenExportFields** | Pointer to **map[string]interface{}** |  | [optional] 
**FrozenUnifiedExportFields** | Pointer to **map[string]interface{}** |  | [optional] 
**IntegrationId** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateFieldMappingDto

`func NewCreateFieldMappingDto(key string, ) *CreateFieldMappingDto`

NewCreateFieldMappingDto instantiates a new CreateFieldMappingDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFieldMappingDtoWithDefaults

`func NewCreateFieldMappingDtoWithDefaults() *CreateFieldMappingDto`

NewCreateFieldMappingDtoWithDefaults instantiates a new CreateFieldMappingDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *CreateFieldMappingDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateFieldMappingDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateFieldMappingDto) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *CreateFieldMappingDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateFieldMappingDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateFieldMappingDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateFieldMappingDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetArchivedAt

`func (o *CreateFieldMappingDto) GetArchivedAt() string`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *CreateFieldMappingDto) GetArchivedAtOk() (*string, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *CreateFieldMappingDto) SetArchivedAt(v string)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *CreateFieldMappingDto) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### GetDataSourceKey

`func (o *CreateFieldMappingDto) GetDataSourceKey() string`

GetDataSourceKey returns the DataSourceKey field if non-nil, zero value otherwise.

### GetDataSourceKeyOk

`func (o *CreateFieldMappingDto) GetDataSourceKeyOk() (*string, bool)`

GetDataSourceKeyOk returns a tuple with the DataSourceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceKey

`func (o *CreateFieldMappingDto) SetDataSourceKey(v string)`

SetDataSourceKey sets DataSourceKey field to given value.

### HasDataSourceKey

`func (o *CreateFieldMappingDto) HasDataSourceKey() bool`

HasDataSourceKey returns a boolean if a field has been set.

### GetDataSourceId

`func (o *CreateFieldMappingDto) GetDataSourceId() string`

GetDataSourceId returns the DataSourceId field if non-nil, zero value otherwise.

### GetDataSourceIdOk

`func (o *CreateFieldMappingDto) GetDataSourceIdOk() (*string, bool)`

GetDataSourceIdOk returns a tuple with the DataSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceId

`func (o *CreateFieldMappingDto) SetDataSourceId(v string)`

SetDataSourceId sets DataSourceId field to given value.

### HasDataSourceId

`func (o *CreateFieldMappingDto) HasDataSourceId() bool`

HasDataSourceId returns a boolean if a field has been set.

### GetAppSchema

`func (o *CreateFieldMappingDto) GetAppSchema() map[string]interface{}`

GetAppSchema returns the AppSchema field if non-nil, zero value otherwise.

### GetAppSchemaOk

`func (o *CreateFieldMappingDto) GetAppSchemaOk() (*map[string]interface{}, bool)`

GetAppSchemaOk returns a tuple with the AppSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppSchema

`func (o *CreateFieldMappingDto) SetAppSchema(v map[string]interface{})`

SetAppSchema sets AppSchema field to given value.

### HasAppSchema

`func (o *CreateFieldMappingDto) HasAppSchema() bool`

HasAppSchema returns a boolean if a field has been set.

### GetDirection

`func (o *CreateFieldMappingDto) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *CreateFieldMappingDto) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *CreateFieldMappingDto) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *CreateFieldMappingDto) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetDefaultImportValue

`func (o *CreateFieldMappingDto) GetDefaultImportValue() map[string]interface{}`

GetDefaultImportValue returns the DefaultImportValue field if non-nil, zero value otherwise.

### GetDefaultImportValueOk

`func (o *CreateFieldMappingDto) GetDefaultImportValueOk() (*map[string]interface{}, bool)`

GetDefaultImportValueOk returns a tuple with the DefaultImportValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultImportValue

`func (o *CreateFieldMappingDto) SetDefaultImportValue(v map[string]interface{})`

SetDefaultImportValue sets DefaultImportValue field to given value.

### HasDefaultImportValue

`func (o *CreateFieldMappingDto) HasDefaultImportValue() bool`

HasDefaultImportValue returns a boolean if a field has been set.

### GetExportValue

`func (o *CreateFieldMappingDto) GetExportValue() map[string]interface{}`

GetExportValue returns the ExportValue field if non-nil, zero value otherwise.

### GetExportValueOk

`func (o *CreateFieldMappingDto) GetExportValueOk() (*map[string]interface{}, bool)`

GetExportValueOk returns a tuple with the ExportValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportValue

`func (o *CreateFieldMappingDto) SetExportValue(v map[string]interface{})`

SetExportValue sets ExportValue field to given value.

### HasExportValue

`func (o *CreateFieldMappingDto) HasExportValue() bool`

HasExportValue returns a boolean if a field has been set.

### GetImportValue

`func (o *CreateFieldMappingDto) GetImportValue() map[string]interface{}`

GetImportValue returns the ImportValue field if non-nil, zero value otherwise.

### GetImportValueOk

`func (o *CreateFieldMappingDto) GetImportValueOk() (*map[string]interface{}, bool)`

GetImportValueOk returns a tuple with the ImportValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportValue

`func (o *CreateFieldMappingDto) SetImportValue(v map[string]interface{})`

SetImportValue sets ImportValue field to given value.

### HasImportValue

`func (o *CreateFieldMappingDto) HasImportValue() bool`

HasImportValue returns a boolean if a field has been set.

### GetDefaultExportValue

`func (o *CreateFieldMappingDto) GetDefaultExportValue() map[string]interface{}`

GetDefaultExportValue returns the DefaultExportValue field if non-nil, zero value otherwise.

### GetDefaultExportValueOk

`func (o *CreateFieldMappingDto) GetDefaultExportValueOk() (*map[string]interface{}, bool)`

GetDefaultExportValueOk returns a tuple with the DefaultExportValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultExportValue

`func (o *CreateFieldMappingDto) SetDefaultExportValue(v map[string]interface{})`

SetDefaultExportValue sets DefaultExportValue field to given value.

### HasDefaultExportValue

`func (o *CreateFieldMappingDto) HasDefaultExportValue() bool`

HasDefaultExportValue returns a boolean if a field has been set.

### GetFrozenImportFields

`func (o *CreateFieldMappingDto) GetFrozenImportFields() map[string]interface{}`

GetFrozenImportFields returns the FrozenImportFields field if non-nil, zero value otherwise.

### GetFrozenImportFieldsOk

`func (o *CreateFieldMappingDto) GetFrozenImportFieldsOk() (*map[string]interface{}, bool)`

GetFrozenImportFieldsOk returns a tuple with the FrozenImportFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrozenImportFields

`func (o *CreateFieldMappingDto) SetFrozenImportFields(v map[string]interface{})`

SetFrozenImportFields sets FrozenImportFields field to given value.

### HasFrozenImportFields

`func (o *CreateFieldMappingDto) HasFrozenImportFields() bool`

HasFrozenImportFields returns a boolean if a field has been set.

### GetFrozenExportFields

`func (o *CreateFieldMappingDto) GetFrozenExportFields() map[string]interface{}`

GetFrozenExportFields returns the FrozenExportFields field if non-nil, zero value otherwise.

### GetFrozenExportFieldsOk

`func (o *CreateFieldMappingDto) GetFrozenExportFieldsOk() (*map[string]interface{}, bool)`

GetFrozenExportFieldsOk returns a tuple with the FrozenExportFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrozenExportFields

`func (o *CreateFieldMappingDto) SetFrozenExportFields(v map[string]interface{})`

SetFrozenExportFields sets FrozenExportFields field to given value.

### HasFrozenExportFields

`func (o *CreateFieldMappingDto) HasFrozenExportFields() bool`

HasFrozenExportFields returns a boolean if a field has been set.

### GetFrozenUnifiedExportFields

`func (o *CreateFieldMappingDto) GetFrozenUnifiedExportFields() map[string]interface{}`

GetFrozenUnifiedExportFields returns the FrozenUnifiedExportFields field if non-nil, zero value otherwise.

### GetFrozenUnifiedExportFieldsOk

`func (o *CreateFieldMappingDto) GetFrozenUnifiedExportFieldsOk() (*map[string]interface{}, bool)`

GetFrozenUnifiedExportFieldsOk returns a tuple with the FrozenUnifiedExportFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrozenUnifiedExportFields

`func (o *CreateFieldMappingDto) SetFrozenUnifiedExportFields(v map[string]interface{})`

SetFrozenUnifiedExportFields sets FrozenUnifiedExportFields field to given value.

### HasFrozenUnifiedExportFields

`func (o *CreateFieldMappingDto) HasFrozenUnifiedExportFields() bool`

HasFrozenUnifiedExportFields returns a boolean if a field has been set.

### GetIntegrationId

`func (o *CreateFieldMappingDto) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *CreateFieldMappingDto) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *CreateFieldMappingDto) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.

### HasIntegrationId

`func (o *CreateFieldMappingDto) HasIntegrationId() bool`

HasIntegrationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


