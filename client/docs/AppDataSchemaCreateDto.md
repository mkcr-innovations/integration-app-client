# AppDataSchemaCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Name** | **string** |  | 
**Schema** | Pointer to **map[string]interface{}** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**ArchivedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewAppDataSchemaCreateDto

`func NewAppDataSchemaCreateDto(key string, name string, ) *AppDataSchemaCreateDto`

NewAppDataSchemaCreateDto instantiates a new AppDataSchemaCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppDataSchemaCreateDtoWithDefaults

`func NewAppDataSchemaCreateDtoWithDefaults() *AppDataSchemaCreateDto`

NewAppDataSchemaCreateDtoWithDefaults instantiates a new AppDataSchemaCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *AppDataSchemaCreateDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AppDataSchemaCreateDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AppDataSchemaCreateDto) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *AppDataSchemaCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppDataSchemaCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppDataSchemaCreateDto) SetName(v string)`

SetName sets Name field to given value.


### GetSchema

`func (o *AppDataSchemaCreateDto) GetSchema() map[string]interface{}`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *AppDataSchemaCreateDto) GetSchemaOk() (*map[string]interface{}, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *AppDataSchemaCreateDto) SetSchema(v map[string]interface{})`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *AppDataSchemaCreateDto) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetCode

`func (o *AppDataSchemaCreateDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AppDataSchemaCreateDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AppDataSchemaCreateDto) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AppDataSchemaCreateDto) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetArchivedAt

`func (o *AppDataSchemaCreateDto) GetArchivedAt() string`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *AppDataSchemaCreateDto) GetArchivedAtOk() (*string, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *AppDataSchemaCreateDto) SetArchivedAt(v string)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *AppDataSchemaCreateDto) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


