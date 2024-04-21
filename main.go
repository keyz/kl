package main

import (
	"log"

	"golang.design/x/hotkey"
	"golang.design/x/hotkey/mainthread"
)

func main() {
	mainthread.Init(loop)
}

func loop() {
	hk := hotkey.New([]hotkey.Modifier{hotkey.ModCmd, hotkey.ModOption}, hotkey.KeyR)

	if err := hk.Register(); err != nil {
		log.Fatalf("[kl] failed to register hotkey: %v", err)
	}

	log.Printf("[kl] listening to hotkey %v...\n", hk)
	defer hk.Unregister()

	for {
		<-hk.Keyup()
		log.Printf("[kl] hotkey pressed: %v\n", hk)
	}
}

func Add(a, b int) int {
	return a + b
}
