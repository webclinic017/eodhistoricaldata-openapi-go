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

// FundamentalsEarningsAnnual struct for FundamentalsEarningsAnnual
type FundamentalsEarningsAnnual struct {
	Date      *string  `json:"date,omitempty"`
	EpsActual *float64 `json:"epsActual,omitempty"`
}

// NewFundamentalsEarningsAnnual instantiates a new FundamentalsEarningsAnnual object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFundamentalsEarningsAnnual() *FundamentalsEarningsAnnual {
	this := FundamentalsEarningsAnnual{}
	return &this
}

// NewFundamentalsEarningsAnnualWithDefaults instantiates a new FundamentalsEarningsAnnual object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFundamentalsEarningsAnnualWithDefaults() *FundamentalsEarningsAnnual {
	this := FundamentalsEarningsAnnual{}
	return &this
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *FundamentalsEarningsAnnual) GetDate() string {
	if o == nil || o.Date == nil {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsEarningsAnnual) GetDateOk() (*string, bool) {
	if o == nil || o.Date == nil {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *FundamentalsEarningsAnnual) HasDate() bool {
	if o != nil && o.Date != nil {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *FundamentalsEarningsAnnual) SetDate(v string) {
	o.Date = &v
}

// GetEpsActual returns the EpsActual field value if set, zero value otherwise.
func (o *FundamentalsEarningsAnnual) GetEpsActual() float64 {
	if o == nil || o.EpsActual == nil {
		var ret float64
		return ret
	}
	return *o.EpsActual
}

// GetEpsActualOk returns a tuple with the EpsActual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsEarningsAnnual) GetEpsActualOk() (*float64, bool) {
	if o == nil || o.EpsActual == nil {
		return nil, false
	}
	return o.EpsActual, true
}

// HasEpsActual returns a boolean if a field has been set.
func (o *FundamentalsEarningsAnnual) HasEpsActual() bool {
	if o != nil && o.EpsActual != nil {
		return true
	}

	return false
}

// SetEpsActual gets a reference to the given float64 and assigns it to the EpsActual field.
func (o *FundamentalsEarningsAnnual) SetEpsActual(v float64) {
	o.EpsActual = &v
}

func (o FundamentalsEarningsAnnual) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Date != nil {
		toSerialize["date"] = o.Date
	}
	if o.EpsActual != nil {
		toSerialize["epsActual"] = o.EpsActual
	}
	return json.Marshal(toSerialize)
}

type NullableFundamentalsEarningsAnnual struct {
	value *FundamentalsEarningsAnnual
	isSet bool
}

func (v NullableFundamentalsEarningsAnnual) Get() *FundamentalsEarningsAnnual {
	return v.value
}

func (v *NullableFundamentalsEarningsAnnual) Set(val *FundamentalsEarningsAnnual) {
	v.value = val
	v.isSet = true
}

func (v NullableFundamentalsEarningsAnnual) IsSet() bool {
	return v.isSet
}

func (v *NullableFundamentalsEarningsAnnual) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFundamentalsEarningsAnnual(val *FundamentalsEarningsAnnual) *NullableFundamentalsEarningsAnnual {
	return &NullableFundamentalsEarningsAnnual{value: val, isSet: true}
}

func (v NullableFundamentalsEarningsAnnual) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFundamentalsEarningsAnnual) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
