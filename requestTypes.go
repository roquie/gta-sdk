package gta_sdk


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
// Requests
