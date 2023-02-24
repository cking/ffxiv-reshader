package main

import (
	"os"
)

func getGamePath(gamePath string) string {
	if gamePath != "" && isDir(gamePath) {
		return gamePath
	}

	env := os.Getenv("XIV_GAME_PATH")
	if env != "" && isDir(env) {
		return env
	}

	return getGamePathNative()
}

func isDir(path string) bool {
	return expect(os.Stat(path)).IsDir()
}
