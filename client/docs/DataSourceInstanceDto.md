# DataSourceInstanceDto

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
**DataSourceId** | Pointer to **string** |  | [optional] 
**DataSourceRevision** | Pointer to **string** |  | [optional] 
**UniversalDataSourceId** | Pointer to **string** |  | [optional] 
**DataSource** | Pointer to [**DataSourceDto**](DataSourceDto.md) |  | [optional] 
**CollectionKey** | Pointer to **string** |  | [optional] 
**CollectionParameters** | Pointer to **map[string]interface{}** |  | [optional] 
**DefaultCollectionKey** | Pointer to **string** |  | [optional] 
**DefaultCollectionParameters** | Pointer to **map[string]interface{}** |  | [optional] 
**Path** | **string** |  | 
**DefaultPath** | **string** |  | 
**CollectionSpec** | Pointer to **map[string]interface{}** |  | [optional] 
**IsCustomized** | **bool** |  | 
**Udm** | Pointer to **string** |  | [optional] 
**Uuid** | **string** |  | 
**Error** | Pointer to **map[string]interface{}** |  | [optional] 
**Subscriptions** | [**DataSourceInstanceDtoSubscriptions**](DataSourceInstanceDtoSubscriptions.md) |  | 
**PullUpdatesIntervalSeconds** | Pointer to **float32** |  | [optional] 
**NextPullUpdatesTimestamp** | Pointer to **float32** |  | [optional] 
**FullSyncIntervalSeconds** | Pointer to **float32** |  | [optional] 
**NextFullSyncTimestamp** | Pointer to **float32** |  | [optional] 

## Methods

### NewDataSourceInstanceDto

`func NewDataSourceInstanceDto(id string, name string, revision string, userId string, connectionId string, integrationId string, path string, defaultPath string, isCustomized bool, uuid string, subscriptions DataSourceInstanceDtoSubscriptions, ) *DataSourceInstanceDto`

NewDataSourceInstanceDto instantiates a new DataSourceInstanceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataSourceInstanceDtoWithDefaults

`func NewDataSourceInstanceDtoWithDefaults() *DataSourceInstanceDto`

NewDataSourceInstanceDtoWithDefaults instantiates a new DataSourceInstanceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DataSourceInstanceDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataSourceInstanceDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataSourceInstanceDto) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *DataSourceInstanceDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataSourceInstanceDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataSourceInstanceDto) SetName(v string)`

SetName sets Name field to given value.


### GetRevision

`func (o *DataSourceInstanceDto) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *DataSourceInstanceDto) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *DataSourceInstanceDto) SetRevision(v string)`

SetRevision sets Revision field to given value.


### GetUserId

`func (o *DataSourceInstanceDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *DataSourceInstanceDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *DataSourceInstanceDto) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetUser

`func (o *DataSourceInstanceDto) GetUser() CustomerDto`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *DataSourceInstanceDto) GetUserOk() (*CustomerDto, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *DataSourceInstanceDto) SetUser(v CustomerDto)`

SetUser sets User field to given value.

### HasUser

`func (o *DataSourceInstanceDto) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetConnectionId

`func (o *DataSourceInstanceDto) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *DataSourceInstanceDto) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *DataSourceInstanceDto) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetConnection

`func (o *DataSourceInstanceDto) GetConnection() ConnectionDto`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *DataSourceInstanceDto) GetConnectionOk() (*ConnectionDto, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *DataSourceInstanceDto) SetConnection(v ConnectionDto)`

SetConnection sets Connection field to given value.

### HasConnection

`func (o *DataSourceInstanceDto) HasConnection() bool`

HasConnection returns a boolean if a field has been set.

### GetIntegrationId

`func (o *DataSourceInstanceDto) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *DataSourceInstanceDto) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *DataSourceInstanceDto) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.


### GetIntegration

`func (o *DataSourceInstanceDto) GetIntegration() IntegrationDto`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *DataSourceInstanceDto) GetIntegrationOk() (*IntegrationDto, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *DataSourceInstanceDto) SetIntegration(v IntegrationDto)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *DataSourceInstanceDto) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetInstanceKey

`func (o *DataSourceInstanceDto) GetInstanceKey() string`

GetInstanceKey returns the InstanceKey field if non-nil, zero value otherwise.

### GetInstanceKeyOk

`func (o *DataSourceInstanceDto) GetInstanceKeyOk() (*string, bool)`

GetInstanceKeyOk returns a tuple with the InstanceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceKey

`func (o *DataSourceInstanceDto) SetInstanceKey(v string)`

SetInstanceKey sets InstanceKey field to given value.

### HasInstanceKey

`func (o *DataSourceInstanceDto) HasInstanceKey() bool`

HasInstanceKey returns a boolean if a field has been set.

### GetDataSourceId

`func (o *DataSourceInstanceDto) GetDataSourceId() string`

GetDataSourceId returns the DataSourceId field if non-nil, zero value otherwise.

### GetDataSourceIdOk

`func (o *DataSourceInstanceDto) GetDataSourceIdOk() (*string, bool)`

GetDataSourceIdOk returns a tuple with the DataSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceId

`func (o *DataSourceInstanceDto) SetDataSourceId(v string)`

SetDataSourceId sets DataSourceId field to given value.

### HasDataSourceId

`func (o *DataSourceInstanceDto) HasDataSourceId() bool`

HasDataSourceId returns a boolean if a field has been set.

### GetDataSourceRevision

`func (o *DataSourceInstanceDto) GetDataSourceRevision() string`

GetDataSourceRevision returns the DataSourceRevision field if non-nil, zero value otherwise.

### GetDataSourceRevisionOk

`func (o *DataSourceInstanceDto) GetDataSourceRevisionOk() (*string, bool)`

GetDataSourceRevisionOk returns a tuple with the DataSourceRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceRevision

`func (o *DataSourceInstanceDto) SetDataSourceRevision(v string)`

SetDataSourceRevision sets DataSourceRevision field to given value.

### HasDataSourceRevision

`func (o *DataSourceInstanceDto) HasDataSourceRevision() bool`

HasDataSourceRevision returns a boolean if a field has been set.

### GetUniversalDataSourceId

`func (o *DataSourceInstanceDto) GetUniversalDataSourceId() string`

GetUniversalDataSourceId returns the UniversalDataSourceId field if non-nil, zero value otherwise.

### GetUniversalDataSourceIdOk

`func (o *DataSourceInstanceDto) GetUniversalDataSourceIdOk() (*string, bool)`

GetUniversalDataSourceIdOk returns a tuple with the UniversalDataSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniversalDataSourceId

`func (o *DataSourceInstanceDto) SetUniversalDataSourceId(v string)`

SetUniversalDataSourceId sets UniversalDataSourceId field to given value.

### HasUniversalDataSourceId

`func (o *DataSourceInstanceDto) HasUniversalDataSourceId() bool`

HasUniversalDataSourceId returns a boolean if a field has been set.

### GetDataSource

`func (o *DataSourceInstanceDto) GetDataSource() DataSourceDto`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *DataSourceInstanceDto) GetDataSourceOk() (*DataSourceDto, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *DataSourceInstanceDto) SetDataSource(v DataSourceDto)`

SetDataSource sets DataSource field to given value.

### HasDataSource

`func (o *DataSourceInstanceDto) HasDataSource() bool`

HasDataSource returns a boolean if a field has been set.

### GetCollectionKey

`func (o *DataSourceInstanceDto) GetCollectionKey() string`

GetCollectionKey returns the CollectionKey field if non-nil, zero value otherwise.

### GetCollectionKeyOk

`func (o *DataSourceInstanceDto) GetCollectionKeyOk() (*string, bool)`

GetCollectionKeyOk returns a tuple with the CollectionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionKey

`func (o *DataSourceInstanceDto) SetCollectionKey(v string)`

SetCollectionKey sets CollectionKey field to given value.

### HasCollectionKey

`func (o *DataSourceInstanceDto) HasCollectionKey() bool`

HasCollectionKey returns a boolean if a field has been set.

### GetCollectionParameters

`func (o *DataSourceInstanceDto) GetCollectionParameters() map[string]interface{}`

GetCollectionParameters returns the CollectionParameters field if non-nil, zero value otherwise.

### GetCollectionParametersOk

`func (o *DataSourceInstanceDto) GetCollectionParametersOk() (*map[string]interface{}, bool)`

GetCollectionParametersOk returns a tuple with the CollectionParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionParameters

`func (o *DataSourceInstanceDto) SetCollectionParameters(v map[string]interface{})`

SetCollectionParameters sets CollectionParameters field to given value.

### HasCollectionParameters

`func (o *DataSourceInstanceDto) HasCollectionParameters() bool`

HasCollectionParameters returns a boolean if a field has been set.

### GetDefaultCollectionKey

`func (o *DataSourceInstanceDto) GetDefaultCollectionKey() string`

GetDefaultCollectionKey returns the DefaultCollectionKey field if non-nil, zero value otherwise.

### GetDefaultCollectionKeyOk

`func (o *DataSourceInstanceDto) GetDefaultCollectionKeyOk() (*string, bool)`

GetDefaultCollectionKeyOk returns a tuple with the DefaultCollectionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCollectionKey

`func (o *DataSourceInstanceDto) SetDefaultCollectionKey(v string)`

SetDefaultCollectionKey sets DefaultCollectionKey field to given value.

### HasDefaultCollectionKey

`func (o *DataSourceInstanceDto) HasDefaultCollectionKey() bool`

HasDefaultCollectionKey returns a boolean if a field has been set.

### GetDefaultCollectionParameters

`func (o *DataSourceInstanceDto) GetDefaultCollectionParameters() map[string]interface{}`

GetDefaultCollectionParameters returns the DefaultCollectionParameters field if non-nil, zero value otherwise.

### GetDefaultCollectionParametersOk

`func (o *DataSourceInstanceDto) GetDefaultCollectionParametersOk() (*map[string]interface{}, bool)`

GetDefaultCollectionParametersOk returns a tuple with the DefaultCollectionParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCollectionParameters

`func (o *DataSourceInstanceDto) SetDefaultCollectionParameters(v map[string]interface{})`

SetDefaultCollectionParameters sets DefaultCollectionParameters field to given value.

### HasDefaultCollectionParameters

`func (o *DataSourceInstanceDto) HasDefaultCollectionParameters() bool`

HasDefaultCollectionParameters returns a boolean if a field has been set.

### GetPath

`func (o *DataSourceInstanceDto) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DataSourceInstanceDto) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DataSourceInstanceDto) SetPath(v string)`

SetPath sets Path field to given value.


### GetDefaultPath

`func (o *DataSourceInstanceDto) GetDefaultPath() string`

GetDefaultPath returns the DefaultPath field if non-nil, zero value otherwise.

### GetDefaultPathOk

`func (o *DataSourceInstanceDto) GetDefaultPathOk() (*string, bool)`

GetDefaultPathOk returns a tuple with the DefaultPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPath

`func (o *DataSourceInstanceDto) SetDefaultPath(v string)`

SetDefaultPath sets DefaultPath field to given value.


### GetCollectionSpec

`func (o *DataSourceInstanceDto) GetCollectionSpec() map[string]interface{}`

GetCollectionSpec returns the CollectionSpec field if non-nil, zero value otherwise.

### GetCollectionSpecOk

`func (o *DataSourceInstanceDto) GetCollectionSpecOk() (*map[string]interface{}, bool)`

GetCollectionSpecOk returns a tuple with the CollectionSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionSpec

`func (o *DataSourceInstanceDto) SetCollectionSpec(v map[string]interface{})`

SetCollectionSpec sets CollectionSpec field to given value.

### HasCollectionSpec

`func (o *DataSourceInstanceDto) HasCollectionSpec() bool`

HasCollectionSpec returns a boolean if a field has been set.

### GetIsCustomized

`func (o *DataSourceInstanceDto) GetIsCustomized() bool`

GetIsCustomized returns the IsCustomized field if non-nil, zero value otherwise.

### GetIsCustomizedOk

`func (o *DataSourceInstanceDto) GetIsCustomizedOk() (*bool, bool)`

GetIsCustomizedOk returns a tuple with the IsCustomized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCustomized

`func (o *DataSourceInstanceDto) SetIsCustomized(v bool)`

SetIsCustomized sets IsCustomized field to given value.


### GetUdm

`func (o *DataSourceInstanceDto) GetUdm() string`

GetUdm returns the Udm field if non-nil, zero value otherwise.

### GetUdmOk

`func (o *DataSourceInstanceDto) GetUdmOk() (*string, bool)`

GetUdmOk returns a tuple with the Udm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdm

`func (o *DataSourceInstanceDto) SetUdm(v string)`

SetUdm sets Udm field to given value.

### HasUdm

`func (o *DataSourceInstanceDto) HasUdm() bool`

HasUdm returns a boolean if a field has been set.

### GetUuid

`func (o *DataSourceInstanceDto) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *DataSourceInstanceDto) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *DataSourceInstanceDto) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetError

`func (o *DataSourceInstanceDto) GetError() map[string]interface{}`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *DataSourceInstanceDto) GetErrorOk() (*map[string]interface{}, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *DataSourceInstanceDto) SetError(v map[string]interface{})`

SetError sets Error field to given value.

### HasError

`func (o *DataSourceInstanceDto) HasError() bool`

HasError returns a boolean if a field has been set.

### GetSubscriptions

`func (o *DataSourceInstanceDto) GetSubscriptions() DataSourceInstanceDtoSubscriptions`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *DataSourceInstanceDto) GetSubscriptionsOk() (*DataSourceInstanceDtoSubscriptions, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *DataSourceInstanceDto) SetSubscriptions(v DataSourceInstanceDtoSubscriptions)`

SetSubscriptions sets Subscriptions field to given value.


### GetPullUpdatesIntervalSeconds

`func (o *DataSourceInstanceDto) GetPullUpdatesIntervalSeconds() float32`

GetPullUpdatesIntervalSeconds returns the PullUpdatesIntervalSeconds field if non-nil, zero value otherwise.

### GetPullUpdatesIntervalSecondsOk

`func (o *DataSourceInstanceDto) GetPullUpdatesIntervalSecondsOk() (*float32, bool)`

GetPullUpdatesIntervalSecondsOk returns a tuple with the PullUpdatesIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullUpdatesIntervalSeconds

`func (o *DataSourceInstanceDto) SetPullUpdatesIntervalSeconds(v float32)`

SetPullUpdatesIntervalSeconds sets PullUpdatesIntervalSeconds field to given value.

### HasPullUpdatesIntervalSeconds

`func (o *DataSourceInstanceDto) HasPullUpdatesIntervalSeconds() bool`

HasPullUpdatesIntervalSeconds returns a boolean if a field has been set.

### GetNextPullUpdatesTimestamp

`func (o *DataSourceInstanceDto) GetNextPullUpdatesTimestamp() float32`

GetNextPullUpdatesTimestamp returns the NextPullUpdatesTimestamp field if non-nil, zero value otherwise.

### GetNextPullUpdatesTimestampOk

`func (o *DataSourceInstanceDto) GetNextPullUpdatesTimestampOk() (*float32, bool)`

GetNextPullUpdatesTimestampOk returns a tuple with the NextPullUpdatesTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPullUpdatesTimestamp

`func (o *DataSourceInstanceDto) SetNextPullUpdatesTimestamp(v float32)`

SetNextPullUpdatesTimestamp sets NextPullUpdatesTimestamp field to given value.

### HasNextPullUpdatesTimestamp

`func (o *DataSourceInstanceDto) HasNextPullUpdatesTimestamp() bool`

HasNextPullUpdatesTimestamp returns a boolean if a field has been set.

### GetFullSyncIntervalSeconds

`func (o *DataSourceInstanceDto) GetFullSyncIntervalSeconds() float32`

GetFullSyncIntervalSeconds returns the FullSyncIntervalSeconds field if non-nil, zero value otherwise.

### GetFullSyncIntervalSecondsOk

`func (o *DataSourceInstanceDto) GetFullSyncIntervalSecondsOk() (*float32, bool)`

GetFullSyncIntervalSecondsOk returns a tuple with the FullSyncIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullSyncIntervalSeconds

`func (o *DataSourceInstanceDto) SetFullSyncIntervalSeconds(v float32)`

SetFullSyncIntervalSeconds sets FullSyncIntervalSeconds field to given value.

### HasFullSyncIntervalSeconds

`func (o *DataSourceInstanceDto) HasFullSyncIntervalSeconds() bool`

HasFullSyncIntervalSeconds returns a boolean if a field has been set.

### GetNextFullSyncTimestamp

`func (o *DataSourceInstanceDto) GetNextFullSyncTimestamp() float32`

GetNextFullSyncTimestamp returns the NextFullSyncTimestamp field if non-nil, zero value otherwise.

### GetNextFullSyncTimestampOk

`func (o *DataSourceInstanceDto) GetNextFullSyncTimestampOk() (*float32, bool)`

GetNextFullSyncTimestampOk returns a tuple with the NextFullSyncTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextFullSyncTimestamp

`func (o *DataSourceInstanceDto) SetNextFullSyncTimestamp(v float32)`

SetNextFullSyncTimestamp sets NextFullSyncTimestamp field to given value.

### HasNextFullSyncTimestamp

`func (o *DataSourceInstanceDto) HasNextFullSyncTimestamp() bool`

HasNextFullSyncTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


