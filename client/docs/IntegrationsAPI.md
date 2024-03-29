# \IntegrationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveIntegration**](IntegrationsAPI.md#ArchiveIntegration) | **Delete** /integrations/{isOrKey} | 
[**CreateIntegration**](IntegrationsAPI.md#CreateIntegration) | **Post** /integrations | 
[**DisconnectIntegration**](IntegrationsAPI.md#DisconnectIntegration) | **Post** /integrations/{idOrKey}/disconnect | 
[**GetConnectorSpec**](IntegrationsAPI.md#GetConnectorSpec) | **Get** /integrations/{idOrKey}/connector-spec | 
[**GetIntegration**](IntegrationsAPI.md#GetIntegration) | **Get** /integrations/{idOrKey} | 
[**GetIntegrationDataCollectionSpecs**](IntegrationsAPI.md#GetIntegrationDataCollectionSpecs) | **Get** /integrations/{idOrKey}/data-collection-specs | 
[**GetIntegrationDataLocationByKey**](IntegrationsAPI.md#GetIntegrationDataLocationByKey) | **Get** /integrations/{idOrKey}/data/{dataLocationKey} | 
[**GetIntegrationDataLocations**](IntegrationsAPI.md#GetIntegrationDataLocations) | **Get** /integrations/{idOrKey}/data | 
[**GetIntegrationDocumentation**](IntegrationsAPI.md#GetIntegrationDocumentation) | **Get** /integrations/{idOrKey}/documentation | 
[**GetIntegrationEventByKey**](IntegrationsAPI.md#GetIntegrationEventByKey) | **Get** /integrations/{idOrKey}/events/{eventKey} | 
[**GetIntegrationEvents**](IntegrationsAPI.md#GetIntegrationEvents) | **Get** /integrations/{idOrKey}/events | 
[**GetIntegrationFlowNodes**](IntegrationsAPI.md#GetIntegrationFlowNodes) | **Get** /integrations/{idOrKey}/flow-nodes | 
[**GetIntegrationOpenapiOperation**](IntegrationsAPI.md#GetIntegrationOpenapiOperation) | **Get** /integrations/{idOrKey}/openapi/request-schema | 
[**GetIntegrationOpenapiPathMethods**](IntegrationsAPI.md#GetIntegrationOpenapiPathMethods) | **Get** /integrations/{idOrKey}/openapi/path-methods | 
[**GetIntegrationOpenapiPaths**](IntegrationsAPI.md#GetIntegrationOpenapiPaths) | **Get** /integrations/{idOrKey}/openapi/paths | 
[**GetIntegrationOpenapiRequests**](IntegrationsAPI.md#GetIntegrationOpenapiRequests) | **Get** /integrations/{idOrKey}/openapi/request-specs | 
[**GetIntegrationOpenapiResponseSchema**](IntegrationsAPI.md#GetIntegrationOpenapiResponseSchema) | **Get** /integrations/{idOrKey}/openapi/response-schema | 
[**GetIntegrationOperationByKey**](IntegrationsAPI.md#GetIntegrationOperationByKey) | **Get** /integrations/{idOrKey}/operations/{operationKey} | 
[**GetIntegrationOperations**](IntegrationsAPI.md#GetIntegrationOperations) | **Get** /integrations/{idOrKey}/operations | 
[**GetIntegrationUdmByKey**](IntegrationsAPI.md#GetIntegrationUdmByKey) | **Get** /integrations/{idOrKey}/udm/{udmKey} | 
[**GetIntegrationUdms**](IntegrationsAPI.md#GetIntegrationUdms) | **Get** /integrations/{idOrKey}/udm | 
[**ListIntegrations**](IntegrationsAPI.md#ListIntegrations) | **Get** /integrations | 
[**PatchIntegration**](IntegrationsAPI.md#PatchIntegration) | **Patch** /integrations/{idOrKey} | 
[**PublishConnectorToIntegration**](IntegrationsAPI.md#PublishConnectorToIntegration) | **Post** /integrations/{idOrKey}/publish-connector | 
[**PutIntegration**](IntegrationsAPI.md#PutIntegration) | **Put** /integrations/{idOrKey} | 
[**ReUploadIntegration**](IntegrationsAPI.md#ReUploadIntegration) | **Post** /integrations/{id} | 
[**ResetParameters**](IntegrationsAPI.md#ResetParameters) | **Post** /integrations/{idOrKey}/reset-parameters | 
[**SetupIntegration**](IntegrationsAPI.md#SetupIntegration) | **Post** /integrations/{idOrKey}/setup | 
[**UpdateIntegrationConnector**](IntegrationsAPI.md#UpdateIntegrationConnector) | **Post** /integrations/{idOrKey}/update-connector | 



## ArchiveIntegration

> ArchiveIntegration(ctx, isOrKey).Execute()



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
	isOrKey := "isOrKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IntegrationsAPI.ArchiveIntegration(context.Background(), isOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.ArchiveIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**isOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiArchiveIntegrationRequest struct via the builder pattern


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


## CreateIntegration

> IntegrationDto CreateIntegration(ctx).CreateIntegrationDto(createIntegrationDto).Execute()



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
	createIntegrationDto := *openapiclient.NewCreateIntegrationDto() // CreateIntegrationDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.CreateIntegration(context.Background()).CreateIntegrationDto(createIntegrationDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.CreateIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIntegration`: IntegrationDto
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.CreateIntegration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createIntegrationDto** | [**CreateIntegrationDto**](CreateIntegrationDto.md) |  | 

### Return type

[**IntegrationDto**](IntegrationDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisconnectIntegration

> DisconnectIntegration(ctx, idOrKey).Execute()



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
	idOrKey := "idOrKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IntegrationsAPI.DisconnectIntegration(context.Background(), idOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.DisconnectIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisconnectIntegrationRequest struct via the builder pattern


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


## GetConnectorSpec

> map[string]interface{} GetConnectorSpec(ctx, idOrKey).Execute()



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
	idOrKey := "idOrKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.GetConnectorSpec(context.Background(), idOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.GetConnectorSpec``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectorSpec`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.GetConnectorSpec`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectorSpecRequest struct via the builder pattern


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


## GetIntegration

> IntegrationDto GetIntegration(ctx, idOrKey).Execute()



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
	idOrKey := "idOrKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.GetIntegration(context.Background(), idOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.GetIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIntegration`: IntegrationDto
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.GetIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IntegrationDto**](IntegrationDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntegrationDataCollectionSpecs

> []map[string]interface{} GetIntegrationDataCollectionSpecs(ctx, idOrKey).Execute()



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
	idOrKey := "idOrKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.GetIntegrationDataCollectionSpecs(context.Background(), idOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.GetIntegrationDataCollectionSpecs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIntegrationDataCollectionSpecs`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.GetIntegrationDataCollectionSpecs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationDataCollectionSpecsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntegrationDataLocationByKey

> map[string]interface{} GetIntegrationDataLocationByKey(ctx, idOrKey, dataLocationKey).Execute()



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
	idOrKey := "idOrKey_example" // string | 
	dataLocationKey := "dataLocationKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.GetIntegrationDataLocationByKey(context.Background(), idOrKey, dataLocationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.GetIntegrationDataLocationByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIntegrationDataLocationByKey`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.GetIntegrationDataLocationByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrKey** | **string** |  | 
**dataLocationKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationDataLocationByKeyRequest struct via the builder pattern


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


## GetIntegrationDataLocations

> []map[string]interface{} GetIntegrationDataLocations(ctx, idOrKey).Execute()



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
	idOrKey := "idOrKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.GetIntegrationDataLocations(context.Background(), idOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.GetIntegrationDataLocations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIntegrationDataLocations`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.GetIntegrationDataLocations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationDataLocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntegrationDocumentation

> map[string]interface{} GetIntegrationDocumentation(ctx, idOrKey).Execute()



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
	idOrKey := "idOrKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.GetIntegrationDocumentation(context.Background(), idOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.GetIntegrationDocumentation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIntegrationDocumentation`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.GetIntegrationDocumentation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationDocumentationRequest struct via the builder pattern


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


## GetIntegrationEventByKey

> map[string]interface{} GetIntegrationEventByKey(ctx, idOrKey, eventKey).Execute()



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
	idOrKey := "idOrKey_example" // string | 
	eventKey := "eventKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.GetIntegrationEventByKey(context.Background(), idOrKey, eventKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.GetIntegrationEventByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIntegrationEventByKey`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.GetIntegrationEventByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrKey** | **string** |  | 
**eventKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationEventByKeyRequest struct via the builder pattern


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


## GetIntegrationEvents

> []map[string]interface{} GetIntegrationEvents(ctx, idOrKey).Execute()



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
	idOrKey := "idOrKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.GetIntegrationEvents(context.Background(), idOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.GetIntegrationEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIntegrationEvents`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.GetIntegrationEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntegrationFlowNodes

> []map[string]interface{} GetIntegrationFlowNodes(ctx, idOrKey).Execute()



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
	idOrKey := "idOrKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.GetIntegrationFlowNodes(context.Background(), idOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.GetIntegrationFlowNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIntegrationFlowNodes`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.GetIntegrationFlowNodes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationFlowNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntegrationOpenapiOperation

> map[string]interface{} GetIntegrationOpenapiOperation(ctx, idOrKey).Path(path).Method(method).Execute()



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
	idOrKey := "idOrKey_example" // string | 
	path := "path_example" // string | 
	method := "method_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.GetIntegrationOpenapiOperation(context.Background(), idOrKey).Path(path).Method(method).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.GetIntegrationOpenapiOperation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIntegrationOpenapiOperation`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.GetIntegrationOpenapiOperation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationOpenapiOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **path** | **string** |  | 
 **method** | **string** |  | 

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


## GetIntegrationOpenapiPathMethods

> []string GetIntegrationOpenapiPathMethods(ctx, idOrKey).Path(path).Execute()



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
	idOrKey := "idOrKey_example" // string | 
	path := "path_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.GetIntegrationOpenapiPathMethods(context.Background(), idOrKey).Path(path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.GetIntegrationOpenapiPathMethods``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIntegrationOpenapiPathMethods`: []string
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.GetIntegrationOpenapiPathMethods`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationOpenapiPathMethodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **path** | **string** |  | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntegrationOpenapiPaths

> []string GetIntegrationOpenapiPaths(ctx, idOrKey).Execute()



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
	idOrKey := "idOrKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.GetIntegrationOpenapiPaths(context.Background(), idOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.GetIntegrationOpenapiPaths``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIntegrationOpenapiPaths`: []string
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.GetIntegrationOpenapiPaths`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationOpenapiPathsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntegrationOpenapiRequests

> []map[string]interface{} GetIntegrationOpenapiRequests(ctx, idOrKey).Execute()



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
	idOrKey := "idOrKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.GetIntegrationOpenapiRequests(context.Background(), idOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.GetIntegrationOpenapiRequests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIntegrationOpenapiRequests`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.GetIntegrationOpenapiRequests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationOpenapiRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntegrationOpenapiResponseSchema

> map[string]interface{} GetIntegrationOpenapiResponseSchema(ctx, idOrKey).Path(path).Method(method).Execute()



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
	idOrKey := "idOrKey_example" // string | 
	path := "path_example" // string | 
	method := "method_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.GetIntegrationOpenapiResponseSchema(context.Background(), idOrKey).Path(path).Method(method).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.GetIntegrationOpenapiResponseSchema``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIntegrationOpenapiResponseSchema`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.GetIntegrationOpenapiResponseSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationOpenapiResponseSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **path** | **string** |  | 
 **method** | **string** |  | 

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


## GetIntegrationOperationByKey

> map[string]interface{} GetIntegrationOperationByKey(ctx, idOrKey, operationKey).Execute()



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
	idOrKey := "idOrKey_example" // string | 
	operationKey := "operationKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.GetIntegrationOperationByKey(context.Background(), idOrKey, operationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.GetIntegrationOperationByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIntegrationOperationByKey`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.GetIntegrationOperationByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrKey** | **string** |  | 
**operationKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationOperationByKeyRequest struct via the builder pattern


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


## GetIntegrationOperations

> []map[string]interface{} GetIntegrationOperations(ctx, idOrKey).Execute()



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
	idOrKey := "idOrKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.GetIntegrationOperations(context.Background(), idOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.GetIntegrationOperations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIntegrationOperations`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.GetIntegrationOperations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntegrationUdmByKey

> map[string]interface{} GetIntegrationUdmByKey(ctx, idOrKey, udmKey).Execute()



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
	idOrKey := "idOrKey_example" // string | 
	udmKey := "udmKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.GetIntegrationUdmByKey(context.Background(), idOrKey, udmKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.GetIntegrationUdmByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIntegrationUdmByKey`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.GetIntegrationUdmByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrKey** | **string** |  | 
**udmKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationUdmByKeyRequest struct via the builder pattern


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


## GetIntegrationUdms

> []map[string]interface{} GetIntegrationUdms(ctx, idOrKey).Execute()



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
	idOrKey := "idOrKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.GetIntegrationUdms(context.Background(), idOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.GetIntegrationUdms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIntegrationUdms`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.GetIntegrationUdms`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationUdmsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIntegrations

> ListIntegrations200Response ListIntegrations(ctx).Execute()



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
	resp, r, err := apiClient.IntegrationsAPI.ListIntegrations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.ListIntegrations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIntegrations`: ListIntegrations200Response
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.ListIntegrations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListIntegrationsRequest struct via the builder pattern


### Return type

[**ListIntegrations200Response**](ListIntegrations200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchIntegration

> IntegrationDto PatchIntegration(ctx, idOrKey).UpdateIntegrationDto(updateIntegrationDto).Execute()



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
	idOrKey := "idOrKey_example" // string | 
	updateIntegrationDto := *openapiclient.NewUpdateIntegrationDto() // UpdateIntegrationDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.PatchIntegration(context.Background(), idOrKey).UpdateIntegrationDto(updateIntegrationDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.PatchIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchIntegration`: IntegrationDto
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.PatchIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateIntegrationDto** | [**UpdateIntegrationDto**](UpdateIntegrationDto.md) |  | 

### Return type

[**IntegrationDto**](IntegrationDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublishConnectorToIntegration

> IntegrationDto PublishConnectorToIntegration(ctx, idOrKey).PublishConnectorToIntegrationDto(publishConnectorToIntegrationDto).Execute()



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
	idOrKey := "idOrKey_example" // string | 
	publishConnectorToIntegrationDto := *openapiclient.NewPublishConnectorToIntegrationDto("ConnectorId_example") // PublishConnectorToIntegrationDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.PublishConnectorToIntegration(context.Background(), idOrKey).PublishConnectorToIntegrationDto(publishConnectorToIntegrationDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.PublishConnectorToIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PublishConnectorToIntegration`: IntegrationDto
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.PublishConnectorToIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublishConnectorToIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **publishConnectorToIntegrationDto** | [**PublishConnectorToIntegrationDto**](PublishConnectorToIntegrationDto.md) |  | 

### Return type

[**IntegrationDto**](IntegrationDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutIntegration

> IntegrationDto PutIntegration(ctx, idOrKey).CreateIntegrationDto(createIntegrationDto).Execute()



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
	idOrKey := "idOrKey_example" // string | 
	createIntegrationDto := *openapiclient.NewCreateIntegrationDto() // CreateIntegrationDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.PutIntegration(context.Background(), idOrKey).CreateIntegrationDto(createIntegrationDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.PutIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutIntegration`: IntegrationDto
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.PutIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createIntegrationDto** | [**CreateIntegrationDto**](CreateIntegrationDto.md) |  | 

### Return type

[**IntegrationDto**](IntegrationDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReUploadIntegration

> IntegrationDto ReUploadIntegration(ctx, id).Execute()



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
	resp, r, err := apiClient.IntegrationsAPI.ReUploadIntegration(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.ReUploadIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReUploadIntegration`: IntegrationDto
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.ReUploadIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReUploadIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IntegrationDto**](IntegrationDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetParameters

> IntegrationDto ResetParameters(ctx, idOrKey).Execute()



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
	idOrKey := "idOrKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.ResetParameters(context.Background(), idOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.ResetParameters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResetParameters`: IntegrationDto
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.ResetParameters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetParametersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IntegrationDto**](IntegrationDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetupIntegration

> IntegrationDto SetupIntegration(ctx, idOrKey).Execute()



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
	idOrKey := "idOrKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.SetupIntegration(context.Background(), idOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.SetupIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetupIntegration`: IntegrationDto
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.SetupIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetupIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IntegrationDto**](IntegrationDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIntegrationConnector

> IntegrationDto UpdateIntegrationConnector(ctx, idOrKey).Execute()



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
	idOrKey := "idOrKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.UpdateIntegrationConnector(context.Background(), idOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.UpdateIntegrationConnector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateIntegrationConnector`: IntegrationDto
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.UpdateIntegrationConnector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIntegrationConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IntegrationDto**](IntegrationDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

