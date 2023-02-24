package main

import (
	"os/exec"
	"path"
	"strings"
)

func run(name string, args ...string) string {
	cmd := exec.Command(name, args...)
	return string(expect(cmd.CombinedOutput()))
}

func getGamePathNative() string {
	xom := path.Join(strings.TrimSpace(run("defaults read dezent.XIV-on-Mac GamePath")), "/game/")
	if isDir(xom) {
		return xom
	}

	panic("Could not find game path. Please specify it as the first argument or set the environment variable XIV_GAME_PATH")
}
