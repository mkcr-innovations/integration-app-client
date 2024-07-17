# \ExternalEventsAPI

All URIs are relative to *https://api.integration.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExternalEventLogRecordsControllerGetExternalEventLogRecord**](ExternalEventsAPI.md#ExternalEventLogRecordsControllerGetExternalEventLogRecord) | **Get** /external-event-log-records/{id} | Get external event log record
[**ExternalEventLogRecordsControllerListExternalEventLogRecords**](ExternalEventsAPI.md#ExternalEventLogRecordsControllerListExternalEventLogRecords) | **Get** /external-event-log-records | List external event log records
[**ExternalEventPullsControllerGetById**](ExternalEventsAPI.md#ExternalEventPullsControllerGetById) | **Get** /external-event-pulls/{id} | Get external event pull
[**ExternalEventPullsControllerGetLogs**](ExternalEventsAPI.md#ExternalEventPullsControllerGetLogs) | **Get** /external-event-pulls/{id}/logs | Get external event pull logs
[**ExternalEventPullsControllerList**](ExternalEventsAPI.md#ExternalEventPullsControllerList) | **Get** /external-event-pulls | List external event pulls
[**ExternalEventSubscriptionsControllerDeleteExternalEventSubscription**](ExternalEventsAPI.md#ExternalEventSubscriptionsControllerDeleteExternalEventSubscription) | **Delete** /external-event-subscriptions/{id} | Delete external event subscription
[**ExternalEventSubscriptionsControllerGetExternalEventSubscription**](ExternalEventsAPI.md#ExternalEventSubscriptionsControllerGetExternalEventSubscription) | **Get** /external-event-subscriptions/{id} | Get external event subscription
[**ExternalEventSubscriptionsControllerListExternalEventSubscriptions**](ExternalEventsAPI.md#ExternalEventSubscriptionsControllerListExternalEventSubscriptions) | **Get** /external-event-subscriptions | List external event subscriptions
[**ExternalEventSubscriptionsControllerPullExternalEventSubscriptionEvents**](ExternalEventsAPI.md#ExternalEventSubscriptionsControllerPullExternalEventSubscriptionEvents) | **Post** /external-event-subscriptions/{id}/pull-events | Triggers pull events for external event subscription
[**ExternalEventSubscriptionsControllerResubscribeToExternalEventSubscription**](ExternalEventsAPI.md#ExternalEventSubscriptionsControllerResubscribeToExternalEventSubscription) | **Post** /external-event-subscriptions/{id}/resubscribe | Resubscribe to external event subscription
[**ExternalEventSubscriptionsControllerSetupExternalEventSubscription**](ExternalEventsAPI.md#ExternalEventSubscriptionsControllerSetupExternalEventSubscription) | **Post** /external-event-subscriptions/{id}/setup | Setup external event subscription
[**ExternalEventSubscriptionsControllerSubscribeToExternalEventSubscription**](ExternalEventsAPI.md#ExternalEventSubscriptionsControllerSubscribeToExternalEventSubscription) | **Post** /external-event-subscriptions/{id}/subscribe | Subscribe to external event subscription
[**ExternalEventSubscriptionsControllerUnsubscribeFromExternalEventSubscription**](ExternalEventsAPI.md#ExternalEventSubscriptionsControllerUnsubscribeFromExternalEventSubscription) | **Post** /external-event-subscriptions/{id}/unsubscribe | Unsubscribe from external event subscription



## ExternalEventLogRecordsControllerGetExternalEventLogRecord

> ExternalEventLogRecordDto ExternalEventLogRecordsControllerGetExternalEventLogRecord(ctx, id).Execute()

Get external event log record

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
	resp, r, err := apiClient.ExternalEventsAPI.ExternalEventLogRecordsControllerGetExternalEventLogRecord(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalEventsAPI.ExternalEventLogRecordsControllerGetExternalEventLogRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExternalEventLogRecordsControllerGetExternalEventLogRecord`: ExternalEventLogRecordDto
	fmt.Fprintf(os.Stdout, "Response from `ExternalEventsAPI.ExternalEventLogRecordsControllerGetExternalEventLogRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExternalEventLogRecordsControllerGetExternalEventLogRecordRequest struct via the builder pattern


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


## ExternalEventLogRecordsControllerListExternalEventLogRecords

> ExternalEventLogRecordsControllerListExternalEventLogRecords200Response ExternalEventLogRecordsControllerListExternalEventLogRecords(ctx).Limit(limit).Cursor(cursor).Search(search).ConnectorId(connectorId).Id(id).UserId(userId).ExternalEventSubscriptionId(externalEventSubscriptionId).ConnectionId(connectionId).IntegrationId(integrationId).Execute()

List external event log records

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
	externalEventSubscriptionId := "externalEventSubscriptionId_example" // string |  (optional)
	connectionId := "connectionId_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalEventsAPI.ExternalEventLogRecordsControllerListExternalEventLogRecords(context.Background()).Limit(limit).Cursor(cursor).Search(search).ConnectorId(connectorId).Id(id).UserId(userId).ExternalEventSubscriptionId(externalEventSubscriptionId).ConnectionId(connectionId).IntegrationId(integrationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalEventsAPI.ExternalEventLogRecordsControllerListExternalEventLogRecords``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExternalEventLogRecordsControllerListExternalEventLogRecords`: ExternalEventLogRecordsControllerListExternalEventLogRecords200Response
	fmt.Fprintf(os.Stdout, "Response from `ExternalEventsAPI.ExternalEventLogRecordsControllerListExternalEventLogRecords`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExternalEventLogRecordsControllerListExternalEventLogRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **float32** |  | 
 **cursor** | **string** |  | 
 **search** | **string** |  | 
 **connectorId** | **string** |  | 
 **id** | **string** |  | 
 **userId** | **string** |  | 
 **externalEventSubscriptionId** | **string** |  | 
 **connectionId** | **string** |  | 
 **integrationId** | **string** |  | 

### Return type

[**ExternalEventLogRecordsControllerListExternalEventLogRecords200Response**](ExternalEventLogRecordsControllerListExternalEventLogRecords200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExternalEventPullsControllerGetById

> ExternalEventPullDto ExternalEventPullsControllerGetById(ctx, id).Execute()

Get external event pull

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
	resp, r, err := apiClient.ExternalEventsAPI.ExternalEventPullsControllerGetById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalEventsAPI.ExternalEventPullsControllerGetById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExternalEventPullsControllerGetById`: ExternalEventPullDto
	fmt.Fprintf(os.Stdout, "Response from `ExternalEventsAPI.ExternalEventPullsControllerGetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExternalEventPullsControllerGetByIdRequest struct via the builder pattern


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


## ExternalEventPullsControllerGetLogs

> map[string]interface{} ExternalEventPullsControllerGetLogs(ctx, id).Execute()

Get external event pull logs

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
	resp, r, err := apiClient.ExternalEventsAPI.ExternalEventPullsControllerGetLogs(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalEventsAPI.ExternalEventPullsControllerGetLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExternalEventPullsControllerGetLogs`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ExternalEventsAPI.ExternalEventPullsControllerGetLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExternalEventPullsControllerGetLogsRequest struct via the builder pattern


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


## ExternalEventPullsControllerList

> ExternalEventPullsControllerList200Response ExternalEventPullsControllerList(ctx).Limit(limit).Cursor(cursor).ExternalEventSubscriptionId(externalEventSubscriptionId).IntegrationId(integrationId).ConnectionId(connectionId).UserId(userId).Status(status).StartedAfter(startedAfter).Execute()

List external event pulls

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
	externalEventSubscriptionId := "externalEventSubscriptionId_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	connectionId := "connectionId_example" // string |  (optional)
	userId := "userId_example" // string |  (optional)
	status := "status_example" // string |  (optional)
	startedAfter := "startedAfter_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalEventsAPI.ExternalEventPullsControllerList(context.Background()).Limit(limit).Cursor(cursor).ExternalEventSubscriptionId(externalEventSubscriptionId).IntegrationId(integrationId).ConnectionId(connectionId).UserId(userId).Status(status).StartedAfter(startedAfter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalEventsAPI.ExternalEventPullsControllerList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExternalEventPullsControllerList`: ExternalEventPullsControllerList200Response
	fmt.Fprintf(os.Stdout, "Response from `ExternalEventsAPI.ExternalEventPullsControllerList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExternalEventPullsControllerListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **float32** |  | 
 **cursor** | **string** |  | 
 **externalEventSubscriptionId** | **string** |  | 
 **integrationId** | **string** |  | 
 **connectionId** | **string** |  | 
 **userId** | **string** |  | 
 **status** | **string** |  | 
 **startedAfter** | **string** |  | 

### Return type

[**ExternalEventPullsControllerList200Response**](ExternalEventPullsControllerList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExternalEventSubscriptionsControllerDeleteExternalEventSubscription

> ExternalEventSubscriptionsControllerDeleteExternalEventSubscription(ctx, id).Execute()

Delete external event subscription

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
	r, err := apiClient.ExternalEventsAPI.ExternalEventSubscriptionsControllerDeleteExternalEventSubscription(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalEventsAPI.ExternalEventSubscriptionsControllerDeleteExternalEventSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExternalEventSubscriptionsControllerDeleteExternalEventSubscriptionRequest struct via the builder pattern


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


## ExternalEventSubscriptionsControllerGetExternalEventSubscription

> ExternalEventSubscriptionDto ExternalEventSubscriptionsControllerGetExternalEventSubscription(ctx, id).Execute()

Get external event subscription

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
	resp, r, err := apiClient.ExternalEventsAPI.ExternalEventSubscriptionsControllerGetExternalEventSubscription(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalEventsAPI.ExternalEventSubscriptionsControllerGetExternalEventSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExternalEventSubscriptionsControllerGetExternalEventSubscription`: ExternalEventSubscriptionDto
	fmt.Fprintf(os.Stdout, "Response from `ExternalEventsAPI.ExternalEventSubscriptionsControllerGetExternalEventSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExternalEventSubscriptionsControllerGetExternalEventSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExternalEventSubscriptionDto**](ExternalEventSubscriptionDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExternalEventSubscriptionsControllerListExternalEventSubscriptions

> ExternalEventSubscriptionsControllerListExternalEventSubscriptions200Response ExternalEventSubscriptionsControllerListExternalEventSubscriptions(ctx).Limit(limit).Cursor(cursor).Search(search).ConnectorId(connectorId).UserId(userId).ConnectionId(connectionId).IntegrationId(integrationId).Execute()

List external event subscriptions

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
	userId := "userId_example" // string |  (optional)
	connectionId := "connectionId_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalEventsAPI.ExternalEventSubscriptionsControllerListExternalEventSubscriptions(context.Background()).Limit(limit).Cursor(cursor).Search(search).ConnectorId(connectorId).UserId(userId).ConnectionId(connectionId).IntegrationId(integrationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalEventsAPI.ExternalEventSubscriptionsControllerListExternalEventSubscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExternalEventSubscriptionsControllerListExternalEventSubscriptions`: ExternalEventSubscriptionsControllerListExternalEventSubscriptions200Response
	fmt.Fprintf(os.Stdout, "Response from `ExternalEventsAPI.ExternalEventSubscriptionsControllerListExternalEventSubscriptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExternalEventSubscriptionsControllerListExternalEventSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **float32** |  | 
 **cursor** | **string** |  | 
 **search** | **string** |  | 
 **connectorId** | **string** |  | 
 **userId** | **string** |  | 
 **connectionId** | **string** |  | 
 **integrationId** | **string** |  | 

### Return type

[**ExternalEventSubscriptionsControllerListExternalEventSubscriptions200Response**](ExternalEventSubscriptionsControllerListExternalEventSubscriptions200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExternalEventSubscriptionsControllerPullExternalEventSubscriptionEvents

> ExternalEventSubscriptionDto ExternalEventSubscriptionsControllerPullExternalEventSubscriptionEvents(ctx, id).Execute()

Triggers pull events for external event subscription

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
	resp, r, err := apiClient.ExternalEventsAPI.ExternalEventSubscriptionsControllerPullExternalEventSubscriptionEvents(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalEventsAPI.ExternalEventSubscriptionsControllerPullExternalEventSubscriptionEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExternalEventSubscriptionsControllerPullExternalEventSubscriptionEvents`: ExternalEventSubscriptionDto
	fmt.Fprintf(os.Stdout, "Response from `ExternalEventsAPI.ExternalEventSubscriptionsControllerPullExternalEventSubscriptionEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExternalEventSubscriptionsControllerPullExternalEventSubscriptionEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExternalEventSubscriptionDto**](ExternalEventSubscriptionDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExternalEventSubscriptionsControllerResubscribeToExternalEventSubscription

> ExternalEventSubscriptionDto ExternalEventSubscriptionsControllerResubscribeToExternalEventSubscription(ctx, id).Execute()

Resubscribe to external event subscription

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
	resp, r, err := apiClient.ExternalEventsAPI.ExternalEventSubscriptionsControllerResubscribeToExternalEventSubscription(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalEventsAPI.ExternalEventSubscriptionsControllerResubscribeToExternalEventSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExternalEventSubscriptionsControllerResubscribeToExternalEventSubscription`: ExternalEventSubscriptionDto
	fmt.Fprintf(os.Stdout, "Response from `ExternalEventsAPI.ExternalEventSubscriptionsControllerResubscribeToExternalEventSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExternalEventSubscriptionsControllerResubscribeToExternalEventSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExternalEventSubscriptionDto**](ExternalEventSubscriptionDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExternalEventSubscriptionsControllerSetupExternalEventSubscription

> ExternalEventSubscriptionDto ExternalEventSubscriptionsControllerSetupExternalEventSubscription(ctx, id).Execute()

Setup external event subscription

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
	resp, r, err := apiClient.ExternalEventsAPI.ExternalEventSubscriptionsControllerSetupExternalEventSubscription(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalEventsAPI.ExternalEventSubscriptionsControllerSetupExternalEventSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExternalEventSubscriptionsControllerSetupExternalEventSubscription`: ExternalEventSubscriptionDto
	fmt.Fprintf(os.Stdout, "Response from `ExternalEventsAPI.ExternalEventSubscriptionsControllerSetupExternalEventSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExternalEventSubscriptionsControllerSetupExternalEventSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExternalEventSubscriptionDto**](ExternalEventSubscriptionDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExternalEventSubscriptionsControllerSubscribeToExternalEventSubscription

> ExternalEventSubscriptionDto ExternalEventSubscriptionsControllerSubscribeToExternalEventSubscription(ctx, id).Execute()

Subscribe to external event subscription

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
	resp, r, err := apiClient.ExternalEventsAPI.ExternalEventSubscriptionsControllerSubscribeToExternalEventSubscription(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalEventsAPI.ExternalEventSubscriptionsControllerSubscribeToExternalEventSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExternalEventSubscriptionsControllerSubscribeToExternalEventSubscription`: ExternalEventSubscriptionDto
	fmt.Fprintf(os.Stdout, "Response from `ExternalEventsAPI.ExternalEventSubscriptionsControllerSubscribeToExternalEventSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExternalEventSubscriptionsControllerSubscribeToExternalEventSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExternalEventSubscriptionDto**](ExternalEventSubscriptionDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExternalEventSubscriptionsControllerUnsubscribeFromExternalEventSubscription

> ExternalEventSubscriptionDto ExternalEventSubscriptionsControllerUnsubscribeFromExternalEventSubscription(ctx, id).Execute()

Unsubscribe from external event subscription

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
	resp, r, err := apiClient.ExternalEventsAPI.ExternalEventSubscriptionsControllerUnsubscribeFromExternalEventSubscription(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalEventsAPI.ExternalEventSubscriptionsControllerUnsubscribeFromExternalEventSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExternalEventSubscriptionsControllerUnsubscribeFromExternalEventSubscription`: ExternalEventSubscriptionDto
	fmt.Fprintf(os.Stdout, "Response from `ExternalEventsAPI.ExternalEventSubscriptionsControllerUnsubscribeFromExternalEventSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExternalEventSubscriptionsControllerUnsubscribeFromExternalEventSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExternalEventSubscriptionDto**](ExternalEventSubscriptionDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

