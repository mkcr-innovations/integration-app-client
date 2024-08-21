# \DataLinksAPI

All URIs are relative to *https://api.integration.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DataLinkTableInstanceControllerArchiveDataLinkTableInstance**](DataLinksAPI.md#DataLinkTableInstanceControllerArchiveDataLinkTableInstance) | **Delete** /data-link-table-instance | Archive data link table instance
[**DataLinkTableInstanceControllerCreateDataLinkTableInstanceLinks**](DataLinksAPI.md#DataLinkTableInstanceControllerCreateDataLinkTableInstanceLinks) | **Post** /data-link-table-instance/links | Create data link table instance link
[**DataLinkTableInstanceControllerDeleteDataLinkTableInstanceLinks**](DataLinksAPI.md#DataLinkTableInstanceControllerDeleteDataLinkTableInstanceLinks) | **Delete** /data-link-table-instance/links | Delete data link table instance link
[**DataLinkTableInstanceControllerGetDataLinkTableInstance**](DataLinksAPI.md#DataLinkTableInstanceControllerGetDataLinkTableInstance) | **Get** /data-link-table-instance | Get data link table instance
[**DataLinkTableInstanceControllerGetDataLinkTableInstanceLinks**](DataLinksAPI.md#DataLinkTableInstanceControllerGetDataLinkTableInstanceLinks) | **Get** /data-link-table-instance/links | Get data link table instance links
[**DataLinkTableInstanceControllerPatchDataLinkTableInstance**](DataLinksAPI.md#DataLinkTableInstanceControllerPatchDataLinkTableInstance) | **Patch** /data-link-table-instance | Path data link table instance
[**DataLinkTableInstanceControllerPostDataLinkTableInstance**](DataLinksAPI.md#DataLinkTableInstanceControllerPostDataLinkTableInstance) | **Post** /data-link-table-instance | Create data link table instance
[**DataLinkTableInstanceControllerPutDataLinkTableInstance**](DataLinksAPI.md#DataLinkTableInstanceControllerPutDataLinkTableInstance) | **Put** /data-link-table-instance | Update data link table instance
[**DataLinkTableInstancesControllerListDataLinkTableInstances**](DataLinksAPI.md#DataLinkTableInstancesControllerListDataLinkTableInstances) | **Get** /data-link-table-instances | List data link table instances
[**DataLinkTablesControllerCreateDataLinkTable**](DataLinksAPI.md#DataLinkTablesControllerCreateDataLinkTable) | **Post** /data-link-tables | Create data link table
[**DataLinkTablesControllerDeleteDataLinkTable**](DataLinksAPI.md#DataLinkTablesControllerDeleteDataLinkTable) | **Delete** /data-link-tables/{idOrKey} | Archive data link table
[**DataLinkTablesControllerGetDataLinkTable**](DataLinksAPI.md#DataLinkTablesControllerGetDataLinkTable) | **Get** /data-link-tables/{idOrKey} | Get data link table
[**DataLinkTablesControllerListDataLinkTables**](DataLinksAPI.md#DataLinkTablesControllerListDataLinkTables) | **Get** /data-link-tables | List data link tables
[**DataLinkTablesControllerPatchDataLinkTable**](DataLinksAPI.md#DataLinkTablesControllerPatchDataLinkTable) | **Patch** /data-link-tables/{idOrKey} | Patch data link table
[**DataLinkTablesControllerPutDataLinkTable**](DataLinksAPI.md#DataLinkTablesControllerPutDataLinkTable) | **Put** /data-link-tables/{idOrKey} | Update data link table
[**DataLinksControllerDeleteDataLinkDto**](DataLinksAPI.md#DataLinksControllerDeleteDataLinkDto) | **Delete** /data-links/{dataLinkId} | Delete data link
[**DataLinksControllerListDataLinks**](DataLinksAPI.md#DataLinksControllerListDataLinks) | **Get** /data-links | List data links



## DataLinkTableInstanceControllerArchiveDataLinkTableInstance

> DataLinkTableInstanceControllerArchiveDataLinkTableInstance(ctx).Id(id).InstanceKey(instanceKey).DataLinkTableId(dataLinkTableId).DataLinkTableKey(dataLinkTableKey).AutoCreate(autoCreate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()

Archive data link table instance

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
	dataLinkTableId := "dataLinkTableId_example" // string |  (optional)
	dataLinkTableKey := "dataLinkTableKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	connectionId := "connectionId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DataLinksAPI.DataLinkTableInstanceControllerArchiveDataLinkTableInstance(context.Background()).Id(id).InstanceKey(instanceKey).DataLinkTableId(dataLinkTableId).DataLinkTableKey(dataLinkTableKey).AutoCreate(autoCreate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataLinksAPI.DataLinkTableInstanceControllerArchiveDataLinkTableInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDataLinkTableInstanceControllerArchiveDataLinkTableInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **instanceKey** | **string** |  | 
 **dataLinkTableId** | **string** |  | 
 **dataLinkTableKey** | **string** |  | 
 **autoCreate** | **bool** |  | 
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


## DataLinkTableInstanceControllerCreateDataLinkTableInstanceLinks

> DataLinkTableInstanceControllerCreateDataLinkTableInstanceLinks(ctx).CreateDataLinkDto(createDataLinkDto).Id(id).InstanceKey(instanceKey).DataLinkTableId(dataLinkTableId).DataLinkTableKey(dataLinkTableKey).AutoCreate(autoCreate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()

Create data link table instance link

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
	createDataLinkDto := *openapiclient.NewCreateDataLinkDto() // CreateDataLinkDto | 
	id := "id_example" // string |  (optional)
	instanceKey := "instanceKey_example" // string |  (optional)
	dataLinkTableId := "dataLinkTableId_example" // string |  (optional)
	dataLinkTableKey := "dataLinkTableKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	connectionId := "connectionId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DataLinksAPI.DataLinkTableInstanceControllerCreateDataLinkTableInstanceLinks(context.Background()).CreateDataLinkDto(createDataLinkDto).Id(id).InstanceKey(instanceKey).DataLinkTableId(dataLinkTableId).DataLinkTableKey(dataLinkTableKey).AutoCreate(autoCreate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataLinksAPI.DataLinkTableInstanceControllerCreateDataLinkTableInstanceLinks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDataLinkTableInstanceControllerCreateDataLinkTableInstanceLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDataLinkDto** | [**CreateDataLinkDto**](CreateDataLinkDto.md) |  | 
 **id** | **string** |  | 
 **instanceKey** | **string** |  | 
 **dataLinkTableId** | **string** |  | 
 **dataLinkTableKey** | **string** |  | 
 **autoCreate** | **bool** |  | 
 **integrationKey** | **string** |  | 
 **integrationId** | **string** |  | 
 **connectionId** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataLinkTableInstanceControllerDeleteDataLinkTableInstanceLinks

> DataLinkTableInstanceControllerDeleteDataLinkTableInstanceLinks(ctx).DeleteDataLinkDto(deleteDataLinkDto).Id(id).InstanceKey(instanceKey).DataLinkTableId(dataLinkTableId).DataLinkTableKey(dataLinkTableKey).AutoCreate(autoCreate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()

Delete data link table instance link

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
	deleteDataLinkDto := *openapiclient.NewDeleteDataLinkDto() // DeleteDataLinkDto | 
	id := "id_example" // string |  (optional)
	instanceKey := "instanceKey_example" // string |  (optional)
	dataLinkTableId := "dataLinkTableId_example" // string |  (optional)
	dataLinkTableKey := "dataLinkTableKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	connectionId := "connectionId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DataLinksAPI.DataLinkTableInstanceControllerDeleteDataLinkTableInstanceLinks(context.Background()).DeleteDataLinkDto(deleteDataLinkDto).Id(id).InstanceKey(instanceKey).DataLinkTableId(dataLinkTableId).DataLinkTableKey(dataLinkTableKey).AutoCreate(autoCreate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataLinksAPI.DataLinkTableInstanceControllerDeleteDataLinkTableInstanceLinks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDataLinkTableInstanceControllerDeleteDataLinkTableInstanceLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteDataLinkDto** | [**DeleteDataLinkDto**](DeleteDataLinkDto.md) |  | 
 **id** | **string** |  | 
 **instanceKey** | **string** |  | 
 **dataLinkTableId** | **string** |  | 
 **dataLinkTableKey** | **string** |  | 
 **autoCreate** | **bool** |  | 
 **integrationKey** | **string** |  | 
 **integrationId** | **string** |  | 
 **connectionId** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataLinkTableInstanceControllerGetDataLinkTableInstance

> DataLinkTableInstanceDto DataLinkTableInstanceControllerGetDataLinkTableInstance(ctx).Id(id).InstanceKey(instanceKey).DataLinkTableId(dataLinkTableId).DataLinkTableKey(dataLinkTableKey).AutoCreate(autoCreate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()

Get data link table instance

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
	dataLinkTableId := "dataLinkTableId_example" // string |  (optional)
	dataLinkTableKey := "dataLinkTableKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	connectionId := "connectionId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataLinksAPI.DataLinkTableInstanceControllerGetDataLinkTableInstance(context.Background()).Id(id).InstanceKey(instanceKey).DataLinkTableId(dataLinkTableId).DataLinkTableKey(dataLinkTableKey).AutoCreate(autoCreate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataLinksAPI.DataLinkTableInstanceControllerGetDataLinkTableInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DataLinkTableInstanceControllerGetDataLinkTableInstance`: DataLinkTableInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `DataLinksAPI.DataLinkTableInstanceControllerGetDataLinkTableInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDataLinkTableInstanceControllerGetDataLinkTableInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **instanceKey** | **string** |  | 
 **dataLinkTableId** | **string** |  | 
 **dataLinkTableKey** | **string** |  | 
 **autoCreate** | **bool** |  | 
 **integrationKey** | **string** |  | 
 **integrationId** | **string** |  | 
 **connectionId** | **string** |  | 

### Return type

[**DataLinkTableInstanceDto**](DataLinkTableInstanceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataLinkTableInstanceControllerGetDataLinkTableInstanceLinks

> DataLinkTableInstanceControllerGetDataLinkTableInstanceLinks(ctx).Id(id).InstanceKey(instanceKey).DataLinkTableId(dataLinkTableId).DataLinkTableKey(dataLinkTableKey).AutoCreate(autoCreate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Direction(direction).AppRecordId(appRecordId).ExternalRecordId(externalRecordId).Limit(limit).Cursor(cursor).Execute()

Get data link table instance links

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
	dataLinkTableId := "dataLinkTableId_example" // string |  (optional)
	dataLinkTableKey := "dataLinkTableKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	connectionId := "connectionId_example" // string |  (optional)
	direction := "direction_example" // string |  (optional)
	appRecordId := "appRecordId_example" // string |  (optional)
	externalRecordId := "externalRecordId_example" // string |  (optional)
	limit := float32(8.14) // float32 |  (optional)
	cursor := "cursor_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DataLinksAPI.DataLinkTableInstanceControllerGetDataLinkTableInstanceLinks(context.Background()).Id(id).InstanceKey(instanceKey).DataLinkTableId(dataLinkTableId).DataLinkTableKey(dataLinkTableKey).AutoCreate(autoCreate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Direction(direction).AppRecordId(appRecordId).ExternalRecordId(externalRecordId).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataLinksAPI.DataLinkTableInstanceControllerGetDataLinkTableInstanceLinks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDataLinkTableInstanceControllerGetDataLinkTableInstanceLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **instanceKey** | **string** |  | 
 **dataLinkTableId** | **string** |  | 
 **dataLinkTableKey** | **string** |  | 
 **autoCreate** | **bool** |  | 
 **integrationKey** | **string** |  | 
 **integrationId** | **string** |  | 
 **connectionId** | **string** |  | 
 **direction** | **string** |  | 
 **appRecordId** | **string** |  | 
 **externalRecordId** | **string** |  | 
 **limit** | **float32** |  | 
 **cursor** | **string** |  | 

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


## DataLinkTableInstanceControllerPatchDataLinkTableInstance

> DataLinkTableInstanceDto DataLinkTableInstanceControllerPatchDataLinkTableInstance(ctx).Body(body).Id(id).InstanceKey(instanceKey).DataLinkTableId(dataLinkTableId).DataLinkTableKey(dataLinkTableKey).AutoCreate(autoCreate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()

Path data link table instance

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
	dataLinkTableId := "dataLinkTableId_example" // string |  (optional)
	dataLinkTableKey := "dataLinkTableKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	connectionId := "connectionId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataLinksAPI.DataLinkTableInstanceControllerPatchDataLinkTableInstance(context.Background()).Body(body).Id(id).InstanceKey(instanceKey).DataLinkTableId(dataLinkTableId).DataLinkTableKey(dataLinkTableKey).AutoCreate(autoCreate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataLinksAPI.DataLinkTableInstanceControllerPatchDataLinkTableInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DataLinkTableInstanceControllerPatchDataLinkTableInstance`: DataLinkTableInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `DataLinksAPI.DataLinkTableInstanceControllerPatchDataLinkTableInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDataLinkTableInstanceControllerPatchDataLinkTableInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 
 **id** | **string** |  | 
 **instanceKey** | **string** |  | 
 **dataLinkTableId** | **string** |  | 
 **dataLinkTableKey** | **string** |  | 
 **autoCreate** | **bool** |  | 
 **integrationKey** | **string** |  | 
 **integrationId** | **string** |  | 
 **connectionId** | **string** |  | 

### Return type

[**DataLinkTableInstanceDto**](DataLinkTableInstanceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataLinkTableInstanceControllerPostDataLinkTableInstance

> DataLinkTableInstanceDto DataLinkTableInstanceControllerPostDataLinkTableInstance(ctx).Body(body).Id(id).InstanceKey(instanceKey).DataLinkTableId(dataLinkTableId).DataLinkTableKey(dataLinkTableKey).AutoCreate(autoCreate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()

Create data link table instance

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
	dataLinkTableId := "dataLinkTableId_example" // string |  (optional)
	dataLinkTableKey := "dataLinkTableKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	connectionId := "connectionId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataLinksAPI.DataLinkTableInstanceControllerPostDataLinkTableInstance(context.Background()).Body(body).Id(id).InstanceKey(instanceKey).DataLinkTableId(dataLinkTableId).DataLinkTableKey(dataLinkTableKey).AutoCreate(autoCreate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataLinksAPI.DataLinkTableInstanceControllerPostDataLinkTableInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DataLinkTableInstanceControllerPostDataLinkTableInstance`: DataLinkTableInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `DataLinksAPI.DataLinkTableInstanceControllerPostDataLinkTableInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDataLinkTableInstanceControllerPostDataLinkTableInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 
 **id** | **string** |  | 
 **instanceKey** | **string** |  | 
 **dataLinkTableId** | **string** |  | 
 **dataLinkTableKey** | **string** |  | 
 **autoCreate** | **bool** |  | 
 **integrationKey** | **string** |  | 
 **integrationId** | **string** |  | 
 **connectionId** | **string** |  | 

### Return type

[**DataLinkTableInstanceDto**](DataLinkTableInstanceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataLinkTableInstanceControllerPutDataLinkTableInstance

> DataLinkTableInstanceDto DataLinkTableInstanceControllerPutDataLinkTableInstance(ctx).Body(body).Id(id).InstanceKey(instanceKey).DataLinkTableId(dataLinkTableId).DataLinkTableKey(dataLinkTableKey).AutoCreate(autoCreate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()

Update data link table instance

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
	dataLinkTableId := "dataLinkTableId_example" // string |  (optional)
	dataLinkTableKey := "dataLinkTableKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	connectionId := "connectionId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataLinksAPI.DataLinkTableInstanceControllerPutDataLinkTableInstance(context.Background()).Body(body).Id(id).InstanceKey(instanceKey).DataLinkTableId(dataLinkTableId).DataLinkTableKey(dataLinkTableKey).AutoCreate(autoCreate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataLinksAPI.DataLinkTableInstanceControllerPutDataLinkTableInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DataLinkTableInstanceControllerPutDataLinkTableInstance`: DataLinkTableInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `DataLinksAPI.DataLinkTableInstanceControllerPutDataLinkTableInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDataLinkTableInstanceControllerPutDataLinkTableInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 
 **id** | **string** |  | 
 **instanceKey** | **string** |  | 
 **dataLinkTableId** | **string** |  | 
 **dataLinkTableKey** | **string** |  | 
 **autoCreate** | **bool** |  | 
 **integrationKey** | **string** |  | 
 **integrationId** | **string** |  | 
 **connectionId** | **string** |  | 

### Return type

[**DataLinkTableInstanceDto**](DataLinkTableInstanceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataLinkTableInstancesControllerListDataLinkTableInstances

> DataLinkTableInstancesControllerListDataLinkTableInstances200Response DataLinkTableInstancesControllerListDataLinkTableInstances(ctx).Limit(limit).Cursor(cursor).Id(id).UserId(userId).DataLinkTableId(dataLinkTableId).ConnectionId(connectionId).IntegrationId(integrationId).Execute()

List data link table instances

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
	id := "id_example" // string |  (optional)
	userId := "userId_example" // string |  (optional)
	dataLinkTableId := "dataLinkTableId_example" // string |  (optional)
	connectionId := "connectionId_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataLinksAPI.DataLinkTableInstancesControllerListDataLinkTableInstances(context.Background()).Limit(limit).Cursor(cursor).Id(id).UserId(userId).DataLinkTableId(dataLinkTableId).ConnectionId(connectionId).IntegrationId(integrationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataLinksAPI.DataLinkTableInstancesControllerListDataLinkTableInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DataLinkTableInstancesControllerListDataLinkTableInstances`: DataLinkTableInstancesControllerListDataLinkTableInstances200Response
	fmt.Fprintf(os.Stdout, "Response from `DataLinksAPI.DataLinkTableInstancesControllerListDataLinkTableInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDataLinkTableInstancesControllerListDataLinkTableInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **float32** |  | 
 **cursor** | **string** |  | 
 **id** | **string** |  | 
 **userId** | **string** |  | 
 **dataLinkTableId** | **string** |  | 
 **connectionId** | **string** |  | 
 **integrationId** | **string** |  | 

### Return type

[**DataLinkTableInstancesControllerListDataLinkTableInstances200Response**](DataLinkTableInstancesControllerListDataLinkTableInstances200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataLinkTablesControllerCreateDataLinkTable

> DataLinkTableDto DataLinkTablesControllerCreateDataLinkTable(ctx).CreateDataLinkTableDto(createDataLinkTableDto).Execute()

Create data link table

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
	createDataLinkTableDto := *openapiclient.NewCreateDataLinkTableDto("Key_example") // CreateDataLinkTableDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataLinksAPI.DataLinkTablesControllerCreateDataLinkTable(context.Background()).CreateDataLinkTableDto(createDataLinkTableDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataLinksAPI.DataLinkTablesControllerCreateDataLinkTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DataLinkTablesControllerCreateDataLinkTable`: DataLinkTableDto
	fmt.Fprintf(os.Stdout, "Response from `DataLinksAPI.DataLinkTablesControllerCreateDataLinkTable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDataLinkTablesControllerCreateDataLinkTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDataLinkTableDto** | [**CreateDataLinkTableDto**](CreateDataLinkTableDto.md) |  | 

### Return type

[**DataLinkTableDto**](DataLinkTableDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataLinkTablesControllerDeleteDataLinkTable

> DataLinkTablesControllerDeleteDataLinkTable(ctx, idOrKey).Execute()

Archive data link table

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
	r, err := apiClient.DataLinksAPI.DataLinkTablesControllerDeleteDataLinkTable(context.Background(), idOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataLinksAPI.DataLinkTablesControllerDeleteDataLinkTable``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDataLinkTablesControllerDeleteDataLinkTableRequest struct via the builder pattern


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


## DataLinkTablesControllerGetDataLinkTable

> DataLinkTableDto DataLinkTablesControllerGetDataLinkTable(ctx, idOrKey).Execute()

Get data link table

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
	resp, r, err := apiClient.DataLinksAPI.DataLinkTablesControllerGetDataLinkTable(context.Background(), idOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataLinksAPI.DataLinkTablesControllerGetDataLinkTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DataLinkTablesControllerGetDataLinkTable`: DataLinkTableDto
	fmt.Fprintf(os.Stdout, "Response from `DataLinksAPI.DataLinkTablesControllerGetDataLinkTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDataLinkTablesControllerGetDataLinkTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DataLinkTableDto**](DataLinkTableDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataLinkTablesControllerListDataLinkTables

> DataLinkTablesControllerListDataLinkTables200Response DataLinkTablesControllerListDataLinkTables(ctx).Limit(limit).Cursor(cursor).Search(search).ConnectorId(connectorId).Execute()

List data link tables

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
	resp, r, err := apiClient.DataLinksAPI.DataLinkTablesControllerListDataLinkTables(context.Background()).Limit(limit).Cursor(cursor).Search(search).ConnectorId(connectorId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataLinksAPI.DataLinkTablesControllerListDataLinkTables``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DataLinkTablesControllerListDataLinkTables`: DataLinkTablesControllerListDataLinkTables200Response
	fmt.Fprintf(os.Stdout, "Response from `DataLinksAPI.DataLinkTablesControllerListDataLinkTables`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDataLinkTablesControllerListDataLinkTablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **float32** |  | 
 **cursor** | **string** |  | 
 **search** | **string** |  | 
 **connectorId** | **string** |  | 

### Return type

[**DataLinkTablesControllerListDataLinkTables200Response**](DataLinkTablesControllerListDataLinkTables200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataLinkTablesControllerPatchDataLinkTable

> DataLinkTableDto DataLinkTablesControllerPatchDataLinkTable(ctx, idOrKey).UpdateDataLinkTableDto(updateDataLinkTableDto).Execute()

Patch data link table

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
	updateDataLinkTableDto := *openapiclient.NewUpdateDataLinkTableDto() // UpdateDataLinkTableDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataLinksAPI.DataLinkTablesControllerPatchDataLinkTable(context.Background(), idOrKey).UpdateDataLinkTableDto(updateDataLinkTableDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataLinksAPI.DataLinkTablesControllerPatchDataLinkTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DataLinkTablesControllerPatchDataLinkTable`: DataLinkTableDto
	fmt.Fprintf(os.Stdout, "Response from `DataLinksAPI.DataLinkTablesControllerPatchDataLinkTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDataLinkTablesControllerPatchDataLinkTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDataLinkTableDto** | [**UpdateDataLinkTableDto**](UpdateDataLinkTableDto.md) |  | 

### Return type

[**DataLinkTableDto**](DataLinkTableDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataLinkTablesControllerPutDataLinkTable

> DataLinkTableDto DataLinkTablesControllerPutDataLinkTable(ctx, idOrKey).CreateDataLinkTableDto(createDataLinkTableDto).Execute()

Update data link table

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
	createDataLinkTableDto := *openapiclient.NewCreateDataLinkTableDto("Key_example") // CreateDataLinkTableDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataLinksAPI.DataLinkTablesControllerPutDataLinkTable(context.Background(), idOrKey).CreateDataLinkTableDto(createDataLinkTableDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataLinksAPI.DataLinkTablesControllerPutDataLinkTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DataLinkTablesControllerPutDataLinkTable`: DataLinkTableDto
	fmt.Fprintf(os.Stdout, "Response from `DataLinksAPI.DataLinkTablesControllerPutDataLinkTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDataLinkTablesControllerPutDataLinkTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDataLinkTableDto** | [**CreateDataLinkTableDto**](CreateDataLinkTableDto.md) |  | 

### Return type

[**DataLinkTableDto**](DataLinkTableDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataLinksControllerDeleteDataLinkDto

> DataLinksControllerDeleteDataLinkDto(ctx, dataLinkId).Execute()

Delete data link

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
	dataLinkId := "dataLinkId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DataLinksAPI.DataLinksControllerDeleteDataLinkDto(context.Background(), dataLinkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataLinksAPI.DataLinksControllerDeleteDataLinkDto``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataLinkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDataLinksControllerDeleteDataLinkDtoRequest struct via the builder pattern


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


## DataLinksControllerListDataLinks

> DataLinksControllerListDataLinks200Response DataLinksControllerListDataLinks(ctx).DataLinkTableInstanceId(dataLinkTableInstanceId).Limit(limit).Cursor(cursor).ExternalRecordId(externalRecordId).AppRecordId(appRecordId).Direction(direction).Execute()

List data links

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
	dataLinkTableInstanceId := "dataLinkTableInstanceId_example" // string | 
	limit := float32(8.14) // float32 |  (optional)
	cursor := "cursor_example" // string |  (optional)
	externalRecordId := "externalRecordId_example" // string |  (optional)
	appRecordId := "appRecordId_example" // string |  (optional)
	direction := "direction_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataLinksAPI.DataLinksControllerListDataLinks(context.Background()).DataLinkTableInstanceId(dataLinkTableInstanceId).Limit(limit).Cursor(cursor).ExternalRecordId(externalRecordId).AppRecordId(appRecordId).Direction(direction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataLinksAPI.DataLinksControllerListDataLinks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DataLinksControllerListDataLinks`: DataLinksControllerListDataLinks200Response
	fmt.Fprintf(os.Stdout, "Response from `DataLinksAPI.DataLinksControllerListDataLinks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDataLinksControllerListDataLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataLinkTableInstanceId** | **string** |  | 
 **limit** | **float32** |  | 
 **cursor** | **string** |  | 
 **externalRecordId** | **string** |  | 
 **appRecordId** | **string** |  | 
 **direction** | **string** |  | 

### Return type

[**DataLinksControllerListDataLinks200Response**](DataLinksControllerListDataLinks200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

