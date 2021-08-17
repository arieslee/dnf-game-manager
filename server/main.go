package main

import (
	_ "dnf-game-manager/boot"
	"dnf-game-manager/router"
)

func main() {
	router.Build()
}
