package main

import (
	"bytes"
	"github.com/schollz/progressbar/v3"
	"io"
	"net/http"
	"os"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func expect[T any](value T, err error) T {
	must(err)
	return value
}

func arg(args []string, i int) string {
	if len(args) > i {
		return args[i]
	}
	return ""
}

func download(url string) []byte {
	resp := fetch(url)
	println(resp.Status)
	bar := progressbar.DefaultBytes(resp.ContentLength)
	buffer := new(bytes.Buffer)
	expect(io.Copy(io.MultiWriter(buffer, bar), resp.Body))
	must(resp.Body.Close())

	return buffer.Bytes()
}

func downloadFile(url, path string) {
	buffer := download(url)
	file := expect(os.Create(path))
	expect(file.Write(buffer))
	must(file.Close())
}

func fetch(url string) *http.Response {
	req := expect(http.NewRequest("GET", url, nil))
	req.Header.Set("User-Agent", "ffxiv-reshader/1")
	return expect(http.DefaultClient.Do(req))
}
