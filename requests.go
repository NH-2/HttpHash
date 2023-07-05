package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func hashResponse(rawUrl string) (string, error) {
	schemeUrl := addScheme(rawUrl)

	err := validateURL(schemeUrl)
	if err != nil {
		return "", fmt.Errorf("failed to parse url: %v", err)
	}

	responseBody, err := getResponseBody(schemeUrl)
	if err != nil {
		return "", fmt.Errorf("error: %v", err)
	}
	binaryHash := md5.Sum(responseBody)
	hexHash := hex.EncodeToString(binaryHash[:])
	return hexHash, nil
}

// adds scheme (`http://`) to the raw URL if missing
func addScheme(rawUrl string) string {
	if rawUrl[0:len("http://")] != "http://" {
		// add scheme to url
		rawUrl = fmt.Sprintf("http://%s", rawUrl)
	}
	return rawUrl
}

// checks that the URL has a valid format.
func validateURL(rawUrl string) error {
	_, err := url.ParseRequestURI(rawUrl)
	if err != nil {
		// if URL is invalid, it would not be parsed
		return err
	}
	return nil
}

func getResponseBody(schemeUrl string) ([]byte, error) {
	resp, err := http.Get(schemeUrl)
	if err != nil {
		return nil, fmt.Errorf("perform GET request to the url %v: %v", schemeUrl, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body %v", err)
	}
	return body, nil
}
