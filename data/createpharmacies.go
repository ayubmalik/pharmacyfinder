package data

import (
	"encoding/csv"
	"fmt"
	"os"
)

// PharmacySummaries takes NHS ODS pharmacy CSV files and creates CSV of active phamacies.
// See https://digital.nhs.uk/services/organisation-data-service/data-downloads/gp-and-gp-practice-related-data
// for source data files.
func PharmacySummaries(dispensaryCsv string, outputCsv string) error {
	d, err := os.Open(dispensaryCsv)
	if err != nil {
		return err
	}
	defer d.Close()

	rows, _ := csv.NewReader(d).ReadAll()

	var summaries []string
	for _, row := range rows {
		// A for active and 1 for type Pharmacy
		if row[12] == "A" && row[13] == "1" {
			s := fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%s,%s,%s", row[0], row[1], row[4], row[5], row[6], row[7], row[8], row[9], row[17], row[23])
			summaries = append(summaries, s)
		}
	}

	return write(outputCsv, summaries)
}

func write(file string, values []string) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()
	last := len(values) - 1
	for i, value := range values {
		nl := "\n"
		if i == last {
			nl = ""
		}
		fmt.Fprintf(f, "%s%s", value, nl)
	}
	return nil
}