package main

import (
	"github.com/victorzhou123/ai-agent/common/log"
	"github.com/victorzhou123/ai-agent/config"
)

const cfgPath = "./config.yml"

func main() {

	exitSig := make(chan struct{})
	defer func() {
		exitSig <- struct{}{}
	}()

	// config
	config.LoadConfig(cfgPath)

	// log init
	log.Init(&config.GetGlobalConfig().Log, exitSig)
}