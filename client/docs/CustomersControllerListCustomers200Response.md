# CustomersControllerListCustomers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | Pointer to **string** |  | [optional] 
**Items** | [**[]CustomerDto**](CustomerDto.md) |  | 

## Methods

### NewCustomersControllerListCustomers200Response

`func NewCustomersControllerListCustomers200Response(items []CustomerDto, ) *CustomersControllerListCustomers200Response`

NewCustomersControllerListCustomers200Response instantiates a new CustomersControllerListCustomers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomersControllerListCustomers200ResponseWithDefaults

`func NewCustomersControllerListCustomers200ResponseWithDefaults() *CustomersControllerListCustomers200Response`

NewCustomersControllerListCustomers200ResponseWithDefaults instantiates a new CustomersControllerListCustomers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursor

`func (o *CustomersControllerListCustomers200Response) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *CustomersControllerListCustomers200Response) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *CustomersControllerListCustomers200Response) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *CustomersControllerListCustomers200Response) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetItems

`func (o *CustomersControllerListCustomers200Response) GetItems() []CustomerDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CustomersControllerListCustomers200Response) GetItemsOk() (*[]CustomerDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CustomersControllerListCustomers200Response) SetItems(v []CustomerDto)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


