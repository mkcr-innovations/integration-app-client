# \FlowsAPI

All URIs are relative to *https://api.integration.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConnectionLevelFlowControllerArchive**](FlowsAPI.md#ConnectionLevelFlowControllerArchive) | **Delete** /connections/{connectionSelector}/flows/{flowSelector} | Archive flow instance for connection
[**ConnectionLevelFlowControllerGet**](FlowsAPI.md#ConnectionLevelFlowControllerGet) | **Get** /connections/{connectionSelector}/flows/{flowSelector} | Get flow instance for connection
[**ConnectionLevelFlowControllerPatch**](FlowsAPI.md#ConnectionLevelFlowControllerPatch) | **Patch** /connections/{connectionSelector}/flows/{flowSelector} | Patch update flow instance for connection
[**ConnectionLevelFlowControllerPut**](FlowsAPI.md#ConnectionLevelFlowControllerPut) | **Put** /connections/{connectionSelector}/flows/{flowSelector} | Update flow instance for connection
[**ConnectionLevelFlowControllerReset**](FlowsAPI.md#ConnectionLevelFlowControllerReset) | **Post** /connections/{connectionSelector}/flows/{flowSelector}/reset | Reset flow instance for connection
[**ConnectionLevelFlowControllerSetup**](FlowsAPI.md#ConnectionLevelFlowControllerSetup) | **Post** /connections/{connectionSelector}/flows/{flowSelector}/setup | Setup flow instance for connection
[**ConnectionLevelFlowsControllerList**](FlowsAPI.md#ConnectionLevelFlowsControllerList) | **Get** /connections/{connectionSelector}/flows | List flow instances for connection
[**FlowByIdControllerApply**](FlowsAPI.md#FlowByIdControllerApply) | **Post** /flows/{id}/apply | Apply flow to integrations
[**FlowByIdControllerArchive**](FlowsAPI.md#FlowByIdControllerArchive) | **Delete** /flows/{id} | Archive flow by id
[**FlowByIdControllerClone**](FlowsAPI.md#FlowByIdControllerClone) | **Post** /flows/{id}/clone | 
[**FlowByIdControllerExport**](FlowsAPI.md#FlowByIdControllerExport) | **Get** /flows/{id}/export | 
[**FlowByIdControllerGet**](FlowsAPI.md#FlowByIdControllerGet) | **Get** /flows/{id} | Get flow by id
[**FlowByIdControllerPatch**](FlowsAPI.md#FlowByIdControllerPatch) | **Patch** /flows/{id} | Patch flow by id
[**FlowByIdControllerPut**](FlowsAPI.md#FlowByIdControllerPut) | **Put** /flows/{id} | Update flow by id
[**FlowByIdControllerReset**](FlowsAPI.md#FlowByIdControllerReset) | **Post** /flows/{id}/reset | Reset flow by id
[**FlowInstanceByIdControllerArchiveFlowInstance**](FlowsAPI.md#FlowInstanceByIdControllerArchiveFlowInstance) | **Delete** /flow-instances/{id} | Archive flow instance by id
[**FlowInstanceByIdControllerCreateFlowInstance**](FlowsAPI.md#FlowInstanceByIdControllerCreateFlowInstance) | **Post** /flow-instances/{id} | Create flow instance by id
[**FlowInstanceByIdControllerExport**](FlowsAPI.md#FlowInstanceByIdControllerExport) | **Get** /flow-instances/{id}/export | 
[**FlowInstanceByIdControllerGetFlowInstance**](FlowsAPI.md#FlowInstanceByIdControllerGetFlowInstance) | **Get** /flow-instances/{id} | Get flow instance by id
[**FlowInstanceByIdControllerPatchFlowInstance**](FlowsAPI.md#FlowInstanceByIdControllerPatchFlowInstance) | **Patch** /flow-instances/{id} | Patch flow instance by id
[**FlowInstanceByIdControllerResetFlowInstance**](FlowsAPI.md#FlowInstanceByIdControllerResetFlowInstance) | **Post** /flow-instances/{id}/reset | Reset flow instance by id
[**FlowInstanceByIdControllerSetupFlowInstance**](FlowsAPI.md#FlowInstanceByIdControllerSetupFlowInstance) | **Post** /flow-instances/{id}/setup | Setup flow instance by id
[**FlowInstanceByIdControllerUpdateFlowInstance**](FlowsAPI.md#FlowInstanceByIdControllerUpdateFlowInstance) | **Put** /flow-instances/{id} | Update flow instance by id
[**FlowInstancesControllerListFlowInstances**](FlowsAPI.md#FlowInstancesControllerListFlowInstances) | **Get** /flow-instances | List flow instances
[**FlowRunsControllerGetFlowNodeRun**](FlowsAPI.md#FlowRunsControllerGetFlowNodeRun) | **Get** /flow-runs/{id}/nodes/{nodeKey}/runs/{nodeRunId} | Get node run
[**FlowRunsControllerGetFlowRun**](FlowsAPI.md#FlowRunsControllerGetFlowRun) | **Get** /flow-runs/{id} | Get flow run
[**FlowRunsControllerGetFlowRunOutput**](FlowsAPI.md#FlowRunsControllerGetFlowRunOutput) | **Get** /flow-runs/{id}/output | Get flow run output
[**FlowRunsControllerGetFlowRunOutputForNode**](FlowsAPI.md#FlowRunsControllerGetFlowRunOutputForNode) | **Get** /flow-runs/{id}/output/{nodeKey} | Get flow run output for specific node
[**FlowRunsControllerGetNodeRunOutput**](FlowsAPI.md#FlowRunsControllerGetNodeRunOutput) | **Get** /flow-runs/{id}/nodes/{nodeKey}/outputs/{outputId} | Get node run output
[**FlowRunsControllerListFlowNodeRuns**](FlowsAPI.md#FlowRunsControllerListFlowNodeRuns) | **Get** /flow-runs/{id}/nodes/{nodeKey}/runs | Get node runs
[**FlowRunsControllerListFlowRuns**](FlowsAPI.md#FlowRunsControllerListFlowRuns) | **Get** /flow-runs | List flow runs
[**FlowRunsControllerListNodeRunOutputs**](FlowsAPI.md#FlowRunsControllerListNodeRunOutputs) | **Get** /flow-runs/{id}/nodes/{nodeKey}/outputs | Get node run outputs
[**FlowRunsControllerStopFlowRun**](FlowsAPI.md#FlowRunsControllerStopFlowRun) | **Post** /flow-runs/{id}/stop | Stop flow run
[**FlowsControllerCreateFlow**](FlowsAPI.md#FlowsControllerCreateFlow) | **Post** /flows | Create flow
[**FlowsControllerListFlows**](FlowsAPI.md#FlowsControllerListFlows) | **Get** /flows | List flows
[**IntegrationLevelFlowControllerArchive**](FlowsAPI.md#IntegrationLevelFlowControllerArchive) | **Delete** /integrations/{integrationSelector}/flows/{flowSelector} | Archive flow for integration
[**IntegrationLevelFlowControllerGet**](FlowsAPI.md#IntegrationLevelFlowControllerGet) | **Get** /integrations/{integrationSelector}/flows/{flowSelector} | Get flow for integration
[**IntegrationLevelFlowControllerPatch**](FlowsAPI.md#IntegrationLevelFlowControllerPatch) | **Patch** /integrations/{integrationSelector}/flows/{flowSelector} | Patch update flow for integration
[**IntegrationLevelFlowControllerPut**](FlowsAPI.md#IntegrationLevelFlowControllerPut) | **Put** /integrations/{integrationSelector}/flows/{flowSelector} | Update flow for integration
[**IntegrationLevelFlowControllerReset**](FlowsAPI.md#IntegrationLevelFlowControllerReset) | **Post** /integrations/{integrationSelector}/flows/{flowSelector}/reset | Reset flow for integration
[**IntegrationLevelFlowsControllerCreate**](FlowsAPI.md#IntegrationLevelFlowsControllerCreate) | **Post** /integrations/{integrationSelector}/flows | Create flow for integration
[**IntegrationLevelFlowsControllerList**](FlowsAPI.md#IntegrationLevelFlowsControllerList) | **Get** /integrations/{integrationSelector}/flows | List flows for integration
[**RunFlowControllerRunFlow**](FlowsAPI.md#RunFlowControllerRunFlow) | **Post** /connections/{connectionSelector}/flows/{flowSelector}/run | Run flow instance for connection



## ConnectionLevelFlowControllerArchive

> ConnectionLevelFlowControllerArchive(ctx, flowSelector, connectionSelector).InstanceKey(instanceKey).Execute()

Archive flow instance for connection

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/integration-app-client/client"
)

func main() {
	flowSelector := "flowSelector_example" // string | Flow ID or Key
	connectionSelector := "connectionSelector_example" // string | Integration Key, Connection ID, or Integration ID
	instanceKey := "instanceKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FlowsAPI.ConnectionLevelFlowControllerArchive(context.Background(), flowSelector, connectionSelector).InstanceKey(instanceKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.ConnectionLevelFlowControllerArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowSelector** | **string** | Flow ID or Key | 
**connectionSelector** | **string** | Integration Key, Connection ID, or Integration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionLevelFlowControllerArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **instanceKey** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionLevelFlowControllerGet

> FlowInstanceDto ConnectionLevelFlowControllerGet(ctx, flowSelector, connectionSelector).InstanceKey(instanceKey).AutoCreate(autoCreate).Execute()

Get flow instance for connection

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/integration-app-client/client"
)

func main() {
	flowSelector := "flowSelector_example" // string | Flow ID or Key
	connectionSelector := "connectionSelector_example" // string | Integration Key, Connection ID, or Integration ID
	instanceKey := "instanceKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.ConnectionLevelFlowControllerGet(context.Background(), flowSelector, connectionSelector).InstanceKey(instanceKey).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.ConnectionLevelFlowControllerGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectionLevelFlowControllerGet`: FlowInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.ConnectionLevelFlowControllerGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowSelector** | **string** | Flow ID or Key | 
**connectionSelector** | **string** | Integration Key, Connection ID, or Integration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionLevelFlowControllerGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **instanceKey** | **string** |  | 
 **autoCreate** | **bool** |  | 

### Return type

[**FlowInstanceDto**](FlowInstanceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionLevelFlowControllerPatch

> FlowInstanceDto ConnectionLevelFlowControllerPatch(ctx, flowSelector, connectionSelector).UpdateFlowInstanceDto(updateFlowInstanceDto).InstanceKey(instanceKey).Execute()

Patch update flow instance for connection

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/integration-app-client/client"
)

func main() {
	flowSelector := "flowSelector_example" // string | Flow ID or Key
	connectionSelector := "connectionSelector_example" // string | Integration Key, Connection ID, or Integration ID
	updateFlowInstanceDto := *openapiclient.NewUpdateFlowInstanceDto() // UpdateFlowInstanceDto | 
	instanceKey := "instanceKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.ConnectionLevelFlowControllerPatch(context.Background(), flowSelector, connectionSelector).UpdateFlowInstanceDto(updateFlowInstanceDto).InstanceKey(instanceKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.ConnectionLevelFlowControllerPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectionLevelFlowControllerPatch`: FlowInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.ConnectionLevelFlowControllerPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowSelector** | **string** | Flow ID or Key | 
**connectionSelector** | **string** | Integration Key, Connection ID, or Integration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionLevelFlowControllerPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateFlowInstanceDto** | [**UpdateFlowInstanceDto**](UpdateFlowInstanceDto.md) |  | 
 **instanceKey** | **string** |  | 

### Return type

[**FlowInstanceDto**](FlowInstanceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionLevelFlowControllerPut

> FlowInstanceDto ConnectionLevelFlowControllerPut(ctx, flowSelector, connectionSelector).UpdateFlowInstanceDto(updateFlowInstanceDto).InstanceKey(instanceKey).Execute()

Update flow instance for connection

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/integration-app-client/client"
)

func main() {
	flowSelector := "flowSelector_example" // string | Flow ID or Key
	connectionSelector := "connectionSelector_example" // string | Integration Key, Connection ID, or Integration ID
	updateFlowInstanceDto := *openapiclient.NewUpdateFlowInstanceDto() // UpdateFlowInstanceDto | 
	instanceKey := "instanceKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.ConnectionLevelFlowControllerPut(context.Background(), flowSelector, connectionSelector).UpdateFlowInstanceDto(updateFlowInstanceDto).InstanceKey(instanceKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.ConnectionLevelFlowControllerPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectionLevelFlowControllerPut`: FlowInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.ConnectionLevelFlowControllerPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowSelector** | **string** | Flow ID or Key | 
**connectionSelector** | **string** | Integration Key, Connection ID, or Integration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionLevelFlowControllerPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateFlowInstanceDto** | [**UpdateFlowInstanceDto**](UpdateFlowInstanceDto.md) |  | 
 **instanceKey** | **string** |  | 

### Return type

[**FlowInstanceDto**](FlowInstanceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionLevelFlowControllerReset

> FlowInstanceDto ConnectionLevelFlowControllerReset(ctx, flowSelector, connectionSelector).InstanceKey(instanceKey).Execute()

Reset flow instance for connection

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/integration-app-client/client"
)

func main() {
	flowSelector := "flowSelector_example" // string | Flow ID or Key
	connectionSelector := "connectionSelector_example" // string | Integration Key, Connection ID, or Integration ID
	instanceKey := "instanceKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.ConnectionLevelFlowControllerReset(context.Background(), flowSelector, connectionSelector).InstanceKey(instanceKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.ConnectionLevelFlowControllerReset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectionLevelFlowControllerReset`: FlowInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.ConnectionLevelFlowControllerReset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowSelector** | **string** | Flow ID or Key | 
**connectionSelector** | **string** | Integration Key, Connection ID, or Integration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionLevelFlowControllerResetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **instanceKey** | **string** |  | 

### Return type

[**FlowInstanceDto**](FlowInstanceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionLevelFlowControllerSetup

> FlowInstanceDto ConnectionLevelFlowControllerSetup(ctx, flowSelector, connectionSelector).InstanceKey(instanceKey).Execute()

Setup flow instance for connection

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/integration-app-client/client"
)

func main() {
	flowSelector := "flowSelector_example" // string | Flow ID or Key
	connectionSelector := "connectionSelector_example" // string | Integration Key, Connection ID, or Integration ID
	instanceKey := "instanceKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.ConnectionLevelFlowControllerSetup(context.Background(), flowSelector, connectionSelector).InstanceKey(instanceKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.ConnectionLevelFlowControllerSetup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectionLevelFlowControllerSetup`: FlowInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.ConnectionLevelFlowControllerSetup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowSelector** | **string** | Flow ID or Key | 
**connectionSelector** | **string** | Integration Key, Connection ID, or Integration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionLevelFlowControllerSetupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **instanceKey** | **string** |  | 

### Return type

[**FlowInstanceDto**](FlowInstanceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionLevelFlowsControllerList

> ConnectionLevelFlowsControllerList(ctx, connectionSelector).Limit(limit).Cursor(cursor).Search(search).ConnectorId(connectorId).IncludeArchived(includeArchived).Id(id).InstanceKey(instanceKey).UserId(userId).FlowKey(flowKey).FlowId(flowId).UniversalFlowId(universalFlowId).Enabled(enabled).Execute()

List flow instances for connection

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/integration-app-client/client"
)

func main() {
	connectionSelector := "connectionSelector_example" // string | Integration Key, Connection ID, or Integration ID
	limit := float32(8.14) // float32 |  (optional)
	cursor := "cursor_example" // string |  (optional)
	search := "search_example" // string |  (optional)
	connectorId := "connectorId_example" // string |  (optional)
	includeArchived := true // bool |  (optional)
	id := "id_example" // string |  (optional)
	instanceKey := "instanceKey_example" // string |  (optional)
	userId := "userId_example" // string |  (optional)
	flowKey := "flowKey_example" // string |  (optional)
	flowId := "flowId_example" // string |  (optional)
	universalFlowId := "universalFlowId_example" // string |  (optional)
	enabled := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FlowsAPI.ConnectionLevelFlowsControllerList(context.Background(), connectionSelector).Limit(limit).Cursor(cursor).Search(search).ConnectorId(connectorId).IncludeArchived(includeArchived).Id(id).InstanceKey(instanceKey).UserId(userId).FlowKey(flowKey).FlowId(flowId).UniversalFlowId(universalFlowId).Enabled(enabled).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.ConnectionLevelFlowsControllerList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionSelector** | **string** | Integration Key, Connection ID, or Integration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionLevelFlowsControllerListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **float32** |  | 
 **cursor** | **string** |  | 
 **search** | **string** |  | 
 **connectorId** | **string** |  | 
 **includeArchived** | **bool** |  | 
 **id** | **string** |  | 
 **instanceKey** | **string** |  | 
 **userId** | **string** |  | 
 **flowKey** | **string** |  | 
 **flowId** | **string** |  | 
 **universalFlowId** | **string** |  | 
 **enabled** | **bool** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowByIdControllerApply

> []FlowDto FlowByIdControllerApply(ctx, id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).AutoCreate(autoCreate).Execute()

Apply flow to integrations

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/integration-app-client/client"
)

func main() {
	id := "id_example" // string | ID
	key := "key_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.FlowByIdControllerApply(context.Background(), id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.FlowByIdControllerApply``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowByIdControllerApply`: []FlowDto
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.FlowByIdControllerApply`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowByIdControllerApplyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **key** | **string** |  | 
 **integrationId** | **string** |  | 
 **integrationKey** | **string** |  | 
 **autoCreate** | **bool** |  | 

### Return type

[**[]FlowDto**](FlowDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowByIdControllerArchive

> FlowByIdControllerArchive(ctx, id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).AutoCreate(autoCreate).Execute()

Archive flow by id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/integration-app-client/client"
)

func main() {
	id := "id_example" // string | ID
	key := "key_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FlowsAPI.FlowByIdControllerArchive(context.Background(), id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.FlowByIdControllerArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowByIdControllerArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **key** | **string** |  | 
 **integrationId** | **string** |  | 
 **integrationKey** | **string** |  | 
 **autoCreate** | **bool** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowByIdControllerClone

> FlowDto FlowByIdControllerClone(ctx, id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).AutoCreate(autoCreate).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/integration-app-client/client"
)

func main() {
	id := "id_example" // string | ID
	key := "key_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.FlowByIdControllerClone(context.Background(), id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.FlowByIdControllerClone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowByIdControllerClone`: FlowDto
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.FlowByIdControllerClone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowByIdControllerCloneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **key** | **string** |  | 
 **integrationId** | **string** |  | 
 **integrationKey** | **string** |  | 
 **autoCreate** | **bool** |  | 

### Return type

[**FlowDto**](FlowDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowByIdControllerExport

> FlowExportDto FlowByIdControllerExport(ctx, id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).AutoCreate(autoCreate).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/integration-app-client/client"
)

func main() {
	id := "id_example" // string | ID
	key := "key_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.FlowByIdControllerExport(context.Background(), id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.FlowByIdControllerExport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowByIdControllerExport`: FlowExportDto
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.FlowByIdControllerExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowByIdControllerExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **key** | **string** |  | 
 **integrationId** | **string** |  | 
 **integrationKey** | **string** |  | 
 **autoCreate** | **bool** |  | 

### Return type

[**FlowExportDto**](FlowExportDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowByIdControllerGet

> FlowDto FlowByIdControllerGet(ctx, id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).AutoCreate(autoCreate).Execute()

Get flow by id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/integration-app-client/client"
)

func main() {
	id := "id_example" // string | ID
	key := "key_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.FlowByIdControllerGet(context.Background(), id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.FlowByIdControllerGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowByIdControllerGet`: FlowDto
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.FlowByIdControllerGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowByIdControllerGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **key** | **string** |  | 
 **integrationId** | **string** |  | 
 **integrationKey** | **string** |  | 
 **autoCreate** | **bool** |  | 

### Return type

[**FlowDto**](FlowDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowByIdControllerPatch

> FlowDto FlowByIdControllerPatch(ctx, id).UpdateFlowDto(updateFlowDto).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).AutoCreate(autoCreate).Execute()

Patch flow by id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/integration-app-client/client"
)

func main() {
	id := "id_example" // string | ID
	updateFlowDto := *openapiclient.NewUpdateFlowDto() // UpdateFlowDto | 
	key := "key_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.FlowByIdControllerPatch(context.Background(), id).UpdateFlowDto(updateFlowDto).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.FlowByIdControllerPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowByIdControllerPatch`: FlowDto
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.FlowByIdControllerPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowByIdControllerPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateFlowDto** | [**UpdateFlowDto**](UpdateFlowDto.md) |  | 
 **key** | **string** |  | 
 **integrationId** | **string** |  | 
 **integrationKey** | **string** |  | 
 **autoCreate** | **bool** |  | 

### Return type

[**FlowDto**](FlowDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowByIdControllerPut

> FlowDto FlowByIdControllerPut(ctx, id).UpdateFlowDto(updateFlowDto).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).AutoCreate(autoCreate).Execute()

Update flow by id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/integration-app-client/client"
)

func main() {
	id := "id_example" // string | ID
	updateFlowDto := *openapiclient.NewUpdateFlowDto() // UpdateFlowDto | 
	key := "key_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.FlowByIdControllerPut(context.Background(), id).UpdateFlowDto(updateFlowDto).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.FlowByIdControllerPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowByIdControllerPut`: FlowDto
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.FlowByIdControllerPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowByIdControllerPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateFlowDto** | [**UpdateFlowDto**](UpdateFlowDto.md) |  | 
 **key** | **string** |  | 
 **integrationId** | **string** |  | 
 **integrationKey** | **string** |  | 
 **autoCreate** | **bool** |  | 

### Return type

[**FlowDto**](FlowDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowByIdControllerReset

> FlowDto FlowByIdControllerReset(ctx, id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).AutoCreate(autoCreate).Execute()

Reset flow by id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/integration-app-client/client"
)

func main() {
	id := "id_example" // string | ID
	key := "key_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.FlowByIdControllerReset(context.Background(), id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.FlowByIdControllerReset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowByIdControllerReset`: FlowDto
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.FlowByIdControllerReset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowByIdControllerResetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **key** | **string** |  | 
 **integrationId** | **string** |  | 
 **integrationKey** | **string** |  | 
 **autoCreate** | **bool** |  | 

### Return type

[**FlowDto**](FlowDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowInstanceByIdControllerArchiveFlowInstance

> FlowInstanceByIdControllerArchiveFlowInstance(ctx, id).FlowId(flowId).FlowKey(flowKey).InstanceKey(instanceKey).AutoCreate(autoCreate).AutoUpdate(autoUpdate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()

Archive flow instance by id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/integration-app-client/client"
)

func main() {
	id := "id_example" // string | ID
	flowId := "flowId_example" // string |  (optional)
	flowKey := "flowKey_example" // string |  (optional)
	instanceKey := "instanceKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)
	autoUpdate := true // bool |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	connectionId := "connectionId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FlowsAPI.FlowInstanceByIdControllerArchiveFlowInstance(context.Background(), id).FlowId(flowId).FlowKey(flowKey).InstanceKey(instanceKey).AutoCreate(autoCreate).AutoUpdate(autoUpdate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.FlowInstanceByIdControllerArchiveFlowInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowInstanceByIdControllerArchiveFlowInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **flowId** | **string** |  | 
 **flowKey** | **string** |  | 
 **instanceKey** | **string** |  | 
 **autoCreate** | **bool** |  | 
 **autoUpdate** | **bool** |  | 
 **integrationKey** | **string** |  | 
 **integrationId** | **string** |  | 
 **connectionId** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowInstanceByIdControllerCreateFlowInstance

> FlowInstanceDto FlowInstanceByIdControllerCreateFlowInstance(ctx, id).UpdateFlowInstanceDto(updateFlowInstanceDto).FlowId(flowId).FlowKey(flowKey).InstanceKey(instanceKey).AutoCreate(autoCreate).AutoUpdate(autoUpdate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()

Create flow instance by id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/integration-app-client/client"
)

func main() {
	id := "id_example" // string | ID
	updateFlowInstanceDto := *openapiclient.NewUpdateFlowInstanceDto() // UpdateFlowInstanceDto | 
	flowId := "flowId_example" // string |  (optional)
	flowKey := "flowKey_example" // string |  (optional)
	instanceKey := "instanceKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)
	autoUpdate := true // bool |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	connectionId := "connectionId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.FlowInstanceByIdControllerCreateFlowInstance(context.Background(), id).UpdateFlowInstanceDto(updateFlowInstanceDto).FlowId(flowId).FlowKey(flowKey).InstanceKey(instanceKey).AutoCreate(autoCreate).AutoUpdate(autoUpdate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.FlowInstanceByIdControllerCreateFlowInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowInstanceByIdControllerCreateFlowInstance`: FlowInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.FlowInstanceByIdControllerCreateFlowInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowInstanceByIdControllerCreateFlowInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateFlowInstanceDto** | [**UpdateFlowInstanceDto**](UpdateFlowInstanceDto.md) |  | 
 **flowId** | **string** |  | 
 **flowKey** | **string** |  | 
 **instanceKey** | **string** |  | 
 **autoCreate** | **bool** |  | 
 **autoUpdate** | **bool** |  | 
 **integrationKey** | **string** |  | 
 **integrationId** | **string** |  | 
 **connectionId** | **string** |  | 

### Return type

[**FlowInstanceDto**](FlowInstanceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowInstanceByIdControllerExport

> map[string]interface{} FlowInstanceByIdControllerExport(ctx, id).FlowId(flowId).FlowKey(flowKey).InstanceKey(instanceKey).AutoCreate(autoCreate).AutoUpdate(autoUpdate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/integration-app-client/client"
)

func main() {
	id := "id_example" // string | ID
	flowId := "flowId_example" // string |  (optional)
	flowKey := "flowKey_example" // string |  (optional)
	instanceKey := "instanceKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)
	autoUpdate := true // bool |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	connectionId := "connectionId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.FlowInstanceByIdControllerExport(context.Background(), id).FlowId(flowId).FlowKey(flowKey).InstanceKey(instanceKey).AutoCreate(autoCreate).AutoUpdate(autoUpdate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.FlowInstanceByIdControllerExport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowInstanceByIdControllerExport`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.FlowInstanceByIdControllerExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowInstanceByIdControllerExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **flowId** | **string** |  | 
 **flowKey** | **string** |  | 
 **instanceKey** | **string** |  | 
 **autoCreate** | **bool** |  | 
 **autoUpdate** | **bool** |  | 
 **integrationKey** | **string** |  | 
 **integrationId** | **string** |  | 
 **connectionId** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowInstanceByIdControllerGetFlowInstance

> map[string]interface{} FlowInstanceByIdControllerGetFlowInstance(ctx, id).FlowId(flowId).FlowKey(flowKey).InstanceKey(instanceKey).AutoCreate(autoCreate).AutoUpdate(autoUpdate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()

Get flow instance by id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/integration-app-client/client"
)

func main() {
	id := "id_example" // string | ID
	flowId := "flowId_example" // string |  (optional)
	flowKey := "flowKey_example" // string |  (optional)
	instanceKey := "instanceKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)
	autoUpdate := true // bool |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	connectionId := "connectionId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.FlowInstanceByIdControllerGetFlowInstance(context.Background(), id).FlowId(flowId).FlowKey(flowKey).InstanceKey(instanceKey).AutoCreate(autoCreate).AutoUpdate(autoUpdate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.FlowInstanceByIdControllerGetFlowInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowInstanceByIdControllerGetFlowInstance`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.FlowInstanceByIdControllerGetFlowInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowInstanceByIdControllerGetFlowInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **flowId** | **string** |  | 
 **flowKey** | **string** |  | 
 **instanceKey** | **string** |  | 
 **autoCreate** | **bool** |  | 
 **autoUpdate** | **bool** |  | 
 **integrationKey** | **string** |  | 
 **integrationId** | **string** |  | 
 **connectionId** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowInstanceByIdControllerPatchFlowInstance

> FlowInstanceDto FlowInstanceByIdControllerPatchFlowInstance(ctx, id).UpdateFlowInstanceDto(updateFlowInstanceDto).FlowId(flowId).FlowKey(flowKey).InstanceKey(instanceKey).AutoCreate(autoCreate).AutoUpdate(autoUpdate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()

Patch flow instance by id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/integration-app-client/client"
)

func main() {
	id := "id_example" // string | ID
	updateFlowInstanceDto := *openapiclient.NewUpdateFlowInstanceDto() // UpdateFlowInstanceDto | 
	flowId := "flowId_example" // string |  (optional)
	flowKey := "flowKey_example" // string |  (optional)
	instanceKey := "instanceKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)
	autoUpdate := true // bool |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	connectionId := "connectionId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.FlowInstanceByIdControllerPatchFlowInstance(context.Background(), id).UpdateFlowInstanceDto(updateFlowInstanceDto).FlowId(flowId).FlowKey(flowKey).InstanceKey(instanceKey).AutoCreate(autoCreate).AutoUpdate(autoUpdate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.FlowInstanceByIdControllerPatchFlowInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowInstanceByIdControllerPatchFlowInstance`: FlowInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.FlowInstanceByIdControllerPatchFlowInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowInstanceByIdControllerPatchFlowInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateFlowInstanceDto** | [**UpdateFlowInstanceDto**](UpdateFlowInstanceDto.md) |  | 
 **flowId** | **string** |  | 
 **flowKey** | **string** |  | 
 **instanceKey** | **string** |  | 
 **autoCreate** | **bool** |  | 
 **autoUpdate** | **bool** |  | 
 **integrationKey** | **string** |  | 
 **integrationId** | **string** |  | 
 **connectionId** | **string** |  | 

### Return type

[**FlowInstanceDto**](FlowInstanceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowInstanceByIdControllerResetFlowInstance

> FlowInstanceDto FlowInstanceByIdControllerResetFlowInstance(ctx, id).FlowId(flowId).FlowKey(flowKey).InstanceKey(instanceKey).AutoCreate(autoCreate).AutoUpdate(autoUpdate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()

Reset flow instance by id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/integration-app-client/client"
)

func main() {
	id := "id_example" // string | ID
	flowId := "flowId_example" // string |  (optional)
	flowKey := "flowKey_example" // string |  (optional)
	instanceKey := "instanceKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)
	autoUpdate := true // bool |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	connectionId := "connectionId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.FlowInstanceByIdControllerResetFlowInstance(context.Background(), id).FlowId(flowId).FlowKey(flowKey).InstanceKey(instanceKey).AutoCreate(autoCreate).AutoUpdate(autoUpdate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.FlowInstanceByIdControllerResetFlowInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowInstanceByIdControllerResetFlowInstance`: FlowInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.FlowInstanceByIdControllerResetFlowInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowInstanceByIdControllerResetFlowInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **flowId** | **string** |  | 
 **flowKey** | **string** |  | 
 **instanceKey** | **string** |  | 
 **autoCreate** | **bool** |  | 
 **autoUpdate** | **bool** |  | 
 **integrationKey** | **string** |  | 
 **integrationId** | **string** |  | 
 **connectionId** | **string** |  | 

### Return type

[**FlowInstanceDto**](FlowInstanceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowInstanceByIdControllerSetupFlowInstance

> FlowInstanceDto FlowInstanceByIdControllerSetupFlowInstance(ctx, id).FlowId(flowId).FlowKey(flowKey).InstanceKey(instanceKey).AutoCreate(autoCreate).AutoUpdate(autoUpdate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()

Setup flow instance by id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/integration-app-client/client"
)

func main() {
	id := "id_example" // string | ID
	flowId := "flowId_example" // string |  (optional)
	flowKey := "flowKey_example" // string |  (optional)
	instanceKey := "instanceKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)
	autoUpdate := true // bool |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	connectionId := "connectionId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.FlowInstanceByIdControllerSetupFlowInstance(context.Background(), id).FlowId(flowId).FlowKey(flowKey).InstanceKey(instanceKey).AutoCreate(autoCreate).AutoUpdate(autoUpdate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.FlowInstanceByIdControllerSetupFlowInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowInstanceByIdControllerSetupFlowInstance`: FlowInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.FlowInstanceByIdControllerSetupFlowInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowInstanceByIdControllerSetupFlowInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **flowId** | **string** |  | 
 **flowKey** | **string** |  | 
 **instanceKey** | **string** |  | 
 **autoCreate** | **bool** |  | 
 **autoUpdate** | **bool** |  | 
 **integrationKey** | **string** |  | 
 **integrationId** | **string** |  | 
 **connectionId** | **string** |  | 

### Return type

[**FlowInstanceDto**](FlowInstanceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowInstanceByIdControllerUpdateFlowInstance

> FlowInstanceDto FlowInstanceByIdControllerUpdateFlowInstance(ctx, id).UpdateFlowInstanceDto(updateFlowInstanceDto).FlowId(flowId).FlowKey(flowKey).InstanceKey(instanceKey).AutoCreate(autoCreate).AutoUpdate(autoUpdate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()

Update flow instance by id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/integration-app-client/client"
)

func main() {
	id := "id_example" // string | ID
	updateFlowInstanceDto := *openapiclient.NewUpdateFlowInstanceDto() // UpdateFlowInstanceDto | 
	flowId := "flowId_example" // string |  (optional)
	flowKey := "flowKey_example" // string |  (optional)
	instanceKey := "instanceKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)
	autoUpdate := true // bool |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	connectionId := "connectionId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.FlowInstanceByIdControllerUpdateFlowInstance(context.Background(), id).UpdateFlowInstanceDto(updateFlowInstanceDto).FlowId(flowId).FlowKey(flowKey).InstanceKey(instanceKey).AutoCreate(autoCreate).AutoUpdate(autoUpdate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.FlowInstanceByIdControllerUpdateFlowInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowInstanceByIdControllerUpdateFlowInstance`: FlowInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.FlowInstanceByIdControllerUpdateFlowInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowInstanceByIdControllerUpdateFlowInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateFlowInstanceDto** | [**UpdateFlowInstanceDto**](UpdateFlowInstanceDto.md) |  | 
 **flowId** | **string** |  | 
 **flowKey** | **string** |  | 
 **instanceKey** | **string** |  | 
 **autoCreate** | **bool** |  | 
 **autoUpdate** | **bool** |  | 
 **integrationKey** | **string** |  | 
 **integrationId** | **string** |  | 
 **connectionId** | **string** |  | 

### Return type

[**FlowInstanceDto**](FlowInstanceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowInstancesControllerListFlowInstances

> FlowInstancesControllerListFlowInstances200Response FlowInstancesControllerListFlowInstances(ctx).Limit(limit).Cursor(cursor).Search(search).ConnectorId(connectorId).IncludeArchived(includeArchived).Id(id).InstanceKey(instanceKey).UserId(userId).FlowKey(flowKey).FlowId(flowId).UniversalFlowId(universalFlowId).Enabled(enabled).ConnectionId(connectionId).IntegrationId(integrationId).IntegrationKey(integrationKey).DependencyInstanceId(dependencyInstanceId).Execute()

List flow instances

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/integration-app-client/client"
)

func main() {
	limit := float32(8.14) // float32 |  (optional)
	cursor := "cursor_example" // string |  (optional)
	search := "search_example" // string |  (optional)
	connectorId := "connectorId_example" // string |  (optional)
	includeArchived := true // bool |  (optional)
	id := "id_example" // string |  (optional)
	instanceKey := "instanceKey_example" // string |  (optional)
	userId := "userId_example" // string |  (optional)
	flowKey := "flowKey_example" // string |  (optional)
	flowId := "flowId_example" // string |  (optional)
	universalFlowId := "universalFlowId_example" // string |  (optional)
	enabled := true // bool |  (optional)
	connectionId := "connectionId_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	dependencyInstanceId := "dependencyInstanceId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.FlowInstancesControllerListFlowInstances(context.Background()).Limit(limit).Cursor(cursor).Search(search).ConnectorId(connectorId).IncludeArchived(includeArchived).Id(id).InstanceKey(instanceKey).UserId(userId).FlowKey(flowKey).FlowId(flowId).UniversalFlowId(universalFlowId).Enabled(enabled).ConnectionId(connectionId).IntegrationId(integrationId).IntegrationKey(integrationKey).DependencyInstanceId(dependencyInstanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.FlowInstancesControllerListFlowInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowInstancesControllerListFlowInstances`: FlowInstancesControllerListFlowInstances200Response
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.FlowInstancesControllerListFlowInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlowInstancesControllerListFlowInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **float32** |  | 
 **cursor** | **string** |  | 
 **search** | **string** |  | 
 **connectorId** | **string** |  | 
 **includeArchived** | **bool** |  | 
 **id** | **string** |  | 
 **instanceKey** | **string** |  | 
 **userId** | **string** |  | 
 **flowKey** | **string** |  | 
 **flowId** | **string** |  | 
 **universalFlowId** | **string** |  | 
 **enabled** | **bool** |  | 
 **connectionId** | **string** |  | 
 **integrationId** | **string** |  | 
 **integrationKey** | **string** |  | 
 **dependencyInstanceId** | **string** |  | 

### Return type

[**FlowInstancesControllerListFlowInstances200Response**](FlowInstancesControllerListFlowInstances200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowRunsControllerGetFlowNodeRun

> map[string]interface{} FlowRunsControllerGetFlowNodeRun(ctx, id, nodeKey, nodeRunId).Execute()

Get node run

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/integration-app-client/client"
)

func main() {
	id := "id_example" // string | 
	nodeKey := "nodeKey_example" // string | 
	nodeRunId := "nodeRunId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.FlowRunsControllerGetFlowNodeRun(context.Background(), id, nodeKey, nodeRunId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.FlowRunsControllerGetFlowNodeRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowRunsControllerGetFlowNodeRun`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.FlowRunsControllerGetFlowNodeRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**nodeKey** | **string** |  | 
**nodeRunId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowRunsControllerGetFlowNodeRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowRunsControllerGetFlowRun

> FlowRunDto FlowRunsControllerGetFlowRun(ctx, id).Execute()

Get flow run

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/integration-app-client/client"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.FlowRunsControllerGetFlowRun(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.FlowRunsControllerGetFlowRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowRunsControllerGetFlowRun`: FlowRunDto
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.FlowRunsControllerGetFlowRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowRunsControllerGetFlowRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FlowRunDto**](FlowRunDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowRunsControllerGetFlowRunOutput

> FlowRunsControllerGetFlowRunOutput(ctx, id).Cursor(cursor).Execute()

Get flow run output

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/integration-app-client/client"
)

func main() {
	id := "id_example" // string | 
	cursor := "cursor_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FlowsAPI.FlowRunsControllerGetFlowRunOutput(context.Background(), id).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.FlowRunsControllerGetFlowRunOutput``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowRunsControllerGetFlowRunOutputRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowRunsControllerGetFlowRunOutputForNode

> FlowRunsControllerGetFlowRunOutputForNode(ctx, id, nodeKey).Cursor(cursor).Execute()

Get flow run output for specific node

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/integration-app-client/client"
)

func main() {
	id := "id_example" // string | 
	nodeKey := "nodeKey_example" // string | 
	cursor := "cursor_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FlowsAPI.FlowRunsControllerGetFlowRunOutputForNode(context.Background(), id, nodeKey).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.FlowRunsControllerGetFlowRunOutputForNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**nodeKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowRunsControllerGetFlowRunOutputForNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowRunsControllerGetNodeRunOutput

> map[string]interface{} FlowRunsControllerGetNodeRunOutput(ctx, id, nodeKey, outputId).Execute()

Get node run output

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/integration-app-client/client"
)

func main() {
	id := "id_example" // string | 
	nodeKey := "nodeKey_example" // string | 
	outputId := "outputId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.FlowRunsControllerGetNodeRunOutput(context.Background(), id, nodeKey, outputId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.FlowRunsControllerGetNodeRunOutput``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowRunsControllerGetNodeRunOutput`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.FlowRunsControllerGetNodeRunOutput`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**nodeKey** | **string** |  | 
**outputId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowRunsControllerGetNodeRunOutputRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowRunsControllerListFlowNodeRuns

> FlowRunsControllerListFlowNodeRuns(ctx, id, nodeKey).Cursor(cursor).Execute()

Get node runs

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/integration-app-client/client"
)

func main() {
	id := "id_example" // string | 
	nodeKey := "nodeKey_example" // string | 
	cursor := "cursor_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FlowsAPI.FlowRunsControllerListFlowNodeRuns(context.Background(), id, nodeKey).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.FlowRunsControllerListFlowNodeRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**nodeKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowRunsControllerListFlowNodeRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowRunsControllerListFlowRuns

> FlowRunsControllerListFlowRuns200Response FlowRunsControllerListFlowRuns(ctx).Limit(limit).Cursor(cursor).Search(search).ConnectorId(connectorId).Id(id).FlowInstanceId(flowInstanceId).StartNodeKey(startNodeKey).FlowId(flowId).UniversalFlowId(universalFlowId).UserId(userId).State(state).IntegrationId(integrationId).ConnectionId(connectionId).StartedAfter(startedAfter).Execute()

List flow runs

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/integration-app-client/client"
)

func main() {
	limit := float32(8.14) // float32 |  (optional)
	cursor := "cursor_example" // string |  (optional)
	search := "search_example" // string |  (optional)
	connectorId := "connectorId_example" // string |  (optional)
	id := "id_example" // string |  (optional)
	flowInstanceId := "flowInstanceId_example" // string |  (optional)
	startNodeKey := "startNodeKey_example" // string |  (optional)
	flowId := "flowId_example" // string |  (optional)
	universalFlowId := "universalFlowId_example" // string |  (optional)
	userId := "userId_example" // string |  (optional)
	state := "state_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	connectionId := "connectionId_example" // string |  (optional)
	startedAfter := "startedAfter_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.FlowRunsControllerListFlowRuns(context.Background()).Limit(limit).Cursor(cursor).Search(search).ConnectorId(connectorId).Id(id).FlowInstanceId(flowInstanceId).StartNodeKey(startNodeKey).FlowId(flowId).UniversalFlowId(universalFlowId).UserId(userId).State(state).IntegrationId(integrationId).ConnectionId(connectionId).StartedAfter(startedAfter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.FlowRunsControllerListFlowRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowRunsControllerListFlowRuns`: FlowRunsControllerListFlowRuns200Response
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.FlowRunsControllerListFlowRuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlowRunsControllerListFlowRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **float32** |  | 
 **cursor** | **string** |  | 
 **search** | **string** |  | 
 **connectorId** | **string** |  | 
 **id** | **string** |  | 
 **flowInstanceId** | **string** |  | 
 **startNodeKey** | **string** |  | 
 **flowId** | **string** |  | 
 **universalFlowId** | **string** |  | 
 **userId** | **string** |  | 
 **state** | **string** |  | 
 **integrationId** | **string** |  | 
 **connectionId** | **string** |  | 
 **startedAfter** | **string** |  | 

### Return type

[**FlowRunsControllerListFlowRuns200Response**](FlowRunsControllerListFlowRuns200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowRunsControllerListNodeRunOutputs

> FlowRunsControllerListNodeRunOutputs(ctx, id, nodeKey).Cursor(cursor).Execute()

Get node run outputs

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/integration-app-client/client"
)

func main() {
	id := "id_example" // string | 
	nodeKey := "nodeKey_example" // string | 
	cursor := "cursor_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FlowsAPI.FlowRunsControllerListNodeRunOutputs(context.Background(), id, nodeKey).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.FlowRunsControllerListNodeRunOutputs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**nodeKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowRunsControllerListNodeRunOutputsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowRunsControllerStopFlowRun

> FlowRunsControllerStopFlowRun(ctx, id).Execute()

Stop flow run

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/integration-app-client/client"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FlowsAPI.FlowRunsControllerStopFlowRun(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.FlowRunsControllerStopFlowRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowRunsControllerStopFlowRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowsControllerCreateFlow

> FlowDto FlowsControllerCreateFlow(ctx).CreateFlowDto(createFlowDto).Execute()

Create flow

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/integration-app-client/client"
)

func main() {
	createFlowDto := *openapiclient.NewCreateFlowDto("Key_example", "Name_example") // CreateFlowDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.FlowsControllerCreateFlow(context.Background()).CreateFlowDto(createFlowDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.FlowsControllerCreateFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowsControllerCreateFlow`: FlowDto
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.FlowsControllerCreateFlow`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlowsControllerCreateFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFlowDto** | [**CreateFlowDto**](CreateFlowDto.md) |  | 

### Return type

[**FlowDto**](FlowDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowsControllerListFlows

> FlowsControllerListFlows200Response FlowsControllerListFlows(ctx).Limit(limit).Cursor(cursor).Search(search).ConnectorId(connectorId).IncludeArchived(includeArchived).UniversalFlowId(universalFlowId).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()

List flows

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/integration-app-client/client"
)

func main() {
	limit := float32(8.14) // float32 |  (optional)
	cursor := "cursor_example" // string |  (optional)
	search := "search_example" // string |  (optional)
	connectorId := "connectorId_example" // string |  (optional)
	includeArchived := true // bool |  (optional)
	universalFlowId := "universalFlowId_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.FlowsControllerListFlows(context.Background()).Limit(limit).Cursor(cursor).Search(search).ConnectorId(connectorId).IncludeArchived(includeArchived).UniversalFlowId(universalFlowId).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.FlowsControllerListFlows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowsControllerListFlows`: FlowsControllerListFlows200Response
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.FlowsControllerListFlows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlowsControllerListFlowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **float32** |  | 
 **cursor** | **string** |  | 
 **search** | **string** |  | 
 **connectorId** | **string** |  | 
 **includeArchived** | **bool** |  | 
 **universalFlowId** | **string** |  | 
 **integrationId** | **string** |  | 
 **integrationKey** | **string** |  | 

### Return type

[**FlowsControllerListFlows200Response**](FlowsControllerListFlows200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationLevelFlowControllerArchive

> IntegrationLevelFlowControllerArchive(ctx, flowSelector, integrationSelector).Execute()

Archive flow for integration

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/integration-app-client/client"
)

func main() {
	flowSelector := "flowSelector_example" // string | Flow ID or Key
	integrationSelector := "integrationSelector_example" // string | Integration ID or Key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FlowsAPI.IntegrationLevelFlowControllerArchive(context.Background(), flowSelector, integrationSelector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.IntegrationLevelFlowControllerArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowSelector** | **string** | Flow ID or Key | 
**integrationSelector** | **string** | Integration ID or Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationLevelFlowControllerArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationLevelFlowControllerGet

> FlowDto IntegrationLevelFlowControllerGet(ctx, flowSelector, integrationSelector).Execute()

Get flow for integration

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/integration-app-client/client"
)

func main() {
	flowSelector := "flowSelector_example" // string | Flow ID or Key
	integrationSelector := "integrationSelector_example" // string | Integration ID or Key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.IntegrationLevelFlowControllerGet(context.Background(), flowSelector, integrationSelector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.IntegrationLevelFlowControllerGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationLevelFlowControllerGet`: FlowDto
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.IntegrationLevelFlowControllerGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowSelector** | **string** | Flow ID or Key | 
**integrationSelector** | **string** | Integration ID or Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationLevelFlowControllerGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FlowDto**](FlowDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationLevelFlowControllerPatch

> FlowDto IntegrationLevelFlowControllerPatch(ctx, flowSelector, integrationSelector).UpdateFlowDto(updateFlowDto).Execute()

Patch update flow for integration

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/integration-app-client/client"
)

func main() {
	flowSelector := "flowSelector_example" // string | Flow ID or Key
	integrationSelector := "integrationSelector_example" // string | Integration ID or Key
	updateFlowDto := *openapiclient.NewUpdateFlowDto() // UpdateFlowDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.IntegrationLevelFlowControllerPatch(context.Background(), flowSelector, integrationSelector).UpdateFlowDto(updateFlowDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.IntegrationLevelFlowControllerPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationLevelFlowControllerPatch`: FlowDto
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.IntegrationLevelFlowControllerPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowSelector** | **string** | Flow ID or Key | 
**integrationSelector** | **string** | Integration ID or Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationLevelFlowControllerPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateFlowDto** | [**UpdateFlowDto**](UpdateFlowDto.md) |  | 

### Return type

[**FlowDto**](FlowDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationLevelFlowControllerPut

> FlowDto IntegrationLevelFlowControllerPut(ctx, flowSelector, integrationSelector).UpdateFlowDto(updateFlowDto).Execute()

Update flow for integration

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/integration-app-client/client"
)

func main() {
	flowSelector := "flowSelector_example" // string | Flow ID or Key
	integrationSelector := "integrationSelector_example" // string | Integration ID or Key
	updateFlowDto := *openapiclient.NewUpdateFlowDto() // UpdateFlowDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.IntegrationLevelFlowControllerPut(context.Background(), flowSelector, integrationSelector).UpdateFlowDto(updateFlowDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.IntegrationLevelFlowControllerPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationLevelFlowControllerPut`: FlowDto
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.IntegrationLevelFlowControllerPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowSelector** | **string** | Flow ID or Key | 
**integrationSelector** | **string** | Integration ID or Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationLevelFlowControllerPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateFlowDto** | [**UpdateFlowDto**](UpdateFlowDto.md) |  | 

### Return type

[**FlowDto**](FlowDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationLevelFlowControllerReset

> FlowDto IntegrationLevelFlowControllerReset(ctx, flowSelector, integrationSelector).Execute()

Reset flow for integration

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/integration-app-client/client"
)

func main() {
	flowSelector := "flowSelector_example" // string | Flow ID or Key
	integrationSelector := "integrationSelector_example" // string | Integration ID or Key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.IntegrationLevelFlowControllerReset(context.Background(), flowSelector, integrationSelector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.IntegrationLevelFlowControllerReset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationLevelFlowControllerReset`: FlowDto
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.IntegrationLevelFlowControllerReset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowSelector** | **string** | Flow ID or Key | 
**integrationSelector** | **string** | Integration ID or Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationLevelFlowControllerResetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FlowDto**](FlowDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationLevelFlowsControllerCreate

> FlowDto IntegrationLevelFlowsControllerCreate(ctx, integrationSelector).CreateIntegrationLevelFlowDto(createIntegrationLevelFlowDto).Execute()

Create flow for integration

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/integration-app-client/client"
)

func main() {
	integrationSelector := "integrationSelector_example" // string | Integration ID or Key
	createIntegrationLevelFlowDto := *openapiclient.NewCreateIntegrationLevelFlowDto("Key_example", "Name_example") // CreateIntegrationLevelFlowDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.IntegrationLevelFlowsControllerCreate(context.Background(), integrationSelector).CreateIntegrationLevelFlowDto(createIntegrationLevelFlowDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.IntegrationLevelFlowsControllerCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationLevelFlowsControllerCreate`: FlowDto
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.IntegrationLevelFlowsControllerCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationSelector** | **string** | Integration ID or Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationLevelFlowsControllerCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createIntegrationLevelFlowDto** | [**CreateIntegrationLevelFlowDto**](CreateIntegrationLevelFlowDto.md) |  | 

### Return type

[**FlowDto**](FlowDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationLevelFlowsControllerList

> FlowsControllerListFlows200Response IntegrationLevelFlowsControllerList(ctx, integrationSelector).Limit(limit).Cursor(cursor).Search(search).ConnectorId(connectorId).IncludeArchived(includeArchived).UniversalFlowId(universalFlowId).Execute()

List flows for integration

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/integration-app-client/client"
)

func main() {
	integrationSelector := "integrationSelector_example" // string | Integration ID or Key
	limit := float32(8.14) // float32 |  (optional)
	cursor := "cursor_example" // string |  (optional)
	search := "search_example" // string |  (optional)
	connectorId := "connectorId_example" // string |  (optional)
	includeArchived := true // bool |  (optional)
	universalFlowId := "universalFlowId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.IntegrationLevelFlowsControllerList(context.Background(), integrationSelector).Limit(limit).Cursor(cursor).Search(search).ConnectorId(connectorId).IncludeArchived(includeArchived).UniversalFlowId(universalFlowId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.IntegrationLevelFlowsControllerList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationLevelFlowsControllerList`: FlowsControllerListFlows200Response
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.IntegrationLevelFlowsControllerList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationSelector** | **string** | Integration ID or Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationLevelFlowsControllerListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **float32** |  | 
 **cursor** | **string** |  | 
 **search** | **string** |  | 
 **connectorId** | **string** |  | 
 **includeArchived** | **bool** |  | 
 **universalFlowId** | **string** |  | 

### Return type

[**FlowsControllerListFlows200Response**](FlowsControllerListFlows200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunFlowControllerRunFlow

> RunFlowControllerRunFlow(ctx, flowSelector, connectionSelector).InstanceKey(instanceKey).AutoCreate(autoCreate).Execute()

Run flow instance for connection

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/integration-app-client/client"
)

func main() {
	flowSelector := "flowSelector_example" // string | Flow ID or Key
	connectionSelector := "connectionSelector_example" // string | Integration Key, Connection ID, or Integration ID
	instanceKey := "instanceKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FlowsAPI.RunFlowControllerRunFlow(context.Background(), flowSelector, connectionSelector).InstanceKey(instanceKey).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.RunFlowControllerRunFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowSelector** | **string** | Flow ID or Key | 
**connectionSelector** | **string** | Integration Key, Connection ID, or Integration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRunFlowControllerRunFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **instanceKey** | **string** |  | 
 **autoCreate** | **bool** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

