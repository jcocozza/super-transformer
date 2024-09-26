package formattransformers

import (
	"fmt"
	"strings"
)

type PlainTextTransformer struct{}

func (pt *PlainTextTransformer) Parse(data []byte, opts *FormatOptions) (map[string]any, error) {
	jsonData := make(map[string]any)
	if opts != nil && opts.SplitByLine {
		lines := strings.Split(string(data), "\n")
		for i, ln := range lines {
			jsonData[fmt.Sprintf("line_%d", i)]	= ln
		}
	} else {
		jsonData["text"] = string(data)
	}
	return jsonData, nil
}
