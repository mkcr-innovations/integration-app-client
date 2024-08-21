# ScenarioDtoElementsInnerElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revision** | Pointer to **string** |  | [optional] 
**IsCustomized** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Config** | Pointer to [**ScenarioDtoElementsInnerElementConfig**](ScenarioDtoElementsInnerElementConfig.md) |  | [optional] 
**IsDeployed** | Pointer to **bool** |  | [optional] 
**IntegrationId** | Pointer to **string** |  | [optional] 
**Integration** | Pointer to [**ScenarioDtoElementsInnerElementIntegration**](ScenarioDtoElementsInnerElementIntegration.md) |  | [optional] 
**AppliedToIntegrations** | Pointer to [**[]ScenarioDtoElementsInnerElementAppliedToIntegrationsInner**](ScenarioDtoElementsInnerElementAppliedToIntegrationsInner.md) |  | [optional] 

## Methods

### NewScenarioDtoElementsInnerElement

`func NewScenarioDtoElementsInnerElement() *ScenarioDtoElementsInnerElement`

NewScenarioDtoElementsInnerElement instantiates a new ScenarioDtoElementsInnerElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScenarioDtoElementsInnerElementWithDefaults

`func NewScenarioDtoElementsInnerElementWithDefaults() *ScenarioDtoElementsInnerElement`

NewScenarioDtoElementsInnerElementWithDefaults instantiates a new ScenarioDtoElementsInnerElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevision

`func (o *ScenarioDtoElementsInnerElement) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ScenarioDtoElementsInnerElement) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ScenarioDtoElementsInnerElement) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *ScenarioDtoElementsInnerElement) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetIsCustomized

`func (o *ScenarioDtoElementsInnerElement) GetIsCustomized() bool`

GetIsCustomized returns the IsCustomized field if non-nil, zero value otherwise.

### GetIsCustomizedOk

`func (o *ScenarioDtoElementsInnerElement) GetIsCustomizedOk() (*bool, bool)`

GetIsCustomizedOk returns a tuple with the IsCustomized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCustomized

`func (o *ScenarioDtoElementsInnerElement) SetIsCustomized(v bool)`

SetIsCustomized sets IsCustomized field to given value.

### HasIsCustomized

`func (o *ScenarioDtoElementsInnerElement) HasIsCustomized() bool`

HasIsCustomized returns a boolean if a field has been set.

### GetId

`func (o *ScenarioDtoElementsInnerElement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScenarioDtoElementsInnerElement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScenarioDtoElementsInnerElement) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ScenarioDtoElementsInnerElement) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *ScenarioDtoElementsInnerElement) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ScenarioDtoElementsInnerElement) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ScenarioDtoElementsInnerElement) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ScenarioDtoElementsInnerElement) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *ScenarioDtoElementsInnerElement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScenarioDtoElementsInnerElement) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScenarioDtoElementsInnerElement) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ScenarioDtoElementsInnerElement) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ScenarioDtoElementsInnerElement) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ScenarioDtoElementsInnerElement) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ScenarioDtoElementsInnerElement) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ScenarioDtoElementsInnerElement) HasType() bool`

HasType returns a boolean if a field has been set.

### GetConfig

`func (o *ScenarioDtoElementsInnerElement) GetConfig() ScenarioDtoElementsInnerElementConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ScenarioDtoElementsInnerElement) GetConfigOk() (*ScenarioDtoElementsInnerElementConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ScenarioDtoElementsInnerElement) SetConfig(v ScenarioDtoElementsInnerElementConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ScenarioDtoElementsInnerElement) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetIsDeployed

`func (o *ScenarioDtoElementsInnerElement) GetIsDeployed() bool`

GetIsDeployed returns the IsDeployed field if non-nil, zero value otherwise.

### GetIsDeployedOk

`func (o *ScenarioDtoElementsInnerElement) GetIsDeployedOk() (*bool, bool)`

GetIsDeployedOk returns a tuple with the IsDeployed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeployed

`func (o *ScenarioDtoElementsInnerElement) SetIsDeployed(v bool)`

SetIsDeployed sets IsDeployed field to given value.

### HasIsDeployed

`func (o *ScenarioDtoElementsInnerElement) HasIsDeployed() bool`

HasIsDeployed returns a boolean if a field has been set.

### GetIntegrationId

`func (o *ScenarioDtoElementsInnerElement) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *ScenarioDtoElementsInnerElement) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *ScenarioDtoElementsInnerElement) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.

### HasIntegrationId

`func (o *ScenarioDtoElementsInnerElement) HasIntegrationId() bool`

HasIntegrationId returns a boolean if a field has been set.

### GetIntegration

`func (o *ScenarioDtoElementsInnerElement) GetIntegration() ScenarioDtoElementsInnerElementIntegration`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *ScenarioDtoElementsInnerElement) GetIntegrationOk() (*ScenarioDtoElementsInnerElementIntegration, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *ScenarioDtoElementsInnerElement) SetIntegration(v ScenarioDtoElementsInnerElementIntegration)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *ScenarioDtoElementsInnerElement) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetAppliedToIntegrations

`func (o *ScenarioDtoElementsInnerElement) GetAppliedToIntegrations() []ScenarioDtoElementsInnerElementAppliedToIntegrationsInner`

GetAppliedToIntegrations returns the AppliedToIntegrations field if non-nil, zero value otherwise.

### GetAppliedToIntegrationsOk

`func (o *ScenarioDtoElementsInnerElement) GetAppliedToIntegrationsOk() (*[]ScenarioDtoElementsInnerElementAppliedToIntegrationsInner, bool)`

GetAppliedToIntegrationsOk returns a tuple with the AppliedToIntegrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedToIntegrations

`func (o *ScenarioDtoElementsInnerElement) SetAppliedToIntegrations(v []ScenarioDtoElementsInnerElementAppliedToIntegrationsInner)`

SetAppliedToIntegrations sets AppliedToIntegrations field to given value.

### HasAppliedToIntegrations

`func (o *ScenarioDtoElementsInnerElement) HasAppliedToIntegrations() bool`

HasAppliedToIntegrations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


