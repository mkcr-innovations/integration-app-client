# \FieldMappingsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyFieldMapping**](FieldMappingsAPI.md#ApplyFieldMapping) | **Post** /field-mappings/{id}/apply | 
[**ApplyFieldMappings**](FieldMappingsAPI.md#ApplyFieldMappings) | **Post** /field-mapping/apply | 
[**ArchiveFieldMapping**](FieldMappingsAPI.md#ArchiveFieldMapping) | **Delete** /field-mappings/{id} | 
[**ArchiveFieldMappings**](FieldMappingsAPI.md#ArchiveFieldMappings) | **Delete** /field-mapping | 
[**CreateFieldMapping**](FieldMappingsAPI.md#CreateFieldMapping) | **Post** /field-mappings | 
[**DeleteFieldMappingInstance**](FieldMappingsAPI.md#DeleteFieldMappingInstance) | **Delete** /field-mapping-instance | 
[**FieldMappingSetup**](FieldMappingsAPI.md#FieldMappingSetup) | **Post** /field-mapping-instance/setup | 
[**GetFieldMapping**](FieldMappingsAPI.md#GetFieldMapping) | **Get** /field-mappings/{id} | 
[**GetFieldMappingAppSchema**](FieldMappingsAPI.md#GetFieldMappingAppSchema) | **Get** /field-mappings/{id}/app-schema | 
[**GetFieldMappingAppSchemas**](FieldMappingsAPI.md#GetFieldMappingAppSchemas) | **Get** /field-mapping/app-schema | 
[**GetFieldMappingInstance**](FieldMappingsAPI.md#GetFieldMappingInstance) | **Get** /field-mapping-instance | 
[**ListFieldMapping**](FieldMappingsAPI.md#ListFieldMapping) | **Get** /field-mapping | 
[**ListFieldMappingInstances**](FieldMappingsAPI.md#ListFieldMappingInstances) | **Get** /field-mapping-instances | 
[**ListFieldMappings**](FieldMappingsAPI.md#ListFieldMappings) | **Get** /field-mappings | 
[**PatchFieldMapping**](FieldMappingsAPI.md#PatchFieldMapping) | **Patch** /field-mappings/{id} | 
[**PatchFieldMappingInstance**](FieldMappingsAPI.md#PatchFieldMappingInstance) | **Patch** /field-mapping-instance | 
[**PatchFieldMappings**](FieldMappingsAPI.md#PatchFieldMappings) | **Patch** /field-mapping | 
[**PostFieldMappingInstance**](FieldMappingsAPI.md#PostFieldMappingInstance) | **Post** /field-mapping-instance | 
[**PutFieldMapping**](FieldMappingsAPI.md#PutFieldMapping) | **Put** /field-mappings/{id} | 
[**PutFieldMappingInstance**](FieldMappingsAPI.md#PutFieldMappingInstance) | **Put** /field-mapping-instance | 
[**PutFieldMappings**](FieldMappingsAPI.md#PutFieldMappings) | **Put** /field-mapping | 
[**ResetFieldMapping**](FieldMappingsAPI.md#ResetFieldMapping) | **Post** /field-mappings/{id}/reset | 
[**ResetFieldMappingInstance**](FieldMappingsAPI.md#ResetFieldMappingInstance) | **Post** /field-mapping-instance/reset | 
[**ResetFieldMappings**](FieldMappingsAPI.md#ResetFieldMappings) | **Post** /field-mapping/reset | 



## ApplyFieldMapping

> []FieldMappingDto ApplyFieldMapping(ctx, id).Id2(id2).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()



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
	id := "id_example" // string | The ID of the field-mapping to apply
	id2 := "id_example" // string |  (optional)
	key := "key_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldMappingsAPI.ApplyFieldMapping(context.Background(), id).Id2(id2).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldMappingsAPI.ApplyFieldMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApplyFieldMapping`: []FieldMappingDto
	fmt.Fprintf(os.Stdout, "Response from `FieldMappingsAPI.ApplyFieldMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the field-mapping to apply | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplyFieldMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id2** | **string** |  | 
 **key** | **string** |  | 
 **integrationId** | **string** |  | 
 **integrationKey** | **string** |  | 

### Return type

[**[]FieldMappingDto**](FieldMappingDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplyFieldMappings

> []FieldMappingDto ApplyFieldMappings(ctx).Id(id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()



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
	resp, r, err := apiClient.FieldMappingsAPI.ApplyFieldMappings(context.Background()).Id(id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldMappingsAPI.ApplyFieldMappings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApplyFieldMappings`: []FieldMappingDto
	fmt.Fprintf(os.Stdout, "Response from `FieldMappingsAPI.ApplyFieldMappings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApplyFieldMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **key** | **string** |  | 
 **integrationId** | **string** |  | 
 **integrationKey** | **string** |  | 

### Return type

[**[]FieldMappingDto**](FieldMappingDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ArchiveFieldMapping

> ArchiveFieldMapping(ctx, id).Id2(id2).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()



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
	id := "id_example" // string | The ID of the field-mapping to retrive
	id2 := "id_example" // string |  (optional)
	key := "key_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FieldMappingsAPI.ArchiveFieldMapping(context.Background(), id).Id2(id2).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldMappingsAPI.ArchiveFieldMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the field-mapping to retrive | 

### Other Parameters

Other parameters are passed through a pointer to a apiArchiveFieldMappingRequest struct via the builder pattern


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


## ArchiveFieldMappings

> ArchiveFieldMappings(ctx).Id(id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()



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
	r, err := apiClient.FieldMappingsAPI.ArchiveFieldMappings(context.Background()).Id(id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldMappingsAPI.ArchiveFieldMappings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArchiveFieldMappingsRequest struct via the builder pattern


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


## CreateFieldMapping

> FieldMappingDto CreateFieldMapping(ctx).CreateFieldMappingDto(createFieldMappingDto).Execute()



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
	createFieldMappingDto := *openapiclient.NewCreateFieldMappingDto("Key_example", "Name_example") // CreateFieldMappingDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldMappingsAPI.CreateFieldMapping(context.Background()).CreateFieldMappingDto(createFieldMappingDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldMappingsAPI.CreateFieldMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFieldMapping`: FieldMappingDto
	fmt.Fprintf(os.Stdout, "Response from `FieldMappingsAPI.CreateFieldMapping`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFieldMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFieldMappingDto** | [**CreateFieldMappingDto**](CreateFieldMappingDto.md) |  | 

### Return type

[**FieldMappingDto**](FieldMappingDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFieldMappingInstance

> DeleteFieldMappingInstance(ctx).Id(id).FieldMappingKey(fieldMappingKey).FieldMappingId(fieldMappingId).InstanceKey(instanceKey).AutoCreate(autoCreate).FlowKey(flowKey).FlowId(flowId).NodeKey(nodeKey).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()



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
	fieldMappingKey := "fieldMappingKey_example" // string |  (optional)
	fieldMappingId := "fieldMappingId_example" // string |  (optional)
	instanceKey := "instanceKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)
	flowKey := "flowKey_example" // string |  (optional)
	flowId := "flowId_example" // string |  (optional)
	nodeKey := "nodeKey_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	connectionId := "connectionId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FieldMappingsAPI.DeleteFieldMappingInstance(context.Background()).Id(id).FieldMappingKey(fieldMappingKey).FieldMappingId(fieldMappingId).InstanceKey(instanceKey).AutoCreate(autoCreate).FlowKey(flowKey).FlowId(flowId).NodeKey(nodeKey).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldMappingsAPI.DeleteFieldMappingInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFieldMappingInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **fieldMappingKey** | **string** |  | 
 **fieldMappingId** | **string** |  | 
 **instanceKey** | **string** |  | 
 **autoCreate** | **bool** |  | 
 **flowKey** | **string** |  | 
 **flowId** | **string** |  | 
 **nodeKey** | **string** |  | 
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


## FieldMappingSetup

> FieldMappingInstanceDto FieldMappingSetup(ctx).Id(id).FieldMappingKey(fieldMappingKey).FieldMappingId(fieldMappingId).InstanceKey(instanceKey).AutoCreate(autoCreate).FlowKey(flowKey).FlowId(flowId).NodeKey(nodeKey).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()



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
	fieldMappingKey := "fieldMappingKey_example" // string |  (optional)
	fieldMappingId := "fieldMappingId_example" // string |  (optional)
	instanceKey := "instanceKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)
	flowKey := "flowKey_example" // string |  (optional)
	flowId := "flowId_example" // string |  (optional)
	nodeKey := "nodeKey_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	connectionId := "connectionId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldMappingsAPI.FieldMappingSetup(context.Background()).Id(id).FieldMappingKey(fieldMappingKey).FieldMappingId(fieldMappingId).InstanceKey(instanceKey).AutoCreate(autoCreate).FlowKey(flowKey).FlowId(flowId).NodeKey(nodeKey).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldMappingsAPI.FieldMappingSetup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FieldMappingSetup`: FieldMappingInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `FieldMappingsAPI.FieldMappingSetup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFieldMappingSetupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **fieldMappingKey** | **string** |  | 
 **fieldMappingId** | **string** |  | 
 **instanceKey** | **string** |  | 
 **autoCreate** | **bool** |  | 
 **flowKey** | **string** |  | 
 **flowId** | **string** |  | 
 **nodeKey** | **string** |  | 
 **integrationKey** | **string** |  | 
 **integrationId** | **string** |  | 
 **connectionId** | **string** |  | 

### Return type

[**FieldMappingInstanceDto**](FieldMappingInstanceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFieldMapping

> FieldMappingDto GetFieldMapping(ctx, id).Id2(id2).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()



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
	id := "id_example" // string | The ID of the field-mapping to retrive
	id2 := "id_example" // string |  (optional)
	key := "key_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldMappingsAPI.GetFieldMapping(context.Background(), id).Id2(id2).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldMappingsAPI.GetFieldMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFieldMapping`: FieldMappingDto
	fmt.Fprintf(os.Stdout, "Response from `FieldMappingsAPI.GetFieldMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the field-mapping to retrive | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFieldMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id2** | **string** |  | 
 **key** | **string** |  | 
 **integrationId** | **string** |  | 
 **integrationKey** | **string** |  | 

### Return type

[**FieldMappingDto**](FieldMappingDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFieldMappingAppSchema

> map[string]interface{} GetFieldMappingAppSchema(ctx, id).Id2(id2).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()



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
	id := "id_example" // string | The ID of the field-mappping
	id2 := "id_example" // string |  (optional)
	key := "key_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldMappingsAPI.GetFieldMappingAppSchema(context.Background(), id).Id2(id2).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldMappingsAPI.GetFieldMappingAppSchema``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFieldMappingAppSchema`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FieldMappingsAPI.GetFieldMappingAppSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the field-mappping | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFieldMappingAppSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id2** | **string** |  | 
 **key** | **string** |  | 
 **integrationId** | **string** |  | 
 **integrationKey** | **string** |  | 

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


## GetFieldMappingAppSchemas

> map[string]interface{} GetFieldMappingAppSchemas(ctx).Id(id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()



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
	resp, r, err := apiClient.FieldMappingsAPI.GetFieldMappingAppSchemas(context.Background()).Id(id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldMappingsAPI.GetFieldMappingAppSchemas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFieldMappingAppSchemas`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FieldMappingsAPI.GetFieldMappingAppSchemas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFieldMappingAppSchemasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **key** | **string** |  | 
 **integrationId** | **string** |  | 
 **integrationKey** | **string** |  | 

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


## GetFieldMappingInstance

> FieldMappingInstanceDto GetFieldMappingInstance(ctx).Id(id).FieldMappingKey(fieldMappingKey).FieldMappingId(fieldMappingId).InstanceKey(instanceKey).AutoCreate(autoCreate).FlowKey(flowKey).FlowId(flowId).NodeKey(nodeKey).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()



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
	fieldMappingKey := "fieldMappingKey_example" // string |  (optional)
	fieldMappingId := "fieldMappingId_example" // string |  (optional)
	instanceKey := "instanceKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)
	flowKey := "flowKey_example" // string |  (optional)
	flowId := "flowId_example" // string |  (optional)
	nodeKey := "nodeKey_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	connectionId := "connectionId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldMappingsAPI.GetFieldMappingInstance(context.Background()).Id(id).FieldMappingKey(fieldMappingKey).FieldMappingId(fieldMappingId).InstanceKey(instanceKey).AutoCreate(autoCreate).FlowKey(flowKey).FlowId(flowId).NodeKey(nodeKey).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldMappingsAPI.GetFieldMappingInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFieldMappingInstance`: FieldMappingInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `FieldMappingsAPI.GetFieldMappingInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFieldMappingInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **fieldMappingKey** | **string** |  | 
 **fieldMappingId** | **string** |  | 
 **instanceKey** | **string** |  | 
 **autoCreate** | **bool** |  | 
 **flowKey** | **string** |  | 
 **flowId** | **string** |  | 
 **nodeKey** | **string** |  | 
 **integrationKey** | **string** |  | 
 **integrationId** | **string** |  | 
 **connectionId** | **string** |  | 

### Return type

[**FieldMappingInstanceDto**](FieldMappingInstanceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFieldMapping

> FieldMappingDto ListFieldMapping(ctx).Id(id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()



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
	resp, r, err := apiClient.FieldMappingsAPI.ListFieldMapping(context.Background()).Id(id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldMappingsAPI.ListFieldMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFieldMapping`: FieldMappingDto
	fmt.Fprintf(os.Stdout, "Response from `FieldMappingsAPI.ListFieldMapping`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFieldMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **key** | **string** |  | 
 **integrationId** | **string** |  | 
 **integrationKey** | **string** |  | 

### Return type

[**FieldMappingDto**](FieldMappingDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFieldMappingInstances

> ListFieldMappingInstances200Response ListFieldMappingInstances(ctx).Id(id).UserId(userId).FieldMappingId(fieldMappingId).ConnectionId(connectionId).IntegrationId(integrationId).IntegrationKey(integrationKey).InstanceKey(instanceKey).DataSourceInstanceId(dataSourceInstanceId).Execute()



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
	fieldMappingId := "fieldMappingId_example" // string |  (optional)
	connectionId := "connectionId_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	instanceKey := "instanceKey_example" // string |  (optional)
	dataSourceInstanceId := "dataSourceInstanceId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldMappingsAPI.ListFieldMappingInstances(context.Background()).Id(id).UserId(userId).FieldMappingId(fieldMappingId).ConnectionId(connectionId).IntegrationId(integrationId).IntegrationKey(integrationKey).InstanceKey(instanceKey).DataSourceInstanceId(dataSourceInstanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldMappingsAPI.ListFieldMappingInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFieldMappingInstances`: ListFieldMappingInstances200Response
	fmt.Fprintf(os.Stdout, "Response from `FieldMappingsAPI.ListFieldMappingInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFieldMappingInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **userId** | **string** |  | 
 **fieldMappingId** | **string** |  | 
 **connectionId** | **string** |  | 
 **integrationId** | **string** |  | 
 **integrationKey** | **string** |  | 
 **instanceKey** | **string** |  | 
 **dataSourceInstanceId** | **string** |  | 

### Return type

[**ListFieldMappingInstances200Response**](ListFieldMappingInstances200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFieldMappings

> ListFieldMappings200Response ListFieldMappings(ctx).IntegrationId(integrationId).UniversalFieldMappingId(universalFieldMappingId).Execute()



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
	universalFieldMappingId := "universalFieldMappingId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldMappingsAPI.ListFieldMappings(context.Background()).IntegrationId(integrationId).UniversalFieldMappingId(universalFieldMappingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldMappingsAPI.ListFieldMappings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFieldMappings`: ListFieldMappings200Response
	fmt.Fprintf(os.Stdout, "Response from `FieldMappingsAPI.ListFieldMappings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFieldMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **integrationId** | **string** |  | 
 **universalFieldMappingId** | **string** |  | 

### Return type

[**ListFieldMappings200Response**](ListFieldMappings200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchFieldMapping

> FieldMappingDto PatchFieldMapping(ctx, id).UpdateFieldMappingDto(updateFieldMappingDto).Id2(id2).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()



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
	id := "id_example" // string | The ID of the field-mapping to retrive
	updateFieldMappingDto := *openapiclient.NewUpdateFieldMappingDto() // UpdateFieldMappingDto | 
	id2 := "id_example" // string |  (optional)
	key := "key_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldMappingsAPI.PatchFieldMapping(context.Background(), id).UpdateFieldMappingDto(updateFieldMappingDto).Id2(id2).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldMappingsAPI.PatchFieldMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchFieldMapping`: FieldMappingDto
	fmt.Fprintf(os.Stdout, "Response from `FieldMappingsAPI.PatchFieldMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the field-mapping to retrive | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchFieldMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateFieldMappingDto** | [**UpdateFieldMappingDto**](UpdateFieldMappingDto.md) |  | 
 **id2** | **string** |  | 
 **key** | **string** |  | 
 **integrationId** | **string** |  | 
 **integrationKey** | **string** |  | 

### Return type

[**FieldMappingDto**](FieldMappingDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchFieldMappingInstance

> FieldMappingInstanceDto PatchFieldMappingInstance(ctx).UpdateFieldMappingInstanceDto(updateFieldMappingInstanceDto).Id(id).FieldMappingKey(fieldMappingKey).FieldMappingId(fieldMappingId).InstanceKey(instanceKey).AutoCreate(autoCreate).FlowKey(flowKey).FlowId(flowId).NodeKey(nodeKey).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()



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
	updateFieldMappingInstanceDto := *openapiclient.NewUpdateFieldMappingInstanceDto() // UpdateFieldMappingInstanceDto | 
	id := "id_example" // string |  (optional)
	fieldMappingKey := "fieldMappingKey_example" // string |  (optional)
	fieldMappingId := "fieldMappingId_example" // string |  (optional)
	instanceKey := "instanceKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)
	flowKey := "flowKey_example" // string |  (optional)
	flowId := "flowId_example" // string |  (optional)
	nodeKey := "nodeKey_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	connectionId := "connectionId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldMappingsAPI.PatchFieldMappingInstance(context.Background()).UpdateFieldMappingInstanceDto(updateFieldMappingInstanceDto).Id(id).FieldMappingKey(fieldMappingKey).FieldMappingId(fieldMappingId).InstanceKey(instanceKey).AutoCreate(autoCreate).FlowKey(flowKey).FlowId(flowId).NodeKey(nodeKey).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldMappingsAPI.PatchFieldMappingInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchFieldMappingInstance`: FieldMappingInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `FieldMappingsAPI.PatchFieldMappingInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchFieldMappingInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateFieldMappingInstanceDto** | [**UpdateFieldMappingInstanceDto**](UpdateFieldMappingInstanceDto.md) |  | 
 **id** | **string** |  | 
 **fieldMappingKey** | **string** |  | 
 **fieldMappingId** | **string** |  | 
 **instanceKey** | **string** |  | 
 **autoCreate** | **bool** |  | 
 **flowKey** | **string** |  | 
 **flowId** | **string** |  | 
 **nodeKey** | **string** |  | 
 **integrationKey** | **string** |  | 
 **integrationId** | **string** |  | 
 **connectionId** | **string** |  | 

### Return type

[**FieldMappingInstanceDto**](FieldMappingInstanceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchFieldMappings

> FieldMappingDto PatchFieldMappings(ctx).UpdateFieldMappingDto(updateFieldMappingDto).Id(id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()



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
	updateFieldMappingDto := *openapiclient.NewUpdateFieldMappingDto() // UpdateFieldMappingDto | 
	id := "id_example" // string |  (optional)
	key := "key_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldMappingsAPI.PatchFieldMappings(context.Background()).UpdateFieldMappingDto(updateFieldMappingDto).Id(id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldMappingsAPI.PatchFieldMappings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchFieldMappings`: FieldMappingDto
	fmt.Fprintf(os.Stdout, "Response from `FieldMappingsAPI.PatchFieldMappings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchFieldMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateFieldMappingDto** | [**UpdateFieldMappingDto**](UpdateFieldMappingDto.md) |  | 
 **id** | **string** |  | 
 **key** | **string** |  | 
 **integrationId** | **string** |  | 
 **integrationKey** | **string** |  | 

### Return type

[**FieldMappingDto**](FieldMappingDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFieldMappingInstance

> FieldMappingInstanceDto PostFieldMappingInstance(ctx).UpdateFieldMappingInstanceDto(updateFieldMappingInstanceDto).Id(id).FieldMappingKey(fieldMappingKey).FieldMappingId(fieldMappingId).InstanceKey(instanceKey).AutoCreate(autoCreate).FlowKey(flowKey).FlowId(flowId).NodeKey(nodeKey).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()



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
	updateFieldMappingInstanceDto := *openapiclient.NewUpdateFieldMappingInstanceDto() // UpdateFieldMappingInstanceDto | 
	id := "id_example" // string |  (optional)
	fieldMappingKey := "fieldMappingKey_example" // string |  (optional)
	fieldMappingId := "fieldMappingId_example" // string |  (optional)
	instanceKey := "instanceKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)
	flowKey := "flowKey_example" // string |  (optional)
	flowId := "flowId_example" // string |  (optional)
	nodeKey := "nodeKey_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	connectionId := "connectionId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldMappingsAPI.PostFieldMappingInstance(context.Background()).UpdateFieldMappingInstanceDto(updateFieldMappingInstanceDto).Id(id).FieldMappingKey(fieldMappingKey).FieldMappingId(fieldMappingId).InstanceKey(instanceKey).AutoCreate(autoCreate).FlowKey(flowKey).FlowId(flowId).NodeKey(nodeKey).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldMappingsAPI.PostFieldMappingInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFieldMappingInstance`: FieldMappingInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `FieldMappingsAPI.PostFieldMappingInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostFieldMappingInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateFieldMappingInstanceDto** | [**UpdateFieldMappingInstanceDto**](UpdateFieldMappingInstanceDto.md) |  | 
 **id** | **string** |  | 
 **fieldMappingKey** | **string** |  | 
 **fieldMappingId** | **string** |  | 
 **instanceKey** | **string** |  | 
 **autoCreate** | **bool** |  | 
 **flowKey** | **string** |  | 
 **flowId** | **string** |  | 
 **nodeKey** | **string** |  | 
 **integrationKey** | **string** |  | 
 **integrationId** | **string** |  | 
 **connectionId** | **string** |  | 

### Return type

[**FieldMappingInstanceDto**](FieldMappingInstanceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutFieldMapping

> FieldMappingDto PutFieldMapping(ctx, id).UpdateFieldMappingDto(updateFieldMappingDto).Id2(id2).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()



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
	id := "id_example" // string | The ID of the field-mapping to retrive
	updateFieldMappingDto := *openapiclient.NewUpdateFieldMappingDto() // UpdateFieldMappingDto | 
	id2 := "id_example" // string |  (optional)
	key := "key_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldMappingsAPI.PutFieldMapping(context.Background(), id).UpdateFieldMappingDto(updateFieldMappingDto).Id2(id2).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldMappingsAPI.PutFieldMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutFieldMapping`: FieldMappingDto
	fmt.Fprintf(os.Stdout, "Response from `FieldMappingsAPI.PutFieldMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the field-mapping to retrive | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutFieldMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateFieldMappingDto** | [**UpdateFieldMappingDto**](UpdateFieldMappingDto.md) |  | 
 **id2** | **string** |  | 
 **key** | **string** |  | 
 **integrationId** | **string** |  | 
 **integrationKey** | **string** |  | 

### Return type

[**FieldMappingDto**](FieldMappingDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutFieldMappingInstance

> FieldMappingInstanceDto PutFieldMappingInstance(ctx).UpdateFieldMappingInstanceDto(updateFieldMappingInstanceDto).Id(id).FieldMappingKey(fieldMappingKey).FieldMappingId(fieldMappingId).InstanceKey(instanceKey).AutoCreate(autoCreate).FlowKey(flowKey).FlowId(flowId).NodeKey(nodeKey).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()



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
	updateFieldMappingInstanceDto := *openapiclient.NewUpdateFieldMappingInstanceDto() // UpdateFieldMappingInstanceDto | 
	id := "id_example" // string |  (optional)
	fieldMappingKey := "fieldMappingKey_example" // string |  (optional)
	fieldMappingId := "fieldMappingId_example" // string |  (optional)
	instanceKey := "instanceKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)
	flowKey := "flowKey_example" // string |  (optional)
	flowId := "flowId_example" // string |  (optional)
	nodeKey := "nodeKey_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	connectionId := "connectionId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldMappingsAPI.PutFieldMappingInstance(context.Background()).UpdateFieldMappingInstanceDto(updateFieldMappingInstanceDto).Id(id).FieldMappingKey(fieldMappingKey).FieldMappingId(fieldMappingId).InstanceKey(instanceKey).AutoCreate(autoCreate).FlowKey(flowKey).FlowId(flowId).NodeKey(nodeKey).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldMappingsAPI.PutFieldMappingInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutFieldMappingInstance`: FieldMappingInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `FieldMappingsAPI.PutFieldMappingInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutFieldMappingInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateFieldMappingInstanceDto** | [**UpdateFieldMappingInstanceDto**](UpdateFieldMappingInstanceDto.md) |  | 
 **id** | **string** |  | 
 **fieldMappingKey** | **string** |  | 
 **fieldMappingId** | **string** |  | 
 **instanceKey** | **string** |  | 
 **autoCreate** | **bool** |  | 
 **flowKey** | **string** |  | 
 **flowId** | **string** |  | 
 **nodeKey** | **string** |  | 
 **integrationKey** | **string** |  | 
 **integrationId** | **string** |  | 
 **connectionId** | **string** |  | 

### Return type

[**FieldMappingInstanceDto**](FieldMappingInstanceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutFieldMappings

> FieldMappingDto PutFieldMappings(ctx).UpdateFieldMappingDto(updateFieldMappingDto).Id(id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()



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
	updateFieldMappingDto := *openapiclient.NewUpdateFieldMappingDto() // UpdateFieldMappingDto | 
	id := "id_example" // string |  (optional)
	key := "key_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldMappingsAPI.PutFieldMappings(context.Background()).UpdateFieldMappingDto(updateFieldMappingDto).Id(id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldMappingsAPI.PutFieldMappings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutFieldMappings`: FieldMappingDto
	fmt.Fprintf(os.Stdout, "Response from `FieldMappingsAPI.PutFieldMappings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutFieldMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateFieldMappingDto** | [**UpdateFieldMappingDto**](UpdateFieldMappingDto.md) |  | 
 **id** | **string** |  | 
 **key** | **string** |  | 
 **integrationId** | **string** |  | 
 **integrationKey** | **string** |  | 

### Return type

[**FieldMappingDto**](FieldMappingDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetFieldMapping

> FieldMappingDto ResetFieldMapping(ctx, id).Id2(id2).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()



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
	id := "id_example" // string | The ID of the field-mapping to reset
	id2 := "id_example" // string |  (optional)
	key := "key_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldMappingsAPI.ResetFieldMapping(context.Background(), id).Id2(id2).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldMappingsAPI.ResetFieldMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResetFieldMapping`: FieldMappingDto
	fmt.Fprintf(os.Stdout, "Response from `FieldMappingsAPI.ResetFieldMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the field-mapping to reset | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetFieldMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id2** | **string** |  | 
 **key** | **string** |  | 
 **integrationId** | **string** |  | 
 **integrationKey** | **string** |  | 

### Return type

[**FieldMappingDto**](FieldMappingDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetFieldMappingInstance

> FieldMappingInstanceDto ResetFieldMappingInstance(ctx).Id(id).FieldMappingKey(fieldMappingKey).FieldMappingId(fieldMappingId).InstanceKey(instanceKey).AutoCreate(autoCreate).FlowKey(flowKey).FlowId(flowId).NodeKey(nodeKey).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()



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
	fieldMappingKey := "fieldMappingKey_example" // string |  (optional)
	fieldMappingId := "fieldMappingId_example" // string |  (optional)
	instanceKey := "instanceKey_example" // string |  (optional)
	autoCreate := true // bool |  (optional)
	flowKey := "flowKey_example" // string |  (optional)
	flowId := "flowId_example" // string |  (optional)
	nodeKey := "nodeKey_example" // string |  (optional)
	integrationKey := "integrationKey_example" // string |  (optional)
	integrationId := "integrationId_example" // string |  (optional)
	connectionId := "connectionId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldMappingsAPI.ResetFieldMappingInstance(context.Background()).Id(id).FieldMappingKey(fieldMappingKey).FieldMappingId(fieldMappingId).InstanceKey(instanceKey).AutoCreate(autoCreate).FlowKey(flowKey).FlowId(flowId).NodeKey(nodeKey).IntegrationKey(integrationKey).IntegrationId(integrationId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldMappingsAPI.ResetFieldMappingInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResetFieldMappingInstance`: FieldMappingInstanceDto
	fmt.Fprintf(os.Stdout, "Response from `FieldMappingsAPI.ResetFieldMappingInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResetFieldMappingInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **fieldMappingKey** | **string** |  | 
 **fieldMappingId** | **string** |  | 
 **instanceKey** | **string** |  | 
 **autoCreate** | **bool** |  | 
 **flowKey** | **string** |  | 
 **flowId** | **string** |  | 
 **nodeKey** | **string** |  | 
 **integrationKey** | **string** |  | 
 **integrationId** | **string** |  | 
 **connectionId** | **string** |  | 

### Return type

[**FieldMappingInstanceDto**](FieldMappingInstanceDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetFieldMappings

> FieldMappingDto ResetFieldMappings(ctx).Id(id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()



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
	resp, r, err := apiClient.FieldMappingsAPI.ResetFieldMappings(context.Background()).Id(id).Key(key).IntegrationId(integrationId).IntegrationKey(integrationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldMappingsAPI.ResetFieldMappings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResetFieldMappings`: FieldMappingDto
	fmt.Fprintf(os.Stdout, "Response from `FieldMappingsAPI.ResetFieldMappings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResetFieldMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **key** | **string** |  | 
 **integrationId** | **string** |  | 
 **integrationKey** | **string** |  | 

### Return type

[**FieldMappingDto**](FieldMappingDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

