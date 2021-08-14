# IncomeStatement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **string** |  | [optional] 
**FilingDate** | Pointer to **string** |  | [optional] 
**CurrencySymbol** | Pointer to **string** |  | [optional] 
**ResearchDevelopment** | Pointer to **NullableString** |  | [optional] 
**EffectOfAccountingCharges** | Pointer to **NullableString** |  | [optional] 
**IncomeBeforeTax** | Pointer to **NullableString** |  | [optional] 
**MinorityInterest** | Pointer to **NullableString** |  | [optional] 
**NetIncome** | Pointer to **NullableString** |  | [optional] 
**SellingGeneralAdministrative** | Pointer to **NullableString** |  | [optional] 
**SellingAndMarketingExpenses** | Pointer to **NullableString** |  | [optional] 
**GrossProfit** | Pointer to **NullableString** |  | [optional] 
**ReconciledDepreciation** | Pointer to **NullableString** |  | [optional] 
**Ebit** | Pointer to **NullableString** |  | [optional] 
**Ebitda** | Pointer to **NullableString** |  | [optional] 
**DepreciationAndAmortization** | Pointer to **NullableString** |  | [optional] 
**NonOperatingIncomeNetOther** | Pointer to **NullableString** |  | [optional] 
**OperatingIncome** | Pointer to **NullableString** |  | [optional] 
**OtherOperatingExpenses** | Pointer to **NullableString** |  | [optional] 
**InterestExpense** | Pointer to **NullableString** |  | [optional] 
**TaxProvision** | Pointer to **NullableString** |  | [optional] 
**InterestIncome** | Pointer to **NullableString** |  | [optional] 
**NetInterestIncome** | Pointer to **NullableString** |  | [optional] 
**ExtraordinaryItems** | Pointer to **NullableString** |  | [optional] 
**NonRecurring** | Pointer to **NullableString** |  | [optional] 
**OtherItems** | Pointer to **NullableString** |  | [optional] 
**IncomeTaxExpense** | Pointer to **NullableString** |  | [optional] 
**TotalRevenue** | Pointer to **NullableString** |  | [optional] 
**TotalOperatingExpenses** | Pointer to **NullableString** |  | [optional] 
**CostOfRevenue** | Pointer to **NullableString** |  | [optional] 
**TotalOtherIncomeExpenseNet** | Pointer to **NullableString** |  | [optional] 
**DiscontinuedOperations** | Pointer to **NullableString** |  | [optional] 
**NetIncomeFromContinuingOps** | Pointer to **NullableString** |  | [optional] 
**NetIncomeApplicableToCommonShares** | Pointer to **NullableString** |  | [optional] 
**PreferredStockAndOtherAdjustments** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewIncomeStatement

`func NewIncomeStatement() *IncomeStatement`

NewIncomeStatement instantiates a new IncomeStatement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncomeStatementWithDefaults

`func NewIncomeStatementWithDefaults() *IncomeStatement`

NewIncomeStatementWithDefaults instantiates a new IncomeStatement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *IncomeStatement) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *IncomeStatement) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *IncomeStatement) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *IncomeStatement) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetFilingDate

`func (o *IncomeStatement) GetFilingDate() string`

GetFilingDate returns the FilingDate field if non-nil, zero value otherwise.

### GetFilingDateOk

`func (o *IncomeStatement) GetFilingDateOk() (*string, bool)`

GetFilingDateOk returns a tuple with the FilingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilingDate

`func (o *IncomeStatement) SetFilingDate(v string)`

SetFilingDate sets FilingDate field to given value.

### HasFilingDate

`func (o *IncomeStatement) HasFilingDate() bool`

HasFilingDate returns a boolean if a field has been set.

### GetCurrencySymbol

`func (o *IncomeStatement) GetCurrencySymbol() string`

GetCurrencySymbol returns the CurrencySymbol field if non-nil, zero value otherwise.

### GetCurrencySymbolOk

`func (o *IncomeStatement) GetCurrencySymbolOk() (*string, bool)`

GetCurrencySymbolOk returns a tuple with the CurrencySymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencySymbol

`func (o *IncomeStatement) SetCurrencySymbol(v string)`

SetCurrencySymbol sets CurrencySymbol field to given value.

### HasCurrencySymbol

`func (o *IncomeStatement) HasCurrencySymbol() bool`

HasCurrencySymbol returns a boolean if a field has been set.

### GetResearchDevelopment

`func (o *IncomeStatement) GetResearchDevelopment() string`

GetResearchDevelopment returns the ResearchDevelopment field if non-nil, zero value otherwise.

### GetResearchDevelopmentOk

`func (o *IncomeStatement) GetResearchDevelopmentOk() (*string, bool)`

GetResearchDevelopmentOk returns a tuple with the ResearchDevelopment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResearchDevelopment

`func (o *IncomeStatement) SetResearchDevelopment(v string)`

SetResearchDevelopment sets ResearchDevelopment field to given value.

### HasResearchDevelopment

`func (o *IncomeStatement) HasResearchDevelopment() bool`

HasResearchDevelopment returns a boolean if a field has been set.

### SetResearchDevelopmentNil

`func (o *IncomeStatement) SetResearchDevelopmentNil(b bool)`

 SetResearchDevelopmentNil sets the value for ResearchDevelopment to be an explicit nil

### UnsetResearchDevelopment
`func (o *IncomeStatement) UnsetResearchDevelopment()`

UnsetResearchDevelopment ensures that no value is present for ResearchDevelopment, not even an explicit nil
### GetEffectOfAccountingCharges

`func (o *IncomeStatement) GetEffectOfAccountingCharges() string`

GetEffectOfAccountingCharges returns the EffectOfAccountingCharges field if non-nil, zero value otherwise.

### GetEffectOfAccountingChargesOk

`func (o *IncomeStatement) GetEffectOfAccountingChargesOk() (*string, bool)`

GetEffectOfAccountingChargesOk returns a tuple with the EffectOfAccountingCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectOfAccountingCharges

`func (o *IncomeStatement) SetEffectOfAccountingCharges(v string)`

SetEffectOfAccountingCharges sets EffectOfAccountingCharges field to given value.

### HasEffectOfAccountingCharges

`func (o *IncomeStatement) HasEffectOfAccountingCharges() bool`

HasEffectOfAccountingCharges returns a boolean if a field has been set.

### SetEffectOfAccountingChargesNil

`func (o *IncomeStatement) SetEffectOfAccountingChargesNil(b bool)`

 SetEffectOfAccountingChargesNil sets the value for EffectOfAccountingCharges to be an explicit nil

### UnsetEffectOfAccountingCharges
`func (o *IncomeStatement) UnsetEffectOfAccountingCharges()`

UnsetEffectOfAccountingCharges ensures that no value is present for EffectOfAccountingCharges, not even an explicit nil
### GetIncomeBeforeTax

`func (o *IncomeStatement) GetIncomeBeforeTax() string`

GetIncomeBeforeTax returns the IncomeBeforeTax field if non-nil, zero value otherwise.

### GetIncomeBeforeTaxOk

`func (o *IncomeStatement) GetIncomeBeforeTaxOk() (*string, bool)`

GetIncomeBeforeTaxOk returns a tuple with the IncomeBeforeTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomeBeforeTax

`func (o *IncomeStatement) SetIncomeBeforeTax(v string)`

SetIncomeBeforeTax sets IncomeBeforeTax field to given value.

### HasIncomeBeforeTax

`func (o *IncomeStatement) HasIncomeBeforeTax() bool`

HasIncomeBeforeTax returns a boolean if a field has been set.

### SetIncomeBeforeTaxNil

`func (o *IncomeStatement) SetIncomeBeforeTaxNil(b bool)`

 SetIncomeBeforeTaxNil sets the value for IncomeBeforeTax to be an explicit nil

### UnsetIncomeBeforeTax
`func (o *IncomeStatement) UnsetIncomeBeforeTax()`

UnsetIncomeBeforeTax ensures that no value is present for IncomeBeforeTax, not even an explicit nil
### GetMinorityInterest

`func (o *IncomeStatement) GetMinorityInterest() string`

GetMinorityInterest returns the MinorityInterest field if non-nil, zero value otherwise.

### GetMinorityInterestOk

`func (o *IncomeStatement) GetMinorityInterestOk() (*string, bool)`

GetMinorityInterestOk returns a tuple with the MinorityInterest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinorityInterest

`func (o *IncomeStatement) SetMinorityInterest(v string)`

SetMinorityInterest sets MinorityInterest field to given value.

### HasMinorityInterest

`func (o *IncomeStatement) HasMinorityInterest() bool`

HasMinorityInterest returns a boolean if a field has been set.

### SetMinorityInterestNil

`func (o *IncomeStatement) SetMinorityInterestNil(b bool)`

 SetMinorityInterestNil sets the value for MinorityInterest to be an explicit nil

### UnsetMinorityInterest
`func (o *IncomeStatement) UnsetMinorityInterest()`

UnsetMinorityInterest ensures that no value is present for MinorityInterest, not even an explicit nil
### GetNetIncome

`func (o *IncomeStatement) GetNetIncome() string`

GetNetIncome returns the NetIncome field if non-nil, zero value otherwise.

### GetNetIncomeOk

`func (o *IncomeStatement) GetNetIncomeOk() (*string, bool)`

GetNetIncomeOk returns a tuple with the NetIncome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetIncome

`func (o *IncomeStatement) SetNetIncome(v string)`

SetNetIncome sets NetIncome field to given value.

### HasNetIncome

`func (o *IncomeStatement) HasNetIncome() bool`

HasNetIncome returns a boolean if a field has been set.

### SetNetIncomeNil

`func (o *IncomeStatement) SetNetIncomeNil(b bool)`

 SetNetIncomeNil sets the value for NetIncome to be an explicit nil

### UnsetNetIncome
`func (o *IncomeStatement) UnsetNetIncome()`

UnsetNetIncome ensures that no value is present for NetIncome, not even an explicit nil
### GetSellingGeneralAdministrative

`func (o *IncomeStatement) GetSellingGeneralAdministrative() string`

GetSellingGeneralAdministrative returns the SellingGeneralAdministrative field if non-nil, zero value otherwise.

### GetSellingGeneralAdministrativeOk

`func (o *IncomeStatement) GetSellingGeneralAdministrativeOk() (*string, bool)`

GetSellingGeneralAdministrativeOk returns a tuple with the SellingGeneralAdministrative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellingGeneralAdministrative

`func (o *IncomeStatement) SetSellingGeneralAdministrative(v string)`

SetSellingGeneralAdministrative sets SellingGeneralAdministrative field to given value.

### HasSellingGeneralAdministrative

`func (o *IncomeStatement) HasSellingGeneralAdministrative() bool`

HasSellingGeneralAdministrative returns a boolean if a field has been set.

### SetSellingGeneralAdministrativeNil

`func (o *IncomeStatement) SetSellingGeneralAdministrativeNil(b bool)`

 SetSellingGeneralAdministrativeNil sets the value for SellingGeneralAdministrative to be an explicit nil

### UnsetSellingGeneralAdministrative
`func (o *IncomeStatement) UnsetSellingGeneralAdministrative()`

UnsetSellingGeneralAdministrative ensures that no value is present for SellingGeneralAdministrative, not even an explicit nil
### GetSellingAndMarketingExpenses

`func (o *IncomeStatement) GetSellingAndMarketingExpenses() string`

GetSellingAndMarketingExpenses returns the SellingAndMarketingExpenses field if non-nil, zero value otherwise.

### GetSellingAndMarketingExpensesOk

`func (o *IncomeStatement) GetSellingAndMarketingExpensesOk() (*string, bool)`

GetSellingAndMarketingExpensesOk returns a tuple with the SellingAndMarketingExpenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellingAndMarketingExpenses

`func (o *IncomeStatement) SetSellingAndMarketingExpenses(v string)`

SetSellingAndMarketingExpenses sets SellingAndMarketingExpenses field to given value.

### HasSellingAndMarketingExpenses

`func (o *IncomeStatement) HasSellingAndMarketingExpenses() bool`

HasSellingAndMarketingExpenses returns a boolean if a field has been set.

### SetSellingAndMarketingExpensesNil

`func (o *IncomeStatement) SetSellingAndMarketingExpensesNil(b bool)`

 SetSellingAndMarketingExpensesNil sets the value for SellingAndMarketingExpenses to be an explicit nil

### UnsetSellingAndMarketingExpenses
`func (o *IncomeStatement) UnsetSellingAndMarketingExpenses()`

UnsetSellingAndMarketingExpenses ensures that no value is present for SellingAndMarketingExpenses, not even an explicit nil
### GetGrossProfit

`func (o *IncomeStatement) GetGrossProfit() string`

GetGrossProfit returns the GrossProfit field if non-nil, zero value otherwise.

### GetGrossProfitOk

`func (o *IncomeStatement) GetGrossProfitOk() (*string, bool)`

GetGrossProfitOk returns a tuple with the GrossProfit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossProfit

`func (o *IncomeStatement) SetGrossProfit(v string)`

SetGrossProfit sets GrossProfit field to given value.

### HasGrossProfit

`func (o *IncomeStatement) HasGrossProfit() bool`

HasGrossProfit returns a boolean if a field has been set.

### SetGrossProfitNil

`func (o *IncomeStatement) SetGrossProfitNil(b bool)`

 SetGrossProfitNil sets the value for GrossProfit to be an explicit nil

### UnsetGrossProfit
`func (o *IncomeStatement) UnsetGrossProfit()`

UnsetGrossProfit ensures that no value is present for GrossProfit, not even an explicit nil
### GetReconciledDepreciation

`func (o *IncomeStatement) GetReconciledDepreciation() string`

GetReconciledDepreciation returns the ReconciledDepreciation field if non-nil, zero value otherwise.

### GetReconciledDepreciationOk

`func (o *IncomeStatement) GetReconciledDepreciationOk() (*string, bool)`

GetReconciledDepreciationOk returns a tuple with the ReconciledDepreciation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReconciledDepreciation

`func (o *IncomeStatement) SetReconciledDepreciation(v string)`

SetReconciledDepreciation sets ReconciledDepreciation field to given value.

### HasReconciledDepreciation

`func (o *IncomeStatement) HasReconciledDepreciation() bool`

HasReconciledDepreciation returns a boolean if a field has been set.

### SetReconciledDepreciationNil

`func (o *IncomeStatement) SetReconciledDepreciationNil(b bool)`

 SetReconciledDepreciationNil sets the value for ReconciledDepreciation to be an explicit nil

### UnsetReconciledDepreciation
`func (o *IncomeStatement) UnsetReconciledDepreciation()`

UnsetReconciledDepreciation ensures that no value is present for ReconciledDepreciation, not even an explicit nil
### GetEbit

`func (o *IncomeStatement) GetEbit() string`

GetEbit returns the Ebit field if non-nil, zero value otherwise.

### GetEbitOk

`func (o *IncomeStatement) GetEbitOk() (*string, bool)`

GetEbitOk returns a tuple with the Ebit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbit

`func (o *IncomeStatement) SetEbit(v string)`

SetEbit sets Ebit field to given value.

### HasEbit

`func (o *IncomeStatement) HasEbit() bool`

HasEbit returns a boolean if a field has been set.

### SetEbitNil

`func (o *IncomeStatement) SetEbitNil(b bool)`

 SetEbitNil sets the value for Ebit to be an explicit nil

### UnsetEbit
`func (o *IncomeStatement) UnsetEbit()`

UnsetEbit ensures that no value is present for Ebit, not even an explicit nil
### GetEbitda

`func (o *IncomeStatement) GetEbitda() string`

GetEbitda returns the Ebitda field if non-nil, zero value otherwise.

### GetEbitdaOk

`func (o *IncomeStatement) GetEbitdaOk() (*string, bool)`

GetEbitdaOk returns a tuple with the Ebitda field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbitda

`func (o *IncomeStatement) SetEbitda(v string)`

SetEbitda sets Ebitda field to given value.

### HasEbitda

`func (o *IncomeStatement) HasEbitda() bool`

HasEbitda returns a boolean if a field has been set.

### SetEbitdaNil

`func (o *IncomeStatement) SetEbitdaNil(b bool)`

 SetEbitdaNil sets the value for Ebitda to be an explicit nil

### UnsetEbitda
`func (o *IncomeStatement) UnsetEbitda()`

UnsetEbitda ensures that no value is present for Ebitda, not even an explicit nil
### GetDepreciationAndAmortization

`func (o *IncomeStatement) GetDepreciationAndAmortization() string`

GetDepreciationAndAmortization returns the DepreciationAndAmortization field if non-nil, zero value otherwise.

### GetDepreciationAndAmortizationOk

`func (o *IncomeStatement) GetDepreciationAndAmortizationOk() (*string, bool)`

GetDepreciationAndAmortizationOk returns a tuple with the DepreciationAndAmortization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepreciationAndAmortization

`func (o *IncomeStatement) SetDepreciationAndAmortization(v string)`

SetDepreciationAndAmortization sets DepreciationAndAmortization field to given value.

### HasDepreciationAndAmortization

`func (o *IncomeStatement) HasDepreciationAndAmortization() bool`

HasDepreciationAndAmortization returns a boolean if a field has been set.

### SetDepreciationAndAmortizationNil

`func (o *IncomeStatement) SetDepreciationAndAmortizationNil(b bool)`

 SetDepreciationAndAmortizationNil sets the value for DepreciationAndAmortization to be an explicit nil

### UnsetDepreciationAndAmortization
`func (o *IncomeStatement) UnsetDepreciationAndAmortization()`

UnsetDepreciationAndAmortization ensures that no value is present for DepreciationAndAmortization, not even an explicit nil
### GetNonOperatingIncomeNetOther

`func (o *IncomeStatement) GetNonOperatingIncomeNetOther() string`

GetNonOperatingIncomeNetOther returns the NonOperatingIncomeNetOther field if non-nil, zero value otherwise.

### GetNonOperatingIncomeNetOtherOk

`func (o *IncomeStatement) GetNonOperatingIncomeNetOtherOk() (*string, bool)`

GetNonOperatingIncomeNetOtherOk returns a tuple with the NonOperatingIncomeNetOther field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonOperatingIncomeNetOther

`func (o *IncomeStatement) SetNonOperatingIncomeNetOther(v string)`

SetNonOperatingIncomeNetOther sets NonOperatingIncomeNetOther field to given value.

### HasNonOperatingIncomeNetOther

`func (o *IncomeStatement) HasNonOperatingIncomeNetOther() bool`

HasNonOperatingIncomeNetOther returns a boolean if a field has been set.

### SetNonOperatingIncomeNetOtherNil

`func (o *IncomeStatement) SetNonOperatingIncomeNetOtherNil(b bool)`

 SetNonOperatingIncomeNetOtherNil sets the value for NonOperatingIncomeNetOther to be an explicit nil

### UnsetNonOperatingIncomeNetOther
`func (o *IncomeStatement) UnsetNonOperatingIncomeNetOther()`

UnsetNonOperatingIncomeNetOther ensures that no value is present for NonOperatingIncomeNetOther, not even an explicit nil
### GetOperatingIncome

`func (o *IncomeStatement) GetOperatingIncome() string`

GetOperatingIncome returns the OperatingIncome field if non-nil, zero value otherwise.

### GetOperatingIncomeOk

`func (o *IncomeStatement) GetOperatingIncomeOk() (*string, bool)`

GetOperatingIncomeOk returns a tuple with the OperatingIncome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingIncome

`func (o *IncomeStatement) SetOperatingIncome(v string)`

SetOperatingIncome sets OperatingIncome field to given value.

### HasOperatingIncome

`func (o *IncomeStatement) HasOperatingIncome() bool`

HasOperatingIncome returns a boolean if a field has been set.

### SetOperatingIncomeNil

`func (o *IncomeStatement) SetOperatingIncomeNil(b bool)`

 SetOperatingIncomeNil sets the value for OperatingIncome to be an explicit nil

### UnsetOperatingIncome
`func (o *IncomeStatement) UnsetOperatingIncome()`

UnsetOperatingIncome ensures that no value is present for OperatingIncome, not even an explicit nil
### GetOtherOperatingExpenses

`func (o *IncomeStatement) GetOtherOperatingExpenses() string`

GetOtherOperatingExpenses returns the OtherOperatingExpenses field if non-nil, zero value otherwise.

### GetOtherOperatingExpensesOk

`func (o *IncomeStatement) GetOtherOperatingExpensesOk() (*string, bool)`

GetOtherOperatingExpensesOk returns a tuple with the OtherOperatingExpenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherOperatingExpenses

`func (o *IncomeStatement) SetOtherOperatingExpenses(v string)`

SetOtherOperatingExpenses sets OtherOperatingExpenses field to given value.

### HasOtherOperatingExpenses

`func (o *IncomeStatement) HasOtherOperatingExpenses() bool`

HasOtherOperatingExpenses returns a boolean if a field has been set.

### SetOtherOperatingExpensesNil

`func (o *IncomeStatement) SetOtherOperatingExpensesNil(b bool)`

 SetOtherOperatingExpensesNil sets the value for OtherOperatingExpenses to be an explicit nil

### UnsetOtherOperatingExpenses
`func (o *IncomeStatement) UnsetOtherOperatingExpenses()`

UnsetOtherOperatingExpenses ensures that no value is present for OtherOperatingExpenses, not even an explicit nil
### GetInterestExpense

`func (o *IncomeStatement) GetInterestExpense() string`

GetInterestExpense returns the InterestExpense field if non-nil, zero value otherwise.

### GetInterestExpenseOk

`func (o *IncomeStatement) GetInterestExpenseOk() (*string, bool)`

GetInterestExpenseOk returns a tuple with the InterestExpense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestExpense

`func (o *IncomeStatement) SetInterestExpense(v string)`

SetInterestExpense sets InterestExpense field to given value.

### HasInterestExpense

`func (o *IncomeStatement) HasInterestExpense() bool`

HasInterestExpense returns a boolean if a field has been set.

### SetInterestExpenseNil

`func (o *IncomeStatement) SetInterestExpenseNil(b bool)`

 SetInterestExpenseNil sets the value for InterestExpense to be an explicit nil

### UnsetInterestExpense
`func (o *IncomeStatement) UnsetInterestExpense()`

UnsetInterestExpense ensures that no value is present for InterestExpense, not even an explicit nil
### GetTaxProvision

`func (o *IncomeStatement) GetTaxProvision() string`

GetTaxProvision returns the TaxProvision field if non-nil, zero value otherwise.

### GetTaxProvisionOk

`func (o *IncomeStatement) GetTaxProvisionOk() (*string, bool)`

GetTaxProvisionOk returns a tuple with the TaxProvision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxProvision

`func (o *IncomeStatement) SetTaxProvision(v string)`

SetTaxProvision sets TaxProvision field to given value.

### HasTaxProvision

`func (o *IncomeStatement) HasTaxProvision() bool`

HasTaxProvision returns a boolean if a field has been set.

### SetTaxProvisionNil

`func (o *IncomeStatement) SetTaxProvisionNil(b bool)`

 SetTaxProvisionNil sets the value for TaxProvision to be an explicit nil

### UnsetTaxProvision
`func (o *IncomeStatement) UnsetTaxProvision()`

UnsetTaxProvision ensures that no value is present for TaxProvision, not even an explicit nil
### GetInterestIncome

`func (o *IncomeStatement) GetInterestIncome() string`

GetInterestIncome returns the InterestIncome field if non-nil, zero value otherwise.

### GetInterestIncomeOk

`func (o *IncomeStatement) GetInterestIncomeOk() (*string, bool)`

GetInterestIncomeOk returns a tuple with the InterestIncome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestIncome

`func (o *IncomeStatement) SetInterestIncome(v string)`

SetInterestIncome sets InterestIncome field to given value.

### HasInterestIncome

`func (o *IncomeStatement) HasInterestIncome() bool`

HasInterestIncome returns a boolean if a field has been set.

### SetInterestIncomeNil

`func (o *IncomeStatement) SetInterestIncomeNil(b bool)`

 SetInterestIncomeNil sets the value for InterestIncome to be an explicit nil

### UnsetInterestIncome
`func (o *IncomeStatement) UnsetInterestIncome()`

UnsetInterestIncome ensures that no value is present for InterestIncome, not even an explicit nil
### GetNetInterestIncome

`func (o *IncomeStatement) GetNetInterestIncome() string`

GetNetInterestIncome returns the NetInterestIncome field if non-nil, zero value otherwise.

### GetNetInterestIncomeOk

`func (o *IncomeStatement) GetNetInterestIncomeOk() (*string, bool)`

GetNetInterestIncomeOk returns a tuple with the NetInterestIncome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetInterestIncome

`func (o *IncomeStatement) SetNetInterestIncome(v string)`

SetNetInterestIncome sets NetInterestIncome field to given value.

### HasNetInterestIncome

`func (o *IncomeStatement) HasNetInterestIncome() bool`

HasNetInterestIncome returns a boolean if a field has been set.

### SetNetInterestIncomeNil

`func (o *IncomeStatement) SetNetInterestIncomeNil(b bool)`

 SetNetInterestIncomeNil sets the value for NetInterestIncome to be an explicit nil

### UnsetNetInterestIncome
`func (o *IncomeStatement) UnsetNetInterestIncome()`

UnsetNetInterestIncome ensures that no value is present for NetInterestIncome, not even an explicit nil
### GetExtraordinaryItems

`func (o *IncomeStatement) GetExtraordinaryItems() string`

GetExtraordinaryItems returns the ExtraordinaryItems field if non-nil, zero value otherwise.

### GetExtraordinaryItemsOk

`func (o *IncomeStatement) GetExtraordinaryItemsOk() (*string, bool)`

GetExtraordinaryItemsOk returns a tuple with the ExtraordinaryItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraordinaryItems

`func (o *IncomeStatement) SetExtraordinaryItems(v string)`

SetExtraordinaryItems sets ExtraordinaryItems field to given value.

### HasExtraordinaryItems

`func (o *IncomeStatement) HasExtraordinaryItems() bool`

HasExtraordinaryItems returns a boolean if a field has been set.

### SetExtraordinaryItemsNil

`func (o *IncomeStatement) SetExtraordinaryItemsNil(b bool)`

 SetExtraordinaryItemsNil sets the value for ExtraordinaryItems to be an explicit nil

### UnsetExtraordinaryItems
`func (o *IncomeStatement) UnsetExtraordinaryItems()`

UnsetExtraordinaryItems ensures that no value is present for ExtraordinaryItems, not even an explicit nil
### GetNonRecurring

`func (o *IncomeStatement) GetNonRecurring() string`

GetNonRecurring returns the NonRecurring field if non-nil, zero value otherwise.

### GetNonRecurringOk

`func (o *IncomeStatement) GetNonRecurringOk() (*string, bool)`

GetNonRecurringOk returns a tuple with the NonRecurring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonRecurring

`func (o *IncomeStatement) SetNonRecurring(v string)`

SetNonRecurring sets NonRecurring field to given value.

### HasNonRecurring

`func (o *IncomeStatement) HasNonRecurring() bool`

HasNonRecurring returns a boolean if a field has been set.

### SetNonRecurringNil

`func (o *IncomeStatement) SetNonRecurringNil(b bool)`

 SetNonRecurringNil sets the value for NonRecurring to be an explicit nil

### UnsetNonRecurring
`func (o *IncomeStatement) UnsetNonRecurring()`

UnsetNonRecurring ensures that no value is present for NonRecurring, not even an explicit nil
### GetOtherItems

`func (o *IncomeStatement) GetOtherItems() string`

GetOtherItems returns the OtherItems field if non-nil, zero value otherwise.

### GetOtherItemsOk

`func (o *IncomeStatement) GetOtherItemsOk() (*string, bool)`

GetOtherItemsOk returns a tuple with the OtherItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherItems

`func (o *IncomeStatement) SetOtherItems(v string)`

SetOtherItems sets OtherItems field to given value.

### HasOtherItems

`func (o *IncomeStatement) HasOtherItems() bool`

HasOtherItems returns a boolean if a field has been set.

### SetOtherItemsNil

`func (o *IncomeStatement) SetOtherItemsNil(b bool)`

 SetOtherItemsNil sets the value for OtherItems to be an explicit nil

### UnsetOtherItems
`func (o *IncomeStatement) UnsetOtherItems()`

UnsetOtherItems ensures that no value is present for OtherItems, not even an explicit nil
### GetIncomeTaxExpense

`func (o *IncomeStatement) GetIncomeTaxExpense() string`

GetIncomeTaxExpense returns the IncomeTaxExpense field if non-nil, zero value otherwise.

### GetIncomeTaxExpenseOk

`func (o *IncomeStatement) GetIncomeTaxExpenseOk() (*string, bool)`

GetIncomeTaxExpenseOk returns a tuple with the IncomeTaxExpense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomeTaxExpense

`func (o *IncomeStatement) SetIncomeTaxExpense(v string)`

SetIncomeTaxExpense sets IncomeTaxExpense field to given value.

### HasIncomeTaxExpense

`func (o *IncomeStatement) HasIncomeTaxExpense() bool`

HasIncomeTaxExpense returns a boolean if a field has been set.

### SetIncomeTaxExpenseNil

`func (o *IncomeStatement) SetIncomeTaxExpenseNil(b bool)`

 SetIncomeTaxExpenseNil sets the value for IncomeTaxExpense to be an explicit nil

### UnsetIncomeTaxExpense
`func (o *IncomeStatement) UnsetIncomeTaxExpense()`

UnsetIncomeTaxExpense ensures that no value is present for IncomeTaxExpense, not even an explicit nil
### GetTotalRevenue

`func (o *IncomeStatement) GetTotalRevenue() string`

GetTotalRevenue returns the TotalRevenue field if non-nil, zero value otherwise.

### GetTotalRevenueOk

`func (o *IncomeStatement) GetTotalRevenueOk() (*string, bool)`

GetTotalRevenueOk returns a tuple with the TotalRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRevenue

`func (o *IncomeStatement) SetTotalRevenue(v string)`

SetTotalRevenue sets TotalRevenue field to given value.

### HasTotalRevenue

`func (o *IncomeStatement) HasTotalRevenue() bool`

HasTotalRevenue returns a boolean if a field has been set.

### SetTotalRevenueNil

`func (o *IncomeStatement) SetTotalRevenueNil(b bool)`

 SetTotalRevenueNil sets the value for TotalRevenue to be an explicit nil

### UnsetTotalRevenue
`func (o *IncomeStatement) UnsetTotalRevenue()`

UnsetTotalRevenue ensures that no value is present for TotalRevenue, not even an explicit nil
### GetTotalOperatingExpenses

`func (o *IncomeStatement) GetTotalOperatingExpenses() string`

GetTotalOperatingExpenses returns the TotalOperatingExpenses field if non-nil, zero value otherwise.

### GetTotalOperatingExpensesOk

`func (o *IncomeStatement) GetTotalOperatingExpensesOk() (*string, bool)`

GetTotalOperatingExpensesOk returns a tuple with the TotalOperatingExpenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalOperatingExpenses

`func (o *IncomeStatement) SetTotalOperatingExpenses(v string)`

SetTotalOperatingExpenses sets TotalOperatingExpenses field to given value.

### HasTotalOperatingExpenses

`func (o *IncomeStatement) HasTotalOperatingExpenses() bool`

HasTotalOperatingExpenses returns a boolean if a field has been set.

### SetTotalOperatingExpensesNil

`func (o *IncomeStatement) SetTotalOperatingExpensesNil(b bool)`

 SetTotalOperatingExpensesNil sets the value for TotalOperatingExpenses to be an explicit nil

### UnsetTotalOperatingExpenses
`func (o *IncomeStatement) UnsetTotalOperatingExpenses()`

UnsetTotalOperatingExpenses ensures that no value is present for TotalOperatingExpenses, not even an explicit nil
### GetCostOfRevenue

`func (o *IncomeStatement) GetCostOfRevenue() string`

GetCostOfRevenue returns the CostOfRevenue field if non-nil, zero value otherwise.

### GetCostOfRevenueOk

`func (o *IncomeStatement) GetCostOfRevenueOk() (*string, bool)`

GetCostOfRevenueOk returns a tuple with the CostOfRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostOfRevenue

`func (o *IncomeStatement) SetCostOfRevenue(v string)`

SetCostOfRevenue sets CostOfRevenue field to given value.

### HasCostOfRevenue

`func (o *IncomeStatement) HasCostOfRevenue() bool`

HasCostOfRevenue returns a boolean if a field has been set.

### SetCostOfRevenueNil

`func (o *IncomeStatement) SetCostOfRevenueNil(b bool)`

 SetCostOfRevenueNil sets the value for CostOfRevenue to be an explicit nil

### UnsetCostOfRevenue
`func (o *IncomeStatement) UnsetCostOfRevenue()`

UnsetCostOfRevenue ensures that no value is present for CostOfRevenue, not even an explicit nil
### GetTotalOtherIncomeExpenseNet

`func (o *IncomeStatement) GetTotalOtherIncomeExpenseNet() string`

GetTotalOtherIncomeExpenseNet returns the TotalOtherIncomeExpenseNet field if non-nil, zero value otherwise.

### GetTotalOtherIncomeExpenseNetOk

`func (o *IncomeStatement) GetTotalOtherIncomeExpenseNetOk() (*string, bool)`

GetTotalOtherIncomeExpenseNetOk returns a tuple with the TotalOtherIncomeExpenseNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalOtherIncomeExpenseNet

`func (o *IncomeStatement) SetTotalOtherIncomeExpenseNet(v string)`

SetTotalOtherIncomeExpenseNet sets TotalOtherIncomeExpenseNet field to given value.

### HasTotalOtherIncomeExpenseNet

`func (o *IncomeStatement) HasTotalOtherIncomeExpenseNet() bool`

HasTotalOtherIncomeExpenseNet returns a boolean if a field has been set.

### SetTotalOtherIncomeExpenseNetNil

`func (o *IncomeStatement) SetTotalOtherIncomeExpenseNetNil(b bool)`

 SetTotalOtherIncomeExpenseNetNil sets the value for TotalOtherIncomeExpenseNet to be an explicit nil

### UnsetTotalOtherIncomeExpenseNet
`func (o *IncomeStatement) UnsetTotalOtherIncomeExpenseNet()`

UnsetTotalOtherIncomeExpenseNet ensures that no value is present for TotalOtherIncomeExpenseNet, not even an explicit nil
### GetDiscontinuedOperations

`func (o *IncomeStatement) GetDiscontinuedOperations() string`

GetDiscontinuedOperations returns the DiscontinuedOperations field if non-nil, zero value otherwise.

### GetDiscontinuedOperationsOk

`func (o *IncomeStatement) GetDiscontinuedOperationsOk() (*string, bool)`

GetDiscontinuedOperationsOk returns a tuple with the DiscontinuedOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscontinuedOperations

`func (o *IncomeStatement) SetDiscontinuedOperations(v string)`

SetDiscontinuedOperations sets DiscontinuedOperations field to given value.

### HasDiscontinuedOperations

`func (o *IncomeStatement) HasDiscontinuedOperations() bool`

HasDiscontinuedOperations returns a boolean if a field has been set.

### SetDiscontinuedOperationsNil

`func (o *IncomeStatement) SetDiscontinuedOperationsNil(b bool)`

 SetDiscontinuedOperationsNil sets the value for DiscontinuedOperations to be an explicit nil

### UnsetDiscontinuedOperations
`func (o *IncomeStatement) UnsetDiscontinuedOperations()`

UnsetDiscontinuedOperations ensures that no value is present for DiscontinuedOperations, not even an explicit nil
### GetNetIncomeFromContinuingOps

`func (o *IncomeStatement) GetNetIncomeFromContinuingOps() string`

GetNetIncomeFromContinuingOps returns the NetIncomeFromContinuingOps field if non-nil, zero value otherwise.

### GetNetIncomeFromContinuingOpsOk

`func (o *IncomeStatement) GetNetIncomeFromContinuingOpsOk() (*string, bool)`

GetNetIncomeFromContinuingOpsOk returns a tuple with the NetIncomeFromContinuingOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetIncomeFromContinuingOps

`func (o *IncomeStatement) SetNetIncomeFromContinuingOps(v string)`

SetNetIncomeFromContinuingOps sets NetIncomeFromContinuingOps field to given value.

### HasNetIncomeFromContinuingOps

`func (o *IncomeStatement) HasNetIncomeFromContinuingOps() bool`

HasNetIncomeFromContinuingOps returns a boolean if a field has been set.

### SetNetIncomeFromContinuingOpsNil

`func (o *IncomeStatement) SetNetIncomeFromContinuingOpsNil(b bool)`

 SetNetIncomeFromContinuingOpsNil sets the value for NetIncomeFromContinuingOps to be an explicit nil

### UnsetNetIncomeFromContinuingOps
`func (o *IncomeStatement) UnsetNetIncomeFromContinuingOps()`

UnsetNetIncomeFromContinuingOps ensures that no value is present for NetIncomeFromContinuingOps, not even an explicit nil
### GetNetIncomeApplicableToCommonShares

`func (o *IncomeStatement) GetNetIncomeApplicableToCommonShares() string`

GetNetIncomeApplicableToCommonShares returns the NetIncomeApplicableToCommonShares field if non-nil, zero value otherwise.

### GetNetIncomeApplicableToCommonSharesOk

`func (o *IncomeStatement) GetNetIncomeApplicableToCommonSharesOk() (*string, bool)`

GetNetIncomeApplicableToCommonSharesOk returns a tuple with the NetIncomeApplicableToCommonShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetIncomeApplicableToCommonShares

`func (o *IncomeStatement) SetNetIncomeApplicableToCommonShares(v string)`

SetNetIncomeApplicableToCommonShares sets NetIncomeApplicableToCommonShares field to given value.

### HasNetIncomeApplicableToCommonShares

`func (o *IncomeStatement) HasNetIncomeApplicableToCommonShares() bool`

HasNetIncomeApplicableToCommonShares returns a boolean if a field has been set.

### SetNetIncomeApplicableToCommonSharesNil

`func (o *IncomeStatement) SetNetIncomeApplicableToCommonSharesNil(b bool)`

 SetNetIncomeApplicableToCommonSharesNil sets the value for NetIncomeApplicableToCommonShares to be an explicit nil

### UnsetNetIncomeApplicableToCommonShares
`func (o *IncomeStatement) UnsetNetIncomeApplicableToCommonShares()`

UnsetNetIncomeApplicableToCommonShares ensures that no value is present for NetIncomeApplicableToCommonShares, not even an explicit nil
### GetPreferredStockAndOtherAdjustments

`func (o *IncomeStatement) GetPreferredStockAndOtherAdjustments() string`

GetPreferredStockAndOtherAdjustments returns the PreferredStockAndOtherAdjustments field if non-nil, zero value otherwise.

### GetPreferredStockAndOtherAdjustmentsOk

`func (o *IncomeStatement) GetPreferredStockAndOtherAdjustmentsOk() (*string, bool)`

GetPreferredStockAndOtherAdjustmentsOk returns a tuple with the PreferredStockAndOtherAdjustments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredStockAndOtherAdjustments

`func (o *IncomeStatement) SetPreferredStockAndOtherAdjustments(v string)`

SetPreferredStockAndOtherAdjustments sets PreferredStockAndOtherAdjustments field to given value.

### HasPreferredStockAndOtherAdjustments

`func (o *IncomeStatement) HasPreferredStockAndOtherAdjustments() bool`

HasPreferredStockAndOtherAdjustments returns a boolean if a field has been set.

### SetPreferredStockAndOtherAdjustmentsNil

`func (o *IncomeStatement) SetPreferredStockAndOtherAdjustmentsNil(b bool)`

 SetPreferredStockAndOtherAdjustmentsNil sets the value for PreferredStockAndOtherAdjustments to be an explicit nil

### UnsetPreferredStockAndOtherAdjustments
`func (o *IncomeStatement) UnsetPreferredStockAndOtherAdjustments()`

UnsetPreferredStockAndOtherAdjustments ensures that no value is present for PreferredStockAndOtherAdjustments, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


