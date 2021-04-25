# FundamentalsGeneral

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Exchange** | Pointer to **string** |  | [optional] 
**CurrencyCode** | Pointer to **string** |  | [optional] 
**CurrencyName** | Pointer to **string** |  | [optional] 
**CurrencySymbol** | Pointer to **string** |  | [optional] 
**CountryName** | Pointer to **string** |  | [optional] 
**CountryISO** | Pointer to **string** |  | [optional] 
**ISIN** | Pointer to **string** |  | [optional] 
**CUSIP** | Pointer to **string** |  | [optional] 
**CIK** | Pointer to **string** |  | [optional] 
**EmployerIdNumber** | Pointer to **string** |  | [optional] 
**FiscalYearEnd** | Pointer to **string** |  | [optional] 
**IPODate** | Pointer to **string** |  | [optional] 
**InternationalDomestic** | Pointer to **string** |  | [optional] 
**Sector** | Pointer to **string** |  | [optional] 
**Industry** | Pointer to **string** |  | [optional] 
**GicSector** | Pointer to **string** |  | [optional] 
**GicGroup** | Pointer to **string** |  | [optional] 
**GicIndustry** | Pointer to **string** |  | [optional] 
**GicSubIndustry** | Pointer to **string** |  | [optional] 
**HomeCategory** | Pointer to **string** |  | [optional] 
**IsDelisted** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Address** | Pointer to **string** |  | [optional] 
**AddressData** | Pointer to [**FundamentalsGeneralAddressData**](FundamentalsGeneralAddressData.md) |  | [optional] 
**Listings** | Pointer to [**map[string]FundamentalsGeneralListings**](FundamentalsGeneralListings.md) |  | [optional] 
**Officers** | Pointer to [**map[string]FundamentalsGeneralOfficers**](FundamentalsGeneralOfficers.md) |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**WebURL** | Pointer to **string** |  | [optional] 
**LogoURL** | Pointer to **string** |  | [optional] 
**FullTimeEmployees** | Pointer to **int64** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewFundamentalsGeneral

`func NewFundamentalsGeneral() *FundamentalsGeneral`

NewFundamentalsGeneral instantiates a new FundamentalsGeneral object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFundamentalsGeneralWithDefaults

`func NewFundamentalsGeneralWithDefaults() *FundamentalsGeneral`

NewFundamentalsGeneralWithDefaults instantiates a new FundamentalsGeneral object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *FundamentalsGeneral) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *FundamentalsGeneral) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *FundamentalsGeneral) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *FundamentalsGeneral) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetType

`func (o *FundamentalsGeneral) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FundamentalsGeneral) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FundamentalsGeneral) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FundamentalsGeneral) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *FundamentalsGeneral) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FundamentalsGeneral) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FundamentalsGeneral) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FundamentalsGeneral) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExchange

`func (o *FundamentalsGeneral) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *FundamentalsGeneral) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *FundamentalsGeneral) SetExchange(v string)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *FundamentalsGeneral) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *FundamentalsGeneral) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *FundamentalsGeneral) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *FundamentalsGeneral) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *FundamentalsGeneral) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetCurrencyName

`func (o *FundamentalsGeneral) GetCurrencyName() string`

GetCurrencyName returns the CurrencyName field if non-nil, zero value otherwise.

### GetCurrencyNameOk

`func (o *FundamentalsGeneral) GetCurrencyNameOk() (*string, bool)`

GetCurrencyNameOk returns a tuple with the CurrencyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyName

`func (o *FundamentalsGeneral) SetCurrencyName(v string)`

SetCurrencyName sets CurrencyName field to given value.

### HasCurrencyName

`func (o *FundamentalsGeneral) HasCurrencyName() bool`

HasCurrencyName returns a boolean if a field has been set.

### GetCurrencySymbol

`func (o *FundamentalsGeneral) GetCurrencySymbol() string`

GetCurrencySymbol returns the CurrencySymbol field if non-nil, zero value otherwise.

### GetCurrencySymbolOk

`func (o *FundamentalsGeneral) GetCurrencySymbolOk() (*string, bool)`

GetCurrencySymbolOk returns a tuple with the CurrencySymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencySymbol

`func (o *FundamentalsGeneral) SetCurrencySymbol(v string)`

SetCurrencySymbol sets CurrencySymbol field to given value.

### HasCurrencySymbol

`func (o *FundamentalsGeneral) HasCurrencySymbol() bool`

HasCurrencySymbol returns a boolean if a field has been set.

### GetCountryName

`func (o *FundamentalsGeneral) GetCountryName() string`

GetCountryName returns the CountryName field if non-nil, zero value otherwise.

### GetCountryNameOk

`func (o *FundamentalsGeneral) GetCountryNameOk() (*string, bool)`

GetCountryNameOk returns a tuple with the CountryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryName

`func (o *FundamentalsGeneral) SetCountryName(v string)`

SetCountryName sets CountryName field to given value.

### HasCountryName

`func (o *FundamentalsGeneral) HasCountryName() bool`

HasCountryName returns a boolean if a field has been set.

### GetCountryISO

`func (o *FundamentalsGeneral) GetCountryISO() string`

GetCountryISO returns the CountryISO field if non-nil, zero value otherwise.

### GetCountryISOOk

`func (o *FundamentalsGeneral) GetCountryISOOk() (*string, bool)`

GetCountryISOOk returns a tuple with the CountryISO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryISO

`func (o *FundamentalsGeneral) SetCountryISO(v string)`

SetCountryISO sets CountryISO field to given value.

### HasCountryISO

`func (o *FundamentalsGeneral) HasCountryISO() bool`

HasCountryISO returns a boolean if a field has been set.

### GetISIN

`func (o *FundamentalsGeneral) GetISIN() string`

GetISIN returns the ISIN field if non-nil, zero value otherwise.

### GetISINOk

`func (o *FundamentalsGeneral) GetISINOk() (*string, bool)`

GetISINOk returns a tuple with the ISIN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetISIN

`func (o *FundamentalsGeneral) SetISIN(v string)`

SetISIN sets ISIN field to given value.

### HasISIN

`func (o *FundamentalsGeneral) HasISIN() bool`

HasISIN returns a boolean if a field has been set.

### GetCUSIP

`func (o *FundamentalsGeneral) GetCUSIP() string`

GetCUSIP returns the CUSIP field if non-nil, zero value otherwise.

### GetCUSIPOk

`func (o *FundamentalsGeneral) GetCUSIPOk() (*string, bool)`

GetCUSIPOk returns a tuple with the CUSIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCUSIP

`func (o *FundamentalsGeneral) SetCUSIP(v string)`

SetCUSIP sets CUSIP field to given value.

### HasCUSIP

`func (o *FundamentalsGeneral) HasCUSIP() bool`

HasCUSIP returns a boolean if a field has been set.

### GetCIK

`func (o *FundamentalsGeneral) GetCIK() string`

GetCIK returns the CIK field if non-nil, zero value otherwise.

### GetCIKOk

`func (o *FundamentalsGeneral) GetCIKOk() (*string, bool)`

GetCIKOk returns a tuple with the CIK field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCIK

`func (o *FundamentalsGeneral) SetCIK(v string)`

SetCIK sets CIK field to given value.

### HasCIK

`func (o *FundamentalsGeneral) HasCIK() bool`

HasCIK returns a boolean if a field has been set.

### GetEmployerIdNumber

`func (o *FundamentalsGeneral) GetEmployerIdNumber() string`

GetEmployerIdNumber returns the EmployerIdNumber field if non-nil, zero value otherwise.

### GetEmployerIdNumberOk

`func (o *FundamentalsGeneral) GetEmployerIdNumberOk() (*string, bool)`

GetEmployerIdNumberOk returns a tuple with the EmployerIdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployerIdNumber

`func (o *FundamentalsGeneral) SetEmployerIdNumber(v string)`

SetEmployerIdNumber sets EmployerIdNumber field to given value.

### HasEmployerIdNumber

`func (o *FundamentalsGeneral) HasEmployerIdNumber() bool`

HasEmployerIdNumber returns a boolean if a field has been set.

### GetFiscalYearEnd

`func (o *FundamentalsGeneral) GetFiscalYearEnd() string`

GetFiscalYearEnd returns the FiscalYearEnd field if non-nil, zero value otherwise.

### GetFiscalYearEndOk

`func (o *FundamentalsGeneral) GetFiscalYearEndOk() (*string, bool)`

GetFiscalYearEndOk returns a tuple with the FiscalYearEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalYearEnd

`func (o *FundamentalsGeneral) SetFiscalYearEnd(v string)`

SetFiscalYearEnd sets FiscalYearEnd field to given value.

### HasFiscalYearEnd

`func (o *FundamentalsGeneral) HasFiscalYearEnd() bool`

HasFiscalYearEnd returns a boolean if a field has been set.

### GetIPODate

`func (o *FundamentalsGeneral) GetIPODate() string`

GetIPODate returns the IPODate field if non-nil, zero value otherwise.

### GetIPODateOk

`func (o *FundamentalsGeneral) GetIPODateOk() (*string, bool)`

GetIPODateOk returns a tuple with the IPODate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPODate

`func (o *FundamentalsGeneral) SetIPODate(v string)`

SetIPODate sets IPODate field to given value.

### HasIPODate

`func (o *FundamentalsGeneral) HasIPODate() bool`

HasIPODate returns a boolean if a field has been set.

### GetInternationalDomestic

`func (o *FundamentalsGeneral) GetInternationalDomestic() string`

GetInternationalDomestic returns the InternationalDomestic field if non-nil, zero value otherwise.

### GetInternationalDomesticOk

`func (o *FundamentalsGeneral) GetInternationalDomesticOk() (*string, bool)`

GetInternationalDomesticOk returns a tuple with the InternationalDomestic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternationalDomestic

`func (o *FundamentalsGeneral) SetInternationalDomestic(v string)`

SetInternationalDomestic sets InternationalDomestic field to given value.

### HasInternationalDomestic

`func (o *FundamentalsGeneral) HasInternationalDomestic() bool`

HasInternationalDomestic returns a boolean if a field has been set.

### GetSector

`func (o *FundamentalsGeneral) GetSector() string`

GetSector returns the Sector field if non-nil, zero value otherwise.

### GetSectorOk

`func (o *FundamentalsGeneral) GetSectorOk() (*string, bool)`

GetSectorOk returns a tuple with the Sector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSector

`func (o *FundamentalsGeneral) SetSector(v string)`

SetSector sets Sector field to given value.

### HasSector

`func (o *FundamentalsGeneral) HasSector() bool`

HasSector returns a boolean if a field has been set.

### GetIndustry

`func (o *FundamentalsGeneral) GetIndustry() string`

GetIndustry returns the Industry field if non-nil, zero value otherwise.

### GetIndustryOk

`func (o *FundamentalsGeneral) GetIndustryOk() (*string, bool)`

GetIndustryOk returns a tuple with the Industry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustry

`func (o *FundamentalsGeneral) SetIndustry(v string)`

SetIndustry sets Industry field to given value.

### HasIndustry

`func (o *FundamentalsGeneral) HasIndustry() bool`

HasIndustry returns a boolean if a field has been set.

### GetGicSector

`func (o *FundamentalsGeneral) GetGicSector() string`

GetGicSector returns the GicSector field if non-nil, zero value otherwise.

### GetGicSectorOk

`func (o *FundamentalsGeneral) GetGicSectorOk() (*string, bool)`

GetGicSectorOk returns a tuple with the GicSector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGicSector

`func (o *FundamentalsGeneral) SetGicSector(v string)`

SetGicSector sets GicSector field to given value.

### HasGicSector

`func (o *FundamentalsGeneral) HasGicSector() bool`

HasGicSector returns a boolean if a field has been set.

### GetGicGroup

`func (o *FundamentalsGeneral) GetGicGroup() string`

GetGicGroup returns the GicGroup field if non-nil, zero value otherwise.

### GetGicGroupOk

`func (o *FundamentalsGeneral) GetGicGroupOk() (*string, bool)`

GetGicGroupOk returns a tuple with the GicGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGicGroup

`func (o *FundamentalsGeneral) SetGicGroup(v string)`

SetGicGroup sets GicGroup field to given value.

### HasGicGroup

`func (o *FundamentalsGeneral) HasGicGroup() bool`

HasGicGroup returns a boolean if a field has been set.

### GetGicIndustry

`func (o *FundamentalsGeneral) GetGicIndustry() string`

GetGicIndustry returns the GicIndustry field if non-nil, zero value otherwise.

### GetGicIndustryOk

`func (o *FundamentalsGeneral) GetGicIndustryOk() (*string, bool)`

GetGicIndustryOk returns a tuple with the GicIndustry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGicIndustry

`func (o *FundamentalsGeneral) SetGicIndustry(v string)`

SetGicIndustry sets GicIndustry field to given value.

### HasGicIndustry

`func (o *FundamentalsGeneral) HasGicIndustry() bool`

HasGicIndustry returns a boolean if a field has been set.

### GetGicSubIndustry

`func (o *FundamentalsGeneral) GetGicSubIndustry() string`

GetGicSubIndustry returns the GicSubIndustry field if non-nil, zero value otherwise.

### GetGicSubIndustryOk

`func (o *FundamentalsGeneral) GetGicSubIndustryOk() (*string, bool)`

GetGicSubIndustryOk returns a tuple with the GicSubIndustry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGicSubIndustry

`func (o *FundamentalsGeneral) SetGicSubIndustry(v string)`

SetGicSubIndustry sets GicSubIndustry field to given value.

### HasGicSubIndustry

`func (o *FundamentalsGeneral) HasGicSubIndustry() bool`

HasGicSubIndustry returns a boolean if a field has been set.

### GetHomeCategory

`func (o *FundamentalsGeneral) GetHomeCategory() string`

GetHomeCategory returns the HomeCategory field if non-nil, zero value otherwise.

### GetHomeCategoryOk

`func (o *FundamentalsGeneral) GetHomeCategoryOk() (*string, bool)`

GetHomeCategoryOk returns a tuple with the HomeCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeCategory

`func (o *FundamentalsGeneral) SetHomeCategory(v string)`

SetHomeCategory sets HomeCategory field to given value.

### HasHomeCategory

`func (o *FundamentalsGeneral) HasHomeCategory() bool`

HasHomeCategory returns a boolean if a field has been set.

### GetIsDelisted

`func (o *FundamentalsGeneral) GetIsDelisted() bool`

GetIsDelisted returns the IsDelisted field if non-nil, zero value otherwise.

### GetIsDelistedOk

`func (o *FundamentalsGeneral) GetIsDelistedOk() (*bool, bool)`

GetIsDelistedOk returns a tuple with the IsDelisted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDelisted

`func (o *FundamentalsGeneral) SetIsDelisted(v bool)`

SetIsDelisted sets IsDelisted field to given value.

### HasIsDelisted

`func (o *FundamentalsGeneral) HasIsDelisted() bool`

HasIsDelisted returns a boolean if a field has been set.

### GetDescription

`func (o *FundamentalsGeneral) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FundamentalsGeneral) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FundamentalsGeneral) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FundamentalsGeneral) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAddress

`func (o *FundamentalsGeneral) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *FundamentalsGeneral) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *FundamentalsGeneral) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *FundamentalsGeneral) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAddressData

`func (o *FundamentalsGeneral) GetAddressData() FundamentalsGeneralAddressData`

GetAddressData returns the AddressData field if non-nil, zero value otherwise.

### GetAddressDataOk

`func (o *FundamentalsGeneral) GetAddressDataOk() (*FundamentalsGeneralAddressData, bool)`

GetAddressDataOk returns a tuple with the AddressData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressData

`func (o *FundamentalsGeneral) SetAddressData(v FundamentalsGeneralAddressData)`

SetAddressData sets AddressData field to given value.

### HasAddressData

`func (o *FundamentalsGeneral) HasAddressData() bool`

HasAddressData returns a boolean if a field has been set.

### GetListings

`func (o *FundamentalsGeneral) GetListings() map[string]FundamentalsGeneralListings`

GetListings returns the Listings field if non-nil, zero value otherwise.

### GetListingsOk

`func (o *FundamentalsGeneral) GetListingsOk() (*map[string]FundamentalsGeneralListings, bool)`

GetListingsOk returns a tuple with the Listings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListings

`func (o *FundamentalsGeneral) SetListings(v map[string]FundamentalsGeneralListings)`

SetListings sets Listings field to given value.

### HasListings

`func (o *FundamentalsGeneral) HasListings() bool`

HasListings returns a boolean if a field has been set.

### GetOfficers

`func (o *FundamentalsGeneral) GetOfficers() map[string]FundamentalsGeneralOfficers`

GetOfficers returns the Officers field if non-nil, zero value otherwise.

### GetOfficersOk

`func (o *FundamentalsGeneral) GetOfficersOk() (*map[string]FundamentalsGeneralOfficers, bool)`

GetOfficersOk returns a tuple with the Officers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficers

`func (o *FundamentalsGeneral) SetOfficers(v map[string]FundamentalsGeneralOfficers)`

SetOfficers sets Officers field to given value.

### HasOfficers

`func (o *FundamentalsGeneral) HasOfficers() bool`

HasOfficers returns a boolean if a field has been set.

### GetPhone

`func (o *FundamentalsGeneral) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *FundamentalsGeneral) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *FundamentalsGeneral) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *FundamentalsGeneral) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetWebURL

`func (o *FundamentalsGeneral) GetWebURL() string`

GetWebURL returns the WebURL field if non-nil, zero value otherwise.

### GetWebURLOk

`func (o *FundamentalsGeneral) GetWebURLOk() (*string, bool)`

GetWebURLOk returns a tuple with the WebURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebURL

`func (o *FundamentalsGeneral) SetWebURL(v string)`

SetWebURL sets WebURL field to given value.

### HasWebURL

`func (o *FundamentalsGeneral) HasWebURL() bool`

HasWebURL returns a boolean if a field has been set.

### GetLogoURL

`func (o *FundamentalsGeneral) GetLogoURL() string`

GetLogoURL returns the LogoURL field if non-nil, zero value otherwise.

### GetLogoURLOk

`func (o *FundamentalsGeneral) GetLogoURLOk() (*string, bool)`

GetLogoURLOk returns a tuple with the LogoURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoURL

`func (o *FundamentalsGeneral) SetLogoURL(v string)`

SetLogoURL sets LogoURL field to given value.

### HasLogoURL

`func (o *FundamentalsGeneral) HasLogoURL() bool`

HasLogoURL returns a boolean if a field has been set.

### GetFullTimeEmployees

`func (o *FundamentalsGeneral) GetFullTimeEmployees() int64`

GetFullTimeEmployees returns the FullTimeEmployees field if non-nil, zero value otherwise.

### GetFullTimeEmployeesOk

`func (o *FundamentalsGeneral) GetFullTimeEmployeesOk() (*int64, bool)`

GetFullTimeEmployeesOk returns a tuple with the FullTimeEmployees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullTimeEmployees

`func (o *FundamentalsGeneral) SetFullTimeEmployees(v int64)`

SetFullTimeEmployees sets FullTimeEmployees field to given value.

### HasFullTimeEmployees

`func (o *FundamentalsGeneral) HasFullTimeEmployees() bool`

HasFullTimeEmployees returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FundamentalsGeneral) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FundamentalsGeneral) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FundamentalsGeneral) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FundamentalsGeneral) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


