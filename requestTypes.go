package gta_sdk

import (
	"time"
)

type RequestModeEnum string
const (
	SYNCHRONOUS = "SYNCHRONOUS"
	ASYNCHRONOUS = "ASYNCHRONOUS"
)

type DestinationTypeEnum string
const (
	AREA = "area"
	CITY = "city"
	GEOCODE = "geocode"
)

type OrderByPriceEnum string
const (
	PRICELOWTOHIGH = "pricelowtohigh"
	PRICEHIGHTOLOW = "pricehightolow"
)

// Auth and request preferences
type Source struct {
	RequestorID				RequestorID				`xml:"RequestorID"`
	RequestorPreferences	RequestorPreferences	`xml:"RequestorPreferences"`
}

type RequestorPreferences struct {
	Language	string			`xml:"Language,attr"`
	Currency	string			`xml:"Currency,attr"`
	Country		string			`xml:"Country,attr"`
	RequestMode	RequestModeEnum	`xml:"RequestMode"`
	ResponseURL	string			`xml:"ResponseUrl,omitempty"`
}

type RequestorID struct {
	Client			string	`xml:"Client,attr"`
	EMailAddress	string	`xml:"EMailAddress,attr"`
	Password		string	`xml:"Password,attr"`
}
// End Auth and request preferences

// Common Response
type Errors struct {
	Errors	[]Error	`xml:"Errors"`
}

type Error struct {
	ErrorId		string	`xml:"ErrorId"`
	ErrorText	string	`xml:"ErrorText"`
}

type Response struct {
	ResponseReference	string			`xml:"ResponseReference,attr"`
	ResponseSequence	string			`xml:"ResponseSequence,attr,omitempty"`
	ResponseDetails		ResponseDetails	`xml:"ResponseDetails"`
}

type ResponseDetails struct {
	Language				string					`xml:"Language,attr"`
	Errors					*Errors					`xml:"Errors,omitempty"`
	ResponseItems			interface{}
}
// End Common Response

// Common Requests
type Request struct {
	Source			Source			`xml:"Source"`
	RequestDetails	RequestDetails	`xml:"RequestDetails"`
}

type RequestDetails struct {
	RequestItems interface{}
}
// End Common Response

//  Request Search Hotel Price
type SearchHotelPriceRequest struct {
	ItemDestination				ItemDestination			`xml:"ItemDestination"`
	ImmediateConfirmationOnly	bool					`xml:"ImmediateConfirmationOnly,omitempty"`
	ItemName					string					`xml:"ItemName,omitempty"`
	ItemCodes					ItemCodes				`xml:"ItemCodes,omitempty"`
	ItemCode					ItemCode				`xml:"ItemCode,omitempty"`
	PeriodOfStay				PeriodOfStay			`xml:"PeriodOfStay"`
	IncludeRecommended			bool					`xml:"IncludeRecommended"`
	RecommendedOnly				bool					`xml:"RecommendedOnly,omitempty"`
	IncludePriceBreakdown		bool					`xml:"IncludePriceBreakdown"`
	IncludeChargeConditions		bool					`xml:"IncludeChargeConditions,omitempty"`
	IncludeChargeableItems		IncludeChargeConditions	`xml:"IncludeChargeableItems"`
	ExcludeChargeableItems		ExcludeChargeableItems	`xml:"ExcludeChargeableItems"`
	Rooms						[]Room					`xml:"Rooms"`
	StarRating					StarRating				`xml:"MinimumRating,omitempty"`
	StarRatingRange				StarRatingRange			`xml:"StarRatingRange,omitempty"`
	LocationCode				string					`xml:"LocationCode,omitempty"`
	FacilityCodes				[]FacilityCodes			`xml:"FacilityCodes,omitempty"`
	OrderBy						OrderByPriceEnum		`xml:"OrderBy,omitempty"`
	NumberOfReturnedItems		int						`xml:"NumberOfReturnedItems,omitempty"`
}

type StarRatingRange struct {
	Min	int	`xml:"Min"`
	Max	int	`xml:"Max"`
}

type IncludeChargeConditions struct {
	DateFormatResponse	bool	`xml:"DateFormatResponse"`
}

type FacilityCodes struct {
	FacilityCode	string	`xml:"FacilityCode"`
}

type StarRating struct {
	MinimumRating	bool	`xml:"MinimumRating,attr"`
	Value			int		`xml:",chardata"`
}

type Room struct {
	Code			string		`xml:"Code,attr"`
	Id				string		`xml:"Id,attr,omitempty"`
	NumberOfRooms	string		`xml:"NumberOfRooms,attr,omitempty"`
	NumberOfCots	string		`xml:"NumberOfCots,attr,omitempty"`
	ExtraBeds		ExtraBeds	`xml:"ExtraBeds"`
}

type ExtraBeds struct {
	Age	int	`xml:"Age"`
}

type ExcludeChargeableItems struct {
	CancellationDeadlineHours	int	`xml:"CancellationDeadlineHours,omitempty"`
	CancellationDeadlineDays	int	`xml:"CancellationDeadlineDays,omitempty"`
}

type PeriodOfStay struct {
	CheckInDate	time.Time	`xml:"CheckInDate"`
	Duration	int			`xml:"Duration"`
}

type ItemCodes struct {
	ItemCodes	[]ItemCode	`xml:"ItemCodes"`
}

type ItemCode struct {
	ItemCode	string	`xml:"ItemCode"`
}

type ItemDestination struct {
	DestinationType	DestinationTypeEnum	`xml:"DestinationType,attr"`
	DestinationCode	string				`xml:"DestinationCode,attr,omitempty"`
	Latitude		string				`xml:"Latitude,omitempty"`
	Longitude		string				`xml:"Longitude,omitempty"`
	RadiusKm		string				`xml:"RadiusKm,omitempty"`
	WestLongitude	string				`xml:"WestLongitude,omitempty"`
	SouthLatitude	string				`xml:"SouthLatitude,omitempty"`
	EastLongitude	string				`xml:"EastLongitude,omitempty"`
	NorthLatitude	string				`xml:"NorthLatitude,omitempty"`
}
// End  Request Search Hotel Price
