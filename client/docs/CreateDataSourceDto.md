# CreateDataSourceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Name** | **string** |  | 
**IntegrationId** | Pointer to **string** |  | [optional] 
**Udm** | Pointer to **string** |  | [optional] 
**PullUpdatesIntervalSeconds** | Pointer to **float32** |  | [optional] 
**FullSyncIntervalSeconds** | Pointer to **float32** |  | [optional] 
**DefaultPath** | Pointer to **string** |  | [optional] 
**CollectionKey** | Pointer to **map[string]interface{}** |  | [optional] 
**CollectionParameters** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCreateDataSourceDto

`func NewCreateDataSourceDto(key string, name string, ) *CreateDataSourceDto`

NewCreateDataSourceDto instantiates a new CreateDataSourceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDataSourceDtoWithDefaults

`func NewCreateDataSourceDtoWithDefaults() *CreateDataSourceDto`

NewCreateDataSourceDtoWithDefaults instantiates a new CreateDataSourceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *CreateDataSourceDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateDataSourceDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateDataSourceDto) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *CreateDataSourceDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateDataSourceDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateDataSourceDto) SetName(v string)`

SetName sets Name field to given value.


### GetIntegrationId

`func (o *CreateDataSourceDto) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *CreateDataSourceDto) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *CreateDataSourceDto) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.

### HasIntegrationId

`func (o *CreateDataSourceDto) HasIntegrationId() bool`

HasIntegrationId returns a boolean if a field has been set.

### GetUdm

`func (o *CreateDataSourceDto) GetUdm() string`

GetUdm returns the Udm field if non-nil, zero value otherwise.

### GetUdmOk

`func (o *CreateDataSourceDto) GetUdmOk() (*string, bool)`

GetUdmOk returns a tuple with the Udm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdm

`func (o *CreateDataSourceDto) SetUdm(v string)`

SetUdm sets Udm field to given value.

### HasUdm

`func (o *CreateDataSourceDto) HasUdm() bool`

HasUdm returns a boolean if a field has been set.

### GetPullUpdatesIntervalSeconds

`func (o *CreateDataSourceDto) GetPullUpdatesIntervalSeconds() float32`

GetPullUpdatesIntervalSeconds returns the PullUpdatesIntervalSeconds field if non-nil, zero value otherwise.

### GetPullUpdatesIntervalSecondsOk

`func (o *CreateDataSourceDto) GetPullUpdatesIntervalSecondsOk() (*float32, bool)`

GetPullUpdatesIntervalSecondsOk returns a tuple with the PullUpdatesIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullUpdatesIntervalSeconds

`func (o *CreateDataSourceDto) SetPullUpdatesIntervalSeconds(v float32)`

SetPullUpdatesIntervalSeconds sets PullUpdatesIntervalSeconds field to given value.

### HasPullUpdatesIntervalSeconds

`func (o *CreateDataSourceDto) HasPullUpdatesIntervalSeconds() bool`

HasPullUpdatesIntervalSeconds returns a boolean if a field has been set.

### GetFullSyncIntervalSeconds

`func (o *CreateDataSourceDto) GetFullSyncIntervalSeconds() float32`

GetFullSyncIntervalSeconds returns the FullSyncIntervalSeconds field if non-nil, zero value otherwise.

### GetFullSyncIntervalSecondsOk

`func (o *CreateDataSourceDto) GetFullSyncIntervalSecondsOk() (*float32, bool)`

GetFullSyncIntervalSecondsOk returns a tuple with the FullSyncIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullSyncIntervalSeconds

`func (o *CreateDataSourceDto) SetFullSyncIntervalSeconds(v float32)`

SetFullSyncIntervalSeconds sets FullSyncIntervalSeconds field to given value.

### HasFullSyncIntervalSeconds

`func (o *CreateDataSourceDto) HasFullSyncIntervalSeconds() bool`

HasFullSyncIntervalSeconds returns a boolean if a field has been set.

### GetDefaultPath

`func (o *CreateDataSourceDto) GetDefaultPath() string`

GetDefaultPath returns the DefaultPath field if non-nil, zero value otherwise.

### GetDefaultPathOk

`func (o *CreateDataSourceDto) GetDefaultPathOk() (*string, bool)`

GetDefaultPathOk returns a tuple with the DefaultPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPath

`func (o *CreateDataSourceDto) SetDefaultPath(v string)`

SetDefaultPath sets DefaultPath field to given value.

### HasDefaultPath

`func (o *CreateDataSourceDto) HasDefaultPath() bool`

HasDefaultPath returns a boolean if a field has been set.

### GetCollectionKey

`func (o *CreateDataSourceDto) GetCollectionKey() map[string]interface{}`

GetCollectionKey returns the CollectionKey field if non-nil, zero value otherwise.

### GetCollectionKeyOk

`func (o *CreateDataSourceDto) GetCollectionKeyOk() (*map[string]interface{}, bool)`

GetCollectionKeyOk returns a tuple with the CollectionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionKey

`func (o *CreateDataSourceDto) SetCollectionKey(v map[string]interface{})`

SetCollectionKey sets CollectionKey field to given value.

### HasCollectionKey

`func (o *CreateDataSourceDto) HasCollectionKey() bool`

HasCollectionKey returns a boolean if a field has been set.

### GetCollectionParameters

`func (o *CreateDataSourceDto) GetCollectionParameters() map[string]interface{}`

GetCollectionParameters returns the CollectionParameters field if non-nil, zero value otherwise.

### GetCollectionParametersOk

`func (o *CreateDataSourceDto) GetCollectionParametersOk() (*map[string]interface{}, bool)`

GetCollectionParametersOk returns a tuple with the CollectionParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionParameters

`func (o *CreateDataSourceDto) SetCollectionParameters(v map[string]interface{})`

SetCollectionParameters sets CollectionParameters field to given value.

### HasCollectionParameters

`func (o *CreateDataSourceDto) HasCollectionParameters() bool`

HasCollectionParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


