package schema_org

// AggregateRating is a generated struct representing the https://schema.org/AggregateRating class.
type AggregateRating struct {
	ItemReviewed *Thing   `json:"itemreviewed"` //https://schema.org/Thing
	RatingCount  *Integer `json:"ratingcount"`  //https://schema.org/Integer
	ReviewCount  *Integer `json:"reviewcount"`  //https://schema.org/Integer
}

// ClothingStore is a generated struct representing the https://schema.org/ClothingStore class.
type ClothingStore struct {
}

// Time is a generated struct representing the https://schema.org/Time class.
type Time struct {
}

// BroadcastChannel is a generated struct representing the https://schema.org/BroadcastChannel class.
type BroadcastChannel struct {
	BroadcastChannelId       *Text                    `json:"broadcastchannelid"`       //https://schema.org/Text
	BroadcastFrequency       *Text                    `json:"broadcastfrequency"`       //https://schema.org/Text
	BroadcastServiceTier     *Text                    `json:"broadcastservicetier"`     //https://schema.org/Text
	InBroadcastLineup        *CableOrSatelliteService `json:"inbroadcastlineup"`        //https://schema.org/CableOrSatelliteService
	Genre                    *URL                     `json:"genre"`                    //https://schema.org/URL
	ProvidesBroadcastService *BroadcastService        `json:"providesbroadcastservice"` //https://schema.org/BroadcastService
}

// BarOrPub is a generated struct representing the https://schema.org/BarOrPub class.
type BarOrPub struct {
}

// BookFormatType is a generated struct representing the https://schema.org/BookFormatType class.
type BookFormatType struct {
}

// CreateAction is a generated struct representing the https://schema.org/CreateAction class.
type CreateAction struct {
}

// MedicalOrganization is a generated struct representing the https://schema.org/MedicalOrganization class.
type MedicalOrganization struct {
	HealthPlanNetworkId    *Text             `json:"healthplannetworkid"`    //https://schema.org/Text
	IsAcceptingNewPatients *Boolean          `json:"isacceptingnewpatients"` //https://schema.org/Boolean
	MedicalSpecialty       *MedicalSpecialty `json:"medicalspecialty"`       //https://schema.org/MedicalSpecialty
}

// OutletStore is a generated struct representing the https://schema.org/OutletStore class.
type OutletStore struct {
}

// BookSeries is a generated struct representing the https://schema.org/BookSeries class.
type BookSeries struct {
}

// PayAction is a generated struct representing the https://schema.org/PayAction class.
type PayAction struct {
	Recipient *Person `json:"recipient"` //https://schema.org/Person
}

// Role is a generated struct representing the https://schema.org/Role class.
type Role struct {
	EndDate       *DateTime `json:"enddate"`       //https://schema.org/DateTime
	NamedPosition *URL      `json:"namedposition"` //https://schema.org/URL
	StartDate     *DateTime `json:"startdate"`     //https://schema.org/DateTime
	RoleName      *URL      `json:"rolename"`      //https://schema.org/URL
}

// UserReview is a generated struct representing the https://schema.org/UserReview class.
type UserReview struct {
}

// AskAction is a generated struct representing the https://schema.org/AskAction class.
type AskAction struct {
	Question *Question `json:"question"` //https://schema.org/Question
}

// BefriendAction is a generated struct representing the https://schema.org/BefriendAction class.
type BefriendAction struct {
}

// BusinessFunction is a generated struct representing the https://schema.org/BusinessFunction class.
type BusinessFunction struct {
}

// MusicAlbumProductionType is a generated struct representing the https://schema.org/MusicAlbumProductionType class.
type MusicAlbumProductionType struct {
}

// PetStore is a generated struct representing the https://schema.org/PetStore class.
type PetStore struct {
}

// ReceiveAction is a generated struct representing the https://schema.org/ReceiveAction class.
type ReceiveAction struct {
	DeliveryMethod *DeliveryMethod `json:"deliverymethod"` //https://schema.org/DeliveryMethod
	Sender         *Person         `json:"sender"`         //https://schema.org/Person
}

// Gene is a generated struct representing the https://schema.org/Gene class.
type Gene struct {
	AlternativeOf         *Gene          `json:"alternativeof"`         //https://schema.org/Gene
	ExpressedIn           *DefinedTerm   `json:"expressedin"`           //https://schema.org/DefinedTerm
	HasBioPolymerSequence *Text          `json:"hasbiopolymersequence"` //https://schema.org/Text
	EncodesBioChemEntity  *BioChemEntity `json:"encodesbiochementity"`  //https://schema.org/BioChemEntity
}

// ItemListOrderType is a generated struct representing the https://schema.org/ItemListOrderType class.
type ItemListOrderType struct {
}

// StatisticalPopulation is a generated struct representing the https://schema.org/StatisticalPopulation class.
type StatisticalPopulation struct {
	PopulationType *Class `json:"populationtype"` //https://schema.org/Class
}

// AnatomicalStructure is a generated struct representing the https://schema.org/AnatomicalStructure class.
type AnatomicalStructure struct {
	AssociatedPathophysiology *Text                `json:"associatedpathophysiology"` //https://schema.org/Text
	BodyLocation              *Text                `json:"bodylocation"`              //https://schema.org/Text
	ConnectedTo               *AnatomicalStructure `json:"connectedto"`               //https://schema.org/AnatomicalStructure
	Diagram                   *ImageObject         `json:"diagram"`                   //https://schema.org/ImageObject
	PartOfSystem              *AnatomicalSystem    `json:"partofsystem"`              //https://schema.org/AnatomicalSystem
	RelatedCondition          *MedicalCondition    `json:"relatedcondition"`          //https://schema.org/MedicalCondition
	RelatedTherapy            *MedicalTherapy      `json:"relatedtherapy"`            //https://schema.org/MedicalTherapy
	SubStructure              *AnatomicalStructure `json:"substructure"`              //https://schema.org/AnatomicalStructure
}

// DanceGroup is a generated struct representing the https://schema.org/DanceGroup class.
type DanceGroup struct {
}

// HowToStep is a generated struct representing the https://schema.org/HowToStep class.
type HowToStep struct {
}

// HyperTocEntry is a generated struct representing the https://schema.org/HyperTocEntry class.
type HyperTocEntry struct {
	AssociatedMedia *MediaObject   `json:"associatedmedia"` //https://schema.org/MediaObject
	TocContinuation *HyperTocEntry `json:"toccontinuation"` //https://schema.org/HyperTocEntry
	Utterances      *Text          `json:"utterances"`      //https://schema.org/Text
}

// InstallAction is a generated struct representing the https://schema.org/InstallAction class.
type InstallAction struct {
}

// ReportageNewsArticle is a generated struct representing the https://schema.org/ReportageNewsArticle class.
type ReportageNewsArticle struct {
}

// CityHall is a generated struct representing the https://schema.org/CityHall class.
type CityHall struct {
}

// CoverArt is a generated struct representing the https://schema.org/CoverArt class.
type CoverArt struct {
}

// PropertyValueSpecification is a generated struct representing the https://schema.org/PropertyValueSpecification class.
type PropertyValueSpecification struct {
	DefaultValue   *Thing   `json:"defaultvalue"`   //https://schema.org/Thing
	MaxValue       *Number  `json:"maxvalue"`       //https://schema.org/Number
	MinValue       *Number  `json:"minvalue"`       //https://schema.org/Number
	MultipleValues *Boolean `json:"multiplevalues"` //https://schema.org/Boolean
	ReadonlyValue  *Boolean `json:"readonlyvalue"`  //https://schema.org/Boolean
	StepValue      *Number  `json:"stepvalue"`      //https://schema.org/Number
	ValueMaxLength *Number  `json:"valuemaxlength"` //https://schema.org/Number
	ValueMinLength *Number  `json:"valueminlength"` //https://schema.org/Number
	ValueName      *Text    `json:"valuename"`      //https://schema.org/Text
	ValuePattern   *Text    `json:"valuepattern"`   //https://schema.org/Text
	ValueRequired  *Boolean `json:"valuerequired"`  //https://schema.org/Boolean
}

// Artery is a generated struct representing the https://schema.org/Artery class.
type Artery struct {
	SupplyTo       *AnatomicalStructure `json:"supplyto"`       //https://schema.org/AnatomicalStructure
	ArterialBranch *AnatomicalStructure `json:"arterialbranch"` //https://schema.org/AnatomicalStructure
}

// FundingAgency is a generated struct representing the https://schema.org/FundingAgency class.
type FundingAgency struct {
}

// InteractAction is a generated struct representing the https://schema.org/InteractAction class.
type InteractAction struct {
}

// StructuredValue is a generated struct representing the https://schema.org/StructuredValue class.
type StructuredValue struct {
}

// PostOffice is a generated struct representing the https://schema.org/PostOffice class.
type PostOffice struct {
}

// BackgroundNewsArticle is a generated struct representing the https://schema.org/BackgroundNewsArticle class.
type BackgroundNewsArticle struct {
}

// PublicationEvent is a generated struct representing the https://schema.org/PublicationEvent class.
type PublicationEvent struct {
	Free        *Boolean          `json:"free"`        //https://schema.org/Boolean
	PublishedBy *Person           `json:"publishedby"` //https://schema.org/Person
	PublishedOn *BroadcastService `json:"publishedon"` //https://schema.org/BroadcastService
}

// WebApplication is a generated struct representing the https://schema.org/WebApplication class.
type WebApplication struct {
	BrowserRequirements *Text `json:"browserrequirements"` //https://schema.org/Text
}

// DataFeed is a generated struct representing the https://schema.org/DataFeed class.
type DataFeed struct {
	DataFeedElement *Thing `json:"datafeedelement"` //https://schema.org/Thing
}

// InfectiousDisease is a generated struct representing the https://schema.org/InfectiousDisease class.
type InfectiousDisease struct {
	InfectiousAgent      *Text                 `json:"infectiousagent"`      //https://schema.org/Text
	InfectiousAgentClass *InfectiousAgentClass `json:"infectiousagentclass"` //https://schema.org/InfectiousAgentClass
	TransmissionMethod   *Text                 `json:"transmissionmethod"`   //https://schema.org/Text
}

// InternetCafe is a generated struct representing the https://schema.org/InternetCafe class.
type InternetCafe struct {
}

// MedicalCondition is a generated struct representing the https://schema.org/MedicalCondition class.
type MedicalCondition struct {
	AssociatedAnatomy     *SuperficialAnatomy    `json:"associatedanatomy"`     //https://schema.org/SuperficialAnatomy
	DifferentialDiagnosis *DDxElement            `json:"differentialdiagnosis"` //https://schema.org/DDxElement
	Drug                  *Drug                  `json:"drug"`                  //https://schema.org/Drug
	Epidemiology          *Text                  `json:"epidemiology"`          //https://schema.org/Text
	ExpectedPrognosis     *Text                  `json:"expectedprognosis"`     //https://schema.org/Text
	NaturalProgression    *Text                  `json:"naturalprogression"`    //https://schema.org/Text
	Pathophysiology       *Text                  `json:"pathophysiology"`       //https://schema.org/Text
	PossibleComplication  *Text                  `json:"possiblecomplication"`  //https://schema.org/Text
	PossibleTreatment     *MedicalTherapy        `json:"possibletreatment"`     //https://schema.org/MedicalTherapy
	PrimaryPrevention     *MedicalTherapy        `json:"primaryprevention"`     //https://schema.org/MedicalTherapy
	RiskFactor            *MedicalRiskFactor     `json:"riskfactor"`            //https://schema.org/MedicalRiskFactor
	SecondaryPrevention   *MedicalTherapy        `json:"secondaryprevention"`   //https://schema.org/MedicalTherapy
	SignOrSymptom         *MedicalSignOrSymptom  `json:"signorsymptom"`         //https://schema.org/MedicalSignOrSymptom
	Stage                 *MedicalConditionStage `json:"stage"`                 //https://schema.org/MedicalConditionStage
	Status                *Text                  `json:"status"`                //https://schema.org/Text
	TypicalTest           *MedicalTest           `json:"typicaltest"`           //https://schema.org/MedicalTest
}

// DrugClass is a generated struct representing the https://schema.org/DrugClass class.
type DrugClass struct {
	Drug *Drug `json:"drug"` //https://schema.org/Drug
}

// NLNonprofitType is a generated struct representing the https://schema.org/NLNonprofitType class.
type NLNonprofitType struct {
}

// Waterfall is a generated struct representing the https://schema.org/Waterfall class.
type Waterfall struct {
}

// Mountain is a generated struct representing the https://schema.org/Mountain class.
type Mountain struct {
}

// PriceComponentTypeEnumeration is a generated struct representing the https://schema.org/PriceComponentTypeEnumeration class.
type PriceComponentTypeEnumeration struct {
}

// SearchAction is a generated struct representing the https://schema.org/SearchAction class.
type SearchAction struct {
	Query *Text `json:"query"` //https://schema.org/Text
}

// WorkersUnion is a generated struct representing the https://schema.org/WorkersUnion class.
type WorkersUnion struct {
}

// MedicalTherapy is a generated struct representing the https://schema.org/MedicalTherapy class.
type MedicalTherapy struct {
	Contraindication      *Text           `json:"contraindication"`      //https://schema.org/Text
	DuplicateTherapy      *MedicalTherapy `json:"duplicatetherapy"`      //https://schema.org/MedicalTherapy
	SeriousAdverseOutcome *MedicalEntity  `json:"seriousadverseoutcome"` //https://schema.org/MedicalEntity
}

// MerchantReturnPolicy is a generated struct representing the https://schema.org/MerchantReturnPolicy class.
type MerchantReturnPolicy struct {
	AdditionalProperty                      *PropertyValue                        `json:"additionalproperty"`                      //https://schema.org/PropertyValue
	ApplicableCountry                       *Text                                 `json:"applicablecountry"`                       //https://schema.org/Text
	CustomerRemorseReturnFees               *ReturnFeesEnumeration                `json:"customerremorsereturnfees"`               //https://schema.org/ReturnFeesEnumeration
	CustomerRemorseReturnLabelSource        *ReturnLabelSourceEnumeration         `json:"customerremorsereturnlabelsource"`        //https://schema.org/ReturnLabelSourceEnumeration
	CustomerRemorseReturnShippingFeesAmount *MonetaryAmount                       `json:"customerremorsereturnshippingfeesamount"` //https://schema.org/MonetaryAmount
	InStoreReturnsOffered                   *Boolean                              `json:"instorereturnsoffered"`                   //https://schema.org/Boolean
	ItemCondition                           *OfferItemCondition                   `json:"itemcondition"`                           //https://schema.org/OfferItemCondition
	ItemDefectReturnFees                    *ReturnFeesEnumeration                `json:"itemdefectreturnfees"`                    //https://schema.org/ReturnFeesEnumeration
	ItemDefectReturnLabelSource             *ReturnLabelSourceEnumeration         `json:"itemdefectreturnlabelsource"`             //https://schema.org/ReturnLabelSourceEnumeration
	ItemDefectReturnShippingFeesAmount      *MonetaryAmount                       `json:"itemdefectreturnshippingfeesamount"`      //https://schema.org/MonetaryAmount
	MerchantReturnDays                      *Integer                              `json:"merchantreturndays"`                      //https://schema.org/Integer
	MerchantReturnLink                      *URL                                  `json:"merchantreturnlink"`                      //https://schema.org/URL
	RefundType                              *RefundTypeEnumeration                `json:"refundtype"`                              //https://schema.org/RefundTypeEnumeration
	RestockingFee                           *Number                               `json:"restockingfee"`                           //https://schema.org/Number
	ReturnFees                              *ReturnFeesEnumeration                `json:"returnfees"`                              //https://schema.org/ReturnFeesEnumeration
	ReturnLabelSource                       *ReturnLabelSourceEnumeration         `json:"returnlabelsource"`                       //https://schema.org/ReturnLabelSourceEnumeration
	ReturnMethod                            *ReturnMethodEnumeration              `json:"returnmethod"`                            //https://schema.org/ReturnMethodEnumeration
	ReturnPolicyCategory                    *MerchantReturnEnumeration            `json:"returnpolicycategory"`                    //https://schema.org/MerchantReturnEnumeration
	ReturnPolicyCountry                     *Text                                 `json:"returnpolicycountry"`                     //https://schema.org/Text
	ReturnPolicySeasonalOverride            *MerchantReturnPolicySeasonalOverride `json:"returnpolicyseasonaloverride"`            //https://schema.org/MerchantReturnPolicySeasonalOverride
	ReturnShippingFeesAmount                *MonetaryAmount                       `json:"returnshippingfeesamount"`                //https://schema.org/MonetaryAmount
}

// QuantitativeValue is a generated struct representing the https://schema.org/QuantitativeValue class.
type QuantitativeValue struct {
	AdditionalProperty *PropertyValue `json:"additionalproperty"` //https://schema.org/PropertyValue
	MaxValue           *Number        `json:"maxvalue"`           //https://schema.org/Number
	MinValue           *Number        `json:"minvalue"`           //https://schema.org/Number
	UnitCode           *URL           `json:"unitcode"`           //https://schema.org/URL
	UnitText           *Text          `json:"unittext"`           //https://schema.org/Text
	Value              *Text          `json:"value"`              //https://schema.org/Text
	ValueReference     *Text          `json:"valuereference"`     //https://schema.org/Text
}

// ComputerStore is a generated struct representing the https://schema.org/ComputerStore class.
type ComputerStore struct {
}

// DownloadAction is a generated struct representing the https://schema.org/DownloadAction class.
type DownloadAction struct {
}

// ReportedDoseSchedule is a generated struct representing the https://schema.org/ReportedDoseSchedule class.
type ReportedDoseSchedule struct {
}

// BusinessAudience is a generated struct representing the https://schema.org/BusinessAudience class.
type BusinessAudience struct {
	NumberOfEmployees *QuantitativeValue `json:"numberofemployees"` //https://schema.org/QuantitativeValue
	YearlyRevenue     *QuantitativeValue `json:"yearlyrevenue"`     //https://schema.org/QuantitativeValue
	YearsInOperation  *QuantitativeValue `json:"yearsinoperation"`  //https://schema.org/QuantitativeValue
}

// EducationalOrganization is a generated struct representing the https://schema.org/EducationalOrganization class.
type EducationalOrganization struct {
	Alumni *Person `json:"alumni"` //https://schema.org/Person
}

// InvestmentOrDeposit is a generated struct representing the https://schema.org/InvestmentOrDeposit class.
type InvestmentOrDeposit struct {
	Amount *Number `json:"amount"` //https://schema.org/Number
}

// PublicationIssue is a generated struct representing the https://schema.org/PublicationIssue class.
type PublicationIssue struct {
	IssueNumber *Text `json:"issuenumber"` //https://schema.org/Text
	PageEnd     *Text `json:"pageend"`     //https://schema.org/Text
	PageStart   *Text `json:"pagestart"`   //https://schema.org/Text
	Pagination  *Text `json:"pagination"`  //https://schema.org/Text
}

// CriticReview is a generated struct representing the https://schema.org/CriticReview class.
type CriticReview struct {
}

// FinancialService is a generated struct representing the https://schema.org/FinancialService class.
type FinancialService struct {
	FeesAndCommissionsSpecification *URL `json:"feesandcommissionsspecification"` //https://schema.org/URL
}

// Vehicle is a generated struct representing the https://schema.org/Vehicle class.
type Vehicle struct {
	AccelerationTime            *QuantitativeValue     `json:"accelerationtime"`            //https://schema.org/QuantitativeValue
	BodyType                    *URL                   `json:"bodytype"`                    //https://schema.org/URL
	CallSign                    *Text                  `json:"callsign"`                    //https://schema.org/Text
	CargoVolume                 *QuantitativeValue     `json:"cargovolume"`                 //https://schema.org/QuantitativeValue
	DateVehicleFirstRegistered  *Date                  `json:"datevehiclefirstregistered"`  //https://schema.org/Date
	DriveWheelConfiguration     *Text                  `json:"drivewheelconfiguration"`     //https://schema.org/Text
	EmissionsCO2                *Number                `json:"emissionsco2"`                //https://schema.org/Number
	FuelCapacity                *QuantitativeValue     `json:"fuelcapacity"`                //https://schema.org/QuantitativeValue
	FuelConsumption             *QuantitativeValue     `json:"fuelconsumption"`             //https://schema.org/QuantitativeValue
	FuelEfficiency              *QuantitativeValue     `json:"fuelefficiency"`              //https://schema.org/QuantitativeValue
	FuelType                    *URL                   `json:"fueltype"`                    //https://schema.org/URL
	KnownVehicleDamages         *Text                  `json:"knownvehicledamages"`         //https://schema.org/Text
	MeetsEmissionStandard       *URL                   `json:"meetsemissionstandard"`       //https://schema.org/URL
	MileageFromOdometer         *QuantitativeValue     `json:"mileagefromodometer"`         //https://schema.org/QuantitativeValue
	ModelDate                   *Date                  `json:"modeldate"`                   //https://schema.org/Date
	NumberOfAirbags             *Text                  `json:"numberofairbags"`             //https://schema.org/Text
	NumberOfAxles               *QuantitativeValue     `json:"numberofaxles"`               //https://schema.org/QuantitativeValue
	NumberOfDoors               *QuantitativeValue     `json:"numberofdoors"`               //https://schema.org/QuantitativeValue
	NumberOfForwardGears        *QuantitativeValue     `json:"numberofforwardgears"`        //https://schema.org/QuantitativeValue
	NumberOfPreviousOwners      *QuantitativeValue     `json:"numberofpreviousowners"`      //https://schema.org/QuantitativeValue
	Payload                     *QuantitativeValue     `json:"payload"`                     //https://schema.org/QuantitativeValue
	ProductionDate              *Date                  `json:"productiondate"`              //https://schema.org/Date
	PurchaseDate                *Date                  `json:"purchasedate"`                //https://schema.org/Date
	SeatingCapacity             *QuantitativeValue     `json:"seatingcapacity"`             //https://schema.org/QuantitativeValue
	Speed                       *QuantitativeValue     `json:"speed"`                       //https://schema.org/QuantitativeValue
	SteeringPosition            *SteeringPositionValue `json:"steeringposition"`            //https://schema.org/SteeringPositionValue
	TongueWeight                *QuantitativeValue     `json:"tongueweight"`                //https://schema.org/QuantitativeValue
	TrailerWeight               *QuantitativeValue     `json:"trailerweight"`               //https://schema.org/QuantitativeValue
	VehicleConfiguration        *Text                  `json:"vehicleconfiguration"`        //https://schema.org/Text
	VehicleEngine               *EngineSpecification   `json:"vehicleengine"`               //https://schema.org/EngineSpecification
	VehicleIdentificationNumber *Text                  `json:"vehicleidentificationnumber"` //https://schema.org/Text
	VehicleInteriorColor        *Text                  `json:"vehicleinteriorcolor"`        //https://schema.org/Text
	VehicleInteriorType         *Text                  `json:"vehicleinteriortype"`         //https://schema.org/Text
	VehicleModelDate            *Date                  `json:"vehiclemodeldate"`            //https://schema.org/Date
	VehicleSeatingCapacity      *QuantitativeValue     `json:"vehicleseatingcapacity"`      //https://schema.org/QuantitativeValue
	VehicleSpecialUsage         *Text                  `json:"vehiclespecialusage"`         //https://schema.org/Text
	VehicleTransmission         *URL                   `json:"vehicletransmission"`         //https://schema.org/URL
	WeightTotal                 *QuantitativeValue     `json:"weighttotal"`                 //https://schema.org/QuantitativeValue
	Wheelbase                   *QuantitativeValue     `json:"wheelbase"`                   //https://schema.org/QuantitativeValue
}

// MoveAction is a generated struct representing the https://schema.org/MoveAction class.
type MoveAction struct {
	FromLocation *Place `json:"fromlocation"` //https://schema.org/Place
	ToLocation   *Place `json:"tolocation"`   //https://schema.org/Place
}

// MusicReleaseFormatType is a generated struct representing the https://schema.org/MusicReleaseFormatType class.
type MusicReleaseFormatType struct {
}

// OceanBodyOfWater is a generated struct representing the https://schema.org/OceanBodyOfWater class.
type OceanBodyOfWater struct {
}

// ChildCare is a generated struct representing the https://schema.org/ChildCare class.
type ChildCare struct {
}

// MedicalGuideline is a generated struct representing the https://schema.org/MedicalGuideline class.
type MedicalGuideline struct {
	EvidenceLevel    *MedicalEvidenceLevel `json:"evidencelevel"`    //https://schema.org/MedicalEvidenceLevel
	EvidenceOrigin   *Text                 `json:"evidenceorigin"`   //https://schema.org/Text
	GuidelineDate    *Date                 `json:"guidelinedate"`    //https://schema.org/Date
	GuidelineSubject *MedicalEntity        `json:"guidelinesubject"` //https://schema.org/MedicalEntity
}

// Motorcycle is a generated struct representing the https://schema.org/Motorcycle class.
type Motorcycle struct {
}

// Service is a generated struct representing the https://schema.org/Service class.
type Service struct {
	AggregateRating  *AggregateRating           `json:"aggregaterating"`  //https://schema.org/AggregateRating
	AvailableChannel *ServiceChannel            `json:"availablechannel"` //https://schema.org/ServiceChannel
	Brand            *Organization              `json:"brand"`            //https://schema.org/Organization
	HasOfferCatalog  *OfferCatalog              `json:"hasoffercatalog"`  //https://schema.org/OfferCatalog
	HoursAvailable   *OpeningHoursSpecification `json:"hoursavailable"`   //https://schema.org/OpeningHoursSpecification
	IsRelatedTo      *Service                   `json:"isrelatedto"`      //https://schema.org/Service
	IsSimilarTo      *Service                   `json:"issimilarto"`      //https://schema.org/Service
	Logo             *URL                       `json:"logo"`             //https://schema.org/URL
	Produces         *Thing                     `json:"produces"`         //https://schema.org/Thing
	ProviderMobility *Text                      `json:"providermobility"` //https://schema.org/Text
	ServiceAudience  *Audience                  `json:"serviceaudience"`  //https://schema.org/Audience
	ServiceType      *Text                      `json:"servicetype"`      //https://schema.org/Text
	Slogan           *Text                      `json:"slogan"`           //https://schema.org/Text
	TermsOfService   *URL                       `json:"termsofservice"`   //https://schema.org/URL
	Audience         *Audience                  `json:"audience"`         //https://schema.org/Audience
	Award            *Text                      `json:"award"`            //https://schema.org/Text
	Broker           *Person                    `json:"broker"`           //https://schema.org/Person
	Category         *URL                       `json:"category"`         //https://schema.org/URL
	Offers           *Offer                     `json:"offers"`           //https://schema.org/Offer
	Provider         *Person                    `json:"provider"`         //https://schema.org/Person
	Review           *Review                    `json:"review"`           //https://schema.org/Review
	ServiceArea      *Place                     `json:"servicearea"`      //https://schema.org/Place
	ServiceOutput    *Thing                     `json:"serviceoutput"`    //https://schema.org/Thing
	AreaServed       *Text                      `json:"areaserved"`       //https://schema.org/Text
}

// EUEnergyEfficiencyEnumeration is a generated struct representing the https://schema.org/EUEnergyEfficiencyEnumeration class.
type EUEnergyEfficiencyEnumeration struct {
}

// EngineSpecification is a generated struct representing the https://schema.org/EngineSpecification class.
type EngineSpecification struct {
	EngineDisplacement *QuantitativeValue `json:"enginedisplacement"` //https://schema.org/QuantitativeValue
	EnginePower        *QuantitativeValue `json:"enginepower"`        //https://schema.org/QuantitativeValue
	EngineType         *URL               `json:"enginetype"`         //https://schema.org/URL
	FuelType           *URL               `json:"fueltype"`           //https://schema.org/URL
	Torque             *QuantitativeValue `json:"torque"`             //https://schema.org/QuantitativeValue
}

// FAQPage is a generated struct representing the https://schema.org/FAQPage class.
type FAQPage struct {
}

// MerchantReturnEnumeration is a generated struct representing the https://schema.org/MerchantReturnEnumeration class.
type MerchantReturnEnumeration struct {
}

// Hackathon is a generated struct representing the https://schema.org/Hackathon class.
type Hackathon struct {
}

// NutritionInformation is a generated struct representing the https://schema.org/NutritionInformation class.
type NutritionInformation struct {
	Calories              *Energy `json:"calories"`              //https://schema.org/Energy
	CarbohydrateContent   *Mass   `json:"carbohydratecontent"`   //https://schema.org/Mass
	CholesterolContent    *Mass   `json:"cholesterolcontent"`    //https://schema.org/Mass
	FatContent            *Mass   `json:"fatcontent"`            //https://schema.org/Mass
	FiberContent          *Mass   `json:"fibercontent"`          //https://schema.org/Mass
	ProteinContent        *Mass   `json:"proteincontent"`        //https://schema.org/Mass
	SaturatedFatContent   *Mass   `json:"saturatedfatcontent"`   //https://schema.org/Mass
	ServingSize           *Text   `json:"servingsize"`           //https://schema.org/Text
	SodiumContent         *Mass   `json:"sodiumcontent"`         //https://schema.org/Mass
	SugarContent          *Mass   `json:"sugarcontent"`          //https://schema.org/Mass
	TransFatContent       *Mass   `json:"transfatcontent"`       //https://schema.org/Mass
	UnsaturatedFatContent *Mass   `json:"unsaturatedfatcontent"` //https://schema.org/Mass
}

// Report is a generated struct representing the https://schema.org/Report class.
type Report struct {
	ReportNumber *Text `json:"reportnumber"` //https://schema.org/Text
}

// UserPageVisits is a generated struct representing the https://schema.org/UserPageVisits class.
type UserPageVisits struct {
}

// VisualArtwork is a generated struct representing the https://schema.org/VisualArtwork class.
type VisualArtwork struct {
	ArtEdition     *Text              `json:"artedition"`     //https://schema.org/Text
	ArtMedium      *URL               `json:"artmedium"`      //https://schema.org/URL
	Artform        *URL               `json:"artform"`        //https://schema.org/URL
	Artist         *Person            `json:"artist"`         //https://schema.org/Person
	Colorist       *Person            `json:"colorist"`       //https://schema.org/Person
	Depth          *QuantitativeValue `json:"depth"`          //https://schema.org/QuantitativeValue
	Height         *QuantitativeValue `json:"height"`         //https://schema.org/QuantitativeValue
	Inker          *Person            `json:"inker"`          //https://schema.org/Person
	Letterer       *Person            `json:"letterer"`       //https://schema.org/Person
	Penciler       *Person            `json:"penciler"`       //https://schema.org/Person
	Surface        *URL               `json:"surface"`        //https://schema.org/URL
	Width          *QuantitativeValue `json:"width"`          //https://schema.org/QuantitativeValue
	ArtworkSurface *URL               `json:"artworksurface"` //https://schema.org/URL
}

// Pharmacy is a generated struct representing the https://schema.org/Pharmacy class.
type Pharmacy struct {
}

// Residence is a generated struct representing the https://schema.org/Residence class.
type Residence struct {
	AccommodationFloorPlan *FloorPlan `json:"accommodationfloorplan"` //https://schema.org/FloorPlan
}

// FoodEvent is a generated struct representing the https://schema.org/FoodEvent class.
type FoodEvent struct {
}

// Quantity is a generated struct representing the https://schema.org/Quantity class.
type Quantity struct {
}

// WinAction is a generated struct representing the https://schema.org/WinAction class.
type WinAction struct {
	Loser *Person `json:"loser"` //https://schema.org/Person
}

// CreativeWorkSeason is a generated struct representing the https://schema.org/CreativeWorkSeason class.
type CreativeWorkSeason struct {
	EndDate           *DateTime           `json:"enddate"`           //https://schema.org/DateTime
	Episodes          *Episode            `json:"episodes"`          //https://schema.org/Episode
	NumberOfEpisodes  *Integer            `json:"numberofepisodes"`  //https://schema.org/Integer
	ProductionCompany *Organization       `json:"productioncompany"` //https://schema.org/Organization
	SeasonNumber      *Text               `json:"seasonnumber"`      //https://schema.org/Text
	StartDate         *DateTime           `json:"startdate"`         //https://schema.org/DateTime
	Trailer           *VideoObject        `json:"trailer"`           //https://schema.org/VideoObject
	Director          *Person             `json:"director"`          //https://schema.org/Person
	Episode           *Episode            `json:"episode"`           //https://schema.org/Episode
	PartOfSeries      *CreativeWorkSeries `json:"partofseries"`      //https://schema.org/CreativeWorkSeries
	Actor             *Person             `json:"actor"`             //https://schema.org/Person
}

// Demand is a generated struct representing the https://schema.org/Demand class.
type Demand struct {
	AcceptedPaymentMethod     *PaymentMethod       `json:"acceptedpaymentmethod"`     //https://schema.org/PaymentMethod
	AdvanceBookingRequirement *QuantitativeValue   `json:"advancebookingrequirement"` //https://schema.org/QuantitativeValue
	Asin                      *URL                 `json:"asin"`                      //https://schema.org/URL
	Availability              *ItemAvailability    `json:"availability"`              //https://schema.org/ItemAvailability
	AvailabilityEnds          *Time                `json:"availabilityends"`          //https://schema.org/Time
	AvailabilityStarts        *Time                `json:"availabilitystarts"`        //https://schema.org/Time
	AvailableAtOrFrom         *Place               `json:"availableatorfrom"`         //https://schema.org/Place
	AvailableDeliveryMethod   *DeliveryMethod      `json:"availabledeliverymethod"`   //https://schema.org/DeliveryMethod
	BusinessFunction          *BusinessFunction    `json:"businessfunction"`          //https://schema.org/BusinessFunction
	DeliveryLeadTime          *QuantitativeValue   `json:"deliveryleadtime"`          //https://schema.org/QuantitativeValue
	EligibleCustomerType      *BusinessEntityType  `json:"eligiblecustomertype"`      //https://schema.org/BusinessEntityType
	EligibleDuration          *QuantitativeValue   `json:"eligibleduration"`          //https://schema.org/QuantitativeValue
	EligibleQuantity          *QuantitativeValue   `json:"eligiblequantity"`          //https://schema.org/QuantitativeValue
	EligibleRegion            *Text                `json:"eligibleregion"`            //https://schema.org/Text
	EligibleTransactionVolume *PriceSpecification  `json:"eligibletransactionvolume"` //https://schema.org/PriceSpecification
	Gtin12                    *Text                `json:"gtin12"`                    //https://schema.org/Text
	Gtin13                    *Text                `json:"gtin13"`                    //https://schema.org/Text
	Gtin14                    *Text                `json:"gtin14"`                    //https://schema.org/Text
	Gtin8                     *Text                `json:"gtin8"`                     //https://schema.org/Text
	IncludesObject            *TypeAndQuantityNode `json:"includesobject"`            //https://schema.org/TypeAndQuantityNode
	IneligibleRegion          *Text                `json:"ineligibleregion"`          //https://schema.org/Text
	InventoryLevel            *QuantitativeValue   `json:"inventorylevel"`            //https://schema.org/QuantitativeValue
	ItemCondition             *OfferItemCondition  `json:"itemcondition"`             //https://schema.org/OfferItemCondition
	Mpn                       *Text                `json:"mpn"`                       //https://schema.org/Text
	PriceSpecification        *PriceSpecification  `json:"pricespecification"`        //https://schema.org/PriceSpecification
	Sku                       *Text                `json:"sku"`                       //https://schema.org/Text
	ValidFrom                 *DateTime            `json:"validfrom"`                 //https://schema.org/DateTime
	ValidThrough              *DateTime            `json:"validthrough"`              //https://schema.org/DateTime
	ItemOffered               *Trip                `json:"itemoffered"`               //https://schema.org/Trip
	SerialNumber              *Text                `json:"serialnumber"`              //https://schema.org/Text
	Warranty                  *WarrantyPromise     `json:"warranty"`                  //https://schema.org/WarrantyPromise
	Seller                    *Person              `json:"seller"`                    //https://schema.org/Person
	AreaServed                *Text                `json:"areaserved"`                //https://schema.org/Text
	Gtin                      *URL                 `json:"gtin"`                      //https://schema.org/URL
}

// EndorseAction is a generated struct representing the https://schema.org/EndorseAction class.
type EndorseAction struct {
	Endorsee *Person `json:"endorsee"` //https://schema.org/Person
}

// HealthPlanFormulary is a generated struct representing the https://schema.org/HealthPlanFormulary class.
type HealthPlanFormulary struct {
	HealthPlanCostSharing    *Boolean `json:"healthplancostsharing"`    //https://schema.org/Boolean
	HealthPlanDrugTier       *Text    `json:"healthplandrugtier"`       //https://schema.org/Text
	OffersPrescriptionByMail *Boolean `json:"offersprescriptionbymail"` //https://schema.org/Boolean
}

// Manuscript is a generated struct representing the https://schema.org/Manuscript class.
type Manuscript struct {
}

// SeekToAction is a generated struct representing the https://schema.org/SeekToAction class.
type SeekToAction struct {
	StartOffset *Number `json:"startoffset"` //https://schema.org/Number
}

// BoatTrip is a generated struct representing the https://schema.org/BoatTrip class.
type BoatTrip struct {
	ArrivalBoatTerminal   *BoatTerminal `json:"arrivalboatterminal"`   //https://schema.org/BoatTerminal
	DepartureBoatTerminal *BoatTerminal `json:"departureboatterminal"` //https://schema.org/BoatTerminal
}

// DigitalDocumentPermission is a generated struct representing the https://schema.org/DigitalDocumentPermission class.
type DigitalDocumentPermission struct {
	Grantee        *Person                        `json:"grantee"`        //https://schema.org/Person
	PermissionType *DigitalDocumentPermissionType `json:"permissiontype"` //https://schema.org/DigitalDocumentPermissionType
}

// Room is a generated struct representing the https://schema.org/Room class.
type Room struct {
}

// SearchResultsPage is a generated struct representing the https://schema.org/SearchResultsPage class.
type SearchResultsPage struct {
}

// LendAction is a generated struct representing the https://schema.org/LendAction class.
type LendAction struct {
	Borrower *Person `json:"borrower"` //https://schema.org/Person
}

// OwnershipInfo is a generated struct representing the https://schema.org/OwnershipInfo class.
type OwnershipInfo struct {
	AcquiredFrom *Person   `json:"acquiredfrom"` //https://schema.org/Person
	OwnedFrom    *DateTime `json:"ownedfrom"`    //https://schema.org/DateTime
	OwnedThrough *DateTime `json:"ownedthrough"` //https://schema.org/DateTime
	TypeOfGood   *Service  `json:"typeofgood"`   //https://schema.org/Service
}

// TouristAttraction is a generated struct representing the https://schema.org/TouristAttraction class.
type TouristAttraction struct {
	AvailableLanguage *Text `json:"availablelanguage"` //https://schema.org/Text
	TouristType       *Text `json:"touristtype"`       //https://schema.org/Text
}

// ComicCoverArt is a generated struct representing the https://schema.org/ComicCoverArt class.
type ComicCoverArt struct {
}

// Language is a generated struct representing the https://schema.org/Language class.
type Language struct {
}

// MediaGallery is a generated struct representing the https://schema.org/MediaGallery class.
type MediaGallery struct {
}

// StatisticalVariable is a generated struct representing the https://schema.org/StatisticalVariable class.
type StatisticalVariable struct {
	MeasuredProperty       *Property            `json:"measuredproperty"`       //https://schema.org/Property
	MeasurementDenominator *StatisticalVariable `json:"measurementdenominator"` //https://schema.org/StatisticalVariable
	MeasurementMethod      *URL                 `json:"measurementmethod"`      //https://schema.org/URL
	MeasurementQualifier   *Enumeration         `json:"measurementqualifier"`   //https://schema.org/Enumeration
	PopulationType         *Class               `json:"populationtype"`         //https://schema.org/Class
	StatType               *URL                 `json:"stattype"`               //https://schema.org/URL
	MeasurementTechnique   *URL                 `json:"measurementtechnique"`   //https://schema.org/URL
}

// MedicalConditionStage is a generated struct representing the https://schema.org/MedicalConditionStage class.
type MedicalConditionStage struct {
	StageAsNumber  *Number `json:"stageasnumber"`  //https://schema.org/Number
	SubStageSuffix *Text   `json:"substagesuffix"` //https://schema.org/Text
}

// PhotographAction is a generated struct representing the https://schema.org/PhotographAction class.
type PhotographAction struct {
}

// VacationRental is a generated struct representing the https://schema.org/VacationRental class.
type VacationRental struct {
}

// Beach is a generated struct representing the https://schema.org/Beach class.
type Beach struct {
}

// BusStation is a generated struct representing the https://schema.org/BusStation class.
type BusStation struct {
}

// MolecularEntity is a generated struct representing the https://schema.org/MolecularEntity class.
type MolecularEntity struct {
	ChemicalRole                *DefinedTerm `json:"chemicalrole"`                //https://schema.org/DefinedTerm
	InChI                       *Text        `json:"inchi"`                       //https://schema.org/Text
	InChIKey                    *Text        `json:"inchikey"`                    //https://schema.org/Text
	IupacName                   *Text        `json:"iupacname"`                   //https://schema.org/Text
	MolecularFormula            *Text        `json:"molecularformula"`            //https://schema.org/Text
	MolecularWeight             *Text        `json:"molecularweight"`             //https://schema.org/Text
	MonoisotopicMolecularWeight *Text        `json:"monoisotopicmolecularweight"` //https://schema.org/Text
	PotentialUse                *DefinedTerm `json:"potentialuse"`                //https://schema.org/DefinedTerm
	Smiles                      *Text        `json:"smiles"`                      //https://schema.org/Text
}

// Property is a generated struct representing the https://schema.org/Property class.
type Property struct {
	DomainIncludes *Class    `json:"domainincludes"` //https://schema.org/Class
	InverseOf      *Property `json:"inverseof"`      //https://schema.org/Property
	RangeIncludes  *Class    `json:"rangeincludes"`  //https://schema.org/Class
	SupersededBy   *Property `json:"supersededby"`   //https://schema.org/Property
}

// Blog is a generated struct representing the https://schema.org/Blog class.
type Blog struct {
	BlogPosts *BlogPosting `json:"blogposts"` //https://schema.org/BlogPosting
	Issn      *Text        `json:"issn"`      //https://schema.org/Text
	BlogPost  *BlogPosting `json:"blogpost"`  //https://schema.org/BlogPosting
}

// Hotel is a generated struct representing the https://schema.org/Hotel class.
type Hotel struct {
}

// Landform is a generated struct representing the https://schema.org/Landform class.
type Landform struct {
}

// LiquorStore is a generated struct representing the https://schema.org/LiquorStore class.
type LiquorStore struct {
}

// MovieSeries is a generated struct representing the https://schema.org/MovieSeries class.
type MovieSeries struct {
	Actors            *Person       `json:"actors"`            //https://schema.org/Person
	Directors         *Person       `json:"directors"`         //https://schema.org/Person
	MusicBy           *Person       `json:"musicby"`           //https://schema.org/Person
	ProductionCompany *Organization `json:"productioncompany"` //https://schema.org/Organization
	Trailer           *VideoObject  `json:"trailer"`           //https://schema.org/VideoObject
	Director          *Person       `json:"director"`          //https://schema.org/Person
	Actor             *Person       `json:"actor"`             //https://schema.org/Person
}

// PronounceableText is a generated struct representing the https://schema.org/PronounceableText class.
type PronounceableText struct {
	PhoneticText       *Text `json:"phonetictext"`       //https://schema.org/Text
	SpeechToTextMarkup *Text `json:"speechtotextmarkup"` //https://schema.org/Text
	TextValue          *Text `json:"textvalue"`          //https://schema.org/Text
	InLanguage         *Text `json:"inlanguage"`         //https://schema.org/Text
}

// ArriveAction is a generated struct representing the https://schema.org/ArriveAction class.
type ArriveAction struct {
}

// GeneralContractor is a generated struct representing the https://schema.org/GeneralContractor class.
type GeneralContractor struct {
}

// RoofingContractor is a generated struct representing the https://schema.org/RoofingContractor class.
type RoofingContractor struct {
}

// DrugCostCategory is a generated struct representing the https://schema.org/DrugCostCategory class.
type DrugCostCategory struct {
}

// GamePlayMode is a generated struct representing the https://schema.org/GamePlayMode class.
type GamePlayMode struct {
}

// Observation is a generated struct representing the https://schema.org/Observation class.
type Observation struct {
	MarginOfError          *QuantitativeValue   `json:"marginoferror"`          //https://schema.org/QuantitativeValue
	MeasuredProperty       *Property            `json:"measuredproperty"`       //https://schema.org/Property
	MeasurementDenominator *StatisticalVariable `json:"measurementdenominator"` //https://schema.org/StatisticalVariable
	MeasurementMethod      *URL                 `json:"measurementmethod"`      //https://schema.org/URL
	MeasurementQualifier   *Enumeration         `json:"measurementqualifier"`   //https://schema.org/Enumeration
	ObservationAbout       *Thing               `json:"observationabout"`       //https://schema.org/Thing
	ObservationDate        *DateTime            `json:"observationdate"`        //https://schema.org/DateTime
	ObservationPeriod      *Text                `json:"observationperiod"`      //https://schema.org/Text
	VariableMeasured       *Text                `json:"variablemeasured"`       //https://schema.org/Text
	MeasurementTechnique   *URL                 `json:"measurementtechnique"`   //https://schema.org/URL
}

// Bridge is a generated struct representing the https://schema.org/Bridge class.
type Bridge struct {
}

// DepartmentStore is a generated struct representing the https://schema.org/DepartmentStore class.
type DepartmentStore struct {
}

// GovernmentPermit is a generated struct representing the https://schema.org/GovernmentPermit class.
type GovernmentPermit struct {
}

// Question is a generated struct representing the https://schema.org/Question class.
type Question struct {
	AcceptedAnswer  *ItemList     `json:"acceptedanswer"`  //https://schema.org/ItemList
	AnswerCount     *Integer      `json:"answercount"`     //https://schema.org/Integer
	EduQuestionType *Text         `json:"eduquestiontype"` //https://schema.org/Text
	ParentItem      *CreativeWork `json:"parentitem"`      //https://schema.org/CreativeWork
	SuggestedAnswer *ItemList     `json:"suggestedanswer"` //https://schema.org/ItemList
}

// ShoeStore is a generated struct representing the https://schema.org/ShoeStore class.
type ShoeStore struct {
}

// CivicStructure is a generated struct representing the https://schema.org/CivicStructure class.
type CivicStructure struct {
	OpeningHours *Text `json:"openinghours"` //https://schema.org/Text
}

// PoliceStation is a generated struct representing the https://schema.org/PoliceStation class.
type PoliceStation struct {
}

// SoftwareSourceCode is a generated struct representing the https://schema.org/SoftwareSourceCode class.
type SoftwareSourceCode struct {
	CodeRepository      *URL                 `json:"coderepository"`      //https://schema.org/URL
	ProgrammingLanguage *Text                `json:"programminglanguage"` //https://schema.org/Text
	Runtime             *Text                `json:"runtime"`             //https://schema.org/Text
	SampleType          *Text                `json:"sampletype"`          //https://schema.org/Text
	TargetProduct       *SoftwareApplication `json:"targetproduct"`       //https://schema.org/SoftwareApplication
	CodeSampleType      *Text                `json:"codesampletype"`      //https://schema.org/Text
	RuntimePlatform     *Text                `json:"runtimeplatform"`     //https://schema.org/Text
}

// CDCPMDRecord is a generated struct representing the https://schema.org/CDCPMDRecord class.
type CDCPMDRecord struct {
	CvdCollectionDate       *Text     `json:"cvdcollectiondate"`       //https://schema.org/Text
	CvdFacilityCounty       *Text     `json:"cvdfacilitycounty"`       //https://schema.org/Text
	CvdFacilityId           *Text     `json:"cvdfacilityid"`           //https://schema.org/Text
	CvdNumBeds              *Number   `json:"cvdnumbeds"`              //https://schema.org/Number
	CvdNumBedsOcc           *Number   `json:"cvdnumbedsocc"`           //https://schema.org/Number
	CvdNumC19Died           *Number   `json:"cvdnumc19died"`           //https://schema.org/Number
	CvdNumC19HOPats         *Number   `json:"cvdnumc19hopats"`         //https://schema.org/Number
	CvdNumC19HospPats       *Number   `json:"cvdnumc19hosppats"`       //https://schema.org/Number
	CvdNumC19MechVentPats   *Number   `json:"cvdnumc19mechventpats"`   //https://schema.org/Number
	CvdNumC19OFMechVentPats *Number   `json:"cvdnumc19ofmechventpats"` //https://schema.org/Number
	CvdNumC19OverflowPats   *Number   `json:"cvdnumc19overflowpats"`   //https://schema.org/Number
	CvdNumICUBeds           *Number   `json:"cvdnumicubeds"`           //https://schema.org/Number
	CvdNumICUBedsOcc        *Number   `json:"cvdnumicubedsocc"`        //https://schema.org/Number
	CvdNumTotBeds           *Number   `json:"cvdnumtotbeds"`           //https://schema.org/Number
	CvdNumVent              *Number   `json:"cvdnumvent"`              //https://schema.org/Number
	CvdNumVentUse           *Number   `json:"cvdnumventuse"`           //https://schema.org/Number
	DatePosted              *DateTime `json:"dateposted"`              //https://schema.org/DateTime
}

// CovidTestingFacility is a generated struct representing the https://schema.org/CovidTestingFacility class.
type CovidTestingFacility struct {
}

// GeoCoordinates is a generated struct representing the https://schema.org/GeoCoordinates class.
type GeoCoordinates struct {
	Address        *Text `json:"address"`        //https://schema.org/Text
	AddressCountry *Text `json:"addresscountry"` //https://schema.org/Text
	Elevation      *Text `json:"elevation"`      //https://schema.org/Text
	Latitude       *Text `json:"latitude"`       //https://schema.org/Text
	Longitude      *Text `json:"longitude"`      //https://schema.org/Text
	PostalCode     *Text `json:"postalcode"`     //https://schema.org/Text
}

// MedicalIntangible is a generated struct representing the https://schema.org/MedicalIntangible class.
type MedicalIntangible struct {
}

// ReserveAction is a generated struct representing the https://schema.org/ReserveAction class.
type ReserveAction struct {
}

// WarrantyScope is a generated struct representing the https://schema.org/WarrantyScope class.
type WarrantyScope struct {
}

// Church is a generated struct representing the https://schema.org/Church class.
type Church struct {
}

// GeoCircle is a generated struct representing the https://schema.org/GeoCircle class.
type GeoCircle struct {
	GeoMidpoint *GeoCoordinates `json:"geomidpoint"` //https://schema.org/GeoCoordinates
	GeoRadius   *Text           `json:"georadius"`   //https://schema.org/Text
}

// Mosque is a generated struct representing the https://schema.org/Mosque class.
type Mosque struct {
}

// Photograph is a generated struct representing the https://schema.org/Photograph class.
type Photograph struct {
}

// Resort is a generated struct representing the https://schema.org/Resort class.
type Resort struct {
}

// WearableMeasurementTypeEnumeration is a generated struct representing the https://schema.org/WearableMeasurementTypeEnumeration class.
type WearableMeasurementTypeEnumeration struct {
}

// ItemList is a generated struct representing the https://schema.org/ItemList class.
type ItemList struct {
	ItemListElement *Thing   `json:"itemlistelement"` //https://schema.org/Thing
	ItemListOrder   *Text    `json:"itemlistorder"`   //https://schema.org/Text
	NumberOfItems   *Integer `json:"numberofitems"`   //https://schema.org/Integer
}

// Series is a generated struct representing the https://schema.org/Series class.
type Series struct {
}

// DisagreeAction is a generated struct representing the https://schema.org/DisagreeAction class.
type DisagreeAction struct {
}

// FoodEstablishment is a generated struct representing the https://schema.org/FoodEstablishment class.
type FoodEstablishment struct {
	AcceptsReservations *URL    `json:"acceptsreservations"` //https://schema.org/URL
	Menu                *URL    `json:"menu"`                //https://schema.org/URL
	ServesCuisine       *Text   `json:"servescuisine"`       //https://schema.org/Text
	StarRating          *Rating `json:"starrating"`          //https://schema.org/Rating
	HasMenu             *URL    `json:"hasmenu"`             //https://schema.org/URL
}

// House is a generated struct representing the https://schema.org/House class.
type House struct {
	NumberOfRooms *QuantitativeValue `json:"numberofrooms"` //https://schema.org/QuantitativeValue
}

// MedicalObservationalStudy is a generated struct representing the https://schema.org/MedicalObservationalStudy class.
type MedicalObservationalStudy struct {
	StudyDesign *MedicalObservationalStudyDesign `json:"studydesign"` //https://schema.org/MedicalObservationalStudyDesign
}

// Quotation is a generated struct representing the https://schema.org/Quotation class.
type Quotation struct {
	SpokenByCharacter *Person `json:"spokenbycharacter"` //https://schema.org/Person
}

// ComicStory is a generated struct representing the https://schema.org/ComicStory class.
type ComicStory struct {
	Artist   *Person `json:"artist"`   //https://schema.org/Person
	Colorist *Person `json:"colorist"` //https://schema.org/Person
	Inker    *Person `json:"inker"`    //https://schema.org/Person
	Letterer *Person `json:"letterer"` //https://schema.org/Person
	Penciler *Person `json:"penciler"` //https://schema.org/Person
}

// SizeGroupEnumeration is a generated struct representing the https://schema.org/SizeGroupEnumeration class.
type SizeGroupEnumeration struct {
}

// Attorney is a generated struct representing the https://schema.org/Attorney class.
type Attorney struct {
}

// Embassy is a generated struct representing the https://schema.org/Embassy class.
type Embassy struct {
}

// MedicalTrialDesign is a generated struct representing the https://schema.org/MedicalTrialDesign class.
type MedicalTrialDesign struct {
}

// Statement is a generated struct representing the https://schema.org/Statement class.
type Statement struct {
}

// TieAction is a generated struct representing the https://schema.org/TieAction class.
type TieAction struct {
}

// ContactPoint is a generated struct representing the https://schema.org/ContactPoint class.
type ContactPoint struct {
	AvailableLanguage *Text                      `json:"availablelanguage"` //https://schema.org/Text
	ContactOption     *ContactPointOption        `json:"contactoption"`     //https://schema.org/ContactPointOption
	ContactType       *Text                      `json:"contacttype"`       //https://schema.org/Text
	Email             *Text                      `json:"email"`             //https://schema.org/Text
	FaxNumber         *Text                      `json:"faxnumber"`         //https://schema.org/Text
	HoursAvailable    *OpeningHoursSpecification `json:"hoursavailable"`    //https://schema.org/OpeningHoursSpecification
	ProductSupported  *Text                      `json:"productsupported"`  //https://schema.org/Text
	Telephone         *Text                      `json:"telephone"`         //https://schema.org/Text
	ServiceArea       *Place                     `json:"servicearea"`       //https://schema.org/Place
	AreaServed        *Text                      `json:"areaserved"`        //https://schema.org/Text
}

// Rating is a generated struct representing the https://schema.org/Rating class.
type Rating struct {
	Author            *Person `json:"author"`            //https://schema.org/Person
	BestRating        *Text   `json:"bestrating"`        //https://schema.org/Text
	RatingExplanation *Text   `json:"ratingexplanation"` //https://schema.org/Text
	RatingValue       *Text   `json:"ratingvalue"`       //https://schema.org/Text
	ReviewAspect      *Text   `json:"reviewaspect"`      //https://schema.org/Text
	WorstRating       *Text   `json:"worstrating"`       //https://schema.org/Text
}

// TouristInformationCenter is a generated struct representing the https://schema.org/TouristInformationCenter class.
type TouristInformationCenter struct {
}

// VideoGame is a generated struct representing the https://schema.org/VideoGame class.
type VideoGame struct {
	Actors       *Person       `json:"actors"`       //https://schema.org/Person
	CheatCode    *CreativeWork `json:"cheatcode"`    //https://schema.org/CreativeWork
	Directors    *Person       `json:"directors"`    //https://schema.org/Person
	GameEdition  *Text         `json:"gameedition"`  //https://schema.org/Text
	GamePlatform *URL          `json:"gameplatform"` //https://schema.org/URL
	GameTip      *CreativeWork `json:"gametip"`      //https://schema.org/CreativeWork
	MusicBy      *Person       `json:"musicby"`      //https://schema.org/Person
	PlayMode     *GamePlayMode `json:"playmode"`     //https://schema.org/GamePlayMode
	Trailer      *VideoObject  `json:"trailer"`      //https://schema.org/VideoObject
	Director     *Person       `json:"director"`     //https://schema.org/Person
	GameServer   *GameServer   `json:"gameserver"`   //https://schema.org/GameServer
	Actor        *Person       `json:"actor"`        //https://schema.org/Person
}

// Action is a generated struct representing the https://schema.org/Action class.
type Action struct {
	ActionStatus *ActionStatusType `json:"actionstatus"` //https://schema.org/ActionStatusType
	Agent        *Person           `json:"agent"`        //https://schema.org/Person
	EndTime      *Time             `json:"endtime"`      //https://schema.org/Time
	Error        *Thing            `json:"error"`        //https://schema.org/Thing
	StartTime    *Time             `json:"starttime"`    //https://schema.org/Time
	Target       *URL              `json:"target"`       //https://schema.org/URL
	Provider     *Person           `json:"provider"`     //https://schema.org/Person
	Result       *Thing            `json:"result"`       //https://schema.org/Thing
	Instrument   *Thing            `json:"instrument"`   //https://schema.org/Thing
	Object       *Thing            `json:"object"`       //https://schema.org/Thing
	Location     *VirtualLocation  `json:"location"`     //https://schema.org/VirtualLocation
	Participant  *Person           `json:"participant"`  //https://schema.org/Person
}

// DigitalPlatformEnumeration is a generated struct representing the https://schema.org/DigitalPlatformEnumeration class.
type DigitalPlatformEnumeration struct {
}

// RepaymentSpecification is a generated struct representing the https://schema.org/RepaymentSpecification class.
type RepaymentSpecification struct {
	DownPayment            *Number         `json:"downpayment"`            //https://schema.org/Number
	EarlyPrepaymentPenalty *MonetaryAmount `json:"earlyprepaymentpenalty"` //https://schema.org/MonetaryAmount
	LoanPaymentAmount      *MonetaryAmount `json:"loanpaymentamount"`      //https://schema.org/MonetaryAmount
	LoanPaymentFrequency   *Number         `json:"loanpaymentfrequency"`   //https://schema.org/Number
	NumberOfLoanPayments   *Number         `json:"numberofloanpayments"`   //https://schema.org/Number
}

// RsvpAction is a generated struct representing the https://schema.org/RsvpAction class.
type RsvpAction struct {
	AdditionalNumberOfGuests *Number           `json:"additionalnumberofguests"` //https://schema.org/Number
	Comment                  *Comment          `json:"comment"`                  //https://schema.org/Comment
	RsvpResponse             *RsvpResponseType `json:"rsvpresponse"`             //https://schema.org/RsvpResponseType
}

// InvestmentFund is a generated struct representing the https://schema.org/InvestmentFund class.
type InvestmentFund struct {
}

// PreOrderAction is a generated struct representing the https://schema.org/PreOrderAction class.
type PreOrderAction struct {
}

// FastFoodRestaurant is a generated struct representing the https://schema.org/FastFoodRestaurant class.
type FastFoodRestaurant struct {
}

// InteractionCounter is a generated struct representing the https://schema.org/InteractionCounter class.
type InteractionCounter struct {
	EndTime              *Time            `json:"endtime"`              //https://schema.org/Time
	InteractionService   *WebSite         `json:"interactionservice"`   //https://schema.org/WebSite
	InteractionType      *Action          `json:"interactiontype"`      //https://schema.org/Action
	StartTime            *Time            `json:"starttime"`            //https://schema.org/Time
	UserInteractionCount *Integer         `json:"userinteractioncount"` //https://schema.org/Integer
	Location             *VirtualLocation `json:"location"`             //https://schema.org/VirtualLocation
}

// ElementarySchool is a generated struct representing the https://schema.org/ElementarySchool class.
type ElementarySchool struct {
}

// SteeringPositionValue is a generated struct representing the https://schema.org/SteeringPositionValue class.
type SteeringPositionValue struct {
}

// CompleteDataFeed is a generated struct representing the https://schema.org/CompleteDataFeed class.
type CompleteDataFeed struct {
}

// GeoShape is a generated struct representing the https://schema.org/GeoShape class.
type GeoShape struct {
	Address        *Text `json:"address"`        //https://schema.org/Text
	AddressCountry *Text `json:"addresscountry"` //https://schema.org/Text
	Box            *Text `json:"box"`            //https://schema.org/Text
	Circle         *Text `json:"circle"`         //https://schema.org/Text
	Elevation      *Text `json:"elevation"`      //https://schema.org/Text
	Line           *Text `json:"line"`           //https://schema.org/Text
	Polygon        *Text `json:"polygon"`        //https://schema.org/Text
	PostalCode     *Text `json:"postalcode"`     //https://schema.org/Text
}

// LifestyleModification is a generated struct representing the https://schema.org/LifestyleModification class.
type LifestyleModification struct {
}

// TVClip is a generated struct representing the https://schema.org/TVClip class.
type TVClip struct {
	PartOfTVSeries *TVSeries `json:"partoftvseries"` //https://schema.org/TVSeries
}

// AutomatedTeller is a generated struct representing the https://schema.org/AutomatedTeller class.
type AutomatedTeller struct {
}

// MedicalBusiness is a generated struct representing the https://schema.org/MedicalBusiness class.
type MedicalBusiness struct {
}

// NewsMediaOrganization is a generated struct representing the https://schema.org/NewsMediaOrganization class.
type NewsMediaOrganization struct {
	ActionableFeedbackPolicy        *URL `json:"actionablefeedbackpolicy"`        //https://schema.org/URL
	CorrectionsPolicy               *URL `json:"correctionspolicy"`               //https://schema.org/URL
	DiversityPolicy                 *URL `json:"diversitypolicy"`                 //https://schema.org/URL
	DiversityStaffingReport         *URL `json:"diversitystaffingreport"`         //https://schema.org/URL
	EthicsPolicy                    *URL `json:"ethicspolicy"`                    //https://schema.org/URL
	Masthead                        *URL `json:"masthead"`                        //https://schema.org/URL
	MissionCoveragePrioritiesPolicy *URL `json:"missioncoverageprioritiespolicy"` //https://schema.org/URL
	NoBylinesPolicy                 *URL `json:"nobylinespolicy"`                 //https://schema.org/URL
	OwnershipFundingInfo            *URL `json:"ownershipfundinginfo"`            //https://schema.org/URL
	UnnamedSourcesPolicy            *URL `json:"unnamedsourcespolicy"`            //https://schema.org/URL
	VerificationFactCheckingPolicy  *URL `json:"verificationfactcheckingpolicy"`  //https://schema.org/URL
}

// Poster is a generated struct representing the https://schema.org/Poster class.
type Poster struct {
}

// Crematorium is a generated struct representing the https://schema.org/Crematorium class.
type Crematorium struct {
}

// OfferForPurchase is a generated struct representing the https://schema.org/OfferForPurchase class.
type OfferForPurchase struct {
}

// PaymentCard is a generated struct representing the https://schema.org/PaymentCard class.
type PaymentCard struct {
	CashBack                      *Number         `json:"cashback"`                      //https://schema.org/Number
	ContactlessPayment            *Boolean        `json:"contactlesspayment"`            //https://schema.org/Boolean
	FloorLimit                    *MonetaryAmount `json:"floorlimit"`                    //https://schema.org/MonetaryAmount
	MonthlyMinimumRepaymentAmount *Number         `json:"monthlyminimumrepaymentamount"` //https://schema.org/Number
}

// PhysicalActivityCategory is a generated struct representing the https://schema.org/PhysicalActivityCategory class.
type PhysicalActivityCategory struct {
}

// OfferShippingDetails is a generated struct representing the https://schema.org/OfferShippingDetails class.
type OfferShippingDetails struct {
	DeliveryTime         *ShippingDeliveryTime `json:"deliverytime"`         //https://schema.org/ShippingDeliveryTime
	Depth                *QuantitativeValue    `json:"depth"`                //https://schema.org/QuantitativeValue
	DoesNotShip          *Boolean              `json:"doesnotship"`          //https://schema.org/Boolean
	Height               *QuantitativeValue    `json:"height"`               //https://schema.org/QuantitativeValue
	ShippingDestination  *DefinedRegion        `json:"shippingdestination"`  //https://schema.org/DefinedRegion
	ShippingLabel        *Text                 `json:"shippinglabel"`        //https://schema.org/Text
	ShippingOrigin       *DefinedRegion        `json:"shippingorigin"`       //https://schema.org/DefinedRegion
	ShippingRate         *MonetaryAmount       `json:"shippingrate"`         //https://schema.org/MonetaryAmount
	ShippingSettingsLink *URL                  `json:"shippingsettingslink"` //https://schema.org/URL
	TransitTimeLabel     *Text                 `json:"transittimelabel"`     //https://schema.org/Text
	Weight               *QuantitativeValue    `json:"weight"`               //https://schema.org/QuantitativeValue
	Width                *QuantitativeValue    `json:"width"`                //https://schema.org/QuantitativeValue
}

// WebAPI is a generated struct representing the https://schema.org/WebAPI class.
type WebAPI struct {
	Documentation *URL `json:"documentation"` //https://schema.org/URL
}

// MedicineSystem is a generated struct representing the https://schema.org/MedicineSystem class.
type MedicineSystem struct {
}

// RsvpResponseType is a generated struct representing the https://schema.org/RsvpResponseType class.
type RsvpResponseType struct {
}

// ScholarlyArticle is a generated struct representing the https://schema.org/ScholarlyArticle class.
type ScholarlyArticle struct {
}

// Consortium is a generated struct representing the https://schema.org/Consortium class.
type Consortium struct {
}

// EntryPoint is a generated struct representing the https://schema.org/EntryPoint class.
type EntryPoint struct {
	ActionPlatform    *URL                 `json:"actionplatform"`    //https://schema.org/URL
	Application       *SoftwareApplication `json:"application"`       //https://schema.org/SoftwareApplication
	ContentType       *Text                `json:"contenttype"`       //https://schema.org/Text
	EncodingType      *Text                `json:"encodingtype"`      //https://schema.org/Text
	HttpMethod        *Text                `json:"httpmethod"`        //https://schema.org/Text
	UrlTemplate       *Text                `json:"urltemplate"`       //https://schema.org/Text
	ActionApplication *SoftwareApplication `json:"actionapplication"` //https://schema.org/SoftwareApplication
}

// FurnitureStore is a generated struct representing the https://schema.org/FurnitureStore class.
type FurnitureStore struct {
}

// LegislativeBuilding is a generated struct representing the https://schema.org/LegislativeBuilding class.
type LegislativeBuilding struct {
}

// MusicAlbumReleaseType is a generated struct representing the https://schema.org/MusicAlbumReleaseType class.
type MusicAlbumReleaseType struct {
}

// ProfessionalService is a generated struct representing the https://schema.org/ProfessionalService class.
type ProfessionalService struct {
}

// VideoGameSeries is a generated struct representing the https://schema.org/VideoGameSeries class.
type VideoGameSeries struct {
	Actors             *Person             `json:"actors"`             //https://schema.org/Person
	CharacterAttribute *Thing              `json:"characterattribute"` //https://schema.org/Thing
	CheatCode          *CreativeWork       `json:"cheatcode"`          //https://schema.org/CreativeWork
	Directors          *Person             `json:"directors"`          //https://schema.org/Person
	Episodes           *Episode            `json:"episodes"`           //https://schema.org/Episode
	GameItem           *Thing              `json:"gameitem"`           //https://schema.org/Thing
	GameLocation       *URL                `json:"gamelocation"`       //https://schema.org/URL
	GamePlatform       *URL                `json:"gameplatform"`       //https://schema.org/URL
	MusicBy            *Person             `json:"musicby"`            //https://schema.org/Person
	NumberOfEpisodes   *Integer            `json:"numberofepisodes"`   //https://schema.org/Integer
	NumberOfPlayers    *QuantitativeValue  `json:"numberofplayers"`    //https://schema.org/QuantitativeValue
	NumberOfSeasons    *Integer            `json:"numberofseasons"`    //https://schema.org/Integer
	PlayMode           *GamePlayMode       `json:"playmode"`           //https://schema.org/GamePlayMode
	ProductionCompany  *Organization       `json:"productioncompany"`  //https://schema.org/Organization
	Quest              *Thing              `json:"quest"`              //https://schema.org/Thing
	Seasons            *CreativeWorkSeason `json:"seasons"`            //https://schema.org/CreativeWorkSeason
	Trailer            *VideoObject        `json:"trailer"`            //https://schema.org/VideoObject
	ContainsSeason     *CreativeWorkSeason `json:"containsseason"`     //https://schema.org/CreativeWorkSeason
	Director           *Person             `json:"director"`           //https://schema.org/Person
	Episode            *Episode            `json:"episode"`            //https://schema.org/Episode
	Season             *URL                `json:"season"`             //https://schema.org/URL
	Actor              *Person             `json:"actor"`              //https://schema.org/Person
}

// WPAdBlock is a generated struct representing the https://schema.org/WPAdBlock class.
type WPAdBlock struct {
}

// PaymentStatusType is a generated struct representing the https://schema.org/PaymentStatusType class.
type PaymentStatusType struct {
}

// ActionAccessSpecification is a generated struct representing the https://schema.org/ActionAccessSpecification class.
type ActionAccessSpecification struct {
	AvailabilityEnds     *Time              `json:"availabilityends"`     //https://schema.org/Time
	AvailabilityStarts   *Time              `json:"availabilitystarts"`   //https://schema.org/Time
	EligibleRegion       *Text              `json:"eligibleregion"`       //https://schema.org/Text
	ExpectsAcceptanceOf  *Offer             `json:"expectsacceptanceof"`  //https://schema.org/Offer
	IneligibleRegion     *Text              `json:"ineligibleregion"`     //https://schema.org/Text
	RequiresSubscription *MediaSubscription `json:"requiressubscription"` //https://schema.org/MediaSubscription
	Category             *URL               `json:"category"`             //https://schema.org/URL
}

// FollowAction is a generated struct representing the https://schema.org/FollowAction class.
type FollowAction struct {
	Followee *Person `json:"followee"` //https://schema.org/Person
}

// MusicRecording is a generated struct representing the https://schema.org/MusicRecording class.
type MusicRecording struct {
	ByArtist    *Person           `json:"byartist"`    //https://schema.org/Person
	InAlbum     *MusicAlbum       `json:"inalbum"`     //https://schema.org/MusicAlbum
	InPlaylist  *MusicPlaylist    `json:"inplaylist"`  //https://schema.org/MusicPlaylist
	IsrcCode    *Text             `json:"isrccode"`    //https://schema.org/Text
	Duration    *Duration         `json:"duration"`    //https://schema.org/Duration
	RecordingOf *MusicComposition `json:"recordingof"` //https://schema.org/MusicComposition
}

// SchoolDistrict is a generated struct representing the https://schema.org/SchoolDistrict class.
type SchoolDistrict struct {
}

// UserInteraction is a generated struct representing the https://schema.org/UserInteraction class.
type UserInteraction struct {
}

// HyperToc is a generated struct representing the https://schema.org/HyperToc class.
type HyperToc struct {
	AssociatedMedia *MediaObject   `json:"associatedmedia"` //https://schema.org/MediaObject
	TocEntry        *HyperTocEntry `json:"tocentry"`        //https://schema.org/HyperTocEntry
}

// JewelryStore is a generated struct representing the https://schema.org/JewelryStore class.
type JewelryStore struct {
}

// MedicalImagingTechnique is a generated struct representing the https://schema.org/MedicalImagingTechnique class.
type MedicalImagingTechnique struct {
}

// Vessel is a generated struct representing the https://schema.org/Vessel class.
type Vessel struct {
}

// BroadcastEvent is a generated struct representing the https://schema.org/BroadcastEvent class.
type BroadcastEvent struct {
	BroadcastOfEvent *Event   `json:"broadcastofevent"` //https://schema.org/Event
	IsLiveBroadcast  *Boolean `json:"islivebroadcast"`  //https://schema.org/Boolean
	SubtitleLanguage *Text    `json:"subtitlelanguage"` //https://schema.org/Text
	VideoFormat      *Text    `json:"videoformat"`      //https://schema.org/Text
}

// Comment is a generated struct representing the https://schema.org/Comment class.
type Comment struct {
	DownvoteCount *Integer      `json:"downvotecount"` //https://schema.org/Integer
	ParentItem    *CreativeWork `json:"parentitem"`    //https://schema.org/CreativeWork
	SharedContent *CreativeWork `json:"sharedcontent"` //https://schema.org/CreativeWork
	UpvoteCount   *Integer      `json:"upvotecount"`   //https://schema.org/Integer
}

// ConstraintNode is a generated struct representing the https://schema.org/ConstraintNode class.
type ConstraintNode struct {
	ConstraintProperty *URL     `json:"constraintproperty"` //https://schema.org/URL
	NumConstraints     *Integer `json:"numconstraints"`     //https://schema.org/Integer
}

// BuyAction is a generated struct representing the https://schema.org/BuyAction class.
type BuyAction struct {
	Vendor          *Person          `json:"vendor"`          //https://schema.org/Person
	WarrantyPromise *WarrantyPromise `json:"warrantypromise"` //https://schema.org/WarrantyPromise
	Seller          *Person          `json:"seller"`          //https://schema.org/Person
}

// GasStation is a generated struct representing the https://schema.org/GasStation class.
type GasStation struct {
}

// SpreadsheetDigitalDocument is a generated struct representing the https://schema.org/SpreadsheetDigitalDocument class.
type SpreadsheetDigitalDocument struct {
}

// ProductCollection is a generated struct representing the https://schema.org/ProductCollection class.
type ProductCollection struct {
	IncludesObject *TypeAndQuantityNode `json:"includesobject"` //https://schema.org/TypeAndQuantityNode
}

// BroadcastService is a generated struct representing the https://schema.org/BroadcastService class.
type BroadcastService struct {
	Area                 *Place            `json:"area"`                 //https://schema.org/Place
	BroadcastAffiliateOf *Organization     `json:"broadcastaffiliateof"` //https://schema.org/Organization
	BroadcastDisplayName *Text             `json:"broadcastdisplayname"` //https://schema.org/Text
	BroadcastFrequency   *Text             `json:"broadcastfrequency"`   //https://schema.org/Text
	BroadcastTimezone    *Text             `json:"broadcasttimezone"`    //https://schema.org/Text
	Broadcaster          *Organization     `json:"broadcaster"`          //https://schema.org/Organization
	CallSign             *Text             `json:"callsign"`             //https://schema.org/Text
	ParentService        *BroadcastService `json:"parentservice"`        //https://schema.org/BroadcastService
	VideoFormat          *Text             `json:"videoformat"`          //https://schema.org/Text
	HasBroadcastChannel  *BroadcastChannel `json:"hasbroadcastchannel"`  //https://schema.org/BroadcastChannel
	InLanguage           *Text             `json:"inlanguage"`           //https://schema.org/Text
}

// MedicalTestPanel is a generated struct representing the https://schema.org/MedicalTestPanel class.
type MedicalTestPanel struct {
	SubTest *MedicalTest `json:"subtest"` //https://schema.org/MedicalTest
}

// Museum is a generated struct representing the https://schema.org/Museum class.
type Museum struct {
}

// DryCleaningOrLaundry is a generated struct representing the https://schema.org/DryCleaningOrLaundry class.
type DryCleaningOrLaundry struct {
}

// GatedResidenceCommunity is a generated struct representing the https://schema.org/GatedResidenceCommunity class.
type GatedResidenceCommunity struct {
}

// LibrarySystem is a generated struct representing the https://schema.org/LibrarySystem class.
type LibrarySystem struct {
}

// MeetingRoom is a generated struct representing the https://schema.org/MeetingRoom class.
type MeetingRoom struct {
}

// MovieClip is a generated struct representing the https://schema.org/MovieClip class.
type MovieClip struct {
}

// PhysicalTherapy is a generated struct representing the https://schema.org/PhysicalTherapy class.
type PhysicalTherapy struct {
}

// BedType is a generated struct representing the https://schema.org/BedType class.
type BedType struct {
}

// DoseSchedule is a generated struct representing the https://schema.org/DoseSchedule class.
type DoseSchedule struct {
	DoseUnit         *Text             `json:"doseunit"`         //https://schema.org/Text
	DoseValue        *QualitativeValue `json:"dosevalue"`        //https://schema.org/QualitativeValue
	TargetPopulation *Text             `json:"targetpopulation"` //https://schema.org/Text
	Frequency        *Text             `json:"frequency"`        //https://schema.org/Text
}

// DietarySupplement is a generated struct representing the https://schema.org/DietarySupplement class.
type DietarySupplement struct {
	ActiveIngredient    *Text                    `json:"activeingredient"`    //https://schema.org/Text
	IsProprietary       *Boolean                 `json:"isproprietary"`       //https://schema.org/Boolean
	LegalStatus         *Text                    `json:"legalstatus"`         //https://schema.org/Text
	MaximumIntake       *MaximumDoseSchedule     `json:"maximumintake"`       //https://schema.org/MaximumDoseSchedule
	MechanismOfAction   *Text                    `json:"mechanismofaction"`   //https://schema.org/Text
	NonProprietaryName  *Text                    `json:"nonproprietaryname"`  //https://schema.org/Text
	ProprietaryName     *Text                    `json:"proprietaryname"`     //https://schema.org/Text
	RecommendedIntake   *RecommendedDoseSchedule `json:"recommendedintake"`   //https://schema.org/RecommendedDoseSchedule
	SafetyConsideration *Text                    `json:"safetyconsideration"` //https://schema.org/Text
	TargetPopulation    *Text                    `json:"targetpopulation"`    //https://schema.org/Text
}

// MediaReviewItem is a generated struct representing the https://schema.org/MediaReviewItem class.
type MediaReviewItem struct {
	MediaItemAppearance *MediaObject `json:"mediaitemappearance"` //https://schema.org/MediaObject
}

// PeopleAudience is a generated struct representing the https://schema.org/PeopleAudience class.
type PeopleAudience struct {
	HealthCondition      *MedicalCondition  `json:"healthcondition"`      //https://schema.org/MedicalCondition
	RequiredGender       *Text              `json:"requiredgender"`       //https://schema.org/Text
	RequiredMaxAge       *Integer           `json:"requiredmaxage"`       //https://schema.org/Integer
	RequiredMinAge       *Integer           `json:"requiredminage"`       //https://schema.org/Integer
	SuggestedAge         *QuantitativeValue `json:"suggestedage"`         //https://schema.org/QuantitativeValue
	SuggestedGender      *Text              `json:"suggestedgender"`      //https://schema.org/Text
	SuggestedMaxAge      *Number            `json:"suggestedmaxage"`      //https://schema.org/Number
	SuggestedMeasurement *QuantitativeValue `json:"suggestedmeasurement"` //https://schema.org/QuantitativeValue
	SuggestedMinAge      *Number            `json:"suggestedminage"`      //https://schema.org/Number
}

// SendAction is a generated struct representing the https://schema.org/SendAction class.
type SendAction struct {
	DeliveryMethod *DeliveryMethod `json:"deliverymethod"` //https://schema.org/DeliveryMethod
	Recipient      *Person         `json:"recipient"`      //https://schema.org/Person
}

// CommunicateAction is a generated struct representing the https://schema.org/CommunicateAction class.
type CommunicateAction struct {
	Language   *Language `json:"language"`   //https://schema.org/Language
	InLanguage *Text     `json:"inlanguage"` //https://schema.org/Text
	About      *Thing    `json:"about"`      //https://schema.org/Thing
	Recipient  *Person   `json:"recipient"`  //https://schema.org/Person
}

// MedicalAudience is a generated struct representing the https://schema.org/MedicalAudience class.
type MedicalAudience struct {
}

// ReturnFeesEnumeration is a generated struct representing the https://schema.org/ReturnFeesEnumeration class.
type ReturnFeesEnumeration struct {
}

// Vein is a generated struct representing the https://schema.org/Vein class.
type Vein struct {
	DrainsTo      *Vessel              `json:"drainsto"`      //https://schema.org/Vessel
	RegionDrained *AnatomicalSystem    `json:"regiondrained"` //https://schema.org/AnatomicalSystem
	Tributary     *AnatomicalStructure `json:"tributary"`     //https://schema.org/AnatomicalStructure
}

// ChooseAction is a generated struct representing the https://schema.org/ChooseAction class.
type ChooseAction struct {
	Option       *Thing `json:"option"`       //https://schema.org/Thing
	ActionOption *Thing `json:"actionoption"` //https://schema.org/Thing
}

// Class is a generated struct representing the https://schema.org/Class class.
type Class struct {
	SupersededBy *Property `json:"supersededby"` //https://schema.org/Property
}

// PerformingGroup is a generated struct representing the https://schema.org/PerformingGroup class.
type PerformingGroup struct {
}

// Substance is a generated struct representing the https://schema.org/Substance class.
type Substance struct {
	ActiveIngredient *Text                `json:"activeingredient"` //https://schema.org/Text
	MaximumIntake    *MaximumDoseSchedule `json:"maximumintake"`    //https://schema.org/MaximumDoseSchedule
}

// Bone is a generated struct representing the https://schema.org/Bone class.
type Bone struct {
}

// EmergencyService is a generated struct representing the https://schema.org/EmergencyService class.
type EmergencyService struct {
}

// EventReservation is a generated struct representing the https://schema.org/EventReservation class.
type EventReservation struct {
}

// HVACBusiness is a generated struct representing the https://schema.org/HVACBusiness class.
type HVACBusiness struct {
}

// MonetaryAmount is a generated struct representing the https://schema.org/MonetaryAmount class.
type MonetaryAmount struct {
	Currency     *Text     `json:"currency"`     //https://schema.org/Text
	MaxValue     *Number   `json:"maxvalue"`     //https://schema.org/Number
	MinValue     *Number   `json:"minvalue"`     //https://schema.org/Number
	ValidFrom    *DateTime `json:"validfrom"`    //https://schema.org/DateTime
	ValidThrough *DateTime `json:"validthrough"` //https://schema.org/DateTime
	Value        *Text     `json:"value"`        //https://schema.org/Text
}

// SpecialAnnouncement is a generated struct representing the https://schema.org/SpecialAnnouncement class.
type SpecialAnnouncement struct {
	AnnouncementLocation        *LocalBusiness     `json:"announcementlocation"`        //https://schema.org/LocalBusiness
	DatePosted                  *DateTime          `json:"dateposted"`                  //https://schema.org/DateTime
	DiseasePreventionInfo       *WebContent        `json:"diseasepreventioninfo"`       //https://schema.org/WebContent
	DiseaseSpreadStatistics     *WebContent        `json:"diseasespreadstatistics"`     //https://schema.org/WebContent
	GettingTestedInfo           *WebContent        `json:"gettingtestedinfo"`           //https://schema.org/WebContent
	GovernmentBenefitsInfo      *GovernmentService `json:"governmentbenefitsinfo"`      //https://schema.org/GovernmentService
	NewsUpdatesAndGuidelines    *WebContent        `json:"newsupdatesandguidelines"`    //https://schema.org/WebContent
	PublicTransportClosuresInfo *WebContent        `json:"publictransportclosuresinfo"` //https://schema.org/WebContent
	QuarantineGuidelines        *WebContent        `json:"quarantineguidelines"`        //https://schema.org/WebContent
	SchoolClosuresInfo          *WebContent        `json:"schoolclosuresinfo"`          //https://schema.org/WebContent
	TravelBans                  *WebContent        `json:"travelbans"`                  //https://schema.org/WebContent
	WebFeed                     *URL               `json:"webfeed"`                     //https://schema.org/URL
	Category                    *URL               `json:"category"`                    //https://schema.org/URL
}

// WearableSizeSystemEnumeration is a generated struct representing the https://schema.org/WearableSizeSystemEnumeration class.
type WearableSizeSystemEnumeration struct {
}

// AboutPage is a generated struct representing the https://schema.org/AboutPage class.
type AboutPage struct {
}

// HealthPlanCostSharingSpecification is a generated struct representing the https://schema.org/HealthPlanCostSharingSpecification class.
type HealthPlanCostSharingSpecification struct {
	HealthPlanCoinsuranceOption *Text               `json:"healthplancoinsuranceoption"` //https://schema.org/Text
	HealthPlanCoinsuranceRate   *Number             `json:"healthplancoinsurancerate"`   //https://schema.org/Number
	HealthPlanCopay             *PriceSpecification `json:"healthplancopay"`             //https://schema.org/PriceSpecification
	HealthPlanCopayOption       *Text               `json:"healthplancopayoption"`       //https://schema.org/Text
	HealthPlanPharmacyCategory  *Text               `json:"healthplanpharmacycategory"`  //https://schema.org/Text
}

// PresentationDigitalDocument is a generated struct representing the https://schema.org/PresentationDigitalDocument class.
type PresentationDigitalDocument struct {
}

// RegisterAction is a generated struct representing the https://schema.org/RegisterAction class.
type RegisterAction struct {
}

// Review is a generated struct representing the https://schema.org/Review class.
type Review struct {
	AssociatedClaimReview *Review     `json:"associatedclaimreview"` //https://schema.org/Review
	AssociatedMediaReview *Review     `json:"associatedmediareview"` //https://schema.org/Review
	ItemReviewed          *Thing      `json:"itemreviewed"`          //https://schema.org/Thing
	NegativeNotes         *WebContent `json:"negativenotes"`         //https://schema.org/WebContent
	PositiveNotes         *WebContent `json:"positivenotes"`         //https://schema.org/WebContent
	ReviewAspect          *Text       `json:"reviewaspect"`          //https://schema.org/Text
	ReviewBody            *Text       `json:"reviewbody"`            //https://schema.org/Text
	ReviewRating          *Rating     `json:"reviewrating"`          //https://schema.org/Rating
	AssociatedReview      *Review     `json:"associatedreview"`      //https://schema.org/Review
}

// UserTweets is a generated struct representing the https://schema.org/UserTweets class.
type UserTweets struct {
}

// GovernmentOffice is a generated struct representing the https://schema.org/GovernmentOffice class.
type GovernmentOffice struct {
}

// PlayAction is a generated struct representing the https://schema.org/PlayAction class.
type PlayAction struct {
	Audience *Audience `json:"audience"` //https://schema.org/Audience
	Event    *Event    `json:"event"`    //https://schema.org/Event
}

// Airport is a generated struct representing the https://schema.org/Airport class.
type Airport struct {
	IataCode *Text `json:"iatacode"` //https://schema.org/Text
	IcaoCode *Text `json:"icaocode"` //https://schema.org/Text
}

// BusStop is a generated struct representing the https://schema.org/BusStop class.
type BusStop struct {
}

// DiscussionForumPosting is a generated struct representing the https://schema.org/DiscussionForumPosting class.
type DiscussionForumPosting struct {
}

// EmploymentAgency is a generated struct representing the https://schema.org/EmploymentAgency class.
type EmploymentAgency struct {
}

// GameAvailabilityEnumeration is a generated struct representing the https://schema.org/GameAvailabilityEnumeration class.
type GameAvailabilityEnumeration struct {
}

// GeospatialGeometry is a generated struct representing the https://schema.org/GeospatialGeometry class.
type GeospatialGeometry struct {
	GeoContains   *Place `json:"geocontains"`   //https://schema.org/Place
	GeoCoveredBy  *Place `json:"geocoveredby"`  //https://schema.org/Place
	GeoCovers     *Place `json:"geocovers"`     //https://schema.org/Place
	GeoCrosses    *Place `json:"geocrosses"`    //https://schema.org/Place
	GeoDisjoint   *Place `json:"geodisjoint"`   //https://schema.org/Place
	GeoEquals     *Place `json:"geoequals"`     //https://schema.org/Place
	GeoIntersects *Place `json:"geointersects"` //https://schema.org/Place
	GeoOverlaps   *Place `json:"geooverlaps"`   //https://schema.org/Place
	GeoTouches    *Place `json:"geotouches"`    //https://schema.org/Place
	GeoWithin     *Place `json:"geowithin"`     //https://schema.org/Place
}

// SellAction is a generated struct representing the https://schema.org/SellAction class.
type SellAction struct {
	Buyer           *Person          `json:"buyer"`           //https://schema.org/Person
	WarrantyPromise *WarrantyPromise `json:"warrantypromise"` //https://schema.org/WarrantyPromise
}

// ComedyEvent is a generated struct representing the https://schema.org/ComedyEvent class.
type ComedyEvent struct {
}

// Florist is a generated struct representing the https://schema.org/Florist class.
type Florist struct {
}

// OccupationalTherapy is a generated struct representing the https://schema.org/OccupationalTherapy class.
type OccupationalTherapy struct {
}

// DataCatalog is a generated struct representing the https://schema.org/DataCatalog class.
type DataCatalog struct {
	MeasurementMethod    *URL     `json:"measurementmethod"`    //https://schema.org/URL
	Dataset              *Dataset `json:"dataset"`              //https://schema.org/Dataset
	MeasurementTechnique *URL     `json:"measurementtechnique"` //https://schema.org/URL
}

// DateTime is a generated struct representing the https://schema.org/DateTime class.
type DateTime struct {
}

// MedicalDevice is a generated struct representing the https://schema.org/MedicalDevice class.
type MedicalDevice struct {
	AdverseOutcome        *MedicalEntity `json:"adverseoutcome"`        //https://schema.org/MedicalEntity
	Contraindication      *Text          `json:"contraindication"`      //https://schema.org/Text
	PostOp                *Text          `json:"postop"`                //https://schema.org/Text
	PreOp                 *Text          `json:"preop"`                 //https://schema.org/Text
	Procedure             *Text          `json:"procedure"`             //https://schema.org/Text
	SeriousAdverseOutcome *MedicalEntity `json:"seriousadverseoutcome"` //https://schema.org/MedicalEntity
}

// Nerve is a generated struct representing the https://schema.org/Nerve class.
type Nerve struct {
	Branch      *AnatomicalStructure `json:"branch"`      //https://schema.org/AnatomicalStructure
	NerveMotor  *Muscle              `json:"nervemotor"`  //https://schema.org/Muscle
	SensoryUnit *SuperficialAnatomy  `json:"sensoryunit"` //https://schema.org/SuperficialAnatomy
	SourcedFrom *BrainStructure      `json:"sourcedfrom"` //https://schema.org/BrainStructure
}

// ParkingFacility is a generated struct representing the https://schema.org/ParkingFacility class.
type ParkingFacility struct {
}

// Table is a generated struct representing the https://schema.org/Table class.
type Table struct {
}

// ChemicalSubstance is a generated struct representing the https://schema.org/ChemicalSubstance class.
type ChemicalSubstance struct {
	ChemicalComposition *Text        `json:"chemicalcomposition"` //https://schema.org/Text
	ChemicalRole        *DefinedTerm `json:"chemicalrole"`        //https://schema.org/DefinedTerm
	PotentialUse        *DefinedTerm `json:"potentialuse"`        //https://schema.org/DefinedTerm
}

// ComicIssue is a generated struct representing the https://schema.org/ComicIssue class.
type ComicIssue struct {
	Artist       *Person `json:"artist"`       //https://schema.org/Person
	Colorist     *Person `json:"colorist"`     //https://schema.org/Person
	Inker        *Person `json:"inker"`        //https://schema.org/Person
	Letterer     *Person `json:"letterer"`     //https://schema.org/Person
	Penciler     *Person `json:"penciler"`     //https://schema.org/Person
	VariantCover *Text   `json:"variantcover"` //https://schema.org/Text
}

// LodgingReservation is a generated struct representing the https://schema.org/LodgingReservation class.
type LodgingReservation struct {
	CheckinTime            *Time              `json:"checkintime"`            //https://schema.org/Time
	CheckoutTime           *Time              `json:"checkouttime"`           //https://schema.org/Time
	LodgingUnitDescription *Text              `json:"lodgingunitdescription"` //https://schema.org/Text
	LodgingUnitType        *Text              `json:"lodgingunittype"`        //https://schema.org/Text
	NumAdults              *QuantitativeValue `json:"numadults"`              //https://schema.org/QuantitativeValue
	NumChildren            *QuantitativeValue `json:"numchildren"`            //https://schema.org/QuantitativeValue
}

// MedicalEntity is a generated struct representing the https://schema.org/MedicalEntity class.
type MedicalEntity struct {
	Code                 *MedicalCode      `json:"code"`                 //https://schema.org/MedicalCode
	Guideline            *MedicalGuideline `json:"guideline"`            //https://schema.org/MedicalGuideline
	LegalStatus          *Text             `json:"legalstatus"`          //https://schema.org/Text
	MedicineSystem       *MedicineSystem   `json:"medicinesystem"`       //https://schema.org/MedicineSystem
	RecognizingAuthority *Organization     `json:"recognizingauthority"` //https://schema.org/Organization
	RelevantSpecialty    *MedicalSpecialty `json:"relevantspecialty"`    //https://schema.org/MedicalSpecialty
	Study                *MedicalStudy     `json:"study"`                //https://schema.org/MedicalStudy
	Funding              *Grant            `json:"funding"`              //https://schema.org/Grant
}

// PaymentMethod is a generated struct representing the https://schema.org/PaymentMethod class.
type PaymentMethod struct {
}

// Researcher is a generated struct representing the https://schema.org/Researcher class.
type Researcher struct {
}

// EmployerAggregateRating is a generated struct representing the https://schema.org/EmployerAggregateRating class.
type EmployerAggregateRating struct {
}

// LinkRole is a generated struct representing the https://schema.org/LinkRole class.
type LinkRole struct {
	LinkRelationship *Text `json:"linkrelationship"` //https://schema.org/Text
	InLanguage       *Text `json:"inlanguage"`       //https://schema.org/Text
}

// MedicalEnumeration is a generated struct representing the https://schema.org/MedicalEnumeration class.
type MedicalEnumeration struct {
}

// TipAction is a generated struct representing the https://schema.org/TipAction class.
type TipAction struct {
	Recipient *Person `json:"recipient"` //https://schema.org/Person
}

// BrokerageAccount is a generated struct representing the https://schema.org/BrokerageAccount class.
type BrokerageAccount struct {
}

// EatAction is a generated struct representing the https://schema.org/EatAction class.
type EatAction struct {
}

// HighSchool is a generated struct representing the https://schema.org/HighSchool class.
type HighSchool struct {
}

// RecommendedDoseSchedule is a generated struct representing the https://schema.org/RecommendedDoseSchedule class.
type RecommendedDoseSchedule struct {
}

// HowToSection is a generated struct representing the https://schema.org/HowToSection class.
type HowToSection struct {
	Steps *Text `json:"steps"` //https://schema.org/Text
}

// UserLikes is a generated struct representing the https://schema.org/UserLikes class.
type UserLikes struct {
}

// Festival is a generated struct representing the https://schema.org/Festival class.
type Festival struct {
}

// OrganizationRole is a generated struct representing the https://schema.org/OrganizationRole class.
type OrganizationRole struct {
	NumberedPosition *Number `json:"numberedposition"` //https://schema.org/Number
}

// Patient is a generated struct representing the https://schema.org/Patient class.
type Patient struct {
	Diagnosis       *MedicalCondition `json:"diagnosis"`       //https://schema.org/MedicalCondition
	Drug            *Drug             `json:"drug"`            //https://schema.org/Drug
	HealthCondition *MedicalCondition `json:"healthcondition"` //https://schema.org/MedicalCondition
}

// SportsTeam is a generated struct representing the https://schema.org/SportsTeam class.
type SportsTeam struct {
	Athlete *Person `json:"athlete"` //https://schema.org/Person
	Coach   *Person `json:"coach"`   //https://schema.org/Person
	Gender  *Text   `json:"gender"`  //https://schema.org/Text
}

// OrganizeAction is a generated struct representing the https://schema.org/OrganizeAction class.
type OrganizeAction struct {
}

// FundingScheme is a generated struct representing the https://schema.org/FundingScheme class.
type FundingScheme struct {
}

// MapCategoryType is a generated struct representing the https://schema.org/MapCategoryType class.
type MapCategoryType struct {
}

// MedicalWebPage is a generated struct representing the https://schema.org/MedicalWebPage class.
type MedicalWebPage struct {
	Aspect          *Text                `json:"aspect"`          //https://schema.org/Text
	MedicalAudience *MedicalAudienceType `json:"medicalaudience"` //https://schema.org/MedicalAudienceType
}

// ServiceChannel is a generated struct representing the https://schema.org/ServiceChannel class.
type ServiceChannel struct {
	AvailableLanguage    *Text          `json:"availablelanguage"`    //https://schema.org/Text
	ProcessingTime       *Duration      `json:"processingtime"`       //https://schema.org/Duration
	ProvidesService      *Service       `json:"providesservice"`      //https://schema.org/Service
	ServiceLocation      *Place         `json:"servicelocation"`      //https://schema.org/Place
	ServicePhone         *ContactPoint  `json:"servicephone"`         //https://schema.org/ContactPoint
	ServicePostalAddress *PostalAddress `json:"servicepostaladdress"` //https://schema.org/PostalAddress
	ServiceSmsNumber     *ContactPoint  `json:"servicesmsnumber"`     //https://schema.org/ContactPoint
	ServiceUrl           *URL           `json:"serviceurl"`           //https://schema.org/URL
}

// BankAccount is a generated struct representing the https://schema.org/BankAccount class.
type BankAccount struct {
	AccountMinimumInflow  *MonetaryAmount `json:"accountminimuminflow"`  //https://schema.org/MonetaryAmount
	AccountOverdraftLimit *MonetaryAmount `json:"accountoverdraftlimit"` //https://schema.org/MonetaryAmount
	BankAccountType       *URL            `json:"bankaccounttype"`       //https://schema.org/URL
}

// AssignAction is a generated struct representing the https://schema.org/AssignAction class.
type AssignAction struct {
}

// MathSolver is a generated struct representing the https://schema.org/MathSolver class.
type MathSolver struct {
	MathExpression *Text `json:"mathexpression"` //https://schema.org/Text
}

// HowToTool is a generated struct representing the https://schema.org/HowToTool class.
type HowToTool struct {
}

// Winery is a generated struct representing the https://schema.org/Winery class.
type Winery struct {
}

// BlogPosting is a generated struct representing the https://schema.org/BlogPosting class.
type BlogPosting struct {
}

// EntertainmentBusiness is a generated struct representing the https://schema.org/EntertainmentBusiness class.
type EntertainmentBusiness struct {
}

// GolfCourse is a generated struct representing the https://schema.org/GolfCourse class.
type GolfCourse struct {
}

// PerformanceRole is a generated struct representing the https://schema.org/PerformanceRole class.
type PerformanceRole struct {
	CharacterName *Text `json:"charactername"` //https://schema.org/Text
}

// CompoundPriceSpecification is a generated struct representing the https://schema.org/CompoundPriceSpecification class.
type CompoundPriceSpecification struct {
	PriceComponent *UnitPriceSpecification `json:"pricecomponent"` //https://schema.org/UnitPriceSpecification
	PriceType      *Text                   `json:"pricetype"`      //https://schema.org/Text
}

// EducationalOccupationalCredential is a generated struct representing the https://schema.org/EducationalOccupationalCredential class.
type EducationalOccupationalCredential struct {
	CompetencyRequired *URL                `json:"competencyrequired"` //https://schema.org/URL
	CredentialCategory *URL                `json:"credentialcategory"` //https://schema.org/URL
	EducationalLevel   *URL                `json:"educationallevel"`   //https://schema.org/URL
	RecognizedBy       *Organization       `json:"recognizedby"`       //https://schema.org/Organization
	ValidFor           *Duration           `json:"validfor"`           //https://schema.org/Duration
	ValidIn            *AdministrativeArea `json:"validin"`            //https://schema.org/AdministrativeArea
}

// DataFeedItem is a generated struct representing the https://schema.org/DataFeedItem class.
type DataFeedItem struct {
	DateDeleted  *DateTime `json:"datedeleted"`  //https://schema.org/DateTime
	DateModified *DateTime `json:"datemodified"` //https://schema.org/DateTime
	Item         *Thing    `json:"item"`         //https://schema.org/Thing
	DateCreated  *DateTime `json:"datecreated"`  //https://schema.org/DateTime
}

// Pond is a generated struct representing the https://schema.org/Pond class.
type Pond struct {
}

// VitalSign is a generated struct representing the https://schema.org/VitalSign class.
type VitalSign struct {
}

// BodyOfWater is a generated struct representing the https://schema.org/BodyOfWater class.
type BodyOfWater struct {
}

// Brewery is a generated struct representing the https://schema.org/Brewery class.
type Brewery struct {
}

// EventSeries is a generated struct representing the https://schema.org/EventSeries class.
type EventSeries struct {
}

// LiveBlogPosting is a generated struct representing the https://schema.org/LiveBlogPosting class.
type LiveBlogPosting struct {
	CoverageEndTime   *DateTime    `json:"coverageendtime"`   //https://schema.org/DateTime
	CoverageStartTime *DateTime    `json:"coveragestarttime"` //https://schema.org/DateTime
	LiveBlogUpdate    *BlogPosting `json:"liveblogupdate"`    //https://schema.org/BlogPosting
}

// RadiationTherapy is a generated struct representing the https://schema.org/RadiationTherapy class.
type RadiationTherapy struct {
}

// DefenceEstablishment is a generated struct representing the https://schema.org/DefenceEstablishment class.
type DefenceEstablishment struct {
}

// Distillery is a generated struct representing the https://schema.org/Distillery class.
type Distillery struct {
}

// MedicalEvidenceLevel is a generated struct representing the https://schema.org/MedicalEvidenceLevel class.
type MedicalEvidenceLevel struct {
}

// MedicalRiskCalculator is a generated struct representing the https://schema.org/MedicalRiskCalculator class.
type MedicalRiskCalculator struct {
}

// OrderStatus is a generated struct representing the https://schema.org/OrderStatus class.
type OrderStatus struct {
}

// ArtGallery is a generated struct representing the https://schema.org/ArtGallery class.
type ArtGallery struct {
}

// Continent is a generated struct representing the https://schema.org/Continent class.
type Continent struct {
}

// EndorsementRating is a generated struct representing the https://schema.org/EndorsementRating class.
type EndorsementRating struct {
}

// HairSalon is a generated struct representing the https://schema.org/HairSalon class.
type HairSalon struct {
}

// RadioEpisode is a generated struct representing the https://schema.org/RadioEpisode class.
type RadioEpisode struct {
}

// TreatmentIndication is a generated struct representing the https://schema.org/TreatmentIndication class.
type TreatmentIndication struct {
}

// GiveAction is a generated struct representing the https://schema.org/GiveAction class.
type GiveAction struct {
	Recipient *Person `json:"recipient"` //https://schema.org/Person
}

// MaximumDoseSchedule is a generated struct representing the https://schema.org/MaximumDoseSchedule class.
type MaximumDoseSchedule struct {
}

// Motel is a generated struct representing the https://schema.org/Motel class.
type Motel struct {
}

// PhysicalActivity is a generated struct representing the https://schema.org/PhysicalActivity class.
type PhysicalActivity struct {
	AssociatedAnatomy *SuperficialAnatomy `json:"associatedanatomy"` //https://schema.org/SuperficialAnatomy
	Epidemiology      *Text               `json:"epidemiology"`      //https://schema.org/Text
	Pathophysiology   *Text               `json:"pathophysiology"`   //https://schema.org/Text
	Category          *URL                `json:"category"`          //https://schema.org/URL
}

// DefinedTermSet is a generated struct representing the https://schema.org/DefinedTermSet class.
type DefinedTermSet struct {
	HasDefinedTerm *DefinedTerm `json:"hasdefinedterm"` //https://schema.org/DefinedTerm
}

// ImageObject is a generated struct representing the https://schema.org/ImageObject class.
type ImageObject struct {
	EmbeddedTextCaption  *Text    `json:"embeddedtextcaption"`  //https://schema.org/Text
	ExifData             *Text    `json:"exifdata"`             //https://schema.org/Text
	RepresentativeOfPage *Boolean `json:"representativeofpage"` //https://schema.org/Boolean
	Caption              *Text    `json:"caption"`              //https://schema.org/Text
}

// MovieRentalStore is a generated struct representing the https://schema.org/MovieRentalStore class.
type MovieRentalStore struct {
}

// OrderAction is a generated struct representing the https://schema.org/OrderAction class.
type OrderAction struct {
	DeliveryMethod *DeliveryMethod `json:"deliverymethod"` //https://schema.org/DeliveryMethod
}

// Painting is a generated struct representing the https://schema.org/Painting class.
type Painting struct {
}

// BusOrCoach is a generated struct representing the https://schema.org/BusOrCoach class.
type BusOrCoach struct {
	AcrissCode *Text              `json:"acrisscode"` //https://schema.org/Text
	RoofLoad   *QuantitativeValue `json:"roofload"`   //https://schema.org/QuantitativeValue
}

// DaySpa is a generated struct representing the https://schema.org/DaySpa class.
type DaySpa struct {
}

// Enumeration is a generated struct representing the https://schema.org/Enumeration class.
type Enumeration struct {
	SupersededBy *Property `json:"supersededby"` //https://schema.org/Property
}

// NonprofitType is a generated struct representing the https://schema.org/NonprofitType class.
type NonprofitType struct {
}

// Reservoir is a generated struct representing the https://schema.org/Reservoir class.
type Reservoir struct {
}

// USNonprofitType is a generated struct representing the https://schema.org/USNonprofitType class.
type USNonprofitType struct {
}

// AutomotiveBusiness is a generated struct representing the https://schema.org/AutomotiveBusiness class.
type AutomotiveBusiness struct {
}

// MedicalContraindication is a generated struct representing the https://schema.org/MedicalContraindication class.
type MedicalContraindication struct {
}

// MerchantReturnPolicySeasonalOverride is a generated struct representing the https://schema.org/MerchantReturnPolicySeasonalOverride class.
type MerchantReturnPolicySeasonalOverride struct {
	EndDate              *DateTime                  `json:"enddate"`              //https://schema.org/DateTime
	MerchantReturnDays   *Integer                   `json:"merchantreturndays"`   //https://schema.org/Integer
	ReturnPolicyCategory *MerchantReturnEnumeration `json:"returnpolicycategory"` //https://schema.org/MerchantReturnEnumeration
	StartDate            *DateTime                  `json:"startdate"`            //https://schema.org/DateTime
}

// WebContent is a generated struct representing the https://schema.org/WebContent class.
type WebContent struct {
}

// BusTrip is a generated struct representing the https://schema.org/BusTrip class.
type BusTrip struct {
	ArrivalBusStop   *BusStop `json:"arrivalbusstop"`   //https://schema.org/BusStop
	BusName          *Text    `json:"busname"`          //https://schema.org/Text
	BusNumber        *Text    `json:"busnumber"`        //https://schema.org/Text
	DepartureBusStop *BusStop `json:"departurebusstop"` //https://schema.org/BusStop
}

// MobilePhoneStore is a generated struct representing the https://schema.org/MobilePhoneStore class.
type MobilePhoneStore struct {
}

// TextDigitalDocument is a generated struct representing the https://schema.org/TextDigitalDocument class.
type TextDigitalDocument struct {
}

// WearableSizeGroupEnumeration is a generated struct representing the https://schema.org/WearableSizeGroupEnumeration class.
type WearableSizeGroupEnumeration struct {
}

// ReservationPackage is a generated struct representing the https://schema.org/ReservationPackage class.
type ReservationPackage struct {
	SubReservation *Reservation `json:"subreservation"` //https://schema.org/Reservation
}

// School is a generated struct representing the https://schema.org/School class.
type School struct {
}

// Store is a generated struct representing the https://schema.org/Store class.
type Store struct {
}

// TVSeries is a generated struct representing the https://schema.org/TVSeries class.
type TVSeries struct {
	Actors            *Person             `json:"actors"`            //https://schema.org/Person
	CountryOfOrigin   *Country            `json:"countryoforigin"`   //https://schema.org/Country
	Directors         *Person             `json:"directors"`         //https://schema.org/Person
	Episodes          *Episode            `json:"episodes"`          //https://schema.org/Episode
	MusicBy           *Person             `json:"musicby"`           //https://schema.org/Person
	NumberOfEpisodes  *Integer            `json:"numberofepisodes"`  //https://schema.org/Integer
	NumberOfSeasons   *Integer            `json:"numberofseasons"`   //https://schema.org/Integer
	ProductionCompany *Organization       `json:"productioncompany"` //https://schema.org/Organization
	Seasons           *CreativeWorkSeason `json:"seasons"`           //https://schema.org/CreativeWorkSeason
	TitleEIDR         *URL                `json:"titleeidr"`         //https://schema.org/URL
	Trailer           *VideoObject        `json:"trailer"`           //https://schema.org/VideoObject
	ContainsSeason    *CreativeWorkSeason `json:"containsseason"`    //https://schema.org/CreativeWorkSeason
	Director          *Person             `json:"director"`          //https://schema.org/Person
	Episode           *Episode            `json:"episode"`           //https://schema.org/Episode
	Season            *URL                `json:"season"`            //https://schema.org/URL
	Actor             *Person             `json:"actor"`             //https://schema.org/Person
}

// TradeAction is a generated struct representing the https://schema.org/TradeAction class.
type TradeAction struct {
	Price              *Text               `json:"price"`              //https://schema.org/Text
	PriceCurrency      *Text               `json:"pricecurrency"`      //https://schema.org/Text
	PriceSpecification *PriceSpecification `json:"pricespecification"` //https://schema.org/PriceSpecification
}

// TrainReservation is a generated struct representing the https://schema.org/TrainReservation class.
type TrainReservation struct {
}

// Grant is a generated struct representing the https://schema.org/Grant class.
type Grant struct {
	Funder     *Person  `json:"funder"`     //https://schema.org/Person
	FundedItem *Product `json:"fundeditem"` //https://schema.org/Product
	Sponsor    *Person  `json:"sponsor"`    //https://schema.org/Person
}

// Joint is a generated struct representing the https://schema.org/Joint class.
type Joint struct {
	BiomechnicalClass *Text `json:"biomechnicalclass"` //https://schema.org/Text
	FunctionalClass   *Text `json:"functionalclass"`   //https://schema.org/Text
	StructuralClass   *Text `json:"structuralclass"`   //https://schema.org/Text
}

// ActivateAction is a generated struct representing the https://schema.org/ActivateAction class.
type ActivateAction struct {
}

// Claim is a generated struct representing the https://schema.org/Claim class.
type Claim struct {
	Appearance       *CreativeWork `json:"appearance"`       //https://schema.org/CreativeWork
	ClaimInterpreter *Person       `json:"claiminterpreter"` //https://schema.org/Person
	FirstAppearance  *CreativeWork `json:"firstappearance"`  //https://schema.org/CreativeWork
}

// BoatTerminal is a generated struct representing the https://schema.org/BoatTerminal class.
type BoatTerminal struct {
}

// HomeGoodsStore is a generated struct representing the https://schema.org/HomeGoodsStore class.
type HomeGoodsStore struct {
}

// InfectiousAgentClass is a generated struct representing the https://schema.org/InfectiousAgentClass class.
type InfectiousAgentClass struct {
}

// RealEstateAgent is a generated struct representing the https://schema.org/RealEstateAgent class.
type RealEstateAgent struct {
}

// SportsEvent is a generated struct representing the https://schema.org/SportsEvent class.
type SportsEvent struct {
	AwayTeam   *SportsTeam `json:"awayteam"`   //https://schema.org/SportsTeam
	HomeTeam   *SportsTeam `json:"hometeam"`   //https://schema.org/SportsTeam
	Sport      *URL        `json:"sport"`      //https://schema.org/URL
	Competitor *SportsTeam `json:"competitor"` //https://schema.org/SportsTeam
}

// TaxiReservation is a generated struct representing the https://schema.org/TaxiReservation class.
type TaxiReservation struct {
	PartySize      *QuantitativeValue `json:"partysize"`      //https://schema.org/QuantitativeValue
	PickupLocation *Place             `json:"pickuplocation"` //https://schema.org/Place
	PickupTime     *DateTime          `json:"pickuptime"`     //https://schema.org/DateTime
}

// ComicSeries is a generated struct representing the https://schema.org/ComicSeries class.
type ComicSeries struct {
}

// ElectronicsStore is a generated struct representing the https://schema.org/ElectronicsStore class.
type ElectronicsStore struct {
}

// RecyclingCenter is a generated struct representing the https://schema.org/RecyclingCenter class.
type RecyclingCenter struct {
}

// UserComments is a generated struct representing the https://schema.org/UserComments class.
type UserComments struct {
	CommentText *Text         `json:"commenttext"` //https://schema.org/Text
	CommentTime *DateTime     `json:"commenttime"` //https://schema.org/DateTime
	Discusses   *CreativeWork `json:"discusses"`   //https://schema.org/CreativeWork
	ReplyToUrl  *URL          `json:"replytourl"`  //https://schema.org/URL
	Creator     *Person       `json:"creator"`     //https://schema.org/Person
}

// DrugPregnancyCategory is a generated struct representing the https://schema.org/DrugPregnancyCategory class.
type DrugPregnancyCategory struct {
}

// EnergyEfficiencyEnumeration is a generated struct representing the https://schema.org/EnergyEfficiencyEnumeration class.
type EnergyEfficiencyEnumeration struct {
}

// FinancialProduct is a generated struct representing the https://schema.org/FinancialProduct class.
type FinancialProduct struct {
	AnnualPercentageRate            *QuantitativeValue `json:"annualpercentagerate"`            //https://schema.org/QuantitativeValue
	FeesAndCommissionsSpecification *URL               `json:"feesandcommissionsspecification"` //https://schema.org/URL
	InterestRate                    *QuantitativeValue `json:"interestrate"`                    //https://schema.org/QuantitativeValue
}

// MotorcycleDealer is a generated struct representing the https://schema.org/MotorcycleDealer class.
type MotorcycleDealer struct {
}

// Playground is a generated struct representing the https://schema.org/Playground class.
type Playground struct {
}

// PostalAddress is a generated struct representing the https://schema.org/PostalAddress class.
type PostalAddress struct {
	AddressCountry      *Text `json:"addresscountry"`      //https://schema.org/Text
	AddressLocality     *Text `json:"addresslocality"`     //https://schema.org/Text
	AddressRegion       *Text `json:"addressregion"`       //https://schema.org/Text
	PostOfficeBoxNumber *Text `json:"postofficeboxnumber"` //https://schema.org/Text
	PostalCode          *Text `json:"postalcode"`          //https://schema.org/Text
	StreetAddress       *Text `json:"streetaddress"`       //https://schema.org/Text
}

// ApplyAction is a generated struct representing the https://schema.org/ApplyAction class.
type ApplyAction struct {
}

// DrugCost is a generated struct representing the https://schema.org/DrugCost class.
type DrugCost struct {
	ApplicableLocation *AdministrativeArea `json:"applicablelocation"` //https://schema.org/AdministrativeArea
	CostCategory       *DrugCostCategory   `json:"costcategory"`       //https://schema.org/DrugCostCategory
	CostCurrency       *Text               `json:"costcurrency"`       //https://schema.org/Text
	CostOrigin         *Text               `json:"costorigin"`         //https://schema.org/Text
	CostPerUnit        *Text               `json:"costperunit"`        //https://schema.org/Text
	DrugUnit           *Text               `json:"drugunit"`           //https://schema.org/Text
}

// Season is a generated struct representing the https://schema.org/Season class.
type Season struct {
}

// DeliveryEvent is a generated struct representing the https://schema.org/DeliveryEvent class.
type DeliveryEvent struct {
	AccessCode        *Text           `json:"accesscode"`        //https://schema.org/Text
	AvailableFrom     *DateTime       `json:"availablefrom"`     //https://schema.org/DateTime
	AvailableThrough  *DateTime       `json:"availablethrough"`  //https://schema.org/DateTime
	HasDeliveryMethod *DeliveryMethod `json:"hasdeliverymethod"` //https://schema.org/DeliveryMethod
}

// DrinkAction is a generated struct representing the https://schema.org/DrinkAction class.
type DrinkAction struct {
}

// PriceSpecification is a generated struct representing the https://schema.org/PriceSpecification class.
type PriceSpecification struct {
	EligibleQuantity          *QuantitativeValue  `json:"eligiblequantity"`          //https://schema.org/QuantitativeValue
	EligibleTransactionVolume *PriceSpecification `json:"eligibletransactionvolume"` //https://schema.org/PriceSpecification
	MaxPrice                  *Number             `json:"maxprice"`                  //https://schema.org/Number
	MinPrice                  *Number             `json:"minprice"`                  //https://schema.org/Number
	Price                     *Text               `json:"price"`                     //https://schema.org/Text
	PriceCurrency             *Text               `json:"pricecurrency"`             //https://schema.org/Text
	ValidFrom                 *DateTime           `json:"validfrom"`                 //https://schema.org/DateTime
	ValidThrough              *DateTime           `json:"validthrough"`              //https://schema.org/DateTime
	ValueAddedTaxIncluded     *Boolean            `json:"valueaddedtaxincluded"`     //https://schema.org/Boolean
}

// TakeAction is a generated struct representing the https://schema.org/TakeAction class.
type TakeAction struct {
}

// EmployerReview is a generated struct representing the https://schema.org/EmployerReview class.
type EmployerReview struct {
}

// LakeBodyOfWater is a generated struct representing the https://schema.org/LakeBodyOfWater class.
type LakeBodyOfWater struct {
}

// OnlineStore is a generated struct representing the https://schema.org/OnlineStore class.
type OnlineStore struct {
}

// CommentAction is a generated struct representing the https://schema.org/CommentAction class.
type CommentAction struct {
	ResultComment *Comment `json:"resultcomment"` //https://schema.org/Comment
}

// WriteAction is a generated struct representing the https://schema.org/WriteAction class.
type WriteAction struct {
	Language   *Language `json:"language"`   //https://schema.org/Language
	InLanguage *Text     `json:"inlanguage"` //https://schema.org/Text
}

// MusicGroup is a generated struct representing the https://schema.org/MusicGroup class.
type MusicGroup struct {
	Albums           *MusicAlbum     `json:"albums"`           //https://schema.org/MusicAlbum
	MusicGroupMember *Person         `json:"musicgroupmember"` //https://schema.org/Person
	Tracks           *MusicRecording `json:"tracks"`           //https://schema.org/MusicRecording
	Album            *MusicAlbum     `json:"album"`            //https://schema.org/MusicAlbum
	Genre            *URL            `json:"genre"`            //https://schema.org/URL
	Track            *MusicRecording `json:"track"`            //https://schema.org/MusicRecording
}

// Occupation is a generated struct representing the https://schema.org/Occupation class.
type Occupation struct {
	EducationRequirements  *Text               `json:"educationrequirements"`  //https://schema.org/Text
	EstimatedSalary        *Number             `json:"estimatedsalary"`        //https://schema.org/Number
	ExperienceRequirements *Text               `json:"experiencerequirements"` //https://schema.org/Text
	OccupationLocation     *AdministrativeArea `json:"occupationlocation"`     //https://schema.org/AdministrativeArea
	OccupationalCategory   *Text               `json:"occupationalcategory"`   //https://schema.org/Text
	Qualifications         *Text               `json:"qualifications"`         //https://schema.org/Text
	Responsibilities       *Text               `json:"responsibilities"`       //https://schema.org/Text
	Skills                 *Text               `json:"skills"`                 //https://schema.org/Text
}

// Optician is a generated struct representing the https://schema.org/Optician class.
type Optician struct {
}

// PawnShop is a generated struct representing the https://schema.org/PawnShop class.
type PawnShop struct {
}

// ScheduleAction is a generated struct representing the https://schema.org/ScheduleAction class.
type ScheduleAction struct {
}

// SelfStorage is a generated struct representing the https://schema.org/SelfStorage class.
type SelfStorage struct {
}

// DataDownload is a generated struct representing the https://schema.org/DataDownload class.
type DataDownload struct {
	MeasurementMethod    *URL `json:"measurementmethod"`    //https://schema.org/URL
	MeasurementTechnique *URL `json:"measurementtechnique"` //https://schema.org/URL
}

// Recipe is a generated struct representing the https://schema.org/Recipe class.
type Recipe struct {
	CookTime           *Duration             `json:"cooktime"`           //https://schema.org/Duration
	CookingMethod      *Text                 `json:"cookingmethod"`      //https://schema.org/Text
	Ingredients        *Text                 `json:"ingredients"`        //https://schema.org/Text
	Nutrition          *NutritionInformation `json:"nutrition"`          //https://schema.org/NutritionInformation
	RecipeCategory     *Text                 `json:"recipecategory"`     //https://schema.org/Text
	RecipeCuisine      *Text                 `json:"recipecuisine"`      //https://schema.org/Text
	RecipeInstructions *Text                 `json:"recipeinstructions"` //https://schema.org/Text
	RecipeYield        *Text                 `json:"recipeyield"`        //https://schema.org/Text
	SuitableForDiet    *RestrictedDiet       `json:"suitablefordiet"`    //https://schema.org/RestrictedDiet
	RecipeIngredient   *Text                 `json:"recipeingredient"`   //https://schema.org/Text
}

// ShippingRateSettings is a generated struct representing the https://schema.org/ShippingRateSettings class.
type ShippingRateSettings struct {
	DoesNotShip           *Boolean        `json:"doesnotship"`           //https://schema.org/Boolean
	FreeShippingThreshold *MonetaryAmount `json:"freeshippingthreshold"` //https://schema.org/MonetaryAmount
	IsUnlabelledFallback  *Boolean        `json:"isunlabelledfallback"`  //https://schema.org/Boolean
	ShippingDestination   *DefinedRegion  `json:"shippingdestination"`   //https://schema.org/DefinedRegion
	ShippingLabel         *Text           `json:"shippinglabel"`         //https://schema.org/Text
	ShippingRate          *MonetaryAmount `json:"shippingrate"`          //https://schema.org/MonetaryAmount
}

// IndividualProduct is a generated struct representing the https://schema.org/IndividualProduct class.
type IndividualProduct struct {
	SerialNumber *Text `json:"serialnumber"` //https://schema.org/Text
}

// MobileApplication is a generated struct representing the https://schema.org/MobileApplication class.
type MobileApplication struct {
	CarrierRequirements *Text `json:"carrierrequirements"` //https://schema.org/Text
}

// AllocateAction is a generated struct representing the https://schema.org/AllocateAction class.
type AllocateAction struct {
}

// DiscoverAction is a generated struct representing the https://schema.org/DiscoverAction class.
type DiscoverAction struct {
}

// DrawAction is a generated struct representing the https://schema.org/DrawAction class.
type DrawAction struct {
}

// EnergyConsumptionDetails is a generated struct representing the https://schema.org/EnergyConsumptionDetails class.
type EnergyConsumptionDetails struct {
	EnergyEfficiencyScaleMax    *EUEnergyEfficiencyEnumeration `json:"energyefficiencyscalemax"`    //https://schema.org/EUEnergyEfficiencyEnumeration
	EnergyEfficiencyScaleMin    *EUEnergyEfficiencyEnumeration `json:"energyefficiencyscalemin"`    //https://schema.org/EUEnergyEfficiencyEnumeration
	HasEnergyEfficiencyCategory *EnergyEfficiencyEnumeration   `json:"hasenergyefficiencycategory"` //https://schema.org/EnergyEfficiencyEnumeration
}

// FilmAction is a generated struct representing the https://schema.org/FilmAction class.
type FilmAction struct {
}

// HowToItem is a generated struct representing the https://schema.org/HowToItem class.
type HowToItem struct {
	RequiredQuantity *Text `json:"requiredquantity"` //https://schema.org/Text
}

// QuantitativeValueDistribution is a generated struct representing the https://schema.org/QuantitativeValueDistribution class.
type QuantitativeValueDistribution struct {
	Median       *Number   `json:"median"`       //https://schema.org/Number
	Percentile10 *Number   `json:"percentile10"` //https://schema.org/Number
	Percentile25 *Number   `json:"percentile25"` //https://schema.org/Number
	Percentile75 *Number   `json:"percentile75"` //https://schema.org/Number
	Percentile90 *Number   `json:"percentile90"` //https://schema.org/Number
	Duration     *Duration `json:"duration"`     //https://schema.org/Duration
}

// TheaterEvent is a generated struct representing the https://schema.org/TheaterEvent class.
type TheaterEvent struct {
}

// ProductGroup is a generated struct representing the https://schema.org/ProductGroup class.
type ProductGroup struct {
	ProductGroupID *Text    `json:"productgroupid"` //https://schema.org/Text
	VariesBy       *Text    `json:"variesby"`       //https://schema.org/Text
	HasVariant     *Product `json:"hasvariant"`     //https://schema.org/Product
}

// SuspendAction is a generated struct representing the https://schema.org/SuspendAction class.
type SuspendAction struct {
}

// DeleteAction is a generated struct representing the https://schema.org/DeleteAction class.
type DeleteAction struct {
}

// HealthTopicContent is a generated struct representing the https://schema.org/HealthTopicContent class.
type HealthTopicContent struct {
	HasHealthAspect *HealthAspectEnumeration `json:"hashealthaspect"` //https://schema.org/HealthAspectEnumeration
}

// MediaSubscription is a generated struct representing the https://schema.org/MediaSubscription class.
type MediaSubscription struct {
	Authenticator       *Organization `json:"authenticator"`       //https://schema.org/Organization
	ExpectsAcceptanceOf *Offer        `json:"expectsacceptanceof"` //https://schema.org/Offer
}

// QuoteAction is a generated struct representing the https://schema.org/QuoteAction class.
type QuoteAction struct {
}

// TrainTrip is a generated struct representing the https://schema.org/TrainTrip class.
type TrainTrip struct {
	ArrivalPlatform   *Text         `json:"arrivalplatform"`   //https://schema.org/Text
	ArrivalStation    *TrainStation `json:"arrivalstation"`    //https://schema.org/TrainStation
	DeparturePlatform *Text         `json:"departureplatform"` //https://schema.org/Text
	DepartureStation  *TrainStation `json:"departurestation"`  //https://schema.org/TrainStation
	TrainName         *Text         `json:"trainname"`         //https://schema.org/Text
	TrainNumber       *Text         `json:"trainnumber"`       //https://schema.org/Text
}

// CurrencyConversionService is a generated struct representing the https://schema.org/CurrencyConversionService class.
type CurrencyConversionService struct {
}

// Legislation is a generated struct representing the https://schema.org/Legislation class.
type Legislation struct {
	LegislationChanges      *Legislation      `json:"legislationchanges"`      //https://schema.org/Legislation
	LegislationConsolidates *Legislation      `json:"legislationconsolidates"` //https://schema.org/Legislation
	LegislationDate         *Date             `json:"legislationdate"`         //https://schema.org/Date
	LegislationDateVersion  *Date             `json:"legislationdateversion"`  //https://schema.org/Date
	LegislationIdentifier   *URL              `json:"legislationidentifier"`   //https://schema.org/URL
	LegislationJurisdiction *Text             `json:"legislationjurisdiction"` //https://schema.org/Text
	LegislationLegalForce   *LegalForceStatus `json:"legislationlegalforce"`   //https://schema.org/LegalForceStatus
	LegislationPassedBy     *Person           `json:"legislationpassedby"`     //https://schema.org/Person
	LegislationResponsible  *Person           `json:"legislationresponsible"`  //https://schema.org/Person
	LegislationTransposes   *Legislation      `json:"legislationtransposes"`   //https://schema.org/Legislation
	LegislationType         *Text             `json:"legislationtype"`         //https://schema.org/Text
	Jurisdiction            *Text             `json:"jurisdiction"`            //https://schema.org/Text
	LegislationApplies      *Legislation      `json:"legislationapplies"`      //https://schema.org/Legislation
}

// Menu is a generated struct representing the https://schema.org/Menu class.
type Menu struct {
	HasMenuItem    *MenuItem    `json:"hasmenuitem"`    //https://schema.org/MenuItem
	HasMenuSection *MenuSection `json:"hasmenusection"` //https://schema.org/MenuSection
}

// OfferCatalog is a generated struct representing the https://schema.org/OfferCatalog class.
type OfferCatalog struct {
}

// OnDemandEvent is a generated struct representing the https://schema.org/OnDemandEvent class.
type OnDemandEvent struct {
}

// AdultEntertainment is a generated struct representing the https://schema.org/AdultEntertainment class.
type AdultEntertainment struct {
}

// Electrician is a generated struct representing the https://schema.org/Electrician class.
type Electrician struct {
}

// NoteDigitalDocument is a generated struct representing the https://schema.org/NoteDigitalDocument class.
type NoteDigitalDocument struct {
}

// CheckAction is a generated struct representing the https://schema.org/CheckAction class.
type CheckAction struct {
}

// Conversation is a generated struct representing the https://schema.org/Conversation class.
type Conversation struct {
}

// PublicToilet is a generated struct representing the https://schema.org/PublicToilet class.
type PublicToilet struct {
}

// TransferAction is a generated struct representing the https://schema.org/TransferAction class.
type TransferAction struct {
	FromLocation *Place `json:"fromlocation"` //https://schema.org/Place
	ToLocation   *Place `json:"tolocation"`   //https://schema.org/Place
}

// CafeOrCoffeeShop is a generated struct representing the https://schema.org/CafeOrCoffeeShop class.
type CafeOrCoffeeShop struct {
}

// Dentist is a generated struct representing the https://schema.org/Dentist class.
type Dentist struct {
}

// OfferForLease is a generated struct representing the https://schema.org/OfferForLease class.
type OfferForLease struct {
}

// ResearchOrganization is a generated struct representing the https://schema.org/ResearchOrganization class.
type ResearchOrganization struct {
}

// BreadcrumbList is a generated struct representing the https://schema.org/BreadcrumbList class.
type BreadcrumbList struct {
}

// Course is a generated struct representing the https://schema.org/Course class.
type Course struct {
	AvailableLanguage             *Text            `json:"availablelanguage"`             //https://schema.org/Text
	CourseCode                    *Text            `json:"coursecode"`                    //https://schema.org/Text
	CoursePrerequisites           *Text            `json:"courseprerequisites"`           //https://schema.org/Text
	EducationalCredentialAwarded  *URL             `json:"educationalcredentialawarded"`  //https://schema.org/URL
	FinancialAidEligible          *Text            `json:"financialaideligible"`          //https://schema.org/Text
	HasCourseInstance             *CourseInstance  `json:"hascourseinstance"`             //https://schema.org/CourseInstance
	NumberOfCredits               *StructuredValue `json:"numberofcredits"`               //https://schema.org/StructuredValue
	OccupationalCredentialAwarded *URL             `json:"occupationalcredentialawarded"` //https://schema.org/URL
	SyllabusSections              *Syllabus        `json:"syllabussections"`              //https://schema.org/Syllabus
	TotalHistoricalEnrollment     *Integer         `json:"totalhistoricalenrollment"`     //https://schema.org/Integer
}

// ProductModel is a generated struct representing the https://schema.org/ProductModel class.
type ProductModel struct {
	PredecessorOf *ProductModel `json:"predecessorof"` //https://schema.org/ProductModel
	SuccessorOf   *ProductModel `json:"successorof"`   //https://schema.org/ProductModel
	IsVariantOf   *ProductModel `json:"isvariantof"`   //https://schema.org/ProductModel
}

// ViewAction is a generated struct representing the https://schema.org/ViewAction class.
type ViewAction struct {
}

// FoodEstablishmentReservation is a generated struct representing the https://schema.org/FoodEstablishmentReservation class.
type FoodEstablishmentReservation struct {
	EndTime   *Time              `json:"endtime"`   //https://schema.org/Time
	PartySize *QuantitativeValue `json:"partysize"` //https://schema.org/QuantitativeValue
	StartTime *Time              `json:"starttime"` //https://schema.org/Time
}

// Guide is a generated struct representing the https://schema.org/Guide class.
type Guide struct {
	ReviewAspect *Text `json:"reviewaspect"` //https://schema.org/Text
}

// PaymentChargeSpecification is a generated struct representing the https://schema.org/PaymentChargeSpecification class.
type PaymentChargeSpecification struct {
	AppliesToDeliveryMethod *DeliveryMethod `json:"appliestodeliverymethod"` //https://schema.org/DeliveryMethod
	AppliesToPaymentMethod  *PaymentMethod  `json:"appliestopaymentmethod"`  //https://schema.org/PaymentMethod
}

// ResearchProject is a generated struct representing the https://schema.org/ResearchProject class.
type ResearchProject struct {
}

// Airline is a generated struct representing the https://schema.org/Airline class.
type Airline struct {
	BoardingPolicy *BoardingPolicyType `json:"boardingpolicy"` //https://schema.org/BoardingPolicyType
	IataCode       *Text               `json:"iatacode"`       //https://schema.org/Text
}

// Country is a generated struct representing the https://schema.org/Country class.
type Country struct {
}

// WebSite is a generated struct representing the https://schema.org/WebSite class.
type WebSite struct {
	Issn *Text `json:"issn"` //https://schema.org/Text
}

// AutoBodyShop is a generated struct representing the https://schema.org/AutoBodyShop class.
type AutoBodyShop struct {
}

// CollectionPage is a generated struct representing the https://schema.org/CollectionPage class.
type CollectionPage struct {
}

// TattooParlor is a generated struct representing the https://schema.org/TattooParlor class.
type TattooParlor struct {
}

// HowToSupply is a generated struct representing the https://schema.org/HowToSupply class.
type HowToSupply struct {
	EstimatedCost *Text `json:"estimatedcost"` //https://schema.org/Text
}

// TVSeason is a generated struct representing the https://schema.org/TVSeason class.
type TVSeason struct {
	CountryOfOrigin *Country  `json:"countryoforigin"` //https://schema.org/Country
	PartOfTVSeries  *TVSeries `json:"partoftvseries"`  //https://schema.org/TVSeries
	TitleEIDR       *URL      `json:"titleeidr"`       //https://schema.org/URL
}

// MonetaryGrant is a generated struct representing the https://schema.org/MonetaryGrant class.
type MonetaryGrant struct {
	Amount *Number `json:"amount"` //https://schema.org/Number
	Funder *Person `json:"funder"` //https://schema.org/Person
}

// SingleFamilyResidence is a generated struct representing the https://schema.org/SingleFamilyResidence class.
type SingleFamilyResidence struct {
	NumberOfRooms *QuantitativeValue `json:"numberofrooms"` //https://schema.org/QuantitativeValue
	Occupancy     *QuantitativeValue `json:"occupancy"`     //https://schema.org/QuantitativeValue
}

// AggregateOffer is a generated struct representing the https://schema.org/AggregateOffer class.
type AggregateOffer struct {
	HighPrice  *Text    `json:"highprice"`  //https://schema.org/Text
	LowPrice   *Text    `json:"lowprice"`   //https://schema.org/Text
	OfferCount *Integer `json:"offercount"` //https://schema.org/Integer
	Offers     *Offer   `json:"offers"`     //https://schema.org/Offer
}

// Answer is a generated struct representing the https://schema.org/Answer class.
type Answer struct {
	AnswerExplanation *WebContent   `json:"answerexplanation"` //https://schema.org/WebContent
	ParentItem        *CreativeWork `json:"parentitem"`        //https://schema.org/CreativeWork
}

// ClaimReview is a generated struct representing the https://schema.org/ClaimReview class.
type ClaimReview struct {
	ClaimReviewed *Text `json:"claimreviewed"` //https://schema.org/Text
}

// ContactPage is a generated struct representing the https://schema.org/ContactPage class.
type ContactPage struct {
}

// DepositAccount is a generated struct representing the https://schema.org/DepositAccount class.
type DepositAccount struct {
}

// LymphaticVessel is a generated struct representing the https://schema.org/LymphaticVessel class.
type LymphaticVessel struct {
	OriginatesFrom *Vessel           `json:"originatesfrom"` //https://schema.org/Vessel
	RegionDrained  *AnatomicalSystem `json:"regiondrained"`  //https://schema.org/AnatomicalSystem
	RunsTo         *Vessel           `json:"runsto"`         //https://schema.org/Vessel
}

// DislikeAction is a generated struct representing the https://schema.org/DislikeAction class.
type DislikeAction struct {
}

// GenderType is a generated struct representing the https://schema.org/GenderType class.
type GenderType struct {
}

// DefinedTerm is a generated struct representing the https://schema.org/DefinedTerm class.
type DefinedTerm struct {
	InDefinedTermSet *URL  `json:"indefinedtermset"` //https://schema.org/URL
	TermCode         *Text `json:"termcode"`         //https://schema.org/Text
}

// ListenAction is a generated struct representing the https://schema.org/ListenAction class.
type ListenAction struct {
}

// ParcelDelivery is a generated struct representing the https://schema.org/ParcelDelivery class.
type ParcelDelivery struct {
	Carrier              *Organization   `json:"carrier"`              //https://schema.org/Organization
	DeliveryAddress      *PostalAddress  `json:"deliveryaddress"`      //https://schema.org/PostalAddress
	DeliveryStatus       *DeliveryEvent  `json:"deliverystatus"`       //https://schema.org/DeliveryEvent
	ExpectedArrivalFrom  *DateTime       `json:"expectedarrivalfrom"`  //https://schema.org/DateTime
	ExpectedArrivalUntil *DateTime       `json:"expectedarrivaluntil"` //https://schema.org/DateTime
	HasDeliveryMethod    *DeliveryMethod `json:"hasdeliverymethod"`    //https://schema.org/DeliveryMethod
	ItemShipped          *Product        `json:"itemshipped"`          //https://schema.org/Product
	OriginAddress        *PostalAddress  `json:"originaddress"`        //https://schema.org/PostalAddress
	PartOfOrder          *Order          `json:"partoforder"`          //https://schema.org/Order
	TrackingNumber       *Text           `json:"trackingnumber"`       //https://schema.org/Text
	TrackingUrl          *URL            `json:"trackingurl"`          //https://schema.org/URL
	Provider             *Person         `json:"provider"`             //https://schema.org/Person
}

// TireShop is a generated struct representing the https://schema.org/TireShop class.
type TireShop struct {
}

// ChildrensEvent is a generated struct representing the https://schema.org/ChildrensEvent class.
type ChildrensEvent struct {
}

// MusicAlbum is a generated struct representing the https://schema.org/MusicAlbum class.
type MusicAlbum struct {
	AlbumProductionType *MusicAlbumProductionType `json:"albumproductiontype"` //https://schema.org/MusicAlbumProductionType
	AlbumReleaseType    *MusicAlbumReleaseType    `json:"albumreleasetype"`    //https://schema.org/MusicAlbumReleaseType
	ByArtist            *Person                   `json:"byartist"`            //https://schema.org/Person
	AlbumRelease        *MusicRelease             `json:"albumrelease"`        //https://schema.org/MusicRelease
}

// SportsClub is a generated struct representing the https://schema.org/SportsClub class.
type SportsClub struct {
}

// TherapeuticProcedure is a generated struct representing the https://schema.org/TherapeuticProcedure class.
type TherapeuticProcedure struct {
	AdverseOutcome *MedicalEntity `json:"adverseoutcome"` //https://schema.org/MedicalEntity
	DoseSchedule   *DoseSchedule  `json:"doseschedule"`   //https://schema.org/DoseSchedule
	Drug           *Drug          `json:"drug"`           //https://schema.org/Drug
}

// MedicalCause is a generated struct representing the https://schema.org/MedicalCause class.
type MedicalCause struct {
	CauseOf *MedicalEntity `json:"causeof"` //https://schema.org/MedicalEntity
}

// JobPosting is a generated struct representing the https://schema.org/JobPosting class.
type JobPosting struct {
	ApplicantLocationRequirements *AdministrativeArea `json:"applicantlocationrequirements"` //https://schema.org/AdministrativeArea
	ApplicationContact            *ContactPoint       `json:"applicationcontact"`            //https://schema.org/ContactPoint
	BaseSalary                    *PriceSpecification `json:"basesalary"`                    //https://schema.org/PriceSpecification
	Benefits                      *Text               `json:"benefits"`                      //https://schema.org/Text
	DatePosted                    *DateTime           `json:"dateposted"`                    //https://schema.org/DateTime
	DirectApply                   *Boolean            `json:"directapply"`                   //https://schema.org/Boolean
	EducationRequirements         *Text               `json:"educationrequirements"`         //https://schema.org/Text
	EligibilityToWorkRequirement  *Text               `json:"eligibilitytoworkrequirement"`  //https://schema.org/Text
	EmployerOverview              *Text               `json:"employeroverview"`              //https://schema.org/Text
	EmploymentType                *Text               `json:"employmenttype"`                //https://schema.org/Text
	EmploymentUnit                *Organization       `json:"employmentunit"`                //https://schema.org/Organization
	EstimatedSalary               *Number             `json:"estimatedsalary"`               //https://schema.org/Number
	ExperienceInPlaceOfEducation  *Boolean            `json:"experienceinplaceofeducation"`  //https://schema.org/Boolean
	ExperienceRequirements        *Text               `json:"experiencerequirements"`        //https://schema.org/Text
	HiringOrganization            *Person             `json:"hiringorganization"`            //https://schema.org/Person
	Incentives                    *Text               `json:"incentives"`                    //https://schema.org/Text
	Industry                      *Text               `json:"industry"`                      //https://schema.org/Text
	JobImmediateStart             *Boolean            `json:"jobimmediatestart"`             //https://schema.org/Boolean
	JobLocation                   *Place              `json:"joblocation"`                   //https://schema.org/Place
	JobLocationType               *Text               `json:"joblocationtype"`               //https://schema.org/Text
	JobStartDate                  *Text               `json:"jobstartdate"`                  //https://schema.org/Text
	OccupationalCategory          *Text               `json:"occupationalcategory"`          //https://schema.org/Text
	PhysicalRequirement           *URL                `json:"physicalrequirement"`           //https://schema.org/URL
	Qualifications                *Text               `json:"qualifications"`                //https://schema.org/Text
	RelevantOccupation            *Occupation         `json:"relevantoccupation"`            //https://schema.org/Occupation
	Responsibilities              *Text               `json:"responsibilities"`              //https://schema.org/Text
	SalaryCurrency                *Text               `json:"salarycurrency"`                //https://schema.org/Text
	SecurityClearanceRequirement  *URL                `json:"securityclearancerequirement"`  //https://schema.org/URL
	SensoryRequirement            *URL                `json:"sensoryrequirement"`            //https://schema.org/URL
	Skills                        *Text               `json:"skills"`                        //https://schema.org/Text
	SpecialCommitments            *Text               `json:"specialcommitments"`            //https://schema.org/Text
	Title                         *Text               `json:"title"`                         //https://schema.org/Text
	TotalJobOpenings              *Integer            `json:"totaljobopenings"`              //https://schema.org/Integer
	ValidThrough                  *DateTime           `json:"validthrough"`                  //https://schema.org/DateTime
	WorkHours                     *Text               `json:"workhours"`                     //https://schema.org/Text
	IncentiveCompensation         *Text               `json:"incentivecompensation"`         //https://schema.org/Text
	JobBenefits                   *Text               `json:"jobbenefits"`                   //https://schema.org/Text
}

// PoliticalParty is a generated struct representing the https://schema.org/PoliticalParty class.
type PoliticalParty struct {
}

// ReviewAction is a generated struct representing the https://schema.org/ReviewAction class.
type ReviewAction struct {
	ResultReview *Review `json:"resultreview"` //https://schema.org/Review
}

// Energy is a generated struct representing the https://schema.org/Energy class.
type Energy struct {
}

// MoneyTransfer is a generated struct representing the https://schema.org/MoneyTransfer class.
type MoneyTransfer struct {
	Amount          *Number `json:"amount"`          //https://schema.org/Number
	BeneficiaryBank *Text   `json:"beneficiarybank"` //https://schema.org/Text
}

// ReviewNewsArticle is a generated struct representing the https://schema.org/ReviewNewsArticle class.
type ReviewNewsArticle struct {
}

// ShortStory is a generated struct representing the https://schema.org/ShortStory class.
type ShortStory struct {
}

// City is a generated struct representing the https://schema.org/City class.
type City struct {
}

// Drug is a generated struct representing the https://schema.org/Drug class.
type Drug struct {
	ActiveIngredient              *Text                  `json:"activeingredient"`              //https://schema.org/Text
	AdministrationRoute           *Text                  `json:"administrationroute"`           //https://schema.org/Text
	AlcoholWarning                *Text                  `json:"alcoholwarning"`                //https://schema.org/Text
	AvailableStrength             *DrugStrength          `json:"availablestrength"`             //https://schema.org/DrugStrength
	BreastfeedingWarning          *Text                  `json:"breastfeedingwarning"`          //https://schema.org/Text
	ClincalPharmacology           *Text                  `json:"clincalpharmacology"`           //https://schema.org/Text
	DosageForm                    *Text                  `json:"dosageform"`                    //https://schema.org/Text
	DoseSchedule                  *DoseSchedule          `json:"doseschedule"`                  //https://schema.org/DoseSchedule
	DrugClass                     *DrugClass             `json:"drugclass"`                     //https://schema.org/DrugClass
	DrugUnit                      *Text                  `json:"drugunit"`                      //https://schema.org/Text
	FoodWarning                   *Text                  `json:"foodwarning"`                   //https://schema.org/Text
	IncludedInHealthInsurancePlan *HealthInsurancePlan   `json:"includedinhealthinsuranceplan"` //https://schema.org/HealthInsurancePlan
	InteractingDrug               *Drug                  `json:"interactingdrug"`               //https://schema.org/Drug
	IsAvailableGenerically        *Boolean               `json:"isavailablegenerically"`        //https://schema.org/Boolean
	IsProprietary                 *Boolean               `json:"isproprietary"`                 //https://schema.org/Boolean
	LabelDetails                  *URL                   `json:"labeldetails"`                  //https://schema.org/URL
	LegalStatus                   *Text                  `json:"legalstatus"`                   //https://schema.org/Text
	MaximumIntake                 *MaximumDoseSchedule   `json:"maximumintake"`                 //https://schema.org/MaximumDoseSchedule
	MechanismOfAction             *Text                  `json:"mechanismofaction"`             //https://schema.org/Text
	NonProprietaryName            *Text                  `json:"nonproprietaryname"`            //https://schema.org/Text
	Overdosage                    *Text                  `json:"overdosage"`                    //https://schema.org/Text
	PregnancyCategory             *DrugPregnancyCategory `json:"pregnancycategory"`             //https://schema.org/DrugPregnancyCategory
	PregnancyWarning              *Text                  `json:"pregnancywarning"`              //https://schema.org/Text
	PrescribingInfo               *URL                   `json:"prescribinginfo"`               //https://schema.org/URL
	PrescriptionStatus            *Text                  `json:"prescriptionstatus"`            //https://schema.org/Text
	ProprietaryName               *Text                  `json:"proprietaryname"`               //https://schema.org/Text
	RelatedDrug                   *Drug                  `json:"relateddrug"`                   //https://schema.org/Drug
	Rxcui                         *Text                  `json:"rxcui"`                         //https://schema.org/Text
	Warning                       *URL                   `json:"warning"`                       //https://schema.org/URL
	ClinicalPharmacology          *Text                  `json:"clinicalpharmacology"`          //https://schema.org/Text
}

// FMRadioChannel is a generated struct representing the https://schema.org/FMRadioChannel class.
type FMRadioChannel struct {
}

// LegislationObject is a generated struct representing the https://schema.org/LegislationObject class.
type LegislationObject struct {
	LegislationLegalValue *LegalValueLevel `json:"legislationlegalvalue"` //https://schema.org/LegalValueLevel
}

// Event is a generated struct representing the https://schema.org/Event class.
type Event struct {
	AggregateRating                 *AggregateRating                `json:"aggregaterating"`                 //https://schema.org/AggregateRating
	Attendees                       *Person                         `json:"attendees"`                       //https://schema.org/Person
	Composer                        *Person                         `json:"composer"`                        //https://schema.org/Person
	Contributor                     *Person                         `json:"contributor"`                     //https://schema.org/Person
	DoorTime                        *Time                           `json:"doortime"`                        //https://schema.org/Time
	EndDate                         *DateTime                       `json:"enddate"`                         //https://schema.org/DateTime
	EventAttendanceMode             *EventAttendanceModeEnumeration `json:"eventattendancemode"`             //https://schema.org/EventAttendanceModeEnumeration
	EventSchedule                   *Schedule                       `json:"eventschedule"`                   //https://schema.org/Schedule
	EventStatus                     *EventStatusType                `json:"eventstatus"`                     //https://schema.org/EventStatusType
	Funder                          *Person                         `json:"funder"`                          //https://schema.org/Person
	Keywords                        *URL                            `json:"keywords"`                        //https://schema.org/URL
	MaximumAttendeeCapacity         *Integer                        `json:"maximumattendeecapacity"`         //https://schema.org/Integer
	MaximumPhysicalAttendeeCapacity *Integer                        `json:"maximumphysicalattendeecapacity"` //https://schema.org/Integer
	MaximumVirtualAttendeeCapacity  *Integer                        `json:"maximumvirtualattendeecapacity"`  //https://schema.org/Integer
	Organizer                       *Person                         `json:"organizer"`                       //https://schema.org/Person
	Performers                      *Person                         `json:"performers"`                      //https://schema.org/Person
	PreviousStartDate               *Date                           `json:"previousstartdate"`               //https://schema.org/Date
	RemainingAttendeeCapacity       *Integer                        `json:"remainingattendeecapacity"`       //https://schema.org/Integer
	StartDate                       *DateTime                       `json:"startdate"`                       //https://schema.org/DateTime
	SubEvents                       *Event                          `json:"subevents"`                       //https://schema.org/Event
	Translator                      *Person                         `json:"translator"`                      //https://schema.org/Person
	TypicalAgeRange                 *Text                           `json:"typicalagerange"`                 //https://schema.org/Text
	WorkPerformed                   *CreativeWork                   `json:"workperformed"`                   //https://schema.org/CreativeWork
	Attendee                        *Person                         `json:"attendee"`                        //https://schema.org/Person
	Audience                        *Audience                       `json:"audience"`                        //https://schema.org/Audience
	Director                        *Person                         `json:"director"`                        //https://schema.org/Person
	Duration                        *Duration                       `json:"duration"`                        //https://schema.org/Duration
	Funding                         *Grant                          `json:"funding"`                         //https://schema.org/Grant
	InLanguage                      *Text                           `json:"inlanguage"`                      //https://schema.org/Text
	IsAccessibleForFree             *Boolean                        `json:"isaccessibleforfree"`             //https://schema.org/Boolean
	Offers                          *Offer                          `json:"offers"`                          //https://schema.org/Offer
	Performer                       *Person                         `json:"performer"`                       //https://schema.org/Person
	RecordedIn                      *CreativeWork                   `json:"recordedin"`                      //https://schema.org/CreativeWork
	Review                          *Review                         `json:"review"`                          //https://schema.org/Review
	Sponsor                         *Person                         `json:"sponsor"`                         //https://schema.org/Person
	SuperEvent                      *Event                          `json:"superevent"`                      //https://schema.org/Event
	About                           *Thing                          `json:"about"`                           //https://schema.org/Thing
	Actor                           *Person                         `json:"actor"`                           //https://schema.org/Person
	SubEvent                        *Event                          `json:"subevent"`                        //https://schema.org/Event
	WorkFeatured                    *CreativeWork                   `json:"workfeatured"`                    //https://schema.org/CreativeWork
	Location                        *VirtualLocation                `json:"location"`                        //https://schema.org/VirtualLocation
}

// VisualArtsEvent is a generated struct representing the https://schema.org/VisualArtsEvent class.
type VisualArtsEvent struct {
}

// ApprovedIndication is a generated struct representing the https://schema.org/ApprovedIndication class.
type ApprovedIndication struct {
}

// BoardingPolicyType is a generated struct representing the https://schema.org/BoardingPolicyType class.
type BoardingPolicyType struct {
}

// BookmarkAction is a generated struct representing the https://schema.org/BookmarkAction class.
type BookmarkAction struct {
}

// Drawing is a generated struct representing the https://schema.org/Drawing class.
type Drawing struct {
}

// LeaveAction is a generated struct representing the https://schema.org/LeaveAction class.
type LeaveAction struct {
	Event *Event `json:"event"` //https://schema.org/Event
}

// CarUsageType is a generated struct representing the https://schema.org/CarUsageType class.
type CarUsageType struct {
}

// CreativeWorkSeries is a generated struct representing the https://schema.org/CreativeWorkSeries class.
type CreativeWorkSeries struct {
	EndDate   *DateTime `json:"enddate"`   //https://schema.org/DateTime
	Issn      *Text     `json:"issn"`      //https://schema.org/Text
	StartDate *DateTime `json:"startdate"` //https://schema.org/DateTime
}

// State is a generated struct representing the https://schema.org/State class.
type State struct {
}

// VeterinaryCare is a generated struct representing the https://schema.org/VeterinaryCare class.
type VeterinaryCare struct {
}

// WebPageElement is a generated struct representing the https://schema.org/WebPageElement class.
type WebPageElement struct {
	CssSelector *CssSelectorType `json:"cssselector"` //https://schema.org/CssSelectorType
	Xpath       *XPathType       `json:"xpath"`       //https://schema.org/XPathType
}

// DrugLegalStatus is a generated struct representing the https://schema.org/DrugLegalStatus class.
type DrugLegalStatus struct {
	ApplicableLocation *AdministrativeArea `json:"applicablelocation"` //https://schema.org/AdministrativeArea
}

// InviteAction is a generated struct representing the https://schema.org/InviteAction class.
type InviteAction struct {
	Event *Event `json:"event"` //https://schema.org/Event
}

// MedicalIndication is a generated struct representing the https://schema.org/MedicalIndication class.
type MedicalIndication struct {
}

// RestrictedDiet is a generated struct representing the https://schema.org/RestrictedDiet class.
type RestrictedDiet struct {
}

// WholesaleStore is a generated struct representing the https://schema.org/WholesaleStore class.
type WholesaleStore struct {
}

// AssessAction is a generated struct representing the https://schema.org/AssessAction class.
type AssessAction struct {
}

// BookStore is a generated struct representing the https://schema.org/BookStore class.
type BookStore struct {
}

// CourseInstance is a generated struct representing the https://schema.org/CourseInstance class.
type CourseInstance struct {
	CourseMode     *URL      `json:"coursemode"`     //https://schema.org/URL
	CourseSchedule *Schedule `json:"courseschedule"` //https://schema.org/Schedule
	CourseWorkload *Text     `json:"courseworkload"` //https://schema.org/Text
	Instructor     *Person   `json:"instructor"`     //https://schema.org/Person
}

// TelevisionStation is a generated struct representing the https://schema.org/TelevisionStation class.
type TelevisionStation struct {
}

// Text is a generated struct representing the https://schema.org/Text class.
type Text struct {
}

// ToyStore is a generated struct representing the https://schema.org/ToyStore class.
type ToyStore struct {
}

// ComputerLanguage is a generated struct representing the https://schema.org/ComputerLanguage class.
type ComputerLanguage struct {
}

// EducationalOccupationalProgram is a generated struct representing the https://schema.org/EducationalOccupationalProgram class.
type EducationalOccupationalProgram struct {
	ApplicationDeadline           *Date                       `json:"applicationdeadline"`           //https://schema.org/Date
	ApplicationStartDate          *Date                       `json:"applicationstartdate"`          //https://schema.org/Date
	DayOfWeek                     *DayOfWeek                  `json:"dayofweek"`                     //https://schema.org/DayOfWeek
	EducationalCredentialAwarded  *URL                        `json:"educationalcredentialawarded"`  //https://schema.org/URL
	EducationalProgramMode        *URL                        `json:"educationalprogrammode"`        //https://schema.org/URL
	EndDate                       *DateTime                   `json:"enddate"`                       //https://schema.org/DateTime
	FinancialAidEligible          *Text                       `json:"financialaideligible"`          //https://schema.org/Text
	HasCourse                     *Course                     `json:"hascourse"`                     //https://schema.org/Course
	MaximumEnrollment             *Integer                    `json:"maximumenrollment"`             //https://schema.org/Integer
	NumberOfCredits               *StructuredValue            `json:"numberofcredits"`               //https://schema.org/StructuredValue
	OccupationalCategory          *Text                       `json:"occupationalcategory"`          //https://schema.org/Text
	OccupationalCredentialAwarded *URL                        `json:"occupationalcredentialawarded"` //https://schema.org/URL
	ProgramPrerequisites          *Text                       `json:"programprerequisites"`          //https://schema.org/Text
	ProgramType                   *Text                       `json:"programtype"`                   //https://schema.org/Text
	SalaryUponCompletion          *MonetaryAmountDistribution `json:"salaryuponcompletion"`          //https://schema.org/MonetaryAmountDistribution
	StartDate                     *DateTime                   `json:"startdate"`                     //https://schema.org/DateTime
	TermDuration                  *Duration                   `json:"termduration"`                  //https://schema.org/Duration
	TermsPerYear                  *Number                     `json:"termsperyear"`                  //https://schema.org/Number
	TimeOfDay                     *Text                       `json:"timeofday"`                     //https://schema.org/Text
	TimeToComplete                *Duration                   `json:"timetocomplete"`                //https://schema.org/Duration
	TrainingSalary                *MonetaryAmountDistribution `json:"trainingsalary"`                //https://schema.org/MonetaryAmountDistribution
	TypicalCreditsPerTerm         *StructuredValue            `json:"typicalcreditsperterm"`         //https://schema.org/StructuredValue
	Offers                        *Offer                      `json:"offers"`                        //https://schema.org/Offer
	Provider                      *Person                     `json:"provider"`                      //https://schema.org/Person
}

// ImagingTest is a generated struct representing the https://schema.org/ImagingTest class.
type ImagingTest struct {
	ImagingTechnique *MedicalImagingTechnique `json:"imagingtechnique"` //https://schema.org/MedicalImagingTechnique
}

// MediaReview is a generated struct representing the https://schema.org/MediaReview class.
type MediaReview struct {
	MediaAuthenticityCategory       *MediaManipulationRatingEnumeration `json:"mediaauthenticitycategory"`       //https://schema.org/MediaManipulationRatingEnumeration
	OriginalMediaContextDescription *Text                               `json:"originalmediacontextdescription"` //https://schema.org/Text
	OriginalMediaLink               *WebPage                            `json:"originalmedialink"`               //https://schema.org/WebPage
}

// ReplyAction is a generated struct representing the https://schema.org/ReplyAction class.
type ReplyAction struct {
	ResultComment *Comment `json:"resultcomment"` //https://schema.org/Comment
}

// Taxi is a generated struct representing the https://schema.org/Taxi class.
type Taxi struct {
}

// HowToTip is a generated struct representing the https://schema.org/HowToTip class.
type HowToTip struct {
}

// ImageObjectSnapshot is a generated struct representing the https://schema.org/ImageObjectSnapshot class.
type ImageObjectSnapshot struct {
}

// MiddleSchool is a generated struct representing the https://schema.org/MiddleSchool class.
type MiddleSchool struct {
}

// ParentAudience is a generated struct representing the https://schema.org/ParentAudience class.
type ParentAudience struct {
	ChildMaxAge *Number `json:"childmaxage"` //https://schema.org/Number
	ChildMinAge *Number `json:"childminage"` //https://schema.org/Number
}

// TravelAction is a generated struct representing the https://schema.org/TravelAction class.
type TravelAction struct {
	Distance *Distance `json:"distance"` //https://schema.org/Distance
}

// Aquarium is a generated struct representing the https://schema.org/Aquarium class.
type Aquarium struct {
}

// ArchiveOrganization is a generated struct representing the https://schema.org/ArchiveOrganization class.
type ArchiveOrganization struct {
	ArchiveHeld *ArchiveComponent `json:"archiveheld"` //https://schema.org/ArchiveComponent
}

// MedicalClinic is a generated struct representing the https://schema.org/MedicalClinic class.
type MedicalClinic struct {
	AvailableService *MedicalTherapy   `json:"availableservice"` //https://schema.org/MedicalTherapy
	MedicalSpecialty *MedicalSpecialty `json:"medicalspecialty"` //https://schema.org/MedicalSpecialty
}

// SubscribeAction is a generated struct representing the https://schema.org/SubscribeAction class.
type SubscribeAction struct {
}

// AudioObject is a generated struct representing the https://schema.org/AudioObject class.
type AudioObject struct {
	EmbeddedTextCaption *Text `json:"embeddedtextcaption"` //https://schema.org/Text
	Transcript          *Text `json:"transcript"`          //https://schema.org/Text
	Caption             *Text `json:"caption"`             //https://schema.org/Text
}

// Courthouse is a generated struct representing the https://schema.org/Courthouse class.
type Courthouse struct {
}

// MusicEvent is a generated struct representing the https://schema.org/MusicEvent class.
type MusicEvent struct {
}

// RiverBodyOfWater is a generated struct representing the https://schema.org/RiverBodyOfWater class.
type RiverBodyOfWater struct {
}

// Thesis is a generated struct representing the https://schema.org/Thesis class.
type Thesis struct {
	InSupportOf *Text `json:"insupportof"` //https://schema.org/Text
}

// ShippingDeliveryTime is a generated struct representing the https://schema.org/ShippingDeliveryTime class.
type ShippingDeliveryTime struct {
	BusinessDays *OpeningHoursSpecification `json:"businessdays"` //https://schema.org/OpeningHoursSpecification
	CutoffTime   *Time                      `json:"cutofftime"`   //https://schema.org/Time
	HandlingTime *QuantitativeValue         `json:"handlingtime"` //https://schema.org/QuantitativeValue
	TransitTime  *QuantitativeValue         `json:"transittime"`  //https://schema.org/QuantitativeValue
}

// TaxiStand is a generated struct representing the https://schema.org/TaxiStand class.
type TaxiStand struct {
}

// _3DModel is a generated struct representing the https://schema.org/3DModel class.
type _3DModel struct {
	IsResizable *Boolean `json:"isresizable"` //https://schema.org/Boolean
}

// BedDetails is a generated struct representing the https://schema.org/BedDetails class.
type BedDetails struct {
	NumberOfBeds *Number `json:"numberofbeds"` //https://schema.org/Number
	TypeOfBed    *Text   `json:"typeofbed"`    //https://schema.org/Text
}

// Campground is a generated struct representing the https://schema.org/Campground class.
type Campground struct {
}

// EmailMessage is a generated struct representing the https://schema.org/EmailMessage class.
type EmailMessage struct {
}

// EventStatusType is a generated struct representing the https://schema.org/EventStatusType class.
type EventStatusType struct {
}

// PaintAction is a generated struct representing the https://schema.org/PaintAction class.
type PaintAction struct {
}

// XPathType is a generated struct representing the https://schema.org/XPathType class.
type XPathType struct {
}

// EventVenue is a generated struct representing the https://schema.org/EventVenue class.
type EventVenue struct {
}

// MusicStore is a generated struct representing the https://schema.org/MusicStore class.
type MusicStore struct {
}

// PlayGameAction is a generated struct representing the https://schema.org/PlayGameAction class.
type PlayGameAction struct {
	GameAvailabilityType *Text `json:"gameavailabilitytype"` //https://schema.org/Text
}

// CampingPitch is a generated struct representing the https://schema.org/CampingPitch class.
type CampingPitch struct {
}

// EducationEvent is a generated struct representing the https://schema.org/EducationEvent class.
type EducationEvent struct {
	Assesses         *Text `json:"assesses"`         //https://schema.org/Text
	EducationalLevel *URL  `json:"educationallevel"` //https://schema.org/URL
	Teaches          *Text `json:"teaches"`          //https://schema.org/Text
}

// MenuSection is a generated struct representing the https://schema.org/MenuSection class.
type MenuSection struct {
	HasMenuItem    *MenuItem    `json:"hasmenuitem"`    //https://schema.org/MenuItem
	HasMenuSection *MenuSection `json:"hasmenusection"` //https://schema.org/MenuSection
}

// Taxon is a generated struct representing the https://schema.org/Taxon class.
type Taxon struct {
	TaxonRank      *URL         `json:"taxonrank"`      //https://schema.org/URL
	ChildTaxon     *URL         `json:"childtaxon"`     //https://schema.org/URL
	HasDefinedTerm *DefinedTerm `json:"hasdefinedterm"` //https://schema.org/DefinedTerm
	ParentTaxon    *URL         `json:"parenttaxon"`    //https://schema.org/URL
}

// UserPlusOnes is a generated struct representing the https://schema.org/UserPlusOnes class.
type UserPlusOnes struct {
}

// Hospital is a generated struct representing the https://schema.org/Hospital class.
type Hospital struct {
	AvailableService        *MedicalTherapy   `json:"availableservice"`        //https://schema.org/MedicalTherapy
	HealthcareReportingData *Dataset          `json:"healthcarereportingdata"` //https://schema.org/Dataset
	MedicalSpecialty        *MedicalSpecialty `json:"medicalspecialty"`        //https://schema.org/MedicalSpecialty
}

// ProfilePage is a generated struct representing the https://schema.org/ProfilePage class.
type ProfilePage struct {
}

// PsychologicalTreatment is a generated struct representing the https://schema.org/PsychologicalTreatment class.
type PsychologicalTreatment struct {
}

// Date is a generated struct representing the https://schema.org/Date class.
type Date struct {
}

// Order is a generated struct representing the https://schema.org/Order class.
type Order struct {
	AcceptedOffer      *Offer          `json:"acceptedoffer"`      //https://schema.org/Offer
	BillingAddress     *PostalAddress  `json:"billingaddress"`     //https://schema.org/PostalAddress
	ConfirmationNumber *Text           `json:"confirmationnumber"` //https://schema.org/Text
	Customer           *Person         `json:"customer"`           //https://schema.org/Person
	Discount           *Text           `json:"discount"`           //https://schema.org/Text
	DiscountCode       *Text           `json:"discountcode"`       //https://schema.org/Text
	DiscountCurrency   *Text           `json:"discountcurrency"`   //https://schema.org/Text
	IsGift             *Boolean        `json:"isgift"`             //https://schema.org/Boolean
	Merchant           *Person         `json:"merchant"`           //https://schema.org/Person
	OrderDate          *DateTime       `json:"orderdate"`          //https://schema.org/DateTime
	OrderDelivery      *ParcelDelivery `json:"orderdelivery"`      //https://schema.org/ParcelDelivery
	OrderNumber        *Text           `json:"ordernumber"`        //https://schema.org/Text
	OrderStatus        *OrderStatus    `json:"orderstatus"`        //https://schema.org/OrderStatus
	OrderedItem        *Service        `json:"ordereditem"`        //https://schema.org/Service
	PartOfInvoice      *Invoice        `json:"partofinvoice"`      //https://schema.org/Invoice
	PaymentDue         *DateTime       `json:"paymentdue"`         //https://schema.org/DateTime
	PaymentMethod      *PaymentMethod  `json:"paymentmethod"`      //https://schema.org/PaymentMethod
	PaymentMethodId    *Text           `json:"paymentmethodid"`    //https://schema.org/Text
	PaymentUrl         *URL            `json:"paymenturl"`         //https://schema.org/URL
	Broker             *Person         `json:"broker"`             //https://schema.org/Person
	PaymentDueDate     *DateTime       `json:"paymentduedate"`     //https://schema.org/DateTime
	Seller             *Person         `json:"seller"`             //https://schema.org/Person
}

// QAPage is a generated struct representing the https://schema.org/QAPage class.
type QAPage struct {
}

// TrainStation is a generated struct representing the https://schema.org/TrainStation class.
type TrainStation struct {
}

// HotelRoom is a generated struct representing the https://schema.org/HotelRoom class.
type HotelRoom struct {
	Bed       *Text              `json:"bed"`       //https://schema.org/Text
	Occupancy *QuantitativeValue `json:"occupancy"` //https://schema.org/QuantitativeValue
}

// PostalCodeRangeSpecification is a generated struct representing the https://schema.org/PostalCodeRangeSpecification class.
type PostalCodeRangeSpecification struct {
	PostalCodeBegin *Text `json:"postalcodebegin"` //https://schema.org/Text
	PostalCodeEnd   *Text `json:"postalcodeend"`   //https://schema.org/Text
}

// PriceTypeEnumeration is a generated struct representing the https://schema.org/PriceTypeEnumeration class.
type PriceTypeEnumeration struct {
}

// Seat is a generated struct representing the https://schema.org/Seat class.
type Seat struct {
	SeatNumber  *Text `json:"seatnumber"`  //https://schema.org/Text
	SeatRow     *Text `json:"seatrow"`     //https://schema.org/Text
	SeatSection *Text `json:"seatsection"` //https://schema.org/Text
	SeatingType *Text `json:"seatingtype"` //https://schema.org/Text
}

// ImageGallery is a generated struct representing the https://schema.org/ImageGallery class.
type ImageGallery struct {
}

// MedicalRiskFactor is a generated struct representing the https://schema.org/MedicalRiskFactor class.
type MedicalRiskFactor struct {
	IncreasesRiskOf *MedicalEntity `json:"increasesriskof"` //https://schema.org/MedicalEntity
}

// NightClub is a generated struct representing the https://schema.org/NightClub class.
type NightClub struct {
}

// SoftwareApplication is a generated struct representing the https://schema.org/SoftwareApplication class.
type SoftwareApplication struct {
	ApplicationCategory    *URL                 `json:"applicationcategory"`    //https://schema.org/URL
	ApplicationSubCategory *URL                 `json:"applicationsubcategory"` //https://schema.org/URL
	ApplicationSuite       *Text                `json:"applicationsuite"`       //https://schema.org/Text
	CountriesNotSupported  *Text                `json:"countriesnotsupported"`  //https://schema.org/Text
	CountriesSupported     *Text                `json:"countriessupported"`     //https://schema.org/Text
	Device                 *Text                `json:"device"`                 //https://schema.org/Text
	DownloadUrl            *URL                 `json:"downloadurl"`            //https://schema.org/URL
	FeatureList            *URL                 `json:"featurelist"`            //https://schema.org/URL
	FileSize               *Text                `json:"filesize"`               //https://schema.org/Text
	InstallUrl             *URL                 `json:"installurl"`             //https://schema.org/URL
	MemoryRequirements     *URL                 `json:"memoryrequirements"`     //https://schema.org/URL
	OperatingSystem        *Text                `json:"operatingsystem"`        //https://schema.org/Text
	Permissions            *Text                `json:"permissions"`            //https://schema.org/Text
	ProcessorRequirements  *Text                `json:"processorrequirements"`  //https://schema.org/Text
	ReleaseNotes           *URL                 `json:"releasenotes"`           //https://schema.org/URL
	Requirements           *URL                 `json:"requirements"`           //https://schema.org/URL
	Screenshot             *URL                 `json:"screenshot"`             //https://schema.org/URL
	SoftwareAddOn          *SoftwareApplication `json:"softwareaddon"`          //https://schema.org/SoftwareApplication
	SoftwareHelp           *CreativeWork        `json:"softwarehelp"`           //https://schema.org/CreativeWork
	SoftwareVersion        *Text                `json:"softwareversion"`        //https://schema.org/Text
	StorageRequirements    *URL                 `json:"storagerequirements"`    //https://schema.org/URL
	SupportingData         *DataFeed            `json:"supportingdata"`         //https://schema.org/DataFeed
	AvailableOnDevice      *Text                `json:"availableondevice"`      //https://schema.org/Text
	SoftwareRequirements   *URL                 `json:"softwarerequirements"`   //https://schema.org/URL
}

// GovernmentService is a generated struct representing the https://schema.org/GovernmentService class.
type GovernmentService struct {
	ServiceOperator *Organization `json:"serviceoperator"` //https://schema.org/Organization
	Jurisdiction    *Text         `json:"jurisdiction"`    //https://schema.org/Text
}

// VideoGallery is a generated struct representing the https://schema.org/VideoGallery class.
type VideoGallery struct {
}

// ActionStatusType is a generated struct representing the https://schema.org/ActionStatusType class.
type ActionStatusType struct {
}

// OpeningHoursSpecification is a generated struct representing the https://schema.org/OpeningHoursSpecification class.
type OpeningHoursSpecification struct {
	Closes       *Time      `json:"closes"`       //https://schema.org/Time
	DayOfWeek    *DayOfWeek `json:"dayofweek"`    //https://schema.org/DayOfWeek
	Opens        *Time      `json:"opens"`        //https://schema.org/Time
	ValidFrom    *DateTime  `json:"validfrom"`    //https://schema.org/DateTime
	ValidThrough *DateTime  `json:"validthrough"` //https://schema.org/DateTime
}

// DatedMoneySpecification is a generated struct representing the https://schema.org/DatedMoneySpecification class.
type DatedMoneySpecification struct {
	Amount    *Number   `json:"amount"`    //https://schema.org/Number
	Currency  *Text     `json:"currency"`  //https://schema.org/Text
	EndDate   *DateTime `json:"enddate"`   //https://schema.org/DateTime
	StartDate *DateTime `json:"startdate"` //https://schema.org/DateTime
}

// GroceryStore is a generated struct representing the https://schema.org/GroceryStore class.
type GroceryStore struct {
}

// MusicComposition is a generated struct representing the https://schema.org/MusicComposition class.
type MusicComposition struct {
	Composer             *Person           `json:"composer"`             //https://schema.org/Person
	FirstPerformance     *Event            `json:"firstperformance"`     //https://schema.org/Event
	IncludedComposition  *MusicComposition `json:"includedcomposition"`  //https://schema.org/MusicComposition
	IswcCode             *Text             `json:"iswccode"`             //https://schema.org/Text
	Lyricist             *Person           `json:"lyricist"`             //https://schema.org/Person
	Lyrics               *CreativeWork     `json:"lyrics"`               //https://schema.org/CreativeWork
	MusicArrangement     *MusicComposition `json:"musicarrangement"`     //https://schema.org/MusicComposition
	MusicCompositionForm *Text             `json:"musiccompositionform"` //https://schema.org/Text
	MusicalKey           *Text             `json:"musicalkey"`           //https://schema.org/Text
	RecordedAs           *MusicRecording   `json:"recordedas"`           //https://schema.org/MusicRecording
}

// GovernmentOrganization is a generated struct representing the https://schema.org/GovernmentOrganization class.
type GovernmentOrganization struct {
}

// HardwareStore is a generated struct representing the https://schema.org/HardwareStore class.
type HardwareStore struct {
}

// PodcastSeason is a generated struct representing the https://schema.org/PodcastSeason class.
type PodcastSeason struct {
}

// EducationalAudience is a generated struct representing the https://schema.org/EducationalAudience class.
type EducationalAudience struct {
	EducationalRole *Text `json:"educationalrole"` //https://schema.org/Text
}

// MedicalObservationalStudyDesign is a generated struct representing the https://schema.org/MedicalObservationalStudyDesign class.
type MedicalObservationalStudyDesign struct {
}

// StadiumOrArena is a generated struct representing the https://schema.org/StadiumOrArena class.
type StadiumOrArena struct {
}

// TVEpisode is a generated struct representing the https://schema.org/TVEpisode class.
type TVEpisode struct {
	CountryOfOrigin  *Country  `json:"countryoforigin"`  //https://schema.org/Country
	PartOfTVSeries   *TVSeries `json:"partoftvseries"`   //https://schema.org/TVSeries
	SubtitleLanguage *Text     `json:"subtitlelanguage"` //https://schema.org/Text
	TitleEIDR        *URL      `json:"titleeidr"`        //https://schema.org/URL
}

// DepartAction is a generated struct representing the https://schema.org/DepartAction class.
type DepartAction struct {
}

// ExchangeRateSpecification is a generated struct representing the https://schema.org/ExchangeRateSpecification class.
type ExchangeRateSpecification struct {
	Currency            *Text                   `json:"currency"`            //https://schema.org/Text
	CurrentExchangeRate *UnitPriceSpecification `json:"currentexchangerate"` //https://schema.org/UnitPriceSpecification
	ExchangeRateSpread  *Number                 `json:"exchangeratespread"`  //https://schema.org/Number
}

// HobbyShop is a generated struct representing the https://schema.org/HobbyShop class.
type HobbyShop struct {
}

// NailSalon is a generated struct representing the https://schema.org/NailSalon class.
type NailSalon struct {
}

// RadioStation is a generated struct representing the https://schema.org/RadioStation class.
type RadioStation struct {
}

// TextObject is a generated struct representing the https://schema.org/TextObject class.
type TextObject struct {
}

// CheckOutAction is a generated struct representing the https://schema.org/CheckOutAction class.
type CheckOutAction struct {
}

// AlignmentObject is a generated struct representing the https://schema.org/AlignmentObject class.
type AlignmentObject struct {
	AlignmentType        *Text `json:"alignmenttype"`        //https://schema.org/Text
	EducationalFramework *Text `json:"educationalframework"` //https://schema.org/Text
	TargetDescription    *Text `json:"targetdescription"`    //https://schema.org/Text
	TargetName           *Text `json:"targetname"`           //https://schema.org/Text
	TargetUrl            *URL  `json:"targeturl"`            //https://schema.org/URL
}

// AnatomicalSystem is a generated struct representing the https://schema.org/AnatomicalSystem class.
type AnatomicalSystem struct {
	AssociatedPathophysiology *Text                `json:"associatedpathophysiology"` //https://schema.org/Text
	ComprisedOf               *AnatomicalSystem    `json:"comprisedof"`               //https://schema.org/AnatomicalSystem
	RelatedCondition          *MedicalCondition    `json:"relatedcondition"`          //https://schema.org/MedicalCondition
	RelatedStructure          *AnatomicalStructure `json:"relatedstructure"`          //https://schema.org/AnatomicalStructure
	RelatedTherapy            *MedicalTherapy      `json:"relatedtherapy"`            //https://schema.org/MedicalTherapy
}

// JoinAction is a generated struct representing the https://schema.org/JoinAction class.
type JoinAction struct {
	Event *Event `json:"event"` //https://schema.org/Event
}

// Ticket is a generated struct representing the https://schema.org/Ticket class.
type Ticket struct {
	DateIssued    *DateTime     `json:"dateissued"`    //https://schema.org/DateTime
	IssuedBy      *Organization `json:"issuedby"`      //https://schema.org/Organization
	PriceCurrency *Text         `json:"pricecurrency"` //https://schema.org/Text
	TicketNumber  *Text         `json:"ticketnumber"`  //https://schema.org/Text
	TicketToken   *URL          `json:"tickettoken"`   //https://schema.org/URL
	TicketedSeat  *Seat         `json:"ticketedseat"`  //https://schema.org/Seat
	TotalPrice    *Text         `json:"totalprice"`    //https://schema.org/Text
	UnderName     *Person       `json:"undername"`     //https://schema.org/Person
}

// Flight is a generated struct representing the https://schema.org/Flight class.
type Flight struct {
	Aircraft                *Vehicle            `json:"aircraft"`                //https://schema.org/Vehicle
	ArrivalAirport          *Airport            `json:"arrivalairport"`          //https://schema.org/Airport
	ArrivalGate             *Text               `json:"arrivalgate"`             //https://schema.org/Text
	ArrivalTerminal         *Text               `json:"arrivalterminal"`         //https://schema.org/Text
	BoardingPolicy          *BoardingPolicyType `json:"boardingpolicy"`          //https://schema.org/BoardingPolicyType
	Carrier                 *Organization       `json:"carrier"`                 //https://schema.org/Organization
	DepartureAirport        *Airport            `json:"departureairport"`        //https://schema.org/Airport
	DepartureGate           *Text               `json:"departuregate"`           //https://schema.org/Text
	DepartureTerminal       *Text               `json:"departureterminal"`       //https://schema.org/Text
	EstimatedFlightDuration *Text               `json:"estimatedflightduration"` //https://schema.org/Text
	FlightDistance          *Text               `json:"flightdistance"`          //https://schema.org/Text
	FlightNumber            *Text               `json:"flightnumber"`            //https://schema.org/Text
	MealService             *Text               `json:"mealservice"`             //https://schema.org/Text
	WebCheckinTime          *DateTime           `json:"webcheckintime"`          //https://schema.org/DateTime
	Seller                  *Person             `json:"seller"`                  //https://schema.org/Person
}

// GameServer is a generated struct representing the https://schema.org/GameServer class.
type GameServer struct {
	PlayersOnline *Integer          `json:"playersonline"` //https://schema.org/Integer
	ServerStatus  *GameServerStatus `json:"serverstatus"`  //https://schema.org/GameServerStatus
	Game          *VideoGame        `json:"game"`          //https://schema.org/VideoGame
}

// Syllabus is a generated struct representing the https://schema.org/Syllabus class.
type Syllabus struct {
}

// URL is a generated struct representing the https://schema.org/URL class.
type URL struct {
}

// WorkBasedProgram is a generated struct representing the https://schema.org/WorkBasedProgram class.
type WorkBasedProgram struct {
	OccupationalCategory *Text                       `json:"occupationalcategory"` //https://schema.org/Text
	TrainingSalary       *MonetaryAmountDistribution `json:"trainingsalary"`       //https://schema.org/MonetaryAmountDistribution
}

// Atlas is a generated struct representing the https://schema.org/Atlas class.
type Atlas struct {
}

// Episode is a generated struct representing the https://schema.org/Episode class.
type Episode struct {
	Actors            *Person             `json:"actors"`            //https://schema.org/Person
	Directors         *Person             `json:"directors"`         //https://schema.org/Person
	EpisodeNumber     *Text               `json:"episodenumber"`     //https://schema.org/Text
	MusicBy           *Person             `json:"musicby"`           //https://schema.org/Person
	PartOfSeason      *CreativeWorkSeason `json:"partofseason"`      //https://schema.org/CreativeWorkSeason
	ProductionCompany *Organization       `json:"productioncompany"` //https://schema.org/Organization
	Trailer           *VideoObject        `json:"trailer"`           //https://schema.org/VideoObject
	Director          *Person             `json:"director"`          //https://schema.org/Person
	Duration          *Duration           `json:"duration"`          //https://schema.org/Duration
	PartOfSeries      *CreativeWorkSeries `json:"partofseries"`      //https://schema.org/CreativeWorkSeries
	Actor             *Person             `json:"actor"`             //https://schema.org/Person
}

// PhysicalExam is a generated struct representing the https://schema.org/PhysicalExam class.
type PhysicalExam struct {
}

// WebPage is a generated struct representing the https://schema.org/WebPage class.
type WebPage struct {
	Breadcrumb         *Text           `json:"breadcrumb"`         //https://schema.org/Text
	LastReviewed       *Date           `json:"lastreviewed"`       //https://schema.org/Date
	PrimaryImageOfPage *ImageObject    `json:"primaryimageofpage"` //https://schema.org/ImageObject
	RelatedLink        *URL            `json:"relatedlink"`        //https://schema.org/URL
	ReviewedBy         *Person         `json:"reviewedby"`         //https://schema.org/Person
	SignificantLinks   *URL            `json:"significantlinks"`   //https://schema.org/URL
	Speakable          *URL            `json:"speakable"`          //https://schema.org/URL
	Specialty          *Specialty      `json:"specialty"`          //https://schema.org/Specialty
	MainContentOfPage  *WebPageElement `json:"maincontentofpage"`  //https://schema.org/WebPageElement
	SignificantLink    *URL            `json:"significantlink"`    //https://schema.org/URL
}

// UserCheckins is a generated struct representing the https://schema.org/UserCheckins class.
type UserCheckins struct {
}

// ControlAction is a generated struct representing the https://schema.org/ControlAction class.
type ControlAction struct {
}

// CookAction is a generated struct representing the https://schema.org/CookAction class.
type CookAction struct {
	FoodEstablishment *Place     `json:"foodestablishment"` //https://schema.org/Place
	FoodEvent         *FoodEvent `json:"foodevent"`         //https://schema.org/FoodEvent
	Recipe            *Recipe    `json:"recipe"`            //https://schema.org/Recipe
}

// EventAttendanceModeEnumeration is a generated struct representing the https://schema.org/EventAttendanceModeEnumeration class.
type EventAttendanceModeEnumeration struct {
}

// MotorcycleRepair is a generated struct representing the https://schema.org/MotorcycleRepair class.
type MotorcycleRepair struct {
}

// RealEstateListing is a generated struct representing the https://schema.org/RealEstateListing class.
type RealEstateListing struct {
	DatePosted  *DateTime          `json:"dateposted"`  //https://schema.org/DateTime
	LeaseLength *QuantitativeValue `json:"leaselength"` //https://schema.org/QuantitativeValue
}

// StatusEnumeration is a generated struct representing the https://schema.org/StatusEnumeration class.
type StatusEnumeration struct {
}

// AutoRepair is a generated struct representing the https://schema.org/AutoRepair class.
type AutoRepair struct {
}

// DeliveryTimeSettings is a generated struct representing the https://schema.org/DeliveryTimeSettings class.
type DeliveryTimeSettings struct {
	DeliveryTime         *ShippingDeliveryTime `json:"deliverytime"`         //https://schema.org/ShippingDeliveryTime
	IsUnlabelledFallback *Boolean              `json:"isunlabelledfallback"` //https://schema.org/Boolean
	ShippingDestination  *DefinedRegion        `json:"shippingdestination"`  //https://schema.org/DefinedRegion
	TransitTimeLabel     *Text                 `json:"transittimelabel"`     //https://schema.org/Text
}

// BioChemEntity is a generated struct representing the https://schema.org/BioChemEntity class.
type BioChemEntity struct {
	AssociatedDisease              *URL           `json:"associateddisease"`              //https://schema.org/URL
	BioChemInteraction             *BioChemEntity `json:"biocheminteraction"`             //https://schema.org/BioChemEntity
	BioChemSimilarity              *BioChemEntity `json:"biochemsimilarity"`              //https://schema.org/BioChemEntity
	BiologicalRole                 *DefinedTerm   `json:"biologicalrole"`                 //https://schema.org/DefinedTerm
	HasMolecularFunction           *URL           `json:"hasmolecularfunction"`           //https://schema.org/URL
	IsInvolvedInBiologicalProcess  *URL           `json:"isinvolvedinbiologicalprocess"`  //https://schema.org/URL
	IsLocatedInSubcellularLocation *URL           `json:"islocatedinsubcellularlocation"` //https://schema.org/URL
	TaxonomicRange                 *URL           `json:"taxonomicrange"`                 //https://schema.org/URL
	Funding                        *Grant         `json:"funding"`                        //https://schema.org/Grant
	HasBioChemEntityPart           *BioChemEntity `json:"hasbiochementitypart"`           //https://schema.org/BioChemEntity
	IsEncodedByBioChemEntity       *Gene          `json:"isencodedbybiochementity"`       //https://schema.org/Gene
	IsPartOfBioChemEntity          *BioChemEntity `json:"ispartofbiochementity"`          //https://schema.org/BioChemEntity
	HasRepresentation              *URL           `json:"hasrepresentation"`              //https://schema.org/URL
}

// LodgingBusiness is a generated struct representing the https://schema.org/LodgingBusiness class.
type LodgingBusiness struct {
	AmenityFeature    *LocationFeatureSpecification `json:"amenityfeature"`    //https://schema.org/LocationFeatureSpecification
	AvailableLanguage *Text                         `json:"availablelanguage"` //https://schema.org/Text
	CheckinTime       *Time                         `json:"checkintime"`       //https://schema.org/Time
	CheckoutTime      *Time                         `json:"checkouttime"`      //https://schema.org/Time
	NumberOfRooms     *QuantitativeValue            `json:"numberofrooms"`     //https://schema.org/QuantitativeValue
	PetsAllowed       *Text                         `json:"petsallowed"`       //https://schema.org/Text
	StarRating        *Rating                       `json:"starrating"`        //https://schema.org/Rating
	Audience          *Audience                     `json:"audience"`          //https://schema.org/Audience
}

// SurgicalProcedure is a generated struct representing the https://schema.org/SurgicalProcedure class.
type SurgicalProcedure struct {
}

// AutoPartsStore is a generated struct representing the https://schema.org/AutoPartsStore class.
type AutoPartsStore struct {
}

// Muscle is a generated struct representing the https://schema.org/Muscle class.
type Muscle struct {
	Antagonist   *Muscle              `json:"antagonist"`   //https://schema.org/Muscle
	BloodSupply  *Vessel              `json:"bloodsupply"`  //https://schema.org/Vessel
	Insertion    *AnatomicalStructure `json:"insertion"`    //https://schema.org/AnatomicalStructure
	MuscleAction *Text                `json:"muscleaction"` //https://schema.org/Text
	Nerve        *Nerve               `json:"nerve"`        //https://schema.org/Nerve
}

// Recommendation is a generated struct representing the https://schema.org/Recommendation class.
type Recommendation struct {
	Category *URL `json:"category"` //https://schema.org/URL
}

// AmpStory is a generated struct representing the https://schema.org/AmpStory class.
type AmpStory struct {
}

// AutoWash is a generated struct representing the https://schema.org/AutoWash class.
type AutoWash struct {
}

// CheckInAction is a generated struct representing the https://schema.org/CheckInAction class.
type CheckInAction struct {
}

// ConvenienceStore is a generated struct representing the https://schema.org/ConvenienceStore class.
type ConvenienceStore struct {
}

// HealthInsurancePlan is a generated struct representing the https://schema.org/HealthInsurancePlan class.
type HealthInsurancePlan struct {
	BenefitsSummaryUrl          *URL                 `json:"benefitssummaryurl"`          //https://schema.org/URL
	HealthPlanDrugOption        *Text                `json:"healthplandrugoption"`        //https://schema.org/Text
	HealthPlanDrugTier          *Text                `json:"healthplandrugtier"`          //https://schema.org/Text
	HealthPlanId                *Text                `json:"healthplanid"`                //https://schema.org/Text
	HealthPlanMarketingUrl      *URL                 `json:"healthplanmarketingurl"`      //https://schema.org/URL
	IncludesHealthPlanFormulary *HealthPlanFormulary `json:"includeshealthplanformulary"` //https://schema.org/HealthPlanFormulary
	IncludesHealthPlanNetwork   *HealthPlanNetwork   `json:"includeshealthplannetwork"`   //https://schema.org/HealthPlanNetwork
	UsesHealthPlanIdStandard    *URL                 `json:"useshealthplanidstandard"`    //https://schema.org/URL
	ContactPoint                *ContactPoint        `json:"contactpoint"`                //https://schema.org/ContactPoint
}

// Play is a generated struct representing the https://schema.org/Play class.
type Play struct {
}

// CreativeWork is a generated struct representing the https://schema.org/CreativeWork class.
type CreativeWork struct {
	Abstract             *Text               `json:"abstract"`             //https://schema.org/Text
	AccessMode           *Text               `json:"accessmode"`           //https://schema.org/Text
	AccessModeSufficient *ItemList           `json:"accessmodesufficient"` //https://schema.org/ItemList
	AccessibilityAPI     *Text               `json:"accessibilityapi"`     //https://schema.org/Text
	AccessibilityControl *Text               `json:"accessibilitycontrol"` //https://schema.org/Text
	AccessibilityFeature *Text               `json:"accessibilityfeature"` //https://schema.org/Text
	AccessibilityHazard  *Text               `json:"accessibilityhazard"`  //https://schema.org/Text
	AccessibilitySummary *Text               `json:"accessibilitysummary"` //https://schema.org/Text
	AccountablePerson    *Person             `json:"accountableperson"`    //https://schema.org/Person
	AcquireLicensePage   *URL                `json:"acquirelicensepage"`   //https://schema.org/URL
	AggregateRating      *AggregateRating    `json:"aggregaterating"`      //https://schema.org/AggregateRating
	AlternativeHeadline  *Text               `json:"alternativeheadline"`  //https://schema.org/Text
	ArchivedAt           *WebPage            `json:"archivedat"`           //https://schema.org/WebPage
	Assesses             *Text               `json:"assesses"`             //https://schema.org/Text
	AssociatedMedia      *MediaObject        `json:"associatedmedia"`      //https://schema.org/MediaObject
	Audio                *MusicRecording     `json:"audio"`                //https://schema.org/MusicRecording
	Author               *Person             `json:"author"`               //https://schema.org/Person
	Awards               *Text               `json:"awards"`               //https://schema.org/Text
	Character            *Person             `json:"character"`            //https://schema.org/Person
	Citation             *Text               `json:"citation"`             //https://schema.org/Text
	Comment              *Comment            `json:"comment"`              //https://schema.org/Comment
	CommentCount         *Integer            `json:"commentcount"`         //https://schema.org/Integer
	ConditionsOfAccess   *Text               `json:"conditionsofaccess"`   //https://schema.org/Text
	ContentRating        *Text               `json:"contentrating"`        //https://schema.org/Text
	ContentReferenceTime *DateTime           `json:"contentreferencetime"` //https://schema.org/DateTime
	Contributor          *Person             `json:"contributor"`          //https://schema.org/Person
	CopyrightHolder      *Person             `json:"copyrightholder"`      //https://schema.org/Person
	CopyrightNotice      *Text               `json:"copyrightnotice"`      //https://schema.org/Text
	CopyrightYear        *Number             `json:"copyrightyear"`        //https://schema.org/Number
	Correction           *URL                `json:"correction"`           //https://schema.org/URL
	CountryOfOrigin      *Country            `json:"countryoforigin"`      //https://schema.org/Country
	CreativeWorkStatus   *Text               `json:"creativeworkstatus"`   //https://schema.org/Text
	CreditText           *Text               `json:"credittext"`           //https://schema.org/Text
	DateModified         *DateTime           `json:"datemodified"`         //https://schema.org/DateTime
	DatePublished        *DateTime           `json:"datepublished"`        //https://schema.org/DateTime
	DiscussionUrl        *URL                `json:"discussionurl"`        //https://schema.org/URL
	EditEIDR             *URL                `json:"editeidr"`             //https://schema.org/URL
	Editor               *Person             `json:"editor"`               //https://schema.org/Person
	EducationalAlignment *AlignmentObject    `json:"educationalalignment"` //https://schema.org/AlignmentObject
	EducationalLevel     *URL                `json:"educationallevel"`     //https://schema.org/URL
	EducationalUse       *Text               `json:"educationaluse"`       //https://schema.org/Text
	Encodings            *MediaObject        `json:"encodings"`            //https://schema.org/MediaObject
	Expires              *DateTime           `json:"expires"`              //https://schema.org/DateTime
	FileFormat           *URL                `json:"fileformat"`           //https://schema.org/URL
	Funder               *Person             `json:"funder"`               //https://schema.org/Person
	Headline             *Text               `json:"headline"`             //https://schema.org/Text
	InteractivityType    *Text               `json:"interactivitytype"`    //https://schema.org/Text
	InterpretedAsClaim   *Claim              `json:"interpretedasclaim"`   //https://schema.org/Claim
	IsBasedOnUrl         *URL                `json:"isbasedonurl"`         //https://schema.org/URL
	IsFamilyFriendly     *Boolean            `json:"isfamilyfriendly"`     //https://schema.org/Boolean
	Keywords             *URL                `json:"keywords"`             //https://schema.org/URL
	LearningResourceType *Text               `json:"learningresourcetype"` //https://schema.org/Text
	License              *URL                `json:"license"`              //https://schema.org/URL
	LocationCreated      *Place              `json:"locationcreated"`      //https://schema.org/Place
	Maintainer           *Person             `json:"maintainer"`           //https://schema.org/Person
	MaterialExtent       *Text               `json:"materialextent"`       //https://schema.org/Text
	Mentions             *Thing              `json:"mentions"`             //https://schema.org/Thing
	Pattern              *Text               `json:"pattern"`              //https://schema.org/Text
	Producer             *Person             `json:"producer"`             //https://schema.org/Person
	Publication          *PublicationEvent   `json:"publication"`          //https://schema.org/PublicationEvent
	Publisher            *Person             `json:"publisher"`            //https://schema.org/Person
	PublisherImprint     *Organization       `json:"publisherimprint"`     //https://schema.org/Organization
	ReleasedEvent        *PublicationEvent   `json:"releasedevent"`        //https://schema.org/PublicationEvent
	Reviews              *Review             `json:"reviews"`              //https://schema.org/Review
	SchemaVersion        *URL                `json:"schemaversion"`        //https://schema.org/URL
	SdDatePublished      *Date               `json:"sddatepublished"`      //https://schema.org/Date
	SdLicense            *URL                `json:"sdlicense"`            //https://schema.org/URL
	SdPublisher          *Person             `json:"sdpublisher"`          //https://schema.org/Person
	Size                 *Text               `json:"size"`                 //https://schema.org/Text
	SourceOrganization   *Organization       `json:"sourceorganization"`   //https://schema.org/Organization
	Spatial              *Place              `json:"spatial"`              //https://schema.org/Place
	Teaches              *Text               `json:"teaches"`              //https://schema.org/Text
	Temporal             *Text               `json:"temporal"`             //https://schema.org/Text
	Text                 *Text               `json:"text"`                 //https://schema.org/Text
	Thumbnail            *ImageObject        `json:"thumbnail"`            //https://schema.org/ImageObject
	ThumbnailUrl         *URL                `json:"thumbnailurl"`         //https://schema.org/URL
	TimeRequired         *Duration           `json:"timerequired"`         //https://schema.org/Duration
	Translator           *Person             `json:"translator"`           //https://schema.org/Person
	TypicalAgeRange      *Text               `json:"typicalagerange"`      //https://schema.org/Text
	Version              *Text               `json:"version"`              //https://schema.org/Text
	Video                *VideoObject        `json:"video"`                //https://schema.org/VideoObject
	Audience             *Audience           `json:"audience"`             //https://schema.org/Audience
	Award                *Text               `json:"award"`                //https://schema.org/Text
	ContentLocation      *Place              `json:"contentlocation"`      //https://schema.org/Place
	Creator              *Person             `json:"creator"`              //https://schema.org/Person
	DateCreated          *DateTime           `json:"datecreated"`          //https://schema.org/DateTime
	EncodingFormat       *URL                `json:"encodingformat"`       //https://schema.org/URL
	ExampleOfWork        *CreativeWork       `json:"exampleofwork"`        //https://schema.org/CreativeWork
	Funding              *Grant              `json:"funding"`              //https://schema.org/Grant
	Genre                *URL                `json:"genre"`                //https://schema.org/URL
	InLanguage           *Text               `json:"inlanguage"`           //https://schema.org/Text
	InteractionStatistic *InteractionCounter `json:"interactionstatistic"` //https://schema.org/InteractionCounter
	IsAccessibleForFree  *Boolean            `json:"isaccessibleforfree"`  //https://schema.org/Boolean
	IsBasedOn            *URL                `json:"isbasedon"`            //https://schema.org/URL
	MainEntity           *Thing              `json:"mainentity"`           //https://schema.org/Thing
	Offers               *Offer              `json:"offers"`               //https://schema.org/Offer
	Provider             *Person             `json:"provider"`             //https://schema.org/Person
	RecordedAt           *Event              `json:"recordedat"`           //https://schema.org/Event
	Review               *Review             `json:"review"`               //https://schema.org/Review
	Sponsor              *Person             `json:"sponsor"`              //https://schema.org/Person
	TemporalCoverage     *URL                `json:"temporalcoverage"`     //https://schema.org/URL
	TranslationOfWork    *CreativeWork       `json:"translationofwork"`    //https://schema.org/CreativeWork
	UsageInfo            *URL                `json:"usageinfo"`            //https://schema.org/URL
	WorkTranslation      *CreativeWork       `json:"worktranslation"`      //https://schema.org/CreativeWork
	About                *Thing              `json:"about"`                //https://schema.org/Thing
	Encoding             *MediaObject        `json:"encoding"`             //https://schema.org/MediaObject
	Material             *URL                `json:"material"`             //https://schema.org/URL
	SpatialCoverage      *Place              `json:"spatialcoverage"`      //https://schema.org/Place
	WorkExample          *CreativeWork       `json:"workexample"`          //https://schema.org/CreativeWork
	HasPart              *CreativeWork       `json:"haspart"`              //https://schema.org/CreativeWork
	Position             *Text               `json:"position"`             //https://schema.org/Text
	IsPartOf             *URL                `json:"ispartof"`             //https://schema.org/URL
	PublishingPrinciples *URL                `json:"publishingprinciples"` //https://schema.org/URL
}

// Mass is a generated struct representing the https://schema.org/Mass class.
type Mass struct {
}

// MedicalStudy is a generated struct representing the https://schema.org/MedicalStudy class.
type MedicalStudy struct {
	HealthCondition *MedicalCondition   `json:"healthcondition"` //https://schema.org/MedicalCondition
	Status          *Text               `json:"status"`          //https://schema.org/Text
	StudyLocation   *AdministrativeArea `json:"studylocation"`   //https://schema.org/AdministrativeArea
	StudySubject    *MedicalEntity      `json:"studysubject"`    //https://schema.org/MedicalEntity
	Sponsor         *Person             `json:"sponsor"`         //https://schema.org/Person
}

// Reservation is a generated struct representing the https://schema.org/Reservation class.
type Reservation struct {
	BookingAgent          *Person                `json:"bookingagent"`          //https://schema.org/Person
	BookingTime           *DateTime              `json:"bookingtime"`           //https://schema.org/DateTime
	ModifiedTime          *DateTime              `json:"modifiedtime"`          //https://schema.org/DateTime
	PriceCurrency         *Text                  `json:"pricecurrency"`         //https://schema.org/Text
	ProgramMembershipUsed *ProgramMembership     `json:"programmembershipused"` //https://schema.org/ProgramMembership
	ReservationFor        *Thing                 `json:"reservationfor"`        //https://schema.org/Thing
	ReservationId         *Text                  `json:"reservationid"`         //https://schema.org/Text
	ReservationStatus     *ReservationStatusType `json:"reservationstatus"`     //https://schema.org/ReservationStatusType
	ReservedTicket        *Ticket                `json:"reservedticket"`        //https://schema.org/Ticket
	TotalPrice            *Text                  `json:"totalprice"`            //https://schema.org/Text
	UnderName             *Person                `json:"undername"`             //https://schema.org/Person
	Broker                *Person                `json:"broker"`                //https://schema.org/Person
	Provider              *Person                `json:"provider"`              //https://schema.org/Person
}

// CollegeOrUniversity is a generated struct representing the https://schema.org/CollegeOrUniversity class.
type CollegeOrUniversity struct {
}

// Offer is a generated struct representing the https://schema.org/Offer class.
type Offer struct {
	AcceptedPaymentMethod     *PaymentMethod            `json:"acceptedpaymentmethod"`     //https://schema.org/PaymentMethod
	AddOn                     *Offer                    `json:"addon"`                     //https://schema.org/Offer
	AdvanceBookingRequirement *QuantitativeValue        `json:"advancebookingrequirement"` //https://schema.org/QuantitativeValue
	AggregateRating           *AggregateRating          `json:"aggregaterating"`           //https://schema.org/AggregateRating
	Asin                      *URL                      `json:"asin"`                      //https://schema.org/URL
	Availability              *ItemAvailability         `json:"availability"`              //https://schema.org/ItemAvailability
	AvailabilityEnds          *Time                     `json:"availabilityends"`          //https://schema.org/Time
	AvailabilityStarts        *Time                     `json:"availabilitystarts"`        //https://schema.org/Time
	AvailableAtOrFrom         *Place                    `json:"availableatorfrom"`         //https://schema.org/Place
	AvailableDeliveryMethod   *DeliveryMethod           `json:"availabledeliverymethod"`   //https://schema.org/DeliveryMethod
	BusinessFunction          *BusinessFunction         `json:"businessfunction"`          //https://schema.org/BusinessFunction
	CheckoutPageURLTemplate   *Text                     `json:"checkoutpageurltemplate"`   //https://schema.org/Text
	DeliveryLeadTime          *QuantitativeValue        `json:"deliveryleadtime"`          //https://schema.org/QuantitativeValue
	EligibleCustomerType      *BusinessEntityType       `json:"eligiblecustomertype"`      //https://schema.org/BusinessEntityType
	EligibleDuration          *QuantitativeValue        `json:"eligibleduration"`          //https://schema.org/QuantitativeValue
	EligibleQuantity          *QuantitativeValue        `json:"eligiblequantity"`          //https://schema.org/QuantitativeValue
	EligibleRegion            *Text                     `json:"eligibleregion"`            //https://schema.org/Text
	EligibleTransactionVolume *PriceSpecification       `json:"eligibletransactionvolume"` //https://schema.org/PriceSpecification
	Gtin12                    *Text                     `json:"gtin12"`                    //https://schema.org/Text
	Gtin13                    *Text                     `json:"gtin13"`                    //https://schema.org/Text
	Gtin14                    *Text                     `json:"gtin14"`                    //https://schema.org/Text
	Gtin8                     *Text                     `json:"gtin8"`                     //https://schema.org/Text
	HasAdultConsideration     *AdultOrientedEnumeration `json:"hasadultconsideration"`     //https://schema.org/AdultOrientedEnumeration
	HasMeasurement            *QuantitativeValue        `json:"hasmeasurement"`            //https://schema.org/QuantitativeValue
	HasMerchantReturnPolicy   *MerchantReturnPolicy     `json:"hasmerchantreturnpolicy"`   //https://schema.org/MerchantReturnPolicy
	IncludesObject            *TypeAndQuantityNode      `json:"includesobject"`            //https://schema.org/TypeAndQuantityNode
	IneligibleRegion          *Text                     `json:"ineligibleregion"`          //https://schema.org/Text
	InventoryLevel            *QuantitativeValue        `json:"inventorylevel"`            //https://schema.org/QuantitativeValue
	IsFamilyFriendly          *Boolean                  `json:"isfamilyfriendly"`          //https://schema.org/Boolean
	ItemCondition             *OfferItemCondition       `json:"itemcondition"`             //https://schema.org/OfferItemCondition
	LeaseLength               *QuantitativeValue        `json:"leaselength"`               //https://schema.org/QuantitativeValue
	MobileUrl                 *Text                     `json:"mobileurl"`                 //https://schema.org/Text
	Mpn                       *Text                     `json:"mpn"`                       //https://schema.org/Text
	Price                     *Text                     `json:"price"`                     //https://schema.org/Text
	PriceCurrency             *Text                     `json:"pricecurrency"`             //https://schema.org/Text
	PriceSpecification        *PriceSpecification       `json:"pricespecification"`        //https://schema.org/PriceSpecification
	PriceValidUntil           *Date                     `json:"pricevaliduntil"`           //https://schema.org/Date
	Reviews                   *Review                   `json:"reviews"`                   //https://schema.org/Review
	ShippingDetails           *OfferShippingDetails     `json:"shippingdetails"`           //https://schema.org/OfferShippingDetails
	Sku                       *Text                     `json:"sku"`                       //https://schema.org/Text
	ValidFrom                 *DateTime                 `json:"validfrom"`                 //https://schema.org/DateTime
	ValidThrough              *DateTime                 `json:"validthrough"`              //https://schema.org/DateTime
	Category                  *URL                      `json:"category"`                  //https://schema.org/URL
	ItemOffered               *Trip                     `json:"itemoffered"`               //https://schema.org/Trip
	OfferedBy                 *Person                   `json:"offeredby"`                 //https://schema.org/Person
	Review                    *Review                   `json:"review"`                    //https://schema.org/Review
	SerialNumber              *Text                     `json:"serialnumber"`              //https://schema.org/Text
	Warranty                  *WarrantyPromise          `json:"warranty"`                  //https://schema.org/WarrantyPromise
	Seller                    *Person                   `json:"seller"`                    //https://schema.org/Person
	AreaServed                *Text                     `json:"areaserved"`                //https://schema.org/Text
	Gtin                      *URL                      `json:"gtin"`                      //https://schema.org/URL
}

// BankOrCreditUnion is a generated struct representing the https://schema.org/BankOrCreditUnion class.
type BankOrCreditUnion struct {
}

// Permit is a generated struct representing the https://schema.org/Permit class.
type Permit struct {
	IssuedBy       *Organization       `json:"issuedby"`       //https://schema.org/Organization
	IssuedThrough  *Service            `json:"issuedthrough"`  //https://schema.org/Service
	PermitAudience *Audience           `json:"permitaudience"` //https://schema.org/Audience
	ValidFor       *Duration           `json:"validfor"`       //https://schema.org/Duration
	ValidFrom      *DateTime           `json:"validfrom"`      //https://schema.org/DateTime
	ValidIn        *AdministrativeArea `json:"validin"`        //https://schema.org/AdministrativeArea
	ValidUntil     *Date               `json:"validuntil"`     //https://schema.org/Date
}

// UserBlocks is a generated struct representing the https://schema.org/UserBlocks class.
type UserBlocks struct {
}

// Audience is a generated struct representing the https://schema.org/Audience class.
type Audience struct {
	AudienceType   *Text               `json:"audiencetype"`   //https://schema.org/Text
	GeographicArea *AdministrativeArea `json:"geographicarea"` //https://schema.org/AdministrativeArea
}

// CategoryCodeSet is a generated struct representing the https://schema.org/CategoryCodeSet class.
type CategoryCodeSet struct {
	HasCategoryCode *CategoryCode `json:"hascategorycode"` //https://schema.org/CategoryCode
}

// MusicRelease is a generated struct representing the https://schema.org/MusicRelease class.
type MusicRelease struct {
	CatalogNumber      *Text                   `json:"catalognumber"`      //https://schema.org/Text
	CreditedTo         *Person                 `json:"creditedto"`         //https://schema.org/Person
	MusicReleaseFormat *MusicReleaseFormatType `json:"musicreleaseformat"` //https://schema.org/MusicReleaseFormatType
	RecordLabel        *Organization           `json:"recordlabel"`        //https://schema.org/Organization
	Duration           *Duration               `json:"duration"`           //https://schema.org/Duration
	ReleaseOf          *MusicAlbum             `json:"releaseof"`          //https://schema.org/MusicAlbum
}

// ReturnLabelSourceEnumeration is a generated struct representing the https://schema.org/ReturnLabelSourceEnumeration class.
type ReturnLabelSourceEnumeration struct {
}

// WarrantyPromise is a generated struct representing the https://schema.org/WarrantyPromise class.
type WarrantyPromise struct {
	DurationOfWarranty *QuantitativeValue `json:"durationofwarranty"` //https://schema.org/QuantitativeValue
	WarrantyScope      *WarrantyScope     `json:"warrantyscope"`      //https://schema.org/WarrantyScope
}

// Apartment is a generated struct representing the https://schema.org/Apartment class.
type Apartment struct {
	NumberOfRooms *QuantitativeValue `json:"numberofrooms"` //https://schema.org/QuantitativeValue
	Occupancy     *QuantitativeValue `json:"occupancy"`     //https://schema.org/QuantitativeValue
}

// HealthAndBeautyBusiness is a generated struct representing the https://schema.org/HealthAndBeautyBusiness class.
type HealthAndBeautyBusiness struct {
}

// MeasurementTypeEnumeration is a generated struct representing the https://schema.org/MeasurementTypeEnumeration class.
type MeasurementTypeEnumeration struct {
}

// UnRegisterAction is a generated struct representing the https://schema.org/UnRegisterAction class.
type UnRegisterAction struct {
}

// DonateAction is a generated struct representing the https://schema.org/DonateAction class.
type DonateAction struct {
	Recipient *Person `json:"recipient"` //https://schema.org/Person
}

// MovingCompany is a generated struct representing the https://schema.org/MovingCompany class.
type MovingCompany struct {
}

// NewsArticle is a generated struct representing the https://schema.org/NewsArticle class.
type NewsArticle struct {
	Dateline     *Text `json:"dateline"`     //https://schema.org/Text
	PrintColumn  *Text `json:"printcolumn"`  //https://schema.org/Text
	PrintEdition *Text `json:"printedition"` //https://schema.org/Text
	PrintPage    *Text `json:"printpage"`    //https://schema.org/Text
	PrintSection *Text `json:"printsection"` //https://schema.org/Text
}

// DDxElement is a generated struct representing the https://schema.org/DDxElement class.
type DDxElement struct {
	Diagnosis          *MedicalCondition     `json:"diagnosis"`          //https://schema.org/MedicalCondition
	DistinguishingSign *MedicalSignOrSymptom `json:"distinguishingsign"` //https://schema.org/MedicalSignOrSymptom
}

// EmployeeRole is a generated struct representing the https://schema.org/EmployeeRole class.
type EmployeeRole struct {
	BaseSalary     *PriceSpecification `json:"basesalary"`     //https://schema.org/PriceSpecification
	SalaryCurrency *Text               `json:"salarycurrency"` //https://schema.org/Text
}

// LearningResource is a generated struct representing the https://schema.org/LearningResource class.
type LearningResource struct {
	Assesses             *Text            `json:"assesses"`             //https://schema.org/Text
	CompetencyRequired   *URL             `json:"competencyrequired"`   //https://schema.org/URL
	EducationalAlignment *AlignmentObject `json:"educationalalignment"` //https://schema.org/AlignmentObject
	EducationalLevel     *URL             `json:"educationallevel"`     //https://schema.org/URL
	EducationalUse       *Text            `json:"educationaluse"`       //https://schema.org/Text
	LearningResourceType *Text            `json:"learningresourcetype"` //https://schema.org/Text
	Teaches              *Text            `json:"teaches"`              //https://schema.org/Text
}

// MedicalTest is a generated struct representing the https://schema.org/MedicalTest class.
type MedicalTest struct {
	AffectedBy     *Drug             `json:"affectedby"`     //https://schema.org/Drug
	NormalRange    *Text             `json:"normalrange"`    //https://schema.org/Text
	SignDetected   *MedicalSign      `json:"signdetected"`   //https://schema.org/MedicalSign
	UsedToDiagnose *MedicalCondition `json:"usedtodiagnose"` //https://schema.org/MedicalCondition
	UsesDevice     *MedicalDevice    `json:"usesdevice"`     //https://schema.org/MedicalDevice
}

// ReturnAction is a generated struct representing the https://schema.org/ReturnAction class.
type ReturnAction struct {
	Recipient *Person `json:"recipient"` //https://schema.org/Person
}

// ReturnMethodEnumeration is a generated struct representing the https://schema.org/ReturnMethodEnumeration class.
type ReturnMethodEnumeration struct {
}

// Number is a generated struct representing the https://schema.org/Number class.
type Number struct {
}

// AccountingService is a generated struct representing the https://schema.org/AccountingService class.
type AccountingService struct {
}

// MedicalSpecialty is a generated struct representing the https://schema.org/MedicalSpecialty class.
type MedicalSpecialty struct {
}

// Physician is a generated struct representing the https://schema.org/Physician class.
type Physician struct {
	AvailableService    *MedicalTherapy   `json:"availableservice"`    //https://schema.org/MedicalTherapy
	HospitalAffiliation *Hospital         `json:"hospitalaffiliation"` //https://schema.org/Hospital
	MedicalSpecialty    *MedicalSpecialty `json:"medicalspecialty"`    //https://schema.org/MedicalSpecialty
}

// RentAction is a generated struct representing the https://schema.org/RentAction class.
type RentAction struct {
	Landlord        *Person          `json:"landlord"`        //https://schema.org/Person
	RealEstateAgent *RealEstateAgent `json:"realestateagent"` //https://schema.org/RealEstateAgent
}

// SportingGoodsStore is a generated struct representing the https://schema.org/SportingGoodsStore class.
type SportingGoodsStore struct {
}

// AuthorizeAction is a generated struct representing the https://schema.org/AuthorizeAction class.
type AuthorizeAction struct {
	Recipient *Person `json:"recipient"` //https://schema.org/Person
}

// Barcode is a generated struct representing the https://schema.org/Barcode class.
type Barcode struct {
}

// PlaceOfWorship is a generated struct representing the https://schema.org/PlaceOfWorship class.
type PlaceOfWorship struct {
}

// Schedule is a generated struct representing the https://schema.org/Schedule class.
type Schedule struct {
	ByDay            *Text     `json:"byday"`            //https://schema.org/Text
	ByMonth          *Integer  `json:"bymonth"`          //https://schema.org/Integer
	ByMonthDay       *Integer  `json:"bymonthday"`       //https://schema.org/Integer
	ByMonthWeek      *Integer  `json:"bymonthweek"`      //https://schema.org/Integer
	EndDate          *DateTime `json:"enddate"`          //https://schema.org/DateTime
	EndTime          *Time     `json:"endtime"`          //https://schema.org/Time
	ExceptDate       *DateTime `json:"exceptdate"`       //https://schema.org/DateTime
	RepeatCount      *Integer  `json:"repeatcount"`      //https://schema.org/Integer
	RepeatFrequency  *Text     `json:"repeatfrequency"`  //https://schema.org/Text
	ScheduleTimezone *Text     `json:"scheduletimezone"` //https://schema.org/Text
	StartDate        *DateTime `json:"startdate"`        //https://schema.org/DateTime
	StartTime        *Time     `json:"starttime"`        //https://schema.org/Time
	Duration         *Duration `json:"duration"`         //https://schema.org/Duration
}

// CableOrSatelliteService is a generated struct representing the https://schema.org/CableOrSatelliteService class.
type CableOrSatelliteService struct {
}

// ContactPointOption is a generated struct representing the https://schema.org/ContactPointOption class.
type ContactPointOption struct {
}

// Hostel is a generated struct representing the https://schema.org/Hostel class.
type Hostel struct {
}

// BorrowAction is a generated struct representing the https://schema.org/BorrowAction class.
type BorrowAction struct {
	Lender *Person `json:"lender"` //https://schema.org/Person
}

// HowTo is a generated struct representing the https://schema.org/HowTo class.
type HowTo struct {
	EstimatedCost *Text     `json:"estimatedcost"` //https://schema.org/Text
	PrepTime      *Duration `json:"preptime"`      //https://schema.org/Duration
	Steps         *Text     `json:"steps"`         //https://schema.org/Text
	Tool          *Text     `json:"tool"`          //https://schema.org/Text
	TotalTime     *Duration `json:"totaltime"`     //https://schema.org/Duration
	PerformTime   *Duration `json:"performtime"`   //https://schema.org/Duration
	Yield         *Text     `json:"yield"`         //https://schema.org/Text
	Step          *Text     `json:"step"`          //https://schema.org/Text
	Supply        *Text     `json:"supply"`        //https://schema.org/Text
}

// MedicalCode is a generated struct representing the https://schema.org/MedicalCode class.
type MedicalCode struct {
	CodeValue    *Text `json:"codevalue"`    //https://schema.org/Text
	CodingSystem *Text `json:"codingsystem"` //https://schema.org/Text
}

// Plumber is a generated struct representing the https://schema.org/Plumber class.
type Plumber struct {
}

// UserDownloads is a generated struct representing the https://schema.org/UserDownloads class.
type UserDownloads struct {
}

// BloodTest is a generated struct representing the https://schema.org/BloodTest class.
type BloodTest struct {
}

// BowlingAlley is a generated struct representing the https://schema.org/BowlingAlley class.
type BowlingAlley struct {
}

// BroadcastFrequencySpecification is a generated struct representing the https://schema.org/BroadcastFrequencySpecification class.
type BroadcastFrequencySpecification struct {
	BroadcastFrequencyValue   *QuantitativeValue `json:"broadcastfrequencyvalue"`   //https://schema.org/QuantitativeValue
	BroadcastSignalModulation *Text              `json:"broadcastsignalmodulation"` //https://schema.org/Text
	BroadcastSubChannel       *Text              `json:"broadcastsubchannel"`       //https://schema.org/Text
}

// Ligament is a generated struct representing the https://schema.org/Ligament class.
type Ligament struct {
}

// LikeAction is a generated struct representing the https://schema.org/LikeAction class.
type LikeAction struct {
}

// TheaterGroup is a generated struct representing the https://schema.org/TheaterGroup class.
type TheaterGroup struct {
}

// Integer is a generated struct representing the https://schema.org/Integer class.
type Integer struct {
}

// MedicalRiskScore is a generated struct representing the https://schema.org/MedicalRiskScore class.
type MedicalRiskScore struct {
	Algorithm *Text `json:"algorithm"` //https://schema.org/Text
}

// Message is a generated struct representing the https://schema.org/Message class.
type Message struct {
	BccRecipient      *Person       `json:"bccrecipient"`      //https://schema.org/Person
	CcRecipient       *Person       `json:"ccrecipient"`       //https://schema.org/Person
	DateRead          *DateTime     `json:"dateread"`          //https://schema.org/DateTime
	DateReceived      *DateTime     `json:"datereceived"`      //https://schema.org/DateTime
	DateSent          *DateTime     `json:"datesent"`          //https://schema.org/DateTime
	MessageAttachment *CreativeWork `json:"messageattachment"` //https://schema.org/CreativeWork
	Sender            *Person       `json:"sender"`            //https://schema.org/Person
	ToRecipient       *Person       `json:"torecipient"`       //https://schema.org/Person
	Recipient         *Person       `json:"recipient"`         //https://schema.org/Person
}

// PropertyValue is a generated struct representing the https://schema.org/PropertyValue class.
type PropertyValue struct {
	MaxValue             *Number `json:"maxvalue"`             //https://schema.org/Number
	MeasurementMethod    *URL    `json:"measurementmethod"`    //https://schema.org/URL
	MinValue             *Number `json:"minvalue"`             //https://schema.org/Number
	PropertyID           *URL    `json:"propertyid"`           //https://schema.org/URL
	UnitCode             *URL    `json:"unitcode"`             //https://schema.org/URL
	UnitText             *Text   `json:"unittext"`             //https://schema.org/Text
	Value                *Text   `json:"value"`                //https://schema.org/Text
	ValueReference       *Text   `json:"valuereference"`       //https://schema.org/Text
	MeasurementTechnique *URL    `json:"measurementtechnique"` //https://schema.org/URL
}

// WPHeader is a generated struct representing the https://schema.org/WPHeader class.
type WPHeader struct {
}

// WantAction is a generated struct representing the https://schema.org/WantAction class.
type WantAction struct {
}

// Code is a generated struct representing the https://schema.org/Code class.
type Code struct {
}

// MensClothingStore is a generated struct representing the https://schema.org/MensClothingStore class.
type MensClothingStore struct {
}

// PerformingArtsTheater is a generated struct representing the https://schema.org/PerformingArtsTheater class.
type PerformingArtsTheater struct {
}

// DefinedRegion is a generated struct representing the https://schema.org/DefinedRegion class.
type DefinedRegion struct {
	AddressCountry   *Text                         `json:"addresscountry"`   //https://schema.org/Text
	AddressRegion    *Text                         `json:"addressregion"`    //https://schema.org/Text
	PostalCode       *Text                         `json:"postalcode"`       //https://schema.org/Text
	PostalCodePrefix *Text                         `json:"postalcodeprefix"` //https://schema.org/Text
	PostalCodeRange  *PostalCodeRangeSpecification `json:"postalcoderange"`  //https://schema.org/PostalCodeRangeSpecification
}

// CategoryCode is a generated struct representing the https://schema.org/CategoryCode class.
type CategoryCode struct {
	CodeValue *Text `json:"codevalue"` //https://schema.org/Text
	InCodeSet *URL  `json:"incodeset"` //https://schema.org/URL
}

// CssSelectorType is a generated struct representing the https://schema.org/CssSelectorType class.
type CssSelectorType struct {
}

// PodcastEpisode is a generated struct representing the https://schema.org/PodcastEpisode class.
type PodcastEpisode struct {
}

// VideoObjectSnapshot is a generated struct representing the https://schema.org/VideoObjectSnapshot class.
type VideoObjectSnapshot struct {
}

// VirtualLocation is a generated struct representing the https://schema.org/VirtualLocation class.
type VirtualLocation struct {
}

// AudioObjectSnapshot is a generated struct representing the https://schema.org/AudioObjectSnapshot class.
type AudioObjectSnapshot struct {
}

// BusReservation is a generated struct representing the https://schema.org/BusReservation class.
type BusReservation struct {
}

// ShoppingCenter is a generated struct representing the https://schema.org/ShoppingCenter class.
type ShoppingCenter struct {
}

// PreventionIndication is a generated struct representing the https://schema.org/PreventionIndication class.
type PreventionIndication struct {
}

// Protein is a generated struct representing the https://schema.org/Protein class.
type Protein struct {
	HasBioPolymerSequence *Text `json:"hasbiopolymersequence"` //https://schema.org/Text
}

// Synagogue is a generated struct representing the https://schema.org/Synagogue class.
type Synagogue struct {
}

// AutoRental is a generated struct representing the https://schema.org/AutoRental class.
type AutoRental struct {
}

// LocationFeatureSpecification is a generated struct representing the https://schema.org/LocationFeatureSpecification class.
type LocationFeatureSpecification struct {
	HoursAvailable *OpeningHoursSpecification `json:"hoursavailable"` //https://schema.org/OpeningHoursSpecification
	ValidFrom      *DateTime                  `json:"validfrom"`      //https://schema.org/DateTime
	ValidThrough   *DateTime                  `json:"validthrough"`   //https://schema.org/DateTime
}

// ScreeningEvent is a generated struct representing the https://schema.org/ScreeningEvent class.
type ScreeningEvent struct {
	SubtitleLanguage *Text  `json:"subtitlelanguage"` //https://schema.org/Text
	VideoFormat      *Text  `json:"videoformat"`      //https://schema.org/Text
	WorkPresented    *Movie `json:"workpresented"`    //https://schema.org/Movie
}

// SubwayStation is a generated struct representing the https://schema.org/SubwayStation class.
type SubwayStation struct {
}

// Thing is a generated struct representing the https://schema.org/Thing class.
type Thing struct {
	AdditionalType            *URL        `json:"additionaltype"`            //https://schema.org/URL
	DisambiguatingDescription *Text       `json:"disambiguatingdescription"` //https://schema.org/Text
	Name                      *Text       `json:"name"`                      //https://schema.org/Text
	PotentialAction           *Action     `json:"potentialaction"`           //https://schema.org/Action
	SameAs                    *URL        `json:"sameas"`                    //https://schema.org/URL
	Url                       *URL        `json:"url"`                       //https://schema.org/URL
	AlternateName             *Text       `json:"alternatename"`             //https://schema.org/Text
	MainEntityOfPage          *URL        `json:"mainentityofpage"`          //https://schema.org/URL
	SubjectOf                 *Event      `json:"subjectof"`                 //https://schema.org/Event
	Image                     *URL        `json:"image"`                     //https://schema.org/URL
	Description               *TextObject `json:"description"`               //https://schema.org/TextObject
	Identifier                *URL        `json:"identifier"`                //https://schema.org/URL
}

// LoseAction is a generated struct representing the https://schema.org/LoseAction class.
type LoseAction struct {
	Winner *Person `json:"winner"` //https://schema.org/Person
}

// VoteAction is a generated struct representing the https://schema.org/VoteAction class.
type VoteAction struct {
	Candidate *Person `json:"candidate"` //https://schema.org/Person
}

// Zoo is a generated struct representing the https://schema.org/Zoo class.
type Zoo struct {
}

// IceCreamShop is a generated struct representing the https://schema.org/IceCreamShop class.
type IceCreamShop struct {
}

// LegalService is a generated struct representing the https://schema.org/LegalService class.
type LegalService struct {
}

// SolveMathAction is a generated struct representing the https://schema.org/SolveMathAction class.
type SolveMathAction struct {
	EduQuestionType *Text `json:"eduquestiontype"` //https://schema.org/Text
}

// ConfirmAction is a generated struct representing the https://schema.org/ConfirmAction class.
type ConfirmAction struct {
}

// MonetaryAmountDistribution is a generated struct representing the https://schema.org/MonetaryAmountDistribution class.
type MonetaryAmountDistribution struct {
	Currency *Text `json:"currency"` //https://schema.org/Text
}

// Movie is a generated struct representing the https://schema.org/Movie class.
type Movie struct {
	Actors            *Person       `json:"actors"`            //https://schema.org/Person
	CountryOfOrigin   *Country      `json:"countryoforigin"`   //https://schema.org/Country
	Directors         *Person       `json:"directors"`         //https://schema.org/Person
	MusicBy           *Person       `json:"musicby"`           //https://schema.org/Person
	ProductionCompany *Organization `json:"productioncompany"` //https://schema.org/Organization
	SubtitleLanguage  *Text         `json:"subtitlelanguage"`  //https://schema.org/Text
	TitleEIDR         *URL          `json:"titleeidr"`         //https://schema.org/URL
	Trailer           *VideoObject  `json:"trailer"`           //https://schema.org/VideoObject
	Director          *Person       `json:"director"`          //https://schema.org/Person
	Duration          *Duration     `json:"duration"`          //https://schema.org/Duration
	Actor             *Person       `json:"actor"`             //https://schema.org/Person
}

// OfferItemCondition is a generated struct representing the https://schema.org/OfferItemCondition class.
type OfferItemCondition struct {
}

// HinduTemple is a generated struct representing the https://schema.org/HinduTemple class.
type HinduTemple struct {
}

// SportsOrganization is a generated struct representing the https://schema.org/SportsOrganization class.
type SportsOrganization struct {
	Sport *URL `json:"sport"` //https://schema.org/URL
}

// Quiz is a generated struct representing the https://schema.org/Quiz class.
type Quiz struct {
}

// SiteNavigationElement is a generated struct representing the https://schema.org/SiteNavigationElement class.
type SiteNavigationElement struct {
}

// AnimalShelter is a generated struct representing the https://schema.org/AnimalShelter class.
type AnimalShelter struct {
}

// Car is a generated struct representing the https://schema.org/Car class.
type Car struct {
	AcrissCode *Text              `json:"acrisscode"` //https://schema.org/Text
	RoofLoad   *QuantitativeValue `json:"roofload"`   //https://schema.org/QuantitativeValue
}

// Corporation is a generated struct representing the https://schema.org/Corporation class.
type Corporation struct {
	TickerSymbol *Text `json:"tickersymbol"` //https://schema.org/Text
}

// Locksmith is a generated struct representing the https://schema.org/Locksmith class.
type Locksmith struct {
}

// Newspaper is a generated struct representing the https://schema.org/Newspaper class.
type Newspaper struct {
}

// PaymentService is a generated struct representing the https://schema.org/PaymentService class.
type PaymentService struct {
}

// DrugPrescriptionStatus is a generated struct representing the https://schema.org/DrugPrescriptionStatus class.
type DrugPrescriptionStatus struct {
}

// FlightReservation is a generated struct representing the https://schema.org/FlightReservation class.
type FlightReservation struct {
	BoardingGroup           *Text `json:"boardinggroup"`           //https://schema.org/Text
	PassengerPriorityStatus *Text `json:"passengerprioritystatus"` //https://schema.org/Text
	PassengerSequenceNumber *Text `json:"passengersequencenumber"` //https://schema.org/Text
	SecurityScreening       *Text `json:"securityscreening"`       //https://schema.org/Text
}

// PerformAction is a generated struct representing the https://schema.org/PerformAction class.
type PerformAction struct {
	EntertainmentBusiness *EntertainmentBusiness `json:"entertainmentbusiness"` //https://schema.org/EntertainmentBusiness
}

// ExercisePlan is a generated struct representing the https://schema.org/ExercisePlan class.
type ExercisePlan struct {
	ActivityDuration   *QuantitativeValue `json:"activityduration"`   //https://schema.org/QuantitativeValue
	ActivityFrequency  *Text              `json:"activityfrequency"`  //https://schema.org/Text
	AdditionalVariable *Text              `json:"additionalvariable"` //https://schema.org/Text
	ExerciseType       *Text              `json:"exercisetype"`       //https://schema.org/Text
	Intensity          *Text              `json:"intensity"`          //https://schema.org/Text
	Repetitions        *QuantitativeValue `json:"repetitions"`        //https://schema.org/QuantitativeValue
	RestPeriods        *Text              `json:"restperiods"`        //https://schema.org/Text
	Workload           *QuantitativeValue `json:"workload"`           //https://schema.org/QuantitativeValue
}

// FoodService is a generated struct representing the https://schema.org/FoodService class.
type FoodService struct {
}

// MedicalDevicePurpose is a generated struct representing the https://schema.org/MedicalDevicePurpose class.
type MedicalDevicePurpose struct {
}

// SuperficialAnatomy is a generated struct representing the https://schema.org/SuperficialAnatomy class.
type SuperficialAnatomy struct {
	AssociatedPathophysiology *Text             `json:"associatedpathophysiology"` //https://schema.org/Text
	RelatedAnatomy            *AnatomicalSystem `json:"relatedanatomy"`            //https://schema.org/AnatomicalSystem
	RelatedCondition          *MedicalCondition `json:"relatedcondition"`          //https://schema.org/MedicalCondition
	RelatedTherapy            *MedicalTherapy   `json:"relatedtherapy"`            //https://schema.org/MedicalTherapy
	Significance              *Text             `json:"significance"`              //https://schema.org/Text
}

// TouristTrip is a generated struct representing the https://schema.org/TouristTrip class.
type TouristTrip struct {
	TouristType *Text `json:"touristtype"` //https://schema.org/Text
}

// Distance is a generated struct representing the https://schema.org/Distance class.
type Distance struct {
}

// HowToDirection is a generated struct representing the https://schema.org/HowToDirection class.
type HowToDirection struct {
	AfterMedia  *URL      `json:"aftermedia"`  //https://schema.org/URL
	BeforeMedia *URL      `json:"beforemedia"` //https://schema.org/URL
	DuringMedia *URL      `json:"duringmedia"` //https://schema.org/URL
	PrepTime    *Duration `json:"preptime"`    //https://schema.org/Duration
	Tool        *Text     `json:"tool"`        //https://schema.org/Text
	TotalTime   *Duration `json:"totaltime"`   //https://schema.org/Duration
	PerformTime *Duration `json:"performtime"` //https://schema.org/Duration
	Supply      *Text     `json:"supply"`      //https://schema.org/Text
}

// BusinessEntityType is a generated struct representing the https://schema.org/BusinessEntityType class.
type BusinessEntityType struct {
}

// Game is a generated struct representing the https://schema.org/Game class.
type Game struct {
	CharacterAttribute *Thing             `json:"characterattribute"` //https://schema.org/Thing
	GameItem           *Thing             `json:"gameitem"`           //https://schema.org/Thing
	GameLocation       *URL               `json:"gamelocation"`       //https://schema.org/URL
	NumberOfPlayers    *QuantitativeValue `json:"numberofplayers"`    //https://schema.org/QuantitativeValue
	Quest              *Thing             `json:"quest"`              //https://schema.org/Thing
}

// LegalForceStatus is a generated struct representing the https://schema.org/LegalForceStatus class.
type LegalForceStatus struct {
}

// TaxiService is a generated struct representing the https://schema.org/TaxiService class.
type TaxiService struct {
}

// Collection is a generated struct representing the https://schema.org/Collection class.
type Collection struct {
	CollectionSize *Integer `json:"collectionsize"` //https://schema.org/Integer
}

// EnergyStarEnergyEfficiencyEnumeration is a generated struct representing the https://schema.org/EnergyStarEnergyEfficiencyEnumeration class.
type EnergyStarEnergyEfficiencyEnumeration struct {
}

// HealthAspectEnumeration is a generated struct representing the https://schema.org/HealthAspectEnumeration class.
type HealthAspectEnumeration struct {
}

// PublicationVolume is a generated struct representing the https://schema.org/PublicationVolume class.
type PublicationVolume struct {
	PageEnd      *Text `json:"pageend"`      //https://schema.org/Text
	PageStart    *Text `json:"pagestart"`    //https://schema.org/Text
	Pagination   *Text `json:"pagination"`   //https://schema.org/Text
	VolumeNumber *Text `json:"volumenumber"` //https://schema.org/Text
}

// RVPark is a generated struct representing the https://schema.org/RVPark class.
type RVPark struct {
}

// ReservationStatusType is a generated struct representing the https://schema.org/ReservationStatusType class.
type ReservationStatusType struct {
}

// UKNonprofitType is a generated struct representing the https://schema.org/UKNonprofitType class.
type UKNonprofitType struct {
}

// BusinessEvent is a generated struct representing the https://schema.org/BusinessEvent class.
type BusinessEvent struct {
}

// DiagnosticProcedure is a generated struct representing the https://schema.org/DiagnosticProcedure class.
type DiagnosticProcedure struct {
}

// MusicVenue is a generated struct representing the https://schema.org/MusicVenue class.
type MusicVenue struct {
}

// Notary is a generated struct representing the https://schema.org/Notary class.
type Notary struct {
}

// AgreeAction is a generated struct representing the https://schema.org/AgreeAction class.
type AgreeAction struct {
}

// ReplaceAction is a generated struct representing the https://schema.org/ReplaceAction class.
type ReplaceAction struct {
	Replacee *Thing `json:"replacee"` //https://schema.org/Thing
	Replacer *Thing `json:"replacer"` //https://schema.org/Thing
}

// TelevisionChannel is a generated struct representing the https://schema.org/TelevisionChannel class.
type TelevisionChannel struct {
}

// AdvertiserContentArticle is a generated struct representing the https://schema.org/AdvertiserContentArticle class.
type AdvertiserContentArticle struct {
}

// BrainStructure is a generated struct representing the https://schema.org/BrainStructure class.
type BrainStructure struct {
}

// RejectAction is a generated struct representing the https://schema.org/RejectAction class.
type RejectAction struct {
}

// SheetMusic is a generated struct representing the https://schema.org/SheetMusic class.
type SheetMusic struct {
}

// BikeStore is a generated struct representing the https://schema.org/BikeStore class.
type BikeStore struct {
}

// SaleEvent is a generated struct representing the https://schema.org/SaleEvent class.
type SaleEvent struct {
}

// BuddhistTemple is a generated struct representing the https://schema.org/BuddhistTemple class.
type BuddhistTemple struct {
}

// DanceEvent is a generated struct representing the https://schema.org/DanceEvent class.
type DanceEvent struct {
}

// MediaObject is a generated struct representing the https://schema.org/MediaObject class.
type MediaObject struct {
	AssociatedArticle    *NewsArticle       `json:"associatedarticle"`    //https://schema.org/NewsArticle
	Bitrate              *Text              `json:"bitrate"`              //https://schema.org/Text
	ContentSize          *Text              `json:"contentsize"`          //https://schema.org/Text
	ContentUrl           *URL               `json:"contenturl"`           //https://schema.org/URL
	EmbedUrl             *URL               `json:"embedurl"`             //https://schema.org/URL
	EndTime              *Time              `json:"endtime"`              //https://schema.org/Time
	Height               *QuantitativeValue `json:"height"`               //https://schema.org/QuantitativeValue
	IneligibleRegion     *Text              `json:"ineligibleregion"`     //https://schema.org/Text
	InterpretedAsClaim   *Claim             `json:"interpretedasclaim"`   //https://schema.org/Claim
	PlayerType           *Text              `json:"playertype"`           //https://schema.org/Text
	ProductionCompany    *Organization      `json:"productioncompany"`    //https://schema.org/Organization
	RegionsAllowed       *Place             `json:"regionsallowed"`       //https://schema.org/Place
	RequiresSubscription *MediaSubscription `json:"requiressubscription"` //https://schema.org/MediaSubscription
	Sha256               *Text              `json:"sha256"`               //https://schema.org/Text
	StartTime            *Time              `json:"starttime"`            //https://schema.org/Time
	UploadDate           *DateTime          `json:"uploaddate"`           //https://schema.org/DateTime
	Width                *QuantitativeValue `json:"width"`                //https://schema.org/QuantitativeValue
	Duration             *Duration          `json:"duration"`             //https://schema.org/Duration
	EncodesCreativeWork  *CreativeWork      `json:"encodescreativework"`  //https://schema.org/CreativeWork
	EncodingFormat       *URL               `json:"encodingformat"`       //https://schema.org/URL
}

// MusicPlaylist is a generated struct representing the https://schema.org/MusicPlaylist class.
type MusicPlaylist struct {
	NumTracks *Integer        `json:"numtracks"` //https://schema.org/Integer
	Tracks    *MusicRecording `json:"tracks"`    //https://schema.org/MusicRecording
	Track     *MusicRecording `json:"track"`     //https://schema.org/MusicRecording
}

// Accommodation is a generated struct representing the https://schema.org/Accommodation class.
type Accommodation struct {
	AccommodationCategory    *Text                         `json:"accommodationcategory"`    //https://schema.org/Text
	AccommodationFloorPlan   *FloorPlan                    `json:"accommodationfloorplan"`   //https://schema.org/FloorPlan
	AmenityFeature           *LocationFeatureSpecification `json:"amenityfeature"`           //https://schema.org/LocationFeatureSpecification
	Bed                      *Text                         `json:"bed"`                      //https://schema.org/Text
	FloorLevel               *Text                         `json:"floorlevel"`               //https://schema.org/Text
	FloorSize                *QuantitativeValue            `json:"floorsize"`                //https://schema.org/QuantitativeValue
	LeaseLength              *QuantitativeValue            `json:"leaselength"`              //https://schema.org/QuantitativeValue
	NumberOfBathroomsTotal   *Integer                      `json:"numberofbathroomstotal"`   //https://schema.org/Integer
	NumberOfBedrooms         *QuantitativeValue            `json:"numberofbedrooms"`         //https://schema.org/QuantitativeValue
	NumberOfFullBathrooms    *Number                       `json:"numberoffullbathrooms"`    //https://schema.org/Number
	NumberOfPartialBathrooms *Number                       `json:"numberofpartialbathrooms"` //https://schema.org/Number
	NumberOfRooms            *QuantitativeValue            `json:"numberofrooms"`            //https://schema.org/QuantitativeValue
	Occupancy                *QuantitativeValue            `json:"occupancy"`                //https://schema.org/QuantitativeValue
	PermittedUsage           *Text                         `json:"permittedusage"`           //https://schema.org/Text
	PetsAllowed              *Text                         `json:"petsallowed"`              //https://schema.org/Text
	TourBookingPage          *URL                          `json:"tourbookingpage"`          //https://schema.org/URL
	YearBuilt                *Number                       `json:"yearbuilt"`                //https://schema.org/Number
}

// Bakery is a generated struct representing the https://schema.org/Bakery class.
type Bakery struct {
}

// ExhibitionEvent is a generated struct representing the https://schema.org/ExhibitionEvent class.
type ExhibitionEvent struct {
}

// MedicalScholarlyArticle is a generated struct representing the https://schema.org/MedicalScholarlyArticle class.
type MedicalScholarlyArticle struct {
	PublicationType *Text `json:"publicationtype"` //https://schema.org/Text
}

// UseAction is a generated struct representing the https://schema.org/UseAction class.
type UseAction struct {
}

// WPFooter is a generated struct representing the https://schema.org/WPFooter class.
type WPFooter struct {
}

// InformAction is a generated struct representing the https://schema.org/InformAction class.
type InformAction struct {
	Event *Event `json:"event"` //https://schema.org/Event
}

// Periodical is a generated struct representing the https://schema.org/Periodical class.
type Periodical struct {
}

// SatiricalArticle is a generated struct representing the https://schema.org/SatiricalArticle class.
type SatiricalArticle struct {
}

// SocialMediaPosting is a generated struct representing the https://schema.org/SocialMediaPosting class.
type SocialMediaPosting struct {
	SharedContent *CreativeWork `json:"sharedcontent"` //https://schema.org/CreativeWork
}

// SomeProducts is a generated struct representing the https://schema.org/SomeProducts class.
type SomeProducts struct {
	InventoryLevel *QuantitativeValue `json:"inventorylevel"` //https://schema.org/QuantitativeValue
}

// DiagnosticLab is a generated struct representing the https://schema.org/DiagnosticLab class.
type DiagnosticLab struct {
	AvailableTest *MedicalTest `json:"availabletest"` //https://schema.org/MedicalTest
}

// ItemAvailability is a generated struct representing the https://schema.org/ItemAvailability class.
type ItemAvailability struct {
}

// MarryAction is a generated struct representing the https://schema.org/MarryAction class.
type MarryAction struct {
}

// MusicVideoObject is a generated struct representing the https://schema.org/MusicVideoObject class.
type MusicVideoObject struct {
}

// Park is a generated struct representing the https://schema.org/Park class.
type Park struct {
}

// HomeAndConstructionBusiness is a generated struct representing the https://schema.org/HomeAndConstructionBusiness class.
type HomeAndConstructionBusiness struct {
}

// MedicalStudyStatus is a generated struct representing the https://schema.org/MedicalStudyStatus class.
type MedicalStudyStatus struct {
}

// Product is a generated struct representing the https://schema.org/Product class.
type Product struct {
	AdditionalProperty          *PropertyValue            `json:"additionalproperty"`          //https://schema.org/PropertyValue
	AggregateRating             *AggregateRating          `json:"aggregaterating"`             //https://schema.org/AggregateRating
	Asin                        *URL                      `json:"asin"`                        //https://schema.org/URL
	Awards                      *Text                     `json:"awards"`                      //https://schema.org/Text
	Brand                       *Organization             `json:"brand"`                       //https://schema.org/Organization
	Color                       *Text                     `json:"color"`                       //https://schema.org/Text
	CountryOfAssembly           *Text                     `json:"countryofassembly"`           //https://schema.org/Text
	CountryOfLastProcessing     *Text                     `json:"countryoflastprocessing"`     //https://schema.org/Text
	CountryOfOrigin             *Country                  `json:"countryoforigin"`             //https://schema.org/Country
	Depth                       *QuantitativeValue        `json:"depth"`                       //https://schema.org/QuantitativeValue
	Gtin12                      *Text                     `json:"gtin12"`                      //https://schema.org/Text
	Gtin13                      *Text                     `json:"gtin13"`                      //https://schema.org/Text
	Gtin14                      *Text                     `json:"gtin14"`                      //https://schema.org/Text
	Gtin8                       *Text                     `json:"gtin8"`                       //https://schema.org/Text
	HasAdultConsideration       *AdultOrientedEnumeration `json:"hasadultconsideration"`       //https://schema.org/AdultOrientedEnumeration
	HasEnergyConsumptionDetails *EnergyConsumptionDetails `json:"hasenergyconsumptiondetails"` //https://schema.org/EnergyConsumptionDetails
	HasMeasurement              *QuantitativeValue        `json:"hasmeasurement"`              //https://schema.org/QuantitativeValue
	HasMerchantReturnPolicy     *MerchantReturnPolicy     `json:"hasmerchantreturnpolicy"`     //https://schema.org/MerchantReturnPolicy
	Height                      *QuantitativeValue        `json:"height"`                      //https://schema.org/QuantitativeValue
	InProductGroupWithID        *Text                     `json:"inproductgroupwithid"`        //https://schema.org/Text
	IsAccessoryOrSparePartFor   *Product                  `json:"isaccessoryorsparepartfor"`   //https://schema.org/Product
	IsConsumableFor             *Product                  `json:"isconsumablefor"`             //https://schema.org/Product
	IsFamilyFriendly            *Boolean                  `json:"isfamilyfriendly"`            //https://schema.org/Boolean
	IsRelatedTo                 *Service                  `json:"isrelatedto"`                 //https://schema.org/Service
	IsSimilarTo                 *Service                  `json:"issimilarto"`                 //https://schema.org/Service
	ItemCondition               *OfferItemCondition       `json:"itemcondition"`               //https://schema.org/OfferItemCondition
	Keywords                    *URL                      `json:"keywords"`                    //https://schema.org/URL
	Logo                        *URL                      `json:"logo"`                        //https://schema.org/URL
	Manufacturer                *Organization             `json:"manufacturer"`                //https://schema.org/Organization
	MobileUrl                   *Text                     `json:"mobileurl"`                   //https://schema.org/Text
	Model                       *Text                     `json:"model"`                       //https://schema.org/Text
	Mpn                         *Text                     `json:"mpn"`                         //https://schema.org/Text
	NegativeNotes               *WebContent               `json:"negativenotes"`               //https://schema.org/WebContent
	Nsn                         *Text                     `json:"nsn"`                         //https://schema.org/Text
	Pattern                     *Text                     `json:"pattern"`                     //https://schema.org/Text
	PositiveNotes               *WebContent               `json:"positivenotes"`               //https://schema.org/WebContent
	ProductID                   *Text                     `json:"productid"`                   //https://schema.org/Text
	ProductionDate              *Date                     `json:"productiondate"`              //https://schema.org/Date
	PurchaseDate                *Date                     `json:"purchasedate"`                //https://schema.org/Date
	ReleaseDate                 *Date                     `json:"releasedate"`                 //https://schema.org/Date
	Reviews                     *Review                   `json:"reviews"`                     //https://schema.org/Review
	Size                        *Text                     `json:"size"`                        //https://schema.org/Text
	Sku                         *Text                     `json:"sku"`                         //https://schema.org/Text
	Slogan                      *Text                     `json:"slogan"`                      //https://schema.org/Text
	Weight                      *QuantitativeValue        `json:"weight"`                      //https://schema.org/QuantitativeValue
	Width                       *QuantitativeValue        `json:"width"`                       //https://schema.org/QuantitativeValue
	Audience                    *Audience                 `json:"audience"`                    //https://schema.org/Audience
	Award                       *Text                     `json:"award"`                       //https://schema.org/Text
	Category                    *URL                      `json:"category"`                    //https://schema.org/URL
	Funding                     *Grant                    `json:"funding"`                     //https://schema.org/Grant
	IsVariantOf                 *ProductModel             `json:"isvariantof"`                 //https://schema.org/ProductModel
	Offers                      *Offer                    `json:"offers"`                      //https://schema.org/Offer
	Review                      *Review                   `json:"review"`                      //https://schema.org/Review
	Material                    *URL                      `json:"material"`                    //https://schema.org/URL
	Gtin                        *URL                      `json:"gtin"`                        //https://schema.org/URL
}

// ArchiveComponent is a generated struct representing the https://schema.org/ArchiveComponent class.
type ArchiveComponent struct {
	ItemLocation   *Text                `json:"itemlocation"`   //https://schema.org/Text
	HoldingArchive *ArchiveOrganization `json:"holdingarchive"` //https://schema.org/ArchiveOrganization
}

// ComedyClub is a generated struct representing the https://schema.org/ComedyClub class.
type ComedyClub struct {
}

// DriveWheelConfigurationValue is a generated struct representing the https://schema.org/DriveWheelConfigurationValue class.
type DriveWheelConfigurationValue struct {
}

// GovernmentBuilding is a generated struct representing the https://schema.org/GovernmentBuilding class.
type GovernmentBuilding struct {
}

// OccupationalExperienceRequirements is a generated struct representing the https://schema.org/OccupationalExperienceRequirements class.
type OccupationalExperienceRequirements struct {
	MonthsOfExperience *Number `json:"monthsofexperience"` //https://schema.org/Number
}

// SkiResort is a generated struct representing the https://schema.org/SkiResort class.
type SkiResort struct {
}

// DigitalDocumentPermissionType is a generated struct representing the https://schema.org/DigitalDocumentPermissionType class.
type DigitalDocumentPermissionType struct {
}

// Intangible is a generated struct representing the https://schema.org/Intangible class.
type Intangible struct {
}

// SizeSpecification is a generated struct representing the https://schema.org/SizeSpecification class.
type SizeSpecification struct {
	HasMeasurement       *QuantitativeValue `json:"hasmeasurement"`       //https://schema.org/QuantitativeValue
	SizeGroup            *Text              `json:"sizegroup"`            //https://schema.org/Text
	SizeSystem           *Text              `json:"sizesystem"`           //https://schema.org/Text
	SuggestedAge         *QuantitativeValue `json:"suggestedage"`         //https://schema.org/QuantitativeValue
	SuggestedGender      *Text              `json:"suggestedgender"`      //https://schema.org/Text
	SuggestedMeasurement *QuantitativeValue `json:"suggestedmeasurement"` //https://schema.org/QuantitativeValue
}

// CreditCard is a generated struct representing the https://schema.org/CreditCard class.
type CreditCard struct {
}

// Person is a generated struct representing the https://schema.org/Person class.
type Person struct {
	AdditionalName            *Text                              `json:"additionalname"`            //https://schema.org/Text
	Address                   *Text                              `json:"address"`                   //https://schema.org/Text
	Affiliation               *Organization                      `json:"affiliation"`               //https://schema.org/Organization
	AgentInteractionStatistic *InteractionCounter                `json:"agentinteractionstatistic"` //https://schema.org/InteractionCounter
	Awards                    *Text                              `json:"awards"`                    //https://schema.org/Text
	BirthDate                 *Date                              `json:"birthdate"`                 //https://schema.org/Date
	BirthPlace                *Place                             `json:"birthplace"`                //https://schema.org/Place
	Brand                     *Organization                      `json:"brand"`                     //https://schema.org/Organization
	CallSign                  *Text                              `json:"callsign"`                  //https://schema.org/Text
	Children                  *Person                            `json:"children"`                  //https://schema.org/Person
	Colleagues                *Person                            `json:"colleagues"`                //https://schema.org/Person
	ContactPoints             *ContactPoint                      `json:"contactpoints"`             //https://schema.org/ContactPoint
	DeathDate                 *Date                              `json:"deathdate"`                 //https://schema.org/Date
	DeathPlace                *Place                             `json:"deathplace"`                //https://schema.org/Place
	Duns                      *Text                              `json:"duns"`                      //https://schema.org/Text
	Email                     *Text                              `json:"email"`                     //https://schema.org/Text
	FamilyName                *Text                              `json:"familyname"`                //https://schema.org/Text
	FaxNumber                 *Text                              `json:"faxnumber"`                 //https://schema.org/Text
	Follows                   *Person                            `json:"follows"`                   //https://schema.org/Person
	Funder                    *Person                            `json:"funder"`                    //https://schema.org/Person
	Gender                    *Text                              `json:"gender"`                    //https://schema.org/Text
	GivenName                 *Text                              `json:"givenname"`                 //https://schema.org/Text
	GlobalLocationNumber      *Text                              `json:"globallocationnumber"`      //https://schema.org/Text
	HasCredential             *EducationalOccupationalCredential `json:"hascredential"`             //https://schema.org/EducationalOccupationalCredential
	HasOccupation             *Occupation                        `json:"hasoccupation"`             //https://schema.org/Occupation
	HasOfferCatalog           *OfferCatalog                      `json:"hasoffercatalog"`           //https://schema.org/OfferCatalog
	HasPOS                    *Place                             `json:"haspos"`                    //https://schema.org/Place
	Height                    *QuantitativeValue                 `json:"height"`                    //https://schema.org/QuantitativeValue
	HomeLocation              *Place                             `json:"homelocation"`              //https://schema.org/Place
	HonorificPrefix           *Text                              `json:"honorificprefix"`           //https://schema.org/Text
	HonorificSuffix           *Text                              `json:"honorificsuffix"`           //https://schema.org/Text
	IsicV4                    *Text                              `json:"isicv4"`                    //https://schema.org/Text
	JobTitle                  *Text                              `json:"jobtitle"`                  //https://schema.org/Text
	Knows                     *Person                            `json:"knows"`                     //https://schema.org/Person
	KnowsAbout                *URL                               `json:"knowsabout"`                //https://schema.org/URL
	KnowsLanguage             *Text                              `json:"knowslanguage"`             //https://schema.org/Text
	Naics                     *Text                              `json:"naics"`                     //https://schema.org/Text
	Nationality               *Country                           `json:"nationality"`               //https://schema.org/Country
	NetWorth                  *PriceSpecification                `json:"networth"`                  //https://schema.org/PriceSpecification
	Owns                      *Product                           `json:"owns"`                      //https://schema.org/Product
	Parents                   *Person                            `json:"parents"`                   //https://schema.org/Person
	PerformerIn               *Event                             `json:"performerin"`               //https://schema.org/Event
	RelatedTo                 *Person                            `json:"relatedto"`                 //https://schema.org/Person
	Seeks                     *Demand                            `json:"seeks"`                     //https://schema.org/Demand
	Siblings                  *Person                            `json:"siblings"`                  //https://schema.org/Person
	Spouse                    *Person                            `json:"spouse"`                    //https://schema.org/Person
	TaxID                     *Text                              `json:"taxid"`                     //https://schema.org/Text
	Telephone                 *Text                              `json:"telephone"`                 //https://schema.org/Text
	VatID                     *Text                              `json:"vatid"`                     //https://schema.org/Text
	Weight                    *QuantitativeValue                 `json:"weight"`                    //https://schema.org/QuantitativeValue
	WorkLocation              *Place                             `json:"worklocation"`              //https://schema.org/Place
	WorksFor                  *Organization                      `json:"worksfor"`                  //https://schema.org/Organization
	AlumniOf                  *Organization                      `json:"alumniof"`                  //https://schema.org/Organization
	Award                     *Text                              `json:"award"`                     //https://schema.org/Text
	Colleague                 *URL                               `json:"colleague"`                 //https://schema.org/URL
	ContactPoint              *ContactPoint                      `json:"contactpoint"`              //https://schema.org/ContactPoint
	Funding                   *Grant                             `json:"funding"`                   //https://schema.org/Grant
	InteractionStatistic      *InteractionCounter                `json:"interactionstatistic"`      //https://schema.org/InteractionCounter
	MakesOffer                *Offer                             `json:"makesoffer"`                //https://schema.org/Offer
	Parent                    *Person                            `json:"parent"`                    //https://schema.org/Person
	Sibling                   *Person                            `json:"sibling"`                   //https://schema.org/Person
	Sponsor                   *Person                            `json:"sponsor"`                   //https://schema.org/Person
	MemberOf                  *ProgramMembership                 `json:"memberof"`                  //https://schema.org/ProgramMembership
	PublishingPrinciples      *URL                               `json:"publishingprinciples"`      //https://schema.org/URL
}

// TypeAndQuantityNode is a generated struct representing the https://schema.org/TypeAndQuantityNode class.
type TypeAndQuantityNode struct {
	AmountOfThisGood *Number           `json:"amountofthisgood"` //https://schema.org/Number
	BusinessFunction *BusinessFunction `json:"businessfunction"` //https://schema.org/BusinessFunction
	TypeOfGood       *Service          `json:"typeofgood"`       //https://schema.org/Service
	UnitCode         *URL              `json:"unitcode"`         //https://schema.org/URL
	UnitText         *Text             `json:"unittext"`         //https://schema.org/Text
}

// UnitPriceSpecification is a generated struct representing the https://schema.org/UnitPriceSpecification class.
type UnitPriceSpecification struct {
	BillingDuration    *QuantitativeValue             `json:"billingduration"`    //https://schema.org/QuantitativeValue
	BillingIncrement   *Number                        `json:"billingincrement"`   //https://schema.org/Number
	BillingStart       *Number                        `json:"billingstart"`       //https://schema.org/Number
	PriceComponentType *PriceComponentTypeEnumeration `json:"pricecomponenttype"` //https://schema.org/PriceComponentTypeEnumeration
	PriceType          *Text                          `json:"pricetype"`          //https://schema.org/Text
	ReferenceQuantity  *QuantitativeValue             `json:"referencequantity"`  //https://schema.org/QuantitativeValue
	UnitCode           *URL                           `json:"unitcode"`           //https://schema.org/URL
	UnitText           *Text                          `json:"unittext"`           //https://schema.org/Text
}

// IgnoreAction is a generated struct representing the https://schema.org/IgnoreAction class.
type IgnoreAction struct {
}

// AutoDealer is a generated struct representing the https://schema.org/AutoDealer class.
type AutoDealer struct {
}

// Float is a generated struct representing the https://schema.org/Float class.
type Float struct {
}

// MovieTheater is a generated struct representing the https://schema.org/MovieTheater class.
type MovieTheater struct {
	ScreenCount *Number `json:"screencount"` //https://schema.org/Number
}

// FindAction is a generated struct representing the https://schema.org/FindAction class.
type FindAction struct {
}

// Invoice is a generated struct representing the https://schema.org/Invoice class.
type Invoice struct {
	AccountId            *Text               `json:"accountid"`            //https://schema.org/Text
	BillingPeriod        *Duration           `json:"billingperiod"`        //https://schema.org/Duration
	ConfirmationNumber   *Text               `json:"confirmationnumber"`   //https://schema.org/Text
	Customer             *Person             `json:"customer"`             //https://schema.org/Person
	MinimumPaymentDue    *PriceSpecification `json:"minimumpaymentdue"`    //https://schema.org/PriceSpecification
	PaymentDue           *DateTime           `json:"paymentdue"`           //https://schema.org/DateTime
	PaymentMethod        *PaymentMethod      `json:"paymentmethod"`        //https://schema.org/PaymentMethod
	PaymentMethodId      *Text               `json:"paymentmethodid"`      //https://schema.org/Text
	PaymentStatus        *Text               `json:"paymentstatus"`        //https://schema.org/Text
	ReferencesOrder      *Order              `json:"referencesorder"`      //https://schema.org/Order
	ScheduledPaymentDate *Date               `json:"scheduledpaymentdate"` //https://schema.org/Date
	TotalPaymentDue      *PriceSpecification `json:"totalpaymentdue"`      //https://schema.org/PriceSpecification
	Broker               *Person             `json:"broker"`               //https://schema.org/Person
	Category             *URL                `json:"category"`             //https://schema.org/URL
	PaymentDueDate       *DateTime           `json:"paymentduedate"`       //https://schema.org/DateTime
	Provider             *Person             `json:"provider"`             //https://schema.org/Person
}

// RadioChannel is a generated struct representing the https://schema.org/RadioChannel class.
type RadioChannel struct {
}

// TravelAgency is a generated struct representing the https://schema.org/TravelAgency class.
type TravelAgency struct {
}

// AdultOrientedEnumeration is a generated struct representing the https://schema.org/AdultOrientedEnumeration class.
type AdultOrientedEnumeration struct {
}

// HealthClub is a generated struct representing the https://schema.org/HealthClub class.
type HealthClub struct {
}

// Map is a generated struct representing the https://schema.org/Map class.
type Map struct {
	MapType *MapCategoryType `json:"maptype"` //https://schema.org/MapCategoryType
}

// TrackAction is a generated struct representing the https://schema.org/TrackAction class.
type TrackAction struct {
	DeliveryMethod *DeliveryMethod `json:"deliverymethod"` //https://schema.org/DeliveryMethod
}

// WPSideBar is a generated struct representing the https://schema.org/WPSideBar class.
type WPSideBar struct {
}

// WearAction is a generated struct representing the https://schema.org/WearAction class.
type WearAction struct {
}

// AchieveAction is a generated struct representing the https://schema.org/AchieveAction class.
type AchieveAction struct {
}

// ProgramMembership is a generated struct representing the https://schema.org/ProgramMembership class.
type ProgramMembership struct {
	HostingOrganization    *Organization      `json:"hostingorganization"`    //https://schema.org/Organization
	Members                *Person            `json:"members"`                //https://schema.org/Person
	MembershipNumber       *Text              `json:"membershipnumber"`       //https://schema.org/Text
	MembershipPointsEarned *QuantitativeValue `json:"membershippointsearned"` //https://schema.org/QuantitativeValue
	ProgramName            *Text              `json:"programname"`            //https://schema.org/Text
	Member                 *Person            `json:"member"`                 //https://schema.org/Person
}

// HousePainter is a generated struct representing the https://schema.org/HousePainter class.
type HousePainter struct {
}

// MediaManipulationRatingEnumeration is a generated struct representing the https://schema.org/MediaManipulationRatingEnumeration class.
type MediaManipulationRatingEnumeration struct {
}

// CheckoutPage is a generated struct representing the https://schema.org/CheckoutPage class.
type CheckoutPage struct {
}

// DeactivateAction is a generated struct representing the https://schema.org/DeactivateAction class.
type DeactivateAction struct {
}

// MedicalSymptom is a generated struct representing the https://schema.org/MedicalSymptom class.
type MedicalSymptom struct {
}

// OpinionNewsArticle is a generated struct representing the https://schema.org/OpinionNewsArticle class.
type OpinionNewsArticle struct {
}

// Specialty is a generated struct representing the https://schema.org/Specialty class.
type Specialty struct {
}

// Suite is a generated struct representing the https://schema.org/Suite class.
type Suite struct {
	Bed           *Text              `json:"bed"`           //https://schema.org/Text
	NumberOfRooms *QuantitativeValue `json:"numberofrooms"` //https://schema.org/QuantitativeValue
	Occupancy     *QuantitativeValue `json:"occupancy"`     //https://schema.org/QuantitativeValue
}

// ShareAction is a generated struct representing the https://schema.org/ShareAction class.
type ShareAction struct {
}

// TouristDestination is a generated struct representing the https://schema.org/TouristDestination class.
type TouristDestination struct {
	IncludesAttraction *TouristAttraction `json:"includesattraction"` //https://schema.org/TouristAttraction
	TouristType        *Text              `json:"touristtype"`        //https://schema.org/Text
}

// CatholicChurch is a generated struct representing the https://schema.org/CatholicChurch class.
type CatholicChurch struct {
}

// Dataset is a generated struct representing the https://schema.org/Dataset class.
type Dataset struct {
	Catalog               *DataCatalog  `json:"catalog"`               //https://schema.org/DataCatalog
	DatasetTimeInterval   *DateTime     `json:"datasettimeinterval"`   //https://schema.org/DateTime
	Distribution          *DataDownload `json:"distribution"`          //https://schema.org/DataDownload
	IncludedDataCatalog   *DataCatalog  `json:"includeddatacatalog"`   //https://schema.org/DataCatalog
	Issn                  *Text         `json:"issn"`                  //https://schema.org/Text
	MeasurementMethod     *URL          `json:"measurementmethod"`     //https://schema.org/URL
	VariableMeasured      *Text         `json:"variablemeasured"`      //https://schema.org/Text
	MeasurementTechnique  *URL          `json:"measurementtechnique"`  //https://schema.org/URL
	IncludedInDataCatalog *DataCatalog  `json:"includedindatacatalog"` //https://schema.org/DataCatalog
}

// DayOfWeek is a generated struct representing the https://schema.org/DayOfWeek class.
type DayOfWeek struct {
}

// MeasurementMethodEnum is a generated struct representing the https://schema.org/MeasurementMethodEnum class.
type MeasurementMethodEnum struct {
}

// OfficeEquipmentStore is a generated struct representing the https://schema.org/OfficeEquipmentStore class.
type OfficeEquipmentStore struct {
}

// RadioSeries is a generated struct representing the https://schema.org/RadioSeries class.
type RadioSeries struct {
	Actors            *Person             `json:"actors"`            //https://schema.org/Person
	Directors         *Person             `json:"directors"`         //https://schema.org/Person
	Episodes          *Episode            `json:"episodes"`          //https://schema.org/Episode
	MusicBy           *Person             `json:"musicby"`           //https://schema.org/Person
	NumberOfEpisodes  *Integer            `json:"numberofepisodes"`  //https://schema.org/Integer
	NumberOfSeasons   *Integer            `json:"numberofseasons"`   //https://schema.org/Integer
	ProductionCompany *Organization       `json:"productioncompany"` //https://schema.org/Organization
	Seasons           *CreativeWorkSeason `json:"seasons"`           //https://schema.org/CreativeWorkSeason
	Trailer           *VideoObject        `json:"trailer"`           //https://schema.org/VideoObject
	ContainsSeason    *CreativeWorkSeason `json:"containsseason"`    //https://schema.org/CreativeWorkSeason
	Director          *Person             `json:"director"`          //https://schema.org/Person
	Episode           *Episode            `json:"episode"`           //https://schema.org/Episode
	Season            *URL                `json:"season"`            //https://schema.org/URL
	Actor             *Person             `json:"actor"`             //https://schema.org/Person
}

// UserPlays is a generated struct representing the https://schema.org/UserPlays class.
type UserPlays struct {
}

// CancelAction is a generated struct representing the https://schema.org/CancelAction class.
type CancelAction struct {
}

// SportsActivityLocation is a generated struct representing the https://schema.org/SportsActivityLocation class.
type SportsActivityLocation struct {
}

// Audiobook is a generated struct representing the https://schema.org/Audiobook class.
type Audiobook struct {
	ReadBy   *Person   `json:"readby"`   //https://schema.org/Person
	Duration *Duration `json:"duration"` //https://schema.org/Duration
}

// HealthPlanNetwork is a generated struct representing the https://schema.org/HealthPlanNetwork class.
type HealthPlanNetwork struct {
	HealthPlanCostSharing *Boolean `json:"healthplancostsharing"` //https://schema.org/Boolean
	HealthPlanNetworkId   *Text    `json:"healthplannetworkid"`   //https://schema.org/Text
	HealthPlanNetworkTier *Text    `json:"healthplannetworktier"` //https://schema.org/Text
}

// MedicalGuidelineRecommendation is a generated struct representing the https://schema.org/MedicalGuidelineRecommendation class.
type MedicalGuidelineRecommendation struct {
	RecommendationStrength *Text `json:"recommendationstrength"` //https://schema.org/Text
}

// ReadAction is a generated struct representing the https://schema.org/ReadAction class.
type ReadAction struct {
}

// VideoGameClip is a generated struct representing the https://schema.org/VideoGameClip class.
type VideoGameClip struct {
}

// AmusementPark is a generated struct representing the https://schema.org/AmusementPark class.
type AmusementPark struct {
}

// PlanAction is a generated struct representing the https://schema.org/PlanAction class.
type PlanAction struct {
	ScheduledTime *DateTime `json:"scheduledtime"` //https://schema.org/DateTime
}

// SocialEvent is a generated struct representing the https://schema.org/SocialEvent class.
type SocialEvent struct {
}

// SpeakableSpecification is a generated struct representing the https://schema.org/SpeakableSpecification class.
type SpeakableSpecification struct {
	CssSelector *CssSelectorType `json:"cssselector"` //https://schema.org/CssSelectorType
	Xpath       *XPathType       `json:"xpath"`       //https://schema.org/XPathType
}

// DataType is a generated struct representing the https://schema.org/DataType class.
type DataType struct {
}

// OrderItem is a generated struct representing the https://schema.org/OrderItem class.
type OrderItem struct {
	OrderDelivery   *ParcelDelivery `json:"orderdelivery"`   //https://schema.org/ParcelDelivery
	OrderItemNumber *Text           `json:"orderitemnumber"` //https://schema.org/Text
	OrderItemStatus *OrderStatus    `json:"orderitemstatus"` //https://schema.org/OrderStatus
	OrderQuantity   *Number         `json:"orderquantity"`   //https://schema.org/Number
	OrderedItem     *Service        `json:"ordereditem"`     //https://schema.org/Service
}

// TechArticle is a generated struct representing the https://schema.org/TechArticle class.
type TechArticle struct {
	Dependencies     *Text `json:"dependencies"`     //https://schema.org/Text
	ProficiencyLevel *Text `json:"proficiencylevel"` //https://schema.org/Text
}

// LandmarksOrHistoricalBuildings is a generated struct representing the https://schema.org/LandmarksOrHistoricalBuildings class.
type LandmarksOrHistoricalBuildings struct {
}

// LoanOrCredit is a generated struct representing the https://schema.org/LoanOrCredit class.
type LoanOrCredit struct {
	Amount             *Number                 `json:"amount"`             //https://schema.org/Number
	Currency           *Text                   `json:"currency"`           //https://schema.org/Text
	GracePeriod        *Duration               `json:"graceperiod"`        //https://schema.org/Duration
	LoanRepaymentForm  *RepaymentSpecification `json:"loanrepaymentform"`  //https://schema.org/RepaymentSpecification
	LoanTerm           *QuantitativeValue      `json:"loanterm"`           //https://schema.org/QuantitativeValue
	LoanType           *URL                    `json:"loantype"`           //https://schema.org/URL
	RecourseLoan       *Boolean                `json:"recourseloan"`       //https://schema.org/Boolean
	RenegotiableLoan   *Boolean                `json:"renegotiableloan"`   //https://schema.org/Boolean
	RequiredCollateral *Thing                  `json:"requiredcollateral"` //https://schema.org/Thing
}

// AMRadioChannel is a generated struct representing the https://schema.org/AMRadioChannel class.
type AMRadioChannel struct {
}

// AnalysisNewsArticle is a generated struct representing the https://schema.org/AnalysisNewsArticle class.
type AnalysisNewsArticle struct {
}

// BedAndBreakfast is a generated struct representing the https://schema.org/BedAndBreakfast class.
type BedAndBreakfast struct {
}

// DigitalDocument is a generated struct representing the https://schema.org/DigitalDocument class.
type DigitalDocument struct {
	HasDigitalDocumentPermission *DigitalDocumentPermission `json:"hasdigitaldocumentpermission"` //https://schema.org/DigitalDocumentPermission
}

// ExerciseAction is a generated struct representing the https://schema.org/ExerciseAction class.
type ExerciseAction struct {
	Course                 *Place                  `json:"course"`                 //https://schema.org/Place
	Diet                   *Diet                   `json:"diet"`                   //https://schema.org/Diet
	Distance               *Distance               `json:"distance"`               //https://schema.org/Distance
	ExercisePlan           *ExercisePlan           `json:"exerciseplan"`           //https://schema.org/ExercisePlan
	ExerciseRelatedDiet    *Diet                   `json:"exerciserelateddiet"`    //https://schema.org/Diet
	ExerciseType           *Text                   `json:"exercisetype"`           //https://schema.org/Text
	FromLocation           *Place                  `json:"fromlocation"`           //https://schema.org/Place
	Opponent               *Person                 `json:"opponent"`               //https://schema.org/Person
	SportsActivityLocation *SportsActivityLocation `json:"sportsactivitylocation"` //https://schema.org/SportsActivityLocation
	SportsEvent            *SportsEvent            `json:"sportsevent"`            //https://schema.org/SportsEvent
	SportsTeam             *SportsTeam             `json:"sportsteam"`             //https://schema.org/SportsTeam
	ToLocation             *Place                  `json:"tolocation"`             //https://schema.org/Place
	ExerciseCourse         *Place                  `json:"exercisecourse"`         //https://schema.org/Place
}

// GovernmentBenefitsType is a generated struct representing the https://schema.org/GovernmentBenefitsType class.
type GovernmentBenefitsType struct {
}

// AddAction is a generated struct representing the https://schema.org/AddAction class.
type AddAction struct {
}

// Volcano is a generated struct representing the https://schema.org/Volcano class.
type Volcano struct {
}

// ApartmentComplex is a generated struct representing the https://schema.org/ApartmentComplex class.
type ApartmentComplex struct {
	NumberOfAccommodationUnits          *QuantitativeValue `json:"numberofaccommodationunits"`          //https://schema.org/QuantitativeValue
	NumberOfAvailableAccommodationUnits *QuantitativeValue `json:"numberofavailableaccommodationunits"` //https://schema.org/QuantitativeValue
	NumberOfBedrooms                    *QuantitativeValue `json:"numberofbedrooms"`                    //https://schema.org/QuantitativeValue
	PetsAllowed                         *Text              `json:"petsallowed"`                         //https://schema.org/Text
	TourBookingPage                     *URL               `json:"tourbookingpage"`                     //https://schema.org/URL
}

// Clip is a generated struct representing the https://schema.org/Clip class.
type Clip struct {
	Actors        *Person             `json:"actors"`        //https://schema.org/Person
	ClipNumber    *Text               `json:"clipnumber"`    //https://schema.org/Text
	Directors     *Person             `json:"directors"`     //https://schema.org/Person
	EndOffset     *Number             `json:"endoffset"`     //https://schema.org/Number
	MusicBy       *Person             `json:"musicby"`       //https://schema.org/Person
	PartOfEpisode *Episode            `json:"partofepisode"` //https://schema.org/Episode
	PartOfSeason  *CreativeWorkSeason `json:"partofseason"`  //https://schema.org/CreativeWorkSeason
	StartOffset   *Number             `json:"startoffset"`   //https://schema.org/Number
	Director      *Person             `json:"director"`      //https://schema.org/Person
	PartOfSeries  *CreativeWorkSeries `json:"partofseries"`  //https://schema.org/CreativeWorkSeries
	Actor         *Person             `json:"actor"`         //https://schema.org/Person
}

// MenuItem is a generated struct representing the https://schema.org/MenuItem class.
type MenuItem struct {
	MenuAddOn       *MenuSection          `json:"menuaddon"`       //https://schema.org/MenuSection
	Nutrition       *NutritionInformation `json:"nutrition"`       //https://schema.org/NutritionInformation
	SuitableForDiet *RestrictedDiet       `json:"suitablefordiet"` //https://schema.org/RestrictedDiet
	Offers          *Offer                `json:"offers"`          //https://schema.org/Offer
}

// LegalValueLevel is a generated struct representing the https://schema.org/LegalValueLevel class.
type LegalValueLevel struct {
}

// LiteraryEvent is a generated struct representing the https://schema.org/LiteraryEvent class.
type LiteraryEvent struct {
}

// SizeSystemEnumeration is a generated struct representing the https://schema.org/SizeSystemEnumeration class.
type SizeSystemEnumeration struct {
}

// TennisComplex is a generated struct representing the https://schema.org/TennisComplex class.
type TennisComplex struct {
}

// Casino is a generated struct representing the https://schema.org/Casino class.
type Casino struct {
}

// RefundTypeEnumeration is a generated struct representing the https://schema.org/RefundTypeEnumeration class.
type RefundTypeEnumeration struct {
}

// Trip is a generated struct representing the https://schema.org/Trip class.
type Trip struct {
	ArrivalTime   *Time   `json:"arrivaltime"`   //https://schema.org/Time
	DepartureTime *Time   `json:"departuretime"` //https://schema.org/Time
	Itinerary     *Place  `json:"itinerary"`     //https://schema.org/Place
	TripOrigin    *Place  `json:"triporigin"`    //https://schema.org/Place
	Offers        *Offer  `json:"offers"`        //https://schema.org/Offer
	PartOfTrip    *Trip   `json:"partoftrip"`    //https://schema.org/Trip
	Provider      *Person `json:"provider"`      //https://schema.org/Person
	SubTrip       *Trip   `json:"subtrip"`       //https://schema.org/Trip
}

// GameServerStatus is a generated struct representing the https://schema.org/GameServerStatus class.
type GameServerStatus struct {
}

// PalliativeProcedure is a generated struct representing the https://schema.org/PalliativeProcedure class.
type PalliativeProcedure struct {
}

// PublicSwimmingPool is a generated struct representing the https://schema.org/PublicSwimmingPool class.
type PublicSwimmingPool struct {
}

// Cemetery is a generated struct representing the https://schema.org/Cemetery class.
type Cemetery struct {
}

// SearchRescueOrganization is a generated struct representing the https://schema.org/SearchRescueOrganization class.
type SearchRescueOrganization struct {
}

// UpdateAction is a generated struct representing the https://schema.org/UpdateAction class.
type UpdateAction struct {
	Collection       *Thing `json:"collection"`       //https://schema.org/Thing
	TargetCollection *Thing `json:"targetcollection"` //https://schema.org/Thing
}

// Book is a generated struct representing the https://schema.org/Book class.
type Book struct {
	Abridged      *Boolean        `json:"abridged"`      //https://schema.org/Boolean
	BookEdition   *Text           `json:"bookedition"`   //https://schema.org/Text
	BookFormat    *BookFormatType `json:"bookformat"`    //https://schema.org/BookFormatType
	Illustrator   *Person         `json:"illustrator"`   //https://schema.org/Person
	Isbn          *Text           `json:"isbn"`          //https://schema.org/Text
	NumberOfPages *Integer        `json:"numberofpages"` //https://schema.org/Integer
}

// FireStation is a generated struct representing the https://schema.org/FireStation class.
type FireStation struct {
}

// CorrectionComment is a generated struct representing the https://schema.org/CorrectionComment class.
type CorrectionComment struct {
}

// FloorPlan is a generated struct representing the https://schema.org/FloorPlan class.
type FloorPlan struct {
	AmenityFeature                      *LocationFeatureSpecification `json:"amenityfeature"`                      //https://schema.org/LocationFeatureSpecification
	FloorSize                           *QuantitativeValue            `json:"floorsize"`                           //https://schema.org/QuantitativeValue
	IsPlanForApartment                  *Accommodation                `json:"isplanforapartment"`                  //https://schema.org/Accommodation
	LayoutImage                         *URL                          `json:"layoutimage"`                         //https://schema.org/URL
	NumberOfAccommodationUnits          *QuantitativeValue            `json:"numberofaccommodationunits"`          //https://schema.org/QuantitativeValue
	NumberOfAvailableAccommodationUnits *QuantitativeValue            `json:"numberofavailableaccommodationunits"` //https://schema.org/QuantitativeValue
	NumberOfBathroomsTotal              *Integer                      `json:"numberofbathroomstotal"`              //https://schema.org/Integer
	NumberOfBedrooms                    *QuantitativeValue            `json:"numberofbedrooms"`                    //https://schema.org/QuantitativeValue
	NumberOfFullBathrooms               *Number                       `json:"numberoffullbathrooms"`               //https://schema.org/Number
	NumberOfPartialBathrooms            *Number                       `json:"numberofpartialbathrooms"`            //https://schema.org/Number
	NumberOfRooms                       *QuantitativeValue            `json:"numberofrooms"`                       //https://schema.org/QuantitativeValue
	PetsAllowed                         *Text                         `json:"petsallowed"`                         //https://schema.org/Text
}

// MedicalProcedure is a generated struct representing the https://schema.org/MedicalProcedure class.
type MedicalProcedure struct {
	BodyLocation  *Text                 `json:"bodylocation"`  //https://schema.org/Text
	Followup      *Text                 `json:"followup"`      //https://schema.org/Text
	HowPerformed  *Text                 `json:"howperformed"`  //https://schema.org/Text
	Preparation   *Text                 `json:"preparation"`   //https://schema.org/Text
	ProcedureType *MedicalProcedureType `json:"proceduretype"` //https://schema.org/MedicalProcedureType
	Status        *Text                 `json:"status"`        //https://schema.org/Text
}

// MedicalSignOrSymptom is a generated struct representing the https://schema.org/MedicalSignOrSymptom class.
type MedicalSignOrSymptom struct {
	PossibleTreatment *MedicalTherapy `json:"possibletreatment"` //https://schema.org/MedicalTherapy
}

// PodcastSeries is a generated struct representing the https://schema.org/PodcastSeries class.
type PodcastSeries struct {
	WebFeed *URL    `json:"webfeed"` //https://schema.org/URL
	Actor   *Person `json:"actor"`   //https://schema.org/Person
}

// AdministrativeArea is a generated struct representing the https://schema.org/AdministrativeArea class.
type AdministrativeArea struct {
}

// BoatReservation is a generated struct representing the https://schema.org/BoatReservation class.
type BoatReservation struct {
}

// DeliveryChargeSpecification is a generated struct representing the https://schema.org/DeliveryChargeSpecification class.
type DeliveryChargeSpecification struct {
	AppliesToDeliveryMethod *DeliveryMethod `json:"appliestodeliverymethod"` //https://schema.org/DeliveryMethod
	EligibleRegion          *Text           `json:"eligibleregion"`          //https://schema.org/Text
	IneligibleRegion        *Text           `json:"ineligibleregion"`        //https://schema.org/Text
	AreaServed              *Text           `json:"areaserved"`              //https://schema.org/Text
}

// Library is a generated struct representing the https://schema.org/Library class.
type Library struct {
}

// MedicalTrial is a generated struct representing the https://schema.org/MedicalTrial class.
type MedicalTrial struct {
	TrialDesign *MedicalTrialDesign `json:"trialdesign"` //https://schema.org/MedicalTrialDesign
}

// RadioClip is a generated struct representing the https://schema.org/RadioClip class.
type RadioClip struct {
}

// ConsumeAction is a generated struct representing the https://schema.org/ConsumeAction class.
type ConsumeAction struct {
	ActionAccessibilityRequirement *ActionAccessSpecification `json:"actionaccessibilityrequirement"` //https://schema.org/ActionAccessSpecification
	ExpectsAcceptanceOf            *Offer                     `json:"expectsacceptanceof"`            //https://schema.org/Offer
}

// Organization is a generated struct representing the https://schema.org/Organization class.
type Organization struct {
	ActionableFeedbackPolicy  *URL                               `json:"actionablefeedbackpolicy"`  //https://schema.org/URL
	Address                   *Text                              `json:"address"`                   //https://schema.org/Text
	AgentInteractionStatistic *InteractionCounter                `json:"agentinteractionstatistic"` //https://schema.org/InteractionCounter
	AggregateRating           *AggregateRating                   `json:"aggregaterating"`           //https://schema.org/AggregateRating
	Awards                    *Text                              `json:"awards"`                    //https://schema.org/Text
	Brand                     *Organization                      `json:"brand"`                     //https://schema.org/Organization
	ContactPoints             *ContactPoint                      `json:"contactpoints"`             //https://schema.org/ContactPoint
	CorrectionsPolicy         *URL                               `json:"correctionspolicy"`         //https://schema.org/URL
	Department                *Organization                      `json:"department"`                //https://schema.org/Organization
	DissolutionDate           *Date                              `json:"dissolutiondate"`           //https://schema.org/Date
	DiversityPolicy           *URL                               `json:"diversitypolicy"`           //https://schema.org/URL
	DiversityStaffingReport   *URL                               `json:"diversitystaffingreport"`   //https://schema.org/URL
	Duns                      *Text                              `json:"duns"`                      //https://schema.org/Text
	Email                     *Text                              `json:"email"`                     //https://schema.org/Text
	Employees                 *Person                            `json:"employees"`                 //https://schema.org/Person
	EthicsPolicy              *URL                               `json:"ethicspolicy"`              //https://schema.org/URL
	Events                    *Event                             `json:"events"`                    //https://schema.org/Event
	FaxNumber                 *Text                              `json:"faxnumber"`                 //https://schema.org/Text
	Founders                  *Person                            `json:"founders"`                  //https://schema.org/Person
	FoundingDate              *Date                              `json:"foundingdate"`              //https://schema.org/Date
	FoundingLocation          *Place                             `json:"foundinglocation"`          //https://schema.org/Place
	Funder                    *Person                            `json:"funder"`                    //https://schema.org/Person
	GlobalLocationNumber      *Text                              `json:"globallocationnumber"`      //https://schema.org/Text
	HasCredential             *EducationalOccupationalCredential `json:"hascredential"`             //https://schema.org/EducationalOccupationalCredential
	HasMerchantReturnPolicy   *MerchantReturnPolicy              `json:"hasmerchantreturnpolicy"`   //https://schema.org/MerchantReturnPolicy
	HasOfferCatalog           *OfferCatalog                      `json:"hasoffercatalog"`           //https://schema.org/OfferCatalog
	HasPOS                    *Place                             `json:"haspos"`                    //https://schema.org/Place
	IsicV4                    *Text                              `json:"isicv4"`                    //https://schema.org/Text
	Iso6523Code               *Text                              `json:"iso6523code"`               //https://schema.org/Text
	Keywords                  *URL                               `json:"keywords"`                  //https://schema.org/URL
	KnowsAbout                *URL                               `json:"knowsabout"`                //https://schema.org/URL
	KnowsLanguage             *Text                              `json:"knowslanguage"`             //https://schema.org/Text
	LegalName                 *Text                              `json:"legalname"`                 //https://schema.org/Text
	LeiCode                   *Text                              `json:"leicode"`                   //https://schema.org/Text
	Logo                      *URL                               `json:"logo"`                      //https://schema.org/URL
	Members                   *Person                            `json:"members"`                   //https://schema.org/Person
	Naics                     *Text                              `json:"naics"`                     //https://schema.org/Text
	NonprofitStatus           *NonprofitType                     `json:"nonprofitstatus"`           //https://schema.org/NonprofitType
	NumberOfEmployees         *QuantitativeValue                 `json:"numberofemployees"`         //https://schema.org/QuantitativeValue
	OwnershipFundingInfo      *URL                               `json:"ownershipfundinginfo"`      //https://schema.org/URL
	Owns                      *Product                           `json:"owns"`                      //https://schema.org/Product
	Reviews                   *Review                            `json:"reviews"`                   //https://schema.org/Review
	Seeks                     *Demand                            `json:"seeks"`                     //https://schema.org/Demand
	Slogan                    *Text                              `json:"slogan"`                    //https://schema.org/Text
	TaxID                     *Text                              `json:"taxid"`                     //https://schema.org/Text
	Telephone                 *Text                              `json:"telephone"`                 //https://schema.org/Text
	UnnamedSourcesPolicy      *URL                               `json:"unnamedsourcespolicy"`      //https://schema.org/URL
	VatID                     *Text                              `json:"vatid"`                     //https://schema.org/Text
	Alumni                    *Person                            `json:"alumni"`                    //https://schema.org/Person
	Award                     *Text                              `json:"award"`                     //https://schema.org/Text
	ContactPoint              *ContactPoint                      `json:"contactpoint"`              //https://schema.org/ContactPoint
	Employee                  *Person                            `json:"employee"`                  //https://schema.org/Person
	Event                     *Event                             `json:"event"`                     //https://schema.org/Event
	Founder                   *Person                            `json:"founder"`                   //https://schema.org/Person
	Funding                   *Grant                             `json:"funding"`                   //https://schema.org/Grant
	InteractionStatistic      *InteractionCounter                `json:"interactionstatistic"`      //https://schema.org/InteractionCounter
	MakesOffer                *Offer                             `json:"makesoffer"`                //https://schema.org/Offer
	Review                    *Review                            `json:"review"`                    //https://schema.org/Review
	ServiceArea               *Place                             `json:"servicearea"`               //https://schema.org/Place
	Sponsor                   *Person                            `json:"sponsor"`                   //https://schema.org/Person
	SubOrganization           *Organization                      `json:"suborganization"`           //https://schema.org/Organization
	MemberOf                  *ProgramMembership                 `json:"memberof"`                  //https://schema.org/ProgramMembership
	ParentOrganization        *Organization                      `json:"parentorganization"`        //https://schema.org/Organization
	AreaServed                *Text                              `json:"areaserved"`                //https://schema.org/Text
	Member                    *Person                            `json:"member"`                    //https://schema.org/Person
	PublishingPrinciples      *URL                               `json:"publishingprinciples"`      //https://schema.org/URL
	Location                  *VirtualLocation                   `json:"location"`                  //https://schema.org/VirtualLocation
}

// DeliveryMethod is a generated struct representing the https://schema.org/DeliveryMethod class.
type DeliveryMethod struct {
}

// MedicalProcedureType is a generated struct representing the https://schema.org/MedicalProcedureType class.
type MedicalProcedureType struct {
}

// Boolean is a generated struct representing the https://schema.org/Boolean class.
type Boolean struct {
}

// Diet is a generated struct representing the https://schema.org/Diet class.
type Diet struct {
	DietFeatures          *Text   `json:"dietfeatures"`          //https://schema.org/Text
	Endorsers             *Person `json:"endorsers"`             //https://schema.org/Person
	ExpertConsiderations  *Text   `json:"expertconsiderations"`  //https://schema.org/Text
	PhysiologicalBenefits *Text   `json:"physiologicalbenefits"` //https://schema.org/Text
	Risks                 *Text   `json:"risks"`                 //https://schema.org/Text
}

// InsertAction is a generated struct representing the https://schema.org/InsertAction class.
type InsertAction struct {
	ToLocation *Place `json:"tolocation"` //https://schema.org/Place
}

// ListItem is a generated struct representing the https://schema.org/ListItem class.
type ListItem struct {
	Item         *Thing    `json:"item"`         //https://schema.org/Thing
	NextItem     *ListItem `json:"nextitem"`     //https://schema.org/ListItem
	PreviousItem *ListItem `json:"previousitem"` //https://schema.org/ListItem
	Position     *Text     `json:"position"`     //https://schema.org/Text
}

// LocalBusiness is a generated struct representing the https://schema.org/LocalBusiness class.
type LocalBusiness struct {
	BranchOf           *Organization `json:"branchof"`           //https://schema.org/Organization
	CurrenciesAccepted *Text         `json:"currenciesaccepted"` //https://schema.org/Text
	OpeningHours       *Text         `json:"openinghours"`       //https://schema.org/Text
	PaymentAccepted    *Text         `json:"paymentaccepted"`    //https://schema.org/Text
	PriceRange         *Text         `json:"pricerange"`         //https://schema.org/Text
}

// MotorizedBicycle is a generated struct representing the https://schema.org/MotorizedBicycle class.
type MotorizedBicycle struct {
}

// ReactAction is a generated struct representing the https://schema.org/ReactAction class.
type ReactAction struct {
}

// BodyMeasurementTypeEnumeration is a generated struct representing the https://schema.org/BodyMeasurementTypeEnumeration class.
type BodyMeasurementTypeEnumeration struct {
}

// Preschool is a generated struct representing the https://schema.org/Preschool class.
type Preschool struct {
}

// Sculpture is a generated struct representing the https://schema.org/Sculpture class.
type Sculpture struct {
}

// VideoObject is a generated struct representing the https://schema.org/VideoObject class.
type VideoObject struct {
	Actors              *Person `json:"actors"`              //https://schema.org/Person
	Directors           *Person `json:"directors"`           //https://schema.org/Person
	EmbeddedTextCaption *Text   `json:"embeddedtextcaption"` //https://schema.org/Text
	MusicBy             *Person `json:"musicby"`             //https://schema.org/Person
	Transcript          *Text   `json:"transcript"`          //https://schema.org/Text
	VideoFrameSize      *Text   `json:"videoframesize"`      //https://schema.org/Text
	VideoQuality        *Text   `json:"videoquality"`        //https://schema.org/Text
	Caption             *Text   `json:"caption"`             //https://schema.org/Text
	Director            *Person `json:"director"`            //https://schema.org/Person
	Actor               *Person `json:"actor"`               //https://schema.org/Person
}

// AskPublicNewsArticle is a generated struct representing the https://schema.org/AskPublicNewsArticle class.
type AskPublicNewsArticle struct {
}

// PathologyTest is a generated struct representing the https://schema.org/PathologyTest class.
type PathologyTest struct {
	TissueSample *Text `json:"tissuesample"` //https://schema.org/Text
}

// Place is a generated struct representing the https://schema.org/Place class.
type Place struct {
	AdditionalProperty               *PropertyValue                `json:"additionalproperty"`               //https://schema.org/PropertyValue
	Address                          *Text                         `json:"address"`                          //https://schema.org/Text
	AggregateRating                  *AggregateRating              `json:"aggregaterating"`                  //https://schema.org/AggregateRating
	AmenityFeature                   *LocationFeatureSpecification `json:"amenityfeature"`                   //https://schema.org/LocationFeatureSpecification
	BranchCode                       *Text                         `json:"branchcode"`                       //https://schema.org/Text
	ContainedIn                      *Place                        `json:"containedin"`                      //https://schema.org/Place
	Events                           *Event                        `json:"events"`                           //https://schema.org/Event
	FaxNumber                        *Text                         `json:"faxnumber"`                        //https://schema.org/Text
	Geo                              *GeoShape                     `json:"geo"`                              //https://schema.org/GeoShape
	GeoContains                      *Place                        `json:"geocontains"`                      //https://schema.org/Place
	GeoCoveredBy                     *Place                        `json:"geocoveredby"`                     //https://schema.org/Place
	GeoCovers                        *Place                        `json:"geocovers"`                        //https://schema.org/Place
	GeoCrosses                       *Place                        `json:"geocrosses"`                       //https://schema.org/Place
	GeoDisjoint                      *Place                        `json:"geodisjoint"`                      //https://schema.org/Place
	GeoEquals                        *Place                        `json:"geoequals"`                        //https://schema.org/Place
	GeoIntersects                    *Place                        `json:"geointersects"`                    //https://schema.org/Place
	GeoOverlaps                      *Place                        `json:"geooverlaps"`                      //https://schema.org/Place
	GeoTouches                       *Place                        `json:"geotouches"`                       //https://schema.org/Place
	GeoWithin                        *Place                        `json:"geowithin"`                        //https://schema.org/Place
	GlobalLocationNumber             *Text                         `json:"globallocationnumber"`             //https://schema.org/Text
	HasDriveThroughService           *Boolean                      `json:"hasdrivethroughservice"`           //https://schema.org/Boolean
	IsicV4                           *Text                         `json:"isicv4"`                           //https://schema.org/Text
	Keywords                         *URL                          `json:"keywords"`                         //https://schema.org/URL
	Latitude                         *Text                         `json:"latitude"`                         //https://schema.org/Text
	Logo                             *URL                          `json:"logo"`                             //https://schema.org/URL
	Longitude                        *Text                         `json:"longitude"`                        //https://schema.org/Text
	Map                              *URL                          `json:"map"`                              //https://schema.org/URL
	Maps                             *URL                          `json:"maps"`                             //https://schema.org/URL
	MaximumAttendeeCapacity          *Integer                      `json:"maximumattendeecapacity"`          //https://schema.org/Integer
	OpeningHoursSpecification        *OpeningHoursSpecification    `json:"openinghoursspecification"`        //https://schema.org/OpeningHoursSpecification
	Photos                           *Photograph                   `json:"photos"`                           //https://schema.org/Photograph
	PublicAccess                     *Boolean                      `json:"publicaccess"`                     //https://schema.org/Boolean
	Reviews                          *Review                       `json:"reviews"`                          //https://schema.org/Review
	Slogan                           *Text                         `json:"slogan"`                           //https://schema.org/Text
	SmokingAllowed                   *Boolean                      `json:"smokingallowed"`                   //https://schema.org/Boolean
	SpecialOpeningHoursSpecification *OpeningHoursSpecification    `json:"specialopeninghoursspecification"` //https://schema.org/OpeningHoursSpecification
	Telephone                        *Text                         `json:"telephone"`                        //https://schema.org/Text
	TourBookingPage                  *URL                          `json:"tourbookingpage"`                  //https://schema.org/URL
	ContainsPlace                    *Place                        `json:"containsplace"`                    //https://schema.org/Place
	Event                            *Event                        `json:"event"`                            //https://schema.org/Event
	IsAccessibleForFree              *Boolean                      `json:"isaccessibleforfree"`              //https://schema.org/Boolean
	Photo                            *Photograph                   `json:"photo"`                            //https://schema.org/Photograph
	Review                           *Review                       `json:"review"`                           //https://schema.org/Review
	ContainedInPlace                 *Place                        `json:"containedinplace"`                 //https://schema.org/Place
	HasMap                           *URL                          `json:"hasmap"`                           //https://schema.org/URL
}

// WatchAction is a generated struct representing the https://schema.org/WatchAction class.
type WatchAction struct {
}

// Brand is a generated struct representing the https://schema.org/Brand class.
type Brand struct {
	AggregateRating *AggregateRating `json:"aggregaterating"` //https://schema.org/AggregateRating
	Logo            *URL             `json:"logo"`            //https://schema.org/URL
	Slogan          *Text            `json:"slogan"`          //https://schema.org/Text
	Review          *Review          `json:"review"`          //https://schema.org/Review
}

// ItemPage is a generated struct representing the https://schema.org/ItemPage class.
type ItemPage struct {
}

// BeautySalon is a generated struct representing the https://schema.org/BeautySalon class.
type BeautySalon struct {
}

// DrugStrength is a generated struct representing the https://schema.org/DrugStrength class.
type DrugStrength struct {
	ActiveIngredient *Text                `json:"activeingredient"` //https://schema.org/Text
	AvailableIn      *AdministrativeArea  `json:"availablein"`      //https://schema.org/AdministrativeArea
	MaximumIntake    *MaximumDoseSchedule `json:"maximumintake"`    //https://schema.org/MaximumDoseSchedule
	StrengthUnit     *Text                `json:"strengthunit"`     //https://schema.org/Text
	StrengthValue    *Number              `json:"strengthvalue"`    //https://schema.org/Number
}

// PrependAction is a generated struct representing the https://schema.org/PrependAction class.
type PrependAction struct {
}

// Article is a generated struct representing the https://schema.org/Article class.
type Article struct {
	ArticleBody    *Text    `json:"articlebody"`    //https://schema.org/Text
	ArticleSection *Text    `json:"articlesection"` //https://schema.org/Text
	Backstory      *Text    `json:"backstory"`      //https://schema.org/Text
	PageEnd        *Text    `json:"pageend"`        //https://schema.org/Text
	PageStart      *Text    `json:"pagestart"`      //https://schema.org/Text
	Pagination     *Text    `json:"pagination"`     //https://schema.org/Text
	Speakable      *URL     `json:"speakable"`      //https://schema.org/URL
	WordCount      *Integer `json:"wordcount"`      //https://schema.org/Integer
}

// ExerciseGym is a generated struct representing the https://schema.org/ExerciseGym class.
type ExerciseGym struct {
}

// NGO is a generated struct representing the https://schema.org/NGO class.
type NGO struct {
}

// RadioSeason is a generated struct representing the https://schema.org/RadioSeason class.
type RadioSeason struct {
}

// Restaurant is a generated struct representing the https://schema.org/Restaurant class.
type Restaurant struct {
}

// AcceptAction is a generated struct representing the https://schema.org/AcceptAction class.
type AcceptAction struct {
}

// InsuranceAgency is a generated struct representing the https://schema.org/InsuranceAgency class.
type InsuranceAgency struct {
}

// OnlineBusiness is a generated struct representing the https://schema.org/OnlineBusiness class.
type OnlineBusiness struct {
}

// QualitativeValue is a generated struct representing the https://schema.org/QualitativeValue class.
type QualitativeValue struct {
	AdditionalProperty *PropertyValue    `json:"additionalproperty"` //https://schema.org/PropertyValue
	Equal              *QualitativeValue `json:"equal"`              //https://schema.org/QualitativeValue
	Greater            *QualitativeValue `json:"greater"`            //https://schema.org/QualitativeValue
	GreaterOrEqual     *QualitativeValue `json:"greaterorequal"`     //https://schema.org/QualitativeValue
	Lesser             *QualitativeValue `json:"lesser"`             //https://schema.org/QualitativeValue
	LesserOrEqual      *QualitativeValue `json:"lesserorequal"`      //https://schema.org/QualitativeValue
	NonEqual           *QualitativeValue `json:"nonequal"`           //https://schema.org/QualitativeValue
	ValueReference     *Text             `json:"valuereference"`     //https://schema.org/Text
}

// RadioBroadcastService is a generated struct representing the https://schema.org/RadioBroadcastService class.
type RadioBroadcastService struct {
}

// AppendAction is a generated struct representing the https://schema.org/AppendAction class.
type AppendAction struct {
}

// MedicalGuidelineContraindication is a generated struct representing the https://schema.org/MedicalGuidelineContraindication class.
type MedicalGuidelineContraindication struct {
}

// SeaBodyOfWater is a generated struct representing the https://schema.org/SeaBodyOfWater class.
type SeaBodyOfWater struct {
}

// Canal is a generated struct representing the https://schema.org/Canal class.
type Canal struct {
}

// Duration is a generated struct representing the https://schema.org/Duration class.
type Duration struct {
}

// GardenStore is a generated struct representing the https://schema.org/GardenStore class.
type GardenStore struct {
}

// MedicalAudienceType is a generated struct representing the https://schema.org/MedicalAudienceType class.
type MedicalAudienceType struct {
}

// MedicalSign is a generated struct representing the https://schema.org/MedicalSign class.
type MedicalSign struct {
	IdentifyingExam *PhysicalExam `json:"identifyingexam"` //https://schema.org/PhysicalExam
	IdentifyingTest *MedicalTest  `json:"identifyingtest"` //https://schema.org/MedicalTest
}

// ResumeAction is a generated struct representing the https://schema.org/ResumeAction class.
type ResumeAction struct {
}

// MedicalRiskEstimator is a generated struct representing the https://schema.org/MedicalRiskEstimator class.
type MedicalRiskEstimator struct {
	EstimatesRiskOf    *MedicalEntity     `json:"estimatesriskof"`    //https://schema.org/MedicalEntity
	IncludedRiskFactor *MedicalRiskFactor `json:"includedriskfactor"` //https://schema.org/MedicalRiskFactor
}

// APIReference is a generated struct representing the https://schema.org/APIReference class.
type APIReference struct {
	Assembly              *Text `json:"assembly"`              //https://schema.org/Text
	AssemblyVersion       *Text `json:"assemblyversion"`       //https://schema.org/Text
	ProgrammingModel      *Text `json:"programmingmodel"`      //https://schema.org/Text
	TargetPlatform        *Text `json:"targetplatform"`        //https://schema.org/Text
	ExecutableLibraryName *Text `json:"executablelibraryname"` //https://schema.org/Text
}

// Chapter is a generated struct representing the https://schema.org/Chapter class.
type Chapter struct {
	PageEnd    *Text `json:"pageend"`    //https://schema.org/Text
	PageStart  *Text `json:"pagestart"`  //https://schema.org/Text
	Pagination *Text `json:"pagination"` //https://schema.org/Text
}

// MortgageLoan is a generated struct representing the https://schema.org/MortgageLoan class.
type MortgageLoan struct {
	DomiciledMortgage         *Boolean        `json:"domiciledmortgage"`         //https://schema.org/Boolean
	LoanMortgageMandateAmount *MonetaryAmount `json:"loanmortgagemandateamount"` //https://schema.org/MonetaryAmount
}

// Project is a generated struct representing the https://schema.org/Project class.
type Project struct {
}

// RentalCarReservation is a generated struct representing the https://schema.org/RentalCarReservation class.
type RentalCarReservation struct {
	DropoffLocation *Place    `json:"dropofflocation"` //https://schema.org/Place
	DropoffTime     *DateTime `json:"dropofftime"`     //https://schema.org/DateTime
	PickupLocation  *Place    `json:"pickuplocation"`  //https://schema.org/Place
	PickupTime      *DateTime `json:"pickuptime"`      //https://schema.org/DateTime
}
