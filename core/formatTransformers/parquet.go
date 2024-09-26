package formattransformers

import (
	"bytes"
	"fmt"

	"github.com/xitongsys/parquet-go/reader"
	"github.com/xitongsys/parquet-go/source"
)

type ParquetTransformer struct{}

type BytesReader struct {
	*bytes.Reader
}

// no op because we don't have to create
func (r *BytesReader) Create(string) (source.ParquetFile, error) { return nil, nil }

// return a new byte reader to read files
func (r *BytesReader) Open(_ string) (source.ParquetFile, error) {
	return &BytesReader{r.Reader}, nil
}

// seek to position
func (r *BytesReader) Seek(offset int64, whence int) (int64, error) {
	return r.Reader.Seek(offset, whence)
}

// length of slices
func (r *BytesReader) Size() (int64, error) {
	return int64(r.Len()), nil
}

// no op since nothing to free?
func (r *BytesReader) Close() error { return nil }

// no op because we never write parquet
func (r *BytesReader) Write(p []byte) (n int, err error) {
	return 0, fmt.Errorf("write op not supported")
}

func (pt *ParquetTransformer) Parse(data []byte, opts *FormatOptions) (map[string]any, error) {
	/* memory reader for raw bytes */
	// pr := source.NewBytesReader(bytes.NewReader(data))

	if data == nil || len(data) == 0 {
		return nil, fmt.Errorf("no data provided")
	}

	fmt.Println("hi")

	byteReader := &BytesReader{bytes.NewReader(data)}

	fmt.Println("bye")

	/* new parquet reader */
	pr, err := reader.NewParquetReader(byteReader, nil, 4)

	if err != nil {
		return nil, fmt.Errorf("failed to create parquet reader: %w", err)
	}
	defer pr.ReadStop()

	/* read data from file */
	numRows := int(pr.GetNumRows())
	if opts != nil && opts.RowLimit > 0 && opts.RowLimit < numRows {
		numRows = opts.RowLimit
	}

	rows := make([]interface{}, numRows)
	if err := pr.Read(&rows); err != nil {
		return nil, fmt.Errorf("failed to read rows: %w", err)
	}

  fmt.Println("num rows read is:", len(rows))

	result := make(map[string]any)
	for i, row := range rows {
    fmt.Println(row)
		result[fmt.Sprintf("row_%d", i)] = row
	}

	fmt.Println("num of rows is", numRows)

  fmt.Println(result)

	return result, nil
}
