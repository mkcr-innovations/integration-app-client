# \ActionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActionInstanceSetup**](ActionsAPI.md#ActionInstanceSetup) | **Post** /action-instance/setup | 
[**ApplyAction**](ActionsAPI.md#ApplyAction) | **Post** /actions/{id}/apply | 
[**ApplyActions**](ActionsAPI.md#ApplyActions) | **Post** /action/apply | 
[**ArchiveAction**](ActionsAPI.md#ArchiveAction) | **Delete** /actions/{id} | 
[**ArchiveActions**](ActionsAPI.md#ArchiveActions) | **Delete** /action | 
[**CreateAction**](ActionsAPI.md#CreateAction) | **Post** /actions | 
[**DeleteActionInstance**](ActionsAPI.md#DeleteActionInstance) | **Delete** /action-instance | 
[**GetAction**](ActionsAPI.md#GetAction) | **Get** /actions/{id} | 
[**GetActionInstance**](ActionsAPI.md#GetActionInstance) | **Get** /action-instance | 
[**GetActions**](ActionsAPI.md#GetActions) | **Get** /action | 
[**ListActionInstances**](ActionsAPI.md#ListActionInstances) | **Get** /action-instances | 
[**ListActions**](ActionsAPI.md#ListActions) | **Get** /actions | 
[**PatchAction**](ActionsAPI.md#PatchAction) | **Patch** /actions/{id} | 
[**PatchActionInstance**](ActionsAPI.md#PatchActionInstance) | **Patch** /action-instance | 
[**PatchActions**](ActionsAPI.md#PatchActions) | **Patch** /action | 
[**PostActionInstance**](ActionsAPI.md#PostActionInstance) | **Post** /action-instance | 
[**PutAction**](ActionsAPI.md#PutAction) | **Put** /actions/{id} | 
[**PutActionInstance**](ActionsAPI.md#PutActionInstance) | **Put** /action-instance | 
[**PutActions**](ActionsAPI.md#PutActions) | **Put** /action | 
[**ResetAction**](ActionsAPI.md#ResetAction) | **Post** /actions/{id}/reset | 
[**ResetActionInstance**](ActionsAPI.md#ResetActionInstance) | **Post** /action-instance/reset | 
[**ResetActions**](ActionsAPI.md#ResetActions) | **Post** /action/reset | 
[**RunActionInstance**](ActionsAPI.md#RunActionInstance) | **Post** /action-instance/run | 



## ActionInstanceSetup

> ActionInstanceDto ActionInstanceSetup(ctx).Id(id).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()



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
	integrationKey := "integrationKey_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	connectionId := "connectionId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionsAPI.ActionInstanceSetup(context.Background()).Id(id).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsAPI.ActionInstanceSetup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActionInstanceSetup`: ActionInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `ActionsAPI.ActionInstanceSetup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiActionInstanceSetupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **integrationKey** | **string** |  | 
 **integrationId** | **string** |  | 
 **connectionId** | **string** |  | 

### Return type

[**ActionInstanceDto**](ActionInstanceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplyAction

> []ActionDto ApplyAction(ctx, id).AutoCreate(autoCreate).Execute()



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
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionsAPI.ApplyAction(context.Background(), id).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsAPI.ApplyAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApplyAction`: []ActionDto
	fmt.Fprintf(os.Stdout, "Response from `ActionsAPI.ApplyAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplyActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **autoCreate** | **bool** |  | 

### Return type

[**[]ActionDto**](ActionDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplyActions

> []ActionDto ApplyActions(ctx).AutoCreate(autoCreate).Execute()



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
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionsAPI.ApplyActions(context.Background()).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsAPI.ApplyActions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApplyActions`: []ActionDto
	fmt.Fprintf(os.Stdout, "Response from `ActionsAPI.ApplyActions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApplyActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoCreate** | **bool** |  | 

### Return type

[**[]ActionDto**](ActionDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ArchiveAction

> ArchiveAction(ctx, id).AutoCreate(autoCreate).Execute()



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
	id := "id_example" // string | The ID of the action to retrive
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ActionsAPI.ArchiveAction(context.Background(), id).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsAPI.ArchiveAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the action to retrive | 

### Other Parameters

Other parameters are passed through a pointer to a apiArchiveActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


## ArchiveActions

> ArchiveActions(ctx).AutoCreate(autoCreate).Execute()



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
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ActionsAPI.ArchiveActions(context.Background()).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsAPI.ArchiveActions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArchiveActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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


## CreateAction

> ActionDto CreateAction(ctx).CreateActionDto(createActionDto).Execute()



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
	createActionDto := *openapiclient.NewCreateActionDto("Key_example", "Name_example") // CreateActionDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionsAPI.CreateAction(context.Background()).CreateActionDto(createActionDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsAPI.CreateAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAction`: ActionDto
	fmt.Fprintf(os.Stdout, "Response from `ActionsAPI.CreateAction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createActionDto** | [**CreateActionDto**](CreateActionDto.md) |  | 

### Return type

[**ActionDto**](ActionDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteActionInstance

> DeleteActionInstance(ctx).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()



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
	integrationKey := "integrationKey_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	connectionId := "connectionId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ActionsAPI.DeleteActionInstance(context.Background()).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsAPI.DeleteActionInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteActionInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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


## GetAction

> ActionDto GetAction(ctx, id).AutoCreate(autoCreate).Execute()



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
	id := "id_example" // string | The ID of the action to retrive
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionsAPI.GetAction(context.Background(), id).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsAPI.GetAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAction`: ActionDto
	fmt.Fprintf(os.Stdout, "Response from `ActionsAPI.GetAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the action to retrive | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **autoCreate** | **bool** |  | 

### Return type

[**ActionDto**](ActionDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActionInstance

> ActionInstanceDto GetActionInstance(ctx).Id(id).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).AutoCreate(autoCreate).ParentId(parentId).Execute()



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
	integrationKey := "integrationKey_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	connectionId := "connectionId_example" // string |  (optional)
	autoCreate := true // bool |  (optional)
	parentId := "parentId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionsAPI.GetActionInstance(context.Background()).Id(id).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).AutoCreate(autoCreate).ParentId(parentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsAPI.GetActionInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActionInstance`: ActionInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `ActionsAPI.GetActionInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetActionInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **integrationKey** | **string** |  | 
 **integrationId** | **string** |  | 
 **connectionId** | **string** |  | 
 **autoCreate** | **bool** |  | 
 **parentId** | **string** |  | 

### Return type

[**ActionInstanceDto**](ActionInstanceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActions

> ActionDto GetActions(ctx).AutoCreate(autoCreate).Execute()



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
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionsAPI.GetActions(context.Background()).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsAPI.GetActions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActions`: ActionDto
	fmt.Fprintf(os.Stdout, "Response from `ActionsAPI.GetActions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoCreate** | **bool** |  | 

### Return type

[**ActionDto**](ActionDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListActionInstances

> ListActionInstances200Response ListActionInstances(ctx).Execute()



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
	resp, r, err := apiClient.ActionsAPI.ListActionInstances(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsAPI.ListActionInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListActionInstances`: ListActionInstances200Response
	fmt.Fprintf(os.Stdout, "Response from `ActionsAPI.ListActionInstances`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListActionInstancesRequest struct via the builder pattern


### Return type

[**ListActionInstances200Response**](ListActionInstances200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListActions

> ListActions200Response ListActions(ctx).IntegrationId(integrationId).IntegrationKey(integrationKey).ParentId(parentId).Execute()



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
	integrationKey := "integrationKey_example" // string |  (optional)
	parentId := "parentId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionsAPI.ListActions(context.Background()).IntegrationId(integrationId).IntegrationKey(integrationKey).ParentId(parentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsAPI.ListActions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListActions`: ListActions200Response
	fmt.Fprintf(os.Stdout, "Response from `ActionsAPI.ListActions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **integrationId** | **string** |  | 
 **integrationKey** | **string** |  | 
 **parentId** | **string** |  | 

### Return type

[**ListActions200Response**](ListActions200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAction

> ActionDto PatchAction(ctx, id).UpdateActionDto(updateActionDto).AutoCreate(autoCreate).Execute()



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
	id := "id_example" // string | The ID of the action to retrive
	updateActionDto := *openapiclient.NewUpdateActionDto() // UpdateActionDto | 
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionsAPI.PatchAction(context.Background(), id).UpdateActionDto(updateActionDto).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsAPI.PatchAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAction`: ActionDto
	fmt.Fprintf(os.Stdout, "Response from `ActionsAPI.PatchAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the action to retrive | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateActionDto** | [**UpdateActionDto**](UpdateActionDto.md) |  | 
 **autoCreate** | **bool** |  | 

### Return type

[**ActionDto**](ActionDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchActionInstance

> ActionInstanceDto PatchActionInstance(ctx).UpdateActionInstanceDto(updateActionInstanceDto).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()



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
	updateActionInstanceDto := *openapiclient.NewUpdateActionInstanceDto() // UpdateActionInstanceDto | 
	integrationKey := "integrationKey_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	connectionId := "connectionId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionsAPI.PatchActionInstance(context.Background()).UpdateActionInstanceDto(updateActionInstanceDto).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsAPI.PatchActionInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchActionInstance`: ActionInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `ActionsAPI.PatchActionInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchActionInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateActionInstanceDto** | [**UpdateActionInstanceDto**](UpdateActionInstanceDto.md) |  | 
 **integrationKey** | **string** |  | 
 **integrationId** | **string** |  | 
 **connectionId** | **string** |  | 

### Return type

[**ActionInstanceDto**](ActionInstanceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchActions

> ActionDto PatchActions(ctx).UpdateActionDto(updateActionDto).AutoCreate(autoCreate).Execute()



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
	updateActionDto := *openapiclient.NewUpdateActionDto() // UpdateActionDto | 
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionsAPI.PatchActions(context.Background()).UpdateActionDto(updateActionDto).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsAPI.PatchActions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchActions`: ActionDto
	fmt.Fprintf(os.Stdout, "Response from `ActionsAPI.PatchActions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateActionDto** | [**UpdateActionDto**](UpdateActionDto.md) |  | 
 **autoCreate** | **bool** |  | 

### Return type

[**ActionDto**](ActionDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostActionInstance

> ActionInstanceDto PostActionInstance(ctx).UpdateActionInstanceDto(updateActionInstanceDto).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()



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
	updateActionInstanceDto := *openapiclient.NewUpdateActionInstanceDto() // UpdateActionInstanceDto | 
	integrationKey := "integrationKey_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	connectionId := "connectionId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionsAPI.PostActionInstance(context.Background()).UpdateActionInstanceDto(updateActionInstanceDto).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsAPI.PostActionInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostActionInstance`: ActionInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `ActionsAPI.PostActionInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostActionInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateActionInstanceDto** | [**UpdateActionInstanceDto**](UpdateActionInstanceDto.md) |  | 
 **integrationKey** | **string** |  | 
 **integrationId** | **string** |  | 
 **connectionId** | **string** |  | 

### Return type

[**ActionInstanceDto**](ActionInstanceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAction

> ActionDto PutAction(ctx, id).UpdateActionDto(updateActionDto).AutoCreate(autoCreate).Execute()



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
	id := "id_example" // string | The ID of the action to retrive
	updateActionDto := *openapiclient.NewUpdateActionDto() // UpdateActionDto | 
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionsAPI.PutAction(context.Background(), id).UpdateActionDto(updateActionDto).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsAPI.PutAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAction`: ActionDto
	fmt.Fprintf(os.Stdout, "Response from `ActionsAPI.PutAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the action to retrive | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateActionDto** | [**UpdateActionDto**](UpdateActionDto.md) |  | 
 **autoCreate** | **bool** |  | 

### Return type

[**ActionDto**](ActionDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutActionInstance

> ActionInstanceDto PutActionInstance(ctx).UpdateActionInstanceDto(updateActionInstanceDto).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()



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
	updateActionInstanceDto := *openapiclient.NewUpdateActionInstanceDto() // UpdateActionInstanceDto | 
	integrationKey := "integrationKey_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	connectionId := "connectionId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionsAPI.PutActionInstance(context.Background()).UpdateActionInstanceDto(updateActionInstanceDto).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsAPI.PutActionInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutActionInstance`: ActionInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `ActionsAPI.PutActionInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutActionInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateActionInstanceDto** | [**UpdateActionInstanceDto**](UpdateActionInstanceDto.md) |  | 
 **integrationKey** | **string** |  | 
 **integrationId** | **string** |  | 
 **connectionId** | **string** |  | 

### Return type

[**ActionInstanceDto**](ActionInstanceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutActions

> ActionDto PutActions(ctx).UpdateActionDto(updateActionDto).AutoCreate(autoCreate).Execute()



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
	updateActionDto := *openapiclient.NewUpdateActionDto() // UpdateActionDto | 
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionsAPI.PutActions(context.Background()).UpdateActionDto(updateActionDto).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsAPI.PutActions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutActions`: ActionDto
	fmt.Fprintf(os.Stdout, "Response from `ActionsAPI.PutActions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateActionDto** | [**UpdateActionDto**](UpdateActionDto.md) |  | 
 **autoCreate** | **bool** |  | 

### Return type

[**ActionDto**](ActionDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetAction

> ActionDto ResetAction(ctx, id).AutoCreate(autoCreate).Execute()



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
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionsAPI.ResetAction(context.Background(), id).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsAPI.ResetAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResetAction`: ActionDto
	fmt.Fprintf(os.Stdout, "Response from `ActionsAPI.ResetAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **autoCreate** | **bool** |  | 

### Return type

[**ActionDto**](ActionDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetActionInstance

> ActionInstanceDto ResetActionInstance(ctx).Id(id).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()



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
	integrationKey := "integrationKey_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	connectionId := "connectionId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionsAPI.ResetActionInstance(context.Background()).Id(id).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsAPI.ResetActionInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResetActionInstance`: ActionInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `ActionsAPI.ResetActionInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResetActionInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **integrationKey** | **string** |  | 
 **integrationId** | **string** |  | 
 **connectionId** | **string** |  | 

### Return type

[**ActionInstanceDto**](ActionInstanceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetActions

> ActionDto ResetActions(ctx).AutoCreate(autoCreate).Execute()



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
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionsAPI.ResetActions(context.Background()).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsAPI.ResetActions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResetActions`: ActionDto
	fmt.Fprintf(os.Stdout, "Response from `ActionsAPI.ResetActions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResetActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoCreate** | **bool** |  | 

### Return type

[**ActionDto**](ActionDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunActionInstance

> RunActionInstance(ctx).RequestBody(requestBody).Id(id).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()



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
	requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | 
	id := "id_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	connectionId := "connectionId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ActionsAPI.RunActionInstance(context.Background()).RequestBody(requestBody).Id(id).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsAPI.RunActionInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRunActionInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]interface{}** |  | 
 **id** | **string** |  | 
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

