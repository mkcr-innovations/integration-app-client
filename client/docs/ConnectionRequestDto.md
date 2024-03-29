# ConnectionRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** |  | 
**Method** | **string** |  | 
**PathParameters** | Pointer to **map[string]interface{}** |  | [optional] 
**Data** | Pointer to **map[string]interface{}** |  | [optional] 
**Query** | Pointer to **map[string]interface{}** |  | [optional] 
**Headers** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewConnectionRequestDto

`func NewConnectionRequestDto(path string, method string, ) *ConnectionRequestDto`

NewConnectionRequestDto instantiates a new ConnectionRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionRequestDtoWithDefaults

`func NewConnectionRequestDtoWithDefaults() *ConnectionRequestDto`

NewConnectionRequestDtoWithDefaults instantiates a new ConnectionRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *ConnectionRequestDto) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ConnectionRequestDto) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ConnectionRequestDto) SetPath(v string)`

SetPath sets Path field to given value.


### GetMethod

`func (o *ConnectionRequestDto) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *ConnectionRequestDto) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *ConnectionRequestDto) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetPathParameters

`func (o *ConnectionRequestDto) GetPathParameters() map[string]interface{}`

GetPathParameters returns the PathParameters field if non-nil, zero value otherwise.

### GetPathParametersOk

`func (o *ConnectionRequestDto) GetPathParametersOk() (*map[string]interface{}, bool)`

GetPathParametersOk returns a tuple with the PathParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathParameters

`func (o *ConnectionRequestDto) SetPathParameters(v map[string]interface{})`

SetPathParameters sets PathParameters field to given value.

### HasPathParameters

`func (o *ConnectionRequestDto) HasPathParameters() bool`

HasPathParameters returns a boolean if a field has been set.

### GetData

`func (o *ConnectionRequestDto) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ConnectionRequestDto) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ConnectionRequestDto) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *ConnectionRequestDto) HasData() bool`

HasData returns a boolean if a field has been set.

### GetQuery

`func (o *ConnectionRequestDto) GetQuery() map[string]interface{}`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *ConnectionRequestDto) GetQueryOk() (*map[string]interface{}, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *ConnectionRequestDto) SetQuery(v map[string]interface{})`

SetQuery sets Query field to given value.

### HasQuery

`func (o *ConnectionRequestDto) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetHeaders

`func (o *ConnectionRequestDto) GetHeaders() map[string]interface{}`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *ConnectionRequestDto) GetHeadersOk() (*map[string]interface{}, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *ConnectionRequestDto) SetHeaders(v map[string]interface{})`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *ConnectionRequestDto) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


