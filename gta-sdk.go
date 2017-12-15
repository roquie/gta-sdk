package gta_sdk

import (
	"net/http"
	"bytes"
	"io/ioutil"
	"encoding/xml"
)

const gtaAPIUrl = ""

func gtaRequestInternal(reqAuth *Source) ([]byte, error) {
	data, err := xml.Marshal(reqAuth)
	FatalError(err)
	req, err := http.NewRequest("POST", gtaAPIUrl, bytes.NewBuffer(data))
	req.Header.Add("Content-Type", "application/xml")

	client := &http.Client{}
	resp, err := client.Do(req)
	FatalError(err)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return body, nil
}