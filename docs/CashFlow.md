# CashFlow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **string** |  | [optional] 
**FilingDate** | Pointer to **string** |  | [optional] 
**Investments** | Pointer to **NullableString** |  | [optional] 
**ChangeToLiabilities** | Pointer to **NullableString** |  | [optional] 
**TotalCashflowsFromInvestingActivities** | Pointer to **NullableString** |  | [optional] 
**NetBorrowings** | Pointer to **NullableString** |  | [optional] 
**TotalCashFromFinancingActivities** | Pointer to **NullableString** |  | [optional] 
**ChangeToOperatingActivities** | Pointer to **NullableString** |  | [optional] 
**NetIncome** | Pointer to **NullableString** |  | [optional] 
**ChangeInCash** | Pointer to **NullableString** |  | [optional] 
**BeginPeriodCashFlow** | Pointer to **NullableString** |  | [optional] 
**EndPeriodCashFlow** | Pointer to **NullableString** |  | [optional] 
**TotalCashFromOperatingActivities** | Pointer to **NullableString** |  | [optional] 
**Depreciation** | Pointer to **NullableString** |  | [optional] 
**OtherCashflowsFromInvestingActivities** | Pointer to **NullableString** |  | [optional] 
**DividendsPaid** | Pointer to **NullableString** |  | [optional] 
**ChangeToInventory** | Pointer to **NullableString** |  | [optional] 
**ChangeToAccountReceivables** | Pointer to **NullableString** |  | [optional] 
**SalePurchaseOfStock** | Pointer to **NullableString** |  | [optional] 
**OtherCashflowsFromFinancingActivities** | Pointer to **NullableString** |  | [optional] 
**ChangeToNetincome** | Pointer to **NullableString** |  | [optional] 
**CapitalExpenditures** | Pointer to **NullableString** |  | [optional] 
**ChangeReceivables** | Pointer to **NullableString** |  | [optional] 
**CashFlowsOtherOperating** | Pointer to **NullableString** |  | [optional] 
**ExchangeRateChanges** | Pointer to **NullableString** |  | [optional] 
**CashAndCashEquivalentsChanges** | Pointer to **NullableString** |  | [optional] 
**ChangeInWorkingCapital** | Pointer to **NullableString** |  | [optional] 
**OtherNonCashItems** | Pointer to **NullableString** |  | [optional] 
**FreeCashFlow** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCashFlow

`func NewCashFlow() *CashFlow`

NewCashFlow instantiates a new CashFlow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCashFlowWithDefaults

`func NewCashFlowWithDefaults() *CashFlow`

NewCashFlowWithDefaults instantiates a new CashFlow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *CashFlow) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CashFlow) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *CashFlow) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *CashFlow) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetFilingDate

`func (o *CashFlow) GetFilingDate() string`

GetFilingDate returns the FilingDate field if non-nil, zero value otherwise.

### GetFilingDateOk

`func (o *CashFlow) GetFilingDateOk() (*string, bool)`

GetFilingDateOk returns a tuple with the FilingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilingDate

`func (o *CashFlow) SetFilingDate(v string)`

SetFilingDate sets FilingDate field to given value.

### HasFilingDate

`func (o *CashFlow) HasFilingDate() bool`

HasFilingDate returns a boolean if a field has been set.

### GetInvestments

`func (o *CashFlow) GetInvestments() string`

GetInvestments returns the Investments field if non-nil, zero value otherwise.

### GetInvestmentsOk

`func (o *CashFlow) GetInvestmentsOk() (*string, bool)`

GetInvestmentsOk returns a tuple with the Investments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestments

`func (o *CashFlow) SetInvestments(v string)`

SetInvestments sets Investments field to given value.

### HasInvestments

`func (o *CashFlow) HasInvestments() bool`

HasInvestments returns a boolean if a field has been set.

### SetInvestmentsNil

`func (o *CashFlow) SetInvestmentsNil(b bool)`

 SetInvestmentsNil sets the value for Investments to be an explicit nil

### UnsetInvestments
`func (o *CashFlow) UnsetInvestments()`

UnsetInvestments ensures that no value is present for Investments, not even an explicit nil
### GetChangeToLiabilities

`func (o *CashFlow) GetChangeToLiabilities() string`

GetChangeToLiabilities returns the ChangeToLiabilities field if non-nil, zero value otherwise.

### GetChangeToLiabilitiesOk

`func (o *CashFlow) GetChangeToLiabilitiesOk() (*string, bool)`

GetChangeToLiabilitiesOk returns a tuple with the ChangeToLiabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeToLiabilities

`func (o *CashFlow) SetChangeToLiabilities(v string)`

SetChangeToLiabilities sets ChangeToLiabilities field to given value.

### HasChangeToLiabilities

`func (o *CashFlow) HasChangeToLiabilities() bool`

HasChangeToLiabilities returns a boolean if a field has been set.

### SetChangeToLiabilitiesNil

`func (o *CashFlow) SetChangeToLiabilitiesNil(b bool)`

 SetChangeToLiabilitiesNil sets the value for ChangeToLiabilities to be an explicit nil

### UnsetChangeToLiabilities
`func (o *CashFlow) UnsetChangeToLiabilities()`

UnsetChangeToLiabilities ensures that no value is present for ChangeToLiabilities, not even an explicit nil
### GetTotalCashflowsFromInvestingActivities

`func (o *CashFlow) GetTotalCashflowsFromInvestingActivities() string`

GetTotalCashflowsFromInvestingActivities returns the TotalCashflowsFromInvestingActivities field if non-nil, zero value otherwise.

### GetTotalCashflowsFromInvestingActivitiesOk

`func (o *CashFlow) GetTotalCashflowsFromInvestingActivitiesOk() (*string, bool)`

GetTotalCashflowsFromInvestingActivitiesOk returns a tuple with the TotalCashflowsFromInvestingActivities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCashflowsFromInvestingActivities

`func (o *CashFlow) SetTotalCashflowsFromInvestingActivities(v string)`

SetTotalCashflowsFromInvestingActivities sets TotalCashflowsFromInvestingActivities field to given value.

### HasTotalCashflowsFromInvestingActivities

`func (o *CashFlow) HasTotalCashflowsFromInvestingActivities() bool`

HasTotalCashflowsFromInvestingActivities returns a boolean if a field has been set.

### SetTotalCashflowsFromInvestingActivitiesNil

`func (o *CashFlow) SetTotalCashflowsFromInvestingActivitiesNil(b bool)`

 SetTotalCashflowsFromInvestingActivitiesNil sets the value for TotalCashflowsFromInvestingActivities to be an explicit nil

### UnsetTotalCashflowsFromInvestingActivities
`func (o *CashFlow) UnsetTotalCashflowsFromInvestingActivities()`

UnsetTotalCashflowsFromInvestingActivities ensures that no value is present for TotalCashflowsFromInvestingActivities, not even an explicit nil
### GetNetBorrowings

`func (o *CashFlow) GetNetBorrowings() string`

GetNetBorrowings returns the NetBorrowings field if non-nil, zero value otherwise.

### GetNetBorrowingsOk

`func (o *CashFlow) GetNetBorrowingsOk() (*string, bool)`

GetNetBorrowingsOk returns a tuple with the NetBorrowings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetBorrowings

`func (o *CashFlow) SetNetBorrowings(v string)`

SetNetBorrowings sets NetBorrowings field to given value.

### HasNetBorrowings

`func (o *CashFlow) HasNetBorrowings() bool`

HasNetBorrowings returns a boolean if a field has been set.

### SetNetBorrowingsNil

`func (o *CashFlow) SetNetBorrowingsNil(b bool)`

 SetNetBorrowingsNil sets the value for NetBorrowings to be an explicit nil

### UnsetNetBorrowings
`func (o *CashFlow) UnsetNetBorrowings()`

UnsetNetBorrowings ensures that no value is present for NetBorrowings, not even an explicit nil
### GetTotalCashFromFinancingActivities

`func (o *CashFlow) GetTotalCashFromFinancingActivities() string`

GetTotalCashFromFinancingActivities returns the TotalCashFromFinancingActivities field if non-nil, zero value otherwise.

### GetTotalCashFromFinancingActivitiesOk

`func (o *CashFlow) GetTotalCashFromFinancingActivitiesOk() (*string, bool)`

GetTotalCashFromFinancingActivitiesOk returns a tuple with the TotalCashFromFinancingActivities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCashFromFinancingActivities

`func (o *CashFlow) SetTotalCashFromFinancingActivities(v string)`

SetTotalCashFromFinancingActivities sets TotalCashFromFinancingActivities field to given value.

### HasTotalCashFromFinancingActivities

`func (o *CashFlow) HasTotalCashFromFinancingActivities() bool`

HasTotalCashFromFinancingActivities returns a boolean if a field has been set.

### SetTotalCashFromFinancingActivitiesNil

`func (o *CashFlow) SetTotalCashFromFinancingActivitiesNil(b bool)`

 SetTotalCashFromFinancingActivitiesNil sets the value for TotalCashFromFinancingActivities to be an explicit nil

### UnsetTotalCashFromFinancingActivities
`func (o *CashFlow) UnsetTotalCashFromFinancingActivities()`

UnsetTotalCashFromFinancingActivities ensures that no value is present for TotalCashFromFinancingActivities, not even an explicit nil
### GetChangeToOperatingActivities

`func (o *CashFlow) GetChangeToOperatingActivities() string`

GetChangeToOperatingActivities returns the ChangeToOperatingActivities field if non-nil, zero value otherwise.

### GetChangeToOperatingActivitiesOk

`func (o *CashFlow) GetChangeToOperatingActivitiesOk() (*string, bool)`

GetChangeToOperatingActivitiesOk returns a tuple with the ChangeToOperatingActivities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeToOperatingActivities

`func (o *CashFlow) SetChangeToOperatingActivities(v string)`

SetChangeToOperatingActivities sets ChangeToOperatingActivities field to given value.

### HasChangeToOperatingActivities

`func (o *CashFlow) HasChangeToOperatingActivities() bool`

HasChangeToOperatingActivities returns a boolean if a field has been set.

### SetChangeToOperatingActivitiesNil

`func (o *CashFlow) SetChangeToOperatingActivitiesNil(b bool)`

 SetChangeToOperatingActivitiesNil sets the value for ChangeToOperatingActivities to be an explicit nil

### UnsetChangeToOperatingActivities
`func (o *CashFlow) UnsetChangeToOperatingActivities()`

UnsetChangeToOperatingActivities ensures that no value is present for ChangeToOperatingActivities, not even an explicit nil
### GetNetIncome

`func (o *CashFlow) GetNetIncome() string`

GetNetIncome returns the NetIncome field if non-nil, zero value otherwise.

### GetNetIncomeOk

`func (o *CashFlow) GetNetIncomeOk() (*string, bool)`

GetNetIncomeOk returns a tuple with the NetIncome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetIncome

`func (o *CashFlow) SetNetIncome(v string)`

SetNetIncome sets NetIncome field to given value.

### HasNetIncome

`func (o *CashFlow) HasNetIncome() bool`

HasNetIncome returns a boolean if a field has been set.

### SetNetIncomeNil

`func (o *CashFlow) SetNetIncomeNil(b bool)`

 SetNetIncomeNil sets the value for NetIncome to be an explicit nil

### UnsetNetIncome
`func (o *CashFlow) UnsetNetIncome()`

UnsetNetIncome ensures that no value is present for NetIncome, not even an explicit nil
### GetChangeInCash

`func (o *CashFlow) GetChangeInCash() string`

GetChangeInCash returns the ChangeInCash field if non-nil, zero value otherwise.

### GetChangeInCashOk

`func (o *CashFlow) GetChangeInCashOk() (*string, bool)`

GetChangeInCashOk returns a tuple with the ChangeInCash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeInCash

`func (o *CashFlow) SetChangeInCash(v string)`

SetChangeInCash sets ChangeInCash field to given value.

### HasChangeInCash

`func (o *CashFlow) HasChangeInCash() bool`

HasChangeInCash returns a boolean if a field has been set.

### SetChangeInCashNil

`func (o *CashFlow) SetChangeInCashNil(b bool)`

 SetChangeInCashNil sets the value for ChangeInCash to be an explicit nil

### UnsetChangeInCash
`func (o *CashFlow) UnsetChangeInCash()`

UnsetChangeInCash ensures that no value is present for ChangeInCash, not even an explicit nil
### GetBeginPeriodCashFlow

`func (o *CashFlow) GetBeginPeriodCashFlow() string`

GetBeginPeriodCashFlow returns the BeginPeriodCashFlow field if non-nil, zero value otherwise.

### GetBeginPeriodCashFlowOk

`func (o *CashFlow) GetBeginPeriodCashFlowOk() (*string, bool)`

GetBeginPeriodCashFlowOk returns a tuple with the BeginPeriodCashFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeginPeriodCashFlow

`func (o *CashFlow) SetBeginPeriodCashFlow(v string)`

SetBeginPeriodCashFlow sets BeginPeriodCashFlow field to given value.

### HasBeginPeriodCashFlow

`func (o *CashFlow) HasBeginPeriodCashFlow() bool`

HasBeginPeriodCashFlow returns a boolean if a field has been set.

### SetBeginPeriodCashFlowNil

`func (o *CashFlow) SetBeginPeriodCashFlowNil(b bool)`

 SetBeginPeriodCashFlowNil sets the value for BeginPeriodCashFlow to be an explicit nil

### UnsetBeginPeriodCashFlow
`func (o *CashFlow) UnsetBeginPeriodCashFlow()`

UnsetBeginPeriodCashFlow ensures that no value is present for BeginPeriodCashFlow, not even an explicit nil
### GetEndPeriodCashFlow

`func (o *CashFlow) GetEndPeriodCashFlow() string`

GetEndPeriodCashFlow returns the EndPeriodCashFlow field if non-nil, zero value otherwise.

### GetEndPeriodCashFlowOk

`func (o *CashFlow) GetEndPeriodCashFlowOk() (*string, bool)`

GetEndPeriodCashFlowOk returns a tuple with the EndPeriodCashFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPeriodCashFlow

`func (o *CashFlow) SetEndPeriodCashFlow(v string)`

SetEndPeriodCashFlow sets EndPeriodCashFlow field to given value.

### HasEndPeriodCashFlow

`func (o *CashFlow) HasEndPeriodCashFlow() bool`

HasEndPeriodCashFlow returns a boolean if a field has been set.

### SetEndPeriodCashFlowNil

`func (o *CashFlow) SetEndPeriodCashFlowNil(b bool)`

 SetEndPeriodCashFlowNil sets the value for EndPeriodCashFlow to be an explicit nil

### UnsetEndPeriodCashFlow
`func (o *CashFlow) UnsetEndPeriodCashFlow()`

UnsetEndPeriodCashFlow ensures that no value is present for EndPeriodCashFlow, not even an explicit nil
### GetTotalCashFromOperatingActivities

`func (o *CashFlow) GetTotalCashFromOperatingActivities() string`

GetTotalCashFromOperatingActivities returns the TotalCashFromOperatingActivities field if non-nil, zero value otherwise.

### GetTotalCashFromOperatingActivitiesOk

`func (o *CashFlow) GetTotalCashFromOperatingActivitiesOk() (*string, bool)`

GetTotalCashFromOperatingActivitiesOk returns a tuple with the TotalCashFromOperatingActivities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCashFromOperatingActivities

`func (o *CashFlow) SetTotalCashFromOperatingActivities(v string)`

SetTotalCashFromOperatingActivities sets TotalCashFromOperatingActivities field to given value.

### HasTotalCashFromOperatingActivities

`func (o *CashFlow) HasTotalCashFromOperatingActivities() bool`

HasTotalCashFromOperatingActivities returns a boolean if a field has been set.

### SetTotalCashFromOperatingActivitiesNil

`func (o *CashFlow) SetTotalCashFromOperatingActivitiesNil(b bool)`

 SetTotalCashFromOperatingActivitiesNil sets the value for TotalCashFromOperatingActivities to be an explicit nil

### UnsetTotalCashFromOperatingActivities
`func (o *CashFlow) UnsetTotalCashFromOperatingActivities()`

UnsetTotalCashFromOperatingActivities ensures that no value is present for TotalCashFromOperatingActivities, not even an explicit nil
### GetDepreciation

`func (o *CashFlow) GetDepreciation() string`

GetDepreciation returns the Depreciation field if non-nil, zero value otherwise.

### GetDepreciationOk

`func (o *CashFlow) GetDepreciationOk() (*string, bool)`

GetDepreciationOk returns a tuple with the Depreciation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepreciation

`func (o *CashFlow) SetDepreciation(v string)`

SetDepreciation sets Depreciation field to given value.

### HasDepreciation

`func (o *CashFlow) HasDepreciation() bool`

HasDepreciation returns a boolean if a field has been set.

### SetDepreciationNil

`func (o *CashFlow) SetDepreciationNil(b bool)`

 SetDepreciationNil sets the value for Depreciation to be an explicit nil

### UnsetDepreciation
`func (o *CashFlow) UnsetDepreciation()`

UnsetDepreciation ensures that no value is present for Depreciation, not even an explicit nil
### GetOtherCashflowsFromInvestingActivities

`func (o *CashFlow) GetOtherCashflowsFromInvestingActivities() string`

GetOtherCashflowsFromInvestingActivities returns the OtherCashflowsFromInvestingActivities field if non-nil, zero value otherwise.

### GetOtherCashflowsFromInvestingActivitiesOk

`func (o *CashFlow) GetOtherCashflowsFromInvestingActivitiesOk() (*string, bool)`

GetOtherCashflowsFromInvestingActivitiesOk returns a tuple with the OtherCashflowsFromInvestingActivities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherCashflowsFromInvestingActivities

`func (o *CashFlow) SetOtherCashflowsFromInvestingActivities(v string)`

SetOtherCashflowsFromInvestingActivities sets OtherCashflowsFromInvestingActivities field to given value.

### HasOtherCashflowsFromInvestingActivities

`func (o *CashFlow) HasOtherCashflowsFromInvestingActivities() bool`

HasOtherCashflowsFromInvestingActivities returns a boolean if a field has been set.

### SetOtherCashflowsFromInvestingActivitiesNil

`func (o *CashFlow) SetOtherCashflowsFromInvestingActivitiesNil(b bool)`

 SetOtherCashflowsFromInvestingActivitiesNil sets the value for OtherCashflowsFromInvestingActivities to be an explicit nil

### UnsetOtherCashflowsFromInvestingActivities
`func (o *CashFlow) UnsetOtherCashflowsFromInvestingActivities()`

UnsetOtherCashflowsFromInvestingActivities ensures that no value is present for OtherCashflowsFromInvestingActivities, not even an explicit nil
### GetDividendsPaid

`func (o *CashFlow) GetDividendsPaid() string`

GetDividendsPaid returns the DividendsPaid field if non-nil, zero value otherwise.

### GetDividendsPaidOk

`func (o *CashFlow) GetDividendsPaidOk() (*string, bool)`

GetDividendsPaidOk returns a tuple with the DividendsPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDividendsPaid

`func (o *CashFlow) SetDividendsPaid(v string)`

SetDividendsPaid sets DividendsPaid field to given value.

### HasDividendsPaid

`func (o *CashFlow) HasDividendsPaid() bool`

HasDividendsPaid returns a boolean if a field has been set.

### SetDividendsPaidNil

`func (o *CashFlow) SetDividendsPaidNil(b bool)`

 SetDividendsPaidNil sets the value for DividendsPaid to be an explicit nil

### UnsetDividendsPaid
`func (o *CashFlow) UnsetDividendsPaid()`

UnsetDividendsPaid ensures that no value is present for DividendsPaid, not even an explicit nil
### GetChangeToInventory

`func (o *CashFlow) GetChangeToInventory() string`

GetChangeToInventory returns the ChangeToInventory field if non-nil, zero value otherwise.

### GetChangeToInventoryOk

`func (o *CashFlow) GetChangeToInventoryOk() (*string, bool)`

GetChangeToInventoryOk returns a tuple with the ChangeToInventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeToInventory

`func (o *CashFlow) SetChangeToInventory(v string)`

SetChangeToInventory sets ChangeToInventory field to given value.

### HasChangeToInventory

`func (o *CashFlow) HasChangeToInventory() bool`

HasChangeToInventory returns a boolean if a field has been set.

### SetChangeToInventoryNil

`func (o *CashFlow) SetChangeToInventoryNil(b bool)`

 SetChangeToInventoryNil sets the value for ChangeToInventory to be an explicit nil

### UnsetChangeToInventory
`func (o *CashFlow) UnsetChangeToInventory()`

UnsetChangeToInventory ensures that no value is present for ChangeToInventory, not even an explicit nil
### GetChangeToAccountReceivables

`func (o *CashFlow) GetChangeToAccountReceivables() string`

GetChangeToAccountReceivables returns the ChangeToAccountReceivables field if non-nil, zero value otherwise.

### GetChangeToAccountReceivablesOk

`func (o *CashFlow) GetChangeToAccountReceivablesOk() (*string, bool)`

GetChangeToAccountReceivablesOk returns a tuple with the ChangeToAccountReceivables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeToAccountReceivables

`func (o *CashFlow) SetChangeToAccountReceivables(v string)`

SetChangeToAccountReceivables sets ChangeToAccountReceivables field to given value.

### HasChangeToAccountReceivables

`func (o *CashFlow) HasChangeToAccountReceivables() bool`

HasChangeToAccountReceivables returns a boolean if a field has been set.

### SetChangeToAccountReceivablesNil

`func (o *CashFlow) SetChangeToAccountReceivablesNil(b bool)`

 SetChangeToAccountReceivablesNil sets the value for ChangeToAccountReceivables to be an explicit nil

### UnsetChangeToAccountReceivables
`func (o *CashFlow) UnsetChangeToAccountReceivables()`

UnsetChangeToAccountReceivables ensures that no value is present for ChangeToAccountReceivables, not even an explicit nil
### GetSalePurchaseOfStock

`func (o *CashFlow) GetSalePurchaseOfStock() string`

GetSalePurchaseOfStock returns the SalePurchaseOfStock field if non-nil, zero value otherwise.

### GetSalePurchaseOfStockOk

`func (o *CashFlow) GetSalePurchaseOfStockOk() (*string, bool)`

GetSalePurchaseOfStockOk returns a tuple with the SalePurchaseOfStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalePurchaseOfStock

`func (o *CashFlow) SetSalePurchaseOfStock(v string)`

SetSalePurchaseOfStock sets SalePurchaseOfStock field to given value.

### HasSalePurchaseOfStock

`func (o *CashFlow) HasSalePurchaseOfStock() bool`

HasSalePurchaseOfStock returns a boolean if a field has been set.

### SetSalePurchaseOfStockNil

`func (o *CashFlow) SetSalePurchaseOfStockNil(b bool)`

 SetSalePurchaseOfStockNil sets the value for SalePurchaseOfStock to be an explicit nil

### UnsetSalePurchaseOfStock
`func (o *CashFlow) UnsetSalePurchaseOfStock()`

UnsetSalePurchaseOfStock ensures that no value is present for SalePurchaseOfStock, not even an explicit nil
### GetOtherCashflowsFromFinancingActivities

`func (o *CashFlow) GetOtherCashflowsFromFinancingActivities() string`

GetOtherCashflowsFromFinancingActivities returns the OtherCashflowsFromFinancingActivities field if non-nil, zero value otherwise.

### GetOtherCashflowsFromFinancingActivitiesOk

`func (o *CashFlow) GetOtherCashflowsFromFinancingActivitiesOk() (*string, bool)`

GetOtherCashflowsFromFinancingActivitiesOk returns a tuple with the OtherCashflowsFromFinancingActivities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherCashflowsFromFinancingActivities

`func (o *CashFlow) SetOtherCashflowsFromFinancingActivities(v string)`

SetOtherCashflowsFromFinancingActivities sets OtherCashflowsFromFinancingActivities field to given value.

### HasOtherCashflowsFromFinancingActivities

`func (o *CashFlow) HasOtherCashflowsFromFinancingActivities() bool`

HasOtherCashflowsFromFinancingActivities returns a boolean if a field has been set.

### SetOtherCashflowsFromFinancingActivitiesNil

`func (o *CashFlow) SetOtherCashflowsFromFinancingActivitiesNil(b bool)`

 SetOtherCashflowsFromFinancingActivitiesNil sets the value for OtherCashflowsFromFinancingActivities to be an explicit nil

### UnsetOtherCashflowsFromFinancingActivities
`func (o *CashFlow) UnsetOtherCashflowsFromFinancingActivities()`

UnsetOtherCashflowsFromFinancingActivities ensures that no value is present for OtherCashflowsFromFinancingActivities, not even an explicit nil
### GetChangeToNetincome

`func (o *CashFlow) GetChangeToNetincome() string`

GetChangeToNetincome returns the ChangeToNetincome field if non-nil, zero value otherwise.

### GetChangeToNetincomeOk

`func (o *CashFlow) GetChangeToNetincomeOk() (*string, bool)`

GetChangeToNetincomeOk returns a tuple with the ChangeToNetincome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeToNetincome

`func (o *CashFlow) SetChangeToNetincome(v string)`

SetChangeToNetincome sets ChangeToNetincome field to given value.

### HasChangeToNetincome

`func (o *CashFlow) HasChangeToNetincome() bool`

HasChangeToNetincome returns a boolean if a field has been set.

### SetChangeToNetincomeNil

`func (o *CashFlow) SetChangeToNetincomeNil(b bool)`

 SetChangeToNetincomeNil sets the value for ChangeToNetincome to be an explicit nil

### UnsetChangeToNetincome
`func (o *CashFlow) UnsetChangeToNetincome()`

UnsetChangeToNetincome ensures that no value is present for ChangeToNetincome, not even an explicit nil
### GetCapitalExpenditures

`func (o *CashFlow) GetCapitalExpenditures() string`

GetCapitalExpenditures returns the CapitalExpenditures field if non-nil, zero value otherwise.

### GetCapitalExpendituresOk

`func (o *CashFlow) GetCapitalExpendituresOk() (*string, bool)`

GetCapitalExpendituresOk returns a tuple with the CapitalExpenditures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapitalExpenditures

`func (o *CashFlow) SetCapitalExpenditures(v string)`

SetCapitalExpenditures sets CapitalExpenditures field to given value.

### HasCapitalExpenditures

`func (o *CashFlow) HasCapitalExpenditures() bool`

HasCapitalExpenditures returns a boolean if a field has been set.

### SetCapitalExpendituresNil

`func (o *CashFlow) SetCapitalExpendituresNil(b bool)`

 SetCapitalExpendituresNil sets the value for CapitalExpenditures to be an explicit nil

### UnsetCapitalExpenditures
`func (o *CashFlow) UnsetCapitalExpenditures()`

UnsetCapitalExpenditures ensures that no value is present for CapitalExpenditures, not even an explicit nil
### GetChangeReceivables

`func (o *CashFlow) GetChangeReceivables() string`

GetChangeReceivables returns the ChangeReceivables field if non-nil, zero value otherwise.

### GetChangeReceivablesOk

`func (o *CashFlow) GetChangeReceivablesOk() (*string, bool)`

GetChangeReceivablesOk returns a tuple with the ChangeReceivables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeReceivables

`func (o *CashFlow) SetChangeReceivables(v string)`

SetChangeReceivables sets ChangeReceivables field to given value.

### HasChangeReceivables

`func (o *CashFlow) HasChangeReceivables() bool`

HasChangeReceivables returns a boolean if a field has been set.

### SetChangeReceivablesNil

`func (o *CashFlow) SetChangeReceivablesNil(b bool)`

 SetChangeReceivablesNil sets the value for ChangeReceivables to be an explicit nil

### UnsetChangeReceivables
`func (o *CashFlow) UnsetChangeReceivables()`

UnsetChangeReceivables ensures that no value is present for ChangeReceivables, not even an explicit nil
### GetCashFlowsOtherOperating

`func (o *CashFlow) GetCashFlowsOtherOperating() string`

GetCashFlowsOtherOperating returns the CashFlowsOtherOperating field if non-nil, zero value otherwise.

### GetCashFlowsOtherOperatingOk

`func (o *CashFlow) GetCashFlowsOtherOperatingOk() (*string, bool)`

GetCashFlowsOtherOperatingOk returns a tuple with the CashFlowsOtherOperating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashFlowsOtherOperating

`func (o *CashFlow) SetCashFlowsOtherOperating(v string)`

SetCashFlowsOtherOperating sets CashFlowsOtherOperating field to given value.

### HasCashFlowsOtherOperating

`func (o *CashFlow) HasCashFlowsOtherOperating() bool`

HasCashFlowsOtherOperating returns a boolean if a field has been set.

### SetCashFlowsOtherOperatingNil

`func (o *CashFlow) SetCashFlowsOtherOperatingNil(b bool)`

 SetCashFlowsOtherOperatingNil sets the value for CashFlowsOtherOperating to be an explicit nil

### UnsetCashFlowsOtherOperating
`func (o *CashFlow) UnsetCashFlowsOtherOperating()`

UnsetCashFlowsOtherOperating ensures that no value is present for CashFlowsOtherOperating, not even an explicit nil
### GetExchangeRateChanges

`func (o *CashFlow) GetExchangeRateChanges() string`

GetExchangeRateChanges returns the ExchangeRateChanges field if non-nil, zero value otherwise.

### GetExchangeRateChangesOk

`func (o *CashFlow) GetExchangeRateChangesOk() (*string, bool)`

GetExchangeRateChangesOk returns a tuple with the ExchangeRateChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRateChanges

`func (o *CashFlow) SetExchangeRateChanges(v string)`

SetExchangeRateChanges sets ExchangeRateChanges field to given value.

### HasExchangeRateChanges

`func (o *CashFlow) HasExchangeRateChanges() bool`

HasExchangeRateChanges returns a boolean if a field has been set.

### SetExchangeRateChangesNil

`func (o *CashFlow) SetExchangeRateChangesNil(b bool)`

 SetExchangeRateChangesNil sets the value for ExchangeRateChanges to be an explicit nil

### UnsetExchangeRateChanges
`func (o *CashFlow) UnsetExchangeRateChanges()`

UnsetExchangeRateChanges ensures that no value is present for ExchangeRateChanges, not even an explicit nil
### GetCashAndCashEquivalentsChanges

`func (o *CashFlow) GetCashAndCashEquivalentsChanges() string`

GetCashAndCashEquivalentsChanges returns the CashAndCashEquivalentsChanges field if non-nil, zero value otherwise.

### GetCashAndCashEquivalentsChangesOk

`func (o *CashFlow) GetCashAndCashEquivalentsChangesOk() (*string, bool)`

GetCashAndCashEquivalentsChangesOk returns a tuple with the CashAndCashEquivalentsChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashAndCashEquivalentsChanges

`func (o *CashFlow) SetCashAndCashEquivalentsChanges(v string)`

SetCashAndCashEquivalentsChanges sets CashAndCashEquivalentsChanges field to given value.

### HasCashAndCashEquivalentsChanges

`func (o *CashFlow) HasCashAndCashEquivalentsChanges() bool`

HasCashAndCashEquivalentsChanges returns a boolean if a field has been set.

### SetCashAndCashEquivalentsChangesNil

`func (o *CashFlow) SetCashAndCashEquivalentsChangesNil(b bool)`

 SetCashAndCashEquivalentsChangesNil sets the value for CashAndCashEquivalentsChanges to be an explicit nil

### UnsetCashAndCashEquivalentsChanges
`func (o *CashFlow) UnsetCashAndCashEquivalentsChanges()`

UnsetCashAndCashEquivalentsChanges ensures that no value is present for CashAndCashEquivalentsChanges, not even an explicit nil
### GetChangeInWorkingCapital

`func (o *CashFlow) GetChangeInWorkingCapital() string`

GetChangeInWorkingCapital returns the ChangeInWorkingCapital field if non-nil, zero value otherwise.

### GetChangeInWorkingCapitalOk

`func (o *CashFlow) GetChangeInWorkingCapitalOk() (*string, bool)`

GetChangeInWorkingCapitalOk returns a tuple with the ChangeInWorkingCapital field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeInWorkingCapital

`func (o *CashFlow) SetChangeInWorkingCapital(v string)`

SetChangeInWorkingCapital sets ChangeInWorkingCapital field to given value.

### HasChangeInWorkingCapital

`func (o *CashFlow) HasChangeInWorkingCapital() bool`

HasChangeInWorkingCapital returns a boolean if a field has been set.

### SetChangeInWorkingCapitalNil

`func (o *CashFlow) SetChangeInWorkingCapitalNil(b bool)`

 SetChangeInWorkingCapitalNil sets the value for ChangeInWorkingCapital to be an explicit nil

### UnsetChangeInWorkingCapital
`func (o *CashFlow) UnsetChangeInWorkingCapital()`

UnsetChangeInWorkingCapital ensures that no value is present for ChangeInWorkingCapital, not even an explicit nil
### GetOtherNonCashItems

`func (o *CashFlow) GetOtherNonCashItems() string`

GetOtherNonCashItems returns the OtherNonCashItems field if non-nil, zero value otherwise.

### GetOtherNonCashItemsOk

`func (o *CashFlow) GetOtherNonCashItemsOk() (*string, bool)`

GetOtherNonCashItemsOk returns a tuple with the OtherNonCashItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherNonCashItems

`func (o *CashFlow) SetOtherNonCashItems(v string)`

SetOtherNonCashItems sets OtherNonCashItems field to given value.

### HasOtherNonCashItems

`func (o *CashFlow) HasOtherNonCashItems() bool`

HasOtherNonCashItems returns a boolean if a field has been set.

### SetOtherNonCashItemsNil

`func (o *CashFlow) SetOtherNonCashItemsNil(b bool)`

 SetOtherNonCashItemsNil sets the value for OtherNonCashItems to be an explicit nil

### UnsetOtherNonCashItems
`func (o *CashFlow) UnsetOtherNonCashItems()`

UnsetOtherNonCashItems ensures that no value is present for OtherNonCashItems, not even an explicit nil
### GetFreeCashFlow

`func (o *CashFlow) GetFreeCashFlow() string`

GetFreeCashFlow returns the FreeCashFlow field if non-nil, zero value otherwise.

### GetFreeCashFlowOk

`func (o *CashFlow) GetFreeCashFlowOk() (*string, bool)`

GetFreeCashFlowOk returns a tuple with the FreeCashFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeCashFlow

`func (o *CashFlow) SetFreeCashFlow(v string)`

SetFreeCashFlow sets FreeCashFlow field to given value.

### HasFreeCashFlow

`func (o *CashFlow) HasFreeCashFlow() bool`

HasFreeCashFlow returns a boolean if a field has been set.

### SetFreeCashFlowNil

`func (o *CashFlow) SetFreeCashFlowNil(b bool)`

 SetFreeCashFlowNil sets the value for FreeCashFlow to be an explicit nil

### UnsetFreeCashFlow
`func (o *CashFlow) UnsetFreeCashFlow()`

UnsetFreeCashFlow ensures that no value is present for FreeCashFlow, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


