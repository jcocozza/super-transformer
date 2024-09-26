package formattransformers

import (
	"bytes"
	"encoding/csv"
	"fmt"
)

type CSVTransformer struct{}

func (cs *CSVTransformer) Parse(data []byte, opts *FormatOptions) (map[string]any, error) {
	reader := csv.NewReader(bytes.NewReader(data))
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed reading csv: %w", err)
	}
	if len(records) == 0 {
		return nil, nil // return nil if no records
	}

	var headers []string = make([]string, len(records[0]))
	startIdx := 0
	numRecords := len(records)
	if opts.HasHeader {
		headers = records[0]
		startIdx = 1
		numRecords = numRecords - 1
	} else {
		for i := range records[0] {
			headers[i] = fmt.Sprintf("row_%d", i)
		}
	}
	csvRecords := make([]map[string]string, numRecords)
	for i, record := range records[startIdx:] {
		csvRecords[i] = make(map[string]string)
		for j, value := range record {
			csvRecords[i][headers[j]] = value
		}
	}
	var jsonData = make(map[string]any)
	jsonData["records"] = csvRecords
	return jsonData, nil
}
