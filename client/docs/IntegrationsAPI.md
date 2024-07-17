# \IntegrationsAPI

All URIs are relative to *https://api.integration.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IntegrationsControllerArchiveIntegration**](IntegrationsAPI.md#IntegrationsControllerArchiveIntegration) | **Delete** /integrations/{idOrKey} | Archive integration
[**IntegrationsControllerCreateIntegration**](IntegrationsAPI.md#IntegrationsControllerCreateIntegration) | **Post** /integrations | Create integration
[**IntegrationsControllerGetConnectorSpec**](IntegrationsAPI.md#IntegrationsControllerGetConnectorSpec) | **Get** /integrations/{idOrKey}/connector-spec | Get integration connector spec
[**IntegrationsControllerGetIntegration**](IntegrationsAPI.md#IntegrationsControllerGetIntegration) | **Get** /integrations/{idOrKey} | Get integration
[**IntegrationsControllerGetIntegrationDataCollectionSpecs**](IntegrationsAPI.md#IntegrationsControllerGetIntegrationDataCollectionSpecs) | **Get** /integrations/{idOrKey}/data-collection-specs | List integration data collection specs
[**IntegrationsControllerGetIntegrationDataLocationByKey**](IntegrationsAPI.md#IntegrationsControllerGetIntegrationDataLocationByKey) | **Get** /integrations/{idOrKey}/data/{dataLocationKey} | Get integration data location
[**IntegrationsControllerGetIntegrationDataLocations**](IntegrationsAPI.md#IntegrationsControllerGetIntegrationDataLocations) | **Get** /integrations/{idOrKey}/data | List integration data locations
[**IntegrationsControllerGetIntegrationDocumentation**](IntegrationsAPI.md#IntegrationsControllerGetIntegrationDocumentation) | **Get** /integrations/{idOrKey}/documentation | Get integration connector documentation
[**IntegrationsControllerGetIntegrationEventByKey**](IntegrationsAPI.md#IntegrationsControllerGetIntegrationEventByKey) | **Get** /integrations/{idOrKey}/events/{eventKey} | Get integration event
[**IntegrationsControllerGetIntegrationEvents**](IntegrationsAPI.md#IntegrationsControllerGetIntegrationEvents) | **Get** /integrations/{idOrKey}/events | List integration events
[**IntegrationsControllerGetIntegrationGlobalWebhooks**](IntegrationsAPI.md#IntegrationsControllerGetIntegrationGlobalWebhooks) | **Get** /integrations/{idOrKey}/global-webhooks | List integration global webhooks
[**IntegrationsControllerGetIntegrationOperationByKey**](IntegrationsAPI.md#IntegrationsControllerGetIntegrationOperationByKey) | **Get** /integrations/{idOrKey}/operations/{operationKey} | Get integration operation
[**IntegrationsControllerGetIntegrationOperations**](IntegrationsAPI.md#IntegrationsControllerGetIntegrationOperations) | **Get** /integrations/{idOrKey}/operations | List integration operations
[**IntegrationsControllerGetIntegrationUdmByKey**](IntegrationsAPI.md#IntegrationsControllerGetIntegrationUdmByKey) | **Get** /integrations/{idOrKey}/udm/{udmKey} | Get integration UDM
[**IntegrationsControllerGetIntegrationUdms**](IntegrationsAPI.md#IntegrationsControllerGetIntegrationUdms) | **Get** /integrations/{idOrKey}/udm | List integration UDMs
[**IntegrationsControllerListIntegrations**](IntegrationsAPI.md#IntegrationsControllerListIntegrations) | **Get** /integrations | List integrations
[**IntegrationsControllerPatchIntegration**](IntegrationsAPI.md#IntegrationsControllerPatchIntegration) | **Patch** /integrations/{idOrKey} | Patch integration
[**IntegrationsControllerPutIntegration**](IntegrationsAPI.md#IntegrationsControllerPutIntegration) | **Put** /integrations/{idOrKey} | Update integration
[**IntegrationsControllerSetupIntegration**](IntegrationsAPI.md#IntegrationsControllerSetupIntegration) | **Post** /integrations/{idOrKey}/setup | Setup integration
[**IntegrationsControllerUploadConnector**](IntegrationsAPI.md#IntegrationsControllerUploadConnector) | **Post** /integrations/{idOrKey}/upload-connector | Upload integration connector



## IntegrationsControllerArchiveIntegration

> IntegrationsControllerArchiveIntegration(ctx, idOrKey).Execute()

Archive integration

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
	r, err := apiClient.IntegrationsAPI.IntegrationsControllerArchiveIntegration(context.Background(), idOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.IntegrationsControllerArchiveIntegration``: %v\n", err)
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

Other parameters are passed through a pointer to a apiIntegrationsControllerArchiveIntegrationRequest struct via the builder pattern


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


## IntegrationsControllerCreateIntegration

> IntegrationDto IntegrationsControllerCreateIntegration(ctx).CreateIntegrationDto(createIntegrationDto).Execute()

Create integration

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
	resp, r, err := apiClient.IntegrationsAPI.IntegrationsControllerCreateIntegration(context.Background()).CreateIntegrationDto(createIntegrationDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.IntegrationsControllerCreateIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationsControllerCreateIntegration`: IntegrationDto
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.IntegrationsControllerCreateIntegration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsControllerCreateIntegrationRequest struct via the builder pattern


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


## IntegrationsControllerGetConnectorSpec

> map[string]interface{} IntegrationsControllerGetConnectorSpec(ctx, idOrKey).Execute()

Get integration connector spec

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
	resp, r, err := apiClient.IntegrationsAPI.IntegrationsControllerGetConnectorSpec(context.Background(), idOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.IntegrationsControllerGetConnectorSpec``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationsControllerGetConnectorSpec`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.IntegrationsControllerGetConnectorSpec`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsControllerGetConnectorSpecRequest struct via the builder pattern


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


## IntegrationsControllerGetIntegration

> IntegrationDto IntegrationsControllerGetIntegration(ctx, idOrKey).Execute()

Get integration

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
	resp, r, err := apiClient.IntegrationsAPI.IntegrationsControllerGetIntegration(context.Background(), idOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.IntegrationsControllerGetIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationsControllerGetIntegration`: IntegrationDto
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.IntegrationsControllerGetIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsControllerGetIntegrationRequest struct via the builder pattern


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


## IntegrationsControllerGetIntegrationDataCollectionSpecs

> []map[string]interface{} IntegrationsControllerGetIntegrationDataCollectionSpecs(ctx, idOrKey).Execute()

List integration data collection specs

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
	resp, r, err := apiClient.IntegrationsAPI.IntegrationsControllerGetIntegrationDataCollectionSpecs(context.Background(), idOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.IntegrationsControllerGetIntegrationDataCollectionSpecs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationsControllerGetIntegrationDataCollectionSpecs`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.IntegrationsControllerGetIntegrationDataCollectionSpecs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsControllerGetIntegrationDataCollectionSpecsRequest struct via the builder pattern


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


## IntegrationsControllerGetIntegrationDataLocationByKey

> map[string]interface{} IntegrationsControllerGetIntegrationDataLocationByKey(ctx, idOrKey, dataLocationKey).Execute()

Get integration data location

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
	resp, r, err := apiClient.IntegrationsAPI.IntegrationsControllerGetIntegrationDataLocationByKey(context.Background(), idOrKey, dataLocationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.IntegrationsControllerGetIntegrationDataLocationByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationsControllerGetIntegrationDataLocationByKey`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.IntegrationsControllerGetIntegrationDataLocationByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrKey** | **string** |  | 
**dataLocationKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsControllerGetIntegrationDataLocationByKeyRequest struct via the builder pattern


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


## IntegrationsControllerGetIntegrationDataLocations

> []map[string]interface{} IntegrationsControllerGetIntegrationDataLocations(ctx, idOrKey).Execute()

List integration data locations

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
	resp, r, err := apiClient.IntegrationsAPI.IntegrationsControllerGetIntegrationDataLocations(context.Background(), idOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.IntegrationsControllerGetIntegrationDataLocations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationsControllerGetIntegrationDataLocations`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.IntegrationsControllerGetIntegrationDataLocations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsControllerGetIntegrationDataLocationsRequest struct via the builder pattern


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


## IntegrationsControllerGetIntegrationDocumentation

> map[string]interface{} IntegrationsControllerGetIntegrationDocumentation(ctx, idOrKey).Execute()

Get integration connector documentation

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
	resp, r, err := apiClient.IntegrationsAPI.IntegrationsControllerGetIntegrationDocumentation(context.Background(), idOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.IntegrationsControllerGetIntegrationDocumentation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationsControllerGetIntegrationDocumentation`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.IntegrationsControllerGetIntegrationDocumentation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsControllerGetIntegrationDocumentationRequest struct via the builder pattern


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


## IntegrationsControllerGetIntegrationEventByKey

> map[string]interface{} IntegrationsControllerGetIntegrationEventByKey(ctx, idOrKey, eventKey).Execute()

Get integration event

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
	resp, r, err := apiClient.IntegrationsAPI.IntegrationsControllerGetIntegrationEventByKey(context.Background(), idOrKey, eventKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.IntegrationsControllerGetIntegrationEventByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationsControllerGetIntegrationEventByKey`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.IntegrationsControllerGetIntegrationEventByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrKey** | **string** |  | 
**eventKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsControllerGetIntegrationEventByKeyRequest struct via the builder pattern


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


## IntegrationsControllerGetIntegrationEvents

> []map[string]interface{} IntegrationsControllerGetIntegrationEvents(ctx, idOrKey).Execute()

List integration events

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
	resp, r, err := apiClient.IntegrationsAPI.IntegrationsControllerGetIntegrationEvents(context.Background(), idOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.IntegrationsControllerGetIntegrationEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationsControllerGetIntegrationEvents`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.IntegrationsControllerGetIntegrationEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsControllerGetIntegrationEventsRequest struct via the builder pattern


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


## IntegrationsControllerGetIntegrationGlobalWebhooks

> []map[string]interface{} IntegrationsControllerGetIntegrationGlobalWebhooks(ctx, idOrKey).Execute()

List integration global webhooks

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
	resp, r, err := apiClient.IntegrationsAPI.IntegrationsControllerGetIntegrationGlobalWebhooks(context.Background(), idOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.IntegrationsControllerGetIntegrationGlobalWebhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationsControllerGetIntegrationGlobalWebhooks`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.IntegrationsControllerGetIntegrationGlobalWebhooks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsControllerGetIntegrationGlobalWebhooksRequest struct via the builder pattern


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


## IntegrationsControllerGetIntegrationOperationByKey

> map[string]interface{} IntegrationsControllerGetIntegrationOperationByKey(ctx, idOrKey, operationKey).Execute()

Get integration operation

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
	resp, r, err := apiClient.IntegrationsAPI.IntegrationsControllerGetIntegrationOperationByKey(context.Background(), idOrKey, operationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.IntegrationsControllerGetIntegrationOperationByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationsControllerGetIntegrationOperationByKey`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.IntegrationsControllerGetIntegrationOperationByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrKey** | **string** |  | 
**operationKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsControllerGetIntegrationOperationByKeyRequest struct via the builder pattern


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


## IntegrationsControllerGetIntegrationOperations

> []map[string]interface{} IntegrationsControllerGetIntegrationOperations(ctx, idOrKey).Execute()

List integration operations

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
	resp, r, err := apiClient.IntegrationsAPI.IntegrationsControllerGetIntegrationOperations(context.Background(), idOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.IntegrationsControllerGetIntegrationOperations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationsControllerGetIntegrationOperations`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.IntegrationsControllerGetIntegrationOperations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsControllerGetIntegrationOperationsRequest struct via the builder pattern


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


## IntegrationsControllerGetIntegrationUdmByKey

> map[string]interface{} IntegrationsControllerGetIntegrationUdmByKey(ctx, idOrKey, udmKey).Execute()

Get integration UDM

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
	resp, r, err := apiClient.IntegrationsAPI.IntegrationsControllerGetIntegrationUdmByKey(context.Background(), idOrKey, udmKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.IntegrationsControllerGetIntegrationUdmByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationsControllerGetIntegrationUdmByKey`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.IntegrationsControllerGetIntegrationUdmByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrKey** | **string** |  | 
**udmKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsControllerGetIntegrationUdmByKeyRequest struct via the builder pattern


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


## IntegrationsControllerGetIntegrationUdms

> []map[string]interface{} IntegrationsControllerGetIntegrationUdms(ctx, idOrKey).Execute()

List integration UDMs

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
	resp, r, err := apiClient.IntegrationsAPI.IntegrationsControllerGetIntegrationUdms(context.Background(), idOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.IntegrationsControllerGetIntegrationUdms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationsControllerGetIntegrationUdms`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.IntegrationsControllerGetIntegrationUdms`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsControllerGetIntegrationUdmsRequest struct via the builder pattern


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


## IntegrationsControllerListIntegrations

> IntegrationsControllerListIntegrations200Response IntegrationsControllerListIntegrations(ctx).Limit(limit).Cursor(cursor).Search(search).ConnectorId(connectorId).IncludeArchived(includeArchived).Execute()

List integrations

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.IntegrationsControllerListIntegrations(context.Background()).Limit(limit).Cursor(cursor).Search(search).ConnectorId(connectorId).IncludeArchived(includeArchived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.IntegrationsControllerListIntegrations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationsControllerListIntegrations`: IntegrationsControllerListIntegrations200Response
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.IntegrationsControllerListIntegrations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsControllerListIntegrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **float32** |  | 
 **cursor** | **string** |  | 
 **search** | **string** |  | 
 **connectorId** | **string** |  | 
 **includeArchived** | **bool** |  | 

### Return type

[**IntegrationsControllerListIntegrations200Response**](IntegrationsControllerListIntegrations200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsControllerPatchIntegration

> IntegrationDto IntegrationsControllerPatchIntegration(ctx, idOrKey).UpdateIntegrationDto(updateIntegrationDto).Execute()

Patch integration

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
	resp, r, err := apiClient.IntegrationsAPI.IntegrationsControllerPatchIntegration(context.Background(), idOrKey).UpdateIntegrationDto(updateIntegrationDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.IntegrationsControllerPatchIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationsControllerPatchIntegration`: IntegrationDto
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.IntegrationsControllerPatchIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsControllerPatchIntegrationRequest struct via the builder pattern


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


## IntegrationsControllerPutIntegration

> IntegrationDto IntegrationsControllerPutIntegration(ctx, idOrKey).CreateIntegrationDto(createIntegrationDto).Execute()

Update integration

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
	resp, r, err := apiClient.IntegrationsAPI.IntegrationsControllerPutIntegration(context.Background(), idOrKey).CreateIntegrationDto(createIntegrationDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.IntegrationsControllerPutIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationsControllerPutIntegration`: IntegrationDto
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.IntegrationsControllerPutIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsControllerPutIntegrationRequest struct via the builder pattern


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


## IntegrationsControllerSetupIntegration

> IntegrationDto IntegrationsControllerSetupIntegration(ctx, idOrKey).Execute()

Setup integration

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
	resp, r, err := apiClient.IntegrationsAPI.IntegrationsControllerSetupIntegration(context.Background(), idOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.IntegrationsControllerSetupIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationsControllerSetupIntegration`: IntegrationDto
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.IntegrationsControllerSetupIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsControllerSetupIntegrationRequest struct via the builder pattern


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


## IntegrationsControllerUploadConnector

> IntegrationDto IntegrationsControllerUploadConnector(ctx, idOrKey).Execute()

Upload integration connector

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
	resp, r, err := apiClient.IntegrationsAPI.IntegrationsControllerUploadConnector(context.Background(), idOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.IntegrationsControllerUploadConnector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationsControllerUploadConnector`: IntegrationDto
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.IntegrationsControllerUploadConnector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsControllerUploadConnectorRequest struct via the builder pattern


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

