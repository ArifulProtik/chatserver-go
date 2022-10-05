package main

import (
	"github.com/ArifulProtik/chatserver-go/config"
	"github.com/ArifulProtik/chatserver-go/internal/logger"
)

func main() {
	cfg, err := config.New("./", "config")
	if err != nil {
		panic(err)
	}

	logger := logger.New(&cfg.AppInfo)
	logger.Info(cfg.AppInfo.Name)
}
