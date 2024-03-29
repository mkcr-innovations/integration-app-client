# \ScenariosAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveScenario**](ScenariosAPI.md#ArchiveScenario) | **Delete** /scenarios/{id} | 
[**ArchiveScenarios**](ScenariosAPI.md#ArchiveScenarios) | **Delete** /scenario | 
[**CreateScenario**](ScenariosAPI.md#CreateScenario) | **Post** /scenarios | 
[**GetScenario**](ScenariosAPI.md#GetScenario) | **Get** /scenarios/{id} | 
[**GetScenarios**](ScenariosAPI.md#GetScenarios) | **Get** /scenario | 
[**ListScenarios**](ScenariosAPI.md#ListScenarios) | **Get** /scenarios | 
[**PatchScenario**](ScenariosAPI.md#PatchScenario) | **Patch** /scenarios/{id} | 
[**PatchScenarios**](ScenariosAPI.md#PatchScenarios) | **Patch** /scenario | 
[**PutScenario**](ScenariosAPI.md#PutScenario) | **Put** /scenarios/{id} | 
[**PutScenarios**](ScenariosAPI.md#PutScenarios) | **Put** /scenario | 



## ArchiveScenario

> ArchiveScenario(ctx, id).Id2(id2).Key(key).ScenarioTemplateId(scenarioTemplateId).Execute()



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
	id2 := "id_example" // string |  (optional)
	key := "key_example" // string |  (optional)
	scenarioTemplateId := "scenarioTemplateId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ScenariosAPI.ArchiveScenario(context.Background(), id).Id2(id2).Key(key).ScenarioTemplateId(scenarioTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScenariosAPI.ArchiveScenario``: %v\n", err)
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

Other parameters are passed through a pointer to a apiArchiveScenarioRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id2** | **string** |  | 
 **key** | **string** |  | 
 **scenarioTemplateId** | **string** |  | 

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


## ArchiveScenarios

> ArchiveScenarios(ctx).Id(id).Key(key).ScenarioTemplateId(scenarioTemplateId).Execute()



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
	scenarioTemplateId := "scenarioTemplateId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ScenariosAPI.ArchiveScenarios(context.Background()).Id(id).Key(key).ScenarioTemplateId(scenarioTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScenariosAPI.ArchiveScenarios``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArchiveScenariosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **key** | **string** |  | 
 **scenarioTemplateId** | **string** |  | 

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


## CreateScenario

> ScenarioDto CreateScenario(ctx).CreateScenarioDto(createScenarioDto).Execute()



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
	createScenarioDto := *openapiclient.NewCreateScenarioDto("Name_example", "Key_example") // CreateScenarioDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScenariosAPI.CreateScenario(context.Background()).CreateScenarioDto(createScenarioDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScenariosAPI.CreateScenario``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateScenario`: ScenarioDto
	fmt.Fprintf(os.Stdout, "Response from `ScenariosAPI.CreateScenario`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateScenarioRequest struct via the builder pattern


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


## GetScenario

> ScenarioDto GetScenario(ctx, id).Id2(id2).Key(key).ScenarioTemplateId(scenarioTemplateId).Execute()



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
	id2 := "id_example" // string |  (optional)
	key := "key_example" // string |  (optional)
	scenarioTemplateId := "scenarioTemplateId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScenariosAPI.GetScenario(context.Background(), id).Id2(id2).Key(key).ScenarioTemplateId(scenarioTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScenariosAPI.GetScenario``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetScenario`: ScenarioDto
	fmt.Fprintf(os.Stdout, "Response from `ScenariosAPI.GetScenario`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetScenarioRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id2** | **string** |  | 
 **key** | **string** |  | 
 **scenarioTemplateId** | **string** |  | 

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


## GetScenarios

> ScenarioDto GetScenarios(ctx).Id(id).Key(key).ScenarioTemplateId(scenarioTemplateId).Execute()



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
	scenarioTemplateId := "scenarioTemplateId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScenariosAPI.GetScenarios(context.Background()).Id(id).Key(key).ScenarioTemplateId(scenarioTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScenariosAPI.GetScenarios``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetScenarios`: ScenarioDto
	fmt.Fprintf(os.Stdout, "Response from `ScenariosAPI.GetScenarios`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetScenariosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **key** | **string** |  | 
 **scenarioTemplateId** | **string** |  | 

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


## ListScenarios

> ListScenarios(ctx).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ScenariosAPI.ListScenarios(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScenariosAPI.ListScenarios``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListScenariosRequest struct via the builder pattern


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


## PatchScenario

> ScenarioDto PatchScenario(ctx, id).UpdateScenarioDto(updateScenarioDto).Id2(id2).Key(key).ScenarioTemplateId(scenarioTemplateId).Execute()



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
	id2 := "id_example" // string |  (optional)
	key := "key_example" // string |  (optional)
	scenarioTemplateId := "scenarioTemplateId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScenariosAPI.PatchScenario(context.Background(), id).UpdateScenarioDto(updateScenarioDto).Id2(id2).Key(key).ScenarioTemplateId(scenarioTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScenariosAPI.PatchScenario``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchScenario`: ScenarioDto
	fmt.Fprintf(os.Stdout, "Response from `ScenariosAPI.PatchScenario`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchScenarioRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateScenarioDto** | [**UpdateScenarioDto**](UpdateScenarioDto.md) |  | 
 **id2** | **string** |  | 
 **key** | **string** |  | 
 **scenarioTemplateId** | **string** |  | 

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


## PatchScenarios

> ScenarioDto PatchScenarios(ctx).UpdateScenarioDto(updateScenarioDto).Id(id).Key(key).ScenarioTemplateId(scenarioTemplateId).Execute()



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
	updateScenarioDto := *openapiclient.NewUpdateScenarioDto() // UpdateScenarioDto | 
	id := "id_example" // string |  (optional)
	key := "key_example" // string |  (optional)
	scenarioTemplateId := "scenarioTemplateId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScenariosAPI.PatchScenarios(context.Background()).UpdateScenarioDto(updateScenarioDto).Id(id).Key(key).ScenarioTemplateId(scenarioTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScenariosAPI.PatchScenarios``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchScenarios`: ScenarioDto
	fmt.Fprintf(os.Stdout, "Response from `ScenariosAPI.PatchScenarios`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchScenariosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateScenarioDto** | [**UpdateScenarioDto**](UpdateScenarioDto.md) |  | 
 **id** | **string** |  | 
 **key** | **string** |  | 
 **scenarioTemplateId** | **string** |  | 

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


## PutScenario

> ScenarioDto PutScenario(ctx, id).UpdateScenarioDto(updateScenarioDto).Id2(id2).Key(key).ScenarioTemplateId(scenarioTemplateId).Execute()



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
	id2 := "id_example" // string |  (optional)
	key := "key_example" // string |  (optional)
	scenarioTemplateId := "scenarioTemplateId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScenariosAPI.PutScenario(context.Background(), id).UpdateScenarioDto(updateScenarioDto).Id2(id2).Key(key).ScenarioTemplateId(scenarioTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScenariosAPI.PutScenario``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutScenario`: ScenarioDto
	fmt.Fprintf(os.Stdout, "Response from `ScenariosAPI.PutScenario`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutScenarioRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateScenarioDto** | [**UpdateScenarioDto**](UpdateScenarioDto.md) |  | 
 **id2** | **string** |  | 
 **key** | **string** |  | 
 **scenarioTemplateId** | **string** |  | 

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


## PutScenarios

> ScenarioDto PutScenarios(ctx).UpdateScenarioDto(updateScenarioDto).Id(id).Key(key).ScenarioTemplateId(scenarioTemplateId).Execute()



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
	updateScenarioDto := *openapiclient.NewUpdateScenarioDto() // UpdateScenarioDto | 
	id := "id_example" // string |  (optional)
	key := "key_example" // string |  (optional)
	scenarioTemplateId := "scenarioTemplateId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScenariosAPI.PutScenarios(context.Background()).UpdateScenarioDto(updateScenarioDto).Id(id).Key(key).ScenarioTemplateId(scenarioTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScenariosAPI.PutScenarios``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutScenarios`: ScenarioDto
	fmt.Fprintf(os.Stdout, "Response from `ScenariosAPI.PutScenarios`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutScenariosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateScenarioDto** | [**UpdateScenarioDto**](UpdateScenarioDto.md) |  | 
 **id** | **string** |  | 
 **key** | **string** |  | 
 **scenarioTemplateId** | **string** |  | 

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

