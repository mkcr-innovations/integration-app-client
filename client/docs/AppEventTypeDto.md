# AppEventTypeDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revision** | **string** |  | 
**PublishedRevision** | Pointer to **string** |  | [optional] 
**GlobalWebhookUri** | **string** |  | 
**Example** | Pointer to **map[string]interface{}** |  | [optional] 
**Schema** | Pointer to **map[string]interface{}** |  | [optional] 
**SubscribeRequest** | Pointer to **map[string]interface{}** |  | [optional] 
**UserIdFormula** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewAppEventTypeDto

`func NewAppEventTypeDto(revision string, globalWebhookUri string, ) *AppEventTypeDto`

NewAppEventTypeDto instantiates a new AppEventTypeDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppEventTypeDtoWithDefaults

`func NewAppEventTypeDtoWithDefaults() *AppEventTypeDto`

NewAppEventTypeDtoWithDefaults instantiates a new AppEventTypeDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevision

`func (o *AppEventTypeDto) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *AppEventTypeDto) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *AppEventTypeDto) SetRevision(v string)`

SetRevision sets Revision field to given value.


### GetPublishedRevision

`func (o *AppEventTypeDto) GetPublishedRevision() string`

GetPublishedRevision returns the PublishedRevision field if non-nil, zero value otherwise.

### GetPublishedRevisionOk

`func (o *AppEventTypeDto) GetPublishedRevisionOk() (*string, bool)`

GetPublishedRevisionOk returns a tuple with the PublishedRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedRevision

`func (o *AppEventTypeDto) SetPublishedRevision(v string)`

SetPublishedRevision sets PublishedRevision field to given value.

### HasPublishedRevision

`func (o *AppEventTypeDto) HasPublishedRevision() bool`

HasPublishedRevision returns a boolean if a field has been set.

### GetGlobalWebhookUri

`func (o *AppEventTypeDto) GetGlobalWebhookUri() string`

GetGlobalWebhookUri returns the GlobalWebhookUri field if non-nil, zero value otherwise.

### GetGlobalWebhookUriOk

`func (o *AppEventTypeDto) GetGlobalWebhookUriOk() (*string, bool)`

GetGlobalWebhookUriOk returns a tuple with the GlobalWebhookUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalWebhookUri

`func (o *AppEventTypeDto) SetGlobalWebhookUri(v string)`

SetGlobalWebhookUri sets GlobalWebhookUri field to given value.


### GetExample

`func (o *AppEventTypeDto) GetExample() map[string]interface{}`

GetExample returns the Example field if non-nil, zero value otherwise.

### GetExampleOk

`func (o *AppEventTypeDto) GetExampleOk() (*map[string]interface{}, bool)`

GetExampleOk returns a tuple with the Example field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExample

`func (o *AppEventTypeDto) SetExample(v map[string]interface{})`

SetExample sets Example field to given value.

### HasExample

`func (o *AppEventTypeDto) HasExample() bool`

HasExample returns a boolean if a field has been set.

### GetSchema

`func (o *AppEventTypeDto) GetSchema() map[string]interface{}`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *AppEventTypeDto) GetSchemaOk() (*map[string]interface{}, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *AppEventTypeDto) SetSchema(v map[string]interface{})`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *AppEventTypeDto) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetSubscribeRequest

`func (o *AppEventTypeDto) GetSubscribeRequest() map[string]interface{}`

GetSubscribeRequest returns the SubscribeRequest field if non-nil, zero value otherwise.

### GetSubscribeRequestOk

`func (o *AppEventTypeDto) GetSubscribeRequestOk() (*map[string]interface{}, bool)`

GetSubscribeRequestOk returns a tuple with the SubscribeRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribeRequest

`func (o *AppEventTypeDto) SetSubscribeRequest(v map[string]interface{})`

SetSubscribeRequest sets SubscribeRequest field to given value.

### HasSubscribeRequest

`func (o *AppEventTypeDto) HasSubscribeRequest() bool`

HasSubscribeRequest returns a boolean if a field has been set.

### GetUserIdFormula

`func (o *AppEventTypeDto) GetUserIdFormula() map[string]interface{}`

GetUserIdFormula returns the UserIdFormula field if non-nil, zero value otherwise.

### GetUserIdFormulaOk

`func (o *AppEventTypeDto) GetUserIdFormulaOk() (*map[string]interface{}, bool)`

GetUserIdFormulaOk returns a tuple with the UserIdFormula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdFormula

`func (o *AppEventTypeDto) SetUserIdFormula(v map[string]interface{})`

SetUserIdFormula sets UserIdFormula field to given value.

### HasUserIdFormula

`func (o *AppEventTypeDto) HasUserIdFormula() bool`

HasUserIdFormula returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


