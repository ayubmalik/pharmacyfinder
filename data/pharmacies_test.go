package data

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/ayubmalik/nhsfinder"
)

func TestCreatePharmacies(t *testing.T) {

	t.Run("Only include active pharmacies", func(t *testing.T) {

		goldenFile := "testdata/pharmacies.golden.csv"
		inputFile := "testdata/sample-edispensary.csv"
		outputFile := "/tmp/pharmacies.csv"
		latLngs := map[string]nhsfinder.LatLng{
			"SK7 5LD":  nhsfinder.LatLng{Lat: 51.000000, Lng: 1.000000},
			"LS2 8PJ":  nhsfinder.LatLng{Lat: 52.000000, Lng: 2.000000},
			"TN27 9AA": nhsfinder.LatLng{Lat: 53.000000, Lng: 3.000000}}

		if err := CreatePharmacies(inputFile, latLngs, outputFile); err != nil {
			t.Fatalf("%v", err)
		}

		expected, _ := ioutil.ReadFile(goldenFile)
		actual, _ := ioutil.ReadFile(outputFile)

		if !bytes.Equal(expected, actual) {
			t.Fatalf("actual/expected contents not equal:\n%s\n\n%s", actual, expected)
		}
	})
}
