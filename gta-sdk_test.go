package gta_sdk

import (
	"testing"
	"encoding/xml"
	"fmt"
)

func TestSourceStruct(t *testing.T) {
	src := &Source{}
	src.RequestorID = RequestorID{Client:"Test client", EMailAddress:"test@email.com", Password:"test pass"}
	src.RequestorPreferences = RequestorPreferences{Country:"Russia", Currency:"RUB", Language:"RU-ru", RequestMode:SYNCHRONOUS}

	b, _ := xml.MarshalIndent(src,"","   ")
	// Print XML and check by eyes
	fmt.Println(string(b))
}
