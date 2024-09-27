package core

import (
	"path/filepath"

	formattransformers "github.com/jcocozza/super-transformer/core/formatTransformers"
)

// Transformer is the interface that orchestrates file -> json
type Transformer interface {
	Parse(data []byte, opts *formattransformers.FormatOptions) (map[string]any, error)
}

func getFileType(path string) string {
	return filepath.Ext(path)
}

func GetTransformer(path string) Transformer {
	var transformer Transformer
	switch getFileType(path) {
	case ".csv":
		transformer = &formattransformers.CSVTransformer{}
	case ".xlsx":
		transformer = &formattransformers.ExcelTransformer{}
	case ".parquet":
		transformer = &formattransformers.ParquetTransformer{}
	default:
		transformer = &formattransformers.PlainTextTransformer{}
	}
	return transformer
}
