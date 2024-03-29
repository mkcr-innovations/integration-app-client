# AppDataSchemaDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Key** | **string** |  | 
**Name** | **string** |  | 
**Schema** | **map[string]interface{}** |  | 
**Code** | Pointer to **string** |  | [optional] 
**ArchivedAt** | Pointer to **string** |  | [optional] 
**Revision** | Pointer to **string** |  | [optional] 

## Methods

### NewAppDataSchemaDto

`func NewAppDataSchemaDto(id string, key string, name string, schema map[string]interface{}, ) *AppDataSchemaDto`

NewAppDataSchemaDto instantiates a new AppDataSchemaDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppDataSchemaDtoWithDefaults

`func NewAppDataSchemaDtoWithDefaults() *AppDataSchemaDto`

NewAppDataSchemaDtoWithDefaults instantiates a new AppDataSchemaDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppDataSchemaDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppDataSchemaDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppDataSchemaDto) SetId(v string)`

SetId sets Id field to given value.


### GetKey

`func (o *AppDataSchemaDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AppDataSchemaDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AppDataSchemaDto) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *AppDataSchemaDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppDataSchemaDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppDataSchemaDto) SetName(v string)`

SetName sets Name field to given value.


### GetSchema

`func (o *AppDataSchemaDto) GetSchema() map[string]interface{}`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *AppDataSchemaDto) GetSchemaOk() (*map[string]interface{}, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *AppDataSchemaDto) SetSchema(v map[string]interface{})`

SetSchema sets Schema field to given value.


### GetCode

`func (o *AppDataSchemaDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AppDataSchemaDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AppDataSchemaDto) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AppDataSchemaDto) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetArchivedAt

`func (o *AppDataSchemaDto) GetArchivedAt() string`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *AppDataSchemaDto) GetArchivedAtOk() (*string, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *AppDataSchemaDto) SetArchivedAt(v string)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *AppDataSchemaDto) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### GetRevision

`func (o *AppDataSchemaDto) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *AppDataSchemaDto) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *AppDataSchemaDto) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *AppDataSchemaDto) HasRevision() bool`

HasRevision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


