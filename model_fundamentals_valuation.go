/*
 * eodhistoricaldata
 *
 * eodhistoricaldata API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// FundamentalsValuation struct for FundamentalsValuation
type FundamentalsValuation struct {
	TrailingPE             *float64 `json:"TrailingPE,omitempty"`
	ForwardPE              *float64 `json:"ForwardPE,omitempty"`
	PriceSalesTTM          *float64 `json:"PriceSalesTTM,omitempty"`
	PriceBookMRQ           *float64 `json:"PriceBookMRQ,omitempty"`
	EnterpriseValueRevenue *float64 `json:"EnterpriseValueRevenue,omitempty"`
	EnterpriseValueEbitda  *float64 `json:"EnterpriseValueEbitda,omitempty"`
}

// NewFundamentalsValuation instantiates a new FundamentalsValuation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFundamentalsValuation() *FundamentalsValuation {
	this := FundamentalsValuation{}
	return &this
}

// NewFundamentalsValuationWithDefaults instantiates a new FundamentalsValuation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFundamentalsValuationWithDefaults() *FundamentalsValuation {
	this := FundamentalsValuation{}
	return &this
}

// GetTrailingPE returns the TrailingPE field value if set, zero value otherwise.
func (o *FundamentalsValuation) GetTrailingPE() float64 {
	if o == nil || o.TrailingPE == nil {
		var ret float64
		return ret
	}
	return *o.TrailingPE
}

// GetTrailingPEOk returns a tuple with the TrailingPE field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsValuation) GetTrailingPEOk() (*float64, bool) {
	if o == nil || o.TrailingPE == nil {
		return nil, false
	}
	return o.TrailingPE, true
}

// HasTrailingPE returns a boolean if a field has been set.
func (o *FundamentalsValuation) HasTrailingPE() bool {
	if o != nil && o.TrailingPE != nil {
		return true
	}

	return false
}

// SetTrailingPE gets a reference to the given float64 and assigns it to the TrailingPE field.
func (o *FundamentalsValuation) SetTrailingPE(v float64) {
	o.TrailingPE = &v
}

// GetForwardPE returns the ForwardPE field value if set, zero value otherwise.
func (o *FundamentalsValuation) GetForwardPE() float64 {
	if o == nil || o.ForwardPE == nil {
		var ret float64
		return ret
	}
	return *o.ForwardPE
}

// GetForwardPEOk returns a tuple with the ForwardPE field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsValuation) GetForwardPEOk() (*float64, bool) {
	if o == nil || o.ForwardPE == nil {
		return nil, false
	}
	return o.ForwardPE, true
}

// HasForwardPE returns a boolean if a field has been set.
func (o *FundamentalsValuation) HasForwardPE() bool {
	if o != nil && o.ForwardPE != nil {
		return true
	}

	return false
}

// SetForwardPE gets a reference to the given float64 and assigns it to the ForwardPE field.
func (o *FundamentalsValuation) SetForwardPE(v float64) {
	o.ForwardPE = &v
}

// GetPriceSalesTTM returns the PriceSalesTTM field value if set, zero value otherwise.
func (o *FundamentalsValuation) GetPriceSalesTTM() float64 {
	if o == nil || o.PriceSalesTTM == nil {
		var ret float64
		return ret
	}
	return *o.PriceSalesTTM
}

// GetPriceSalesTTMOk returns a tuple with the PriceSalesTTM field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsValuation) GetPriceSalesTTMOk() (*float64, bool) {
	if o == nil || o.PriceSalesTTM == nil {
		return nil, false
	}
	return o.PriceSalesTTM, true
}

// HasPriceSalesTTM returns a boolean if a field has been set.
func (o *FundamentalsValuation) HasPriceSalesTTM() bool {
	if o != nil && o.PriceSalesTTM != nil {
		return true
	}

	return false
}

// SetPriceSalesTTM gets a reference to the given float64 and assigns it to the PriceSalesTTM field.
func (o *FundamentalsValuation) SetPriceSalesTTM(v float64) {
	o.PriceSalesTTM = &v
}

// GetPriceBookMRQ returns the PriceBookMRQ field value if set, zero value otherwise.
func (o *FundamentalsValuation) GetPriceBookMRQ() float64 {
	if o == nil || o.PriceBookMRQ == nil {
		var ret float64
		return ret
	}
	return *o.PriceBookMRQ
}

// GetPriceBookMRQOk returns a tuple with the PriceBookMRQ field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsValuation) GetPriceBookMRQOk() (*float64, bool) {
	if o == nil || o.PriceBookMRQ == nil {
		return nil, false
	}
	return o.PriceBookMRQ, true
}

// HasPriceBookMRQ returns a boolean if a field has been set.
func (o *FundamentalsValuation) HasPriceBookMRQ() bool {
	if o != nil && o.PriceBookMRQ != nil {
		return true
	}

	return false
}

// SetPriceBookMRQ gets a reference to the given float64 and assigns it to the PriceBookMRQ field.
func (o *FundamentalsValuation) SetPriceBookMRQ(v float64) {
	o.PriceBookMRQ = &v
}

// GetEnterpriseValueRevenue returns the EnterpriseValueRevenue field value if set, zero value otherwise.
func (o *FundamentalsValuation) GetEnterpriseValueRevenue() float64 {
	if o == nil || o.EnterpriseValueRevenue == nil {
		var ret float64
		return ret
	}
	return *o.EnterpriseValueRevenue
}

// GetEnterpriseValueRevenueOk returns a tuple with the EnterpriseValueRevenue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsValuation) GetEnterpriseValueRevenueOk() (*float64, bool) {
	if o == nil || o.EnterpriseValueRevenue == nil {
		return nil, false
	}
	return o.EnterpriseValueRevenue, true
}

// HasEnterpriseValueRevenue returns a boolean if a field has been set.
func (o *FundamentalsValuation) HasEnterpriseValueRevenue() bool {
	if o != nil && o.EnterpriseValueRevenue != nil {
		return true
	}

	return false
}

// SetEnterpriseValueRevenue gets a reference to the given float64 and assigns it to the EnterpriseValueRevenue field.
func (o *FundamentalsValuation) SetEnterpriseValueRevenue(v float64) {
	o.EnterpriseValueRevenue = &v
}

// GetEnterpriseValueEbitda returns the EnterpriseValueEbitda field value if set, zero value otherwise.
func (o *FundamentalsValuation) GetEnterpriseValueEbitda() float64 {
	if o == nil || o.EnterpriseValueEbitda == nil {
		var ret float64
		return ret
	}
	return *o.EnterpriseValueEbitda
}

// GetEnterpriseValueEbitdaOk returns a tuple with the EnterpriseValueEbitda field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsValuation) GetEnterpriseValueEbitdaOk() (*float64, bool) {
	if o == nil || o.EnterpriseValueEbitda == nil {
		return nil, false
	}
	return o.EnterpriseValueEbitda, true
}

// HasEnterpriseValueEbitda returns a boolean if a field has been set.
func (o *FundamentalsValuation) HasEnterpriseValueEbitda() bool {
	if o != nil && o.EnterpriseValueEbitda != nil {
		return true
	}

	return false
}

// SetEnterpriseValueEbitda gets a reference to the given float64 and assigns it to the EnterpriseValueEbitda field.
func (o *FundamentalsValuation) SetEnterpriseValueEbitda(v float64) {
	o.EnterpriseValueEbitda = &v
}

func (o FundamentalsValuation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TrailingPE != nil {
		toSerialize["TrailingPE"] = o.TrailingPE
	}
	if o.ForwardPE != nil {
		toSerialize["ForwardPE"] = o.ForwardPE
	}
	if o.PriceSalesTTM != nil {
		toSerialize["PriceSalesTTM"] = o.PriceSalesTTM
	}
	if o.PriceBookMRQ != nil {
		toSerialize["PriceBookMRQ"] = o.PriceBookMRQ
	}
	if o.EnterpriseValueRevenue != nil {
		toSerialize["EnterpriseValueRevenue"] = o.EnterpriseValueRevenue
	}
	if o.EnterpriseValueEbitda != nil {
		toSerialize["EnterpriseValueEbitda"] = o.EnterpriseValueEbitda
	}
	return json.Marshal(toSerialize)
}

type NullableFundamentalsValuation struct {
	value *FundamentalsValuation
	isSet bool
}

func (v NullableFundamentalsValuation) Get() *FundamentalsValuation {
	return v.value
}

func (v *NullableFundamentalsValuation) Set(val *FundamentalsValuation) {
	v.value = val
	v.isSet = true
}

func (v NullableFundamentalsValuation) IsSet() bool {
	return v.isSet
}

func (v *NullableFundamentalsValuation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFundamentalsValuation(val *FundamentalsValuation) *NullableFundamentalsValuation {
	return &NullableFundamentalsValuation{value: val, isSet: true}
}

func (v NullableFundamentalsValuation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFundamentalsValuation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
