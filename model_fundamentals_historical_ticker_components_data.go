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

// FundamentalsHistoricalTickerComponentsData struct for FundamentalsHistoricalTickerComponentsData
type FundamentalsHistoricalTickerComponentsData struct {
	Code        *string `json:"Code,omitempty"`
	Name        *string `json:"Name,omitempty"`
	StartDate   *string `json:"StartDate,omitempty"`
	EndDate     *string `json:"EndDate,omitempty"`
	IsActiveNow *int32  `json:"IsActiveNow,omitempty"`
	IsDelisted  *int32  `json:"IsDelisted,omitempty"`
}

// NewFundamentalsHistoricalTickerComponentsData instantiates a new FundamentalsHistoricalTickerComponentsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFundamentalsHistoricalTickerComponentsData() *FundamentalsHistoricalTickerComponentsData {
	this := FundamentalsHistoricalTickerComponentsData{}
	return &this
}

// NewFundamentalsHistoricalTickerComponentsDataWithDefaults instantiates a new FundamentalsHistoricalTickerComponentsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFundamentalsHistoricalTickerComponentsDataWithDefaults() *FundamentalsHistoricalTickerComponentsData {
	this := FundamentalsHistoricalTickerComponentsData{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *FundamentalsHistoricalTickerComponentsData) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsHistoricalTickerComponentsData) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *FundamentalsHistoricalTickerComponentsData) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *FundamentalsHistoricalTickerComponentsData) SetCode(v string) {
	o.Code = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FundamentalsHistoricalTickerComponentsData) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsHistoricalTickerComponentsData) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FundamentalsHistoricalTickerComponentsData) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FundamentalsHistoricalTickerComponentsData) SetName(v string) {
	o.Name = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *FundamentalsHistoricalTickerComponentsData) GetStartDate() string {
	if o == nil || o.StartDate == nil {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsHistoricalTickerComponentsData) GetStartDateOk() (*string, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *FundamentalsHistoricalTickerComponentsData) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *FundamentalsHistoricalTickerComponentsData) SetStartDate(v string) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *FundamentalsHistoricalTickerComponentsData) GetEndDate() string {
	if o == nil || o.EndDate == nil {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsHistoricalTickerComponentsData) GetEndDateOk() (*string, bool) {
	if o == nil || o.EndDate == nil {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *FundamentalsHistoricalTickerComponentsData) HasEndDate() bool {
	if o != nil && o.EndDate != nil {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *FundamentalsHistoricalTickerComponentsData) SetEndDate(v string) {
	o.EndDate = &v
}

// GetIsActiveNow returns the IsActiveNow field value if set, zero value otherwise.
func (o *FundamentalsHistoricalTickerComponentsData) GetIsActiveNow() int32 {
	if o == nil || o.IsActiveNow == nil {
		var ret int32
		return ret
	}
	return *o.IsActiveNow
}

// GetIsActiveNowOk returns a tuple with the IsActiveNow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsHistoricalTickerComponentsData) GetIsActiveNowOk() (*int32, bool) {
	if o == nil || o.IsActiveNow == nil {
		return nil, false
	}
	return o.IsActiveNow, true
}

// HasIsActiveNow returns a boolean if a field has been set.
func (o *FundamentalsHistoricalTickerComponentsData) HasIsActiveNow() bool {
	if o != nil && o.IsActiveNow != nil {
		return true
	}

	return false
}

// SetIsActiveNow gets a reference to the given int32 and assigns it to the IsActiveNow field.
func (o *FundamentalsHistoricalTickerComponentsData) SetIsActiveNow(v int32) {
	o.IsActiveNow = &v
}

// GetIsDelisted returns the IsDelisted field value if set, zero value otherwise.
func (o *FundamentalsHistoricalTickerComponentsData) GetIsDelisted() int32 {
	if o == nil || o.IsDelisted == nil {
		var ret int32
		return ret
	}
	return *o.IsDelisted
}

// GetIsDelistedOk returns a tuple with the IsDelisted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsHistoricalTickerComponentsData) GetIsDelistedOk() (*int32, bool) {
	if o == nil || o.IsDelisted == nil {
		return nil, false
	}
	return o.IsDelisted, true
}

// HasIsDelisted returns a boolean if a field has been set.
func (o *FundamentalsHistoricalTickerComponentsData) HasIsDelisted() bool {
	if o != nil && o.IsDelisted != nil {
		return true
	}

	return false
}

// SetIsDelisted gets a reference to the given int32 and assigns it to the IsDelisted field.
func (o *FundamentalsHistoricalTickerComponentsData) SetIsDelisted(v int32) {
	o.IsDelisted = &v
}

func (o FundamentalsHistoricalTickerComponentsData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["Code"] = o.Code
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.StartDate != nil {
		toSerialize["StartDate"] = o.StartDate
	}
	if o.EndDate != nil {
		toSerialize["EndDate"] = o.EndDate
	}
	if o.IsActiveNow != nil {
		toSerialize["IsActiveNow"] = o.IsActiveNow
	}
	if o.IsDelisted != nil {
		toSerialize["IsDelisted"] = o.IsDelisted
	}
	return json.Marshal(toSerialize)
}

type NullableFundamentalsHistoricalTickerComponentsData struct {
	value *FundamentalsHistoricalTickerComponentsData
	isSet bool
}

func (v NullableFundamentalsHistoricalTickerComponentsData) Get() *FundamentalsHistoricalTickerComponentsData {
	return v.value
}

func (v *NullableFundamentalsHistoricalTickerComponentsData) Set(val *FundamentalsHistoricalTickerComponentsData) {
	v.value = val
	v.isSet = true
}

func (v NullableFundamentalsHistoricalTickerComponentsData) IsSet() bool {
	return v.isSet
}

func (v *NullableFundamentalsHistoricalTickerComponentsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFundamentalsHistoricalTickerComponentsData(val *FundamentalsHistoricalTickerComponentsData) *NullableFundamentalsHistoricalTickerComponentsData {
	return &NullableFundamentalsHistoricalTickerComponentsData{value: val, isSet: true}
}

func (v NullableFundamentalsHistoricalTickerComponentsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFundamentalsHistoricalTickerComponentsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}