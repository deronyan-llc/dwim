package generated

// HealthTopicContent is a generated struct representing the https://schema.org/HealthTopicContent class.
type HealthTopicContent struct {
	HasHealthAspect *HealthAspectEnumeration `json:"hashealthaspect"`
}

// MedicalObservationalStudyDesign is a generated struct representing the https://schema.org/MedicalObservationalStudyDesign class.
type MedicalObservationalStudyDesign struct {
}

// MerchantReturnPolicy is a generated struct representing the https://schema.org/MerchantReturnPolicy class.
type MerchantReturnPolicy struct {
	AdditionalProperty                      *PropertyValue                        `json:"additionalproperty"`
	ApplicableCountry                       *Text                                 `json:"applicablecountry"`
	CustomerRemorseReturnFees               *ReturnFeesEnumeration                `json:"customerremorsereturnfees"`
	CustomerRemorseReturnLabelSource        *ReturnLabelSourceEnumeration         `json:"customerremorsereturnlabelsource"`
	CustomerRemorseReturnShippingFeesAmount *MonetaryAmount                       `json:"customerremorsereturnshippingfeesamount"`
	InStoreReturnsOffered                   *Boolean                              `json:"instorereturnsoffered"`
	ItemCondition                           *OfferItemCondition                   `json:"itemcondition"`
	ItemDefectReturnFees                    *ReturnFeesEnumeration                `json:"itemdefectreturnfees"`
	ItemDefectReturnLabelSource             *ReturnLabelSourceEnumeration         `json:"itemdefectreturnlabelsource"`
	ItemDefectReturnShippingFeesAmount      *MonetaryAmount                       `json:"itemdefectreturnshippingfeesamount"`
	MerchantReturnDays                      *Integer                              `json:"merchantreturndays"`
	MerchantReturnLink                      *URL                                  `json:"merchantreturnlink"`
	RefundType                              *RefundTypeEnumeration                `json:"refundtype"`
	RestockingFee                           *Number                               `json:"restockingfee"`
	ReturnFees                              *ReturnFeesEnumeration                `json:"returnfees"`
	ReturnLabelSource                       *ReturnLabelSourceEnumeration         `json:"returnlabelsource"`
	ReturnMethod                            *ReturnMethodEnumeration              `json:"returnmethod"`
	ReturnPolicyCategory                    *MerchantReturnEnumeration            `json:"returnpolicycategory"`
	ReturnPolicyCountry                     *Text                                 `json:"returnpolicycountry"`
	ReturnPolicySeasonalOverride            *MerchantReturnPolicySeasonalOverride `json:"returnpolicyseasonaloverride"`
	ReturnShippingFeesAmount                *MonetaryAmount                       `json:"returnshippingfeesamount"`
}

// PaintAction is a generated struct representing the https://schema.org/PaintAction class.
type PaintAction struct {
}

// Property is a generated struct representing the https://schema.org/Property class.
type Property struct {
	DomainIncludes *Class    `json:"domainincludes"`
	InverseOf      *Property `json:"inverseof"`
	RangeIncludes  *Class    `json:"rangeincludes"`
	SupersededBy   *Property `json:"supersededby"`
}

// ApartmentComplex is a generated struct representing the https://schema.org/ApartmentComplex class.
type ApartmentComplex struct {
	NumberOfAccommodationUnits          *QuantitativeValue `json:"numberofaccommodationunits"`
	NumberOfAvailableAccommodationUnits *QuantitativeValue `json:"numberofavailableaccommodationunits"`
	NumberOfBedrooms                    *QuantitativeValue `json:"numberofbedrooms"`
	PetsAllowed                         *Text              `json:"petsallowed"`
	TourBookingPage                     *URL               `json:"tourbookingpage"`
}

// CollectionPage is a generated struct representing the https://schema.org/CollectionPage class.
type CollectionPage struct {
}

// LymphaticVessel is a generated struct representing the https://schema.org/LymphaticVessel class.
type LymphaticVessel struct {
	OriginatesFrom *Vessel           `json:"originatesfrom"`
	RegionDrained  *AnatomicalSystem `json:"regiondrained"`
	RunsTo         *Vessel           `json:"runsto"`
}

// OfferShippingDetails is a generated struct representing the https://schema.org/OfferShippingDetails class.
type OfferShippingDetails struct {
	DeliveryTime         *ShippingDeliveryTime `json:"deliverytime"`
	Depth                *QuantitativeValue    `json:"depth"`
	DoesNotShip          *Boolean              `json:"doesnotship"`
	Height               *QuantitativeValue    `json:"height"`
	ShippingDestination  *DefinedRegion        `json:"shippingdestination"`
	ShippingLabel        *Text                 `json:"shippinglabel"`
	ShippingOrigin       *DefinedRegion        `json:"shippingorigin"`
	ShippingRate         *MonetaryAmount       `json:"shippingrate"`
	ShippingSettingsLink *URL                  `json:"shippingsettingslink"`
	TransitTimeLabel     *Text                 `json:"transittimelabel"`
	Weight               *QuantitativeValue    `json:"weight"`
	Width                *QuantitativeValue    `json:"width"`
}

// RentalCarReservation is a generated struct representing the https://schema.org/RentalCarReservation class.
type RentalCarReservation struct {
	DropoffLocation *Place    `json:"dropofflocation"`
	DropoffTime     *DateTime `json:"dropofftime"`
	PickupLocation  *Place    `json:"pickuplocation"`
	PickupTime      *DateTime `json:"pickuptime"`
}

// ScreeningEvent is a generated struct representing the https://schema.org/ScreeningEvent class.
type ScreeningEvent struct {
	SubtitleLanguage *Text  `json:"subtitlelanguage"`
	VideoFormat      *Text  `json:"videoformat"`
	WorkPresented    *Movie `json:"workpresented"`
}

// WPHeader is a generated struct representing the https://schema.org/WPHeader class.
type WPHeader struct {
}

// WebAPI is a generated struct representing the https://schema.org/WebAPI class.
type WebAPI struct {
	Documentation *URL `json:"documentation"`
}

// AnatomicalStructure is a generated struct representing the https://schema.org/AnatomicalStructure class.
type AnatomicalStructure struct {
	AssociatedPathophysiology *Text                `json:"associatedpathophysiology"`
	BodyLocation              *Text                `json:"bodylocation"`
	ConnectedTo               *AnatomicalStructure `json:"connectedto"`
	Diagram                   *ImageObject         `json:"diagram"`
	PartOfSystem              *AnatomicalSystem    `json:"partofsystem"`
	RelatedCondition          *MedicalCondition    `json:"relatedcondition"`
	RelatedTherapy            *MedicalTherapy      `json:"relatedtherapy"`
	SubStructure              *AnatomicalStructure `json:"substructure"`
}

// Legislation is a generated struct representing the https://schema.org/Legislation class.
type Legislation struct {
	LegislationChanges      *Legislation      `json:"legislationchanges"`
	LegislationConsolidates *Legislation      `json:"legislationconsolidates"`
	LegislationDate         *Date             `json:"legislationdate"`
	LegislationDateVersion  *Date             `json:"legislationdateversion"`
	LegislationIdentifier   *URL              `json:"legislationidentifier"`
	LegislationJurisdiction *Text             `json:"legislationjurisdiction"`
	LegislationLegalForce   *LegalForceStatus `json:"legislationlegalforce"`
	LegislationPassedBy     *Person           `json:"legislationpassedby"`
	LegislationResponsible  *Person           `json:"legislationresponsible"`
	LegislationTransposes   *Legislation      `json:"legislationtransposes"`
	LegislationType         *Text             `json:"legislationtype"`
	Jurisdiction            *Text             `json:"jurisdiction"`
	LegislationApplies      *Legislation      `json:"legislationapplies"`
}

// BreadcrumbList is a generated struct representing the https://schema.org/BreadcrumbList class.
type BreadcrumbList struct {
}

// DatedMoneySpecification is a generated struct representing the https://schema.org/DatedMoneySpecification class.
type DatedMoneySpecification struct {
	Amount    *Number   `json:"amount"`
	Currency  *Text     `json:"currency"`
	EndDate   *DateTime `json:"enddate"`
	StartDate *DateTime `json:"startdate"`
}

// FindAction is a generated struct representing the https://schema.org/FindAction class.
type FindAction struct {
}

// GamePlayMode is a generated struct representing the https://schema.org/GamePlayMode class.
type GamePlayMode struct {
}

// DisagreeAction is a generated struct representing the https://schema.org/DisagreeAction class.
type DisagreeAction struct {
}

// HyperTocEntry is a generated struct representing the https://schema.org/HyperTocEntry class.
type HyperTocEntry struct {
	AssociatedMedia *MediaObject   `json:"associatedmedia"`
	TocContinuation *HyperTocEntry `json:"toccontinuation"`
	Utterances      *Text          `json:"utterances"`
}

// MedicalScholarlyArticle is a generated struct representing the https://schema.org/MedicalScholarlyArticle class.
type MedicalScholarlyArticle struct {
	PublicationType *Text `json:"publicationtype"`
}

// ReturnFeesEnumeration is a generated struct representing the https://schema.org/ReturnFeesEnumeration class.
type ReturnFeesEnumeration struct {
}

// VideoGame is a generated struct representing the https://schema.org/VideoGame class.
type VideoGame struct {
	Actors       *Person       `json:"actors"`
	CheatCode    *CreativeWork `json:"cheatcode"`
	Directors    *Person       `json:"directors"`
	GameEdition  *Text         `json:"gameedition"`
	GamePlatform *URL          `json:"gameplatform"`
	GameTip      *CreativeWork `json:"gametip"`
	MusicBy      *Person       `json:"musicby"`
	PlayMode     *GamePlayMode `json:"playmode"`
	Trailer      *VideoObject  `json:"trailer"`
	Director     *Person       `json:"director"`
	GameServer   *GameServer   `json:"gameserver"`
	Actor        *Person       `json:"actor"`
}

// PoliticalParty is a generated struct representing the https://schema.org/PoliticalParty class.
type PoliticalParty struct {
}

// ResumeAction is a generated struct representing the https://schema.org/ResumeAction class.
type ResumeAction struct {
}

// ShortStory is a generated struct representing the https://schema.org/ShortStory class.
type ShortStory struct {
}

// EnergyConsumptionDetails is a generated struct representing the https://schema.org/EnergyConsumptionDetails class.
type EnergyConsumptionDetails struct {
	EnergyEfficiencyScaleMax    *EUEnergyEfficiencyEnumeration `json:"energyefficiencyscalemax"`
	EnergyEfficiencyScaleMin    *EUEnergyEfficiencyEnumeration `json:"energyefficiencyscalemin"`
	HasEnergyEfficiencyCategory *EnergyEfficiencyEnumeration   `json:"hasenergyefficiencycategory"`
}

// InviteAction is a generated struct representing the https://schema.org/InviteAction class.
type InviteAction struct {
	Event *Event `json:"event"`
}

// MedicalWebPage is a generated struct representing the https://schema.org/MedicalWebPage class.
type MedicalWebPage struct {
	Aspect          *Text                `json:"aspect"`
	MedicalAudience *MedicalAudienceType `json:"medicalaudience"`
}

// ResearchProject is a generated struct representing the https://schema.org/ResearchProject class.
type ResearchProject struct {
}

// SearchAction is a generated struct representing the https://schema.org/SearchAction class.
type SearchAction struct {
	Query *Text `json:"query"`
}

// ShippingDeliveryTime is a generated struct representing the https://schema.org/ShippingDeliveryTime class.
type ShippingDeliveryTime struct {
	BusinessDays *OpeningHoursSpecification `json:"businessdays"`
	CutoffTime   *Time                      `json:"cutofftime"`
	HandlingTime *QuantitativeValue         `json:"handlingtime"`
	TransitTime  *QuantitativeValue         `json:"transittime"`
}

// TieAction is a generated struct representing the https://schema.org/TieAction class.
type TieAction struct {
}

// BackgroundNewsArticle is a generated struct representing the https://schema.org/BackgroundNewsArticle class.
type BackgroundNewsArticle struct {
}

// GeneralContractor is a generated struct representing the https://schema.org/GeneralContractor class.
type GeneralContractor struct {
}

// MobileApplication is a generated struct representing the https://schema.org/MobileApplication class.
type MobileApplication struct {
	CarrierRequirements *Text `json:"carrierrequirements"`
}

// Trip is a generated struct representing the https://schema.org/Trip class.
type Trip struct {
	ArrivalTime   *Time   `json:"arrivaltime"`
	DepartureTime *Time   `json:"departuretime"`
	Itinerary     *Place  `json:"itinerary"`
	TripOrigin    *Place  `json:"triporigin"`
	Offers        *Offer  `json:"offers"`
	PartOfTrip    *Trip   `json:"partoftrip"`
	Provider      *Person `json:"provider"`
	SubTrip       *Trip   `json:"subtrip"`
}

// AutoBodyShop is a generated struct representing the https://schema.org/AutoBodyShop class.
type AutoBodyShop struct {
}

// DeliveryChargeSpecification is a generated struct representing the https://schema.org/DeliveryChargeSpecification class.
type DeliveryChargeSpecification struct {
	AppliesToDeliveryMethod *DeliveryMethod `json:"appliestodeliverymethod"`
	EligibleRegion          *Text           `json:"eligibleregion"`
	IneligibleRegion        *Text           `json:"ineligibleregion"`
	AreaServed              *Text           `json:"areaserved"`
}

// Intangible is a generated struct representing the https://schema.org/Intangible class.
type Intangible struct {
}

// ReportageNewsArticle is a generated struct representing the https://schema.org/ReportageNewsArticle class.
type ReportageNewsArticle struct {
}

// Beach is a generated struct representing the https://schema.org/Beach class.
type Beach struct {
}

// LikeAction is a generated struct representing the https://schema.org/LikeAction class.
type LikeAction struct {
}

// TechArticle is a generated struct representing the https://schema.org/TechArticle class.
type TechArticle struct {
	Dependencies     *Text `json:"dependencies"`
	ProficiencyLevel *Text `json:"proficiencylevel"`
}

// ChildrensEvent is a generated struct representing the https://schema.org/ChildrensEvent class.
type ChildrensEvent struct {
}

// StadiumOrArena is a generated struct representing the https://schema.org/StadiumOrArena class.
type StadiumOrArena struct {
}

// EngineSpecification is a generated struct representing the https://schema.org/EngineSpecification class.
type EngineSpecification struct {
	EngineDisplacement *QuantitativeValue `json:"enginedisplacement"`
	EnginePower        *QuantitativeValue `json:"enginepower"`
	EngineType         *URL               `json:"enginetype"`
	FuelType           *URL               `json:"fueltype"`
	Torque             *QuantitativeValue `json:"torque"`
}

// LegalForceStatus is a generated struct representing the https://schema.org/LegalForceStatus class.
type LegalForceStatus struct {
}

// AMRadioChannel is a generated struct representing the https://schema.org/AMRadioChannel class.
type AMRadioChannel struct {
}

// BeautySalon is a generated struct representing the https://schema.org/BeautySalon class.
type BeautySalon struct {
}

// Enumeration is a generated struct representing the https://schema.org/Enumeration class.
type Enumeration struct {
	SupersededBy *Property `json:"supersededby"`
}

// ListenAction is a generated struct representing the https://schema.org/ListenAction class.
type ListenAction struct {
}

// Code is a generated struct representing the https://schema.org/Code class.
type Code struct {
}

// CourseInstance is a generated struct representing the https://schema.org/CourseInstance class.
type CourseInstance struct {
	CourseMode     *URL      `json:"coursemode"`
	CourseSchedule *Schedule `json:"courseschedule"`
	CourseWorkload *Text     `json:"courseworkload"`
	Instructor     *Person   `json:"instructor"`
}

// DefinedTermSet is a generated struct representing the https://schema.org/DefinedTermSet class.
type DefinedTermSet struct {
	HasDefinedTerm *DefinedTerm `json:"hasdefinedterm"`
}

// PsychologicalTreatment is a generated struct representing the https://schema.org/PsychologicalTreatment class.
type PsychologicalTreatment struct {
}

// Episode is a generated struct representing the https://schema.org/Episode class.
type Episode struct {
	Actors            *Person             `json:"actors"`
	Directors         *Person             `json:"directors"`
	EpisodeNumber     *Text               `json:"episodenumber"`
	MusicBy           *Person             `json:"musicby"`
	PartOfSeason      *CreativeWorkSeason `json:"partofseason"`
	ProductionCompany *Organization       `json:"productioncompany"`
	Trailer           *VideoObject        `json:"trailer"`
	Director          *Person             `json:"director"`
	Duration          *Duration           `json:"duration"`
	PartOfSeries      *CreativeWorkSeries `json:"partofseries"`
	Actor             *Person             `json:"actor"`
}

// UseAction is a generated struct representing the https://schema.org/UseAction class.
type UseAction struct {
}

// DrugCost is a generated struct representing the https://schema.org/DrugCost class.
type DrugCost struct {
	ApplicableLocation *AdministrativeArea `json:"applicablelocation"`
	CostCategory       *DrugCostCategory   `json:"costcategory"`
	CostCurrency       *Text               `json:"costcurrency"`
	CostOrigin         *Text               `json:"costorigin"`
	CostPerUnit        *Text               `json:"costperunit"`
	DrugUnit           *Text               `json:"drugunit"`
}

// House is a generated struct representing the https://schema.org/House class.
type House struct {
	NumberOfRooms *QuantitativeValue `json:"numberofrooms"`
}

// Motel is a generated struct representing the https://schema.org/Motel class.
type Motel struct {
}

// NGO is a generated struct representing the https://schema.org/NGO class.
type NGO struct {
}

// _3DModel is a generated struct representing the https://schema.org/3DModel class.
type _3DModel struct {
	IsResizable *Boolean `json:"isresizable"`
}

// UserTweets is a generated struct representing the https://schema.org/UserTweets class.
type UserTweets struct {
}

// Course is a generated struct representing the https://schema.org/Course class.
type Course struct {
	AvailableLanguage             *Text            `json:"availablelanguage"`
	CourseCode                    *Text            `json:"coursecode"`
	CoursePrerequisites           *Text            `json:"courseprerequisites"`
	EducationalCredentialAwarded  *URL             `json:"educationalcredentialawarded"`
	FinancialAidEligible          *Text            `json:"financialaideligible"`
	HasCourseInstance             *CourseInstance  `json:"hascourseinstance"`
	NumberOfCredits               *StructuredValue `json:"numberofcredits"`
	OccupationalCredentialAwarded *URL             `json:"occupationalcredentialawarded"`
	SyllabusSections              *Syllabus        `json:"syllabussections"`
	TotalHistoricalEnrollment     *Integer         `json:"totalhistoricalenrollment"`
}

// EmergencyService is a generated struct representing the https://schema.org/EmergencyService class.
type EmergencyService struct {
}

// FinancialService is a generated struct representing the https://schema.org/FinancialService class.
type FinancialService struct {
	FeesAndCommissionsSpecification *URL `json:"feesandcommissionsspecification"`
}

// LegalValueLevel is a generated struct representing the https://schema.org/LegalValueLevel class.
type LegalValueLevel struct {
}

// PodcastSeries is a generated struct representing the https://schema.org/PodcastSeries class.
type PodcastSeries struct {
	WebFeed *URL    `json:"webfeed"`
	Actor   *Person `json:"actor"`
}

// MedicalStudyStatus is a generated struct representing the https://schema.org/MedicalStudyStatus class.
type MedicalStudyStatus struct {
}

// ActivateAction is a generated struct representing the https://schema.org/ActivateAction class.
type ActivateAction struct {
}

// BuyAction is a generated struct representing the https://schema.org/BuyAction class.
type BuyAction struct {
	Vendor          *Person          `json:"vendor"`
	WarrantyPromise *WarrantyPromise `json:"warrantypromise"`
	Seller          *Person          `json:"seller"`
}

// CafeOrCoffeeShop is a generated struct representing the https://schema.org/CafeOrCoffeeShop class.
type CafeOrCoffeeShop struct {
}

// MobilePhoneStore is a generated struct representing the https://schema.org/MobilePhoneStore class.
type MobilePhoneStore struct {
}

// CableOrSatelliteService is a generated struct representing the https://schema.org/CableOrSatelliteService class.
type CableOrSatelliteService struct {
}

// MedicalObservationalStudy is a generated struct representing the https://schema.org/MedicalObservationalStudy class.
type MedicalObservationalStudy struct {
	StudyDesign *MedicalObservationalStudyDesign `json:"studydesign"`
}

// Seat is a generated struct representing the https://schema.org/Seat class.
type Seat struct {
	SeatNumber  *Text `json:"seatnumber"`
	SeatRow     *Text `json:"seatrow"`
	SeatSection *Text `json:"seatsection"`
	SeatingType *Text `json:"seatingtype"`
}

// SportingGoodsStore is a generated struct representing the https://schema.org/SportingGoodsStore class.
type SportingGoodsStore struct {
}

// VisualArtwork is a generated struct representing the https://schema.org/VisualArtwork class.
type VisualArtwork struct {
	ArtEdition     *Text              `json:"artedition"`
	ArtMedium      *URL               `json:"artmedium"`
	Artform        *URL               `json:"artform"`
	Artist         *Person            `json:"artist"`
	Colorist       *Person            `json:"colorist"`
	Depth          *QuantitativeValue `json:"depth"`
	Height         *QuantitativeValue `json:"height"`
	Inker          *Person            `json:"inker"`
	Letterer       *Person            `json:"letterer"`
	Penciler       *Person            `json:"penciler"`
	Surface        *URL               `json:"surface"`
	Width          *QuantitativeValue `json:"width"`
	ArtworkSurface *URL               `json:"artworksurface"`
}

// Campground is a generated struct representing the https://schema.org/Campground class.
type Campground struct {
}

// DigitalPlatformEnumeration is a generated struct representing the https://schema.org/DigitalPlatformEnumeration class.
type DigitalPlatformEnumeration struct {
}

// GolfCourse is a generated struct representing the https://schema.org/GolfCourse class.
type GolfCourse struct {
}

// Room is a generated struct representing the https://schema.org/Room class.
type Room struct {
}

// HealthClub is a generated struct representing the https://schema.org/HealthClub class.
type HealthClub struct {
}

// Quantity is a generated struct representing the https://schema.org/Quantity class.
type Quantity struct {
}

// MaximumDoseSchedule is a generated struct representing the https://schema.org/MaximumDoseSchedule class.
type MaximumDoseSchedule struct {
}

// SurgicalProcedure is a generated struct representing the https://schema.org/SurgicalProcedure class.
type SurgicalProcedure struct {
}

// Article is a generated struct representing the https://schema.org/Article class.
type Article struct {
	ArticleBody    *Text    `json:"articlebody"`
	ArticleSection *Text    `json:"articlesection"`
	Backstory      *Text    `json:"backstory"`
	PageEnd        *Text    `json:"pageend"`
	PageStart      *Text    `json:"pagestart"`
	Pagination     *Text    `json:"pagination"`
	Speakable      *URL     `json:"speakable"`
	WordCount      *Integer `json:"wordcount"`
}

// DoseSchedule is a generated struct representing the https://schema.org/DoseSchedule class.
type DoseSchedule struct {
	DoseUnit         *Text             `json:"doseunit"`
	DoseValue        *QualitativeValue `json:"dosevalue"`
	TargetPopulation *Text             `json:"targetpopulation"`
	Frequency        *Text             `json:"frequency"`
}

// PriceSpecification is a generated struct representing the https://schema.org/PriceSpecification class.
type PriceSpecification struct {
	EligibleQuantity          *QuantitativeValue  `json:"eligiblequantity"`
	EligibleTransactionVolume *PriceSpecification `json:"eligibletransactionvolume"`
	MaxPrice                  *Number             `json:"maxprice"`
	MinPrice                  *Number             `json:"minprice"`
	Price                     *Text               `json:"price"`
	PriceCurrency             *Text               `json:"pricecurrency"`
	ValidFrom                 *DateTime           `json:"validfrom"`
	ValidThrough              *DateTime           `json:"validthrough"`
	ValueAddedTaxIncluded     *Boolean            `json:"valueaddedtaxincluded"`
}

// WantAction is a generated struct representing the https://schema.org/WantAction class.
type WantAction struct {
}

// CancelAction is a generated struct representing the https://schema.org/CancelAction class.
type CancelAction struct {
}

// Demand is a generated struct representing the https://schema.org/Demand class.
type Demand struct {
	AcceptedPaymentMethod     *PaymentMethod       `json:"acceptedpaymentmethod"`
	AdvanceBookingRequirement *QuantitativeValue   `json:"advancebookingrequirement"`
	Asin                      *URL                 `json:"asin"`
	Availability              *ItemAvailability    `json:"availability"`
	AvailabilityEnds          *Time                `json:"availabilityends"`
	AvailabilityStarts        *Time                `json:"availabilitystarts"`
	AvailableAtOrFrom         *Place               `json:"availableatorfrom"`
	AvailableDeliveryMethod   *DeliveryMethod      `json:"availabledeliverymethod"`
	BusinessFunction          *BusinessFunction    `json:"businessfunction"`
	DeliveryLeadTime          *QuantitativeValue   `json:"deliveryleadtime"`
	EligibleCustomerType      *BusinessEntityType  `json:"eligiblecustomertype"`
	EligibleDuration          *QuantitativeValue   `json:"eligibleduration"`
	EligibleQuantity          *QuantitativeValue   `json:"eligiblequantity"`
	EligibleRegion            *Text                `json:"eligibleregion"`
	EligibleTransactionVolume *PriceSpecification  `json:"eligibletransactionvolume"`
	Gtin12                    *Text                `json:"gtin12"`
	Gtin13                    *Text                `json:"gtin13"`
	Gtin14                    *Text                `json:"gtin14"`
	Gtin8                     *Text                `json:"gtin8"`
	IncludesObject            *TypeAndQuantityNode `json:"includesobject"`
	IneligibleRegion          *Text                `json:"ineligibleregion"`
	InventoryLevel            *QuantitativeValue   `json:"inventorylevel"`
	ItemCondition             *OfferItemCondition  `json:"itemcondition"`
	Mpn                       *Text                `json:"mpn"`
	PriceSpecification        *PriceSpecification  `json:"pricespecification"`
	Sku                       *Text                `json:"sku"`
	ValidFrom                 *DateTime            `json:"validfrom"`
	ValidThrough              *DateTime            `json:"validthrough"`
	ItemOffered               *Trip                `json:"itemoffered"`
	SerialNumber              *Text                `json:"serialnumber"`
	Warranty                  *WarrantyPromise     `json:"warranty"`
	Seller                    *Person              `json:"seller"`
	AreaServed                *Text                `json:"areaserved"`
	Gtin                      *URL                 `json:"gtin"`
}

// QuoteAction is a generated struct representing the https://schema.org/QuoteAction class.
type QuoteAction struct {
}

// SportsTeam is a generated struct representing the https://schema.org/SportsTeam class.
type SportsTeam struct {
	Athlete *Person `json:"athlete"`
	Coach   *Person `json:"coach"`
	Gender  *Text   `json:"gender"`
}

// TransferAction is a generated struct representing the https://schema.org/TransferAction class.
type TransferAction struct {
	FromLocation *Place `json:"fromlocation"`
	ToLocation   *Place `json:"tolocation"`
}

// DigitalDocument is a generated struct representing the https://schema.org/DigitalDocument class.
type DigitalDocument struct {
	HasDigitalDocumentPermission *DigitalDocumentPermission `json:"hasdigitaldocumentpermission"`
}

// MedicalImagingTechnique is a generated struct representing the https://schema.org/MedicalImagingTechnique class.
type MedicalImagingTechnique struct {
}

// AddAction is a generated struct representing the https://schema.org/AddAction class.
type AddAction struct {
}

// HousePainter is a generated struct representing the https://schema.org/HousePainter class.
type HousePainter struct {
}

// HowTo is a generated struct representing the https://schema.org/HowTo class.
type HowTo struct {
	EstimatedCost *Text     `json:"estimatedcost"`
	PrepTime      *Duration `json:"preptime"`
	Steps         *Text     `json:"steps"`
	Tool          *Text     `json:"tool"`
	TotalTime     *Duration `json:"totaltime"`
	PerformTime   *Duration `json:"performtime"`
	Yield         *Text     `json:"yield"`
	Step          *Text     `json:"step"`
	Supply        *Text     `json:"supply"`
}

// LendAction is a generated struct representing the https://schema.org/LendAction class.
type LendAction struct {
	Borrower *Person `json:"borrower"`
}

// RsvpAction is a generated struct representing the https://schema.org/RsvpAction class.
type RsvpAction struct {
	AdditionalNumberOfGuests *Number           `json:"additionalnumberofguests"`
	Comment                  *Comment          `json:"comment"`
	RsvpResponse             *RsvpResponseType `json:"rsvpresponse"`
}

// BoardingPolicyType is a generated struct representing the https://schema.org/BoardingPolicyType class.
type BoardingPolicyType struct {
}

// FAQPage is a generated struct representing the https://schema.org/FAQPage class.
type FAQPage struct {
}

// TakeAction is a generated struct representing the https://schema.org/TakeAction class.
type TakeAction struct {
}

// MotorcycleRepair is a generated struct representing the https://schema.org/MotorcycleRepair class.
type MotorcycleRepair struct {
}

// Recommendation is a generated struct representing the https://schema.org/Recommendation class.
type Recommendation struct {
	Category *URL `json:"category"`
}

// Airline is a generated struct representing the https://schema.org/Airline class.
type Airline struct {
	BoardingPolicy *BoardingPolicyType `json:"boardingpolicy"`
	IataCode       *Text               `json:"iatacode"`
}

// ComicStory is a generated struct representing the https://schema.org/ComicStory class.
type ComicStory struct {
	Artist   *Person `json:"artist"`
	Colorist *Person `json:"colorist"`
	Inker    *Person `json:"inker"`
	Letterer *Person `json:"letterer"`
	Penciler *Person `json:"penciler"`
}

// CurrencyConversionService is a generated struct representing the https://schema.org/CurrencyConversionService class.
type CurrencyConversionService struct {
}

// EducationEvent is a generated struct representing the https://schema.org/EducationEvent class.
type EducationEvent struct {
	Assesses         *Text `json:"assesses"`
	EducationalLevel *URL  `json:"educationallevel"`
	Teaches          *Text `json:"teaches"`
}

// ExhibitionEvent is a generated struct representing the https://schema.org/ExhibitionEvent class.
type ExhibitionEvent struct {
}

// PreventionIndication is a generated struct representing the https://schema.org/PreventionIndication class.
type PreventionIndication struct {
}

// Thing is a generated struct representing the https://schema.org/Thing class.
type Thing struct {
	AdditionalType            *URL        `json:"additionaltype"`
	DisambiguatingDescription *Text       `json:"disambiguatingdescription"`
	Name                      *Text       `json:"name"`
	PotentialAction           *Action     `json:"potentialaction"`
	SameAs                    *URL        `json:"sameas"`
	Url                       *URL        `json:"url"`
	AlternateName             *Text       `json:"alternatename"`
	MainEntityOfPage          *URL        `json:"mainentityofpage"`
	SubjectOf                 *Event      `json:"subjectof"`
	Image                     *URL        `json:"image"`
	Description               *TextObject `json:"description"`
	Identifier                *URL        `json:"identifier"`
}

// AutoWash is a generated struct representing the https://schema.org/AutoWash class.
type AutoWash struct {
}

// BankOrCreditUnion is a generated struct representing the https://schema.org/BankOrCreditUnion class.
type BankOrCreditUnion struct {
}

// InstallAction is a generated struct representing the https://schema.org/InstallAction class.
type InstallAction struct {
}

// MediaReview is a generated struct representing the https://schema.org/MediaReview class.
type MediaReview struct {
	MediaAuthenticityCategory       *MediaManipulationRatingEnumeration `json:"mediaauthenticitycategory"`
	OriginalMediaContextDescription *Text                               `json:"originalmediacontextdescription"`
	OriginalMediaLink               *WebPage                            `json:"originalmedialink"`
}

// MotorizedBicycle is a generated struct representing the https://schema.org/MotorizedBicycle class.
type MotorizedBicycle struct {
}

// PayAction is a generated struct representing the https://schema.org/PayAction class.
type PayAction struct {
	Recipient *Person `json:"recipient"`
}

// SkiResort is a generated struct representing the https://schema.org/SkiResort class.
type SkiResort struct {
}

// ViewAction is a generated struct representing the https://schema.org/ViewAction class.
type ViewAction struct {
}

// WebSite is a generated struct representing the https://schema.org/WebSite class.
type WebSite struct {
	Issn *Text `json:"issn"`
}

// ReviewAction is a generated struct representing the https://schema.org/ReviewAction class.
type ReviewAction struct {
	ResultReview *Review `json:"resultreview"`
}

// Corporation is a generated struct representing the https://schema.org/Corporation class.
type Corporation struct {
	TickerSymbol *Text `json:"tickersymbol"`
}

// MedicalEnumeration is a generated struct representing the https://schema.org/MedicalEnumeration class.
type MedicalEnumeration struct {
}

// OrderItem is a generated struct representing the https://schema.org/OrderItem class.
type OrderItem struct {
	OrderDelivery   *ParcelDelivery `json:"orderdelivery"`
	OrderItemNumber *Text           `json:"orderitemnumber"`
	OrderItemStatus *OrderStatus    `json:"orderitemstatus"`
	OrderQuantity   *Number         `json:"orderquantity"`
	OrderedItem     *Service        `json:"ordereditem"`
}

// Claim is a generated struct representing the https://schema.org/Claim class.
type Claim struct {
	Appearance       *CreativeWork `json:"appearance"`
	ClaimInterpreter *Person       `json:"claiminterpreter"`
	FirstAppearance  *CreativeWork `json:"firstappearance"`
}

// ComicCoverArt is a generated struct representing the https://schema.org/ComicCoverArt class.
type ComicCoverArt struct {
}

// MedicalRiskCalculator is a generated struct representing the https://schema.org/MedicalRiskCalculator class.
type MedicalRiskCalculator struct {
}

// PodcastEpisode is a generated struct representing the https://schema.org/PodcastEpisode class.
type PodcastEpisode struct {
}

// SearchRescueOrganization is a generated struct representing the https://schema.org/SearchRescueOrganization class.
type SearchRescueOrganization struct {
}

// Text is a generated struct representing the https://schema.org/Text class.
type Text struct {
}

// DataFeed is a generated struct representing the https://schema.org/DataFeed class.
type DataFeed struct {
	DataFeedElement *Thing `json:"datafeedelement"`
}

// Distance is a generated struct representing the https://schema.org/Distance class.
type Distance struct {
}

// HowToStep is a generated struct representing the https://schema.org/HowToStep class.
type HowToStep struct {
}

// MedicalContraindication is a generated struct representing the https://schema.org/MedicalContraindication class.
type MedicalContraindication struct {
}

// Project is a generated struct representing the https://schema.org/Project class.
type Project struct {
}

// Store is a generated struct representing the https://schema.org/Store class.
type Store struct {
}

// AggregateRating is a generated struct representing the https://schema.org/AggregateRating class.
type AggregateRating struct {
	ItemReviewed *Thing   `json:"itemreviewed"`
	RatingCount  *Integer `json:"ratingcount"`
	ReviewCount  *Integer `json:"reviewcount"`
}

// BikeStore is a generated struct representing the https://schema.org/BikeStore class.
type BikeStore struct {
}

// BookFormatType is a generated struct representing the https://schema.org/BookFormatType class.
type BookFormatType struct {
}

// SatiricalArticle is a generated struct representing the https://schema.org/SatiricalArticle class.
type SatiricalArticle struct {
}

// UserDownloads is a generated struct representing the https://schema.org/UserDownloads class.
type UserDownloads struct {
}

// AnatomicalSystem is a generated struct representing the https://schema.org/AnatomicalSystem class.
type AnatomicalSystem struct {
	AssociatedPathophysiology *Text                `json:"associatedpathophysiology"`
	ComprisedOf               *AnatomicalSystem    `json:"comprisedof"`
	RelatedCondition          *MedicalCondition    `json:"relatedcondition"`
	RelatedStructure          *AnatomicalStructure `json:"relatedstructure"`
	RelatedTherapy            *MedicalTherapy      `json:"relatedtherapy"`
}

// PropertyValueSpecification is a generated struct representing the https://schema.org/PropertyValueSpecification class.
type PropertyValueSpecification struct {
	DefaultValue   *Thing   `json:"defaultvalue"`
	MaxValue       *Number  `json:"maxvalue"`
	MinValue       *Number  `json:"minvalue"`
	MultipleValues *Boolean `json:"multiplevalues"`
	ReadonlyValue  *Boolean `json:"readonlyvalue"`
	StepValue      *Number  `json:"stepvalue"`
	ValueMaxLength *Number  `json:"valuemaxlength"`
	ValueMinLength *Number  `json:"valueminlength"`
	ValueName      *Text    `json:"valuename"`
	ValuePattern   *Text    `json:"valuepattern"`
	ValueRequired  *Boolean `json:"valuerequired"`
}

// MonetaryAmount is a generated struct representing the https://schema.org/MonetaryAmount class.
type MonetaryAmount struct {
	Currency     *Text     `json:"currency"`
	MaxValue     *Number   `json:"maxvalue"`
	MinValue     *Number   `json:"minvalue"`
	ValidFrom    *DateTime `json:"validfrom"`
	ValidThrough *DateTime `json:"validthrough"`
	Value        *Text     `json:"value"`
}

// LegalService is a generated struct representing the https://schema.org/LegalService class.
type LegalService struct {
}

// MedicalGuidelineRecommendation is a generated struct representing the https://schema.org/MedicalGuidelineRecommendation class.
type MedicalGuidelineRecommendation struct {
	RecommendationStrength *Text `json:"recommendationstrength"`
}

// VeterinaryCare is a generated struct representing the https://schema.org/VeterinaryCare class.
type VeterinaryCare struct {
}

// AmpStory is a generated struct representing the https://schema.org/AmpStory class.
type AmpStory struct {
}

// Attorney is a generated struct representing the https://schema.org/Attorney class.
type Attorney struct {
}

// MensClothingStore is a generated struct representing the https://schema.org/MensClothingStore class.
type MensClothingStore struct {
}

// OutletStore is a generated struct representing the https://schema.org/OutletStore class.
type OutletStore struct {
}

// SubwayStation is a generated struct representing the https://schema.org/SubwayStation class.
type SubwayStation struct {
}

// ConfirmAction is a generated struct representing the https://schema.org/ConfirmAction class.
type ConfirmAction struct {
}

// FundingAgency is a generated struct representing the https://schema.org/FundingAgency class.
type FundingAgency struct {
}

// ItemList is a generated struct representing the https://schema.org/ItemList class.
type ItemList struct {
	ItemListElement *Thing   `json:"itemlistelement"`
	ItemListOrder   *Text    `json:"itemlistorder"`
	NumberOfItems   *Integer `json:"numberofitems"`
}

// MedicalConditionStage is a generated struct representing the https://schema.org/MedicalConditionStage class.
type MedicalConditionStage struct {
	StageAsNumber  *Number `json:"stageasnumber"`
	SubStageSuffix *Text   `json:"substagesuffix"`
}

// EatAction is a generated struct representing the https://schema.org/EatAction class.
type EatAction struct {
}

// Integer is a generated struct representing the https://schema.org/Integer class.
type Integer struct {
}

// InvestmentFund is a generated struct representing the https://schema.org/InvestmentFund class.
type InvestmentFund struct {
}

// MediaGallery is a generated struct representing the https://schema.org/MediaGallery class.
type MediaGallery struct {
}

// Nerve is a generated struct representing the https://schema.org/Nerve class.
type Nerve struct {
	Branch      *AnatomicalStructure `json:"branch"`
	NerveMotor  *Muscle              `json:"nervemotor"`
	SensoryUnit *SuperficialAnatomy  `json:"sensoryunit"`
	SourcedFrom *BrainStructure      `json:"sourcedfrom"`
}

// ReceiveAction is a generated struct representing the https://schema.org/ReceiveAction class.
type ReceiveAction struct {
	DeliveryMethod *DeliveryMethod `json:"deliverymethod"`
	Sender         *Person         `json:"sender"`
}

// ScholarlyArticle is a generated struct representing the https://schema.org/ScholarlyArticle class.
type ScholarlyArticle struct {
}

// DepartmentStore is a generated struct representing the https://schema.org/DepartmentStore class.
type DepartmentStore struct {
}

// MedicalEvidenceLevel is a generated struct representing the https://schema.org/MedicalEvidenceLevel class.
type MedicalEvidenceLevel struct {
}

// PriceComponentTypeEnumeration is a generated struct representing the https://schema.org/PriceComponentTypeEnumeration class.
type PriceComponentTypeEnumeration struct {
}

// VideoGallery is a generated struct representing the https://schema.org/VideoGallery class.
type VideoGallery struct {
}

// AutoDealer is a generated struct representing the https://schema.org/AutoDealer class.
type AutoDealer struct {
}

// PublicSwimmingPool is a generated struct representing the https://schema.org/PublicSwimmingPool class.
type PublicSwimmingPool struct {
}

// TVClip is a generated struct representing the https://schema.org/TVClip class.
type TVClip struct {
	PartOfTVSeries *TVSeries `json:"partoftvseries"`
}

// WarrantyScope is a generated struct representing the https://schema.org/WarrantyScope class.
type WarrantyScope struct {
}

// MedicalSignOrSymptom is a generated struct representing the https://schema.org/MedicalSignOrSymptom class.
type MedicalSignOrSymptom struct {
	PossibleTreatment *MedicalTherapy `json:"possibletreatment"`
}

// MusicGroup is a generated struct representing the https://schema.org/MusicGroup class.
type MusicGroup struct {
	Albums           *MusicAlbum     `json:"albums"`
	MusicGroupMember *Person         `json:"musicgroupmember"`
	Tracks           *MusicRecording `json:"tracks"`
	Album            *MusicAlbum     `json:"album"`
	Genre            *URL            `json:"genre"`
	Track            *MusicRecording `json:"track"`
}

// Permit is a generated struct representing the https://schema.org/Permit class.
type Permit struct {
	IssuedBy       *Organization       `json:"issuedby"`
	IssuedThrough  *Service            `json:"issuedthrough"`
	PermitAudience *Audience           `json:"permitaudience"`
	ValidFor       *Duration           `json:"validfor"`
	ValidFrom      *DateTime           `json:"validfrom"`
	ValidIn        *AdministrativeArea `json:"validin"`
	ValidUntil     *Date               `json:"validuntil"`
}

// SportsActivityLocation is a generated struct representing the https://schema.org/SportsActivityLocation class.
type SportsActivityLocation struct {
}

// UserPageVisits is a generated struct representing the https://schema.org/UserPageVisits class.
type UserPageVisits struct {
}

// Church is a generated struct representing the https://schema.org/Church class.
type Church struct {
}

// Pharmacy is a generated struct representing the https://schema.org/Pharmacy class.
type Pharmacy struct {
}

// ProgramMembership is a generated struct representing the https://schema.org/ProgramMembership class.
type ProgramMembership struct {
	HostingOrganization    *Organization      `json:"hostingorganization"`
	Members                *Person            `json:"members"`
	MembershipNumber       *Text              `json:"membershipnumber"`
	MembershipPointsEarned *QuantitativeValue `json:"membershippointsearned"`
	ProgramName            *Text              `json:"programname"`
	Member                 *Person            `json:"member"`
}

// TaxiStand is a generated struct representing the https://schema.org/TaxiStand class.
type TaxiStand struct {
}

// TypeAndQuantityNode is a generated struct representing the https://schema.org/TypeAndQuantityNode class.
type TypeAndQuantityNode struct {
	AmountOfThisGood *Number           `json:"amountofthisgood"`
	BusinessFunction *BusinessFunction `json:"businessfunction"`
	TypeOfGood       *Service          `json:"typeofgood"`
	UnitCode         *URL              `json:"unitcode"`
	UnitText         *Text             `json:"unittext"`
}

// GatedResidenceCommunity is a generated struct representing the https://schema.org/GatedResidenceCommunity class.
type GatedResidenceCommunity struct {
}

// LiveBlogPosting is a generated struct representing the https://schema.org/LiveBlogPosting class.
type LiveBlogPosting struct {
	CoverageEndTime   *DateTime    `json:"coverageendtime"`
	CoverageStartTime *DateTime    `json:"coveragestarttime"`
	LiveBlogUpdate    *BlogPosting `json:"liveblogupdate"`
}

// Periodical is a generated struct representing the https://schema.org/Periodical class.
type Periodical struct {
}

// RadiationTherapy is a generated struct representing the https://schema.org/RadiationTherapy class.
type RadiationTherapy struct {
}

// Canal is a generated struct representing the https://schema.org/Canal class.
type Canal struct {
}

// MedicalBusiness is a generated struct representing the https://schema.org/MedicalBusiness class.
type MedicalBusiness struct {
}

// NewsArticle is a generated struct representing the https://schema.org/NewsArticle class.
type NewsArticle struct {
	Dateline     *Text `json:"dateline"`
	PrintColumn  *Text `json:"printcolumn"`
	PrintEdition *Text `json:"printedition"`
	PrintPage    *Text `json:"printpage"`
	PrintSection *Text `json:"printsection"`
}

// OfferForLease is a generated struct representing the https://schema.org/OfferForLease class.
type OfferForLease struct {
}

// BusinessEvent is a generated struct representing the https://schema.org/BusinessEvent class.
type BusinessEvent struct {
}

// ItemListOrderType is a generated struct representing the https://schema.org/ItemListOrderType class.
type ItemListOrderType struct {
}

// Photograph is a generated struct representing the https://schema.org/Photograph class.
type Photograph struct {
}

// RentAction is a generated struct representing the https://schema.org/RentAction class.
type RentAction struct {
	Landlord        *Person          `json:"landlord"`
	RealEstateAgent *RealEstateAgent `json:"realestateagent"`
}

// GameAvailabilityEnumeration is a generated struct representing the https://schema.org/GameAvailabilityEnumeration class.
type GameAvailabilityEnumeration struct {
}

// ImagingTest is a generated struct representing the https://schema.org/ImagingTest class.
type ImagingTest struct {
	ImagingTechnique *MedicalImagingTechnique `json:"imagingtechnique"`
}

// MedicalEntity is a generated struct representing the https://schema.org/MedicalEntity class.
type MedicalEntity struct {
	Code                 *MedicalCode      `json:"code"`
	Guideline            *MedicalGuideline `json:"guideline"`
	LegalStatus          *Text             `json:"legalstatus"`
	MedicineSystem       *MedicineSystem   `json:"medicinesystem"`
	RecognizingAuthority *Organization     `json:"recognizingauthority"`
	RelevantSpecialty    *MedicalSpecialty `json:"relevantspecialty"`
	Study                *MedicalStudy     `json:"study"`
	Funding              *Grant            `json:"funding"`
}

// OrderAction is a generated struct representing the https://schema.org/OrderAction class.
type OrderAction struct {
	DeliveryMethod *DeliveryMethod `json:"deliverymethod"`
}

// UserComments is a generated struct representing the https://schema.org/UserComments class.
type UserComments struct {
	CommentText *Text         `json:"commenttext"`
	CommentTime *DateTime     `json:"commenttime"`
	Discusses   *CreativeWork `json:"discusses"`
	ReplyToUrl  *URL          `json:"replytourl"`
	Creator     *Person       `json:"creator"`
}

// Apartment is a generated struct representing the https://schema.org/Apartment class.
type Apartment struct {
	NumberOfRooms *QuantitativeValue `json:"numberofrooms"`
	Occupancy     *QuantitativeValue `json:"occupancy"`
}

// BedType is a generated struct representing the https://schema.org/BedType class.
type BedType struct {
}

// Mountain is a generated struct representing the https://schema.org/Mountain class.
type Mountain struct {
}

// StatusEnumeration is a generated struct representing the https://schema.org/StatusEnumeration class.
type StatusEnumeration struct {
}

// TVSeason is a generated struct representing the https://schema.org/TVSeason class.
type TVSeason struct {
	CountryOfOrigin *Country  `json:"countryoforigin"`
	PartOfTVSeries  *TVSeries `json:"partoftvseries"`
	TitleEIDR       *URL      `json:"titleeidr"`
}

// VideoGameSeries is a generated struct representing the https://schema.org/VideoGameSeries class.
type VideoGameSeries struct {
	Actors             *Person             `json:"actors"`
	CharacterAttribute *Thing              `json:"characterattribute"`
	CheatCode          *CreativeWork       `json:"cheatcode"`
	Directors          *Person             `json:"directors"`
	Episodes           *Episode            `json:"episodes"`
	GameItem           *Thing              `json:"gameitem"`
	GameLocation       *URL                `json:"gamelocation"`
	GamePlatform       *URL                `json:"gameplatform"`
	MusicBy            *Person             `json:"musicby"`
	NumberOfEpisodes   *Integer            `json:"numberofepisodes"`
	NumberOfPlayers    *QuantitativeValue  `json:"numberofplayers"`
	NumberOfSeasons    *Integer            `json:"numberofseasons"`
	PlayMode           *GamePlayMode       `json:"playmode"`
	ProductionCompany  *Organization       `json:"productioncompany"`
	Quest              *Thing              `json:"quest"`
	Seasons            *CreativeWorkSeason `json:"seasons"`
	Trailer            *VideoObject        `json:"trailer"`
	ContainsSeason     *CreativeWorkSeason `json:"containsseason"`
	Director           *Person             `json:"director"`
	Episode            *Episode            `json:"episode"`
	Season             *URL                `json:"season"`
	Actor              *Person             `json:"actor"`
}

// WebContent is a generated struct representing the https://schema.org/WebContent class.
type WebContent struct {
}

// Clip is a generated struct representing the https://schema.org/Clip class.
type Clip struct {
	Actors        *Person             `json:"actors"`
	ClipNumber    *Text               `json:"clipnumber"`
	Directors     *Person             `json:"directors"`
	EndOffset     *Number             `json:"endoffset"`
	MusicBy       *Person             `json:"musicby"`
	PartOfEpisode *Episode            `json:"partofepisode"`
	PartOfSeason  *CreativeWorkSeason `json:"partofseason"`
	StartOffset   *Number             `json:"startoffset"`
	Director      *Person             `json:"director"`
	PartOfSeries  *CreativeWorkSeries `json:"partofseries"`
	Actor         *Person             `json:"actor"`
}

// QAPage is a generated struct representing the https://schema.org/QAPage class.
type QAPage struct {
}

// ContactPoint is a generated struct representing the https://schema.org/ContactPoint class.
type ContactPoint struct {
	AvailableLanguage *Text                      `json:"availablelanguage"`
	ContactOption     *ContactPointOption        `json:"contactoption"`
	ContactType       *Text                      `json:"contacttype"`
	Email             *Text                      `json:"email"`
	FaxNumber         *Text                      `json:"faxnumber"`
	HoursAvailable    *OpeningHoursSpecification `json:"hoursavailable"`
	ProductSupported  *Text                      `json:"productsupported"`
	Telephone         *Text                      `json:"telephone"`
	ServiceArea       *Place                     `json:"servicearea"`
	AreaServed        *Text                      `json:"areaserved"`
}

// MedicalTrialDesign is a generated struct representing the https://schema.org/MedicalTrialDesign class.
type MedicalTrialDesign struct {
}

// MusicStore is a generated struct representing the https://schema.org/MusicStore class.
type MusicStore struct {
}

// PlanAction is a generated struct representing the https://schema.org/PlanAction class.
type PlanAction struct {
	ScheduledTime *DateTime `json:"scheduledtime"`
}

// ProductModel is a generated struct representing the https://schema.org/ProductModel class.
type ProductModel struct {
	PredecessorOf *ProductModel `json:"predecessorof"`
	SuccessorOf   *ProductModel `json:"successorof"`
	IsVariantOf   *ProductModel `json:"isvariantof"`
}

// ReservationStatusType is a generated struct representing the https://schema.org/ReservationStatusType class.
type ReservationStatusType struct {
}

// MovieSeries is a generated struct representing the https://schema.org/MovieSeries class.
type MovieSeries struct {
	Actors            *Person       `json:"actors"`
	Directors         *Person       `json:"directors"`
	MusicBy           *Person       `json:"musicby"`
	ProductionCompany *Organization `json:"productioncompany"`
	Trailer           *VideoObject  `json:"trailer"`
	Director          *Person       `json:"director"`
	Actor             *Person       `json:"actor"`
}

// ParcelDelivery is a generated struct representing the https://schema.org/ParcelDelivery class.
type ParcelDelivery struct {
	Carrier              *Organization   `json:"carrier"`
	DeliveryAddress      *PostalAddress  `json:"deliveryaddress"`
	DeliveryStatus       *DeliveryEvent  `json:"deliverystatus"`
	ExpectedArrivalFrom  *DateTime       `json:"expectedarrivalfrom"`
	ExpectedArrivalUntil *DateTime       `json:"expectedarrivaluntil"`
	HasDeliveryMethod    *DeliveryMethod `json:"hasdeliverymethod"`
	ItemShipped          *Product        `json:"itemshipped"`
	OriginAddress        *PostalAddress  `json:"originaddress"`
	PartOfOrder          *Order          `json:"partoforder"`
	TrackingNumber       *Text           `json:"trackingnumber"`
	TrackingUrl          *URL            `json:"trackingurl"`
	Provider             *Person         `json:"provider"`
}

// RsvpResponseType is a generated struct representing the https://schema.org/RsvpResponseType class.
type RsvpResponseType struct {
}

// UserCheckins is a generated struct representing the https://schema.org/UserCheckins class.
type UserCheckins struct {
}

// DiscoverAction is a generated struct representing the https://schema.org/DiscoverAction class.
type DiscoverAction struct {
}

// Language is a generated struct representing the https://schema.org/Language class.
type Language struct {
}

// PetStore is a generated struct representing the https://schema.org/PetStore class.
type PetStore struct {
}

// SpecialAnnouncement is a generated struct representing the https://schema.org/SpecialAnnouncement class.
type SpecialAnnouncement struct {
	AnnouncementLocation        *LocalBusiness     `json:"announcementlocation"`
	DatePosted                  *DateTime          `json:"dateposted"`
	DiseasePreventionInfo       *WebContent        `json:"diseasepreventioninfo"`
	DiseaseSpreadStatistics     *WebContent        `json:"diseasespreadstatistics"`
	GettingTestedInfo           *WebContent        `json:"gettingtestedinfo"`
	GovernmentBenefitsInfo      *GovernmentService `json:"governmentbenefitsinfo"`
	NewsUpdatesAndGuidelines    *WebContent        `json:"newsupdatesandguidelines"`
	PublicTransportClosuresInfo *WebContent        `json:"publictransportclosuresinfo"`
	QuarantineGuidelines        *WebContent        `json:"quarantineguidelines"`
	SchoolClosuresInfo          *WebContent        `json:"schoolclosuresinfo"`
	TravelBans                  *WebContent        `json:"travelbans"`
	WebFeed                     *URL               `json:"webfeed"`
	Category                    *URL               `json:"category"`
}

// TravelAction is a generated struct representing the https://schema.org/TravelAction class.
type TravelAction struct {
	Distance *Distance `json:"distance"`
}

// AutomotiveBusiness is a generated struct representing the https://schema.org/AutomotiveBusiness class.
type AutomotiveBusiness struct {
}

// LegislativeBuilding is a generated struct representing the https://schema.org/LegislativeBuilding class.
type LegislativeBuilding struct {
}

// CheckInAction is a generated struct representing the https://schema.org/CheckInAction class.
type CheckInAction struct {
}

// CompoundPriceSpecification is a generated struct representing the https://schema.org/CompoundPriceSpecification class.
type CompoundPriceSpecification struct {
	PriceComponent *UnitPriceSpecification `json:"pricecomponent"`
	PriceType      *Text                   `json:"pricetype"`
}

// GovernmentPermit is a generated struct representing the https://schema.org/GovernmentPermit class.
type GovernmentPermit struct {
}

// Researcher is a generated struct representing the https://schema.org/Researcher class.
type Researcher struct {
}

// TrainReservation is a generated struct representing the https://schema.org/TrainReservation class.
type TrainReservation struct {
}

// MusicReleaseFormatType is a generated struct representing the https://schema.org/MusicReleaseFormatType class.
type MusicReleaseFormatType struct {
}

// RVPark is a generated struct representing the https://schema.org/RVPark class.
type RVPark struct {
}

// StructuredValue is a generated struct representing the https://schema.org/StructuredValue class.
type StructuredValue struct {
}

// ComputerLanguage is a generated struct representing the https://schema.org/ComputerLanguage class.
type ComputerLanguage struct {
}

// DaySpa is a generated struct representing the https://schema.org/DaySpa class.
type DaySpa struct {
}

// DigitalDocumentPermissionType is a generated struct representing the https://schema.org/DigitalDocumentPermissionType class.
type DigitalDocumentPermissionType struct {
}

// City is a generated struct representing the https://schema.org/City class.
type City struct {
}

// ComicIssue is a generated struct representing the https://schema.org/ComicIssue class.
type ComicIssue struct {
	Artist       *Person `json:"artist"`
	Colorist     *Person `json:"colorist"`
	Inker        *Person `json:"inker"`
	Letterer     *Person `json:"letterer"`
	Penciler     *Person `json:"penciler"`
	VariantCover *Text   `json:"variantcover"`
}

// MusicVideoObject is a generated struct representing the https://schema.org/MusicVideoObject class.
type MusicVideoObject struct {
}

// ReplaceAction is a generated struct representing the https://schema.org/ReplaceAction class.
type ReplaceAction struct {
	Replacee *Thing `json:"replacee"`
	Replacer *Thing `json:"replacer"`
}

// InformAction is a generated struct representing the https://schema.org/InformAction class.
type InformAction struct {
	Event *Event `json:"event"`
}

// CatholicChurch is a generated struct representing the https://schema.org/CatholicChurch class.
type CatholicChurch struct {
}

// DepositAccount is a generated struct representing the https://schema.org/DepositAccount class.
type DepositAccount struct {
}

// MusicRelease is a generated struct representing the https://schema.org/MusicRelease class.
type MusicRelease struct {
	CatalogNumber      *Text                   `json:"catalognumber"`
	CreditedTo         *Person                 `json:"creditedto"`
	MusicReleaseFormat *MusicReleaseFormatType `json:"musicreleaseformat"`
	RecordLabel        *Organization           `json:"recordlabel"`
	Duration           *Duration               `json:"duration"`
	ReleaseOf          *MusicAlbum             `json:"releaseof"`
}

// PriceTypeEnumeration is a generated struct representing the https://schema.org/PriceTypeEnumeration class.
type PriceTypeEnumeration struct {
}

// SizeSystemEnumeration is a generated struct representing the https://schema.org/SizeSystemEnumeration class.
type SizeSystemEnumeration struct {
}

// OrderStatus is a generated struct representing the https://schema.org/OrderStatus class.
type OrderStatus struct {
}

// PhysicalTherapy is a generated struct representing the https://schema.org/PhysicalTherapy class.
type PhysicalTherapy struct {
}

// RadioChannel is a generated struct representing the https://schema.org/RadioChannel class.
type RadioChannel struct {
}

// ReserveAction is a generated struct representing the https://schema.org/ReserveAction class.
type ReserveAction struct {
}

// AmusementPark is a generated struct representing the https://schema.org/AmusementPark class.
type AmusementPark struct {
}

// CheckOutAction is a generated struct representing the https://schema.org/CheckOutAction class.
type CheckOutAction struct {
}

// CommentAction is a generated struct representing the https://schema.org/CommentAction class.
type CommentAction struct {
	ResultComment *Comment `json:"resultcomment"`
}

// DDxElement is a generated struct representing the https://schema.org/DDxElement class.
type DDxElement struct {
	Diagnosis          *MedicalCondition     `json:"diagnosis"`
	DistinguishingSign *MedicalSignOrSymptom `json:"distinguishingsign"`
}

// AdultEntertainment is a generated struct representing the https://schema.org/AdultEntertainment class.
type AdultEntertainment struct {
}

// Airport is a generated struct representing the https://schema.org/Airport class.
type Airport struct {
	IataCode *Text `json:"iatacode"`
	IcaoCode *Text `json:"icaocode"`
}

// Question is a generated struct representing the https://schema.org/Question class.
type Question struct {
	AcceptedAnswer  *ItemList     `json:"acceptedanswer"`
	AnswerCount     *Integer      `json:"answercount"`
	EduQuestionType *Text         `json:"eduquestiontype"`
	ParentItem      *CreativeWork `json:"parentitem"`
	SuggestedAnswer *ItemList     `json:"suggestedanswer"`
}

// ReplyAction is a generated struct representing the https://schema.org/ReplyAction class.
type ReplyAction struct {
	ResultComment *Comment `json:"resultcomment"`
}

// WPAdBlock is a generated struct representing the https://schema.org/WPAdBlock class.
type WPAdBlock struct {
}

// BusStop is a generated struct representing the https://schema.org/BusStop class.
type BusStop struct {
}

// DataType is a generated struct representing the https://schema.org/DataType class.
type DataType struct {
}

// EnergyStarEnergyEfficiencyEnumeration is a generated struct representing the https://schema.org/EnergyStarEnergyEfficiencyEnumeration class.
type EnergyStarEnergyEfficiencyEnumeration struct {
}

// Volcano is a generated struct representing the https://schema.org/Volcano class.
type Volcano struct {
}

// Offer is a generated struct representing the https://schema.org/Offer class.
type Offer struct {
	AcceptedPaymentMethod     *PaymentMethod            `json:"acceptedpaymentmethod"`
	AddOn                     *Offer                    `json:"addon"`
	AdvanceBookingRequirement *QuantitativeValue        `json:"advancebookingrequirement"`
	AggregateRating           *AggregateRating          `json:"aggregaterating"`
	Asin                      *URL                      `json:"asin"`
	Availability              *ItemAvailability         `json:"availability"`
	AvailabilityEnds          *Time                     `json:"availabilityends"`
	AvailabilityStarts        *Time                     `json:"availabilitystarts"`
	AvailableAtOrFrom         *Place                    `json:"availableatorfrom"`
	AvailableDeliveryMethod   *DeliveryMethod           `json:"availabledeliverymethod"`
	BusinessFunction          *BusinessFunction         `json:"businessfunction"`
	CheckoutPageURLTemplate   *Text                     `json:"checkoutpageurltemplate"`
	DeliveryLeadTime          *QuantitativeValue        `json:"deliveryleadtime"`
	EligibleCustomerType      *BusinessEntityType       `json:"eligiblecustomertype"`
	EligibleDuration          *QuantitativeValue        `json:"eligibleduration"`
	EligibleQuantity          *QuantitativeValue        `json:"eligiblequantity"`
	EligibleRegion            *Text                     `json:"eligibleregion"`
	EligibleTransactionVolume *PriceSpecification       `json:"eligibletransactionvolume"`
	Gtin12                    *Text                     `json:"gtin12"`
	Gtin13                    *Text                     `json:"gtin13"`
	Gtin14                    *Text                     `json:"gtin14"`
	Gtin8                     *Text                     `json:"gtin8"`
	HasAdultConsideration     *AdultOrientedEnumeration `json:"hasadultconsideration"`
	HasMeasurement            *QuantitativeValue        `json:"hasmeasurement"`
	HasMerchantReturnPolicy   *MerchantReturnPolicy     `json:"hasmerchantreturnpolicy"`
	IncludesObject            *TypeAndQuantityNode      `json:"includesobject"`
	IneligibleRegion          *Text                     `json:"ineligibleregion"`
	InventoryLevel            *QuantitativeValue        `json:"inventorylevel"`
	IsFamilyFriendly          *Boolean                  `json:"isfamilyfriendly"`
	ItemCondition             *OfferItemCondition       `json:"itemcondition"`
	LeaseLength               *QuantitativeValue        `json:"leaselength"`
	MobileUrl                 *Text                     `json:"mobileurl"`
	Mpn                       *Text                     `json:"mpn"`
	Price                     *Text                     `json:"price"`
	PriceCurrency             *Text                     `json:"pricecurrency"`
	PriceSpecification        *PriceSpecification       `json:"pricespecification"`
	PriceValidUntil           *Date                     `json:"pricevaliduntil"`
	Reviews                   *Review                   `json:"reviews"`
	ShippingDetails           *OfferShippingDetails     `json:"shippingdetails"`
	Sku                       *Text                     `json:"sku"`
	ValidFrom                 *DateTime                 `json:"validfrom"`
	ValidThrough              *DateTime                 `json:"validthrough"`
	Category                  *URL                      `json:"category"`
	ItemOffered               *Trip                     `json:"itemoffered"`
	OfferedBy                 *Person                   `json:"offeredby"`
	Review                    *Review                   `json:"review"`
	SerialNumber              *Text                     `json:"serialnumber"`
	Warranty                  *WarrantyPromise          `json:"warranty"`
	Seller                    *Person                   `json:"seller"`
	AreaServed                *Text                     `json:"areaserved"`
	Gtin                      *URL                      `json:"gtin"`
}

// WPFooter is a generated struct representing the https://schema.org/WPFooter class.
type WPFooter struct {
}

// Book is a generated struct representing the https://schema.org/Book class.
type Book struct {
	Abridged      *Boolean        `json:"abridged"`
	BookEdition   *Text           `json:"bookedition"`
	BookFormat    *BookFormatType `json:"bookformat"`
	Illustrator   *Person         `json:"illustrator"`
	Isbn          *Text           `json:"isbn"`
	NumberOfPages *Integer        `json:"numberofpages"`
}

// Class is a generated struct representing the https://schema.org/Class class.
type Class struct {
	SupersededBy *Property `json:"supersededby"`
}

// ElectronicsStore is a generated struct representing the https://schema.org/ElectronicsStore class.
type ElectronicsStore struct {
}

// HealthAspectEnumeration is a generated struct representing the https://schema.org/HealthAspectEnumeration class.
type HealthAspectEnumeration struct {
}

// Person is a generated struct representing the https://schema.org/Person class.
type Person struct {
	AdditionalName            *Text                              `json:"additionalname"`
	Address                   *Text                              `json:"address"`
	Affiliation               *Organization                      `json:"affiliation"`
	AgentInteractionStatistic *InteractionCounter                `json:"agentinteractionstatistic"`
	Awards                    *Text                              `json:"awards"`
	BirthDate                 *Date                              `json:"birthdate"`
	BirthPlace                *Place                             `json:"birthplace"`
	Brand                     *Organization                      `json:"brand"`
	CallSign                  *Text                              `json:"callsign"`
	Children                  *Person                            `json:"children"`
	Colleagues                *Person                            `json:"colleagues"`
	ContactPoints             *ContactPoint                      `json:"contactpoints"`
	DeathDate                 *Date                              `json:"deathdate"`
	DeathPlace                *Place                             `json:"deathplace"`
	Duns                      *Text                              `json:"duns"`
	Email                     *Text                              `json:"email"`
	FamilyName                *Text                              `json:"familyname"`
	FaxNumber                 *Text                              `json:"faxnumber"`
	Follows                   *Person                            `json:"follows"`
	Funder                    *Person                            `json:"funder"`
	Gender                    *Text                              `json:"gender"`
	GivenName                 *Text                              `json:"givenname"`
	GlobalLocationNumber      *Text                              `json:"globallocationnumber"`
	HasCredential             *EducationalOccupationalCredential `json:"hascredential"`
	HasOccupation             *Occupation                        `json:"hasoccupation"`
	HasOfferCatalog           *OfferCatalog                      `json:"hasoffercatalog"`
	HasPOS                    *Place                             `json:"haspos"`
	Height                    *QuantitativeValue                 `json:"height"`
	HomeLocation              *Place                             `json:"homelocation"`
	HonorificPrefix           *Text                              `json:"honorificprefix"`
	HonorificSuffix           *Text                              `json:"honorificsuffix"`
	IsicV4                    *Text                              `json:"isicv4"`
	JobTitle                  *Text                              `json:"jobtitle"`
	Knows                     *Person                            `json:"knows"`
	KnowsAbout                *URL                               `json:"knowsabout"`
	KnowsLanguage             *Text                              `json:"knowslanguage"`
	Naics                     *Text                              `json:"naics"`
	Nationality               *Country                           `json:"nationality"`
	NetWorth                  *PriceSpecification                `json:"networth"`
	Owns                      *Product                           `json:"owns"`
	Parents                   *Person                            `json:"parents"`
	PerformerIn               *Event                             `json:"performerin"`
	RelatedTo                 *Person                            `json:"relatedto"`
	Seeks                     *Demand                            `json:"seeks"`
	Siblings                  *Person                            `json:"siblings"`
	Spouse                    *Person                            `json:"spouse"`
	TaxID                     *Text                              `json:"taxid"`
	Telephone                 *Text                              `json:"telephone"`
	VatID                     *Text                              `json:"vatid"`
	Weight                    *QuantitativeValue                 `json:"weight"`
	WorkLocation              *Place                             `json:"worklocation"`
	WorksFor                  *Organization                      `json:"worksfor"`
	AlumniOf                  *Organization                      `json:"alumniof"`
	Award                     *Text                              `json:"award"`
	Colleague                 *URL                               `json:"colleague"`
	ContactPoint              *ContactPoint                      `json:"contactpoint"`
	Funding                   *Grant                             `json:"funding"`
	InteractionStatistic      *InteractionCounter                `json:"interactionstatistic"`
	MakesOffer                *Offer                             `json:"makesoffer"`
	Parent                    *Person                            `json:"parent"`
	Sibling                   *Person                            `json:"sibling"`
	Sponsor                   *Person                            `json:"sponsor"`
	MemberOf                  *ProgramMembership                 `json:"memberof"`
	PublishingPrinciples      *URL                               `json:"publishingprinciples"`
}

// SaleEvent is a generated struct representing the https://schema.org/SaleEvent class.
type SaleEvent struct {
}

// SelfStorage is a generated struct representing the https://schema.org/SelfStorage class.
type SelfStorage struct {
}

// ContactPointOption is a generated struct representing the https://schema.org/ContactPointOption class.
type ContactPointOption struct {
}

// Diet is a generated struct representing the https://schema.org/Diet class.
type Diet struct {
	DietFeatures          *Text   `json:"dietfeatures"`
	Endorsers             *Person `json:"endorsers"`
	ExpertConsiderations  *Text   `json:"expertconsiderations"`
	PhysiologicalBenefits *Text   `json:"physiologicalbenefits"`
	Risks                 *Text   `json:"risks"`
}

// Vessel is a generated struct representing the https://schema.org/Vessel class.
type Vessel struct {
}

// AggregateOffer is a generated struct representing the https://schema.org/AggregateOffer class.
type AggregateOffer struct {
	HighPrice  *Text    `json:"highprice"`
	LowPrice   *Text    `json:"lowprice"`
	OfferCount *Integer `json:"offercount"`
	Offers     *Offer   `json:"offers"`
}

// BioChemEntity is a generated struct representing the https://schema.org/BioChemEntity class.
type BioChemEntity struct {
	AssociatedDisease              *URL           `json:"associateddisease"`
	BioChemInteraction             *BioChemEntity `json:"biocheminteraction"`
	BioChemSimilarity              *BioChemEntity `json:"biochemsimilarity"`
	BiologicalRole                 *DefinedTerm   `json:"biologicalrole"`
	HasMolecularFunction           *URL           `json:"hasmolecularfunction"`
	IsInvolvedInBiologicalProcess  *URL           `json:"isinvolvedinbiologicalprocess"`
	IsLocatedInSubcellularLocation *URL           `json:"islocatedinsubcellularlocation"`
	TaxonomicRange                 *URL           `json:"taxonomicrange"`
	Funding                        *Grant         `json:"funding"`
	HasBioChemEntityPart           *BioChemEntity `json:"hasbiochementitypart"`
	IsEncodedByBioChemEntity       *Gene          `json:"isencodedbybiochementity"`
	IsPartOfBioChemEntity          *BioChemEntity `json:"ispartofbiochementity"`
	HasRepresentation              *URL           `json:"hasrepresentation"`
}

// Atlas is a generated struct representing the https://schema.org/Atlas class.
type Atlas struct {
}

// FastFoodRestaurant is a generated struct representing the https://schema.org/FastFoodRestaurant class.
type FastFoodRestaurant struct {
}

// JewelryStore is a generated struct representing the https://schema.org/JewelryStore class.
type JewelryStore struct {
}

// OfferCatalog is a generated struct representing the https://schema.org/OfferCatalog class.
type OfferCatalog struct {
}

// QuantitativeValue is a generated struct representing the https://schema.org/QuantitativeValue class.
type QuantitativeValue struct {
	AdditionalProperty *PropertyValue `json:"additionalproperty"`
	MaxValue           *Number        `json:"maxvalue"`
	MinValue           *Number        `json:"minvalue"`
	UnitCode           *URL           `json:"unitcode"`
	UnitText           *Text          `json:"unittext"`
	Value              *Text          `json:"value"`
	ValueReference     *Text          `json:"valuereference"`
}

// TrainStation is a generated struct representing the https://schema.org/TrainStation class.
type TrainStation struct {
}

// BodyMeasurementTypeEnumeration is a generated struct representing the https://schema.org/BodyMeasurementTypeEnumeration class.
type BodyMeasurementTypeEnumeration struct {
}

// DeactivateAction is a generated struct representing the https://schema.org/DeactivateAction class.
type DeactivateAction struct {
}

// Gene is a generated struct representing the https://schema.org/Gene class.
type Gene struct {
	AlternativeOf         *Gene          `json:"alternativeof"`
	ExpressedIn           *DefinedTerm   `json:"expressedin"`
	HasBioPolymerSequence *Text          `json:"hasbiopolymersequence"`
	EncodesBioChemEntity  *BioChemEntity `json:"encodesbiochementity"`
}

// HighSchool is a generated struct representing the https://schema.org/HighSchool class.
type HighSchool struct {
}

// LandmarksOrHistoricalBuildings is a generated struct representing the https://schema.org/LandmarksOrHistoricalBuildings class.
type LandmarksOrHistoricalBuildings struct {
}

// Schedule is a generated struct representing the https://schema.org/Schedule class.
type Schedule struct {
	ByDay            *Text     `json:"byday"`
	ByMonth          *Integer  `json:"bymonth"`
	ByMonthDay       *Integer  `json:"bymonthday"`
	ByMonthWeek      *Integer  `json:"bymonthweek"`
	EndDate          *DateTime `json:"enddate"`
	EndTime          *Time     `json:"endtime"`
	ExceptDate       *DateTime `json:"exceptdate"`
	RepeatCount      *Integer  `json:"repeatcount"`
	RepeatFrequency  *Text     `json:"repeatfrequency"`
	ScheduleTimezone *Text     `json:"scheduletimezone"`
	StartDate        *DateTime `json:"startdate"`
	StartTime        *Time     `json:"starttime"`
	Duration         *Duration `json:"duration"`
}

// Cemetery is a generated struct representing the https://schema.org/Cemetery class.
type Cemetery struct {
}

// PreOrderAction is a generated struct representing the https://schema.org/PreOrderAction class.
type PreOrderAction struct {
}

// WinAction is a generated struct representing the https://schema.org/WinAction class.
type WinAction struct {
	Loser *Person `json:"loser"`
}

// Crematorium is a generated struct representing the https://schema.org/Crematorium class.
type Crematorium struct {
}

// MusicVenue is a generated struct representing the https://schema.org/MusicVenue class.
type MusicVenue struct {
}

// ControlAction is a generated struct representing the https://schema.org/ControlAction class.
type ControlAction struct {
}

// SeaBodyOfWater is a generated struct representing the https://schema.org/SeaBodyOfWater class.
type SeaBodyOfWater struct {
}

// BedAndBreakfast is a generated struct representing the https://schema.org/BedAndBreakfast class.
type BedAndBreakfast struct {
}

// Country is a generated struct representing the https://schema.org/Country class.
type Country struct {
}

// ExchangeRateSpecification is a generated struct representing the https://schema.org/ExchangeRateSpecification class.
type ExchangeRateSpecification struct {
	Currency            *Text                   `json:"currency"`
	CurrentExchangeRate *UnitPriceSpecification `json:"currentexchangerate"`
	ExchangeRateSpread  *Number                 `json:"exchangeratespread"`
}

// HairSalon is a generated struct representing the https://schema.org/HairSalon class.
type HairSalon struct {
}

// TreatmentIndication is a generated struct representing the https://schema.org/TreatmentIndication class.
type TreatmentIndication struct {
}

// DateTime is a generated struct representing the https://schema.org/DateTime class.
type DateTime struct {
}

// DrugStrength is a generated struct representing the https://schema.org/DrugStrength class.
type DrugStrength struct {
	ActiveIngredient *Text                `json:"activeingredient"`
	AvailableIn      *AdministrativeArea  `json:"availablein"`
	MaximumIntake    *MaximumDoseSchedule `json:"maximumintake"`
	StrengthUnit     *Text                `json:"strengthunit"`
	StrengthValue    *Number              `json:"strengthvalue"`
}

// DryCleaningOrLaundry is a generated struct representing the https://schema.org/DryCleaningOrLaundry class.
type DryCleaningOrLaundry struct {
}

// FinancialProduct is a generated struct representing the https://schema.org/FinancialProduct class.
type FinancialProduct struct {
	AnnualPercentageRate            *QuantitativeValue `json:"annualpercentagerate"`
	FeesAndCommissionsSpecification *URL               `json:"feesandcommissionsspecification"`
	InterestRate                    *QuantitativeValue `json:"interestrate"`
}

// HealthAndBeautyBusiness is a generated struct representing the https://schema.org/HealthAndBeautyBusiness class.
type HealthAndBeautyBusiness struct {
}

// Muscle is a generated struct representing the https://schema.org/Muscle class.
type Muscle struct {
	Antagonist   *Muscle              `json:"antagonist"`
	BloodSupply  *Vessel              `json:"bloodsupply"`
	Insertion    *AnatomicalStructure `json:"insertion"`
	MuscleAction *Text                `json:"muscleaction"`
	Nerve        *Nerve               `json:"nerve"`
}

// ProfessionalService is a generated struct representing the https://schema.org/ProfessionalService class.
type ProfessionalService struct {
}

// UnitPriceSpecification is a generated struct representing the https://schema.org/UnitPriceSpecification class.
type UnitPriceSpecification struct {
	BillingDuration    *QuantitativeValue             `json:"billingduration"`
	BillingIncrement   *Number                        `json:"billingincrement"`
	BillingStart       *Number                        `json:"billingstart"`
	PriceComponentType *PriceComponentTypeEnumeration `json:"pricecomponenttype"`
	PriceType          *Text                          `json:"pricetype"`
	ReferenceQuantity  *QuantitativeValue             `json:"referencequantity"`
	UnitCode           *URL                           `json:"unitcode"`
	UnitText           *Text                          `json:"unittext"`
}

// ArchiveComponent is a generated struct representing the https://schema.org/ArchiveComponent class.
type ArchiveComponent struct {
	ItemLocation   *Text                `json:"itemlocation"`
	HoldingArchive *ArchiveOrganization `json:"holdingarchive"`
}

// CreditCard is a generated struct representing the https://schema.org/CreditCard class.
type CreditCard struct {
}

// AnimalShelter is a generated struct representing the https://schema.org/AnimalShelter class.
type AnimalShelter struct {
}

// DataFeedItem is a generated struct representing the https://schema.org/DataFeedItem class.
type DataFeedItem struct {
	DateDeleted  *DateTime `json:"datedeleted"`
	DateModified *DateTime `json:"datemodified"`
	Item         *Thing    `json:"item"`
	DateCreated  *DateTime `json:"datecreated"`
}

// GasStation is a generated struct representing the https://schema.org/GasStation class.
type GasStation struct {
}

// InfectiousDisease is a generated struct representing the https://schema.org/InfectiousDisease class.
type InfectiousDisease struct {
	InfectiousAgent      *Text                 `json:"infectiousagent"`
	InfectiousAgentClass *InfectiousAgentClass `json:"infectiousagentclass"`
	TransmissionMethod   *Text                 `json:"transmissionmethod"`
}

// LeaveAction is a generated struct representing the https://schema.org/LeaveAction class.
type LeaveAction struct {
	Event *Event `json:"event"`
}

// EntertainmentBusiness is a generated struct representing the https://schema.org/EntertainmentBusiness class.
type EntertainmentBusiness struct {
}

// HowToItem is a generated struct representing the https://schema.org/HowToItem class.
type HowToItem struct {
	RequiredQuantity *Text `json:"requiredquantity"`
}

// Dentist is a generated struct representing the https://schema.org/Dentist class.
type Dentist struct {
}

// MedicalTherapy is a generated struct representing the https://schema.org/MedicalTherapy class.
type MedicalTherapy struct {
	Contraindication      *Text           `json:"contraindication"`
	DuplicateTherapy      *MedicalTherapy `json:"duplicatetherapy"`
	SeriousAdverseOutcome *MedicalEntity  `json:"seriousadverseoutcome"`
}

// ProductCollection is a generated struct representing the https://schema.org/ProductCollection class.
type ProductCollection struct {
	IncludesObject *TypeAndQuantityNode `json:"includesobject"`
}

// SuspendAction is a generated struct representing the https://schema.org/SuspendAction class.
type SuspendAction struct {
}

// DeleteAction is a generated struct representing the https://schema.org/DeleteAction class.
type DeleteAction struct {
}

// Mosque is a generated struct representing the https://schema.org/Mosque class.
type Mosque struct {
}

// PerformingGroup is a generated struct representing the https://schema.org/PerformingGroup class.
type PerformingGroup struct {
}

// VisualArtsEvent is a generated struct representing the https://schema.org/VisualArtsEvent class.
type VisualArtsEvent struct {
}

// ConstraintNode is a generated struct representing the https://schema.org/ConstraintNode class.
type ConstraintNode struct {
	ConstraintProperty *URL     `json:"constraintproperty"`
	NumConstraints     *Integer `json:"numconstraints"`
}

// GardenStore is a generated struct representing the https://schema.org/GardenStore class.
type GardenStore struct {
}

// RadioBroadcastService is a generated struct representing the https://schema.org/RadioBroadcastService class.
type RadioBroadcastService struct {
}

// AssessAction is a generated struct representing the https://schema.org/AssessAction class.
type AssessAction struct {
}

// BowlingAlley is a generated struct representing the https://schema.org/BowlingAlley class.
type BowlingAlley struct {
}

// MedicalDevice is a generated struct representing the https://schema.org/MedicalDevice class.
type MedicalDevice struct {
	AdverseOutcome        *MedicalEntity `json:"adverseoutcome"`
	Contraindication      *Text          `json:"contraindication"`
	PostOp                *Text          `json:"postop"`
	PreOp                 *Text          `json:"preop"`
	Procedure             *Text          `json:"procedure"`
	SeriousAdverseOutcome *MedicalEntity `json:"seriousadverseoutcome"`
}

// PropertyValue is a generated struct representing the https://schema.org/PropertyValue class.
type PropertyValue struct {
	MaxValue             *Number `json:"maxvalue"`
	MeasurementMethod    *URL    `json:"measurementmethod"`
	MinValue             *Number `json:"minvalue"`
	PropertyID           *URL    `json:"propertyid"`
	UnitCode             *URL    `json:"unitcode"`
	UnitText             *Text   `json:"unittext"`
	Value                *Text   `json:"value"`
	ValueReference       *Text   `json:"valuereference"`
	MeasurementTechnique *URL    `json:"measurementtechnique"`
}

// Residence is a generated struct representing the https://schema.org/Residence class.
type Residence struct {
	AccommodationFloorPlan *FloorPlan `json:"accommodationfloorplan"`
}

// Hospital is a generated struct representing the https://schema.org/Hospital class.
type Hospital struct {
	AvailableService        *MedicalTherapy   `json:"availableservice"`
	HealthcareReportingData *Dataset          `json:"healthcarereportingdata"`
	MedicalSpecialty        *MedicalSpecialty `json:"medicalspecialty"`
}

// Report is a generated struct representing the https://schema.org/Report class.
type Report struct {
	ReportNumber *Text `json:"reportnumber"`
}

// MedicalAudienceType is a generated struct representing the https://schema.org/MedicalAudienceType class.
type MedicalAudienceType struct {
}

// SocialMediaPosting is a generated struct representing the https://schema.org/SocialMediaPosting class.
type SocialMediaPosting struct {
	SharedContent *CreativeWork `json:"sharedcontent"`
}

// GovernmentBuilding is a generated struct representing the https://schema.org/GovernmentBuilding class.
type GovernmentBuilding struct {
}

// MedicalCode is a generated struct representing the https://schema.org/MedicalCode class.
type MedicalCode struct {
	CodeValue    *Text `json:"codevalue"`
	CodingSystem *Text `json:"codingsystem"`
}

// Answer is a generated struct representing the https://schema.org/Answer class.
type Answer struct {
	AnswerExplanation *WebContent   `json:"answerexplanation"`
	ParentItem        *CreativeWork `json:"parentitem"`
}

// BusReservation is a generated struct representing the https://schema.org/BusReservation class.
type BusReservation struct {
}

// ContactPage is a generated struct representing the https://schema.org/ContactPage class.
type ContactPage struct {
}

// InteractAction is a generated struct representing the https://schema.org/InteractAction class.
type InteractAction struct {
}

// URL is a generated struct representing the https://schema.org/URL class.
type URL struct {
}

// BusTrip is a generated struct representing the https://schema.org/BusTrip class.
type BusTrip struct {
	ArrivalBusStop   *BusStop `json:"arrivalbusstop"`
	BusName          *Text    `json:"busname"`
	BusNumber        *Text    `json:"busnumber"`
	DepartureBusStop *BusStop `json:"departurebusstop"`
}

// BusinessEntityType is a generated struct representing the https://schema.org/BusinessEntityType class.
type BusinessEntityType struct {
}

// DayOfWeek is a generated struct representing the https://schema.org/DayOfWeek class.
type DayOfWeek struct {
}

// DrugCostCategory is a generated struct representing the https://schema.org/DrugCostCategory class.
type DrugCostCategory struct {
}

// FloorPlan is a generated struct representing the https://schema.org/FloorPlan class.
type FloorPlan struct {
	AmenityFeature                      *LocationFeatureSpecification `json:"amenityfeature"`
	FloorSize                           *QuantitativeValue            `json:"floorsize"`
	IsPlanForApartment                  *Accommodation                `json:"isplanforapartment"`
	LayoutImage                         *URL                          `json:"layoutimage"`
	NumberOfAccommodationUnits          *QuantitativeValue            `json:"numberofaccommodationunits"`
	NumberOfAvailableAccommodationUnits *QuantitativeValue            `json:"numberofavailableaccommodationunits"`
	NumberOfBathroomsTotal              *Integer                      `json:"numberofbathroomstotal"`
	NumberOfBedrooms                    *QuantitativeValue            `json:"numberofbedrooms"`
	NumberOfFullBathrooms               *Number                       `json:"numberoffullbathrooms"`
	NumberOfPartialBathrooms            *Number                       `json:"numberofpartialbathrooms"`
	NumberOfRooms                       *QuantitativeValue            `json:"numberofrooms"`
	PetsAllowed                         *Text                         `json:"petsallowed"`
}

// LodgingReservation is a generated struct representing the https://schema.org/LodgingReservation class.
type LodgingReservation struct {
	CheckinTime            *Time              `json:"checkintime"`
	CheckoutTime           *Time              `json:"checkouttime"`
	LodgingUnitDescription *Text              `json:"lodgingunitdescription"`
	LodgingUnitType        *Text              `json:"lodgingunittype"`
	NumAdults              *QuantitativeValue `json:"numadults"`
	NumChildren            *QuantitativeValue `json:"numchildren"`
}

// MedicalIntangible is a generated struct representing the https://schema.org/MedicalIntangible class.
type MedicalIntangible struct {
}

// TrainTrip is a generated struct representing the https://schema.org/TrainTrip class.
type TrainTrip struct {
	ArrivalPlatform   *Text         `json:"arrivalplatform"`
	ArrivalStation    *TrainStation `json:"arrivalstation"`
	DeparturePlatform *Text         `json:"departureplatform"`
	DepartureStation  *TrainStation `json:"departurestation"`
	TrainName         *Text         `json:"trainname"`
	TrainNumber       *Text         `json:"trainnumber"`
}

// CompleteDataFeed is a generated struct representing the https://schema.org/CompleteDataFeed class.
type CompleteDataFeed struct {
}

// CssSelectorType is a generated struct representing the https://schema.org/CssSelectorType class.
type CssSelectorType struct {
}

// HealthPlanNetwork is a generated struct representing the https://schema.org/HealthPlanNetwork class.
type HealthPlanNetwork struct {
	HealthPlanCostSharing *Boolean `json:"healthplancostsharing"`
	HealthPlanNetworkId   *Text    `json:"healthplannetworkid"`
	HealthPlanNetworkTier *Text    `json:"healthplannetworktier"`
}

// Hostel is a generated struct representing the https://schema.org/Hostel class.
type Hostel struct {
}

// ReservationPackage is a generated struct representing the https://schema.org/ReservationPackage class.
type ReservationPackage struct {
	SubReservation *Reservation `json:"subreservation"`
}

// AudioObject is a generated struct representing the https://schema.org/AudioObject class.
type AudioObject struct {
	EmbeddedTextCaption *Text `json:"embeddedtextcaption"`
	Transcript          *Text `json:"transcript"`
	Caption             *Text `json:"caption"`
}

// State is a generated struct representing the https://schema.org/State class.
type State struct {
}

// TouristAttraction is a generated struct representing the https://schema.org/TouristAttraction class.
type TouristAttraction struct {
	AvailableLanguage *Text `json:"availablelanguage"`
	TouristType       *Text `json:"touristtype"`
}

// TouristDestination is a generated struct representing the https://schema.org/TouristDestination class.
type TouristDestination struct {
	IncludesAttraction *TouristAttraction `json:"includesattraction"`
	TouristType        *Text              `json:"touristtype"`
}

// StatisticalPopulation is a generated struct representing the https://schema.org/StatisticalPopulation class.
type StatisticalPopulation struct {
	PopulationType *Class `json:"populationtype"`
}

// DrugClass is a generated struct representing the https://schema.org/DrugClass class.
type DrugClass struct {
	Drug *Drug `json:"drug"`
}

// GeospatialGeometry is a generated struct representing the https://schema.org/GeospatialGeometry class.
type GeospatialGeometry struct {
	GeoContains   *Place `json:"geocontains"`
	GeoCoveredBy  *Place `json:"geocoveredby"`
	GeoCovers     *Place `json:"geocovers"`
	GeoCrosses    *Place `json:"geocrosses"`
	GeoDisjoint   *Place `json:"geodisjoint"`
	GeoEquals     *Place `json:"geoequals"`
	GeoIntersects *Place `json:"geointersects"`
	GeoOverlaps   *Place `json:"geooverlaps"`
	GeoTouches    *Place `json:"geotouches"`
	GeoWithin     *Place `json:"geowithin"`
}

// HardwareStore is a generated struct representing the https://schema.org/HardwareStore class.
type HardwareStore struct {
}

// HowToDirection is a generated struct representing the https://schema.org/HowToDirection class.
type HowToDirection struct {
	AfterMedia  *URL      `json:"aftermedia"`
	BeforeMedia *URL      `json:"beforemedia"`
	DuringMedia *URL      `json:"duringmedia"`
	PrepTime    *Duration `json:"preptime"`
	Tool        *Text     `json:"tool"`
	TotalTime   *Duration `json:"totaltime"`
	PerformTime *Duration `json:"performtime"`
	Supply      *Text     `json:"supply"`
}

// MotorcycleDealer is a generated struct representing the https://schema.org/MotorcycleDealer class.
type MotorcycleDealer struct {
}

// MusicAlbum is a generated struct representing the https://schema.org/MusicAlbum class.
type MusicAlbum struct {
	AlbumProductionType *MusicAlbumProductionType `json:"albumproductiontype"`
	AlbumReleaseType    *MusicAlbumReleaseType    `json:"albumreleasetype"`
	ByArtist            *Person                   `json:"byartist"`
	AlbumRelease        *MusicRelease             `json:"albumrelease"`
}

// SingleFamilyResidence is a generated struct representing the https://schema.org/SingleFamilyResidence class.
type SingleFamilyResidence struct {
	NumberOfRooms *QuantitativeValue `json:"numberofrooms"`
	Occupancy     *QuantitativeValue `json:"occupancy"`
}

// Boolean is a generated struct representing the https://schema.org/Boolean class.
type Boolean struct {
}

// Brewery is a generated struct representing the https://schema.org/Brewery class.
type Brewery struct {
}

// Mass is a generated struct representing the https://schema.org/Mass class.
type Mass struct {
}

// PaymentCard is a generated struct representing the https://schema.org/PaymentCard class.
type PaymentCard struct {
	CashBack                      *Number         `json:"cashback"`
	ContactlessPayment            *Boolean        `json:"contactlesspayment"`
	FloorLimit                    *MonetaryAmount `json:"floorlimit"`
	MonthlyMinimumRepaymentAmount *Number         `json:"monthlyminimumrepaymentamount"`
}

// PronounceableText is a generated struct representing the https://schema.org/PronounceableText class.
type PronounceableText struct {
	PhoneticText       *Text `json:"phonetictext"`
	SpeechToTextMarkup *Text `json:"speechtotextmarkup"`
	TextValue          *Text `json:"textvalue"`
	InLanguage         *Text `json:"inlanguage"`
}

// Ticket is a generated struct representing the https://schema.org/Ticket class.
type Ticket struct {
	DateIssued    *DateTime     `json:"dateissued"`
	IssuedBy      *Organization `json:"issuedby"`
	PriceCurrency *Text         `json:"pricecurrency"`
	TicketNumber  *Text         `json:"ticketnumber"`
	TicketToken   *URL          `json:"tickettoken"`
	TicketedSeat  *Seat         `json:"ticketedseat"`
	TotalPrice    *Text         `json:"totalprice"`
	UnderName     *Person       `json:"undername"`
}

// DrawAction is a generated struct representing the https://schema.org/DrawAction class.
type DrawAction struct {
}

// DrugPregnancyCategory is a generated struct representing the https://schema.org/DrugPregnancyCategory class.
type DrugPregnancyCategory struct {
}

// ImageObjectSnapshot is a generated struct representing the https://schema.org/ImageObjectSnapshot class.
type ImageObjectSnapshot struct {
}

// OnlineStore is a generated struct representing the https://schema.org/OnlineStore class.
type OnlineStore struct {
}

// ProfilePage is a generated struct representing the https://schema.org/ProfilePage class.
type ProfilePage struct {
}

// DefenceEstablishment is a generated struct representing the https://schema.org/DefenceEstablishment class.
type DefenceEstablishment struct {
}

// Patient is a generated struct representing the https://schema.org/Patient class.
type Patient struct {
	Diagnosis       *MedicalCondition `json:"diagnosis"`
	Drug            *Drug             `json:"drug"`
	HealthCondition *MedicalCondition `json:"healthcondition"`
}

// PaymentMethod is a generated struct representing the https://schema.org/PaymentMethod class.
type PaymentMethod struct {
}

// BusinessAudience is a generated struct representing the https://schema.org/BusinessAudience class.
type BusinessAudience struct {
	NumberOfEmployees *QuantitativeValue `json:"numberofemployees"`
	YearlyRevenue     *QuantitativeValue `json:"yearlyrevenue"`
	YearsInOperation  *QuantitativeValue `json:"yearsinoperation"`
}

// GameServerStatus is a generated struct representing the https://schema.org/GameServerStatus class.
type GameServerStatus struct {
}

// GeoCoordinates is a generated struct representing the https://schema.org/GeoCoordinates class.
type GeoCoordinates struct {
	Address        *Text `json:"address"`
	AddressCountry *Text `json:"addresscountry"`
	Elevation      *Text `json:"elevation"`
	Latitude       *Text `json:"latitude"`
	Longitude      *Text `json:"longitude"`
	PostalCode     *Text `json:"postalcode"`
}

// JoinAction is a generated struct representing the https://schema.org/JoinAction class.
type JoinAction struct {
	Event *Event `json:"event"`
}

// ClaimReview is a generated struct representing the https://schema.org/ClaimReview class.
type ClaimReview struct {
	ClaimReviewed *Text `json:"claimreviewed"`
}

// FollowAction is a generated struct representing the https://schema.org/FollowAction class.
type FollowAction struct {
	Followee *Person `json:"followee"`
}

// MeasurementMethodEnum is a generated struct representing the https://schema.org/MeasurementMethodEnum class.
type MeasurementMethodEnum struct {
}

// TheaterGroup is a generated struct representing the https://schema.org/TheaterGroup class.
type TheaterGroup struct {
}

// AuthorizeAction is a generated struct representing the https://schema.org/AuthorizeAction class.
type AuthorizeAction struct {
	Recipient *Person `json:"recipient"`
}

// MedicalTestPanel is a generated struct representing the https://schema.org/MedicalTestPanel class.
type MedicalTestPanel struct {
	SubTest *MedicalTest `json:"subtest"`
}

// PaymentStatusType is a generated struct representing the https://schema.org/PaymentStatusType class.
type PaymentStatusType struct {
}

// PhysicalActivity is a generated struct representing the https://schema.org/PhysicalActivity class.
type PhysicalActivity struct {
	AssociatedAnatomy *SuperficialAnatomy `json:"associatedanatomy"`
	Epidemiology      *Text               `json:"epidemiology"`
	Pathophysiology   *Text               `json:"pathophysiology"`
	Category          *URL                `json:"category"`
}

// Season is a generated struct representing the https://schema.org/Season class.
type Season struct {
}

// SendAction is a generated struct representing the https://schema.org/SendAction class.
type SendAction struct {
	DeliveryMethod *DeliveryMethod `json:"deliverymethod"`
	Recipient      *Person         `json:"recipient"`
}

// WPSideBar is a generated struct representing the https://schema.org/WPSideBar class.
type WPSideBar struct {
}

// LocalBusiness is a generated struct representing the https://schema.org/LocalBusiness class.
type LocalBusiness struct {
	BranchOf           *Organization `json:"branchof"`
	CurrenciesAccepted *Text         `json:"currenciesaccepted"`
	OpeningHours       *Text         `json:"openinghours"`
	PaymentAccepted    *Text         `json:"paymentaccepted"`
	PriceRange         *Text         `json:"pricerange"`
}

// OfficeEquipmentStore is a generated struct representing the https://schema.org/OfficeEquipmentStore class.
type OfficeEquipmentStore struct {
}

// RepaymentSpecification is a generated struct representing the https://schema.org/RepaymentSpecification class.
type RepaymentSpecification struct {
	DownPayment            *Number         `json:"downpayment"`
	EarlyPrepaymentPenalty *MonetaryAmount `json:"earlyprepaymentpenalty"`
	LoanPaymentAmount      *MonetaryAmount `json:"loanpaymentamount"`
	LoanPaymentFrequency   *Number         `json:"loanpaymentfrequency"`
	NumberOfLoanPayments   *Number         `json:"numberofloanpayments"`
}

// Table is a generated struct representing the https://schema.org/Table class.
type Table struct {
}

// CommunicateAction is a generated struct representing the https://schema.org/CommunicateAction class.
type CommunicateAction struct {
	Language   *Language `json:"language"`
	InLanguage *Text     `json:"inlanguage"`
	About      *Thing    `json:"about"`
	Recipient  *Person   `json:"recipient"`
}

// XPathType is a generated struct representing the https://schema.org/XPathType class.
type XPathType struct {
}

// GovernmentService is a generated struct representing the https://schema.org/GovernmentService class.
type GovernmentService struct {
	ServiceOperator *Organization `json:"serviceoperator"`
	Jurisdiction    *Text         `json:"jurisdiction"`
}

// TouristInformationCenter is a generated struct representing the https://schema.org/TouristInformationCenter class.
type TouristInformationCenter struct {
}

// Event is a generated struct representing the https://schema.org/Event class.
type Event struct {
	AggregateRating                 *AggregateRating                `json:"aggregaterating"`
	Attendees                       *Person                         `json:"attendees"`
	Composer                        *Person                         `json:"composer"`
	Contributor                     *Person                         `json:"contributor"`
	DoorTime                        *Time                           `json:"doortime"`
	EndDate                         *DateTime                       `json:"enddate"`
	EventAttendanceMode             *EventAttendanceModeEnumeration `json:"eventattendancemode"`
	EventSchedule                   *Schedule                       `json:"eventschedule"`
	EventStatus                     *EventStatusType                `json:"eventstatus"`
	Funder                          *Person                         `json:"funder"`
	Keywords                        *URL                            `json:"keywords"`
	MaximumAttendeeCapacity         *Integer                        `json:"maximumattendeecapacity"`
	MaximumPhysicalAttendeeCapacity *Integer                        `json:"maximumphysicalattendeecapacity"`
	MaximumVirtualAttendeeCapacity  *Integer                        `json:"maximumvirtualattendeecapacity"`
	Organizer                       *Person                         `json:"organizer"`
	Performers                      *Person                         `json:"performers"`
	PreviousStartDate               *Date                           `json:"previousstartdate"`
	RemainingAttendeeCapacity       *Integer                        `json:"remainingattendeecapacity"`
	StartDate                       *DateTime                       `json:"startdate"`
	SubEvents                       *Event                          `json:"subevents"`
	Translator                      *Person                         `json:"translator"`
	TypicalAgeRange                 *Text                           `json:"typicalagerange"`
	WorkPerformed                   *CreativeWork                   `json:"workperformed"`
	Attendee                        *Person                         `json:"attendee"`
	Audience                        *Audience                       `json:"audience"`
	Director                        *Person                         `json:"director"`
	Duration                        *Duration                       `json:"duration"`
	Funding                         *Grant                          `json:"funding"`
	InLanguage                      *Text                           `json:"inlanguage"`
	IsAccessibleForFree             *Boolean                        `json:"isaccessibleforfree"`
	Offers                          *Offer                          `json:"offers"`
	Performer                       *Person                         `json:"performer"`
	RecordedIn                      *CreativeWork                   `json:"recordedin"`
	Review                          *Review                         `json:"review"`
	Sponsor                         *Person                         `json:"sponsor"`
	SuperEvent                      *Event                          `json:"superevent"`
	About                           *Thing                          `json:"about"`
	Actor                           *Person                         `json:"actor"`
	SubEvent                        *Event                          `json:"subevent"`
	WorkFeatured                    *CreativeWork                   `json:"workfeatured"`
	Location                        *VirtualLocation                `json:"location"`
}

// HowToTool is a generated struct representing the https://schema.org/HowToTool class.
type HowToTool struct {
}

// MerchantReturnEnumeration is a generated struct representing the https://schema.org/MerchantReturnEnumeration class.
type MerchantReturnEnumeration struct {
}

// MovieRentalStore is a generated struct representing the https://schema.org/MovieRentalStore class.
type MovieRentalStore struct {
}

// WorkBasedProgram is a generated struct representing the https://schema.org/WorkBasedProgram class.
type WorkBasedProgram struct {
	OccupationalCategory *Text                       `json:"occupationalcategory"`
	TrainingSalary       *MonetaryAmountDistribution `json:"trainingsalary"`
}

// BroadcastFrequencySpecification is a generated struct representing the https://schema.org/BroadcastFrequencySpecification class.
type BroadcastFrequencySpecification struct {
	BroadcastFrequencyValue   *QuantitativeValue `json:"broadcastfrequencyvalue"`
	BroadcastSignalModulation *Text              `json:"broadcastsignalmodulation"`
	BroadcastSubChannel       *Text              `json:"broadcastsubchannel"`
}

// CreativeWorkSeason is a generated struct representing the https://schema.org/CreativeWorkSeason class.
type CreativeWorkSeason struct {
	EndDate           *DateTime           `json:"enddate"`
	Episodes          *Episode            `json:"episodes"`
	NumberOfEpisodes  *Integer            `json:"numberofepisodes"`
	ProductionCompany *Organization       `json:"productioncompany"`
	SeasonNumber      *Text               `json:"seasonnumber"`
	StartDate         *DateTime           `json:"startdate"`
	Trailer           *VideoObject        `json:"trailer"`
	Director          *Person             `json:"director"`
	Episode           *Episode            `json:"episode"`
	PartOfSeries      *CreativeWorkSeries `json:"partofseries"`
	Actor             *Person             `json:"actor"`
}

// FundingScheme is a generated struct representing the https://schema.org/FundingScheme class.
type FundingScheme struct {
}

// RealEstateListing is a generated struct representing the https://schema.org/RealEstateListing class.
type RealEstateListing struct {
	DatePosted  *DateTime          `json:"dateposted"`
	LeaseLength *QuantitativeValue `json:"leaselength"`
}

// DrugPrescriptionStatus is a generated struct representing the https://schema.org/DrugPrescriptionStatus class.
type DrugPrescriptionStatus struct {
}

// LiteraryEvent is a generated struct representing the https://schema.org/LiteraryEvent class.
type LiteraryEvent struct {
}

// OccupationalExperienceRequirements is a generated struct representing the https://schema.org/OccupationalExperienceRequirements class.
type OccupationalExperienceRequirements struct {
	MonthsOfExperience *Number `json:"monthsofexperience"`
}

// BookSeries is a generated struct representing the https://schema.org/BookSeries class.
type BookSeries struct {
}

// GovernmentBenefitsType is a generated struct representing the https://schema.org/GovernmentBenefitsType class.
type GovernmentBenefitsType struct {
}

// MiddleSchool is a generated struct representing the https://schema.org/MiddleSchool class.
type MiddleSchool struct {
}

// PhysicalActivityCategory is a generated struct representing the https://schema.org/PhysicalActivityCategory class.
type PhysicalActivityCategory struct {
}

// RiverBodyOfWater is a generated struct representing the https://schema.org/RiverBodyOfWater class.
type RiverBodyOfWater struct {
}

// Action is a generated struct representing the https://schema.org/Action class.
type Action struct {
	ActionStatus *ActionStatusType `json:"actionstatus"`
	Agent        *Person           `json:"agent"`
	EndTime      *Time             `json:"endtime"`
	Error        *Thing            `json:"error"`
	StartTime    *Time             `json:"starttime"`
	Target       *URL              `json:"target"`
	Provider     *Person           `json:"provider"`
	Result       *Thing            `json:"result"`
	Instrument   *Thing            `json:"instrument"`
	Object       *Thing            `json:"object"`
	Location     *VirtualLocation  `json:"location"`
	Participant  *Person           `json:"participant"`
}

// CreativeWorkSeries is a generated struct representing the https://schema.org/CreativeWorkSeries class.
type CreativeWorkSeries struct {
	EndDate   *DateTime `json:"enddate"`
	Issn      *Text     `json:"issn"`
	StartDate *DateTime `json:"startdate"`
}

// Museum is a generated struct representing the https://schema.org/Museum class.
type Museum struct {
}

// APIReference is a generated struct representing the https://schema.org/APIReference class.
type APIReference struct {
	Assembly              *Text `json:"assembly"`
	AssemblyVersion       *Text `json:"assemblyversion"`
	ProgrammingModel      *Text `json:"programmingmodel"`
	TargetPlatform        *Text `json:"targetplatform"`
	ExecutableLibraryName *Text `json:"executablelibraryname"`
}

// Bridge is a generated struct representing the https://schema.org/Bridge class.
type Bridge struct {
}

// DiscussionForumPosting is a generated struct representing the https://schema.org/DiscussionForumPosting class.
type DiscussionForumPosting struct {
}

// Taxon is a generated struct representing the https://schema.org/Taxon class.
type Taxon struct {
	TaxonRank      *URL         `json:"taxonrank"`
	ChildTaxon     *URL         `json:"childtaxon"`
	HasDefinedTerm *DefinedTerm `json:"hasdefinedterm"`
	ParentTaxon    *URL         `json:"parenttaxon"`
}

// SomeProducts is a generated struct representing the https://schema.org/SomeProducts class.
type SomeProducts struct {
	InventoryLevel *QuantitativeValue `json:"inventorylevel"`
}

// AudioObjectSnapshot is a generated struct representing the https://schema.org/AudioObjectSnapshot class.
type AudioObjectSnapshot struct {
}

// ClothingStore is a generated struct representing the https://schema.org/ClothingStore class.
type ClothingStore struct {
}

// DataCatalog is a generated struct representing the https://schema.org/DataCatalog class.
type DataCatalog struct {
	MeasurementMethod    *URL     `json:"measurementmethod"`
	Dataset              *Dataset `json:"dataset"`
	MeasurementTechnique *URL     `json:"measurementtechnique"`
}

// FurnitureStore is a generated struct representing the https://schema.org/FurnitureStore class.
type FurnitureStore struct {
}

// MedicalRiskEstimator is a generated struct representing the https://schema.org/MedicalRiskEstimator class.
type MedicalRiskEstimator struct {
	EstimatesRiskOf    *MedicalEntity     `json:"estimatesriskof"`
	IncludedRiskFactor *MedicalRiskFactor `json:"includedriskfactor"`
}

// OrganizationRole is a generated struct representing the https://schema.org/OrganizationRole class.
type OrganizationRole struct {
	NumberedPosition *Number `json:"numberedposition"`
}

// Place is a generated struct representing the https://schema.org/Place class.
type Place struct {
	AdditionalProperty               *PropertyValue                `json:"additionalproperty"`
	Address                          *Text                         `json:"address"`
	AggregateRating                  *AggregateRating              `json:"aggregaterating"`
	AmenityFeature                   *LocationFeatureSpecification `json:"amenityfeature"`
	BranchCode                       *Text                         `json:"branchcode"`
	ContainedIn                      *Place                        `json:"containedin"`
	Events                           *Event                        `json:"events"`
	FaxNumber                        *Text                         `json:"faxnumber"`
	Geo                              *GeoShape                     `json:"geo"`
	GeoContains                      *Place                        `json:"geocontains"`
	GeoCoveredBy                     *Place                        `json:"geocoveredby"`
	GeoCovers                        *Place                        `json:"geocovers"`
	GeoCrosses                       *Place                        `json:"geocrosses"`
	GeoDisjoint                      *Place                        `json:"geodisjoint"`
	GeoEquals                        *Place                        `json:"geoequals"`
	GeoIntersects                    *Place                        `json:"geointersects"`
	GeoOverlaps                      *Place                        `json:"geooverlaps"`
	GeoTouches                       *Place                        `json:"geotouches"`
	GeoWithin                        *Place                        `json:"geowithin"`
	GlobalLocationNumber             *Text                         `json:"globallocationnumber"`
	HasDriveThroughService           *Boolean                      `json:"hasdrivethroughservice"`
	IsicV4                           *Text                         `json:"isicv4"`
	Keywords                         *URL                          `json:"keywords"`
	Latitude                         *Text                         `json:"latitude"`
	Logo                             *URL                          `json:"logo"`
	Longitude                        *Text                         `json:"longitude"`
	Map                              *URL                          `json:"map"`
	Maps                             *URL                          `json:"maps"`
	MaximumAttendeeCapacity          *Integer                      `json:"maximumattendeecapacity"`
	OpeningHoursSpecification        *OpeningHoursSpecification    `json:"openinghoursspecification"`
	Photos                           *Photograph                   `json:"photos"`
	PublicAccess                     *Boolean                      `json:"publicaccess"`
	Reviews                          *Review                       `json:"reviews"`
	Slogan                           *Text                         `json:"slogan"`
	SmokingAllowed                   *Boolean                      `json:"smokingallowed"`
	SpecialOpeningHoursSpecification *OpeningHoursSpecification    `json:"specialopeninghoursspecification"`
	Telephone                        *Text                         `json:"telephone"`
	TourBookingPage                  *URL                          `json:"tourbookingpage"`
	ContainsPlace                    *Place                        `json:"containsplace"`
	Event                            *Event                        `json:"event"`
	IsAccessibleForFree              *Boolean                      `json:"isaccessibleforfree"`
	Photo                            *Photograph                   `json:"photo"`
	Review                           *Review                       `json:"review"`
	ContainedInPlace                 *Place                        `json:"containedinplace"`
	HasMap                           *URL                          `json:"hasmap"`
}

// TipAction is a generated struct representing the https://schema.org/TipAction class.
type TipAction struct {
	Recipient *Person `json:"recipient"`
}

// Winery is a generated struct representing the https://schema.org/Winery class.
type Winery struct {
}

// FMRadioChannel is a generated struct representing the https://schema.org/FMRadioChannel class.
type FMRadioChannel struct {
}

// IgnoreAction is a generated struct representing the https://schema.org/IgnoreAction class.
type IgnoreAction struct {
}

// Invoice is a generated struct representing the https://schema.org/Invoice class.
type Invoice struct {
	AccountId            *Text               `json:"accountid"`
	BillingPeriod        *Duration           `json:"billingperiod"`
	ConfirmationNumber   *Text               `json:"confirmationnumber"`
	Customer             *Person             `json:"customer"`
	MinimumPaymentDue    *PriceSpecification `json:"minimumpaymentdue"`
	PaymentDue           *DateTime           `json:"paymentdue"`
	PaymentMethod        *PaymentMethod      `json:"paymentmethod"`
	PaymentMethodId      *Text               `json:"paymentmethodid"`
	PaymentStatus        *Text               `json:"paymentstatus"`
	ReferencesOrder      *Order              `json:"referencesorder"`
	ScheduledPaymentDate *Date               `json:"scheduledpaymentdate"`
	TotalPaymentDue      *PriceSpecification `json:"totalpaymentdue"`
	Broker               *Person             `json:"broker"`
	Category             *URL                `json:"category"`
	PaymentDueDate       *DateTime           `json:"paymentduedate"`
	Provider             *Person             `json:"provider"`
}

// MedicalRiskFactor is a generated struct representing the https://schema.org/MedicalRiskFactor class.
type MedicalRiskFactor struct {
	IncreasesRiskOf *MedicalEntity `json:"increasesriskof"`
}

// MusicComposition is a generated struct representing the https://schema.org/MusicComposition class.
type MusicComposition struct {
	Composer             *Person           `json:"composer"`
	FirstPerformance     *Event            `json:"firstperformance"`
	IncludedComposition  *MusicComposition `json:"includedcomposition"`
	IswcCode             *Text             `json:"iswccode"`
	Lyricist             *Person           `json:"lyricist"`
	Lyrics               *CreativeWork     `json:"lyrics"`
	MusicArrangement     *MusicComposition `json:"musicarrangement"`
	MusicCompositionForm *Text             `json:"musiccompositionform"`
	MusicalKey           *Text             `json:"musicalkey"`
	RecordedAs           *MusicRecording   `json:"recordedas"`
}

// Service is a generated struct representing the https://schema.org/Service class.
type Service struct {
	AggregateRating  *AggregateRating           `json:"aggregaterating"`
	AvailableChannel *ServiceChannel            `json:"availablechannel"`
	Brand            *Organization              `json:"brand"`
	HasOfferCatalog  *OfferCatalog              `json:"hasoffercatalog"`
	HoursAvailable   *OpeningHoursSpecification `json:"hoursavailable"`
	IsRelatedTo      *Service                   `json:"isrelatedto"`
	IsSimilarTo      *Service                   `json:"issimilarto"`
	Logo             *URL                       `json:"logo"`
	Produces         *Thing                     `json:"produces"`
	ProviderMobility *Text                      `json:"providermobility"`
	ServiceAudience  *Audience                  `json:"serviceaudience"`
	ServiceType      *Text                      `json:"servicetype"`
	Slogan           *Text                      `json:"slogan"`
	TermsOfService   *URL                       `json:"termsofservice"`
	Audience         *Audience                  `json:"audience"`
	Award            *Text                      `json:"award"`
	Broker           *Person                    `json:"broker"`
	Category         *URL                       `json:"category"`
	Offers           *Offer                     `json:"offers"`
	Provider         *Person                    `json:"provider"`
	Review           *Review                    `json:"review"`
	ServiceArea      *Place                     `json:"servicearea"`
	ServiceOutput    *Thing                     `json:"serviceoutput"`
	AreaServed       *Text                      `json:"areaserved"`
}

// UnRegisterAction is a generated struct representing the https://schema.org/UnRegisterAction class.
type UnRegisterAction struct {
}

// FoodEstablishment is a generated struct representing the https://schema.org/FoodEstablishment class.
type FoodEstablishment struct {
	AcceptsReservations *URL    `json:"acceptsreservations"`
	Menu                *URL    `json:"menu"`
	ServesCuisine       *Text   `json:"servescuisine"`
	StarRating          *Rating `json:"starrating"`
	HasMenu             *URL    `json:"hasmenu"`
}

// ShareAction is a generated struct representing the https://schema.org/ShareAction class.
type ShareAction struct {
}

// AdvertiserContentArticle is a generated struct representing the https://schema.org/AdvertiserContentArticle class.
type AdvertiserContentArticle struct {
}

// Collection is a generated struct representing the https://schema.org/Collection class.
type Collection struct {
	CollectionSize *Integer `json:"collectionsize"`
}

// CorrectionComment is a generated struct representing the https://schema.org/CorrectionComment class.
type CorrectionComment struct {
}

// SellAction is a generated struct representing the https://schema.org/SellAction class.
type SellAction struct {
	Buyer           *Person          `json:"buyer"`
	WarrantyPromise *WarrantyPromise `json:"warrantypromise"`
}

// CreateAction is a generated struct representing the https://schema.org/CreateAction class.
type CreateAction struct {
}

// FlightReservation is a generated struct representing the https://schema.org/FlightReservation class.
type FlightReservation struct {
	BoardingGroup           *Text `json:"boardinggroup"`
	PassengerPriorityStatus *Text `json:"passengerprioritystatus"`
	PassengerSequenceNumber *Text `json:"passengersequencenumber"`
	SecurityScreening       *Text `json:"securityscreening"`
}

// JobPosting is a generated struct representing the https://schema.org/JobPosting class.
type JobPosting struct {
	ApplicantLocationRequirements *AdministrativeArea `json:"applicantlocationrequirements"`
	ApplicationContact            *ContactPoint       `json:"applicationcontact"`
	BaseSalary                    *PriceSpecification `json:"basesalary"`
	Benefits                      *Text               `json:"benefits"`
	DatePosted                    *DateTime           `json:"dateposted"`
	DirectApply                   *Boolean            `json:"directapply"`
	EducationRequirements         *Text               `json:"educationrequirements"`
	EligibilityToWorkRequirement  *Text               `json:"eligibilitytoworkrequirement"`
	EmployerOverview              *Text               `json:"employeroverview"`
	EmploymentType                *Text               `json:"employmenttype"`
	EmploymentUnit                *Organization       `json:"employmentunit"`
	EstimatedSalary               *Number             `json:"estimatedsalary"`
	ExperienceInPlaceOfEducation  *Boolean            `json:"experienceinplaceofeducation"`
	ExperienceRequirements        *Text               `json:"experiencerequirements"`
	HiringOrganization            *Person             `json:"hiringorganization"`
	Incentives                    *Text               `json:"incentives"`
	Industry                      *Text               `json:"industry"`
	JobImmediateStart             *Boolean            `json:"jobimmediatestart"`
	JobLocation                   *Place              `json:"joblocation"`
	JobLocationType               *Text               `json:"joblocationtype"`
	JobStartDate                  *Text               `json:"jobstartdate"`
	OccupationalCategory          *Text               `json:"occupationalcategory"`
	PhysicalRequirement           *URL                `json:"physicalrequirement"`
	Qualifications                *Text               `json:"qualifications"`
	RelevantOccupation            *Occupation         `json:"relevantoccupation"`
	Responsibilities              *Text               `json:"responsibilities"`
	SalaryCurrency                *Text               `json:"salarycurrency"`
	SecurityClearanceRequirement  *URL                `json:"securityclearancerequirement"`
	SensoryRequirement            *URL                `json:"sensoryrequirement"`
	Skills                        *Text               `json:"skills"`
	SpecialCommitments            *Text               `json:"specialcommitments"`
	Title                         *Text               `json:"title"`
	TotalJobOpenings              *Integer            `json:"totaljobopenings"`
	ValidThrough                  *DateTime           `json:"validthrough"`
	WorkHours                     *Text               `json:"workhours"`
	IncentiveCompensation         *Text               `json:"incentivecompensation"`
	JobBenefits                   *Text               `json:"jobbenefits"`
}

// MonetaryAmountDistribution is a generated struct representing the https://schema.org/MonetaryAmountDistribution class.
type MonetaryAmountDistribution struct {
	Currency *Text `json:"currency"`
}

// MediaReviewItem is a generated struct representing the https://schema.org/MediaReviewItem class.
type MediaReviewItem struct {
	MediaItemAppearance *MediaObject `json:"mediaitemappearance"`
}

// DepartAction is a generated struct representing the https://schema.org/DepartAction class.
type DepartAction struct {
}

// TextDigitalDocument is a generated struct representing the https://schema.org/TextDigitalDocument class.
type TextDigitalDocument struct {
}

// VideoObjectSnapshot is a generated struct representing the https://schema.org/VideoObjectSnapshot class.
type VideoObjectSnapshot struct {
}

// OccupationalTherapy is a generated struct representing the https://schema.org/OccupationalTherapy class.
type OccupationalTherapy struct {
}

// SportsClub is a generated struct representing the https://schema.org/SportsClub class.
type SportsClub struct {
}

// Vehicle is a generated struct representing the https://schema.org/Vehicle class.
type Vehicle struct {
	AccelerationTime            *QuantitativeValue     `json:"accelerationtime"`
	BodyType                    *URL                   `json:"bodytype"`
	CallSign                    *Text                  `json:"callsign"`
	CargoVolume                 *QuantitativeValue     `json:"cargovolume"`
	DateVehicleFirstRegistered  *Date                  `json:"datevehiclefirstregistered"`
	DriveWheelConfiguration     *Text                  `json:"drivewheelconfiguration"`
	EmissionsCO2                *Number                `json:"emissionsco2"`
	FuelCapacity                *QuantitativeValue     `json:"fuelcapacity"`
	FuelConsumption             *QuantitativeValue     `json:"fuelconsumption"`
	FuelEfficiency              *QuantitativeValue     `json:"fuelefficiency"`
	FuelType                    *URL                   `json:"fueltype"`
	KnownVehicleDamages         *Text                  `json:"knownvehicledamages"`
	MeetsEmissionStandard       *URL                   `json:"meetsemissionstandard"`
	MileageFromOdometer         *QuantitativeValue     `json:"mileagefromodometer"`
	ModelDate                   *Date                  `json:"modeldate"`
	NumberOfAirbags             *Text                  `json:"numberofairbags"`
	NumberOfAxles               *QuantitativeValue     `json:"numberofaxles"`
	NumberOfDoors               *QuantitativeValue     `json:"numberofdoors"`
	NumberOfForwardGears        *QuantitativeValue     `json:"numberofforwardgears"`
	NumberOfPreviousOwners      *QuantitativeValue     `json:"numberofpreviousowners"`
	Payload                     *QuantitativeValue     `json:"payload"`
	ProductionDate              *Date                  `json:"productiondate"`
	PurchaseDate                *Date                  `json:"purchasedate"`
	SeatingCapacity             *QuantitativeValue     `json:"seatingcapacity"`
	Speed                       *QuantitativeValue     `json:"speed"`
	SteeringPosition            *SteeringPositionValue `json:"steeringposition"`
	TongueWeight                *QuantitativeValue     `json:"tongueweight"`
	TrailerWeight               *QuantitativeValue     `json:"trailerweight"`
	VehicleConfiguration        *Text                  `json:"vehicleconfiguration"`
	VehicleEngine               *EngineSpecification   `json:"vehicleengine"`
	VehicleIdentificationNumber *Text                  `json:"vehicleidentificationnumber"`
	VehicleInteriorColor        *Text                  `json:"vehicleinteriorcolor"`
	VehicleInteriorType         *Text                  `json:"vehicleinteriortype"`
	VehicleModelDate            *Date                  `json:"vehiclemodeldate"`
	VehicleSeatingCapacity      *QuantitativeValue     `json:"vehicleseatingcapacity"`
	VehicleSpecialUsage         *Text                  `json:"vehiclespecialusage"`
	VehicleTransmission         *URL                   `json:"vehicletransmission"`
	WeightTotal                 *QuantitativeValue     `json:"weighttotal"`
	Wheelbase                   *QuantitativeValue     `json:"wheelbase"`
}

// FilmAction is a generated struct representing the https://schema.org/FilmAction class.
type FilmAction struct {
}

// GroceryStore is a generated struct representing the https://schema.org/GroceryStore class.
type GroceryStore struct {
}

// ShippingRateSettings is a generated struct representing the https://schema.org/ShippingRateSettings class.
type ShippingRateSettings struct {
	DoesNotShip           *Boolean        `json:"doesnotship"`
	FreeShippingThreshold *MonetaryAmount `json:"freeshippingthreshold"`
	IsUnlabelledFallback  *Boolean        `json:"isunlabelledfallback"`
	ShippingDestination   *DefinedRegion  `json:"shippingdestination"`
	ShippingLabel         *Text           `json:"shippinglabel"`
	ShippingRate          *MonetaryAmount `json:"shippingrate"`
}

// AutomatedTeller is a generated struct representing the https://schema.org/AutomatedTeller class.
type AutomatedTeller struct {
}

// EventAttendanceModeEnumeration is a generated struct representing the https://schema.org/EventAttendanceModeEnumeration class.
type EventAttendanceModeEnumeration struct {
}

// GovernmentOrganization is a generated struct representing the https://schema.org/GovernmentOrganization class.
type GovernmentOrganization struct {
}

// Audience is a generated struct representing the https://schema.org/Audience class.
type Audience struct {
	AudienceType   *Text               `json:"audiencetype"`
	GeographicArea *AdministrativeArea `json:"geographicarea"`
}

// Park is a generated struct representing the https://schema.org/Park class.
type Park struct {
}

// ElementarySchool is a generated struct representing the https://schema.org/ElementarySchool class.
type ElementarySchool struct {
}

// EntryPoint is a generated struct representing the https://schema.org/EntryPoint class.
type EntryPoint struct {
	ActionPlatform    *URL                 `json:"actionplatform"`
	Application       *SoftwareApplication `json:"application"`
	ContentType       *Text                `json:"contenttype"`
	EncodingType      *Text                `json:"encodingtype"`
	HttpMethod        *Text                `json:"httpmethod"`
	UrlTemplate       *Text                `json:"urltemplate"`
	ActionApplication *SoftwareApplication `json:"actionapplication"`
}

// MarryAction is a generated struct representing the https://schema.org/MarryAction class.
type MarryAction struct {
}

// MedicalOrganization is a generated struct representing the https://schema.org/MedicalOrganization class.
type MedicalOrganization struct {
	HealthPlanNetworkId    *Text             `json:"healthplannetworkid"`
	IsAcceptingNewPatients *Boolean          `json:"isacceptingnewpatients"`
	MedicalSpecialty       *MedicalSpecialty `json:"medicalspecialty"`
}

// ReturnAction is a generated struct representing the https://schema.org/ReturnAction class.
type ReturnAction struct {
	Recipient *Person `json:"recipient"`
}

// AchieveAction is a generated struct representing the https://schema.org/AchieveAction class.
type AchieveAction struct {
}

// HowToSection is a generated struct representing the https://schema.org/HowToSection class.
type HowToSection struct {
	Steps *Text `json:"steps"`
}

// NLNonprofitType is a generated struct representing the https://schema.org/NLNonprofitType class.
type NLNonprofitType struct {
}

// PrependAction is a generated struct representing the https://schema.org/PrependAction class.
type PrependAction struct {
}

// RegisterAction is a generated struct representing the https://schema.org/RegisterAction class.
type RegisterAction struct {
}

// SolveMathAction is a generated struct representing the https://schema.org/SolveMathAction class.
type SolveMathAction struct {
	EduQuestionType *Text `json:"eduquestiontype"`
}

// ArriveAction is a generated struct representing the https://schema.org/ArriveAction class.
type ArriveAction struct {
}

// MedicalClinic is a generated struct representing the https://schema.org/MedicalClinic class.
type MedicalClinic struct {
	AvailableService *MedicalTherapy   `json:"availableservice"`
	MedicalSpecialty *MedicalSpecialty `json:"medicalspecialty"`
}

// OrganizeAction is a generated struct representing the https://schema.org/OrganizeAction class.
type OrganizeAction struct {
}

// Statement is a generated struct representing the https://schema.org/Statement class.
type Statement struct {
}

// Thesis is a generated struct representing the https://schema.org/Thesis class.
type Thesis struct {
	InSupportOf *Text `json:"insupportof"`
}

// UserInteraction is a generated struct representing the https://schema.org/UserInteraction class.
type UserInteraction struct {
}

// BlogPosting is a generated struct representing the https://schema.org/BlogPosting class.
type BlogPosting struct {
}

// CDCPMDRecord is a generated struct representing the https://schema.org/CDCPMDRecord class.
type CDCPMDRecord struct {
	CvdCollectionDate       *Text     `json:"cvdcollectiondate"`
	CvdFacilityCounty       *Text     `json:"cvdfacilitycounty"`
	CvdFacilityId           *Text     `json:"cvdfacilityid"`
	CvdNumBeds              *Number   `json:"cvdnumbeds"`
	CvdNumBedsOcc           *Number   `json:"cvdnumbedsocc"`
	CvdNumC19Died           *Number   `json:"cvdnumc19died"`
	CvdNumC19HOPats         *Number   `json:"cvdnumc19hopats"`
	CvdNumC19HospPats       *Number   `json:"cvdnumc19hosppats"`
	CvdNumC19MechVentPats   *Number   `json:"cvdnumc19mechventpats"`
	CvdNumC19OFMechVentPats *Number   `json:"cvdnumc19ofmechventpats"`
	CvdNumC19OverflowPats   *Number   `json:"cvdnumc19overflowpats"`
	CvdNumICUBeds           *Number   `json:"cvdnumicubeds"`
	CvdNumICUBedsOcc        *Number   `json:"cvdnumicubedsocc"`
	CvdNumTotBeds           *Number   `json:"cvdnumtotbeds"`
	CvdNumVent              *Number   `json:"cvdnumvent"`
	CvdNumVentUse           *Number   `json:"cvdnumventuse"`
	DatePosted              *DateTime `json:"dateposted"`
}

// DeliveryEvent is a generated struct representing the https://schema.org/DeliveryEvent class.
type DeliveryEvent struct {
	AccessCode        *Text           `json:"accesscode"`
	AvailableFrom     *DateTime       `json:"availablefrom"`
	AvailableThrough  *DateTime       `json:"availablethrough"`
	HasDeliveryMethod *DeliveryMethod `json:"hasdeliverymethod"`
}

// UserPlusOnes is a generated struct representing the https://schema.org/UserPlusOnes class.
type UserPlusOnes struct {
}

// EventStatusType is a generated struct representing the https://schema.org/EventStatusType class.
type EventStatusType struct {
}

// Ligament is a generated struct representing the https://schema.org/Ligament class.
type Ligament struct {
}

// MortgageLoan is a generated struct representing the https://schema.org/MortgageLoan class.
type MortgageLoan struct {
	DomiciledMortgage         *Boolean        `json:"domiciledmortgage"`
	LoanMortgageMandateAmount *MonetaryAmount `json:"loanmortgagemandateamount"`
}

// VacationRental is a generated struct representing the https://schema.org/VacationRental class.
type VacationRental struct {
}

// MediaObject is a generated struct representing the https://schema.org/MediaObject class.
type MediaObject struct {
	AssociatedArticle    *NewsArticle       `json:"associatedarticle"`
	Bitrate              *Text              `json:"bitrate"`
	ContentSize          *Text              `json:"contentsize"`
	ContentUrl           *URL               `json:"contenturl"`
	EmbedUrl             *URL               `json:"embedurl"`
	EndTime              *Time              `json:"endtime"`
	Height               *QuantitativeValue `json:"height"`
	IneligibleRegion     *Text              `json:"ineligibleregion"`
	InterpretedAsClaim   *Claim             `json:"interpretedasclaim"`
	PlayerType           *Text              `json:"playertype"`
	ProductionCompany    *Organization      `json:"productioncompany"`
	RegionsAllowed       *Place             `json:"regionsallowed"`
	RequiresSubscription *MediaSubscription `json:"requiressubscription"`
	Sha256               *Text              `json:"sha256"`
	StartTime            *Time              `json:"starttime"`
	UploadDate           *DateTime          `json:"uploaddate"`
	Width                *QuantitativeValue `json:"width"`
	Duration             *Duration          `json:"duration"`
	EncodesCreativeWork  *CreativeWork      `json:"encodescreativework"`
	EncodingFormat       *URL               `json:"encodingformat"`
}

// Sculpture is a generated struct representing the https://schema.org/Sculpture class.
type Sculpture struct {
}

// Specialty is a generated struct representing the https://schema.org/Specialty class.
type Specialty struct {
}

// PaymentService is a generated struct representing the https://schema.org/PaymentService class.
type PaymentService struct {
}

// Newspaper is a generated struct representing the https://schema.org/Newspaper class.
type Newspaper struct {
}

// ReturnMethodEnumeration is a generated struct representing the https://schema.org/ReturnMethodEnumeration class.
type ReturnMethodEnumeration struct {
}

// ShoppingCenter is a generated struct representing the https://schema.org/ShoppingCenter class.
type ShoppingCenter struct {
}

// BarOrPub is a generated struct representing the https://schema.org/BarOrPub class.
type BarOrPub struct {
}

// BedDetails is a generated struct representing the https://schema.org/BedDetails class.
type BedDetails struct {
	NumberOfBeds *Number `json:"numberofbeds"`
	TypeOfBed    *Text   `json:"typeofbed"`
}

// BoatTerminal is a generated struct representing the https://schema.org/BoatTerminal class.
type BoatTerminal struct {
}

// BrainStructure is a generated struct representing the https://schema.org/BrainStructure class.
type BrainStructure struct {
}

// ReviewNewsArticle is a generated struct representing the https://schema.org/ReviewNewsArticle class.
type ReviewNewsArticle struct {
}

// Taxi is a generated struct representing the https://schema.org/Taxi class.
type Taxi struct {
}

// Brand is a generated struct representing the https://schema.org/Brand class.
type Brand struct {
	AggregateRating *AggregateRating `json:"aggregaterating"`
	Logo            *URL             `json:"logo"`
	Slogan          *Text            `json:"slogan"`
	Review          *Review          `json:"review"`
}

// EnergyEfficiencyEnumeration is a generated struct representing the https://schema.org/EnergyEfficiencyEnumeration class.
type EnergyEfficiencyEnumeration struct {
}

// Menu is a generated struct representing the https://schema.org/Menu class.
type Menu struct {
	HasMenuItem    *MenuItem    `json:"hasmenuitem"`
	HasMenuSection *MenuSection `json:"hasmenusection"`
}

// SoftwareSourceCode is a generated struct representing the https://schema.org/SoftwareSourceCode class.
type SoftwareSourceCode struct {
	CodeRepository      *URL                 `json:"coderepository"`
	ProgrammingLanguage *Text                `json:"programminglanguage"`
	Runtime             *Text                `json:"runtime"`
	SampleType          *Text                `json:"sampletype"`
	TargetProduct       *SoftwareApplication `json:"targetproduct"`
	CodeSampleType      *Text                `json:"codesampletype"`
	RuntimePlatform     *Text                `json:"runtimeplatform"`
}

// Game is a generated struct representing the https://schema.org/Game class.
type Game struct {
	CharacterAttribute *Thing             `json:"characterattribute"`
	GameItem           *Thing             `json:"gameitem"`
	GameLocation       *URL               `json:"gamelocation"`
	NumberOfPlayers    *QuantitativeValue `json:"numberofplayers"`
	Quest              *Thing             `json:"quest"`
}

// GenderType is a generated struct representing the https://schema.org/GenderType class.
type GenderType struct {
}

// MerchantReturnPolicySeasonalOverride is a generated struct representing the https://schema.org/MerchantReturnPolicySeasonalOverride class.
type MerchantReturnPolicySeasonalOverride struct {
	EndDate              *DateTime                  `json:"enddate"`
	MerchantReturnDays   *Integer                   `json:"merchantreturndays"`
	ReturnPolicyCategory *MerchantReturnEnumeration `json:"returnpolicycategory"`
	StartDate            *DateTime                  `json:"startdate"`
}

// BrokerageAccount is a generated struct representing the https://schema.org/BrokerageAccount class.
type BrokerageAccount struct {
}

// HealthInsurancePlan is a generated struct representing the https://schema.org/HealthInsurancePlan class.
type HealthInsurancePlan struct {
	BenefitsSummaryUrl          *URL                 `json:"benefitssummaryurl"`
	HealthPlanDrugOption        *Text                `json:"healthplandrugoption"`
	HealthPlanDrugTier          *Text                `json:"healthplandrugtier"`
	HealthPlanId                *Text                `json:"healthplanid"`
	HealthPlanMarketingUrl      *URL                 `json:"healthplanmarketingurl"`
	IncludesHealthPlanFormulary *HealthPlanFormulary `json:"includeshealthplanformulary"`
	IncludesHealthPlanNetwork   *HealthPlanNetwork   `json:"includeshealthplannetwork"`
	UsesHealthPlanIdStandard    *URL                 `json:"useshealthplanidstandard"`
	ContactPoint                *ContactPoint        `json:"contactpoint"`
}

// InfectiousAgentClass is a generated struct representing the https://schema.org/InfectiousAgentClass class.
type InfectiousAgentClass struct {
}

// TrackAction is a generated struct representing the https://schema.org/TrackAction class.
type TrackAction struct {
	DeliveryMethod *DeliveryMethod `json:"deliverymethod"`
}

// UKNonprofitType is a generated struct representing the https://schema.org/UKNonprofitType class.
type UKNonprofitType struct {
}

// AutoRepair is a generated struct representing the https://schema.org/AutoRepair class.
type AutoRepair struct {
}

// MedicalGuideline is a generated struct representing the https://schema.org/MedicalGuideline class.
type MedicalGuideline struct {
	EvidenceLevel    *MedicalEvidenceLevel `json:"evidencelevel"`
	EvidenceOrigin   *Text                 `json:"evidenceorigin"`
	GuidelineDate    *Date                 `json:"guidelinedate"`
	GuidelineSubject *MedicalEntity        `json:"guidelinesubject"`
}

// SuperficialAnatomy is a generated struct representing the https://schema.org/SuperficialAnatomy class.
type SuperficialAnatomy struct {
	AssociatedPathophysiology *Text             `json:"associatedpathophysiology"`
	RelatedAnatomy            *AnatomicalSystem `json:"relatedanatomy"`
	RelatedCondition          *MedicalCondition `json:"relatedcondition"`
	RelatedTherapy            *MedicalTherapy   `json:"relatedtherapy"`
	Significance              *Text             `json:"significance"`
}

// Barcode is a generated struct representing the https://schema.org/Barcode class.
type Barcode struct {
}

// BroadcastService is a generated struct representing the https://schema.org/BroadcastService class.
type BroadcastService struct {
	Area                 *Place            `json:"area"`
	BroadcastAffiliateOf *Organization     `json:"broadcastaffiliateof"`
	BroadcastDisplayName *Text             `json:"broadcastdisplayname"`
	BroadcastFrequency   *Text             `json:"broadcastfrequency"`
	BroadcastTimezone    *Text             `json:"broadcasttimezone"`
	Broadcaster          *Organization     `json:"broadcaster"`
	CallSign             *Text             `json:"callsign"`
	ParentService        *BroadcastService `json:"parentservice"`
	VideoFormat          *Text             `json:"videoformat"`
	HasBroadcastChannel  *BroadcastChannel `json:"hasbroadcastchannel"`
	InLanguage           *Text             `json:"inlanguage"`
}

// DietarySupplement is a generated struct representing the https://schema.org/DietarySupplement class.
type DietarySupplement struct {
	ActiveIngredient    *Text                    `json:"activeingredient"`
	IsProprietary       *Boolean                 `json:"isproprietary"`
	LegalStatus         *Text                    `json:"legalstatus"`
	MaximumIntake       *MaximumDoseSchedule     `json:"maximumintake"`
	MechanismOfAction   *Text                    `json:"mechanismofaction"`
	NonProprietaryName  *Text                    `json:"nonproprietaryname"`
	ProprietaryName     *Text                    `json:"proprietaryname"`
	RecommendedIntake   *RecommendedDoseSchedule `json:"recommendedintake"`
	SafetyConsideration *Text                    `json:"safetyconsideration"`
	TargetPopulation    *Text                    `json:"targetpopulation"`
}

// PostalAddress is a generated struct representing the https://schema.org/PostalAddress class.
type PostalAddress struct {
	AddressCountry      *Text `json:"addresscountry"`
	AddressLocality     *Text `json:"addresslocality"`
	AddressRegion       *Text `json:"addressregion"`
	PostOfficeBoxNumber *Text `json:"postofficeboxnumber"`
	PostalCode          *Text `json:"postalcode"`
	StreetAddress       *Text `json:"streetaddress"`
}

// InternetCafe is a generated struct representing the https://schema.org/InternetCafe class.
type InternetCafe struct {
}

// MapCategoryType is a generated struct representing the https://schema.org/MapCategoryType class.
type MapCategoryType struct {
}

// ParentAudience is a generated struct representing the https://schema.org/ParentAudience class.
type ParentAudience struct {
	ChildMaxAge *Number `json:"childmaxage"`
	ChildMinAge *Number `json:"childminage"`
}

// Review is a generated struct representing the https://schema.org/Review class.
type Review struct {
	AssociatedClaimReview *Review     `json:"associatedclaimreview"`
	AssociatedMediaReview *Review     `json:"associatedmediareview"`
	ItemReviewed          *Thing      `json:"itemreviewed"`
	NegativeNotes         *WebContent `json:"negativenotes"`
	PositiveNotes         *WebContent `json:"positivenotes"`
	ReviewAspect          *Text       `json:"reviewaspect"`
	ReviewBody            *Text       `json:"reviewbody"`
	ReviewRating          *Rating     `json:"reviewrating"`
	AssociatedReview      *Review     `json:"associatedreview"`
}

// VoteAction is a generated struct representing the https://schema.org/VoteAction class.
type VoteAction struct {
	Candidate *Person `json:"candidate"`
}

// GameServer is a generated struct representing the https://schema.org/GameServer class.
type GameServer struct {
	PlayersOnline *Integer          `json:"playersonline"`
	ServerStatus  *GameServerStatus `json:"serverstatus"`
	Game          *VideoGame        `json:"game"`
}

// ProductGroup is a generated struct representing the https://schema.org/ProductGroup class.
type ProductGroup struct {
	ProductGroupID *Text    `json:"productgroupid"`
	VariesBy       *Text    `json:"variesby"`
	HasVariant     *Product `json:"hasvariant"`
}

// TheaterEvent is a generated struct representing the https://schema.org/TheaterEvent class.
type TheaterEvent struct {
}

// EmployerReview is a generated struct representing the https://schema.org/EmployerReview class.
type EmployerReview struct {
}

// SoftwareApplication is a generated struct representing the https://schema.org/SoftwareApplication class.
type SoftwareApplication struct {
	ApplicationCategory    *URL                 `json:"applicationcategory"`
	ApplicationSubCategory *URL                 `json:"applicationsubcategory"`
	ApplicationSuite       *Text                `json:"applicationsuite"`
	CountriesNotSupported  *Text                `json:"countriesnotsupported"`
	CountriesSupported     *Text                `json:"countriessupported"`
	Device                 *Text                `json:"device"`
	DownloadUrl            *URL                 `json:"downloadurl"`
	FeatureList            *URL                 `json:"featurelist"`
	FileSize               *Text                `json:"filesize"`
	InstallUrl             *URL                 `json:"installurl"`
	MemoryRequirements     *URL                 `json:"memoryrequirements"`
	OperatingSystem        *Text                `json:"operatingsystem"`
	Permissions            *Text                `json:"permissions"`
	ProcessorRequirements  *Text                `json:"processorrequirements"`
	ReleaseNotes           *URL                 `json:"releasenotes"`
	Requirements           *URL                 `json:"requirements"`
	Screenshot             *URL                 `json:"screenshot"`
	SoftwareAddOn          *SoftwareApplication `json:"softwareaddon"`
	SoftwareHelp           *CreativeWork        `json:"softwarehelp"`
	SoftwareVersion        *Text                `json:"softwareversion"`
	StorageRequirements    *URL                 `json:"storagerequirements"`
	SupportingData         *DataFeed            `json:"supportingdata"`
	AvailableOnDevice      *Text                `json:"availableondevice"`
	SoftwareRequirements   *URL                 `json:"softwarerequirements"`
}

// DrinkAction is a generated struct representing the https://schema.org/DrinkAction class.
type DrinkAction struct {
}

// MenuSection is a generated struct representing the https://schema.org/MenuSection class.
type MenuSection struct {
	HasMenuItem    *MenuItem    `json:"hasmenuitem"`
	HasMenuSection *MenuSection `json:"hasmenusection"`
}

// TattooParlor is a generated struct representing the https://schema.org/TattooParlor class.
type TattooParlor struct {
}

// AppendAction is a generated struct representing the https://schema.org/AppendAction class.
type AppendAction struct {
}

// BroadcastChannel is a generated struct representing the https://schema.org/BroadcastChannel class.
type BroadcastChannel struct {
	BroadcastChannelId       *Text                    `json:"broadcastchannelid"`
	BroadcastFrequency       *Text                    `json:"broadcastfrequency"`
	BroadcastServiceTier     *Text                    `json:"broadcastservicetier"`
	InBroadcastLineup        *CableOrSatelliteService `json:"inbroadcastlineup"`
	Genre                    *URL                     `json:"genre"`
	ProvidesBroadcastService *BroadcastService        `json:"providesbroadcastservice"`
}

// PawnShop is a generated struct representing the https://schema.org/PawnShop class.
type PawnShop struct {
}

// Playground is a generated struct representing the https://schema.org/Playground class.
type Playground struct {
}

// QuantitativeValueDistribution is a generated struct representing the https://schema.org/QuantitativeValueDistribution class.
type QuantitativeValueDistribution struct {
	Median       *Number   `json:"median"`
	Percentile10 *Number   `json:"percentile10"`
	Percentile25 *Number   `json:"percentile25"`
	Percentile75 *Number   `json:"percentile75"`
	Percentile90 *Number   `json:"percentile90"`
	Duration     *Duration `json:"duration"`
}

// CovidTestingFacility is a generated struct representing the https://schema.org/CovidTestingFacility class.
type CovidTestingFacility struct {
}

// DiagnosticProcedure is a generated struct representing the https://schema.org/DiagnosticProcedure class.
type DiagnosticProcedure struct {
}

// Notary is a generated struct representing the https://schema.org/Notary class.
type Notary struct {
}

// Reservoir is a generated struct representing the https://schema.org/Reservoir class.
type Reservoir struct {
}

// UserPlays is a generated struct representing the https://schema.org/UserPlays class.
type UserPlays struct {
}

// ComicSeries is a generated struct representing the https://schema.org/ComicSeries class.
type ComicSeries struct {
}

// VideoObject is a generated struct representing the https://schema.org/VideoObject class.
type VideoObject struct {
	Actors              *Person `json:"actors"`
	Directors           *Person `json:"directors"`
	EmbeddedTextCaption *Text   `json:"embeddedtextcaption"`
	MusicBy             *Person `json:"musicby"`
	Transcript          *Text   `json:"transcript"`
	VideoFrameSize      *Text   `json:"videoframesize"`
	VideoQuality        *Text   `json:"videoquality"`
	Caption             *Text   `json:"caption"`
	Director            *Person `json:"director"`
	Actor               *Person `json:"actor"`
}

// Flight is a generated struct representing the https://schema.org/Flight class.
type Flight struct {
	Aircraft                *Vehicle            `json:"aircraft"`
	ArrivalAirport          *Airport            `json:"arrivalairport"`
	ArrivalGate             *Text               `json:"arrivalgate"`
	ArrivalTerminal         *Text               `json:"arrivalterminal"`
	BoardingPolicy          *BoardingPolicyType `json:"boardingpolicy"`
	Carrier                 *Organization       `json:"carrier"`
	DepartureAirport        *Airport            `json:"departureairport"`
	DepartureGate           *Text               `json:"departuregate"`
	DepartureTerminal       *Text               `json:"departureterminal"`
	EstimatedFlightDuration *Text               `json:"estimatedflightduration"`
	FlightDistance          *Text               `json:"flightdistance"`
	FlightNumber            *Text               `json:"flightnumber"`
	MealService             *Text               `json:"mealservice"`
	WebCheckinTime          *DateTime           `json:"webcheckintime"`
	Seller                  *Person             `json:"seller"`
}

// HealthPlanFormulary is a generated struct representing the https://schema.org/HealthPlanFormulary class.
type HealthPlanFormulary struct {
	HealthPlanCostSharing    *Boolean `json:"healthplancostsharing"`
	HealthPlanDrugTier       *Text    `json:"healthplandrugtier"`
	OffersPrescriptionByMail *Boolean `json:"offersprescriptionbymail"`
}

// Resort is a generated struct representing the https://schema.org/Resort class.
type Resort struct {
}

// UserBlocks is a generated struct representing the https://schema.org/UserBlocks class.
type UserBlocks struct {
}

// ApprovedIndication is a generated struct representing the https://schema.org/ApprovedIndication class.
type ApprovedIndication struct {
}

// AskPublicNewsArticle is a generated struct representing the https://schema.org/AskPublicNewsArticle class.
type AskPublicNewsArticle struct {
}

// BusinessFunction is a generated struct representing the https://schema.org/BusinessFunction class.
type BusinessFunction struct {
}

// NewsMediaOrganization is a generated struct representing the https://schema.org/NewsMediaOrganization class.
type NewsMediaOrganization struct {
	ActionableFeedbackPolicy        *URL `json:"actionablefeedbackpolicy"`
	CorrectionsPolicy               *URL `json:"correctionspolicy"`
	DiversityPolicy                 *URL `json:"diversitypolicy"`
	DiversityStaffingReport         *URL `json:"diversitystaffingreport"`
	EthicsPolicy                    *URL `json:"ethicspolicy"`
	Masthead                        *URL `json:"masthead"`
	MissionCoveragePrioritiesPolicy *URL `json:"missioncoverageprioritiespolicy"`
	NoBylinesPolicy                 *URL `json:"nobylinespolicy"`
	OwnershipFundingInfo            *URL `json:"ownershipfundinginfo"`
	UnnamedSourcesPolicy            *URL `json:"unnamedsourcespolicy"`
	VerificationFactCheckingPolicy  *URL `json:"verificationfactcheckingpolicy"`
}

// AskAction is a generated struct representing the https://schema.org/AskAction class.
type AskAction struct {
	Question *Question `json:"question"`
}

// CarUsageType is a generated struct representing the https://schema.org/CarUsageType class.
type CarUsageType struct {
}

// Embassy is a generated struct representing the https://schema.org/Embassy class.
type Embassy struct {
}

// LodgingBusiness is a generated struct representing the https://schema.org/LodgingBusiness class.
type LodgingBusiness struct {
	AmenityFeature    *LocationFeatureSpecification `json:"amenityfeature"`
	AvailableLanguage *Text                         `json:"availablelanguage"`
	CheckinTime       *Time                         `json:"checkintime"`
	CheckoutTime      *Time                         `json:"checkouttime"`
	NumberOfRooms     *QuantitativeValue            `json:"numberofrooms"`
	PetsAllowed       *Text                         `json:"petsallowed"`
	StarRating        *Rating                       `json:"starrating"`
	Audience          *Audience                     `json:"audience"`
}

// ChemicalSubstance is a generated struct representing the https://schema.org/ChemicalSubstance class.
type ChemicalSubstance struct {
	ChemicalComposition *Text        `json:"chemicalcomposition"`
	ChemicalRole        *DefinedTerm `json:"chemicalrole"`
	PotentialUse        *DefinedTerm `json:"potentialuse"`
}

// HobbyShop is a generated struct representing the https://schema.org/HobbyShop class.
type HobbyShop struct {
}

// LocationFeatureSpecification is a generated struct representing the https://schema.org/LocationFeatureSpecification class.
type LocationFeatureSpecification struct {
	HoursAvailable *OpeningHoursSpecification `json:"hoursavailable"`
	ValidFrom      *DateTime                  `json:"validfrom"`
	ValidThrough   *DateTime                  `json:"validthrough"`
}

// MedicalCause is a generated struct representing the https://schema.org/MedicalCause class.
type MedicalCause struct {
	CauseOf *MedicalEntity `json:"causeof"`
}

// Car is a generated struct representing the https://schema.org/Car class.
type Car struct {
	AcrissCode *Text              `json:"acrisscode"`
	RoofLoad   *QuantitativeValue `json:"roofload"`
}

// DefinedTerm is a generated struct representing the https://schema.org/DefinedTerm class.
type DefinedTerm struct {
	InDefinedTermSet *URL  `json:"indefinedtermset"`
	TermCode         *Text `json:"termcode"`
}

// DriveWheelConfigurationValue is a generated struct representing the https://schema.org/DriveWheelConfigurationValue class.
type DriveWheelConfigurationValue struct {
}

// MedicalSign is a generated struct representing the https://schema.org/MedicalSign class.
type MedicalSign struct {
	IdentifyingExam *PhysicalExam `json:"identifyingexam"`
	IdentifyingTest *MedicalTest  `json:"identifyingtest"`
}

// SocialEvent is a generated struct representing the https://schema.org/SocialEvent class.
type SocialEvent struct {
}

// DefinedRegion is a generated struct representing the https://schema.org/DefinedRegion class.
type DefinedRegion struct {
	AddressCountry   *Text                         `json:"addresscountry"`
	AddressRegion    *Text                         `json:"addressregion"`
	PostalCode       *Text                         `json:"postalcode"`
	PostalCodePrefix *Text                         `json:"postalcodeprefix"`
	PostalCodeRange  *PostalCodeRangeSpecification `json:"postalcoderange"`
}

// PeopleAudience is a generated struct representing the https://schema.org/PeopleAudience class.
type PeopleAudience struct {
	HealthCondition      *MedicalCondition  `json:"healthcondition"`
	RequiredGender       *Text              `json:"requiredgender"`
	RequiredMaxAge       *Integer           `json:"requiredmaxage"`
	RequiredMinAge       *Integer           `json:"requiredminage"`
	SuggestedAge         *QuantitativeValue `json:"suggestedage"`
	SuggestedGender      *Text              `json:"suggestedgender"`
	SuggestedMaxAge      *Number            `json:"suggestedmaxage"`
	SuggestedMeasurement *QuantitativeValue `json:"suggestedmeasurement"`
	SuggestedMinAge      *Number            `json:"suggestedminage"`
}

// EventVenue is a generated struct representing the https://schema.org/EventVenue class.
type EventVenue struct {
}

// ArtGallery is a generated struct representing the https://schema.org/ArtGallery class.
type ArtGallery struct {
}

// BorrowAction is a generated struct representing the https://schema.org/BorrowAction class.
type BorrowAction struct {
	Lender *Person `json:"lender"`
}

// Play is a generated struct representing the https://schema.org/Play class.
type Play struct {
}

// Audiobook is a generated struct representing the https://schema.org/Audiobook class.
type Audiobook struct {
	ReadBy   *Person   `json:"readby"`
	Duration *Duration `json:"duration"`
}

// Occupation is a generated struct representing the https://schema.org/Occupation class.
type Occupation struct {
	EducationRequirements  *Text               `json:"educationrequirements"`
	EstimatedSalary        *Number             `json:"estimatedsalary"`
	ExperienceRequirements *Text               `json:"experiencerequirements"`
	OccupationLocation     *AdministrativeArea `json:"occupationlocation"`
	OccupationalCategory   *Text               `json:"occupationalcategory"`
	Qualifications         *Text               `json:"qualifications"`
	Responsibilities       *Text               `json:"responsibilities"`
	Skills                 *Text               `json:"skills"`
}

// Recipe is a generated struct representing the https://schema.org/Recipe class.
type Recipe struct {
	CookTime           *Duration             `json:"cooktime"`
	CookingMethod      *Text                 `json:"cookingmethod"`
	Ingredients        *Text                 `json:"ingredients"`
	Nutrition          *NutritionInformation `json:"nutrition"`
	RecipeCategory     *Text                 `json:"recipecategory"`
	RecipeCuisine      *Text                 `json:"recipecuisine"`
	RecipeInstructions *Text                 `json:"recipeinstructions"`
	RecipeYield        *Text                 `json:"recipeyield"`
	SuitableForDiet    *RestrictedDiet       `json:"suitablefordiet"`
	RecipeIngredient   *Text                 `json:"recipeingredient"`
}

// BoatReservation is a generated struct representing the https://schema.org/BoatReservation class.
type BoatReservation struct {
}

// BoatTrip is a generated struct representing the https://schema.org/BoatTrip class.
type BoatTrip struct {
	ArrivalBoatTerminal   *BoatTerminal `json:"arrivalboatterminal"`
	DepartureBoatTerminal *BoatTerminal `json:"departureboatterminal"`
}

// Float is a generated struct representing the https://schema.org/Float class.
type Float struct {
}

// NailSalon is a generated struct representing the https://schema.org/NailSalon class.
type NailSalon struct {
}

// Bakery is a generated struct representing the https://schema.org/Bakery class.
type Bakery struct {
}

// Message is a generated struct representing the https://schema.org/Message class.
type Message struct {
	BccRecipient      *Person       `json:"bccrecipient"`
	CcRecipient       *Person       `json:"ccrecipient"`
	DateRead          *DateTime     `json:"dateread"`
	DateReceived      *DateTime     `json:"datereceived"`
	DateSent          *DateTime     `json:"datesent"`
	MessageAttachment *CreativeWork `json:"messageattachment"`
	Sender            *Person       `json:"sender"`
	ToRecipient       *Person       `json:"torecipient"`
	Recipient         *Person       `json:"recipient"`
}

// PerformingArtsTheater is a generated struct representing the https://schema.org/PerformingArtsTheater class.
type PerformingArtsTheater struct {
}

// Time is a generated struct representing the https://schema.org/Time class.
type Time struct {
}

// Date is a generated struct representing the https://schema.org/Date class.
type Date struct {
}

// EducationalOccupationalProgram is a generated struct representing the https://schema.org/EducationalOccupationalProgram class.
type EducationalOccupationalProgram struct {
	ApplicationDeadline           *Date                       `json:"applicationdeadline"`
	ApplicationStartDate          *Date                       `json:"applicationstartdate"`
	DayOfWeek                     *DayOfWeek                  `json:"dayofweek"`
	EducationalCredentialAwarded  *URL                        `json:"educationalcredentialawarded"`
	EducationalProgramMode        *URL                        `json:"educationalprogrammode"`
	EndDate                       *DateTime                   `json:"enddate"`
	FinancialAidEligible          *Text                       `json:"financialaideligible"`
	HasCourse                     *Course                     `json:"hascourse"`
	MaximumEnrollment             *Integer                    `json:"maximumenrollment"`
	NumberOfCredits               *StructuredValue            `json:"numberofcredits"`
	OccupationalCategory          *Text                       `json:"occupationalcategory"`
	OccupationalCredentialAwarded *URL                        `json:"occupationalcredentialawarded"`
	ProgramPrerequisites          *Text                       `json:"programprerequisites"`
	ProgramType                   *Text                       `json:"programtype"`
	SalaryUponCompletion          *MonetaryAmountDistribution `json:"salaryuponcompletion"`
	StartDate                     *DateTime                   `json:"startdate"`
	TermDuration                  *Duration                   `json:"termduration"`
	TermsPerYear                  *Number                     `json:"termsperyear"`
	TimeOfDay                     *Text                       `json:"timeofday"`
	TimeToComplete                *Duration                   `json:"timetocomplete"`
	TrainingSalary                *MonetaryAmountDistribution `json:"trainingsalary"`
	TypicalCreditsPerTerm         *StructuredValue            `json:"typicalcreditsperterm"`
	Offers                        *Offer                      `json:"offers"`
	Provider                      *Person                     `json:"provider"`
}

// MenuItem is a generated struct representing the https://schema.org/MenuItem class.
type MenuItem struct {
	MenuAddOn       *MenuSection          `json:"menuaddon"`
	Nutrition       *NutritionInformation `json:"nutrition"`
	SuitableForDiet *RestrictedDiet       `json:"suitablefordiet"`
	Offers          *Offer                `json:"offers"`
}

// TherapeuticProcedure is a generated struct representing the https://schema.org/TherapeuticProcedure class.
type TherapeuticProcedure struct {
	AdverseOutcome *MedicalEntity `json:"adverseoutcome"`
	DoseSchedule   *DoseSchedule  `json:"doseschedule"`
	Drug           *Drug          `json:"drug"`
}

// Aquarium is a generated struct representing the https://schema.org/Aquarium class.
type Aquarium struct {
}

// MeetingRoom is a generated struct representing the https://schema.org/MeetingRoom class.
type MeetingRoom struct {
}

// PaymentChargeSpecification is a generated struct representing the https://schema.org/PaymentChargeSpecification class.
type PaymentChargeSpecification struct {
	AppliesToDeliveryMethod *DeliveryMethod `json:"appliestodeliverymethod"`
	AppliesToPaymentMethod  *PaymentMethod  `json:"appliestopaymentmethod"`
}

// DonateAction is a generated struct representing the https://schema.org/DonateAction class.
type DonateAction struct {
	Recipient *Person `json:"recipient"`
}

// OwnershipInfo is a generated struct representing the https://schema.org/OwnershipInfo class.
type OwnershipInfo struct {
	AcquiredFrom *Person   `json:"acquiredfrom"`
	OwnedFrom    *DateTime `json:"ownedfrom"`
	OwnedThrough *DateTime `json:"ownedthrough"`
	TypeOfGood   *Service  `json:"typeofgood"`
}

// Quiz is a generated struct representing the https://schema.org/Quiz class.
type Quiz struct {
}

// Series is a generated struct representing the https://schema.org/Series class.
type Series struct {
}

// WebPageElement is a generated struct representing the https://schema.org/WebPageElement class.
type WebPageElement struct {
	CssSelector *CssSelectorType `json:"cssselector"`
	Xpath       *XPathType       `json:"xpath"`
}

// WholesaleStore is a generated struct representing the https://schema.org/WholesaleStore class.
type WholesaleStore struct {
}

// BefriendAction is a generated struct representing the https://schema.org/BefriendAction class.
type BefriendAction struct {
}

// Casino is a generated struct representing the https://schema.org/Casino class.
type Casino struct {
}

// CoverArt is a generated struct representing the https://schema.org/CoverArt class.
type CoverArt struct {
}

// DrugLegalStatus is a generated struct representing the https://schema.org/DrugLegalStatus class.
type DrugLegalStatus struct {
	ApplicableLocation *AdministrativeArea `json:"applicablelocation"`
}

// ItemPage is a generated struct representing the https://schema.org/ItemPage class.
type ItemPage struct {
}

// LoanOrCredit is a generated struct representing the https://schema.org/LoanOrCredit class.
type LoanOrCredit struct {
	Amount             *Number                 `json:"amount"`
	Currency           *Text                   `json:"currency"`
	GracePeriod        *Duration               `json:"graceperiod"`
	LoanRepaymentForm  *RepaymentSpecification `json:"loanrepaymentform"`
	LoanTerm           *QuantitativeValue      `json:"loanterm"`
	LoanType           *URL                    `json:"loantype"`
	RecourseLoan       *Boolean                `json:"recourseloan"`
	RenegotiableLoan   *Boolean                `json:"renegotiableloan"`
	RequiredCollateral *Thing                  `json:"requiredcollateral"`
}

// ServiceChannel is a generated struct representing the https://schema.org/ServiceChannel class.
type ServiceChannel struct {
	AvailableLanguage    *Text          `json:"availablelanguage"`
	ProcessingTime       *Duration      `json:"processingtime"`
	ProvidesService      *Service       `json:"providesservice"`
	ServiceLocation      *Place         `json:"servicelocation"`
	ServicePhone         *ContactPoint  `json:"servicephone"`
	ServicePostalAddress *PostalAddress `json:"servicepostaladdress"`
	ServiceSmsNumber     *ContactPoint  `json:"servicesmsnumber"`
	ServiceUrl           *URL           `json:"serviceurl"`
}

// WearableSizeSystemEnumeration is a generated struct representing the https://schema.org/WearableSizeSystemEnumeration class.
type WearableSizeSystemEnumeration struct {
}

// AboutPage is a generated struct representing the https://schema.org/AboutPage class.
type AboutPage struct {
}

// GiveAction is a generated struct representing the https://schema.org/GiveAction class.
type GiveAction struct {
	Recipient *Person `json:"recipient"`
}

// MedicalRiskScore is a generated struct representing the https://schema.org/MedicalRiskScore class.
type MedicalRiskScore struct {
	Algorithm *Text `json:"algorithm"`
}

// Duration is a generated struct representing the https://schema.org/Duration class.
type Duration struct {
}

// MoneyTransfer is a generated struct representing the https://schema.org/MoneyTransfer class.
type MoneyTransfer struct {
	Amount          *Number `json:"amount"`
	BeneficiaryBank *Text   `json:"beneficiarybank"`
}

// HealthPlanCostSharingSpecification is a generated struct representing the https://schema.org/HealthPlanCostSharingSpecification class.
type HealthPlanCostSharingSpecification struct {
	HealthPlanCoinsuranceOption *Text               `json:"healthplancoinsuranceoption"`
	HealthPlanCoinsuranceRate   *Number             `json:"healthplancoinsurancerate"`
	HealthPlanCopay             *PriceSpecification `json:"healthplancopay"`
	HealthPlanCopayOption       *Text               `json:"healthplancopayoption"`
	HealthPlanPharmacyCategory  *Text               `json:"healthplanpharmacycategory"`
}

// Product is a generated struct representing the https://schema.org/Product class.
type Product struct {
	AdditionalProperty          *PropertyValue            `json:"additionalproperty"`
	AggregateRating             *AggregateRating          `json:"aggregaterating"`
	Asin                        *URL                      `json:"asin"`
	Awards                      *Text                     `json:"awards"`
	Brand                       *Organization             `json:"brand"`
	Color                       *Text                     `json:"color"`
	CountryOfAssembly           *Text                     `json:"countryofassembly"`
	CountryOfLastProcessing     *Text                     `json:"countryoflastprocessing"`
	CountryOfOrigin             *Country                  `json:"countryoforigin"`
	Depth                       *QuantitativeValue        `json:"depth"`
	Gtin12                      *Text                     `json:"gtin12"`
	Gtin13                      *Text                     `json:"gtin13"`
	Gtin14                      *Text                     `json:"gtin14"`
	Gtin8                       *Text                     `json:"gtin8"`
	HasAdultConsideration       *AdultOrientedEnumeration `json:"hasadultconsideration"`
	HasEnergyConsumptionDetails *EnergyConsumptionDetails `json:"hasenergyconsumptiondetails"`
	HasMeasurement              *QuantitativeValue        `json:"hasmeasurement"`
	HasMerchantReturnPolicy     *MerchantReturnPolicy     `json:"hasmerchantreturnpolicy"`
	Height                      *QuantitativeValue        `json:"height"`
	InProductGroupWithID        *Text                     `json:"inproductgroupwithid"`
	IsAccessoryOrSparePartFor   *Product                  `json:"isaccessoryorsparepartfor"`
	IsConsumableFor             *Product                  `json:"isconsumablefor"`
	IsFamilyFriendly            *Boolean                  `json:"isfamilyfriendly"`
	IsRelatedTo                 *Service                  `json:"isrelatedto"`
	IsSimilarTo                 *Service                  `json:"issimilarto"`
	ItemCondition               *OfferItemCondition       `json:"itemcondition"`
	Keywords                    *URL                      `json:"keywords"`
	Logo                        *URL                      `json:"logo"`
	Manufacturer                *Organization             `json:"manufacturer"`
	MobileUrl                   *Text                     `json:"mobileurl"`
	Model                       *Text                     `json:"model"`
	Mpn                         *Text                     `json:"mpn"`
	NegativeNotes               *WebContent               `json:"negativenotes"`
	Nsn                         *Text                     `json:"nsn"`
	Pattern                     *Text                     `json:"pattern"`
	PositiveNotes               *WebContent               `json:"positivenotes"`
	ProductID                   *Text                     `json:"productid"`
	ProductionDate              *Date                     `json:"productiondate"`
	PurchaseDate                *Date                     `json:"purchasedate"`
	ReleaseDate                 *Date                     `json:"releasedate"`
	Reviews                     *Review                   `json:"reviews"`
	Size                        *Text                     `json:"size"`
	Sku                         *Text                     `json:"sku"`
	Slogan                      *Text                     `json:"slogan"`
	Weight                      *QuantitativeValue        `json:"weight"`
	Width                       *QuantitativeValue        `json:"width"`
	Audience                    *Audience                 `json:"audience"`
	Award                       *Text                     `json:"award"`
	Category                    *URL                      `json:"category"`
	Funding                     *Grant                    `json:"funding"`
	IsVariantOf                 *ProductModel             `json:"isvariantof"`
	Offers                      *Offer                    `json:"offers"`
	Review                      *Review                   `json:"review"`
	Material                    *URL                      `json:"material"`
	Gtin                        *URL                      `json:"gtin"`
}

// Protein is a generated struct representing the https://schema.org/Protein class.
type Protein struct {
	HasBioPolymerSequence *Text `json:"hasbiopolymersequence"`
}

// MeasurementTypeEnumeration is a generated struct representing the https://schema.org/MeasurementTypeEnumeration class.
type MeasurementTypeEnumeration struct {
}

// StatisticalVariable is a generated struct representing the https://schema.org/StatisticalVariable class.
type StatisticalVariable struct {
	MeasuredProperty       *Property            `json:"measuredproperty"`
	MeasurementDenominator *StatisticalVariable `json:"measurementdenominator"`
	MeasurementMethod      *URL                 `json:"measurementmethod"`
	MeasurementQualifier   *Enumeration         `json:"measurementqualifier"`
	PopulationType         *Class               `json:"populationtype"`
	StatType               *URL                 `json:"stattype"`
	MeasurementTechnique   *URL                 `json:"measurementtechnique"`
}

// GovernmentOffice is a generated struct representing the https://schema.org/GovernmentOffice class.
type GovernmentOffice struct {
}

// PublicationEvent is a generated struct representing the https://schema.org/PublicationEvent class.
type PublicationEvent struct {
	Free        *Boolean          `json:"free"`
	PublishedBy *Person           `json:"publishedby"`
	PublishedOn *BroadcastService `json:"publishedon"`
}

// RecommendedDoseSchedule is a generated struct representing the https://schema.org/RecommendedDoseSchedule class.
type RecommendedDoseSchedule struct {
}

// WatchAction is a generated struct representing the https://schema.org/WatchAction class.
type WatchAction struct {
}

// SearchResultsPage is a generated struct representing the https://schema.org/SearchResultsPage class.
type SearchResultsPage struct {
}

// ComputerStore is a generated struct representing the https://schema.org/ComputerStore class.
type ComputerStore struct {
}

// Conversation is a generated struct representing the https://schema.org/Conversation class.
type Conversation struct {
}

// DanceEvent is a generated struct representing the https://schema.org/DanceEvent class.
type DanceEvent struct {
}

// HVACBusiness is a generated struct representing the https://schema.org/HVACBusiness class.
type HVACBusiness struct {
}

// MedicalTest is a generated struct representing the https://schema.org/MedicalTest class.
type MedicalTest struct {
	AffectedBy     *Drug             `json:"affectedby"`
	NormalRange    *Text             `json:"normalrange"`
	SignDetected   *MedicalSign      `json:"signdetected"`
	UsedToDiagnose *MedicalCondition `json:"usedtodiagnose"`
	UsesDevice     *MedicalDevice    `json:"usesdevice"`
}

// OnDemandEvent is a generated struct representing the https://schema.org/OnDemandEvent class.
type OnDemandEvent struct {
}

// Quotation is a generated struct representing the https://schema.org/Quotation class.
type Quotation struct {
	SpokenByCharacter *Person `json:"spokenbycharacter"`
}

// UpdateAction is a generated struct representing the https://schema.org/UpdateAction class.
type UpdateAction struct {
	Collection       *Thing `json:"collection"`
	TargetCollection *Thing `json:"targetcollection"`
}

// HyperToc is a generated struct representing the https://schema.org/HyperToc class.
type HyperToc struct {
	AssociatedMedia *MediaObject   `json:"associatedmedia"`
	TocEntry        *HyperTocEntry `json:"tocentry"`
}

// CheckAction is a generated struct representing the https://schema.org/CheckAction class.
type CheckAction struct {
}

// MovieTheater is a generated struct representing the https://schema.org/MovieTheater class.
type MovieTheater struct {
	ScreenCount *Number `json:"screencount"`
}

// MusicRecording is a generated struct representing the https://schema.org/MusicRecording class.
type MusicRecording struct {
	ByArtist    *Person           `json:"byartist"`
	InAlbum     *MusicAlbum       `json:"inalbum"`
	InPlaylist  *MusicPlaylist    `json:"inplaylist"`
	IsrcCode    *Text             `json:"isrccode"`
	Duration    *Duration         `json:"duration"`
	RecordingOf *MusicComposition `json:"recordingof"`
}

// PlayAction is a generated struct representing the https://schema.org/PlayAction class.
type PlayAction struct {
	Audience *Audience `json:"audience"`
	Event    *Event    `json:"event"`
}

// AllocateAction is a generated struct representing the https://schema.org/AllocateAction class.
type AllocateAction struct {
}

// BankAccount is a generated struct representing the https://schema.org/BankAccount class.
type BankAccount struct {
	AccountMinimumInflow  *MonetaryAmount `json:"accountminimuminflow"`
	AccountOverdraftLimit *MonetaryAmount `json:"accountoverdraftlimit"`
	BankAccountType       *URL            `json:"bankaccounttype"`
}

// FireStation is a generated struct representing the https://schema.org/FireStation class.
type FireStation struct {
}

// Hackathon is a generated struct representing the https://schema.org/Hackathon class.
type Hackathon struct {
}

// HinduTemple is a generated struct representing the https://schema.org/HinduTemple class.
type HinduTemple struct {
}

// MedicalIndication is a generated struct representing the https://schema.org/MedicalIndication class.
type MedicalIndication struct {
}

// ParkingFacility is a generated struct representing the https://schema.org/ParkingFacility class.
type ParkingFacility struct {
}

// PodcastSeason is a generated struct representing the https://schema.org/PodcastSeason class.
type PodcastSeason struct {
}

// SizeSpecification is a generated struct representing the https://schema.org/SizeSpecification class.
type SizeSpecification struct {
	HasMeasurement       *QuantitativeValue `json:"hasmeasurement"`
	SizeGroup            *Text              `json:"sizegroup"`
	SizeSystem           *Text              `json:"sizesystem"`
	SuggestedAge         *QuantitativeValue `json:"suggestedage"`
	SuggestedGender      *Text              `json:"suggestedgender"`
	SuggestedMeasurement *QuantitativeValue `json:"suggestedmeasurement"`
}

// Locksmith is a generated struct representing the https://schema.org/Locksmith class.
type Locksmith struct {
}

// PostalCodeRangeSpecification is a generated struct representing the https://schema.org/PostalCodeRangeSpecification class.
type PostalCodeRangeSpecification struct {
	PostalCodeBegin *Text `json:"postalcodebegin"`
	PostalCodeEnd   *Text `json:"postalcodeend"`
}

// School is a generated struct representing the https://schema.org/School class.
type School struct {
}

// Waterfall is a generated struct representing the https://schema.org/Waterfall class.
type Waterfall struct {
}

// EmploymentAgency is a generated struct representing the https://schema.org/EmploymentAgency class.
type EmploymentAgency struct {
}

// MedicalSpecialty is a generated struct representing the https://schema.org/MedicalSpecialty class.
type MedicalSpecialty struct {
}

// MedicalTrial is a generated struct representing the https://schema.org/MedicalTrial class.
type MedicalTrial struct {
	TrialDesign *MedicalTrialDesign `json:"trialdesign"`
}

// ImageObject is a generated struct representing the https://schema.org/ImageObject class.
type ImageObject struct {
	EmbeddedTextCaption  *Text    `json:"embeddedtextcaption"`
	ExifData             *Text    `json:"exifdata"`
	RepresentativeOfPage *Boolean `json:"representativeofpage"`
	Caption              *Text    `json:"caption"`
}

// PublicationIssue is a generated struct representing the https://schema.org/PublicationIssue class.
type PublicationIssue struct {
	IssueNumber *Text `json:"issuenumber"`
	PageEnd     *Text `json:"pageend"`
	PageStart   *Text `json:"pagestart"`
	Pagination  *Text `json:"pagination"`
}

// BodyOfWater is a generated struct representing the https://schema.org/BodyOfWater class.
type BodyOfWater struct {
}

// Poster is a generated struct representing the https://schema.org/Poster class.
type Poster struct {
}

// Reservation is a generated struct representing the https://schema.org/Reservation class.
type Reservation struct {
	BookingAgent          *Person                `json:"bookingagent"`
	BookingTime           *DateTime              `json:"bookingtime"`
	ModifiedTime          *DateTime              `json:"modifiedtime"`
	PriceCurrency         *Text                  `json:"pricecurrency"`
	ProgramMembershipUsed *ProgramMembership     `json:"programmembershipused"`
	ReservationFor        *Thing                 `json:"reservationfor"`
	ReservationId         *Text                  `json:"reservationid"`
	ReservationStatus     *ReservationStatusType `json:"reservationstatus"`
	ReservedTicket        *Ticket                `json:"reservedticket"`
	TotalPrice            *Text                  `json:"totalprice"`
	UnderName             *Person                `json:"undername"`
	Broker                *Person                `json:"broker"`
	Provider              *Person                `json:"provider"`
}

// ScheduleAction is a generated struct representing the https://schema.org/ScheduleAction class.
type ScheduleAction struct {
}

// Synagogue is a generated struct representing the https://schema.org/Synagogue class.
type Synagogue struct {
}

// AlignmentObject is a generated struct representing the https://schema.org/AlignmentObject class.
type AlignmentObject struct {
	AlignmentType        *Text `json:"alignmenttype"`
	EducationalFramework *Text `json:"educationalframework"`
	TargetDescription    *Text `json:"targetdescription"`
	TargetName           *Text `json:"targetname"`
	TargetUrl            *URL  `json:"targeturl"`
}

// Consortium is a generated struct representing the https://schema.org/Consortium class.
type Consortium struct {
}

// InvestmentOrDeposit is a generated struct representing the https://schema.org/InvestmentOrDeposit class.
type InvestmentOrDeposit struct {
	Amount *Number `json:"amount"`
}

// PublicToilet is a generated struct representing the https://schema.org/PublicToilet class.
type PublicToilet struct {
}

// TaxiService is a generated struct representing the https://schema.org/TaxiService class.
type TaxiService struct {
}

// UserReview is a generated struct representing the https://schema.org/UserReview class.
type UserReview struct {
}

// BusStation is a generated struct representing the https://schema.org/BusStation class.
type BusStation struct {
}

// CollegeOrUniversity is a generated struct representing the https://schema.org/CollegeOrUniversity class.
type CollegeOrUniversity struct {
}

// ConvenienceStore is a generated struct representing the https://schema.org/ConvenienceStore class.
type ConvenienceStore struct {
}

// MedicalCondition is a generated struct representing the https://schema.org/MedicalCondition class.
type MedicalCondition struct {
	AssociatedAnatomy     *SuperficialAnatomy    `json:"associatedanatomy"`
	DifferentialDiagnosis *DDxElement            `json:"differentialdiagnosis"`
	Drug                  *Drug                  `json:"drug"`
	Epidemiology          *Text                  `json:"epidemiology"`
	ExpectedPrognosis     *Text                  `json:"expectedprognosis"`
	NaturalProgression    *Text                  `json:"naturalprogression"`
	Pathophysiology       *Text                  `json:"pathophysiology"`
	PossibleComplication  *Text                  `json:"possiblecomplication"`
	PossibleTreatment     *MedicalTherapy        `json:"possibletreatment"`
	PrimaryPrevention     *MedicalTherapy        `json:"primaryprevention"`
	RiskFactor            *MedicalRiskFactor     `json:"riskfactor"`
	SecondaryPrevention   *MedicalTherapy        `json:"secondaryprevention"`
	SignOrSymptom         *MedicalSignOrSymptom  `json:"signorsymptom"`
	Stage                 *MedicalConditionStage `json:"stage"`
	Status                *Text                  `json:"status"`
	TypicalTest           *MedicalTest           `json:"typicaltest"`
}

// Comment is a generated struct representing the https://schema.org/Comment class.
type Comment struct {
	DownvoteCount *Integer      `json:"downvotecount"`
	ParentItem    *CreativeWork `json:"parentitem"`
	SharedContent *CreativeWork `json:"sharedcontent"`
	UpvoteCount   *Integer      `json:"upvotecount"`
}

// SpeakableSpecification is a generated struct representing the https://schema.org/SpeakableSpecification class.
type SpeakableSpecification struct {
	CssSelector *CssSelectorType `json:"cssselector"`
	Xpath       *XPathType       `json:"xpath"`
}

// TireShop is a generated struct representing the https://schema.org/TireShop class.
type TireShop struct {
}

// ReportedDoseSchedule is a generated struct representing the https://schema.org/ReportedDoseSchedule class.
type ReportedDoseSchedule struct {
}

// EventSeries is a generated struct representing the https://schema.org/EventSeries class.
type EventSeries struct {
}

// HowToTip is a generated struct representing the https://schema.org/HowToTip class.
type HowToTip struct {
}

// MedicalProcedure is a generated struct representing the https://schema.org/MedicalProcedure class.
type MedicalProcedure struct {
	BodyLocation  *Text                 `json:"bodylocation"`
	Followup      *Text                 `json:"followup"`
	HowPerformed  *Text                 `json:"howperformed"`
	Preparation   *Text                 `json:"preparation"`
	ProcedureType *MedicalProcedureType `json:"proceduretype"`
	Status        *Text                 `json:"status"`
}

// MedicineSystem is a generated struct representing the https://schema.org/MedicineSystem class.
type MedicineSystem struct {
}

// PalliativeProcedure is a generated struct representing the https://schema.org/PalliativeProcedure class.
type PalliativeProcedure struct {
}

// RealEstateAgent is a generated struct representing the https://schema.org/RealEstateAgent class.
type RealEstateAgent struct {
}

// RefundTypeEnumeration is a generated struct representing the https://schema.org/RefundTypeEnumeration class.
type RefundTypeEnumeration struct {
}

// VirtualLocation is a generated struct representing the https://schema.org/VirtualLocation class.
type VirtualLocation struct {
}

// Library is a generated struct representing the https://schema.org/Library class.
type Library struct {
}

// PresentationDigitalDocument is a generated struct representing the https://schema.org/PresentationDigitalDocument class.
type PresentationDigitalDocument struct {
}

// SchoolDistrict is a generated struct representing the https://schema.org/SchoolDistrict class.
type SchoolDistrict struct {
}

// CityHall is a generated struct representing the https://schema.org/CityHall class.
type CityHall struct {
}

// Continent is a generated struct representing the https://schema.org/Continent class.
type Continent struct {
}

// NoteDigitalDocument is a generated struct representing the https://schema.org/NoteDigitalDocument class.
type NoteDigitalDocument struct {
}

// Rating is a generated struct representing the https://schema.org/Rating class.
type Rating struct {
	Author            *Person `json:"author"`
	BestRating        *Text   `json:"bestrating"`
	RatingExplanation *Text   `json:"ratingexplanation"`
	RatingValue       *Text   `json:"ratingvalue"`
	ReviewAspect      *Text   `json:"reviewaspect"`
	WorstRating       *Text   `json:"worstrating"`
}

// TravelAgency is a generated struct representing the https://schema.org/TravelAgency class.
type TravelAgency struct {
}

// VitalSign is a generated struct representing the https://schema.org/VitalSign class.
type VitalSign struct {
}

// Accommodation is a generated struct representing the https://schema.org/Accommodation class.
type Accommodation struct {
	AccommodationCategory    *Text                         `json:"accommodationcategory"`
	AccommodationFloorPlan   *FloorPlan                    `json:"accommodationfloorplan"`
	AmenityFeature           *LocationFeatureSpecification `json:"amenityfeature"`
	Bed                      *Text                         `json:"bed"`
	FloorLevel               *Text                         `json:"floorlevel"`
	FloorSize                *QuantitativeValue            `json:"floorsize"`
	LeaseLength              *QuantitativeValue            `json:"leaselength"`
	NumberOfBathroomsTotal   *Integer                      `json:"numberofbathroomstotal"`
	NumberOfBedrooms         *QuantitativeValue            `json:"numberofbedrooms"`
	NumberOfFullBathrooms    *Number                       `json:"numberoffullbathrooms"`
	NumberOfPartialBathrooms *Number                       `json:"numberofpartialbathrooms"`
	NumberOfRooms            *QuantitativeValue            `json:"numberofrooms"`
	Occupancy                *QuantitativeValue            `json:"occupancy"`
	PermittedUsage           *Text                         `json:"permittedusage"`
	PetsAllowed              *Text                         `json:"petsallowed"`
	TourBookingPage          *URL                          `json:"tourbookingpage"`
	YearBuilt                *Number                       `json:"yearbuilt"`
}

// AssignAction is a generated struct representing the https://schema.org/AssignAction class.
type AssignAction struct {
}

// MedicalSymptom is a generated struct representing the https://schema.org/MedicalSymptom class.
type MedicalSymptom struct {
}

// RecyclingCenter is a generated struct representing the https://schema.org/RecyclingCenter class.
type RecyclingCenter struct {
}

// WearableSizeGroupEnumeration is a generated struct representing the https://schema.org/WearableSizeGroupEnumeration class.
type WearableSizeGroupEnumeration struct {
}

// AccountingService is a generated struct representing the https://schema.org/AccountingService class.
type AccountingService struct {
}

// AdministrativeArea is a generated struct representing the https://schema.org/AdministrativeArea class.
type AdministrativeArea struct {
}

// Map is a generated struct representing the https://schema.org/Map class.
type Map struct {
	MapType *MapCategoryType `json:"maptype"`
}

// Number is a generated struct representing the https://schema.org/Number class.
type Number struct {
}

// RadioStation is a generated struct representing the https://schema.org/RadioStation class.
type RadioStation struct {
}

// SheetMusic is a generated struct representing the https://schema.org/SheetMusic class.
type SheetMusic struct {
}

// CampingPitch is a generated struct representing the https://schema.org/CampingPitch class.
type CampingPitch struct {
}

// GeoShape is a generated struct representing the https://schema.org/GeoShape class.
type GeoShape struct {
	Address        *Text `json:"address"`
	AddressCountry *Text `json:"addresscountry"`
	Box            *Text `json:"box"`
	Circle         *Text `json:"circle"`
	Elevation      *Text `json:"elevation"`
	Line           *Text `json:"line"`
	Polygon        *Text `json:"polygon"`
	PostalCode     *Text `json:"postalcode"`
}

// Blog is a generated struct representing the https://schema.org/Blog class.
type Blog struct {
	BlogPosts *BlogPosting `json:"blogposts"`
	Issn      *Text        `json:"issn"`
	BlogPost  *BlogPosting `json:"blogpost"`
}

// BroadcastEvent is a generated struct representing the https://schema.org/BroadcastEvent class.
type BroadcastEvent struct {
	BroadcastOfEvent *Event   `json:"broadcastofevent"`
	IsLiveBroadcast  *Boolean `json:"islivebroadcast"`
	SubtitleLanguage *Text    `json:"subtitlelanguage"`
	VideoFormat      *Text    `json:"videoformat"`
}

// Energy is a generated struct representing the https://schema.org/Energy class.
type Energy struct {
}

// LifestyleModification is a generated struct representing the https://schema.org/LifestyleModification class.
type LifestyleModification struct {
}

// Movie is a generated struct representing the https://schema.org/Movie class.
type Movie struct {
	Actors            *Person       `json:"actors"`
	CountryOfOrigin   *Country      `json:"countryoforigin"`
	Directors         *Person       `json:"directors"`
	MusicBy           *Person       `json:"musicby"`
	ProductionCompany *Organization `json:"productioncompany"`
	SubtitleLanguage  *Text         `json:"subtitlelanguage"`
	TitleEIDR         *URL          `json:"titleeidr"`
	Trailer           *VideoObject  `json:"trailer"`
	Director          *Person       `json:"director"`
	Duration          *Duration     `json:"duration"`
	Actor             *Person       `json:"actor"`
}

// RadioEpisode is a generated struct representing the https://schema.org/RadioEpisode class.
type RadioEpisode struct {
}

// SeekToAction is a generated struct representing the https://schema.org/SeekToAction class.
type SeekToAction struct {
	StartOffset *Number `json:"startoffset"`
}

// ConsumeAction is a generated struct representing the https://schema.org/ConsumeAction class.
type ConsumeAction struct {
	ActionAccessibilityRequirement *ActionAccessSpecification `json:"actionaccessibilityrequirement"`
	ExpectsAcceptanceOf            *Offer                     `json:"expectsacceptanceof"`
}

// DanceGroup is a generated struct representing the https://schema.org/DanceGroup class.
type DanceGroup struct {
}

// EventReservation is a generated struct representing the https://schema.org/EventReservation class.
type EventReservation struct {
}

// PublicationVolume is a generated struct representing the https://schema.org/PublicationVolume class.
type PublicationVolume struct {
	PageEnd      *Text `json:"pageend"`
	PageStart    *Text `json:"pagestart"`
	Pagination   *Text `json:"pagination"`
	VolumeNumber *Text `json:"volumenumber"`
}

// ResearchOrganization is a generated struct representing the https://schema.org/ResearchOrganization class.
type ResearchOrganization struct {
}

// ArchiveOrganization is a generated struct representing the https://schema.org/ArchiveOrganization class.
type ArchiveOrganization struct {
	ArchiveHeld *ArchiveComponent `json:"archiveheld"`
}

// Hotel is a generated struct representing the https://schema.org/Hotel class.
type Hotel struct {
}

// NutritionInformation is a generated struct representing the https://schema.org/NutritionInformation class.
type NutritionInformation struct {
	Calories              *Energy `json:"calories"`
	CarbohydrateContent   *Mass   `json:"carbohydratecontent"`
	CholesterolContent    *Mass   `json:"cholesterolcontent"`
	FatContent            *Mass   `json:"fatcontent"`
	FiberContent          *Mass   `json:"fibercontent"`
	ProteinContent        *Mass   `json:"proteincontent"`
	SaturatedFatContent   *Mass   `json:"saturatedfatcontent"`
	ServingSize           *Text   `json:"servingsize"`
	SodiumContent         *Mass   `json:"sodiumcontent"`
	SugarContent          *Mass   `json:"sugarcontent"`
	TransFatContent       *Mass   `json:"transfatcontent"`
	UnsaturatedFatContent *Mass   `json:"unsaturatedfatcontent"`
}

// RejectAction is a generated struct representing the https://schema.org/RejectAction class.
type RejectAction struct {
}

// EmployerAggregateRating is a generated struct representing the https://schema.org/EmployerAggregateRating class.
type EmployerAggregateRating struct {
}

// PlayGameAction is a generated struct representing the https://schema.org/PlayGameAction class.
type PlayGameAction struct {
	GameAvailabilityType *Text `json:"gameavailabilitytype"`
}

// BookStore is a generated struct representing the https://schema.org/BookStore class.
type BookStore struct {
}

// BookmarkAction is a generated struct representing the https://schema.org/BookmarkAction class.
type BookmarkAction struct {
}

// LiquorStore is a generated struct representing the https://schema.org/LiquorStore class.
type LiquorStore struct {
}

// LoseAction is a generated struct representing the https://schema.org/LoseAction class.
type LoseAction struct {
	Winner *Person `json:"winner"`
}

// MovieClip is a generated struct representing the https://schema.org/MovieClip class.
type MovieClip struct {
}

// ChildCare is a generated struct representing the https://schema.org/ChildCare class.
type ChildCare struct {
}

// Drug is a generated struct representing the https://schema.org/Drug class.
type Drug struct {
	ActiveIngredient              *Text                  `json:"activeingredient"`
	AdministrationRoute           *Text                  `json:"administrationroute"`
	AlcoholWarning                *Text                  `json:"alcoholwarning"`
	AvailableStrength             *DrugStrength          `json:"availablestrength"`
	BreastfeedingWarning          *Text                  `json:"breastfeedingwarning"`
	ClincalPharmacology           *Text                  `json:"clincalpharmacology"`
	DosageForm                    *Text                  `json:"dosageform"`
	DoseSchedule                  *DoseSchedule          `json:"doseschedule"`
	DrugClass                     *DrugClass             `json:"drugclass"`
	DrugUnit                      *Text                  `json:"drugunit"`
	FoodWarning                   *Text                  `json:"foodwarning"`
	IncludedInHealthInsurancePlan *HealthInsurancePlan   `json:"includedinhealthinsuranceplan"`
	InteractingDrug               *Drug                  `json:"interactingdrug"`
	IsAvailableGenerically        *Boolean               `json:"isavailablegenerically"`
	IsProprietary                 *Boolean               `json:"isproprietary"`
	LabelDetails                  *URL                   `json:"labeldetails"`
	LegalStatus                   *Text                  `json:"legalstatus"`
	MaximumIntake                 *MaximumDoseSchedule   `json:"maximumintake"`
	MechanismOfAction             *Text                  `json:"mechanismofaction"`
	NonProprietaryName            *Text                  `json:"nonproprietaryname"`
	Overdosage                    *Text                  `json:"overdosage"`
	PregnancyCategory             *DrugPregnancyCategory `json:"pregnancycategory"`
	PregnancyWarning              *Text                  `json:"pregnancywarning"`
	PrescribingInfo               *URL                   `json:"prescribinginfo"`
	PrescriptionStatus            *Text                  `json:"prescriptionstatus"`
	ProprietaryName               *Text                  `json:"proprietaryname"`
	RelatedDrug                   *Drug                  `json:"relateddrug"`
	Rxcui                         *Text                  `json:"rxcui"`
	Warning                       *URL                   `json:"warning"`
	ClinicalPharmacology          *Text                  `json:"clinicalpharmacology"`
}

// PoliceStation is a generated struct representing the https://schema.org/PoliceStation class.
type PoliceStation struct {
}

// SiteNavigationElement is a generated struct representing the https://schema.org/SiteNavigationElement class.
type SiteNavigationElement struct {
}

// SubscribeAction is a generated struct representing the https://schema.org/SubscribeAction class.
type SubscribeAction struct {
}

// GeoCircle is a generated struct representing the https://schema.org/GeoCircle class.
type GeoCircle struct {
	GeoMidpoint *GeoCoordinates `json:"geomidpoint"`
	GeoRadius   *Text           `json:"georadius"`
}

// Guide is a generated struct representing the https://schema.org/Guide class.
type Guide struct {
	ReviewAspect *Text `json:"reviewaspect"`
}

// InteractionCounter is a generated struct representing the https://schema.org/InteractionCounter class.
type InteractionCounter struct {
	EndTime              *Time            `json:"endtime"`
	InteractionService   *WebSite         `json:"interactionservice"`
	InteractionType      *Action          `json:"interactiontype"`
	StartTime            *Time            `json:"starttime"`
	UserInteractionCount *Integer         `json:"userinteractioncount"`
	Location             *VirtualLocation `json:"location"`
}

// ShoeStore is a generated struct representing the https://schema.org/ShoeStore class.
type ShoeStore struct {
}

// TVSeries is a generated struct representing the https://schema.org/TVSeries class.
type TVSeries struct {
	Actors            *Person             `json:"actors"`
	CountryOfOrigin   *Country            `json:"countryoforigin"`
	Directors         *Person             `json:"directors"`
	Episodes          *Episode            `json:"episodes"`
	MusicBy           *Person             `json:"musicby"`
	NumberOfEpisodes  *Integer            `json:"numberofepisodes"`
	NumberOfSeasons   *Integer            `json:"numberofseasons"`
	ProductionCompany *Organization       `json:"productioncompany"`
	Seasons           *CreativeWorkSeason `json:"seasons"`
	TitleEIDR         *URL                `json:"titleeidr"`
	Trailer           *VideoObject        `json:"trailer"`
	ContainsSeason    *CreativeWorkSeason `json:"containsseason"`
	Director          *Person             `json:"director"`
	Episode           *Episode            `json:"episode"`
	Season            *URL                `json:"season"`
	Actor             *Person             `json:"actor"`
}

// ToyStore is a generated struct representing the https://schema.org/ToyStore class.
type ToyStore struct {
}

// CategoryCode is a generated struct representing the https://schema.org/CategoryCode class.
type CategoryCode struct {
	CodeValue *Text `json:"codevalue"`
	InCodeSet *URL  `json:"incodeset"`
}

// ComedyClub is a generated struct representing the https://schema.org/ComedyClub class.
type ComedyClub struct {
}

// TelevisionChannel is a generated struct representing the https://schema.org/TelevisionChannel class.
type TelevisionChannel struct {
}

// CriticReview is a generated struct representing the https://schema.org/CriticReview class.
type CriticReview struct {
}

// ItemAvailability is a generated struct representing the https://schema.org/ItemAvailability class.
type ItemAvailability struct {
}

// MedicalProcedureType is a generated struct representing the https://schema.org/MedicalProcedureType class.
type MedicalProcedureType struct {
}

// OpeningHoursSpecification is a generated struct representing the https://schema.org/OpeningHoursSpecification class.
type OpeningHoursSpecification struct {
	Closes       *Time      `json:"closes"`
	DayOfWeek    *DayOfWeek `json:"dayofweek"`
	Opens        *Time      `json:"opens"`
	ValidFrom    *DateTime  `json:"validfrom"`
	ValidThrough *DateTime  `json:"validthrough"`
}

// AdultOrientedEnumeration is a generated struct representing the https://schema.org/AdultOrientedEnumeration class.
type AdultOrientedEnumeration struct {
}

// AnalysisNewsArticle is a generated struct representing the https://schema.org/AnalysisNewsArticle class.
type AnalysisNewsArticle struct {
}

// EducationalOccupationalCredential is a generated struct representing the https://schema.org/EducationalOccupationalCredential class.
type EducationalOccupationalCredential struct {
	CompetencyRequired *URL                `json:"competencyrequired"`
	CredentialCategory *URL                `json:"credentialcategory"`
	EducationalLevel   *URL                `json:"educationallevel"`
	RecognizedBy       *Organization       `json:"recognizedby"`
	ValidFor           *Duration           `json:"validfor"`
	ValidIn            *AdministrativeArea `json:"validin"`
}

// MedicalAudience is a generated struct representing the https://schema.org/MedicalAudience class.
type MedicalAudience struct {
}

// TelevisionStation is a generated struct representing the https://schema.org/TelevisionStation class.
type TelevisionStation struct {
}

// Vein is a generated struct representing the https://schema.org/Vein class.
type Vein struct {
	DrainsTo      *Vessel              `json:"drainsto"`
	RegionDrained *AnatomicalSystem    `json:"regiondrained"`
	Tributary     *AnatomicalStructure `json:"tributary"`
}

// WebApplication is a generated struct representing the https://schema.org/WebApplication class.
type WebApplication struct {
	BrowserRequirements *Text `json:"browserrequirements"`
}

// DataDownload is a generated struct representing the https://schema.org/DataDownload class.
type DataDownload struct {
	MeasurementMethod    *URL `json:"measurementmethod"`
	MeasurementTechnique *URL `json:"measurementtechnique"`
}

// LibrarySystem is a generated struct representing the https://schema.org/LibrarySystem class.
type LibrarySystem struct {
}

// Pond is a generated struct representing the https://schema.org/Pond class.
type Pond struct {
}

// RadioClip is a generated struct representing the https://schema.org/RadioClip class.
type RadioClip struct {
}

// ReadAction is a generated struct representing the https://schema.org/ReadAction class.
type ReadAction struct {
}

// ExerciseAction is a generated struct representing the https://schema.org/ExerciseAction class.
type ExerciseAction struct {
	Course                 *Place                  `json:"course"`
	Diet                   *Diet                   `json:"diet"`
	Distance               *Distance               `json:"distance"`
	ExercisePlan           *ExercisePlan           `json:"exerciseplan"`
	ExerciseRelatedDiet    *Diet                   `json:"exerciserelateddiet"`
	ExerciseType           *Text                   `json:"exercisetype"`
	FromLocation           *Place                  `json:"fromlocation"`
	Opponent               *Person                 `json:"opponent"`
	SportsActivityLocation *SportsActivityLocation `json:"sportsactivitylocation"`
	SportsEvent            *SportsEvent            `json:"sportsevent"`
	SportsTeam             *SportsTeam             `json:"sportsteam"`
	ToLocation             *Place                  `json:"tolocation"`
	ExerciseCourse         *Place                  `json:"exercisecourse"`
}

// PhysicalExam is a generated struct representing the https://schema.org/PhysicalExam class.
type PhysicalExam struct {
}

// EUEnergyEfficiencyEnumeration is a generated struct representing the https://schema.org/EUEnergyEfficiencyEnumeration class.
type EUEnergyEfficiencyEnumeration struct {
}

// FoodEvent is a generated struct representing the https://schema.org/FoodEvent class.
type FoodEvent struct {
}

// ImageGallery is a generated struct representing the https://schema.org/ImageGallery class.
type ImageGallery struct {
}

// MedicalDevicePurpose is a generated struct representing the https://schema.org/MedicalDevicePurpose class.
type MedicalDevicePurpose struct {
}

// PhotographAction is a generated struct representing the https://schema.org/PhotographAction class.
type PhotographAction struct {
}

// PostOffice is a generated struct representing the https://schema.org/PostOffice class.
type PostOffice struct {
}

// SportsOrganization is a generated struct representing the https://schema.org/SportsOrganization class.
type SportsOrganization struct {
	Sport *URL `json:"sport"`
}

// Bone is a generated struct representing the https://schema.org/Bone class.
type Bone struct {
}

// Courthouse is a generated struct representing the https://schema.org/Courthouse class.
type Courthouse struct {
}

// Grant is a generated struct representing the https://schema.org/Grant class.
type Grant struct {
	Funder     *Person  `json:"funder"`
	FundedItem *Product `json:"fundeditem"`
	Sponsor    *Person  `json:"sponsor"`
}

// OnlineBusiness is a generated struct representing the https://schema.org/OnlineBusiness class.
type OnlineBusiness struct {
}

// Plumber is a generated struct representing the https://schema.org/Plumber class.
type Plumber struct {
}

// DigitalDocumentPermission is a generated struct representing the https://schema.org/DigitalDocumentPermission class.
type DigitalDocumentPermission struct {
	Grantee        *Person                        `json:"grantee"`
	PermissionType *DigitalDocumentPermissionType `json:"permissiontype"`
}

// MediaSubscription is a generated struct representing the https://schema.org/MediaSubscription class.
type MediaSubscription struct {
	Authenticator       *Organization `json:"authenticator"`
	ExpectsAcceptanceOf *Offer        `json:"expectsacceptanceof"`
}

// Order is a generated struct representing the https://schema.org/Order class.
type Order struct {
	AcceptedOffer      *Offer          `json:"acceptedoffer"`
	BillingAddress     *PostalAddress  `json:"billingaddress"`
	ConfirmationNumber *Text           `json:"confirmationnumber"`
	Customer           *Person         `json:"customer"`
	Discount           *Text           `json:"discount"`
	DiscountCode       *Text           `json:"discountcode"`
	DiscountCurrency   *Text           `json:"discountcurrency"`
	IsGift             *Boolean        `json:"isgift"`
	Merchant           *Person         `json:"merchant"`
	OrderDate          *DateTime       `json:"orderdate"`
	OrderDelivery      *ParcelDelivery `json:"orderdelivery"`
	OrderNumber        *Text           `json:"ordernumber"`
	OrderStatus        *OrderStatus    `json:"orderstatus"`
	OrderedItem        *Service        `json:"ordereditem"`
	PartOfInvoice      *Invoice        `json:"partofinvoice"`
	PaymentDue         *DateTime       `json:"paymentdue"`
	PaymentMethod      *PaymentMethod  `json:"paymentmethod"`
	PaymentMethodId    *Text           `json:"paymentmethodid"`
	PaymentUrl         *URL            `json:"paymenturl"`
	Broker             *Person         `json:"broker"`
	PaymentDueDate     *DateTime       `json:"paymentduedate"`
	Seller             *Person         `json:"seller"`
}

// ReactAction is a generated struct representing the https://schema.org/ReactAction class.
type ReactAction struct {
}

// HowToSupply is a generated struct representing the https://schema.org/HowToSupply class.
type HowToSupply struct {
	EstimatedCost *Text `json:"estimatedcost"`
}

// Motorcycle is a generated struct representing the https://schema.org/Motorcycle class.
type Motorcycle struct {
}

// OfferItemCondition is a generated struct representing the https://schema.org/OfferItemCondition class.
type OfferItemCondition struct {
}

// RadioSeason is a generated struct representing the https://schema.org/RadioSeason class.
type RadioSeason struct {
}

// TaxiReservation is a generated struct representing the https://schema.org/TaxiReservation class.
type TaxiReservation struct {
	PartySize      *QuantitativeValue `json:"partysize"`
	PickupLocation *Place             `json:"pickuplocation"`
	PickupTime     *DateTime          `json:"pickuptime"`
}

// WearableMeasurementTypeEnumeration is a generated struct representing the https://schema.org/WearableMeasurementTypeEnumeration class.
type WearableMeasurementTypeEnumeration struct {
}

// MedicalGuidelineContraindication is a generated struct representing the https://schema.org/MedicalGuidelineContraindication class.
type MedicalGuidelineContraindication struct {
}

// PathologyTest is a generated struct representing the https://schema.org/PathologyTest class.
type PathologyTest struct {
	TissueSample *Text `json:"tissuesample"`
}

// SpreadsheetDigitalDocument is a generated struct representing the https://schema.org/SpreadsheetDigitalDocument class.
type SpreadsheetDigitalDocument struct {
}

// Landform is a generated struct representing the https://schema.org/Landform class.
type Landform struct {
}

// OpinionNewsArticle is a generated struct representing the https://schema.org/OpinionNewsArticle class.
type OpinionNewsArticle struct {
}

// Substance is a generated struct representing the https://schema.org/Substance class.
type Substance struct {
	ActiveIngredient *Text                `json:"activeingredient"`
	MaximumIntake    *MaximumDoseSchedule `json:"maximumintake"`
}

// WebPage is a generated struct representing the https://schema.org/WebPage class.
type WebPage struct {
	Breadcrumb         *Text           `json:"breadcrumb"`
	LastReviewed       *Date           `json:"lastreviewed"`
	PrimaryImageOfPage *ImageObject    `json:"primaryimageofpage"`
	RelatedLink        *URL            `json:"relatedlink"`
	ReviewedBy         *Person         `json:"reviewedby"`
	SignificantLinks   *URL            `json:"significantlinks"`
	Speakable          *URL            `json:"speakable"`
	Specialty          *Specialty      `json:"specialty"`
	MainContentOfPage  *WebPageElement `json:"maincontentofpage"`
	SignificantLink    *URL            `json:"significantlink"`
}

// AutoPartsStore is a generated struct representing the https://schema.org/AutoPartsStore class.
type AutoPartsStore struct {
}

// BloodTest is a generated struct representing the https://schema.org/BloodTest class.
type BloodTest struct {
}

// CookAction is a generated struct representing the https://schema.org/CookAction class.
type CookAction struct {
	FoodEstablishment *Place     `json:"foodestablishment"`
	FoodEvent         *FoodEvent `json:"foodevent"`
	Recipe            *Recipe    `json:"recipe"`
}

// DiagnosticLab is a generated struct representing the https://schema.org/DiagnosticLab class.
type DiagnosticLab struct {
	AvailableTest *MedicalTest `json:"availabletest"`
}

// InsertAction is a generated struct representing the https://schema.org/InsertAction class.
type InsertAction struct {
	ToLocation *Place `json:"tolocation"`
}

// LegislationObject is a generated struct representing the https://schema.org/LegislationObject class.
type LegislationObject struct {
	LegislationLegalValue *LegalValueLevel `json:"legislationlegalvalue"`
}

// PlaceOfWorship is a generated struct representing the https://schema.org/PlaceOfWorship class.
type PlaceOfWorship struct {
}

// HotelRoom is a generated struct representing the https://schema.org/HotelRoom class.
type HotelRoom struct {
	Bed       *Text              `json:"bed"`
	Occupancy *QuantitativeValue `json:"occupancy"`
}

// Painting is a generated struct representing the https://schema.org/Painting class.
type Painting struct {
}

// Zoo is a generated struct representing the https://schema.org/Zoo class.
type Zoo struct {
}

// Artery is a generated struct representing the https://schema.org/Artery class.
type Artery struct {
	SupplyTo       *AnatomicalStructure `json:"supplyto"`
	ArterialBranch *AnatomicalStructure `json:"arterialbranch"`
}

// MusicAlbumReleaseType is a generated struct representing the https://schema.org/MusicAlbumReleaseType class.
type MusicAlbumReleaseType struct {
}

// RadioSeries is a generated struct representing the https://schema.org/RadioSeries class.
type RadioSeries struct {
	Actors            *Person             `json:"actors"`
	Directors         *Person             `json:"directors"`
	Episodes          *Episode            `json:"episodes"`
	MusicBy           *Person             `json:"musicby"`
	NumberOfEpisodes  *Integer            `json:"numberofepisodes"`
	NumberOfSeasons   *Integer            `json:"numberofseasons"`
	ProductionCompany *Organization       `json:"productioncompany"`
	Seasons           *CreativeWorkSeason `json:"seasons"`
	Trailer           *VideoObject        `json:"trailer"`
	ContainsSeason    *CreativeWorkSeason `json:"containsseason"`
	Director          *Person             `json:"director"`
	Episode           *Episode            `json:"episode"`
	Season            *URL                `json:"season"`
	Actor             *Person             `json:"actor"`
}

// RoofingContractor is a generated struct representing the https://schema.org/RoofingContractor class.
type RoofingContractor struct {
}

// TennisComplex is a generated struct representing the https://schema.org/TennisComplex class.
type TennisComplex struct {
}

// Distillery is a generated struct representing the https://schema.org/Distillery class.
type Distillery struct {
}

// EducationalAudience is a generated struct representing the https://schema.org/EducationalAudience class.
type EducationalAudience struct {
	EducationalRole *Text `json:"educationalrole"`
}

// CategoryCodeSet is a generated struct representing the https://schema.org/CategoryCodeSet class.
type CategoryCodeSet struct {
	HasCategoryCode *CategoryCode `json:"hascategorycode"`
}

// EndorsementRating is a generated struct representing the https://schema.org/EndorsementRating class.
type EndorsementRating struct {
}

// HomeGoodsStore is a generated struct representing the https://schema.org/HomeGoodsStore class.
type HomeGoodsStore struct {
}

// LakeBodyOfWater is a generated struct representing the https://schema.org/LakeBodyOfWater class.
type LakeBodyOfWater struct {
}

// OfferForPurchase is a generated struct representing the https://schema.org/OfferForPurchase class.
type OfferForPurchase struct {
}

// CheckoutPage is a generated struct representing the https://schema.org/CheckoutPage class.
type CheckoutPage struct {
}

// MonetaryGrant is a generated struct representing the https://schema.org/MonetaryGrant class.
type MonetaryGrant struct {
	Amount *Number `json:"amount"`
	Funder *Person `json:"funder"`
}

// Optician is a generated struct representing the https://schema.org/Optician class.
type Optician struct {
}

// PerformAction is a generated struct representing the https://schema.org/PerformAction class.
type PerformAction struct {
	EntertainmentBusiness *EntertainmentBusiness `json:"entertainmentbusiness"`
}

// MusicAlbumProductionType is a generated struct representing the https://schema.org/MusicAlbumProductionType class.
type MusicAlbumProductionType struct {
}

// DeliveryMethod is a generated struct representing the https://schema.org/DeliveryMethod class.
type DeliveryMethod struct {
}

// Restaurant is a generated struct representing the https://schema.org/Restaurant class.
type Restaurant struct {
}

// Manuscript is a generated struct representing the https://schema.org/Manuscript class.
type Manuscript struct {
}

// MoveAction is a generated struct representing the https://schema.org/MoveAction class.
type MoveAction struct {
	FromLocation *Place `json:"fromlocation"`
	ToLocation   *Place `json:"tolocation"`
}

// USNonprofitType is a generated struct representing the https://schema.org/USNonprofitType class.
type USNonprofitType struct {
}

// EducationalOrganization is a generated struct representing the https://schema.org/EducationalOrganization class.
type EducationalOrganization struct {
	Alumni *Person `json:"alumni"`
}

// LinkRole is a generated struct representing the https://schema.org/LinkRole class.
type LinkRole struct {
	LinkRelationship *Text `json:"linkrelationship"`
	InLanguage       *Text `json:"inlanguage"`
}

// MediaManipulationRatingEnumeration is a generated struct representing the https://schema.org/MediaManipulationRatingEnumeration class.
type MediaManipulationRatingEnumeration struct {
}

// PerformanceRole is a generated struct representing the https://schema.org/PerformanceRole class.
type PerformanceRole struct {
	CharacterName *Text `json:"charactername"`
}

// SizeGroupEnumeration is a generated struct representing the https://schema.org/SizeGroupEnumeration class.
type SizeGroupEnumeration struct {
}

// WarrantyPromise is a generated struct representing the https://schema.org/WarrantyPromise class.
type WarrantyPromise struct {
	DurationOfWarranty *QuantitativeValue `json:"durationofwarranty"`
	WarrantyScope      *WarrantyScope     `json:"warrantyscope"`
}

// ComedyEvent is a generated struct representing the https://schema.org/ComedyEvent class.
type ComedyEvent struct {
}

// SteeringPositionValue is a generated struct representing the https://schema.org/SteeringPositionValue class.
type SteeringPositionValue struct {
}

// TVEpisode is a generated struct representing the https://schema.org/TVEpisode class.
type TVEpisode struct {
	CountryOfOrigin  *Country  `json:"countryoforigin"`
	PartOfTVSeries   *TVSeries `json:"partoftvseries"`
	SubtitleLanguage *Text     `json:"subtitlelanguage"`
	TitleEIDR        *URL      `json:"titleeidr"`
}

// Chapter is a generated struct representing the https://schema.org/Chapter class.
type Chapter struct {
	PageEnd    *Text `json:"pageend"`
	PageStart  *Text `json:"pagestart"`
	Pagination *Text `json:"pagination"`
}

// Dataset is a generated struct representing the https://schema.org/Dataset class.
type Dataset struct {
	Catalog               *DataCatalog  `json:"catalog"`
	DatasetTimeInterval   *DateTime     `json:"datasettimeinterval"`
	Distribution          *DataDownload `json:"distribution"`
	IncludedDataCatalog   *DataCatalog  `json:"includeddatacatalog"`
	Issn                  *Text         `json:"issn"`
	MeasurementMethod     *URL          `json:"measurementmethod"`
	VariableMeasured      *Text         `json:"variablemeasured"`
	MeasurementTechnique  *URL          `json:"measurementtechnique"`
	IncludedInDataCatalog *DataCatalog  `json:"includedindatacatalog"`
}

// VideoGameClip is a generated struct representing the https://schema.org/VideoGameClip class.
type VideoGameClip struct {
}

// DownloadAction is a generated struct representing the https://schema.org/DownloadAction class.
type DownloadAction struct {
}

// ReturnLabelSourceEnumeration is a generated struct representing the https://schema.org/ReturnLabelSourceEnumeration class.
type ReturnLabelSourceEnumeration struct {
}

// WorkersUnion is a generated struct representing the https://schema.org/WorkersUnion class.
type WorkersUnion struct {
}

// IceCreamShop is a generated struct representing the https://schema.org/IceCreamShop class.
type IceCreamShop struct {
}

// Physician is a generated struct representing the https://schema.org/Physician class.
type Physician struct {
	AvailableService    *MedicalTherapy   `json:"availableservice"`
	HospitalAffiliation *Hospital         `json:"hospitalaffiliation"`
	MedicalSpecialty    *MedicalSpecialty `json:"medicalspecialty"`
}

// MathSolver is a generated struct representing the https://schema.org/MathSolver class.
type MathSolver struct {
	MathExpression *Text `json:"mathexpression"`
}

// AutoRental is a generated struct representing the https://schema.org/AutoRental class.
type AutoRental struct {
}

// ExercisePlan is a generated struct representing the https://schema.org/ExercisePlan class.
type ExercisePlan struct {
	ActivityDuration   *QuantitativeValue `json:"activityduration"`
	ActivityFrequency  *Text              `json:"activityfrequency"`
	AdditionalVariable *Text              `json:"additionalvariable"`
	ExerciseType       *Text              `json:"exercisetype"`
	Intensity          *Text              `json:"intensity"`
	Repetitions        *QuantitativeValue `json:"repetitions"`
	RestPeriods        *Text              `json:"restperiods"`
	Workload           *QuantitativeValue `json:"workload"`
}

// FoodService is a generated struct representing the https://schema.org/FoodService class.
type FoodService struct {
}

// MedicalStudy is a generated struct representing the https://schema.org/MedicalStudy class.
type MedicalStudy struct {
	HealthCondition *MedicalCondition   `json:"healthcondition"`
	Status          *Text               `json:"status"`
	StudyLocation   *AdministrativeArea `json:"studylocation"`
	StudySubject    *MedicalEntity      `json:"studysubject"`
	Sponsor         *Person             `json:"sponsor"`
}

// MusicEvent is a generated struct representing the https://schema.org/MusicEvent class.
type MusicEvent struct {
}

// TextObject is a generated struct representing the https://schema.org/TextObject class.
type TextObject struct {
}

// UserLikes is a generated struct representing the https://schema.org/UserLikes class.
type UserLikes struct {
}

// AcceptAction is a generated struct representing the https://schema.org/AcceptAction class.
type AcceptAction struct {
}

// ActionStatusType is a generated struct representing the https://schema.org/ActionStatusType class.
type ActionStatusType struct {
}

// RestrictedDiet is a generated struct representing the https://schema.org/RestrictedDiet class.
type RestrictedDiet struct {
}

// DeliveryTimeSettings is a generated struct representing the https://schema.org/DeliveryTimeSettings class.
type DeliveryTimeSettings struct {
	DeliveryTime         *ShippingDeliveryTime `json:"deliverytime"`
	IsUnlabelledFallback *Boolean              `json:"isunlabelledfallback"`
	ShippingDestination  *DefinedRegion        `json:"shippingdestination"`
	TransitTimeLabel     *Text                 `json:"transittimelabel"`
}

// DislikeAction is a generated struct representing the https://schema.org/DislikeAction class.
type DislikeAction struct {
}

// HomeAndConstructionBusiness is a generated struct representing the https://schema.org/HomeAndConstructionBusiness class.
type HomeAndConstructionBusiness struct {
}

// MolecularEntity is a generated struct representing the https://schema.org/MolecularEntity class.
type MolecularEntity struct {
	ChemicalRole                *DefinedTerm `json:"chemicalrole"`
	InChI                       *Text        `json:"inchi"`
	InChIKey                    *Text        `json:"inchikey"`
	IupacName                   *Text        `json:"iupacname"`
	MolecularFormula            *Text        `json:"molecularformula"`
	MolecularWeight             *Text        `json:"molecularweight"`
	MonoisotopicMolecularWeight *Text        `json:"monoisotopicmolecularweight"`
	PotentialUse                *DefinedTerm `json:"potentialuse"`
	Smiles                      *Text        `json:"smiles"`
}

// Electrician is a generated struct representing the https://schema.org/Electrician class.
type Electrician struct {
}

// Observation is a generated struct representing the https://schema.org/Observation class.
type Observation struct {
	MarginOfError          *QuantitativeValue   `json:"marginoferror"`
	MeasuredProperty       *Property            `json:"measuredproperty"`
	MeasurementDenominator *StatisticalVariable `json:"measurementdenominator"`
	MeasurementMethod      *URL                 `json:"measurementmethod"`
	MeasurementQualifier   *Enumeration         `json:"measurementqualifier"`
	ObservationAbout       *Thing               `json:"observationabout"`
	ObservationDate        *DateTime            `json:"observationdate"`
	ObservationPeriod      *Text                `json:"observationperiod"`
	VariableMeasured       *Text                `json:"variablemeasured"`
	MeasurementTechnique   *URL                 `json:"measurementtechnique"`
}

// Organization is a generated struct representing the https://schema.org/Organization class.
type Organization struct {
	ActionableFeedbackPolicy  *URL                               `json:"actionablefeedbackpolicy"`
	Address                   *Text                              `json:"address"`
	AgentInteractionStatistic *InteractionCounter                `json:"agentinteractionstatistic"`
	AggregateRating           *AggregateRating                   `json:"aggregaterating"`
	Awards                    *Text                              `json:"awards"`
	Brand                     *Organization                      `json:"brand"`
	ContactPoints             *ContactPoint                      `json:"contactpoints"`
	CorrectionsPolicy         *URL                               `json:"correctionspolicy"`
	Department                *Organization                      `json:"department"`
	DissolutionDate           *Date                              `json:"dissolutiondate"`
	DiversityPolicy           *URL                               `json:"diversitypolicy"`
	DiversityStaffingReport   *URL                               `json:"diversitystaffingreport"`
	Duns                      *Text                              `json:"duns"`
	Email                     *Text                              `json:"email"`
	Employees                 *Person                            `json:"employees"`
	EthicsPolicy              *URL                               `json:"ethicspolicy"`
	Events                    *Event                             `json:"events"`
	FaxNumber                 *Text                              `json:"faxnumber"`
	Founders                  *Person                            `json:"founders"`
	FoundingDate              *Date                              `json:"foundingdate"`
	FoundingLocation          *Place                             `json:"foundinglocation"`
	Funder                    *Person                            `json:"funder"`
	GlobalLocationNumber      *Text                              `json:"globallocationnumber"`
	HasCredential             *EducationalOccupationalCredential `json:"hascredential"`
	HasMerchantReturnPolicy   *MerchantReturnPolicy              `json:"hasmerchantreturnpolicy"`
	HasOfferCatalog           *OfferCatalog                      `json:"hasoffercatalog"`
	HasPOS                    *Place                             `json:"haspos"`
	IsicV4                    *Text                              `json:"isicv4"`
	Iso6523Code               *Text                              `json:"iso6523code"`
	Keywords                  *URL                               `json:"keywords"`
	KnowsAbout                *URL                               `json:"knowsabout"`
	KnowsLanguage             *Text                              `json:"knowslanguage"`
	LegalName                 *Text                              `json:"legalname"`
	LeiCode                   *Text                              `json:"leicode"`
	Logo                      *URL                               `json:"logo"`
	Members                   *Person                            `json:"members"`
	Naics                     *Text                              `json:"naics"`
	NonprofitStatus           *NonprofitType                     `json:"nonprofitstatus"`
	NumberOfEmployees         *QuantitativeValue                 `json:"numberofemployees"`
	OwnershipFundingInfo      *URL                               `json:"ownershipfundinginfo"`
	Owns                      *Product                           `json:"owns"`
	Reviews                   *Review                            `json:"reviews"`
	Seeks                     *Demand                            `json:"seeks"`
	Slogan                    *Text                              `json:"slogan"`
	TaxID                     *Text                              `json:"taxid"`
	Telephone                 *Text                              `json:"telephone"`
	UnnamedSourcesPolicy      *URL                               `json:"unnamedsourcespolicy"`
	VatID                     *Text                              `json:"vatid"`
	Alumni                    *Person                            `json:"alumni"`
	Award                     *Text                              `json:"award"`
	ContactPoint              *ContactPoint                      `json:"contactpoint"`
	Employee                  *Person                            `json:"employee"`
	Event                     *Event                             `json:"event"`
	Founder                   *Person                            `json:"founder"`
	Funding                   *Grant                             `json:"funding"`
	InteractionStatistic      *InteractionCounter                `json:"interactionstatistic"`
	MakesOffer                *Offer                             `json:"makesoffer"`
	Review                    *Review                            `json:"review"`
	ServiceArea               *Place                             `json:"servicearea"`
	Sponsor                   *Person                            `json:"sponsor"`
	SubOrganization           *Organization                      `json:"suborganization"`
	MemberOf                  *ProgramMembership                 `json:"memberof"`
	ParentOrganization        *Organization                      `json:"parentorganization"`
	AreaServed                *Text                              `json:"areaserved"`
	Member                    *Person                            `json:"member"`
	PublishingPrinciples      *URL                               `json:"publishingprinciples"`
	Location                  *VirtualLocation                   `json:"location"`
}

// Syllabus is a generated struct representing the https://schema.org/Syllabus class.
type Syllabus struct {
}

// TradeAction is a generated struct representing the https://schema.org/TradeAction class.
type TradeAction struct {
	Price              *Text               `json:"price"`
	PriceCurrency      *Text               `json:"pricecurrency"`
	PriceSpecification *PriceSpecification `json:"pricespecification"`
}

// AgreeAction is a generated struct representing the https://schema.org/AgreeAction class.
type AgreeAction struct {
}

// EmailMessage is a generated struct representing the https://schema.org/EmailMessage class.
type EmailMessage struct {
}

// Festival is a generated struct representing the https://schema.org/Festival class.
type Festival struct {
}

// LearningResource is a generated struct representing the https://schema.org/LearningResource class.
type LearningResource struct {
	Assesses             *Text            `json:"assesses"`
	CompetencyRequired   *URL             `json:"competencyrequired"`
	EducationalAlignment *AlignmentObject `json:"educationalalignment"`
	EducationalLevel     *URL             `json:"educationallevel"`
	EducationalUse       *Text            `json:"educationaluse"`
	LearningResourceType *Text            `json:"learningresourcetype"`
	Teaches              *Text            `json:"teaches"`
}

// NightClub is a generated struct representing the https://schema.org/NightClub class.
type NightClub struct {
}

// ActionAccessSpecification is a generated struct representing the https://schema.org/ActionAccessSpecification class.
type ActionAccessSpecification struct {
	AvailabilityEnds     *Time              `json:"availabilityends"`
	AvailabilityStarts   *Time              `json:"availabilitystarts"`
	EligibleRegion       *Text              `json:"eligibleregion"`
	ExpectsAcceptanceOf  *Offer             `json:"expectsacceptanceof"`
	IneligibleRegion     *Text              `json:"ineligibleregion"`
	RequiresSubscription *MediaSubscription `json:"requiressubscription"`
	Category             *URL               `json:"category"`
}

// Drawing is a generated struct representing the https://schema.org/Drawing class.
type Drawing struct {
}

// ApplyAction is a generated struct representing the https://schema.org/ApplyAction class.
type ApplyAction struct {
}

// InsuranceAgency is a generated struct representing the https://schema.org/InsuranceAgency class.
type InsuranceAgency struct {
}

// MovingCompany is a generated struct representing the https://schema.org/MovingCompany class.
type MovingCompany struct {
}

// NonprofitType is a generated struct representing the https://schema.org/NonprofitType class.
type NonprofitType struct {
}

// OceanBodyOfWater is a generated struct representing the https://schema.org/OceanBodyOfWater class.
type OceanBodyOfWater struct {
}

// SportsEvent is a generated struct representing the https://schema.org/SportsEvent class.
type SportsEvent struct {
	AwayTeam   *SportsTeam `json:"awayteam"`
	HomeTeam   *SportsTeam `json:"hometeam"`
	Sport      *URL        `json:"sport"`
	Competitor *SportsTeam `json:"competitor"`
}

// BuddhistTemple is a generated struct representing the https://schema.org/BuddhistTemple class.
type BuddhistTemple struct {
}

// CreativeWork is a generated struct representing the https://schema.org/CreativeWork class.
type CreativeWork struct {
	Abstract             *Text               `json:"abstract"`
	AccessMode           *Text               `json:"accessmode"`
	AccessModeSufficient *ItemList           `json:"accessmodesufficient"`
	AccessibilityAPI     *Text               `json:"accessibilityapi"`
	AccessibilityControl *Text               `json:"accessibilitycontrol"`
	AccessibilityFeature *Text               `json:"accessibilityfeature"`
	AccessibilityHazard  *Text               `json:"accessibilityhazard"`
	AccessibilitySummary *Text               `json:"accessibilitysummary"`
	AccountablePerson    *Person             `json:"accountableperson"`
	AcquireLicensePage   *URL                `json:"acquirelicensepage"`
	AggregateRating      *AggregateRating    `json:"aggregaterating"`
	AlternativeHeadline  *Text               `json:"alternativeheadline"`
	ArchivedAt           *WebPage            `json:"archivedat"`
	Assesses             *Text               `json:"assesses"`
	AssociatedMedia      *MediaObject        `json:"associatedmedia"`
	Audio                *MusicRecording     `json:"audio"`
	Author               *Person             `json:"author"`
	Awards               *Text               `json:"awards"`
	Character            *Person             `json:"character"`
	Citation             *Text               `json:"citation"`
	Comment              *Comment            `json:"comment"`
	CommentCount         *Integer            `json:"commentcount"`
	ConditionsOfAccess   *Text               `json:"conditionsofaccess"`
	ContentRating        *Text               `json:"contentrating"`
	ContentReferenceTime *DateTime           `json:"contentreferencetime"`
	Contributor          *Person             `json:"contributor"`
	CopyrightHolder      *Person             `json:"copyrightholder"`
	CopyrightNotice      *Text               `json:"copyrightnotice"`
	CopyrightYear        *Number             `json:"copyrightyear"`
	Correction           *URL                `json:"correction"`
	CountryOfOrigin      *Country            `json:"countryoforigin"`
	CreativeWorkStatus   *Text               `json:"creativeworkstatus"`
	CreditText           *Text               `json:"credittext"`
	DateModified         *DateTime           `json:"datemodified"`
	DatePublished        *DateTime           `json:"datepublished"`
	DiscussionUrl        *URL                `json:"discussionurl"`
	EditEIDR             *URL                `json:"editeidr"`
	Editor               *Person             `json:"editor"`
	EducationalAlignment *AlignmentObject    `json:"educationalalignment"`
	EducationalLevel     *URL                `json:"educationallevel"`
	EducationalUse       *Text               `json:"educationaluse"`
	Encodings            *MediaObject        `json:"encodings"`
	Expires              *DateTime           `json:"expires"`
	FileFormat           *URL                `json:"fileformat"`
	Funder               *Person             `json:"funder"`
	Headline             *Text               `json:"headline"`
	InteractivityType    *Text               `json:"interactivitytype"`
	InterpretedAsClaim   *Claim              `json:"interpretedasclaim"`
	IsBasedOnUrl         *URL                `json:"isbasedonurl"`
	IsFamilyFriendly     *Boolean            `json:"isfamilyfriendly"`
	Keywords             *URL                `json:"keywords"`
	LearningResourceType *Text               `json:"learningresourcetype"`
	License              *URL                `json:"license"`
	LocationCreated      *Place              `json:"locationcreated"`
	Maintainer           *Person             `json:"maintainer"`
	MaterialExtent       *Text               `json:"materialextent"`
	Mentions             *Thing              `json:"mentions"`
	Pattern              *Text               `json:"pattern"`
	Producer             *Person             `json:"producer"`
	Publication          *PublicationEvent   `json:"publication"`
	Publisher            *Person             `json:"publisher"`
	PublisherImprint     *Organization       `json:"publisherimprint"`
	ReleasedEvent        *PublicationEvent   `json:"releasedevent"`
	Reviews              *Review             `json:"reviews"`
	SchemaVersion        *URL                `json:"schemaversion"`
	SdDatePublished      *Date               `json:"sddatepublished"`
	SdLicense            *URL                `json:"sdlicense"`
	SdPublisher          *Person             `json:"sdpublisher"`
	Size                 *Text               `json:"size"`
	SourceOrganization   *Organization       `json:"sourceorganization"`
	Spatial              *Place              `json:"spatial"`
	Teaches              *Text               `json:"teaches"`
	Temporal             *Text               `json:"temporal"`
	Text                 *Text               `json:"text"`
	Thumbnail            *ImageObject        `json:"thumbnail"`
	ThumbnailUrl         *URL                `json:"thumbnailurl"`
	TimeRequired         *Duration           `json:"timerequired"`
	Translator           *Person             `json:"translator"`
	TypicalAgeRange      *Text               `json:"typicalagerange"`
	Version              *Text               `json:"version"`
	Video                *VideoObject        `json:"video"`
	Audience             *Audience           `json:"audience"`
	Award                *Text               `json:"award"`
	ContentLocation      *Place              `json:"contentlocation"`
	Creator              *Person             `json:"creator"`
	DateCreated          *DateTime           `json:"datecreated"`
	EncodingFormat       *URL                `json:"encodingformat"`
	ExampleOfWork        *CreativeWork       `json:"exampleofwork"`
	Funding              *Grant              `json:"funding"`
	Genre                *URL                `json:"genre"`
	InLanguage           *Text               `json:"inlanguage"`
	InteractionStatistic *InteractionCounter `json:"interactionstatistic"`
	IsAccessibleForFree  *Boolean            `json:"isaccessibleforfree"`
	IsBasedOn            *URL                `json:"isbasedon"`
	MainEntity           *Thing              `json:"mainentity"`
	Offers               *Offer              `json:"offers"`
	Provider             *Person             `json:"provider"`
	RecordedAt           *Event              `json:"recordedat"`
	Review               *Review             `json:"review"`
	Sponsor              *Person             `json:"sponsor"`
	TemporalCoverage     *URL                `json:"temporalcoverage"`
	TranslationOfWork    *CreativeWork       `json:"translationofwork"`
	UsageInfo            *URL                `json:"usageinfo"`
	WorkTranslation      *CreativeWork       `json:"worktranslation"`
	About                *Thing              `json:"about"`
	Encoding             *MediaObject        `json:"encoding"`
	Material             *URL                `json:"material"`
	SpatialCoverage      *Place              `json:"spatialcoverage"`
	WorkExample          *CreativeWork       `json:"workexample"`
	HasPart              *CreativeWork       `json:"haspart"`
	Position             *Text               `json:"position"`
	IsPartOf             *URL                `json:"ispartof"`
	PublishingPrinciples *URL                `json:"publishingprinciples"`
}

// Florist is a generated struct representing the https://schema.org/Florist class.
type Florist struct {
}

// Joint is a generated struct representing the https://schema.org/Joint class.
type Joint struct {
	BiomechnicalClass *Text `json:"biomechnicalclass"`
	FunctionalClass   *Text `json:"functionalclass"`
	StructuralClass   *Text `json:"structuralclass"`
}

// ListItem is a generated struct representing the https://schema.org/ListItem class.
type ListItem struct {
	Item         *Thing    `json:"item"`
	NextItem     *ListItem `json:"nextitem"`
	PreviousItem *ListItem `json:"previousitem"`
	Position     *Text     `json:"position"`
}

// TouristTrip is a generated struct representing the https://schema.org/TouristTrip class.
type TouristTrip struct {
	TouristType *Text `json:"touristtype"`
}

// BusOrCoach is a generated struct representing the https://schema.org/BusOrCoach class.
type BusOrCoach struct {
	AcrissCode *Text              `json:"acrisscode"`
	RoofLoad   *QuantitativeValue `json:"roofload"`
}

// ChooseAction is a generated struct representing the https://schema.org/ChooseAction class.
type ChooseAction struct {
	Option       *Thing `json:"option"`
	ActionOption *Thing `json:"actionoption"`
}

// EndorseAction is a generated struct representing the https://schema.org/EndorseAction class.
type EndorseAction struct {
	Endorsee *Person `json:"endorsee"`
}

// IndividualProduct is a generated struct representing the https://schema.org/IndividualProduct class.
type IndividualProduct struct {
	SerialNumber *Text `json:"serialnumber"`
}

// MusicPlaylist is a generated struct representing the https://schema.org/MusicPlaylist class.
type MusicPlaylist struct {
	NumTracks *Integer        `json:"numtracks"`
	Tracks    *MusicRecording `json:"tracks"`
	Track     *MusicRecording `json:"track"`
}

// QualitativeValue is a generated struct representing the https://schema.org/QualitativeValue class.
type QualitativeValue struct {
	AdditionalProperty *PropertyValue    `json:"additionalproperty"`
	Equal              *QualitativeValue `json:"equal"`
	Greater            *QualitativeValue `json:"greater"`
	GreaterOrEqual     *QualitativeValue `json:"greaterorequal"`
	Lesser             *QualitativeValue `json:"lesser"`
	LesserOrEqual      *QualitativeValue `json:"lesserorequal"`
	NonEqual           *QualitativeValue `json:"nonequal"`
	ValueReference     *Text             `json:"valuereference"`
}

// Suite is a generated struct representing the https://schema.org/Suite class.
type Suite struct {
	Bed           *Text              `json:"bed"`
	NumberOfRooms *QuantitativeValue `json:"numberofrooms"`
	Occupancy     *QuantitativeValue `json:"occupancy"`
}

// CivicStructure is a generated struct representing the https://schema.org/CivicStructure class.
type CivicStructure struct {
	OpeningHours *Text `json:"openinghours"`
}

// EmployeeRole is a generated struct representing the https://schema.org/EmployeeRole class.
type EmployeeRole struct {
	BaseSalary     *PriceSpecification `json:"basesalary"`
	SalaryCurrency *Text               `json:"salarycurrency"`
}

// Preschool is a generated struct representing the https://schema.org/Preschool class.
type Preschool struct {
}

// WearAction is a generated struct representing the https://schema.org/WearAction class.
type WearAction struct {
}

// FoodEstablishmentReservation is a generated struct representing the https://schema.org/FoodEstablishmentReservation class.
type FoodEstablishmentReservation struct {
	EndTime   *Time              `json:"endtime"`
	PartySize *QuantitativeValue `json:"partysize"`
	StartTime *Time              `json:"starttime"`
}

// Role is a generated struct representing the https://schema.org/Role class.
type Role struct {
	EndDate       *DateTime `json:"enddate"`
	NamedPosition *URL      `json:"namedposition"`
	StartDate     *DateTime `json:"startdate"`
	RoleName      *URL      `json:"rolename"`
}

// ExerciseGym is a generated struct representing the https://schema.org/ExerciseGym class.
type ExerciseGym struct {
}

// WriteAction is a generated struct representing the https://schema.org/WriteAction class.
type WriteAction struct {
	Language   *Language `json:"language"`
	InLanguage *Text     `json:"inlanguage"`
}
