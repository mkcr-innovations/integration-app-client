# CreateIntegrationLevelFieldMappingDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Name** | **string** |  | 
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

## Methods

### NewCreateIntegrationLevelFieldMappingDto

`func NewCreateIntegrationLevelFieldMappingDto(key string, name string, ) *CreateIntegrationLevelFieldMappingDto`

NewCreateIntegrationLevelFieldMappingDto instantiates a new CreateIntegrationLevelFieldMappingDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIntegrationLevelFieldMappingDtoWithDefaults

`func NewCreateIntegrationLevelFieldMappingDtoWithDefaults() *CreateIntegrationLevelFieldMappingDto`

NewCreateIntegrationLevelFieldMappingDtoWithDefaults instantiates a new CreateIntegrationLevelFieldMappingDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *CreateIntegrationLevelFieldMappingDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateIntegrationLevelFieldMappingDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateIntegrationLevelFieldMappingDto) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *CreateIntegrationLevelFieldMappingDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateIntegrationLevelFieldMappingDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateIntegrationLevelFieldMappingDto) SetName(v string)`

SetName sets Name field to given value.


### GetArchivedAt

`func (o *CreateIntegrationLevelFieldMappingDto) GetArchivedAt() string`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *CreateIntegrationLevelFieldMappingDto) GetArchivedAtOk() (*string, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *CreateIntegrationLevelFieldMappingDto) SetArchivedAt(v string)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *CreateIntegrationLevelFieldMappingDto) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### GetDataSourceKey

`func (o *CreateIntegrationLevelFieldMappingDto) GetDataSourceKey() string`

GetDataSourceKey returns the DataSourceKey field if non-nil, zero value otherwise.

### GetDataSourceKeyOk

`func (o *CreateIntegrationLevelFieldMappingDto) GetDataSourceKeyOk() (*string, bool)`

GetDataSourceKeyOk returns a tuple with the DataSourceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceKey

`func (o *CreateIntegrationLevelFieldMappingDto) SetDataSourceKey(v string)`

SetDataSourceKey sets DataSourceKey field to given value.

### HasDataSourceKey

`func (o *CreateIntegrationLevelFieldMappingDto) HasDataSourceKey() bool`

HasDataSourceKey returns a boolean if a field has been set.

### GetDataSourceId

`func (o *CreateIntegrationLevelFieldMappingDto) GetDataSourceId() string`

GetDataSourceId returns the DataSourceId field if non-nil, zero value otherwise.

### GetDataSourceIdOk

`func (o *CreateIntegrationLevelFieldMappingDto) GetDataSourceIdOk() (*string, bool)`

GetDataSourceIdOk returns a tuple with the DataSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceId

`func (o *CreateIntegrationLevelFieldMappingDto) SetDataSourceId(v string)`

SetDataSourceId sets DataSourceId field to given value.

### HasDataSourceId

`func (o *CreateIntegrationLevelFieldMappingDto) HasDataSourceId() bool`

HasDataSourceId returns a boolean if a field has been set.

### GetAppSchema

`func (o *CreateIntegrationLevelFieldMappingDto) GetAppSchema() map[string]interface{}`

GetAppSchema returns the AppSchema field if non-nil, zero value otherwise.

### GetAppSchemaOk

`func (o *CreateIntegrationLevelFieldMappingDto) GetAppSchemaOk() (*map[string]interface{}, bool)`

GetAppSchemaOk returns a tuple with the AppSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppSchema

`func (o *CreateIntegrationLevelFieldMappingDto) SetAppSchema(v map[string]interface{})`

SetAppSchema sets AppSchema field to given value.

### HasAppSchema

`func (o *CreateIntegrationLevelFieldMappingDto) HasAppSchema() bool`

HasAppSchema returns a boolean if a field has been set.

### GetDirection

`func (o *CreateIntegrationLevelFieldMappingDto) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *CreateIntegrationLevelFieldMappingDto) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *CreateIntegrationLevelFieldMappingDto) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *CreateIntegrationLevelFieldMappingDto) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetDefaultImportValue

`func (o *CreateIntegrationLevelFieldMappingDto) GetDefaultImportValue() map[string]interface{}`

GetDefaultImportValue returns the DefaultImportValue field if non-nil, zero value otherwise.

### GetDefaultImportValueOk

`func (o *CreateIntegrationLevelFieldMappingDto) GetDefaultImportValueOk() (*map[string]interface{}, bool)`

GetDefaultImportValueOk returns a tuple with the DefaultImportValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultImportValue

`func (o *CreateIntegrationLevelFieldMappingDto) SetDefaultImportValue(v map[string]interface{})`

SetDefaultImportValue sets DefaultImportValue field to given value.

### HasDefaultImportValue

`func (o *CreateIntegrationLevelFieldMappingDto) HasDefaultImportValue() bool`

HasDefaultImportValue returns a boolean if a field has been set.

### GetExportValue

`func (o *CreateIntegrationLevelFieldMappingDto) GetExportValue() map[string]interface{}`

GetExportValue returns the ExportValue field if non-nil, zero value otherwise.

### GetExportValueOk

`func (o *CreateIntegrationLevelFieldMappingDto) GetExportValueOk() (*map[string]interface{}, bool)`

GetExportValueOk returns a tuple with the ExportValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportValue

`func (o *CreateIntegrationLevelFieldMappingDto) SetExportValue(v map[string]interface{})`

SetExportValue sets ExportValue field to given value.

### HasExportValue

`func (o *CreateIntegrationLevelFieldMappingDto) HasExportValue() bool`

HasExportValue returns a boolean if a field has been set.

### GetImportValue

`func (o *CreateIntegrationLevelFieldMappingDto) GetImportValue() map[string]interface{}`

GetImportValue returns the ImportValue field if non-nil, zero value otherwise.

### GetImportValueOk

`func (o *CreateIntegrationLevelFieldMappingDto) GetImportValueOk() (*map[string]interface{}, bool)`

GetImportValueOk returns a tuple with the ImportValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportValue

`func (o *CreateIntegrationLevelFieldMappingDto) SetImportValue(v map[string]interface{})`

SetImportValue sets ImportValue field to given value.

### HasImportValue

`func (o *CreateIntegrationLevelFieldMappingDto) HasImportValue() bool`

HasImportValue returns a boolean if a field has been set.

### GetDefaultExportValue

`func (o *CreateIntegrationLevelFieldMappingDto) GetDefaultExportValue() map[string]interface{}`

GetDefaultExportValue returns the DefaultExportValue field if non-nil, zero value otherwise.

### GetDefaultExportValueOk

`func (o *CreateIntegrationLevelFieldMappingDto) GetDefaultExportValueOk() (*map[string]interface{}, bool)`

GetDefaultExportValueOk returns a tuple with the DefaultExportValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultExportValue

`func (o *CreateIntegrationLevelFieldMappingDto) SetDefaultExportValue(v map[string]interface{})`

SetDefaultExportValue sets DefaultExportValue field to given value.

### HasDefaultExportValue

`func (o *CreateIntegrationLevelFieldMappingDto) HasDefaultExportValue() bool`

HasDefaultExportValue returns a boolean if a field has been set.

### GetFrozenImportFields

`func (o *CreateIntegrationLevelFieldMappingDto) GetFrozenImportFields() map[string]interface{}`

GetFrozenImportFields returns the FrozenImportFields field if non-nil, zero value otherwise.

### GetFrozenImportFieldsOk

`func (o *CreateIntegrationLevelFieldMappingDto) GetFrozenImportFieldsOk() (*map[string]interface{}, bool)`

GetFrozenImportFieldsOk returns a tuple with the FrozenImportFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrozenImportFields

`func (o *CreateIntegrationLevelFieldMappingDto) SetFrozenImportFields(v map[string]interface{})`

SetFrozenImportFields sets FrozenImportFields field to given value.

### HasFrozenImportFields

`func (o *CreateIntegrationLevelFieldMappingDto) HasFrozenImportFields() bool`

HasFrozenImportFields returns a boolean if a field has been set.

### GetFrozenExportFields

`func (o *CreateIntegrationLevelFieldMappingDto) GetFrozenExportFields() map[string]interface{}`

GetFrozenExportFields returns the FrozenExportFields field if non-nil, zero value otherwise.

### GetFrozenExportFieldsOk

`func (o *CreateIntegrationLevelFieldMappingDto) GetFrozenExportFieldsOk() (*map[string]interface{}, bool)`

GetFrozenExportFieldsOk returns a tuple with the FrozenExportFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrozenExportFields

`func (o *CreateIntegrationLevelFieldMappingDto) SetFrozenExportFields(v map[string]interface{})`

SetFrozenExportFields sets FrozenExportFields field to given value.

### HasFrozenExportFields

`func (o *CreateIntegrationLevelFieldMappingDto) HasFrozenExportFields() bool`

HasFrozenExportFields returns a boolean if a field has been set.

### GetFrozenUnifiedExportFields

`func (o *CreateIntegrationLevelFieldMappingDto) GetFrozenUnifiedExportFields() map[string]interface{}`

GetFrozenUnifiedExportFields returns the FrozenUnifiedExportFields field if non-nil, zero value otherwise.

### GetFrozenUnifiedExportFieldsOk

`func (o *CreateIntegrationLevelFieldMappingDto) GetFrozenUnifiedExportFieldsOk() (*map[string]interface{}, bool)`

GetFrozenUnifiedExportFieldsOk returns a tuple with the FrozenUnifiedExportFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrozenUnifiedExportFields

`func (o *CreateIntegrationLevelFieldMappingDto) SetFrozenUnifiedExportFields(v map[string]interface{})`

SetFrozenUnifiedExportFields sets FrozenUnifiedExportFields field to given value.

### HasFrozenUnifiedExportFields

`func (o *CreateIntegrationLevelFieldMappingDto) HasFrozenUnifiedExportFields() bool`

HasFrozenUnifiedExportFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


