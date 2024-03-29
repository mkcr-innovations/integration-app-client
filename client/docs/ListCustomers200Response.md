# ListCustomers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]CustomerDto**](CustomerDto.md) |  | 
**Cursor** | Pointer to **string** |  | [optional] 

## Methods

### NewListCustomers200Response

`func NewListCustomers200Response(items []CustomerDto, ) *ListCustomers200Response`

NewListCustomers200Response instantiates a new ListCustomers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCustomers200ResponseWithDefaults

`func NewListCustomers200ResponseWithDefaults() *ListCustomers200Response`

NewListCustomers200ResponseWithDefaults instantiates a new ListCustomers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ListCustomers200Response) GetItems() []CustomerDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListCustomers200Response) GetItemsOk() (*[]CustomerDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListCustomers200Response) SetItems(v []CustomerDto)`

SetItems sets Items field to given value.


### GetCursor

`func (o *ListCustomers200Response) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *ListCustomers200Response) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *ListCustomers200Response) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *ListCustomers200Response) HasCursor() bool`

HasCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


