# FieldMappingExportDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**IntegrationKeys** | Pointer to **[]string** |  | [optional] 
**DataSourceKey** | Pointer to **string** |  | [optional] 
**AppSchema** | Pointer to **map[string]interface{}** |  | [optional] 
**Direction** | Pointer to **string** |  | [optional] 
**DefaultImportValue** | Pointer to **map[string]interface{}** |  | [optional] 
**DefaultExportValue** | Pointer to **map[string]interface{}** |  | [optional] 
**ExportValue** | Pointer to **map[string]interface{}** |  | [optional] 
**ImportValue** | Pointer to **map[string]interface{}** |  | [optional] 
**FrozenImportFields** | Pointer to **map[string]interface{}** |  | [optional] 
**FrozenExportFields** | Pointer to **map[string]interface{}** |  | [optional] 
**FrozenUnifiedExportFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewFieldMappingExportDto

`func NewFieldMappingExportDto(name string, ) *FieldMappingExportDto`

NewFieldMappingExportDto instantiates a new FieldMappingExportDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldMappingExportDtoWithDefaults

`func NewFieldMappingExportDtoWithDefaults() *FieldMappingExportDto`

NewFieldMappingExportDtoWithDefaults instantiates a new FieldMappingExportDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *FieldMappingExportDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FieldMappingExportDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FieldMappingExportDto) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *FieldMappingExportDto) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *FieldMappingExportDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FieldMappingExportDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FieldMappingExportDto) SetName(v string)`

SetName sets Name field to given value.


### GetIntegrationKeys

`func (o *FieldMappingExportDto) GetIntegrationKeys() []string`

GetIntegrationKeys returns the IntegrationKeys field if non-nil, zero value otherwise.

### GetIntegrationKeysOk

`func (o *FieldMappingExportDto) GetIntegrationKeysOk() (*[]string, bool)`

GetIntegrationKeysOk returns a tuple with the IntegrationKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationKeys

`func (o *FieldMappingExportDto) SetIntegrationKeys(v []string)`

SetIntegrationKeys sets IntegrationKeys field to given value.

### HasIntegrationKeys

`func (o *FieldMappingExportDto) HasIntegrationKeys() bool`

HasIntegrationKeys returns a boolean if a field has been set.

### GetDataSourceKey

`func (o *FieldMappingExportDto) GetDataSourceKey() string`

GetDataSourceKey returns the DataSourceKey field if non-nil, zero value otherwise.

### GetDataSourceKeyOk

`func (o *FieldMappingExportDto) GetDataSourceKeyOk() (*string, bool)`

GetDataSourceKeyOk returns a tuple with the DataSourceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceKey

`func (o *FieldMappingExportDto) SetDataSourceKey(v string)`

SetDataSourceKey sets DataSourceKey field to given value.

### HasDataSourceKey

`func (o *FieldMappingExportDto) HasDataSourceKey() bool`

HasDataSourceKey returns a boolean if a field has been set.

### GetAppSchema

`func (o *FieldMappingExportDto) GetAppSchema() map[string]interface{}`

GetAppSchema returns the AppSchema field if non-nil, zero value otherwise.

### GetAppSchemaOk

`func (o *FieldMappingExportDto) GetAppSchemaOk() (*map[string]interface{}, bool)`

GetAppSchemaOk returns a tuple with the AppSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppSchema

`func (o *FieldMappingExportDto) SetAppSchema(v map[string]interface{})`

SetAppSchema sets AppSchema field to given value.

### HasAppSchema

`func (o *FieldMappingExportDto) HasAppSchema() bool`

HasAppSchema returns a boolean if a field has been set.

### GetDirection

`func (o *FieldMappingExportDto) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *FieldMappingExportDto) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *FieldMappingExportDto) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *FieldMappingExportDto) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetDefaultImportValue

`func (o *FieldMappingExportDto) GetDefaultImportValue() map[string]interface{}`

GetDefaultImportValue returns the DefaultImportValue field if non-nil, zero value otherwise.

### GetDefaultImportValueOk

`func (o *FieldMappingExportDto) GetDefaultImportValueOk() (*map[string]interface{}, bool)`

GetDefaultImportValueOk returns a tuple with the DefaultImportValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultImportValue

`func (o *FieldMappingExportDto) SetDefaultImportValue(v map[string]interface{})`

SetDefaultImportValue sets DefaultImportValue field to given value.

### HasDefaultImportValue

`func (o *FieldMappingExportDto) HasDefaultImportValue() bool`

HasDefaultImportValue returns a boolean if a field has been set.

### GetDefaultExportValue

`func (o *FieldMappingExportDto) GetDefaultExportValue() map[string]interface{}`

GetDefaultExportValue returns the DefaultExportValue field if non-nil, zero value otherwise.

### GetDefaultExportValueOk

`func (o *FieldMappingExportDto) GetDefaultExportValueOk() (*map[string]interface{}, bool)`

GetDefaultExportValueOk returns a tuple with the DefaultExportValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultExportValue

`func (o *FieldMappingExportDto) SetDefaultExportValue(v map[string]interface{})`

SetDefaultExportValue sets DefaultExportValue field to given value.

### HasDefaultExportValue

`func (o *FieldMappingExportDto) HasDefaultExportValue() bool`

HasDefaultExportValue returns a boolean if a field has been set.

### GetExportValue

`func (o *FieldMappingExportDto) GetExportValue() map[string]interface{}`

GetExportValue returns the ExportValue field if non-nil, zero value otherwise.

### GetExportValueOk

`func (o *FieldMappingExportDto) GetExportValueOk() (*map[string]interface{}, bool)`

GetExportValueOk returns a tuple with the ExportValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportValue

`func (o *FieldMappingExportDto) SetExportValue(v map[string]interface{})`

SetExportValue sets ExportValue field to given value.

### HasExportValue

`func (o *FieldMappingExportDto) HasExportValue() bool`

HasExportValue returns a boolean if a field has been set.

### GetImportValue

`func (o *FieldMappingExportDto) GetImportValue() map[string]interface{}`

GetImportValue returns the ImportValue field if non-nil, zero value otherwise.

### GetImportValueOk

`func (o *FieldMappingExportDto) GetImportValueOk() (*map[string]interface{}, bool)`

GetImportValueOk returns a tuple with the ImportValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportValue

`func (o *FieldMappingExportDto) SetImportValue(v map[string]interface{})`

SetImportValue sets ImportValue field to given value.

### HasImportValue

`func (o *FieldMappingExportDto) HasImportValue() bool`

HasImportValue returns a boolean if a field has been set.

### GetFrozenImportFields

`func (o *FieldMappingExportDto) GetFrozenImportFields() map[string]interface{}`

GetFrozenImportFields returns the FrozenImportFields field if non-nil, zero value otherwise.

### GetFrozenImportFieldsOk

`func (o *FieldMappingExportDto) GetFrozenImportFieldsOk() (*map[string]interface{}, bool)`

GetFrozenImportFieldsOk returns a tuple with the FrozenImportFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrozenImportFields

`func (o *FieldMappingExportDto) SetFrozenImportFields(v map[string]interface{})`

SetFrozenImportFields sets FrozenImportFields field to given value.

### HasFrozenImportFields

`func (o *FieldMappingExportDto) HasFrozenImportFields() bool`

HasFrozenImportFields returns a boolean if a field has been set.

### GetFrozenExportFields

`func (o *FieldMappingExportDto) GetFrozenExportFields() map[string]interface{}`

GetFrozenExportFields returns the FrozenExportFields field if non-nil, zero value otherwise.

### GetFrozenExportFieldsOk

`func (o *FieldMappingExportDto) GetFrozenExportFieldsOk() (*map[string]interface{}, bool)`

GetFrozenExportFieldsOk returns a tuple with the FrozenExportFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrozenExportFields

`func (o *FieldMappingExportDto) SetFrozenExportFields(v map[string]interface{})`

SetFrozenExportFields sets FrozenExportFields field to given value.

### HasFrozenExportFields

`func (o *FieldMappingExportDto) HasFrozenExportFields() bool`

HasFrozenExportFields returns a boolean if a field has been set.

### GetFrozenUnifiedExportFields

`func (o *FieldMappingExportDto) GetFrozenUnifiedExportFields() map[string]interface{}`

GetFrozenUnifiedExportFields returns the FrozenUnifiedExportFields field if non-nil, zero value otherwise.

### GetFrozenUnifiedExportFieldsOk

`func (o *FieldMappingExportDto) GetFrozenUnifiedExportFieldsOk() (*map[string]interface{}, bool)`

GetFrozenUnifiedExportFieldsOk returns a tuple with the FrozenUnifiedExportFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrozenUnifiedExportFields

`func (o *FieldMappingExportDto) SetFrozenUnifiedExportFields(v map[string]interface{})`

SetFrozenUnifiedExportFields sets FrozenUnifiedExportFields field to given value.

### HasFrozenUnifiedExportFields

`func (o *FieldMappingExportDto) HasFrozenUnifiedExportFields() bool`

HasFrozenUnifiedExportFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


