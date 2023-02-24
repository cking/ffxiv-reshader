package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path"
	"regexp"
	"strings"
)

type reshadePackList map[string]func(shaders, presets string)

var reshadeShaders = reshadePackList{
	// WinUaeMaskGlow
	"gshadesucks": func(shaders, presets string) {
		f := expect(os.CreateTemp("", "gshadesucks"))
		must(f.Close())
		downloadFile("https://kagamine.tech/shade/gshade.zip", f.Name())

		extractPack(f.Name(), regexp.MustCompile(`^gshade-(shaders|presets)/`), func(s string) string {
			switch s {
			case "gshade-shaders/":
				return shaders + "/"
			case "gshade-presets/":
				return presets + "/"
			default:
				panic("Unexpected match: " + s)
			}
		})

		// fixing compiling by removing 3 files
		must(os.Remove(path.Join(shaders, "Shaders", "WinUaeMaskGlow.fx")))
		must(os.Remove(path.Join(shaders, "Shaders", "NLM_Sharp.fx")))
		must(os.Remove(path.Join(shaders, "Shaders", "SmartDeNoise.fx")))

		must(os.Remove(f.Name()))
	},

	"crosire/reshade-shaders": func(shaders, presets string) {
		f := expect(os.CreateTemp("", "crosire"))
		must(f.Close())
		downloadFile("https://github.com/crosire/reshade-shaders/archive/refs/heads/slim.zip", f.Name())
		extractPack(f.Name(), regexp.MustCompile(`^[^/]+/(Shaders|Textures)/`), func(s string) string { return path.Join(shaders, shiftPath(s)) + "/" })
		must(os.Remove(f.Name()))
	},

	"SweetFX": func(shaders, presets string) {
		f := expect(os.CreateTemp("", "sweetfx"))
		must(f.Close())
		downloadFile("https://github.com/CeeJayDK/SweetFX/archive/refs/heads/master.zip", f.Name())
		extractPack(f.Name(), regexp.MustCompile(`^[^/]+/(Shaders|Textures)/`), func(s string) string { return path.Join(shaders, shiftPath(s)) + "/" })
		must(os.Remove(f.Name()))
	},

	"qUINT": func(shaders, presets string) {
		f := expect(os.CreateTemp("", "quint"))
		must(f.Close())
		downloadFile("https://github.com/martymcmodding/qUINT/archive/refs/heads/master.zip", f.Name())
		extractPack(f.Name(), regexp.MustCompile(`^[^/]+/(Shaders|Textures)/`), func(s string) string { return path.Join(shaders, shiftPath(s)) + "/" })
		must(os.Remove(f.Name()))
	},

	"OtisFX": func(shaders, presets string) {
		f := expect(os.CreateTemp("", "otisfx"))
		must(f.Close())
		downloadFile("https://github.com/FransBouma/OtisFX/archive/refs/heads/master.zip", f.Name())
		extractPack(f.Name(), regexp.MustCompile(`^[^/]+/(Shaders|Textures)/`), func(s string) string { return path.Join(shaders, shiftPath(s)) + "/" })
		must(os.Remove(f.Name()))
	},

	/*
		Broken with reshade:
			- Flair
			- RadiantGI

		"AstrayFX": func(shaders, presets string) {
			f := expect(os.CreateTemp("", "astrayfx"))
			must(f.Close())
			downloadFile("https://github.com/BlueSkyDefender/AstrayFX/archive/refs/heads/master.zip", f.Name())
			extractPack(f.Name(), regexp.MustCompile(`^[^/]+/(Shaders|Textures)/`), func(s string) string { return path.Join(shaders, shiftPath(s)) + "/" })
			must(os.Remove(f.Name()))
		},
	*/

	"Depth3D": func(shaders, presets string) {
		f := expect(os.CreateTemp("", "depth3d"))
		must(f.Close())
		downloadFile("https://github.com/BlueSkyDefender/Depth3D/archive/refs/heads/master.zip", f.Name())
		extractPack(f.Name(), regexp.MustCompile(`^[^/]+/(Shaders|Textures)/`), func(s string) string { return path.Join(shaders, shiftPath(s)) + "/" })
		must(os.Remove(f.Name()))
	},

	"fubax": func(shaders, presets string) {
		f := expect(os.CreateTemp("", "fubax"))
		must(f.Close())
		downloadFile("https://github.com/Fubaxiusz/fubax-shaders/archive/refs/heads/master.zip", f.Name())
		extractPack(f.Name(), regexp.MustCompile(`^[^/]+/(Shaders|Textures)/`), func(s string) string { return path.Join(shaders, shiftPath(s)) + "/" })
		must(os.Remove(f.Name()))
	},

	"Daodan317081/reshade-shaders": func(shaders, presets string) {
		f := expect(os.CreateTemp("", "daodan"))
		must(f.Close())
		downloadFile("https://github.com/Daodan317081/reshade-shaders/archive/refs/heads/master.zip", f.Name())
		extractPack(f.Name(), regexp.MustCompile(`^[^/]+/(Shaders|Textures)/`), func(s string) string { return path.Join(shaders, shiftPath(s)) + "/" })
		must(os.Remove(f.Name()))
	},

	"brussell1/Shaders": func(shaders, presets string) {
		f := expect(os.CreateTemp("", "brussel"))
		must(f.Close())
		downloadFile("https://github.com/brussell1/Shaders/archive/refs/heads/master.zip", f.Name())
		extractPack(f.Name(), regexp.MustCompile(`^[^/]+/(Shaders|Textures|Other)/`), func(s string) string { return path.Join(shaders, shiftPath(s)) + "/" })
		must(os.Remove(f.Name()))
	},

	/*
		Broken with reshade:
			- AspectRatioSuite
			- GloomAO
			- GrainSpread

		"FXShaders": func(shaders, presets string) {
			f := expect(os.CreateTemp("", "fxshaders"))
			must(f.Close())
			downloadFile("https://github.com/luluco250/FXShaders/archive/refs/heads/master.zip", f.Name())
			extractPack(f.Name(), regexp.MustCompile(`^[^/]+/(Shaders|Textures)/`), func(s string) string { return path.Join(shaders, shiftPath(s)) + "/" })
			must(os.Remove(f.Name()))
		},
	*/

	"prod80": func(shaders, presets string) {
		f := expect(os.CreateTemp("", "prod80"))
		must(f.Close())
		downloadFile("https://github.com/prod80/prod80-ReShade-Repository/archive/refs/heads/master.zip", f.Name())
		extractPack(f.Name(), regexp.MustCompile(`^[^/]+/(Shaders|Textures)/`), func(s string) string { return path.Join(shaders, shiftPath(s)) + "/" })
		must(os.Remove(f.Name()))
	},

	"CorgiFX": func(shaders, presets string) {
		f := expect(os.CreateTemp("", "corgifx"))
		must(f.Close())
		downloadFile("https://github.com/originalnicodr/CorgiFX/archive/refs/heads/master.zip", f.Name())
		extractPack(f.Name(), regexp.MustCompile(`^[^/]+/(Shaders|Textures)/`), func(s string) string { return path.Join(shaders, shiftPath(s)) + "/" })
		must(os.Remove(f.Name()))
	},

	/*
		1.3 GB :(

		"MLUT": func(shaders, presets string) {
			f := expect(os.CreateTemp("", "MLUT"))
			must(f.Close())
			downloadFile("https://github.com/TheGordinho/MLUT/archive/refs/heads/master.zip", f.Name())
			extractPack(f.Name(), regexp.MustCompile(`^[^/]+/(Shaders|Textures)/`), func(s string) string { return path.Join(shaders, shiftPath(s))  + "/" }, shaders, presets)
			must(os.Remove(f.Name()))
		},
	*/

	"InsaneShaders": func(shaders, presets string) {
		f := expect(os.CreateTemp("", "insane"))
		must(f.Close())
		downloadFile("https://github.com/LordOfLunacy/Insane-Shaders/archive/refs/heads/master.zip", f.Name())
		extractPack(f.Name(), regexp.MustCompile(`^[^/]+/(Shaders|Textures)/`), func(s string) string { return path.Join(shaders, shiftPath(s)) + "/" })
		must(os.Remove(f.Name()))
	},

	/*
		Broken with reshade:
			- NTSC_XOT
			- NTSCCustom

		"RSRetroArch": func(shaders, presets string) {
			f := expect(os.CreateTemp("", "retro"))
			must(f.Close())
			downloadFile("https://github.com/Matsilagi/RSRetroArch/archive/refs/heads/main.zip", f.Name())
			extractPack(f.Name(), regexp.MustCompile(`^[^/]+/(Shaders|Textures)/`), func(s string) string { return path.Join(shaders, shiftPath(s)) + "/" })
			must(os.Remove(f.Name()))
		},
	*/

	"KosRud/Shaders": func(shaders, presets string) {
		f := expect(os.CreateTemp("", "kosrud"))
		must(f.Close())
		downloadFile("https://github.com/KosRud/Shaders/archive/refs/heads/master.zip", f.Name())
		extractPack(f.Name(), regexp.MustCompile(`^[^/]+/reshade/(Shaders|Textures)/`), func(s string) string { return path.Join(shaders, shiftPath(shiftPath(s))) + "/" })
		must(os.Remove(f.Name()))
	},

	"CobraFX": func(shaders, presets string) {
		f := expect(os.CreateTemp("", "cobrafx"))
		must(f.Close())
		downloadFile("https://github.com/LordKobra/CobraFX/archive/refs/heads/master.zip", f.Name())
		extractPack(f.Name(), regexp.MustCompile(`^[^/]+/(Shaders|Textures)/`), func(s string) string { return path.Join(shaders, shiftPath(s)) + "/" })
		must(os.Remove(f.Name()))
	},

	"WarpFX": func(shaders, presets string) {
		f := expect(os.CreateTemp("", "warpfx"))
		must(f.Close())
		downloadFile("https://github.com/Radegast-FFXIV/Warp-FX/archive/refs/heads/master.zip", f.Name())
		extractPack(f.Name(), regexp.MustCompile(`^[^/]+/(Shaders|Textures)/`), func(s string) string { return path.Join(shaders, shiftPath(s)) + "/" })
		must(os.Remove(f.Name()))
	},

	/*
		Broken with reshade:
			- VRToolkit

		"VRToolKit": func(shaders, presets string) {
			f := expect(os.CreateTemp("", "vr"))
			must(f.Close())
			downloadFile("https://github.com/retroluxfilm/reshade-vrtoolkit/archive/refs/heads/main.zip", f.Name())
			extractPack(f.Name(), regexp.MustCompile(`^[^/]+/(Shaders|Textures|Presets)/`), func(s string) string {
				s = shiftPath(s)
				if strings.HasPrefix(s, "Presets") {
					return path.Join(presets, s[len("Presets/"):]) + "/"
				}
				return path.Join(shaders, s) + "/"
			})
			must(os.Remove(f.Name()))
		},
	*/

	/*
		Broken with reshade:
			- dh_Lain

		"DH": func(shaders, presets string) {
			f := expect(os.CreateTemp("", "dh"))
			must(f.Close())
			downloadFile("https://github.com/AlucardDH/dh-reshade-shaders/archive/refs/heads/master.zip", f.Name())
			extractPack(f.Name(), regexp.MustCompile(`^[^/]+/(Shaders|Textures)/`), func(s string) string { return path.Join(shaders, shiftPath(s)) + "/" })
			must(os.Remove(f.Name()))
		},
	*/

	"FastEffects": func(shaders, presets string) {
		f := expect(os.CreateTemp("", "fasteffects"))
		must(f.Close())
		downloadFile("https://github.com/rj200/Glamarye_Fast_Effects_for_ReShade/archive/refs/heads/main.zip", f.Name())
		extractPack(f.Name(), regexp.MustCompile(`^[^/]+/(Shaders|Textures)/`), func(s string) string { return path.Join(shaders, shiftPath(s)) + "/" })
		must(os.Remove(f.Name()))
	},

	"Pirate-Shaders": func(shaders, presets string) {
		f := expect(os.CreateTemp("", "pirate-shaders"))
		must(f.Close())
		downloadFile("https://github.com/Heathen/Pirate-Shaders/archive/refs/heads/master.zip", f.Name())
		extractPack(f.Name(), regexp.MustCompile(`^[^/]+/reshade-shaders/`), func(s string) string { return shaders + "/" })
		must(os.Remove(f.Name()))
	},
}

func log(format string, args ...interface{}) {
	on := os.Getenv("VERBOSE") != ""
	if !on {
		return
	}
	fmt.Printf(format, args...)
}

func extractPack(fileName string, re *regexp.Regexp, replace func(s string) string) {
	reader := expect(zip.OpenReader(fileName))
	for _, file := range reader.File {
		dest := file.Name
		if !re.MatchString(dest) {
			log(" ! Skipping %v \n", dest)
			continue
		}

		dest = re.ReplaceAllStringFunc(dest, replace)

		if file.FileInfo().IsDir() {
			must(os.MkdirAll(dest, 0755))
			continue
		}

		if _, err := os.Stat(dest); os.IsExist(err) {
			panic("File already exists: " + dest)
		}

		f := expect(file.Open())
		out := expect(os.Create(dest))
		expect(io.Copy(out, f))
	}
}

func shiftPath(p string) string {
	if !strings.Contains(p, "/") {
		return ""
	}

	return p[strings.Index(p, "/")+1:]
}

func latestReshade() string {
	resp := fetch("https://reshade.me")
	body := expect(io.ReadAll(resp.Body))
	must(resp.Body.Close())
	re := regexp.MustCompile(`ReShade_Setup_([\d.]+)\.exe`)
	match := re.FindAllStringSubmatch(string(body), -1)
	if len(match) == 0 || match[0][1] == "" {
		panic("Could not extract version info")
	}

	return match[0][1]
}

func downloadShadersAndPresets(shaders, presets string) {
	if _, err := os.Stat(shaders); os.IsNotExist(err) {
		must(os.MkdirAll(shaders, 0755))
	}

	if _, err := os.Stat(presets); os.IsNotExist(err) {
		must(os.MkdirAll(presets, 0755))
	}

	n := len(reshadeShaders)
	i := 0
	for name, f := range reshadeShaders {
		i++
		fmt.Printf("- Downloading %v (%v/%v)\n", name, i, n)

		f(shaders, presets)
	}
}
