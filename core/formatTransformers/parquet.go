package formattransformers

import (
	"bytes"
	"context"
	"fmt"

	"github.com/apache/arrow/go/v14/arrow/array"
	"github.com/apache/arrow/go/v14/arrow/memory"
	"github.com/apache/arrow/go/v14/parquet/file"
	"github.com/apache/arrow/go/v14/parquet/pqarrow"
)

type ParquetTransformer struct{}

func (pt *ParquetTransformer) Parse(data []byte, opts *FormatOptions) (map[string]any, error) {
	if data == nil || len(data) == 0 {
		return nil, fmt.Errorf("no data provided")
	}

	reader := bytes.NewReader(data)

	fr, err := file.NewParquetReader(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to create parquet reader: %w", err)
	}
	defer fr.Close()

	alloc := memory.NewGoAllocator()

	props := pqarrow.ArrowReadProperties{}

	arrowReader, err := pqarrow.NewFileReader(fr, props, alloc)
	if err != nil {
		return nil, fmt.Errorf("failed to create arrow reader: %w", err)
	}

	ctx := context.Background()

	// read table
	tbl, err := arrowReader.ReadTable(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to read parquet as arrow table: %w")
	}
	defer tbl.Release()

	numRows := tbl.NumRows()
	numCols := tbl.NumCols()

	if opts != nil && opts.RowLimit > 0 && opts.RowLimit < int(numRows) {
		numRows = int64(opts.RowLimit)
	}

	result := make(map[string]any)

	// init slices for each column
	// iterate over the Arrow table and extract values
	for rowIdx := int64(0); rowIdx < numRows; rowIdx++ {
		row := make(map[string]any)
		for colIdx := 0; colIdx < int(numCols); colIdx++ {
			col := tbl.Column(colIdx)

			// accessing the column data as an array
			chunk := col.Data().Chunk(0)

			// switch on the type of the array to extract the correct value
			switch chunk := chunk.(type) {
			case *array.Int64:
				row[col.Name()] = chunk.Value(int(rowIdx))
			case *array.Float64:
				row[col.Name()] = chunk.Value(int(rowIdx))
			case *array.String:
				row[col.Name()] = chunk.Value(int(rowIdx))
			case *array.Boolean:
				row[col.Name()] = chunk.Value(int(rowIdx))
			case *array.Int32:
				row[col.Name()] = chunk.Value(int(rowIdx))
			case *array.Int16:
				row[col.Name()] = chunk.Value(int(rowIdx))
			case *array.Timestamp:
				row[col.Name()] = chunk.Value(int(rowIdx))
			case *array.Date32:
				row[col.Name()] = chunk.Value(int(rowIdx))
			case *array.Date64:
				row[col.Name()] = chunk.Value(int(rowIdx))
			default:
				row[col.Name()] = nil
			}
		}
		result[fmt.Sprintf("row_%d", rowIdx)] = row
	}

	return result, nil

}
