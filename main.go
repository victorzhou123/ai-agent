package main

import (
	"fmt"

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
		fmt.Printf("LoadConfig err: %s", err)
		return
	}

	lg := config.GetGlobalConfig()
	fmt.Printf("lg: %v\n", lg)

	// log init
	log.Init(&config.GetGlobalConfig().Log, exitSig)

	// web server
	server.StartWebServer(config.GetGlobalConfig())
}
