package gta_sdk

import (
	"encoding/xml"
	"time"
)

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
	XMLName					xml.Name				`xml:"RoomCategory"`
	Id						string					`xml:"Id,attr"`
	Description				string					`xml:"Description,cdata"`
	ItemPrice				ItemPrice				`xml:"ItemPrice"`
	Confirmation			Confirmation			`xml:"Confirmation,cdata"`
	SharingBedding			bool					`xml:"SharingBedding"`
	Meals					Meals					`xml:"Meals"`
	HotelRoomPrices			[]HotelRoomPrice		`xml:"HotelRoomPrices"`
	Offer					Offer					`xml:"Offer,omitempty"`
	EssentialInformation	EssentialInformation	`xml:"EssentialInformation,omitempty"`
	ChargeConditions		ChargeConditions		`xml:"ChargeConditions,omitempty"`
}

type Offer struct {
	Code	string	`xml:"Code,attr"`
}

type ChargeConditions struct {
	ChargeCondition	[]ChargeCondition	`xml:"ChargeCondition"`
}

type ChargeCondition struct {
	Type		string		`xml:"Type,attr"`
	Condition	[]Condition	`xml:"Condition"`
}

type Condition struct {
	Allowable		bool		`xml:"Allowable,attr,omitempty"`
	Charge			bool		`xml:"Charge,attr,omitempty"`
	ChargeAmount	string		`xml:"ChargeAmount,attr,omitempty"`
	Currency		string		`xml:"Currency,attr,omitempty"`
	FromDate		time.Time	`xml:"FromDate,attr,omitempty"`
	ToDate			time.Time	`xml:"ToDate,attr,omitempty"`
}

type EssentialInformation struct {
	Information	string	`xml:"Information"`
}

type Information struct {
	Text		string		`xml:"Text,cdata"`
	DateRange	DataRange	`xml:"DateRange"`
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
	Price		Price		`xml:"Price"`
}

type Price struct {
	Gross	string	`xml:"Gross,attr"`
	Night	string	`xml:"Night,attr"`
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
