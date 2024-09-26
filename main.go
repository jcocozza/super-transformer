package main

import (
	"fmt"
	"os"

	"github.com/jcocozza/super-transformer/core"
	formattransformers "github.com/jcocozza/super-transformer/core/formatTransformers"
)

func main() {
	/*
		path := "/Users/josephcocozza/Downloads/airtravel.csv"
		t := core.GetTransformer(path)
		data, err := os.ReadFile(path)
		if err != nil {
			panic(err)
		}
		d, err := t.Parse(data, &formattransformers.FormatOptions{HasHeader: true})
		if err != nil {
			panic(err)
		}
		b, _ := core.EncodeJson(d)
		fmt.Println(string(b))
	*/

	WD, _ := os.Getwd()
	FILE := "mtcars.parquet"

	path := fmt.Sprintf("%s/data/%s", WD, FILE)

	fmt.Println(FILE)

	t := core.GetTransformer(path)
	data, err := os.ReadFile(path)
	fmt.Println(data)
	if err != nil {
		panic(err)
	}
	d, err := t.Parse(data, &formattransformers.FormatOptions{})
	if err != nil {
		panic(err)
	}
	b, _ := core.EncodeJson(d)
	fmt.Println(string(b))
}
