# ExternalEventLogRecordsControllerListExternalEventLogRecords200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | Pointer to **string** |  | [optional] 
**Items** | [**[]ExternalEventLogRecordDto**](ExternalEventLogRecordDto.md) |  | 

## Methods

### NewExternalEventLogRecordsControllerListExternalEventLogRecords200Response

`func NewExternalEventLogRecordsControllerListExternalEventLogRecords200Response(items []ExternalEventLogRecordDto, ) *ExternalEventLogRecordsControllerListExternalEventLogRecords200Response`

NewExternalEventLogRecordsControllerListExternalEventLogRecords200Response instantiates a new ExternalEventLogRecordsControllerListExternalEventLogRecords200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalEventLogRecordsControllerListExternalEventLogRecords200ResponseWithDefaults

`func NewExternalEventLogRecordsControllerListExternalEventLogRecords200ResponseWithDefaults() *ExternalEventLogRecordsControllerListExternalEventLogRecords200Response`

NewExternalEventLogRecordsControllerListExternalEventLogRecords200ResponseWithDefaults instantiates a new ExternalEventLogRecordsControllerListExternalEventLogRecords200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursor

`func (o *ExternalEventLogRecordsControllerListExternalEventLogRecords200Response) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *ExternalEventLogRecordsControllerListExternalEventLogRecords200Response) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *ExternalEventLogRecordsControllerListExternalEventLogRecords200Response) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *ExternalEventLogRecordsControllerListExternalEventLogRecords200Response) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetItems

`func (o *ExternalEventLogRecordsControllerListExternalEventLogRecords200Response) GetItems() []ExternalEventLogRecordDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ExternalEventLogRecordsControllerListExternalEventLogRecords200Response) GetItemsOk() (*[]ExternalEventLogRecordDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ExternalEventLogRecordsControllerListExternalEventLogRecords200Response) SetItems(v []ExternalEventLogRecordDto)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


