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

// Fundamentals struct for Fundamentals
type Fundamentals struct {
	General                    *FundamentalsGeneral                            `json:"General,omitempty"`
	Components                 *FundamentalsComponents                         `json:"Components,omitempty"`
	HistoricalTickerComponents *[]FundamentalsHistoricalTickerComponentsData   `json:"HistoricalTickerComponents,omitempty"`
	Highlights                 *FundamentalsHighlights                         `json:"Highlights,omitempty"`
	Valuation                  *FundamentalsValuation                          `json:"Valuation,omitempty"`
	SharesStats                *FundamentalsSharesStats                        `json:"SharesStats,omitempty"`
	Technicals                 *FundamentalsTechnicals                         `json:"Technicals,omitempty"`
	SplitsDividends            *FundamentalsSplitsDividends                    `json:"SplitsDividends,omitempty"`
	AnalystRatings             *FundamentalsAnalystRatings                     `json:"AnalystRatings,omitempty"`
	Holders                    *FundamentalsHolders                            `json:"Holders,omitempty"`
	InsiderTransactions        *map[string]FundamentalsInsiderTransactionsData `json:"InsiderTransactions,omitempty"`
	ESGScores                  *FundamentalsESGScores                          `json:"ESGScores,omitempty"`
	OutstandingShares          *FundamentalsOutstandingShares                  `json:"outstandingShares,omitempty"`
	Earnings                   *FundamentalsEarnings                           `json:"Earnings,omitempty"`
	Financials                 *FundamentalsFinancials                         `json:"Financials,omitempty"`
}

// NewFundamentals instantiates a new Fundamentals object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFundamentals() *Fundamentals {
	this := Fundamentals{}
	return &this
}

// NewFundamentalsWithDefaults instantiates a new Fundamentals object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFundamentalsWithDefaults() *Fundamentals {
	this := Fundamentals{}
	return &this
}

// GetGeneral returns the General field value if set, zero value otherwise.
func (o *Fundamentals) GetGeneral() FundamentalsGeneral {
	if o == nil || o.General == nil {
		var ret FundamentalsGeneral
		return ret
	}
	return *o.General
}

// GetGeneralOk returns a tuple with the General field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fundamentals) GetGeneralOk() (*FundamentalsGeneral, bool) {
	if o == nil || o.General == nil {
		return nil, false
	}
	return o.General, true
}

// HasGeneral returns a boolean if a field has been set.
func (o *Fundamentals) HasGeneral() bool {
	if o != nil && o.General != nil {
		return true
	}

	return false
}

// SetGeneral gets a reference to the given FundamentalsGeneral and assigns it to the General field.
func (o *Fundamentals) SetGeneral(v FundamentalsGeneral) {
	o.General = &v
}

// GetComponents returns the Components field value if set, zero value otherwise.
func (o *Fundamentals) GetComponents() FundamentalsComponents {
	if o == nil || o.Components == nil {
		var ret FundamentalsComponents
		return ret
	}
	return *o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fundamentals) GetComponentsOk() (*FundamentalsComponents, bool) {
	if o == nil || o.Components == nil {
		return nil, false
	}
	return o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *Fundamentals) HasComponents() bool {
	if o != nil && o.Components != nil {
		return true
	}

	return false
}

// SetComponents gets a reference to the given FundamentalsComponents and assigns it to the Components field.
func (o *Fundamentals) SetComponents(v FundamentalsComponents) {
	o.Components = &v
}

// GetHistoricalTickerComponents returns the HistoricalTickerComponents field value if set, zero value otherwise.
func (o *Fundamentals) GetHistoricalTickerComponents() []FundamentalsHistoricalTickerComponentsData {
	if o == nil || o.HistoricalTickerComponents == nil {
		var ret []FundamentalsHistoricalTickerComponentsData
		return ret
	}
	return *o.HistoricalTickerComponents
}

// GetHistoricalTickerComponentsOk returns a tuple with the HistoricalTickerComponents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fundamentals) GetHistoricalTickerComponentsOk() (*[]FundamentalsHistoricalTickerComponentsData, bool) {
	if o == nil || o.HistoricalTickerComponents == nil {
		return nil, false
	}
	return o.HistoricalTickerComponents, true
}

// HasHistoricalTickerComponents returns a boolean if a field has been set.
func (o *Fundamentals) HasHistoricalTickerComponents() bool {
	if o != nil && o.HistoricalTickerComponents != nil {
		return true
	}

	return false
}

// SetHistoricalTickerComponents gets a reference to the given []FundamentalsHistoricalTickerComponentsData and assigns it to the HistoricalTickerComponents field.
func (o *Fundamentals) SetHistoricalTickerComponents(v []FundamentalsHistoricalTickerComponentsData) {
	o.HistoricalTickerComponents = &v
}

// GetHighlights returns the Highlights field value if set, zero value otherwise.
func (o *Fundamentals) GetHighlights() FundamentalsHighlights {
	if o == nil || o.Highlights == nil {
		var ret FundamentalsHighlights
		return ret
	}
	return *o.Highlights
}

// GetHighlightsOk returns a tuple with the Highlights field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fundamentals) GetHighlightsOk() (*FundamentalsHighlights, bool) {
	if o == nil || o.Highlights == nil {
		return nil, false
	}
	return o.Highlights, true
}

// HasHighlights returns a boolean if a field has been set.
func (o *Fundamentals) HasHighlights() bool {
	if o != nil && o.Highlights != nil {
		return true
	}

	return false
}

// SetHighlights gets a reference to the given FundamentalsHighlights and assigns it to the Highlights field.
func (o *Fundamentals) SetHighlights(v FundamentalsHighlights) {
	o.Highlights = &v
}

// GetValuation returns the Valuation field value if set, zero value otherwise.
func (o *Fundamentals) GetValuation() FundamentalsValuation {
	if o == nil || o.Valuation == nil {
		var ret FundamentalsValuation
		return ret
	}
	return *o.Valuation
}

// GetValuationOk returns a tuple with the Valuation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fundamentals) GetValuationOk() (*FundamentalsValuation, bool) {
	if o == nil || o.Valuation == nil {
		return nil, false
	}
	return o.Valuation, true
}

// HasValuation returns a boolean if a field has been set.
func (o *Fundamentals) HasValuation() bool {
	if o != nil && o.Valuation != nil {
		return true
	}

	return false
}

// SetValuation gets a reference to the given FundamentalsValuation and assigns it to the Valuation field.
func (o *Fundamentals) SetValuation(v FundamentalsValuation) {
	o.Valuation = &v
}

// GetSharesStats returns the SharesStats field value if set, zero value otherwise.
func (o *Fundamentals) GetSharesStats() FundamentalsSharesStats {
	if o == nil || o.SharesStats == nil {
		var ret FundamentalsSharesStats
		return ret
	}
	return *o.SharesStats
}

// GetSharesStatsOk returns a tuple with the SharesStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fundamentals) GetSharesStatsOk() (*FundamentalsSharesStats, bool) {
	if o == nil || o.SharesStats == nil {
		return nil, false
	}
	return o.SharesStats, true
}

// HasSharesStats returns a boolean if a field has been set.
func (o *Fundamentals) HasSharesStats() bool {
	if o != nil && o.SharesStats != nil {
		return true
	}

	return false
}

// SetSharesStats gets a reference to the given FundamentalsSharesStats and assigns it to the SharesStats field.
func (o *Fundamentals) SetSharesStats(v FundamentalsSharesStats) {
	o.SharesStats = &v
}

// GetTechnicals returns the Technicals field value if set, zero value otherwise.
func (o *Fundamentals) GetTechnicals() FundamentalsTechnicals {
	if o == nil || o.Technicals == nil {
		var ret FundamentalsTechnicals
		return ret
	}
	return *o.Technicals
}

// GetTechnicalsOk returns a tuple with the Technicals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fundamentals) GetTechnicalsOk() (*FundamentalsTechnicals, bool) {
	if o == nil || o.Technicals == nil {
		return nil, false
	}
	return o.Technicals, true
}

// HasTechnicals returns a boolean if a field has been set.
func (o *Fundamentals) HasTechnicals() bool {
	if o != nil && o.Technicals != nil {
		return true
	}

	return false
}

// SetTechnicals gets a reference to the given FundamentalsTechnicals and assigns it to the Technicals field.
func (o *Fundamentals) SetTechnicals(v FundamentalsTechnicals) {
	o.Technicals = &v
}

// GetSplitsDividends returns the SplitsDividends field value if set, zero value otherwise.
func (o *Fundamentals) GetSplitsDividends() FundamentalsSplitsDividends {
	if o == nil || o.SplitsDividends == nil {
		var ret FundamentalsSplitsDividends
		return ret
	}
	return *o.SplitsDividends
}

// GetSplitsDividendsOk returns a tuple with the SplitsDividends field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fundamentals) GetSplitsDividendsOk() (*FundamentalsSplitsDividends, bool) {
	if o == nil || o.SplitsDividends == nil {
		return nil, false
	}
	return o.SplitsDividends, true
}

// HasSplitsDividends returns a boolean if a field has been set.
func (o *Fundamentals) HasSplitsDividends() bool {
	if o != nil && o.SplitsDividends != nil {
		return true
	}

	return false
}

// SetSplitsDividends gets a reference to the given FundamentalsSplitsDividends and assigns it to the SplitsDividends field.
func (o *Fundamentals) SetSplitsDividends(v FundamentalsSplitsDividends) {
	o.SplitsDividends = &v
}

// GetAnalystRatings returns the AnalystRatings field value if set, zero value otherwise.
func (o *Fundamentals) GetAnalystRatings() FundamentalsAnalystRatings {
	if o == nil || o.AnalystRatings == nil {
		var ret FundamentalsAnalystRatings
		return ret
	}
	return *o.AnalystRatings
}

// GetAnalystRatingsOk returns a tuple with the AnalystRatings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fundamentals) GetAnalystRatingsOk() (*FundamentalsAnalystRatings, bool) {
	if o == nil || o.AnalystRatings == nil {
		return nil, false
	}
	return o.AnalystRatings, true
}

// HasAnalystRatings returns a boolean if a field has been set.
func (o *Fundamentals) HasAnalystRatings() bool {
	if o != nil && o.AnalystRatings != nil {
		return true
	}

	return false
}

// SetAnalystRatings gets a reference to the given FundamentalsAnalystRatings and assigns it to the AnalystRatings field.
func (o *Fundamentals) SetAnalystRatings(v FundamentalsAnalystRatings) {
	o.AnalystRatings = &v
}

// GetHolders returns the Holders field value if set, zero value otherwise.
func (o *Fundamentals) GetHolders() FundamentalsHolders {
	if o == nil || o.Holders == nil {
		var ret FundamentalsHolders
		return ret
	}
	return *o.Holders
}

// GetHoldersOk returns a tuple with the Holders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fundamentals) GetHoldersOk() (*FundamentalsHolders, bool) {
	if o == nil || o.Holders == nil {
		return nil, false
	}
	return o.Holders, true
}

// HasHolders returns a boolean if a field has been set.
func (o *Fundamentals) HasHolders() bool {
	if o != nil && o.Holders != nil {
		return true
	}

	return false
}

// SetHolders gets a reference to the given FundamentalsHolders and assigns it to the Holders field.
func (o *Fundamentals) SetHolders(v FundamentalsHolders) {
	o.Holders = &v
}

// GetInsiderTransactions returns the InsiderTransactions field value if set, zero value otherwise.
func (o *Fundamentals) GetInsiderTransactions() map[string]FundamentalsInsiderTransactionsData {
	if o == nil || o.InsiderTransactions == nil {
		var ret map[string]FundamentalsInsiderTransactionsData
		return ret
	}
	return *o.InsiderTransactions
}

// GetInsiderTransactionsOk returns a tuple with the InsiderTransactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fundamentals) GetInsiderTransactionsOk() (*map[string]FundamentalsInsiderTransactionsData, bool) {
	if o == nil || o.InsiderTransactions == nil {
		return nil, false
	}
	return o.InsiderTransactions, true
}

// HasInsiderTransactions returns a boolean if a field has been set.
func (o *Fundamentals) HasInsiderTransactions() bool {
	if o != nil && o.InsiderTransactions != nil {
		return true
	}

	return false
}

// SetInsiderTransactions gets a reference to the given map[string]FundamentalsInsiderTransactionsData and assigns it to the InsiderTransactions field.
func (o *Fundamentals) SetInsiderTransactions(v map[string]FundamentalsInsiderTransactionsData) {
	o.InsiderTransactions = &v
}

// GetESGScores returns the ESGScores field value if set, zero value otherwise.
func (o *Fundamentals) GetESGScores() FundamentalsESGScores {
	if o == nil || o.ESGScores == nil {
		var ret FundamentalsESGScores
		return ret
	}
	return *o.ESGScores
}

// GetESGScoresOk returns a tuple with the ESGScores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fundamentals) GetESGScoresOk() (*FundamentalsESGScores, bool) {
	if o == nil || o.ESGScores == nil {
		return nil, false
	}
	return o.ESGScores, true
}

// HasESGScores returns a boolean if a field has been set.
func (o *Fundamentals) HasESGScores() bool {
	if o != nil && o.ESGScores != nil {
		return true
	}

	return false
}

// SetESGScores gets a reference to the given FundamentalsESGScores and assigns it to the ESGScores field.
func (o *Fundamentals) SetESGScores(v FundamentalsESGScores) {
	o.ESGScores = &v
}

// GetOutstandingShares returns the OutstandingShares field value if set, zero value otherwise.
func (o *Fundamentals) GetOutstandingShares() FundamentalsOutstandingShares {
	if o == nil || o.OutstandingShares == nil {
		var ret FundamentalsOutstandingShares
		return ret
	}
	return *o.OutstandingShares
}

// GetOutstandingSharesOk returns a tuple with the OutstandingShares field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fundamentals) GetOutstandingSharesOk() (*FundamentalsOutstandingShares, bool) {
	if o == nil || o.OutstandingShares == nil {
		return nil, false
	}
	return o.OutstandingShares, true
}

// HasOutstandingShares returns a boolean if a field has been set.
func (o *Fundamentals) HasOutstandingShares() bool {
	if o != nil && o.OutstandingShares != nil {
		return true
	}

	return false
}

// SetOutstandingShares gets a reference to the given FundamentalsOutstandingShares and assigns it to the OutstandingShares field.
func (o *Fundamentals) SetOutstandingShares(v FundamentalsOutstandingShares) {
	o.OutstandingShares = &v
}

// GetEarnings returns the Earnings field value if set, zero value otherwise.
func (o *Fundamentals) GetEarnings() FundamentalsEarnings {
	if o == nil || o.Earnings == nil {
		var ret FundamentalsEarnings
		return ret
	}
	return *o.Earnings
}

// GetEarningsOk returns a tuple with the Earnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fundamentals) GetEarningsOk() (*FundamentalsEarnings, bool) {
	if o == nil || o.Earnings == nil {
		return nil, false
	}
	return o.Earnings, true
}

// HasEarnings returns a boolean if a field has been set.
func (o *Fundamentals) HasEarnings() bool {
	if o != nil && o.Earnings != nil {
		return true
	}

	return false
}

// SetEarnings gets a reference to the given FundamentalsEarnings and assigns it to the Earnings field.
func (o *Fundamentals) SetEarnings(v FundamentalsEarnings) {
	o.Earnings = &v
}

// GetFinancials returns the Financials field value if set, zero value otherwise.
func (o *Fundamentals) GetFinancials() FundamentalsFinancials {
	if o == nil || o.Financials == nil {
		var ret FundamentalsFinancials
		return ret
	}
	return *o.Financials
}

// GetFinancialsOk returns a tuple with the Financials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fundamentals) GetFinancialsOk() (*FundamentalsFinancials, bool) {
	if o == nil || o.Financials == nil {
		return nil, false
	}
	return o.Financials, true
}

// HasFinancials returns a boolean if a field has been set.
func (o *Fundamentals) HasFinancials() bool {
	if o != nil && o.Financials != nil {
		return true
	}

	return false
}

// SetFinancials gets a reference to the given FundamentalsFinancials and assigns it to the Financials field.
func (o *Fundamentals) SetFinancials(v FundamentalsFinancials) {
	o.Financials = &v
}

func (o Fundamentals) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.General != nil {
		toSerialize["General"] = o.General
	}
	if o.Components != nil {
		toSerialize["Components"] = o.Components
	}
	if o.HistoricalTickerComponents != nil {
		toSerialize["HistoricalTickerComponents"] = o.HistoricalTickerComponents
	}
	if o.Highlights != nil {
		toSerialize["Highlights"] = o.Highlights
	}
	if o.Valuation != nil {
		toSerialize["Valuation"] = o.Valuation
	}
	if o.SharesStats != nil {
		toSerialize["SharesStats"] = o.SharesStats
	}
	if o.Technicals != nil {
		toSerialize["Technicals"] = o.Technicals
	}
	if o.SplitsDividends != nil {
		toSerialize["SplitsDividends"] = o.SplitsDividends
	}
	if o.AnalystRatings != nil {
		toSerialize["AnalystRatings"] = o.AnalystRatings
	}
	if o.Holders != nil {
		toSerialize["Holders"] = o.Holders
	}
	if o.InsiderTransactions != nil {
		toSerialize["InsiderTransactions"] = o.InsiderTransactions
	}
	if o.ESGScores != nil {
		toSerialize["ESGScores"] = o.ESGScores
	}
	if o.OutstandingShares != nil {
		toSerialize["outstandingShares"] = o.OutstandingShares
	}
	if o.Earnings != nil {
		toSerialize["Earnings"] = o.Earnings
	}
	if o.Financials != nil {
		toSerialize["Financials"] = o.Financials
	}
	return json.Marshal(toSerialize)
}

type NullableFundamentals struct {
	value *Fundamentals
	isSet bool
}

func (v NullableFundamentals) Get() *Fundamentals {
	return v.value
}

func (v *NullableFundamentals) Set(val *Fundamentals) {
	v.value = val
	v.isSet = true
}

func (v NullableFundamentals) IsSet() bool {
	return v.isSet
}

func (v *NullableFundamentals) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFundamentals(val *Fundamentals) *NullableFundamentals {
	return &NullableFundamentals{value: val, isSet: true}
}

func (v NullableFundamentals) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFundamentals) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
