# UpdateDataSourceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Udm** | Pointer to **string** |  | [optional] 
**PullUpdatesIntervalSeconds** | Pointer to **float32** |  | [optional] 
**FullSyncIntervalSeconds** | Pointer to **float32** |  | [optional] 
**DefaultPath** | Pointer to **string** |  | [optional] 
**CollectionKey** | Pointer to **map[string]interface{}** |  | [optional] 
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

`func (o *UpdateDataSourceDto) GetCollectionKey() map[string]interface{}`

GetCollectionKey returns the CollectionKey field if non-nil, zero value otherwise.

### GetCollectionKeyOk

`func (o *UpdateDataSourceDto) GetCollectionKeyOk() (*map[string]interface{}, bool)`

GetCollectionKeyOk returns a tuple with the CollectionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionKey

`func (o *UpdateDataSourceDto) SetCollectionKey(v map[string]interface{})`

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


