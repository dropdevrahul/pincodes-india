package pincodes

import (
	"embed"
	"encoding/csv"
	"strconv"
	"strings"
)

//go:embed pincode.csv
var f embed.FS

// "CircleName","RegionName","DivisionName","OfficeName","Pincode","OfficeType","Delivery",
// "District","StateName","Latitude","Longitude"
type Pincode struct {
	CircleName   string
	RegionName   string
	DivisionName string
	OfficeName   string
	Pincode      int
	OfficeType   string
	Delivery     string
	District     string
	StateName    string
	Latitude     string
	Longitude    string
}

func Load() (map[int]Pincode, error) {
	data, err := f.Open("pincode.csv")
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(data)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	res := map[int]Pincode{}

	for i, row := range records {
		if i == 0 {
			continue
		}

		record, err := rowToRecord(row)
		if err != nil {
			return nil, err
		}

		res[int(record.Pincode)] = *record
	}

	return res, nil
}

func rowToRecord(row []string) (*Pincode, error) {
	pin, err := strconv.Atoi(row[4])
	if err != nil {
		return nil, err
	}

	result := &Pincode{
		CircleName:   row[0],
		RegionName:   row[1],
		DivisionName: row[2],
		OfficeName:   row[3],
		Pincode:      pin,
		OfficeType:   row[5],
		Delivery:     row[6],
		District:     row[7],
		StateName:    row[8],
		Latitude:     strings.TrimSpace(row[9]),
		Longitude:    strings.TrimSpace(row[10]),
	}

	return result, nil
}
