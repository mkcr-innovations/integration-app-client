# CreateAppEventTypeDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Name** | **string** |  | 
**Schema** | Pointer to **map[string]interface{}** |  | [optional] 
**Example** | Pointer to **map[string]interface{}** |  | [optional] 
**UserIdFormula** | Pointer to **map[string]interface{}** |  | [optional] 
**SubscribeRequest** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCreateAppEventTypeDto

`func NewCreateAppEventTypeDto(key string, name string, ) *CreateAppEventTypeDto`

NewCreateAppEventTypeDto instantiates a new CreateAppEventTypeDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAppEventTypeDtoWithDefaults

`func NewCreateAppEventTypeDtoWithDefaults() *CreateAppEventTypeDto`

NewCreateAppEventTypeDtoWithDefaults instantiates a new CreateAppEventTypeDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *CreateAppEventTypeDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateAppEventTypeDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateAppEventTypeDto) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *CreateAppEventTypeDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAppEventTypeDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAppEventTypeDto) SetName(v string)`

SetName sets Name field to given value.


### GetSchema

`func (o *CreateAppEventTypeDto) GetSchema() map[string]interface{}`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *CreateAppEventTypeDto) GetSchemaOk() (*map[string]interface{}, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *CreateAppEventTypeDto) SetSchema(v map[string]interface{})`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *CreateAppEventTypeDto) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetExample

`func (o *CreateAppEventTypeDto) GetExample() map[string]interface{}`

GetExample returns the Example field if non-nil, zero value otherwise.

### GetExampleOk

`func (o *CreateAppEventTypeDto) GetExampleOk() (*map[string]interface{}, bool)`

GetExampleOk returns a tuple with the Example field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExample

`func (o *CreateAppEventTypeDto) SetExample(v map[string]interface{})`

SetExample sets Example field to given value.

### HasExample

`func (o *CreateAppEventTypeDto) HasExample() bool`

HasExample returns a boolean if a field has been set.

### GetUserIdFormula

`func (o *CreateAppEventTypeDto) GetUserIdFormula() map[string]interface{}`

GetUserIdFormula returns the UserIdFormula field if non-nil, zero value otherwise.

### GetUserIdFormulaOk

`func (o *CreateAppEventTypeDto) GetUserIdFormulaOk() (*map[string]interface{}, bool)`

GetUserIdFormulaOk returns a tuple with the UserIdFormula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdFormula

`func (o *CreateAppEventTypeDto) SetUserIdFormula(v map[string]interface{})`

SetUserIdFormula sets UserIdFormula field to given value.

### HasUserIdFormula

`func (o *CreateAppEventTypeDto) HasUserIdFormula() bool`

HasUserIdFormula returns a boolean if a field has been set.

### GetSubscribeRequest

`func (o *CreateAppEventTypeDto) GetSubscribeRequest() map[string]interface{}`

GetSubscribeRequest returns the SubscribeRequest field if non-nil, zero value otherwise.

### GetSubscribeRequestOk

`func (o *CreateAppEventTypeDto) GetSubscribeRequestOk() (*map[string]interface{}, bool)`

GetSubscribeRequestOk returns a tuple with the SubscribeRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribeRequest

`func (o *CreateAppEventTypeDto) SetSubscribeRequest(v map[string]interface{})`

SetSubscribeRequest sets SubscribeRequest field to given value.

### HasSubscribeRequest

`func (o *CreateAppEventTypeDto) HasSubscribeRequest() bool`

HasSubscribeRequest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


