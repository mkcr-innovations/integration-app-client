# \ExternalEventsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetExternalEventLogRecord**](ExternalEventsAPI.md#GetExternalEventLogRecord) | **Get** /external-event-log-records/{id} | 
[**GetExternalEventPull**](ExternalEventsAPI.md#GetExternalEventPull) | **Get** /external-event-pulls/{id} | 
[**GetLogs**](ExternalEventsAPI.md#GetLogs) | **Get** /external-event-pulls/{id}/logs | 
[**List**](ExternalEventsAPI.md#List) | **Get** /external-event-pulls | 
[**ListExternalEventLogRecords**](ExternalEventsAPI.md#ListExternalEventLogRecords) | **Get** /external-event-log-records | 



## GetExternalEventLogRecord

> ExternalEventLogRecordDto GetExternalEventLogRecord(ctx, id).Execute()



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
	resp, r, err := apiClient.ExternalEventsAPI.GetExternalEventLogRecord(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalEventsAPI.GetExternalEventLogRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExternalEventLogRecord`: ExternalEventLogRecordDto
	fmt.Fprintf(os.Stdout, "Response from `ExternalEventsAPI.GetExternalEventLogRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalEventLogRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExternalEventLogRecordDto**](ExternalEventLogRecordDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExternalEventPull

> ExternalEventPullDto GetExternalEventPull(ctx, id).Execute()



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
	resp, r, err := apiClient.ExternalEventsAPI.GetExternalEventPull(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalEventsAPI.GetExternalEventPull``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExternalEventPull`: ExternalEventPullDto
	fmt.Fprintf(os.Stdout, "Response from `ExternalEventsAPI.GetExternalEventPull`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalEventPullRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExternalEventPullDto**](ExternalEventPullDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogs

> map[string]interface{} GetLogs(ctx, id).Execute()



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
	resp, r, err := apiClient.ExternalEventsAPI.GetLogs(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalEventsAPI.GetLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogs`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ExternalEventsAPI.GetLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogsRequest struct via the builder pattern


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


## List

> List200Response List(ctx).ExternalEventSubscriptionId(externalEventSubscriptionId).IntegrationId(integrationId).ConnectionId(connectionId).UserId(userId).Status(status).StartedAfter(startedAfter).Execute()



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
	externalEventSubscriptionId := "externalEventSubscriptionId_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	connectionId := "connectionId_example" // string |  (optional)
	userId := "userId_example" // string |  (optional)
	status := "status_example" // string |  (optional)
	startedAfter := "startedAfter_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalEventsAPI.List(context.Background()).ExternalEventSubscriptionId(externalEventSubscriptionId).IntegrationId(integrationId).ConnectionId(connectionId).UserId(userId).Status(status).StartedAfter(startedAfter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalEventsAPI.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: List200Response
	fmt.Fprintf(os.Stdout, "Response from `ExternalEventsAPI.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **externalEventSubscriptionId** | **string** |  | 
 **integrationId** | **string** |  | 
 **connectionId** | **string** |  | 
 **userId** | **string** |  | 
 **status** | **string** |  | 
 **startedAfter** | **string** |  | 

### Return type

[**List200Response**](List200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListExternalEventLogRecords

> ListExternalEventLogRecords200Response ListExternalEventLogRecords(ctx).Id(id).UserId(userId).ExternalEventSubscriptionId(externalEventSubscriptionId).ConnectionId(connectionId).IntegrationId(integrationId).Execute()



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
	externalEventSubscriptionId := "externalEventSubscriptionId_example" // string |  (optional)
	connectionId := "connectionId_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalEventsAPI.ListExternalEventLogRecords(context.Background()).Id(id).UserId(userId).ExternalEventSubscriptionId(externalEventSubscriptionId).ConnectionId(connectionId).IntegrationId(integrationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalEventsAPI.ListExternalEventLogRecords``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListExternalEventLogRecords`: ListExternalEventLogRecords200Response
	fmt.Fprintf(os.Stdout, "Response from `ExternalEventsAPI.ListExternalEventLogRecords`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListExternalEventLogRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **userId** | **string** |  | 
 **externalEventSubscriptionId** | **string** |  | 
 **connectionId** | **string** |  | 
 **integrationId** | **string** |  | 

### Return type

[**ListExternalEventLogRecords200Response**](ListExternalEventLogRecords200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

