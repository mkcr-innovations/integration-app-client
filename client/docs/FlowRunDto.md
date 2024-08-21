# FlowRunDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**FlowInstanceId** | Pointer to **string** |  | [optional] 
**UniversalFlowId** | Pointer to **string** |  | [optional] 
**FlowInstance** | Pointer to [**FlowInstanceDto**](FlowInstanceDto.md) |  | [optional] 
**IntegrationId** | Pointer to **string** |  | [optional] 
**Integration** | Pointer to [**IntegrationDto**](IntegrationDto.md) |  | [optional] 
**ConnectionId** | Pointer to **string** |  | [optional] 
**Connection** | Pointer to [**ConnectionDto**](ConnectionDto.md) |  | [optional] 
**StartNodeKey** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**CustomerDto**](CustomerDto.md) |  | [optional] 
**Input** | Pointer to **map[string]interface{}** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**StartTime** | Pointer to **string** |  | [optional] 
**EndTime** | Pointer to **string** |  | [optional] 
**Errors** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Nodes** | Pointer to **map[string]interface{}** |  | [optional] 
**LaunchedBy** | Pointer to [**FlorRunLaunchedByDto**](FlorRunLaunchedByDto.md) |  | [optional] 

## Methods

### NewFlowRunDto

`func NewFlowRunDto(id string, ) *FlowRunDto`

NewFlowRunDto instantiates a new FlowRunDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowRunDtoWithDefaults

`func NewFlowRunDtoWithDefaults() *FlowRunDto`

NewFlowRunDtoWithDefaults instantiates a new FlowRunDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FlowRunDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FlowRunDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FlowRunDto) SetId(v string)`

SetId sets Id field to given value.


### GetFlowInstanceId

`func (o *FlowRunDto) GetFlowInstanceId() string`

GetFlowInstanceId returns the FlowInstanceId field if non-nil, zero value otherwise.

### GetFlowInstanceIdOk

`func (o *FlowRunDto) GetFlowInstanceIdOk() (*string, bool)`

GetFlowInstanceIdOk returns a tuple with the FlowInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowInstanceId

`func (o *FlowRunDto) SetFlowInstanceId(v string)`

SetFlowInstanceId sets FlowInstanceId field to given value.

### HasFlowInstanceId

`func (o *FlowRunDto) HasFlowInstanceId() bool`

HasFlowInstanceId returns a boolean if a field has been set.

### GetUniversalFlowId

`func (o *FlowRunDto) GetUniversalFlowId() string`

GetUniversalFlowId returns the UniversalFlowId field if non-nil, zero value otherwise.

### GetUniversalFlowIdOk

`func (o *FlowRunDto) GetUniversalFlowIdOk() (*string, bool)`

GetUniversalFlowIdOk returns a tuple with the UniversalFlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniversalFlowId

`func (o *FlowRunDto) SetUniversalFlowId(v string)`

SetUniversalFlowId sets UniversalFlowId field to given value.

### HasUniversalFlowId

`func (o *FlowRunDto) HasUniversalFlowId() bool`

HasUniversalFlowId returns a boolean if a field has been set.

### GetFlowInstance

`func (o *FlowRunDto) GetFlowInstance() FlowInstanceDto`

GetFlowInstance returns the FlowInstance field if non-nil, zero value otherwise.

### GetFlowInstanceOk

`func (o *FlowRunDto) GetFlowInstanceOk() (*FlowInstanceDto, bool)`

GetFlowInstanceOk returns a tuple with the FlowInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowInstance

`func (o *FlowRunDto) SetFlowInstance(v FlowInstanceDto)`

SetFlowInstance sets FlowInstance field to given value.

### HasFlowInstance

`func (o *FlowRunDto) HasFlowInstance() bool`

HasFlowInstance returns a boolean if a field has been set.

### GetIntegrationId

`func (o *FlowRunDto) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *FlowRunDto) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *FlowRunDto) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.

### HasIntegrationId

`func (o *FlowRunDto) HasIntegrationId() bool`

HasIntegrationId returns a boolean if a field has been set.

### GetIntegration

`func (o *FlowRunDto) GetIntegration() IntegrationDto`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *FlowRunDto) GetIntegrationOk() (*IntegrationDto, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *FlowRunDto) SetIntegration(v IntegrationDto)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *FlowRunDto) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetConnectionId

`func (o *FlowRunDto) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *FlowRunDto) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *FlowRunDto) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *FlowRunDto) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetConnection

`func (o *FlowRunDto) GetConnection() ConnectionDto`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *FlowRunDto) GetConnectionOk() (*ConnectionDto, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *FlowRunDto) SetConnection(v ConnectionDto)`

SetConnection sets Connection field to given value.

### HasConnection

`func (o *FlowRunDto) HasConnection() bool`

HasConnection returns a boolean if a field has been set.

### GetStartNodeKey

`func (o *FlowRunDto) GetStartNodeKey() string`

GetStartNodeKey returns the StartNodeKey field if non-nil, zero value otherwise.

### GetStartNodeKeyOk

`func (o *FlowRunDto) GetStartNodeKeyOk() (*string, bool)`

GetStartNodeKeyOk returns a tuple with the StartNodeKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartNodeKey

`func (o *FlowRunDto) SetStartNodeKey(v string)`

SetStartNodeKey sets StartNodeKey field to given value.

### HasStartNodeKey

`func (o *FlowRunDto) HasStartNodeKey() bool`

HasStartNodeKey returns a boolean if a field has been set.

### GetUserId

`func (o *FlowRunDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *FlowRunDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *FlowRunDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *FlowRunDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUser

`func (o *FlowRunDto) GetUser() CustomerDto`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *FlowRunDto) GetUserOk() (*CustomerDto, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *FlowRunDto) SetUser(v CustomerDto)`

SetUser sets User field to given value.

### HasUser

`func (o *FlowRunDto) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetInput

`func (o *FlowRunDto) GetInput() map[string]interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *FlowRunDto) GetInputOk() (*map[string]interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *FlowRunDto) SetInput(v map[string]interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *FlowRunDto) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetState

`func (o *FlowRunDto) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FlowRunDto) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FlowRunDto) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *FlowRunDto) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStartTime

`func (o *FlowRunDto) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *FlowRunDto) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *FlowRunDto) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *FlowRunDto) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *FlowRunDto) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *FlowRunDto) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *FlowRunDto) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *FlowRunDto) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetErrors

`func (o *FlowRunDto) GetErrors() []map[string]interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *FlowRunDto) GetErrorsOk() (*[]map[string]interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *FlowRunDto) SetErrors(v []map[string]interface{})`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *FlowRunDto) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetNodes

`func (o *FlowRunDto) GetNodes() map[string]interface{}`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *FlowRunDto) GetNodesOk() (*map[string]interface{}, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *FlowRunDto) SetNodes(v map[string]interface{})`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *FlowRunDto) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetLaunchedBy

`func (o *FlowRunDto) GetLaunchedBy() FlorRunLaunchedByDto`

GetLaunchedBy returns the LaunchedBy field if non-nil, zero value otherwise.

### GetLaunchedByOk

`func (o *FlowRunDto) GetLaunchedByOk() (*FlorRunLaunchedByDto, bool)`

GetLaunchedByOk returns a tuple with the LaunchedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchedBy

`func (o *FlowRunDto) SetLaunchedBy(v FlorRunLaunchedByDto)`

SetLaunchedBy sets LaunchedBy field to given value.

### HasLaunchedBy

`func (o *FlowRunDto) HasLaunchedBy() bool`

HasLaunchedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


