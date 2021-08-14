# FundamentalsFinancial

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **string** |  | [optional] 
**FilingDate** | Pointer to **string** |  | [optional] 
**CurrencySymbol** | Pointer to **NullableString** |  | [optional] 
**TotalAssets** | Pointer to **NullableString** |  | [optional] 
**IntangibleAssets** | Pointer to **NullableString** |  | [optional] 
**EarningAssets** | Pointer to **NullableString** |  | [optional] 
**OtherCurrentAssets** | Pointer to **NullableString** |  | [optional] 
**TotalLiab** | Pointer to **NullableString** |  | [optional] 
**TotalStockholderEquity** | Pointer to **NullableString** |  | [optional] 
**DeferredLongTermLiab** | Pointer to **NullableString** |  | [optional] 
**OtherCurrentLiab** | Pointer to **NullableString** |  | [optional] 
**CommonStock** | Pointer to **NullableString** |  | [optional] 
**RetainedEarnings** | Pointer to **NullableString** |  | [optional] 
**OtherLiab** | Pointer to **NullableString** |  | [optional] 
**GoodWill** | Pointer to **NullableString** |  | [optional] 
**OtherAssets** | Pointer to **NullableString** |  | [optional] 
**Cash** | Pointer to **NullableString** |  | [optional] 
**TotalCurrentLiabilities** | Pointer to **NullableString** |  | [optional] 
**NetDebt** | Pointer to **NullableString** |  | [optional] 
**ShortTermDebt** | Pointer to **NullableString** |  | [optional] 
**ShortLongTermDebt** | Pointer to **NullableString** |  | [optional] 
**ShortLongTermDebtTotal** | Pointer to **NullableString** |  | [optional] 
**OtherStockholderEquity** | Pointer to **NullableString** |  | [optional] 
**PropertyPlantEquipment** | Pointer to **NullableString** |  | [optional] 
**TotalCurrentAssets** | Pointer to **NullableString** |  | [optional] 
**LongTermInvestments** | Pointer to **NullableString** |  | [optional] 
**NetTangibleAssets** | Pointer to **NullableString** |  | [optional] 
**ShortTermInvestments** | Pointer to **NullableString** |  | [optional] 
**NetReceivables** | Pointer to **NullableString** |  | [optional] 
**LongTermDebt** | Pointer to **NullableString** |  | [optional] 
**Inventory** | Pointer to **NullableString** |  | [optional] 
**AccountsPayable** | Pointer to **NullableString** |  | [optional] 
**TotalPermanentEquity** | Pointer to **NullableString** |  | [optional] 
**NoncontrollingInterestInConsolidatedEntity** | Pointer to **NullableString** |  | [optional] 
**TemporaryEquityRedeemableNoncontrollingInterests** | Pointer to **NullableString** |  | [optional] 
**AccumulatedOtherComprehensiveIncome** | Pointer to **NullableString** |  | [optional] 
**AdditionalPaidInCapital** | Pointer to **NullableString** |  | [optional] 
**CommonStockTotalEquity** | Pointer to **NullableString** |  | [optional] 
**PreferredStockTotalEquity** | Pointer to **NullableString** |  | [optional] 
**RetainedEarningsTotalEquity** | Pointer to **NullableString** |  | [optional] 
**TreasuryStock** | Pointer to **NullableString** |  | [optional] 
**AccumulatedAmortization** | Pointer to **NullableString** |  | [optional] 
**NonCurrrentAssetsOther** | Pointer to **NullableString** |  | [optional] 
**DeferredLongTermAssetCharges** | Pointer to **NullableString** |  | [optional] 
**NonCurrentAssetsTotal** | Pointer to **NullableString** |  | [optional] 
**CapitalLeaseObligations** | Pointer to **NullableString** |  | [optional] 
**LongTermDebtTotal** | Pointer to **NullableString** |  | [optional] 
**NonCurrentLiabilitiesOther** | Pointer to **NullableString** |  | [optional] 
**NonCurrentLiabilitiesTotal** | Pointer to **NullableString** |  | [optional] 
**NegativeGoodwill** | Pointer to **NullableString** |  | [optional] 
**Warrants** | Pointer to **NullableString** |  | [optional] 
**PreferredStockRedeemable** | Pointer to **NullableString** |  | [optional] 
**CapitalSurpluse** | Pointer to **NullableString** |  | [optional] 
**LiabilitiesAndStockholdersEquity** | Pointer to **NullableString** |  | [optional] 
**CashAndShortTermInvestments** | Pointer to **NullableString** |  | [optional] 
**PropertyPlantAndEquipmentGross** | Pointer to **NullableString** |  | [optional] 
**AccumulatedDepreciation** | Pointer to **NullableString** |  | [optional] 
**NetWorkingCapital** | Pointer to **NullableString** |  | [optional] 
**NetInvestedCapital** | Pointer to **NullableString** |  | [optional] 
**CommonStockSharesOutstanding** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewFundamentalsFinancial

`func NewFundamentalsFinancial() *FundamentalsFinancial`

NewFundamentalsFinancial instantiates a new FundamentalsFinancial object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFundamentalsFinancialWithDefaults

`func NewFundamentalsFinancialWithDefaults() *FundamentalsFinancial`

NewFundamentalsFinancialWithDefaults instantiates a new FundamentalsFinancial object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *FundamentalsFinancial) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *FundamentalsFinancial) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *FundamentalsFinancial) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *FundamentalsFinancial) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetFilingDate

`func (o *FundamentalsFinancial) GetFilingDate() string`

GetFilingDate returns the FilingDate field if non-nil, zero value otherwise.

### GetFilingDateOk

`func (o *FundamentalsFinancial) GetFilingDateOk() (*string, bool)`

GetFilingDateOk returns a tuple with the FilingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilingDate

`func (o *FundamentalsFinancial) SetFilingDate(v string)`

SetFilingDate sets FilingDate field to given value.

### HasFilingDate

`func (o *FundamentalsFinancial) HasFilingDate() bool`

HasFilingDate returns a boolean if a field has been set.

### GetCurrencySymbol

`func (o *FundamentalsFinancial) GetCurrencySymbol() string`

GetCurrencySymbol returns the CurrencySymbol field if non-nil, zero value otherwise.

### GetCurrencySymbolOk

`func (o *FundamentalsFinancial) GetCurrencySymbolOk() (*string, bool)`

GetCurrencySymbolOk returns a tuple with the CurrencySymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencySymbol

`func (o *FundamentalsFinancial) SetCurrencySymbol(v string)`

SetCurrencySymbol sets CurrencySymbol field to given value.

### HasCurrencySymbol

`func (o *FundamentalsFinancial) HasCurrencySymbol() bool`

HasCurrencySymbol returns a boolean if a field has been set.

### SetCurrencySymbolNil

`func (o *FundamentalsFinancial) SetCurrencySymbolNil(b bool)`

 SetCurrencySymbolNil sets the value for CurrencySymbol to be an explicit nil

### UnsetCurrencySymbol
`func (o *FundamentalsFinancial) UnsetCurrencySymbol()`

UnsetCurrencySymbol ensures that no value is present for CurrencySymbol, not even an explicit nil
### GetTotalAssets

`func (o *FundamentalsFinancial) GetTotalAssets() string`

GetTotalAssets returns the TotalAssets field if non-nil, zero value otherwise.

### GetTotalAssetsOk

`func (o *FundamentalsFinancial) GetTotalAssetsOk() (*string, bool)`

GetTotalAssetsOk returns a tuple with the TotalAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAssets

`func (o *FundamentalsFinancial) SetTotalAssets(v string)`

SetTotalAssets sets TotalAssets field to given value.

### HasTotalAssets

`func (o *FundamentalsFinancial) HasTotalAssets() bool`

HasTotalAssets returns a boolean if a field has been set.

### SetTotalAssetsNil

`func (o *FundamentalsFinancial) SetTotalAssetsNil(b bool)`

 SetTotalAssetsNil sets the value for TotalAssets to be an explicit nil

### UnsetTotalAssets
`func (o *FundamentalsFinancial) UnsetTotalAssets()`

UnsetTotalAssets ensures that no value is present for TotalAssets, not even an explicit nil
### GetIntangibleAssets

`func (o *FundamentalsFinancial) GetIntangibleAssets() string`

GetIntangibleAssets returns the IntangibleAssets field if non-nil, zero value otherwise.

### GetIntangibleAssetsOk

`func (o *FundamentalsFinancial) GetIntangibleAssetsOk() (*string, bool)`

GetIntangibleAssetsOk returns a tuple with the IntangibleAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntangibleAssets

`func (o *FundamentalsFinancial) SetIntangibleAssets(v string)`

SetIntangibleAssets sets IntangibleAssets field to given value.

### HasIntangibleAssets

`func (o *FundamentalsFinancial) HasIntangibleAssets() bool`

HasIntangibleAssets returns a boolean if a field has been set.

### SetIntangibleAssetsNil

`func (o *FundamentalsFinancial) SetIntangibleAssetsNil(b bool)`

 SetIntangibleAssetsNil sets the value for IntangibleAssets to be an explicit nil

### UnsetIntangibleAssets
`func (o *FundamentalsFinancial) UnsetIntangibleAssets()`

UnsetIntangibleAssets ensures that no value is present for IntangibleAssets, not even an explicit nil
### GetEarningAssets

`func (o *FundamentalsFinancial) GetEarningAssets() string`

GetEarningAssets returns the EarningAssets field if non-nil, zero value otherwise.

### GetEarningAssetsOk

`func (o *FundamentalsFinancial) GetEarningAssetsOk() (*string, bool)`

GetEarningAssetsOk returns a tuple with the EarningAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarningAssets

`func (o *FundamentalsFinancial) SetEarningAssets(v string)`

SetEarningAssets sets EarningAssets field to given value.

### HasEarningAssets

`func (o *FundamentalsFinancial) HasEarningAssets() bool`

HasEarningAssets returns a boolean if a field has been set.

### SetEarningAssetsNil

`func (o *FundamentalsFinancial) SetEarningAssetsNil(b bool)`

 SetEarningAssetsNil sets the value for EarningAssets to be an explicit nil

### UnsetEarningAssets
`func (o *FundamentalsFinancial) UnsetEarningAssets()`

UnsetEarningAssets ensures that no value is present for EarningAssets, not even an explicit nil
### GetOtherCurrentAssets

`func (o *FundamentalsFinancial) GetOtherCurrentAssets() string`

GetOtherCurrentAssets returns the OtherCurrentAssets field if non-nil, zero value otherwise.

### GetOtherCurrentAssetsOk

`func (o *FundamentalsFinancial) GetOtherCurrentAssetsOk() (*string, bool)`

GetOtherCurrentAssetsOk returns a tuple with the OtherCurrentAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherCurrentAssets

`func (o *FundamentalsFinancial) SetOtherCurrentAssets(v string)`

SetOtherCurrentAssets sets OtherCurrentAssets field to given value.

### HasOtherCurrentAssets

`func (o *FundamentalsFinancial) HasOtherCurrentAssets() bool`

HasOtherCurrentAssets returns a boolean if a field has been set.

### SetOtherCurrentAssetsNil

`func (o *FundamentalsFinancial) SetOtherCurrentAssetsNil(b bool)`

 SetOtherCurrentAssetsNil sets the value for OtherCurrentAssets to be an explicit nil

### UnsetOtherCurrentAssets
`func (o *FundamentalsFinancial) UnsetOtherCurrentAssets()`

UnsetOtherCurrentAssets ensures that no value is present for OtherCurrentAssets, not even an explicit nil
### GetTotalLiab

`func (o *FundamentalsFinancial) GetTotalLiab() string`

GetTotalLiab returns the TotalLiab field if non-nil, zero value otherwise.

### GetTotalLiabOk

`func (o *FundamentalsFinancial) GetTotalLiabOk() (*string, bool)`

GetTotalLiabOk returns a tuple with the TotalLiab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLiab

`func (o *FundamentalsFinancial) SetTotalLiab(v string)`

SetTotalLiab sets TotalLiab field to given value.

### HasTotalLiab

`func (o *FundamentalsFinancial) HasTotalLiab() bool`

HasTotalLiab returns a boolean if a field has been set.

### SetTotalLiabNil

`func (o *FundamentalsFinancial) SetTotalLiabNil(b bool)`

 SetTotalLiabNil sets the value for TotalLiab to be an explicit nil

### UnsetTotalLiab
`func (o *FundamentalsFinancial) UnsetTotalLiab()`

UnsetTotalLiab ensures that no value is present for TotalLiab, not even an explicit nil
### GetTotalStockholderEquity

`func (o *FundamentalsFinancial) GetTotalStockholderEquity() string`

GetTotalStockholderEquity returns the TotalStockholderEquity field if non-nil, zero value otherwise.

### GetTotalStockholderEquityOk

`func (o *FundamentalsFinancial) GetTotalStockholderEquityOk() (*string, bool)`

GetTotalStockholderEquityOk returns a tuple with the TotalStockholderEquity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalStockholderEquity

`func (o *FundamentalsFinancial) SetTotalStockholderEquity(v string)`

SetTotalStockholderEquity sets TotalStockholderEquity field to given value.

### HasTotalStockholderEquity

`func (o *FundamentalsFinancial) HasTotalStockholderEquity() bool`

HasTotalStockholderEquity returns a boolean if a field has been set.

### SetTotalStockholderEquityNil

`func (o *FundamentalsFinancial) SetTotalStockholderEquityNil(b bool)`

 SetTotalStockholderEquityNil sets the value for TotalStockholderEquity to be an explicit nil

### UnsetTotalStockholderEquity
`func (o *FundamentalsFinancial) UnsetTotalStockholderEquity()`

UnsetTotalStockholderEquity ensures that no value is present for TotalStockholderEquity, not even an explicit nil
### GetDeferredLongTermLiab

`func (o *FundamentalsFinancial) GetDeferredLongTermLiab() string`

GetDeferredLongTermLiab returns the DeferredLongTermLiab field if non-nil, zero value otherwise.

### GetDeferredLongTermLiabOk

`func (o *FundamentalsFinancial) GetDeferredLongTermLiabOk() (*string, bool)`

GetDeferredLongTermLiabOk returns a tuple with the DeferredLongTermLiab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeferredLongTermLiab

`func (o *FundamentalsFinancial) SetDeferredLongTermLiab(v string)`

SetDeferredLongTermLiab sets DeferredLongTermLiab field to given value.

### HasDeferredLongTermLiab

`func (o *FundamentalsFinancial) HasDeferredLongTermLiab() bool`

HasDeferredLongTermLiab returns a boolean if a field has been set.

### SetDeferredLongTermLiabNil

`func (o *FundamentalsFinancial) SetDeferredLongTermLiabNil(b bool)`

 SetDeferredLongTermLiabNil sets the value for DeferredLongTermLiab to be an explicit nil

### UnsetDeferredLongTermLiab
`func (o *FundamentalsFinancial) UnsetDeferredLongTermLiab()`

UnsetDeferredLongTermLiab ensures that no value is present for DeferredLongTermLiab, not even an explicit nil
### GetOtherCurrentLiab

`func (o *FundamentalsFinancial) GetOtherCurrentLiab() string`

GetOtherCurrentLiab returns the OtherCurrentLiab field if non-nil, zero value otherwise.

### GetOtherCurrentLiabOk

`func (o *FundamentalsFinancial) GetOtherCurrentLiabOk() (*string, bool)`

GetOtherCurrentLiabOk returns a tuple with the OtherCurrentLiab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherCurrentLiab

`func (o *FundamentalsFinancial) SetOtherCurrentLiab(v string)`

SetOtherCurrentLiab sets OtherCurrentLiab field to given value.

### HasOtherCurrentLiab

`func (o *FundamentalsFinancial) HasOtherCurrentLiab() bool`

HasOtherCurrentLiab returns a boolean if a field has been set.

### SetOtherCurrentLiabNil

`func (o *FundamentalsFinancial) SetOtherCurrentLiabNil(b bool)`

 SetOtherCurrentLiabNil sets the value for OtherCurrentLiab to be an explicit nil

### UnsetOtherCurrentLiab
`func (o *FundamentalsFinancial) UnsetOtherCurrentLiab()`

UnsetOtherCurrentLiab ensures that no value is present for OtherCurrentLiab, not even an explicit nil
### GetCommonStock

`func (o *FundamentalsFinancial) GetCommonStock() string`

GetCommonStock returns the CommonStock field if non-nil, zero value otherwise.

### GetCommonStockOk

`func (o *FundamentalsFinancial) GetCommonStockOk() (*string, bool)`

GetCommonStockOk returns a tuple with the CommonStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonStock

`func (o *FundamentalsFinancial) SetCommonStock(v string)`

SetCommonStock sets CommonStock field to given value.

### HasCommonStock

`func (o *FundamentalsFinancial) HasCommonStock() bool`

HasCommonStock returns a boolean if a field has been set.

### SetCommonStockNil

`func (o *FundamentalsFinancial) SetCommonStockNil(b bool)`

 SetCommonStockNil sets the value for CommonStock to be an explicit nil

### UnsetCommonStock
`func (o *FundamentalsFinancial) UnsetCommonStock()`

UnsetCommonStock ensures that no value is present for CommonStock, not even an explicit nil
### GetRetainedEarnings

`func (o *FundamentalsFinancial) GetRetainedEarnings() string`

GetRetainedEarnings returns the RetainedEarnings field if non-nil, zero value otherwise.

### GetRetainedEarningsOk

`func (o *FundamentalsFinancial) GetRetainedEarningsOk() (*string, bool)`

GetRetainedEarningsOk returns a tuple with the RetainedEarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainedEarnings

`func (o *FundamentalsFinancial) SetRetainedEarnings(v string)`

SetRetainedEarnings sets RetainedEarnings field to given value.

### HasRetainedEarnings

`func (o *FundamentalsFinancial) HasRetainedEarnings() bool`

HasRetainedEarnings returns a boolean if a field has been set.

### SetRetainedEarningsNil

`func (o *FundamentalsFinancial) SetRetainedEarningsNil(b bool)`

 SetRetainedEarningsNil sets the value for RetainedEarnings to be an explicit nil

### UnsetRetainedEarnings
`func (o *FundamentalsFinancial) UnsetRetainedEarnings()`

UnsetRetainedEarnings ensures that no value is present for RetainedEarnings, not even an explicit nil
### GetOtherLiab

`func (o *FundamentalsFinancial) GetOtherLiab() string`

GetOtherLiab returns the OtherLiab field if non-nil, zero value otherwise.

### GetOtherLiabOk

`func (o *FundamentalsFinancial) GetOtherLiabOk() (*string, bool)`

GetOtherLiabOk returns a tuple with the OtherLiab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherLiab

`func (o *FundamentalsFinancial) SetOtherLiab(v string)`

SetOtherLiab sets OtherLiab field to given value.

### HasOtherLiab

`func (o *FundamentalsFinancial) HasOtherLiab() bool`

HasOtherLiab returns a boolean if a field has been set.

### SetOtherLiabNil

`func (o *FundamentalsFinancial) SetOtherLiabNil(b bool)`

 SetOtherLiabNil sets the value for OtherLiab to be an explicit nil

### UnsetOtherLiab
`func (o *FundamentalsFinancial) UnsetOtherLiab()`

UnsetOtherLiab ensures that no value is present for OtherLiab, not even an explicit nil
### GetGoodWill

`func (o *FundamentalsFinancial) GetGoodWill() string`

GetGoodWill returns the GoodWill field if non-nil, zero value otherwise.

### GetGoodWillOk

`func (o *FundamentalsFinancial) GetGoodWillOk() (*string, bool)`

GetGoodWillOk returns a tuple with the GoodWill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoodWill

`func (o *FundamentalsFinancial) SetGoodWill(v string)`

SetGoodWill sets GoodWill field to given value.

### HasGoodWill

`func (o *FundamentalsFinancial) HasGoodWill() bool`

HasGoodWill returns a boolean if a field has been set.

### SetGoodWillNil

`func (o *FundamentalsFinancial) SetGoodWillNil(b bool)`

 SetGoodWillNil sets the value for GoodWill to be an explicit nil

### UnsetGoodWill
`func (o *FundamentalsFinancial) UnsetGoodWill()`

UnsetGoodWill ensures that no value is present for GoodWill, not even an explicit nil
### GetOtherAssets

`func (o *FundamentalsFinancial) GetOtherAssets() string`

GetOtherAssets returns the OtherAssets field if non-nil, zero value otherwise.

### GetOtherAssetsOk

`func (o *FundamentalsFinancial) GetOtherAssetsOk() (*string, bool)`

GetOtherAssetsOk returns a tuple with the OtherAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherAssets

`func (o *FundamentalsFinancial) SetOtherAssets(v string)`

SetOtherAssets sets OtherAssets field to given value.

### HasOtherAssets

`func (o *FundamentalsFinancial) HasOtherAssets() bool`

HasOtherAssets returns a boolean if a field has been set.

### SetOtherAssetsNil

`func (o *FundamentalsFinancial) SetOtherAssetsNil(b bool)`

 SetOtherAssetsNil sets the value for OtherAssets to be an explicit nil

### UnsetOtherAssets
`func (o *FundamentalsFinancial) UnsetOtherAssets()`

UnsetOtherAssets ensures that no value is present for OtherAssets, not even an explicit nil
### GetCash

`func (o *FundamentalsFinancial) GetCash() string`

GetCash returns the Cash field if non-nil, zero value otherwise.

### GetCashOk

`func (o *FundamentalsFinancial) GetCashOk() (*string, bool)`

GetCashOk returns a tuple with the Cash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCash

`func (o *FundamentalsFinancial) SetCash(v string)`

SetCash sets Cash field to given value.

### HasCash

`func (o *FundamentalsFinancial) HasCash() bool`

HasCash returns a boolean if a field has been set.

### SetCashNil

`func (o *FundamentalsFinancial) SetCashNil(b bool)`

 SetCashNil sets the value for Cash to be an explicit nil

### UnsetCash
`func (o *FundamentalsFinancial) UnsetCash()`

UnsetCash ensures that no value is present for Cash, not even an explicit nil
### GetTotalCurrentLiabilities

`func (o *FundamentalsFinancial) GetTotalCurrentLiabilities() string`

GetTotalCurrentLiabilities returns the TotalCurrentLiabilities field if non-nil, zero value otherwise.

### GetTotalCurrentLiabilitiesOk

`func (o *FundamentalsFinancial) GetTotalCurrentLiabilitiesOk() (*string, bool)`

GetTotalCurrentLiabilitiesOk returns a tuple with the TotalCurrentLiabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCurrentLiabilities

`func (o *FundamentalsFinancial) SetTotalCurrentLiabilities(v string)`

SetTotalCurrentLiabilities sets TotalCurrentLiabilities field to given value.

### HasTotalCurrentLiabilities

`func (o *FundamentalsFinancial) HasTotalCurrentLiabilities() bool`

HasTotalCurrentLiabilities returns a boolean if a field has been set.

### SetTotalCurrentLiabilitiesNil

`func (o *FundamentalsFinancial) SetTotalCurrentLiabilitiesNil(b bool)`

 SetTotalCurrentLiabilitiesNil sets the value for TotalCurrentLiabilities to be an explicit nil

### UnsetTotalCurrentLiabilities
`func (o *FundamentalsFinancial) UnsetTotalCurrentLiabilities()`

UnsetTotalCurrentLiabilities ensures that no value is present for TotalCurrentLiabilities, not even an explicit nil
### GetNetDebt

`func (o *FundamentalsFinancial) GetNetDebt() string`

GetNetDebt returns the NetDebt field if non-nil, zero value otherwise.

### GetNetDebtOk

`func (o *FundamentalsFinancial) GetNetDebtOk() (*string, bool)`

GetNetDebtOk returns a tuple with the NetDebt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetDebt

`func (o *FundamentalsFinancial) SetNetDebt(v string)`

SetNetDebt sets NetDebt field to given value.

### HasNetDebt

`func (o *FundamentalsFinancial) HasNetDebt() bool`

HasNetDebt returns a boolean if a field has been set.

### SetNetDebtNil

`func (o *FundamentalsFinancial) SetNetDebtNil(b bool)`

 SetNetDebtNil sets the value for NetDebt to be an explicit nil

### UnsetNetDebt
`func (o *FundamentalsFinancial) UnsetNetDebt()`

UnsetNetDebt ensures that no value is present for NetDebt, not even an explicit nil
### GetShortTermDebt

`func (o *FundamentalsFinancial) GetShortTermDebt() string`

GetShortTermDebt returns the ShortTermDebt field if non-nil, zero value otherwise.

### GetShortTermDebtOk

`func (o *FundamentalsFinancial) GetShortTermDebtOk() (*string, bool)`

GetShortTermDebtOk returns a tuple with the ShortTermDebt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortTermDebt

`func (o *FundamentalsFinancial) SetShortTermDebt(v string)`

SetShortTermDebt sets ShortTermDebt field to given value.

### HasShortTermDebt

`func (o *FundamentalsFinancial) HasShortTermDebt() bool`

HasShortTermDebt returns a boolean if a field has been set.

### SetShortTermDebtNil

`func (o *FundamentalsFinancial) SetShortTermDebtNil(b bool)`

 SetShortTermDebtNil sets the value for ShortTermDebt to be an explicit nil

### UnsetShortTermDebt
`func (o *FundamentalsFinancial) UnsetShortTermDebt()`

UnsetShortTermDebt ensures that no value is present for ShortTermDebt, not even an explicit nil
### GetShortLongTermDebt

`func (o *FundamentalsFinancial) GetShortLongTermDebt() string`

GetShortLongTermDebt returns the ShortLongTermDebt field if non-nil, zero value otherwise.

### GetShortLongTermDebtOk

`func (o *FundamentalsFinancial) GetShortLongTermDebtOk() (*string, bool)`

GetShortLongTermDebtOk returns a tuple with the ShortLongTermDebt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortLongTermDebt

`func (o *FundamentalsFinancial) SetShortLongTermDebt(v string)`

SetShortLongTermDebt sets ShortLongTermDebt field to given value.

### HasShortLongTermDebt

`func (o *FundamentalsFinancial) HasShortLongTermDebt() bool`

HasShortLongTermDebt returns a boolean if a field has been set.

### SetShortLongTermDebtNil

`func (o *FundamentalsFinancial) SetShortLongTermDebtNil(b bool)`

 SetShortLongTermDebtNil sets the value for ShortLongTermDebt to be an explicit nil

### UnsetShortLongTermDebt
`func (o *FundamentalsFinancial) UnsetShortLongTermDebt()`

UnsetShortLongTermDebt ensures that no value is present for ShortLongTermDebt, not even an explicit nil
### GetShortLongTermDebtTotal

`func (o *FundamentalsFinancial) GetShortLongTermDebtTotal() string`

GetShortLongTermDebtTotal returns the ShortLongTermDebtTotal field if non-nil, zero value otherwise.

### GetShortLongTermDebtTotalOk

`func (o *FundamentalsFinancial) GetShortLongTermDebtTotalOk() (*string, bool)`

GetShortLongTermDebtTotalOk returns a tuple with the ShortLongTermDebtTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortLongTermDebtTotal

`func (o *FundamentalsFinancial) SetShortLongTermDebtTotal(v string)`

SetShortLongTermDebtTotal sets ShortLongTermDebtTotal field to given value.

### HasShortLongTermDebtTotal

`func (o *FundamentalsFinancial) HasShortLongTermDebtTotal() bool`

HasShortLongTermDebtTotal returns a boolean if a field has been set.

### SetShortLongTermDebtTotalNil

`func (o *FundamentalsFinancial) SetShortLongTermDebtTotalNil(b bool)`

 SetShortLongTermDebtTotalNil sets the value for ShortLongTermDebtTotal to be an explicit nil

### UnsetShortLongTermDebtTotal
`func (o *FundamentalsFinancial) UnsetShortLongTermDebtTotal()`

UnsetShortLongTermDebtTotal ensures that no value is present for ShortLongTermDebtTotal, not even an explicit nil
### GetOtherStockholderEquity

`func (o *FundamentalsFinancial) GetOtherStockholderEquity() string`

GetOtherStockholderEquity returns the OtherStockholderEquity field if non-nil, zero value otherwise.

### GetOtherStockholderEquityOk

`func (o *FundamentalsFinancial) GetOtherStockholderEquityOk() (*string, bool)`

GetOtherStockholderEquityOk returns a tuple with the OtherStockholderEquity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherStockholderEquity

`func (o *FundamentalsFinancial) SetOtherStockholderEquity(v string)`

SetOtherStockholderEquity sets OtherStockholderEquity field to given value.

### HasOtherStockholderEquity

`func (o *FundamentalsFinancial) HasOtherStockholderEquity() bool`

HasOtherStockholderEquity returns a boolean if a field has been set.

### SetOtherStockholderEquityNil

`func (o *FundamentalsFinancial) SetOtherStockholderEquityNil(b bool)`

 SetOtherStockholderEquityNil sets the value for OtherStockholderEquity to be an explicit nil

### UnsetOtherStockholderEquity
`func (o *FundamentalsFinancial) UnsetOtherStockholderEquity()`

UnsetOtherStockholderEquity ensures that no value is present for OtherStockholderEquity, not even an explicit nil
### GetPropertyPlantEquipment

`func (o *FundamentalsFinancial) GetPropertyPlantEquipment() string`

GetPropertyPlantEquipment returns the PropertyPlantEquipment field if non-nil, zero value otherwise.

### GetPropertyPlantEquipmentOk

`func (o *FundamentalsFinancial) GetPropertyPlantEquipmentOk() (*string, bool)`

GetPropertyPlantEquipmentOk returns a tuple with the PropertyPlantEquipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyPlantEquipment

`func (o *FundamentalsFinancial) SetPropertyPlantEquipment(v string)`

SetPropertyPlantEquipment sets PropertyPlantEquipment field to given value.

### HasPropertyPlantEquipment

`func (o *FundamentalsFinancial) HasPropertyPlantEquipment() bool`

HasPropertyPlantEquipment returns a boolean if a field has been set.

### SetPropertyPlantEquipmentNil

`func (o *FundamentalsFinancial) SetPropertyPlantEquipmentNil(b bool)`

 SetPropertyPlantEquipmentNil sets the value for PropertyPlantEquipment to be an explicit nil

### UnsetPropertyPlantEquipment
`func (o *FundamentalsFinancial) UnsetPropertyPlantEquipment()`

UnsetPropertyPlantEquipment ensures that no value is present for PropertyPlantEquipment, not even an explicit nil
### GetTotalCurrentAssets

`func (o *FundamentalsFinancial) GetTotalCurrentAssets() string`

GetTotalCurrentAssets returns the TotalCurrentAssets field if non-nil, zero value otherwise.

### GetTotalCurrentAssetsOk

`func (o *FundamentalsFinancial) GetTotalCurrentAssetsOk() (*string, bool)`

GetTotalCurrentAssetsOk returns a tuple with the TotalCurrentAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCurrentAssets

`func (o *FundamentalsFinancial) SetTotalCurrentAssets(v string)`

SetTotalCurrentAssets sets TotalCurrentAssets field to given value.

### HasTotalCurrentAssets

`func (o *FundamentalsFinancial) HasTotalCurrentAssets() bool`

HasTotalCurrentAssets returns a boolean if a field has been set.

### SetTotalCurrentAssetsNil

`func (o *FundamentalsFinancial) SetTotalCurrentAssetsNil(b bool)`

 SetTotalCurrentAssetsNil sets the value for TotalCurrentAssets to be an explicit nil

### UnsetTotalCurrentAssets
`func (o *FundamentalsFinancial) UnsetTotalCurrentAssets()`

UnsetTotalCurrentAssets ensures that no value is present for TotalCurrentAssets, not even an explicit nil
### GetLongTermInvestments

`func (o *FundamentalsFinancial) GetLongTermInvestments() string`

GetLongTermInvestments returns the LongTermInvestments field if non-nil, zero value otherwise.

### GetLongTermInvestmentsOk

`func (o *FundamentalsFinancial) GetLongTermInvestmentsOk() (*string, bool)`

GetLongTermInvestmentsOk returns a tuple with the LongTermInvestments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongTermInvestments

`func (o *FundamentalsFinancial) SetLongTermInvestments(v string)`

SetLongTermInvestments sets LongTermInvestments field to given value.

### HasLongTermInvestments

`func (o *FundamentalsFinancial) HasLongTermInvestments() bool`

HasLongTermInvestments returns a boolean if a field has been set.

### SetLongTermInvestmentsNil

`func (o *FundamentalsFinancial) SetLongTermInvestmentsNil(b bool)`

 SetLongTermInvestmentsNil sets the value for LongTermInvestments to be an explicit nil

### UnsetLongTermInvestments
`func (o *FundamentalsFinancial) UnsetLongTermInvestments()`

UnsetLongTermInvestments ensures that no value is present for LongTermInvestments, not even an explicit nil
### GetNetTangibleAssets

`func (o *FundamentalsFinancial) GetNetTangibleAssets() string`

GetNetTangibleAssets returns the NetTangibleAssets field if non-nil, zero value otherwise.

### GetNetTangibleAssetsOk

`func (o *FundamentalsFinancial) GetNetTangibleAssetsOk() (*string, bool)`

GetNetTangibleAssetsOk returns a tuple with the NetTangibleAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetTangibleAssets

`func (o *FundamentalsFinancial) SetNetTangibleAssets(v string)`

SetNetTangibleAssets sets NetTangibleAssets field to given value.

### HasNetTangibleAssets

`func (o *FundamentalsFinancial) HasNetTangibleAssets() bool`

HasNetTangibleAssets returns a boolean if a field has been set.

### SetNetTangibleAssetsNil

`func (o *FundamentalsFinancial) SetNetTangibleAssetsNil(b bool)`

 SetNetTangibleAssetsNil sets the value for NetTangibleAssets to be an explicit nil

### UnsetNetTangibleAssets
`func (o *FundamentalsFinancial) UnsetNetTangibleAssets()`

UnsetNetTangibleAssets ensures that no value is present for NetTangibleAssets, not even an explicit nil
### GetShortTermInvestments

`func (o *FundamentalsFinancial) GetShortTermInvestments() string`

GetShortTermInvestments returns the ShortTermInvestments field if non-nil, zero value otherwise.

### GetShortTermInvestmentsOk

`func (o *FundamentalsFinancial) GetShortTermInvestmentsOk() (*string, bool)`

GetShortTermInvestmentsOk returns a tuple with the ShortTermInvestments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortTermInvestments

`func (o *FundamentalsFinancial) SetShortTermInvestments(v string)`

SetShortTermInvestments sets ShortTermInvestments field to given value.

### HasShortTermInvestments

`func (o *FundamentalsFinancial) HasShortTermInvestments() bool`

HasShortTermInvestments returns a boolean if a field has been set.

### SetShortTermInvestmentsNil

`func (o *FundamentalsFinancial) SetShortTermInvestmentsNil(b bool)`

 SetShortTermInvestmentsNil sets the value for ShortTermInvestments to be an explicit nil

### UnsetShortTermInvestments
`func (o *FundamentalsFinancial) UnsetShortTermInvestments()`

UnsetShortTermInvestments ensures that no value is present for ShortTermInvestments, not even an explicit nil
### GetNetReceivables

`func (o *FundamentalsFinancial) GetNetReceivables() string`

GetNetReceivables returns the NetReceivables field if non-nil, zero value otherwise.

### GetNetReceivablesOk

`func (o *FundamentalsFinancial) GetNetReceivablesOk() (*string, bool)`

GetNetReceivablesOk returns a tuple with the NetReceivables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetReceivables

`func (o *FundamentalsFinancial) SetNetReceivables(v string)`

SetNetReceivables sets NetReceivables field to given value.

### HasNetReceivables

`func (o *FundamentalsFinancial) HasNetReceivables() bool`

HasNetReceivables returns a boolean if a field has been set.

### SetNetReceivablesNil

`func (o *FundamentalsFinancial) SetNetReceivablesNil(b bool)`

 SetNetReceivablesNil sets the value for NetReceivables to be an explicit nil

### UnsetNetReceivables
`func (o *FundamentalsFinancial) UnsetNetReceivables()`

UnsetNetReceivables ensures that no value is present for NetReceivables, not even an explicit nil
### GetLongTermDebt

`func (o *FundamentalsFinancial) GetLongTermDebt() string`

GetLongTermDebt returns the LongTermDebt field if non-nil, zero value otherwise.

### GetLongTermDebtOk

`func (o *FundamentalsFinancial) GetLongTermDebtOk() (*string, bool)`

GetLongTermDebtOk returns a tuple with the LongTermDebt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongTermDebt

`func (o *FundamentalsFinancial) SetLongTermDebt(v string)`

SetLongTermDebt sets LongTermDebt field to given value.

### HasLongTermDebt

`func (o *FundamentalsFinancial) HasLongTermDebt() bool`

HasLongTermDebt returns a boolean if a field has been set.

### SetLongTermDebtNil

`func (o *FundamentalsFinancial) SetLongTermDebtNil(b bool)`

 SetLongTermDebtNil sets the value for LongTermDebt to be an explicit nil

### UnsetLongTermDebt
`func (o *FundamentalsFinancial) UnsetLongTermDebt()`

UnsetLongTermDebt ensures that no value is present for LongTermDebt, not even an explicit nil
### GetInventory

`func (o *FundamentalsFinancial) GetInventory() string`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *FundamentalsFinancial) GetInventoryOk() (*string, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *FundamentalsFinancial) SetInventory(v string)`

SetInventory sets Inventory field to given value.

### HasInventory

`func (o *FundamentalsFinancial) HasInventory() bool`

HasInventory returns a boolean if a field has been set.

### SetInventoryNil

`func (o *FundamentalsFinancial) SetInventoryNil(b bool)`

 SetInventoryNil sets the value for Inventory to be an explicit nil

### UnsetInventory
`func (o *FundamentalsFinancial) UnsetInventory()`

UnsetInventory ensures that no value is present for Inventory, not even an explicit nil
### GetAccountsPayable

`func (o *FundamentalsFinancial) GetAccountsPayable() string`

GetAccountsPayable returns the AccountsPayable field if non-nil, zero value otherwise.

### GetAccountsPayableOk

`func (o *FundamentalsFinancial) GetAccountsPayableOk() (*string, bool)`

GetAccountsPayableOk returns a tuple with the AccountsPayable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountsPayable

`func (o *FundamentalsFinancial) SetAccountsPayable(v string)`

SetAccountsPayable sets AccountsPayable field to given value.

### HasAccountsPayable

`func (o *FundamentalsFinancial) HasAccountsPayable() bool`

HasAccountsPayable returns a boolean if a field has been set.

### SetAccountsPayableNil

`func (o *FundamentalsFinancial) SetAccountsPayableNil(b bool)`

 SetAccountsPayableNil sets the value for AccountsPayable to be an explicit nil

### UnsetAccountsPayable
`func (o *FundamentalsFinancial) UnsetAccountsPayable()`

UnsetAccountsPayable ensures that no value is present for AccountsPayable, not even an explicit nil
### GetTotalPermanentEquity

`func (o *FundamentalsFinancial) GetTotalPermanentEquity() string`

GetTotalPermanentEquity returns the TotalPermanentEquity field if non-nil, zero value otherwise.

### GetTotalPermanentEquityOk

`func (o *FundamentalsFinancial) GetTotalPermanentEquityOk() (*string, bool)`

GetTotalPermanentEquityOk returns a tuple with the TotalPermanentEquity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPermanentEquity

`func (o *FundamentalsFinancial) SetTotalPermanentEquity(v string)`

SetTotalPermanentEquity sets TotalPermanentEquity field to given value.

### HasTotalPermanentEquity

`func (o *FundamentalsFinancial) HasTotalPermanentEquity() bool`

HasTotalPermanentEquity returns a boolean if a field has been set.

### SetTotalPermanentEquityNil

`func (o *FundamentalsFinancial) SetTotalPermanentEquityNil(b bool)`

 SetTotalPermanentEquityNil sets the value for TotalPermanentEquity to be an explicit nil

### UnsetTotalPermanentEquity
`func (o *FundamentalsFinancial) UnsetTotalPermanentEquity()`

UnsetTotalPermanentEquity ensures that no value is present for TotalPermanentEquity, not even an explicit nil
### GetNoncontrollingInterestInConsolidatedEntity

`func (o *FundamentalsFinancial) GetNoncontrollingInterestInConsolidatedEntity() string`

GetNoncontrollingInterestInConsolidatedEntity returns the NoncontrollingInterestInConsolidatedEntity field if non-nil, zero value otherwise.

### GetNoncontrollingInterestInConsolidatedEntityOk

`func (o *FundamentalsFinancial) GetNoncontrollingInterestInConsolidatedEntityOk() (*string, bool)`

GetNoncontrollingInterestInConsolidatedEntityOk returns a tuple with the NoncontrollingInterestInConsolidatedEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoncontrollingInterestInConsolidatedEntity

`func (o *FundamentalsFinancial) SetNoncontrollingInterestInConsolidatedEntity(v string)`

SetNoncontrollingInterestInConsolidatedEntity sets NoncontrollingInterestInConsolidatedEntity field to given value.

### HasNoncontrollingInterestInConsolidatedEntity

`func (o *FundamentalsFinancial) HasNoncontrollingInterestInConsolidatedEntity() bool`

HasNoncontrollingInterestInConsolidatedEntity returns a boolean if a field has been set.

### SetNoncontrollingInterestInConsolidatedEntityNil

`func (o *FundamentalsFinancial) SetNoncontrollingInterestInConsolidatedEntityNil(b bool)`

 SetNoncontrollingInterestInConsolidatedEntityNil sets the value for NoncontrollingInterestInConsolidatedEntity to be an explicit nil

### UnsetNoncontrollingInterestInConsolidatedEntity
`func (o *FundamentalsFinancial) UnsetNoncontrollingInterestInConsolidatedEntity()`

UnsetNoncontrollingInterestInConsolidatedEntity ensures that no value is present for NoncontrollingInterestInConsolidatedEntity, not even an explicit nil
### GetTemporaryEquityRedeemableNoncontrollingInterests

`func (o *FundamentalsFinancial) GetTemporaryEquityRedeemableNoncontrollingInterests() string`

GetTemporaryEquityRedeemableNoncontrollingInterests returns the TemporaryEquityRedeemableNoncontrollingInterests field if non-nil, zero value otherwise.

### GetTemporaryEquityRedeemableNoncontrollingInterestsOk

`func (o *FundamentalsFinancial) GetTemporaryEquityRedeemableNoncontrollingInterestsOk() (*string, bool)`

GetTemporaryEquityRedeemableNoncontrollingInterestsOk returns a tuple with the TemporaryEquityRedeemableNoncontrollingInterests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryEquityRedeemableNoncontrollingInterests

`func (o *FundamentalsFinancial) SetTemporaryEquityRedeemableNoncontrollingInterests(v string)`

SetTemporaryEquityRedeemableNoncontrollingInterests sets TemporaryEquityRedeemableNoncontrollingInterests field to given value.

### HasTemporaryEquityRedeemableNoncontrollingInterests

`func (o *FundamentalsFinancial) HasTemporaryEquityRedeemableNoncontrollingInterests() bool`

HasTemporaryEquityRedeemableNoncontrollingInterests returns a boolean if a field has been set.

### SetTemporaryEquityRedeemableNoncontrollingInterestsNil

`func (o *FundamentalsFinancial) SetTemporaryEquityRedeemableNoncontrollingInterestsNil(b bool)`

 SetTemporaryEquityRedeemableNoncontrollingInterestsNil sets the value for TemporaryEquityRedeemableNoncontrollingInterests to be an explicit nil

### UnsetTemporaryEquityRedeemableNoncontrollingInterests
`func (o *FundamentalsFinancial) UnsetTemporaryEquityRedeemableNoncontrollingInterests()`

UnsetTemporaryEquityRedeemableNoncontrollingInterests ensures that no value is present for TemporaryEquityRedeemableNoncontrollingInterests, not even an explicit nil
### GetAccumulatedOtherComprehensiveIncome

`func (o *FundamentalsFinancial) GetAccumulatedOtherComprehensiveIncome() string`

GetAccumulatedOtherComprehensiveIncome returns the AccumulatedOtherComprehensiveIncome field if non-nil, zero value otherwise.

### GetAccumulatedOtherComprehensiveIncomeOk

`func (o *FundamentalsFinancial) GetAccumulatedOtherComprehensiveIncomeOk() (*string, bool)`

GetAccumulatedOtherComprehensiveIncomeOk returns a tuple with the AccumulatedOtherComprehensiveIncome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccumulatedOtherComprehensiveIncome

`func (o *FundamentalsFinancial) SetAccumulatedOtherComprehensiveIncome(v string)`

SetAccumulatedOtherComprehensiveIncome sets AccumulatedOtherComprehensiveIncome field to given value.

### HasAccumulatedOtherComprehensiveIncome

`func (o *FundamentalsFinancial) HasAccumulatedOtherComprehensiveIncome() bool`

HasAccumulatedOtherComprehensiveIncome returns a boolean if a field has been set.

### SetAccumulatedOtherComprehensiveIncomeNil

`func (o *FundamentalsFinancial) SetAccumulatedOtherComprehensiveIncomeNil(b bool)`

 SetAccumulatedOtherComprehensiveIncomeNil sets the value for AccumulatedOtherComprehensiveIncome to be an explicit nil

### UnsetAccumulatedOtherComprehensiveIncome
`func (o *FundamentalsFinancial) UnsetAccumulatedOtherComprehensiveIncome()`

UnsetAccumulatedOtherComprehensiveIncome ensures that no value is present for AccumulatedOtherComprehensiveIncome, not even an explicit nil
### GetAdditionalPaidInCapital

`func (o *FundamentalsFinancial) GetAdditionalPaidInCapital() string`

GetAdditionalPaidInCapital returns the AdditionalPaidInCapital field if non-nil, zero value otherwise.

### GetAdditionalPaidInCapitalOk

`func (o *FundamentalsFinancial) GetAdditionalPaidInCapitalOk() (*string, bool)`

GetAdditionalPaidInCapitalOk returns a tuple with the AdditionalPaidInCapital field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalPaidInCapital

`func (o *FundamentalsFinancial) SetAdditionalPaidInCapital(v string)`

SetAdditionalPaidInCapital sets AdditionalPaidInCapital field to given value.

### HasAdditionalPaidInCapital

`func (o *FundamentalsFinancial) HasAdditionalPaidInCapital() bool`

HasAdditionalPaidInCapital returns a boolean if a field has been set.

### SetAdditionalPaidInCapitalNil

`func (o *FundamentalsFinancial) SetAdditionalPaidInCapitalNil(b bool)`

 SetAdditionalPaidInCapitalNil sets the value for AdditionalPaidInCapital to be an explicit nil

### UnsetAdditionalPaidInCapital
`func (o *FundamentalsFinancial) UnsetAdditionalPaidInCapital()`

UnsetAdditionalPaidInCapital ensures that no value is present for AdditionalPaidInCapital, not even an explicit nil
### GetCommonStockTotalEquity

`func (o *FundamentalsFinancial) GetCommonStockTotalEquity() string`

GetCommonStockTotalEquity returns the CommonStockTotalEquity field if non-nil, zero value otherwise.

### GetCommonStockTotalEquityOk

`func (o *FundamentalsFinancial) GetCommonStockTotalEquityOk() (*string, bool)`

GetCommonStockTotalEquityOk returns a tuple with the CommonStockTotalEquity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonStockTotalEquity

`func (o *FundamentalsFinancial) SetCommonStockTotalEquity(v string)`

SetCommonStockTotalEquity sets CommonStockTotalEquity field to given value.

### HasCommonStockTotalEquity

`func (o *FundamentalsFinancial) HasCommonStockTotalEquity() bool`

HasCommonStockTotalEquity returns a boolean if a field has been set.

### SetCommonStockTotalEquityNil

`func (o *FundamentalsFinancial) SetCommonStockTotalEquityNil(b bool)`

 SetCommonStockTotalEquityNil sets the value for CommonStockTotalEquity to be an explicit nil

### UnsetCommonStockTotalEquity
`func (o *FundamentalsFinancial) UnsetCommonStockTotalEquity()`

UnsetCommonStockTotalEquity ensures that no value is present for CommonStockTotalEquity, not even an explicit nil
### GetPreferredStockTotalEquity

`func (o *FundamentalsFinancial) GetPreferredStockTotalEquity() string`

GetPreferredStockTotalEquity returns the PreferredStockTotalEquity field if non-nil, zero value otherwise.

### GetPreferredStockTotalEquityOk

`func (o *FundamentalsFinancial) GetPreferredStockTotalEquityOk() (*string, bool)`

GetPreferredStockTotalEquityOk returns a tuple with the PreferredStockTotalEquity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredStockTotalEquity

`func (o *FundamentalsFinancial) SetPreferredStockTotalEquity(v string)`

SetPreferredStockTotalEquity sets PreferredStockTotalEquity field to given value.

### HasPreferredStockTotalEquity

`func (o *FundamentalsFinancial) HasPreferredStockTotalEquity() bool`

HasPreferredStockTotalEquity returns a boolean if a field has been set.

### SetPreferredStockTotalEquityNil

`func (o *FundamentalsFinancial) SetPreferredStockTotalEquityNil(b bool)`

 SetPreferredStockTotalEquityNil sets the value for PreferredStockTotalEquity to be an explicit nil

### UnsetPreferredStockTotalEquity
`func (o *FundamentalsFinancial) UnsetPreferredStockTotalEquity()`

UnsetPreferredStockTotalEquity ensures that no value is present for PreferredStockTotalEquity, not even an explicit nil
### GetRetainedEarningsTotalEquity

`func (o *FundamentalsFinancial) GetRetainedEarningsTotalEquity() string`

GetRetainedEarningsTotalEquity returns the RetainedEarningsTotalEquity field if non-nil, zero value otherwise.

### GetRetainedEarningsTotalEquityOk

`func (o *FundamentalsFinancial) GetRetainedEarningsTotalEquityOk() (*string, bool)`

GetRetainedEarningsTotalEquityOk returns a tuple with the RetainedEarningsTotalEquity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainedEarningsTotalEquity

`func (o *FundamentalsFinancial) SetRetainedEarningsTotalEquity(v string)`

SetRetainedEarningsTotalEquity sets RetainedEarningsTotalEquity field to given value.

### HasRetainedEarningsTotalEquity

`func (o *FundamentalsFinancial) HasRetainedEarningsTotalEquity() bool`

HasRetainedEarningsTotalEquity returns a boolean if a field has been set.

### SetRetainedEarningsTotalEquityNil

`func (o *FundamentalsFinancial) SetRetainedEarningsTotalEquityNil(b bool)`

 SetRetainedEarningsTotalEquityNil sets the value for RetainedEarningsTotalEquity to be an explicit nil

### UnsetRetainedEarningsTotalEquity
`func (o *FundamentalsFinancial) UnsetRetainedEarningsTotalEquity()`

UnsetRetainedEarningsTotalEquity ensures that no value is present for RetainedEarningsTotalEquity, not even an explicit nil
### GetTreasuryStock

`func (o *FundamentalsFinancial) GetTreasuryStock() string`

GetTreasuryStock returns the TreasuryStock field if non-nil, zero value otherwise.

### GetTreasuryStockOk

`func (o *FundamentalsFinancial) GetTreasuryStockOk() (*string, bool)`

GetTreasuryStockOk returns a tuple with the TreasuryStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTreasuryStock

`func (o *FundamentalsFinancial) SetTreasuryStock(v string)`

SetTreasuryStock sets TreasuryStock field to given value.

### HasTreasuryStock

`func (o *FundamentalsFinancial) HasTreasuryStock() bool`

HasTreasuryStock returns a boolean if a field has been set.

### SetTreasuryStockNil

`func (o *FundamentalsFinancial) SetTreasuryStockNil(b bool)`

 SetTreasuryStockNil sets the value for TreasuryStock to be an explicit nil

### UnsetTreasuryStock
`func (o *FundamentalsFinancial) UnsetTreasuryStock()`

UnsetTreasuryStock ensures that no value is present for TreasuryStock, not even an explicit nil
### GetAccumulatedAmortization

`func (o *FundamentalsFinancial) GetAccumulatedAmortization() string`

GetAccumulatedAmortization returns the AccumulatedAmortization field if non-nil, zero value otherwise.

### GetAccumulatedAmortizationOk

`func (o *FundamentalsFinancial) GetAccumulatedAmortizationOk() (*string, bool)`

GetAccumulatedAmortizationOk returns a tuple with the AccumulatedAmortization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccumulatedAmortization

`func (o *FundamentalsFinancial) SetAccumulatedAmortization(v string)`

SetAccumulatedAmortization sets AccumulatedAmortization field to given value.

### HasAccumulatedAmortization

`func (o *FundamentalsFinancial) HasAccumulatedAmortization() bool`

HasAccumulatedAmortization returns a boolean if a field has been set.

### SetAccumulatedAmortizationNil

`func (o *FundamentalsFinancial) SetAccumulatedAmortizationNil(b bool)`

 SetAccumulatedAmortizationNil sets the value for AccumulatedAmortization to be an explicit nil

### UnsetAccumulatedAmortization
`func (o *FundamentalsFinancial) UnsetAccumulatedAmortization()`

UnsetAccumulatedAmortization ensures that no value is present for AccumulatedAmortization, not even an explicit nil
### GetNonCurrrentAssetsOther

`func (o *FundamentalsFinancial) GetNonCurrrentAssetsOther() string`

GetNonCurrrentAssetsOther returns the NonCurrrentAssetsOther field if non-nil, zero value otherwise.

### GetNonCurrrentAssetsOtherOk

`func (o *FundamentalsFinancial) GetNonCurrrentAssetsOtherOk() (*string, bool)`

GetNonCurrrentAssetsOtherOk returns a tuple with the NonCurrrentAssetsOther field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonCurrrentAssetsOther

`func (o *FundamentalsFinancial) SetNonCurrrentAssetsOther(v string)`

SetNonCurrrentAssetsOther sets NonCurrrentAssetsOther field to given value.

### HasNonCurrrentAssetsOther

`func (o *FundamentalsFinancial) HasNonCurrrentAssetsOther() bool`

HasNonCurrrentAssetsOther returns a boolean if a field has been set.

### SetNonCurrrentAssetsOtherNil

`func (o *FundamentalsFinancial) SetNonCurrrentAssetsOtherNil(b bool)`

 SetNonCurrrentAssetsOtherNil sets the value for NonCurrrentAssetsOther to be an explicit nil

### UnsetNonCurrrentAssetsOther
`func (o *FundamentalsFinancial) UnsetNonCurrrentAssetsOther()`

UnsetNonCurrrentAssetsOther ensures that no value is present for NonCurrrentAssetsOther, not even an explicit nil
### GetDeferredLongTermAssetCharges

`func (o *FundamentalsFinancial) GetDeferredLongTermAssetCharges() string`

GetDeferredLongTermAssetCharges returns the DeferredLongTermAssetCharges field if non-nil, zero value otherwise.

### GetDeferredLongTermAssetChargesOk

`func (o *FundamentalsFinancial) GetDeferredLongTermAssetChargesOk() (*string, bool)`

GetDeferredLongTermAssetChargesOk returns a tuple with the DeferredLongTermAssetCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeferredLongTermAssetCharges

`func (o *FundamentalsFinancial) SetDeferredLongTermAssetCharges(v string)`

SetDeferredLongTermAssetCharges sets DeferredLongTermAssetCharges field to given value.

### HasDeferredLongTermAssetCharges

`func (o *FundamentalsFinancial) HasDeferredLongTermAssetCharges() bool`

HasDeferredLongTermAssetCharges returns a boolean if a field has been set.

### SetDeferredLongTermAssetChargesNil

`func (o *FundamentalsFinancial) SetDeferredLongTermAssetChargesNil(b bool)`

 SetDeferredLongTermAssetChargesNil sets the value for DeferredLongTermAssetCharges to be an explicit nil

### UnsetDeferredLongTermAssetCharges
`func (o *FundamentalsFinancial) UnsetDeferredLongTermAssetCharges()`

UnsetDeferredLongTermAssetCharges ensures that no value is present for DeferredLongTermAssetCharges, not even an explicit nil
### GetNonCurrentAssetsTotal

`func (o *FundamentalsFinancial) GetNonCurrentAssetsTotal() string`

GetNonCurrentAssetsTotal returns the NonCurrentAssetsTotal field if non-nil, zero value otherwise.

### GetNonCurrentAssetsTotalOk

`func (o *FundamentalsFinancial) GetNonCurrentAssetsTotalOk() (*string, bool)`

GetNonCurrentAssetsTotalOk returns a tuple with the NonCurrentAssetsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonCurrentAssetsTotal

`func (o *FundamentalsFinancial) SetNonCurrentAssetsTotal(v string)`

SetNonCurrentAssetsTotal sets NonCurrentAssetsTotal field to given value.

### HasNonCurrentAssetsTotal

`func (o *FundamentalsFinancial) HasNonCurrentAssetsTotal() bool`

HasNonCurrentAssetsTotal returns a boolean if a field has been set.

### SetNonCurrentAssetsTotalNil

`func (o *FundamentalsFinancial) SetNonCurrentAssetsTotalNil(b bool)`

 SetNonCurrentAssetsTotalNil sets the value for NonCurrentAssetsTotal to be an explicit nil

### UnsetNonCurrentAssetsTotal
`func (o *FundamentalsFinancial) UnsetNonCurrentAssetsTotal()`

UnsetNonCurrentAssetsTotal ensures that no value is present for NonCurrentAssetsTotal, not even an explicit nil
### GetCapitalLeaseObligations

`func (o *FundamentalsFinancial) GetCapitalLeaseObligations() string`

GetCapitalLeaseObligations returns the CapitalLeaseObligations field if non-nil, zero value otherwise.

### GetCapitalLeaseObligationsOk

`func (o *FundamentalsFinancial) GetCapitalLeaseObligationsOk() (*string, bool)`

GetCapitalLeaseObligationsOk returns a tuple with the CapitalLeaseObligations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapitalLeaseObligations

`func (o *FundamentalsFinancial) SetCapitalLeaseObligations(v string)`

SetCapitalLeaseObligations sets CapitalLeaseObligations field to given value.

### HasCapitalLeaseObligations

`func (o *FundamentalsFinancial) HasCapitalLeaseObligations() bool`

HasCapitalLeaseObligations returns a boolean if a field has been set.

### SetCapitalLeaseObligationsNil

`func (o *FundamentalsFinancial) SetCapitalLeaseObligationsNil(b bool)`

 SetCapitalLeaseObligationsNil sets the value for CapitalLeaseObligations to be an explicit nil

### UnsetCapitalLeaseObligations
`func (o *FundamentalsFinancial) UnsetCapitalLeaseObligations()`

UnsetCapitalLeaseObligations ensures that no value is present for CapitalLeaseObligations, not even an explicit nil
### GetLongTermDebtTotal

`func (o *FundamentalsFinancial) GetLongTermDebtTotal() string`

GetLongTermDebtTotal returns the LongTermDebtTotal field if non-nil, zero value otherwise.

### GetLongTermDebtTotalOk

`func (o *FundamentalsFinancial) GetLongTermDebtTotalOk() (*string, bool)`

GetLongTermDebtTotalOk returns a tuple with the LongTermDebtTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongTermDebtTotal

`func (o *FundamentalsFinancial) SetLongTermDebtTotal(v string)`

SetLongTermDebtTotal sets LongTermDebtTotal field to given value.

### HasLongTermDebtTotal

`func (o *FundamentalsFinancial) HasLongTermDebtTotal() bool`

HasLongTermDebtTotal returns a boolean if a field has been set.

### SetLongTermDebtTotalNil

`func (o *FundamentalsFinancial) SetLongTermDebtTotalNil(b bool)`

 SetLongTermDebtTotalNil sets the value for LongTermDebtTotal to be an explicit nil

### UnsetLongTermDebtTotal
`func (o *FundamentalsFinancial) UnsetLongTermDebtTotal()`

UnsetLongTermDebtTotal ensures that no value is present for LongTermDebtTotal, not even an explicit nil
### GetNonCurrentLiabilitiesOther

`func (o *FundamentalsFinancial) GetNonCurrentLiabilitiesOther() string`

GetNonCurrentLiabilitiesOther returns the NonCurrentLiabilitiesOther field if non-nil, zero value otherwise.

### GetNonCurrentLiabilitiesOtherOk

`func (o *FundamentalsFinancial) GetNonCurrentLiabilitiesOtherOk() (*string, bool)`

GetNonCurrentLiabilitiesOtherOk returns a tuple with the NonCurrentLiabilitiesOther field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonCurrentLiabilitiesOther

`func (o *FundamentalsFinancial) SetNonCurrentLiabilitiesOther(v string)`

SetNonCurrentLiabilitiesOther sets NonCurrentLiabilitiesOther field to given value.

### HasNonCurrentLiabilitiesOther

`func (o *FundamentalsFinancial) HasNonCurrentLiabilitiesOther() bool`

HasNonCurrentLiabilitiesOther returns a boolean if a field has been set.

### SetNonCurrentLiabilitiesOtherNil

`func (o *FundamentalsFinancial) SetNonCurrentLiabilitiesOtherNil(b bool)`

 SetNonCurrentLiabilitiesOtherNil sets the value for NonCurrentLiabilitiesOther to be an explicit nil

### UnsetNonCurrentLiabilitiesOther
`func (o *FundamentalsFinancial) UnsetNonCurrentLiabilitiesOther()`

UnsetNonCurrentLiabilitiesOther ensures that no value is present for NonCurrentLiabilitiesOther, not even an explicit nil
### GetNonCurrentLiabilitiesTotal

`func (o *FundamentalsFinancial) GetNonCurrentLiabilitiesTotal() string`

GetNonCurrentLiabilitiesTotal returns the NonCurrentLiabilitiesTotal field if non-nil, zero value otherwise.

### GetNonCurrentLiabilitiesTotalOk

`func (o *FundamentalsFinancial) GetNonCurrentLiabilitiesTotalOk() (*string, bool)`

GetNonCurrentLiabilitiesTotalOk returns a tuple with the NonCurrentLiabilitiesTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonCurrentLiabilitiesTotal

`func (o *FundamentalsFinancial) SetNonCurrentLiabilitiesTotal(v string)`

SetNonCurrentLiabilitiesTotal sets NonCurrentLiabilitiesTotal field to given value.

### HasNonCurrentLiabilitiesTotal

`func (o *FundamentalsFinancial) HasNonCurrentLiabilitiesTotal() bool`

HasNonCurrentLiabilitiesTotal returns a boolean if a field has been set.

### SetNonCurrentLiabilitiesTotalNil

`func (o *FundamentalsFinancial) SetNonCurrentLiabilitiesTotalNil(b bool)`

 SetNonCurrentLiabilitiesTotalNil sets the value for NonCurrentLiabilitiesTotal to be an explicit nil

### UnsetNonCurrentLiabilitiesTotal
`func (o *FundamentalsFinancial) UnsetNonCurrentLiabilitiesTotal()`

UnsetNonCurrentLiabilitiesTotal ensures that no value is present for NonCurrentLiabilitiesTotal, not even an explicit nil
### GetNegativeGoodwill

`func (o *FundamentalsFinancial) GetNegativeGoodwill() string`

GetNegativeGoodwill returns the NegativeGoodwill field if non-nil, zero value otherwise.

### GetNegativeGoodwillOk

`func (o *FundamentalsFinancial) GetNegativeGoodwillOk() (*string, bool)`

GetNegativeGoodwillOk returns a tuple with the NegativeGoodwill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegativeGoodwill

`func (o *FundamentalsFinancial) SetNegativeGoodwill(v string)`

SetNegativeGoodwill sets NegativeGoodwill field to given value.

### HasNegativeGoodwill

`func (o *FundamentalsFinancial) HasNegativeGoodwill() bool`

HasNegativeGoodwill returns a boolean if a field has been set.

### SetNegativeGoodwillNil

`func (o *FundamentalsFinancial) SetNegativeGoodwillNil(b bool)`

 SetNegativeGoodwillNil sets the value for NegativeGoodwill to be an explicit nil

### UnsetNegativeGoodwill
`func (o *FundamentalsFinancial) UnsetNegativeGoodwill()`

UnsetNegativeGoodwill ensures that no value is present for NegativeGoodwill, not even an explicit nil
### GetWarrants

`func (o *FundamentalsFinancial) GetWarrants() string`

GetWarrants returns the Warrants field if non-nil, zero value otherwise.

### GetWarrantsOk

`func (o *FundamentalsFinancial) GetWarrantsOk() (*string, bool)`

GetWarrantsOk returns a tuple with the Warrants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrants

`func (o *FundamentalsFinancial) SetWarrants(v string)`

SetWarrants sets Warrants field to given value.

### HasWarrants

`func (o *FundamentalsFinancial) HasWarrants() bool`

HasWarrants returns a boolean if a field has been set.

### SetWarrantsNil

`func (o *FundamentalsFinancial) SetWarrantsNil(b bool)`

 SetWarrantsNil sets the value for Warrants to be an explicit nil

### UnsetWarrants
`func (o *FundamentalsFinancial) UnsetWarrants()`

UnsetWarrants ensures that no value is present for Warrants, not even an explicit nil
### GetPreferredStockRedeemable

`func (o *FundamentalsFinancial) GetPreferredStockRedeemable() string`

GetPreferredStockRedeemable returns the PreferredStockRedeemable field if non-nil, zero value otherwise.

### GetPreferredStockRedeemableOk

`func (o *FundamentalsFinancial) GetPreferredStockRedeemableOk() (*string, bool)`

GetPreferredStockRedeemableOk returns a tuple with the PreferredStockRedeemable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredStockRedeemable

`func (o *FundamentalsFinancial) SetPreferredStockRedeemable(v string)`

SetPreferredStockRedeemable sets PreferredStockRedeemable field to given value.

### HasPreferredStockRedeemable

`func (o *FundamentalsFinancial) HasPreferredStockRedeemable() bool`

HasPreferredStockRedeemable returns a boolean if a field has been set.

### SetPreferredStockRedeemableNil

`func (o *FundamentalsFinancial) SetPreferredStockRedeemableNil(b bool)`

 SetPreferredStockRedeemableNil sets the value for PreferredStockRedeemable to be an explicit nil

### UnsetPreferredStockRedeemable
`func (o *FundamentalsFinancial) UnsetPreferredStockRedeemable()`

UnsetPreferredStockRedeemable ensures that no value is present for PreferredStockRedeemable, not even an explicit nil
### GetCapitalSurpluse

`func (o *FundamentalsFinancial) GetCapitalSurpluse() string`

GetCapitalSurpluse returns the CapitalSurpluse field if non-nil, zero value otherwise.

### GetCapitalSurpluseOk

`func (o *FundamentalsFinancial) GetCapitalSurpluseOk() (*string, bool)`

GetCapitalSurpluseOk returns a tuple with the CapitalSurpluse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapitalSurpluse

`func (o *FundamentalsFinancial) SetCapitalSurpluse(v string)`

SetCapitalSurpluse sets CapitalSurpluse field to given value.

### HasCapitalSurpluse

`func (o *FundamentalsFinancial) HasCapitalSurpluse() bool`

HasCapitalSurpluse returns a boolean if a field has been set.

### SetCapitalSurpluseNil

`func (o *FundamentalsFinancial) SetCapitalSurpluseNil(b bool)`

 SetCapitalSurpluseNil sets the value for CapitalSurpluse to be an explicit nil

### UnsetCapitalSurpluse
`func (o *FundamentalsFinancial) UnsetCapitalSurpluse()`

UnsetCapitalSurpluse ensures that no value is present for CapitalSurpluse, not even an explicit nil
### GetLiabilitiesAndStockholdersEquity

`func (o *FundamentalsFinancial) GetLiabilitiesAndStockholdersEquity() string`

GetLiabilitiesAndStockholdersEquity returns the LiabilitiesAndStockholdersEquity field if non-nil, zero value otherwise.

### GetLiabilitiesAndStockholdersEquityOk

`func (o *FundamentalsFinancial) GetLiabilitiesAndStockholdersEquityOk() (*string, bool)`

GetLiabilitiesAndStockholdersEquityOk returns a tuple with the LiabilitiesAndStockholdersEquity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiabilitiesAndStockholdersEquity

`func (o *FundamentalsFinancial) SetLiabilitiesAndStockholdersEquity(v string)`

SetLiabilitiesAndStockholdersEquity sets LiabilitiesAndStockholdersEquity field to given value.

### HasLiabilitiesAndStockholdersEquity

`func (o *FundamentalsFinancial) HasLiabilitiesAndStockholdersEquity() bool`

HasLiabilitiesAndStockholdersEquity returns a boolean if a field has been set.

### SetLiabilitiesAndStockholdersEquityNil

`func (o *FundamentalsFinancial) SetLiabilitiesAndStockholdersEquityNil(b bool)`

 SetLiabilitiesAndStockholdersEquityNil sets the value for LiabilitiesAndStockholdersEquity to be an explicit nil

### UnsetLiabilitiesAndStockholdersEquity
`func (o *FundamentalsFinancial) UnsetLiabilitiesAndStockholdersEquity()`

UnsetLiabilitiesAndStockholdersEquity ensures that no value is present for LiabilitiesAndStockholdersEquity, not even an explicit nil
### GetCashAndShortTermInvestments

`func (o *FundamentalsFinancial) GetCashAndShortTermInvestments() string`

GetCashAndShortTermInvestments returns the CashAndShortTermInvestments field if non-nil, zero value otherwise.

### GetCashAndShortTermInvestmentsOk

`func (o *FundamentalsFinancial) GetCashAndShortTermInvestmentsOk() (*string, bool)`

GetCashAndShortTermInvestmentsOk returns a tuple with the CashAndShortTermInvestments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashAndShortTermInvestments

`func (o *FundamentalsFinancial) SetCashAndShortTermInvestments(v string)`

SetCashAndShortTermInvestments sets CashAndShortTermInvestments field to given value.

### HasCashAndShortTermInvestments

`func (o *FundamentalsFinancial) HasCashAndShortTermInvestments() bool`

HasCashAndShortTermInvestments returns a boolean if a field has been set.

### SetCashAndShortTermInvestmentsNil

`func (o *FundamentalsFinancial) SetCashAndShortTermInvestmentsNil(b bool)`

 SetCashAndShortTermInvestmentsNil sets the value for CashAndShortTermInvestments to be an explicit nil

### UnsetCashAndShortTermInvestments
`func (o *FundamentalsFinancial) UnsetCashAndShortTermInvestments()`

UnsetCashAndShortTermInvestments ensures that no value is present for CashAndShortTermInvestments, not even an explicit nil
### GetPropertyPlantAndEquipmentGross

`func (o *FundamentalsFinancial) GetPropertyPlantAndEquipmentGross() string`

GetPropertyPlantAndEquipmentGross returns the PropertyPlantAndEquipmentGross field if non-nil, zero value otherwise.

### GetPropertyPlantAndEquipmentGrossOk

`func (o *FundamentalsFinancial) GetPropertyPlantAndEquipmentGrossOk() (*string, bool)`

GetPropertyPlantAndEquipmentGrossOk returns a tuple with the PropertyPlantAndEquipmentGross field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyPlantAndEquipmentGross

`func (o *FundamentalsFinancial) SetPropertyPlantAndEquipmentGross(v string)`

SetPropertyPlantAndEquipmentGross sets PropertyPlantAndEquipmentGross field to given value.

### HasPropertyPlantAndEquipmentGross

`func (o *FundamentalsFinancial) HasPropertyPlantAndEquipmentGross() bool`

HasPropertyPlantAndEquipmentGross returns a boolean if a field has been set.

### SetPropertyPlantAndEquipmentGrossNil

`func (o *FundamentalsFinancial) SetPropertyPlantAndEquipmentGrossNil(b bool)`

 SetPropertyPlantAndEquipmentGrossNil sets the value for PropertyPlantAndEquipmentGross to be an explicit nil

### UnsetPropertyPlantAndEquipmentGross
`func (o *FundamentalsFinancial) UnsetPropertyPlantAndEquipmentGross()`

UnsetPropertyPlantAndEquipmentGross ensures that no value is present for PropertyPlantAndEquipmentGross, not even an explicit nil
### GetAccumulatedDepreciation

`func (o *FundamentalsFinancial) GetAccumulatedDepreciation() string`

GetAccumulatedDepreciation returns the AccumulatedDepreciation field if non-nil, zero value otherwise.

### GetAccumulatedDepreciationOk

`func (o *FundamentalsFinancial) GetAccumulatedDepreciationOk() (*string, bool)`

GetAccumulatedDepreciationOk returns a tuple with the AccumulatedDepreciation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccumulatedDepreciation

`func (o *FundamentalsFinancial) SetAccumulatedDepreciation(v string)`

SetAccumulatedDepreciation sets AccumulatedDepreciation field to given value.

### HasAccumulatedDepreciation

`func (o *FundamentalsFinancial) HasAccumulatedDepreciation() bool`

HasAccumulatedDepreciation returns a boolean if a field has been set.

### SetAccumulatedDepreciationNil

`func (o *FundamentalsFinancial) SetAccumulatedDepreciationNil(b bool)`

 SetAccumulatedDepreciationNil sets the value for AccumulatedDepreciation to be an explicit nil

### UnsetAccumulatedDepreciation
`func (o *FundamentalsFinancial) UnsetAccumulatedDepreciation()`

UnsetAccumulatedDepreciation ensures that no value is present for AccumulatedDepreciation, not even an explicit nil
### GetNetWorkingCapital

`func (o *FundamentalsFinancial) GetNetWorkingCapital() string`

GetNetWorkingCapital returns the NetWorkingCapital field if non-nil, zero value otherwise.

### GetNetWorkingCapitalOk

`func (o *FundamentalsFinancial) GetNetWorkingCapitalOk() (*string, bool)`

GetNetWorkingCapitalOk returns a tuple with the NetWorkingCapital field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetWorkingCapital

`func (o *FundamentalsFinancial) SetNetWorkingCapital(v string)`

SetNetWorkingCapital sets NetWorkingCapital field to given value.

### HasNetWorkingCapital

`func (o *FundamentalsFinancial) HasNetWorkingCapital() bool`

HasNetWorkingCapital returns a boolean if a field has been set.

### SetNetWorkingCapitalNil

`func (o *FundamentalsFinancial) SetNetWorkingCapitalNil(b bool)`

 SetNetWorkingCapitalNil sets the value for NetWorkingCapital to be an explicit nil

### UnsetNetWorkingCapital
`func (o *FundamentalsFinancial) UnsetNetWorkingCapital()`

UnsetNetWorkingCapital ensures that no value is present for NetWorkingCapital, not even an explicit nil
### GetNetInvestedCapital

`func (o *FundamentalsFinancial) GetNetInvestedCapital() string`

GetNetInvestedCapital returns the NetInvestedCapital field if non-nil, zero value otherwise.

### GetNetInvestedCapitalOk

`func (o *FundamentalsFinancial) GetNetInvestedCapitalOk() (*string, bool)`

GetNetInvestedCapitalOk returns a tuple with the NetInvestedCapital field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetInvestedCapital

`func (o *FundamentalsFinancial) SetNetInvestedCapital(v string)`

SetNetInvestedCapital sets NetInvestedCapital field to given value.

### HasNetInvestedCapital

`func (o *FundamentalsFinancial) HasNetInvestedCapital() bool`

HasNetInvestedCapital returns a boolean if a field has been set.

### SetNetInvestedCapitalNil

`func (o *FundamentalsFinancial) SetNetInvestedCapitalNil(b bool)`

 SetNetInvestedCapitalNil sets the value for NetInvestedCapital to be an explicit nil

### UnsetNetInvestedCapital
`func (o *FundamentalsFinancial) UnsetNetInvestedCapital()`

UnsetNetInvestedCapital ensures that no value is present for NetInvestedCapital, not even an explicit nil
### GetCommonStockSharesOutstanding

`func (o *FundamentalsFinancial) GetCommonStockSharesOutstanding() string`

GetCommonStockSharesOutstanding returns the CommonStockSharesOutstanding field if non-nil, zero value otherwise.

### GetCommonStockSharesOutstandingOk

`func (o *FundamentalsFinancial) GetCommonStockSharesOutstandingOk() (*string, bool)`

GetCommonStockSharesOutstandingOk returns a tuple with the CommonStockSharesOutstanding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonStockSharesOutstanding

`func (o *FundamentalsFinancial) SetCommonStockSharesOutstanding(v string)`

SetCommonStockSharesOutstanding sets CommonStockSharesOutstanding field to given value.

### HasCommonStockSharesOutstanding

`func (o *FundamentalsFinancial) HasCommonStockSharesOutstanding() bool`

HasCommonStockSharesOutstanding returns a boolean if a field has been set.

### SetCommonStockSharesOutstandingNil

`func (o *FundamentalsFinancial) SetCommonStockSharesOutstandingNil(b bool)`

 SetCommonStockSharesOutstandingNil sets the value for CommonStockSharesOutstanding to be an explicit nil

### UnsetCommonStockSharesOutstanding
`func (o *FundamentalsFinancial) UnsetCommonStockSharesOutstanding()`

UnsetCommonStockSharesOutstanding ensures that no value is present for CommonStockSharesOutstanding, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


