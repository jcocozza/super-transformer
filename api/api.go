package api

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/jcocozza/super-transformer/core"
	formattransformers "github.com/jcocozza/super-transformer/core/formatTransformers"
)

func getURLParams(r *http.Request) *formattransformers.FormatOptions {
	opts := &formattransformers.FormatOptions{}
	hasHeader := r.FormValue("has_header")
	splitByLine := r.FormValue("split_by_line")
	rowLimitStr := r.FormValue("row_limit")
	if hasHeader == "true" {
		opts.HasHeader = true
	} else {
		opts.HasHeader = false
	}
	if splitByLine == "true" {
		opts.SplitByLine = true
	} else {
		opts.SplitByLine = false
	}
	rowLimit, err := strconv.Atoi(rowLimitStr)
	if err != nil {
		opts.RowLimit = -1
	} else {
		opts.RowLimit = rowLimit
	}

	return opts
}


func uploadFile(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)
	file, handler, err := r.FormFile("update_file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t := core.GetTransformer(handler.Filename)
	opts := getURLParams(r)
	d, err := t.Parse(fileBytes, opts)
	json, err := core.EncodeJson(d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

const html string = `
<html>
	<body>
		<form enctype="multipart/form-data" action="/upload" method="POST">
			<input type="file" name="update_file" />
			<br>
			split by line:
			<input type="checkbox" name="split_by_line" value="true" />
			<br>
			has header:
			<input type="checkbox" name="has_header" value="true" />
			<br>
			row limit:
			<input type="number" name="row_limit" increment="1" min="1" />
			<br>
			<input type="submit" value="Upload" />
		</form>
	</body>
</html>`

func getFile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, html)
}

func Api() {
	http.HandleFunc("/", getFile)
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":8086", nil)
}
