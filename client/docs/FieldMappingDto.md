# FieldMappingDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Key** | **string** |  | 
**Name** | **string** |  | 
**ArchivedAt** | Pointer to **string** |  | [optional] 
**Customized** | Pointer to **bool** |  | [optional] 
**Revision** | **string** |  | 
**UniversalFieldMappingId** | Pointer to **string** |  | [optional] 
**UniversalFieldMappingRevision** | Pointer to **string** |  | [optional] 
**IntegrationId** | Pointer to **string** |  | [optional] 
**DataSourceKey** | Pointer to **string** |  | [optional] 
**DataSourceId** | Pointer to **string** |  | [optional] 
**AppSchema** | Pointer to **map[string]interface{}** |  | [optional] 
**Direction** | Pointer to **string** |  | [optional] 
**DefaultImportValue** | Pointer to **map[string]interface{}** |  | [optional] 
**DefaultExportValue** | Pointer to **map[string]interface{}** |  | [optional] 
**ExportValue** | Pointer to **map[string]interface{}** |  | [optional] 
**ImportValue** | Pointer to **map[string]interface{}** |  | [optional] 
**FrozenImportFields** | Pointer to **map[string]interface{}** |  | [optional] 
**FrozenExportFields** | Pointer to **map[string]interface{}** |  | [optional] 
**FrozenUnifiedExportFields** | Pointer to **map[string]interface{}** |  | [optional] 
**AppliedToIntegrations** | Pointer to [**[]FiledMappingWithIntegrationDto**](FiledMappingWithIntegrationDto.md) |  | [optional] 

## Methods

### NewFieldMappingDto

`func NewFieldMappingDto(id string, key string, name string, revision string, ) *FieldMappingDto`

NewFieldMappingDto instantiates a new FieldMappingDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldMappingDtoWithDefaults

`func NewFieldMappingDtoWithDefaults() *FieldMappingDto`

NewFieldMappingDtoWithDefaults instantiates a new FieldMappingDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FieldMappingDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FieldMappingDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FieldMappingDto) SetId(v string)`

SetId sets Id field to given value.


### GetKey

`func (o *FieldMappingDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FieldMappingDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FieldMappingDto) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *FieldMappingDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FieldMappingDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FieldMappingDto) SetName(v string)`

SetName sets Name field to given value.


### GetArchivedAt

`func (o *FieldMappingDto) GetArchivedAt() string`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *FieldMappingDto) GetArchivedAtOk() (*string, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *FieldMappingDto) SetArchivedAt(v string)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *FieldMappingDto) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### GetCustomized

`func (o *FieldMappingDto) GetCustomized() bool`

GetCustomized returns the Customized field if non-nil, zero value otherwise.

### GetCustomizedOk

`func (o *FieldMappingDto) GetCustomizedOk() (*bool, bool)`

GetCustomizedOk returns a tuple with the Customized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomized

`func (o *FieldMappingDto) SetCustomized(v bool)`

SetCustomized sets Customized field to given value.

### HasCustomized

`func (o *FieldMappingDto) HasCustomized() bool`

HasCustomized returns a boolean if a field has been set.

### GetRevision

`func (o *FieldMappingDto) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *FieldMappingDto) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *FieldMappingDto) SetRevision(v string)`

SetRevision sets Revision field to given value.


### GetUniversalFieldMappingId

`func (o *FieldMappingDto) GetUniversalFieldMappingId() string`

GetUniversalFieldMappingId returns the UniversalFieldMappingId field if non-nil, zero value otherwise.

### GetUniversalFieldMappingIdOk

`func (o *FieldMappingDto) GetUniversalFieldMappingIdOk() (*string, bool)`

GetUniversalFieldMappingIdOk returns a tuple with the UniversalFieldMappingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniversalFieldMappingId

`func (o *FieldMappingDto) SetUniversalFieldMappingId(v string)`

SetUniversalFieldMappingId sets UniversalFieldMappingId field to given value.

### HasUniversalFieldMappingId

`func (o *FieldMappingDto) HasUniversalFieldMappingId() bool`

HasUniversalFieldMappingId returns a boolean if a field has been set.

### GetUniversalFieldMappingRevision

`func (o *FieldMappingDto) GetUniversalFieldMappingRevision() string`

GetUniversalFieldMappingRevision returns the UniversalFieldMappingRevision field if non-nil, zero value otherwise.

### GetUniversalFieldMappingRevisionOk

`func (o *FieldMappingDto) GetUniversalFieldMappingRevisionOk() (*string, bool)`

GetUniversalFieldMappingRevisionOk returns a tuple with the UniversalFieldMappingRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniversalFieldMappingRevision

`func (o *FieldMappingDto) SetUniversalFieldMappingRevision(v string)`

SetUniversalFieldMappingRevision sets UniversalFieldMappingRevision field to given value.

### HasUniversalFieldMappingRevision

`func (o *FieldMappingDto) HasUniversalFieldMappingRevision() bool`

HasUniversalFieldMappingRevision returns a boolean if a field has been set.

### GetIntegrationId

`func (o *FieldMappingDto) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *FieldMappingDto) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *FieldMappingDto) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.

### HasIntegrationId

`func (o *FieldMappingDto) HasIntegrationId() bool`

HasIntegrationId returns a boolean if a field has been set.

### GetDataSourceKey

`func (o *FieldMappingDto) GetDataSourceKey() string`

GetDataSourceKey returns the DataSourceKey field if non-nil, zero value otherwise.

### GetDataSourceKeyOk

`func (o *FieldMappingDto) GetDataSourceKeyOk() (*string, bool)`

GetDataSourceKeyOk returns a tuple with the DataSourceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceKey

`func (o *FieldMappingDto) SetDataSourceKey(v string)`

SetDataSourceKey sets DataSourceKey field to given value.

### HasDataSourceKey

`func (o *FieldMappingDto) HasDataSourceKey() bool`

HasDataSourceKey returns a boolean if a field has been set.

### GetDataSourceId

`func (o *FieldMappingDto) GetDataSourceId() string`

GetDataSourceId returns the DataSourceId field if non-nil, zero value otherwise.

### GetDataSourceIdOk

`func (o *FieldMappingDto) GetDataSourceIdOk() (*string, bool)`

GetDataSourceIdOk returns a tuple with the DataSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceId

`func (o *FieldMappingDto) SetDataSourceId(v string)`

SetDataSourceId sets DataSourceId field to given value.

### HasDataSourceId

`func (o *FieldMappingDto) HasDataSourceId() bool`

HasDataSourceId returns a boolean if a field has been set.

### GetAppSchema

`func (o *FieldMappingDto) GetAppSchema() map[string]interface{}`

GetAppSchema returns the AppSchema field if non-nil, zero value otherwise.

### GetAppSchemaOk

`func (o *FieldMappingDto) GetAppSchemaOk() (*map[string]interface{}, bool)`

GetAppSchemaOk returns a tuple with the AppSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppSchema

`func (o *FieldMappingDto) SetAppSchema(v map[string]interface{})`

SetAppSchema sets AppSchema field to given value.

### HasAppSchema

`func (o *FieldMappingDto) HasAppSchema() bool`

HasAppSchema returns a boolean if a field has been set.

### GetDirection

`func (o *FieldMappingDto) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *FieldMappingDto) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *FieldMappingDto) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *FieldMappingDto) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetDefaultImportValue

`func (o *FieldMappingDto) GetDefaultImportValue() map[string]interface{}`

GetDefaultImportValue returns the DefaultImportValue field if non-nil, zero value otherwise.

### GetDefaultImportValueOk

`func (o *FieldMappingDto) GetDefaultImportValueOk() (*map[string]interface{}, bool)`

GetDefaultImportValueOk returns a tuple with the DefaultImportValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultImportValue

`func (o *FieldMappingDto) SetDefaultImportValue(v map[string]interface{})`

SetDefaultImportValue sets DefaultImportValue field to given value.

### HasDefaultImportValue

`func (o *FieldMappingDto) HasDefaultImportValue() bool`

HasDefaultImportValue returns a boolean if a field has been set.

### GetDefaultExportValue

`func (o *FieldMappingDto) GetDefaultExportValue() map[string]interface{}`

GetDefaultExportValue returns the DefaultExportValue field if non-nil, zero value otherwise.

### GetDefaultExportValueOk

`func (o *FieldMappingDto) GetDefaultExportValueOk() (*map[string]interface{}, bool)`

GetDefaultExportValueOk returns a tuple with the DefaultExportValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultExportValue

`func (o *FieldMappingDto) SetDefaultExportValue(v map[string]interface{})`

SetDefaultExportValue sets DefaultExportValue field to given value.

### HasDefaultExportValue

`func (o *FieldMappingDto) HasDefaultExportValue() bool`

HasDefaultExportValue returns a boolean if a field has been set.

### GetExportValue

`func (o *FieldMappingDto) GetExportValue() map[string]interface{}`

GetExportValue returns the ExportValue field if non-nil, zero value otherwise.

### GetExportValueOk

`func (o *FieldMappingDto) GetExportValueOk() (*map[string]interface{}, bool)`

GetExportValueOk returns a tuple with the ExportValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportValue

`func (o *FieldMappingDto) SetExportValue(v map[string]interface{})`

SetExportValue sets ExportValue field to given value.

### HasExportValue

`func (o *FieldMappingDto) HasExportValue() bool`

HasExportValue returns a boolean if a field has been set.

### GetImportValue

`func (o *FieldMappingDto) GetImportValue() map[string]interface{}`

GetImportValue returns the ImportValue field if non-nil, zero value otherwise.

### GetImportValueOk

`func (o *FieldMappingDto) GetImportValueOk() (*map[string]interface{}, bool)`

GetImportValueOk returns a tuple with the ImportValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportValue

`func (o *FieldMappingDto) SetImportValue(v map[string]interface{})`

SetImportValue sets ImportValue field to given value.

### HasImportValue

`func (o *FieldMappingDto) HasImportValue() bool`

HasImportValue returns a boolean if a field has been set.

### GetFrozenImportFields

`func (o *FieldMappingDto) GetFrozenImportFields() map[string]interface{}`

GetFrozenImportFields returns the FrozenImportFields field if non-nil, zero value otherwise.

### GetFrozenImportFieldsOk

`func (o *FieldMappingDto) GetFrozenImportFieldsOk() (*map[string]interface{}, bool)`

GetFrozenImportFieldsOk returns a tuple with the FrozenImportFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrozenImportFields

`func (o *FieldMappingDto) SetFrozenImportFields(v map[string]interface{})`

SetFrozenImportFields sets FrozenImportFields field to given value.

### HasFrozenImportFields

`func (o *FieldMappingDto) HasFrozenImportFields() bool`

HasFrozenImportFields returns a boolean if a field has been set.

### GetFrozenExportFields

`func (o *FieldMappingDto) GetFrozenExportFields() map[string]interface{}`

GetFrozenExportFields returns the FrozenExportFields field if non-nil, zero value otherwise.

### GetFrozenExportFieldsOk

`func (o *FieldMappingDto) GetFrozenExportFieldsOk() (*map[string]interface{}, bool)`

GetFrozenExportFieldsOk returns a tuple with the FrozenExportFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrozenExportFields

`func (o *FieldMappingDto) SetFrozenExportFields(v map[string]interface{})`

SetFrozenExportFields sets FrozenExportFields field to given value.

### HasFrozenExportFields

`func (o *FieldMappingDto) HasFrozenExportFields() bool`

HasFrozenExportFields returns a boolean if a field has been set.

### GetFrozenUnifiedExportFields

`func (o *FieldMappingDto) GetFrozenUnifiedExportFields() map[string]interface{}`

GetFrozenUnifiedExportFields returns the FrozenUnifiedExportFields field if non-nil, zero value otherwise.

### GetFrozenUnifiedExportFieldsOk

`func (o *FieldMappingDto) GetFrozenUnifiedExportFieldsOk() (*map[string]interface{}, bool)`

GetFrozenUnifiedExportFieldsOk returns a tuple with the FrozenUnifiedExportFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrozenUnifiedExportFields

`func (o *FieldMappingDto) SetFrozenUnifiedExportFields(v map[string]interface{})`

SetFrozenUnifiedExportFields sets FrozenUnifiedExportFields field to given value.

### HasFrozenUnifiedExportFields

`func (o *FieldMappingDto) HasFrozenUnifiedExportFields() bool`

HasFrozenUnifiedExportFields returns a boolean if a field has been set.

### GetAppliedToIntegrations

`func (o *FieldMappingDto) GetAppliedToIntegrations() []FiledMappingWithIntegrationDto`

GetAppliedToIntegrations returns the AppliedToIntegrations field if non-nil, zero value otherwise.

### GetAppliedToIntegrationsOk

`func (o *FieldMappingDto) GetAppliedToIntegrationsOk() (*[]FiledMappingWithIntegrationDto, bool)`

GetAppliedToIntegrationsOk returns a tuple with the AppliedToIntegrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedToIntegrations

`func (o *FieldMappingDto) SetAppliedToIntegrations(v []FiledMappingWithIntegrationDto)`

SetAppliedToIntegrations sets AppliedToIntegrations field to given value.

### HasAppliedToIntegrations

`func (o *FieldMappingDto) HasAppliedToIntegrations() bool`

HasAppliedToIntegrations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


