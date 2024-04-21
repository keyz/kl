package main

import (
	"os"

	"github.com/fatih/color"
	"golang.design/x/hotkey"
	"golang.design/x/hotkey/mainthread"
)

func main() {
	mainthread.Init(loop)
}

func loop() {
	hk := hotkey.New([]hotkey.Modifier{hotkey.ModCmd, hotkey.ModOption}, hotkey.KeyR)

	if err := hk.Register(); err != nil {
		color.Red("[kl] failed to register hotkey: %v", err)
		os.Exit(1)
	}

	color.Green("[kl] listening to hotkey %v...\n", hk)
	defer hk.Unregister()

	for {
		<-hk.Keyup()
		color.Green("[kl] hotkey pressed: %v\n", hk)
	}
}

func Add(a, b int) int {
	return a + b
}
