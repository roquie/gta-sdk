package gta_sdk

import (
	"log"
)

type GtaResponseError struct {
	ErrorId	string
	Message	string
}

func (re *GtaResponseError) Error() string {
	return re.Message
}

func FatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ErrorResponse(gtaErrors *Errors) *[]GtaResponseError {
	if gtaErrors == nil || len(gtaErrors.Errors) == 0 {
		res := make([]GtaResponseError, 1)
		res[0].Message = "Undefined error"
		return &res
	}
	res := make([]GtaResponseError, len(gtaErrors.Errors))
	for i, item := range gtaErrors.Errors {
		res[i].Message = item.ErrorText
		res[i].ErrorId = item.ErrorId
	}
	return &res
}
