package test

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/tmconsulting/gta-golang-sdk"
	"time"
)

func getTestData(fName string) ([]byte, error) {
	cwd, _ := os.Getwd()
	fullFileName := filepath.Join(cwd, "data", fName)
	data, err := ioutil.ReadFile(fullFileName)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func getRequestHeaders() *gta_sdk.Request {
	src := &gta_sdk.Request{}
	src.Source.RequestorID = gta_sdk.RequestorID{Client:"1479", EMailAddress:"client@net.com", Password:"xxx"}
	src.Source.RequestorPreferences = gta_sdk.RequestorPreferences{Country:"GB", Currency:"GBP", Language:"en", RequestMode:gta_sdk.SYNCHRONOUS, ResponseURL:""}
	return src
}

func TestResponseStruct(t *testing.T) {
	resp := &gta_sdk.Response{}
	resp.ResponseSequence = "1"
	resp.ResponseReference = "123456"
	resp.ResponseDetails.Language = "RU-ru"
	d, err := xml.Marshal(resp)
	if err != nil {
		t.Error(err)
	}
	eData, err := getTestData("response.xml")
	if err != nil {
		t.Error(err)
	}

	err = xml.Unmarshal(d, resp)
	if err != nil {
		t.Error(err)
	}

	testObject := &gta_sdk.Response{}
	err = xml.Unmarshal(eData, testObject)
	if err != nil {
		t.Error(err)
	}

	if gta_sdk.Response(*resp) != gta_sdk.Response(*testObject) {
		t.Fatal("Error deserialize response")
	}
	t.Log("TestResponseStruct: Ok")
}

func TestRequestStruct(t *testing.T) {
	src := getRequestHeaders()
	buf := gta_sdk.SerializeRequest(&src.Source, &src.RequestDetails)

	mokBytes, err := getTestData("request.xml")
	if err != nil {
		t.Error(err)
	}

	testObject := &gta_sdk.Request{}
	err = xml.Unmarshal(mokBytes, testObject)
	if err != nil {
		t.Error(err)
	}

	err = xml.Unmarshal(buf.Bytes(), src)
	if err != nil {
		t.Error(err)
	}

	if gta_sdk.Request(*src) != gta_sdk.Request(*testObject) {
		t.Fatal("Error serialize request")
	}

	t.Log("TestRequestStruct: Ok")
}

func TestSearchHotelPriceRequest(t *testing.T) {
	src := &gta_sdk.SearchHotelPriceRequest{}
	rh := getRequestHeaders()
	rh.RequestDetails.RequestItems = *src

	src.ItemDestination.DestinationType = gta_sdk.CITY
	src.ItemDestination.DestinationCode = "AMS"
	src.ImmediateConfirmationOnly = true
	date := time.Time{}
	src.PeriodOfStay.CheckInDate = date.AddDate(2017,9, 30)
	src.PeriodOfStay.Duration = 4
	src.IncludeRecommended = true
	src.IncludePriceBreakdown = true
	src.IncludeChargeConditions = true
	src.ExcludeChargeableItems.CancellationDeadlineHours = 72
	src.Rooms = make([]gta_sdk.Room, 3)
	src.Rooms[0].Code = "DB"
	src.Rooms[0].NumberOfRooms = "1"
	src.Rooms[0].ExtraBeds.Age = 5
	src.Rooms[1].Code = "TB"
	src.Rooms[1].NumberOfRooms = "2"
	src.Rooms[1].ExtraBeds.Age = 10
	src.Rooms[2].Code = "SB"
	src.StarRating.MinimumRating = true
	src.StarRating.Value = 3
	src.LocationCode = "G1"
	src.FacilityCodes = make([]gta_sdk.FacilityCodes, 2)
	src.FacilityCodes[0].FacilityCode = "*AC"
	src.FacilityCodes[1].FacilityCode = "*LS"
	src.OrderBy = gta_sdk.PRICELOWTOHIGH
	src.NumberOfReturnedItems = 10

	buf := gta_sdk.SerializeRequest(&rh.Source, &rh.RequestDetails)

	mokBytes, err := getTestData("HotelPriceSearchRequestAtCityLevel.xml")
	if err != nil {
		t.Error(err)
	}

	testObject := &gta_sdk.SearchHotelPriceRequest{}
	err = xml.Unmarshal(mokBytes, testObject)
	if err != nil {
		t.Error(err)
	}

	err = xml.Unmarshal(buf.Bytes(), rh)
	if err != nil {
		t.Error(err)
	}

	/*
	if rh.RequestDetails.RequestItems != testObject {
		t.Fatal("Error serialize request")
	}
	*/

	t.Log("TestRequestStruct: Ok")


//	t.Log(string(buf.Bytes()))
}
