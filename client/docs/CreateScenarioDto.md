# CreateScenarioDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Key** | **string** |  | 
**ShortDescription** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Elements** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Todos** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewCreateScenarioDto

`func NewCreateScenarioDto(name string, key string, ) *CreateScenarioDto`

NewCreateScenarioDto instantiates a new CreateScenarioDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateScenarioDtoWithDefaults

`func NewCreateScenarioDtoWithDefaults() *CreateScenarioDto`

NewCreateScenarioDtoWithDefaults instantiates a new CreateScenarioDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateScenarioDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateScenarioDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateScenarioDto) SetName(v string)`

SetName sets Name field to given value.


### GetKey

`func (o *CreateScenarioDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateScenarioDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateScenarioDto) SetKey(v string)`

SetKey sets Key field to given value.


### GetShortDescription

`func (o *CreateScenarioDto) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *CreateScenarioDto) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *CreateScenarioDto) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.

### HasShortDescription

`func (o *CreateScenarioDto) HasShortDescription() bool`

HasShortDescription returns a boolean if a field has been set.

### GetDescription

`func (o *CreateScenarioDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateScenarioDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateScenarioDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateScenarioDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetElements

`func (o *CreateScenarioDto) GetElements() []map[string]interface{}`

GetElements returns the Elements field if non-nil, zero value otherwise.

### GetElementsOk

`func (o *CreateScenarioDto) GetElementsOk() (*[]map[string]interface{}, bool)`

GetElementsOk returns a tuple with the Elements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElements

`func (o *CreateScenarioDto) SetElements(v []map[string]interface{})`

SetElements sets Elements field to given value.

### HasElements

`func (o *CreateScenarioDto) HasElements() bool`

HasElements returns a boolean if a field has been set.

### GetTodos

`func (o *CreateScenarioDto) GetTodos() []map[string]interface{}`

GetTodos returns the Todos field if non-nil, zero value otherwise.

### GetTodosOk

`func (o *CreateScenarioDto) GetTodosOk() (*[]map[string]interface{}, bool)`

GetTodosOk returns a tuple with the Todos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTodos

`func (o *CreateScenarioDto) SetTodos(v []map[string]interface{})`

SetTodos sets Todos field to given value.

### HasTodos

`func (o *CreateScenarioDto) HasTodos() bool`

HasTodos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


