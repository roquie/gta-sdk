package gta_sdk

import "time"

type RequestMode string

const (
	SYNCHRONOUS = "SYNCHRONOUS"
	ASYNCHRONOUS = "ASYNCHRONOUS"
)

// Auth and request preferences
type Source struct {
	RequestorID				RequestorID				`xml:"RequestorID"`
	RequestorPreferences	RequestorPreferences	`xml:"RequestorPreferences"`
}

type RequestorPreferences struct {
	Language	string		`xml:"Language,attr"`
	Currency	string		`xml:"Currency,attr"`
	Country		string		`xml:"Country,attr"`
	RequestMode	RequestMode	`xml:"RequestMode"`
	ResponseURL	string		`xml:"ResponseUrl,omitempty"`
}

type RequestorID struct {
	Client			string	`xml:"Client,attr"`
	EMailAddress	string	`xml:"EMailAddress,attr"`
	Password		string	`xml:"Password,attr"`
}
// End Auth and request preferences

// Response
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
	SearchCountryResponse	SearchCountryResponse	`xml:"SearchCountryResponse"`
}

type SearchCountryResponse struct {
	Iso	string	`xml:"ISO,attr"`
}
// End Response

// Requests
type Request struct {
	Source			Source			`xml:"Source"`
	RequestDetails	RequestDetails	`xml:"RequestDetails"`
}

type RequestDetails struct {
	RequestItems interface{}
}

// Search Hotel Price
type SearchHotelPriceRequest struct {
	ItemDestination				ItemDestination			`xml:"ItemDestination"`
	ItemCodes					ItemCodes				`xml:"ItemCodes,omitempty"`
	ImmediateConfirmationOnly	bool					`xml:"ImmediateConfirmationOnly"`
	PeriodOfStay				PeriodOfStay			`xml:"PeriodOfStay"`
	IncludeRecommended			bool					`xml:"IncludeRecommended"`
	IncludePriceBreakdown		bool					`xml:"IncludePriceBreakdown"`
	IncludeChargeableItems		bool					`xml:"IncludeChargeableItems"`
	ExcludeChargeableItems		ExcludeChargeableItems	`xml:"ExcludeChargeableItems"`
	StarRating					StarRating				`xml:"MinimumRating"`
	LocationCode				string					`xml:"LocationCode"`
	FacilityCodes				[]FacilityCodes			`xml:"FacilityCodes"`
	OrderBy						string					`xml:"OrderBy"`
	NumberOfReturnedItems		int						`xml:"NumberOfReturnedItems"`
}

type FacilityCodes struct {
	FacilityCode	string	`xml:"FacilityCode"`
}

type StarRating struct {
	MinimumRating	bool	`xml:"MinimumRating,attr"`
}

type Rooms struct {
	Rooms	[]Room	`xml:"Rooms"`
}

type Room struct {
	Code			string		`xml:"Code,attr"`
	NumberOfRooms	int			`xml:"NumberOfRooms,attr,omitempty"`
	ExtraBeds		ExtraBeds	`xml:"ExtraBeds"`
}

type ExtraBeds struct {
	Age	int	`xml:"Age"`
}

type ExcludeChargeableItems struct {
	CancellationDeadlineHours	int	`xml:"CancellationDeadlineHours"`
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
	DestinationType	string	`xml:"DestinationType,attr"`
	DestinationCode	string	`xml:"DestinationCode,attr,omitempty"`
	Latitude		string	`xml:"Latitude,omitempty"`
	Longitude		string	`xml:"Longitude,omitempty"`
	RadiusKm		string	`xml:"RadiusKm,omitempty"`
	WestLongitude	string	`xml:"WestLongitude,omitempty"`
	SouthLatitude	string	`xml:"SouthLatitude,omitempty"`
	EastLongitude	string	`xml:"EastLongitude,omitempty"`
	NorthLatitude	string	`xml:"NorthLatitude,omitempty"`
}

// Requests
