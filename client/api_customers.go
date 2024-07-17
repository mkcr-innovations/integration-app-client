/*
Integration.app API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


type CustomersAPI interface {

	/*
	CustomersControllerCreateCustomer Create customer

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCustomersControllerCreateCustomerRequest
	*/
	CustomersControllerCreateCustomer(ctx context.Context) ApiCustomersControllerCreateCustomerRequest

	// CustomersControllerCreateCustomerExecute executes the request
	//  @return CustomerDto
	CustomersControllerCreateCustomerExecute(r ApiCustomersControllerCreateCustomerRequest) (*CustomerDto, *http.Response, error)

	/*
	CustomersControllerDeleteCustomer Delete customer

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@return ApiCustomersControllerDeleteCustomerRequest
	*/
	CustomersControllerDeleteCustomer(ctx context.Context, id string) ApiCustomersControllerDeleteCustomerRequest

	// CustomersControllerDeleteCustomerExecute executes the request
	CustomersControllerDeleteCustomerExecute(r ApiCustomersControllerDeleteCustomerRequest) (*http.Response, error)

	/*
	CustomersControllerGetCustomer Get customer

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@return ApiCustomersControllerGetCustomerRequest
	*/
	CustomersControllerGetCustomer(ctx context.Context, id string) ApiCustomersControllerGetCustomerRequest

	// CustomersControllerGetCustomerExecute executes the request
	//  @return CustomerDto
	CustomersControllerGetCustomerExecute(r ApiCustomersControllerGetCustomerRequest) (*CustomerDto, *http.Response, error)

	/*
	CustomersControllerListCustomers List customers

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCustomersControllerListCustomersRequest
	*/
	CustomersControllerListCustomers(ctx context.Context) ApiCustomersControllerListCustomersRequest

	// CustomersControllerListCustomersExecute executes the request
	//  @return CustomersControllerListCustomers200Response
	CustomersControllerListCustomersExecute(r ApiCustomersControllerListCustomersRequest) (*CustomersControllerListCustomers200Response, *http.Response, error)

	/*
	CustomersControllerPatchCustomer Patch customer

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@return ApiCustomersControllerPatchCustomerRequest
	*/
	CustomersControllerPatchCustomer(ctx context.Context, id string) ApiCustomersControllerPatchCustomerRequest

	// CustomersControllerPatchCustomerExecute executes the request
	//  @return CustomerDto
	CustomersControllerPatchCustomerExecute(r ApiCustomersControllerPatchCustomerRequest) (*CustomerDto, *http.Response, error)

	/*
	CustomersControllerPutCustomer Update customer

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@return ApiCustomersControllerPutCustomerRequest
	*/
	CustomersControllerPutCustomer(ctx context.Context, id string) ApiCustomersControllerPutCustomerRequest

	// CustomersControllerPutCustomerExecute executes the request
	//  @return CustomerDto
	CustomersControllerPutCustomerExecute(r ApiCustomersControllerPutCustomerRequest) (*CustomerDto, *http.Response, error)
}

// CustomersAPIService CustomersAPI service
type CustomersAPIService service

type ApiCustomersControllerCreateCustomerRequest struct {
	ctx context.Context
	ApiService CustomersAPI
	createCustomerDto *CreateCustomerDto
}

func (r ApiCustomersControllerCreateCustomerRequest) CreateCustomerDto(createCustomerDto CreateCustomerDto) ApiCustomersControllerCreateCustomerRequest {
	r.createCustomerDto = &createCustomerDto
	return r
}

func (r ApiCustomersControllerCreateCustomerRequest) Execute() (*CustomerDto, *http.Response, error) {
	return r.ApiService.CustomersControllerCreateCustomerExecute(r)
}

/*
CustomersControllerCreateCustomer Create customer

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCustomersControllerCreateCustomerRequest
*/
func (a *CustomersAPIService) CustomersControllerCreateCustomer(ctx context.Context) ApiCustomersControllerCreateCustomerRequest {
	return ApiCustomersControllerCreateCustomerRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CustomerDto
func (a *CustomersAPIService) CustomersControllerCreateCustomerExecute(r ApiCustomersControllerCreateCustomerRequest) (*CustomerDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CustomerDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomersAPIService.CustomersControllerCreateCustomer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createCustomerDto == nil {
		return localVarReturnValue, nil, reportError("createCustomerDto is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createCustomerDto
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiCustomersControllerDeleteCustomerRequest struct {
	ctx context.Context
	ApiService CustomersAPI
	id string
}

func (r ApiCustomersControllerDeleteCustomerRequest) Execute() (*http.Response, error) {
	return r.ApiService.CustomersControllerDeleteCustomerExecute(r)
}

/*
CustomersControllerDeleteCustomer Delete customer

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiCustomersControllerDeleteCustomerRequest
*/
func (a *CustomersAPIService) CustomersControllerDeleteCustomer(ctx context.Context, id string) ApiCustomersControllerDeleteCustomerRequest {
	return ApiCustomersControllerDeleteCustomerRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *CustomersAPIService) CustomersControllerDeleteCustomerExecute(r ApiCustomersControllerDeleteCustomerRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomersAPIService.CustomersControllerDeleteCustomer")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customers/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiCustomersControllerGetCustomerRequest struct {
	ctx context.Context
	ApiService CustomersAPI
	id string
}

func (r ApiCustomersControllerGetCustomerRequest) Execute() (*CustomerDto, *http.Response, error) {
	return r.ApiService.CustomersControllerGetCustomerExecute(r)
}

/*
CustomersControllerGetCustomer Get customer

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiCustomersControllerGetCustomerRequest
*/
func (a *CustomersAPIService) CustomersControllerGetCustomer(ctx context.Context, id string) ApiCustomersControllerGetCustomerRequest {
	return ApiCustomersControllerGetCustomerRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return CustomerDto
func (a *CustomersAPIService) CustomersControllerGetCustomerExecute(r ApiCustomersControllerGetCustomerRequest) (*CustomerDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CustomerDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomersAPIService.CustomersControllerGetCustomer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customers/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiCustomersControllerListCustomersRequest struct {
	ctx context.Context
	ApiService CustomersAPI
	limit *float32
	cursor *string
	search *string
	connectorId *string
	isTest *bool
}

func (r ApiCustomersControllerListCustomersRequest) Limit(limit float32) ApiCustomersControllerListCustomersRequest {
	r.limit = &limit
	return r
}

func (r ApiCustomersControllerListCustomersRequest) Cursor(cursor string) ApiCustomersControllerListCustomersRequest {
	r.cursor = &cursor
	return r
}

func (r ApiCustomersControllerListCustomersRequest) Search(search string) ApiCustomersControllerListCustomersRequest {
	r.search = &search
	return r
}

func (r ApiCustomersControllerListCustomersRequest) ConnectorId(connectorId string) ApiCustomersControllerListCustomersRequest {
	r.connectorId = &connectorId
	return r
}

func (r ApiCustomersControllerListCustomersRequest) IsTest(isTest bool) ApiCustomersControllerListCustomersRequest {
	r.isTest = &isTest
	return r
}

func (r ApiCustomersControllerListCustomersRequest) Execute() (*CustomersControllerListCustomers200Response, *http.Response, error) {
	return r.ApiService.CustomersControllerListCustomersExecute(r)
}

/*
CustomersControllerListCustomers List customers

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCustomersControllerListCustomersRequest
*/
func (a *CustomersAPIService) CustomersControllerListCustomers(ctx context.Context) ApiCustomersControllerListCustomersRequest {
	return ApiCustomersControllerListCustomersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CustomersControllerListCustomers200Response
func (a *CustomersAPIService) CustomersControllerListCustomersExecute(r ApiCustomersControllerListCustomersRequest) (*CustomersControllerListCustomers200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CustomersControllerListCustomers200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomersAPIService.CustomersControllerListCustomers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.cursor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cursor", r.cursor, "")
	}
	if r.search != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search", r.search, "")
	}
	if r.connectorId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "connectorId", r.connectorId, "")
	}
	if r.isTest != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "isTest", r.isTest, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiCustomersControllerPatchCustomerRequest struct {
	ctx context.Context
	ApiService CustomersAPI
	id string
	updateCustomerAsAdminDto *UpdateCustomerAsAdminDto
}

func (r ApiCustomersControllerPatchCustomerRequest) UpdateCustomerAsAdminDto(updateCustomerAsAdminDto UpdateCustomerAsAdminDto) ApiCustomersControllerPatchCustomerRequest {
	r.updateCustomerAsAdminDto = &updateCustomerAsAdminDto
	return r
}

func (r ApiCustomersControllerPatchCustomerRequest) Execute() (*CustomerDto, *http.Response, error) {
	return r.ApiService.CustomersControllerPatchCustomerExecute(r)
}

/*
CustomersControllerPatchCustomer Patch customer

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiCustomersControllerPatchCustomerRequest
*/
func (a *CustomersAPIService) CustomersControllerPatchCustomer(ctx context.Context, id string) ApiCustomersControllerPatchCustomerRequest {
	return ApiCustomersControllerPatchCustomerRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return CustomerDto
func (a *CustomersAPIService) CustomersControllerPatchCustomerExecute(r ApiCustomersControllerPatchCustomerRequest) (*CustomerDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CustomerDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomersAPIService.CustomersControllerPatchCustomer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customers/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateCustomerAsAdminDto == nil {
		return localVarReturnValue, nil, reportError("updateCustomerAsAdminDto is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.updateCustomerAsAdminDto
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiCustomersControllerPutCustomerRequest struct {
	ctx context.Context
	ApiService CustomersAPI
	id string
	updateCustomerAsAdminDto *UpdateCustomerAsAdminDto
}

func (r ApiCustomersControllerPutCustomerRequest) UpdateCustomerAsAdminDto(updateCustomerAsAdminDto UpdateCustomerAsAdminDto) ApiCustomersControllerPutCustomerRequest {
	r.updateCustomerAsAdminDto = &updateCustomerAsAdminDto
	return r
}

func (r ApiCustomersControllerPutCustomerRequest) Execute() (*CustomerDto, *http.Response, error) {
	return r.ApiService.CustomersControllerPutCustomerExecute(r)
}

/*
CustomersControllerPutCustomer Update customer

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiCustomersControllerPutCustomerRequest
*/
func (a *CustomersAPIService) CustomersControllerPutCustomer(ctx context.Context, id string) ApiCustomersControllerPutCustomerRequest {
	return ApiCustomersControllerPutCustomerRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return CustomerDto
func (a *CustomersAPIService) CustomersControllerPutCustomerExecute(r ApiCustomersControllerPutCustomerRequest) (*CustomerDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CustomerDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomersAPIService.CustomersControllerPutCustomer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customers/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateCustomerAsAdminDto == nil {
		return localVarReturnValue, nil, reportError("updateCustomerAsAdminDto is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.updateCustomerAsAdminDto
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
