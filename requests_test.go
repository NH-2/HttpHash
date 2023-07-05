package main

import (
	"testing"
)

func Test_addScheme(t *testing.T) {

	// test when scheme is missing
	rawUrl := "example.com"
	schemeUrl := addScheme(rawUrl)

	if schemeUrl != "http://example.com" {
		t.Errorf("missing scheme was not added correctly")
	}

	// test when scheme is present
	rawUrl = "http://example.com"
	schemeUrl = addScheme(rawUrl)

	if schemeUrl != "http://example.com" {
		t.Errorf("URL must not be modified if scheme is present")
	}
}

func Test_validateUrl(t *testing.T) {
	validUrl := "http://example.com"
	invalidUrlHost := "http://example com/hello"

	err := validateURL(validUrl)
	if err != nil {
		t.Errorf("valid URL should not raise an error")
	}
	err = validateURL(invalidUrlHost)
	if err == nil {
		t.Errorf("invalid URL should raise an error")
	}
}
