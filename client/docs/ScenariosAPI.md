# \ScenariosAPI

All URIs are relative to *https://api.integration.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ScenarioByIdControllerArchive**](ScenariosAPI.md#ScenarioByIdControllerArchive) | **Delete** /scenarios/{id} | Archive scenario
[**ScenarioByIdControllerExport**](ScenariosAPI.md#ScenarioByIdControllerExport) | **Get** /scenarios/{id}/export | 
[**ScenarioByIdControllerGet**](ScenariosAPI.md#ScenarioByIdControllerGet) | **Get** /scenarios/{id} | Get scenario
[**ScenarioByIdControllerPatch**](ScenariosAPI.md#ScenarioByIdControllerPatch) | **Patch** /scenarios/{id} | Patch scenario
[**ScenarioByIdControllerPut**](ScenariosAPI.md#ScenarioByIdControllerPut) | **Put** /scenarios/{id} | Put scenario
[**ScenariosControllerCreate**](ScenariosAPI.md#ScenariosControllerCreate) | **Post** /scenarios | Create scenario
[**ScenariosControllerList**](ScenariosAPI.md#ScenariosControllerList) | **Get** /scenarios | List scenarios



## ScenarioByIdControllerArchive

> ScenarioByIdControllerArchive(ctx, id).Execute()

Archive scenario

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
	r, err := apiClient.ScenariosAPI.ScenarioByIdControllerArchive(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScenariosAPI.ScenarioByIdControllerArchive``: %v\n", err)
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

Other parameters are passed through a pointer to a apiScenarioByIdControllerArchiveRequest struct via the builder pattern


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


## ScenarioByIdControllerExport

> ScenarioExportDto ScenarioByIdControllerExport(ctx, id).Execute()



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
	resp, r, err := apiClient.ScenariosAPI.ScenarioByIdControllerExport(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScenariosAPI.ScenarioByIdControllerExport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScenarioByIdControllerExport`: ScenarioExportDto
	fmt.Fprintf(os.Stdout, "Response from `ScenariosAPI.ScenarioByIdControllerExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiScenarioByIdControllerExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ScenarioExportDto**](ScenarioExportDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScenarioByIdControllerGet

> ScenarioDto ScenarioByIdControllerGet(ctx, id).Execute()

Get scenario

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
	resp, r, err := apiClient.ScenariosAPI.ScenarioByIdControllerGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScenariosAPI.ScenarioByIdControllerGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScenarioByIdControllerGet`: ScenarioDto
	fmt.Fprintf(os.Stdout, "Response from `ScenariosAPI.ScenarioByIdControllerGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiScenarioByIdControllerGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ScenarioDto**](ScenarioDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScenarioByIdControllerPatch

> ScenarioDto ScenarioByIdControllerPatch(ctx, id).UpdateScenarioDto(updateScenarioDto).Execute()

Patch scenario

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
	updateScenarioDto := *openapiclient.NewUpdateScenarioDto() // UpdateScenarioDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScenariosAPI.ScenarioByIdControllerPatch(context.Background(), id).UpdateScenarioDto(updateScenarioDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScenariosAPI.ScenarioByIdControllerPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScenarioByIdControllerPatch`: ScenarioDto
	fmt.Fprintf(os.Stdout, "Response from `ScenariosAPI.ScenarioByIdControllerPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiScenarioByIdControllerPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateScenarioDto** | [**UpdateScenarioDto**](UpdateScenarioDto.md) |  | 

### Return type

[**ScenarioDto**](ScenarioDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScenarioByIdControllerPut

> ScenarioDto ScenarioByIdControllerPut(ctx, id).UpdateScenarioDto(updateScenarioDto).Execute()

Put scenario

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
	updateScenarioDto := *openapiclient.NewUpdateScenarioDto() // UpdateScenarioDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScenariosAPI.ScenarioByIdControllerPut(context.Background(), id).UpdateScenarioDto(updateScenarioDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScenariosAPI.ScenarioByIdControllerPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScenarioByIdControllerPut`: ScenarioDto
	fmt.Fprintf(os.Stdout, "Response from `ScenariosAPI.ScenarioByIdControllerPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiScenarioByIdControllerPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateScenarioDto** | [**UpdateScenarioDto**](UpdateScenarioDto.md) |  | 

### Return type

[**ScenarioDto**](ScenarioDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScenariosControllerCreate

> ScenarioDto ScenariosControllerCreate(ctx).CreateScenarioDto(createScenarioDto).Execute()

Create scenario

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
	createScenarioDto := *openapiclient.NewCreateScenarioDto() // CreateScenarioDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScenariosAPI.ScenariosControllerCreate(context.Background()).CreateScenarioDto(createScenarioDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScenariosAPI.ScenariosControllerCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScenariosControllerCreate`: ScenarioDto
	fmt.Fprintf(os.Stdout, "Response from `ScenariosAPI.ScenariosControllerCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiScenariosControllerCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createScenarioDto** | [**CreateScenarioDto**](CreateScenarioDto.md) |  | 

### Return type

[**ScenarioDto**](ScenarioDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScenariosControllerList

> ScenariosControllerList200Response ScenariosControllerList(ctx).Limit(limit).Cursor(cursor).Search(search).ConnectorId(connectorId).Execute()

List scenarios

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScenariosAPI.ScenariosControllerList(context.Background()).Limit(limit).Cursor(cursor).Search(search).ConnectorId(connectorId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScenariosAPI.ScenariosControllerList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScenariosControllerList`: ScenariosControllerList200Response
	fmt.Fprintf(os.Stdout, "Response from `ScenariosAPI.ScenariosControllerList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiScenariosControllerListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **float32** |  | 
 **cursor** | **string** |  | 
 **search** | **string** |  | 
 **connectorId** | **string** |  | 

### Return type

[**ScenariosControllerList200Response**](ScenariosControllerList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

