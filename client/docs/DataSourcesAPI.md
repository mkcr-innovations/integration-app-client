# \DataSourcesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyDataSource**](DataSourcesAPI.md#ApplyDataSource) | **Post** /data-sources/{id}/apply | 
[**ApplyDataSources**](DataSourcesAPI.md#ApplyDataSources) | **Post** /data-source/apply | 
[**ArchiveDataSource**](DataSourcesAPI.md#ArchiveDataSource) | **Delete** /data-sources/{id} | 
[**ArchiveDataSourceInstance**](DataSourcesAPI.md#ArchiveDataSourceInstance) | **Delete** /data-source-instance | 
[**ArchiveDataSources**](DataSourcesAPI.md#ArchiveDataSources) | **Delete** /data-source | 
[**CreateDataSource**](DataSourcesAPI.md#CreateDataSource) | **Post** /data-sources | 
[**CreateDataSourceInstance**](DataSourcesAPI.md#CreateDataSourceInstance) | **Post** /data-source-instance | 
[**GetDataSource**](DataSourcesAPI.md#GetDataSource) | **Get** /data-sources/{id} | 
[**GetDataSourceEvent**](DataSourcesAPI.md#GetDataSourceEvent) | **Get** /data-source-events/{id} | 
[**GetDataSourceInstance**](DataSourcesAPI.md#GetDataSourceInstance) | **Get** /data-source-instance | 
[**GetDataSourceSync**](DataSourcesAPI.md#GetDataSourceSync) | **Get** /data-source-syncs/{id} | 
[**GetDataSourceSyncLogs**](DataSourcesAPI.md#GetDataSourceSyncLogs) | **Get** /data-source-syncs/{id}/logs | 
[**ListDataSource**](DataSourcesAPI.md#ListDataSource) | **Get** /data-source | 
[**ListDataSourceEvents**](DataSourcesAPI.md#ListDataSourceEvents) | **Get** /data-source-events | 
[**ListDataSourceInstances**](DataSourcesAPI.md#ListDataSourceInstances) | **Get** /data-source-instances | 
[**ListDataSourceSyncs**](DataSourcesAPI.md#ListDataSourceSyncs) | **Get** /data-source-syncs | 
[**ListDataSources**](DataSourcesAPI.md#ListDataSources) | **Get** /data-sources | 
[**PatchDataSource**](DataSourcesAPI.md#PatchDataSource) | **Patch** /data-sources/{id} | 
[**PatchDataSourceInstance**](DataSourcesAPI.md#PatchDataSourceInstance) | **Patch** /data-source-instance | 
[**PatchDataSources**](DataSourcesAPI.md#PatchDataSources) | **Patch** /data-source | 
[**PutDataSource**](DataSourcesAPI.md#PutDataSource) | **Put** /data-sources/{id} | 
[**PutDataSourceInstance**](DataSourcesAPI.md#PutDataSourceInstance) | **Put** /data-source-instance | 
[**PutDataSources**](DataSourcesAPI.md#PutDataSources) | **Put** /data-source | 
[**ResetDataSource**](DataSourcesAPI.md#ResetDataSource) | **Post** /data-sources/{id}/reset | 
[**ResetDataSourceInstance**](DataSourcesAPI.md#ResetDataSourceInstance) | **Post** /data-source-instance/reset | 
[**ResetDataSources**](DataSourcesAPI.md#ResetDataSources) | **Post** /data-source/reset | 
[**SetupDataSourceInstance**](DataSourcesAPI.md#SetupDataSourceInstance) | **Post** /data-source-instance/setup | 



## ApplyDataSource

> []DataSourceDto ApplyDataSource(ctx, id).Id2(id2).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()



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
	id := "id_example" // string | The ID of the data-source to apply
	id2 := "id_example" // string |  (optional)
	key := "key_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataSourcesAPI.ApplyDataSource(context.Background(), id).Id2(id2).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.ApplyDataSource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApplyDataSource`: []DataSourceDto
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.ApplyDataSource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the data-source to apply | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplyDataSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id2** | **string** |  | 
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


## ApplyDataSources

> []DataSourceDto ApplyDataSources(ctx).Id(id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()



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
	key := "key_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataSourcesAPI.ApplyDataSources(context.Background()).Id(id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.ApplyDataSources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApplyDataSources`: []DataSourceDto
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.ApplyDataSources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApplyDataSourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
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


## ArchiveDataSource

> ArchiveDataSource(ctx, id).Id2(id2).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()



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
	id := "id_example" // string | The ID of the data-source to retrive
	id2 := "id_example" // string |  (optional)
	key := "key_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DataSourcesAPI.ArchiveDataSource(context.Background(), id).Id2(id2).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.ArchiveDataSource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the data-source to retrive | 

### Other Parameters

Other parameters are passed through a pointer to a apiArchiveDataSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id2** | **string** |  | 
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


## ArchiveDataSourceInstance

> ArchiveDataSourceInstance(ctx).Id(id).DataSourceKey(dataSourceKey).DataSourceId(dataSourceId).InstanceKey(instanceKey).AutoCreate(autoCreate).FlowKey(flowKey).FlowId(flowId).NodeKey(nodeKey).Udm(udm).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()



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
	r, err := apiClient.DataSourcesAPI.ArchiveDataSourceInstance(context.Background()).Id(id).DataSourceKey(dataSourceKey).DataSourceId(dataSourceId).InstanceKey(instanceKey).AutoCreate(autoCreate).FlowKey(flowKey).FlowId(flowId).NodeKey(nodeKey).Udm(udm).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.ArchiveDataSourceInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArchiveDataSourceInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
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


## ArchiveDataSources

> ArchiveDataSources(ctx).Id(id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()



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
	key := "key_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DataSourcesAPI.ArchiveDataSources(context.Background()).Id(id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.ArchiveDataSources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArchiveDataSourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
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


## CreateDataSource

> DataSourceDto CreateDataSource(ctx).CreateDataSourceDto(createDataSourceDto).Execute()



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
	createDataSourceDto := *openapiclient.NewCreateDataSourceDto("Key_example", "Name_example") // CreateDataSourceDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataSourcesAPI.CreateDataSource(context.Background()).CreateDataSourceDto(createDataSourceDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.CreateDataSource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDataSource`: DataSourceDto
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.CreateDataSource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDataSourceRequest struct via the builder pattern


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


## CreateDataSourceInstance

> DataSourceInstanceDto CreateDataSourceInstance(ctx).UpdateDataSourceInstanceDto(updateDataSourceInstanceDto).Id(id).DataSourceKey(dataSourceKey).DataSourceId(dataSourceId).InstanceKey(instanceKey).AutoCreate(autoCreate).FlowKey(flowKey).FlowId(flowId).NodeKey(nodeKey).Udm(udm).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()



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
	updateDataSourceInstanceDto := *openapiclient.NewUpdateDataSourceInstanceDto() // UpdateDataSourceInstanceDto | 
	id := "id_example" // string |  (optional)
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
	resp, r, err := apiClient.DataSourcesAPI.CreateDataSourceInstance(context.Background()).UpdateDataSourceInstanceDto(updateDataSourceInstanceDto).Id(id).DataSourceKey(dataSourceKey).DataSourceId(dataSourceId).InstanceKey(instanceKey).AutoCreate(autoCreate).FlowKey(flowKey).FlowId(flowId).NodeKey(nodeKey).Udm(udm).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.CreateDataSourceInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDataSourceInstance`: DataSourceInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.CreateDataSourceInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDataSourceInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateDataSourceInstanceDto** | [**UpdateDataSourceInstanceDto**](UpdateDataSourceInstanceDto.md) |  | 
 **id** | **string** |  | 
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


## GetDataSource

> DataSourceDto GetDataSource(ctx, id).Id2(id2).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()



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
	id := "id_example" // string | The ID of the data-source to retrive
	id2 := "id_example" // string |  (optional)
	key := "key_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataSourcesAPI.GetDataSource(context.Background(), id).Id2(id2).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.GetDataSource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDataSource`: DataSourceDto
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.GetDataSource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the data-source to retrive | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id2** | **string** |  | 
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


## GetDataSourceEvent

> DataSourceEventDto GetDataSourceEvent(ctx, id).Execute()



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
	resp, r, err := apiClient.DataSourcesAPI.GetDataSourceEvent(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.GetDataSourceEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDataSourceEvent`: DataSourceEventDto
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.GetDataSourceEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataSourceEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DataSourceEventDto**](DataSourceEventDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataSourceInstance

> DataSourceInstanceDto GetDataSourceInstance(ctx).Id(id).DataSourceKey(dataSourceKey).DataSourceId(dataSourceId).InstanceKey(instanceKey).AutoCreate(autoCreate).FlowKey(flowKey).FlowId(flowId).NodeKey(nodeKey).Udm(udm).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()



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
	resp, r, err := apiClient.DataSourcesAPI.GetDataSourceInstance(context.Background()).Id(id).DataSourceKey(dataSourceKey).DataSourceId(dataSourceId).InstanceKey(instanceKey).AutoCreate(autoCreate).FlowKey(flowKey).FlowId(flowId).NodeKey(nodeKey).Udm(udm).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.GetDataSourceInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDataSourceInstance`: DataSourceInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.GetDataSourceInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDataSourceInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
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


## GetDataSourceSync

> DataSourceSyncDto GetDataSourceSync(ctx, id).Execute()



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
	resp, r, err := apiClient.DataSourcesAPI.GetDataSourceSync(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.GetDataSourceSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDataSourceSync`: DataSourceSyncDto
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.GetDataSourceSync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataSourceSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DataSourceSyncDto**](DataSourceSyncDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataSourceSyncLogs

> map[string]interface{} GetDataSourceSyncLogs(ctx, id).Execute()



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
	resp, r, err := apiClient.DataSourcesAPI.GetDataSourceSyncLogs(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.GetDataSourceSyncLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDataSourceSyncLogs`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.GetDataSourceSyncLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataSourceSyncLogsRequest struct via the builder pattern


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


## ListDataSource

> DataSourceDto ListDataSource(ctx).Id(id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()



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
	key := "key_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataSourcesAPI.ListDataSource(context.Background()).Id(id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.ListDataSource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDataSource`: DataSourceDto
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.ListDataSource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDataSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
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


## ListDataSourceEvents

> ListDataSourceEvents200Response ListDataSourceEvents(ctx).Id(id).UserId(userId).DataSourceId(dataSourceId).ConnectionId(connectionId).IntegrationId(integrationId).DataSourceInstanceId(dataSourceInstanceId).StartedAfter(startedAfter).Execute()



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
	dataSourceId := "dataSourceId_example" // string |  (optional)
	connectionId := "connectionId_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	dataSourceInstanceId := "dataSourceInstanceId_example" // string |  (optional)
	startedAfter := "startedAfter_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataSourcesAPI.ListDataSourceEvents(context.Background()).Id(id).UserId(userId).DataSourceId(dataSourceId).ConnectionId(connectionId).IntegrationId(integrationId).DataSourceInstanceId(dataSourceInstanceId).StartedAfter(startedAfter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.ListDataSourceEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDataSourceEvents`: ListDataSourceEvents200Response
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.ListDataSourceEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDataSourceEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **userId** | **string** |  | 
 **dataSourceId** | **string** |  | 
 **connectionId** | **string** |  | 
 **integrationId** | **string** |  | 
 **dataSourceInstanceId** | **string** |  | 
 **startedAfter** | **string** |  | 

### Return type

[**ListDataSourceEvents200Response**](ListDataSourceEvents200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDataSourceInstances

> ListDataSourceInstances200Response ListDataSourceInstances(ctx).Id(id).UserId(userId).ConnectionId(connectionId).IntegrationKey(integrationKey).IntegrationId(integrationId).DataSourceId(dataSourceId).UniversalDataSourceId(universalDataSourceId).Udm(udm).InstanceKey(instanceKey).Execute()



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
	connectionId := "connectionId_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	dataSourceId := "dataSourceId_example" // string |  (optional)
	universalDataSourceId := "universalDataSourceId_example" // string |  (optional)
	udm := "udm_example" // string |  (optional)
	instanceKey := "instanceKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataSourcesAPI.ListDataSourceInstances(context.Background()).Id(id).UserId(userId).ConnectionId(connectionId).IntegrationKey(integrationKey).IntegrationId(integrationId).DataSourceId(dataSourceId).UniversalDataSourceId(universalDataSourceId).Udm(udm).InstanceKey(instanceKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.ListDataSourceInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDataSourceInstances`: ListDataSourceInstances200Response
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.ListDataSourceInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDataSourceInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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

[**ListDataSourceInstances200Response**](ListDataSourceInstances200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDataSourceSyncs

> ListDataSourceSyncs200Response ListDataSourceSyncs(ctx).DataSourceId(dataSourceId).DataSourceInstanceId(dataSourceInstanceId).IntegrationId(integrationId).ConnectionId(connectionId).UserId(userId).Status(status).StartedAfter(startedAfter).Execute()



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
	dataSourceId := "dataSourceId_example" // string |  (optional)
	dataSourceInstanceId := "dataSourceInstanceId_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	connectionId := "connectionId_example" // string |  (optional)
	userId := "userId_example" // string |  (optional)
	status := "status_example" // string |  (optional)
	startedAfter := "startedAfter_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataSourcesAPI.ListDataSourceSyncs(context.Background()).DataSourceId(dataSourceId).DataSourceInstanceId(dataSourceInstanceId).IntegrationId(integrationId).ConnectionId(connectionId).UserId(userId).Status(status).StartedAfter(startedAfter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.ListDataSourceSyncs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDataSourceSyncs`: ListDataSourceSyncs200Response
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.ListDataSourceSyncs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDataSourceSyncsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataSourceId** | **string** |  | 
 **dataSourceInstanceId** | **string** |  | 
 **integrationId** | **string** |  | 
 **connectionId** | **string** |  | 
 **userId** | **string** |  | 
 **status** | **string** |  | 
 **startedAfter** | **string** |  | 

### Return type

[**ListDataSourceSyncs200Response**](ListDataSourceSyncs200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDataSources

> ListDataSources200Response ListDataSources(ctx).IntegrationId(integrationId).UniversalDataSourceId(universalDataSourceId).Execute()



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
	integrationId := "integrationId_example" // string |  (optional)
	universalDataSourceId := "universalDataSourceId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataSourcesAPI.ListDataSources(context.Background()).IntegrationId(integrationId).UniversalDataSourceId(universalDataSourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.ListDataSources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDataSources`: ListDataSources200Response
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.ListDataSources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDataSourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **integrationId** | **string** |  | 
 **universalDataSourceId** | **string** |  | 

### Return type

[**ListDataSources200Response**](ListDataSources200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchDataSource

> DataSourceDto PatchDataSource(ctx, id).UpdateDataSourceDto(updateDataSourceDto).Id2(id2).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()



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
	id := "id_example" // string | The ID of the data-source to retrive
	updateDataSourceDto := *openapiclient.NewUpdateDataSourceDto() // UpdateDataSourceDto | 
	id2 := "id_example" // string |  (optional)
	key := "key_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataSourcesAPI.PatchDataSource(context.Background(), id).UpdateDataSourceDto(updateDataSourceDto).Id2(id2).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.PatchDataSource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchDataSource`: DataSourceDto
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.PatchDataSource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the data-source to retrive | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchDataSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDataSourceDto** | [**UpdateDataSourceDto**](UpdateDataSourceDto.md) |  | 
 **id2** | **string** |  | 
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


## PatchDataSourceInstance

> DataSourceInstanceDto PatchDataSourceInstance(ctx).UpdateDataSourceInstanceDto(updateDataSourceInstanceDto).Id(id).DataSourceKey(dataSourceKey).DataSourceId(dataSourceId).InstanceKey(instanceKey).AutoCreate(autoCreate).FlowKey(flowKey).FlowId(flowId).NodeKey(nodeKey).Udm(udm).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()



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
	updateDataSourceInstanceDto := *openapiclient.NewUpdateDataSourceInstanceDto() // UpdateDataSourceInstanceDto | 
	id := "id_example" // string |  (optional)
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
	resp, r, err := apiClient.DataSourcesAPI.PatchDataSourceInstance(context.Background()).UpdateDataSourceInstanceDto(updateDataSourceInstanceDto).Id(id).DataSourceKey(dataSourceKey).DataSourceId(dataSourceId).InstanceKey(instanceKey).AutoCreate(autoCreate).FlowKey(flowKey).FlowId(flowId).NodeKey(nodeKey).Udm(udm).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.PatchDataSourceInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchDataSourceInstance`: DataSourceInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.PatchDataSourceInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchDataSourceInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateDataSourceInstanceDto** | [**UpdateDataSourceInstanceDto**](UpdateDataSourceInstanceDto.md) |  | 
 **id** | **string** |  | 
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


## PatchDataSources

> DataSourceDto PatchDataSources(ctx).UpdateDataSourceDto(updateDataSourceDto).Id(id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()



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
	updateDataSourceDto := *openapiclient.NewUpdateDataSourceDto() // UpdateDataSourceDto | 
	id := "id_example" // string |  (optional)
	key := "key_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataSourcesAPI.PatchDataSources(context.Background()).UpdateDataSourceDto(updateDataSourceDto).Id(id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.PatchDataSources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchDataSources`: DataSourceDto
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.PatchDataSources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchDataSourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateDataSourceDto** | [**UpdateDataSourceDto**](UpdateDataSourceDto.md) |  | 
 **id** | **string** |  | 
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


## PutDataSource

> DataSourceDto PutDataSource(ctx, id).UpdateDataSourceDto(updateDataSourceDto).Id2(id2).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()



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
	id := "id_example" // string | The ID of the data-source to retrive
	updateDataSourceDto := *openapiclient.NewUpdateDataSourceDto() // UpdateDataSourceDto | 
	id2 := "id_example" // string |  (optional)
	key := "key_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataSourcesAPI.PutDataSource(context.Background(), id).UpdateDataSourceDto(updateDataSourceDto).Id2(id2).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.PutDataSource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutDataSource`: DataSourceDto
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.PutDataSource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the data-source to retrive | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutDataSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDataSourceDto** | [**UpdateDataSourceDto**](UpdateDataSourceDto.md) |  | 
 **id2** | **string** |  | 
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


## PutDataSourceInstance

> DataSourceInstanceDto PutDataSourceInstance(ctx).UpdateDataSourceInstanceDto(updateDataSourceInstanceDto).Id(id).DataSourceKey(dataSourceKey).DataSourceId(dataSourceId).InstanceKey(instanceKey).AutoCreate(autoCreate).FlowKey(flowKey).FlowId(flowId).NodeKey(nodeKey).Udm(udm).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()



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
	updateDataSourceInstanceDto := *openapiclient.NewUpdateDataSourceInstanceDto() // UpdateDataSourceInstanceDto | 
	id := "id_example" // string |  (optional)
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
	resp, r, err := apiClient.DataSourcesAPI.PutDataSourceInstance(context.Background()).UpdateDataSourceInstanceDto(updateDataSourceInstanceDto).Id(id).DataSourceKey(dataSourceKey).DataSourceId(dataSourceId).InstanceKey(instanceKey).AutoCreate(autoCreate).FlowKey(flowKey).FlowId(flowId).NodeKey(nodeKey).Udm(udm).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.PutDataSourceInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutDataSourceInstance`: DataSourceInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.PutDataSourceInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutDataSourceInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateDataSourceInstanceDto** | [**UpdateDataSourceInstanceDto**](UpdateDataSourceInstanceDto.md) |  | 
 **id** | **string** |  | 
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


## PutDataSources

> DataSourceDto PutDataSources(ctx).UpdateDataSourceDto(updateDataSourceDto).Id(id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()



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
	updateDataSourceDto := *openapiclient.NewUpdateDataSourceDto() // UpdateDataSourceDto | 
	id := "id_example" // string |  (optional)
	key := "key_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataSourcesAPI.PutDataSources(context.Background()).UpdateDataSourceDto(updateDataSourceDto).Id(id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.PutDataSources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutDataSources`: DataSourceDto
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.PutDataSources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutDataSourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateDataSourceDto** | [**UpdateDataSourceDto**](UpdateDataSourceDto.md) |  | 
 **id** | **string** |  | 
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


## ResetDataSource

> DataSourceDto ResetDataSource(ctx, id).Id2(id2).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()



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
	id := "id_example" // string | The ID of the data-source to reset
	id2 := "id_example" // string |  (optional)
	key := "key_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataSourcesAPI.ResetDataSource(context.Background(), id).Id2(id2).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.ResetDataSource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResetDataSource`: DataSourceDto
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.ResetDataSource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the data-source to reset | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetDataSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id2** | **string** |  | 
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


## ResetDataSourceInstance

> DataSourceInstanceDto ResetDataSourceInstance(ctx).Id(id).DataSourceKey(dataSourceKey).DataSourceId(dataSourceId).InstanceKey(instanceKey).AutoCreate(autoCreate).FlowKey(flowKey).FlowId(flowId).NodeKey(nodeKey).Udm(udm).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()



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
	resp, r, err := apiClient.DataSourcesAPI.ResetDataSourceInstance(context.Background()).Id(id).DataSourceKey(dataSourceKey).DataSourceId(dataSourceId).InstanceKey(instanceKey).AutoCreate(autoCreate).FlowKey(flowKey).FlowId(flowId).NodeKey(nodeKey).Udm(udm).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.ResetDataSourceInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResetDataSourceInstance`: DataSourceInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.ResetDataSourceInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResetDataSourceInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
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


## ResetDataSources

> DataSourceDto ResetDataSources(ctx).Id(id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()



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
	key := "key_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataSourcesAPI.ResetDataSources(context.Background()).Id(id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.ResetDataSources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResetDataSources`: DataSourceDto
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.ResetDataSources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResetDataSourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
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


## SetupDataSourceInstance

> DataSourceInstanceDto SetupDataSourceInstance(ctx).Id(id).DataSourceKey(dataSourceKey).DataSourceId(dataSourceId).InstanceKey(instanceKey).AutoCreate(autoCreate).FlowKey(flowKey).FlowId(flowId).NodeKey(nodeKey).Udm(udm).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()



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
	resp, r, err := apiClient.DataSourcesAPI.SetupDataSourceInstance(context.Background()).Id(id).DataSourceKey(dataSourceKey).DataSourceId(dataSourceId).InstanceKey(instanceKey).AutoCreate(autoCreate).FlowKey(flowKey).FlowId(flowId).NodeKey(nodeKey).Udm(udm).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.SetupDataSourceInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetupDataSourceInstance`: DataSourceInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.SetupDataSourceInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetupDataSourceInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
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

