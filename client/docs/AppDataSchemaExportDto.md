# AppDataSchemaExportDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Schema** | Pointer to **map[string]interface{}** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 

## Methods

### NewAppDataSchemaExportDto

`func NewAppDataSchemaExportDto(name string, ) *AppDataSchemaExportDto`

NewAppDataSchemaExportDto instantiates a new AppDataSchemaExportDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppDataSchemaExportDtoWithDefaults

`func NewAppDataSchemaExportDtoWithDefaults() *AppDataSchemaExportDto`

NewAppDataSchemaExportDtoWithDefaults instantiates a new AppDataSchemaExportDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *AppDataSchemaExportDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AppDataSchemaExportDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AppDataSchemaExportDto) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *AppDataSchemaExportDto) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *AppDataSchemaExportDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppDataSchemaExportDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppDataSchemaExportDto) SetName(v string)`

SetName sets Name field to given value.


### GetSchema

`func (o *AppDataSchemaExportDto) GetSchema() map[string]interface{}`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *AppDataSchemaExportDto) GetSchemaOk() (*map[string]interface{}, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *AppDataSchemaExportDto) SetSchema(v map[string]interface{})`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *AppDataSchemaExportDto) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetCode

`func (o *AppDataSchemaExportDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AppDataSchemaExportDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AppDataSchemaExportDto) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AppDataSchemaExportDto) HasCode() bool`

HasCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


