# \ConnectionsAPI

All URIs are relative to *https://api.integration.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConnectionsControllerArchiveConnection**](ConnectionsAPI.md#ConnectionsControllerArchiveConnection) | **Delete** /connections/{connectionIdOrKey} | Archive connection
[**ConnectionsControllerCreateConnection**](ConnectionsAPI.md#ConnectionsControllerCreateConnection) | **Post** /connections | Create connection
[**ConnectionsControllerDataLocationMethod**](ConnectionsAPI.md#ConnectionsControllerDataLocationMethod) | **Post** /connections/{connectionIdOrKey}/data/{dataLocationKey}/{methodKey} | Get connection data location method
[**ConnectionsControllerGetConnection**](ConnectionsAPI.md#ConnectionsControllerGetConnection) | **Get** /connections/{connectionIdOrKey} | Get connection
[**ConnectionsControllerGetDataLocationSpec**](ConnectionsAPI.md#ConnectionsControllerGetDataLocationSpec) | **Get** /connections/{connectionIdOrKey}/data/{dataLocationKey} | Get connection data location
[**ConnectionsControllerGetOperation**](ConnectionsAPI.md#ConnectionsControllerGetOperation) | **Get** /connections/{connectionIdOrKey}/operations/{operationKey} | Get connection operation
[**ConnectionsControllerListConnections**](ConnectionsAPI.md#ConnectionsControllerListConnections) | **Get** /connections | List connections
[**ConnectionsControllerListDataLocations**](ConnectionsAPI.md#ConnectionsControllerListDataLocations) | **Get** /connections/{connectionIdOrKey}/data | List connection data locations
[**ConnectionsControllerListOperations**](ConnectionsAPI.md#ConnectionsControllerListOperations) | **Get** /connections/{connectionIdOrKey}/operations | List connection operations
[**ConnectionsControllerPatchConnection**](ConnectionsAPI.md#ConnectionsControllerPatchConnection) | **Patch** /connections/{connectionIdOrKey} | Update connection
[**ConnectionsControllerProxyRequestDelete**](ConnectionsAPI.md#ConnectionsControllerProxyRequestDelete) | **Delete** /connections/{connectionIdOrKey}/proxy/{path} | Connection proxy request
[**ConnectionsControllerProxyRequestGet**](ConnectionsAPI.md#ConnectionsControllerProxyRequestGet) | **Get** /connections/{connectionIdOrKey}/proxy/{path} | Connection proxy request
[**ConnectionsControllerProxyRequestHead**](ConnectionsAPI.md#ConnectionsControllerProxyRequestHead) | **Head** /connections/{connectionIdOrKey}/proxy/{path} | Connection proxy request
[**ConnectionsControllerProxyRequestOptions**](ConnectionsAPI.md#ConnectionsControllerProxyRequestOptions) | **Options** /connections/{connectionIdOrKey}/proxy/{path} | Connection proxy request
[**ConnectionsControllerProxyRequestPatch**](ConnectionsAPI.md#ConnectionsControllerProxyRequestPatch) | **Patch** /connections/{connectionIdOrKey}/proxy/{path} | Connection proxy request
[**ConnectionsControllerProxyRequestPost**](ConnectionsAPI.md#ConnectionsControllerProxyRequestPost) | **Post** /connections/{connectionIdOrKey}/proxy/{path} | Connection proxy request
[**ConnectionsControllerProxyRequestPut**](ConnectionsAPI.md#ConnectionsControllerProxyRequestPut) | **Put** /connections/{connectionIdOrKey}/proxy/{path} | Connection proxy request
[**ConnectionsControllerRefreshConnectionCredentials**](ConnectionsAPI.md#ConnectionsControllerRefreshConnectionCredentials) | **Post** /connections/{connectionIdOrKey}/refresh-credentials | Refresh connection credentials
[**ConnectionsControllerRunOperation**](ConnectionsAPI.md#ConnectionsControllerRunOperation) | **Post** /connections/{connectionIdOrKey}/operations/{operationKey}/run | Run connection operation
[**ConnectionsControllerTestConnection**](ConnectionsAPI.md#ConnectionsControllerTestConnection) | **Post** /connections/{connectionIdOrKey}/test | Test connection



## ConnectionsControllerArchiveConnection

> ConnectionsControllerArchiveConnection(ctx, connectionIdOrKey).Execute()

Archive connection

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConnectionsAPI.ConnectionsControllerArchiveConnection(context.Background(), connectionIdOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.ConnectionsControllerArchiveConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionIdOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionsControllerArchiveConnectionRequest struct via the builder pattern


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


## ConnectionsControllerCreateConnection

> ConnectionDto ConnectionsControllerCreateConnection(ctx).CreateConnectionDto(createConnectionDto).Execute()

Create connection

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
	createConnectionDto := *openapiclient.NewCreateConnectionDto() // CreateConnectionDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsAPI.ConnectionsControllerCreateConnection(context.Background()).CreateConnectionDto(createConnectionDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.ConnectionsControllerCreateConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectionsControllerCreateConnection`: ConnectionDto
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.ConnectionsControllerCreateConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConnectionsControllerCreateConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createConnectionDto** | [**CreateConnectionDto**](CreateConnectionDto.md) |  | 

### Return type

[**ConnectionDto**](ConnectionDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionsControllerDataLocationMethod

> map[string]interface{} ConnectionsControllerDataLocationMethod(ctx, connectionIdOrKey, dataLocationKey, methodKey).Execute()

Get connection data location method

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
	dataLocationKey := "dataLocationKey_example" // string | 
	methodKey := "methodKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsAPI.ConnectionsControllerDataLocationMethod(context.Background(), connectionIdOrKey, dataLocationKey, methodKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.ConnectionsControllerDataLocationMethod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectionsControllerDataLocationMethod`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.ConnectionsControllerDataLocationMethod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionIdOrKey** | **string** |  | 
**dataLocationKey** | **string** |  | 
**methodKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionsControllerDataLocationMethodRequest struct via the builder pattern


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


## ConnectionsControllerGetConnection

> ConnectionDto ConnectionsControllerGetConnection(ctx, connectionIdOrKey).Execute()

Get connection

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsAPI.ConnectionsControllerGetConnection(context.Background(), connectionIdOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.ConnectionsControllerGetConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectionsControllerGetConnection`: ConnectionDto
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.ConnectionsControllerGetConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionIdOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionsControllerGetConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConnectionDto**](ConnectionDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionsControllerGetDataLocationSpec

> map[string]interface{} ConnectionsControllerGetDataLocationSpec(ctx, connectionIdOrKey, dataLocationKey).Execute()

Get connection data location

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
	dataLocationKey := "dataLocationKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsAPI.ConnectionsControllerGetDataLocationSpec(context.Background(), connectionIdOrKey, dataLocationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.ConnectionsControllerGetDataLocationSpec``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectionsControllerGetDataLocationSpec`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.ConnectionsControllerGetDataLocationSpec`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionIdOrKey** | **string** |  | 
**dataLocationKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionsControllerGetDataLocationSpecRequest struct via the builder pattern


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


## ConnectionsControllerGetOperation

> map[string]interface{} ConnectionsControllerGetOperation(ctx, connectionIdOrKey, operationKey).Execute()

Get connection operation

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
	operationKey := "operationKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsAPI.ConnectionsControllerGetOperation(context.Background(), connectionIdOrKey, operationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.ConnectionsControllerGetOperation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectionsControllerGetOperation`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.ConnectionsControllerGetOperation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionIdOrKey** | **string** |  | 
**operationKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionsControllerGetOperationRequest struct via the builder pattern


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


## ConnectionsControllerListConnections

> ConnectionsControllerListConnections200Response ConnectionsControllerListConnections(ctx).Limit(limit).Cursor(cursor).Search(search).ConnectorId(connectorId).IncludeArchived(includeArchived).Id(id).UserId(userId).Name(name).IsTest(isTest).Disconnected(disconnected).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()

List connections

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
	userId := "userId_example" // string |  (optional)
	name := "name_example" // string |  (optional)
	isTest := true // bool |  (optional)
	disconnected := true // bool |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsAPI.ConnectionsControllerListConnections(context.Background()).Limit(limit).Cursor(cursor).Search(search).ConnectorId(connectorId).IncludeArchived(includeArchived).Id(id).UserId(userId).Name(name).IsTest(isTest).Disconnected(disconnected).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.ConnectionsControllerListConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectionsControllerListConnections`: ConnectionsControllerListConnections200Response
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.ConnectionsControllerListConnections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConnectionsControllerListConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **float32** |  | 
 **cursor** | **string** |  | 
 **search** | **string** |  | 
 **connectorId** | **string** |  | 
 **includeArchived** | **bool** |  | 
 **id** | **string** |  | 
 **userId** | **string** |  | 
 **name** | **string** |  | 
 **isTest** | **bool** |  | 
 **disconnected** | **bool** |  | 
 **integrationId** | **string** |  | 
 **integrationKey** | **string** |  | 

### Return type

[**ConnectionsControllerListConnections200Response**](ConnectionsControllerListConnections200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionsControllerListDataLocations

> []map[string]interface{} ConnectionsControllerListDataLocations(ctx, connectionIdOrKey).Execute()

List connection data locations

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsAPI.ConnectionsControllerListDataLocations(context.Background(), connectionIdOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.ConnectionsControllerListDataLocations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectionsControllerListDataLocations`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.ConnectionsControllerListDataLocations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionIdOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionsControllerListDataLocationsRequest struct via the builder pattern


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


## ConnectionsControllerListOperations

> []map[string]interface{} ConnectionsControllerListOperations(ctx, connectionIdOrKey).Execute()

List connection operations

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsAPI.ConnectionsControllerListOperations(context.Background(), connectionIdOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.ConnectionsControllerListOperations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectionsControllerListOperations`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.ConnectionsControllerListOperations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionIdOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionsControllerListOperationsRequest struct via the builder pattern


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


## ConnectionsControllerPatchConnection

> ConnectionDto ConnectionsControllerPatchConnection(ctx, connectionIdOrKey).UpdateConnectionDto(updateConnectionDto).Execute()

Update connection

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
	updateConnectionDto := *openapiclient.NewUpdateConnectionDto() // UpdateConnectionDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsAPI.ConnectionsControllerPatchConnection(context.Background(), connectionIdOrKey).UpdateConnectionDto(updateConnectionDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.ConnectionsControllerPatchConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectionsControllerPatchConnection`: ConnectionDto
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.ConnectionsControllerPatchConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionIdOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionsControllerPatchConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateConnectionDto** | [**UpdateConnectionDto**](UpdateConnectionDto.md) |  | 

### Return type

[**ConnectionDto**](ConnectionDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionsControllerProxyRequestDelete

> ConnectionsControllerProxyRequestDelete(ctx, connectionIdOrKey, path).Execute()

Connection proxy request

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
	path := "path_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConnectionsAPI.ConnectionsControllerProxyRequestDelete(context.Background(), connectionIdOrKey, path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.ConnectionsControllerProxyRequestDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionIdOrKey** | **string** |  | 
**path** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionsControllerProxyRequestDeleteRequest struct via the builder pattern


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


## ConnectionsControllerProxyRequestGet

> ConnectionsControllerProxyRequestGet(ctx, connectionIdOrKey, path).Execute()

Connection proxy request

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
	path := "path_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConnectionsAPI.ConnectionsControllerProxyRequestGet(context.Background(), connectionIdOrKey, path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.ConnectionsControllerProxyRequestGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionIdOrKey** | **string** |  | 
**path** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionsControllerProxyRequestGetRequest struct via the builder pattern


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


## ConnectionsControllerProxyRequestHead

> ConnectionsControllerProxyRequestHead(ctx, connectionIdOrKey, path).Execute()

Connection proxy request

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
	path := "path_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConnectionsAPI.ConnectionsControllerProxyRequestHead(context.Background(), connectionIdOrKey, path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.ConnectionsControllerProxyRequestHead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionIdOrKey** | **string** |  | 
**path** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionsControllerProxyRequestHeadRequest struct via the builder pattern


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


## ConnectionsControllerProxyRequestOptions

> ConnectionsControllerProxyRequestOptions(ctx, connectionIdOrKey, path).Execute()

Connection proxy request

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
	path := "path_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConnectionsAPI.ConnectionsControllerProxyRequestOptions(context.Background(), connectionIdOrKey, path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.ConnectionsControllerProxyRequestOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionIdOrKey** | **string** |  | 
**path** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionsControllerProxyRequestOptionsRequest struct via the builder pattern


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


## ConnectionsControllerProxyRequestPatch

> ConnectionsControllerProxyRequestPatch(ctx, connectionIdOrKey, path).Execute()

Connection proxy request

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
	path := "path_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConnectionsAPI.ConnectionsControllerProxyRequestPatch(context.Background(), connectionIdOrKey, path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.ConnectionsControllerProxyRequestPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionIdOrKey** | **string** |  | 
**path** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionsControllerProxyRequestPatchRequest struct via the builder pattern


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


## ConnectionsControllerProxyRequestPost

> ConnectionsControllerProxyRequestPost(ctx, connectionIdOrKey, path).Execute()

Connection proxy request

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
	path := "path_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConnectionsAPI.ConnectionsControllerProxyRequestPost(context.Background(), connectionIdOrKey, path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.ConnectionsControllerProxyRequestPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionIdOrKey** | **string** |  | 
**path** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionsControllerProxyRequestPostRequest struct via the builder pattern


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


## ConnectionsControllerProxyRequestPut

> ConnectionsControllerProxyRequestPut(ctx, connectionIdOrKey, path).Execute()

Connection proxy request

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
	path := "path_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConnectionsAPI.ConnectionsControllerProxyRequestPut(context.Background(), connectionIdOrKey, path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.ConnectionsControllerProxyRequestPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionIdOrKey** | **string** |  | 
**path** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionsControllerProxyRequestPutRequest struct via the builder pattern


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


## ConnectionsControllerRefreshConnectionCredentials

> ConnectionsControllerRefreshConnectionCredentials(ctx, connectionIdOrKey).Execute()

Refresh connection credentials

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConnectionsAPI.ConnectionsControllerRefreshConnectionCredentials(context.Background(), connectionIdOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.ConnectionsControllerRefreshConnectionCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionIdOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionsControllerRefreshConnectionCredentialsRequest struct via the builder pattern


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


## ConnectionsControllerRunOperation

> OperationRunResponseDto ConnectionsControllerRunOperation(ctx, connectionIdOrKey, operationKey).Execute()

Run connection operation

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
	operationKey := "operationKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsAPI.ConnectionsControllerRunOperation(context.Background(), connectionIdOrKey, operationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.ConnectionsControllerRunOperation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectionsControllerRunOperation`: OperationRunResponseDto
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.ConnectionsControllerRunOperation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionIdOrKey** | **string** |  | 
**operationKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionsControllerRunOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationRunResponseDto**](OperationRunResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionsControllerTestConnection

> ConnectionsControllerTestConnection(ctx, connectionIdOrKey).Execute()

Test connection

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConnectionsAPI.ConnectionsControllerTestConnection(context.Background(), connectionIdOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.ConnectionsControllerTestConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionIdOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionsControllerTestConnectionRequest struct via the builder pattern


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

