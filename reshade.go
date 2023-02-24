package main

import (
	"io"
	"net/http"
	"regexp"
)

func latestReshade() string {
	resp := expect(http.Get("https://reshade.me"))
	body := expect(io.ReadAll(resp.Body))
	must(resp.Body.Close())
	re := regexp.MustCompile(`ReShade_Setup_([\d.]+)\.exe`)
	match := re.FindAllStringSubmatch(string(body), -1)
	if len(match) == 0 || match[0][1] == "" {
		panic("Could not extract version info")
	}

	return match[0][1]
}
