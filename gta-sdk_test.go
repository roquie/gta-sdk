package gta_sdk

import (
	"bufio"
	"encoding/xml"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
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
	fullFileName := filepath.Join(cwd, "test_data", fName)
	data, err := ioutil.ReadFile(fullFileName)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func getRequestHeaders() *Request {
	src := &Request{}
	src.Source.RequestorID = RequestorID{Client:"Test client", EMailAddress:"test@email.com", Password:"test pass"}
	src.Source.RequestorPreferences = RequestorPreferences{Country:"Russia", Currency:"RUB", Language:"RU-ru", RequestMode:SYNCHRONOUS, ResponseURL:""}
	return src
}

func TestResponseStruct(t *testing.T) {
	resp := &Response{}
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
	buf := serializeRequest(&src.Source, &src.RequestDetails)

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
	src := &RequestDetails{}
	rh := getRequestHeaders()
	rh.RequestDetails = *src
	buf := serializeRequest(&rh.Source, &rh.RequestDetails)

	t.Log(string(buf.Bytes()))
}
