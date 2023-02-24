package main

import (
	"io"
	"net/http"
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
	resp := expect(http.Get(url))
	body := expect(io.ReadAll(resp.Body))
	must(resp.Body.Close())

	return body
}
