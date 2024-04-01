# ScenarioDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**AppId** | **string** |  | 
**ScenarioTemplateId** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Key** | **string** |  | 
**ShortDescription** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Elements** | Pointer to [**[]ScenarioDtoElementsInner**](ScenarioDtoElementsInner.md) |  | [optional] 
**Todos** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ArchivedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewScenarioDto

`func NewScenarioDto(id string, appId string, name string, key string, ) *ScenarioDto`

NewScenarioDto instantiates a new ScenarioDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScenarioDtoWithDefaults

`func NewScenarioDtoWithDefaults() *ScenarioDto`

NewScenarioDtoWithDefaults instantiates a new ScenarioDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScenarioDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScenarioDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScenarioDto) SetId(v string)`

SetId sets Id field to given value.


### GetAppId

`func (o *ScenarioDto) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *ScenarioDto) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *ScenarioDto) SetAppId(v string)`

SetAppId sets AppId field to given value.


### GetScenarioTemplateId

`func (o *ScenarioDto) GetScenarioTemplateId() string`

GetScenarioTemplateId returns the ScenarioTemplateId field if non-nil, zero value otherwise.

### GetScenarioTemplateIdOk

`func (o *ScenarioDto) GetScenarioTemplateIdOk() (*string, bool)`

GetScenarioTemplateIdOk returns a tuple with the ScenarioTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenarioTemplateId

`func (o *ScenarioDto) SetScenarioTemplateId(v string)`

SetScenarioTemplateId sets ScenarioTemplateId field to given value.

### HasScenarioTemplateId

`func (o *ScenarioDto) HasScenarioTemplateId() bool`

HasScenarioTemplateId returns a boolean if a field has been set.

### GetName

`func (o *ScenarioDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScenarioDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScenarioDto) SetName(v string)`

SetName sets Name field to given value.


### GetKey

`func (o *ScenarioDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ScenarioDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ScenarioDto) SetKey(v string)`

SetKey sets Key field to given value.


### GetShortDescription

`func (o *ScenarioDto) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *ScenarioDto) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *ScenarioDto) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.

### HasShortDescription

`func (o *ScenarioDto) HasShortDescription() bool`

HasShortDescription returns a boolean if a field has been set.

### GetDescription

`func (o *ScenarioDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ScenarioDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ScenarioDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ScenarioDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetElements

`func (o *ScenarioDto) GetElements() []ScenarioDtoElementsInner`

GetElements returns the Elements field if non-nil, zero value otherwise.

### GetElementsOk

`func (o *ScenarioDto) GetElementsOk() (*[]ScenarioDtoElementsInner, bool)`

GetElementsOk returns a tuple with the Elements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElements

`func (o *ScenarioDto) SetElements(v []ScenarioDtoElementsInner)`

SetElements sets Elements field to given value.

### HasElements

`func (o *ScenarioDto) HasElements() bool`

HasElements returns a boolean if a field has been set.

### GetTodos

`func (o *ScenarioDto) GetTodos() []map[string]interface{}`

GetTodos returns the Todos field if non-nil, zero value otherwise.

### GetTodosOk

`func (o *ScenarioDto) GetTodosOk() (*[]map[string]interface{}, bool)`

GetTodosOk returns a tuple with the Todos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTodos

`func (o *ScenarioDto) SetTodos(v []map[string]interface{})`

SetTodos sets Todos field to given value.

### HasTodos

`func (o *ScenarioDto) HasTodos() bool`

HasTodos returns a boolean if a field has been set.

### GetArchivedAt

`func (o *ScenarioDto) GetArchivedAt() string`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *ScenarioDto) GetArchivedAtOk() (*string, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *ScenarioDto) SetArchivedAt(v string)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *ScenarioDto) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


