package formattransformers

import (
	"bytes"
	"fmt"

	"github.com/xuri/excelize/v2"
)

type ExcelTransformer struct{}

func (dx *ExcelTransformer) Parse(data []byte, opts *FormatOptions) (map[string]any, error) {
	f, err := excelize.OpenReader(bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("failed to parse excel file: %w", err)
	}
	sheetNames := f.GetSheetList()
	if len(sheetNames) == 0 {
		return nil, fmt.Errorf("no sheets found")
	}
	allData := make(map[string][]map[string]any)
	for _, sheet := range sheetNames {
		rows, err := f.GetRows(sheet)
		if err != nil {
			return nil, err
		}

		var data []map[string]any
		if len(rows) > 0 {
			headers := rows[0]
			for _, row := range rows[1:] {
				entry := make(map[string]interface{})
				for i, cell := range row {
					if i < len(headers) {
						entry[headers[i]] = cell
					}
				}
				data = append(data, entry)
			}
		}
		allData[sheet] = data
	}
	var jsonData = make(map[string]any)
	jsonData["sheets"] = allData
	return jsonData, nil
}
