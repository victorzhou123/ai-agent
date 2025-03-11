package main

import (
	"github.com/victorzhou123/ai-agent/common/log"
	"github.com/victorzhou123/ai-agent/config"
	"github.com/victorzhou123/ai-agent/server"
)

const cfgPath = "./config/config.yaml"

func main() {

	exitSig := make(chan struct{})
	defer func() {
		exitSig <- struct{}{}
	}()

	// config
	if err := config.LoadConfig(cfgPath); err != nil {
		return
	}

	// log init
	log.Init(&config.GetGlobalConfig().Log, exitSig)

	// web server
	if err := server.StartWebServer(config.GetGlobalConfig()); err != nil {
		log.Fatalf("start web server error: %s", err.Error())
	}
}
