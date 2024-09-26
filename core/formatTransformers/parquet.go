package formattransformers

import (
	"bytes"
	"fmt"

	"github.com/xitongsys/parquet-go/reader"
	"github.com/xitongsys/parquet-go/source"
)

type ParquetTransformer struct{}

// chatgpt claims a fixed schema can help
/*
type MTCars struct {
    MPG   float64 `parquet:"name=mpg, type=DOUBLE"`
    Cyl   int32   `parquet:"name=cyl, type=INT32"`
    Disp  float64 `parquet:"name=disp, type=DOUBLE"`
    HP    int32   `parquet:"name=hp, type=INT32"`
    Drat  float64 `parquet:"name=drat, type=DOUBLE"`
    WT    float64 `parquet:"name=wt, type=DOUBLE"`
    QSec  float64 `parquet:"name=qsec, type=DOUBLE"`
    VS    int32   `parquet:"name=vs, type=INT32"`
    AM    int32   `parquet:"name=am, type=INT32"`
    Gear  int32   `parquet:"name=gear, type=INT32"`
    Carb  int32   `parquet:"name=carb, type=INT32"`
}
*/

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
	if data == nil || len(data) == 0 {
		return nil, fmt.Errorf("no data provided")
	}

	byteReader := &BytesReader{bytes.NewReader(data)}

	/* new parquet reader */
	pr, err := reader.NewParquetReader(byteReader, nil, 1)
	if err != nil {
		return nil, fmt.Errorf("failed to create parquet reader: %w", err)
	}
	defer pr.ReadStop()

  /* ensure schema is legit */
  schema := pr.SchemaHandler
  fmt.Println("schema is:", schema)

	/* read data from file */
	numRows := int(pr.GetNumRows())
	if opts != nil && opts.RowLimit > 0 && opts.RowLimit < numRows {
		numRows = opts.RowLimit
	}
  fmt.Println("num rows to read is:", numRows)

	rows := make([]interface{}, numRows)
  //err = pr.Read(&rows)
  
  // go 1 at a time
  val, err := pr.ReadByNumber(4)
  fmt.Println(val)

  if err != nil {
		return nil, fmt.Errorf("failed to read rows: %w", err)
	}
  fmt.Println("rows read:", len(rows))
  fmt.Printf("row content: %+v\n", rows)

	result := make(map[string]any)
  for i, row := range rows {
    
    fmt.Println(row)
		result[fmt.Sprintf("row_%d", i)] = row
  }

  fmt.Println(result)

	return result, nil
}
