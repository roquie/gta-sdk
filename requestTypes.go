package gta_sdk

type RequestMode string

const (
	SYNCHRONOUS = "SYNCHRONOUS"
	ASYNCHRONOUS = "ASYNCHRONOUS"
)

type Source struct {
	RequestorID				RequestorID				`xml:"RequestorID"`
	RequestorPreferences	RequestorPreferences	`xml:"RequestorPreferences"`
}

type RequestorPreferences struct {
	Language	string		`xml:"Language,attr"`
	Currency	string		`xml:"Currency,attr"`
	Country		string		`xml:"Country,attr"`
	RequestMode	RequestMode	`xml:"RequestMode"`
}

type RequestorID struct {
	Client			string	`xml:"Client,attr"`
	EMailAddress	string	`xml:"EMailAddress,attr"`
	Password		string	`xml:"Password,attr"`
}
