package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/jcocozza/super-transformer/core"
	formattransformers "github.com/jcocozza/super-transformer/core/formatTransformers"
)


func ParseArgs(args []string) {
	hasHeader := flag.Bool("has-header", false, "[CSV] include if the csv file has a header")
	splitByLine := flag.Bool("split-line", false, "[Plain Text] include if you want to return the plain text line by line")
	flag.Parse()
	if len(args) < 1 {
		fmt.Println("expected file path")
	}
	path := os.Args[1]
	t := core.GetTransformer(path)
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	d, err := t.Parse(data, &formattransformers.FormatOptions{SplitByLine: *splitByLine, HasHeader: *hasHeader})
	if err != nil {
		panic(err)
	}
	b, _ := core.EncodeJson(d)
	fmt.Println(string(b))
}
