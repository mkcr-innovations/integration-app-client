# AppDataSchemaInstanceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**AppId** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**User** | Pointer to **map[string]interface{}** |  | [optional] 
**AppDataSchemaId** | Pointer to **string** |  | [optional] 
**AppDataSchemaRevision** | Pointer to **string** |  | [optional] 
**AppDataSchema** | Pointer to **map[string]interface{}** |  | [optional] 
**InstanceKey** | Pointer to **string** |  | [optional] 
**Schema** | Pointer to **map[string]interface{}** |  | [optional] 
**Error** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewAppDataSchemaInstanceDto

`func NewAppDataSchemaInstanceDto(id string, ) *AppDataSchemaInstanceDto`

NewAppDataSchemaInstanceDto instantiates a new AppDataSchemaInstanceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppDataSchemaInstanceDtoWithDefaults

`func NewAppDataSchemaInstanceDtoWithDefaults() *AppDataSchemaInstanceDto`

NewAppDataSchemaInstanceDtoWithDefaults instantiates a new AppDataSchemaInstanceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppDataSchemaInstanceDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppDataSchemaInstanceDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppDataSchemaInstanceDto) SetId(v string)`

SetId sets Id field to given value.


### GetAppId

`func (o *AppDataSchemaInstanceDto) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *AppDataSchemaInstanceDto) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *AppDataSchemaInstanceDto) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *AppDataSchemaInstanceDto) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetUserId

`func (o *AppDataSchemaInstanceDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AppDataSchemaInstanceDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AppDataSchemaInstanceDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *AppDataSchemaInstanceDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUser

`func (o *AppDataSchemaInstanceDto) GetUser() map[string]interface{}`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AppDataSchemaInstanceDto) GetUserOk() (*map[string]interface{}, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AppDataSchemaInstanceDto) SetUser(v map[string]interface{})`

SetUser sets User field to given value.

### HasUser

`func (o *AppDataSchemaInstanceDto) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetAppDataSchemaId

`func (o *AppDataSchemaInstanceDto) GetAppDataSchemaId() string`

GetAppDataSchemaId returns the AppDataSchemaId field if non-nil, zero value otherwise.

### GetAppDataSchemaIdOk

`func (o *AppDataSchemaInstanceDto) GetAppDataSchemaIdOk() (*string, bool)`

GetAppDataSchemaIdOk returns a tuple with the AppDataSchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDataSchemaId

`func (o *AppDataSchemaInstanceDto) SetAppDataSchemaId(v string)`

SetAppDataSchemaId sets AppDataSchemaId field to given value.

### HasAppDataSchemaId

`func (o *AppDataSchemaInstanceDto) HasAppDataSchemaId() bool`

HasAppDataSchemaId returns a boolean if a field has been set.

### GetAppDataSchemaRevision

`func (o *AppDataSchemaInstanceDto) GetAppDataSchemaRevision() string`

GetAppDataSchemaRevision returns the AppDataSchemaRevision field if non-nil, zero value otherwise.

### GetAppDataSchemaRevisionOk

`func (o *AppDataSchemaInstanceDto) GetAppDataSchemaRevisionOk() (*string, bool)`

GetAppDataSchemaRevisionOk returns a tuple with the AppDataSchemaRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDataSchemaRevision

`func (o *AppDataSchemaInstanceDto) SetAppDataSchemaRevision(v string)`

SetAppDataSchemaRevision sets AppDataSchemaRevision field to given value.

### HasAppDataSchemaRevision

`func (o *AppDataSchemaInstanceDto) HasAppDataSchemaRevision() bool`

HasAppDataSchemaRevision returns a boolean if a field has been set.

### GetAppDataSchema

`func (o *AppDataSchemaInstanceDto) GetAppDataSchema() map[string]interface{}`

GetAppDataSchema returns the AppDataSchema field if non-nil, zero value otherwise.

### GetAppDataSchemaOk

`func (o *AppDataSchemaInstanceDto) GetAppDataSchemaOk() (*map[string]interface{}, bool)`

GetAppDataSchemaOk returns a tuple with the AppDataSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDataSchema

`func (o *AppDataSchemaInstanceDto) SetAppDataSchema(v map[string]interface{})`

SetAppDataSchema sets AppDataSchema field to given value.

### HasAppDataSchema

`func (o *AppDataSchemaInstanceDto) HasAppDataSchema() bool`

HasAppDataSchema returns a boolean if a field has been set.

### GetInstanceKey

`func (o *AppDataSchemaInstanceDto) GetInstanceKey() string`

GetInstanceKey returns the InstanceKey field if non-nil, zero value otherwise.

### GetInstanceKeyOk

`func (o *AppDataSchemaInstanceDto) GetInstanceKeyOk() (*string, bool)`

GetInstanceKeyOk returns a tuple with the InstanceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceKey

`func (o *AppDataSchemaInstanceDto) SetInstanceKey(v string)`

SetInstanceKey sets InstanceKey field to given value.

### HasInstanceKey

`func (o *AppDataSchemaInstanceDto) HasInstanceKey() bool`

HasInstanceKey returns a boolean if a field has been set.

### GetSchema

`func (o *AppDataSchemaInstanceDto) GetSchema() map[string]interface{}`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *AppDataSchemaInstanceDto) GetSchemaOk() (*map[string]interface{}, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *AppDataSchemaInstanceDto) SetSchema(v map[string]interface{})`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *AppDataSchemaInstanceDto) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetError

`func (o *AppDataSchemaInstanceDto) GetError() map[string]interface{}`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AppDataSchemaInstanceDto) GetErrorOk() (*map[string]interface{}, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AppDataSchemaInstanceDto) SetError(v map[string]interface{})`

SetError sets Error field to given value.

### HasError

`func (o *AppDataSchemaInstanceDto) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


