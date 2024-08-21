# \ActionsAPI

All URIs are relative to *https://api.integration.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActionByIdControllerApply**](ActionsAPI.md#ActionByIdControllerApply) | **Post** /actions/{id}/apply | Apply action to integrations
[**ActionByIdControllerArchive**](ActionsAPI.md#ActionByIdControllerArchive) | **Delete** /actions/{id} | Delete action by id
[**ActionByIdControllerClone**](ActionsAPI.md#ActionByIdControllerClone) | **Post** /actions/{id}/clone | 
[**ActionByIdControllerExport**](ActionsAPI.md#ActionByIdControllerExport) | **Get** /actions/{id}/export | 
[**ActionByIdControllerGet**](ActionsAPI.md#ActionByIdControllerGet) | **Get** /actions/{id} | Get action by id
[**ActionByIdControllerPatch**](ActionsAPI.md#ActionByIdControllerPatch) | **Patch** /actions/{id} | Patch action by id
[**ActionByIdControllerPut**](ActionsAPI.md#ActionByIdControllerPut) | **Put** /actions/{id} | Update action by id
[**ActionByIdControllerReset**](ActionsAPI.md#ActionByIdControllerReset) | **Post** /actions/{id}/reset | Reset action by id
[**ActionInstancesControllerListActionInstances**](ActionsAPI.md#ActionInstancesControllerListActionInstances) | **Get** /action-instances | List action instances
[**ActionsControllerCreateAction**](ActionsAPI.md#ActionsControllerCreateAction) | **Post** /actions | Create action
[**ActionsControllerListActions**](ActionsAPI.md#ActionsControllerListActions) | **Get** /actions | List actions
[**ConnectionLevelActionControllerArchive**](ActionsAPI.md#ConnectionLevelActionControllerArchive) | **Delete** /connections/{connectionSelector}/actions/{actionSelector} | Archive action instance for connection
[**ConnectionLevelActionControllerGet**](ActionsAPI.md#ConnectionLevelActionControllerGet) | **Get** /connections/{connectionSelector}/actions/{actionSelector} | Get action instance for connection
[**ConnectionLevelActionControllerPatch**](ActionsAPI.md#ConnectionLevelActionControllerPatch) | **Patch** /connections/{connectionSelector}/actions/{actionSelector} | Patch update action instance for connection
[**ConnectionLevelActionControllerPut**](ActionsAPI.md#ConnectionLevelActionControllerPut) | **Put** /connections/{connectionSelector}/actions/{actionSelector} | Create or Replace action instance
[**ConnectionLevelActionControllerReset**](ActionsAPI.md#ConnectionLevelActionControllerReset) | **Post** /connections/{connectionSelector}/actions/{actionSelector}/reset | Reset action instance for connection
[**ConnectionLevelActionControllerRun**](ActionsAPI.md#ConnectionLevelActionControllerRun) | **Post** /connections/{connectionSelector}/actions/{actionSelector}/run | Run action instance for connection
[**ConnectionLevelActionControllerSetup**](ActionsAPI.md#ConnectionLevelActionControllerSetup) | **Post** /connections/{connectionSelector}/actions/{actionSelector}/setup | Setup action instance for connection
[**ConnectionLevelActionsControllerList**](ActionsAPI.md#ConnectionLevelActionsControllerList) | **Get** /connections/{connectionSelector}/actions | List action instances for connection
[**IntegrationLevelActionControllerArchive**](ActionsAPI.md#IntegrationLevelActionControllerArchive) | **Delete** /integrations/{integrationSelector}/actions/{actionSelector} | Archive action for integration
[**IntegrationLevelActionControllerExport**](ActionsAPI.md#IntegrationLevelActionControllerExport) | **Get** /integrations/{integrationSelector}/actions/{actionSelector}/export | 
[**IntegrationLevelActionControllerGet**](ActionsAPI.md#IntegrationLevelActionControllerGet) | **Get** /integrations/{integrationSelector}/actions/{actionSelector} | Get action for integration
[**IntegrationLevelActionControllerPatch**](ActionsAPI.md#IntegrationLevelActionControllerPatch) | **Patch** /integrations/{integrationSelector}/actions/{actionSelector} | Patch update action for integration
[**IntegrationLevelActionControllerPut**](ActionsAPI.md#IntegrationLevelActionControllerPut) | **Put** /integrations/{integrationSelector}/actions/{actionSelector} | Update action for integration
[**IntegrationLevelActionControllerReset**](ActionsAPI.md#IntegrationLevelActionControllerReset) | **Post** /integrations/{integrationSelector}/actions/{actionSelector}/reset | Reset action for integration
[**IntegrationLevelActionsControllerCreate**](ActionsAPI.md#IntegrationLevelActionsControllerCreate) | **Post** /integrations/{integrationSelector}/actions | Create action for integration
[**IntegrationLevelActionsControllerList**](ActionsAPI.md#IntegrationLevelActionsControllerList) | **Get** /integrations/{integrationSelector}/actions | List actions for integration



## ActionByIdControllerApply

> []ActionDto ActionByIdControllerApply(ctx, id).ApplyToIntegrationsDto(applyToIntegrationsDto).Execute()

Apply action to integrations

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
	applyToIntegrationsDto := *openapiclient.NewApplyToIntegrationsDto() // ApplyToIntegrationsDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionsAPI.ActionByIdControllerApply(context.Background(), id).ApplyToIntegrationsDto(applyToIntegrationsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsAPI.ActionByIdControllerApply``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActionByIdControllerApply`: []ActionDto
	fmt.Fprintf(os.Stdout, "Response from `ActionsAPI.ActionByIdControllerApply`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionByIdControllerApplyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applyToIntegrationsDto** | [**ApplyToIntegrationsDto**](ApplyToIntegrationsDto.md) |  | 

### Return type

[**[]ActionDto**](ActionDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionByIdControllerArchive

> ActionByIdControllerArchive(ctx, id).Execute()

Delete action by id

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ActionsAPI.ActionByIdControllerArchive(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsAPI.ActionByIdControllerArchive``: %v\n", err)
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

Other parameters are passed through a pointer to a apiActionByIdControllerArchiveRequest struct via the builder pattern


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


## ActionByIdControllerClone

> ActionDto ActionByIdControllerClone(ctx, id).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionsAPI.ActionByIdControllerClone(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsAPI.ActionByIdControllerClone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActionByIdControllerClone`: ActionDto
	fmt.Fprintf(os.Stdout, "Response from `ActionsAPI.ActionByIdControllerClone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionByIdControllerCloneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ActionByIdControllerExport

> ActionExportDto ActionByIdControllerExport(ctx, id).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionsAPI.ActionByIdControllerExport(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsAPI.ActionByIdControllerExport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActionByIdControllerExport`: ActionExportDto
	fmt.Fprintf(os.Stdout, "Response from `ActionsAPI.ActionByIdControllerExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionByIdControllerExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ActionExportDto**](ActionExportDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionByIdControllerGet

> ActionDto ActionByIdControllerGet(ctx, id).Execute()

Get action by id

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionsAPI.ActionByIdControllerGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsAPI.ActionByIdControllerGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActionByIdControllerGet`: ActionDto
	fmt.Fprintf(os.Stdout, "Response from `ActionsAPI.ActionByIdControllerGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionByIdControllerGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ActionByIdControllerPatch

> ActionDto ActionByIdControllerPatch(ctx, id).UpdateActionDto(updateActionDto).Execute()

Patch action by id

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
	updateActionDto := *openapiclient.NewUpdateActionDto() // UpdateActionDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionsAPI.ActionByIdControllerPatch(context.Background(), id).UpdateActionDto(updateActionDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsAPI.ActionByIdControllerPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActionByIdControllerPatch`: ActionDto
	fmt.Fprintf(os.Stdout, "Response from `ActionsAPI.ActionByIdControllerPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionByIdControllerPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateActionDto** | [**UpdateActionDto**](UpdateActionDto.md) |  | 

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


## ActionByIdControllerPut

> ActionDto ActionByIdControllerPut(ctx, id).UpdateActionDto(updateActionDto).Execute()

Update action by id

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
	updateActionDto := *openapiclient.NewUpdateActionDto() // UpdateActionDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionsAPI.ActionByIdControllerPut(context.Background(), id).UpdateActionDto(updateActionDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsAPI.ActionByIdControllerPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActionByIdControllerPut`: ActionDto
	fmt.Fprintf(os.Stdout, "Response from `ActionsAPI.ActionByIdControllerPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionByIdControllerPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateActionDto** | [**UpdateActionDto**](UpdateActionDto.md) |  | 

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


## ActionByIdControllerReset

> ActionDto ActionByIdControllerReset(ctx, id).Execute()

Reset action by id

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionsAPI.ActionByIdControllerReset(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsAPI.ActionByIdControllerReset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActionByIdControllerReset`: ActionDto
	fmt.Fprintf(os.Stdout, "Response from `ActionsAPI.ActionByIdControllerReset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionByIdControllerResetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ActionInstancesControllerListActionInstances

> ActionInstancesControllerListActionInstances200Response ActionInstancesControllerListActionInstances(ctx).Limit(limit).Cursor(cursor).Search(search).ConnectorId(connectorId).IncludeArchived(includeArchived).Execute()

List action instances

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
	resp, r, err := apiClient.ActionsAPI.ActionInstancesControllerListActionInstances(context.Background()).Limit(limit).Cursor(cursor).Search(search).ConnectorId(connectorId).IncludeArchived(includeArchived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsAPI.ActionInstancesControllerListActionInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActionInstancesControllerListActionInstances`: ActionInstancesControllerListActionInstances200Response
	fmt.Fprintf(os.Stdout, "Response from `ActionsAPI.ActionInstancesControllerListActionInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiActionInstancesControllerListActionInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **float32** |  | 
 **cursor** | **string** |  | 
 **search** | **string** |  | 
 **connectorId** | **string** |  | 
 **includeArchived** | **bool** |  | 

### Return type

[**ActionInstancesControllerListActionInstances200Response**](ActionInstancesControllerListActionInstances200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsControllerCreateAction

> ActionDto ActionsControllerCreateAction(ctx).CreateActionDto(createActionDto).Execute()

Create action

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
	createActionDto := *openapiclient.NewCreateActionDto("Key_example") // CreateActionDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionsAPI.ActionsControllerCreateAction(context.Background()).CreateActionDto(createActionDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsAPI.ActionsControllerCreateAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActionsControllerCreateAction`: ActionDto
	fmt.Fprintf(os.Stdout, "Response from `ActionsAPI.ActionsControllerCreateAction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiActionsControllerCreateActionRequest struct via the builder pattern


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


## ActionsControllerListActions

> ActionsControllerListActions200Response ActionsControllerListActions(ctx).Limit(limit).Cursor(cursor).Search(search).ConnectorId(connectorId).IncludeArchived(includeArchived).IntegrationId(integrationId).IntegrationKey(integrationKey).ParentId(parentId).Execute()

List actions

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
	integrationId := "integrationId_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	parentId := "parentId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionsAPI.ActionsControllerListActions(context.Background()).Limit(limit).Cursor(cursor).Search(search).ConnectorId(connectorId).IncludeArchived(includeArchived).IntegrationId(integrationId).IntegrationKey(integrationKey).ParentId(parentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsAPI.ActionsControllerListActions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActionsControllerListActions`: ActionsControllerListActions200Response
	fmt.Fprintf(os.Stdout, "Response from `ActionsAPI.ActionsControllerListActions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiActionsControllerListActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **float32** |  | 
 **cursor** | **string** |  | 
 **search** | **string** |  | 
 **connectorId** | **string** |  | 
 **includeArchived** | **bool** |  | 
 **integrationId** | **string** |  | 
 **integrationKey** | **string** |  | 
 **parentId** | **string** |  | 

### Return type

[**ActionsControllerListActions200Response**](ActionsControllerListActions200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionLevelActionControllerArchive

> ConnectionLevelActionControllerArchive(ctx, actionSelector, connectionSelector).InstanceKey(instanceKey).AutoCreate(autoCreate).Execute()

Archive action instance for connection

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
	actionSelector := "actionSelector_example" // string | Action Key or Id
	connectionSelector := "connectionSelector_example" // string | Integration Key, Connection ID, or Integration ID
	instanceKey := "instanceKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ActionsAPI.ConnectionLevelActionControllerArchive(context.Background(), actionSelector, connectionSelector).InstanceKey(instanceKey).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsAPI.ConnectionLevelActionControllerArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actionSelector** | **string** | Action Key or Id | 
**connectionSelector** | **string** | Integration Key, Connection ID, or Integration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionLevelActionControllerArchiveRequest struct via the builder pattern


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


## ConnectionLevelActionControllerGet

> ActionInstanceDto ConnectionLevelActionControllerGet(ctx, actionSelector, connectionSelector).InstanceKey(instanceKey).AutoCreate(autoCreate).Execute()

Get action instance for connection

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
	actionSelector := "actionSelector_example" // string | Action Key or Id
	connectionSelector := "connectionSelector_example" // string | Integration Key, Connection ID, or Integration ID
	instanceKey := "instanceKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionsAPI.ConnectionLevelActionControllerGet(context.Background(), actionSelector, connectionSelector).InstanceKey(instanceKey).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsAPI.ConnectionLevelActionControllerGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectionLevelActionControllerGet`: ActionInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `ActionsAPI.ConnectionLevelActionControllerGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actionSelector** | **string** | Action Key or Id | 
**connectionSelector** | **string** | Integration Key, Connection ID, or Integration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionLevelActionControllerGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **instanceKey** | **string** |  | 
 **autoCreate** | **bool** |  | 

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


## ConnectionLevelActionControllerPatch

> ActionInstanceDto ConnectionLevelActionControllerPatch(ctx, actionSelector, connectionSelector).UpdateActionInstanceDto(updateActionInstanceDto).InstanceKey(instanceKey).AutoCreate(autoCreate).Execute()

Patch update action instance for connection

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
	actionSelector := "actionSelector_example" // string | Action Key or Id
	connectionSelector := "connectionSelector_example" // string | Integration Key, Connection ID, or Integration ID
	updateActionInstanceDto := *openapiclient.NewUpdateActionInstanceDto() // UpdateActionInstanceDto | 
	instanceKey := "instanceKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionsAPI.ConnectionLevelActionControllerPatch(context.Background(), actionSelector, connectionSelector).UpdateActionInstanceDto(updateActionInstanceDto).InstanceKey(instanceKey).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsAPI.ConnectionLevelActionControllerPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectionLevelActionControllerPatch`: ActionInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `ActionsAPI.ConnectionLevelActionControllerPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actionSelector** | **string** | Action Key or Id | 
**connectionSelector** | **string** | Integration Key, Connection ID, or Integration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionLevelActionControllerPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateActionInstanceDto** | [**UpdateActionInstanceDto**](UpdateActionInstanceDto.md) |  | 
 **instanceKey** | **string** |  | 
 **autoCreate** | **bool** |  | 

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


## ConnectionLevelActionControllerPut

> ActionInstanceDto ConnectionLevelActionControllerPut(ctx, actionSelector, connectionSelector).UpdateActionInstanceDto(updateActionInstanceDto).InstanceKey(instanceKey).AutoCreate(autoCreate).Execute()

Create or Replace action instance

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
	actionSelector := "actionSelector_example" // string | Action Key or Id
	connectionSelector := "connectionSelector_example" // string | Integration Key, Connection ID, or Integration ID
	updateActionInstanceDto := *openapiclient.NewUpdateActionInstanceDto() // UpdateActionInstanceDto | 
	instanceKey := "instanceKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionsAPI.ConnectionLevelActionControllerPut(context.Background(), actionSelector, connectionSelector).UpdateActionInstanceDto(updateActionInstanceDto).InstanceKey(instanceKey).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsAPI.ConnectionLevelActionControllerPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectionLevelActionControllerPut`: ActionInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `ActionsAPI.ConnectionLevelActionControllerPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actionSelector** | **string** | Action Key or Id | 
**connectionSelector** | **string** | Integration Key, Connection ID, or Integration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionLevelActionControllerPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateActionInstanceDto** | [**UpdateActionInstanceDto**](UpdateActionInstanceDto.md) |  | 
 **instanceKey** | **string** |  | 
 **autoCreate** | **bool** |  | 

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


## ConnectionLevelActionControllerReset

> ActionInstanceDto ConnectionLevelActionControllerReset(ctx, actionSelector, connectionSelector).InstanceKey(instanceKey).AutoCreate(autoCreate).Execute()

Reset action instance for connection

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
	actionSelector := "actionSelector_example" // string | Action Key or Id
	connectionSelector := "connectionSelector_example" // string | Integration Key, Connection ID, or Integration ID
	instanceKey := "instanceKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionsAPI.ConnectionLevelActionControllerReset(context.Background(), actionSelector, connectionSelector).InstanceKey(instanceKey).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsAPI.ConnectionLevelActionControllerReset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectionLevelActionControllerReset`: ActionInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `ActionsAPI.ConnectionLevelActionControllerReset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actionSelector** | **string** | Action Key or Id | 
**connectionSelector** | **string** | Integration Key, Connection ID, or Integration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionLevelActionControllerResetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **instanceKey** | **string** |  | 
 **autoCreate** | **bool** |  | 

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


## ConnectionLevelActionControllerRun

> ConnectionLevelActionControllerRun(ctx, actionSelector, connectionSelector).RequestBody(requestBody).InstanceKey(instanceKey).AutoCreate(autoCreate).Execute()

Run action instance for connection

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
	actionSelector := "actionSelector_example" // string | Action Key or Id
	connectionSelector := "connectionSelector_example" // string | Integration Key, Connection ID, or Integration ID
	requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | Request body to be passed
	instanceKey := "instanceKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ActionsAPI.ConnectionLevelActionControllerRun(context.Background(), actionSelector, connectionSelector).RequestBody(requestBody).InstanceKey(instanceKey).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsAPI.ConnectionLevelActionControllerRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actionSelector** | **string** | Action Key or Id | 
**connectionSelector** | **string** | Integration Key, Connection ID, or Integration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionLevelActionControllerRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestBody** | **map[string]interface{}** | Request body to be passed | 
 **instanceKey** | **string** |  | 
 **autoCreate** | **bool** |  | 

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


## ConnectionLevelActionControllerSetup

> ActionInstanceDto ConnectionLevelActionControllerSetup(ctx, actionSelector, connectionSelector).InstanceKey(instanceKey).AutoCreate(autoCreate).Execute()

Setup action instance for connection

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
	actionSelector := "actionSelector_example" // string | Action Key or Id
	connectionSelector := "connectionSelector_example" // string | Integration Key, Connection ID, or Integration ID
	instanceKey := "instanceKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionsAPI.ConnectionLevelActionControllerSetup(context.Background(), actionSelector, connectionSelector).InstanceKey(instanceKey).AutoCreate(autoCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsAPI.ConnectionLevelActionControllerSetup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectionLevelActionControllerSetup`: ActionInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `ActionsAPI.ConnectionLevelActionControllerSetup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actionSelector** | **string** | Action Key or Id | 
**connectionSelector** | **string** | Integration Key, Connection ID, or Integration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionLevelActionControllerSetupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **instanceKey** | **string** |  | 
 **autoCreate** | **bool** |  | 

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


## ConnectionLevelActionsControllerList

> ConnectionLevelActionsControllerList(ctx, connectionSelector).Limit(limit).Cursor(cursor).Search(search).ConnectorId(connectorId).IncludeArchived(includeArchived).ParentId(parentId).UniversalParentId(universalParentId).Execute()

List action instances for connection

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
	parentId := "parentId_example" // string |  (optional)
	universalParentId := "universalParentId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ActionsAPI.ConnectionLevelActionsControllerList(context.Background(), connectionSelector).Limit(limit).Cursor(cursor).Search(search).ConnectorId(connectorId).IncludeArchived(includeArchived).ParentId(parentId).UniversalParentId(universalParentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsAPI.ConnectionLevelActionsControllerList``: %v\n", err)
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

Other parameters are passed through a pointer to a apiConnectionLevelActionsControllerListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **float32** |  | 
 **cursor** | **string** |  | 
 **search** | **string** |  | 
 **connectorId** | **string** |  | 
 **includeArchived** | **bool** |  | 
 **parentId** | **string** |  | 
 **universalParentId** | **string** |  | 

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


## IntegrationLevelActionControllerArchive

> IntegrationLevelActionControllerArchive(ctx, actionSelector, integrationSelector).Execute()

Archive action for integration

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
	actionSelector := "actionSelector_example" // string | Action Key or Id
	integrationSelector := "integrationSelector_example" // string | Integration ID or Key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ActionsAPI.IntegrationLevelActionControllerArchive(context.Background(), actionSelector, integrationSelector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsAPI.IntegrationLevelActionControllerArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actionSelector** | **string** | Action Key or Id | 
**integrationSelector** | **string** | Integration ID or Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationLevelActionControllerArchiveRequest struct via the builder pattern


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


## IntegrationLevelActionControllerExport

> ActionExportDto IntegrationLevelActionControllerExport(ctx, actionSelector, integrationSelector).Execute()



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
	actionSelector := "actionSelector_example" // string | 
	integrationSelector := "integrationSelector_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionsAPI.IntegrationLevelActionControllerExport(context.Background(), actionSelector, integrationSelector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsAPI.IntegrationLevelActionControllerExport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationLevelActionControllerExport`: ActionExportDto
	fmt.Fprintf(os.Stdout, "Response from `ActionsAPI.IntegrationLevelActionControllerExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actionSelector** | **string** |  | 
**integrationSelector** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationLevelActionControllerExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ActionExportDto**](ActionExportDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationLevelActionControllerGet

> ActionDto IntegrationLevelActionControllerGet(ctx, actionSelector, integrationSelector).Execute()

Get action for integration

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
	actionSelector := "actionSelector_example" // string | Action Key or Id
	integrationSelector := "integrationSelector_example" // string | Integration ID or Key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionsAPI.IntegrationLevelActionControllerGet(context.Background(), actionSelector, integrationSelector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsAPI.IntegrationLevelActionControllerGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationLevelActionControllerGet`: ActionDto
	fmt.Fprintf(os.Stdout, "Response from `ActionsAPI.IntegrationLevelActionControllerGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actionSelector** | **string** | Action Key or Id | 
**integrationSelector** | **string** | Integration ID or Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationLevelActionControllerGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## IntegrationLevelActionControllerPatch

> ActionDto IntegrationLevelActionControllerPatch(ctx, actionSelector, integrationSelector).UpdateActionDto(updateActionDto).Execute()

Patch update action for integration

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
	actionSelector := "actionSelector_example" // string | Action Key or Id
	integrationSelector := "integrationSelector_example" // string | Integration ID or Key
	updateActionDto := *openapiclient.NewUpdateActionDto() // UpdateActionDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionsAPI.IntegrationLevelActionControllerPatch(context.Background(), actionSelector, integrationSelector).UpdateActionDto(updateActionDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsAPI.IntegrationLevelActionControllerPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationLevelActionControllerPatch`: ActionDto
	fmt.Fprintf(os.Stdout, "Response from `ActionsAPI.IntegrationLevelActionControllerPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actionSelector** | **string** | Action Key or Id | 
**integrationSelector** | **string** | Integration ID or Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationLevelActionControllerPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateActionDto** | [**UpdateActionDto**](UpdateActionDto.md) |  | 

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


## IntegrationLevelActionControllerPut

> ActionDto IntegrationLevelActionControllerPut(ctx, actionSelector, integrationSelector).UpdateActionDto(updateActionDto).Execute()

Update action for integration

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
	actionSelector := "actionSelector_example" // string | Action Key or Id
	integrationSelector := "integrationSelector_example" // string | Integration ID or Key
	updateActionDto := *openapiclient.NewUpdateActionDto() // UpdateActionDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionsAPI.IntegrationLevelActionControllerPut(context.Background(), actionSelector, integrationSelector).UpdateActionDto(updateActionDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsAPI.IntegrationLevelActionControllerPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationLevelActionControllerPut`: ActionDto
	fmt.Fprintf(os.Stdout, "Response from `ActionsAPI.IntegrationLevelActionControllerPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actionSelector** | **string** | Action Key or Id | 
**integrationSelector** | **string** | Integration ID or Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationLevelActionControllerPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateActionDto** | [**UpdateActionDto**](UpdateActionDto.md) |  | 

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


## IntegrationLevelActionControllerReset

> ActionDto IntegrationLevelActionControllerReset(ctx, actionSelector, integrationSelector).Execute()

Reset action for integration

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
	actionSelector := "actionSelector_example" // string | Action Key or Id
	integrationSelector := "integrationSelector_example" // string | Integration ID or Key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionsAPI.IntegrationLevelActionControllerReset(context.Background(), actionSelector, integrationSelector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsAPI.IntegrationLevelActionControllerReset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationLevelActionControllerReset`: ActionDto
	fmt.Fprintf(os.Stdout, "Response from `ActionsAPI.IntegrationLevelActionControllerReset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actionSelector** | **string** | Action Key or Id | 
**integrationSelector** | **string** | Integration ID or Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationLevelActionControllerResetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## IntegrationLevelActionsControllerCreate

> ActionDto IntegrationLevelActionsControllerCreate(ctx, integrationSelector).CreateIntegrationLevelActionDto(createIntegrationLevelActionDto).Execute()

Create action for integration

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
	createIntegrationLevelActionDto := *openapiclient.NewCreateIntegrationLevelActionDto("Key_example") // CreateIntegrationLevelActionDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionsAPI.IntegrationLevelActionsControllerCreate(context.Background(), integrationSelector).CreateIntegrationLevelActionDto(createIntegrationLevelActionDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsAPI.IntegrationLevelActionsControllerCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationLevelActionsControllerCreate`: ActionDto
	fmt.Fprintf(os.Stdout, "Response from `ActionsAPI.IntegrationLevelActionsControllerCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationSelector** | **string** | Integration ID or Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationLevelActionsControllerCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createIntegrationLevelActionDto** | [**CreateIntegrationLevelActionDto**](CreateIntegrationLevelActionDto.md) |  | 

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


## IntegrationLevelActionsControllerList

> ActionsControllerListActions200Response IntegrationLevelActionsControllerList(ctx, integrationSelector).Limit(limit).Cursor(cursor).Search(search).ConnectorId(connectorId).IncludeArchived(includeArchived).ParentId(parentId).Execute()

List actions for integration

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
	parentId := "parentId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionsAPI.IntegrationLevelActionsControllerList(context.Background(), integrationSelector).Limit(limit).Cursor(cursor).Search(search).ConnectorId(connectorId).IncludeArchived(includeArchived).ParentId(parentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsAPI.IntegrationLevelActionsControllerList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationLevelActionsControllerList`: ActionsControllerListActions200Response
	fmt.Fprintf(os.Stdout, "Response from `ActionsAPI.IntegrationLevelActionsControllerList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationSelector** | **string** | Integration ID or Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationLevelActionsControllerListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **float32** |  | 
 **cursor** | **string** |  | 
 **search** | **string** |  | 
 **connectorId** | **string** |  | 
 **includeArchived** | **bool** |  | 
 **parentId** | **string** |  | 

### Return type

[**ActionsControllerListActions200Response**](ActionsControllerListActions200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

