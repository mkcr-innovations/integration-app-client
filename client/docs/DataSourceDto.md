# DataSourceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Key** | **string** |  | 
**Name** | **string** |  | 
**ArchivedAt** | Pointer to **string** |  | [optional] 
**Revision** | **string** |  | 
**IntegrationId** | Pointer to **string** |  | [optional] 
**Customized** | Pointer to **bool** |  | [optional] 
**UniversalDataSourceId** | Pointer to **string** |  | [optional] 
**UniversalDataSourceRevision** | Pointer to **string** |  | [optional] 
**CollectionKey** | Pointer to **string** |  | [optional] 
**CollectionParameters** | Pointer to **map[string]interface{}** |  | [optional] 
**Udm** | Pointer to **string** |  | [optional] 
**PullUpdatesIntervalSeconds** | Pointer to **float32** |  | [optional] 
**FullSyncIntervalSeconds** | Pointer to **float32** |  | [optional] 
**AppliedToIntegrations** | Pointer to [**[]DataSourceWithIntegrationDto**](DataSourceWithIntegrationDto.md) |  | [optional] 
**DefaultPath** | Pointer to **string** |  | [optional] 

## Methods

### NewDataSourceDto

`func NewDataSourceDto(id string, key string, name string, revision string, ) *DataSourceDto`

NewDataSourceDto instantiates a new DataSourceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataSourceDtoWithDefaults

`func NewDataSourceDtoWithDefaults() *DataSourceDto`

NewDataSourceDtoWithDefaults instantiates a new DataSourceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DataSourceDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataSourceDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataSourceDto) SetId(v string)`

SetId sets Id field to given value.


### GetKey

`func (o *DataSourceDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *DataSourceDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *DataSourceDto) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *DataSourceDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataSourceDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataSourceDto) SetName(v string)`

SetName sets Name field to given value.


### GetArchivedAt

`func (o *DataSourceDto) GetArchivedAt() string`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *DataSourceDto) GetArchivedAtOk() (*string, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *DataSourceDto) SetArchivedAt(v string)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *DataSourceDto) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### GetRevision

`func (o *DataSourceDto) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *DataSourceDto) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *DataSourceDto) SetRevision(v string)`

SetRevision sets Revision field to given value.


### GetIntegrationId

`func (o *DataSourceDto) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *DataSourceDto) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *DataSourceDto) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.

### HasIntegrationId

`func (o *DataSourceDto) HasIntegrationId() bool`

HasIntegrationId returns a boolean if a field has been set.

### GetCustomized

`func (o *DataSourceDto) GetCustomized() bool`

GetCustomized returns the Customized field if non-nil, zero value otherwise.

### GetCustomizedOk

`func (o *DataSourceDto) GetCustomizedOk() (*bool, bool)`

GetCustomizedOk returns a tuple with the Customized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomized

`func (o *DataSourceDto) SetCustomized(v bool)`

SetCustomized sets Customized field to given value.

### HasCustomized

`func (o *DataSourceDto) HasCustomized() bool`

HasCustomized returns a boolean if a field has been set.

### GetUniversalDataSourceId

`func (o *DataSourceDto) GetUniversalDataSourceId() string`

GetUniversalDataSourceId returns the UniversalDataSourceId field if non-nil, zero value otherwise.

### GetUniversalDataSourceIdOk

`func (o *DataSourceDto) GetUniversalDataSourceIdOk() (*string, bool)`

GetUniversalDataSourceIdOk returns a tuple with the UniversalDataSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniversalDataSourceId

`func (o *DataSourceDto) SetUniversalDataSourceId(v string)`

SetUniversalDataSourceId sets UniversalDataSourceId field to given value.

### HasUniversalDataSourceId

`func (o *DataSourceDto) HasUniversalDataSourceId() bool`

HasUniversalDataSourceId returns a boolean if a field has been set.

### GetUniversalDataSourceRevision

`func (o *DataSourceDto) GetUniversalDataSourceRevision() string`

GetUniversalDataSourceRevision returns the UniversalDataSourceRevision field if non-nil, zero value otherwise.

### GetUniversalDataSourceRevisionOk

`func (o *DataSourceDto) GetUniversalDataSourceRevisionOk() (*string, bool)`

GetUniversalDataSourceRevisionOk returns a tuple with the UniversalDataSourceRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniversalDataSourceRevision

`func (o *DataSourceDto) SetUniversalDataSourceRevision(v string)`

SetUniversalDataSourceRevision sets UniversalDataSourceRevision field to given value.

### HasUniversalDataSourceRevision

`func (o *DataSourceDto) HasUniversalDataSourceRevision() bool`

HasUniversalDataSourceRevision returns a boolean if a field has been set.

### GetCollectionKey

`func (o *DataSourceDto) GetCollectionKey() string`

GetCollectionKey returns the CollectionKey field if non-nil, zero value otherwise.

### GetCollectionKeyOk

`func (o *DataSourceDto) GetCollectionKeyOk() (*string, bool)`

GetCollectionKeyOk returns a tuple with the CollectionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionKey

`func (o *DataSourceDto) SetCollectionKey(v string)`

SetCollectionKey sets CollectionKey field to given value.

### HasCollectionKey

`func (o *DataSourceDto) HasCollectionKey() bool`

HasCollectionKey returns a boolean if a field has been set.

### GetCollectionParameters

`func (o *DataSourceDto) GetCollectionParameters() map[string]interface{}`

GetCollectionParameters returns the CollectionParameters field if non-nil, zero value otherwise.

### GetCollectionParametersOk

`func (o *DataSourceDto) GetCollectionParametersOk() (*map[string]interface{}, bool)`

GetCollectionParametersOk returns a tuple with the CollectionParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionParameters

`func (o *DataSourceDto) SetCollectionParameters(v map[string]interface{})`

SetCollectionParameters sets CollectionParameters field to given value.

### HasCollectionParameters

`func (o *DataSourceDto) HasCollectionParameters() bool`

HasCollectionParameters returns a boolean if a field has been set.

### GetUdm

`func (o *DataSourceDto) GetUdm() string`

GetUdm returns the Udm field if non-nil, zero value otherwise.

### GetUdmOk

`func (o *DataSourceDto) GetUdmOk() (*string, bool)`

GetUdmOk returns a tuple with the Udm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdm

`func (o *DataSourceDto) SetUdm(v string)`

SetUdm sets Udm field to given value.

### HasUdm

`func (o *DataSourceDto) HasUdm() bool`

HasUdm returns a boolean if a field has been set.

### GetPullUpdatesIntervalSeconds

`func (o *DataSourceDto) GetPullUpdatesIntervalSeconds() float32`

GetPullUpdatesIntervalSeconds returns the PullUpdatesIntervalSeconds field if non-nil, zero value otherwise.

### GetPullUpdatesIntervalSecondsOk

`func (o *DataSourceDto) GetPullUpdatesIntervalSecondsOk() (*float32, bool)`

GetPullUpdatesIntervalSecondsOk returns a tuple with the PullUpdatesIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullUpdatesIntervalSeconds

`func (o *DataSourceDto) SetPullUpdatesIntervalSeconds(v float32)`

SetPullUpdatesIntervalSeconds sets PullUpdatesIntervalSeconds field to given value.

### HasPullUpdatesIntervalSeconds

`func (o *DataSourceDto) HasPullUpdatesIntervalSeconds() bool`

HasPullUpdatesIntervalSeconds returns a boolean if a field has been set.

### GetFullSyncIntervalSeconds

`func (o *DataSourceDto) GetFullSyncIntervalSeconds() float32`

GetFullSyncIntervalSeconds returns the FullSyncIntervalSeconds field if non-nil, zero value otherwise.

### GetFullSyncIntervalSecondsOk

`func (o *DataSourceDto) GetFullSyncIntervalSecondsOk() (*float32, bool)`

GetFullSyncIntervalSecondsOk returns a tuple with the FullSyncIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullSyncIntervalSeconds

`func (o *DataSourceDto) SetFullSyncIntervalSeconds(v float32)`

SetFullSyncIntervalSeconds sets FullSyncIntervalSeconds field to given value.

### HasFullSyncIntervalSeconds

`func (o *DataSourceDto) HasFullSyncIntervalSeconds() bool`

HasFullSyncIntervalSeconds returns a boolean if a field has been set.

### GetAppliedToIntegrations

`func (o *DataSourceDto) GetAppliedToIntegrations() []DataSourceWithIntegrationDto`

GetAppliedToIntegrations returns the AppliedToIntegrations field if non-nil, zero value otherwise.

### GetAppliedToIntegrationsOk

`func (o *DataSourceDto) GetAppliedToIntegrationsOk() (*[]DataSourceWithIntegrationDto, bool)`

GetAppliedToIntegrationsOk returns a tuple with the AppliedToIntegrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedToIntegrations

`func (o *DataSourceDto) SetAppliedToIntegrations(v []DataSourceWithIntegrationDto)`

SetAppliedToIntegrations sets AppliedToIntegrations field to given value.

### HasAppliedToIntegrations

`func (o *DataSourceDto) HasAppliedToIntegrations() bool`

HasAppliedToIntegrations returns a boolean if a field has been set.

### GetDefaultPath

`func (o *DataSourceDto) GetDefaultPath() string`

GetDefaultPath returns the DefaultPath field if non-nil, zero value otherwise.

### GetDefaultPathOk

`func (o *DataSourceDto) GetDefaultPathOk() (*string, bool)`

GetDefaultPathOk returns a tuple with the DefaultPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPath

`func (o *DataSourceDto) SetDefaultPath(v string)`

SetDefaultPath sets DefaultPath field to given value.

### HasDefaultPath

`func (o *DataSourceDto) HasDefaultPath() bool`

HasDefaultPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


