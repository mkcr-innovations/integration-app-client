# \AppDataSchemasAPI

All URIs are relative to *https://api.integration.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppDataSchemaByIdControllerArchiveAppDataSchema**](AppDataSchemasAPI.md#AppDataSchemaByIdControllerArchiveAppDataSchema) | **Delete** /app-data-schemas/{id} | Archive app data schema
[**AppDataSchemaByIdControllerExport**](AppDataSchemasAPI.md#AppDataSchemaByIdControllerExport) | **Get** /app-data-schemas/{id}/export | 
[**AppDataSchemaByIdControllerGetAppDataSchema**](AppDataSchemasAPI.md#AppDataSchemaByIdControllerGetAppDataSchema) | **Get** /app-data-schemas/{id} | Get app data schema
[**AppDataSchemaByIdControllerPatchAppDataSchema**](AppDataSchemasAPI.md#AppDataSchemaByIdControllerPatchAppDataSchema) | **Patch** /app-data-schemas/{id} | Patch app data schema
[**AppDataSchemaByIdControllerPutAppDataSchema**](AppDataSchemasAPI.md#AppDataSchemaByIdControllerPutAppDataSchema) | **Put** /app-data-schemas/{id} | Update app data schema
[**AppDataSchemaInstanceControllerArchiveAppDataSchemaInstance**](AppDataSchemasAPI.md#AppDataSchemaInstanceControllerArchiveAppDataSchemaInstance) | **Delete** /app-data-schema-instance | Archive app data schema instance
[**AppDataSchemaInstanceControllerCreateAppDataSchemaInstance**](AppDataSchemasAPI.md#AppDataSchemaInstanceControllerCreateAppDataSchemaInstance) | **Post** /app-data-schema-instance | Create app data schema instance
[**AppDataSchemaInstanceControllerGetAppDataSchemaInstance**](AppDataSchemasAPI.md#AppDataSchemaInstanceControllerGetAppDataSchemaInstance) | **Get** /app-data-schema-instance | Get app data schema instance
[**AppDataSchemaInstanceControllerPatchAppDataSchemaInstance**](AppDataSchemasAPI.md#AppDataSchemaInstanceControllerPatchAppDataSchemaInstance) | **Patch** /app-data-schema-instance | Patch app data schema instance
[**AppDataSchemaInstanceControllerPutAppDataSchemaInstance**](AppDataSchemasAPI.md#AppDataSchemaInstanceControllerPutAppDataSchemaInstance) | **Put** /app-data-schema-instance | Update app data schema instance
[**AppDataSchemaInstanceControllerSetupAppDataSchemaInstance**](AppDataSchemasAPI.md#AppDataSchemaInstanceControllerSetupAppDataSchemaInstance) | **Post** /app-data-schema-instance/setup | Setup app data schema instance
[**AppDataSchemaInstancesControllerListAppDataSchemaInstances**](AppDataSchemasAPI.md#AppDataSchemaInstancesControllerListAppDataSchemaInstances) | **Get** /app-data-schema-instances | List app data schema instances
[**AppDataSchemasControllerCreateAppDataSchema**](AppDataSchemasAPI.md#AppDataSchemasControllerCreateAppDataSchema) | **Post** /app-data-schemas | Create app data schema
[**AppDataSchemasControllerListAppDataSchemas**](AppDataSchemasAPI.md#AppDataSchemasControllerListAppDataSchemas) | **Get** /app-data-schemas | List app data schemas



## AppDataSchemaByIdControllerArchiveAppDataSchema

> AppDataSchemaByIdControllerArchiveAppDataSchema(ctx, id).Execute()

Archive app data schema

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AppDataSchemasAPI.AppDataSchemaByIdControllerArchiveAppDataSchema(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppDataSchemasAPI.AppDataSchemaByIdControllerArchiveAppDataSchema``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAppDataSchemaByIdControllerArchiveAppDataSchemaRequest struct via the builder pattern


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


## AppDataSchemaByIdControllerExport

> AppDataSchemaExportDto AppDataSchemaByIdControllerExport(ctx, id).Key(key).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppDataSchemasAPI.AppDataSchemaByIdControllerExport(context.Background(), id).Key(key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppDataSchemasAPI.AppDataSchemaByIdControllerExport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppDataSchemaByIdControllerExport`: AppDataSchemaExportDto
	fmt.Fprintf(os.Stdout, "Response from `AppDataSchemasAPI.AppDataSchemaByIdControllerExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppDataSchemaByIdControllerExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **key** | **string** |  | 

### Return type

[**AppDataSchemaExportDto**](AppDataSchemaExportDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppDataSchemaByIdControllerGetAppDataSchema

> AppDataSchemaDto AppDataSchemaByIdControllerGetAppDataSchema(ctx, id).Execute()

Get app data schema

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppDataSchemasAPI.AppDataSchemaByIdControllerGetAppDataSchema(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppDataSchemasAPI.AppDataSchemaByIdControllerGetAppDataSchema``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppDataSchemaByIdControllerGetAppDataSchema`: AppDataSchemaDto
	fmt.Fprintf(os.Stdout, "Response from `AppDataSchemasAPI.AppDataSchemaByIdControllerGetAppDataSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppDataSchemaByIdControllerGetAppDataSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## AppDataSchemaByIdControllerPatchAppDataSchema

> AppDataSchemaDto AppDataSchemaByIdControllerPatchAppDataSchema(ctx, id).AppDataSchemaCreateDto(appDataSchemaCreateDto).Execute()

Patch app data schema

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
	appDataSchemaCreateDto := *openapiclient.NewAppDataSchemaCreateDto("Key_example", "Name_example") // AppDataSchemaCreateDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppDataSchemasAPI.AppDataSchemaByIdControllerPatchAppDataSchema(context.Background(), id).AppDataSchemaCreateDto(appDataSchemaCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppDataSchemasAPI.AppDataSchemaByIdControllerPatchAppDataSchema``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppDataSchemaByIdControllerPatchAppDataSchema`: AppDataSchemaDto
	fmt.Fprintf(os.Stdout, "Response from `AppDataSchemasAPI.AppDataSchemaByIdControllerPatchAppDataSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppDataSchemaByIdControllerPatchAppDataSchemaRequest struct via the builder pattern


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


## AppDataSchemaByIdControllerPutAppDataSchema

> AppDataSchemaDto AppDataSchemaByIdControllerPutAppDataSchema(ctx, id).AppDataSchemaCreateDto(appDataSchemaCreateDto).Execute()

Update app data schema

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
	appDataSchemaCreateDto := *openapiclient.NewAppDataSchemaCreateDto("Key_example", "Name_example") // AppDataSchemaCreateDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppDataSchemasAPI.AppDataSchemaByIdControllerPutAppDataSchema(context.Background(), id).AppDataSchemaCreateDto(appDataSchemaCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppDataSchemasAPI.AppDataSchemaByIdControllerPutAppDataSchema``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppDataSchemaByIdControllerPutAppDataSchema`: AppDataSchemaDto
	fmt.Fprintf(os.Stdout, "Response from `AppDataSchemasAPI.AppDataSchemaByIdControllerPutAppDataSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppDataSchemaByIdControllerPutAppDataSchemaRequest struct via the builder pattern


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


## AppDataSchemaInstanceControllerArchiveAppDataSchemaInstance

> AppDataSchemaInstanceControllerArchiveAppDataSchemaInstance(ctx).Id(id).InstanceKey(instanceKey).AppDataSchemaId(appDataSchemaId).AppDataSchemaKey(appDataSchemaKey).AutoCreate(autoCreate).Execute()

Archive app data schema instance

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
	r, err := apiClient.AppDataSchemasAPI.AppDataSchemaInstanceControllerArchiveAppDataSchemaInstance(context.Background()).Id(id).InstanceKey(instanceKey).AppDataSchemaId(appDataSchemaId).AppDataSchemaKey(appDataSchemaKey).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppDataSchemasAPI.AppDataSchemaInstanceControllerArchiveAppDataSchemaInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppDataSchemaInstanceControllerArchiveAppDataSchemaInstanceRequest struct via the builder pattern


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


## AppDataSchemaInstanceControllerCreateAppDataSchemaInstance

> AppDataSchemaInstanceDto AppDataSchemaInstanceControllerCreateAppDataSchemaInstance(ctx).Body(body).Id(id).InstanceKey(instanceKey).AppDataSchemaId(appDataSchemaId).AppDataSchemaKey(appDataSchemaKey).AutoCreate(autoCreate).Execute()

Create app data schema instance

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
	resp, r, err := apiClient.AppDataSchemasAPI.AppDataSchemaInstanceControllerCreateAppDataSchemaInstance(context.Background()).Body(body).Id(id).InstanceKey(instanceKey).AppDataSchemaId(appDataSchemaId).AppDataSchemaKey(appDataSchemaKey).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppDataSchemasAPI.AppDataSchemaInstanceControllerCreateAppDataSchemaInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppDataSchemaInstanceControllerCreateAppDataSchemaInstance`: AppDataSchemaInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `AppDataSchemasAPI.AppDataSchemaInstanceControllerCreateAppDataSchemaInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppDataSchemaInstanceControllerCreateAppDataSchemaInstanceRequest struct via the builder pattern


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


## AppDataSchemaInstanceControllerGetAppDataSchemaInstance

> AppDataSchemaInstanceDto AppDataSchemaInstanceControllerGetAppDataSchemaInstance(ctx).Id(id).InstanceKey(instanceKey).AppDataSchemaId(appDataSchemaId).AppDataSchemaKey(appDataSchemaKey).AutoCreate(autoCreate).Execute()

Get app data schema instance

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
	resp, r, err := apiClient.AppDataSchemasAPI.AppDataSchemaInstanceControllerGetAppDataSchemaInstance(context.Background()).Id(id).InstanceKey(instanceKey).AppDataSchemaId(appDataSchemaId).AppDataSchemaKey(appDataSchemaKey).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppDataSchemasAPI.AppDataSchemaInstanceControllerGetAppDataSchemaInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppDataSchemaInstanceControllerGetAppDataSchemaInstance`: AppDataSchemaInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `AppDataSchemasAPI.AppDataSchemaInstanceControllerGetAppDataSchemaInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppDataSchemaInstanceControllerGetAppDataSchemaInstanceRequest struct via the builder pattern


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


## AppDataSchemaInstanceControllerPatchAppDataSchemaInstance

> AppDataSchemaInstanceDto AppDataSchemaInstanceControllerPatchAppDataSchemaInstance(ctx).Body(body).Id(id).InstanceKey(instanceKey).AppDataSchemaId(appDataSchemaId).AppDataSchemaKey(appDataSchemaKey).AutoCreate(autoCreate).Execute()

Patch app data schema instance

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
	resp, r, err := apiClient.AppDataSchemasAPI.AppDataSchemaInstanceControllerPatchAppDataSchemaInstance(context.Background()).Body(body).Id(id).InstanceKey(instanceKey).AppDataSchemaId(appDataSchemaId).AppDataSchemaKey(appDataSchemaKey).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppDataSchemasAPI.AppDataSchemaInstanceControllerPatchAppDataSchemaInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppDataSchemaInstanceControllerPatchAppDataSchemaInstance`: AppDataSchemaInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `AppDataSchemasAPI.AppDataSchemaInstanceControllerPatchAppDataSchemaInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppDataSchemaInstanceControllerPatchAppDataSchemaInstanceRequest struct via the builder pattern


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


## AppDataSchemaInstanceControllerPutAppDataSchemaInstance

> AppDataSchemaInstanceDto AppDataSchemaInstanceControllerPutAppDataSchemaInstance(ctx).Body(body).Id(id).InstanceKey(instanceKey).AppDataSchemaId(appDataSchemaId).AppDataSchemaKey(appDataSchemaKey).AutoCreate(autoCreate).Execute()

Update app data schema instance

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
	resp, r, err := apiClient.AppDataSchemasAPI.AppDataSchemaInstanceControllerPutAppDataSchemaInstance(context.Background()).Body(body).Id(id).InstanceKey(instanceKey).AppDataSchemaId(appDataSchemaId).AppDataSchemaKey(appDataSchemaKey).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppDataSchemasAPI.AppDataSchemaInstanceControllerPutAppDataSchemaInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppDataSchemaInstanceControllerPutAppDataSchemaInstance`: AppDataSchemaInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `AppDataSchemasAPI.AppDataSchemaInstanceControllerPutAppDataSchemaInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppDataSchemaInstanceControllerPutAppDataSchemaInstanceRequest struct via the builder pattern


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


## AppDataSchemaInstanceControllerSetupAppDataSchemaInstance

> AppDataSchemaInstanceDto AppDataSchemaInstanceControllerSetupAppDataSchemaInstance(ctx).Id(id).InstanceKey(instanceKey).AppDataSchemaId(appDataSchemaId).AppDataSchemaKey(appDataSchemaKey).AutoCreate(autoCreate).Execute()

Setup app data schema instance

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
	resp, r, err := apiClient.AppDataSchemasAPI.AppDataSchemaInstanceControllerSetupAppDataSchemaInstance(context.Background()).Id(id).InstanceKey(instanceKey).AppDataSchemaId(appDataSchemaId).AppDataSchemaKey(appDataSchemaKey).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppDataSchemasAPI.AppDataSchemaInstanceControllerSetupAppDataSchemaInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppDataSchemaInstanceControllerSetupAppDataSchemaInstance`: AppDataSchemaInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `AppDataSchemasAPI.AppDataSchemaInstanceControllerSetupAppDataSchemaInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppDataSchemaInstanceControllerSetupAppDataSchemaInstanceRequest struct via the builder pattern


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


## AppDataSchemaInstancesControllerListAppDataSchemaInstances

> AppDataSchemaInstancesControllerListAppDataSchemaInstances200Response AppDataSchemaInstancesControllerListAppDataSchemaInstances(ctx).Limit(limit).Cursor(cursor).Search(search).ConnectorId(connectorId).Id(id).UserId(userId).AppDataSchemaId(appDataSchemaId).InstanceKey(instanceKey).Execute()

List app data schema instances

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
	userId := "userId_example" // string |  (optional)
	appDataSchemaId := "appDataSchemaId_example" // string |  (optional)
	instanceKey := "instanceKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppDataSchemasAPI.AppDataSchemaInstancesControllerListAppDataSchemaInstances(context.Background()).Limit(limit).Cursor(cursor).Search(search).ConnectorId(connectorId).Id(id).UserId(userId).AppDataSchemaId(appDataSchemaId).InstanceKey(instanceKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppDataSchemasAPI.AppDataSchemaInstancesControllerListAppDataSchemaInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppDataSchemaInstancesControllerListAppDataSchemaInstances`: AppDataSchemaInstancesControllerListAppDataSchemaInstances200Response
	fmt.Fprintf(os.Stdout, "Response from `AppDataSchemasAPI.AppDataSchemaInstancesControllerListAppDataSchemaInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppDataSchemaInstancesControllerListAppDataSchemaInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **float32** |  | 
 **cursor** | **string** |  | 
 **search** | **string** |  | 
 **connectorId** | **string** |  | 
 **id** | **string** |  | 
 **userId** | **string** |  | 
 **appDataSchemaId** | **string** |  | 
 **instanceKey** | **string** |  | 

### Return type

[**AppDataSchemaInstancesControllerListAppDataSchemaInstances200Response**](AppDataSchemaInstancesControllerListAppDataSchemaInstances200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppDataSchemasControllerCreateAppDataSchema

> AppDataSchemaDto AppDataSchemasControllerCreateAppDataSchema(ctx).AppDataSchemaCreateDto(appDataSchemaCreateDto).Execute()

Create app data schema

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
	resp, r, err := apiClient.AppDataSchemasAPI.AppDataSchemasControllerCreateAppDataSchema(context.Background()).AppDataSchemaCreateDto(appDataSchemaCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppDataSchemasAPI.AppDataSchemasControllerCreateAppDataSchema``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppDataSchemasControllerCreateAppDataSchema`: AppDataSchemaDto
	fmt.Fprintf(os.Stdout, "Response from `AppDataSchemasAPI.AppDataSchemasControllerCreateAppDataSchema`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppDataSchemasControllerCreateAppDataSchemaRequest struct via the builder pattern


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


## AppDataSchemasControllerListAppDataSchemas

> AppDataSchemasControllerListAppDataSchemas200Response AppDataSchemasControllerListAppDataSchemas(ctx).Limit(limit).Cursor(cursor).Search(search).ConnectorId(connectorId).Execute()

List app data schemas

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
	resp, r, err := apiClient.AppDataSchemasAPI.AppDataSchemasControllerListAppDataSchemas(context.Background()).Limit(limit).Cursor(cursor).Search(search).ConnectorId(connectorId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppDataSchemasAPI.AppDataSchemasControllerListAppDataSchemas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppDataSchemasControllerListAppDataSchemas`: AppDataSchemasControllerListAppDataSchemas200Response
	fmt.Fprintf(os.Stdout, "Response from `AppDataSchemasAPI.AppDataSchemasControllerListAppDataSchemas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppDataSchemasControllerListAppDataSchemasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **float32** |  | 
 **cursor** | **string** |  | 
 **search** | **string** |  | 
 **connectorId** | **string** |  | 

### Return type

[**AppDataSchemasControllerListAppDataSchemas200Response**](AppDataSchemasControllerListAppDataSchemas200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

