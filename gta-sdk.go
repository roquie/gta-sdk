package gta_sdk

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

const gtaApiUrl = ""

func serializeRequest(reqAuth *Source, reqDetails *RequestDetails) *bytes.Buffer {
	res := 	&Request{Source:*reqAuth, RequestDetails:*reqDetails}
	bItem, err := xml.Marshal(res)
	FatalError(err)
	data := xml.Header + string(bItem)
	return bytes.NewBuffer([]byte(data))
}

func deSerializeResponse(data []byte) *Response {
	res := &Response{}
	err := xml.Unmarshal(data, res)
	FatalError(err)
	return res
}

func gtaRequestInternal(reqAuth *Source, reqDetails *RequestDetails) (*Response, *[]GtaResponseError) {
	data := serializeRequest(reqAuth, reqDetails)

	req, err := http.NewRequest("POST", gtaApiUrl, data)
	FatalError(err)
	req.Header.Add("Content-Type", "application/xml")

	client := &http.Client{}
	resp, err := client.Do(req)
	FatalError(err)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	respDetails := deSerializeResponse(body)
	if respDetails.ResponseDetails.Errors != nil {
		return nil, ErrorResponse(respDetails.ResponseDetails.Errors)
	}
	return respDetails, nil
}
