package main

import (
	"net/http"
	"os"
	"testing"
)

var testApp Config

func TestMain(m *testing.M) {
	os.Exit(m.Run()) // Run tests
}

var jsonToReturn = `
{
	"ts": 1654782060772,
	"tsj": 1654782056216,
	"date": "Jun 9th 2022, 09:40:56 am NY",
	"items": [
	  {
		"curr": "USD",
		"xauPrice": 1849,
		"xagPrice": 21.9115,
		"chgXau": -3.735,
		"chgXag": -0.1425,
		"pcXau": -0.2016,
		"pcXag": -0.6461,
		"xauClose": 1852.735,
		"xagClose": 22.054
	  }
	]
  }
`

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}
