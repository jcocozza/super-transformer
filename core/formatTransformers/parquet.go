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

	/* new parquet reader */
	pr, err := reader.NewParquetReader(byteReader, new(interface{}), 4)

	if err != nil {
		return nil, fmt.Errorf("failed to create parquet reader: %w", err)
	}
	defer pr.ReadStop()

	/* read data from file */
	numRows := int(pr.GetNumRows())

	fmt.Println("num of rows is", numRows)

	return nil, nil
}
