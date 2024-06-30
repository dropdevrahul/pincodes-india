package pincodes_test

import (
	"testing"

	"github.com/dropdevrahul/pincodes-india/pincodes-in-go"
	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	result, err := pincodes.Load()

	assert.NoError(t, err)

	assert.Equal(t, len(result), 19300)

	assert.Equal(t, result[110001], pincodes.Pincode{
		CircleName: "Delhi Circle", RegionName: "DivReportingCircle", DivisionName: "New Delhi Central Division",
		OfficeName: "Parliament House SO", Pincode: 110001, OfficeType: "PO", Delivery: "Non Delivery",
		District: "NEW DELHI", StateName: "DELHI", Latitude: "28.6165000", Longitude: "77.2009722",
	})
}
