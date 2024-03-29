# \AppEventsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveAppEventSubscription**](AppEventsAPI.md#ArchiveAppEventSubscription) | **Delete** /app-event-subscription | 
[**CreateAppEventSubscription**](AppEventsAPI.md#CreateAppEventSubscription) | **Post** /app-event-subscription | 
[**CreateAppEventType**](AppEventsAPI.md#CreateAppEventType) | **Post** /app-event-types | 
[**DeleteAppEventType**](AppEventsAPI.md#DeleteAppEventType) | **Delete** /app-event-types/{id} | 
[**GetAppEvent**](AppEventsAPI.md#GetAppEvent) | **Get** /app-events/{id} | 
[**GetAppEventSubscription**](AppEventsAPI.md#GetAppEventSubscription) | **Get** /app-event-subscription | 
[**GetAppEventSubscriptionSchema**](AppEventsAPI.md#GetAppEventSubscriptionSchema) | **Get** /app-event-subscription/schema | 
[**GetAppEventType**](AppEventsAPI.md#GetAppEventType) | **Get** /app-event-types/{id} | 
[**GetAppEventTypeSchema**](AppEventsAPI.md#GetAppEventTypeSchema) | **Get** /app-event-types/{id}/schema | 
[**ListAppEventSubscriptions**](AppEventsAPI.md#ListAppEventSubscriptions) | **Get** /app-event-subscriptions | 
[**ListAppEventTypes**](AppEventsAPI.md#ListAppEventTypes) | **Get** /app-event-types | 
[**ListAppEvents**](AppEventsAPI.md#ListAppEvents) | **Get** /app-events | 
[**PatchAppEventSubscription**](AppEventsAPI.md#PatchAppEventSubscription) | **Patch** /app-event-subscription | 
[**PatchAppEventType**](AppEventsAPI.md#PatchAppEventType) | **Patch** /app-event-types/{id} | 
[**PutAppEventType**](AppEventsAPI.md#PutAppEventType) | **Put** /app-event-types/{id} | 
[**SubscribeToAppEventSubscription**](AppEventsAPI.md#SubscribeToAppEventSubscription) | **Post** /app-event-subscription/subscribe | 
[**UpdateAppEventSubscription**](AppEventsAPI.md#UpdateAppEventSubscription) | **Put** /app-event-subscription | 



## ArchiveAppEventSubscription

> ArchiveAppEventSubscription(ctx).Id(id).InstanceKey(instanceKey).AppEventTypeId(appEventTypeId).AppEventTypeKey(appEventTypeKey).AutoCreate(autoCreate).UserId(userId).Execute()



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
	r, err := apiClient.AppEventsAPI.ArchiveAppEventSubscription(context.Background()).Id(id).InstanceKey(instanceKey).AppEventTypeId(appEventTypeId).AppEventTypeKey(appEventTypeKey).AutoCreate(autoCreate).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEventsAPI.ArchiveAppEventSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArchiveAppEventSubscriptionRequest struct via the builder pattern


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


## CreateAppEventSubscription

> AppEventSubscriptionDto CreateAppEventSubscription(ctx).Body(body).Id(id).InstanceKey(instanceKey).AppEventTypeId(appEventTypeId).AppEventTypeKey(appEventTypeKey).AutoCreate(autoCreate).UserId(userId).Execute()



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
	resp, r, err := apiClient.AppEventsAPI.CreateAppEventSubscription(context.Background()).Body(body).Id(id).InstanceKey(instanceKey).AppEventTypeId(appEventTypeId).AppEventTypeKey(appEventTypeKey).AutoCreate(autoCreate).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEventsAPI.CreateAppEventSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAppEventSubscription`: AppEventSubscriptionDto
	fmt.Fprintf(os.Stdout, "Response from `AppEventsAPI.CreateAppEventSubscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAppEventSubscriptionRequest struct via the builder pattern


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


## CreateAppEventType

> AppEventTypeDto CreateAppEventType(ctx).CreateAppEventTypeDto(createAppEventTypeDto).Execute()



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
	createAppEventTypeDto := *openapiclient.NewCreateAppEventTypeDto("Key_example", "Name_example") // CreateAppEventTypeDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppEventsAPI.CreateAppEventType(context.Background()).CreateAppEventTypeDto(createAppEventTypeDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEventsAPI.CreateAppEventType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAppEventType`: AppEventTypeDto
	fmt.Fprintf(os.Stdout, "Response from `AppEventsAPI.CreateAppEventType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAppEventTypeRequest struct via the builder pattern


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


## DeleteAppEventType

> DeleteAppEventType(ctx, id).Execute()



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
	r, err := apiClient.AppEventsAPI.DeleteAppEventType(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEventsAPI.DeleteAppEventType``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteAppEventTypeRequest struct via the builder pattern


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


## GetAppEvent

> AppEventDto GetAppEvent(ctx, id).Execute()



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
	resp, r, err := apiClient.AppEventsAPI.GetAppEvent(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEventsAPI.GetAppEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppEvent`: AppEventDto
	fmt.Fprintf(os.Stdout, "Response from `AppEventsAPI.GetAppEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppEventRequest struct via the builder pattern


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


## GetAppEventSubscription

> AppEventSubscriptionDto GetAppEventSubscription(ctx).Id(id).InstanceKey(instanceKey).AppEventTypeId(appEventTypeId).AppEventTypeKey(appEventTypeKey).AutoCreate(autoCreate).UserId(userId).Execute()



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
	resp, r, err := apiClient.AppEventsAPI.GetAppEventSubscription(context.Background()).Id(id).InstanceKey(instanceKey).AppEventTypeId(appEventTypeId).AppEventTypeKey(appEventTypeKey).AutoCreate(autoCreate).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEventsAPI.GetAppEventSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppEventSubscription`: AppEventSubscriptionDto
	fmt.Fprintf(os.Stdout, "Response from `AppEventsAPI.GetAppEventSubscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAppEventSubscriptionRequest struct via the builder pattern


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


## GetAppEventSubscriptionSchema

> map[string]interface{} GetAppEventSubscriptionSchema(ctx).Id(id).InstanceKey(instanceKey).AppEventTypeId(appEventTypeId).AppEventTypeKey(appEventTypeKey).AutoCreate(autoCreate).UserId(userId).Execute()



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
	resp, r, err := apiClient.AppEventsAPI.GetAppEventSubscriptionSchema(context.Background()).Id(id).InstanceKey(instanceKey).AppEventTypeId(appEventTypeId).AppEventTypeKey(appEventTypeKey).AutoCreate(autoCreate).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEventsAPI.GetAppEventSubscriptionSchema``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppEventSubscriptionSchema`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AppEventsAPI.GetAppEventSubscriptionSchema`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAppEventSubscriptionSchemaRequest struct via the builder pattern


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


## GetAppEventType

> AppEventTypeDto GetAppEventType(ctx, id).Execute()



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
	resp, r, err := apiClient.AppEventsAPI.GetAppEventType(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEventsAPI.GetAppEventType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppEventType`: AppEventTypeDto
	fmt.Fprintf(os.Stdout, "Response from `AppEventsAPI.GetAppEventType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppEventTypeRequest struct via the builder pattern


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


## GetAppEventTypeSchema

> map[string]interface{} GetAppEventTypeSchema(ctx, id).Execute()



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
	resp, r, err := apiClient.AppEventsAPI.GetAppEventTypeSchema(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEventsAPI.GetAppEventTypeSchema``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppEventTypeSchema`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AppEventsAPI.GetAppEventTypeSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppEventTypeSchemaRequest struct via the builder pattern


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


## ListAppEventSubscriptions

> ListAppEventSubscriptions200Response ListAppEventSubscriptions(ctx).Id(id).UserId(userId).AppEventTypeId(appEventTypeId).IsSubscribed(isSubscribed).InstanceKey(instanceKey).Execute()



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
	appEventTypeId := "appEventTypeId_example" // string |  (optional)
	isSubscribed := true // bool |  (optional)
	instanceKey := "instanceKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppEventsAPI.ListAppEventSubscriptions(context.Background()).Id(id).UserId(userId).AppEventTypeId(appEventTypeId).IsSubscribed(isSubscribed).InstanceKey(instanceKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEventsAPI.ListAppEventSubscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAppEventSubscriptions`: ListAppEventSubscriptions200Response
	fmt.Fprintf(os.Stdout, "Response from `AppEventsAPI.ListAppEventSubscriptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAppEventSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **userId** | **string** |  | 
 **appEventTypeId** | **string** |  | 
 **isSubscribed** | **bool** |  | 
 **instanceKey** | **string** |  | 

### Return type

[**ListAppEventSubscriptions200Response**](ListAppEventSubscriptions200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAppEventTypes

> ListAppEventTypes200Response ListAppEventTypes(ctx).Execute()



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
	resp, r, err := apiClient.AppEventsAPI.ListAppEventTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEventsAPI.ListAppEventTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAppEventTypes`: ListAppEventTypes200Response
	fmt.Fprintf(os.Stdout, "Response from `AppEventsAPI.ListAppEventTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAppEventTypesRequest struct via the builder pattern


### Return type

[**ListAppEventTypes200Response**](ListAppEventTypes200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAppEvents

> ListAppEvents200Response ListAppEvents(ctx).Id(id).UserId(userId).AppEventTypeId(appEventTypeId).AppEventSubscriptionId(appEventSubscriptionId).StartDatetime(startDatetime).EndDatetime(endDatetime).InstanceKey(instanceKey).Execute()



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
	appEventTypeId := "appEventTypeId_example" // string |  (optional)
	appEventSubscriptionId := "appEventSubscriptionId_example" // string |  (optional)
	startDatetime := "startDatetime_example" // string |  (optional)
	endDatetime := "endDatetime_example" // string |  (optional)
	instanceKey := "instanceKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppEventsAPI.ListAppEvents(context.Background()).Id(id).UserId(userId).AppEventTypeId(appEventTypeId).AppEventSubscriptionId(appEventSubscriptionId).StartDatetime(startDatetime).EndDatetime(endDatetime).InstanceKey(instanceKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEventsAPI.ListAppEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAppEvents`: ListAppEvents200Response
	fmt.Fprintf(os.Stdout, "Response from `AppEventsAPI.ListAppEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAppEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **userId** | **string** |  | 
 **appEventTypeId** | **string** |  | 
 **appEventSubscriptionId** | **string** |  | 
 **startDatetime** | **string** |  | 
 **endDatetime** | **string** |  | 
 **instanceKey** | **string** |  | 

### Return type

[**ListAppEvents200Response**](ListAppEvents200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAppEventSubscription

> map[string]interface{} PatchAppEventSubscription(ctx).Body(body).Id(id).InstanceKey(instanceKey).AppEventTypeId(appEventTypeId).AppEventTypeKey(appEventTypeKey).AutoCreate(autoCreate).UserId(userId).Execute()



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
	resp, r, err := apiClient.AppEventsAPI.PatchAppEventSubscription(context.Background()).Body(body).Id(id).InstanceKey(instanceKey).AppEventTypeId(appEventTypeId).AppEventTypeKey(appEventTypeKey).AutoCreate(autoCreate).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEventsAPI.PatchAppEventSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAppEventSubscription`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AppEventsAPI.PatchAppEventSubscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchAppEventSubscriptionRequest struct via the builder pattern


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


## PatchAppEventType

> AppEventTypeDto PatchAppEventType(ctx, id).UpdateAppEventTypeDto(updateAppEventTypeDto).Execute()



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
	resp, r, err := apiClient.AppEventsAPI.PatchAppEventType(context.Background(), id).UpdateAppEventTypeDto(updateAppEventTypeDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEventsAPI.PatchAppEventType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAppEventType`: AppEventTypeDto
	fmt.Fprintf(os.Stdout, "Response from `AppEventsAPI.PatchAppEventType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAppEventTypeRequest struct via the builder pattern


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


## PutAppEventType

> AppEventTypeDto PutAppEventType(ctx, id).UpdateAppEventTypeDto(updateAppEventTypeDto).Execute()



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
	resp, r, err := apiClient.AppEventsAPI.PutAppEventType(context.Background(), id).UpdateAppEventTypeDto(updateAppEventTypeDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEventsAPI.PutAppEventType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAppEventType`: AppEventTypeDto
	fmt.Fprintf(os.Stdout, "Response from `AppEventsAPI.PutAppEventType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAppEventTypeRequest struct via the builder pattern


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


## SubscribeToAppEventSubscription

> SubscribeToAppEventSubscription(ctx).Id(id).InstanceKey(instanceKey).AppEventTypeId(appEventTypeId).AppEventTypeKey(appEventTypeKey).AutoCreate(autoCreate).UserId(userId).Execute()



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
	r, err := apiClient.AppEventsAPI.SubscribeToAppEventSubscription(context.Background()).Id(id).InstanceKey(instanceKey).AppEventTypeId(appEventTypeId).AppEventTypeKey(appEventTypeKey).AutoCreate(autoCreate).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEventsAPI.SubscribeToAppEventSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscribeToAppEventSubscriptionRequest struct via the builder pattern


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


## UpdateAppEventSubscription

> AppEventSubscriptionDto UpdateAppEventSubscription(ctx).Body(body).Id(id).InstanceKey(instanceKey).AppEventTypeId(appEventTypeId).AppEventTypeKey(appEventTypeKey).AutoCreate(autoCreate).UserId(userId).Execute()



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
	resp, r, err := apiClient.AppEventsAPI.UpdateAppEventSubscription(context.Background()).Body(body).Id(id).InstanceKey(instanceKey).AppEventTypeId(appEventTypeId).AppEventTypeKey(appEventTypeKey).AutoCreate(autoCreate).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEventsAPI.UpdateAppEventSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAppEventSubscription`: AppEventSubscriptionDto
	fmt.Fprintf(os.Stdout, "Response from `AppEventsAPI.UpdateAppEventSubscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppEventSubscriptionRequest struct via the builder pattern


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

