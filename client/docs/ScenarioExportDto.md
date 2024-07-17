# ScenarioExportDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Elements** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewScenarioExportDto

`func NewScenarioExportDto(name string, ) *ScenarioExportDto`

NewScenarioExportDto instantiates a new ScenarioExportDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScenarioExportDtoWithDefaults

`func NewScenarioExportDtoWithDefaults() *ScenarioExportDto`

NewScenarioExportDtoWithDefaults instantiates a new ScenarioExportDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ScenarioExportDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ScenarioExportDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ScenarioExportDto) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ScenarioExportDto) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *ScenarioExportDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScenarioExportDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScenarioExportDto) SetName(v string)`

SetName sets Name field to given value.


### GetElements

`func (o *ScenarioExportDto) GetElements() map[string]interface{}`

GetElements returns the Elements field if non-nil, zero value otherwise.

### GetElementsOk

`func (o *ScenarioExportDto) GetElementsOk() (*map[string]interface{}, bool)`

GetElementsOk returns a tuple with the Elements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElements

`func (o *ScenarioExportDto) SetElements(v map[string]interface{})`

SetElements sets Elements field to given value.

### HasElements

`func (o *ScenarioExportDto) HasElements() bool`

HasElements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


