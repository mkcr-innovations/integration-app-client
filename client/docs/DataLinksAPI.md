# \DataLinksAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveDataLinkTableInstance**](DataLinksAPI.md#ArchiveDataLinkTableInstance) | **Delete** /data-link-table-instance | 
[**CreateDataLinkTable**](DataLinksAPI.md#CreateDataLinkTable) | **Post** /data-link-tables | 
[**CreateDataLinkTableInstanceLinks**](DataLinksAPI.md#CreateDataLinkTableInstanceLinks) | **Post** /data-link-table-instance/links | 
[**DeleteDataLinkDto**](DataLinksAPI.md#DeleteDataLinkDto) | **Delete** /data-links/{dataLinkId} | 
[**DeleteDataLinkTable**](DataLinksAPI.md#DeleteDataLinkTable) | **Delete** /data-link-tables/{idOrKey} | 
[**DeleteDataLinkTableInstanceLinks**](DataLinksAPI.md#DeleteDataLinkTableInstanceLinks) | **Delete** /data-link-table-instance/links | 
[**GetDataLinkTable**](DataLinksAPI.md#GetDataLinkTable) | **Get** /data-link-tables/{idOrKey} | 
[**GetDataLinkTableInstance**](DataLinksAPI.md#GetDataLinkTableInstance) | **Get** /data-link-table-instance | 
[**GetDataLinkTableInstanceLinks**](DataLinksAPI.md#GetDataLinkTableInstanceLinks) | **Get** /data-link-table-instance/links | 
[**ListDataLinkTableInstances**](DataLinksAPI.md#ListDataLinkTableInstances) | **Get** /data-link-table-instances | 
[**ListDataLinkTables**](DataLinksAPI.md#ListDataLinkTables) | **Get** /data-link-tables | 
[**ListDataLinks**](DataLinksAPI.md#ListDataLinks) | **Get** /data-links | 
[**PatchDataLinkTable**](DataLinksAPI.md#PatchDataLinkTable) | **Patch** /data-link-tables/{idOrKey} | 
[**PatchDataLinkTableInstance**](DataLinksAPI.md#PatchDataLinkTableInstance) | **Patch** /data-link-table-instance | 
[**PostDataLinkTableInstance**](DataLinksAPI.md#PostDataLinkTableInstance) | **Post** /data-link-table-instance | 
[**PutDataLinkTable**](DataLinksAPI.md#PutDataLinkTable) | **Put** /data-link-tables/{idOrKey} | 
[**PutDataLinkTableInstance**](DataLinksAPI.md#PutDataLinkTableInstance) | **Put** /data-link-table-instance | 



## ArchiveDataLinkTableInstance

> ArchiveDataLinkTableInstance(ctx).Id(id).InstanceKey(instanceKey).DataLinkTableId(dataLinkTableId).DataLinkTableKey(dataLinkTableKey).AutoCreate(autoCreate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()



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
	r, err := apiClient.DataLinksAPI.ArchiveDataLinkTableInstance(context.Background()).Id(id).InstanceKey(instanceKey).DataLinkTableId(dataLinkTableId).DataLinkTableKey(dataLinkTableKey).AutoCreate(autoCreate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataLinksAPI.ArchiveDataLinkTableInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArchiveDataLinkTableInstanceRequest struct via the builder pattern


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


## CreateDataLinkTable

> DataLinkTableDto CreateDataLinkTable(ctx).CreateDataLinkTableDto(createDataLinkTableDto).Execute()



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
	createDataLinkTableDto := *openapiclient.NewCreateDataLinkTableDto("Key_example", "Name_example") // CreateDataLinkTableDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataLinksAPI.CreateDataLinkTable(context.Background()).CreateDataLinkTableDto(createDataLinkTableDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataLinksAPI.CreateDataLinkTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDataLinkTable`: DataLinkTableDto
	fmt.Fprintf(os.Stdout, "Response from `DataLinksAPI.CreateDataLinkTable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDataLinkTableRequest struct via the builder pattern


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


## CreateDataLinkTableInstanceLinks

> CreateDataLinkTableInstanceLinks(ctx).CreateDataLinkDto(createDataLinkDto).Id(id).InstanceKey(instanceKey).DataLinkTableId(dataLinkTableId).DataLinkTableKey(dataLinkTableKey).AutoCreate(autoCreate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()



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
	createDataLinkDto := *openapiclient.NewCreateDataLinkDto("Direction_example", "AppRecordId_example", "ExternalRecordId_example") // CreateDataLinkDto | 
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
	r, err := apiClient.DataLinksAPI.CreateDataLinkTableInstanceLinks(context.Background()).CreateDataLinkDto(createDataLinkDto).Id(id).InstanceKey(instanceKey).DataLinkTableId(dataLinkTableId).DataLinkTableKey(dataLinkTableKey).AutoCreate(autoCreate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataLinksAPI.CreateDataLinkTableInstanceLinks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDataLinkTableInstanceLinksRequest struct via the builder pattern


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


## DeleteDataLinkDto

> DeleteDataLinkDto(ctx, dataLinkId).Execute()



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
	r, err := apiClient.DataLinksAPI.DeleteDataLinkDto(context.Background(), dataLinkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataLinksAPI.DeleteDataLinkDto``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteDataLinkDtoRequest struct via the builder pattern


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


## DeleteDataLinkTable

> DeleteDataLinkTable(ctx, idOrKey).Execute()



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
	r, err := apiClient.DataLinksAPI.DeleteDataLinkTable(context.Background(), idOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataLinksAPI.DeleteDataLinkTable``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteDataLinkTableRequest struct via the builder pattern


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


## DeleteDataLinkTableInstanceLinks

> DeleteDataLinkTableInstanceLinks(ctx).DeleteDataLinkDto(deleteDataLinkDto).Id(id).InstanceKey(instanceKey).DataLinkTableId(dataLinkTableId).DataLinkTableKey(dataLinkTableKey).AutoCreate(autoCreate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()



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
	deleteDataLinkDto := *openapiclient.NewDeleteDataLinkDto("Direction_example", "AppRecordId_example", "ExternalRecordId_example") // DeleteDataLinkDto | 
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
	r, err := apiClient.DataLinksAPI.DeleteDataLinkTableInstanceLinks(context.Background()).DeleteDataLinkDto(deleteDataLinkDto).Id(id).InstanceKey(instanceKey).DataLinkTableId(dataLinkTableId).DataLinkTableKey(dataLinkTableKey).AutoCreate(autoCreate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataLinksAPI.DeleteDataLinkTableInstanceLinks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDataLinkTableInstanceLinksRequest struct via the builder pattern


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


## GetDataLinkTable

> DataLinkTableDto GetDataLinkTable(ctx, idOrKey).Execute()



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
	resp, r, err := apiClient.DataLinksAPI.GetDataLinkTable(context.Background(), idOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataLinksAPI.GetDataLinkTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDataLinkTable`: DataLinkTableDto
	fmt.Fprintf(os.Stdout, "Response from `DataLinksAPI.GetDataLinkTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataLinkTableRequest struct via the builder pattern


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


## GetDataLinkTableInstance

> DataLinkTableInstanceDto GetDataLinkTableInstance(ctx).Id(id).InstanceKey(instanceKey).DataLinkTableId(dataLinkTableId).DataLinkTableKey(dataLinkTableKey).AutoCreate(autoCreate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()



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
	resp, r, err := apiClient.DataLinksAPI.GetDataLinkTableInstance(context.Background()).Id(id).InstanceKey(instanceKey).DataLinkTableId(dataLinkTableId).DataLinkTableKey(dataLinkTableKey).AutoCreate(autoCreate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataLinksAPI.GetDataLinkTableInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDataLinkTableInstance`: DataLinkTableInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `DataLinksAPI.GetDataLinkTableInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDataLinkTableInstanceRequest struct via the builder pattern


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


## GetDataLinkTableInstanceLinks

> GetDataLinkTableInstanceLinks(ctx).Id(id).InstanceKey(instanceKey).DataLinkTableId(dataLinkTableId).DataLinkTableKey(dataLinkTableKey).AutoCreate(autoCreate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Direction(direction).AppRecordId(appRecordId).ExternalRecordId(externalRecordId).Limit(limit).Cursor(cursor).Execute()



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
	r, err := apiClient.DataLinksAPI.GetDataLinkTableInstanceLinks(context.Background()).Id(id).InstanceKey(instanceKey).DataLinkTableId(dataLinkTableId).DataLinkTableKey(dataLinkTableKey).AutoCreate(autoCreate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Direction(direction).AppRecordId(appRecordId).ExternalRecordId(externalRecordId).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataLinksAPI.GetDataLinkTableInstanceLinks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDataLinkTableInstanceLinksRequest struct via the builder pattern


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


## ListDataLinkTableInstances

> ListDataLinkTableInstances200Response ListDataLinkTableInstances(ctx).Id(id).UserId(userId).DataLinkTableId(dataLinkTableId).ConnectionId(connectionId).IntegrationId(integrationId).Execute()



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
	dataLinkTableId := "dataLinkTableId_example" // string |  (optional)
	connectionId := "connectionId_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataLinksAPI.ListDataLinkTableInstances(context.Background()).Id(id).UserId(userId).DataLinkTableId(dataLinkTableId).ConnectionId(connectionId).IntegrationId(integrationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataLinksAPI.ListDataLinkTableInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDataLinkTableInstances`: ListDataLinkTableInstances200Response
	fmt.Fprintf(os.Stdout, "Response from `DataLinksAPI.ListDataLinkTableInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDataLinkTableInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **userId** | **string** |  | 
 **dataLinkTableId** | **string** |  | 
 **connectionId** | **string** |  | 
 **integrationId** | **string** |  | 

### Return type

[**ListDataLinkTableInstances200Response**](ListDataLinkTableInstances200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDataLinkTables

> ListDataLinkTables200Response ListDataLinkTables(ctx).Execute()



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
	resp, r, err := apiClient.DataLinksAPI.ListDataLinkTables(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataLinksAPI.ListDataLinkTables``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDataLinkTables`: ListDataLinkTables200Response
	fmt.Fprintf(os.Stdout, "Response from `DataLinksAPI.ListDataLinkTables`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDataLinkTablesRequest struct via the builder pattern


### Return type

[**ListDataLinkTables200Response**](ListDataLinkTables200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDataLinks

> ListDataLinks200Response ListDataLinks(ctx).DataLinkTableInstanceId(dataLinkTableInstanceId).ExternalRecordId(externalRecordId).AppRecordId(appRecordId).Direction(direction).Execute()



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
	externalRecordId := "externalRecordId_example" // string |  (optional)
	appRecordId := "appRecordId_example" // string |  (optional)
	direction := "direction_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataLinksAPI.ListDataLinks(context.Background()).DataLinkTableInstanceId(dataLinkTableInstanceId).ExternalRecordId(externalRecordId).AppRecordId(appRecordId).Direction(direction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataLinksAPI.ListDataLinks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDataLinks`: ListDataLinks200Response
	fmt.Fprintf(os.Stdout, "Response from `DataLinksAPI.ListDataLinks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDataLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataLinkTableInstanceId** | **string** |  | 
 **externalRecordId** | **string** |  | 
 **appRecordId** | **string** |  | 
 **direction** | **string** |  | 

### Return type

[**ListDataLinks200Response**](ListDataLinks200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchDataLinkTable

> DataLinkTableDto PatchDataLinkTable(ctx, idOrKey).UpdateDataLinkTableDto(updateDataLinkTableDto).Execute()



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
	resp, r, err := apiClient.DataLinksAPI.PatchDataLinkTable(context.Background(), idOrKey).UpdateDataLinkTableDto(updateDataLinkTableDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataLinksAPI.PatchDataLinkTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchDataLinkTable`: DataLinkTableDto
	fmt.Fprintf(os.Stdout, "Response from `DataLinksAPI.PatchDataLinkTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchDataLinkTableRequest struct via the builder pattern


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


## PatchDataLinkTableInstance

> DataLinkTableInstanceDto PatchDataLinkTableInstance(ctx).Body(body).Id(id).InstanceKey(instanceKey).DataLinkTableId(dataLinkTableId).DataLinkTableKey(dataLinkTableKey).AutoCreate(autoCreate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()



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
	resp, r, err := apiClient.DataLinksAPI.PatchDataLinkTableInstance(context.Background()).Body(body).Id(id).InstanceKey(instanceKey).DataLinkTableId(dataLinkTableId).DataLinkTableKey(dataLinkTableKey).AutoCreate(autoCreate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataLinksAPI.PatchDataLinkTableInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchDataLinkTableInstance`: DataLinkTableInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `DataLinksAPI.PatchDataLinkTableInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchDataLinkTableInstanceRequest struct via the builder pattern


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


## PostDataLinkTableInstance

> DataLinkTableInstanceDto PostDataLinkTableInstance(ctx).Body(body).Id(id).InstanceKey(instanceKey).DataLinkTableId(dataLinkTableId).DataLinkTableKey(dataLinkTableKey).AutoCreate(autoCreate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()



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
	resp, r, err := apiClient.DataLinksAPI.PostDataLinkTableInstance(context.Background()).Body(body).Id(id).InstanceKey(instanceKey).DataLinkTableId(dataLinkTableId).DataLinkTableKey(dataLinkTableKey).AutoCreate(autoCreate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataLinksAPI.PostDataLinkTableInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostDataLinkTableInstance`: DataLinkTableInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `DataLinksAPI.PostDataLinkTableInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostDataLinkTableInstanceRequest struct via the builder pattern


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


## PutDataLinkTable

> DataLinkTableDto PutDataLinkTable(ctx, idOrKey).CreateDataLinkTableDto(createDataLinkTableDto).Execute()



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
	createDataLinkTableDto := *openapiclient.NewCreateDataLinkTableDto("Key_example", "Name_example") // CreateDataLinkTableDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataLinksAPI.PutDataLinkTable(context.Background(), idOrKey).CreateDataLinkTableDto(createDataLinkTableDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataLinksAPI.PutDataLinkTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutDataLinkTable`: DataLinkTableDto
	fmt.Fprintf(os.Stdout, "Response from `DataLinksAPI.PutDataLinkTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutDataLinkTableRequest struct via the builder pattern


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


## PutDataLinkTableInstance

> DataLinkTableInstanceDto PutDataLinkTableInstance(ctx).Body(body).Id(id).InstanceKey(instanceKey).DataLinkTableId(dataLinkTableId).DataLinkTableKey(dataLinkTableKey).AutoCreate(autoCreate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()



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
	resp, r, err := apiClient.DataLinksAPI.PutDataLinkTableInstance(context.Background()).Body(body).Id(id).InstanceKey(instanceKey).DataLinkTableId(dataLinkTableId).DataLinkTableKey(dataLinkTableKey).AutoCreate(autoCreate).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataLinksAPI.PutDataLinkTableInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutDataLinkTableInstance`: DataLinkTableInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `DataLinksAPI.PutDataLinkTableInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutDataLinkTableInstanceRequest struct via the builder pattern


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

