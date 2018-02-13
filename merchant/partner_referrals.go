package merchant

import (
	"encoding/json"
	"errors"
	"net/url"
	"strings"
	"time"
)

type CustomerTypeData string

const (
	CustomerTypeConsumer CustomerTypeData = "CONSUMER"
	CustomerTypeMerchant CustomerTypeData = "MERCHANT"
)

type NameOfAPartyData struct {
	Prefix            string `json:"prefix,omitempty"`
	GivenName         string `json:"given_name,omitempty"`
	Surname           string `json:"surname,omitempty"`
	MiddleName        string `json:"middle_name,omitempty"`
	Suffix            string `json:"suffix,omitempty"`
	AlternameFullName string `json:"alternate_full_name,omitempty"`
}

type PhoneDetailsData struct {
	CountryCode     string `json:"country_code,omitempty"`
	NationalNumber  string `json:"national_number,omitempty"`
	ExtensionNumber string `json:"extension_number,omitempty"`
}

type PhoneTypeData string

const (
	PhoneTypeFax    PhoneTypeData = "FAX"
	PhoneTypeHome   PhoneTypeData = "HOME"
	PhoneTypeMobile PhoneTypeData = "MOBILE"
	PhoneTypeOther  PhoneTypeData = "OTHER"
	PhoneTypePager  PhoneTypeData = "PAGER"
)

type OnboardingCommonUserPhoneData struct {
	PhoneNumberDetails *PhoneDetailsData `json:"phone_number_details,omitempty"`
	PhoneType          PhoneTypeData     `json:"phone_type,omitempty"`
}

type SimplePostalAddressData struct {
	Line1       string `json:"line1,omitempty"`
	Line2       string `json:"line2,omitempty"`
	City        string `json:"city,omitempty"`
	State       string `json:"state,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
	PostalCode  string `json:"postal_code,omitempty"`
}

type EventTypeData string

const (
	EventTypeBirth         EventTypeData = "BIRTH"
	EventTypeEstablished   EventTypeData = "ESTABLISHED"
	EventTypeIncorporation EventTypeData = "INCORPORATION"
	EventTypeOperation     EventTypeData = "OPERATION"
)

type DateData struct {
	EventType EventTypeData `json:"event_type,omitempty"`
	EventDate time.Time     `json:"event_date,omitempty"`
}

func (d *DateData) MarshalJSON() ([]byte, error) {
	data := struct {
		EventType EventTypeData `json:"event_type,omitempty"`
		EventDate string        `json:"event_date,omitempty"`
	}{
		EventType: d.EventType,
		EventDate: d.EventDate.Format(time.RFC3339Nano),
	}
	if d.EventDate.IsZero() {
		data.EventDate = ""
	}
	return json.Marshal(&data)
}

func (d *DateData) UnmarshalJSON(b []byte) error {
	temp := struct {
		EventType EventTypeData `json:"event_type,omitempty"`
		EventDate string        `json:"event_date,omitempty"`
	}{}
	err := json.Unmarshal(b, &temp)
	if err != nil {
		return err
	}
	d.EventType = temp.EventType
	if temp.EventDate != "" {
		d.EventDate, err = time.Parse(time.RFC3339Nano, temp.EventDate)
		if err != nil {
			return err
		}
	} else {
		d.EventDate = time.Time{}
	}
	return nil
}

type IdentityDocumentType string

const (
	IdentityTypeSocialSecurityNumber           IdentityDocumentType = "SOCIAL_SECURITY_NUMBER"
	IdentityTypeEmploymentIdentificationNumber IdentityDocumentType = "EMPLOYMENT_IDENTIFICATION_NUMBER"
	IdentityTypeTaxIdentificationNumber        IdentityDocumentType = "TAX_IDENTIFICATION_NUMBER"
	IdentityTypePassportNumber                 IdentityDocumentType = "PASSPORT_NUMBER"
	IdentityTypePensionFundID                  IdentityDocumentType = "PENSION_FUND_ID"
	IdentityTypeMedicalInsuranceID             IdentityDocumentType = "MEDICAL_INSURANCE_ID"
	IdentityTypeCNPJ                           IdentityDocumentType = "CNPJ"
	IdentityTypeCPF                            IdentityDocumentType = "CPF"
)

type IdentityDocumentData struct {
	Type              IdentityDocumentType `json:"type,omitempty"`
	Value             string               `json:"value,omitempty"`
	PartialValue      bool                 `json:"partial_value"`
	IssuerCountryCode string               `json:"issuer_country_code,omitempty"`
}

type RelationData string

const (
	RelationMother RelationData = "MOTHER"
)

type MerchantRelationData struct {
	Name                     *NameOfAPartyData `json:"name,omitempty"`
	Relation                 RelationData      `json:"relation,omitempty"`
	CountryCodeOfNationality string            `json:"country_code_of_nationality,omitempty"`
}

type PersonDetailsData struct {
	EmailAddress              string                          `json:"email_address,omitempty"`
	Name                      *NameOfAPartyData               `json:"name,omitempty"`
	PhoneContacts             []OnboardingCommonUserPhoneData `json:"phone_contacts,omitempty"`
	HomeAddress               *SimplePostalAddressData        `json:"home_address,omitempty"`
	DateOfBirth               *DateData                       `json:"date_of_birth,omitempty"`
	NationalityCountryCode    string                          `json:"nationality_country_code,omitempty"`
	IdentityDocuments         []IdentityDocumentData          `json:"identity_documents,omitempty"`
	AccountOwnerRelationships []MerchantRelationData          `json:"account_owner_relationships,omitempty"`
}

type BusinessTypeData string

const (
	BusinessTypeIndividual                  BusinessTypeData = "INDIVIDUAL"
	BusinessTypeProprietorship              BusinessTypeData = "PROPRIETORSHIP"
	BusinessTypePartnership                 BusinessTypeData = "PARTNERSHIP"
	BusinessTypeCorporation                 BusinessTypeData = "CORPORATION"
	BusinessTypeNonprofit                   BusinessTypeData = "NONPROFIT"
	BusinessTypeGovernment                  BusinessTypeData = "GOVERNMENT"
	BusinessTypePublicCompany               BusinessTypeData = "PUBLIC_COMPANY"
	BusinessTypeRegisteredCooperative       BusinessTypeData = "REGISTERED_COOPERATIVE"
	BusinessTypeProprietoryCompany          BusinessTypeData = "PROPRIETORY_COMPANY"
	BusinessTypeAssociation                 BusinessTypeData = "ASSOCIATION"
	BusinessTypePrivateCorporation          BusinessTypeData = "PRIVATE_CORPORATION"
	BusinessTypeLimitedPartnership          BusinessTypeData = "LIMITED_PARTNERSHIP"
	BusinessTypeLimitedLiabilityProprietors BusinessTypeData = "LIMITED_LIABILITY_PROPRIETORS"
	BusinessTypeLimitedLiabilityPartnership BusinessTypeData = "LIMITED_LIABILITY_PARTNERSHIP"
	BusinessTypePublicCorporation           BusinessTypeData = "PUBLIC_CORPORATION"
	BusinessTypeOtherCorporateBody          BusinessTypeData = "OTHER_CORPORATE_BODY"
)

type BusinessCategory interface {
	getCategoryCode() string
	getSubCategoryCode() string
}

type SubCategoryArts string

const (
	ArtSubCategoryAntiques       SubCategoryArts = "2000"
	ArtSubCategoryArtCraft       SubCategoryArts = "2001"
	ArtSubCategoryArtDealers     SubCategoryArts = "2002"
	ArtSubCategoryCameraSupplies SubCategoryArts = "2003"
	ArtSubCategoryDigitalArt     SubCategoryArts = "2004"
	ArtSubCategoryMemorabilia    SubCategoryArts = "2005"
	ArtSubCategoryMusic          SubCategoryArts = "2006"
	ArtSubCategorySewing         SubCategoryArts = "2007"
	ArtSubCategoryStampCoin      SubCategoryArts = "2008"
	ArtSubCategoryStationary     SubCategoryArts = "2009"
	ArtSubCategoryVintage        SubCategoryArts = "2010"
)

type CategoryArts struct {
	SubCat SubCategoryArts
}

func (c CategoryArts) getCategoryCode() string {
	return "1000"
}

func (c CategoryArts) getSubCategoryCode() string {
	return string(c.SubCat)
}

type SubCategoryBaby string

const (
	BabySubCategoryClothing      SubCategoryBaby = "2011"
	BabySubCategoryFurniture     SubCategoryBaby = "2012"
	BabySubCategoryProductsOther SubCategoryBaby = "2013"
	BabySubCategorySafety        SubCategoryBaby = "2014"
)

type CategoryBaby struct {
	SubCat SubCategoryBaby
}

func (c CategoryBaby) getCategoryCode() string {
	return "1001"
}

func (c CategoryBaby) getSubCategoryCode() string {
	return string(c.SubCat)
}

type SubCategoryBeauty string

const (
	BeautySubCategoryBathBody   SubCategoryBeauty = "2015"
	BeautySubCategoryFragrances SubCategoryBeauty = "2016"
	BeautySubCategoryMakeup     SubCategoryBeauty = "2017"
)

type CategoryBeauty struct {
	SubCat SubCategoryBeauty
}

func (c CategoryBeauty) getCategoryCode() string {
	return "1002"
}

func (c CategoryBeauty) getSubCategoryCode() string {
	return string(c.SubCat)
}

type SubCategoryBooks string

const (
	BookSubCategoryAudio             SubCategoryBooks = "2018"
	BookSubCategoryDigital           SubCategoryBooks = "2019"
	BookSubCategoryEducational       SubCategoryBooks = "2020"
	BookSubCategoryFictionNonFiction SubCategoryBooks = "2021"
	BookSubCategoryMagazines         SubCategoryBooks = "2022"
	BookSubCategoryPublishing        SubCategoryBooks = "2023"
	BookSubCategoryRareUsed          SubCategoryBooks = "2024"
)

type CategoryBooks struct {
	SubCat SubCategoryBooks
}

func (c CategoryBooks) getCategoryCode() string {
	return "1003"
}

func (c CategoryBooks) getSubCategoryCode() string {
	return string(c.SubCat)
}

type SubCategoryBusiness string

const (
	BusinessSubCategoryAccounting          SubCategoryBusiness = "2025"
	BusinessSubCategoryAdvertising         SubCategoryBusiness = "2026"
	BusinessSubCategoryAgricultural        SubCategoryBusiness = "2027"
	BusinessSubCategoryArchitectural       SubCategoryBusiness = "2028"
	BusinessSubCategoryChemicals           SubCategoryBusiness = "2029"
	BusinessSubCategoryPhotography         SubCategoryBusiness = "2030"
	BusinessSubCategoryConstruction        SubCategoryBusiness = "2031"
	BusinessSubCategoryConsulting          SubCategoryBusiness = "2032"
	BusinessSubCategoryEducational         SubCategoryBusiness = "2033"
	BusinessSubCategoryEquipmentRental     SubCategoryBusiness = "2034"
	BusinessSubCategoryEquipmentRepair     SubCategoryBusiness = "2035"
	BusinessSubCategoryHiring              SubCategoryBusiness = "2036"
	BusinessSubCategoryIndustrial          SubCategoryBusiness = "2037"
	BusinessSubCategoryMailingLists        SubCategoryBusiness = "2038"
	BusinessSubCategoryMarketing           SubCategoryBusiness = "2039"
	BusinessSubCategoryMultiLevelMarketing SubCategoryBusiness = "2040"
	BusinessSubCategoryOfficeFurniture     SubCategoryBusiness = "2041"
	BusinessSubCategoryOfficeSupplies      SubCategoryBusiness = "2042"
	BusinessSubCategoryPublishing          SubCategoryBusiness = "2043"
	BusinessSubCategoryCopyReproduction    SubCategoryBusiness = "2044"
	BusinessSubCategoryShipping            SubCategoryBusiness = "2045"
	BusinessSubCategorySecretarial         SubCategoryBusiness = "2046"
	BusinessSubCategoryWholesale           SubCategoryBusiness = "2047"
)

type CategoryBusiness struct {
	SubCat SubCategoryBusiness
}

func (c CategoryBusiness) getCategoryCode() string {
	return "1004"
}

func (c CategoryBusiness) getSubCategoryCode() string {
	return string(c.SubCat)
}

type SubCategoryClothing string

const (
	ClothingSubCategoryChildrens       SubCategoryClothing = "2048"
	ClothingSubCategoryMens            SubCategoryClothing = "2049"
	ClothingSubCategoryWomens          SubCategoryClothing = "2050"
	ClothingSubCategoryShoes           SubCategoryClothing = "2051"
	ClothingSubCategoryMiltary         SubCategoryClothing = "2052"
	ClothingSubCategoryAccessories     SubCategoryClothing = "2053"
	ClothingSubCategoryRetailJewelry   SubCategoryClothing = "2054"
	ClothingSubCategoryWholesaleStones SubCategoryClothing = "2055"
	ClothingSubCategoryFashion         SubCategoryClothing = "2056"
)

type CategoryClothing struct {
	SubCat SubCategoryClothing
}

func (c CategoryClothing) getCategoryCode() string {
	return "1005"
}

func (c CategoryClothing) getSubCategoryCode() string {
	return string(c.SubCat)
}

type SubCategoryComputers string

const (
	ComputersSubCategoryDataProcessing         SubCategoryComputers = "2057"
	ComputersSubCategoryDesktopLaptopNotebooks SubCategoryComputers = "2058"
	ComputersSubCategoryDigitalContent         SubCategoryComputers = "2059"
	ComputersSubCategoryECommerce              SubCategoryComputers = "2060"
	ComputersSubCategoryMaintenance            SubCategoryComputers = "2061"
	ComputersSubCategoryMonitors               SubCategoryComputers = "2062"
	ComputersSubCategoryNetworking             SubCategoryComputers = "2063"
	ComputersSubCategoryOnlineGaming           SubCategoryComputers = "2064"
	ComputersSubCategoryParts                  SubCategoryComputers = "2065"
	ComputersSubCategoryPeripherals            SubCategoryComputers = "2066"
	ComputersSubCategorySoftware               SubCategoryComputers = "2067"
	ComputersSubCategoryTraining               SubCategoryComputers = "2068"
	ComputersSubCategoryWebHosting             SubCategoryComputers = "2069"
)

type CategoryComputers struct {
	SubCat SubCategoryComputers
}

func (c CategoryComputers) getCategoryCode() string {
	return "1006"
}

func (c CategoryComputers) getSubCategoryCode() string {
	return string(c.SubCat)
}

type SubCategoryEducation string

const (
	EducationSubCategoryBusiness   SubCategoryEducation = "2070"
	EducationSubCategoryDaycare    SubCategoryEducation = "2071"
	EducationSubCategoryColleges   SubCategoryEducation = "2072"
	EducationSubCategoryDance      SubCategoryEducation = "2073"
	EducationSubCategoryElementary SubCategoryEducation = "2074"
	EducationSubCategoryVocational SubCategoryEducation = "2075"
)

type CategoryEducation struct {
	SubCat SubCategoryEducation
}

func (c CategoryEducation) getCategoryCode() string {
	return "1007"
}

func (c CategoryEducation) getSubCategoryCode() string {
	return string(c.SubCat)
}

type SubCategoryElectronics string

const (
	ElectronicsSubCategoryCameras                    SubCategoryElectronics = "2076" // No 2077
	ElectronicsSubCategoryCellPhones                 SubCategoryElectronics = "2078"
	ElectronicsSubCategoryAccessories                SubCategoryElectronics = "2079"
	ElectronicsSubCategoryHomeAudio                  SubCategoryElectronics = "2080"
	ElectronicsSubCategoryHomeElectronics            SubCategoryElectronics = "2081"
	ElectronicsSubCategorySecurity                   SubCategoryElectronics = "2082"
	ElectronicsSubCategoryTelecommunicationEquipment SubCategoryElectronics = "2083"
	ElectronicsSubCategoryTelecommunicationServices  SubCategoryElectronics = "2084"
	ElectronicsSubCategoryTelephoneCards             SubCategoryElectronics = "2085"
)

type CategoryElectronics struct {
	SubCat SubCategoryElectronics
}

func (c CategoryElectronics) getCategoryCode() string {
	return "1008"
}

func (c CategoryElectronics) getSubCategoryCode() string {
	return string(c.SubCat)
}

type SubCategoryEntertainment string

const (
	EntertainmentSubCategoryMemorabilia         SubCategoryEntertainment = "2086"
	EntertainmentSubCategoryMovieTickets        SubCategoryEntertainment = "2087"
	EntertainmentSubCategoryMoviesDVDs          SubCategoryEntertainment = "2088"
	EntertainmentSubCategoryMusicCDs            SubCategoryEntertainment = "2089"
	EntertainmentSubCategoryCableTV             SubCategoryEntertainment = "2090"
	EntertainmentSubCategoryAdultDigitalContent SubCategoryEntertainment = "2091"
	EntertainmentSubCategoryConcertTickets      SubCategoryEntertainment = "2092"
	EntertainmentSubCategoryTheaterTickets      SubCategoryEntertainment = "2093"
	EntertainmentSubCategoryToysGames           SubCategoryEntertainment = "2094"
	EntertainmentSubCategorySlotGames           SubCategoryEntertainment = "2095"
	EntertainmentSubCategoryDigitalContent      SubCategoryEntertainment = "2096"
	EntertainmentSubCategoryEntertainers        SubCategoryEntertainment = "2097"
	EntertainmentSubCategoryGambling            SubCategoryEntertainment = "2098"
	EntertainmentSubCategoryOnlineGames         SubCategoryEntertainment = "2099"
	EntertainmentSubCategoryVideoGames          SubCategoryEntertainment = "2100"
)

type CategoryEntertainment struct {
	SubCat SubCategoryEntertainment
}

func (c CategoryEntertainment) getCategoryCode() string {
	return "1009"
}

func (c CategoryEntertainment) getSubCategoryCode() string {
	return string(c.SubCat)
}

type SubCategoryFinancial string

const (
	FinancialSubCategoryAccounting              SubCategoryFinancial = "2101"
	FinancialSubCategoryCollectionAgency        SubCategoryFinancial = "2102"
	FinancialSubCategoryCommodities             SubCategoryFinancial = "2103"
	FinancialSubCategoryConsumerCreditReporting SubCategoryFinancial = "2104"
	FinancialSubCategoryDebtCounseling          SubCategoryFinancial = "2105"
	FinancialSubCategoryCreditUnion             SubCategoryFinancial = "2106"
	FinancialSubCategoryCurrencyDealerExchange  SubCategoryFinancial = "2107"
	FinancialSubCategoryEscrow                  SubCategoryFinancial = "2108"
	FinancialSubCategoryFinance                 SubCategoryFinancial = "2109"
	FinancialSubCategoryFinancialAdvice         SubCategoryFinancial = "2110"
	FinancialSubCategoryInsuranceAutoHome       SubCategoryFinancial = "2111"
	FinancialSubCategoryInsuranceLifeAnnuity    SubCategoryFinancial = "2112"
	FinancialSubCategoryInvestmentsGeneral      SubCategoryFinancial = "2113"
	FinancialSubCategoryMoneyService            SubCategoryFinancial = "2114"
	FinancialSubCategoryMortgageBrokers         SubCategoryFinancial = "2115"
	FinancialSubCategoryOnlineGamingCurrency    SubCategoryFinancial = "2116"
	FinancialSubCategoryPaycheckLender          SubCategoryFinancial = "2117"
	FinancialSubCategoryPrepaidCards            SubCategoryFinancial = "2118"
	FinancialSubCategoryRealEstate              SubCategoryFinancial = "2119"
	FinancialSubCategoryRemittance              SubCategoryFinancial = "2120"
	FinancialSubCategoryRentalProperty          SubCategoryFinancial = "2121"
	FinancialSubCategorySecurityBrokers         SubCategoryFinancial = "2122"
	FinancialSubCategoryWireTransfer            SubCategoryFinancial = "2123"
)

type CategoryFinancial struct {
	SubCat SubCategoryFinancial
}

func (c CategoryFinancial) getCategoryCode() string {
	return "1010"
}

func (c CategoryFinancial) getSubCategoryCode() string {
	return string(c.SubCat)
}

type SubCategoryFood string

const (
	FoodSubCategoryAlcoholicBeverages  SubCategoryFood = "2124"
	FoodSubCategoryCatering            SubCategoryFood = "2125"
	FoodSubCategoryCoffeeTea           SubCategoryFood = "2126"
	FoodSubCategoryGourmet             SubCategoryFood = "2127"
	FoodSubCategorySpecialty           SubCategoryFood = "2128"
	FoodSubCategoryRestaurant          SubCategoryFood = "2129"
	FoodSubCategoryTobacco             SubCategoryFood = "2130"
	FoodSubCategoryVitaminsSupplements SubCategoryFood = "2131"
)

type CategoryFood struct {
	SubCat SubCategoryFood
}

func (c CategoryFood) getCategoryCode() string {
	return "1011"
}

func (c CategoryFood) getSubCategoryCode() string {
	return string(c.SubCat)
}

type SubCategoryGifts string

const (
	GiftsSubCategoryFlorist       SubCategoryGifts = "2132"
	GiftsSubCategoryGiftShop      SubCategoryGifts = "2133"
	GiftsSubCategoryGourmetFood   SubCategoryGifts = "2134"
	GiftsSubCategoryNurseryPlants SubCategoryGifts = "2135"
	GiftsSubCategoryPartySupplies SubCategoryGifts = "2136"
)

type CategoryGifts struct {
	SubCat SubCategoryGifts
}

func (c CategoryGifts) getCategoryCode() string {
	return "1012"
}

func (c CategoryGifts) getSubCategoryCode() string {
	return string(c.SubCat)
}

type SubCategoryGovernment string

const (
	GovernmentSubCategoryServices SubCategoryGovernment = "2137"
)

type CategoryGovernment struct {
	SubCat SubCategoryGovernment
}

func (c CategoryGovernment) getCategoryCode() string {
	return "1013"
}

func (c CategoryGovernment) getSubCategoryCode() string {
	return string(c.SubCat)
}

type SubCategoryHealth string

const (
	HealthSubCategoryDrugstoreExcludingPerscription SubCategoryHealth = "2138"
	HealthSubCategoryDrugstoreIncludingPerscription SubCategoryHealth = "2139"
	HealthSubCategoryDental                         SubCategoryHealth = "2140"
	HealthSubCategoryMedicalCare                    SubCategoryHealth = "2141"
	HealthSubCategoryMedicalEquipment               SubCategoryHealth = "2142"
	HealthSubCategoryVision                         SubCategoryHealth = "2143"
	HealthSubCategoryVitaminsSupplements            SubCategoryHealth = "2144"
)

type CategoryHealth struct {
	SubCat SubCategoryHealth
}

func (c CategoryHealth) getCategoryCode() string {
	return "1014"
}

func (c CategoryHealth) getSubCategoryCode() string {
	return string(c.SubCat)
}

type SubCategoryHome string

const (
	HomeSubCategoryAntiques            SubCategoryHome = "2145"
	HomeSubCategoryAppliances          SubCategoryHome = "2146"
	HomeSubCategoryArtDealers          SubCategoryHome = "2147"
	HomeSubCategoryBedBath             SubCategoryHome = "2148"
	HomeSubCategoryConstruction        SubCategoryHome = "2149"
	HomeSubCategoryDrapery             SubCategoryHome = "2150"
	HomeSubCategoryExterminating       SubCategoryHome = "2151"
	HomeSubCategoryFireplace           SubCategoryHome = "2152"
	HomeSubCategoryFurniture           SubCategoryHome = "2153"
	HomeSubCategoryGarden              SubCategoryHome = "2154"
	HomeSubCategoryGlassPaintWallpaper SubCategoryHome = "2155"
	HomeSubCategoryHardwareTools       SubCategoryHome = "2156"
	HomeSubCategoryHomeDecor           SubCategoryHome = "2157"
	HomeSubCategoryHousewares          SubCategoryHome = "2158"
	HomeSubCategoryKitchenware         SubCategoryHome = "2159"
	HomeSubCategoryLandscaping         SubCategoryHome = "2160"
	HomeSubCategoryRugsCarpets         SubCategoryHome = "2161"
	HomeSubCategorySecurity            SubCategoryHome = "2162"
	HomeSubCategorySwimmingPools       SubCategoryHome = "2163"
)

type CategoryHome struct {
	SubCat SubCategoryHome
}

func (c CategoryHome) getCategoryCode() string {
	return "1015"
}

func (c CategoryHome) getSubCategoryCode() string {
	return string(c.SubCat)
}

type SubCategoryNonprofit string

const (
	NonprofitSubCategoryCharity     SubCategoryNonprofit = "2164"
	NonprofitSubCategoryPolitical   SubCategoryNonprofit = "2165"
	NonprofitSubCategoryReligious   SubCategoryNonprofit = "2166"
	NonprofitSubCategoryOther       SubCategoryNonprofit = "2167"
	NonprofitSubCategoryPersonal    SubCategoryNonprofit = "2168"
	NonprofitSubCategoryEducational SubCategoryNonprofit = "2169"
)

type CategoryNonprofit struct {
	SubCat SubCategoryNonprofit
}

func (c CategoryNonprofit) getCategoryCode() string {
	return "1016"
}

func (c CategoryNonprofit) getSubCategoryCode() string {
	return string(c.SubCat)
}

type SubCategoryPets string

const (
	PetsSubCategoryMedicationSupplements SubCategoryPets = "2171" // No 2170
	PetsSubCategoryShopsFoodSupplies     SubCategoryPets = "2172"
	PetsSubCategorySpecialty             SubCategoryPets = "2173"
	PetsSubCategoryVeterinary            SubCategoryPets = "2174"
)

type CategoryPets struct {
	SubCat SubCategoryPets
}

func (c CategoryPets) getCategoryCode() string {
	return "1017"
}

func (c CategoryPets) getSubCategoryCode() string {
	return string(c.SubCat)
}

type SubCategoryReligion string

const (
	ReligionSubCategoryMembershipServices SubCategoryReligion = "2175"
	ReligionSubCategoryMerchandise        SubCategoryReligion = "2176"
	ReligionSubCategoryServicesOther      SubCategoryReligion = "2177"
)

type CategoryReligion struct {
	SubCat SubCategoryReligion
}

func (c CategoryReligion) getCategoryCode() string {
	return "1018"
}

func (c CategoryReligion) getSubCategoryCode() string {
	return string(c.SubCat)
}

type SubCategoryRetail string

const (
	RetailSubCategoryChemicals       SubCategoryRetail = "2178"
	RetailSubCategoryDepartment      SubCategoryRetail = "2179"
	RetailSubCategoryDiscount        SubCategoryRetail = "2180"
	RetailSubCategoryDurableGoods    SubCategoryRetail = "2181"
	RetailSubCategoryNonDurableGoods SubCategoryRetail = "2182"
	RetailSubCategoryUsedSecondhand  SubCategoryRetail = "2183"
	RetailSubCategoryVariety         SubCategoryRetail = "2184"
)

type CategoryRetail struct {
	SubCat SubCategoryRetail
}

func (c CategoryRetail) getCategoryCode() string {
	return "1019"
}

func (c CategoryRetail) getSubCategoryCode() string {
	return string(c.SubCat)
}

type SubCategoryServicesOther string

const (
	OtherSubCategoryAdvertising                 SubCategoryServicesOther = "2185"
	OtherSubCategoryShoppingServices            SubCategoryServicesOther = "2186"
	OtherSubCategoryCareerServices              SubCategoryServicesOther = "2187"
	OtherSubCategoryCarpentry                   SubCategoryServicesOther = "2188"
	OtherSubCategoryChildCare                   SubCategoryServicesOther = "2189"
	OtherSubCategoryCleaningMaintenance         SubCategoryServicesOther = "2190"
	OtherSubCategoryCommercialPhotography       SubCategoryServicesOther = "2191"
	OtherSubCategoryComputerDataProcessing      SubCategoryServicesOther = "2192"
	OtherSubCategoryComputerNetwork             SubCategoryServicesOther = "2193"
	OtherSubCategoryConsulting                  SubCategoryServicesOther = "2194"
	OtherSubCategoryCounseling                  SubCategoryServicesOther = "2195"
	OtherSubCategoryCourier                     SubCategoryServicesOther = "2196"
	OtherSubCategoryDental                      SubCategoryServicesOther = "2197"
	OtherSubCategoryECommerce                   SubCategoryServicesOther = "2198"
	OtherSubCategoryElectricalRepair            SubCategoryServicesOther = "2199"
	OtherSubCategoryEntertainment               SubCategoryServicesOther = "2200"
	OtherSubCategoryEquipmentRental             SubCategoryServicesOther = "2201"
	OtherSubCategoryEventPlanning               SubCategoryServicesOther = "2202"
	OtherSubCategoryGambling                    SubCategoryServicesOther = "2203"
	OtherSubCategoryGeneralContractors          SubCategoryServicesOther = "2204"
	OtherSubCategoryGraphicDesign               SubCategoryServicesOther = "2205"
	OtherSubCategoryHealthSpas                  SubCategoryServicesOther = "2206"
	OtherSubCategoryIDPassport                  SubCategoryServicesOther = "2207"
	OtherSubCategoryImportExport                SubCategoryServicesOther = "2208"
	OtherSubCategoryInformationRetrieval        SubCategoryServicesOther = "2209"
	OtherSubCategoryInsuranceAutoHome           SubCategoryServicesOther = "2210"
	OtherSubCategoryInsuranceLifeAnnuity        SubCategoryServicesOther = "2211"
	OtherSubCategoryLandscaping                 SubCategoryServicesOther = "2212"
	OtherSubCategoryLegalServices               SubCategoryServicesOther = "2213"
	OtherSubCategoryLocalDelivery               SubCategoryServicesOther = "2214"
	OtherSubCategoryLottery                     SubCategoryServicesOther = "2215"
	OtherSubCategoryMedicalCare                 SubCategoryServicesOther = "2216"
	OtherSubCategoryMembershipClubs             SubCategoryServicesOther = "2217"
	OtherSubCategoryMiscPublishing              SubCategoryServicesOther = "2218"
	OtherSubCategoryMovingStorage               SubCategoryServicesOther = "2219"
	OtherSubCategoryOnlineDating                SubCategoryServicesOther = "2220"
	OtherSubCategoryPhotofinishing              SubCategoryServicesOther = "2221"
	OtherSubCategoryPhotographicPortraits       SubCategoryServicesOther = "2222"
	OtherSubCategoryProtectiveServices          SubCategoryServicesOther = "2223"
	OtherSubCategoryQuickCopyReproduction       SubCategoryServicesOther = "2224"
	OtherSubCategoryRadioTelevisionRepair       SubCategoryServicesOther = "2225"
	OtherSubCategoryRealEstate                  SubCategoryServicesOther = "2226"
	OtherSubCategoryRentalProperty              SubCategoryServicesOther = "2227"
	OtherSubCategoryReupholsteryFurnitureRepair SubCategoryServicesOther = "2228"
	OtherSubCategoryServicesOther               SubCategoryServicesOther = "2229"
	OtherSubCategoryShipping                    SubCategoryServicesOther = "2230"
	OtherSubCategorySwimmingPool                SubCategoryServicesOther = "2231"
	OtherSubCategoryTailors                     SubCategoryServicesOther = "2232"
	OtherSubCategoryTelecommunicationService    SubCategoryServicesOther = "2233"
	OtherSubCategoryUtilities                   SubCategoryServicesOther = "2234"
	OtherSubCategoryVisionCare                  SubCategoryServicesOther = "2235"
	OtherSubCategoryWatchClockJewelryRepair     SubCategoryServicesOther = "2236"
)

type CategoryServicesOther struct {
	SubCat SubCategoryServicesOther
}

func (c CategoryServicesOther) getCategoryCode() string {
	return "1020"
}

func (c CategoryServicesOther) getSubCategoryCode() string {
	return string(c.SubCat)
}

type SubCategorySports string

const (
	SportsSubCategoryAthleticShoes      SubCategorySports = "2237"
	SportsSubCategoryBicycleShop        SubCategorySports = "2238"
	SportsSubCategoryBoating            SubCategorySports = "2239"
	SportsSubCategoryCamping            SubCategorySports = "2240"
	SportsSubCategoryDanceSchools       SubCategorySports = "2241"
	SportsSubCategoryExerciseFitness    SubCategorySports = "2242"
	SportsSubCategoryFanGear            SubCategorySports = "2243"
	SportsSubCategoryFirearmAccessories SubCategorySports = "2244"
	SportsSubCategoryFirearms           SubCategorySports = "2245"
	SportsSubCategoryHunting            SubCategorySports = "2246"
	SportsSubCategoryKnives             SubCategorySports = "2247"
	SportsSubCategoryMartialArtsWeapons SubCategorySports = "2248"
	SportsSubCategorySportGames         SubCategorySports = "2249"
	SportsSubCategorySportingEquipment  SubCategorySports = "2250"
	SportsSubCategorySwimmingPools      SubCategorySports = "2251"
)

type CategorySports struct {
	SubCat SubCategorySports
}

func (c CategorySports) getCategoryCode() string {
	return "1021"
}

func (c CategorySports) getSubCategoryCode() string {
	return string(c.SubCat)
}

type SubCategoryToys string

const (
	ToysSubCategoryArtsCrafts          SubCategoryToys = "2252"
	ToysSubCategoryCameraSupplies      SubCategoryToys = "2253"
	ToysSubCategoryHobbyToyGameShop    SubCategoryToys = "2254"
	ToysSubCategoryMemorabilia         SubCategoryToys = "2255"
	ToysSubCategoryMusic               SubCategoryToys = "2256"
	ToysSubCategoryStampCoin           SubCategoryToys = "2257"
	ToysSubCategoryStationary          SubCategoryToys = "2258"
	ToysSubCategoryVintageCollectibles SubCategoryToys = "2259"
	ToysSubCategoryVideoGamesSystems   SubCategoryToys = "2260"
)

type CategoryToys struct {
	SubCat SubCategoryToys
}

func (c CategoryToys) getCategoryCode() string {
	return "1022"
}

func (c CategoryToys) getSubCategoryCode() string {
	return string(c.SubCat)
}

type SubCategoryTravel string

const (
	TravelSubCategoryAirline                SubCategoryTravel = "2261"
	TravelSubCategoryAutoRental             SubCategoryTravel = "2262"
	TravelSubCategoryBusLine                SubCategoryTravel = "2263"
	TravelSubCategoryCruises                SubCategoryTravel = "2264"
	TravelSubCategoryLodging                SubCategoryTravel = "2265"
	TravelSubCategoryLuggage                SubCategoryTravel = "2266"
	TravelSubCategoryRecreationalServices   SubCategoryTravel = "2267"
	TravelSubCategorySportingCamps          SubCategoryTravel = "2268"
	TravelSubCategoryTaxicabsLimousines     SubCategoryTravel = "2269"
	TravelSubCategoryTimeshares             SubCategoryTravel = "2270"
	TravelSubCategoryTours                  SubCategoryTravel = "2271"
	TravelSubCategoryTrailerParks           SubCategoryTravel = "2272"
	TravelSubCategoryTransportationServices SubCategoryTravel = "2273"
	TravelSubCategoryTravelAgency           SubCategoryTravel = "2274"
)

type CategoryTravel struct {
	SubCat SubCategoryTravel
}

func (c CategoryTravel) getCategoryCode() string {
	return "1023"
}

func (c CategoryTravel) getSubCategoryCode() string {
	return string(c.SubCat)
}

type SubCategoryVehicleSales string

const (
	VehicleSalesSubCategoryAutoDealerNewUsed                SubCategoryVehicleSales = "2275"
	VehicleSalesSubCategoryAutoDealerUsed                   SubCategoryVehicleSales = "2276"
	VehicleSalesSubCategoryAviation                         SubCategoryVehicleSales = "2277"
	VehicleSalesSubCategoryBoatDealer                       SubCategoryVehicleSales = "2278"
	VehicleSalesSubCategoryMobileHomeDealer                 SubCategoryVehicleSales = "2279"
	VehicleSalesSubCategoryMotorcycleDealer                 SubCategoryVehicleSales = "2280"
	VehicleSalesSubCategoryRecreationalUtilityTrailerDealer SubCategoryVehicleSales = "2281"
	VehicleSalesSubCategoryRecreationalVehicleDealer        SubCategoryVehicleSales = "2282"
	VehicleSalesSubCategoryVintageCollectibles              SubCategoryVehicleSales = "2283"
)

type CategoryVehicleSales struct {
	SubCat SubCategoryVehicleSales
}

func (c CategoryVehicleSales) getCategoryCode() string {
	return "1024"
}

func (c CategoryVehicleSales) getSubCategoryCode() string {
	return string(c.SubCat)
}

type SubCategoryVehicleService string

const (
	VehicleServiceSubCategoryNewParts          SubCategoryVehicleService = "2284"
	VehicleServiceSubCategoryUsedParts         SubCategoryVehicleService = "2285"
	VehicleServiceSubCategoryAudioVideo        SubCategoryVehicleService = "2286"
	VehicleServiceSubCategoryBodyRepairPaint   SubCategoryVehicleService = "2287"
	VehicleServiceSubCategoryAutoRental        SubCategoryVehicleService = "2288"
	VehicleServiceSubCategoryAutoService       SubCategoryVehicleService = "2289"
	VehicleServiceSubCategoryTireSupplyService SubCategoryVehicleService = "2290"
	VehicleServiceSubCategoryBoatRental        SubCategoryVehicleService = "2291"
	VehicleServiceSubCategoryCarWash           SubCategoryVehicleService = "2292"
	VehicleServiceSubCategoryMotorHomeRental   SubCategoryVehicleService = "2293"
	VehicleServiceSubCategoryToolsEquipment    SubCategoryVehicleService = "2294"
	VehicleServiceSubCategoryTowingService     SubCategoryVehicleService = "2295"
	VehicleServiceSubCategoryTruckRental       SubCategoryVehicleService = "2296"
	VehicleServiceSubCategoryAccessories       SubCategoryVehicleService = "2297"
)

type CategoryVehicleService struct {
	SubCat SubCategoryVehicleService
}

func (c CategoryVehicleService) getCategoryCode() string {
	return "1025"
}

func (c CategoryVehicleService) getSubCategoryCode() string {
	return string(c.SubCat)
}

type BusinessNameTypeData string

const (
	BusinessNameTypeLegal            BusinessNameTypeData = "LEGAL"
	BusinessNameTypeDoingBusinessAs  BusinessNameTypeData = "DOING_BUSINESS_AS"
	BusinessNameTypeStockTradingName BusinessNameTypeData = "STOCK_TRADING_NAME"
)

type BusinessNameData struct {
	Type BusinessNameTypeData `json:"type,omitempty"`
	Name string               `json:"name,omitempty"`
}

type CurrencyData struct {
	Currency string `json:"currency,omitempty"`
	Value    string `json:"value,omitempty"`
}

type CurrencyRangeData struct {
	MinimumAmount *CurrencyData `json:"minimum_amount,omitempty"`
	MaximumAmount *CurrencyData `json:"maximum_amount,omitempty"`
}

type EmailRoleData string

const (
	EmailRoleCustomerService EmailRoleData = "CUSTOMER_SERVICE"
)

type EmailData struct {
	EmailAddress string        `json:"email_address,omitempty"`
	Role         EmailRoleData `json:"role,omitempty"`
}

type BusinessDetailsData struct {
	PhoneContacts             []OnboardingCommonUserPhoneData `json:"phone_contacts,omitempty"`
	BusinessAddress           *SimplePostalAddressData        `json:"business_address,omitempty"`
	BusinessType              BusinessTypeData                `json:"business_type,omitempty"`
	Category                  BusinessCategory                `json:"category,omitempty"`
	Names                     []BusinessNameData              `json:"names,omitempty"`
	BusinessDescription       string                          `json:"business_description,omitempty"`
	EventDates                []DateData                      `json:"event_dates,omitempty"`
	WebsiteURLS               []string                        `json:"website_urls,omitempty"`
	AnnualSalesVolumeRange    *CurrencyRangeData              `json:"annual_sales_volume_range,omitempty"`
	AverageMonthlyVolumeRange *CurrencyRangeData              `json:"average_monthly_volume_range,omitempty"`
	IdentityDocuments         []IdentityDocumentData          `json:"identity_documents,omitempty"`
	EmailContacts             []EmailData                     `json:"email_contacts,omitempty"`
}

func (b *BusinessDetailsData) MarshalJSON() ([]byte, error) {
	data := struct {
		PhoneContacts             []OnboardingCommonUserPhoneData `json:"phone_contacts,omitempty"`
		BusinessAddress           *SimplePostalAddressData        `json:"business_address,omitempty"`
		BusinessType              BusinessTypeData                `json:"business_type,omitempty"`
		Category                  string                          `json:"category,omitempty"`
		SubCategory               string                          `json:"sub_category,omitempty"`
		Names                     []BusinessNameData              `json:"names,omitempty"`
		BusinessDescription       string                          `json:"business_description,omitempty"`
		EventDates                []DateData                      `json:"event_dates,omitempty"`
		WebsiteURLS               []string                        `json:"website_urls,omitempty"`
		AnnualSalesVolumeRange    *CurrencyRangeData              `json:"annual_sales_volume_range,omitempty"`
		AverageMonthlyVolumeRange *CurrencyRangeData              `json:"average_monthly_volume_range,omitempty"`
		IdentityDocuments         []IdentityDocumentData          `json:"identity_documents,omitempty"`
		EmailContacts             []EmailData                     `json:"email_contacts,omitempty"`
	}{
		PhoneContacts:             b.PhoneContacts,
		BusinessAddress:           b.BusinessAddress,
		BusinessType:              b.BusinessType,
		Category:                  b.Category.getCategoryCode(),
		SubCategory:               b.Category.getSubCategoryCode(),
		Names:                     b.Names,
		BusinessDescription:       b.BusinessDescription,
		EventDates:                b.EventDates,
		WebsiteURLS:               b.WebsiteURLS,
		AnnualSalesVolumeRange:    b.AnnualSalesVolumeRange,
		AverageMonthlyVolumeRange: b.AverageMonthlyVolumeRange,
		IdentityDocuments:         b.IdentityDocuments,
		EmailContacts:             b.EmailContacts,
	}
	return json.Marshal(&data)
}

func (b *BusinessDetailsData) UnmarshalJSON(bs []byte) error {
	data := struct {
		PhoneContacts             []OnboardingCommonUserPhoneData `json:"phone_contacts,omitempty"`
		BusinessAddress           *SimplePostalAddressData        `json:"business_address,omitempty"`
		BusinessType              BusinessTypeData                `json:"business_type,omitempty"`
		Category                  string                          `json:"category,omitempty"`
		SubCategory               string                          `json:"sub_category,omitempty"`
		Names                     []BusinessNameData              `json:"names,omitempty"`
		BusinessDescription       string                          `json:"business_description,omitempty"`
		EventDates                []DateData                      `json:"event_dates,omitempty"`
		WebsiteURLS               []string                        `json:"website_urls"`
		AnnualSalesVolumeRange    *CurrencyRangeData              `json:"annual_sales_volume_range,omitempty"`
		AverageMonthlyVolumeRange *CurrencyRangeData              `json:"average_monthly_volume_range,omitempty"`
		IdentityDocuments         []IdentityDocumentData          `json:"identity_documents,omitempty"`
		EmailContacts             []EmailData                     `json:"email_contacts,omitempty"`
	}{}
	err := json.Unmarshal(bs, &data)
	if err != nil {
		return err
	}
	b.PhoneContacts = data.PhoneContacts
	b.BusinessAddress = data.BusinessAddress
	b.BusinessType = data.BusinessType
	// TODO: Add categories
	b.Names = data.Names
	b.BusinessDescription = data.BusinessDescription
	b.EventDates = data.EventDates
	b.WebsiteURLS = data.WebsiteURLS
	b.AnnualSalesVolumeRange = data.AnnualSalesVolumeRange
	b.AverageMonthlyVolumeRange = data.AverageMonthlyVolumeRange
	b.IdentityDocuments = data.IdentityDocuments
	b.EmailContacts = data.EmailContacts
	return nil
}

type BankAccountTypeData string

const (
	BankAccountTypeChecking BankAccountTypeData = "CHECKING"
	BankAccountTypeSavings  BankAccountTypeData = "SAVINGS"
)

type BankDetailsData struct {
	NickName       string                   `json:"nick_name,omitempty"`
	AccountNumber  string                   `json:"account_number,omitempty"`
	AccountType    BankAccountTypeData      `json:"account_type,omitempty"`
	CurrencyCode   string                   `json:"currency_code,omitempty"`
	Identifiers    []string                 `json:"identifiers,omitempty"`
	BranchLocation *SimplePostalAddressData `json:"branch_location,omitempty"`
	MandateAgreed  *bool                    `json:"mandate_agreed,omitempty"`
}

type FinancialInstrumentDataType struct {
	BankDetails *BankDetailsData `json:"bank_details,omitempty"`
}

type AccountIdentifierTypeData string

const (
	AccountIdentifierTypePayerID AccountIdentifierTypeData = "PAYER_ID"
)

type AccountIdentifierData struct {
	Type  AccountIdentifierTypeData `json:"type,omitempty"`
	Value string                    `json:"value,omitempty"`
}

type PartnerSpecificIdentifierTypeData string

const (
	PartnerSpecificIdentifierTypeTrackingID       PartnerSpecificIdentifierTypeData = "TRACKING_ID"
	PartnerSpecificIdentifierTypeAccountLinkingID PartnerSpecificIdentifierTypeData = "ACCOUNT_LINKING_ID"
)

type PartnerSpecificIdentifierData struct {
	Type  PartnerSpecificIdentifierTypeData `json:"type,omitempty"`
	Value string                            `json:"value,omitempty"`
}

type UserData struct {
	CustomerType               CustomerTypeData                `json:"customer_type,omitempty"`
	PersonDetails              *PersonDetailsData              `json:"person_details,omitempty"`
	BusinessDetails            *BusinessDetailsData            `json:"business_details,omitempty"`
	FinancialInstrumentData    *FinancialInstrumentDataType    `json:"financial_instrument_data,omitempty"`
	PreferredLanguageCode      string                          `json:"preferred_language_code,omitempty"`
	PrimaryCurrencyCode        string                          `json:"primary_currency_code,omitempty"`
	ReferralUserPayerID        *AccountIdentifierData          `json:"referral_user_payer_id,omitempty"`
	PartnerSpecificIdentifiers []PartnerSpecificIdentifierData `json:"partner_specific_identifiers,omitempty"`
}

type CapabilityData string

const (
	CapabilityApiIntegration             CapabilityData = "API_INTEGRATION"
	CapabilityBankAddition               CapabilityData = "BANK_ADDITION"
	CapabilityBillingAgreement           CapabilityData = "BILLING_AGREEMENT"
	CapabilityContextualMarketingConsent CapabilityData = "CONTEXTUAL_MARKETING_CONSENT"
)

type ClassicIntegrationTypeData string

const (
	ClassicIntegrationThirdParty              ClassicIntegrationTypeData = "THIRD_PARTY"
	ClassicIntegrationFirstPartyIntegrated    ClassicIntegrationTypeData = "FIRST_PARTY_INTEGRATED"
	ClassicIntegrationFirstPartyNonIntegrated ClassicIntegrationTypeData = "FIRST_PARTY_NON_INTEGRATED"
)

type IntegrationMethodData string

const (
	IntegrationMethodBrainTree IntegrationMethodData = "BRAINTREE"
	IntegrationMethodPaypal    IntegrationMethodData = "PAYPAL"
)

type IntegrationTypeData string

const (
	IntegrationTypeThirdParty IntegrationTypeData = "THIRD_PARTY"
)

type RestAPIIntegrationData struct {
	IntegrationMethod IntegrationMethodData `json:"integration_method,omitempty"`
	IntegrationType   IntegrationTypeData   `json:"integration_type,omitempty"`
}

type ReferralDataClassicPermissionData string

const (
	ExpressCheckout                ReferralDataClassicPermissionData = "EXPRESS_CHECKOUT"
	Refund                         ReferralDataClassicPermissionData = "REFUND"
	DirectPayment                  ReferralDataClassicPermissionData = "DIRECT_PAYMENT"
	AuthCapture                    ReferralDataClassicPermissionData = "AUTH_CAPTURE"
	ButtonManager                  ReferralDataClassicPermissionData = "BUTTON_MANAGER"
	AccountBalance                 ReferralDataClassicPermissionData = "ACCOUNT_BALANCE"
	TransactionDetails             ReferralDataClassicPermissionData = "TRANSACTION_DETAILS"
	TransactionSearch              ReferralDataClassicPermissionData = "TRANSACTION_SEARCH"
	ReferenceTransaction           ReferralDataClassicPermissionData = "REFERENCE_TRANSACTION"
	RecurringPayments              ReferralDataClassicPermissionData = "RECURRING_PAYMENTS"
	ManagePendingTransactionStatus ReferralDataClassicPermissionData = "MANAGE_PENDING_TRANSACTION_STATUS"
	NonReferencedCredit            ReferralDataClassicPermissionData = "NON_REFERENCED_CREDIT"
	EncryptedWebsitePayments       ReferralDataClassicPermissionData = "ENCRYPTED_WEBSITE_PAYMENTS"
	MobileCheckout                 ReferralDataClassicPermissionData = "MOBILE_CHECKOUT"
	AirTravel                      ReferralDataClassicPermissionData = "AIR_TRAVEL"
	Invoicing                      ReferralDataClassicPermissionData = "INVOICING"
	AccessBasicPersonalData        ReferralDataClassicPermissionData = "ACCESS_BASIC_PERSONAL_DATA"
)

type SupportedClassicPermissionsData struct {
	ReferralDataClassicPermissions *ReferralDataClassicPermissionData `json:"referral_data-classic_permission_enum,omitempty"`
}

type ClassicThirdPartyDetailsData struct {
	PermissionList []SupportedClassicPermissionsData `json:"permission_list,omitempty"`
}

type ClassicFirstPartyDetailsData string

const (
	ClassicFirstPartySignature   ClassicFirstPartyDetailsData = "SIGNATURE"
	ClassicFirstPartyCertificate ClassicFirstPartyDetailsData = "CERTIFICATE"
)

type ReferralDataRestFeaturesData string

const (
	ReferralDataRestFeaturesPayment                   ReferralDataRestFeaturesData = "PAYMENT"
	ReferralDataRestFeaturesRefund                    ReferralDataRestFeaturesData = "REFUND"
	ReferralDataRestFeaturesFuturePayment             ReferralDataRestFeaturesData = "FUTURE_PAYMENT"
	ReferralDataRestFeaturesDirectPayment             ReferralDataRestFeaturesData = "DIRECT_PAYMENT"
	ReferralDataRestFeaturesPartnerFee                ReferralDataRestFeaturesData = "PARTNER_FEE"
	ReferralDataRestFeaturesDelayDisbursement         ReferralDataRestFeaturesData = "DELAY_FUNDS_DISBURSEMENT"
	ReferralDataRestFeaturesSweepFunds                ReferralDataRestFeaturesData = "SWEEP_FUNDS_EXTERNAL_SINK"
	ReferralDataRestFeaturesAdvancedTransactionSearch ReferralDataRestFeaturesData = "ADVANCED_TRANSACTIONS_SEARCH"
	ReferralDataRestFeaturesReadDispute               ReferralDataRestFeaturesData = "READ_SELLER_DISPUTE"
	ReferralDataRestFeaturesUpdateDispute             ReferralDataRestFeaturesData = "UPDATE_SELLER_DISPUTE"
)

type RestThirdPartyDetailsData struct {
	PartnerClientID string                         `json:"partner_client_id,omitempty"`
	FeatureList     []ReferralDataRestFeaturesData `json:"feature_list,omitempty"`
}

type IntegrationDetailsData struct {
	PartnerID                 string                        `json:"partner_id"`
	ClassicAPIIntegrationType *ClassicIntegrationTypeData   `json:"classic_api_integration_type,omitempty"`
	RestAPIIntegration        *RestAPIIntegrationData       `json:"rest_api_integration,omitempty"`
	ClassicThirdPartyDetails  *ClassicThirdPartyDetailsData `json:"classic_third_party_details,omitempty"`
	ClassicFirstPartyDetails  *ClassicFirstPartyDetailsData `json:"classic_first_party_details,omitempty"`
	RestThirdPartyDetails     *RestThirdPartyDetailsData    `json:"rest_third_party_details,omitempty"`
}

type BillingExperiencePreferenceData struct {
	ExperienceID      string `json:"experience_id,omitempty"`
	BillingContextSet *bool  `json:"billing_context_set,omitempty"`
}

type BillingAgreementData struct {
	Description                 string                           `json:"description,omitempty"`
	BillingExperiencePreference *BillingExperiencePreferenceData `json:"billing_experience_preference,omitempty"`
	MerchantCustomData          string                           `json:"merchant_custom_data,omitempty"`
	ApprovalURL                 string                           `json:"approval_url,omitempty"`
	ECToken                     string                           `json:"ec_token,omitempty"`
}

type CustomerCapabilitiesData struct {
	Capability               CapabilityData          `json:"capability,omitempty"`
	ApiIntegrationPreference *IntegrationDetailsData `json:"api_integration_preference,omitempty"`
	BillingAgreement         *BillingAgreementData   `json:"billing_agreement,omitempty"`
}

type WebExperiencePreferenceData struct {
	PartnerLogoURL          string `json:"partner_logo_url,omitempty"`
	ReturnURL               string `json:"return_url,omitempty"`
	ReturnURLDescription    string `json:"return_url_description,omitempty"`
	ActionRenewalURL        string `json:"action_renewal_url,omitempty"`
	ShowAddCreditCard       *bool  `json:"show_add_credit_card,omitempty"`
	ShowMobileConfirm       *bool  `json:"show_mobile_confirm,omitempty"`
	UseMiniBrowser          *bool  `json:"use_mini_browser,omitempty"`
	UseHuaEmailConfirmation *bool  `json:"use_hua_email_confirmation,omitempty"`
}

type LegalConsentTypeData string

const (
	LegalConsentTypeShareDataConsent LegalConsentTypeData = "SHARE_DATA_CONSENT"
)

type LegalConsentData struct {
	Type    LegalConsentTypeData `json:"type"`
	Granted bool                 `json:"granted"`
}

type ReferralDataProductNameData string

const (
	ReferralDataExpressCheckout ReferralDataProductNameData = "EXPRESS_CHECKOUT"
)

type CreatePartnerReferralParams struct {
	CustomerData            *UserData                     `json:"customer_data,omitempty"`
	RequestedCapabilities   []CustomerCapabilitiesData    `json:"requested_capabilities,omitempty"`
	WebExperiencePreference *WebExperiencePreferenceData  `json:"web_experience_preference,omitempty"`
	CollectedConsents       []LegalConsentData            `json:"collected_consents,omitempty"`
	Products                []ReferralDataProductNameData `json:"products,omitempty"`
}

type CreatePartnerReferralResponse struct {
	RedirectURL       string
	PartnerReferralID string
}

func (c *CreatePartnerReferralResponse) UnmarshalJSON(b []byte) error {
	response := struct {
		Links []struct {
			Href        string
			Rel         string
			Method      string
			Description string
		}
	}{}
	err := json.Unmarshal(b, &response)
	if err != nil {
		return err
	}
	for _, v := range response.Links {
		if v.Rel == "self" {
			u, err := url.Parse(v.Href)
			if err != nil {
				return err
			}
			const prefix = "/v1/customer/partner-referrals/"
			if !strings.HasPrefix(u.Path, prefix) {
				return errors.New("Bad path in partner referral ID: " + u.Path)
			}
			c.PartnerReferralID = u.Path[len(prefix):]
		} else if v.Rel == "action_url" {
			c.RedirectURL = v.Href
		}
	}
	return nil
}

type GetPartnerReferralResponse struct {
	PartnerReferralID string
	SubmitterPayerID  string
	ReferralData      *CreatePartnerReferralParams
	RedirectURL       string
}

func (g *GetPartnerReferralResponse) UnmarshalJSON(b []byte) error {
	response := struct {
		PartnerReferralID string                       `json:"partner_referral_id"`
		SubmitterPayerID  string                       `json:"submitter_payer_id"`
		ReferralData      *CreatePartnerReferralParams `json:"referral_data"`
		Links             []struct {
			Href        string `json:"href"`
			Rel         string `json:"rel"`
			Method      string `json:"method"`
			Description string `json:"description"`
		} `json:"links"`
	}{}
	err := json.Unmarshal(b, &response)
	if err != nil {
		return err
	}
	g.PartnerReferralID = response.PartnerReferralID
	g.SubmitterPayerID = response.SubmitterPayerID
	g.ReferralData = response.ReferralData
	for _, v := range response.Links {
		if v.Rel == "action_url" {
			g.RedirectURL = v.Href
			break
		}
	}
	return nil
}
