package fsvenues

import (
	"os"
	"testing"
)

func Test_GetVenue(t *testing.T) {
	fs := NewFSVenuesClient(os.Getenv("FOURSQUARE_ID"), os.Getenv("FOURSQUARE_SECRET"))
	params := make(map[string]string)
	params["ll"] = "40.7,-74"
	if v, e := fs.GetVenue(params); e == nil {
		t.Log("Passed", v)
	} else {
		t.Error("Failed", e)
	}
}

func Test_ExploreVenues(t *testing.T) {
	fs := NewFSVenuesClient(os.Getenv("FOURSQUARE_ID"), os.Getenv("FOURSQUARE_SECRET"))
	params := make(map[string]string)
	params["ll"] = "40.7,-74"
	if v, e := fs.ExploreVenues(params); e == nil {
		t.Log("Passed", v)
	} else {
		t.Error("Failed", e)
	}
}

func Test_GetVenues(t *testing.T) {
	fs := NewFSVenuesClient(os.Getenv("FOURSQUARE_ID"), os.Getenv("FOURSQUARE_SECRET"))
	params := make(map[string]string)
	params["ll"] = "40.7,-74"
	if v, e := fs.GetVenues(params); e == nil {
		t.Log("Passed", v)
	} else {
		t.Error("Failed", e)
	}
}
