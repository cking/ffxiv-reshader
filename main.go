package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"os"
	"path"
)

func getZipOffset(buffer []byte) int {
	magicBytes := []byte{0x50, 0x4b, 0x03, 0x04}
	return bytes.Index(buffer, magicBytes)
}

func main() {
	gamePath := getGamePath(arg(os.Args, 1))
	if _, err := os.Stat(path.Join(gamePath, "ffxiv_dx11.exe")); os.IsNotExist(err) {
		gamePath = path.Join(gamePath, "game")
	}

	fmt.Printf("Installing ReShade into: [%v]\n", gamePath)

	reshadeVersion := latestReshade()
	reshade := download("https://reshade.me/downloads/ReShade_Setup_" + reshadeVersion + "_Addon.exe")
	offset := getZipOffset(reshade)
	reader := expect(zip.NewReader(bytes.NewReader(reshade[offset:]), int64(len(reshade[offset:]))))

	for _, file := range reader.File {
		if file.FileInfo().Name() == "ReShade64.dll" {
			f := expect(file.Open())
			out := expect(os.Create(path.Join(gamePath, "dxgi.dll")))
			_ = expect(io.Copy(out, f))
			must(out.Close())
			must(f.Close())
		}
	}

	fmt.Printf("ReShade version [%v] installed\n", reshadeVersion)

	println("Downloading ReShade presets and shaders...")
	downloadShadersAndPresets(path.Join(gamePath, "reshade-shaders"), path.Join(gamePath, "reshade-presets"))
	println("Done, enjoy!")
}
