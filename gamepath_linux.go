package main

import (
	"os"
	"path"
)

func getGamePathNative() string {
	home := os.Getenv("HOME")
	xl := path.Join(home, ".xlcore", "ffxiv")
	if isDir(xl) {
		return xl
	}

	panic("Could not find game path. Please specify it as the first argument or set the environment variable XIV_GAME_PATH")
}
