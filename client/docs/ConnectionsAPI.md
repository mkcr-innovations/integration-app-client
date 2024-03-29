# \ConnectionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveConnection**](ConnectionsAPI.md#ArchiveConnection) | **Delete** /connections/{connectionIdOrKey} | 
[**CreateConnection**](ConnectionsAPI.md#CreateConnection) | **Post** /connections | 
[**DataLocationMethod**](ConnectionsAPI.md#DataLocationMethod) | **Post** /connections/{connectionIdOrKey}/data/{dataLocationKey}/{methodKey} | 
[**GetConnection**](ConnectionsAPI.md#GetConnection) | **Get** /connections/{connectionIdOrKey} | 
[**GetDataLocationSpec**](ConnectionsAPI.md#GetDataLocationSpec) | **Get** /connections/{connectionIdOrKey}/data/{dataLocationKey} | 
[**GetOperation**](ConnectionsAPI.md#GetOperation) | **Get** /connections/{connectionIdOrKey}/operations/{operationKey} | 
[**ListConnections**](ConnectionsAPI.md#ListConnections) | **Get** /connections | 
[**ListDataLocations**](ConnectionsAPI.md#ListDataLocations) | **Get** /connections/{connectionIdOrKey}/data | 
[**ListOperations**](ConnectionsAPI.md#ListOperations) | **Get** /connections/{connectionIdOrKey}/operations | 
[**MakeConnectionRequestDelete**](ConnectionsAPI.md#MakeConnectionRequestDelete) | **Delete** /connections/{connectionIdOrKey}/request | 
[**MakeConnectionRequestGet**](ConnectionsAPI.md#MakeConnectionRequestGet) | **Get** /connections/{connectionIdOrKey}/request | 
[**MakeConnectionRequestHead**](ConnectionsAPI.md#MakeConnectionRequestHead) | **Head** /connections/{connectionIdOrKey}/request | 
[**MakeConnectionRequestOptions**](ConnectionsAPI.md#MakeConnectionRequestOptions) | **Options** /connections/{connectionIdOrKey}/request | 
[**MakeConnectionRequestPatch**](ConnectionsAPI.md#MakeConnectionRequestPatch) | **Patch** /connections/{connectionIdOrKey}/request | 
[**MakeConnectionRequestPost**](ConnectionsAPI.md#MakeConnectionRequestPost) | **Post** /connections/{connectionIdOrKey}/request | 
[**MakeConnectionRequestPut**](ConnectionsAPI.md#MakeConnectionRequestPut) | **Put** /connections/{connectionIdOrKey}/request | 
[**PatchConnection**](ConnectionsAPI.md#PatchConnection) | **Patch** /connections/{connectionIdOrKey} | 
[**ProxyRequestDelete**](ConnectionsAPI.md#ProxyRequestDelete) | **Delete** /connections/{connectionIdOrKey}/proxy/{path} | 
[**ProxyRequestGet**](ConnectionsAPI.md#ProxyRequestGet) | **Get** /connections/{connectionIdOrKey}/proxy/{path} | 
[**ProxyRequestHead**](ConnectionsAPI.md#ProxyRequestHead) | **Head** /connections/{connectionIdOrKey}/proxy/{path} | 
[**ProxyRequestOptions**](ConnectionsAPI.md#ProxyRequestOptions) | **Options** /connections/{connectionIdOrKey}/proxy/{path} | 
[**ProxyRequestPatch**](ConnectionsAPI.md#ProxyRequestPatch) | **Patch** /connections/{connectionIdOrKey}/proxy/{path} | 
[**ProxyRequestPost**](ConnectionsAPI.md#ProxyRequestPost) | **Post** /connections/{connectionIdOrKey}/proxy/{path} | 
[**ProxyRequestPut**](ConnectionsAPI.md#ProxyRequestPut) | **Put** /connections/{connectionIdOrKey}/proxy/{path} | 
[**RefreshConnectionCredentials**](ConnectionsAPI.md#RefreshConnectionCredentials) | **Post** /connections/{connectionIdOrKey}/refresh-credentials | 
[**RunOperation**](ConnectionsAPI.md#RunOperation) | **Post** /connections/{connectionIdOrKey}/operations/{operationKey}/run | 
[**TestConnection**](ConnectionsAPI.md#TestConnection) | **Post** /connections/{connectionIdOrKey}/test | 



## ArchiveConnection

> ArchiveConnection(ctx, connectionIdOrKey).Execute()



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
	r, err := apiClient.ConnectionsAPI.ArchiveConnection(context.Background(), connectionIdOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.ArchiveConnection``: %v\n", err)
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

Other parameters are passed through a pointer to a apiArchiveConnectionRequest struct via the builder pattern


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


## CreateConnection

> ConnectionDto CreateConnection(ctx).CreateConnectionDto(createConnectionDto).Execute()



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
	resp, r, err := apiClient.ConnectionsAPI.CreateConnection(context.Background()).CreateConnectionDto(createConnectionDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.CreateConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateConnection`: ConnectionDto
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.CreateConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateConnectionRequest struct via the builder pattern


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


## DataLocationMethod

> map[string]interface{} DataLocationMethod(ctx, connectionIdOrKey, dataLocationKey, methodKey).Execute()



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
	resp, r, err := apiClient.ConnectionsAPI.DataLocationMethod(context.Background(), connectionIdOrKey, dataLocationKey, methodKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.DataLocationMethod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DataLocationMethod`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.DataLocationMethod`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiDataLocationMethodRequest struct via the builder pattern


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


## GetConnection

> ConnectionDto GetConnection(ctx, connectionIdOrKey).Execute()





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
	resp, r, err := apiClient.ConnectionsAPI.GetConnection(context.Background(), connectionIdOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.GetConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnection`: ConnectionDto
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.GetConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionIdOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionRequest struct via the builder pattern


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


## GetDataLocationSpec

> map[string]interface{} GetDataLocationSpec(ctx, connectionIdOrKey, dataLocationKey).Execute()



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
	resp, r, err := apiClient.ConnectionsAPI.GetDataLocationSpec(context.Background(), connectionIdOrKey, dataLocationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.GetDataLocationSpec``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDataLocationSpec`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.GetDataLocationSpec`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionIdOrKey** | **string** |  | 
**dataLocationKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataLocationSpecRequest struct via the builder pattern


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


## GetOperation

> map[string]interface{} GetOperation(ctx, connectionIdOrKey, operationKey).Execute()



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
	resp, r, err := apiClient.ConnectionsAPI.GetOperation(context.Background(), connectionIdOrKey, operationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.GetOperation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOperation`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.GetOperation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionIdOrKey** | **string** |  | 
**operationKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOperationRequest struct via the builder pattern


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


## ListConnections

> ListConnections200Response ListConnections(ctx).Id(id).UserId(userId).Name(name).IsTest(isTest).Disconnected(disconnected).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()



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
	name := "name_example" // string |  (optional)
	isTest := true // bool |  (optional)
	disconnected := true // bool |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsAPI.ListConnections(context.Background()).Id(id).UserId(userId).Name(name).IsTest(isTest).Disconnected(disconnected).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.ListConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListConnections`: ListConnections200Response
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.ListConnections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **userId** | **string** |  | 
 **name** | **string** |  | 
 **isTest** | **bool** |  | 
 **disconnected** | **bool** |  | 
 **integrationId** | **string** |  | 
 **integrationKey** | **string** |  | 

### Return type

[**ListConnections200Response**](ListConnections200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDataLocations

> []map[string]interface{} ListDataLocations(ctx, connectionIdOrKey).Execute()



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
	resp, r, err := apiClient.ConnectionsAPI.ListDataLocations(context.Background(), connectionIdOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.ListDataLocations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDataLocations`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.ListDataLocations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionIdOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDataLocationsRequest struct via the builder pattern


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


## ListOperations

> []map[string]interface{} ListOperations(ctx, connectionIdOrKey).Execute()



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
	resp, r, err := apiClient.ConnectionsAPI.ListOperations(context.Background(), connectionIdOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.ListOperations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOperations`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.ListOperations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionIdOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOperationsRequest struct via the builder pattern


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


## MakeConnectionRequestDelete

> map[string]interface{} MakeConnectionRequestDelete(ctx, connectionIdOrKey).ConnectionRequestDto(connectionRequestDto).Execute()



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
	connectionRequestDto := *openapiclient.NewConnectionRequestDto("Path_example", "Method_example") // ConnectionRequestDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsAPI.MakeConnectionRequestDelete(context.Background(), connectionIdOrKey).ConnectionRequestDto(connectionRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.MakeConnectionRequestDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MakeConnectionRequestDelete`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.MakeConnectionRequestDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionIdOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMakeConnectionRequestDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connectionRequestDto** | [**ConnectionRequestDto**](ConnectionRequestDto.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MakeConnectionRequestGet

> map[string]interface{} MakeConnectionRequestGet(ctx, connectionIdOrKey).ConnectionRequestDto(connectionRequestDto).Execute()



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
	connectionRequestDto := *openapiclient.NewConnectionRequestDto("Path_example", "Method_example") // ConnectionRequestDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsAPI.MakeConnectionRequestGet(context.Background(), connectionIdOrKey).ConnectionRequestDto(connectionRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.MakeConnectionRequestGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MakeConnectionRequestGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.MakeConnectionRequestGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionIdOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMakeConnectionRequestGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connectionRequestDto** | [**ConnectionRequestDto**](ConnectionRequestDto.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MakeConnectionRequestHead

> map[string]interface{} MakeConnectionRequestHead(ctx, connectionIdOrKey).ConnectionRequestDto(connectionRequestDto).Execute()



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
	connectionRequestDto := *openapiclient.NewConnectionRequestDto("Path_example", "Method_example") // ConnectionRequestDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsAPI.MakeConnectionRequestHead(context.Background(), connectionIdOrKey).ConnectionRequestDto(connectionRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.MakeConnectionRequestHead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MakeConnectionRequestHead`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.MakeConnectionRequestHead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionIdOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMakeConnectionRequestHeadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connectionRequestDto** | [**ConnectionRequestDto**](ConnectionRequestDto.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MakeConnectionRequestOptions

> map[string]interface{} MakeConnectionRequestOptions(ctx, connectionIdOrKey).ConnectionRequestDto(connectionRequestDto).Execute()



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
	connectionRequestDto := *openapiclient.NewConnectionRequestDto("Path_example", "Method_example") // ConnectionRequestDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsAPI.MakeConnectionRequestOptions(context.Background(), connectionIdOrKey).ConnectionRequestDto(connectionRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.MakeConnectionRequestOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MakeConnectionRequestOptions`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.MakeConnectionRequestOptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionIdOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMakeConnectionRequestOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connectionRequestDto** | [**ConnectionRequestDto**](ConnectionRequestDto.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MakeConnectionRequestPatch

> map[string]interface{} MakeConnectionRequestPatch(ctx, connectionIdOrKey).ConnectionRequestDto(connectionRequestDto).Execute()



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
	connectionRequestDto := *openapiclient.NewConnectionRequestDto("Path_example", "Method_example") // ConnectionRequestDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsAPI.MakeConnectionRequestPatch(context.Background(), connectionIdOrKey).ConnectionRequestDto(connectionRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.MakeConnectionRequestPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MakeConnectionRequestPatch`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.MakeConnectionRequestPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionIdOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMakeConnectionRequestPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connectionRequestDto** | [**ConnectionRequestDto**](ConnectionRequestDto.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MakeConnectionRequestPost

> map[string]interface{} MakeConnectionRequestPost(ctx, connectionIdOrKey).ConnectionRequestDto(connectionRequestDto).Execute()



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
	connectionRequestDto := *openapiclient.NewConnectionRequestDto("Path_example", "Method_example") // ConnectionRequestDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsAPI.MakeConnectionRequestPost(context.Background(), connectionIdOrKey).ConnectionRequestDto(connectionRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.MakeConnectionRequestPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MakeConnectionRequestPost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.MakeConnectionRequestPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionIdOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMakeConnectionRequestPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connectionRequestDto** | [**ConnectionRequestDto**](ConnectionRequestDto.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MakeConnectionRequestPut

> map[string]interface{} MakeConnectionRequestPut(ctx, connectionIdOrKey).ConnectionRequestDto(connectionRequestDto).Execute()



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
	connectionRequestDto := *openapiclient.NewConnectionRequestDto("Path_example", "Method_example") // ConnectionRequestDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsAPI.MakeConnectionRequestPut(context.Background(), connectionIdOrKey).ConnectionRequestDto(connectionRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.MakeConnectionRequestPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MakeConnectionRequestPut`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.MakeConnectionRequestPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionIdOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMakeConnectionRequestPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connectionRequestDto** | [**ConnectionRequestDto**](ConnectionRequestDto.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchConnection

> ConnectionDto PatchConnection(ctx, connectionIdOrKey).UpdateConnectionDto(updateConnectionDto).Execute()



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
	resp, r, err := apiClient.ConnectionsAPI.PatchConnection(context.Background(), connectionIdOrKey).UpdateConnectionDto(updateConnectionDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.PatchConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchConnection`: ConnectionDto
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.PatchConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionIdOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchConnectionRequest struct via the builder pattern


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


## ProxyRequestDelete

> ProxyRequestDelete(ctx, connectionIdOrKey, path).Execute()



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
	r, err := apiClient.ConnectionsAPI.ProxyRequestDelete(context.Background(), connectionIdOrKey, path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.ProxyRequestDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiProxyRequestDeleteRequest struct via the builder pattern


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


## ProxyRequestGet

> ProxyRequestGet(ctx, connectionIdOrKey, path).Execute()



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
	r, err := apiClient.ConnectionsAPI.ProxyRequestGet(context.Background(), connectionIdOrKey, path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.ProxyRequestGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiProxyRequestGetRequest struct via the builder pattern


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


## ProxyRequestHead

> ProxyRequestHead(ctx, connectionIdOrKey, path).Execute()



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
	r, err := apiClient.ConnectionsAPI.ProxyRequestHead(context.Background(), connectionIdOrKey, path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.ProxyRequestHead``: %v\n", err)
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

Other parameters are passed through a pointer to a apiProxyRequestHeadRequest struct via the builder pattern


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


## ProxyRequestOptions

> ProxyRequestOptions(ctx, connectionIdOrKey, path).Execute()



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
	r, err := apiClient.ConnectionsAPI.ProxyRequestOptions(context.Background(), connectionIdOrKey, path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.ProxyRequestOptions``: %v\n", err)
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

Other parameters are passed through a pointer to a apiProxyRequestOptionsRequest struct via the builder pattern


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


## ProxyRequestPatch

> ProxyRequestPatch(ctx, connectionIdOrKey, path).Execute()



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
	r, err := apiClient.ConnectionsAPI.ProxyRequestPatch(context.Background(), connectionIdOrKey, path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.ProxyRequestPatch``: %v\n", err)
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

Other parameters are passed through a pointer to a apiProxyRequestPatchRequest struct via the builder pattern


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


## ProxyRequestPost

> ProxyRequestPost(ctx, connectionIdOrKey, path).Execute()



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
	r, err := apiClient.ConnectionsAPI.ProxyRequestPost(context.Background(), connectionIdOrKey, path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.ProxyRequestPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiProxyRequestPostRequest struct via the builder pattern


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


## ProxyRequestPut

> ProxyRequestPut(ctx, connectionIdOrKey, path).Execute()



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
	r, err := apiClient.ConnectionsAPI.ProxyRequestPut(context.Background(), connectionIdOrKey, path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.ProxyRequestPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiProxyRequestPutRequest struct via the builder pattern


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


## RefreshConnectionCredentials

> RefreshConnectionCredentials(ctx, connectionIdOrKey).Execute()



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
	r, err := apiClient.ConnectionsAPI.RefreshConnectionCredentials(context.Background(), connectionIdOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.RefreshConnectionCredentials``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRefreshConnectionCredentialsRequest struct via the builder pattern


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


## RunOperation

> OperationRunResponseDto RunOperation(ctx, connectionIdOrKey, operationKey).Execute()



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
	resp, r, err := apiClient.ConnectionsAPI.RunOperation(context.Background(), connectionIdOrKey, operationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.RunOperation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RunOperation`: OperationRunResponseDto
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.RunOperation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionIdOrKey** | **string** |  | 
**operationKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRunOperationRequest struct via the builder pattern


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


## TestConnection

> TestConnection(ctx, connectionIdOrKey).Execute()



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
	r, err := apiClient.ConnectionsAPI.TestConnection(context.Background(), connectionIdOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.TestConnection``: %v\n", err)
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

Other parameters are passed through a pointer to a apiTestConnectionRequest struct via the builder pattern


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

