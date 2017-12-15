package gta_sdk

import (
	"log"
	"encoding/xml"
)

type GtaResponseError interface {
	Error() 	string
	RawError()	string
}

func FatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ErrorResponse(bd []byte) (GtaResponseError, error) {
	var v GtaResponseError
	err := xml.Unmarshal(bd, v)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return v, nil
}
