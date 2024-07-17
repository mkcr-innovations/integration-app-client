# UpdateDataSourceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ArchivedAt** | Pointer to **string** |  | [optional] 
**Udm** | Pointer to **string** |  | [optional] 
**PullUpdatesIntervalSeconds** | Pointer to **float32** |  | [optional] 
**FullSyncIntervalSeconds** | Pointer to **float32** |  | [optional] 
**DefaultPath** | Pointer to **string** |  | [optional] 
**CollectionKey** | Pointer to **string** |  | [optional] 
**CollectionParameters** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewUpdateDataSourceDto

`func NewUpdateDataSourceDto() *UpdateDataSourceDto`

NewUpdateDataSourceDto instantiates a new UpdateDataSourceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDataSourceDtoWithDefaults

`func NewUpdateDataSourceDtoWithDefaults() *UpdateDataSourceDto`

NewUpdateDataSourceDtoWithDefaults instantiates a new UpdateDataSourceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *UpdateDataSourceDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UpdateDataSourceDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UpdateDataSourceDto) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *UpdateDataSourceDto) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *UpdateDataSourceDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateDataSourceDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateDataSourceDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateDataSourceDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetArchivedAt

`func (o *UpdateDataSourceDto) GetArchivedAt() string`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *UpdateDataSourceDto) GetArchivedAtOk() (*string, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *UpdateDataSourceDto) SetArchivedAt(v string)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *UpdateDataSourceDto) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### GetUdm

`func (o *UpdateDataSourceDto) GetUdm() string`

GetUdm returns the Udm field if non-nil, zero value otherwise.

### GetUdmOk

`func (o *UpdateDataSourceDto) GetUdmOk() (*string, bool)`

GetUdmOk returns a tuple with the Udm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdm

`func (o *UpdateDataSourceDto) SetUdm(v string)`

SetUdm sets Udm field to given value.

### HasUdm

`func (o *UpdateDataSourceDto) HasUdm() bool`

HasUdm returns a boolean if a field has been set.

### GetPullUpdatesIntervalSeconds

`func (o *UpdateDataSourceDto) GetPullUpdatesIntervalSeconds() float32`

GetPullUpdatesIntervalSeconds returns the PullUpdatesIntervalSeconds field if non-nil, zero value otherwise.

### GetPullUpdatesIntervalSecondsOk

`func (o *UpdateDataSourceDto) GetPullUpdatesIntervalSecondsOk() (*float32, bool)`

GetPullUpdatesIntervalSecondsOk returns a tuple with the PullUpdatesIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullUpdatesIntervalSeconds

`func (o *UpdateDataSourceDto) SetPullUpdatesIntervalSeconds(v float32)`

SetPullUpdatesIntervalSeconds sets PullUpdatesIntervalSeconds field to given value.

### HasPullUpdatesIntervalSeconds

`func (o *UpdateDataSourceDto) HasPullUpdatesIntervalSeconds() bool`

HasPullUpdatesIntervalSeconds returns a boolean if a field has been set.

### GetFullSyncIntervalSeconds

`func (o *UpdateDataSourceDto) GetFullSyncIntervalSeconds() float32`

GetFullSyncIntervalSeconds returns the FullSyncIntervalSeconds field if non-nil, zero value otherwise.

### GetFullSyncIntervalSecondsOk

`func (o *UpdateDataSourceDto) GetFullSyncIntervalSecondsOk() (*float32, bool)`

GetFullSyncIntervalSecondsOk returns a tuple with the FullSyncIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullSyncIntervalSeconds

`func (o *UpdateDataSourceDto) SetFullSyncIntervalSeconds(v float32)`

SetFullSyncIntervalSeconds sets FullSyncIntervalSeconds field to given value.

### HasFullSyncIntervalSeconds

`func (o *UpdateDataSourceDto) HasFullSyncIntervalSeconds() bool`

HasFullSyncIntervalSeconds returns a boolean if a field has been set.

### GetDefaultPath

`func (o *UpdateDataSourceDto) GetDefaultPath() string`

GetDefaultPath returns the DefaultPath field if non-nil, zero value otherwise.

### GetDefaultPathOk

`func (o *UpdateDataSourceDto) GetDefaultPathOk() (*string, bool)`

GetDefaultPathOk returns a tuple with the DefaultPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPath

`func (o *UpdateDataSourceDto) SetDefaultPath(v string)`

SetDefaultPath sets DefaultPath field to given value.

### HasDefaultPath

`func (o *UpdateDataSourceDto) HasDefaultPath() bool`

HasDefaultPath returns a boolean if a field has been set.

### GetCollectionKey

`func (o *UpdateDataSourceDto) GetCollectionKey() string`

GetCollectionKey returns the CollectionKey field if non-nil, zero value otherwise.

### GetCollectionKeyOk

`func (o *UpdateDataSourceDto) GetCollectionKeyOk() (*string, bool)`

GetCollectionKeyOk returns a tuple with the CollectionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionKey

`func (o *UpdateDataSourceDto) SetCollectionKey(v string)`

SetCollectionKey sets CollectionKey field to given value.

### HasCollectionKey

`func (o *UpdateDataSourceDto) HasCollectionKey() bool`

HasCollectionKey returns a boolean if a field has been set.

### GetCollectionParameters

`func (o *UpdateDataSourceDto) GetCollectionParameters() map[string]interface{}`

GetCollectionParameters returns the CollectionParameters field if non-nil, zero value otherwise.

### GetCollectionParametersOk

`func (o *UpdateDataSourceDto) GetCollectionParametersOk() (*map[string]interface{}, bool)`

GetCollectionParametersOk returns a tuple with the CollectionParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionParameters

`func (o *UpdateDataSourceDto) SetCollectionParameters(v map[string]interface{})`

SetCollectionParameters sets CollectionParameters field to given value.

### HasCollectionParameters

`func (o *UpdateDataSourceDto) HasCollectionParameters() bool`

HasCollectionParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


