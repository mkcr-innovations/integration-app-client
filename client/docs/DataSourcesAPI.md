# \DataSourcesAPI

All URIs are relative to *https://api.integration.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConnectionLevelDataSourceControllerArchive**](DataSourcesAPI.md#ConnectionLevelDataSourceControllerArchive) | **Delete** /connections/{connectionSelector}/data-sources/{dataSourceSelector} | Archive data source instance for connection
[**ConnectionLevelDataSourceControllerCreate**](DataSourcesAPI.md#ConnectionLevelDataSourceControllerCreate) | **Post** /connections/{connectionSelector}/data-sources/{dataSourceSelector} | Create data source instance for connection
[**ConnectionLevelDataSourceControllerGet**](DataSourcesAPI.md#ConnectionLevelDataSourceControllerGet) | **Get** /connections/{connectionSelector}/data-sources/{dataSourceSelector} | Get data source instance for connection
[**ConnectionLevelDataSourceControllerPatch**](DataSourcesAPI.md#ConnectionLevelDataSourceControllerPatch) | **Patch** /connections/{connectionSelector}/data-sources/{dataSourceSelector} | Patch update data source instance for connection
[**ConnectionLevelDataSourceControllerPut**](DataSourcesAPI.md#ConnectionLevelDataSourceControllerPut) | **Put** /connections/{connectionSelector}/data-sources/{dataSourceSelector} | Update data source instance for connection
[**ConnectionLevelDataSourceControllerReset**](DataSourcesAPI.md#ConnectionLevelDataSourceControllerReset) | **Post** /connections/{connectionSelector}/data-sources/{dataSourceSelector}/reset | Reset data source instance for connection
[**ConnectionLevelDataSourceControllerSetup**](DataSourcesAPI.md#ConnectionLevelDataSourceControllerSetup) | **Post** /connections/{connectionSelector}/data-sources/{dataSourceSelector}/setup | Setup data source instance for connection
[**ConnectionLevelDataSourcesControllerGet**](DataSourcesAPI.md#ConnectionLevelDataSourcesControllerGet) | **Get** /connections/{connectionSelector}/data-sources | List data source instances for connection
[**DataSourceByIdControllerApply**](DataSourcesAPI.md#DataSourceByIdControllerApply) | **Post** /data-sources/{id}/apply | Apply data source to integrations
[**DataSourceByIdControllerArchive**](DataSourcesAPI.md#DataSourceByIdControllerArchive) | **Delete** /data-sources/{id} | Archive data source
[**DataSourceByIdControllerClone**](DataSourcesAPI.md#DataSourceByIdControllerClone) | **Post** /data-sources/{id}/clone | 
[**DataSourceByIdControllerExport**](DataSourcesAPI.md#DataSourceByIdControllerExport) | **Get** /data-sources/{id}/export | 
[**DataSourceByIdControllerGet**](DataSourcesAPI.md#DataSourceByIdControllerGet) | **Get** /data-sources/{id} | Get data source
[**DataSourceByIdControllerPatch**](DataSourcesAPI.md#DataSourceByIdControllerPatch) | **Patch** /data-sources/{id} | Patch data source
[**DataSourceByIdControllerPut**](DataSourcesAPI.md#DataSourceByIdControllerPut) | **Put** /data-sources/{id} | Update data source
[**DataSourceByIdControllerReset**](DataSourcesAPI.md#DataSourceByIdControllerReset) | **Post** /data-sources/{id}/reset | Reset data source
[**DataSourceInstanceByIdControllerArchive**](DataSourcesAPI.md#DataSourceInstanceByIdControllerArchive) | **Delete** /data-source-instances/{id} | Archive data source instance
[**DataSourceInstanceByIdControllerCreate**](DataSourcesAPI.md#DataSourceInstanceByIdControllerCreate) | **Post** /data-source-instances/{id} | Create data source instance
[**DataSourceInstanceByIdControllerGet**](DataSourcesAPI.md#DataSourceInstanceByIdControllerGet) | **Get** /data-source-instances/{id} | Get data source instance
[**DataSourceInstanceByIdControllerPatch**](DataSourcesAPI.md#DataSourceInstanceByIdControllerPatch) | **Patch** /data-source-instances/{id} | Patch data source instance
[**DataSourceInstanceByIdControllerPut**](DataSourcesAPI.md#DataSourceInstanceByIdControllerPut) | **Put** /data-source-instances/{id} | Update data source instance
[**DataSourceInstanceByIdControllerReset**](DataSourcesAPI.md#DataSourceInstanceByIdControllerReset) | **Post** /data-source-instances/{id}/reset | Reset data source instance
[**DataSourceInstanceByIdControllerSetup**](DataSourcesAPI.md#DataSourceInstanceByIdControllerSetup) | **Post** /data-source-instances/{id}/setup | Setup data source instance
[**DataSourceInstancesControllerList**](DataSourcesAPI.md#DataSourceInstancesControllerList) | **Get** /data-source-instances | List data source instances
[**DataSourcesControllerCreate**](DataSourcesAPI.md#DataSourcesControllerCreate) | **Post** /data-sources | Create data source
[**DataSourcesControllerList**](DataSourcesAPI.md#DataSourcesControllerList) | **Get** /data-sources | List data sources
[**IntegrationLevelDataSourceControllerArchive**](DataSourcesAPI.md#IntegrationLevelDataSourceControllerArchive) | **Delete** /integrations/{integrationSelector}/data-sources/{dataSourceSelector} | Archive data source for integration
[**IntegrationLevelDataSourceControllerGet**](DataSourcesAPI.md#IntegrationLevelDataSourceControllerGet) | **Get** /integrations/{integrationSelector}/data-sources/{dataSourceSelector} | Get data source for integration
[**IntegrationLevelDataSourceControllerPatch**](DataSourcesAPI.md#IntegrationLevelDataSourceControllerPatch) | **Patch** /integrations/{integrationSelector}/data-sources/{dataSourceSelector} | Patch update data source for integration
[**IntegrationLevelDataSourceControllerPut**](DataSourcesAPI.md#IntegrationLevelDataSourceControllerPut) | **Put** /integrations/{integrationSelector}/data-sources/{dataSourceSelector} | Update data source for integration
[**IntegrationLevelDataSourceControllerReset**](DataSourcesAPI.md#IntegrationLevelDataSourceControllerReset) | **Post** /integrations/{integrationSelector}/data-sources/{dataSourceSelector}/reset | Reset data source for integration
[**IntegrationLevelDataSourcesControllerCreate**](DataSourcesAPI.md#IntegrationLevelDataSourcesControllerCreate) | **Post** /integrations/{integrationSelector}/data-sources | Create data source for integration
[**IntegrationLevelDataSourcesControllerList**](DataSourcesAPI.md#IntegrationLevelDataSourcesControllerList) | **Get** /integrations/{integrationSelector}/data-sources | List data sources for integration



## ConnectionLevelDataSourceControllerArchive

> ConnectionLevelDataSourceControllerArchive(ctx, dataSourceSelector, connectionSelector).InstanceKey(instanceKey).AutoCreate(autoCreate).Execute()

Archive data source instance for connection

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
	dataSourceSelector := "dataSourceSelector_example" // string | Data Source ID or Key
	connectionSelector := "connectionSelector_example" // string | Integration Key, Connection ID, or Integration ID
	instanceKey := "instanceKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DataSourcesAPI.ConnectionLevelDataSourceControllerArchive(context.Background(), dataSourceSelector, connectionSelector).InstanceKey(instanceKey).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.ConnectionLevelDataSourceControllerArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataSourceSelector** | **string** | Data Source ID or Key | 
**connectionSelector** | **string** | Integration Key, Connection ID, or Integration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionLevelDataSourceControllerArchiveRequest struct via the builder pattern


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


## ConnectionLevelDataSourceControllerCreate

> DataSourceInstanceDto ConnectionLevelDataSourceControllerCreate(ctx, connectionSelector, dataSourceSelector).InstanceKey(instanceKey).Execute()

Create data source instance for connection

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
	dataSourceSelector := "dataSourceSelector_example" // string | Data Source ID or Key
	instanceKey := "instanceKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataSourcesAPI.ConnectionLevelDataSourceControllerCreate(context.Background(), connectionSelector, dataSourceSelector).InstanceKey(instanceKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.ConnectionLevelDataSourceControllerCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectionLevelDataSourceControllerCreate`: DataSourceInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.ConnectionLevelDataSourceControllerCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionSelector** | **string** | Integration Key, Connection ID, or Integration ID | 
**dataSourceSelector** | **string** | Data Source ID or Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionLevelDataSourceControllerCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **instanceKey** | **string** |  | 

### Return type

[**DataSourceInstanceDto**](DataSourceInstanceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionLevelDataSourceControllerGet

> DataSourceInstanceDto ConnectionLevelDataSourceControllerGet(ctx, dataSourceSelector, connectionSelector).InstanceKey(instanceKey).AutoCreate(autoCreate).Execute()

Get data source instance for connection

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
	dataSourceSelector := "dataSourceSelector_example" // string | Data Source ID or Key
	connectionSelector := "connectionSelector_example" // string | Integration Key, Connection ID, or Integration ID
	instanceKey := "instanceKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataSourcesAPI.ConnectionLevelDataSourceControllerGet(context.Background(), dataSourceSelector, connectionSelector).InstanceKey(instanceKey).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.ConnectionLevelDataSourceControllerGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectionLevelDataSourceControllerGet`: DataSourceInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.ConnectionLevelDataSourceControllerGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataSourceSelector** | **string** | Data Source ID or Key | 
**connectionSelector** | **string** | Integration Key, Connection ID, or Integration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionLevelDataSourceControllerGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **instanceKey** | **string** |  | 
 **autoCreate** | **bool** |  | 

### Return type

[**DataSourceInstanceDto**](DataSourceInstanceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionLevelDataSourceControllerPatch

> DataSourceInstanceDto ConnectionLevelDataSourceControllerPatch(ctx, dataSourceSelector, connectionSelector).UpdateDataSourceInstanceDto(updateDataSourceInstanceDto).InstanceKey(instanceKey).AutoCreate(autoCreate).Execute()

Patch update data source instance for connection

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
	dataSourceSelector := "dataSourceSelector_example" // string | Data Source ID or Key
	connectionSelector := "connectionSelector_example" // string | Integration Key, Connection ID, or Integration ID
	updateDataSourceInstanceDto := *openapiclient.NewUpdateDataSourceInstanceDto() // UpdateDataSourceInstanceDto | 
	instanceKey := "instanceKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataSourcesAPI.ConnectionLevelDataSourceControllerPatch(context.Background(), dataSourceSelector, connectionSelector).UpdateDataSourceInstanceDto(updateDataSourceInstanceDto).InstanceKey(instanceKey).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.ConnectionLevelDataSourceControllerPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectionLevelDataSourceControllerPatch`: DataSourceInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.ConnectionLevelDataSourceControllerPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataSourceSelector** | **string** | Data Source ID or Key | 
**connectionSelector** | **string** | Integration Key, Connection ID, or Integration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionLevelDataSourceControllerPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateDataSourceInstanceDto** | [**UpdateDataSourceInstanceDto**](UpdateDataSourceInstanceDto.md) |  | 
 **instanceKey** | **string** |  | 
 **autoCreate** | **bool** |  | 

### Return type

[**DataSourceInstanceDto**](DataSourceInstanceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionLevelDataSourceControllerPut

> DataSourceInstanceDto ConnectionLevelDataSourceControllerPut(ctx, dataSourceSelector, connectionSelector).UpdateDataSourceInstanceDto(updateDataSourceInstanceDto).InstanceKey(instanceKey).AutoCreate(autoCreate).Execute()

Update data source instance for connection

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
	dataSourceSelector := "dataSourceSelector_example" // string | Data Source ID or Key
	connectionSelector := "connectionSelector_example" // string | Integration Key, Connection ID, or Integration ID
	updateDataSourceInstanceDto := *openapiclient.NewUpdateDataSourceInstanceDto() // UpdateDataSourceInstanceDto | 
	instanceKey := "instanceKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataSourcesAPI.ConnectionLevelDataSourceControllerPut(context.Background(), dataSourceSelector, connectionSelector).UpdateDataSourceInstanceDto(updateDataSourceInstanceDto).InstanceKey(instanceKey).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.ConnectionLevelDataSourceControllerPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectionLevelDataSourceControllerPut`: DataSourceInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.ConnectionLevelDataSourceControllerPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataSourceSelector** | **string** | Data Source ID or Key | 
**connectionSelector** | **string** | Integration Key, Connection ID, or Integration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionLevelDataSourceControllerPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateDataSourceInstanceDto** | [**UpdateDataSourceInstanceDto**](UpdateDataSourceInstanceDto.md) |  | 
 **instanceKey** | **string** |  | 
 **autoCreate** | **bool** |  | 

### Return type

[**DataSourceInstanceDto**](DataSourceInstanceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionLevelDataSourceControllerReset

> DataSourceInstanceDto ConnectionLevelDataSourceControllerReset(ctx, dataSourceSelector, connectionSelector).InstanceKey(instanceKey).AutoCreate(autoCreate).Execute()

Reset data source instance for connection

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
	dataSourceSelector := "dataSourceSelector_example" // string | Data Source ID or Key
	connectionSelector := "connectionSelector_example" // string | Integration Key, Connection ID, or Integration ID
	instanceKey := "instanceKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataSourcesAPI.ConnectionLevelDataSourceControllerReset(context.Background(), dataSourceSelector, connectionSelector).InstanceKey(instanceKey).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.ConnectionLevelDataSourceControllerReset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectionLevelDataSourceControllerReset`: DataSourceInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.ConnectionLevelDataSourceControllerReset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataSourceSelector** | **string** | Data Source ID or Key | 
**connectionSelector** | **string** | Integration Key, Connection ID, or Integration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionLevelDataSourceControllerResetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **instanceKey** | **string** |  | 
 **autoCreate** | **bool** |  | 

### Return type

[**DataSourceInstanceDto**](DataSourceInstanceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionLevelDataSourceControllerSetup

> DataSourceInstanceDto ConnectionLevelDataSourceControllerSetup(ctx, dataSourceSelector, connectionSelector).InstanceKey(instanceKey).AutoCreate(autoCreate).Execute()

Setup data source instance for connection

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
	dataSourceSelector := "dataSourceSelector_example" // string | Data Source ID or Key
	connectionSelector := "connectionSelector_example" // string | Integration Key, Connection ID, or Integration ID
	instanceKey := "instanceKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataSourcesAPI.ConnectionLevelDataSourceControllerSetup(context.Background(), dataSourceSelector, connectionSelector).InstanceKey(instanceKey).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.ConnectionLevelDataSourceControllerSetup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectionLevelDataSourceControllerSetup`: DataSourceInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.ConnectionLevelDataSourceControllerSetup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataSourceSelector** | **string** | Data Source ID or Key | 
**connectionSelector** | **string** | Integration Key, Connection ID, or Integration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionLevelDataSourceControllerSetupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **instanceKey** | **string** |  | 
 **autoCreate** | **bool** |  | 

### Return type

[**DataSourceInstanceDto**](DataSourceInstanceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionLevelDataSourcesControllerGet

> ConnectionLevelDataSourcesControllerGet(ctx, connectionSelector).Limit(limit).Cursor(cursor).Search(search).ConnectorId(connectorId).IncludeArchived(includeArchived).DataSourceId(dataSourceId).Execute()

List data source instances for connection

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
	dataSourceId := "dataSourceId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DataSourcesAPI.ConnectionLevelDataSourcesControllerGet(context.Background(), connectionSelector).Limit(limit).Cursor(cursor).Search(search).ConnectorId(connectorId).IncludeArchived(includeArchived).DataSourceId(dataSourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.ConnectionLevelDataSourcesControllerGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiConnectionLevelDataSourcesControllerGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **float32** |  | 
 **cursor** | **string** |  | 
 **search** | **string** |  | 
 **connectorId** | **string** |  | 
 **includeArchived** | **bool** |  | 
 **dataSourceId** | **string** |  | 

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


## DataSourceByIdControllerApply

> []DataSourceDto DataSourceByIdControllerApply(ctx, id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()

Apply data source to integrations

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
	resp, r, err := apiClient.DataSourcesAPI.DataSourceByIdControllerApply(context.Background(), id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.DataSourceByIdControllerApply``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DataSourceByIdControllerApply`: []DataSourceDto
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.DataSourceByIdControllerApply`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDataSourceByIdControllerApplyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **key** | **string** |  | 
 **integrationId** | **string** |  | 
 **integrationKey** | **string** |  | 

### Return type

[**[]DataSourceDto**](DataSourceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataSourceByIdControllerArchive

> DataSourceByIdControllerArchive(ctx, id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()

Archive data source

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
	r, err := apiClient.DataSourcesAPI.DataSourceByIdControllerArchive(context.Background(), id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.DataSourceByIdControllerArchive``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDataSourceByIdControllerArchiveRequest struct via the builder pattern


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


## DataSourceByIdControllerClone

> DataSourceDto DataSourceByIdControllerClone(ctx, id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()



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
	resp, r, err := apiClient.DataSourcesAPI.DataSourceByIdControllerClone(context.Background(), id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.DataSourceByIdControllerClone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DataSourceByIdControllerClone`: DataSourceDto
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.DataSourceByIdControllerClone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDataSourceByIdControllerCloneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **key** | **string** |  | 
 **integrationId** | **string** |  | 
 **integrationKey** | **string** |  | 

### Return type

[**DataSourceDto**](DataSourceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataSourceByIdControllerExport

> DataSourceExportDto DataSourceByIdControllerExport(ctx, id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()



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
	resp, r, err := apiClient.DataSourcesAPI.DataSourceByIdControllerExport(context.Background(), id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.DataSourceByIdControllerExport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DataSourceByIdControllerExport`: DataSourceExportDto
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.DataSourceByIdControllerExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDataSourceByIdControllerExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **key** | **string** |  | 
 **integrationId** | **string** |  | 
 **integrationKey** | **string** |  | 

### Return type

[**DataSourceExportDto**](DataSourceExportDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataSourceByIdControllerGet

> DataSourceDto DataSourceByIdControllerGet(ctx, id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()

Get data source

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
	resp, r, err := apiClient.DataSourcesAPI.DataSourceByIdControllerGet(context.Background(), id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.DataSourceByIdControllerGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DataSourceByIdControllerGet`: DataSourceDto
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.DataSourceByIdControllerGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDataSourceByIdControllerGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **key** | **string** |  | 
 **integrationId** | **string** |  | 
 **integrationKey** | **string** |  | 

### Return type

[**DataSourceDto**](DataSourceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataSourceByIdControllerPatch

> DataSourceDto DataSourceByIdControllerPatch(ctx, id).UpdateDataSourceDto(updateDataSourceDto).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()

Patch data source

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
	updateDataSourceDto := *openapiclient.NewUpdateDataSourceDto() // UpdateDataSourceDto | 
	key := "key_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataSourcesAPI.DataSourceByIdControllerPatch(context.Background(), id).UpdateDataSourceDto(updateDataSourceDto).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.DataSourceByIdControllerPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DataSourceByIdControllerPatch`: DataSourceDto
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.DataSourceByIdControllerPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDataSourceByIdControllerPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDataSourceDto** | [**UpdateDataSourceDto**](UpdateDataSourceDto.md) |  | 
 **key** | **string** |  | 
 **integrationId** | **string** |  | 
 **integrationKey** | **string** |  | 

### Return type

[**DataSourceDto**](DataSourceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataSourceByIdControllerPut

> DataSourceDto DataSourceByIdControllerPut(ctx, id).UpdateDataSourceDto(updateDataSourceDto).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()

Update data source

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
	updateDataSourceDto := *openapiclient.NewUpdateDataSourceDto() // UpdateDataSourceDto | 
	key := "key_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataSourcesAPI.DataSourceByIdControllerPut(context.Background(), id).UpdateDataSourceDto(updateDataSourceDto).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.DataSourceByIdControllerPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DataSourceByIdControllerPut`: DataSourceDto
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.DataSourceByIdControllerPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDataSourceByIdControllerPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDataSourceDto** | [**UpdateDataSourceDto**](UpdateDataSourceDto.md) |  | 
 **key** | **string** |  | 
 **integrationId** | **string** |  | 
 **integrationKey** | **string** |  | 

### Return type

[**DataSourceDto**](DataSourceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataSourceByIdControllerReset

> DataSourceDto DataSourceByIdControllerReset(ctx, id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()

Reset data source

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
	resp, r, err := apiClient.DataSourcesAPI.DataSourceByIdControllerReset(context.Background(), id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.DataSourceByIdControllerReset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DataSourceByIdControllerReset`: DataSourceDto
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.DataSourceByIdControllerReset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDataSourceByIdControllerResetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **key** | **string** |  | 
 **integrationId** | **string** |  | 
 **integrationKey** | **string** |  | 

### Return type

[**DataSourceDto**](DataSourceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataSourceInstanceByIdControllerArchive

> DataSourceInstanceByIdControllerArchive(ctx, id).DataSourceKey(dataSourceKey).DataSourceId(dataSourceId).InstanceKey(instanceKey).AutoCreate(autoCreate).FlowKey(flowKey).FlowId(flowId).NodeKey(nodeKey).Udm(udm).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()

Archive data source instance

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
	dataSourceKey := "dataSourceKey_example" // string |  (optional)
	dataSourceId := "dataSourceId_example" // string |  (optional)
	instanceKey := "instanceKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)
	flowKey := "flowKey_example" // string |  (optional)
	flowId := "flowId_example" // string |  (optional)
	nodeKey := "nodeKey_example" // string |  (optional)
	udm := "udm_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	connectionId := "connectionId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DataSourcesAPI.DataSourceInstanceByIdControllerArchive(context.Background(), id).DataSourceKey(dataSourceKey).DataSourceId(dataSourceId).InstanceKey(instanceKey).AutoCreate(autoCreate).FlowKey(flowKey).FlowId(flowId).NodeKey(nodeKey).Udm(udm).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.DataSourceInstanceByIdControllerArchive``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDataSourceInstanceByIdControllerArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dataSourceKey** | **string** |  | 
 **dataSourceId** | **string** |  | 
 **instanceKey** | **string** |  | 
 **autoCreate** | **bool** |  | 
 **flowKey** | **string** |  | 
 **flowId** | **string** |  | 
 **nodeKey** | **string** |  | 
 **udm** | **string** |  | 
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


## DataSourceInstanceByIdControllerCreate

> DataSourceInstanceDto DataSourceInstanceByIdControllerCreate(ctx, id).UpdateDataSourceInstanceDto(updateDataSourceInstanceDto).DataSourceKey(dataSourceKey).DataSourceId(dataSourceId).InstanceKey(instanceKey).AutoCreate(autoCreate).FlowKey(flowKey).FlowId(flowId).NodeKey(nodeKey).Udm(udm).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()

Create data source instance

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
	updateDataSourceInstanceDto := *openapiclient.NewUpdateDataSourceInstanceDto() // UpdateDataSourceInstanceDto | 
	dataSourceKey := "dataSourceKey_example" // string |  (optional)
	dataSourceId := "dataSourceId_example" // string |  (optional)
	instanceKey := "instanceKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)
	flowKey := "flowKey_example" // string |  (optional)
	flowId := "flowId_example" // string |  (optional)
	nodeKey := "nodeKey_example" // string |  (optional)
	udm := "udm_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	connectionId := "connectionId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataSourcesAPI.DataSourceInstanceByIdControllerCreate(context.Background(), id).UpdateDataSourceInstanceDto(updateDataSourceInstanceDto).DataSourceKey(dataSourceKey).DataSourceId(dataSourceId).InstanceKey(instanceKey).AutoCreate(autoCreate).FlowKey(flowKey).FlowId(flowId).NodeKey(nodeKey).Udm(udm).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.DataSourceInstanceByIdControllerCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DataSourceInstanceByIdControllerCreate`: DataSourceInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.DataSourceInstanceByIdControllerCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDataSourceInstanceByIdControllerCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDataSourceInstanceDto** | [**UpdateDataSourceInstanceDto**](UpdateDataSourceInstanceDto.md) |  | 
 **dataSourceKey** | **string** |  | 
 **dataSourceId** | **string** |  | 
 **instanceKey** | **string** |  | 
 **autoCreate** | **bool** |  | 
 **flowKey** | **string** |  | 
 **flowId** | **string** |  | 
 **nodeKey** | **string** |  | 
 **udm** | **string** |  | 
 **integrationKey** | **string** |  | 
 **integrationId** | **string** |  | 
 **connectionId** | **string** |  | 

### Return type

[**DataSourceInstanceDto**](DataSourceInstanceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataSourceInstanceByIdControllerGet

> DataSourceInstanceDto DataSourceInstanceByIdControllerGet(ctx, id).DataSourceKey(dataSourceKey).DataSourceId(dataSourceId).InstanceKey(instanceKey).AutoCreate(autoCreate).FlowKey(flowKey).FlowId(flowId).NodeKey(nodeKey).Udm(udm).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()

Get data source instance

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
	dataSourceKey := "dataSourceKey_example" // string |  (optional)
	dataSourceId := "dataSourceId_example" // string |  (optional)
	instanceKey := "instanceKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)
	flowKey := "flowKey_example" // string |  (optional)
	flowId := "flowId_example" // string |  (optional)
	nodeKey := "nodeKey_example" // string |  (optional)
	udm := "udm_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	connectionId := "connectionId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataSourcesAPI.DataSourceInstanceByIdControllerGet(context.Background(), id).DataSourceKey(dataSourceKey).DataSourceId(dataSourceId).InstanceKey(instanceKey).AutoCreate(autoCreate).FlowKey(flowKey).FlowId(flowId).NodeKey(nodeKey).Udm(udm).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.DataSourceInstanceByIdControllerGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DataSourceInstanceByIdControllerGet`: DataSourceInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.DataSourceInstanceByIdControllerGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDataSourceInstanceByIdControllerGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dataSourceKey** | **string** |  | 
 **dataSourceId** | **string** |  | 
 **instanceKey** | **string** |  | 
 **autoCreate** | **bool** |  | 
 **flowKey** | **string** |  | 
 **flowId** | **string** |  | 
 **nodeKey** | **string** |  | 
 **udm** | **string** |  | 
 **integrationKey** | **string** |  | 
 **integrationId** | **string** |  | 
 **connectionId** | **string** |  | 

### Return type

[**DataSourceInstanceDto**](DataSourceInstanceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataSourceInstanceByIdControllerPatch

> DataSourceInstanceDto DataSourceInstanceByIdControllerPatch(ctx, id).UpdateDataSourceInstanceDto(updateDataSourceInstanceDto).DataSourceKey(dataSourceKey).DataSourceId(dataSourceId).InstanceKey(instanceKey).AutoCreate(autoCreate).FlowKey(flowKey).FlowId(flowId).NodeKey(nodeKey).Udm(udm).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()

Patch data source instance

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
	updateDataSourceInstanceDto := *openapiclient.NewUpdateDataSourceInstanceDto() // UpdateDataSourceInstanceDto | 
	dataSourceKey := "dataSourceKey_example" // string |  (optional)
	dataSourceId := "dataSourceId_example" // string |  (optional)
	instanceKey := "instanceKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)
	flowKey := "flowKey_example" // string |  (optional)
	flowId := "flowId_example" // string |  (optional)
	nodeKey := "nodeKey_example" // string |  (optional)
	udm := "udm_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	connectionId := "connectionId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataSourcesAPI.DataSourceInstanceByIdControllerPatch(context.Background(), id).UpdateDataSourceInstanceDto(updateDataSourceInstanceDto).DataSourceKey(dataSourceKey).DataSourceId(dataSourceId).InstanceKey(instanceKey).AutoCreate(autoCreate).FlowKey(flowKey).FlowId(flowId).NodeKey(nodeKey).Udm(udm).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.DataSourceInstanceByIdControllerPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DataSourceInstanceByIdControllerPatch`: DataSourceInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.DataSourceInstanceByIdControllerPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDataSourceInstanceByIdControllerPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDataSourceInstanceDto** | [**UpdateDataSourceInstanceDto**](UpdateDataSourceInstanceDto.md) |  | 
 **dataSourceKey** | **string** |  | 
 **dataSourceId** | **string** |  | 
 **instanceKey** | **string** |  | 
 **autoCreate** | **bool** |  | 
 **flowKey** | **string** |  | 
 **flowId** | **string** |  | 
 **nodeKey** | **string** |  | 
 **udm** | **string** |  | 
 **integrationKey** | **string** |  | 
 **integrationId** | **string** |  | 
 **connectionId** | **string** |  | 

### Return type

[**DataSourceInstanceDto**](DataSourceInstanceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataSourceInstanceByIdControllerPut

> DataSourceInstanceDto DataSourceInstanceByIdControllerPut(ctx, id).UpdateDataSourceInstanceDto(updateDataSourceInstanceDto).DataSourceKey(dataSourceKey).DataSourceId(dataSourceId).InstanceKey(instanceKey).AutoCreate(autoCreate).FlowKey(flowKey).FlowId(flowId).NodeKey(nodeKey).Udm(udm).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()

Update data source instance

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
	updateDataSourceInstanceDto := *openapiclient.NewUpdateDataSourceInstanceDto() // UpdateDataSourceInstanceDto | 
	dataSourceKey := "dataSourceKey_example" // string |  (optional)
	dataSourceId := "dataSourceId_example" // string |  (optional)
	instanceKey := "instanceKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)
	flowKey := "flowKey_example" // string |  (optional)
	flowId := "flowId_example" // string |  (optional)
	nodeKey := "nodeKey_example" // string |  (optional)
	udm := "udm_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	connectionId := "connectionId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataSourcesAPI.DataSourceInstanceByIdControllerPut(context.Background(), id).UpdateDataSourceInstanceDto(updateDataSourceInstanceDto).DataSourceKey(dataSourceKey).DataSourceId(dataSourceId).InstanceKey(instanceKey).AutoCreate(autoCreate).FlowKey(flowKey).FlowId(flowId).NodeKey(nodeKey).Udm(udm).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.DataSourceInstanceByIdControllerPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DataSourceInstanceByIdControllerPut`: DataSourceInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.DataSourceInstanceByIdControllerPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDataSourceInstanceByIdControllerPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDataSourceInstanceDto** | [**UpdateDataSourceInstanceDto**](UpdateDataSourceInstanceDto.md) |  | 
 **dataSourceKey** | **string** |  | 
 **dataSourceId** | **string** |  | 
 **instanceKey** | **string** |  | 
 **autoCreate** | **bool** |  | 
 **flowKey** | **string** |  | 
 **flowId** | **string** |  | 
 **nodeKey** | **string** |  | 
 **udm** | **string** |  | 
 **integrationKey** | **string** |  | 
 **integrationId** | **string** |  | 
 **connectionId** | **string** |  | 

### Return type

[**DataSourceInstanceDto**](DataSourceInstanceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataSourceInstanceByIdControllerReset

> DataSourceInstanceDto DataSourceInstanceByIdControllerReset(ctx, id).DataSourceKey(dataSourceKey).DataSourceId(dataSourceId).InstanceKey(instanceKey).AutoCreate(autoCreate).FlowKey(flowKey).FlowId(flowId).NodeKey(nodeKey).Udm(udm).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()

Reset data source instance

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
	dataSourceKey := "dataSourceKey_example" // string |  (optional)
	dataSourceId := "dataSourceId_example" // string |  (optional)
	instanceKey := "instanceKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)
	flowKey := "flowKey_example" // string |  (optional)
	flowId := "flowId_example" // string |  (optional)
	nodeKey := "nodeKey_example" // string |  (optional)
	udm := "udm_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	connectionId := "connectionId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataSourcesAPI.DataSourceInstanceByIdControllerReset(context.Background(), id).DataSourceKey(dataSourceKey).DataSourceId(dataSourceId).InstanceKey(instanceKey).AutoCreate(autoCreate).FlowKey(flowKey).FlowId(flowId).NodeKey(nodeKey).Udm(udm).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.DataSourceInstanceByIdControllerReset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DataSourceInstanceByIdControllerReset`: DataSourceInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.DataSourceInstanceByIdControllerReset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDataSourceInstanceByIdControllerResetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dataSourceKey** | **string** |  | 
 **dataSourceId** | **string** |  | 
 **instanceKey** | **string** |  | 
 **autoCreate** | **bool** |  | 
 **flowKey** | **string** |  | 
 **flowId** | **string** |  | 
 **nodeKey** | **string** |  | 
 **udm** | **string** |  | 
 **integrationKey** | **string** |  | 
 **integrationId** | **string** |  | 
 **connectionId** | **string** |  | 

### Return type

[**DataSourceInstanceDto**](DataSourceInstanceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataSourceInstanceByIdControllerSetup

> DataSourceInstanceDto DataSourceInstanceByIdControllerSetup(ctx, id).DataSourceKey(dataSourceKey).DataSourceId(dataSourceId).InstanceKey(instanceKey).AutoCreate(autoCreate).FlowKey(flowKey).FlowId(flowId).NodeKey(nodeKey).Udm(udm).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()

Setup data source instance

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
	dataSourceKey := "dataSourceKey_example" // string |  (optional)
	dataSourceId := "dataSourceId_example" // string |  (optional)
	instanceKey := "instanceKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)
	flowKey := "flowKey_example" // string |  (optional)
	flowId := "flowId_example" // string |  (optional)
	nodeKey := "nodeKey_example" // string |  (optional)
	udm := "udm_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	connectionId := "connectionId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataSourcesAPI.DataSourceInstanceByIdControllerSetup(context.Background(), id).DataSourceKey(dataSourceKey).DataSourceId(dataSourceId).InstanceKey(instanceKey).AutoCreate(autoCreate).FlowKey(flowKey).FlowId(flowId).NodeKey(nodeKey).Udm(udm).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.DataSourceInstanceByIdControllerSetup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DataSourceInstanceByIdControllerSetup`: DataSourceInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.DataSourceInstanceByIdControllerSetup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDataSourceInstanceByIdControllerSetupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dataSourceKey** | **string** |  | 
 **dataSourceId** | **string** |  | 
 **instanceKey** | **string** |  | 
 **autoCreate** | **bool** |  | 
 **flowKey** | **string** |  | 
 **flowId** | **string** |  | 
 **nodeKey** | **string** |  | 
 **udm** | **string** |  | 
 **integrationKey** | **string** |  | 
 **integrationId** | **string** |  | 
 **connectionId** | **string** |  | 

### Return type

[**DataSourceInstanceDto**](DataSourceInstanceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataSourceInstancesControllerList

> DataSourceInstancesControllerList200Response DataSourceInstancesControllerList(ctx).Limit(limit).Cursor(cursor).Search(search).ConnectorId(connectorId).IncludeArchived(includeArchived).Id(id).UserId(userId).ConnectionId(connectionId).IntegrationKey(integrationKey).IntegrationId(integrationId).DataSourceId(dataSourceId).UniversalDataSourceId(universalDataSourceId).Udm(udm).InstanceKey(instanceKey).Execute()

List data source instances

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
	connectionId := "connectionId_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	dataSourceId := "dataSourceId_example" // string |  (optional)
	universalDataSourceId := "universalDataSourceId_example" // string |  (optional)
	udm := "udm_example" // string |  (optional)
	instanceKey := "instanceKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataSourcesAPI.DataSourceInstancesControllerList(context.Background()).Limit(limit).Cursor(cursor).Search(search).ConnectorId(connectorId).IncludeArchived(includeArchived).Id(id).UserId(userId).ConnectionId(connectionId).IntegrationKey(integrationKey).IntegrationId(integrationId).DataSourceId(dataSourceId).UniversalDataSourceId(universalDataSourceId).Udm(udm).InstanceKey(instanceKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.DataSourceInstancesControllerList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DataSourceInstancesControllerList`: DataSourceInstancesControllerList200Response
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.DataSourceInstancesControllerList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDataSourceInstancesControllerListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **float32** |  | 
 **cursor** | **string** |  | 
 **search** | **string** |  | 
 **connectorId** | **string** |  | 
 **includeArchived** | **bool** |  | 
 **id** | **string** |  | 
 **userId** | **string** |  | 
 **connectionId** | **string** |  | 
 **integrationKey** | **string** |  | 
 **integrationId** | **string** |  | 
 **dataSourceId** | **string** |  | 
 **universalDataSourceId** | **string** |  | 
 **udm** | **string** |  | 
 **instanceKey** | **string** |  | 

### Return type

[**DataSourceInstancesControllerList200Response**](DataSourceInstancesControllerList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataSourcesControllerCreate

> DataSourceDto DataSourcesControllerCreate(ctx).CreateDataSourceDto(createDataSourceDto).Execute()

Create data source

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
	createDataSourceDto := *openapiclient.NewCreateDataSourceDto("Key_example") // CreateDataSourceDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataSourcesAPI.DataSourcesControllerCreate(context.Background()).CreateDataSourceDto(createDataSourceDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.DataSourcesControllerCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DataSourcesControllerCreate`: DataSourceDto
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.DataSourcesControllerCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDataSourcesControllerCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDataSourceDto** | [**CreateDataSourceDto**](CreateDataSourceDto.md) |  | 

### Return type

[**DataSourceDto**](DataSourceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataSourcesControllerList

> DataSourcesControllerList200Response DataSourcesControllerList(ctx).Limit(limit).Cursor(cursor).Search(search).ConnectorId(connectorId).IncludeArchived(includeArchived).UniversalDataSourceId(universalDataSourceId).IntegrationId(integrationId).Execute()

List data sources

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
	universalDataSourceId := "universalDataSourceId_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataSourcesAPI.DataSourcesControllerList(context.Background()).Limit(limit).Cursor(cursor).Search(search).ConnectorId(connectorId).IncludeArchived(includeArchived).UniversalDataSourceId(universalDataSourceId).IntegrationId(integrationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.DataSourcesControllerList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DataSourcesControllerList`: DataSourcesControllerList200Response
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.DataSourcesControllerList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDataSourcesControllerListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **float32** |  | 
 **cursor** | **string** |  | 
 **search** | **string** |  | 
 **connectorId** | **string** |  | 
 **includeArchived** | **bool** |  | 
 **universalDataSourceId** | **string** |  | 
 **integrationId** | **string** |  | 

### Return type

[**DataSourcesControllerList200Response**](DataSourcesControllerList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationLevelDataSourceControllerArchive

> IntegrationLevelDataSourceControllerArchive(ctx, dataSourceSelector, integrationSelector).Execute()

Archive data source for integration

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
	dataSourceSelector := "dataSourceSelector_example" // string | Data Source ID or Key
	integrationSelector := "integrationSelector_example" // string | Integration ID or Key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DataSourcesAPI.IntegrationLevelDataSourceControllerArchive(context.Background(), dataSourceSelector, integrationSelector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.IntegrationLevelDataSourceControllerArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataSourceSelector** | **string** | Data Source ID or Key | 
**integrationSelector** | **string** | Integration ID or Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationLevelDataSourceControllerArchiveRequest struct via the builder pattern


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


## IntegrationLevelDataSourceControllerGet

> DataSourceDto IntegrationLevelDataSourceControllerGet(ctx, dataSourceSelector, integrationSelector).Execute()

Get data source for integration

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
	dataSourceSelector := "dataSourceSelector_example" // string | Data Source ID or Key
	integrationSelector := "integrationSelector_example" // string | Integration ID or Key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataSourcesAPI.IntegrationLevelDataSourceControllerGet(context.Background(), dataSourceSelector, integrationSelector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.IntegrationLevelDataSourceControllerGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationLevelDataSourceControllerGet`: DataSourceDto
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.IntegrationLevelDataSourceControllerGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataSourceSelector** | **string** | Data Source ID or Key | 
**integrationSelector** | **string** | Integration ID or Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationLevelDataSourceControllerGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DataSourceDto**](DataSourceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationLevelDataSourceControllerPatch

> DataSourceDto IntegrationLevelDataSourceControllerPatch(ctx, dataSourceSelector, integrationSelector).UpdateDataSourceDto(updateDataSourceDto).Execute()

Patch update data source for integration

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
	dataSourceSelector := "dataSourceSelector_example" // string | Data Source ID or Key
	integrationSelector := "integrationSelector_example" // string | Integration ID or Key
	updateDataSourceDto := *openapiclient.NewUpdateDataSourceDto() // UpdateDataSourceDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataSourcesAPI.IntegrationLevelDataSourceControllerPatch(context.Background(), dataSourceSelector, integrationSelector).UpdateDataSourceDto(updateDataSourceDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.IntegrationLevelDataSourceControllerPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationLevelDataSourceControllerPatch`: DataSourceDto
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.IntegrationLevelDataSourceControllerPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataSourceSelector** | **string** | Data Source ID or Key | 
**integrationSelector** | **string** | Integration ID or Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationLevelDataSourceControllerPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateDataSourceDto** | [**UpdateDataSourceDto**](UpdateDataSourceDto.md) |  | 

### Return type

[**DataSourceDto**](DataSourceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationLevelDataSourceControllerPut

> DataSourceDto IntegrationLevelDataSourceControllerPut(ctx, dataSourceSelector, integrationSelector).UpdateDataSourceDto(updateDataSourceDto).Execute()

Update data source for integration

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
	dataSourceSelector := "dataSourceSelector_example" // string | Data Source ID or Key
	integrationSelector := "integrationSelector_example" // string | Integration ID or Key
	updateDataSourceDto := *openapiclient.NewUpdateDataSourceDto() // UpdateDataSourceDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataSourcesAPI.IntegrationLevelDataSourceControllerPut(context.Background(), dataSourceSelector, integrationSelector).UpdateDataSourceDto(updateDataSourceDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.IntegrationLevelDataSourceControllerPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationLevelDataSourceControllerPut`: DataSourceDto
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.IntegrationLevelDataSourceControllerPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataSourceSelector** | **string** | Data Source ID or Key | 
**integrationSelector** | **string** | Integration ID or Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationLevelDataSourceControllerPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateDataSourceDto** | [**UpdateDataSourceDto**](UpdateDataSourceDto.md) |  | 

### Return type

[**DataSourceDto**](DataSourceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationLevelDataSourceControllerReset

> DataSourceDto IntegrationLevelDataSourceControllerReset(ctx, dataSourceSelector, integrationSelector).Execute()

Reset data source for integration

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
	dataSourceSelector := "dataSourceSelector_example" // string | Data Source ID or Key
	integrationSelector := "integrationSelector_example" // string | Integration ID or Key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataSourcesAPI.IntegrationLevelDataSourceControllerReset(context.Background(), dataSourceSelector, integrationSelector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.IntegrationLevelDataSourceControllerReset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationLevelDataSourceControllerReset`: DataSourceDto
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.IntegrationLevelDataSourceControllerReset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataSourceSelector** | **string** | Data Source ID or Key | 
**integrationSelector** | **string** | Integration ID or Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationLevelDataSourceControllerResetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DataSourceDto**](DataSourceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationLevelDataSourcesControllerCreate

> DataSourceDto IntegrationLevelDataSourcesControllerCreate(ctx, integrationSelector).CreateIntegrationLevelDataSourceDto(createIntegrationLevelDataSourceDto).Execute()

Create data source for integration

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
	createIntegrationLevelDataSourceDto := *openapiclient.NewCreateIntegrationLevelDataSourceDto("Key_example") // CreateIntegrationLevelDataSourceDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataSourcesAPI.IntegrationLevelDataSourcesControllerCreate(context.Background(), integrationSelector).CreateIntegrationLevelDataSourceDto(createIntegrationLevelDataSourceDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.IntegrationLevelDataSourcesControllerCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationLevelDataSourcesControllerCreate`: DataSourceDto
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.IntegrationLevelDataSourcesControllerCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationSelector** | **string** | Integration ID or Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationLevelDataSourcesControllerCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createIntegrationLevelDataSourceDto** | [**CreateIntegrationLevelDataSourceDto**](CreateIntegrationLevelDataSourceDto.md) |  | 

### Return type

[**DataSourceDto**](DataSourceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationLevelDataSourcesControllerList

> DataSourcesControllerList200Response IntegrationLevelDataSourcesControllerList(ctx, integrationSelector).Limit(limit).Cursor(cursor).Search(search).ConnectorId(connectorId).IncludeArchived(includeArchived).UniversalDataSourceId(universalDataSourceId).Execute()

List data sources for integration

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
	universalDataSourceId := "universalDataSourceId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataSourcesAPI.IntegrationLevelDataSourcesControllerList(context.Background(), integrationSelector).Limit(limit).Cursor(cursor).Search(search).ConnectorId(connectorId).IncludeArchived(includeArchived).UniversalDataSourceId(universalDataSourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.IntegrationLevelDataSourcesControllerList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationLevelDataSourcesControllerList`: DataSourcesControllerList200Response
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.IntegrationLevelDataSourcesControllerList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationSelector** | **string** | Integration ID or Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationLevelDataSourcesControllerListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **float32** |  | 
 **cursor** | **string** |  | 
 **search** | **string** |  | 
 **connectorId** | **string** |  | 
 **includeArchived** | **bool** |  | 
 **universalDataSourceId** | **string** |  | 

### Return type

[**DataSourcesControllerList200Response**](DataSourcesControllerList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

