# ExchangeDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**OperatingMIC** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Timezone** | Pointer to **string** |  | [optional] 
**IsOpen** | Pointer to **bool** |  | [optional] 
**TradingHours** | Pointer to [**ExchangeDetailsTradingHours**](ExchangeDetailsTradingHours.md) |  | [optional] 
**ExchangeHolidays** | Pointer to [**map[string]ExchangeDetailsExchangeHolidays**](ExchangeDetailsExchangeHolidays.md) |  | [optional] 
**ActiveTickers** | Pointer to **int32** |  | [optional] 
**UpdatedTickers** | Pointer to **int32** |  | [optional] 

## Methods

### NewExchangeDetails

`func NewExchangeDetails() *ExchangeDetails`

NewExchangeDetails instantiates a new ExchangeDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeDetailsWithDefaults

`func NewExchangeDetailsWithDefaults() *ExchangeDetails`

NewExchangeDetailsWithDefaults instantiates a new ExchangeDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ExchangeDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExchangeDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExchangeDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExchangeDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *ExchangeDetails) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ExchangeDetails) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ExchangeDetails) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ExchangeDetails) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetOperatingMIC

`func (o *ExchangeDetails) GetOperatingMIC() string`

GetOperatingMIC returns the OperatingMIC field if non-nil, zero value otherwise.

### GetOperatingMICOk

`func (o *ExchangeDetails) GetOperatingMICOk() (*string, bool)`

GetOperatingMICOk returns a tuple with the OperatingMIC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingMIC

`func (o *ExchangeDetails) SetOperatingMIC(v string)`

SetOperatingMIC sets OperatingMIC field to given value.

### HasOperatingMIC

`func (o *ExchangeDetails) HasOperatingMIC() bool`

HasOperatingMIC returns a boolean if a field has been set.

### GetCountry

`func (o *ExchangeDetails) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ExchangeDetails) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ExchangeDetails) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ExchangeDetails) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCurrency

`func (o *ExchangeDetails) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ExchangeDetails) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ExchangeDetails) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ExchangeDetails) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetTimezone

`func (o *ExchangeDetails) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ExchangeDetails) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ExchangeDetails) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *ExchangeDetails) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetIsOpen

`func (o *ExchangeDetails) GetIsOpen() bool`

GetIsOpen returns the IsOpen field if non-nil, zero value otherwise.

### GetIsOpenOk

`func (o *ExchangeDetails) GetIsOpenOk() (*bool, bool)`

GetIsOpenOk returns a tuple with the IsOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOpen

`func (o *ExchangeDetails) SetIsOpen(v bool)`

SetIsOpen sets IsOpen field to given value.

### HasIsOpen

`func (o *ExchangeDetails) HasIsOpen() bool`

HasIsOpen returns a boolean if a field has been set.

### GetTradingHours

`func (o *ExchangeDetails) GetTradingHours() ExchangeDetailsTradingHours`

GetTradingHours returns the TradingHours field if non-nil, zero value otherwise.

### GetTradingHoursOk

`func (o *ExchangeDetails) GetTradingHoursOk() (*ExchangeDetailsTradingHours, bool)`

GetTradingHoursOk returns a tuple with the TradingHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingHours

`func (o *ExchangeDetails) SetTradingHours(v ExchangeDetailsTradingHours)`

SetTradingHours sets TradingHours field to given value.

### HasTradingHours

`func (o *ExchangeDetails) HasTradingHours() bool`

HasTradingHours returns a boolean if a field has been set.

### GetExchangeHolidays

`func (o *ExchangeDetails) GetExchangeHolidays() map[string]ExchangeDetailsExchangeHolidays`

GetExchangeHolidays returns the ExchangeHolidays field if non-nil, zero value otherwise.

### GetExchangeHolidaysOk

`func (o *ExchangeDetails) GetExchangeHolidaysOk() (*map[string]ExchangeDetailsExchangeHolidays, bool)`

GetExchangeHolidaysOk returns a tuple with the ExchangeHolidays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeHolidays

`func (o *ExchangeDetails) SetExchangeHolidays(v map[string]ExchangeDetailsExchangeHolidays)`

SetExchangeHolidays sets ExchangeHolidays field to given value.

### HasExchangeHolidays

`func (o *ExchangeDetails) HasExchangeHolidays() bool`

HasExchangeHolidays returns a boolean if a field has been set.

### GetActiveTickers

`func (o *ExchangeDetails) GetActiveTickers() int32`

GetActiveTickers returns the ActiveTickers field if non-nil, zero value otherwise.

### GetActiveTickersOk

`func (o *ExchangeDetails) GetActiveTickersOk() (*int32, bool)`

GetActiveTickersOk returns a tuple with the ActiveTickers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveTickers

`func (o *ExchangeDetails) SetActiveTickers(v int32)`

SetActiveTickers sets ActiveTickers field to given value.

### HasActiveTickers

`func (o *ExchangeDetails) HasActiveTickers() bool`

HasActiveTickers returns a boolean if a field has been set.

### GetUpdatedTickers

`func (o *ExchangeDetails) GetUpdatedTickers() int32`

GetUpdatedTickers returns the UpdatedTickers field if non-nil, zero value otherwise.

### GetUpdatedTickersOk

`func (o *ExchangeDetails) GetUpdatedTickersOk() (*int32, bool)`

GetUpdatedTickersOk returns a tuple with the UpdatedTickers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTickers

`func (o *ExchangeDetails) SetUpdatedTickers(v int32)`

SetUpdatedTickers sets UpdatedTickers field to given value.

### HasUpdatedTickers

`func (o *ExchangeDetails) HasUpdatedTickers() bool`

HasUpdatedTickers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


