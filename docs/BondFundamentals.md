# BondFundamentals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Isin** | Pointer to **string** |  | [optional] 
**Cusip** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**WKN** | Pointer to **string** |  | [optional] 
**Sedol** | Pointer to **string** |  | [optional] 
**FIGI** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Coupon** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**LastTradeDate** | Pointer to **string** |  | [optional] 
**YieldToMaturity** | Pointer to **string** |  | [optional] 
**Callable** | Pointer to **string** |  | [optional] 
**NextCallDate** | Pointer to **string** |  | [optional] 
**MinimumSettlementAmount** | Pointer to **string** |  | [optional] 
**ParIntegralMultiple** | Pointer to **string** |  | [optional] 
**Classification** | Pointer to [**BondFundamentalsClassification**](BondFundamentalsClassification.md) |  | [optional] 
**Rating** | Pointer to [**BondFundamentalsRating**](BondFundamentalsRating.md) |  | [optional] 
**IssueData** | Pointer to [**BondFundamentalsIssueData**](BondFundamentalsIssueData.md) |  | [optional] 

## Methods

### NewBondFundamentals

`func NewBondFundamentals() *BondFundamentals`

NewBondFundamentals instantiates a new BondFundamentals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBondFundamentalsWithDefaults

`func NewBondFundamentalsWithDefaults() *BondFundamentals`

NewBondFundamentalsWithDefaults instantiates a new BondFundamentals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsin

`func (o *BondFundamentals) GetIsin() string`

GetIsin returns the Isin field if non-nil, zero value otherwise.

### GetIsinOk

`func (o *BondFundamentals) GetIsinOk() (*string, bool)`

GetIsinOk returns a tuple with the Isin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsin

`func (o *BondFundamentals) SetIsin(v string)`

SetIsin sets Isin field to given value.

### HasIsin

`func (o *BondFundamentals) HasIsin() bool`

HasIsin returns a boolean if a field has been set.

### GetCusip

`func (o *BondFundamentals) GetCusip() string`

GetCusip returns the Cusip field if non-nil, zero value otherwise.

### GetCusipOk

`func (o *BondFundamentals) GetCusipOk() (*string, bool)`

GetCusipOk returns a tuple with the Cusip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCusip

`func (o *BondFundamentals) SetCusip(v string)`

SetCusip sets Cusip field to given value.

### HasCusip

`func (o *BondFundamentals) HasCusip() bool`

HasCusip returns a boolean if a field has been set.

### GetName

`func (o *BondFundamentals) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BondFundamentals) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BondFundamentals) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BondFundamentals) HasName() bool`

HasName returns a boolean if a field has been set.

### GetWKN

`func (o *BondFundamentals) GetWKN() string`

GetWKN returns the WKN field if non-nil, zero value otherwise.

### GetWKNOk

`func (o *BondFundamentals) GetWKNOk() (*string, bool)`

GetWKNOk returns a tuple with the WKN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWKN

`func (o *BondFundamentals) SetWKN(v string)`

SetWKN sets WKN field to given value.

### HasWKN

`func (o *BondFundamentals) HasWKN() bool`

HasWKN returns a boolean if a field has been set.

### GetSedol

`func (o *BondFundamentals) GetSedol() string`

GetSedol returns the Sedol field if non-nil, zero value otherwise.

### GetSedolOk

`func (o *BondFundamentals) GetSedolOk() (*string, bool)`

GetSedolOk returns a tuple with the Sedol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSedol

`func (o *BondFundamentals) SetSedol(v string)`

SetSedol sets Sedol field to given value.

### HasSedol

`func (o *BondFundamentals) HasSedol() bool`

HasSedol returns a boolean if a field has been set.

### GetFIGI

`func (o *BondFundamentals) GetFIGI() string`

GetFIGI returns the FIGI field if non-nil, zero value otherwise.

### GetFIGIOk

`func (o *BondFundamentals) GetFIGIOk() (*string, bool)`

GetFIGIOk returns a tuple with the FIGI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFIGI

`func (o *BondFundamentals) SetFIGI(v string)`

SetFIGI sets FIGI field to given value.

### HasFIGI

`func (o *BondFundamentals) HasFIGI() bool`

HasFIGI returns a boolean if a field has been set.

### GetCurrency

`func (o *BondFundamentals) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BondFundamentals) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BondFundamentals) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *BondFundamentals) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCoupon

`func (o *BondFundamentals) GetCoupon() string`

GetCoupon returns the Coupon field if non-nil, zero value otherwise.

### GetCouponOk

`func (o *BondFundamentals) GetCouponOk() (*string, bool)`

GetCouponOk returns a tuple with the Coupon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoupon

`func (o *BondFundamentals) SetCoupon(v string)`

SetCoupon sets Coupon field to given value.

### HasCoupon

`func (o *BondFundamentals) HasCoupon() bool`

HasCoupon returns a boolean if a field has been set.

### GetPrice

`func (o *BondFundamentals) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *BondFundamentals) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *BondFundamentals) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *BondFundamentals) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetLastTradeDate

`func (o *BondFundamentals) GetLastTradeDate() string`

GetLastTradeDate returns the LastTradeDate field if non-nil, zero value otherwise.

### GetLastTradeDateOk

`func (o *BondFundamentals) GetLastTradeDateOk() (*string, bool)`

GetLastTradeDateOk returns a tuple with the LastTradeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTradeDate

`func (o *BondFundamentals) SetLastTradeDate(v string)`

SetLastTradeDate sets LastTradeDate field to given value.

### HasLastTradeDate

`func (o *BondFundamentals) HasLastTradeDate() bool`

HasLastTradeDate returns a boolean if a field has been set.

### GetYieldToMaturity

`func (o *BondFundamentals) GetYieldToMaturity() string`

GetYieldToMaturity returns the YieldToMaturity field if non-nil, zero value otherwise.

### GetYieldToMaturityOk

`func (o *BondFundamentals) GetYieldToMaturityOk() (*string, bool)`

GetYieldToMaturityOk returns a tuple with the YieldToMaturity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYieldToMaturity

`func (o *BondFundamentals) SetYieldToMaturity(v string)`

SetYieldToMaturity sets YieldToMaturity field to given value.

### HasYieldToMaturity

`func (o *BondFundamentals) HasYieldToMaturity() bool`

HasYieldToMaturity returns a boolean if a field has been set.

### GetCallable

`func (o *BondFundamentals) GetCallable() string`

GetCallable returns the Callable field if non-nil, zero value otherwise.

### GetCallableOk

`func (o *BondFundamentals) GetCallableOk() (*string, bool)`

GetCallableOk returns a tuple with the Callable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallable

`func (o *BondFundamentals) SetCallable(v string)`

SetCallable sets Callable field to given value.

### HasCallable

`func (o *BondFundamentals) HasCallable() bool`

HasCallable returns a boolean if a field has been set.

### GetNextCallDate

`func (o *BondFundamentals) GetNextCallDate() string`

GetNextCallDate returns the NextCallDate field if non-nil, zero value otherwise.

### GetNextCallDateOk

`func (o *BondFundamentals) GetNextCallDateOk() (*string, bool)`

GetNextCallDateOk returns a tuple with the NextCallDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCallDate

`func (o *BondFundamentals) SetNextCallDate(v string)`

SetNextCallDate sets NextCallDate field to given value.

### HasNextCallDate

`func (o *BondFundamentals) HasNextCallDate() bool`

HasNextCallDate returns a boolean if a field has been set.

### GetMinimumSettlementAmount

`func (o *BondFundamentals) GetMinimumSettlementAmount() string`

GetMinimumSettlementAmount returns the MinimumSettlementAmount field if non-nil, zero value otherwise.

### GetMinimumSettlementAmountOk

`func (o *BondFundamentals) GetMinimumSettlementAmountOk() (*string, bool)`

GetMinimumSettlementAmountOk returns a tuple with the MinimumSettlementAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumSettlementAmount

`func (o *BondFundamentals) SetMinimumSettlementAmount(v string)`

SetMinimumSettlementAmount sets MinimumSettlementAmount field to given value.

### HasMinimumSettlementAmount

`func (o *BondFundamentals) HasMinimumSettlementAmount() bool`

HasMinimumSettlementAmount returns a boolean if a field has been set.

### GetParIntegralMultiple

`func (o *BondFundamentals) GetParIntegralMultiple() string`

GetParIntegralMultiple returns the ParIntegralMultiple field if non-nil, zero value otherwise.

### GetParIntegralMultipleOk

`func (o *BondFundamentals) GetParIntegralMultipleOk() (*string, bool)`

GetParIntegralMultipleOk returns a tuple with the ParIntegralMultiple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParIntegralMultiple

`func (o *BondFundamentals) SetParIntegralMultiple(v string)`

SetParIntegralMultiple sets ParIntegralMultiple field to given value.

### HasParIntegralMultiple

`func (o *BondFundamentals) HasParIntegralMultiple() bool`

HasParIntegralMultiple returns a boolean if a field has been set.

### GetClassification

`func (o *BondFundamentals) GetClassification() BondFundamentalsClassification`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *BondFundamentals) GetClassificationOk() (*BondFundamentalsClassification, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassification

`func (o *BondFundamentals) SetClassification(v BondFundamentalsClassification)`

SetClassification sets Classification field to given value.

### HasClassification

`func (o *BondFundamentals) HasClassification() bool`

HasClassification returns a boolean if a field has been set.

### GetRating

`func (o *BondFundamentals) GetRating() BondFundamentalsRating`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *BondFundamentals) GetRatingOk() (*BondFundamentalsRating, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *BondFundamentals) SetRating(v BondFundamentalsRating)`

SetRating sets Rating field to given value.

### HasRating

`func (o *BondFundamentals) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetIssueData

`func (o *BondFundamentals) GetIssueData() BondFundamentalsIssueData`

GetIssueData returns the IssueData field if non-nil, zero value otherwise.

### GetIssueDataOk

`func (o *BondFundamentals) GetIssueDataOk() (*BondFundamentalsIssueData, bool)`

GetIssueDataOk returns a tuple with the IssueData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueData

`func (o *BondFundamentals) SetIssueData(v BondFundamentalsIssueData)`

SetIssueData sets IssueData field to given value.

### HasIssueData

`func (o *BondFundamentals) HasIssueData() bool`

HasIssueData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


