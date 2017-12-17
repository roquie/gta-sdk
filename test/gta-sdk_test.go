package test

import (
	"bufio"
	"encoding/xml"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/tmconsulting/gta-golang-sdk"
)

func crlfBuf(src []byte) string {
	s := bufio.NewScanner(strings.NewReader(string(src)))
	res := ""
	for s.Scan() {
		res += strings.TrimRight(strings.TrimRight(s.Text(), "\n"), "\r")
	}
	return res
}

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
	src.Source.RequestorID = gta_sdk.RequestorID{Client:"Test client", EMailAddress:"test@email.com", Password:"test pass"}
	src.Source.RequestorPreferences = gta_sdk.RequestorPreferences{Country:"Russia", Currency:"RUB", Language:"RU-ru", RequestMode:gta_sdk.SYNCHRONOUS, ResponseURL:""}
	return src
}

func TestResponseStruct(t *testing.T) {
	resp := &gta_sdk.Response{}
	resp.ResponseSequence = "1"
	resp.ResponseReference = "123456"
	resp.ResponseDetails.Language = "RU-ru"
	d, _ := xml.Marshal(resp)
	eData, err := getTestData("response.xml")
	if err != nil {
		t.Error(err)
	}
	sc := crlfBuf(d)
	fd := crlfBuf(eData)
	if sc != fd {
		t.Fatal("Error deserialize response")
	}
	t.Log("TestResponseStruct: Ok")
}

func TestRequestStruct(t *testing.T) {
	src := getRequestHeaders()
	buf := gta_sdk.SerializeRequest(&src.Source, &src.RequestDetails)

	eData, err := getTestData("request.xml")
	if err != nil {
		t.Error(err)
	}
	sc := crlfBuf(buf.Bytes())
	fd := crlfBuf(eData)
	if sc != fd {
		t.Fatal("Error serialize request")
	}
	t.Log("TestRequestStruct: Ok")
}

func TestSearchHotelPriceRequest(t *testing.T) {
	src := &gta_sdk.RequestDetails{}
	rh := getRequestHeaders()
	rh.RequestDetails = *src
	buf := gta_sdk.SerializeRequest(&rh.Source, &rh.RequestDetails)

	t.Log(string(buf.Bytes()))
}
