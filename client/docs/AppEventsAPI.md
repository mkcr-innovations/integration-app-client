# \AppEventsAPI

All URIs are relative to *https://api.integration.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppEventSubscriptionControllerArchiveAppEventSubscription**](AppEventsAPI.md#AppEventSubscriptionControllerArchiveAppEventSubscription) | **Delete** /app-event-subscription | Archive app event subscription
[**AppEventSubscriptionControllerCreateAppEventSubscription**](AppEventsAPI.md#AppEventSubscriptionControllerCreateAppEventSubscription) | **Post** /app-event-subscription | Create app event subscription
[**AppEventSubscriptionControllerGetAppEventSubscription**](AppEventsAPI.md#AppEventSubscriptionControllerGetAppEventSubscription) | **Get** /app-event-subscription | Get app event subscription
[**AppEventSubscriptionControllerGetAppEventSubscriptionSchema**](AppEventsAPI.md#AppEventSubscriptionControllerGetAppEventSubscriptionSchema) | **Get** /app-event-subscription/schema | Get app event subscription schema
[**AppEventSubscriptionControllerPatchAppEventSubscription**](AppEventsAPI.md#AppEventSubscriptionControllerPatchAppEventSubscription) | **Patch** /app-event-subscription | Patch app event subscription
[**AppEventSubscriptionControllerSubscribeToAppEventSubscription**](AppEventsAPI.md#AppEventSubscriptionControllerSubscribeToAppEventSubscription) | **Post** /app-event-subscription/subscribe | Subscribe to app event subscription
[**AppEventSubscriptionControllerUpdateAppEventSubscription**](AppEventsAPI.md#AppEventSubscriptionControllerUpdateAppEventSubscription) | **Put** /app-event-subscription | Update app event subscription
[**AppEventSubscriptionsControllerListAppEventSubscriptions**](AppEventsAPI.md#AppEventSubscriptionsControllerListAppEventSubscriptions) | **Get** /app-event-subscriptions | List app event subscriptions
[**AppEventTypesControllerCreateAppEventType**](AppEventsAPI.md#AppEventTypesControllerCreateAppEventType) | **Post** /app-event-types | Create app event type
[**AppEventTypesControllerDeleteAppEventType**](AppEventsAPI.md#AppEventTypesControllerDeleteAppEventType) | **Delete** /app-event-types/{id} | Archive app event type
[**AppEventTypesControllerExportAppEventType**](AppEventsAPI.md#AppEventTypesControllerExportAppEventType) | **Get** /app-event-types/{id}/export | Export App Event Type to JSON
[**AppEventTypesControllerGetAppEventType**](AppEventsAPI.md#AppEventTypesControllerGetAppEventType) | **Get** /app-event-types/{id} | Get app event type
[**AppEventTypesControllerGetAppEventTypeSchema**](AppEventsAPI.md#AppEventTypesControllerGetAppEventTypeSchema) | **Get** /app-event-types/{id}/schema | Get app event type schema
[**AppEventTypesControllerListAppEventTypes**](AppEventsAPI.md#AppEventTypesControllerListAppEventTypes) | **Get** /app-event-types | List app event types
[**AppEventTypesControllerPatchAppEventType**](AppEventsAPI.md#AppEventTypesControllerPatchAppEventType) | **Patch** /app-event-types/{id} | Patch app event type
[**AppEventTypesControllerPutAppEventType**](AppEventsAPI.md#AppEventTypesControllerPutAppEventType) | **Put** /app-event-types/{id} | Update app event type
[**AppEventsControllerGetAppEvent**](AppEventsAPI.md#AppEventsControllerGetAppEvent) | **Get** /app-events/{id} | Get app event
[**AppEventsControllerListAppEvents**](AppEventsAPI.md#AppEventsControllerListAppEvents) | **Get** /app-events | List app events



## AppEventSubscriptionControllerArchiveAppEventSubscription

> AppEventSubscriptionControllerArchiveAppEventSubscription(ctx).Id(id).InstanceKey(instanceKey).AppEventTypeId(appEventTypeId).AppEventTypeKey(appEventTypeKey).AutoCreate(autoCreate).UserId(userId).Execute()

Archive app event subscription

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
	appEventTypeId := "appEventTypeId_example" // string |  (optional)
	appEventTypeKey := "appEventTypeKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)
	userId := "userId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AppEventsAPI.AppEventSubscriptionControllerArchiveAppEventSubscription(context.Background()).Id(id).InstanceKey(instanceKey).AppEventTypeId(appEventTypeId).AppEventTypeKey(appEventTypeKey).AutoCreate(autoCreate).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEventsAPI.AppEventSubscriptionControllerArchiveAppEventSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppEventSubscriptionControllerArchiveAppEventSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **instanceKey** | **string** |  | 
 **appEventTypeId** | **string** |  | 
 **appEventTypeKey** | **string** |  | 
 **autoCreate** | **bool** |  | 
 **userId** | **string** |  | 

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


## AppEventSubscriptionControllerCreateAppEventSubscription

> AppEventSubscriptionDto AppEventSubscriptionControllerCreateAppEventSubscription(ctx).Body(body).Id(id).InstanceKey(instanceKey).AppEventTypeId(appEventTypeId).AppEventTypeKey(appEventTypeKey).AutoCreate(autoCreate).UserId(userId).Execute()

Create app event subscription

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
	appEventTypeId := "appEventTypeId_example" // string |  (optional)
	appEventTypeKey := "appEventTypeKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)
	userId := "userId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppEventsAPI.AppEventSubscriptionControllerCreateAppEventSubscription(context.Background()).Body(body).Id(id).InstanceKey(instanceKey).AppEventTypeId(appEventTypeId).AppEventTypeKey(appEventTypeKey).AutoCreate(autoCreate).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEventsAPI.AppEventSubscriptionControllerCreateAppEventSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppEventSubscriptionControllerCreateAppEventSubscription`: AppEventSubscriptionDto
	fmt.Fprintf(os.Stdout, "Response from `AppEventsAPI.AppEventSubscriptionControllerCreateAppEventSubscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppEventSubscriptionControllerCreateAppEventSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 
 **id** | **string** |  | 
 **instanceKey** | **string** |  | 
 **appEventTypeId** | **string** |  | 
 **appEventTypeKey** | **string** |  | 
 **autoCreate** | **bool** |  | 
 **userId** | **string** |  | 

### Return type

[**AppEventSubscriptionDto**](AppEventSubscriptionDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppEventSubscriptionControllerGetAppEventSubscription

> AppEventSubscriptionDto AppEventSubscriptionControllerGetAppEventSubscription(ctx).Id(id).InstanceKey(instanceKey).AppEventTypeId(appEventTypeId).AppEventTypeKey(appEventTypeKey).AutoCreate(autoCreate).UserId(userId).Execute()

Get app event subscription

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
	appEventTypeId := "appEventTypeId_example" // string |  (optional)
	appEventTypeKey := "appEventTypeKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)
	userId := "userId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppEventsAPI.AppEventSubscriptionControllerGetAppEventSubscription(context.Background()).Id(id).InstanceKey(instanceKey).AppEventTypeId(appEventTypeId).AppEventTypeKey(appEventTypeKey).AutoCreate(autoCreate).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEventsAPI.AppEventSubscriptionControllerGetAppEventSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppEventSubscriptionControllerGetAppEventSubscription`: AppEventSubscriptionDto
	fmt.Fprintf(os.Stdout, "Response from `AppEventsAPI.AppEventSubscriptionControllerGetAppEventSubscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppEventSubscriptionControllerGetAppEventSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **instanceKey** | **string** |  | 
 **appEventTypeId** | **string** |  | 
 **appEventTypeKey** | **string** |  | 
 **autoCreate** | **bool** |  | 
 **userId** | **string** |  | 

### Return type

[**AppEventSubscriptionDto**](AppEventSubscriptionDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppEventSubscriptionControllerGetAppEventSubscriptionSchema

> map[string]interface{} AppEventSubscriptionControllerGetAppEventSubscriptionSchema(ctx).Id(id).InstanceKey(instanceKey).AppEventTypeId(appEventTypeId).AppEventTypeKey(appEventTypeKey).AutoCreate(autoCreate).UserId(userId).Execute()

Get app event subscription schema

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
	appEventTypeId := "appEventTypeId_example" // string |  (optional)
	appEventTypeKey := "appEventTypeKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)
	userId := "userId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppEventsAPI.AppEventSubscriptionControllerGetAppEventSubscriptionSchema(context.Background()).Id(id).InstanceKey(instanceKey).AppEventTypeId(appEventTypeId).AppEventTypeKey(appEventTypeKey).AutoCreate(autoCreate).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEventsAPI.AppEventSubscriptionControllerGetAppEventSubscriptionSchema``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppEventSubscriptionControllerGetAppEventSubscriptionSchema`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AppEventsAPI.AppEventSubscriptionControllerGetAppEventSubscriptionSchema`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppEventSubscriptionControllerGetAppEventSubscriptionSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **instanceKey** | **string** |  | 
 **appEventTypeId** | **string** |  | 
 **appEventTypeKey** | **string** |  | 
 **autoCreate** | **bool** |  | 
 **userId** | **string** |  | 

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


## AppEventSubscriptionControllerPatchAppEventSubscription

> map[string]interface{} AppEventSubscriptionControllerPatchAppEventSubscription(ctx).Body(body).Id(id).InstanceKey(instanceKey).AppEventTypeId(appEventTypeId).AppEventTypeKey(appEventTypeKey).AutoCreate(autoCreate).UserId(userId).Execute()

Patch app event subscription

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
	appEventTypeId := "appEventTypeId_example" // string |  (optional)
	appEventTypeKey := "appEventTypeKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)
	userId := "userId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppEventsAPI.AppEventSubscriptionControllerPatchAppEventSubscription(context.Background()).Body(body).Id(id).InstanceKey(instanceKey).AppEventTypeId(appEventTypeId).AppEventTypeKey(appEventTypeKey).AutoCreate(autoCreate).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEventsAPI.AppEventSubscriptionControllerPatchAppEventSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppEventSubscriptionControllerPatchAppEventSubscription`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AppEventsAPI.AppEventSubscriptionControllerPatchAppEventSubscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppEventSubscriptionControllerPatchAppEventSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 
 **id** | **string** |  | 
 **instanceKey** | **string** |  | 
 **appEventTypeId** | **string** |  | 
 **appEventTypeKey** | **string** |  | 
 **autoCreate** | **bool** |  | 
 **userId** | **string** |  | 

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


## AppEventSubscriptionControllerSubscribeToAppEventSubscription

> AppEventSubscriptionControllerSubscribeToAppEventSubscription(ctx).Id(id).InstanceKey(instanceKey).AppEventTypeId(appEventTypeId).AppEventTypeKey(appEventTypeKey).AutoCreate(autoCreate).UserId(userId).Execute()

Subscribe to app event subscription

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
	appEventTypeId := "appEventTypeId_example" // string |  (optional)
	appEventTypeKey := "appEventTypeKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)
	userId := "userId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AppEventsAPI.AppEventSubscriptionControllerSubscribeToAppEventSubscription(context.Background()).Id(id).InstanceKey(instanceKey).AppEventTypeId(appEventTypeId).AppEventTypeKey(appEventTypeKey).AutoCreate(autoCreate).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEventsAPI.AppEventSubscriptionControllerSubscribeToAppEventSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppEventSubscriptionControllerSubscribeToAppEventSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **instanceKey** | **string** |  | 
 **appEventTypeId** | **string** |  | 
 **appEventTypeKey** | **string** |  | 
 **autoCreate** | **bool** |  | 
 **userId** | **string** |  | 

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


## AppEventSubscriptionControllerUpdateAppEventSubscription

> AppEventSubscriptionDto AppEventSubscriptionControllerUpdateAppEventSubscription(ctx).Body(body).Id(id).InstanceKey(instanceKey).AppEventTypeId(appEventTypeId).AppEventTypeKey(appEventTypeKey).AutoCreate(autoCreate).UserId(userId).Execute()

Update app event subscription

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
	appEventTypeId := "appEventTypeId_example" // string |  (optional)
	appEventTypeKey := "appEventTypeKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)
	userId := "userId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppEventsAPI.AppEventSubscriptionControllerUpdateAppEventSubscription(context.Background()).Body(body).Id(id).InstanceKey(instanceKey).AppEventTypeId(appEventTypeId).AppEventTypeKey(appEventTypeKey).AutoCreate(autoCreate).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEventsAPI.AppEventSubscriptionControllerUpdateAppEventSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppEventSubscriptionControllerUpdateAppEventSubscription`: AppEventSubscriptionDto
	fmt.Fprintf(os.Stdout, "Response from `AppEventsAPI.AppEventSubscriptionControllerUpdateAppEventSubscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppEventSubscriptionControllerUpdateAppEventSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 
 **id** | **string** |  | 
 **instanceKey** | **string** |  | 
 **appEventTypeId** | **string** |  | 
 **appEventTypeKey** | **string** |  | 
 **autoCreate** | **bool** |  | 
 **userId** | **string** |  | 

### Return type

[**AppEventSubscriptionDto**](AppEventSubscriptionDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppEventSubscriptionsControllerListAppEventSubscriptions

> AppEventSubscriptionsControllerListAppEventSubscriptions200Response AppEventSubscriptionsControllerListAppEventSubscriptions(ctx).Limit(limit).Cursor(cursor).Search(search).ConnectorId(connectorId).Id(id).UserId(userId).AppEventTypeId(appEventTypeId).IsSubscribed(isSubscribed).InstanceKey(instanceKey).Execute()

List app event subscriptions

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
	appEventTypeId := "appEventTypeId_example" // string |  (optional)
	isSubscribed := true // bool |  (optional)
	instanceKey := "instanceKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppEventsAPI.AppEventSubscriptionsControllerListAppEventSubscriptions(context.Background()).Limit(limit).Cursor(cursor).Search(search).ConnectorId(connectorId).Id(id).UserId(userId).AppEventTypeId(appEventTypeId).IsSubscribed(isSubscribed).InstanceKey(instanceKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEventsAPI.AppEventSubscriptionsControllerListAppEventSubscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppEventSubscriptionsControllerListAppEventSubscriptions`: AppEventSubscriptionsControllerListAppEventSubscriptions200Response
	fmt.Fprintf(os.Stdout, "Response from `AppEventsAPI.AppEventSubscriptionsControllerListAppEventSubscriptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppEventSubscriptionsControllerListAppEventSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **float32** |  | 
 **cursor** | **string** |  | 
 **search** | **string** |  | 
 **connectorId** | **string** |  | 
 **id** | **string** |  | 
 **userId** | **string** |  | 
 **appEventTypeId** | **string** |  | 
 **isSubscribed** | **bool** |  | 
 **instanceKey** | **string** |  | 

### Return type

[**AppEventSubscriptionsControllerListAppEventSubscriptions200Response**](AppEventSubscriptionsControllerListAppEventSubscriptions200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppEventTypesControllerCreateAppEventType

> AppEventTypeDto AppEventTypesControllerCreateAppEventType(ctx).CreateAppEventTypeDto(createAppEventTypeDto).Execute()

Create app event type

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
	createAppEventTypeDto := *openapiclient.NewCreateAppEventTypeDto() // CreateAppEventTypeDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppEventsAPI.AppEventTypesControllerCreateAppEventType(context.Background()).CreateAppEventTypeDto(createAppEventTypeDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEventsAPI.AppEventTypesControllerCreateAppEventType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppEventTypesControllerCreateAppEventType`: AppEventTypeDto
	fmt.Fprintf(os.Stdout, "Response from `AppEventsAPI.AppEventTypesControllerCreateAppEventType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppEventTypesControllerCreateAppEventTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAppEventTypeDto** | [**CreateAppEventTypeDto**](CreateAppEventTypeDto.md) |  | 

### Return type

[**AppEventTypeDto**](AppEventTypeDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppEventTypesControllerDeleteAppEventType

> AppEventTypesControllerDeleteAppEventType(ctx, id).Execute()

Archive app event type

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
	r, err := apiClient.AppEventsAPI.AppEventTypesControllerDeleteAppEventType(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEventsAPI.AppEventTypesControllerDeleteAppEventType``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAppEventTypesControllerDeleteAppEventTypeRequest struct via the builder pattern


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


## AppEventTypesControllerExportAppEventType

> AppEventTypeExportDto AppEventTypesControllerExportAppEventType(ctx, id).Execute()

Export App Event Type to JSON

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
	resp, r, err := apiClient.AppEventsAPI.AppEventTypesControllerExportAppEventType(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEventsAPI.AppEventTypesControllerExportAppEventType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppEventTypesControllerExportAppEventType`: AppEventTypeExportDto
	fmt.Fprintf(os.Stdout, "Response from `AppEventsAPI.AppEventTypesControllerExportAppEventType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppEventTypesControllerExportAppEventTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppEventTypeExportDto**](AppEventTypeExportDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppEventTypesControllerGetAppEventType

> AppEventTypeDto AppEventTypesControllerGetAppEventType(ctx, id).Execute()

Get app event type

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
	resp, r, err := apiClient.AppEventsAPI.AppEventTypesControllerGetAppEventType(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEventsAPI.AppEventTypesControllerGetAppEventType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppEventTypesControllerGetAppEventType`: AppEventTypeDto
	fmt.Fprintf(os.Stdout, "Response from `AppEventsAPI.AppEventTypesControllerGetAppEventType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppEventTypesControllerGetAppEventTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppEventTypeDto**](AppEventTypeDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppEventTypesControllerGetAppEventTypeSchema

> map[string]interface{} AppEventTypesControllerGetAppEventTypeSchema(ctx, id).Execute()

Get app event type schema

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
	resp, r, err := apiClient.AppEventsAPI.AppEventTypesControllerGetAppEventTypeSchema(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEventsAPI.AppEventTypesControllerGetAppEventTypeSchema``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppEventTypesControllerGetAppEventTypeSchema`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AppEventsAPI.AppEventTypesControllerGetAppEventTypeSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppEventTypesControllerGetAppEventTypeSchemaRequest struct via the builder pattern


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


## AppEventTypesControllerListAppEventTypes

> AppEventTypesControllerListAppEventTypes200Response AppEventTypesControllerListAppEventTypes(ctx).Limit(limit).Cursor(cursor).Search(search).ConnectorId(connectorId).IncludeArchived(includeArchived).Execute()

List app event types

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppEventsAPI.AppEventTypesControllerListAppEventTypes(context.Background()).Limit(limit).Cursor(cursor).Search(search).ConnectorId(connectorId).IncludeArchived(includeArchived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEventsAPI.AppEventTypesControllerListAppEventTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppEventTypesControllerListAppEventTypes`: AppEventTypesControllerListAppEventTypes200Response
	fmt.Fprintf(os.Stdout, "Response from `AppEventsAPI.AppEventTypesControllerListAppEventTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppEventTypesControllerListAppEventTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **float32** |  | 
 **cursor** | **string** |  | 
 **search** | **string** |  | 
 **connectorId** | **string** |  | 
 **includeArchived** | **bool** |  | 

### Return type

[**AppEventTypesControllerListAppEventTypes200Response**](AppEventTypesControllerListAppEventTypes200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppEventTypesControllerPatchAppEventType

> AppEventTypeDto AppEventTypesControllerPatchAppEventType(ctx, id).UpdateAppEventTypeDto(updateAppEventTypeDto).Execute()

Patch app event type

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
	updateAppEventTypeDto := *openapiclient.NewUpdateAppEventTypeDto() // UpdateAppEventTypeDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppEventsAPI.AppEventTypesControllerPatchAppEventType(context.Background(), id).UpdateAppEventTypeDto(updateAppEventTypeDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEventsAPI.AppEventTypesControllerPatchAppEventType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppEventTypesControllerPatchAppEventType`: AppEventTypeDto
	fmt.Fprintf(os.Stdout, "Response from `AppEventsAPI.AppEventTypesControllerPatchAppEventType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppEventTypesControllerPatchAppEventTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAppEventTypeDto** | [**UpdateAppEventTypeDto**](UpdateAppEventTypeDto.md) |  | 

### Return type

[**AppEventTypeDto**](AppEventTypeDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppEventTypesControllerPutAppEventType

> AppEventTypeDto AppEventTypesControllerPutAppEventType(ctx, id).UpdateAppEventTypeDto(updateAppEventTypeDto).Execute()

Update app event type

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
	updateAppEventTypeDto := *openapiclient.NewUpdateAppEventTypeDto() // UpdateAppEventTypeDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppEventsAPI.AppEventTypesControllerPutAppEventType(context.Background(), id).UpdateAppEventTypeDto(updateAppEventTypeDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEventsAPI.AppEventTypesControllerPutAppEventType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppEventTypesControllerPutAppEventType`: AppEventTypeDto
	fmt.Fprintf(os.Stdout, "Response from `AppEventsAPI.AppEventTypesControllerPutAppEventType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppEventTypesControllerPutAppEventTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAppEventTypeDto** | [**UpdateAppEventTypeDto**](UpdateAppEventTypeDto.md) |  | 

### Return type

[**AppEventTypeDto**](AppEventTypeDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppEventsControllerGetAppEvent

> AppEventDto AppEventsControllerGetAppEvent(ctx, id).Execute()

Get app event

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
	resp, r, err := apiClient.AppEventsAPI.AppEventsControllerGetAppEvent(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEventsAPI.AppEventsControllerGetAppEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppEventsControllerGetAppEvent`: AppEventDto
	fmt.Fprintf(os.Stdout, "Response from `AppEventsAPI.AppEventsControllerGetAppEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppEventsControllerGetAppEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppEventDto**](AppEventDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppEventsControllerListAppEvents

> AppEventsControllerListAppEvents200Response AppEventsControllerListAppEvents(ctx).Limit(limit).Cursor(cursor).Search(search).ConnectorId(connectorId).Id(id).UserId(userId).AppEventTypeId(appEventTypeId).AppEventSubscriptionId(appEventSubscriptionId).StartDatetime(startDatetime).EndDatetime(endDatetime).InstanceKey(instanceKey).Execute()

List app events

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
	appEventTypeId := "appEventTypeId_example" // string |  (optional)
	appEventSubscriptionId := "appEventSubscriptionId_example" // string |  (optional)
	startDatetime := "startDatetime_example" // string |  (optional)
	endDatetime := "endDatetime_example" // string |  (optional)
	instanceKey := "instanceKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppEventsAPI.AppEventsControllerListAppEvents(context.Background()).Limit(limit).Cursor(cursor).Search(search).ConnectorId(connectorId).Id(id).UserId(userId).AppEventTypeId(appEventTypeId).AppEventSubscriptionId(appEventSubscriptionId).StartDatetime(startDatetime).EndDatetime(endDatetime).InstanceKey(instanceKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEventsAPI.AppEventsControllerListAppEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppEventsControllerListAppEvents`: AppEventsControllerListAppEvents200Response
	fmt.Fprintf(os.Stdout, "Response from `AppEventsAPI.AppEventsControllerListAppEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppEventsControllerListAppEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **float32** |  | 
 **cursor** | **string** |  | 
 **search** | **string** |  | 
 **connectorId** | **string** |  | 
 **id** | **string** |  | 
 **userId** | **string** |  | 
 **appEventTypeId** | **string** |  | 
 **appEventSubscriptionId** | **string** |  | 
 **startDatetime** | **string** |  | 
 **endDatetime** | **string** |  | 
 **instanceKey** | **string** |  | 

### Return type

[**AppEventsControllerListAppEvents200Response**](AppEventsControllerListAppEvents200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

