# DataSourceSyncDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**UserId** | **string** |  | 
**User** | Pointer to [**CustomerDto**](CustomerDto.md) |  | [optional] 
**DataSourceId** | **string** |  | 
**DataSource** | Pointer to [**DataSourceDto**](DataSourceDto.md) |  | [optional] 
**DataSourceInstanceId** | **string** |  | 
**DataSourceInstance** | Pointer to [**DataSourceInstanceDto**](DataSourceInstanceDto.md) |  | [optional] 
**ConnectionId** | **string** |  | 
**Connection** | Pointer to [**ConnectionDto**](ConnectionDto.md) |  | [optional] 
**IntegrationId** | **string** |  | 
**Integration** | Pointer to [**IntegrationDto**](IntegrationDto.md) |  | [optional] 
**Status** | **string** |  | 
**StartDatetime** | **string** |  | 
**EndDatetime** | **string** |  | 
**IsFullScan** | **bool** |  | 
**CollectedEventIds** | **[]string** |  | 
**Error** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewDataSourceSyncDto

`func NewDataSourceSyncDto(id string, userId string, dataSourceId string, dataSourceInstanceId string, connectionId string, integrationId string, status string, startDatetime string, endDatetime string, isFullScan bool, collectedEventIds []string, ) *DataSourceSyncDto`

NewDataSourceSyncDto instantiates a new DataSourceSyncDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataSourceSyncDtoWithDefaults

`func NewDataSourceSyncDtoWithDefaults() *DataSourceSyncDto`

NewDataSourceSyncDtoWithDefaults instantiates a new DataSourceSyncDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DataSourceSyncDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataSourceSyncDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataSourceSyncDto) SetId(v string)`

SetId sets Id field to given value.


### GetUserId

`func (o *DataSourceSyncDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *DataSourceSyncDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *DataSourceSyncDto) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetUser

`func (o *DataSourceSyncDto) GetUser() CustomerDto`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *DataSourceSyncDto) GetUserOk() (*CustomerDto, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *DataSourceSyncDto) SetUser(v CustomerDto)`

SetUser sets User field to given value.

### HasUser

`func (o *DataSourceSyncDto) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetDataSourceId

`func (o *DataSourceSyncDto) GetDataSourceId() string`

GetDataSourceId returns the DataSourceId field if non-nil, zero value otherwise.

### GetDataSourceIdOk

`func (o *DataSourceSyncDto) GetDataSourceIdOk() (*string, bool)`

GetDataSourceIdOk returns a tuple with the DataSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceId

`func (o *DataSourceSyncDto) SetDataSourceId(v string)`

SetDataSourceId sets DataSourceId field to given value.


### GetDataSource

`func (o *DataSourceSyncDto) GetDataSource() DataSourceDto`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *DataSourceSyncDto) GetDataSourceOk() (*DataSourceDto, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *DataSourceSyncDto) SetDataSource(v DataSourceDto)`

SetDataSource sets DataSource field to given value.

### HasDataSource

`func (o *DataSourceSyncDto) HasDataSource() bool`

HasDataSource returns a boolean if a field has been set.

### GetDataSourceInstanceId

`func (o *DataSourceSyncDto) GetDataSourceInstanceId() string`

GetDataSourceInstanceId returns the DataSourceInstanceId field if non-nil, zero value otherwise.

### GetDataSourceInstanceIdOk

`func (o *DataSourceSyncDto) GetDataSourceInstanceIdOk() (*string, bool)`

GetDataSourceInstanceIdOk returns a tuple with the DataSourceInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceInstanceId

`func (o *DataSourceSyncDto) SetDataSourceInstanceId(v string)`

SetDataSourceInstanceId sets DataSourceInstanceId field to given value.


### GetDataSourceInstance

`func (o *DataSourceSyncDto) GetDataSourceInstance() DataSourceInstanceDto`

GetDataSourceInstance returns the DataSourceInstance field if non-nil, zero value otherwise.

### GetDataSourceInstanceOk

`func (o *DataSourceSyncDto) GetDataSourceInstanceOk() (*DataSourceInstanceDto, bool)`

GetDataSourceInstanceOk returns a tuple with the DataSourceInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceInstance

`func (o *DataSourceSyncDto) SetDataSourceInstance(v DataSourceInstanceDto)`

SetDataSourceInstance sets DataSourceInstance field to given value.

### HasDataSourceInstance

`func (o *DataSourceSyncDto) HasDataSourceInstance() bool`

HasDataSourceInstance returns a boolean if a field has been set.

### GetConnectionId

`func (o *DataSourceSyncDto) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *DataSourceSyncDto) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *DataSourceSyncDto) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetConnection

`func (o *DataSourceSyncDto) GetConnection() ConnectionDto`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *DataSourceSyncDto) GetConnectionOk() (*ConnectionDto, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *DataSourceSyncDto) SetConnection(v ConnectionDto)`

SetConnection sets Connection field to given value.

### HasConnection

`func (o *DataSourceSyncDto) HasConnection() bool`

HasConnection returns a boolean if a field has been set.

### GetIntegrationId

`func (o *DataSourceSyncDto) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *DataSourceSyncDto) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *DataSourceSyncDto) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.


### GetIntegration

`func (o *DataSourceSyncDto) GetIntegration() IntegrationDto`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *DataSourceSyncDto) GetIntegrationOk() (*IntegrationDto, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *DataSourceSyncDto) SetIntegration(v IntegrationDto)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *DataSourceSyncDto) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetStatus

`func (o *DataSourceSyncDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DataSourceSyncDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DataSourceSyncDto) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStartDatetime

`func (o *DataSourceSyncDto) GetStartDatetime() string`

GetStartDatetime returns the StartDatetime field if non-nil, zero value otherwise.

### GetStartDatetimeOk

`func (o *DataSourceSyncDto) GetStartDatetimeOk() (*string, bool)`

GetStartDatetimeOk returns a tuple with the StartDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDatetime

`func (o *DataSourceSyncDto) SetStartDatetime(v string)`

SetStartDatetime sets StartDatetime field to given value.


### GetEndDatetime

`func (o *DataSourceSyncDto) GetEndDatetime() string`

GetEndDatetime returns the EndDatetime field if non-nil, zero value otherwise.

### GetEndDatetimeOk

`func (o *DataSourceSyncDto) GetEndDatetimeOk() (*string, bool)`

GetEndDatetimeOk returns a tuple with the EndDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDatetime

`func (o *DataSourceSyncDto) SetEndDatetime(v string)`

SetEndDatetime sets EndDatetime field to given value.


### GetIsFullScan

`func (o *DataSourceSyncDto) GetIsFullScan() bool`

GetIsFullScan returns the IsFullScan field if non-nil, zero value otherwise.

### GetIsFullScanOk

`func (o *DataSourceSyncDto) GetIsFullScanOk() (*bool, bool)`

GetIsFullScanOk returns a tuple with the IsFullScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFullScan

`func (o *DataSourceSyncDto) SetIsFullScan(v bool)`

SetIsFullScan sets IsFullScan field to given value.


### GetCollectedEventIds

`func (o *DataSourceSyncDto) GetCollectedEventIds() []string`

GetCollectedEventIds returns the CollectedEventIds field if non-nil, zero value otherwise.

### GetCollectedEventIdsOk

`func (o *DataSourceSyncDto) GetCollectedEventIdsOk() (*[]string, bool)`

GetCollectedEventIdsOk returns a tuple with the CollectedEventIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectedEventIds

`func (o *DataSourceSyncDto) SetCollectedEventIds(v []string)`

SetCollectedEventIds sets CollectedEventIds field to given value.


### GetError

`func (o *DataSourceSyncDto) GetError() map[string]interface{}`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *DataSourceSyncDto) GetErrorOk() (*map[string]interface{}, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *DataSourceSyncDto) SetError(v map[string]interface{})`

SetError sets Error field to given value.

### HasError

`func (o *DataSourceSyncDto) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


