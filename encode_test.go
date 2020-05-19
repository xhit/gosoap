package gosoap

import (
	"testing"
)

var (
	mapParamsTests = []struct {
		Params Params
		Err    string
	}{
		{
			Params: Params{"": ""},
			Err:    "error expected: xml: start tag with no name",
		},
	}

	arrayParamsTests = []struct {
		Params ArrayParams
		Err    string
	}{
		{
			Params: ArrayParams{{"", ""}},
			Err:    "error expected: xml: start tag with no name",
		},
	}
)

func TestClient_MarshalXML(t *testing.T) {
	soap, err := SoapClient("http://ec.europa.eu/taxation_customs/vies/checkVatService.wsdl")
	if err != nil {
		t.Errorf("error not expected: %s", err)
	}

	for _, test := range mapParamsTests {
		_, err = soap.Call("checkVat", test.Params)
		if err == nil {
			t.Errorf(test.Err)
		}
	}
}

func TestClient_MarshalXML2(t *testing.T) {
	soap, err := SoapClient("http://ec.europa.eu/taxation_customs/vies/checkVatService.wsdl")
	if err != nil {
		t.Errorf("error not expected: %s", err)
	}

	for _, test := range arrayParamsTests {
		_, err = soap.Call("checkVat", test.Params)
		if err == nil {
			t.Errorf(test.Err)
		}
	}
}
