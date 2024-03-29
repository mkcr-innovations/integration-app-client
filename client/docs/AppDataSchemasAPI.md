# \AppDataSchemasAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveAppDataSchema**](AppDataSchemasAPI.md#ArchiveAppDataSchema) | **Delete** /app-data-schemas/{id} | 
[**ArchiveAppDataSchemaInstance**](AppDataSchemasAPI.md#ArchiveAppDataSchemaInstance) | **Delete** /app-data-schema-instance | 
[**ArchiveAppDataSchemas**](AppDataSchemasAPI.md#ArchiveAppDataSchemas) | **Delete** /app-data-schema | 
[**CreateAppDataSchema**](AppDataSchemasAPI.md#CreateAppDataSchema) | **Post** /app-data-schemas | 
[**CreateAppDataSchemaInstance**](AppDataSchemasAPI.md#CreateAppDataSchemaInstance) | **Post** /app-data-schema-instance | 
[**GetAppDataSchema**](AppDataSchemasAPI.md#GetAppDataSchema) | **Get** /app-data-schemas/{id} | 
[**GetAppDataSchemaInstance**](AppDataSchemasAPI.md#GetAppDataSchemaInstance) | **Get** /app-data-schema-instance | 
[**GetAppDataSchemas**](AppDataSchemasAPI.md#GetAppDataSchemas) | **Get** /app-data-schema | 
[**ListAppDataSchemaInstances**](AppDataSchemasAPI.md#ListAppDataSchemaInstances) | **Get** /app-data-schema-instances | 
[**ListAppDataSchemas**](AppDataSchemasAPI.md#ListAppDataSchemas) | **Get** /app-data-schemas | 
[**PatchAppDataSchema**](AppDataSchemasAPI.md#PatchAppDataSchema) | **Patch** /app-data-schemas/{id} | 
[**PatchAppDataSchemaInstance**](AppDataSchemasAPI.md#PatchAppDataSchemaInstance) | **Patch** /app-data-schema-instance | 
[**PatchAppDataSchemas**](AppDataSchemasAPI.md#PatchAppDataSchemas) | **Patch** /app-data-schema | 
[**PutAppDataSchema**](AppDataSchemasAPI.md#PutAppDataSchema) | **Put** /app-data-schemas/{id} | 
[**PutAppDataSchemaInstance**](AppDataSchemasAPI.md#PutAppDataSchemaInstance) | **Put** /app-data-schema-instance | 
[**PutAppDataSchemas**](AppDataSchemasAPI.md#PutAppDataSchemas) | **Put** /app-data-schema | 
[**SetupAppDataSchemaInstance**](AppDataSchemasAPI.md#SetupAppDataSchemaInstance) | **Post** /app-data-schema-instance/setup | 



## ArchiveAppDataSchema

> ArchiveAppDataSchema(ctx, id).AutoCreate(autoCreate).Execute()



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
	id := "id_example" // string | The ID of the app-data-schema to retrive
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AppDataSchemasAPI.ArchiveAppDataSchema(context.Background(), id).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppDataSchemasAPI.ArchiveAppDataSchema``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the app-data-schema to retrive | 

### Other Parameters

Other parameters are passed through a pointer to a apiArchiveAppDataSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


## ArchiveAppDataSchemaInstance

> ArchiveAppDataSchemaInstance(ctx).Id(id).InstanceKey(instanceKey).AppDataSchemaId(appDataSchemaId).AppDataSchemaKey(appDataSchemaKey).AutoCreate(autoCreate).Execute()



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
	instanceKey := "instanceKey_example" // string |  (optional)
	appDataSchemaId := "appDataSchemaId_example" // string |  (optional)
	appDataSchemaKey := "appDataSchemaKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AppDataSchemasAPI.ArchiveAppDataSchemaInstance(context.Background()).Id(id).InstanceKey(instanceKey).AppDataSchemaId(appDataSchemaId).AppDataSchemaKey(appDataSchemaKey).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppDataSchemasAPI.ArchiveAppDataSchemaInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArchiveAppDataSchemaInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **instanceKey** | **string** |  | 
 **appDataSchemaId** | **string** |  | 
 **appDataSchemaKey** | **string** |  | 
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


## ArchiveAppDataSchemas

> ArchiveAppDataSchemas(ctx).AutoCreate(autoCreate).Execute()



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
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AppDataSchemasAPI.ArchiveAppDataSchemas(context.Background()).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppDataSchemasAPI.ArchiveAppDataSchemas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArchiveAppDataSchemasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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


## CreateAppDataSchema

> AppDataSchemaDto CreateAppDataSchema(ctx).AppDataSchemaCreateDto(appDataSchemaCreateDto).Execute()



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
	appDataSchemaCreateDto := *openapiclient.NewAppDataSchemaCreateDto("Key_example", "Name_example") // AppDataSchemaCreateDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppDataSchemasAPI.CreateAppDataSchema(context.Background()).AppDataSchemaCreateDto(appDataSchemaCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppDataSchemasAPI.CreateAppDataSchema``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAppDataSchema`: AppDataSchemaDto
	fmt.Fprintf(os.Stdout, "Response from `AppDataSchemasAPI.CreateAppDataSchema`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAppDataSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appDataSchemaCreateDto** | [**AppDataSchemaCreateDto**](AppDataSchemaCreateDto.md) |  | 

### Return type

[**AppDataSchemaDto**](AppDataSchemaDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAppDataSchemaInstance

> AppDataSchemaInstanceDto CreateAppDataSchemaInstance(ctx).Body(body).Id(id).InstanceKey(instanceKey).AppDataSchemaId(appDataSchemaId).AppDataSchemaKey(appDataSchemaKey).AutoCreate(autoCreate).Execute()



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
	body := map[string]interface{}{ ... } // map[string]interface{} | 
	id := "id_example" // string |  (optional)
	instanceKey := "instanceKey_example" // string |  (optional)
	appDataSchemaId := "appDataSchemaId_example" // string |  (optional)
	appDataSchemaKey := "appDataSchemaKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppDataSchemasAPI.CreateAppDataSchemaInstance(context.Background()).Body(body).Id(id).InstanceKey(instanceKey).AppDataSchemaId(appDataSchemaId).AppDataSchemaKey(appDataSchemaKey).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppDataSchemasAPI.CreateAppDataSchemaInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAppDataSchemaInstance`: AppDataSchemaInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `AppDataSchemasAPI.CreateAppDataSchemaInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAppDataSchemaInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 
 **id** | **string** |  | 
 **instanceKey** | **string** |  | 
 **appDataSchemaId** | **string** |  | 
 **appDataSchemaKey** | **string** |  | 
 **autoCreate** | **bool** |  | 

### Return type

[**AppDataSchemaInstanceDto**](AppDataSchemaInstanceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppDataSchema

> AppDataSchemaDto GetAppDataSchema(ctx, id).AutoCreate(autoCreate).Execute()



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
	id := "id_example" // string | The ID of the app-data-schema to retrive
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppDataSchemasAPI.GetAppDataSchema(context.Background(), id).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppDataSchemasAPI.GetAppDataSchema``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppDataSchema`: AppDataSchemaDto
	fmt.Fprintf(os.Stdout, "Response from `AppDataSchemasAPI.GetAppDataSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the app-data-schema to retrive | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppDataSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **autoCreate** | **bool** |  | 

### Return type

[**AppDataSchemaDto**](AppDataSchemaDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppDataSchemaInstance

> AppDataSchemaInstanceDto GetAppDataSchemaInstance(ctx).Id(id).InstanceKey(instanceKey).AppDataSchemaId(appDataSchemaId).AppDataSchemaKey(appDataSchemaKey).AutoCreate(autoCreate).Execute()



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
	instanceKey := "instanceKey_example" // string |  (optional)
	appDataSchemaId := "appDataSchemaId_example" // string |  (optional)
	appDataSchemaKey := "appDataSchemaKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppDataSchemasAPI.GetAppDataSchemaInstance(context.Background()).Id(id).InstanceKey(instanceKey).AppDataSchemaId(appDataSchemaId).AppDataSchemaKey(appDataSchemaKey).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppDataSchemasAPI.GetAppDataSchemaInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppDataSchemaInstance`: AppDataSchemaInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `AppDataSchemasAPI.GetAppDataSchemaInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAppDataSchemaInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **instanceKey** | **string** |  | 
 **appDataSchemaId** | **string** |  | 
 **appDataSchemaKey** | **string** |  | 
 **autoCreate** | **bool** |  | 

### Return type

[**AppDataSchemaInstanceDto**](AppDataSchemaInstanceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppDataSchemas

> AppDataSchemaDto GetAppDataSchemas(ctx).AutoCreate(autoCreate).Execute()



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
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppDataSchemasAPI.GetAppDataSchemas(context.Background()).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppDataSchemasAPI.GetAppDataSchemas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppDataSchemas`: AppDataSchemaDto
	fmt.Fprintf(os.Stdout, "Response from `AppDataSchemasAPI.GetAppDataSchemas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAppDataSchemasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoCreate** | **bool** |  | 

### Return type

[**AppDataSchemaDto**](AppDataSchemaDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAppDataSchemaInstances

> ListAppDataSchemaInstances200Response ListAppDataSchemaInstances(ctx).Id(id).UserId(userId).AppDataSchemaId(appDataSchemaId).InstanceKey(instanceKey).Execute()



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
	appDataSchemaId := "appDataSchemaId_example" // string |  (optional)
	instanceKey := "instanceKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppDataSchemasAPI.ListAppDataSchemaInstances(context.Background()).Id(id).UserId(userId).AppDataSchemaId(appDataSchemaId).InstanceKey(instanceKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppDataSchemasAPI.ListAppDataSchemaInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAppDataSchemaInstances`: ListAppDataSchemaInstances200Response
	fmt.Fprintf(os.Stdout, "Response from `AppDataSchemasAPI.ListAppDataSchemaInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAppDataSchemaInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **userId** | **string** |  | 
 **appDataSchemaId** | **string** |  | 
 **instanceKey** | **string** |  | 

### Return type

[**ListAppDataSchemaInstances200Response**](ListAppDataSchemaInstances200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAppDataSchemas

> ListAppDataSchemas200Response ListAppDataSchemas(ctx).Execute()



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
	resp, r, err := apiClient.AppDataSchemasAPI.ListAppDataSchemas(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppDataSchemasAPI.ListAppDataSchemas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAppDataSchemas`: ListAppDataSchemas200Response
	fmt.Fprintf(os.Stdout, "Response from `AppDataSchemasAPI.ListAppDataSchemas`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAppDataSchemasRequest struct via the builder pattern


### Return type

[**ListAppDataSchemas200Response**](ListAppDataSchemas200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAppDataSchema

> AppDataSchemaDto PatchAppDataSchema(ctx, id).AppDataSchemaCreateDto(appDataSchemaCreateDto).AutoCreate(autoCreate).Execute()



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
	id := "id_example" // string | The ID of the app-data-schema to retrive
	appDataSchemaCreateDto := *openapiclient.NewAppDataSchemaCreateDto("Key_example", "Name_example") // AppDataSchemaCreateDto | 
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppDataSchemasAPI.PatchAppDataSchema(context.Background(), id).AppDataSchemaCreateDto(appDataSchemaCreateDto).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppDataSchemasAPI.PatchAppDataSchema``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAppDataSchema`: AppDataSchemaDto
	fmt.Fprintf(os.Stdout, "Response from `AppDataSchemasAPI.PatchAppDataSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the app-data-schema to retrive | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAppDataSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appDataSchemaCreateDto** | [**AppDataSchemaCreateDto**](AppDataSchemaCreateDto.md) |  | 
 **autoCreate** | **bool** |  | 

### Return type

[**AppDataSchemaDto**](AppDataSchemaDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAppDataSchemaInstance

> AppDataSchemaInstanceDto PatchAppDataSchemaInstance(ctx).Body(body).Id(id).InstanceKey(instanceKey).AppDataSchemaId(appDataSchemaId).AppDataSchemaKey(appDataSchemaKey).AutoCreate(autoCreate).Execute()



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
	body := map[string]interface{}{ ... } // map[string]interface{} | 
	id := "id_example" // string |  (optional)
	instanceKey := "instanceKey_example" // string |  (optional)
	appDataSchemaId := "appDataSchemaId_example" // string |  (optional)
	appDataSchemaKey := "appDataSchemaKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppDataSchemasAPI.PatchAppDataSchemaInstance(context.Background()).Body(body).Id(id).InstanceKey(instanceKey).AppDataSchemaId(appDataSchemaId).AppDataSchemaKey(appDataSchemaKey).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppDataSchemasAPI.PatchAppDataSchemaInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAppDataSchemaInstance`: AppDataSchemaInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `AppDataSchemasAPI.PatchAppDataSchemaInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchAppDataSchemaInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 
 **id** | **string** |  | 
 **instanceKey** | **string** |  | 
 **appDataSchemaId** | **string** |  | 
 **appDataSchemaKey** | **string** |  | 
 **autoCreate** | **bool** |  | 

### Return type

[**AppDataSchemaInstanceDto**](AppDataSchemaInstanceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAppDataSchemas

> AppDataSchemaDto PatchAppDataSchemas(ctx).AppDataSchemaCreateDto(appDataSchemaCreateDto).AutoCreate(autoCreate).Execute()



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
	appDataSchemaCreateDto := *openapiclient.NewAppDataSchemaCreateDto("Key_example", "Name_example") // AppDataSchemaCreateDto | 
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppDataSchemasAPI.PatchAppDataSchemas(context.Background()).AppDataSchemaCreateDto(appDataSchemaCreateDto).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppDataSchemasAPI.PatchAppDataSchemas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAppDataSchemas`: AppDataSchemaDto
	fmt.Fprintf(os.Stdout, "Response from `AppDataSchemasAPI.PatchAppDataSchemas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchAppDataSchemasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appDataSchemaCreateDto** | [**AppDataSchemaCreateDto**](AppDataSchemaCreateDto.md) |  | 
 **autoCreate** | **bool** |  | 

### Return type

[**AppDataSchemaDto**](AppDataSchemaDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAppDataSchema

> AppDataSchemaDto PutAppDataSchema(ctx, id).AppDataSchemaCreateDto(appDataSchemaCreateDto).AutoCreate(autoCreate).Execute()



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
	id := "id_example" // string | The ID of the app-data-schema to retrive
	appDataSchemaCreateDto := *openapiclient.NewAppDataSchemaCreateDto("Key_example", "Name_example") // AppDataSchemaCreateDto | 
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppDataSchemasAPI.PutAppDataSchema(context.Background(), id).AppDataSchemaCreateDto(appDataSchemaCreateDto).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppDataSchemasAPI.PutAppDataSchema``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAppDataSchema`: AppDataSchemaDto
	fmt.Fprintf(os.Stdout, "Response from `AppDataSchemasAPI.PutAppDataSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the app-data-schema to retrive | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAppDataSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appDataSchemaCreateDto** | [**AppDataSchemaCreateDto**](AppDataSchemaCreateDto.md) |  | 
 **autoCreate** | **bool** |  | 

### Return type

[**AppDataSchemaDto**](AppDataSchemaDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAppDataSchemaInstance

> AppDataSchemaInstanceDto PutAppDataSchemaInstance(ctx).Body(body).Id(id).InstanceKey(instanceKey).AppDataSchemaId(appDataSchemaId).AppDataSchemaKey(appDataSchemaKey).AutoCreate(autoCreate).Execute()



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
	body := map[string]interface{}{ ... } // map[string]interface{} | 
	id := "id_example" // string |  (optional)
	instanceKey := "instanceKey_example" // string |  (optional)
	appDataSchemaId := "appDataSchemaId_example" // string |  (optional)
	appDataSchemaKey := "appDataSchemaKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppDataSchemasAPI.PutAppDataSchemaInstance(context.Background()).Body(body).Id(id).InstanceKey(instanceKey).AppDataSchemaId(appDataSchemaId).AppDataSchemaKey(appDataSchemaKey).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppDataSchemasAPI.PutAppDataSchemaInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAppDataSchemaInstance`: AppDataSchemaInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `AppDataSchemasAPI.PutAppDataSchemaInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutAppDataSchemaInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 
 **id** | **string** |  | 
 **instanceKey** | **string** |  | 
 **appDataSchemaId** | **string** |  | 
 **appDataSchemaKey** | **string** |  | 
 **autoCreate** | **bool** |  | 

### Return type

[**AppDataSchemaInstanceDto**](AppDataSchemaInstanceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAppDataSchemas

> AppDataSchemaDto PutAppDataSchemas(ctx).AppDataSchemaCreateDto(appDataSchemaCreateDto).AutoCreate(autoCreate).Execute()



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
	appDataSchemaCreateDto := *openapiclient.NewAppDataSchemaCreateDto("Key_example", "Name_example") // AppDataSchemaCreateDto | 
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppDataSchemasAPI.PutAppDataSchemas(context.Background()).AppDataSchemaCreateDto(appDataSchemaCreateDto).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppDataSchemasAPI.PutAppDataSchemas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAppDataSchemas`: AppDataSchemaDto
	fmt.Fprintf(os.Stdout, "Response from `AppDataSchemasAPI.PutAppDataSchemas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutAppDataSchemasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appDataSchemaCreateDto** | [**AppDataSchemaCreateDto**](AppDataSchemaCreateDto.md) |  | 
 **autoCreate** | **bool** |  | 

### Return type

[**AppDataSchemaDto**](AppDataSchemaDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetupAppDataSchemaInstance

> AppDataSchemaInstanceDto SetupAppDataSchemaInstance(ctx).Id(id).InstanceKey(instanceKey).AppDataSchemaId(appDataSchemaId).AppDataSchemaKey(appDataSchemaKey).AutoCreate(autoCreate).Execute()



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
	instanceKey := "instanceKey_example" // string |  (optional)
	appDataSchemaId := "appDataSchemaId_example" // string |  (optional)
	appDataSchemaKey := "appDataSchemaKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppDataSchemasAPI.SetupAppDataSchemaInstance(context.Background()).Id(id).InstanceKey(instanceKey).AppDataSchemaId(appDataSchemaId).AppDataSchemaKey(appDataSchemaKey).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppDataSchemasAPI.SetupAppDataSchemaInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetupAppDataSchemaInstance`: AppDataSchemaInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `AppDataSchemasAPI.SetupAppDataSchemaInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetupAppDataSchemaInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **instanceKey** | **string** |  | 
 **appDataSchemaId** | **string** |  | 
 **appDataSchemaKey** | **string** |  | 
 **autoCreate** | **bool** |  | 

### Return type

[**AppDataSchemaInstanceDto**](AppDataSchemaInstanceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

