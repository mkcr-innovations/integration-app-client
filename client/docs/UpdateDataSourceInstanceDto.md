# UpdateDataSourceInstanceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** |  | [optional] 
**CollectionKey** | Pointer to **string** |  | [optional] 
**CollectionParameters** | Pointer to **map[string]interface{}** |  | [optional] 
**PullUpdatesIntervalSeconds** | Pointer to **float32** |  | [optional] 
**FullSyncIntervalSeconds** | Pointer to **float32** |  | [optional] 

## Methods

### NewUpdateDataSourceInstanceDto

`func NewUpdateDataSourceInstanceDto() *UpdateDataSourceInstanceDto`

NewUpdateDataSourceInstanceDto instantiates a new UpdateDataSourceInstanceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDataSourceInstanceDtoWithDefaults

`func NewUpdateDataSourceInstanceDtoWithDefaults() *UpdateDataSourceInstanceDto`

NewUpdateDataSourceInstanceDtoWithDefaults instantiates a new UpdateDataSourceInstanceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *UpdateDataSourceInstanceDto) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *UpdateDataSourceInstanceDto) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *UpdateDataSourceInstanceDto) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *UpdateDataSourceInstanceDto) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetCollectionKey

`func (o *UpdateDataSourceInstanceDto) GetCollectionKey() string`

GetCollectionKey returns the CollectionKey field if non-nil, zero value otherwise.

### GetCollectionKeyOk

`func (o *UpdateDataSourceInstanceDto) GetCollectionKeyOk() (*string, bool)`

GetCollectionKeyOk returns a tuple with the CollectionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionKey

`func (o *UpdateDataSourceInstanceDto) SetCollectionKey(v string)`

SetCollectionKey sets CollectionKey field to given value.

### HasCollectionKey

`func (o *UpdateDataSourceInstanceDto) HasCollectionKey() bool`

HasCollectionKey returns a boolean if a field has been set.

### GetCollectionParameters

`func (o *UpdateDataSourceInstanceDto) GetCollectionParameters() map[string]interface{}`

GetCollectionParameters returns the CollectionParameters field if non-nil, zero value otherwise.

### GetCollectionParametersOk

`func (o *UpdateDataSourceInstanceDto) GetCollectionParametersOk() (*map[string]interface{}, bool)`

GetCollectionParametersOk returns a tuple with the CollectionParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionParameters

`func (o *UpdateDataSourceInstanceDto) SetCollectionParameters(v map[string]interface{})`

SetCollectionParameters sets CollectionParameters field to given value.

### HasCollectionParameters

`func (o *UpdateDataSourceInstanceDto) HasCollectionParameters() bool`

HasCollectionParameters returns a boolean if a field has been set.

### GetPullUpdatesIntervalSeconds

`func (o *UpdateDataSourceInstanceDto) GetPullUpdatesIntervalSeconds() float32`

GetPullUpdatesIntervalSeconds returns the PullUpdatesIntervalSeconds field if non-nil, zero value otherwise.

### GetPullUpdatesIntervalSecondsOk

`func (o *UpdateDataSourceInstanceDto) GetPullUpdatesIntervalSecondsOk() (*float32, bool)`

GetPullUpdatesIntervalSecondsOk returns a tuple with the PullUpdatesIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullUpdatesIntervalSeconds

`func (o *UpdateDataSourceInstanceDto) SetPullUpdatesIntervalSeconds(v float32)`

SetPullUpdatesIntervalSeconds sets PullUpdatesIntervalSeconds field to given value.

### HasPullUpdatesIntervalSeconds

`func (o *UpdateDataSourceInstanceDto) HasPullUpdatesIntervalSeconds() bool`

HasPullUpdatesIntervalSeconds returns a boolean if a field has been set.

### GetFullSyncIntervalSeconds

`func (o *UpdateDataSourceInstanceDto) GetFullSyncIntervalSeconds() float32`

GetFullSyncIntervalSeconds returns the FullSyncIntervalSeconds field if non-nil, zero value otherwise.

### GetFullSyncIntervalSecondsOk

`func (o *UpdateDataSourceInstanceDto) GetFullSyncIntervalSecondsOk() (*float32, bool)`

GetFullSyncIntervalSecondsOk returns a tuple with the FullSyncIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullSyncIntervalSeconds

`func (o *UpdateDataSourceInstanceDto) SetFullSyncIntervalSeconds(v float32)`

SetFullSyncIntervalSeconds sets FullSyncIntervalSeconds field to given value.

### HasFullSyncIntervalSeconds

`func (o *UpdateDataSourceInstanceDto) HasFullSyncIntervalSeconds() bool`

HasFullSyncIntervalSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


