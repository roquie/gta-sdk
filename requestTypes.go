package gta_sdk

import (
	"encoding/xml"
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
	ImmediateConfirmationOnly	XmlBoolTag				`xml:"ImmediateConfirmationOnly,omitempty"`
	ItemName					string					`xml:"ItemName,omitempty"`
	ItemCodes					ItemCodes				`xml:"ItemCodes,omitempty"`
	PeriodOfStay				struct{
		CheckInDate	string		`xml:"CheckInDate"`
		Duration	int			`xml:"Duration"`
	}	`xml:"PeriodOfStay,omitempty"`
	IncludeRecommended			XmlBoolTag				`xml:"IncludeRecommended,omitempty"`
	RecommendedOnly				XmlBoolTag				`xml:"RecommendedOnly,omitempty"`
	IncludePriceBreakdown		XmlBoolTag				`xml:"IncludePriceBreakdown,omitempty"`
	IncludeChargeConditions		XmlBoolTag				`xml:"IncludeChargeConditions,omitempty"`
	IncludeChargeableItems		IncludeChargeableItems	`xml:"IncludeChargeableItems,omitempty"`
	ExcludeChargeableItems		struct{
		CancellationDeadlineHours	int	`xml:"CancellationDeadlineHours,omitempty"`
		CancellationDeadlineDays	int	`xml:"CancellationDeadlineDays,omitempty"`
	}	`xml:"ExcludeChargeableItems,omitempty"`
	Rooms						struct{
		Rooms []Room	`xml:"Room"`
	}	`xml:"Rooms,omitempty"`
	StarRating					struct{
		MinimumRating	bool	`xml:"MinimumRating,attr"`
		Value			int		`xml:",chardata"`
	}	`xml:"StarRating,omitempty"`
	StarRatingRange				StarRatingRange			`xml:"StarRatingRange,omitempty"`
	LocationCode				string					`xml:"LocationCode,omitempty"`
	FacilityCodes				struct{
		FacilityCodes	[]string	`xml:"FacilityCode"`
	}	`xml:"FacilityCodes,omitempty"`
	OrderBy						OrderByPriceEnum		`xml:"OrderBy,omitempty"`
	NumberOfReturnedItems		int						`xml:"NumberOfReturnedItems,omitempty"`
}

type IncludeChargeableItems struct {
	DateFormatResponse	bool	`xml:"DateFormatResponse,omitempty"`
}

func (ici IncludeChargeableItems) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if !ici.DateFormatResponse {
		e.EncodeElement("", start)
	}
	return nil
}

func (ici IncludeChargeableItems) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	d.Skip()
	ici.DateFormatResponse = true
	return nil
}

type XmlBoolTag struct {
	Value	bool	`xml:"-"`
}

func (x XmlBoolTag) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if x.Value {
		e.EncodeElement("", start)
	}
	return nil
}

func (x *XmlBoolTag) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	d.Skip()
	x.Value = true
	return nil
}

type ItemCodes struct {
	ItemCodes	[]string	`xml:"ItemCode"`
}

func (ic ItemCodes) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if ic.ItemCodes == nil || len(ic.ItemCodes) == 0 {
		e.EncodeElement(nil, start)
	} else {
		e.EncodeToken(start)
		for _, item := range ic.ItemCodes {
			e.EncodeElement(item, xml.StartElement{Name: xml.Name{Local: "ItemCode"}})
		}
		e.EncodeToken(xml.EndElement{Name: start.Name})
	}
	return nil
}

type StarRatingRange struct {
	Min	int	`xml:"Min,omitempty"`
	Max	int	`xml:"Max,omitempty"`
}

func (sr StarRatingRange) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if sr.Min == 0 && sr.Max == 0 {
		e.EncodeElement(nil, start)
	} else {
		e.EncodeToken(start)
		e.EncodeElement(sr.Min, xml.StartElement{Name: xml.Name{Local: "Min"}})
		e.EncodeElement(sr.Max, xml.StartElement{Name: xml.Name{Local: "Max"}})
		e.EncodeToken(xml.EndElement{Name: start.Name})
	}
	return nil
}

type Room struct {
	Code			string		`xml:"Code,attr"`
	Id				string		`xml:"Id,attr,omitempty"`
	NumberOfRooms	string		`xml:"NumberOfRooms,attr,omitempty"`
	NumberOfCots	string		`xml:"NumberOfCots,attr,omitempty"`
	ExtraBeds		ExtraBeds	`xml:"ExtraBeds,omitempty"`
}

type ExtraBeds struct {
	Age	int	`xml:"Age,omitempty"`
}

func (eb ExtraBeds) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if eb.Age == 0 {
		e.EncodeElement(nil, start)
	} else {
		e.EncodeToken(start)
		e.EncodeElement(eb.Age, xml.StartElement{Name: xml.Name{Local: "Age"}})
		e.EncodeToken(xml.EndElement{Name: start.Name})
	}
	return nil
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
