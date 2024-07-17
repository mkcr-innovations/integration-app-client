# DataSourceExportDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**IntegrationKeys** | Pointer to **[]string** |  | [optional] 
**CollectionKey** | Pointer to **string** |  | [optional] 
**CollectionParameters** | Pointer to **map[string]interface{}** |  | [optional] 
**Udm** | Pointer to **string** |  | [optional] 
**PullUpdatesIntervalSeconds** | Pointer to **float32** |  | [optional] 
**FullSyncIntervalSeconds** | Pointer to **float32** |  | [optional] 

## Methods

### NewDataSourceExportDto

`func NewDataSourceExportDto(name string, ) *DataSourceExportDto`

NewDataSourceExportDto instantiates a new DataSourceExportDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataSourceExportDtoWithDefaults

`func NewDataSourceExportDtoWithDefaults() *DataSourceExportDto`

NewDataSourceExportDtoWithDefaults instantiates a new DataSourceExportDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *DataSourceExportDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *DataSourceExportDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *DataSourceExportDto) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *DataSourceExportDto) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *DataSourceExportDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataSourceExportDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataSourceExportDto) SetName(v string)`

SetName sets Name field to given value.


### GetIntegrationKeys

`func (o *DataSourceExportDto) GetIntegrationKeys() []string`

GetIntegrationKeys returns the IntegrationKeys field if non-nil, zero value otherwise.

### GetIntegrationKeysOk

`func (o *DataSourceExportDto) GetIntegrationKeysOk() (*[]string, bool)`

GetIntegrationKeysOk returns a tuple with the IntegrationKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationKeys

`func (o *DataSourceExportDto) SetIntegrationKeys(v []string)`

SetIntegrationKeys sets IntegrationKeys field to given value.

### HasIntegrationKeys

`func (o *DataSourceExportDto) HasIntegrationKeys() bool`

HasIntegrationKeys returns a boolean if a field has been set.

### GetCollectionKey

`func (o *DataSourceExportDto) GetCollectionKey() string`

GetCollectionKey returns the CollectionKey field if non-nil, zero value otherwise.

### GetCollectionKeyOk

`func (o *DataSourceExportDto) GetCollectionKeyOk() (*string, bool)`

GetCollectionKeyOk returns a tuple with the CollectionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionKey

`func (o *DataSourceExportDto) SetCollectionKey(v string)`

SetCollectionKey sets CollectionKey field to given value.

### HasCollectionKey

`func (o *DataSourceExportDto) HasCollectionKey() bool`

HasCollectionKey returns a boolean if a field has been set.

### GetCollectionParameters

`func (o *DataSourceExportDto) GetCollectionParameters() map[string]interface{}`

GetCollectionParameters returns the CollectionParameters field if non-nil, zero value otherwise.

### GetCollectionParametersOk

`func (o *DataSourceExportDto) GetCollectionParametersOk() (*map[string]interface{}, bool)`

GetCollectionParametersOk returns a tuple with the CollectionParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionParameters

`func (o *DataSourceExportDto) SetCollectionParameters(v map[string]interface{})`

SetCollectionParameters sets CollectionParameters field to given value.

### HasCollectionParameters

`func (o *DataSourceExportDto) HasCollectionParameters() bool`

HasCollectionParameters returns a boolean if a field has been set.

### GetUdm

`func (o *DataSourceExportDto) GetUdm() string`

GetUdm returns the Udm field if non-nil, zero value otherwise.

### GetUdmOk

`func (o *DataSourceExportDto) GetUdmOk() (*string, bool)`

GetUdmOk returns a tuple with the Udm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdm

`func (o *DataSourceExportDto) SetUdm(v string)`

SetUdm sets Udm field to given value.

### HasUdm

`func (o *DataSourceExportDto) HasUdm() bool`

HasUdm returns a boolean if a field has been set.

### GetPullUpdatesIntervalSeconds

`func (o *DataSourceExportDto) GetPullUpdatesIntervalSeconds() float32`

GetPullUpdatesIntervalSeconds returns the PullUpdatesIntervalSeconds field if non-nil, zero value otherwise.

### GetPullUpdatesIntervalSecondsOk

`func (o *DataSourceExportDto) GetPullUpdatesIntervalSecondsOk() (*float32, bool)`

GetPullUpdatesIntervalSecondsOk returns a tuple with the PullUpdatesIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullUpdatesIntervalSeconds

`func (o *DataSourceExportDto) SetPullUpdatesIntervalSeconds(v float32)`

SetPullUpdatesIntervalSeconds sets PullUpdatesIntervalSeconds field to given value.

### HasPullUpdatesIntervalSeconds

`func (o *DataSourceExportDto) HasPullUpdatesIntervalSeconds() bool`

HasPullUpdatesIntervalSeconds returns a boolean if a field has been set.

### GetFullSyncIntervalSeconds

`func (o *DataSourceExportDto) GetFullSyncIntervalSeconds() float32`

GetFullSyncIntervalSeconds returns the FullSyncIntervalSeconds field if non-nil, zero value otherwise.

### GetFullSyncIntervalSecondsOk

`func (o *DataSourceExportDto) GetFullSyncIntervalSecondsOk() (*float32, bool)`

GetFullSyncIntervalSecondsOk returns a tuple with the FullSyncIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullSyncIntervalSeconds

`func (o *DataSourceExportDto) SetFullSyncIntervalSeconds(v float32)`

SetFullSyncIntervalSeconds sets FullSyncIntervalSeconds field to given value.

### HasFullSyncIntervalSeconds

`func (o *DataSourceExportDto) HasFullSyncIntervalSeconds() bool`

HasFullSyncIntervalSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


