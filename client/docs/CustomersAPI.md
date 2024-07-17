# \CustomersAPI

All URIs are relative to *https://api.integration.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomersControllerCreateCustomer**](CustomersAPI.md#CustomersControllerCreateCustomer) | **Post** /customers | Create customer
[**CustomersControllerDeleteCustomer**](CustomersAPI.md#CustomersControllerDeleteCustomer) | **Delete** /customers/{id} | Delete customer
[**CustomersControllerGetCustomer**](CustomersAPI.md#CustomersControllerGetCustomer) | **Get** /customers/{id} | Get customer
[**CustomersControllerListCustomers**](CustomersAPI.md#CustomersControllerListCustomers) | **Get** /customers | List customers
[**CustomersControllerPatchCustomer**](CustomersAPI.md#CustomersControllerPatchCustomer) | **Patch** /customers/{id} | Patch customer
[**CustomersControllerPutCustomer**](CustomersAPI.md#CustomersControllerPutCustomer) | **Put** /customers/{id} | Update customer



## CustomersControllerCreateCustomer

> CustomerDto CustomersControllerCreateCustomer(ctx).CreateCustomerDto(createCustomerDto).Execute()

Create customer

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
	createCustomerDto := *openapiclient.NewCreateCustomerDto("InternalId_example") // CreateCustomerDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersAPI.CustomersControllerCreateCustomer(context.Background()).CreateCustomerDto(createCustomerDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.CustomersControllerCreateCustomer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomersControllerCreateCustomer`: CustomerDto
	fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.CustomersControllerCreateCustomer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomersControllerCreateCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCustomerDto** | [**CreateCustomerDto**](CreateCustomerDto.md) |  | 

### Return type

[**CustomerDto**](CustomerDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomersControllerDeleteCustomer

> CustomersControllerDeleteCustomer(ctx, id).Execute()

Delete customer

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
	r, err := apiClient.CustomersAPI.CustomersControllerDeleteCustomer(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.CustomersControllerDeleteCustomer``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCustomersControllerDeleteCustomerRequest struct via the builder pattern


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


## CustomersControllerGetCustomer

> CustomerDto CustomersControllerGetCustomer(ctx, id).Execute()

Get customer

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
	resp, r, err := apiClient.CustomersAPI.CustomersControllerGetCustomer(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.CustomersControllerGetCustomer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomersControllerGetCustomer`: CustomerDto
	fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.CustomersControllerGetCustomer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomersControllerGetCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomerDto**](CustomerDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomersControllerListCustomers

> CustomersControllerListCustomers200Response CustomersControllerListCustomers(ctx).Limit(limit).Cursor(cursor).Search(search).ConnectorId(connectorId).IsTest(isTest).Execute()

List customers

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
	isTest := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersAPI.CustomersControllerListCustomers(context.Background()).Limit(limit).Cursor(cursor).Search(search).ConnectorId(connectorId).IsTest(isTest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.CustomersControllerListCustomers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomersControllerListCustomers`: CustomersControllerListCustomers200Response
	fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.CustomersControllerListCustomers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomersControllerListCustomersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **float32** |  | 
 **cursor** | **string** |  | 
 **search** | **string** |  | 
 **connectorId** | **string** |  | 
 **isTest** | **bool** |  | 

### Return type

[**CustomersControllerListCustomers200Response**](CustomersControllerListCustomers200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomersControllerPatchCustomer

> CustomerDto CustomersControllerPatchCustomer(ctx, id).UpdateCustomerAsAdminDto(updateCustomerAsAdminDto).Execute()

Patch customer

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
	updateCustomerAsAdminDto := *openapiclient.NewUpdateCustomerAsAdminDto() // UpdateCustomerAsAdminDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersAPI.CustomersControllerPatchCustomer(context.Background(), id).UpdateCustomerAsAdminDto(updateCustomerAsAdminDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.CustomersControllerPatchCustomer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomersControllerPatchCustomer`: CustomerDto
	fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.CustomersControllerPatchCustomer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomersControllerPatchCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateCustomerAsAdminDto** | [**UpdateCustomerAsAdminDto**](UpdateCustomerAsAdminDto.md) |  | 

### Return type

[**CustomerDto**](CustomerDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomersControllerPutCustomer

> CustomerDto CustomersControllerPutCustomer(ctx, id).UpdateCustomerAsAdminDto(updateCustomerAsAdminDto).Execute()

Update customer

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
	updateCustomerAsAdminDto := *openapiclient.NewUpdateCustomerAsAdminDto() // UpdateCustomerAsAdminDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersAPI.CustomersControllerPutCustomer(context.Background(), id).UpdateCustomerAsAdminDto(updateCustomerAsAdminDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.CustomersControllerPutCustomer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomersControllerPutCustomer`: CustomerDto
	fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.CustomersControllerPutCustomer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomersControllerPutCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateCustomerAsAdminDto** | [**UpdateCustomerAsAdminDto**](UpdateCustomerAsAdminDto.md) |  | 

### Return type

[**CustomerDto**](CustomerDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

