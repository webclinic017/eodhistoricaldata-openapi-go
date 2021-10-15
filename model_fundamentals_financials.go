/*
eodhistoricaldata

eodhistoricaldata API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// FundamentalsFinancials struct for FundamentalsFinancials
type FundamentalsFinancials struct {
	BalanceSheet    *FundamentalsFinancialsBalanceSheet    `json:"Balance_Sheet,omitempty"`
	CashFlow        *FundamentalsFinancialsCashFlow        `json:"Cash_Flow,omitempty"`
	IncomeStatement *FundamentalsFinancialsIncomeStatement `json:"Income_Statement,omitempty"`
}

// NewFundamentalsFinancials instantiates a new FundamentalsFinancials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFundamentalsFinancials() *FundamentalsFinancials {
	this := FundamentalsFinancials{}
	return &this
}

// NewFundamentalsFinancialsWithDefaults instantiates a new FundamentalsFinancials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFundamentalsFinancialsWithDefaults() *FundamentalsFinancials {
	this := FundamentalsFinancials{}
	return &this
}

// GetBalanceSheet returns the BalanceSheet field value if set, zero value otherwise.
func (o *FundamentalsFinancials) GetBalanceSheet() FundamentalsFinancialsBalanceSheet {
	if o == nil || o.BalanceSheet == nil {
		var ret FundamentalsFinancialsBalanceSheet
		return ret
	}
	return *o.BalanceSheet
}

// GetBalanceSheetOk returns a tuple with the BalanceSheet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsFinancials) GetBalanceSheetOk() (*FundamentalsFinancialsBalanceSheet, bool) {
	if o == nil || o.BalanceSheet == nil {
		return nil, false
	}
	return o.BalanceSheet, true
}

// HasBalanceSheet returns a boolean if a field has been set.
func (o *FundamentalsFinancials) HasBalanceSheet() bool {
	if o != nil && o.BalanceSheet != nil {
		return true
	}

	return false
}

// SetBalanceSheet gets a reference to the given FundamentalsFinancialsBalanceSheet and assigns it to the BalanceSheet field.
func (o *FundamentalsFinancials) SetBalanceSheet(v FundamentalsFinancialsBalanceSheet) {
	o.BalanceSheet = &v
}

// GetCashFlow returns the CashFlow field value if set, zero value otherwise.
func (o *FundamentalsFinancials) GetCashFlow() FundamentalsFinancialsCashFlow {
	if o == nil || o.CashFlow == nil {
		var ret FundamentalsFinancialsCashFlow
		return ret
	}
	return *o.CashFlow
}

// GetCashFlowOk returns a tuple with the CashFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsFinancials) GetCashFlowOk() (*FundamentalsFinancialsCashFlow, bool) {
	if o == nil || o.CashFlow == nil {
		return nil, false
	}
	return o.CashFlow, true
}

// HasCashFlow returns a boolean if a field has been set.
func (o *FundamentalsFinancials) HasCashFlow() bool {
	if o != nil && o.CashFlow != nil {
		return true
	}

	return false
}

// SetCashFlow gets a reference to the given FundamentalsFinancialsCashFlow and assigns it to the CashFlow field.
func (o *FundamentalsFinancials) SetCashFlow(v FundamentalsFinancialsCashFlow) {
	o.CashFlow = &v
}

// GetIncomeStatement returns the IncomeStatement field value if set, zero value otherwise.
func (o *FundamentalsFinancials) GetIncomeStatement() FundamentalsFinancialsIncomeStatement {
	if o == nil || o.IncomeStatement == nil {
		var ret FundamentalsFinancialsIncomeStatement
		return ret
	}
	return *o.IncomeStatement
}

// GetIncomeStatementOk returns a tuple with the IncomeStatement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsFinancials) GetIncomeStatementOk() (*FundamentalsFinancialsIncomeStatement, bool) {
	if o == nil || o.IncomeStatement == nil {
		return nil, false
	}
	return o.IncomeStatement, true
}

// HasIncomeStatement returns a boolean if a field has been set.
func (o *FundamentalsFinancials) HasIncomeStatement() bool {
	if o != nil && o.IncomeStatement != nil {
		return true
	}

	return false
}

// SetIncomeStatement gets a reference to the given FundamentalsFinancialsIncomeStatement and assigns it to the IncomeStatement field.
func (o *FundamentalsFinancials) SetIncomeStatement(v FundamentalsFinancialsIncomeStatement) {
	o.IncomeStatement = &v
}

func (o FundamentalsFinancials) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BalanceSheet != nil {
		toSerialize["Balance_Sheet"] = o.BalanceSheet
	}
	if o.CashFlow != nil {
		toSerialize["Cash_Flow"] = o.CashFlow
	}
	if o.IncomeStatement != nil {
		toSerialize["Income_Statement"] = o.IncomeStatement
	}
	return json.Marshal(toSerialize)
}

type NullableFundamentalsFinancials struct {
	value *FundamentalsFinancials
	isSet bool
}

func (v NullableFundamentalsFinancials) Get() *FundamentalsFinancials {
	return v.value
}

func (v *NullableFundamentalsFinancials) Set(val *FundamentalsFinancials) {
	v.value = val
	v.isSet = true
}

func (v NullableFundamentalsFinancials) IsSet() bool {
	return v.isSet
}

func (v *NullableFundamentalsFinancials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFundamentalsFinancials(val *FundamentalsFinancials) *NullableFundamentalsFinancials {
	return &NullableFundamentalsFinancials{value: val, isSet: true}
}

func (v NullableFundamentalsFinancials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFundamentalsFinancials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
