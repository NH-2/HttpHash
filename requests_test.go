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
