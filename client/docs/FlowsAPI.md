# \FlowsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyFlow**](FlowsAPI.md#ApplyFlow) | **Post** /flows/{id}/apply | 
[**ApplyFlows**](FlowsAPI.md#ApplyFlows) | **Post** /flow/apply | 
[**ArchiveFlow**](FlowsAPI.md#ArchiveFlow) | **Delete** /flows/{id} | 
[**ArchiveFlowInstance**](FlowsAPI.md#ArchiveFlowInstance) | **Delete** /flow-instance | 
[**ArchiveFlows**](FlowsAPI.md#ArchiveFlows) | **Delete** /flow | 
[**CreateFlow**](FlowsAPI.md#CreateFlow) | **Post** /flows | 
[**CreateFlowInstance**](FlowsAPI.md#CreateFlowInstance) | **Post** /flow-instance | 
[**GetFlow**](FlowsAPI.md#GetFlow) | **Get** /flows/{id} | 
[**GetFlowInstance**](FlowsAPI.md#GetFlowInstance) | **Get** /flow-instance | 
[**GetFlowNodeRun**](FlowsAPI.md#GetFlowNodeRun) | **Get** /flow-runs/{id}/nodes/{nodeKey}/runs/{nodeRunId} | 
[**GetFlowRun**](FlowsAPI.md#GetFlowRun) | **Get** /flow-runs/{id} | 
[**GetFlowRunOutput**](FlowsAPI.md#GetFlowRunOutput) | **Get** /flow-runs/{id}/output | 
[**GetFlows**](FlowsAPI.md#GetFlows) | **Get** /flow | 
[**GetNodeRunOutput**](FlowsAPI.md#GetNodeRunOutput) | **Get** /flow-runs/{id}/nodes/{nodeKey}/outputs/{outputId} | 
[**ListFlowInstances**](FlowsAPI.md#ListFlowInstances) | **Get** /flow-instances | 
[**ListFlowNodeRuns**](FlowsAPI.md#ListFlowNodeRuns) | **Get** /flow-runs/{id}/nodes/{nodeKey}/runs | 
[**ListFlowRuns**](FlowsAPI.md#ListFlowRuns) | **Get** /flow-runs | 
[**ListFlows**](FlowsAPI.md#ListFlows) | **Get** /flows | 
[**ListNodeRunOutputs**](FlowsAPI.md#ListNodeRunOutputs) | **Get** /flow-runs/{id}/nodes/{nodeKey}/outputs | 
[**PatchFlow**](FlowsAPI.md#PatchFlow) | **Patch** /flows/{id} | 
[**PatchFlowInstance**](FlowsAPI.md#PatchFlowInstance) | **Patch** /flow-instance | 
[**PatchFlows**](FlowsAPI.md#PatchFlows) | **Patch** /flow | 
[**PutFlow**](FlowsAPI.md#PutFlow) | **Put** /flows/{id} | 
[**PutFlows**](FlowsAPI.md#PutFlows) | **Put** /flow | 
[**ResetFlow**](FlowsAPI.md#ResetFlow) | **Post** /flows/{id}/reset | 
[**ResetFlowInstance**](FlowsAPI.md#ResetFlowInstance) | **Post** /flow-instance/reset | 
[**ResetFlows**](FlowsAPI.md#ResetFlows) | **Post** /flow/reset | 
[**RunFlow**](FlowsAPI.md#RunFlow) | **Post** /connections/{connectionIdOrKey}/flows/{flowKey}/run | 
[**SetupFlowInstance**](FlowsAPI.md#SetupFlowInstance) | **Post** /flow-instance/setup | 
[**StopFlowRun**](FlowsAPI.md#StopFlowRun) | **Post** /flow-runs/{id}/stop | 
[**UpdateFlowInstance**](FlowsAPI.md#UpdateFlowInstance) | **Put** /flow-instance | 



## ApplyFlow

> []FlowDto ApplyFlow(ctx, id).Id2(id2).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).AutoCreate(autoCreate).Execute()



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
	id := "id_example" // string | The ID of the flow to apply
	id2 := "id_example" // string |  (optional)
	key := "key_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.ApplyFlow(context.Background(), id).Id2(id2).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.ApplyFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApplyFlow`: []FlowDto
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.ApplyFlow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the flow to apply | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplyFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id2** | **string** |  | 
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


## ApplyFlows

> []FlowDto ApplyFlows(ctx).Id(id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).AutoCreate(autoCreate).Execute()



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
	id := "id_example" // string |  (optional)
	key := "key_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.ApplyFlows(context.Background()).Id(id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.ApplyFlows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApplyFlows`: []FlowDto
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.ApplyFlows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApplyFlowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
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


## ArchiveFlow

> ArchiveFlow(ctx, id).Id2(id2).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).AutoCreate(autoCreate).Execute()



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
	id := "id_example" // string | The ID of the flow to retrive
	id2 := "id_example" // string |  (optional)
	key := "key_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FlowsAPI.ArchiveFlow(context.Background(), id).Id2(id2).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.ArchiveFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the flow to retrive | 

### Other Parameters

Other parameters are passed through a pointer to a apiArchiveFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id2** | **string** |  | 
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


## ArchiveFlowInstance

> ArchiveFlowInstance(ctx).Id(id).FlowId(flowId).FlowKey(flowKey).InstanceKey(instanceKey).AutoCreate(autoCreate).AutoUpdate(autoUpdate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()



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
	id := "id_example" // string |  (optional)
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
	r, err := apiClient.FlowsAPI.ArchiveFlowInstance(context.Background()).Id(id).FlowId(flowId).FlowKey(flowKey).InstanceKey(instanceKey).AutoCreate(autoCreate).AutoUpdate(autoUpdate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.ArchiveFlowInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArchiveFlowInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
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


## ArchiveFlows

> ArchiveFlows(ctx).Id(id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).AutoCreate(autoCreate).Execute()



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
	id := "id_example" // string |  (optional)
	key := "key_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FlowsAPI.ArchiveFlows(context.Background()).Id(id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.ArchiveFlows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArchiveFlowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
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


## CreateFlow

> FlowDto CreateFlow(ctx).CreateFlowDto(createFlowDto).Execute()



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
	resp, r, err := apiClient.FlowsAPI.CreateFlow(context.Background()).CreateFlowDto(createFlowDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.CreateFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFlow`: FlowDto
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.CreateFlow`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFlowRequest struct via the builder pattern


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


## CreateFlowInstance

> FlowInstanceDto CreateFlowInstance(ctx).UpdateFlowInstanceRequestDto(updateFlowInstanceRequestDto).Id(id).FlowId(flowId).FlowKey(flowKey).InstanceKey(instanceKey).AutoCreate(autoCreate).AutoUpdate(autoUpdate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()



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
	updateFlowInstanceRequestDto := *openapiclient.NewUpdateFlowInstanceRequestDto() // UpdateFlowInstanceRequestDto | 
	id := "id_example" // string |  (optional)
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
	resp, r, err := apiClient.FlowsAPI.CreateFlowInstance(context.Background()).UpdateFlowInstanceRequestDto(updateFlowInstanceRequestDto).Id(id).FlowId(flowId).FlowKey(flowKey).InstanceKey(instanceKey).AutoCreate(autoCreate).AutoUpdate(autoUpdate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.CreateFlowInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFlowInstance`: FlowInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.CreateFlowInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFlowInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateFlowInstanceRequestDto** | [**UpdateFlowInstanceRequestDto**](UpdateFlowInstanceRequestDto.md) |  | 
 **id** | **string** |  | 
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


## GetFlow

> FlowDto GetFlow(ctx, id).Id2(id2).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).AutoCreate(autoCreate).Execute()



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
	id := "id_example" // string | The ID of the flow to retrive
	id2 := "id_example" // string |  (optional)
	key := "key_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.GetFlow(context.Background(), id).Id2(id2).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.GetFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFlow`: FlowDto
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.GetFlow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the flow to retrive | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id2** | **string** |  | 
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


## GetFlowInstance

> map[string]interface{} GetFlowInstance(ctx).Id(id).FlowId(flowId).FlowKey(flowKey).InstanceKey(instanceKey).AutoCreate(autoCreate).AutoUpdate(autoUpdate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()



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
	id := "id_example" // string |  (optional)
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
	resp, r, err := apiClient.FlowsAPI.GetFlowInstance(context.Background()).Id(id).FlowId(flowId).FlowKey(flowKey).InstanceKey(instanceKey).AutoCreate(autoCreate).AutoUpdate(autoUpdate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.GetFlowInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFlowInstance`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.GetFlowInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFlowInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
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


## GetFlowNodeRun

> map[string]interface{} GetFlowNodeRun(ctx, id, nodeKey, nodeRunId).Execute()



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
	resp, r, err := apiClient.FlowsAPI.GetFlowNodeRun(context.Background(), id, nodeKey, nodeRunId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.GetFlowNodeRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFlowNodeRun`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.GetFlowNodeRun`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetFlowNodeRunRequest struct via the builder pattern


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


## GetFlowRun

> FlowRunDto GetFlowRun(ctx, id).Execute()



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
	resp, r, err := apiClient.FlowsAPI.GetFlowRun(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.GetFlowRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFlowRun`: FlowRunDto
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.GetFlowRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFlowRunRequest struct via the builder pattern


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


## GetFlowRunOutput

> GetFlowRunOutput(ctx, id).Cursor(cursor).Execute()



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
	r, err := apiClient.FlowsAPI.GetFlowRunOutput(context.Background(), id).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.GetFlowRunOutput``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetFlowRunOutputRequest struct via the builder pattern


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


## GetFlows

> FlowDto GetFlows(ctx).Id(id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).AutoCreate(autoCreate).Execute()



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
	id := "id_example" // string |  (optional)
	key := "key_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.GetFlows(context.Background()).Id(id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.GetFlows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFlows`: FlowDto
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.GetFlows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFlowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
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


## GetNodeRunOutput

> map[string]interface{} GetNodeRunOutput(ctx, id, nodeKey, outputId).Execute()



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
	resp, r, err := apiClient.FlowsAPI.GetNodeRunOutput(context.Background(), id, nodeKey, outputId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.GetNodeRunOutput``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNodeRunOutput`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.GetNodeRunOutput`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetNodeRunOutputRequest struct via the builder pattern


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


## ListFlowInstances

> ListFlowInstances200Response ListFlowInstances(ctx).Id(id).UserId(userId).FlowId(flowId).UniversalFlowId(universalFlowId).FlowKey(flowKey).ConnectionId(connectionId).IntegrationKey(integrationKey).IntegrationId(integrationId).InstanceKey(instanceKey).Enabled(enabled).DependencyInstanceId(dependencyInstanceId).Execute()



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
	id := "id_example" // string |  (optional)
	userId := "userId_example" // string |  (optional)
	flowId := "flowId_example" // string |  (optional)
	universalFlowId := "universalFlowId_example" // string |  (optional)
	flowKey := "flowKey_example" // string |  (optional)
	connectionId := "connectionId_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	instanceKey := "instanceKey_example" // string |  (optional)
	enabled := true // bool |  (optional)
	dependencyInstanceId := "dependencyInstanceId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.ListFlowInstances(context.Background()).Id(id).UserId(userId).FlowId(flowId).UniversalFlowId(universalFlowId).FlowKey(flowKey).ConnectionId(connectionId).IntegrationKey(integrationKey).IntegrationId(integrationId).InstanceKey(instanceKey).Enabled(enabled).DependencyInstanceId(dependencyInstanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.ListFlowInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFlowInstances`: ListFlowInstances200Response
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.ListFlowInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFlowInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **userId** | **string** |  | 
 **flowId** | **string** |  | 
 **universalFlowId** | **string** |  | 
 **flowKey** | **string** |  | 
 **connectionId** | **string** |  | 
 **integrationKey** | **string** |  | 
 **integrationId** | **string** |  | 
 **instanceKey** | **string** |  | 
 **enabled** | **bool** |  | 
 **dependencyInstanceId** | **string** |  | 

### Return type

[**ListFlowInstances200Response**](ListFlowInstances200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFlowNodeRuns

> ListFlowNodeRuns(ctx, id, nodeKey).Cursor(cursor).Execute()



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
	r, err := apiClient.FlowsAPI.ListFlowNodeRuns(context.Background(), id, nodeKey).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.ListFlowNodeRuns``: %v\n", err)
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

Other parameters are passed through a pointer to a apiListFlowNodeRunsRequest struct via the builder pattern


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


## ListFlowRuns

> ListFlowRuns200Response ListFlowRuns(ctx).Id(id).FlowInstanceId(flowInstanceId).StartNodeKey(startNodeKey).FlowId(flowId).UniversalFlowId(universalFlowId).UserId(userId).State(state).IntegrationId(integrationId).ConnectionId(connectionId).StartedAfter(startedAfter).Execute()



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
	resp, r, err := apiClient.FlowsAPI.ListFlowRuns(context.Background()).Id(id).FlowInstanceId(flowInstanceId).StartNodeKey(startNodeKey).FlowId(flowId).UniversalFlowId(universalFlowId).UserId(userId).State(state).IntegrationId(integrationId).ConnectionId(connectionId).StartedAfter(startedAfter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.ListFlowRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFlowRuns`: ListFlowRuns200Response
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.ListFlowRuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFlowRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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

[**ListFlowRuns200Response**](ListFlowRuns200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFlows

> ListFlows200Response ListFlows(ctx).IntegrationKey(integrationKey).IntegrationId(integrationId).UniversalFlowId(universalFlowId).Execute()



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
	integrationKey := "integrationKey_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	universalFlowId := "universalFlowId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.ListFlows(context.Background()).IntegrationKey(integrationKey).IntegrationId(integrationId).UniversalFlowId(universalFlowId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.ListFlows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFlows`: ListFlows200Response
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.ListFlows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFlowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **integrationKey** | **string** |  | 
 **integrationId** | **string** |  | 
 **universalFlowId** | **string** |  | 

### Return type

[**ListFlows200Response**](ListFlows200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNodeRunOutputs

> ListNodeRunOutputs(ctx, id, nodeKey).Cursor(cursor).Execute()



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
	r, err := apiClient.FlowsAPI.ListNodeRunOutputs(context.Background(), id, nodeKey).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.ListNodeRunOutputs``: %v\n", err)
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

Other parameters are passed through a pointer to a apiListNodeRunOutputsRequest struct via the builder pattern


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


## PatchFlow

> FlowDto PatchFlow(ctx, id).UpdateFlowDto(updateFlowDto).Id2(id2).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).AutoCreate(autoCreate).Execute()



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
	id := "id_example" // string | The ID of the flow to retrive
	updateFlowDto := *openapiclient.NewUpdateFlowDto() // UpdateFlowDto | 
	id2 := "id_example" // string |  (optional)
	key := "key_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.PatchFlow(context.Background(), id).UpdateFlowDto(updateFlowDto).Id2(id2).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.PatchFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchFlow`: FlowDto
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.PatchFlow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the flow to retrive | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateFlowDto** | [**UpdateFlowDto**](UpdateFlowDto.md) |  | 
 **id2** | **string** |  | 
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


## PatchFlowInstance

> FlowInstanceDto PatchFlowInstance(ctx).UpdateFlowInstanceRequestDto(updateFlowInstanceRequestDto).Id(id).FlowId(flowId).FlowKey(flowKey).InstanceKey(instanceKey).AutoCreate(autoCreate).AutoUpdate(autoUpdate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()



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
	updateFlowInstanceRequestDto := *openapiclient.NewUpdateFlowInstanceRequestDto() // UpdateFlowInstanceRequestDto | 
	id := "id_example" // string |  (optional)
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
	resp, r, err := apiClient.FlowsAPI.PatchFlowInstance(context.Background()).UpdateFlowInstanceRequestDto(updateFlowInstanceRequestDto).Id(id).FlowId(flowId).FlowKey(flowKey).InstanceKey(instanceKey).AutoCreate(autoCreate).AutoUpdate(autoUpdate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.PatchFlowInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchFlowInstance`: FlowInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.PatchFlowInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchFlowInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateFlowInstanceRequestDto** | [**UpdateFlowInstanceRequestDto**](UpdateFlowInstanceRequestDto.md) |  | 
 **id** | **string** |  | 
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


## PatchFlows

> FlowDto PatchFlows(ctx).UpdateFlowDto(updateFlowDto).Id(id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).AutoCreate(autoCreate).Execute()



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
	updateFlowDto := *openapiclient.NewUpdateFlowDto() // UpdateFlowDto | 
	id := "id_example" // string |  (optional)
	key := "key_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.PatchFlows(context.Background()).UpdateFlowDto(updateFlowDto).Id(id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.PatchFlows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchFlows`: FlowDto
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.PatchFlows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchFlowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateFlowDto** | [**UpdateFlowDto**](UpdateFlowDto.md) |  | 
 **id** | **string** |  | 
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


## PutFlow

> FlowDto PutFlow(ctx, id).UpdateFlowDto(updateFlowDto).Id2(id2).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).AutoCreate(autoCreate).Execute()



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
	id := "id_example" // string | The ID of the flow to retrive
	updateFlowDto := *openapiclient.NewUpdateFlowDto() // UpdateFlowDto | 
	id2 := "id_example" // string |  (optional)
	key := "key_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.PutFlow(context.Background(), id).UpdateFlowDto(updateFlowDto).Id2(id2).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.PutFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutFlow`: FlowDto
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.PutFlow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the flow to retrive | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateFlowDto** | [**UpdateFlowDto**](UpdateFlowDto.md) |  | 
 **id2** | **string** |  | 
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


## PutFlows

> FlowDto PutFlows(ctx).UpdateFlowDto(updateFlowDto).Id(id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).AutoCreate(autoCreate).Execute()



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
	updateFlowDto := *openapiclient.NewUpdateFlowDto() // UpdateFlowDto | 
	id := "id_example" // string |  (optional)
	key := "key_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.PutFlows(context.Background()).UpdateFlowDto(updateFlowDto).Id(id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.PutFlows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutFlows`: FlowDto
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.PutFlows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutFlowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateFlowDto** | [**UpdateFlowDto**](UpdateFlowDto.md) |  | 
 **id** | **string** |  | 
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


## ResetFlow

> FlowDto ResetFlow(ctx, id).Id2(id2).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).AutoCreate(autoCreate).Execute()



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
	id := "id_example" // string | The ID of the flow to reset
	id2 := "id_example" // string |  (optional)
	key := "key_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.ResetFlow(context.Background(), id).Id2(id2).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.ResetFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResetFlow`: FlowDto
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.ResetFlow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the flow to reset | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id2** | **string** |  | 
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


## ResetFlowInstance

> FlowInstanceDto ResetFlowInstance(ctx).Id(id).FlowId(flowId).FlowKey(flowKey).InstanceKey(instanceKey).AutoCreate(autoCreate).AutoUpdate(autoUpdate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()



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
	id := "id_example" // string |  (optional)
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
	resp, r, err := apiClient.FlowsAPI.ResetFlowInstance(context.Background()).Id(id).FlowId(flowId).FlowKey(flowKey).InstanceKey(instanceKey).AutoCreate(autoCreate).AutoUpdate(autoUpdate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.ResetFlowInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResetFlowInstance`: FlowInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.ResetFlowInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResetFlowInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
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


## ResetFlows

> FlowDto ResetFlows(ctx).Id(id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).AutoCreate(autoCreate).Execute()



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
	id := "id_example" // string |  (optional)
	key := "key_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.ResetFlows(context.Background()).Id(id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.ResetFlows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResetFlows`: FlowDto
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.ResetFlows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResetFlowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
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


## RunFlow

> RunFlow(ctx, connectionIdOrKey, flowKey).RunFlowPayload(runFlowPayload).Execute()



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
	connectionIdOrKey := "connectionIdOrKey_example" // string | 
	flowKey := "flowKey_example" // string | 
	runFlowPayload := *openapiclient.NewRunFlowPayload() // RunFlowPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FlowsAPI.RunFlow(context.Background(), connectionIdOrKey, flowKey).RunFlowPayload(runFlowPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.RunFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionIdOrKey** | **string** |  | 
**flowKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRunFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **runFlowPayload** | [**RunFlowPayload**](RunFlowPayload.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetupFlowInstance

> FlowInstanceDto SetupFlowInstance(ctx).Id(id).FlowId(flowId).FlowKey(flowKey).InstanceKey(instanceKey).AutoCreate(autoCreate).AutoUpdate(autoUpdate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()



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
	id := "id_example" // string |  (optional)
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
	resp, r, err := apiClient.FlowsAPI.SetupFlowInstance(context.Background()).Id(id).FlowId(flowId).FlowKey(flowKey).InstanceKey(instanceKey).AutoCreate(autoCreate).AutoUpdate(autoUpdate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.SetupFlowInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetupFlowInstance`: FlowInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.SetupFlowInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetupFlowInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
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


## StopFlowRun

> StopFlowRun(ctx, id).Execute()



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
	r, err := apiClient.FlowsAPI.StopFlowRun(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.StopFlowRun``: %v\n", err)
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

Other parameters are passed through a pointer to a apiStopFlowRunRequest struct via the builder pattern


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


## UpdateFlowInstance

> FlowInstanceDto UpdateFlowInstance(ctx).UpdateFlowInstanceRequestDto(updateFlowInstanceRequestDto).Id(id).FlowId(flowId).FlowKey(flowKey).InstanceKey(instanceKey).AutoCreate(autoCreate).AutoUpdate(autoUpdate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()



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
	updateFlowInstanceRequestDto := *openapiclient.NewUpdateFlowInstanceRequestDto() // UpdateFlowInstanceRequestDto | 
	id := "id_example" // string |  (optional)
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
	resp, r, err := apiClient.FlowsAPI.UpdateFlowInstance(context.Background()).UpdateFlowInstanceRequestDto(updateFlowInstanceRequestDto).Id(id).FlowId(flowId).FlowKey(flowKey).InstanceKey(instanceKey).AutoCreate(autoCreate).AutoUpdate(autoUpdate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.UpdateFlowInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFlowInstance`: FlowInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.UpdateFlowInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFlowInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateFlowInstanceRequestDto** | [**UpdateFlowInstanceRequestDto**](UpdateFlowInstanceRequestDto.md) |  | 
 **id** | **string** |  | 
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

