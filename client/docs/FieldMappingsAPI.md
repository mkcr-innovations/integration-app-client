# \FieldMappingsAPI

All URIs are relative to *https://api.integration.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConnectionLevelFieldMappingControllerArchive**](FieldMappingsAPI.md#ConnectionLevelFieldMappingControllerArchive) | **Delete** /connections/{connectionSelector}/field-mappings/{fieldMappingSelector} | Archive field mapping instance for connection
[**ConnectionLevelFieldMappingControllerCreate**](FieldMappingsAPI.md#ConnectionLevelFieldMappingControllerCreate) | **Post** /connections/{connectionSelector}/field-mappings/{fieldMappingSelector} | Create field mapping instance for connection
[**ConnectionLevelFieldMappingControllerGet**](FieldMappingsAPI.md#ConnectionLevelFieldMappingControllerGet) | **Get** /connections/{connectionSelector}/field-mappings/{fieldMappingSelector} | Get field mapping instance for connection
[**ConnectionLevelFieldMappingControllerPatch**](FieldMappingsAPI.md#ConnectionLevelFieldMappingControllerPatch) | **Patch** /connections/{connectionSelector}/field-mappings/{fieldMappingSelector} | Patch update field mapping instance for connection
[**ConnectionLevelFieldMappingControllerPut**](FieldMappingsAPI.md#ConnectionLevelFieldMappingControllerPut) | **Put** /connections/{connectionSelector}/field-mappings/{fieldMappingSelector} | Update field mapping instance for connection
[**ConnectionLevelFieldMappingControllerReset**](FieldMappingsAPI.md#ConnectionLevelFieldMappingControllerReset) | **Post** /connections/{connectionSelector}/field-mappings/{fieldMappingSelector}/reset | Reset field mapping instance for connection
[**ConnectionLevelFieldMappingControllerSetup**](FieldMappingsAPI.md#ConnectionLevelFieldMappingControllerSetup) | **Post** /connections/{connectionSelector}/field-mappings/{fieldMappingSelector}/setup | Setup field mapping instance for connection
[**ConnectionLevelFieldMappingsControllerList**](FieldMappingsAPI.md#ConnectionLevelFieldMappingsControllerList) | **Get** /connections/{connectionSelector}/field-mappings | List field mapping instances for connection
[**FieldMappingByIdControllerApply**](FieldMappingsAPI.md#FieldMappingByIdControllerApply) | **Post** /field-mappings/{id}/apply | Apply field mapping to integrations
[**FieldMappingByIdControllerArchive**](FieldMappingsAPI.md#FieldMappingByIdControllerArchive) | **Delete** /field-mappings/{id} | Archive field mapping
[**FieldMappingByIdControllerClone**](FieldMappingsAPI.md#FieldMappingByIdControllerClone) | **Post** /field-mappings/{id}/clone | 
[**FieldMappingByIdControllerExport**](FieldMappingsAPI.md#FieldMappingByIdControllerExport) | **Get** /field-mappings/{id}/export | 
[**FieldMappingByIdControllerGet**](FieldMappingsAPI.md#FieldMappingByIdControllerGet) | **Get** /field-mappings/{id} | Get field mapping
[**FieldMappingByIdControllerGetAppSchema**](FieldMappingsAPI.md#FieldMappingByIdControllerGetAppSchema) | **Get** /field-mappings/{id}/app-schema | Get field mapping app schema
[**FieldMappingByIdControllerPatch**](FieldMappingsAPI.md#FieldMappingByIdControllerPatch) | **Patch** /field-mappings/{id} | Patch field mapping
[**FieldMappingByIdControllerPut**](FieldMappingsAPI.md#FieldMappingByIdControllerPut) | **Put** /field-mappings/{id} | Update field mapping
[**FieldMappingByIdControllerReset**](FieldMappingsAPI.md#FieldMappingByIdControllerReset) | **Post** /field-mappings/{id}/reset | Reset field mapping
[**FieldMappingInstanceByIdControllerDeleteFieldMappingInstance**](FieldMappingsAPI.md#FieldMappingInstanceByIdControllerDeleteFieldMappingInstance) | **Delete** /field-mapping-instances/{id} | Delete field mapping instance
[**FieldMappingInstanceByIdControllerGetFieldMappingInstance**](FieldMappingsAPI.md#FieldMappingInstanceByIdControllerGetFieldMappingInstance) | **Get** /field-mapping-instances/{id} | Get field mapping instance
[**FieldMappingInstanceByIdControllerPatchFieldMappingInstance**](FieldMappingsAPI.md#FieldMappingInstanceByIdControllerPatchFieldMappingInstance) | **Patch** /field-mapping-instances/{id} | Patch field mapping instance
[**FieldMappingInstanceByIdControllerPostFieldMappingInstance**](FieldMappingsAPI.md#FieldMappingInstanceByIdControllerPostFieldMappingInstance) | **Post** /field-mapping-instances/{id} | Create field mapping instance
[**FieldMappingInstanceByIdControllerPutFieldMappingInstance**](FieldMappingsAPI.md#FieldMappingInstanceByIdControllerPutFieldMappingInstance) | **Put** /field-mapping-instances/{id} | Update field mapping instance
[**FieldMappingInstanceByIdControllerResetFieldMappingInstance**](FieldMappingsAPI.md#FieldMappingInstanceByIdControllerResetFieldMappingInstance) | **Post** /field-mapping-instances/{id}/reset | Reset field mapping instance
[**FieldMappingInstanceByIdControllerSetup**](FieldMappingsAPI.md#FieldMappingInstanceByIdControllerSetup) | **Post** /field-mapping-instances/{id}/setup | Setup field mapping instance
[**FieldMappingInstancesControllerListFieldMappingInstances**](FieldMappingsAPI.md#FieldMappingInstancesControllerListFieldMappingInstances) | **Get** /field-mapping-instances | List field mapping instances
[**FieldMappingsControllerCreate**](FieldMappingsAPI.md#FieldMappingsControllerCreate) | **Post** /field-mappings | Create field mapping
[**FieldMappingsControllerList**](FieldMappingsAPI.md#FieldMappingsControllerList) | **Get** /field-mappings | List field mappings
[**IntegrationLevelFieldMappingControllerArchive**](FieldMappingsAPI.md#IntegrationLevelFieldMappingControllerArchive) | **Delete** /integrations/{integrationSelector}/field-mappings/{fieldMappingSelector} | Archive filed mapping for integration
[**IntegrationLevelFieldMappingControllerGet**](FieldMappingsAPI.md#IntegrationLevelFieldMappingControllerGet) | **Get** /integrations/{integrationSelector}/field-mappings/{fieldMappingSelector} | Get filed mapping for integration
[**IntegrationLevelFieldMappingControllerGetAppSchema**](FieldMappingsAPI.md#IntegrationLevelFieldMappingControllerGetAppSchema) | **Get** /integrations/{integrationSelector}/field-mappings/{fieldMappingSelector}/app-schema | Get field mapping app schema for integration
[**IntegrationLevelFieldMappingControllerPatch**](FieldMappingsAPI.md#IntegrationLevelFieldMappingControllerPatch) | **Patch** /integrations/{integrationSelector}/field-mappings/{fieldMappingSelector} | Patch update filed mapping for integration
[**IntegrationLevelFieldMappingControllerPut**](FieldMappingsAPI.md#IntegrationLevelFieldMappingControllerPut) | **Put** /integrations/{integrationSelector}/field-mappings/{fieldMappingSelector} | Update filed mapping for integration
[**IntegrationLevelFieldMappingControllerReset**](FieldMappingsAPI.md#IntegrationLevelFieldMappingControllerReset) | **Post** /integrations/{integrationSelector}/field-mappings/{fieldMappingSelector}/reset | Reset filed mapping for integration
[**IntegrationLevelFieldMappingsControllerCreate**](FieldMappingsAPI.md#IntegrationLevelFieldMappingsControllerCreate) | **Post** /integrations/{integrationSelector}/field-mappings | Create filed mapping for integration
[**IntegrationLevelFieldMappingsControllerList**](FieldMappingsAPI.md#IntegrationLevelFieldMappingsControllerList) | **Get** /integrations/{integrationSelector}/field-mappings | List filed mappings for integration



## ConnectionLevelFieldMappingControllerArchive

> ConnectionLevelFieldMappingControllerArchive(ctx, fieldMappingSelector, connectionSelector).InstanceKey(instanceKey).AutoCreate(autoCreate).Execute()

Archive field mapping instance for connection

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
	fieldMappingSelector := "fieldMappingSelector_example" // string | Field Mapping ID or Key
	connectionSelector := "connectionSelector_example" // string | Integration Key, Connection ID, or Integration ID
	instanceKey := "instanceKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FieldMappingsAPI.ConnectionLevelFieldMappingControllerArchive(context.Background(), fieldMappingSelector, connectionSelector).InstanceKey(instanceKey).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldMappingsAPI.ConnectionLevelFieldMappingControllerArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldMappingSelector** | **string** | Field Mapping ID or Key | 
**connectionSelector** | **string** | Integration Key, Connection ID, or Integration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionLevelFieldMappingControllerArchiveRequest struct via the builder pattern


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


## ConnectionLevelFieldMappingControllerCreate

> FieldMappingInstanceDto ConnectionLevelFieldMappingControllerCreate(ctx, connectionSelector, fieldMappingSelector).InstanceKey(instanceKey).Execute()

Create field mapping instance for connection

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
	fieldMappingSelector := "fieldMappingSelector_example" // string | Field Mapping ID or Key
	instanceKey := "instanceKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldMappingsAPI.ConnectionLevelFieldMappingControllerCreate(context.Background(), connectionSelector, fieldMappingSelector).InstanceKey(instanceKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldMappingsAPI.ConnectionLevelFieldMappingControllerCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectionLevelFieldMappingControllerCreate`: FieldMappingInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `FieldMappingsAPI.ConnectionLevelFieldMappingControllerCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionSelector** | **string** | Integration Key, Connection ID, or Integration ID | 
**fieldMappingSelector** | **string** | Field Mapping ID or Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionLevelFieldMappingControllerCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **instanceKey** | **string** |  | 

### Return type

[**FieldMappingInstanceDto**](FieldMappingInstanceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionLevelFieldMappingControllerGet

> FieldMappingInstanceDto ConnectionLevelFieldMappingControllerGet(ctx, fieldMappingSelector, connectionSelector).InstanceKey(instanceKey).AutoCreate(autoCreate).Execute()

Get field mapping instance for connection

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
	fieldMappingSelector := "fieldMappingSelector_example" // string | Field Mapping ID or Key
	connectionSelector := "connectionSelector_example" // string | Integration Key, Connection ID, or Integration ID
	instanceKey := "instanceKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldMappingsAPI.ConnectionLevelFieldMappingControllerGet(context.Background(), fieldMappingSelector, connectionSelector).InstanceKey(instanceKey).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldMappingsAPI.ConnectionLevelFieldMappingControllerGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectionLevelFieldMappingControllerGet`: FieldMappingInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `FieldMappingsAPI.ConnectionLevelFieldMappingControllerGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldMappingSelector** | **string** | Field Mapping ID or Key | 
**connectionSelector** | **string** | Integration Key, Connection ID, or Integration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionLevelFieldMappingControllerGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **instanceKey** | **string** |  | 
 **autoCreate** | **bool** |  | 

### Return type

[**FieldMappingInstanceDto**](FieldMappingInstanceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionLevelFieldMappingControllerPatch

> FieldMappingInstanceDto ConnectionLevelFieldMappingControllerPatch(ctx, fieldMappingSelector, connectionSelector).UpdateFieldMappingInstanceDto(updateFieldMappingInstanceDto).InstanceKey(instanceKey).AutoCreate(autoCreate).Execute()

Patch update field mapping instance for connection

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
	fieldMappingSelector := "fieldMappingSelector_example" // string | Field Mapping ID or Key
	connectionSelector := "connectionSelector_example" // string | Integration Key, Connection ID, or Integration ID
	updateFieldMappingInstanceDto := *openapiclient.NewUpdateFieldMappingInstanceDto() // UpdateFieldMappingInstanceDto | 
	instanceKey := "instanceKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldMappingsAPI.ConnectionLevelFieldMappingControllerPatch(context.Background(), fieldMappingSelector, connectionSelector).UpdateFieldMappingInstanceDto(updateFieldMappingInstanceDto).InstanceKey(instanceKey).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldMappingsAPI.ConnectionLevelFieldMappingControllerPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectionLevelFieldMappingControllerPatch`: FieldMappingInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `FieldMappingsAPI.ConnectionLevelFieldMappingControllerPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldMappingSelector** | **string** | Field Mapping ID or Key | 
**connectionSelector** | **string** | Integration Key, Connection ID, or Integration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionLevelFieldMappingControllerPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateFieldMappingInstanceDto** | [**UpdateFieldMappingInstanceDto**](UpdateFieldMappingInstanceDto.md) |  | 
 **instanceKey** | **string** |  | 
 **autoCreate** | **bool** |  | 

### Return type

[**FieldMappingInstanceDto**](FieldMappingInstanceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionLevelFieldMappingControllerPut

> FieldMappingInstanceDto ConnectionLevelFieldMappingControllerPut(ctx, fieldMappingSelector, connectionSelector).UpdateFieldMappingInstanceDto(updateFieldMappingInstanceDto).InstanceKey(instanceKey).AutoCreate(autoCreate).Execute()

Update field mapping instance for connection

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
	fieldMappingSelector := "fieldMappingSelector_example" // string | Field Mapping ID or Key
	connectionSelector := "connectionSelector_example" // string | Integration Key, Connection ID, or Integration ID
	updateFieldMappingInstanceDto := *openapiclient.NewUpdateFieldMappingInstanceDto() // UpdateFieldMappingInstanceDto | 
	instanceKey := "instanceKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldMappingsAPI.ConnectionLevelFieldMappingControllerPut(context.Background(), fieldMappingSelector, connectionSelector).UpdateFieldMappingInstanceDto(updateFieldMappingInstanceDto).InstanceKey(instanceKey).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldMappingsAPI.ConnectionLevelFieldMappingControllerPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectionLevelFieldMappingControllerPut`: FieldMappingInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `FieldMappingsAPI.ConnectionLevelFieldMappingControllerPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldMappingSelector** | **string** | Field Mapping ID or Key | 
**connectionSelector** | **string** | Integration Key, Connection ID, or Integration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionLevelFieldMappingControllerPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateFieldMappingInstanceDto** | [**UpdateFieldMappingInstanceDto**](UpdateFieldMappingInstanceDto.md) |  | 
 **instanceKey** | **string** |  | 
 **autoCreate** | **bool** |  | 

### Return type

[**FieldMappingInstanceDto**](FieldMappingInstanceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionLevelFieldMappingControllerReset

> FieldMappingInstanceDto ConnectionLevelFieldMappingControllerReset(ctx, fieldMappingSelector, connectionSelector).InstanceKey(instanceKey).AutoCreate(autoCreate).Execute()

Reset field mapping instance for connection

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
	fieldMappingSelector := "fieldMappingSelector_example" // string | Field Mapping ID or Key
	connectionSelector := "connectionSelector_example" // string | Integration Key, Connection ID, or Integration ID
	instanceKey := "instanceKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldMappingsAPI.ConnectionLevelFieldMappingControllerReset(context.Background(), fieldMappingSelector, connectionSelector).InstanceKey(instanceKey).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldMappingsAPI.ConnectionLevelFieldMappingControllerReset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectionLevelFieldMappingControllerReset`: FieldMappingInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `FieldMappingsAPI.ConnectionLevelFieldMappingControllerReset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldMappingSelector** | **string** | Field Mapping ID or Key | 
**connectionSelector** | **string** | Integration Key, Connection ID, or Integration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionLevelFieldMappingControllerResetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **instanceKey** | **string** |  | 
 **autoCreate** | **bool** |  | 

### Return type

[**FieldMappingInstanceDto**](FieldMappingInstanceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionLevelFieldMappingControllerSetup

> FieldMappingInstanceDto ConnectionLevelFieldMappingControllerSetup(ctx, fieldMappingSelector, connectionSelector).InstanceKey(instanceKey).AutoCreate(autoCreate).Execute()

Setup field mapping instance for connection

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
	fieldMappingSelector := "fieldMappingSelector_example" // string | Field Mapping ID or Key
	connectionSelector := "connectionSelector_example" // string | Integration Key, Connection ID, or Integration ID
	instanceKey := "instanceKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldMappingsAPI.ConnectionLevelFieldMappingControllerSetup(context.Background(), fieldMappingSelector, connectionSelector).InstanceKey(instanceKey).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldMappingsAPI.ConnectionLevelFieldMappingControllerSetup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectionLevelFieldMappingControllerSetup`: FieldMappingInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `FieldMappingsAPI.ConnectionLevelFieldMappingControllerSetup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldMappingSelector** | **string** | Field Mapping ID or Key | 
**connectionSelector** | **string** | Integration Key, Connection ID, or Integration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionLevelFieldMappingControllerSetupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **instanceKey** | **string** |  | 
 **autoCreate** | **bool** |  | 

### Return type

[**FieldMappingInstanceDto**](FieldMappingInstanceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionLevelFieldMappingsControllerList

> ConnectionLevelFieldMappingsControllerList(ctx, connectionSelector).Limit(limit).Cursor(cursor).Search(search).ConnectorId(connectorId).IncludeArchived(includeArchived).Id(id).InstanceKey(instanceKey).UserId(userId).FieldMappingId(fieldMappingId).UniversalFieldMappingId(universalFieldMappingId).DataSourceInstanceId(dataSourceInstanceId).Execute()

List field mapping instances for connection

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
	fieldMappingId := "fieldMappingId_example" // string |  (optional)
	universalFieldMappingId := "universalFieldMappingId_example" // string |  (optional)
	dataSourceInstanceId := "dataSourceInstanceId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FieldMappingsAPI.ConnectionLevelFieldMappingsControllerList(context.Background(), connectionSelector).Limit(limit).Cursor(cursor).Search(search).ConnectorId(connectorId).IncludeArchived(includeArchived).Id(id).InstanceKey(instanceKey).UserId(userId).FieldMappingId(fieldMappingId).UniversalFieldMappingId(universalFieldMappingId).DataSourceInstanceId(dataSourceInstanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldMappingsAPI.ConnectionLevelFieldMappingsControllerList``: %v\n", err)
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

Other parameters are passed through a pointer to a apiConnectionLevelFieldMappingsControllerListRequest struct via the builder pattern


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
 **fieldMappingId** | **string** |  | 
 **universalFieldMappingId** | **string** |  | 
 **dataSourceInstanceId** | **string** |  | 

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


## FieldMappingByIdControllerApply

> []FieldMappingDto FieldMappingByIdControllerApply(ctx, id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()

Apply field mapping to integrations

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldMappingsAPI.FieldMappingByIdControllerApply(context.Background(), id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldMappingsAPI.FieldMappingByIdControllerApply``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FieldMappingByIdControllerApply`: []FieldMappingDto
	fmt.Fprintf(os.Stdout, "Response from `FieldMappingsAPI.FieldMappingByIdControllerApply`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFieldMappingByIdControllerApplyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **key** | **string** |  | 
 **integrationId** | **string** |  | 
 **integrationKey** | **string** |  | 

### Return type

[**[]FieldMappingDto**](FieldMappingDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FieldMappingByIdControllerArchive

> FieldMappingByIdControllerArchive(ctx, id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()

Archive field mapping

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FieldMappingsAPI.FieldMappingByIdControllerArchive(context.Background(), id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldMappingsAPI.FieldMappingByIdControllerArchive``: %v\n", err)
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

Other parameters are passed through a pointer to a apiFieldMappingByIdControllerArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **key** | **string** |  | 
 **integrationId** | **string** |  | 
 **integrationKey** | **string** |  | 

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


## FieldMappingByIdControllerClone

> FieldMappingDto FieldMappingByIdControllerClone(ctx, id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldMappingsAPI.FieldMappingByIdControllerClone(context.Background(), id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldMappingsAPI.FieldMappingByIdControllerClone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FieldMappingByIdControllerClone`: FieldMappingDto
	fmt.Fprintf(os.Stdout, "Response from `FieldMappingsAPI.FieldMappingByIdControllerClone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFieldMappingByIdControllerCloneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **key** | **string** |  | 
 **integrationId** | **string** |  | 
 **integrationKey** | **string** |  | 

### Return type

[**FieldMappingDto**](FieldMappingDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FieldMappingByIdControllerExport

> FieldMappingExportDto FieldMappingByIdControllerExport(ctx, id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldMappingsAPI.FieldMappingByIdControllerExport(context.Background(), id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldMappingsAPI.FieldMappingByIdControllerExport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FieldMappingByIdControllerExport`: FieldMappingExportDto
	fmt.Fprintf(os.Stdout, "Response from `FieldMappingsAPI.FieldMappingByIdControllerExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFieldMappingByIdControllerExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **key** | **string** |  | 
 **integrationId** | **string** |  | 
 **integrationKey** | **string** |  | 

### Return type

[**FieldMappingExportDto**](FieldMappingExportDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FieldMappingByIdControllerGet

> FieldMappingDto FieldMappingByIdControllerGet(ctx, id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()

Get field mapping

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldMappingsAPI.FieldMappingByIdControllerGet(context.Background(), id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldMappingsAPI.FieldMappingByIdControllerGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FieldMappingByIdControllerGet`: FieldMappingDto
	fmt.Fprintf(os.Stdout, "Response from `FieldMappingsAPI.FieldMappingByIdControllerGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFieldMappingByIdControllerGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **key** | **string** |  | 
 **integrationId** | **string** |  | 
 **integrationKey** | **string** |  | 

### Return type

[**FieldMappingDto**](FieldMappingDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FieldMappingByIdControllerGetAppSchema

> map[string]interface{} FieldMappingByIdControllerGetAppSchema(ctx, id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()

Get field mapping app schema

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldMappingsAPI.FieldMappingByIdControllerGetAppSchema(context.Background(), id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldMappingsAPI.FieldMappingByIdControllerGetAppSchema``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FieldMappingByIdControllerGetAppSchema`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FieldMappingsAPI.FieldMappingByIdControllerGetAppSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFieldMappingByIdControllerGetAppSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **key** | **string** |  | 
 **integrationId** | **string** |  | 
 **integrationKey** | **string** |  | 

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


## FieldMappingByIdControllerPatch

> FieldMappingDto FieldMappingByIdControllerPatch(ctx, id).UpdateFieldMappingDto(updateFieldMappingDto).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()

Patch field mapping

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
	updateFieldMappingDto := *openapiclient.NewUpdateFieldMappingDto() // UpdateFieldMappingDto | 
	key := "key_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldMappingsAPI.FieldMappingByIdControllerPatch(context.Background(), id).UpdateFieldMappingDto(updateFieldMappingDto).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldMappingsAPI.FieldMappingByIdControllerPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FieldMappingByIdControllerPatch`: FieldMappingDto
	fmt.Fprintf(os.Stdout, "Response from `FieldMappingsAPI.FieldMappingByIdControllerPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFieldMappingByIdControllerPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateFieldMappingDto** | [**UpdateFieldMappingDto**](UpdateFieldMappingDto.md) |  | 
 **key** | **string** |  | 
 **integrationId** | **string** |  | 
 **integrationKey** | **string** |  | 

### Return type

[**FieldMappingDto**](FieldMappingDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FieldMappingByIdControllerPut

> FieldMappingDto FieldMappingByIdControllerPut(ctx, id).UpdateFieldMappingDto(updateFieldMappingDto).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()

Update field mapping

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
	updateFieldMappingDto := *openapiclient.NewUpdateFieldMappingDto() // UpdateFieldMappingDto | 
	key := "key_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldMappingsAPI.FieldMappingByIdControllerPut(context.Background(), id).UpdateFieldMappingDto(updateFieldMappingDto).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldMappingsAPI.FieldMappingByIdControllerPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FieldMappingByIdControllerPut`: FieldMappingDto
	fmt.Fprintf(os.Stdout, "Response from `FieldMappingsAPI.FieldMappingByIdControllerPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFieldMappingByIdControllerPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateFieldMappingDto** | [**UpdateFieldMappingDto**](UpdateFieldMappingDto.md) |  | 
 **key** | **string** |  | 
 **integrationId** | **string** |  | 
 **integrationKey** | **string** |  | 

### Return type

[**FieldMappingDto**](FieldMappingDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FieldMappingByIdControllerReset

> FieldMappingDto FieldMappingByIdControllerReset(ctx, id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()

Reset field mapping

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldMappingsAPI.FieldMappingByIdControllerReset(context.Background(), id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldMappingsAPI.FieldMappingByIdControllerReset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FieldMappingByIdControllerReset`: FieldMappingDto
	fmt.Fprintf(os.Stdout, "Response from `FieldMappingsAPI.FieldMappingByIdControllerReset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFieldMappingByIdControllerResetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **key** | **string** |  | 
 **integrationId** | **string** |  | 
 **integrationKey** | **string** |  | 

### Return type

[**FieldMappingDto**](FieldMappingDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FieldMappingInstanceByIdControllerDeleteFieldMappingInstance

> FieldMappingInstanceByIdControllerDeleteFieldMappingInstance(ctx, id).FieldMappingKey(fieldMappingKey).FieldMappingId(fieldMappingId).InstanceKey(instanceKey).AutoCreate(autoCreate).FlowKey(flowKey).FlowId(flowId).NodeKey(nodeKey).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()

Delete field mapping instance

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
	fieldMappingKey := "fieldMappingKey_example" // string |  (optional)
	fieldMappingId := "fieldMappingId_example" // string |  (optional)
	instanceKey := "instanceKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)
	flowKey := "flowKey_example" // string |  (optional)
	flowId := "flowId_example" // string |  (optional)
	nodeKey := "nodeKey_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	connectionId := "connectionId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FieldMappingsAPI.FieldMappingInstanceByIdControllerDeleteFieldMappingInstance(context.Background(), id).FieldMappingKey(fieldMappingKey).FieldMappingId(fieldMappingId).InstanceKey(instanceKey).AutoCreate(autoCreate).FlowKey(flowKey).FlowId(flowId).NodeKey(nodeKey).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldMappingsAPI.FieldMappingInstanceByIdControllerDeleteFieldMappingInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiFieldMappingInstanceByIdControllerDeleteFieldMappingInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldMappingKey** | **string** |  | 
 **fieldMappingId** | **string** |  | 
 **instanceKey** | **string** |  | 
 **autoCreate** | **bool** |  | 
 **flowKey** | **string** |  | 
 **flowId** | **string** |  | 
 **nodeKey** | **string** |  | 
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


## FieldMappingInstanceByIdControllerGetFieldMappingInstance

> FieldMappingInstanceDto FieldMappingInstanceByIdControllerGetFieldMappingInstance(ctx, id).FieldMappingKey(fieldMappingKey).FieldMappingId(fieldMappingId).InstanceKey(instanceKey).AutoCreate(autoCreate).FlowKey(flowKey).FlowId(flowId).NodeKey(nodeKey).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()

Get field mapping instance

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
	fieldMappingKey := "fieldMappingKey_example" // string |  (optional)
	fieldMappingId := "fieldMappingId_example" // string |  (optional)
	instanceKey := "instanceKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)
	flowKey := "flowKey_example" // string |  (optional)
	flowId := "flowId_example" // string |  (optional)
	nodeKey := "nodeKey_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	connectionId := "connectionId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldMappingsAPI.FieldMappingInstanceByIdControllerGetFieldMappingInstance(context.Background(), id).FieldMappingKey(fieldMappingKey).FieldMappingId(fieldMappingId).InstanceKey(instanceKey).AutoCreate(autoCreate).FlowKey(flowKey).FlowId(flowId).NodeKey(nodeKey).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldMappingsAPI.FieldMappingInstanceByIdControllerGetFieldMappingInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FieldMappingInstanceByIdControllerGetFieldMappingInstance`: FieldMappingInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `FieldMappingsAPI.FieldMappingInstanceByIdControllerGetFieldMappingInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFieldMappingInstanceByIdControllerGetFieldMappingInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldMappingKey** | **string** |  | 
 **fieldMappingId** | **string** |  | 
 **instanceKey** | **string** |  | 
 **autoCreate** | **bool** |  | 
 **flowKey** | **string** |  | 
 **flowId** | **string** |  | 
 **nodeKey** | **string** |  | 
 **integrationKey** | **string** |  | 
 **integrationId** | **string** |  | 
 **connectionId** | **string** |  | 

### Return type

[**FieldMappingInstanceDto**](FieldMappingInstanceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FieldMappingInstanceByIdControllerPatchFieldMappingInstance

> FieldMappingInstanceDto FieldMappingInstanceByIdControllerPatchFieldMappingInstance(ctx, id).UpdateFieldMappingInstanceDto(updateFieldMappingInstanceDto).FieldMappingKey(fieldMappingKey).FieldMappingId(fieldMappingId).InstanceKey(instanceKey).AutoCreate(autoCreate).FlowKey(flowKey).FlowId(flowId).NodeKey(nodeKey).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()

Patch field mapping instance

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
	updateFieldMappingInstanceDto := *openapiclient.NewUpdateFieldMappingInstanceDto() // UpdateFieldMappingInstanceDto | 
	fieldMappingKey := "fieldMappingKey_example" // string |  (optional)
	fieldMappingId := "fieldMappingId_example" // string |  (optional)
	instanceKey := "instanceKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)
	flowKey := "flowKey_example" // string |  (optional)
	flowId := "flowId_example" // string |  (optional)
	nodeKey := "nodeKey_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	connectionId := "connectionId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldMappingsAPI.FieldMappingInstanceByIdControllerPatchFieldMappingInstance(context.Background(), id).UpdateFieldMappingInstanceDto(updateFieldMappingInstanceDto).FieldMappingKey(fieldMappingKey).FieldMappingId(fieldMappingId).InstanceKey(instanceKey).AutoCreate(autoCreate).FlowKey(flowKey).FlowId(flowId).NodeKey(nodeKey).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldMappingsAPI.FieldMappingInstanceByIdControllerPatchFieldMappingInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FieldMappingInstanceByIdControllerPatchFieldMappingInstance`: FieldMappingInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `FieldMappingsAPI.FieldMappingInstanceByIdControllerPatchFieldMappingInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFieldMappingInstanceByIdControllerPatchFieldMappingInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateFieldMappingInstanceDto** | [**UpdateFieldMappingInstanceDto**](UpdateFieldMappingInstanceDto.md) |  | 
 **fieldMappingKey** | **string** |  | 
 **fieldMappingId** | **string** |  | 
 **instanceKey** | **string** |  | 
 **autoCreate** | **bool** |  | 
 **flowKey** | **string** |  | 
 **flowId** | **string** |  | 
 **nodeKey** | **string** |  | 
 **integrationKey** | **string** |  | 
 **integrationId** | **string** |  | 
 **connectionId** | **string** |  | 

### Return type

[**FieldMappingInstanceDto**](FieldMappingInstanceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FieldMappingInstanceByIdControllerPostFieldMappingInstance

> FieldMappingInstanceDto FieldMappingInstanceByIdControllerPostFieldMappingInstance(ctx, id).UpdateFieldMappingInstanceDto(updateFieldMappingInstanceDto).FieldMappingKey(fieldMappingKey).FieldMappingId(fieldMappingId).InstanceKey(instanceKey).AutoCreate(autoCreate).FlowKey(flowKey).FlowId(flowId).NodeKey(nodeKey).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()

Create field mapping instance

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
	updateFieldMappingInstanceDto := *openapiclient.NewUpdateFieldMappingInstanceDto() // UpdateFieldMappingInstanceDto | 
	fieldMappingKey := "fieldMappingKey_example" // string |  (optional)
	fieldMappingId := "fieldMappingId_example" // string |  (optional)
	instanceKey := "instanceKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)
	flowKey := "flowKey_example" // string |  (optional)
	flowId := "flowId_example" // string |  (optional)
	nodeKey := "nodeKey_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	connectionId := "connectionId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldMappingsAPI.FieldMappingInstanceByIdControllerPostFieldMappingInstance(context.Background(), id).UpdateFieldMappingInstanceDto(updateFieldMappingInstanceDto).FieldMappingKey(fieldMappingKey).FieldMappingId(fieldMappingId).InstanceKey(instanceKey).AutoCreate(autoCreate).FlowKey(flowKey).FlowId(flowId).NodeKey(nodeKey).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldMappingsAPI.FieldMappingInstanceByIdControllerPostFieldMappingInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FieldMappingInstanceByIdControllerPostFieldMappingInstance`: FieldMappingInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `FieldMappingsAPI.FieldMappingInstanceByIdControllerPostFieldMappingInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFieldMappingInstanceByIdControllerPostFieldMappingInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateFieldMappingInstanceDto** | [**UpdateFieldMappingInstanceDto**](UpdateFieldMappingInstanceDto.md) |  | 
 **fieldMappingKey** | **string** |  | 
 **fieldMappingId** | **string** |  | 
 **instanceKey** | **string** |  | 
 **autoCreate** | **bool** |  | 
 **flowKey** | **string** |  | 
 **flowId** | **string** |  | 
 **nodeKey** | **string** |  | 
 **integrationKey** | **string** |  | 
 **integrationId** | **string** |  | 
 **connectionId** | **string** |  | 

### Return type

[**FieldMappingInstanceDto**](FieldMappingInstanceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FieldMappingInstanceByIdControllerPutFieldMappingInstance

> FieldMappingInstanceDto FieldMappingInstanceByIdControllerPutFieldMappingInstance(ctx, id).UpdateFieldMappingInstanceDto(updateFieldMappingInstanceDto).FieldMappingKey(fieldMappingKey).FieldMappingId(fieldMappingId).InstanceKey(instanceKey).AutoCreate(autoCreate).FlowKey(flowKey).FlowId(flowId).NodeKey(nodeKey).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()

Update field mapping instance

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
	updateFieldMappingInstanceDto := *openapiclient.NewUpdateFieldMappingInstanceDto() // UpdateFieldMappingInstanceDto | 
	fieldMappingKey := "fieldMappingKey_example" // string |  (optional)
	fieldMappingId := "fieldMappingId_example" // string |  (optional)
	instanceKey := "instanceKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)
	flowKey := "flowKey_example" // string |  (optional)
	flowId := "flowId_example" // string |  (optional)
	nodeKey := "nodeKey_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	connectionId := "connectionId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldMappingsAPI.FieldMappingInstanceByIdControllerPutFieldMappingInstance(context.Background(), id).UpdateFieldMappingInstanceDto(updateFieldMappingInstanceDto).FieldMappingKey(fieldMappingKey).FieldMappingId(fieldMappingId).InstanceKey(instanceKey).AutoCreate(autoCreate).FlowKey(flowKey).FlowId(flowId).NodeKey(nodeKey).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldMappingsAPI.FieldMappingInstanceByIdControllerPutFieldMappingInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FieldMappingInstanceByIdControllerPutFieldMappingInstance`: FieldMappingInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `FieldMappingsAPI.FieldMappingInstanceByIdControllerPutFieldMappingInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFieldMappingInstanceByIdControllerPutFieldMappingInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateFieldMappingInstanceDto** | [**UpdateFieldMappingInstanceDto**](UpdateFieldMappingInstanceDto.md) |  | 
 **fieldMappingKey** | **string** |  | 
 **fieldMappingId** | **string** |  | 
 **instanceKey** | **string** |  | 
 **autoCreate** | **bool** |  | 
 **flowKey** | **string** |  | 
 **flowId** | **string** |  | 
 **nodeKey** | **string** |  | 
 **integrationKey** | **string** |  | 
 **integrationId** | **string** |  | 
 **connectionId** | **string** |  | 

### Return type

[**FieldMappingInstanceDto**](FieldMappingInstanceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FieldMappingInstanceByIdControllerResetFieldMappingInstance

> FieldMappingInstanceDto FieldMappingInstanceByIdControllerResetFieldMappingInstance(ctx, id).FieldMappingKey(fieldMappingKey).FieldMappingId(fieldMappingId).InstanceKey(instanceKey).AutoCreate(autoCreate).FlowKey(flowKey).FlowId(flowId).NodeKey(nodeKey).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()

Reset field mapping instance

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
	fieldMappingKey := "fieldMappingKey_example" // string |  (optional)
	fieldMappingId := "fieldMappingId_example" // string |  (optional)
	instanceKey := "instanceKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)
	flowKey := "flowKey_example" // string |  (optional)
	flowId := "flowId_example" // string |  (optional)
	nodeKey := "nodeKey_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	connectionId := "connectionId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldMappingsAPI.FieldMappingInstanceByIdControllerResetFieldMappingInstance(context.Background(), id).FieldMappingKey(fieldMappingKey).FieldMappingId(fieldMappingId).InstanceKey(instanceKey).AutoCreate(autoCreate).FlowKey(flowKey).FlowId(flowId).NodeKey(nodeKey).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldMappingsAPI.FieldMappingInstanceByIdControllerResetFieldMappingInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FieldMappingInstanceByIdControllerResetFieldMappingInstance`: FieldMappingInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `FieldMappingsAPI.FieldMappingInstanceByIdControllerResetFieldMappingInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFieldMappingInstanceByIdControllerResetFieldMappingInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldMappingKey** | **string** |  | 
 **fieldMappingId** | **string** |  | 
 **instanceKey** | **string** |  | 
 **autoCreate** | **bool** |  | 
 **flowKey** | **string** |  | 
 **flowId** | **string** |  | 
 **nodeKey** | **string** |  | 
 **integrationKey** | **string** |  | 
 **integrationId** | **string** |  | 
 **connectionId** | **string** |  | 

### Return type

[**FieldMappingInstanceDto**](FieldMappingInstanceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FieldMappingInstanceByIdControllerSetup

> FieldMappingInstanceDto FieldMappingInstanceByIdControllerSetup(ctx, id).FieldMappingKey(fieldMappingKey).FieldMappingId(fieldMappingId).InstanceKey(instanceKey).AutoCreate(autoCreate).FlowKey(flowKey).FlowId(flowId).NodeKey(nodeKey).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()

Setup field mapping instance

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
	fieldMappingKey := "fieldMappingKey_example" // string |  (optional)
	fieldMappingId := "fieldMappingId_example" // string |  (optional)
	instanceKey := "instanceKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)
	flowKey := "flowKey_example" // string |  (optional)
	flowId := "flowId_example" // string |  (optional)
	nodeKey := "nodeKey_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	connectionId := "connectionId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldMappingsAPI.FieldMappingInstanceByIdControllerSetup(context.Background(), id).FieldMappingKey(fieldMappingKey).FieldMappingId(fieldMappingId).InstanceKey(instanceKey).AutoCreate(autoCreate).FlowKey(flowKey).FlowId(flowId).NodeKey(nodeKey).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldMappingsAPI.FieldMappingInstanceByIdControllerSetup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FieldMappingInstanceByIdControllerSetup`: FieldMappingInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `FieldMappingsAPI.FieldMappingInstanceByIdControllerSetup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFieldMappingInstanceByIdControllerSetupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldMappingKey** | **string** |  | 
 **fieldMappingId** | **string** |  | 
 **instanceKey** | **string** |  | 
 **autoCreate** | **bool** |  | 
 **flowKey** | **string** |  | 
 **flowId** | **string** |  | 
 **nodeKey** | **string** |  | 
 **integrationKey** | **string** |  | 
 **integrationId** | **string** |  | 
 **connectionId** | **string** |  | 

### Return type

[**FieldMappingInstanceDto**](FieldMappingInstanceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FieldMappingInstancesControllerListFieldMappingInstances

> FieldMappingInstancesControllerListFieldMappingInstances200Response FieldMappingInstancesControllerListFieldMappingInstances(ctx).Limit(limit).Cursor(cursor).Search(search).ConnectorId(connectorId).IncludeArchived(includeArchived).Id(id).InstanceKey(instanceKey).UserId(userId).ConnectionId(connectionId).IntegrationId(integrationId).IntegrationKey(integrationKey).FieldMappingId(fieldMappingId).UniversalFieldMappingId(universalFieldMappingId).DataSourceInstanceId(dataSourceInstanceId).Execute()

List field mapping instances

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
	connectionId := "connectionId_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	fieldMappingId := "fieldMappingId_example" // string |  (optional)
	universalFieldMappingId := "universalFieldMappingId_example" // string |  (optional)
	dataSourceInstanceId := "dataSourceInstanceId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldMappingsAPI.FieldMappingInstancesControllerListFieldMappingInstances(context.Background()).Limit(limit).Cursor(cursor).Search(search).ConnectorId(connectorId).IncludeArchived(includeArchived).Id(id).InstanceKey(instanceKey).UserId(userId).ConnectionId(connectionId).IntegrationId(integrationId).IntegrationKey(integrationKey).FieldMappingId(fieldMappingId).UniversalFieldMappingId(universalFieldMappingId).DataSourceInstanceId(dataSourceInstanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldMappingsAPI.FieldMappingInstancesControllerListFieldMappingInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FieldMappingInstancesControllerListFieldMappingInstances`: FieldMappingInstancesControllerListFieldMappingInstances200Response
	fmt.Fprintf(os.Stdout, "Response from `FieldMappingsAPI.FieldMappingInstancesControllerListFieldMappingInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFieldMappingInstancesControllerListFieldMappingInstancesRequest struct via the builder pattern


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
 **connectionId** | **string** |  | 
 **integrationId** | **string** |  | 
 **integrationKey** | **string** |  | 
 **fieldMappingId** | **string** |  | 
 **universalFieldMappingId** | **string** |  | 
 **dataSourceInstanceId** | **string** |  | 

### Return type

[**FieldMappingInstancesControllerListFieldMappingInstances200Response**](FieldMappingInstancesControllerListFieldMappingInstances200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FieldMappingsControllerCreate

> FieldMappingDto FieldMappingsControllerCreate(ctx).CreateFieldMappingDto(createFieldMappingDto).Execute()

Create field mapping

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
	createFieldMappingDto := *openapiclient.NewCreateFieldMappingDto("Key_example") // CreateFieldMappingDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldMappingsAPI.FieldMappingsControllerCreate(context.Background()).CreateFieldMappingDto(createFieldMappingDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldMappingsAPI.FieldMappingsControllerCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FieldMappingsControllerCreate`: FieldMappingDto
	fmt.Fprintf(os.Stdout, "Response from `FieldMappingsAPI.FieldMappingsControllerCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFieldMappingsControllerCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFieldMappingDto** | [**CreateFieldMappingDto**](CreateFieldMappingDto.md) |  | 

### Return type

[**FieldMappingDto**](FieldMappingDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FieldMappingsControllerList

> FieldMappingsControllerList200Response FieldMappingsControllerList(ctx).Limit(limit).Cursor(cursor).Search(search).ConnectorId(connectorId).IncludeArchived(includeArchived).UniversalFieldMappingId(universalFieldMappingId).IntegrationId(integrationId).Execute()

List field mappings

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
	universalFieldMappingId := "universalFieldMappingId_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldMappingsAPI.FieldMappingsControllerList(context.Background()).Limit(limit).Cursor(cursor).Search(search).ConnectorId(connectorId).IncludeArchived(includeArchived).UniversalFieldMappingId(universalFieldMappingId).IntegrationId(integrationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldMappingsAPI.FieldMappingsControllerList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FieldMappingsControllerList`: FieldMappingsControllerList200Response
	fmt.Fprintf(os.Stdout, "Response from `FieldMappingsAPI.FieldMappingsControllerList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFieldMappingsControllerListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **float32** |  | 
 **cursor** | **string** |  | 
 **search** | **string** |  | 
 **connectorId** | **string** |  | 
 **includeArchived** | **bool** |  | 
 **universalFieldMappingId** | **string** |  | 
 **integrationId** | **string** |  | 

### Return type

[**FieldMappingsControllerList200Response**](FieldMappingsControllerList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationLevelFieldMappingControllerArchive

> IntegrationLevelFieldMappingControllerArchive(ctx, fieldMappingSelector, integrationSelector).Execute()

Archive filed mapping for integration

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
	fieldMappingSelector := "fieldMappingSelector_example" // string | Field Mapping ID or Key
	integrationSelector := "integrationSelector_example" // string | Integration ID or Key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FieldMappingsAPI.IntegrationLevelFieldMappingControllerArchive(context.Background(), fieldMappingSelector, integrationSelector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldMappingsAPI.IntegrationLevelFieldMappingControllerArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldMappingSelector** | **string** | Field Mapping ID or Key | 
**integrationSelector** | **string** | Integration ID or Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationLevelFieldMappingControllerArchiveRequest struct via the builder pattern


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


## IntegrationLevelFieldMappingControllerGet

> FieldMappingDto IntegrationLevelFieldMappingControllerGet(ctx, fieldMappingSelector, integrationSelector).Execute()

Get filed mapping for integration

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
	fieldMappingSelector := "fieldMappingSelector_example" // string | Field Mapping ID or Key
	integrationSelector := "integrationSelector_example" // string | Integration ID or Key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldMappingsAPI.IntegrationLevelFieldMappingControllerGet(context.Background(), fieldMappingSelector, integrationSelector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldMappingsAPI.IntegrationLevelFieldMappingControllerGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationLevelFieldMappingControllerGet`: FieldMappingDto
	fmt.Fprintf(os.Stdout, "Response from `FieldMappingsAPI.IntegrationLevelFieldMappingControllerGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldMappingSelector** | **string** | Field Mapping ID or Key | 
**integrationSelector** | **string** | Integration ID or Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationLevelFieldMappingControllerGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FieldMappingDto**](FieldMappingDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationLevelFieldMappingControllerGetAppSchema

> map[string]interface{} IntegrationLevelFieldMappingControllerGetAppSchema(ctx, fieldMappingSelector, integrationSelector).Execute()

Get field mapping app schema for integration

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
	fieldMappingSelector := "fieldMappingSelector_example" // string | Field Mapping ID or Key
	integrationSelector := "integrationSelector_example" // string | Integration ID or Key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldMappingsAPI.IntegrationLevelFieldMappingControllerGetAppSchema(context.Background(), fieldMappingSelector, integrationSelector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldMappingsAPI.IntegrationLevelFieldMappingControllerGetAppSchema``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationLevelFieldMappingControllerGetAppSchema`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FieldMappingsAPI.IntegrationLevelFieldMappingControllerGetAppSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldMappingSelector** | **string** | Field Mapping ID or Key | 
**integrationSelector** | **string** | Integration ID or Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationLevelFieldMappingControllerGetAppSchemaRequest struct via the builder pattern


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


## IntegrationLevelFieldMappingControllerPatch

> FieldMappingDto IntegrationLevelFieldMappingControllerPatch(ctx, fieldMappingSelector, integrationSelector).UpdateFieldMappingDto(updateFieldMappingDto).Execute()

Patch update filed mapping for integration

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
	fieldMappingSelector := "fieldMappingSelector_example" // string | Field Mapping ID or Key
	integrationSelector := "integrationSelector_example" // string | Integration ID or Key
	updateFieldMappingDto := *openapiclient.NewUpdateFieldMappingDto() // UpdateFieldMappingDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldMappingsAPI.IntegrationLevelFieldMappingControllerPatch(context.Background(), fieldMappingSelector, integrationSelector).UpdateFieldMappingDto(updateFieldMappingDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldMappingsAPI.IntegrationLevelFieldMappingControllerPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationLevelFieldMappingControllerPatch`: FieldMappingDto
	fmt.Fprintf(os.Stdout, "Response from `FieldMappingsAPI.IntegrationLevelFieldMappingControllerPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldMappingSelector** | **string** | Field Mapping ID or Key | 
**integrationSelector** | **string** | Integration ID or Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationLevelFieldMappingControllerPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateFieldMappingDto** | [**UpdateFieldMappingDto**](UpdateFieldMappingDto.md) |  | 

### Return type

[**FieldMappingDto**](FieldMappingDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationLevelFieldMappingControllerPut

> FieldMappingDto IntegrationLevelFieldMappingControllerPut(ctx, fieldMappingSelector, integrationSelector).UpdateFieldMappingDto(updateFieldMappingDto).Execute()

Update filed mapping for integration

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
	fieldMappingSelector := "fieldMappingSelector_example" // string | Field Mapping ID or Key
	integrationSelector := "integrationSelector_example" // string | Integration ID or Key
	updateFieldMappingDto := *openapiclient.NewUpdateFieldMappingDto() // UpdateFieldMappingDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldMappingsAPI.IntegrationLevelFieldMappingControllerPut(context.Background(), fieldMappingSelector, integrationSelector).UpdateFieldMappingDto(updateFieldMappingDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldMappingsAPI.IntegrationLevelFieldMappingControllerPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationLevelFieldMappingControllerPut`: FieldMappingDto
	fmt.Fprintf(os.Stdout, "Response from `FieldMappingsAPI.IntegrationLevelFieldMappingControllerPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldMappingSelector** | **string** | Field Mapping ID or Key | 
**integrationSelector** | **string** | Integration ID or Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationLevelFieldMappingControllerPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateFieldMappingDto** | [**UpdateFieldMappingDto**](UpdateFieldMappingDto.md) |  | 

### Return type

[**FieldMappingDto**](FieldMappingDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationLevelFieldMappingControllerReset

> FieldMappingDto IntegrationLevelFieldMappingControllerReset(ctx, fieldMappingSelector, integrationSelector).Execute()

Reset filed mapping for integration

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
	fieldMappingSelector := "fieldMappingSelector_example" // string | Field Mapping ID or Key
	integrationSelector := "integrationSelector_example" // string | Integration ID or Key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldMappingsAPI.IntegrationLevelFieldMappingControllerReset(context.Background(), fieldMappingSelector, integrationSelector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldMappingsAPI.IntegrationLevelFieldMappingControllerReset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationLevelFieldMappingControllerReset`: FieldMappingDto
	fmt.Fprintf(os.Stdout, "Response from `FieldMappingsAPI.IntegrationLevelFieldMappingControllerReset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldMappingSelector** | **string** | Field Mapping ID or Key | 
**integrationSelector** | **string** | Integration ID or Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationLevelFieldMappingControllerResetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FieldMappingDto**](FieldMappingDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationLevelFieldMappingsControllerCreate

> FieldMappingDto IntegrationLevelFieldMappingsControllerCreate(ctx, integrationSelector).CreateIntegrationLevelFieldMappingDto(createIntegrationLevelFieldMappingDto).Execute()

Create filed mapping for integration

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
	createIntegrationLevelFieldMappingDto := *openapiclient.NewCreateIntegrationLevelFieldMappingDto("Key_example") // CreateIntegrationLevelFieldMappingDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldMappingsAPI.IntegrationLevelFieldMappingsControllerCreate(context.Background(), integrationSelector).CreateIntegrationLevelFieldMappingDto(createIntegrationLevelFieldMappingDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldMappingsAPI.IntegrationLevelFieldMappingsControllerCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationLevelFieldMappingsControllerCreate`: FieldMappingDto
	fmt.Fprintf(os.Stdout, "Response from `FieldMappingsAPI.IntegrationLevelFieldMappingsControllerCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationSelector** | **string** | Integration ID or Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationLevelFieldMappingsControllerCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createIntegrationLevelFieldMappingDto** | [**CreateIntegrationLevelFieldMappingDto**](CreateIntegrationLevelFieldMappingDto.md) |  | 

### Return type

[**FieldMappingDto**](FieldMappingDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationLevelFieldMappingsControllerList

> FieldMappingsControllerList200Response IntegrationLevelFieldMappingsControllerList(ctx, integrationSelector).Limit(limit).Cursor(cursor).Search(search).ConnectorId(connectorId).IncludeArchived(includeArchived).UniversalFieldMappingId(universalFieldMappingId).Execute()

List filed mappings for integration

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
	universalFieldMappingId := "universalFieldMappingId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldMappingsAPI.IntegrationLevelFieldMappingsControllerList(context.Background(), integrationSelector).Limit(limit).Cursor(cursor).Search(search).ConnectorId(connectorId).IncludeArchived(includeArchived).UniversalFieldMappingId(universalFieldMappingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldMappingsAPI.IntegrationLevelFieldMappingsControllerList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationLevelFieldMappingsControllerList`: FieldMappingsControllerList200Response
	fmt.Fprintf(os.Stdout, "Response from `FieldMappingsAPI.IntegrationLevelFieldMappingsControllerList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationSelector** | **string** | Integration ID or Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationLevelFieldMappingsControllerListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **float32** |  | 
 **cursor** | **string** |  | 
 **search** | **string** |  | 
 **connectorId** | **string** |  | 
 **includeArchived** | **bool** |  | 
 **universalFieldMappingId** | **string** |  | 

### Return type

[**FieldMappingsControllerList200Response**](FieldMappingsControllerList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

