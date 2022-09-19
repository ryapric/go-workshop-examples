package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetSalesRecords(t *testing.T) {
	rawData := `
	[
		{
			"date": "2022-01-01",
			"region": "East",
			"customer_id": "1a",
			"revenue": 10000
		},
		{
			"date": "2022-01-02",
			"region": "West",
			"customer_id": "2b",
			"revenue": 25000
		}
]
`

	want := AllSalesRecords{
		SalesRecord{
			Date:       "2022-01-01",
			Region:     "East",
			CustomerID: "1a",
			Revenue:    10_000,
		},
		SalesRecord{
			Date:       "2022-01-02",
			Region:     "West",
			CustomerID: "2b",
			Revenue:    25_000,
		},
	}
	got := getSalesRecordsFromJSONBytes([]byte(rawData))

	if !cmp.Equal(want, got) {
		t.Errorf("\nwant: %v\ngot: %v\n", want, got)
	}
}
