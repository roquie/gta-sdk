package gta_sdk

import (
	"encoding/xml"
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
	XMLName				xml.Name		`xml:"Response"`
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
	XMLName			xml.Name		`xml:"Request"`
	Source			Source			`xml:"Source"`
	RequestDetails	RequestDetails	`xml:"RequestDetails"`
}

type RequestDetails struct {
	RequestItems interface{}
}
// End Common Response

//  Request Search Hotel Price
type SearchHotelPriceRequest struct {
	XMLName   					xml.Name 				`xml:"SearchHotelPriceRequest"`
	ItemDestination				ItemDestination			`xml:"ItemDestination"`
	ImmediateConfirmationOnly	bool					`xml:"ImmediateConfirmationOnly,omitempty"`
	ItemName					string					`xml:"ItemName,omitempty"`
	ItemCodes					ItemCodes				`xml:"ItemCodes,omitempty"`
	ItemCode					ItemCode				`xml:"ItemCode,omitempty"`
	PeriodOfStay				PeriodOfStay			`xml:"PeriodOfStay"`
	IncludeRecommended			bool					`xml:"IncludeRecommended"`
	RecommendedOnly				bool					`xml:"RecommendedOnly,omitempty"`
	IncludePriceBreakdown		bool					`xml:"IncludePriceBreakdown"`
	IncludeChargeConditions		IncludeChargeConditions	`xml:"IncludeChargeConditions,omitempty"`
	IncludeChargeableItems		bool					`xml:"IncludeChargeableItems"`
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

// Response Search Hotel Price
type SearchHotelPriceResponse struct {
	HotelDetails	HotelDetails	`xml:"HotelDetails"`
}

type HotelDetails struct {
	Hotel	[]Hotel	`xml:"Hotel"`
}

type Hotel struct {
	HasExtraInfo	bool			`xml:"HasExtraInfo,attr"`
	HasMap			bool			`xml:"HasMap,attr"`
	HasPictures		bool			`xml:"HasPictures,attr"`
	Recommended		bool			`xml:"Recommended,attr,omitempty"`
	City			City			`xml:"City,cdata"`
	Item			Item			`xml:"Item,cdata"`
	LocationDetails	LocationDetails	`xml:"LocationDetails,cdata"`
	StarRating		int				`xml:"StarRating"`
	HotelRooms		[]HotelRoom		`xml:"HotelRooms"`
	RoomCategories	[]RoomCategory	`xml:"RoomCategories"`
}

type RoomCategory struct {
	XMLName			xml.Name			`xml:"RoomCategory"`
	Id				string				`xml:"Id,attr"`
	Description		string				`xml:"Description,cdata"`
	ItemPrice		ItemPrice			`xml:"ItemPrice"`
	Confirmation	Confirmation		`xml:"Confirmation,cdata"`
	SharingBedding	bool				`xml:"SharingBedding"`
	Meals			Meals				`xml:"Meals"`
	HotelRoomPrices	[]HotelRoomPrice	`xml:"HotelRoomPrices"`
}

type HotelRoomPrice struct {
	XMLName		xml.Name	`xml:"HotelRoom"`
	Code		string		`xml:"Code,attr"`
	RoomPrice	RoomPrice	`xml:"RoomPrice"`
}

type RoomPrice struct {
	Gross		string			`xml:"Gross,attr"`
	PriceRanges	[]PriceRange	`xml:"PriceRanges"`
}

type PriceRange struct {
	DateRange	DataRange	`xml:"DateRange"`
}

type DataRange struct {
	FromDate	time.Time	`xml:"FromDate"`
	ToDate		time.Time	`xml:"ToDate"`
}

type Meals struct {
	Basis		Basis		`xml:"Basis,cdata"`
	Breakfast	Breakfast	`xml:"Breakfast,cdata"`
}

type Breakfast struct {
	Code	string	`xml:"Code,attr"`
}

type Basis struct {
	Code	string	`xml:"Code,attr"`
}

type Confirmation struct {
	Code	string	`xml:"Code,attr"`
}

type ItemPrice struct {
	Currency				string	`xml:"Currency,attr"`
	CommissionIndicator		string	`xml:"CommissionIndicator,attr"`
	CommissionPercentage	string	`xml:"CommissionPercentage,attr"`
	NoOfferDiscount			bool	`xml:"NoOfferDiscount,attr,omitempty"`
	IncludedOfferDiscount	bool	`xml:"IncludedOfferDiscount,attr,omitempty"`
	Rsp						float64	`xml:"RSP,attr,omitempty"`
}

type HotelRoom struct {
	XMLName				xml.Name	`xml:"HotelRoom"`
	Code				string		`xml:"Code,attr"`
	Id					string		`xml:"Id,attr,omitempty"`
	ExtraBed			bool		`xml:"ExtraBed,attr,omitempty"`
	NumberOfCots		string		`xml:"NumberOfCots,attr,omitempty"`
	NumberOfExtraBeds	string		`xml:"NumberOfExtraBeds,attr,omitempty"`
	NumberOfRooms		string		`xml:"NumberOfRooms,attr"`
	SharingBedding		bool		`xml:"SharingBedding,attr,omitempty"`
}

type LocationDetails struct {
	Location	Location	`xml:"Location"`
}

type Location struct {
	Code	string	`xml:"Code,attr"`
}

type Item struct {
	Code	string	`xml:"Code,attr"`
}

type City struct {
	Code	string	`xml:"Code,attr"`
}

// End Response Search Hotel Price
